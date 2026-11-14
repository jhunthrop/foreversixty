# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 61.7. Weights run: 3.4s. Verify run: 4.4s. 199 eligible items had no known source.

Race sweep: human 61.7 ± 0.0, gnome 61.4 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.162 ± 0.005, crit=0.055 ± 0.001 per rating point (14 rating = 1%, 0.773 per %), hit=0.152 ± 0.006 per rating point (10 rating = 1%, 1.519 per %), spell_haste=0.432 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.832 ± 0.001, fire_power=0.168 ± 0.000, spirit=not significant (0.180 ± 0.077)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.2 spell_power points (2.77 DPS) | yes | Pristine Circlet (253949, -0.96 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -2.20 DPS) [crafted]; Flame Circlet (253953, -2.21 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.4 spell_power points (0.11 DPS) | yes | Tarnished Locket (279870, -0.32 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.5 spell_power points (1.95 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.74 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Black Whelp Cloak (7283, -0.52 DPS) [crafted]; Caretaker's Cape (20428, -0.62 DPS, sim-verified) [rep]; Pearl-clasped Cloak (5542, -0.68 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.5 spell_power points (3.77 DPS) | yes | Filigreed Shadow Gown (253907, -0.95 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.85 DPS) [crafted]; Scholarly Robes (2034, -2.21 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Repurposed Hair Band (281256, -0.04 DPS) [quest]; Bright Bracers (3647, -0.05 DPS) [world_drop]; Windsong Bangles (263336, -0.17 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.12 DPS) | yes | Shadow Gloves (253919, -0.24 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.30 DPS) [world]; Pristine Gloves (253913, -0.76 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.5 spell_power points (1.96 DPS) | yes | Pristine Sash (253925, -0.55 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.60 DPS) [crafted]; Novice Ardent's Sash (253887, -1.04 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.8 spell_power points (3.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.96 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.16 DPS) [dungeon]; Colorful Kilt (10048, -1.50 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (2.31 DPS) | yes | Shadow Boots (253895, -0.81 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.83 DPS) [crafted]; Pristine Boots (253889, -1.26 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.3 spell_power points (1.61 DPS) | yes | Band of Purification (12996, -1.28 DPS) [world_drop]; Furlfeather Ring (286745, -1.31 DPS) [world]; Lavishly Jeweled Ring (1156, -1.32 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.51 DPS) | yes | Furlfeather Ring (286745, -1.21 DPS) [world]; Lavishly Jeweled Ring (1156, -1.22 DPS) [dungeon]; Band of Purification (12996, -1.53 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.64 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.58 DPS) [dungeon]; Witching Stave (1484, -1.80 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.2 spell_power points (0.96 DPS) | yes | Antipodean Rod (2879, -0.66 DPS) [world_drop]; Onu's Patience (281292, -0.69 DPS) [quest]; Mathystra Channeling Orb (281291, -0.72 DPS) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.30 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (human, 25520000132001000-0000000000000000000-0000000000000000)

Set DPS (verified): 102.8. Weights run: 3.5s. Verify run: 5.4s. 326 eligible items had no known source.

Race sweep: human 102.8 ± 0.1, gnome 101.9 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.355 ± 0.007, crit=0.079 ± 0.001 per rating point (14 rating = 1%, 1.100 per %), hit=0.155 ± 0.007 per rating point (10 rating = 1%, 1.546 per %), spell_haste=not significant (0.139 ± 0.107), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.369 ± 0.114)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Filigreed Shadow Circlet (253981, +0.00 DPS) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.39 DPS) | yes | - |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Moonlit Amice (11884, -0.56 DPS) [quest]; Invoker's Mantle (215365, -0.56 DPS) [crafted]; Death Speaker Mantle (6685, -0.83 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Caretaker's Cape (19533, -0.28 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.62 DPS) | yes | Necrology Robes (2292, -0.43 DPS) [dungeon]; Shadow Gown (253967, -0.66 DPS) [crafted]; Green Silk Armor (7065, -1.11 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Phoenix Bindings (210781, -1.84 DPS) [crafted]; Windsong Bangles (263336, -2.23 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Darkness (7047, +0.00 DPS) [crafted]; Serpent Gloves (5970, -0.10 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.0 spell_power points (3.06 DPS) | yes | Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.83 DPS) [dungeon]; Invoker's Cord (215366, -1.11 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gaze Dreamer Pants (6903, +0.00 DPS) [dungeon]; Filigreed Shadow Leggings (253943, -0.68 DPS) [crafted]; Necromancer Leggings (2277, -0.91 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.95 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.01 DPS) [crafted]; Glinteye Slippers (273024, -1.01 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.95 DPS) | yes | Ring of Power Regulation (273030, -0.28 DPS) [dungeon]; Minor Channeling Ring (1449, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Minor Channeling Ring (1449, -0.28 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.11 DPS) [dungeon]; Spine of the Devourer (273040, -1.11 DPS) [world]; Tendonscraper Dagger (273053, -1.11 DPS) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.50 DPS) | yes | Orb of Mystic Insight (249394, -0.56 DPS) [crafted]; Orb of Mistmantle (13031, -1.14 DPS) [world_drop]; Eye of Paleth (2943, -1.39 DPS) [quest] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.27 DPS) [crafted]; Wand of Eventide (5214, -0.29 DPS) [world_drop]; Gravestone Scepter (7001, -0.29 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 25521000132001351-0000000000000000000-0000000000000000)

Set DPS (verified): 180.1. Weights run: 4.0s. Verify run: 4.4s. 455 eligible items had no known source.

Race sweep: human 180.1 ± 0.1, gnome 178.7 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.147 ± 0.005, crit=0.106 ± 0.001 per rating point (14 rating = 1%, 1.486 per %), hit=0.148 ± 0.007 per rating point (10 rating = 1%, 1.483 per %), spell_haste=not significant (0.180 ± 0.060), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.873 ± 0.001, fire_power=0.127 ± 0.000, spirit=not significant (0.233 ± 0.084)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.83 DPS) | yes | Miner's Hat of the Deep (9429, -2.37 DPS, sim-verified) [dungeon]; Living Cowl (5608, -2.98 DPS) [world]; Papal Fez (9431, -3.13 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.6 spell_power points (3.59 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.23 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.54 DPS) [dungeon]; Necklace of Calisea (1714, -2.60 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (3.85 DPS) | yes | Mantle of Doan (7712, -0.22 DPS) [dungeon]; Green Silken Shoulders (7057, -0.26 DPS) [crafted]; Inquisitor's Shawl (19507, -0.53 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19532, -0.46 DPS) [rep]; Long Silken Cloak (4326, -0.62 DPS) [crafted]; Guardian Cloak (5965, -0.62 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (8.97 DPS) | yes | Dreamweave Vest (10021, -0.64 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.14 DPS) [world_drop]; Shadoweave Robe (10004, -1.53 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Arcane Runed Bracers (4744, -0.33 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.75 DPS) [quest]; Earthen Silk Cuffs (254019, -1.43 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.2 spell_power points (7.54 DPS) | yes | Shadoweave Gloves (10023, -1.08 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.95 DPS) [crafted]; Black Handwraps (254027, -2.60 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.25 DPS) [crafted]; Black Cord (254043, -0.28 DPS, sim-verified) [crafted]; Belt of Arugal (6392, -1.49 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 18.3 spell_power points (6.84 DPS) | yes | Red Mageweave Pants (10009, -1.18 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -1.46 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.58 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 13.3 spell_power points (4.96 DPS) | yes | Glinteye Slippers (273024, -1.83 DPS) [dungeon]; Wingborne Boots (15104, -1.86 DPS) [quest]; Gilded Slippers (254001, -1.89 DPS, sim-verified) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.0 spell_power points (4.83 DPS) | yes | Lorekeeper's Ring (19524, -1.47 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.85 DPS) [quest]; Reedknot Ring (9622, -2.22 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.7 spell_power points (4.01 DPS) | yes | Lorekeeper's Ring (19524, -0.42 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -1.02 DPS) [quest]; Reedknot Ring (9622, -1.39 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -7.75 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -8.07 DPS) [dungeon]; Swinetusk Shank (6691, -9.91 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Shadoweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 455, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 25522020135001351-0000000000000000000-0040000000000000)

Set DPS (verified): 273.9. Weights run: 4.1s. Verify run: 7.2s. 594 eligible items had no known source.

Race sweep: human 273.9 ± 0.1, gnome 271.9 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.196 ± 0.008, crit=0.160 ± 0.002 per rating point (14 rating = 1%, 2.242 per %), hit=0.212 ± 0.010 per rating point (10 rating = 1%, 2.120 per %), spell_haste=0.345 ± 0.071, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.888 ± 0.001, fire_power=0.112 ± 0.000, spirit=not significant (0.042 ± 0.096)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.63 DPS) | yes | Dreamweave Circlet (10041, -1.02 DPS, sim-verified) [crafted]; Shadoweave Mask (10025, -1.61 DPS) [crafted]; Red Mageweave Headband (10033, -1.76 DPS) [crafted] |
| neck | Talvash's Enhancing Necklace (7673) | Restoring the Necklace [quest] | 12.2 spell_power points (5.25 DPS) | yes | Glowing Eye of Mordresh (10769, -1.56 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.62 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.68 DPS) [dungeon] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | 16.9 spell_power points (7.29 DPS) | yes | Kentic Amice (11624, -0.05 DPS) [dungeon]; Rotgrip Mantle (17732, -0.17 DPS) [dungeon]; Shadoweave Shoulders (10028, -0.78 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.49 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.77 DPS) [dungeon]; Caretaker's Cape (19531, -3.52 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.4 spell_power points (10.07 DPS) | yes | Dreamweave Vest (10021, -1.31 DPS) [crafted]; Elemental Raiment (9434, -1.37 DPS, sim-verified) [world_drop]; Runecloth Tunic (13857, -1.82 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.1 spell_power points (5.65 DPS) | yes | Spidertank Oilrag (9448, -1.78 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.96 DPS) [crafted]; Arcane Runed Bracers (4744, -2.68 DPS, sim-verified) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 19.2 spell_power points (8.26 DPS) | yes | Dreamweave Gloves (10019, -0.04 DPS) [crafted]; Shadoweave Gloves (10023, -1.33 DPS) [crafted]; Black Mageweave Gloves (10003, -1.79 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.6 spell_power points (8.03 DPS) | yes | Ban'thok Sash (11662, -1.10 DPS) [dungeon]; Satyrmane Sash (17755, -1.15 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.66 DPS) [rep] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.4 spell_power points (10.96 DPS) | yes | Spellshock Leggings (9484, -0.28 DPS, sim-verified) [dungeon]; Wizardweave Leggings (14132, -2.77 DPS) [crafted]; Shadoweave Pants (10002, -2.92 DPS) [crafted] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 19.5 spell_power points (8.41 DPS) | yes | Black Slippers (254007, -2.72 DPS, sim-verified) [crafted]; Shadoweave Boots (10031, -2.87 DPS) [crafted]; Gilded Sandals (254107, -2.91 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.4 spell_power points (5.78 DPS) | yes | Band of the Unicorn (7553, -0.18 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.61 DPS) [rep]; Philanthropist's Ring (281635, -1.05 DPS) [quest] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 13.2 spell_power points (5.67 DPS) | yes | Lorekeeper's Ring (19523, -0.50 DPS) [rep]; Band of the Unicorn (7553, -0.75 DPS, sim-verified) [world_drop]; Philanthropist's Ring (281635, -0.94 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.60 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.64 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -3.93 DPS, sim-verified) [world_drop] |
| off_hand | Umbral Crystal (13029) | World drop [world_drop] | 15.1 spell_power points (6.50 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.29 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.47 DPS) [dungeon]; Orb of Dar'Orahil (15108, -0.73 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.43 DPS) [crafted]; Burning Sliver (5249, -0.86 DPS) [quest]; Twisted Nether Wand (249144, -0.86 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Talvash's Enhancing Necklace; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; feet: Black Sandals; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Umbral Crystal; ranged: Lesser Eternal Wand

No-known-source sample (15 of 594, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522020135201351-0002000000000000000-0050050000000000)

Set DPS (verified): 432.2. Weights run: 4.1s. Verify run: 20.7s. 1460 eligible items had no known source.

Race sweep: human 432.2 ± 0.2, gnome 429.4 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.245 ± 0.013, crit=0.221 ± 0.003 per rating point (14 rating = 1%, 3.088 per %), hit=0.294 ± 0.017 per rating point (10 rating = 1%, 2.940 per %), spell_haste=2.568 ± 0.125, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.882 ± 0.001, fire_power=0.118 ± 0.000, spirit=not significant (0.383 ± 0.165)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 40.9 spell_power points (18.58 DPS) | yes | Field Marshal's Coronal (231584, -1.38 DPS) [pvp]; Crimson Felt Hat (18727, -4.66 DPS, sim-verified) [dungeon]; Felcloth Hood (14111, -5.45 DPS) [crafted] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 19.6 spell_power points (8.90 DPS) | yes | Archlight Talisman (15856, -2.34 DPS) [quest]; Jeweled Amulet of Cainwyn (1443, -2.43 DPS) [world_drop]; Lady Maye's Pendant (14558, -3.51 DPS, sim-verified) [world_drop] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 27.7 spell_power points (12.57 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Burial Shawl (18681, -1.71 DPS) [dungeon]; Felcloth Shoulders (14112, -1.83 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystalline Threaded Cape (20697, -0.63 DPS, sim-verified) [world]; Spritecaster Cape (11623, -1.42 DPS) [dungeon]; Shroud of the Exile (15421, -1.71 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (21.89 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.69 DPS) [pvp]; Robe of Winter Night (14136, -5.46 DPS, sim-verified) [crafted]; Felcloth Robe (14106, -7.47 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.6 spell_power points (12.10 DPS) | yes | Netherpearl Cuffs (254067, -5.24 DPS) [crafted]; Runecloth Cuffs (254123, -5.99 DPS) [crafted]; Wyrmthalak's Shackles (13958, -6.31 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 28.3 spell_power points (12.84 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Black Gloves (254101, -2.81 DPS, sim-verified) [crafted]; Earth Warder's Gloves (21318, -2.85 DPS) [quest] |
| waist | Stormpike Cloth Girdle (19094) | Stormpike Guard [rep] | 20.5 spell_power points (9.29 DPS) | yes | Black Waistcord (254087, -0.67 DPS) [crafted]; Highlander's Cloth Girdle (20047, -0.86 DPS) [rep]; Thuzadin Sash (18740, -3.49 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.6 spell_power points (18.45 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -3.16 DPS) [crafted]; Skyshroud Leggings (13170, -3.49 DPS) [dungeon] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 25.1 spell_power points (11.38 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -1.08 DPS) [pvp]; Omnicast Boots (11822, -1.57 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -1.67 DPS) [dungeon]; Blessed Band of Light (272407, -2.73 DPS) [vendor]; Naglering (11669, -11.53 DPS, sim-verified) [dungeon] |
| finger2 | Songstone of Ironforge (12543) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -1.68 DPS) [vendor]; Brainlash (6440, -1.97 DPS) [dungeon]; The Postmaster's Seal (13392, -2.65 DPS, sim-verified) [dungeon] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.39 DPS, sim-verified) [vendor]; Uther's Strength (11302, -5.45 DPS) [world_drop]; Second Wind (11819, -5.90 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, -0.34 DPS) [dungeon]; Grand Marshal's Stave (234571, -2.54 DPS) [pvp]; Ironbark Staff (20069, -14.13 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.17 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.37 DPS) [pvp]; Tome of Shadow Force (19309, -0.75 DPS) [rep]; Book of the Dead (13353, -3.35 DPS) [dungeon] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.56 DPS) [dungeon]; Brilliant Wand (249385, -0.69 DPS) [crafted]; Oblivion's Touch (18761, -1.06 DPS, sim-verified) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Stormpike Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Venomspew Footpads; finger1: Maiden's Circle; finger2: Songstone of Ironforge; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 05520020135201351-0050203001000000000-0500000000000000)

Set DPS (verified): 772.0. Weights run: 3.7s. Verify run: 14.9s. 1460 eligible items had no known source.

Race sweep: human 772.0 ± 0.5, gnome 770.8 ± 0.5 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.351 ± 0.016, crit=0.700 ± 0.012 per rating point (14 rating = 1%, 9.804 per %), hit=0.757 ± 0.036 per rating point (10 rating = 1%, 7.570 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, fire_power=not significant (0.000 ± 0.000), spirit=not significant (0.093 ± 0.035)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 41.8 spell_power points (27.37 DPS) | yes | Field Marshal's Coronal (231584, -0.89 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -3.05 DPS) [pvp]; Crimson Felt Hat (18727, -4.48 DPS, sim-verified) [dungeon] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 22.8 spell_power points (14.94 DPS) | yes | Nacreous Shell Necklace (22403, -2.25 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.37 DPS) [world_drop]; Emberfury Talisman (12929, -4.59 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.4 spell_power points (20.55 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.26 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.28 DPS) [pvp]; Felcloth Shoulders (14112, -3.33 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.26 DPS) [world]; Spritecaster Cape (11623, -6.37 DPS) [dungeon]; Shroud of the Exile (15421, -6.70 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.2 spell_power points (32.21 DPS) | yes | Robe of Winter Night (14136, -2.37 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -5.73 DPS) [pvp]; Felcloth Robe (14106, -8.62 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.5 spell_power points (16.68 DPS) | yes | Netherpearl Cuffs (254067, -6.25 DPS) [crafted]; Runecloth Cuffs (254123, -7.17 DPS) [crafted]; Manacle Cuffs (11962, -8.65 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 32.1 spell_power points (21.05 DPS) | yes | Marshal's Dreadweave Gloves (231586, -0.01 DPS) [pvp]; Hands of Power (13253, -2.93 DPS, sim-verified) [dungeon]; Black Gloves (254101, -5.81 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 25.9 spell_power points (16.98 DPS) | yes | Ban'thok Sash (11662, -1.86 DPS) [dungeon]; Black Waistcord (254087, -2.46 DPS) [crafted]; Stormpike Cloth Girdle (19094, -2.88 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 55.1 spell_power points (36.12 DPS) | yes | Marshal's Dreadweave Leggings (231587, -7.50 DPS) [pvp]; Skyshroud Leggings (13170, -13.70 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -14.78 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.2 spell_power points (15.86 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.39 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.36 DPS) [quest]; Blessed Band of Light (272407, -1.77 DPS) [vendor]; Naglering (11669, -13.42 DPS, sim-verified) [dungeon] |
| finger2 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.23 DPS) [quest]; Blessed Band of Light (272407, -1.63 DPS) [vendor]; Don Julio's Band (19325, -2.25 DPS) [rep] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (+11.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -6.83 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -7.33 DPS) [crafted]; Serenity Field (272439, -7.63 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.75 DPS) [pvp]; Headmaster's Charge (13937, -13.81 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -22.00 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.85 DPS) [dungeon]; Brilliant Wand (249385, -0.85 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -1.15 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Don Mauricio's Band of Domination; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1460, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.6. Weights run: 3.4s. Verify run: 3.6s. 193 eligible items had no known source.

Race sweep: undead 61.3 ± 0.0, orc 60.6 ± 0.0, troll 59.8 ± 0.0 (dps); undead beats the guide race orc by 0.7 (1.21%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.162 ± 0.005, crit=0.055 ± 0.001 per rating point (14 rating = 1%, 0.773 per %), hit=0.152 ± 0.006 per rating point (10 rating = 1%, 1.519 per %), spell_haste=0.432 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.832 ± 0.001, fire_power=0.168 ± 0.000, spirit=not significant (0.180 ± 0.077)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.2 spell_power points (2.77 DPS) | yes | Pristine Circlet (253949, -0.96 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -2.20 DPS) [crafted]; Flame Circlet (253953, -2.21 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.5 spell_power points (0.16 DPS) | yes | Scholarly Pendant (277203, -0.05 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.5 spell_power points (1.95 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.74 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (20427, -0.40 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.52 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.68 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.5 spell_power points (3.77 DPS) | yes | Filigreed Shadow Gown (253907, -1.05 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.85 DPS) [crafted]; Scholarly Robes (2034, -2.21 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.91 DPS) | yes | Owlbeard Bracers (16981, -0.57 DPS, sim-verified) [quest]; Windsong Bangles (263336, -0.61 DPS) [quest]; Mindthrust Bracers (1974, -0.66 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.12 DPS) | yes | Shadow Gloves (253919, -0.23 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.30 DPS) [world]; Blight Gloves (279877, -0.61 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.5 spell_power points (1.96 DPS) | yes | Pristine Sash (253925, -0.56 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.60 DPS) [crafted]; Novice Ardent's Sash (253887, -1.04 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.8 spell_power points (3.28 DPS) | yes | Filigreed Pristine Leggings (253937, -0.96 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.16 DPS) [dungeon]; Colorful Kilt (10048, -1.50 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (2.31 DPS) | yes | Shadow Boots (253895, -0.58 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.83 DPS) [crafted]; Pristine Boots (253889, -1.26 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.51 DPS) | yes | Furlfeather Ring (286745, -1.21 DPS) [world]; Lavishly Jeweled Ring (1156, -1.22 DPS) [dungeon]; Loop of Sacrifice (281673, -1.27 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 1.1 spell_power points (0.33 DPS) | yes | Furlfeather Ring (286745, -0.02 DPS) [world]; Lavishly Jeweled Ring (1156, -0.03 DPS) [dungeon]; Loop of Sacrifice (281673, -0.08 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 29.8 spell_power points (9.02 DPS) | yes | Evocator's Blade (2567, -0.17 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.61 DPS) [quest]; Crescent Staff (6505, -0.81 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.30 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Staff of the Blessed Seer; ranged: Sizzle Stick

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade

### Band 30 (orc, 25520000132001000-0000000000000000000-0000000000000000)

Set DPS (verified): 103.3. Weights run: 3.5s. Verify run: 6.4s. 319 eligible items had no known source.

Race sweep: undead 103.6 ± 0.1, orc 103.3 ± 0.1, troll 101.9 ± 0.1 (dps); undead beats the guide race orc by 0.3 (0.26%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.355 ± 0.007, crit=0.079 ± 0.001 per rating point (14 rating = 1%, 1.100 per %), hit=0.155 ± 0.007 per rating point (10 rating = 1%, 1.546 per %), spell_haste=not significant (0.139 ± 0.107), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.369 ± 0.114)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Circlet (253981, -0.20 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.39 DPS) | yes | Roadwatcher's Confidence (281265, -1.22 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Death Speaker Mantle (6685, -0.83 DPS) [dungeon]; Invoker's Mantle (215365, -0.91 DPS, sim-verified) [crafted]; Chestnut Mantle (17695, -1.11 DPS) [quest] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Battle Healer's Cloak (19529, -0.28 DPS) [rep]; Hillman's Cloak (3719, -0.32 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.62 DPS) | yes | Necrology Robes (2292, -0.43 DPS) [dungeon]; Shadow Gown (253967, -0.66 DPS) [crafted]; Green Silk Armor (7065, -1.11 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Phoenix Bindings (210781, -1.84 DPS) [crafted]; Windsong Bangles (263336, -2.23 DPS) [quest]; Tabitha's Cuffs (251486, -2.51 DPS, sim-verified) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent Gloves (5970, -0.10 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world]; Hands of Darkness (7047, -0.46 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warsong Sash (16975, -0.21 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.83 DPS) [dungeon] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gaze Dreamer Pants (6903, -0.50 DPS, sim-verified) [dungeon]; Filigreed Shadow Leggings (253943, -0.68 DPS) [crafted]; Necromancer Leggings (2277, -0.91 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.95 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.01 DPS) [crafted]; Glinteye Slippers (273024, -1.98 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.95 DPS) | yes | Solid Gold Band (279845, -0.83 DPS) [quest]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon]; Ring of Power Regulation (273030, -2.44 DPS, sim-verified) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of Power Regulation (273030, -0.30 DPS, sim-verified) [dungeon]; Solid Gold Band (279845, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.11 DPS) [dungeon]; Spine of the Devourer (273040, -1.11 DPS) [world]; Tendonscraper Dagger (273053, -1.85 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.50 DPS) | yes | Orb of Mystic Insight (249394, -0.56 DPS) [crafted]; Orb of Mistmantle (13031, -1.14 DPS) [world_drop]; Orb of Souls (249395, -1.39 DPS) [crafted] |
| ranged | Gravestone Scepter (7001) | Blackfathom Villainy [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Wand of Eventide (5214, -0.08 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Gravestone Scepter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (orc, 25521000132001351-0000000000000000000-0000000000000000)

Set DPS (verified): 180.6. Weights run: 4.0s. Verify run: 3.8s. 448 eligible items had no known source.

Race sweep: orc 180.6 ± 0.1, undead 180.5 ± 0.1, troll 178.1 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.151 ± 0.009, crit=0.243 ± 0.003 per rating point (14 rating = 1%, 3.402 per %), hit=0.302 ± 0.014 per rating point (10 rating = 1%, 3.022 per %), spell_haste=0.744 ± 0.114, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.886 ± 0.001, fire_power=0.114 ± 0.000, spirit=not significant (0.164 ± 0.077)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.74 DPS) | yes | Miner's Hat of the Deep (9429, -3.32 DPS, sim-verified) [dungeon]; Living Cowl (5608, -3.33 DPS) [world]; Augural Shroud (2620, -3.53 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.1 spell_power points (3.81 DPS) | yes | Glowing Eye of Mordresh (10769, -0.39 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.47 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -2.89 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (4.31 DPS) | yes | Mantle of Doan (7712, -0.48 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.58 DPS) [dungeon]; Green Silken Shoulders (7057, -0.61 DPS, sim-verified) [crafted] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.51 DPS) [crafted]; Guardian Cloak (5965, -0.51 DPS) [crafted]; Battle Healer's Cloak (19528, -1.02 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.7 spell_power points (9.87 DPS) | yes | Elemental Raiment (9434, -1.13 DPS) [world_drop]; Dreamweave Vest (10021, -1.16 DPS, sim-verified) [crafted]; Shadoweave Robe (10004, -1.65 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.74 DPS) | yes | Radiant Silver Bracers (4545, -1.58 DPS) [quest]; Earthen Silk Cuffs (254019, -1.74 DPS) [crafted]; Condor Bracers (15864, -1.75 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.7 spell_power points (8.22 DPS) | yes | Black Mageweave Gloves (10003, -1.98 DPS) [crafted]; Shadoweave Gloves (10023, -2.41 DPS, sim-verified) [crafted]; Black Handwraps (254027, -2.62 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.6 spell_power points (6.08 DPS) | yes | Black Cord (254043, -0.00 DPS) [crafted]; Star Belt (4329, -0.39 DPS) [crafted]; Warsong Sash (16975, -1.50 DPS) [quest] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 18.6 spell_power points (7.74 DPS) | yes | Red Mageweave Pants (10009, -1.52 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -1.80 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.14 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 13.2 spell_power points (5.50 DPS) | yes | Gilded Slippers (254001, -1.89 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -2.18 DPS) [dungeon]; Spidersilk Boots (4320, -2.34 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.2 spell_power points (5.48 DPS) | yes | Advisor's Ring (19520, -1.73 DPS) [rep]; Reedknot Ring (9622, -2.56 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (4.47 DPS) | yes | Advisor's Ring (19520, -0.89 DPS, sim-verified) [rep]; Reedknot Ring (9622, -1.56 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.16 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -9.77 DPS) [dungeon]; Swinetusk Shank (6691, -11.87 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.08 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Shadoweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 25522020135001351-0000000000000000000-0040000000000000)

Set DPS (verified): 274.3. Weights run: 4.1s. Verify run: 7.5s. 587 eligible items had no known source.

Race sweep: undead 274.8 ± 0.2, orc 274.3 ± 0.1, troll 271.4 ± 0.1 (dps); undead beats the guide race orc by 0.5 (0.18%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.182 ± 0.012, crit=0.314 ± 0.004 per rating point (14 rating = 1%, 4.390 per %), hit=0.408 ± 0.019 per rating point (10 rating = 1%, 4.083 per %), spell_haste=0.472 ± 0.084, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.892 ± 0.001, fire_power=0.108 ± 0.000, spirit=not significant (0.289 ± 0.088)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.35 DPS) | yes | Shadoweave Mask (10025, -1.72 DPS) [crafted]; Dreamweave Circlet (10041, -1.92 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -1.99 DPS) [crafted] |
| neck | Jarkal's Enhancing Necklace (7888) | Necklace Recovery, Take 3 [quest] | 12.2 spell_power points (5.57 DPS) | yes | Triune Amulet (7722, -0.99 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.46 DPS) [quest]; Mindburst Medallion (11196, -1.67 DPS) [quest] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | 18.4 spell_power points (8.40 DPS) | yes | Kentic Amice (11624, -0.12 DPS) [dungeon]; Rotgrip Mantle (17732, -0.95 DPS) [dungeon]; Shadoweave Shoulders (10028, -1.46 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.02 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.91 DPS) [crafted]; Nightfall Drape (12465, -3.58 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.5 spell_power points (11.23 DPS) | yes | Dreamweave Vest (10021, -0.87 DPS, sim-verified) [crafted]; Embrace of the Wind Serpent (12462, -1.39 DPS) [world]; Elemental Raiment (9434, -1.62 DPS) [world_drop] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 14.3 spell_power points (6.55 DPS) | yes | Nethergeld Cuffs (254061, -2.08 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -2.44 DPS) [dungeon]; Bloodband Bracers (11469, -3.12 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 21.1 spell_power points (9.65 DPS) | yes | Dreamweave Gloves (10019, -0.43 DPS, sim-verified) [crafted]; Runecloth Gloves (13863, -2.22 DPS) [crafted]; Shadoweave Gloves (10023, -2.29 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.6 spell_power points (8.51 DPS) | yes | Ban'thok Sash (11662, -0.32 DPS) [dungeon]; Ghostweave Cord (254073, -0.65 DPS) [crafted]; Satyrmane Sash (17755, -1.27 DPS) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.4 spell_power points (11.62 DPS) | yes | Spellshock Leggings (9484, -0.27 DPS, sim-verified) [dungeon]; Wizardweave Leggings (14132, -2.92 DPS) [crafted]; Shadoweave Pants (10002, -3.04 DPS) [crafted] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 19.5 spell_power points (8.92 DPS) | yes | Shadoweave Boots (10031, -1.67 DPS, sim-verified) [crafted]; Mistwalker Boots (10629, -2.42 DPS) [dungeon]; Black Slippers (254007, -2.50 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 14.3 spell_power points (6.53 DPS) | yes | Band of the Unicorn (7553, -0.58 DPS) [world_drop]; Advisor's Ring (19519, -1.04 DPS) [rep]; Cyclopean Band (11824, -1.30 DPS) [dungeon] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 13.4 spell_power points (6.13 DPS) | yes | Band of the Unicorn (7553, -0.19 DPS) [world_drop]; Advisor's Ring (19519, -0.64 DPS) [rep]; Cyclopean Band (11824, -0.90 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -0.10 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.36 DPS) [quest]; Rune of the Guard Captain (19120, -0.37 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.16 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.44 DPS) [quest]; Evonice's Landin' Pilla (18951, -1.68 DPS, sim-verified) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, -4.06 DPS, sim-verified) [dungeon] |
| off_hand | Mordresh's Lifeless Skull (10770) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (7.19 DPS) | yes | Umbral Crystal (13029, -0.25 DPS) [world_drop]; Orb of the Forgotten Seer (7685, -0.79 DPS) [dungeon]; Orb of Dar'Orahil (15108, -1.06 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.46 DPS) [crafted]; Gravestone Scepter (7001, -0.81 DPS) [quest]; Twisted Nether Wand (249144, -0.92 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Jarkal's Enhancing Necklace; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Underworld Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Mordresh's Lifeless Skull; ranged: Lesser Eternal Wand

No-known-source sample (15 of 587, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522020135201351-0002000000000000000-0050050000000000)

Set DPS (verified): 427.1. Weights run: 4.1s. Verify run: 19.0s. 1454 eligible items had no known source.

Race sweep: undead 427.7 ± 0.2, orc 427.1 ± 0.2, troll 421.8 ± 0.2 (dps); undead beats the guide race orc by 0.7 (0.16%) ± 0.3, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.216 ± 0.018, crit=0.469 ± 0.006 per rating point (14 rating = 1%, 6.571 per %), hit=0.577 ± 0.028 per rating point (10 rating = 1%, 5.773 per %), spell_haste=5.726 ± 0.253, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.878 ± 0.001, fire_power=0.122 ± 0.000, spirit=not significant (-0.033 ± 0.132)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 38.8 spell_power points (18.65 DPS) | yes | Warlord's Dreadweave Hood (231590, -0.76 DPS) [pvp]; Crimson Felt Hat (18727, -1.86 DPS, sim-verified) [dungeon]; Champion's Dreadweave Cowl (227090, -3.52 DPS) [pvp] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 19.3 spell_power points (9.28 DPS) | yes | Nacreous Shell Necklace (22403, -1.34 DPS) [dungeon]; Lady Maye's Pendant (14558, -2.01 DPS) [world_drop]; Archlight Talisman (15856, -2.47 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.66 DPS, sim-verified) [crafted]; Burial Shawl (18681, -0.74 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -1.27 DPS) [world]; Spritecaster Cape (11623, -3.95 DPS) [dungeon]; Shroud of the Exile (15421, -4.49 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.9 spell_power points (23.07 DPS) | yes | Robe of Winter Night (14136, -4.07 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -5.17 DPS) [pvp]; Felcloth Robe (14106, -7.86 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.7 spell_power points (11.42 DPS) | yes | Runecloth Cuffs (254123, -5.08 DPS) [crafted]; Netherpearl Cuffs (254067, -5.20 DPS) [crafted]; Manacle Cuffs (11962, -6.49 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 24.9 spell_power points (12.00 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Gloves of Spell Mastery (14146, -0.31 DPS) [crafted]; Earth Warder's Gloves (21318, -1.42 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 21.9 spell_power points (10.52 DPS) | yes | Frostwolf Cloth Belt (19090, -0.82 DPS) [rep]; Ban'thok Sash (11662, -0.93 DPS) [dungeon]; Black Waistcord (254087, -1.56 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.49 DPS, sim-verified) [dungeon]; Legionnaire's Dreadweave Legguards (227097, -0.62 DPS) [pvp] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.26 DPS) [dungeon]; Omnicast Boots (11822, -0.97 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.94 DPS) [dungeon]; Blessed Band of Light (272407, -1.80 DPS) [vendor]; Naglering (11669, -9.39 DPS, sim-verified) [dungeon] |
| finger2 | Eye of Orgrimmar (12545) | The Princess Saved? [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.94 DPS) [dungeon]; Blessed Band of Light (272407, -1.80 DPS) [vendor]; Deep Rooted Ring (19109, -2.57 DPS) [world] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+5.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, -3.83 DPS, sim-verified) [dungeon]; Serenity Field (272439, -5.29 DPS) [vendor]; Uther's Strength (11302, -5.77 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -0.94 DPS) [pvp]; Mindfang (20214, -3.13 DPS) [rep]; Ironbark Staff (20220, -12.17 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (13.95 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.57 DPS) [pvp]; Tome of Shadow Force (19309, -0.85 DPS) [rep]; Spirit of Aquementas (11904, -3.85 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.52 DPS) [dungeon]; Oblivion's Touch (18761, -0.75 DPS) [dungeon]; Brilliant Wand (249385, -0.75 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Venomspew Footpads; finger1: Maiden's Circle; finger2: Eye of Orgrimmar; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 05520020135201351-0050203001000000000-0500000000000000)

Set DPS (verified): 759.0. Weights run: 3.7s. Verify run: 16.4s. 1454 eligible items had no known source.

Race sweep: undead 760.4 ± 0.5, orc 759.0 ± 0.5, troll 750.9 ± 0.5 (dps); undead beats the guide race orc by 1.4 (0.19%) ± 0.7, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.351 ± 0.016, crit=0.700 ± 0.012 per rating point (14 rating = 1%, 9.804 per %), hit=0.757 ± 0.036 per rating point (10 rating = 1%, 7.570 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, fire_power=not significant (0.000 ± 0.000), spirit=not significant (0.093 ± 0.035)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 41.8 spell_power points (27.37 DPS) | yes | Warlord's Dreadweave Hood (231590, -0.89 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -3.05 DPS) [pvp]; Crimson Felt Hat (18727, -4.80 DPS, sim-verified) [dungeon] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 22.8 spell_power points (14.94 DPS) | yes | Nacreous Shell Necklace (22403, -2.25 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.37 DPS) [world_drop]; Emberfury Talisman (12929, -4.59 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.4 spell_power points (20.55 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.26 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -3.28 DPS) [pvp]; Felcloth Shoulders (14112, -3.51 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.88 DPS, sim-verified) [world]; Spritecaster Cape (11623, -6.37 DPS) [dungeon]; Shroud of the Exile (15421, -6.70 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.2 spell_power points (32.21 DPS) | yes | Robe of Winter Night (14136, -3.24 DPS) [crafted]; Warlord's Dreadweave Robe (231591, -5.73 DPS) [pvp]; Felcloth Robe (14106, -8.62 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.5 spell_power points (16.68 DPS) | yes | Netherpearl Cuffs (254067, -6.25 DPS) [crafted]; Runecloth Cuffs (254123, -7.17 DPS) [crafted]; Manacle Cuffs (11962, -8.65 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 32.1 spell_power points (21.05 DPS) | yes | General's Dreadweave Gloves (231589, -0.01 DPS) [pvp]; Hands of Power (13253, -3.03 DPS, sim-verified) [dungeon]; Black Gloves (254101, -5.81 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 25.9 spell_power points (16.98 DPS) | yes | Ban'thok Sash (11662, -1.86 DPS) [dungeon]; Black Waistcord (254087, -2.46 DPS) [crafted]; Frostwolf Cloth Belt (19090, -2.88 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 35.6 spell_power points (23.32 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.91 DPS) [dungeon]; Legionnaire's Dreadweave Legguards (227097, -1.99 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.2 spell_power points (15.86 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.39 DPS) [dungeon]; Black Sandals (254113, -0.69 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.36 DPS) [quest]; Blessed Band of Light (272407, -1.77 DPS) [vendor]; Naglering (11669, -14.90 DPS, sim-verified) [dungeon] |
| finger2 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.23 DPS) [quest]; Blessed Band of Light (272407, -1.63 DPS) [vendor]; Don Julio's Band (19325, -2.25 DPS) [rep] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+12.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -5.34 DPS, sim-verified) [crafted]; Smolderweb's Eye (13213, -6.83 DPS) [dungeon]; Serenity Field (272439, -7.21 DPS) [vendor] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.75 DPS) [pvp]; Headmaster's Charge (13937, -14.66 DPS, sim-verified) [dungeon]; Whiteout Staff (19101, -21.79 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.85 DPS) [dungeon]; Brilliant Wand (249385, -0.85 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -1.15 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Don Mauricio's Band of Domination; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1454, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

