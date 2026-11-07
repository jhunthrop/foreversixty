# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 020000000000000000-00000000000000000-40500000000000000)

Set DPS (verified): 38.4. Weights run: 2.4s. Verify run: 4.0s. 186 eligible items had no known source.

Race sweep: human 38.4 ± 0.0, dwarf 38.4 ± 0.0, gnome 38.3 ± 0.0, night-elf 38.2 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.269 ± 0.022, crit=0.262 ± 0.007 per rating point (14 rating = 1%, 3.674 per %), hit=0.904 ± 0.026 per rating point (10 rating = 1%, 9.037 per %), spell_haste=not significant (0.507 ± 0.189), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003, spirit=not significant (-0.676 ± 0.310), mp5=-2.759 ± 0.354

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (0.76 DPS) | yes | Pristine Circlet (253949, -0.45 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.67 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.4 spell_power points (0.52 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.24 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Black Whelp Cloak (7283, -0.07 DPS) [crafted]; Caretaker's Cape (20428, -0.07 DPS) [rep]; Pearl-clasped Cloak (5542, -0.08 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 15.3 spell_power points (1.07 DPS) | yes | Filigreed Shadow Gown (253907, -0.19 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.62 DPS) [crafted]; Gray Woolen Robe (2585, -0.69 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.3 spell_power points (0.09 DPS) | yes | Windsong Bangles (263336, -0.02 DPS) [quest]; Repurposed Hair Band (281256, -0.06 DPS) [quest]; Bright Bracers (3647, -0.27 DPS, sim-verified) [world_drop] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 7.8 spell_power points (0.54 DPS) | yes | Serpent Gloves (5970, -0.06 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.13 DPS) [world]; Pristine Gloves (253913, -0.21 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 8.1 spell_power points (0.56 DPS) | yes | Novice Arcanist's Sash (253885, -0.23 DPS) [crafted]; Pristine Sash (253925, -0.27 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.37 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 12.6 spell_power points (0.88 DPS) | yes | Silk-threaded Trousers (1929, -0.39 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.45 DPS, sim-verified) [crafted]; Colorful Kilt (10048, -0.53 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.1 spell_power points (0.56 DPS) | yes | Shadow Boots (253895, -0.09 DPS) [crafted]; Red Woolen Boots (4313, -0.28 DPS) [crafted]; Pristine Boots (253889, -0.30 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.5 spell_power points (0.38 DPS) | yes | Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon]; Furlfeather Ring (286745, -0.32 DPS) [world]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.35 DPS) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Furlfeather Ring (286745, -0.28 DPS) [world]; Volcanic Rock Ring (12053, -0.29 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | 28.5 spell_power points (1.98 DPS) | yes | Witching Stave (1484, -0.01 DPS) [dungeon]; Staff of the Blessed Seer (2271, -0.04 DPS) [dungeon]; Emberstone Staff (5201, -0.32 DPS) [dungeon] |
| off_hand | Tear of Grief (5611) | For Love Eternal [quest] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mathystra Channeling Orb (281291, +0.00 DPS) [quest]; Spellbinder Orb (15926, -0.04 DPS) [dungeon]; Dwarven Tome (279898, -0.18 DPS, sim-verified) [quest] |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -1.02 DPS, sim-verified) [world_drop]; Deepblaze (279896, -2.98 DPS) [quest]; Sizzle Stick (8071, -3.54 DPS) [quest] |

**New at 20:** head: Shadow Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Tear of Grief; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes

### Band 30 (human, 020000000000000000-00000000000000000-50542001200000000)

Set DPS (verified): 62.8. Weights run: 2.4s. Verify run: 4.0s. 308 eligible items had no known source.

Race sweep: gnome 63.4 ± 0.1, human 62.8 ± 0.1, dwarf 62.6 ± 0.0, night-elf 62.5 ± 0.1 (dps); gnome beats the guide race human by 0.6 (0.95%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.342 ± 0.019, crit=0.192 ± 0.005 per rating point (14 rating = 1%, 2.687 per %), hit=0.586 ± 0.020 per rating point (10 rating = 1%, 5.863 per %), spell_haste=not significant (0.187 ± 0.136), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=-3.616 ± 0.218, mp5=-2.388 ± 0.273

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 14.0 spell_power points (1.26 DPS) | yes | Holy Shroud (2721, -0.27 DPS) [world_drop]; Shadow Circlet (253955, -0.27 DPS) [crafted]; Shadow Hood (4323, -0.29 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystal Starfire Medallion (5003, +0.00 DPS) [world_drop]; Seal of the Expedition (271920, +0.00 DPS) [rep]; Archmagister's Faceted Pendant (273049, +0.00 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (1.09 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Invoker's Mantle (215365, -0.30 DPS) [crafted]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, +0.00 DPS) [crafted]; Vine Pruner's Cloak (279835, +0.00 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Necrology Robes (2292, +0.00 DPS) [dungeon]; Green Silk Armor (7065, +0.00 DPS) [crafted]; Shadow Gown (253967, +0.00 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.81 DPS) | yes | Glowing Magical Bracelets (13106, -0.57 DPS) [world_drop]; Nightsky Wristbands (6407, -0.63 DPS) [world_drop]; Stonecloth Bindings (14416, -0.66 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 11.4 spell_power points (1.03 DPS) | yes | Hands of Darkness (7047, -0.22 DPS) [crafted]; Shadow Gloves (253919, -0.30 DPS) [crafted]; Serpent Gloves (5970, -0.40 DPS) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Arugal (6392, -0.18 DPS) [dungeon]; Invoker's Cord (215366, -0.30 DPS) [crafted]; Crimson Silk Belt (7055, -0.33 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 16.4 spell_power points (1.48 DPS) | yes | Necromancer Leggings (2277, -0.24 DPS) [world_drop]; Filigreed Shadow Leggings (253943, -0.30 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.40 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 15.4 spell_power points (1.39 DPS) | yes | Gilded Slippers (254001, -0.54 DPS) [crafted]; Spidersilk Boots (4320, -0.63 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Minor Channeling Ring (1449, -0.03 DPS) [quest]; Electrocutioner Lagnut (9447, -0.27 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Minor Channeling Ring (1449, -0.12 DPS) [quest]; Electrocutioner Lagnut (9447, -0.36 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, -0.20 DPS) [crafted]; Staff of the Shade (2549, -0.27 DPS) [dungeon]; Tendonscraper Dagger (273053, -0.36 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 9.1 spell_power points (0.82 DPS) | yes | Orb of Mistmantle (13031, -0.28 DPS) [world_drop]; Eye of Paleth (2943, -0.46 DPS) [quest]; Orb of Souls (249395, -0.46 DPS) [crafted] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, -0.92 DPS) [world]; Gyromatic Icemaker (9489, -1.50 DPS) [dungeon]; Gravestone Scepter (7001, -1.70 DPS) [quest] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Ring of Power Regulation; finger2: Lorekeeper's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 308, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 020000000000000000-00000000000000000-52542001201300040)

Set DPS (verified): 93.1. Weights run: 2.4s. Verify run: 2.7s. 402 eligible items had no known source.

Race sweep: gnome 93.9 ± 0.1, human 93.1 ± 0.1, dwarf 90.7 ± 0.1, night-elf 90.4 ± 0.1 (dps); gnome beats the guide race human by 0.8 (0.84%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.026 ± 0.027, crit=0.235 ± 0.006 per rating point (14 rating = 1%, 3.293 per %), hit=0.673 ± 0.030 per rating point (10 rating = 1%, 6.726 per %), spell_haste=not significant (0.237 ± 0.262), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.254 ± 0.223), mp5=not significant (1.127 ± 0.433)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 27.7 spell_power points (3.61 DPS) | yes | Corpseshroud (10574, -0.87 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.87 DPS) [crafted]; Augural Shroud (2620, -1.30 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.2 spell_power points (1.98 DPS) | yes | Triune Amulet (7722, -0.71 DPS) [dungeon]; Necklace of Calisea (1714, -0.81 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.89 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.3 spell_power points (2.65 DPS) | yes | Bloodmage Mantle (7684, -0.27 DPS) [dungeon]; Death Speaker Mantle (6685, -0.40 DPS) [dungeon]; Green Silken Shoulders (7057, -1.33 DPS, sim-verified) [crafted] |
| back | Expeditionary's Cape (271922) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.15 DPS) [crafted]; Guardian Cloak (5965, -0.15 DPS) [crafted]; Blackforge Cape (6424, -0.40 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 30.8 spell_power points (4.01 DPS) | yes | Robe of Power (7054, -0.32 DPS) [crafted]; Robe of the Magi (1716, -0.86 DPS, sim-verified) [world_drop]; Beguiler Robes (7728, -0.97 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.4 spell_power points (1.61 DPS) | yes | Arcane Runed Bracers (4744, -0.44 DPS) [quest]; Mindthrust Bracers (1974, -0.50 DPS) [dungeon]; Windchaser Cuffs (14429, -1.10 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.9 spell_power points (3.11 DPS) | yes | Shadoweave Gloves (10023, -0.23 DPS) [crafted]; Red Mageweave Gloves (10018, -0.34 DPS) [crafted]; Black Handwraps (254027, -0.35 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | 23.7 spell_power points (3.09 DPS) | yes | Deathmage Sash (10771, -0.55 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.73 DPS) [rep]; Gilded Cord (254037, -0.78 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.3 spell_power points (3.42 DPS) | yes | Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Necromancer Leggings (2277, -0.65 DPS) [world_drop]; Shadow Leggings (253993, -0.81 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.8 spell_power points (3.35 DPS) | yes | Black Slippers (254007, -1.24 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.38 DPS) [crafted]; Acidic Walkers (9454, -1.50 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.1 spell_power points (2.49 DPS) | yes | Lorekeeper's Ring (19524, -0.88 DPS) [rep]; Snake Hoop (6750, -1.04 DPS) [quest]; Ring of Power Regulation (273030, -1.27 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.97 DPS) | yes | Lorekeeper's Ring (19524, -0.36 DPS) [rep]; Snake Hoop (6750, -0.52 DPS) [quest]; Ring of Power Regulation (273030, -0.75 DPS) [dungeon] |
| trinket1 | Enriched Seal (271905) | Theramore Expeditionary Force [rep] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -5.23 DPS) [dungeon]; Illusionary Rod (7713, -5.65 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -6.23 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | Uldaman: Ancient Treasure [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Rod (9381, +0.00 DPS) [dungeon]; Twisted Nether Wand (249144, -0.02 DPS) [crafted]; Jaina's Firestarter (13064, -0.89 DPS, sim-verified) [world_drop] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Expeditionary's Cape; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Black Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Enriched Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Umbral Wand

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (human, 025003000000000000-00000000000000000-52542001201300051)

Set DPS (verified): 218.4. Weights run: 2.5s. Verify run: 4.8s. 517 eligible items had no known source.

Race sweep: gnome 226.8 ± 0.3, human 218.4 ± 0.3, night-elf 217.8 ± 0.3, dwarf 216.9 ± 0.3 (dps); gnome beats the guide race human by 8.3 (3.82%) ± 0.4, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.463 ± 0.029, crit=0.478 ± 0.012 per rating point (14 rating = 1%, 6.691 per %), hit=0.706 ± 0.034 per rating point (10 rating = 1%, 7.057 per %), spell_haste=not significant (0.558 ± 0.326), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.168 ± 0.226), mp5=not significant (0.517 ± 0.495)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 50.7 spell_power points (12.02 DPS) | yes | Red Mageweave Headband (10033, -0.58 DPS) [crafted]; Chief Architect's Monocle (11839, -2.66 DPS) [dungeon]; Shadoweave Mask (10025, -2.87 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.5 spell_power points (4.86 DPS) | yes | Mindburst Medallion (11196, -1.35 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -1.39 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.79 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.3 spell_power points (9.32 DPS) | yes | Kentic Amice (11624, -1.50 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.46 DPS) [crafted]; Netherpearl Shoulders (254055, -2.52 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Expeditionary's Cape (271921, -0.54 DPS) [rep]; Big Voodoo Cloak (8216, -1.09 DPS) [crafted]; Runecloth Cloak (13860, -1.82 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.3 spell_power points (11.44 DPS) | yes | Robes of Insight (940, -1.54 DPS, sim-verified) [world_drop]; Runecloth Robe (13858, -2.94 DPS) [crafted]; Knight's Dreadweave Vest (220886, -3.19 DPS) [vendor] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 23.2 spell_power points (5.51 DPS) | yes | Aristocratic Cuffs (12546, -0.31 DPS) [dungeon]; Bloodband Bracers (11469, -1.20 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.35 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 33.6 spell_power points (7.97 DPS) | yes | Red Mageweave Gloves (10018, -1.90 DPS) [crafted]; Runecloth Gloves (13863, -2.01 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.38 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 35.2 spell_power points (8.33 DPS) | yes | Black Waistcord (254087, -0.71 DPS) [crafted]; Deathmage Sash (10771, -1.47 DPS) [dungeon]; Dawnspire Cord (12466, -1.65 DPS, sim-verified) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 43.6 spell_power points (10.32 DPS) | yes | Spellshock Leggings (9484, -1.32 DPS, sim-verified) [dungeon]; Knight's Dreadweave Leggings (220888, -1.73 DPS) [vendor]; Red Mageweave Pants (10009, -2.84 DPS) [crafted] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 33.2 spell_power points (7.86 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.17 DPS) [vendor]; Gilded Sandals (254107, -2.13 DPS) [crafted]; Southsea Mojo Boots (20641, -2.15 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 21.9 spell_power points (5.20 DPS) | yes | Cyclopean Band (11824, -0.64 DPS) [dungeon]; Mindseye Circle (10634, -1.04 DPS) [dungeon]; Philanthropist's Ring (281635, -1.10 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 21.3 spell_power points (5.05 DPS) | yes | Cyclopean Band (11824, -0.49 DPS) [dungeon]; Mindseye Circle (10634, -0.89 DPS) [dungeon]; Philanthropist's Ring (281635, -0.95 DPS) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (+4.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Uther's Strength (11302, -0.08 DPS) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -2.89 DPS) [quest]; Glowing Brightwood Staff (812, -3.01 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -3.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Woestave (20082) | Blood of Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Pyric Caduceus (11748, -2.11 DPS) [dungeon]; Lesser Eternal Wand (249232, -3.03 DPS) [crafted]; Noxious Shooter (17745, -6.30 DPS, sim-verified) [dungeon] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Ban'thok Sash; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Underworld Band; trinket1: Enriched Seal; trinket2: Frozen Heart of the Mountain; ranged: Woestave

No-known-source sample (15 of 517, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (human, 025105031030000000-00000000000000000-52542001201300051)

Set DPS (verified): 409.0. Weights run: 2.7s. Verify run: 24.9s. 1359 eligible items had no known source.

Race sweep: gnome 410.1 ± 0.5, human 409.0 ± 0.5, night-elf 399.7 ± 0.3, dwarf 399.4 ± 0.3 (dps); gnome beats the guide race human by 1.1 (0.27%) ± 0.7, beyond error.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.364 ± 0.035, crit=0.714 ± 0.016 per rating point (14 rating = 1%, 9.991 per %), hit=1.256 ± 0.045 per rating point (10 rating = 1%, 12.563 per %), spell_haste=not significant (0.163 ± 0.227), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.049 ± 0.490), mp5=not significant (2.016 ± 0.533)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; The Postmaster's Band (13390, -4.65 DPS, sim-verified) [dungeon] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 32.5 spell_power points (6.11 DPS) | yes | Dark Advisor's Pendant (18691, -0.55 DPS) [dungeon]; Lady Maye's Pendant (14558, -1.23 DPS) [world_drop]; Archlight Talisman (15856, -1.29 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Theramore Shoulderpads (271953, +0.00 DPS) [rep] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.64 DPS) [world]; Fel Cape (279269, -2.73 DPS) [crafted]; Shroud of Arcane Mastery (22330, -4.19 DPS, sim-verified) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; The Postmaster's Tunic (13388, -8.20 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 32.9 spell_power points (6.20 DPS) | yes | Marshal's Satin Bracers (17606, -1.32 DPS) [pvp]; Runecloth Cuffs (254123, -1.56 DPS) [crafted]; Magiskull Cuffs (13107, -1.83 DPS) [world_drop] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Theramore Handwraps (271945, -3.22 DPS, sim-verified) [rep] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Clutch of Andros (13956, +0.00 DPS) [dungeon]; Dustfeather Sash (12589, -0.39 DPS) [dungeon]; Ban'thok Sash (11662, -5.10 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, +0.00 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Dragonrider Boots (18102, +0.00 DPS) [dungeon]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.77 DPS) [quest]; Deep Rooted Ring (19109, -1.13 DPS) [world]; Naglering (11669, -12.71 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Deep Rooted Ring (19109, -0.36 DPS) [world]; Seal of Rivendare (13345, -0.82 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.51 DPS) [vendor]; Shard of the Splithooves (10659, -1.86 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.01 DPS) [crafted] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.13 DPS) [crafted]; Enriched Seal (271903, -0.36 DPS) [rep] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Staff of Jordan (873, -0.30 DPS) [world_drop]; Ironbark Staff (20069, -5.63 DPS, sim-verified) [rep] |
| off_hand | Tome of Shadow Force (19309) | Stormpike Guard [rep] | 31.0 spell_power points (5.84 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.38 DPS) [rep] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.82 DPS) [dungeon]; Brilliant Wand (249385, -1.03 DPS) [crafted]; Oblivion's Touch (18761, -1.12 DPS) [dungeon] |

**New at 60:** head: Magister's Crown; neck: Diana's Pearl Necklace; shoulder: Magister's Mantle; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Burst of Knowledge; main_hand: Hand of Edward the Odd; off_hand: Tome of Shadow Force; ranged: Torch of Light

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (human, 005105031030000000-00000000000000000-50052201211312251)

Set DPS (verified): 783.4. Weights run: 3.0s. Verify run: 20.7s. 1359 eligible items had no known source.

Race sweep: human 783.4 ± 0.5, night-elf 782.8 ± 0.5, dwarf 782.1 ± 0.5, gnome 780.1 ± 0.5 (dps); the guide race human is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.270 ± 0.008, crit=0.948 ± 0.017 per rating point (14 rating = 1%, 13.274 per %), hit=1.189 ± 0.053 per rating point (10 rating = 1%, 11.885 per %), spell_haste=1.292 ± 0.255, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 32.7 spell_power points (15.57 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -0.26 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -0.18 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 33.8 spell_power points (16.08 DPS) | yes | Field Marshal's Satin Epaulets (231621, -1.61 DPS) [vendor]; Felcloth Shoulders (14112, -3.66 DPS, sim-verified) [crafted]; Argent Shoulders (19059, -4.18 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fel Cape (279269, -0.84 DPS) [crafted]; Crystalline Threaded Cape (20697, -4.27 DPS) [world]; Spritecaster Cape (11623, -6.87 DPS) [dungeon] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 43.2 spell_power points (20.59 DPS) | yes | Field Marshal's Satin Vestments (17605, -1.27 DPS) [vendor]; Field Marshal's Satin Robe (231618, -1.27 DPS) [vendor]; Felcloth Robe (14106, -3.63 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.2 spell_power points (11.50 DPS) | yes | Netherpearl Cuffs (254067, -4.41 DPS) [crafted]; Runecloth Cuffs (254123, -4.98 DPS) [crafted]; Arena Wristguards (18709, -5.18 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 38.3 spell_power points (18.21 DPS) | yes | Hands of Power (13253, -5.02 DPS, sim-verified) [dungeon]; Marshal's Satin Gloves (17608, -5.59 DPS) [vendor]; Marshal's Satin Grips (231617, -5.59 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 28.9 spell_power points (13.76 DPS) | yes | Ban'thok Sash (11662, -0.98 DPS, sim-verified) [dungeon]; Black Waistcord (254087, -3.55 DPS) [crafted]; Stormpike Cloth Girdle (19094, -3.90 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 61.3 spell_power points (29.16 DPS) | yes | Marshal's Satin Pants (17603, -11.22 DPS) [vendor]; Marshal's Satin Leggings (231619, -11.22 DPS) [vendor]; Skyshroud Leggings (13170, -12.24 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Maleki's Footwraps (18735, +0.00 DPS) [dungeon]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.42 DPS) [dungeon]; Maiden's Circle (13001, -2.51 DPS) [world_drop]; Naglering (11669, -13.43 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Songstone of Ironforge (12543, -0.95 DPS) [quest]; Maiden's Circle (13001, -0.95 DPS) [world_drop] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.91 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -4.76 DPS) [dungeon]; Uther's Strength (11302, -7.62 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Uther's Strength (11302, -2.24 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -4.19 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.30 DPS) [pvp]; Kindling Stave (11750, -16.21 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -16.84 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Woestave (20082, +0.00 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Earthen Silk Slippers; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

## Horde

### Band 20 (undead, 020000000000000000-00000000000000000-40500000000000000)

Set DPS (verified): 41.1. Weights run: 2.4s. Verify run: 3.6s. 179 eligible items had no known source.

Race sweep: undead 41.1 ± 0.0, troll 37.3 ± 0.0 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.669 ± 0.018, crit=0.148 ± 0.003 per rating point (14 rating = 1%, 2.073 per %), hit=0.451 ± 0.016 per rating point (10 rating = 1%, 4.512 per %), spell_haste=not significant (0.437 ± 0.204), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.827 ± 0.246), mp5=not significant (1.176 ± 0.295)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (1.00 DPS) | yes | Pristine Circlet (253949, -0.45 DPS) [crafted]; Shadow Goggles (4373, -0.66 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.70 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Roadwatcher's Confidence (281265, -0.14 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.0 spell_power points (1.00 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.30 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.64 DPS) [crafted]; Slime-encrusted Pads (6461, -0.68 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (20427, -0.24 DPS) [rep]; Pearl-clasped Cloak (5542, -0.30 DPS) [crafted]; Spirit Cloak (4792, -0.36 DPS) [vendor] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 17.3 spell_power points (1.57 DPS) | yes | Filigreed Shadow Gown (253907, -0.26 DPS, sim-verified) [crafted]; Scholarly Robes (2034, -0.40 DPS) [dungeon]; Robe of the Moccasin (6465, -0.55 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 6.9 spell_power points (0.62 DPS) | yes | Tabitha's Cuffs (251486, -0.35 DPS, sim-verified) [quest]; Bright Bracers (3647, -0.38 DPS) [world_drop]; Crystalline Cuffs (14148, -0.40 DPS) [dungeon] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 9.0 spell_power points (0.82 DPS) | yes | Serpent Gloves (5970, -0.18 DPS) [dungeon]; Magefist Gloves (12977, -0.21 DPS) [world_drop]; Pristine Gloves (253913, -0.27 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 9.7 spell_power points (0.88 DPS) | yes | Novice Ardent's Sash (253887, -0.29 DPS) [crafted]; Pristine Sash (253925, -0.30 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.33 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 18.3 spell_power points (1.66 DPS) | yes | Filigreed Pristine Leggings (253937, -0.50 DPS, sim-verified) [crafted]; Darkweave Breeches (12987, -0.79 DPS) [world_drop]; Colorful Kilt (10048, -0.83 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.7 spell_power points (0.88 DPS) | yes | Shadow Boots (253895, -0.15 DPS) [crafted]; Red Woolen Boots (4313, -0.31 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.42 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 7.4 spell_power points (0.67 DPS) | yes | Loop of Sacrifice (281673, -0.36 DPS) [quest]; Volcanic Rock Ring (12053, -0.48 DPS) [world_drop]; Band of Purification (12996, -0.65 DPS, sim-verified) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | sim-verified (41.1 DPS) | yes | Loop of Sacrifice (281673, -0.06 DPS) [quest]; Volcanic Rock Ring (12053, -0.18 DPS) [world_drop]; Band of Purification (12996, -0.25 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 36.3 spell_power points (3.29 DPS) | yes | Evocator's Blade (2567, -0.25 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.39 DPS, sim-verified) [quest]; Crescent Staff (6505, -0.51 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skycaller (12984, -0.92 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -3.31 DPS) [quest]; Greater Magic Wand (11288, -3.42 DPS) [crafted] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of the Blessed Seer; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (undead, 020000000000000000-00000000000000000-50542001200000000)

Set DPS (verified): 68.7. Weights run: 2.4s. Verify run: 3.8s. 301 eligible items had no known source.

Race sweep: undead 68.7 ± 0.1, troll 61.9 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.706 ± 0.017, crit=0.164 ± 0.004 per rating point (14 rating = 1%, 2.291 per %), hit=0.428 ± 0.018 per rating point (10 rating = 1%, 4.280 per %), spell_haste=not significant (0.029 ± 0.202), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.686 ± 0.212), mp5=3.762 ± 0.300

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | sim-verified (68.7 DPS) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Shadow Hood (4323, +0.00 DPS) [crafted]; Enchanter's Cowl (4322, -0.12 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 5.6 spell_power points (0.73 DPS) | yes | Crystal Starfire Medallion (5003, -0.09 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.17 DPS) [world]; Pendant of Myzrael (4614, -0.19 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.4 spell_power points (2.02 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Magician's Mantle (12998, -0.53 DPS) [world_drop]; Slime-encrusted Pads (6461, -0.54 DPS) [dungeon] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, +0.00 DPS) [dungeon]; Darkspear Raider's Cloak (272078, -0.00 DPS) [vendor]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted] |
| chest | Shadow Gown (253967) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Beguiler Robes (7728, +0.00 DPS) [dungeon]; Green Silk Armor (7065, -0.33 DPS) [crafted]; Necrology Robes (2292, -0.42 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 20.3 spell_power points (2.67 DPS) | yes | Mindthrust Bracers (1974, -0.72 DPS) [dungeon]; Nightsky Wristbands (6407, -1.84 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.93 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 13.9 spell_power points (1.83 DPS) | yes | Hands of Darkness (7047, -0.20 DPS) [crafted]; Gloves of Old (9395, -0.29 DPS) [world_drop]; Gloves of Meditation (4318, -0.55 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Belt of Arugal (6392, +0.00 DPS) [dungeon]; Warsong Sash (16975, -0.28 DPS) [quest]; Crimson Silk Belt (7055, -0.29 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Leech Pants (6910, +0.00 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.55 DPS) [dungeon]; Filigreed Shadow Leggings (253943, -0.58 DPS) [crafted] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 20.7 spell_power points (2.72 DPS) | yes | Gilded Slippers (254001, -0.79 DPS) [crafted]; Acidic Walkers (9454, -0.96 DPS) [dungeon]; Glinteye Slippers (273024, -1.26 DPS) [dungeon] |
| finger1 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 17.3 spell_power points (2.27 DPS) | yes | Darkspear Signet (272071, -0.29 DPS) [vendor]; Snake Hoop (6750, -1.02 DPS) [quest]; Electrocutioner Lagnut (9447, -1.25 DPS) [dungeon] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Darkspear Signet (272071, +0.00 DPS) [vendor]; Snake Hoop (6750, -0.66 DPS) [quest]; Electrocutioner Lagnut (9447, -0.88 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, -0.13 DPS) [crafted]; Soulstaff (249393, -0.16 DPS) [crafted]; Rod of the Sleepwalker (1155, -0.42 DPS) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.2 spell_power points (1.48 DPS) | yes | Orb of Mistmantle (13031, -0.33 DPS) [world_drop]; Orb of Souls (249395, -0.41 DPS) [crafted]; Witch's Finger (16887, -0.73 DPS) [quest] |
| ranged | Starfaller (13063) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Mana Concentration (273043, +0.00 DPS) [world]; Gravestone Scepter (7001, -0.99 DPS) [quest]; Greater Mystic Wand (217287, -1.39 DPS) [crafted] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Shadow Gown; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Ring of Power Regulation; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Starfaller

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (undead, 020000000000000000-00000000000000000-52542001201300040)

Set DPS (verified): 104.4. Weights run: 2.4s. Verify run: 2.6s. 395 eligible items had no known source.

Race sweep: undead 104.4 ± 0.1, troll 89.2 ± 0.1 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.026 ± 0.027, crit=0.235 ± 0.006 per rating point (14 rating = 1%, 3.293 per %), hit=0.673 ± 0.030 per rating point (10 rating = 1%, 6.726 per %), spell_haste=not significant (0.237 ± 0.262), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.254 ± 0.223), mp5=not significant (1.127 ± 0.433)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 27.7 spell_power points (3.61 DPS) | yes | Augural Shroud (2620, -0.75 DPS, sim-verified) [world]; Corpseshroud (10574, -0.87 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.87 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.2 spell_power points (1.98 DPS) | yes | Triune Amulet (7722, -0.71 DPS) [dungeon]; Necklace of Calisea (1714, -0.81 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.94 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.3 spell_power points (2.65 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Bloodmage Mantle (7684, -0.27 DPS) [dungeon]; Death Speaker Mantle (6685, -0.40 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.15 DPS) [crafted]; Guardian Cloak (5965, -0.15 DPS) [crafted]; Blackforge Cape (6424, -0.40 DPS) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 30.8 spell_power points (4.01 DPS) | yes | Robe of the Magi (1716, -0.18 DPS) [world_drop]; Robe of Power (7054, -0.32 DPS) [crafted]; Beguiler Robes (7728, -0.97 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.4 spell_power points (1.61 DPS) | yes | Radiant Silver Bracers (4545, -0.02 DPS) [quest]; Windchaser Cuffs (14429, -0.41 DPS) [world_drop]; Mindthrust Bracers (1974, -0.50 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.9 spell_power points (3.11 DPS) | yes | Shadoweave Gloves (10023, -0.23 DPS) [crafted]; Red Mageweave Gloves (10018, -0.34 DPS) [crafted]; Black Handwraps (254027, -0.35 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | 23.7 spell_power points (3.09 DPS) | yes | Deathmage Sash (10771, -0.18 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.73 DPS) [rep]; Gilded Cord (254037, -0.78 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.3 spell_power points (3.42 DPS) | yes | Shadow Leggings (253993, -0.50 DPS, sim-verified) [crafted]; Crimson Silk Pantaloons (7062, -0.65 DPS) [crafted]; Necromancer Leggings (2277, -0.65 DPS) [world_drop] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.8 spell_power points (3.35 DPS) | yes | Black Slippers (254007, -0.51 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.38 DPS) [crafted]; Acidic Walkers (9454, -1.50 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.1 spell_power points (2.49 DPS) | yes | Advisor's Ring (19520, -0.88 DPS) [rep]; Snake Hoop (6750, -1.04 DPS) [quest]; Ring of Power Regulation (273030, -1.27 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.97 DPS) | yes | Advisor's Ring (19520, -0.36 DPS) [rep]; Snake Hoop (6750, -0.52 DPS) [quest]; Ring of Power Regulation (273030, -0.75 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -5.23 DPS) [dungeon]; Illusionary Rod (7713, -5.34 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -6.23 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Earthen Rod (9381, -2.16 DPS) [dungeon]; Twisted Nether Wand (249144, -2.46 DPS) [crafted]; Umbral Wand (5216, -2.52 DPS) [dungeon] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Black Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Darkspear Voodoo Seal; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger; 5743 Prismstone Ring

### Band 50 (undead, 025003000000000000-00000000000000000-52542001201300051)

Set DPS (verified): 250.7. Weights run: 2.5s. Verify run: 5.0s. 510 eligible items had no known source.

Race sweep: undead 250.7 ± 0.3, troll 214.8 ± 0.3 (dps); the guide race undead is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.463 ± 0.029, crit=0.478 ± 0.012 per rating point (14 rating = 1%, 6.691 per %), hit=0.706 ± 0.034 per rating point (10 rating = 1%, 7.057 per %), spell_haste=not significant (0.558 ± 0.326), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.168 ± 0.226), mp5=not significant (0.517 ± 0.495)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 50.7 spell_power points (12.02 DPS) | yes | Red Mageweave Headband (10033, -0.58 DPS) [crafted]; Chief Architect's Monocle (11839, -2.66 DPS) [dungeon]; Shadoweave Mask (10025, -2.87 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.5 spell_power points (4.86 DPS) | yes | Glowing Eye of Mordresh (10769, -0.42 DPS) [dungeon]; Mindburst Medallion (11196, -1.35 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -1.39 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.3 spell_power points (9.32 DPS) | yes | Kentic Amice (11624, -1.50 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.46 DPS) [crafted]; Netherpearl Shoulders (254055, -2.52 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Runecloth Cloak (13860, -0.49 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.54 DPS) [vendor]; Deep Woodlands Cloak (19121, -0.62 DPS, sim-verified) [quest] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.3 spell_power points (11.44 DPS) | yes | Runecloth Robe (13858, -2.94 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -3.19 DPS) [vendor]; Robes of Insight (940, -3.31 DPS, sim-verified) [world_drop] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 23.2 spell_power points (5.51 DPS) | yes | Aristocratic Cuffs (12546, -0.31 DPS) [dungeon]; Bloodband Bracers (11469, -1.20 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.35 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 33.6 spell_power points (7.97 DPS) | yes | Red Mageweave Gloves (10018, -1.90 DPS) [crafted]; Runecloth Gloves (13863, -2.01 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -3.27 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 35.2 spell_power points (8.33 DPS) | yes | Dawnspire Cord (12466, -0.32 DPS) [dungeon]; Black Waistcord (254087, -0.71 DPS) [crafted]; Deathmage Sash (10771, -1.47 DPS) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 43.6 spell_power points (10.32 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -1.73 DPS) [vendor]; Spellshock Leggings (9484, -1.91 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -2.84 DPS) [crafted] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 33.2 spell_power points (7.86 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -2.01 DPS, sim-verified) [vendor]; Gilded Sandals (254107, -2.13 DPS) [crafted]; Southsea Mojo Boots (20641, -2.15 DPS) [quest] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 21.9 spell_power points (5.20 DPS) | yes | Cyclopean Band (11824, -0.64 DPS) [dungeon]; Mindseye Circle (10634, -1.04 DPS) [dungeon]; Philanthropist's Ring (281635, -1.10 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 21.3 spell_power points (5.05 DPS) | yes | Cyclopean Band (11824, -0.49 DPS) [dungeon]; Mindseye Circle (10634, -0.89 DPS) [dungeon]; Philanthropist's Ring (281635, -0.95 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.08 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.33 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -2.89 DPS) [quest]; Glowing Brightwood Staff (812, -3.37 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -3.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Woestave (20082) | Blood of Morphaz [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -2.20 DPS) [crafted]; Pyric Caduceus (11748, -2.29 DPS) [dungeon]; Noxious Shooter (17745, -4.65 DPS, sim-verified) [dungeon] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Ban'thok Sash; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Underworld Band; trinket1: Darkspear Voodoo Seal; trinket2: Frozen Heart of the Mountain; ranged: Woestave

No-known-source sample (15 of 510, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60 (undead, 025105031030000000-00000000000000000-52542001201300051)

Set DPS (verified): 434.1. Weights run: 2.7s. Verify run: 24.2s. 1354 eligible items had no known source.

Race sweep: undead 434.1 ± 0.5, troll 383.7 ± 0.5 (dps); the guide race undead is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.364 ± 0.035, crit=0.714 ± 0.016 per rating point (14 rating = 1%, 9.991 per %), hit=1.256 ± 0.045 per rating point (10 rating = 1%, 12.563 per %), spell_haste=not significant (0.163 ± 0.227), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.049 ± 0.490), mp5=not significant (2.016 ± 0.533)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | The Postmaster's Band (13390) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -10.26 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lady Maye's Pendant (14558, +0.00 DPS) [world_drop]; Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Champion's Satin Epaulets (227130, -0.05 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.47 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.64 DPS) [world]; Fel Cape (279269, -2.73 DPS) [crafted]; Shroud of Arcane Mastery (22330, -4.13 DPS, sim-verified) [dungeon] |
| chest | The Postmaster's Tunic (13388) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robe of Winter Night (14136, -6.13 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 32.9 spell_power points (6.20 DPS) | yes | General's Satin Bracers (17619, -1.32 DPS) [pvp]; Runecloth Cuffs (254123, -1.56 DPS) [crafted]; Magiskull Cuffs (13107, -1.83 DPS) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Gloves (17620, -0.36 DPS) [vendor]; General's Satin Grips (231613, -0.36 DPS) [vendor]; Raider Handwraps (272097, -3.70 DPS, sim-verified) [vendor] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Clutch of Andros (13956, +0.00 DPS) [dungeon]; Dustfeather Sash (12589, -0.39 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 53.9 spell_power points (10.15 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Legionnaire's Satin Leggings (227131, -0.07 DPS) [vendor]; Skyshroud Leggings (13170, -6.66 DPS, sim-verified) [dungeon] |
| feet | The Postmaster's Treads (13391) | Stratholme: Postmaster Malown [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Maleki's Footwraps (18735, -5.74 DPS, sim-verified) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.77 DPS) [quest]; Deep Rooted Ring (19109, -1.13 DPS) [world]; Naglering (11669, -13.99 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Deep Rooted Ring (19109, -0.36 DPS) [world]; Seal of Rivendare (13345, -0.82 DPS) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Burst of Knowledge (11832, -16.51 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -1.51 DPS) [vendor]; Shard of the Splithooves (10659, -1.86 DPS) [quest]; Burst of Knowledge (11832, -1.88 DPS) [dungeon] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Staff of Jordan (873, -0.30 DPS) [world_drop]; Ironbark Staff (20220, -6.11 DPS, sim-verified) [rep] |
| off_hand | Tome of Shadow Force (19309) | Frostwolf Clan [rep] | 31.0 spell_power points (5.84 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Therazane's Touch (19315, -0.38 DPS) [rep] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.13 DPS) [dungeon]; Oblivion's Touch (18761, -0.41 DPS) [dungeon]; Torch of Light (279246, -0.85 DPS) [crafted] |

**New at 60:** head: The Postmaster's Band; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: The Postmaster's Tunic; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Outrider's Silk Leggings; feet: The Postmaster's Treads; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Second Wind; trinket2: Weakness Analyzer; main_hand: Hand of Edward the Odd; off_hand: Tome of Shadow Force; ranged: Brilliant Wand

No-known-source sample (15 of 1354, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 60, raid preset (undead, 005105031030000000-00000000000000000-50052201211312251)

Set DPS (verified): 769.6. Weights run: 3.0s. Verify run: 19.5s. 1354 eligible items had no known source.

Race sweep: undead 769.6 ± 0.5, troll 761.8 ± 0.5 (dps); the guide race undead is best.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.270 ± 0.008, crit=0.948 ± 0.017 per rating point (14 rating = 1%, 13.274 per %), hit=1.189 ± 0.053 per rating point (10 rating = 1%, 11.885 per %), spell_haste=1.292 ± 0.255, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 32.7 spell_power points (15.57 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -0.26 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -0.18 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 33.8 spell_power points (16.08 DPS) | yes | Warlord's Satin Epaulets (231611, -1.61 DPS) [vendor]; Felcloth Shoulders (14112, -3.29 DPS, sim-verified) [crafted]; Argent Shoulders (19059, -4.18 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fel Cape (279269, -0.84 DPS) [crafted]; Crystalline Threaded Cape (20697, -4.27 DPS) [world]; Spritecaster Cape (11623, -6.87 DPS) [dungeon] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 43.2 spell_power points (20.59 DPS) | yes | Warlord's Satin Robes (231612, -1.27 DPS) [pvp]; Felcloth Robe (14106, -3.25 DPS, sim-verified) [crafted]; Warlord's Satin Tunic (231632, -6.99 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.2 spell_power points (11.50 DPS) | yes | Netherpearl Cuffs (254067, -4.41 DPS) [crafted]; Runecloth Cuffs (254123, -4.98 DPS) [crafted]; Arena Wristguards (18709, -5.18 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 38.3 spell_power points (18.21 DPS) | yes | Hands of Power (13253, -5.25 DPS, sim-verified) [dungeon]; General's Satin Gloves (17620, -5.59 DPS) [vendor]; General's Satin Grips (231613, -5.59 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 28.9 spell_power points (13.76 DPS) | yes | Ban'thok Sash (11662, -0.97 DPS) [dungeon]; Black Waistcord (254087, -3.55 DPS) [crafted]; Frostwolf Cloth Belt (19090, -3.90 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Satin Leggings (231614, +0.00 DPS) [pvp]; Felcloth Pants (14107, -1.85 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; Maleki's Footwraps (18735, +0.00 DPS) [dungeon]; General's Satin Treads (231610, +0.00 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.42 DPS) [dungeon]; Maiden's Circle (13001, -2.51 DPS) [world_drop]; Naglering (11669, -13.21 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Eye of Orgrimmar (12545, -0.95 DPS) [quest]; Maiden's Circle (13001, -0.95 DPS) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.62 DPS) [crafted]; Rune of the Guard Captain (19120, -1.75 DPS) [quest]; Uther's Strength (11302, -2.86 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.80 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -5.38 DPS) [crafted]; Rune of the Guard Captain (19120, -6.51 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.30 DPS) [pvp]; Staff of Jordan (873, -16.84 DPS) [world_drop]; Kindling Stave (11750, -20.78 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Woestave (20082, +0.00 DPS) [quest]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Winter Night; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Earthen Silk Slippers; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1354, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

