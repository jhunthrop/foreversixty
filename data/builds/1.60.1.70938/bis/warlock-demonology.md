# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-2005003100000000000-0000000000000000)

Set DPS (verified): 56.9. Weights run: 3.7s. Verify run: 3.0s. 199 eligible items had no known source.

Race sweep: human 56.9 ± 0.0, gnome 56.8 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.148 ± 0.005, crit=0.058 ± 0.001 per rating point (14 rating = 1%, 0.816 per %), hit=0.156 ± 0.007 per rating point (10 rating = 1%, 1.562 per %), spell_haste=0.580 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.830 ± 0.001, fire_power=0.170 ± 0.000, spirit=not significant (0.117 ± 0.048)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.1 spell_power points (2.55 DPS) | yes | Pristine Circlet (253949, -0.88 DPS, sim-verified) [crafted]; Flame Circlet (253953, -2.03 DPS) [crafted]; Shadow Goggles (4373, -2.15 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.2 spell_power points (0.07 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.3 spell_power points (1.77 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.21 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.65 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.32 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.41 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.57 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.4 spell_power points (3.46 DPS) | yes | Filigreed Shadow Gown (253907, -1.02 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.75 DPS) [crafted]; Gray Woolen Robe (2585, -2.13 DPS) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.28 DPS) | yes | Mindthrust Bracers (1974, -0.07 DPS) [dungeon]; Bright Bracers (3647, -0.11 DPS) [world_drop]; Repurposed Hair Band (281256, -0.13 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.96 DPS) | yes | Shadow Gloves (253919, -0.21 DPS) [crafted]; Gnoll Casting Gloves (892, -0.28 DPS) [world]; Pristine Gloves (253913, -0.72 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.4 spell_power points (1.79 DPS) | yes | Pristine Sash (253925, -0.51 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted]; Novice Ardent's Sash (253887, -1.01 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.5 spell_power points (2.93 DPS) | yes | Filigreed Pristine Leggings (253937, -0.88 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.97 DPS) [dungeon]; Colorful Kilt (10048, -1.37 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (2.12 DPS) | yes | Shadow Boots (253895, -0.41 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.84 DPS) [crafted]; Pristine Boots (253889, -1.16 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.3 spell_power points (1.48 DPS) | yes | Furlfeather Ring (286745, -1.20 DPS) [world]; Lavishly Jeweled Ring (1156, -1.23 DPS) [dungeon]; Band of Purification (12996, -1.29 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.40 DPS) | yes | Lavishly Jeweled Ring (1156, -1.15 DPS) [dungeon]; Band of Purification (12996, -1.20 DPS) [world_drop]; Furlfeather Ring (286745, -1.22 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.73 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.37 DPS) [dungeon]; Witching Stave (1484, -1.61 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.1 spell_power points (0.87 DPS) | yes | Antipodean Rod (2879, -0.59 DPS) [world_drop]; Mathystra Channeling Orb (281291, -0.67 DPS) [quest]; Onu's Patience (281292, -0.71 DPS) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.28 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (human, 00000000000000000-2035113101210001000-0000000000000000)

Set DPS (verified): 91.9. Weights run: 3.6s. Verify run: 3.3s. 326 eligible items had no known source.

Race sweep: human 91.9 ± 0.1, gnome 91.5 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.079 ± 0.004, crit=0.040 ± 0.001 per rating point (14 rating = 1%, 0.562 per %), hit=0.092 ± 0.004 per rating point (10 rating = 1%, 0.924 per %), spell_haste=0.281 ± 0.051, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.810 ± 0.001, fire_power=0.190 ± 0.001, spirit=not significant (0.069 ± 0.041)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.4 spell_power points (3.24 DPS) | yes | Filigreed Shadow Circlet (253981, -0.02 DPS) [crafted]; Silk Headband (7050, -0.69 DPS) [crafted]; Shadow Circlet (253955, -0.71 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 0.6 spell_power points (0.17 DPS) | yes | Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Pendant of Myzrael (4614, -0.05 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.7 spell_power points (2.76 DPS) | yes | Invoker's Mantle (215365, -0.49 DPS, sim-verified) [crafted]; Moonlit Amice (11884, -0.77 DPS) [quest]; Death Speaker Mantle (6685, -0.81 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted]; Caretaker's Cape (19533, -0.21 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.69 DPS) | yes | Necrology Robes (2292, -0.36 DPS) [dungeon]; Shadow Gown (253967, -0.47 DPS) [crafted]; Beguiler Robes (7728, -0.71 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.56 DPS) | yes | Phoenix Bindings (210781, -2.24 DPS, sim-verified) [crafted]; Windsong Bangles (263336, -2.27 DPS) [quest]; Nightsky Wristbands (6407, -2.36 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.8 spell_power points (2.23 DPS) | yes | Hands of Darkness (7047, -0.06 DPS) [crafted]; Serpent Gloves (5970, -0.24 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.52 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.2 spell_power points (3.19 DPS) | yes | Belt of Arugal (6392, -0.47 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.92 DPS) [dungeon]; Invoker's Cord (215366, -1.09 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.6 spell_power points (3.58 DPS) | yes | Shadow Leggings (253993, -0.11 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.84 DPS) [crafted]; Necromancer Leggings (2277, -1.04 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.4 spell_power points (3.23 DPS) | yes | Gilded Slippers (254001, -1.12 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.12 DPS) [dungeon]; Spidersilk Boots (4320, -1.15 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.99 DPS) | yes | Ring of Power Regulation (273030, -0.28 DPS) [dungeon]; Minor Channeling Ring (1449, -0.52 DPS) [quest]; Runemender's Seal (273031, -0.88 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.4 spell_power points (1.82 DPS) | yes | Ring of Power Regulation (273030, -0.11 DPS) [dungeon]; Minor Channeling Ring (1449, -0.35 DPS) [quest]; Runemender's Seal (273031, -0.71 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.21 DPS) [dungeon]; Spine of the Devourer (273040, -1.21 DPS) [world]; Tendonscraper Dagger (273053, -1.46 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.56 DPS) | yes | Orb of Mystic Insight (249394, -0.43 DPS) [crafted]; Orb of Mistmantle (13031, -1.10 DPS) [world_drop]; Orb of Souls (249395, -1.30 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, -0.26 DPS) [quest]; Wand of Eventide (5214, -0.28 DPS) [world_drop]; Greater Mystic Wand (217287, -0.28 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 10000000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 158.9. Weights run: 3.7s. Verify run: 6.7s. 455 eligible items had no known source.

Race sweep: human 158.9 ± 0.1, gnome 158.4 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.155 ± 0.011, crit=0.144 ± 0.003 per rating point (14 rating = 1%, 2.022 per %), hit=0.368 ± 0.015 per rating point (10 rating = 1%, 3.685 per %), spell_haste=2.872 ± 0.195, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.839 ± 0.001, fire_power=0.161 ± 0.000, spirit=not significant (0.368 ± 0.093)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.93 DPS) | yes | Miner's Hat of the Deep (9429, -1.43 DPS, sim-verified) [dungeon]; Papal Fez (9431, -2.33 DPS) [dungeon]; Holy Shroud (2721, -2.57 DPS) [world_drop] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 10.6 spell_power points (3.49 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.52 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.77 DPS) [dungeon]; Amberglow Talisman (10824, -2.28 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.23 DPS) [crafted]; Mantle of Doan (7712, -0.36 DPS, sim-verified) [dungeon]; Inquisitor's Shawl (19507, -0.46 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19532, -0.45 DPS) [rep]; Long Silken Cloak (4326, -0.80 DPS) [crafted]; Guardian Cloak (5965, -0.80 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.8 spell_power points (8.17 DPS) | yes | Dreamweave Vest (10021, -0.07 DPS) [crafted]; Elemental Raiment (9434, -1.24 DPS) [world_drop]; Shadoweave Robe (10004, -1.51 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Arcane Runed Bracers (4744, -0.33 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.66 DPS) [quest]; Earthen Silk Cuffs (254019, -1.04 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.2 spell_power points (6.99 DPS) | yes | Shadoweave Gloves (10023, -1.51 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.04 DPS) [crafted]; Black Handwraps (254027, -2.76 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.05 DPS) [crafted]; Black Cord (254043, -0.53 DPS, sim-verified) [crafted]; Gilded Cord (254037, -1.05 DPS) [crafted] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.6 spell_power points (5.81 DPS) | yes | Red Mageweave Pants (10009, -0.58 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.76 DPS) [dungeon]; Shadow Leggings (253993, -0.97 DPS) [crafted] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 13.5 spell_power points (4.44 DPS) | yes | Wingborne Boots (15104, -1.25 DPS) [quest]; Gilded Slippers (254001, -1.29 DPS) [crafted]; Thoughtcast Boots (10578, -1.29 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.5 spell_power points (4.13 DPS) | yes | Lorekeeper's Ring (19524, -1.16 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.49 DPS) [quest]; Reedknot Ring (9622, -1.82 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (3.55 DPS) | yes | Lorekeeper's Ring (19524, -0.25 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.92 DPS) [quest]; Reedknot Ring (9622, -1.25 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -7.52 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -8.56 DPS) [dungeon]; Swinetusk Shank (6691, -10.17 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Gravestone Scepter (7001, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Shadoweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 455, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 25400000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 244.0. Weights run: 3.7s. Verify run: 5.4s. 594 eligible items had no known source.

Race sweep: human 244.0 ± 0.1, gnome 243.5 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.132 ± 0.015, crit=0.187 ± 0.004 per rating point (14 rating = 1%, 2.622 per %), hit=0.475 ± 0.019 per rating point (10 rating = 1%, 4.747 per %), spell_haste=2.641 ± 0.262, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.772 ± 0.001, fire_power=0.228 ± 0.000, spirit=not significant (0.250 ± 0.152)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.03 DPS) | yes | Dreamweave Circlet (10041, -0.98 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -1.99 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.23 DPS) [crafted] |
| neck | Talvash's Enhancing Necklace (7673) | Restoring the Necklace [quest] | 12.1 spell_power points (4.51 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.16 DPS) [quest]; Triune Amulet (7722, -1.38 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -1.51 DPS) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.2 spell_power points (6.39 DPS) | yes | Netherpearl Shoulders (254055, -0.36 DPS, sim-verified) [crafted]; Rotgrip Mantle (17732, -0.68 DPS) [dungeon]; Shadoweave Shoulders (10028, -1.52 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -1.54 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.71 DPS) [dungeon]; Caretaker's Cape (19531, -2.89 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (8.93 DPS) | yes | Elemental Raiment (9434, -1.13 DPS) [world_drop]; Dreamweave Vest (10021, -1.51 DPS, sim-verified) [crafted]; Embrace of the Wind Serpent (12462, -1.57 DPS) [world] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 12.2 spell_power points (4.53 DPS) | yes | Nethergeld Cuffs (254061, -1.12 DPS, sim-verified) [crafted]; Arcane Runed Bracers (4744, -1.19 DPS) [quest]; Spidertank Oilrag (9448, -1.19 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (7.53 DPS) | yes | Black Gloves (254101, -1.18 DPS, sim-verified) [crafted]; Runecloth Gloves (13863, -1.80 DPS) [crafted]; Black Mageweave Gloves (10003, -1.96 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.1 spell_power points (6.71 DPS) | yes | Black Waistcord (254087, -0.82 DPS) [crafted]; Ghostweave Cord (254073, -0.98 DPS, sim-verified) [crafted]; Satyrmane Sash (17755, -1.02 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 24.3 spell_power points (9.03 DPS) | yes | Felcloth Pants (14107, -0.79 DPS, sim-verified) [crafted]; Wizardweave Leggings (14132, -1.98 DPS) [crafted]; Senior Designer's Pantaloons (11841, -2.56 DPS) [dungeon] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 16.6 spell_power points (6.17 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.01 DPS, sim-verified) [vendor]; Mistwalker Boots (10629, -1.16 DPS) [dungeon]; Shadoweave Boots (10031, -1.23 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 13.6 spell_power points (5.04 DPS) | yes | Lorekeeper's Ring (19523, -0.58 DPS) [rep]; Underworld Band (1980, -0.78 DPS) [world_drop]; Cyclopean Band (11824, -0.98 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.83 DPS) | yes | Lorekeeper's Ring (19523, -0.37 DPS) [rep]; Underworld Band (1980, -0.57 DPS) [world_drop]; Cyclopean Band (11824, -0.77 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.47 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.66 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -4.82 DPS, sim-verified) [world_drop] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (5.20 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.15 DPS) [dungeon]; Umbral Crystal (13029, -0.33 DPS) [world_drop]; Orb of Dar'Orahil (15108, -0.94 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.37 DPS) [crafted]; Burning Sliver (5249, -0.74 DPS) [quest]; Twisted Nether Wand (249144, -0.74 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Talvash's Enhancing Necklace; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Black Sandals; finger1: Brainlash; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Orb of the Forgotten Seer; ranged: Lesser Eternal Wand

No-known-source sample (15 of 594, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522000000000000-2035113101220001350-0050000000000000)

Set DPS (verified): 415.4. Weights run: 3.5s. Verify run: 18.1s. 1460 eligible items had no known source.

Race sweep: human 415.4 ± 0.2, gnome 414.3 ± 0.3 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.173 ± 0.028, crit=0.298 ± 0.010 per rating point (14 rating = 1%, 4.175 per %), hit=0.764 ± 0.035 per rating point (10 rating = 1%, 7.637 per %), spell_haste=3.296 ± 0.480, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.776 ± 0.001, fire_power=0.224 ± 0.000, spirit=not significant (0.235 ± 0.250)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 43.7 spell_power points (18.85 DPS) | yes | Field Marshal's Coronal (231584, -3.25 DPS) [pvp]; Crimson Felt Hat (18727, -4.53 DPS, sim-verified) [dungeon]; Lieutenant Commander's Dreadweave Cowl (227093, -6.64 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 18.0 spell_power points (7.78 DPS) | yes | Dark Advisor's Pendant (18691, -1.20 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -1.61 DPS) [world_drop]; Archlight Talisman (15856, -1.85 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 26.6 spell_power points (11.50 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.55 DPS, sim-verified) [crafted]; Burial Shawl (18681, -1.67 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.76 DPS, sim-verified) [world]; Spritecaster Cape (11623, -3.70 DPS) [dungeon]; Shroud of the Exile (15421, -4.16 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (20.52 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.92 DPS) [pvp]; Robe of Winter Night (14136, -6.44 DPS, sim-verified) [crafted]; Knight-Captain's Dreadweave Tunic (227096, -8.24 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.0 spell_power points (10.80 DPS) | yes | Runecloth Cuffs (254123, -5.31 DPS) [crafted]; Netherpearl Cuffs (254067, -5.42 DPS) [crafted]; Wyrmthalak's Shackles (13958, -6.26 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.4 spell_power points (11.41 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -2.05 DPS) [pvp]; Earth Warder's Gloves (21318, -2.88 DPS, sim-verified) [quest] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 21.4 spell_power points (9.22 DPS) | yes | Highlander's Cloth Girdle (20047, -0.93 DPS) [rep]; Stormpike Cloth Girdle (19094, -1.43 DPS, sim-verified) [rep]; Valconian Sash (12083, -1.89 DPS) [quest] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 42.1 spell_power points (18.16 DPS) | yes | Marshal's Dreadweave Leggings (231587, -0.77 DPS) [pvp]; Skyshroud Leggings (13170, -4.32 DPS) [dungeon]; Ghostweave Pants (14144, -4.91 DPS) [crafted] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 23.9 spell_power points (10.30 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.78 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.98 DPS) [pvp] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.75 DPS) [quest]; The Postmaster's Seal (13392, -1.99 DPS) [dungeon]; Naglering (11669, -11.00 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.61 DPS) [quest]; The Postmaster's Seal (13392, -1.85 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -2.35 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+8.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -4.47 DPS) [dungeon]; Serenity Field (272439, -4.75 DPS) [vendor]; Uther's Strength (11302, -5.18 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS) [quest]; Uther's Strength (11302, -0.38 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, -3.44 DPS) [dungeon]; Grand Marshal's Stave (234571, -4.03 DPS) [pvp]; Ironbark Staff (20069, -14.37 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (12.51 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.75 DPS) [pvp]; Tome of Shadow Force (19309, -2.21 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.45 DPS) [quest] |
| ranged | Oblivion's Touch (18761) | Scholomance: Death Knight Darkreaver [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Brilliant Wand (249385, +0.00 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; legs: Sentinel's Silk Leggings; feet: Venomspew Footpads; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Oblivion's Touch

No-known-source sample (15 of 1460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 05000000000000000-2035103021220001351-0550005000000000)

Set DPS (verified): 817.6. Weights run: 4.0s. Verify run: 21.3s. 1460 eligible items had no known source.

Race sweep: human 817.6 ± 0.8, gnome 815.2 ± 0.8 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.417 ± 0.043, crit=0.927 ± 0.026 per rating point (14 rating = 1%, 12.985 per %), hit=1.212 ± 0.084 per rating point (10 rating = 1%, 12.118 per %), spell_haste=8.707 ± 1.066, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.889 ± 0.001, fire_power=0.111 ± 0.000, spirit=not significant (0.129 ± 0.179)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 46.8 spell_power points (28.59 DPS) | yes | Field Marshal's Coronal (231584, -2.92 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -3.24 DPS) [pvp]; Crimson Felt Hat (18727, -5.43 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 24.5 spell_power points (14.94 DPS) | yes | Dark Advisor's Pendant (18691, -2.04 DPS) [dungeon]; Star of Mystaria (12103, -2.98 DPS) [dungeon]; Emberfury Talisman (12929, -3.34 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 35.4 spell_power points (21.63 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -2.03 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.05 DPS) [pvp]; Shroud of the Nathrezim (18720, -6.66 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.98 DPS) [world]; Spritecaster Cape (11623, -8.66 DPS) [dungeon]; Cloak of Flames (3475, -10.76 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.8 spell_power points (30.39 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.73 DPS) [pvp]; Robe of Winter Night (14136, -5.58 DPS, sim-verified) [crafted]; Knight-Captain's Dreadweave Tunic (227096, -10.02 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.2 spell_power points (16.03 DPS) | yes | Runecloth Cuffs (254123, -6.76 DPS) [crafted]; Netherpearl Cuffs (254067, -6.79 DPS) [crafted]; Arena Wristguards (18709, -7.63 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 39.1 spell_power points (23.91 DPS) | yes | Marshal's Dreadweave Gloves (231586, -4.06 DPS) [pvp]; Hands of Power (13253, -6.42 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Handwraps (227100, -10.06 DPS) [pvp] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 29.5 spell_power points (18.01 DPS) | yes | Stormpike Cloth Girdle (19094, -4.47 DPS) [rep]; Ban'thok Sash (11662, -5.10 DPS, sim-verified) [dungeon]; Black Waistcord (254087, -5.40 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 62.1 spell_power points (37.96 DPS) | yes | Marshal's Dreadweave Leggings (231587, -10.52 DPS) [pvp]; Skyshroud Leggings (13170, -16.58 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -17.54 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.0 spell_power points (15.28 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.20 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -2.45 DPS) [dungeon]; Maiden's Circle (13001, -3.85 DPS) [world_drop]; Naglering (11669, -17.02 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.68 DPS) [dungeon]; Maiden's Circle (13001, -2.08 DPS) [world_drop]; Songstone of Ironforge (12543, -2.55 DPS) [quest] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -3.59 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -3.61 DPS, sim-verified) [crafted]; Serenity Field (272439, -6.72 DPS) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.24 DPS) [pvp]; Headmaster's Charge (13937, -16.32 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -19.80 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.71 DPS) [dungeon]; Brilliant Wand (249385, -0.71 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -1.27 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Channeler's Ring; finger2: Don Julio's Band; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 00000000000000000-2005003100000000000-0000000000000000)

Set DPS (verified): 56.1. Weights run: 3.7s. Verify run: 3.3s. 193 eligible items had no known source.

Race sweep: undead 56.6 ± 0.0, orc 56.1 ± 0.0, troll 55.4 ± 0.0 (dps); undead beats the guide race orc by 0.5 (0.90%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.148 ± 0.005, crit=0.058 ± 0.001 per rating point (14 rating = 1%, 0.816 per %), hit=0.156 ± 0.007 per rating point (10 rating = 1%, 1.562 per %), spell_haste=0.580 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.830 ± 0.001, fire_power=0.170 ± 0.000, spirit=not significant (0.117 ± 0.048)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.1 spell_power points (2.55 DPS) | yes | Pristine Circlet (253949, -0.88 DPS, sim-verified) [crafted]; Flame Circlet (253953, -2.03 DPS) [crafted]; Shadow Goggles (4373, -2.15 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.4 spell_power points (0.10 DPS) | yes | Scholarly Pendant (277203, -0.03 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.3 spell_power points (1.77 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.37 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.65 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (20427, -0.33 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.41 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.57 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.4 spell_power points (3.46 DPS) | yes | Filigreed Shadow Gown (253907, -1.16 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.75 DPS) [crafted]; Gray Woolen Robe (2585, -2.13 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.84 DPS) | yes | Windsong Bangles (263336, -0.56 DPS) [quest]; Mindthrust Bracers (1974, -0.63 DPS) [dungeon]; Owlbeard Bracers (16981, -0.67 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.96 DPS) | yes | Shadow Gloves (253919, -0.27 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.28 DPS) [world]; Blight Gloves (279877, -0.56 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.4 spell_power points (1.79 DPS) | yes | Pristine Sash (253925, -0.51 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted]; Novice Ardent's Sash (253887, -1.01 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.5 spell_power points (2.93 DPS) | yes | Filigreed Pristine Leggings (253937, -0.88 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.97 DPS) [dungeon]; Colorful Kilt (10048, -1.37 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (2.12 DPS) | yes | Shadow Boots (253895, -0.71 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.84 DPS) [crafted]; Pristine Boots (253889, -1.16 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.40 DPS) | yes | Lavishly Jeweled Ring (1156, -1.15 DPS) [dungeon]; Loop of Sacrifice (281673, -1.19 DPS) [quest]; Band of Purification (12996, -1.20 DPS) [world_drop] |
| finger2 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | 1.0 spell_power points (0.28 DPS) | yes | Lavishly Jeweled Ring (1156, -0.03 DPS) [dungeon]; Loop of Sacrifice (281673, -0.07 DPS) [quest]; Band of Purification (12996, -0.08 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.11 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -0.48 DPS) [quest]; Crescent Staff (6505, -0.52 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 1.0 spell_power points (0.29 DPS) | yes | Seer's Fine Stein (7608, -0.13 DPS) [world_drop]; Grayson's Torch (1172, -0.15 DPS) [quest]; Spellbinder Orb (15926, -0.16 DPS) [dungeon] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.28 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Furlfeather Ring; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Sizzle Stick

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade

### Band 30 (orc, 00000000000000000-2035113101210001000-0000000000000000)

Set DPS (verified): 92.7. Weights run: 3.6s. Verify run: 3.5s. 319 eligible items had no known source.

Race sweep: undead 93.1 ± 0.1, orc 92.7 ± 0.1, troll 91.7 ± 0.1 (dps); undead beats the guide race orc by 0.4 (0.46%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.079 ± 0.004, crit=0.040 ± 0.001 per rating point (14 rating = 1%, 0.562 per %), hit=0.092 ± 0.004 per rating point (10 rating = 1%, 0.924 per %), spell_haste=0.281 ± 0.051, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.810 ± 0.001, fire_power=0.190 ± 0.001, spirit=not significant (0.069 ± 0.041)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.4 spell_power points (3.24 DPS) | yes | Filigreed Shadow Circlet (253981, -0.33 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.69 DPS) [crafted]; Shadow Circlet (253955, -0.71 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 0.6 spell_power points (0.17 DPS) | yes | Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.03 DPS) [world]; Pendant of Myzrael (4614, -0.05 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.7 spell_power points (2.76 DPS) | yes | Death Speaker Mantle (6685, -0.81 DPS) [dungeon]; Magician's Mantle (12998, -1.14 DPS) [world_drop]; Invoker's Mantle (215365, -1.28 DPS, sim-verified) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.08 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Battle Healer's Cloak (19529, -0.28 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.69 DPS) | yes | Necrology Robes (2292, -0.36 DPS) [dungeon]; Shadow Gown (253967, -0.47 DPS) [crafted]; Beguiler Robes (7728, -0.71 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.56 DPS) | yes | Phoenix Bindings (210781, -1.85 DPS) [crafted]; Tabitha's Cuffs (251486, -1.99 DPS, sim-verified) [quest]; Owlbeard Bracers (16981, -2.23 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.8 spell_power points (2.23 DPS) | yes | Hands of Darkness (7047, -0.06 DPS) [crafted]; Serpent Gloves (5970, -0.24 DPS) [dungeon]; Jutebraid Gloves (10654, -0.41 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.2 spell_power points (3.19 DPS) | yes | Warsong Sash (16975, -0.07 DPS) [quest]; Belt of Arugal (6392, -0.47 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.92 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.6 spell_power points (3.58 DPS) | yes | Shadow Leggings (253993, -0.11 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.84 DPS) [crafted]; Necromancer Leggings (2277, -1.04 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.4 spell_power points (3.23 DPS) | yes | Glinteye Slippers (273024, -1.12 DPS) [dungeon]; Gilded Slippers (254001, -1.13 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.15 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.99 DPS) | yes | Ring of Power Regulation (273030, -0.28 DPS) [dungeon]; Solid Gold Band (279845, -0.75 DPS) [quest]; Runemender's Seal (273031, -0.88 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.4 spell_power points (1.82 DPS) | yes | Ring of Power Regulation (273030, -0.11 DPS) [dungeon]; Solid Gold Band (279845, -0.58 DPS) [quest]; Runemender's Seal (273031, -0.71 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.21 DPS) [dungeon]; Spine of the Devourer (273040, -1.21 DPS) [world]; Tendonscraper Dagger (273053, -1.47 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.56 DPS) | yes | Orb of Mystic Insight (249394, -0.53 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.10 DPS) [world_drop]; Orb of Souls (249395, -1.30 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, -0.26 DPS) [quest]; Wand of Eventide (5214, -0.28 DPS) [world_drop]; Greater Mystic Wand (217287, -0.28 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (orc, 10000000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 160.0. Weights run: 3.7s. Verify run: 5.7s. 448 eligible items had no known source.

Race sweep: undead 160.1 ± 0.1, orc 160.0 ± 0.1, troll 158.5 ± 0.1 (dps); undead beats the guide race orc by 0.1 (0.05%) ± 0.2, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.033 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.013 per %), hit=0.176 ± 0.008 per rating point (10 rating = 1%, 1.763 per %), spell_haste=0.965 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.821 ± 0.001, fire_power=0.179 ± 0.000, spirit=not significant (0.134 ± 0.107)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.44 DPS) | yes | Miner's Hat of the Deep (9429, -2.59 DPS) [dungeon]; Living Cowl (5608, -2.67 DPS, sim-verified) [world]; Holy Shroud (2721, -2.82 DPS) [world_drop] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Prodigious Shadowshard Pendant (17773, -0.28 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -0.51 DPS) [dungeon]; Amberglow Talisman (10824, -2.33 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.3 spell_power points (2.85 DPS) | yes | Green Silken Shoulders (7057, -0.29 DPS) [crafted]; Mantle of Doan (7712, -0.31 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.57 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.50 DPS) [crafted]; Guardian Cloak (5965, -0.50 DPS) [crafted]; Battle Healer's Cloak (19528, -0.54 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (7.01 DPS) | yes | Elemental Raiment (9434, -0.81 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -0.83 DPS) [crafted]; Shadoweave Robe (10004, -1.44 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.76 DPS) | yes | Condor Bracers (15864, -0.91 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.33 DPS) [crafted]; Radiant Silver Bracers (4545, -1.45 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (5.84 DPS) | yes | Shadoweave Gloves (10023, -1.52 DPS) [crafted]; Black Mageweave Gloves (10003, -1.69 DPS, sim-verified) [crafted]; Black Handwraps (254027, -2.25 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.1 spell_power points (4.33 DPS) | yes | Star Belt (4329, -0.18 DPS) [crafted]; Black Cord (254043, -0.48 DPS) [crafted]; Warsong Sash (16975, -0.96 DPS) [quest] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.2 spell_power points (5.28 DPS) | yes | Red Mageweave Pants (10009, -0.87 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.23 DPS) [dungeon]; Shadow Leggings (253993, -1.49 DPS) [crafted] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.4 spell_power points (3.50 DPS) | yes | Gilded Slippers (254001, -1.12 DPS) [crafted]; Spidersilk Boots (4320, -1.32 DPS) [crafted]; Glinteye Slippers (273024, -1.56 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 11.7 spell_power points (3.57 DPS) | yes | Advisor's Ring (19520, -0.81 DPS) [rep]; Reedknot Ring (9622, -1.43 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.2 spell_power points (3.11 DPS) | yes | Advisor's Ring (19520, -0.36 DPS) [rep]; Reedknot Ring (9622, -0.97 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -6.12 DPS) [dungeon]; Illusionary Rod (7713, -7.13 DPS, sim-verified) [dungeon]; Swinetusk Shank (6691, -7.39 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.21 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Shadoweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 25400000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 245.6. Weights run: 3.7s. Verify run: 5.0s. 587 eligible items had no known source.

Race sweep: undead 246.0 ± 0.2, orc 245.6 ± 0.1, troll 242.8 ± 0.2 (dps); undead beats the guide race orc by 0.4 (0.15%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.132 ± 0.015, crit=0.187 ± 0.004 per rating point (14 rating = 1%, 2.622 per %), hit=0.475 ± 0.019 per rating point (10 rating = 1%, 4.747 per %), spell_haste=2.641 ± 0.262, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.772 ± 0.001, fire_power=0.228 ± 0.000, spirit=not significant (0.250 ± 0.152)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.03 DPS) | yes | Dreamweave Circlet (10041, -1.51 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -1.99 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.23 DPS) [crafted] |
| neck | Jarkal's Enhancing Necklace (7888) | Necklace Recovery, Take 3 [quest] | 12.1 spell_power points (4.51 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.16 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.51 DPS) [dungeon]; Triune Amulet (7722, -1.84 DPS, sim-verified) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.2 spell_power points (6.39 DPS) | yes | Rotgrip Mantle (17732, -0.68 DPS) [dungeon]; Netherpearl Shoulders (254055, -0.88 DPS, sim-verified) [crafted]; Shadoweave Shoulders (10028, -1.52 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.29 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.32 DPS) [crafted]; Nightfall Drape (12465, -2.71 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (8.93 DPS) | yes | Dreamweave Vest (10021, -0.50 DPS) [crafted]; Elemental Raiment (9434, -1.13 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -1.57 DPS) [world] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 12.2 spell_power points (4.53 DPS) | yes | Nethergeld Cuffs (254061, -1.13 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -1.19 DPS) [dungeon]; Condor Bracers (15864, -1.93 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.3 spell_power points (7.53 DPS) | yes | Black Gloves (254101, -1.25 DPS, sim-verified) [crafted]; Runecloth Gloves (13863, -1.80 DPS) [crafted]; Black Mageweave Gloves (10003, -1.96 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.1 spell_power points (6.71 DPS) | yes | Black Waistcord (254087, -0.82 DPS) [crafted]; Satyrmane Sash (17755, -1.02 DPS) [dungeon]; Ghostweave Cord (254073, -1.90 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 24.3 spell_power points (9.03 DPS) | yes | Felcloth Pants (14107, -1.24 DPS, sim-verified) [crafted]; Wizardweave Leggings (14132, -1.98 DPS) [crafted]; Senior Designer's Pantaloons (11841, -2.56 DPS) [dungeon] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 16.6 spell_power points (6.17 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -1.09 DPS, sim-verified) [vendor]; Mistwalker Boots (10629, -1.16 DPS) [dungeon]; Shadoweave Boots (10031, -1.23 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 13.6 spell_power points (5.04 DPS) | yes | Advisor's Ring (19519, -0.58 DPS) [rep]; Underworld Band (1980, -0.78 DPS) [world_drop]; Cyclopean Band (11824, -0.98 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.83 DPS) | yes | Underworld Band (1980, -0.57 DPS) [world_drop]; Advisor's Ring (19519, -0.77 DPS, sim-verified) [rep]; Cyclopean Band (11824, -0.77 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.35 DPS) [quest]; Ankh of Life (1713, -0.47 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.66 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.99 DPS) [quest]; Ankh of Life (1713, -1.11 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.22 DPS, sim-verified) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -5.03 DPS, sim-verified) [world_drop] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (5.20 DPS) | yes | Umbral Crystal (13029, -0.33 DPS) [world_drop]; Orb of Dar'Orahil (15108, -0.94 DPS) [quest]; Mordresh's Lifeless Skull (10770, -1.19 DPS, sim-verified) [dungeon] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.37 DPS) [crafted]; Thunderwood (13062, -0.74 DPS) [world_drop]; Twisted Nether Wand (249144, -0.74 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Jarkal's Enhancing Necklace; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Black Sandals; finger1: Brainlash; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Orb of the Forgotten Seer; ranged: Lesser Eternal Wand

No-known-source sample (15 of 587, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522000000000000-2035113101220001350-0050000000000000)

Set DPS (verified): 414.9. Weights run: 3.5s. Verify run: 18.6s. 1454 eligible items had no known source.

Race sweep: orc 414.9 ± 0.2, undead 414.4 ± 0.2, troll 409.5 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.173 ± 0.028, crit=0.298 ± 0.010 per rating point (14 rating = 1%, 4.175 per %), hit=0.764 ± 0.035 per rating point (10 rating = 1%, 7.637 per %), spell_haste=3.296 ± 0.480, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.776 ± 0.001, fire_power=0.224 ± 0.000, spirit=not significant (0.235 ± 0.250)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 43.7 spell_power points (18.85 DPS) | yes | Warlord's Dreadweave Hood (231590, -3.25 DPS) [pvp]; Crimson Felt Hat (18727, -5.25 DPS, sim-verified) [dungeon]; Champion's Dreadweave Cowl (227090, -6.64 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 18.0 spell_power points (7.78 DPS) | yes | Dark Advisor's Pendant (18691, -0.48 DPS) [dungeon]; Lady Maye's Pendant (14558, -1.61 DPS) [world_drop]; Archlight Talisman (15856, -1.85 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 26.6 spell_power points (11.50 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Burial Shawl (18681, -1.67 DPS) [dungeon]; Mantle of the Timbermaw (19050, -1.73 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.84 DPS, sim-verified) [world]; Spritecaster Cape (11623, -3.70 DPS) [dungeon]; Shroud of the Exile (15421, -4.16 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (20.52 DPS) | yes | Warlord's Dreadweave Robe (231591, -4.92 DPS) [pvp]; Robe of Winter Night (14136, -7.20 DPS, sim-verified) [crafted]; Legionnaire's Dreadweave Tunic (227094, -8.24 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.0 spell_power points (10.80 DPS) | yes | Runecloth Cuffs (254123, -5.31 DPS) [crafted]; Netherpearl Cuffs (254067, -5.42 DPS) [crafted]; Wyrmthalak's Shackles (13958, -6.26 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.4 spell_power points (11.41 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -2.05 DPS) [pvp]; Earth Warder's Gloves (21318, -3.09 DPS, sim-verified) [quest] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 21.4 spell_power points (9.22 DPS) | yes | Frostwolf Cloth Belt (19090, -0.71 DPS) [rep]; Defiler's Cloth Girdle (20163, -0.93 DPS) [rep]; Valconian Sash (12083, -1.89 DPS) [quest] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 33.6 spell_power points (14.51 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -1.27 DPS) [crafted]; Skyshroud Leggings (13170, -1.58 DPS, sim-verified) [dungeon] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 23.9 spell_power points (10.30 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.78 DPS) [dungeon]; Dragonrider Boots (18102, -1.34 DPS) [dungeon] |
| finger1 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.75 DPS) [quest]; The Postmaster's Seal (13392, -1.99 DPS) [dungeon]; Naglering (11669, -11.26 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.43 DPS, sim-verified) [quest]; The Postmaster's Seal (13392, -1.85 DPS) [dungeon]; Don Mauricio's Band of Domination (22433, -2.35 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+8.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -4.47 DPS) [dungeon]; Serenity Field (272439, -4.75 DPS) [vendor]; Uther's Strength (11302, -5.18 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Eye of the Beast (13968, +0.00 DPS) [quest]; Uther's Strength (11302, -0.38 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20220, -1.61 DPS) [rep]; Headmaster's Charge (13937, -3.44 DPS) [dungeon]; Blade of Eternal Darkness (17780, -13.94 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (12.51 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.75 DPS) [pvp]; Tome of Shadow Force (19309, -2.22 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.45 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.37 DPS) [dungeon]; Oblivion's Touch (18761, -0.71 DPS) [dungeon]; Brilliant Wand (249385, -0.71 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; legs: Outrider's Silk Leggings; feet: Venomspew Footpads; finger1: Channeler's Ring; finger2: Maiden's Circle; trinket1: Weakness Analyzer; trinket2: Frozen Heart of the Mountain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 05000000000000000-2035103021220001351-0550005000000000)

Set DPS (verified): 805.2. Weights run: 4.0s. Verify run: 21.4s. 1454 eligible items had no known source.

Race sweep: orc 805.2 ± 0.8, troll 799.4 ± 0.8, undead 799.0 ± 0.8 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.272 ± 0.030, crit=0.517 ± 0.017 per rating point (14 rating = 1%, 7.234 per %), hit=0.655 ± 0.046 per rating point (10 rating = 1%, 6.552 per %), spell_haste=4.006 ± 0.584, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.883 ± 0.002, fire_power=0.117 ± 0.000, spirit=not significant (0.130 ± 0.153)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 41.2 spell_power points (21.82 DPS) | yes | Warlord's Dreadweave Hood (231590, -1.44 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -4.29 DPS) [pvp]; Crimson Felt Hat (18727, -12.39 DPS, sim-verified) [dungeon] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lady Maye's Pendant (14558, -0.83 DPS) [world_drop]; Archlight Talisman (15856, -1.59 DPS) [quest]; Dark Advisor's Pendant (18691, -2.83 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.8 spell_power points (14.69 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Burial Shawl (18681, -1.81 DPS) [dungeon]; Argent Shoulders (19059, -9.88 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -4.40 DPS) [dungeon]; Shroud of the Exile (15421, -4.81 DPS) [quest]; Crystalline Threaded Cape (20697, -12.08 DPS, sim-verified) [world] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.4 spell_power points (25.63 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.25 DPS) [pvp]; Robe of Winter Night (14136, -6.36 DPS, sim-verified) [crafted]; Felcloth Robe (14106, -8.82 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.1 spell_power points (13.27 DPS) | yes | Netherpearl Cuffs (254067, -5.85 DPS) [crafted]; Runecloth Cuffs (254123, -6.01 DPS) [crafted]; Manacle Cuffs (11962, -7.41 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -2.17 DPS) [pvp]; Hands of Power (13253, -7.54 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 22.9 spell_power points (12.09 DPS) | yes | Frostwolf Cloth Belt (19090, -1.14 DPS) [rep]; Black Waistcord (254087, -1.93 DPS) [crafted]; Ban'thok Sash (11662, -4.79 DPS, sim-verified) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 34.5 spell_power points (18.23 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -1.55 DPS) [pvp]; Skyshroud Leggings (13170, -6.46 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 23.3 spell_power points (12.30 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.48 DPS) [dungeon]; Venomspew Footpads (275606, -7.70 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.41 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.30 DPS) [dungeon]; Naglering (11669, -19.25 DPS, sim-verified) [dungeon] |
| finger2 | Channeler's Ring (272406) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.15 DPS) [dungeon]; The Postmaster's Seal (13392, -2.66 DPS) [dungeon]; Eye of Orgrimmar (12545, -7.96 DPS, sim-verified) [quest] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Second Wind (11819, -21.11 DPS, sim-verified) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, -5.27 DPS, sim-verified) [dungeon]; Serenity Field (272439, -5.82 DPS) [vendor]; Smolderweb's Eye (13213, -6.06 DPS) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.55 DPS) [pvp]; Whiteout Staff (19101, -14.91 DPS) [rep]; Headmaster's Charge (13937, -20.02 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.72 DPS) [dungeon]; Oblivion's Touch (18761, -0.77 DPS) [dungeon]; Brilliant Wand (249385, -0.77 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Channeler's Ring; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

