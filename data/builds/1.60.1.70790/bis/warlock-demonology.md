# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-2035100000000000000-0000000000000000)

Set DPS (verified): 56.7. Weights run: 3.1s. Verify run: 3.6s. 182 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.127 ± 0.005, crit=0.059 ± 0.001 per rating point (14 rating = 1%, 0.821 per %), hit=0.157 ± 0.007 per rating point (10 rating = 1%, 1.572 per %), spell_haste=0.516 ± 0.078, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.831 ± 0.001, fire_power=0.169 ± 0.000, spirit=not significant (0.169 ± 0.051)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.1 spell_power points (2.54 DPS) | yes | Pristine Circlet (253949, -0.87 DPS) [crafted]; Flame Circlet (253953, -2.02 DPS) [crafted]; Shadow Goggles (4373, -2.08 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.09 DPS) | yes | - |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.1 spell_power points (1.71 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.18 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.60 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.37 DPS) [rep]; Black Whelp Cloak (7283, -0.47 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.64 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.3 spell_power points (3.41 DPS) | yes | Filigreed Shadow Gown (253907, -1.01 DPS) [crafted]; Filigreed Pristine Gown (253901, -1.70 DPS) [crafted]; Scholarly Robes (2034, -2.01 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (56.7 DPS) | yes | Windsong Bangles (263336, +0.00 DPS) [quest]; Repurposed Hair Band (281256, -0.01 DPS) [quest]; Bright Bracers (3647, -0.04 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.94 DPS) | yes | Shadow Gloves (253919, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.28 DPS) [world]; Pristine Gloves (253913, -0.73 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.3 spell_power points (1.76 DPS) | yes | Pristine Sash (253925, -0.51 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.54 DPS) [crafted]; Novice Ardent's Sash (253887, -0.95 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.6 spell_power points (2.94 DPS) | yes | Filigreed Pristine Leggings (253937, -0.87 DPS) [crafted]; Silk-threaded Trousers (1929, -1.00 DPS) [dungeon]; Colorful Kilt (10048, -1.32 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.5 spell_power points (2.08 DPS) | yes | Shadow Boots (253895, -0.59 DPS) [crafted]; Red Woolen Boots (4313, -0.74 DPS) [crafted]; Pristine Boots (253889, -1.15 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.3 spell_power points (1.46 DPS) | yes | Band of Purification (12996, -1.18 DPS) [world_drop]; Furlfeather Ring (286745, -1.18 DPS) [world]; Lavishly Jeweled Ring (1156, -1.25 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.39 DPS) | yes | Band of Purification (12996, -1.11 DPS) [world_drop]; Furlfeather Ring (286745, -1.11 DPS) [world]; Lavishly Jeweled Ring (1156, -1.18 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Blessed Seer (2271, +0.00 DPS) [dungeon]; Emberstone Staff (5201, -1.42 DPS) [dungeon]; Witching Stave (1484, -1.67 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.2 spell_power points (0.88 DPS) | yes | Antipodean Rod (2879, -0.60 DPS) [world_drop]; Onu's Patience (281292, -0.65 DPS) [quest]; Grayson's Torch (1172, -0.69 DPS) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.28 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 182, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 00000000000000000-2035113101210001000-0000000000000000)

Set DPS (verified): 91.9. Weights run: 3.1s. Verify run: 2.8s. 309 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.091 ± 0.004, crit=0.040 ± 0.001 per rating point (14 rating = 1%, 0.562 per %), hit=0.092 ± 0.004 per rating point (10 rating = 1%, 0.924 per %), spell_haste=0.281 ± 0.051, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.810 ± 0.001, fire_power=0.190 ± 0.001, spirit=not significant (0.069 ± 0.041)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.4 spell_power points (3.24 DPS) | yes | Filigreed Shadow Circlet (253981, -0.02 DPS) [crafted]; Silk Headband (7050, -0.69 DPS) [crafted]; Shadow Circlet (253955, -0.71 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 0.6 spell_power points (0.18 DPS) | yes | Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Seal of the Expedition (271920, -0.05 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (2.79 DPS) | yes | Invoker's Mantle (215365, -0.49 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.80 DPS) [dungeon]; Moonlit Amice (11884, -0.80 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted]; Caretaker's Cape (19533, -0.21 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.69 DPS) | yes | Necrology Robes (2292, -0.34 DPS) [dungeon]; Shadow Gown (253967, -0.44 DPS) [crafted]; Beguiler Robes (7728, -0.67 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.56 DPS) | yes | Phoenix Bindings (210781, -2.24 DPS, sim-verified) [crafted]; Windsong Bangles (263336, -2.27 DPS) [quest]; Nightsky Wristbands (6407, -2.34 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.9 spell_power points (2.25 DPS) | yes | Hands of Darkness (7047, -0.08 DPS) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.55 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.3 spell_power points (3.20 DPS) | yes | Belt of Arugal (6392, -0.47 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.93 DPS) [dungeon]; Invoker's Cord (215366, -1.08 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.6 spell_power points (3.58 DPS) | yes | Shadow Leggings (253993, -0.09 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.82 DPS) [crafted]; Necromancer Leggings (2277, -1.00 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.4 spell_power points (3.25 DPS) | yes | Gilded Slippers (254001, -1.12 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.14 DPS) [dungeon]; Spidersilk Boots (4320, -1.16 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.99 DPS) | yes | Ring of Power Regulation (273030, -0.28 DPS) [dungeon]; Minor Channeling Ring (1449, -0.52 DPS) [quest]; Electrocutioner Lagnut (9447, -1.00 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.5 spell_power points (1.83 DPS) | yes | Ring of Power Regulation (273030, -0.13 DPS) [dungeon]; Minor Channeling Ring (1449, -0.36 DPS) [quest]; Electrocutioner Lagnut (9447, -0.84 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.21 DPS) [dungeon]; Spine of the Devourer (273040, -1.21 DPS) [world]; Tendonscraper Dagger (273053, -1.46 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.56 DPS) | yes | Orb of Mystic Insight (249394, -0.41 DPS) [crafted]; Orb of Mistmantle (13031, -1.10 DPS) [world_drop]; Orb of Souls (249395, -1.30 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, -0.26 DPS) [quest]; Wand of Eventide (5214, -0.28 DPS) [world_drop]; Greater Mystic Wand (217287, -0.28 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 10000000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 163.3. Weights run: 3.2s. Verify run: 2.5s. 399 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.162 ± 0.011, crit=0.146 ± 0.003 per rating point (14 rating = 1%, 2.051 per %), hit=0.383 ± 0.016 per rating point (10 rating = 1%, 3.834 per %), spell_haste=2.661 ± 0.196, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.839 ± 0.001, fire_power=0.161 ± 0.000, spirit=not significant (0.071 ± 0.102)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.98 DPS) | yes | Living Cowl (5608, -2.65 DPS, sim-verified) [world]; Augural Shroud (2620, -2.79 DPS) [world]; Filigreed Shadow Circlet (253981, -3.07 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.5 spell_power points (2.82 DPS) | yes | Glowing Eye of Mordresh (10769, -0.80 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.28 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.29 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (3.47 DPS) | yes | Green Silken Shoulders (7057, -0.22 DPS) [crafted]; Inquisitor's Shawl (19507, -0.45 DPS) [dungeon]; Berylline Pads (4197, -0.61 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.21 DPS) [crafted]; Caretaker's Cape (19532, -0.36 DPS) [rep]; Long Silken Cloak (4326, -0.60 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (7.75 DPS) | yes | Elemental Raiment (9434, -0.95 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.95 DPS) [crafted]; Shadoweave Robe (10004, -1.73 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Arcane Runed Bracers (4744, -0.39 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.66 DPS) [quest]; Earthen Silk Cuffs (254019, -1.54 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (6.36 DPS) | yes | Black Mageweave Gloves (10003, -1.38 DPS) [crafted]; Shadoweave Gloves (10023, -1.78 DPS, sim-verified) [crafted]; Black Handwraps (254027, -2.08 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.6 spell_power points (4.87 DPS) | yes | Black Cord (254043, -0.39 DPS) [crafted]; Star Belt (4329, -0.45 DPS) [crafted]; Belt of Arugal (6392, -1.60 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.6 spell_power points (5.86 DPS) | yes | Red Mageweave Pants (10009, -0.56 DPS) [crafted]; Shadow Leggings (253993, -1.46 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.65 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.5 spell_power points (8.14 DPS) | yes | Black Slippers (254007, -4.25 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -5.34 DPS) [crafted]; Spidersilk Boots (4320, -5.60 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.6 spell_power points (4.17 DPS) | yes | Lorekeeper's Ring (19524, -1.18 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.51 DPS) [quest]; Reedknot Ring (9622, -1.85 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (3.59 DPS) | yes | Lorekeeper's Ring (19524, -0.59 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.93 DPS) [quest]; Reedknot Ring (9622, -1.27 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.82 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -8.04 DPS) [dungeon]; Swinetusk Shank (6691, -9.78 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.25 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 399, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 25400000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 238.4. Weights run: 3.2s. Verify run: 4.9s. 506 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.114 ± 0.015, crit=0.189 ± 0.004 per rating point (14 rating = 1%, 2.649 per %), hit=0.489 ± 0.021 per rating point (10 rating = 1%, 4.886 per %), spell_haste=2.201 ± 0.264, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.772 ± 0.001, fire_power=0.228 ± 0.000, spirit=not significant (0.207 ± 0.149)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.04 DPS) | yes | Dreamweave Circlet (10041, -1.33 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.13 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.23 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.4 spell_power points (3.51 DPS) | yes | Mindburst Medallion (11196, -0.64 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.26 DPS, sim-verified) [dungeon]; Lei of Lilies (1315, -2.36 DPS) [world_drop] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Rotgrip Mantle (17732, -0.08 DPS) [dungeon]; Shadoweave Shoulders (10028, -0.80 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -1.53 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.58 DPS) [dungeon]; Caretaker's Cape (19531, -2.86 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.7 spell_power points (8.82 DPS) | yes | Acumen Robes (17775, -0.91 DPS) [quest]; Elemental Raiment (9434, -1.01 DPS) [world_drop]; Dreamweave Vest (10021, -1.29 DPS, sim-verified) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 11.9 spell_power points (4.41 DPS) | yes | Arcane Runed Bracers (4744, -1.06 DPS) [quest]; Spidertank Oilrag (9448, -1.06 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.13 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.9 spell_power points (7.40 DPS) | yes | Black Gloves (254101, -1.65 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.82 DPS) [crafted]; Runecloth Gloves (13863, -1.87 DPS) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Black Waistcord (254087, -0.22 DPS) [crafted]; Satyrmane Sash (17755, -0.42 DPS) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Wizardweave Leggings (14132, -0.90 DPS) [crafted]; Senior Designer's Pantaloons (11841, -1.83 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.4 spell_power points (9.47 DPS) | yes | Black Sandals (254113, -2.86 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -4.29 DPS) [vendor]; Shadoweave Boots (10031, -4.68 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.84 DPS) | yes | Underworld Band (1980, -0.61 DPS) [world_drop]; Cyclopean Band (11824, -0.88 DPS) [dungeon]; Philanthropist's Ring (281635, -0.90 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.46 DPS) | yes | Underworld Band (1980, -0.23 DPS) [world_drop]; Cyclopean Band (11824, -0.51 DPS) [dungeon]; Philanthropist's Ring (281635, -0.53 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.71 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.87 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -1.44 DPS, sim-verified) [world]; Kindling Stave (11750, -2.85 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.36 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.37 DPS) [crafted]; Burning Sliver (5249, -0.74 DPS) [quest]; Twisted Nether Wand (249144, -0.74 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Ghostweave Cord; legs: Felcloth Pants; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; ranged: Lesser Eternal Wand

No-known-source sample (15 of 506, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522000000000000-2035113101220001350-0050000000000000)

Set DPS (verified): 408.8. Weights run: 3.4s. Verify run: 19.1s. 1299 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.449 ± 0.026, crit=0.207 ± 0.007 per rating point (14 rating = 1%, 2.901 per %), hit=0.616 ± 0.026 per rating point (10 rating = 1%, 6.159 per %), spell_haste=3.025 ± 0.347, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.001, fire_power=0.235 ± 0.000, spirit=not significant (0.195 ± 0.156)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 35.2 spell_power points (15.17 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.37 DPS) [pvp]; Red Mageweave Headband (10033, -7.48 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -2.04 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burial Shawl (18681, +0.00 DPS) [dungeon]; Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.27 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -3.40 DPS) [dungeon]; Crystalline Threaded Cape (20697, -4.54 DPS, sim-verified) [world]; Fel Cape (279269, -4.91 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.0 spell_power points (21.60 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.13 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -6.93 DPS) [pvp]; Robe of Winter Night (14136, -8.20 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 27.0 spell_power points (11.64 DPS) | yes | Runecloth Cuffs (254123, -4.95 DPS) [crafted]; Netherpearl Cuffs (254067, -5.56 DPS) [crafted]; Nethergeld Cuffs (254061, -6.84 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Handwraps (227100, +0.00 DPS) [pvp]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp] |
| waist | Stormpike Cloth Girdle (19094) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Highlander's Cloth Girdle (20047, -1.25 DPS) [rep]; Wisdom of the Timbermaw (19047, -1.49 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 42.3 spell_power points (18.25 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -2.03 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -3.65 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.4 spell_power points (10.96 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -1.26 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.17 DPS) [quest]; Deep Rooted Ring (19109, -2.51 DPS) [world]; Naglering (11669, -11.66 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Deep Rooted Ring (19109, -0.53 DPS) [world]; Channeler's Ring (272406, -1.60 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+7.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.59 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.79 DPS) [crafted]; Ankh of Life (1713, -4.17 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.53 DPS, sim-verified) [vendor]; Uther's Strength (11302, -6.90 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -7.10 DPS) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.47 DPS) [pvp]; Sageclaw (20070, -5.16 DPS) [rep]; Ironbark Staff (20069, -17.14 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (12.51 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Tome of Shadow Force (19309, -2.23 DPS, sim-verified) [rep]; Grand Marshal's Tome of Restoration (234590, -3.33 DPS) [pvp] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.49 DPS) [dungeon]; Bonecreeper Stylus (13938, -1.29 DPS, sim-verified) [dungeon]; Greater Eternal Wand (249237, -1.36 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Stormpike Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 05000000000000000-2035103021220001351-0550005000000000)

Set DPS (verified): 795.2. Weights run: 3.6s. Verify run: 21.7s. 1299 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.470 ± 0.046, crit=0.975 ± 0.028 per rating point (14 rating = 1%, 13.653 per %), hit=1.196 ± 0.090 per rating point (10 rating = 1%, 11.965 per %), spell_haste=11.991 ± 1.136, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.887 ± 0.002, fire_power=0.113 ± 0.000, spirit=not significant (-0.131 ± 0.199)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cap of the Scarlet Savant (12752, +0.00 DPS) [quest]; Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -0.30 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.8 spell_power points (21.89 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.25 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.98 DPS) [pvp]; Shroud of the Nathrezim (18720, -7.48 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fel Cape (279269, -2.84 DPS) [crafted]; Crystalline Threaded Cape (20697, -5.86 DPS) [world]; Shroud of Arcane Mastery (22330, -8.69 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.2 spell_power points (29.91 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.14 DPS) [pvp]; Robe of Winter Night (14136, -6.33 DPS, sim-verified) [crafted]; Knight-Captain's Dreadweave Tunic (227096, -9.43 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.8 spell_power points (15.34 DPS) | yes | Runecloth Cuffs (254123, -5.99 DPS) [crafted]; Netherpearl Cuffs (254067, -6.52 DPS) [crafted]; Arena Wristguards (18709, -7.21 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 41.0 spell_power points (24.41 DPS) | yes | Marshal's Dreadweave Gloves (231586, -4.87 DPS) [pvp]; Hands of Power (13253, -5.61 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Handwraps (227100, -10.79 DPS) [pvp] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 30.5 spell_power points (18.14 DPS) | yes | Ban'thok Sash (11662, -0.80 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -4.63 DPS) [rep]; Black Waistcord (254087, -5.59 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 64.0 spell_power points (38.11 DPS) | yes | Marshal's Dreadweave Leggings (231587, -10.76 DPS) [pvp]; Skyshroud Leggings (13170, -14.06 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -17.80 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Maleki's Footwraps (18735, +0.00 DPS) [dungeon]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.57 DPS) [dungeon]; Maiden's Circle (13001, -2.58 DPS) [world_drop]; Naglering (11669, -14.57 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Songstone of Ironforge (12543, -0.40 DPS) [quest]; Maiden's Circle (13001, -0.40 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -1.64 DPS, sim-verified) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, -2.51 DPS, sim-verified) [dungeon]; Serenity Field (272439, -4.76 DPS) [vendor]; Uther's Strength (11302, -9.53 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (-5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Uther's Strength (11302, -2.84 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -1.93 DPS) [pvp]; Staff of Jordan (873, -21.83 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -25.64 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.77 DPS) [dungeon]; Greater Eternal Wand (249237, -1.96 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 00000000000000000-2035100000000000000-0000000000000000)

Set DPS (verified): 56.0. Weights run: 3.1s. Verify run: 3.3s. 176 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.303 ± 0.011, crit=0.025 ± 0.001 per rating point (14 rating = 1%, 0.353 per %), hit=0.059 ± 0.002 per rating point (10 rating = 1%, 0.595 per %), spell_haste=0.257 ± 0.028, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.813 ± 0.001, fire_power=0.187 ± 0.001, spirit=0.153 ± 0.025

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.9 spell_power points (2.34 DPS) | yes | Pristine Circlet (253949, -0.88 DPS, sim-verified) [crafted]; Flame Circlet (253953, -1.81 DPS) [crafted]; Shadow Goggles (4373, -2.10 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.5 spell_power points (0.12 DPS) | yes | Scholarly Pendant (277203, -0.04 DPS) [quest] |
| shoulder | Magician's Mantle (12998) (or Reinforced Woolen Shoulders (4315)) | World drop [world_drop] | 5.0 spell_power points (1.31 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.26 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Black Whelp Cloak (7283, -0.42 DPS) [crafted]; Battle Healer's Cloak (20427, -0.49 DPS, sim-verified) [rep]; Mystic Shawl (3449, -0.68 DPS) [quest] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 11.4 spell_power points (2.98 DPS) | yes | Filigreed Shadow Gown (253907, -1.11 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.55 DPS) [crafted]; Scholarly Robes (2034, -1.72 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.79 DPS) | yes | Owlbeard Bracers (16981, -0.52 DPS) [quest]; Windsong Bangles (263336, -0.56 DPS, sim-verified) [quest]; Crystalline Cuffs (14148, -0.67 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.84 DPS) | yes | Gnoll Casting Gloves (892, -0.28 DPS, sim-verified) [world]; Shadow Gloves (253919, -0.34 DPS) [crafted]; Blight Gloves (279877, -0.52 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 5.7 spell_power points (1.49 DPS) | yes | Novice Arcanist's Sash (253885, -0.44 DPS) [crafted]; Pristine Sash (253925, -0.51 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.85 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 9.6 spell_power points (2.50 DPS) | yes | Filigreed Pristine Leggings (253937, -0.77 DPS) [crafted]; Silk-threaded Trousers (1929, -0.83 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.99 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.84 DPS) | yes | Shadow Boots (253895, -0.57 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.59 DPS) [crafted]; Pristine Boots (253889, -1.05 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.31 DPS) | yes | Band of Purification (12996, -1.07 DPS) [world_drop]; Demon Band (12054, -1.23 DPS) [world_drop]; Horn Ring (17692, -1.23 DPS) [quest] |
| finger2 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | 1.0 spell_power points (0.26 DPS) | yes | Band of Purification (12996, -0.02 DPS) [world_drop]; Demon Band (12054, -0.18 DPS) [world_drop]; Horn Ring (17692, -0.18 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.15 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -0.32 DPS) [quest]; Crescent Staff (6505, -0.72 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 1.1 spell_power points (0.29 DPS) | yes | Grayson's Torch (1172, -0.13 DPS) [quest]; Nightglow Concoction (3451, -0.17 DPS) [quest]; Mystic's Sphere (15946, -0.17 DPS) [world_drop] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.26 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Furlfeather Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Sizzle Stick

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (orc, 00000000000000000-2035113101210001000-0000000000000000)

Set DPS (verified): 92.8. Weights run: 3.1s. Verify run: 2.6s. 302 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.091 ± 0.004, crit=0.040 ± 0.001 per rating point (14 rating = 1%, 0.562 per %), hit=0.092 ± 0.004 per rating point (10 rating = 1%, 0.924 per %), spell_haste=0.281 ± 0.051, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.810 ± 0.001, fire_power=0.190 ± 0.001, spirit=not significant (0.069 ± 0.041)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.4 spell_power points (3.24 DPS) | yes | Filigreed Shadow Circlet (253981, -0.33 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.69 DPS) [crafted]; Shadow Circlet (253955, -0.71 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 0.6 spell_power points (0.18 DPS) | yes | Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Darkspear Warding Pendant (272075, -0.05 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 spell_power points (2.79 DPS) | yes | Death Speaker Mantle (6685, -0.80 DPS) [dungeon]; Magician's Mantle (12998, -1.14 DPS) [world_drop]; Invoker's Mantle (215365, -1.28 DPS, sim-verified) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.08 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Battle Healer's Cloak (19529, -0.28 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.69 DPS) | yes | Necrology Robes (2292, -0.34 DPS) [dungeon]; Shadow Gown (253967, -0.44 DPS) [crafted]; Beguiler Robes (7728, -0.67 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.56 DPS) | yes | Phoenix Bindings (210781, -1.85 DPS) [crafted]; Tabitha's Cuffs (251486, -1.99 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -2.22 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.9 spell_power points (2.25 DPS) | yes | Hands of Darkness (7047, -0.08 DPS) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Jutebraid Gloves (10654, -0.42 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.3 spell_power points (3.20 DPS) | yes | Warsong Sash (16975, -0.08 DPS) [quest]; Belt of Arugal (6392, -0.47 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.93 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.6 spell_power points (3.58 DPS) | yes | Shadow Leggings (253993, -0.09 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.82 DPS) [crafted]; Necromancer Leggings (2277, -1.00 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.4 spell_power points (3.25 DPS) | yes | Gilded Slippers (254001, -1.13 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.14 DPS) [dungeon]; Spidersilk Boots (4320, -1.16 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.99 DPS) | yes | Ring of Power Regulation (273030, -0.28 DPS) [dungeon]; Solid Gold Band (279845, -0.75 DPS) [quest]; Electrocutioner Lagnut (9447, -1.00 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.5 spell_power points (1.83 DPS) | yes | Ring of Power Regulation (273030, -0.13 DPS) [dungeon]; Solid Gold Band (279845, -0.60 DPS) [quest]; Electrocutioner Lagnut (9447, -0.84 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.21 DPS) [dungeon]; Spine of the Devourer (273040, -1.21 DPS) [world]; Tendonscraper Dagger (273053, -1.47 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.56 DPS) | yes | Orb of Mystic Insight (249394, -0.53 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.10 DPS) [world_drop]; Orb of Souls (249395, -1.30 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, -0.26 DPS) [quest]; Wand of Eventide (5214, -0.28 DPS) [world_drop]; Greater Mystic Wand (217287, -0.28 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (orc, 10000000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 164.3. Weights run: 3.2s. Verify run: 1.8s. 392 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.162 ± 0.011, crit=0.146 ± 0.003 per rating point (14 rating = 1%, 2.051 per %), hit=0.383 ± 0.016 per rating point (10 rating = 1%, 3.834 per %), spell_haste=2.661 ± 0.196, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.839 ± 0.001, fire_power=0.161 ± 0.000, spirit=not significant (0.071 ± 0.102)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.98 DPS) | yes | Living Cowl (5608, -2.67 DPS, sim-verified) [world]; Augural Shroud (2620, -2.79 DPS) [world]; Filigreed Shadow Circlet (253981, -3.07 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.5 spell_power points (2.82 DPS) | yes | Glowing Eye of Mordresh (10769, -0.07 DPS) [dungeon]; Necklace of Calisea (1714, -2.28 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.29 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (3.47 DPS) | yes | Green Silken Shoulders (7057, -0.22 DPS) [crafted]; Inquisitor's Shawl (19507, -0.45 DPS) [dungeon]; Berylline Pads (4197, -0.61 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.21 DPS) [crafted]; Guardian Cloak (5965, -0.21 DPS) [crafted]; Battle Healer's Cloak (19528, -0.36 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (7.75 DPS) | yes | Elemental Raiment (9434, -0.53 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.95 DPS) [crafted]; Shadoweave Robe (10004, -1.73 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.99 DPS) | yes | Condor Bracers (15864, -0.73 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.23 DPS) [quest]; Earthen Silk Cuffs (254019, -1.54 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (6.36 DPS) | yes | Black Mageweave Gloves (10003, -1.38 DPS) [crafted]; Shadoweave Gloves (10023, -1.94 DPS, sim-verified) [crafted]; Black Handwraps (254027, -2.08 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.6 spell_power points (4.87 DPS) | yes | Black Cord (254043, -0.39 DPS) [crafted]; Star Belt (4329, -0.45 DPS) [crafted]; Warsong Sash (16975, -1.21 DPS) [quest] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.6 spell_power points (5.86 DPS) | yes | Red Mageweave Pants (10009, -0.56 DPS) [crafted]; Shadow Leggings (253993, -1.46 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.65 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.5 spell_power points (8.14 DPS) | yes | Black Slippers (254007, -4.20 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -5.34 DPS) [crafted]; Spidersilk Boots (4320, -5.60 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.6 spell_power points (4.17 DPS) | yes | Advisor's Ring (19520, -1.18 DPS) [rep]; Reedknot Ring (9622, -1.85 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (3.59 DPS) | yes | Advisor's Ring (19520, -0.60 DPS) [rep]; Reedknot Ring (9622, -1.27 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.63 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -8.04 DPS) [dungeon]; Swinetusk Shank (6691, -9.78 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.25 DPS) [quest]; Greater Mystic Wand (217287, -0.33 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 25400000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 240.8. Weights run: 3.2s. Verify run: 5.6s. 499 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.114 ± 0.015, crit=0.189 ± 0.004 per rating point (14 rating = 1%, 2.649 per %), hit=0.489 ± 0.021 per rating point (10 rating = 1%, 4.886 per %), spell_haste=2.201 ± 0.264, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.772 ± 0.001, fire_power=0.228 ± 0.000, spirit=not significant (0.207 ± 0.149)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.04 DPS) | yes | Dreamweave Circlet (10041, -0.89 DPS) [crafted]; Red Mageweave Headband (10033, -2.13 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.23 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.4 spell_power points (3.51 DPS) | yes | Glowing Eye of Mordresh (10769, -0.57 DPS) [dungeon]; Mindburst Medallion (11196, -0.64 DPS) [quest]; Lei of Lilies (1315, -2.36 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherpearl Shoulders (254055, +0.00 DPS) [crafted]; Shadoweave Shoulders (10028, -0.72 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.08 DPS) [quest]; Runecloth Cloak (13860, -2.24 DPS) [crafted]; Nightfall Drape (12465, -2.58 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.7 spell_power points (8.82 DPS) | yes | Dreamweave Vest (10021, -0.67 DPS) [crafted]; Acumen Robes (17775, -0.91 DPS) [quest]; Elemental Raiment (9434, -1.01 DPS) [world_drop] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 11.9 spell_power points (4.41 DPS) | yes | Spidertank Oilrag (9448, -1.06 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.13 DPS) [crafted]; Condor Bracers (15864, -1.81 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.9 spell_power points (7.40 DPS) | yes | Black Gloves (254101, -0.91 DPS) [crafted]; Black Mageweave Gloves (10003, -1.82 DPS) [crafted]; Runecloth Gloves (13863, -1.87 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.20 DPS) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshock Leggings (9484, +0.00 DPS) [dungeon]; Wizardweave Leggings (14132, -0.90 DPS) [crafted]; Senior Designer's Pantaloons (11841, -1.83 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.4 spell_power points (9.47 DPS) | yes | Black Sandals (254113, -3.34 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.29 DPS) [vendor]; Shadoweave Boots (10031, -4.68 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.84 DPS) | yes | Underworld Band (1980, -0.61 DPS) [world_drop]; Cyclopean Band (11824, -0.88 DPS) [dungeon]; Philanthropist's Ring (281635, -0.90 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.46 DPS) | yes | Underworld Band (1980, -0.23 DPS) [world_drop]; Cyclopean Band (11824, -0.51 DPS) [dungeon]; Philanthropist's Ring (281635, -0.53 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.36 DPS) [quest]; Ankh of Life (1713, -0.71 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.87 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.96 DPS) [quest]; Ankh of Life (1713, -1.31 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.46 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -2.37 DPS) [world]; Kindling Stave (11750, -2.85 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.36 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.37 DPS) [crafted]; Thunderwood (13062, -0.74 DPS) [world_drop]; Twisted Nether Wand (249144, -0.74 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Lesser Eternal Wand

No-known-source sample (15 of 499, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522000000000000-2035113101220001350-0050000000000000)

Set DPS (verified): 410.6. Weights run: 3.4s. Verify run: 20.2s. 1293 eligible items had no known source.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.449 ± 0.026, crit=0.207 ± 0.007 per rating point (14 rating = 1%, 2.901 per %), hit=0.616 ± 0.026 per rating point (10 rating = 1%, 6.159 per %), spell_haste=3.025 ± 0.347, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.765 ± 0.001, fire_power=0.235 ± 0.000, spirit=not significant (0.195 ± 0.156)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 35.2 spell_power points (15.17 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.37 DPS) [pvp]; Red Mageweave Headband (10033, -5.39 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -2.04 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.27 DPS) [crafted]; Burial Shawl (18681, -2.22 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.54 DPS, sim-verified) [world]; Spritecaster Cape (11623, -3.40 DPS) [dungeon]; Deep Woodlands Cloak (19121, -4.19 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.0 spell_power points (21.60 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.13 DPS) [pvp]; Robe of Winter Night (14136, -6.82 DPS, sim-verified) [crafted]; Legionnaire's Dreadweave Tunic (227094, -6.93 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.0 spell_power points (11.64 DPS) | yes | Runecloth Cuffs (254123, -4.95 DPS) [crafted]; Netherpearl Cuffs (254067, -5.56 DPS) [crafted]; Nethergeld Cuffs (254061, -6.84 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Blood Guard's Dreadweave Handwraps (227099, +0.00 DPS) [pvp]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp] |
| waist | Frostwolf Cloth Belt (19090) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Defiler's Cloth Girdle (20163, -1.25 DPS) [rep]; Wisdom of the Timbermaw (19047, -1.49 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.39 DPS) [dungeon]; Legionnaire's Dreadweave Legguards (227097, -2.01 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.08 DPS) [dungeon]; Omnicast Boots (11822, -1.67 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.17 DPS) [quest]; Deep Rooted Ring (19109, -2.51 DPS) [world]; Naglering (11669, -9.61 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Deep Rooted Ring (19109, -0.53 DPS) [world]; Channeler's Ring (272406, -1.60 DPS) [vendor] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+7.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.59 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.79 DPS) [crafted]; Rune of the Guard Captain (19120, -3.32 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -3.53 DPS, sim-verified) [vendor]; Uther's Strength (11302, -6.90 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -7.10 DPS) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.47 DPS) [pvp]; Mindfang (20214, -5.16 DPS) [rep]; Ironbark Staff (20220, -15.14 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (12.51 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Tome of Shadow Force (19309, -2.22 DPS, sim-verified) [rep]; High Warlord's Tome of Mending (234564, -3.33 DPS) [pvp] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.49 DPS) [dungeon]; Greater Eternal Wand (249237, -1.36 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Frostwolf Cloth Belt; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 05000000000000000-2035103021220001351-0550005000000000)

Set DPS (verified): 788.1. Weights run: 3.6s. Verify run: 19.1s. 1293 eligible items had no known source.

6 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.470 ± 0.046, crit=0.975 ± 0.028 per rating point (14 rating = 1%, 13.653 per %), hit=1.196 ± 0.090 per rating point (10 rating = 1%, 11.965 per %), spell_haste=11.991 ± 1.136, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.887 ± 0.002, fire_power=0.113 ± 0.000, spirit=not significant (-0.131 ± 0.199)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cap of the Scarlet Savant (12752, +0.00 DPS) [quest]; Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -0.30 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 36.8 spell_power points (21.89 DPS) | yes | Warlord's Dreadweave Mantle (231592, -2.25 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -2.98 DPS) [pvp]; Shroud of the Nathrezim (18720, -5.52 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.86 DPS) [world]; Fel Cape (279269, -6.71 DPS, sim-verified) [crafted]; Shroud of Arcane Mastery (22330, -8.69 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 50.2 spell_power points (29.91 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.14 DPS) [pvp]; Robe of Winter Night (14136, -5.90 DPS, sim-verified) [crafted]; Legionnaire's Dreadweave Tunic (227094, -9.43 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.8 spell_power points (15.34 DPS) | yes | Runecloth Cuffs (254123, -5.99 DPS) [crafted]; Netherpearl Cuffs (254067, -6.52 DPS) [crafted]; Arena Wristguards (18709, -7.21 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 41.0 spell_power points (24.41 DPS) | yes | General's Dreadweave Gloves (231589, -4.87 DPS) [pvp]; Hands of Power (13253, -7.97 DPS, sim-verified) [dungeon]; Blood Guard's Dreadweave Handwraps (227099, -10.79 DPS) [pvp] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 30.5 spell_power points (18.14 DPS) | yes | Ban'thok Sash (11662, -0.80 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -4.63 DPS) [rep]; Black Waistcord (254087, -5.59 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -1.68 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maleki's Footwraps (18735, +0.00 DPS) [dungeon]; General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.07 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.57 DPS) [dungeon]; Maiden's Circle (13001, -2.58 DPS) [world_drop]; Naglering (11669, -16.03 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Eye of Orgrimmar (12545, -0.40 DPS) [quest]; Maiden's Circle (13001, -0.40 DPS) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+12.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.73 DPS) [crafted]; Rune of the Guard Captain (19120, -2.16 DPS) [quest]; Uther's Strength (11302, -3.57 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.11 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -6.69 DPS) [crafted]; Rune of the Guard Captain (19120, -8.11 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -1.93 DPS) [pvp]; Staff of Jordan (873, -22.54 DPS, sim-verified) [world_drop]; Whiteout Staff (19101, -24.31 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.77 DPS) [dungeon]; Greater Eternal Wand (249237, -1.96 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Don Julio's Band; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

