# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000000-0050230100000000)

Set DPS (verified): 62.9. Weights run: 3.9s. Verify run: 4.8s. 199 eligible items had no known source.

Race sweep: gnome 62.9 ± 0.0, human 62.9 ± 0.0 (dps); gnome beats the guide race human by 0.0 (0.01%) ± 0.1, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.106 ± 0.003, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.365 per %), hit=0.061 ± 0.003 per rating point (10 rating = 1%, 0.605 per %), spell_haste=0.259 ± 0.026, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.800 ± 0.001, fire_power=0.200 ± 0.000, spirit=0.142 ± 0.035

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.8 spell_power points (2.60 DPS) | yes | Pristine Circlet (253949, -0.97 DPS, sim-verified) [crafted]; Flame Circlet (253953, -1.95 DPS) [crafted]; Shadow Goggles (4373, -2.19 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.08 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.0 spell_power points (1.76 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.58 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.45 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.46 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.67 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 11.7 spell_power points (3.47 DPS) | yes | Filigreed Shadow Gown (253907, -1.05 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.71 DPS) [crafted]; Scholarly Robes (2034, -2.08 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Repurposed Hair Band (281256, -0.01 DPS) [quest]; Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.18 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.07 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Shadow Gloves (253919, -0.32 DPS) [crafted]; Pristine Gloves (253913, -0.79 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.0 spell_power points (1.78 DPS) | yes | Novice Arcanist's Sash (253885, -0.50 DPS) [crafted]; Pristine Sash (253925, -0.56 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.97 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.0 spell_power points (2.96 DPS) | yes | Silk-threaded Trousers (1929, -0.89 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.97 DPS, sim-verified) [crafted]; Colorful Kilt (10048, -1.27 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.4 spell_power points (2.19 DPS) | yes | Shadow Boots (253895, -0.55 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.80 DPS) [crafted]; Pristine Boots (253889, -1.21 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.2 spell_power points (1.54 DPS) | yes | Furlfeather Ring (286745, -1.24 DPS) [world]; Band of Purification (12996, -1.29 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.35 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.48 DPS) | yes | Furlfeather Ring (286745, -1.22 DPS, sim-verified) [world]; Band of Purification (12996, -1.23 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.29 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.94 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.46 DPS) [dungeon]; Witching Stave (1484, -2.01 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.1 spell_power points (0.93 DPS) | yes | Antipodean Rod (2879, -0.57 DPS) [world_drop]; Onu's Patience (281292, -0.72 DPS) [quest]; Grayson's Torch (1172, -0.76 DPS) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.30 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (human, 00000000000000000-0000000000000000000-1050235103001000)

Set DPS (verified): 105.5. Weights run: 3.7s. Verify run: 4.2s. 326 eligible items had no known source.

Race sweep: human 105.5 ± 0.1, gnome 105.1 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.082 ± 0.005, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.988 per %), hit=0.110 ± 0.005 per rating point (10 rating = 1%, 1.100 per %), spell_haste=0.250 ± 0.062, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.794 ± 0.001, fire_power=0.206 ± 0.000, spirit=not significant (0.053 ± 0.052)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.3 spell_power points (3.64 DPS) | yes | Filigreed Shadow Circlet (253981, -0.06 DPS) [crafted]; Silk Headband (7050, -0.74 DPS) [crafted]; Shadow Circlet (253955, -0.83 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 0.5 spell_power points (0.17 DPS) | yes | Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Seal of the Expedition (271920, -0.04 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.37 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.7 spell_power points (3.13 DPS) | yes | Invoker's Mantle (215365, -0.76 DPS, sim-verified) [crafted]; Moonlit Amice (11884, -0.88 DPS) [quest]; Death Speaker Mantle (6685, -0.91 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest]; Heavy Woolen Cloak (4311, -0.25 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (4.18 DPS) | yes | Necrology Robes (2292, -0.24 DPS, sim-verified) [dungeon]; Shadow Gown (253967, -0.61 DPS) [crafted]; Beguiler Robes (7728, -0.84 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.89 DPS) | yes | Phoenix Bindings (210781, -2.46 DPS, sim-verified) [crafted]; Windsong Bangles (263336, -2.57 DPS) [quest]; Glowing Magical Bracelets (13106, -2.68 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.7 spell_power points (2.48 DPS) | yes | Shilly Mitts (9609, -0.08 DPS) [quest]; Hands of Darkness (7047, -0.10 DPS) [crafted]; Serpent Gloves (5970, -0.23 DPS) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.2 spell_power points (3.62 DPS) | yes | Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -1.04 DPS) [dungeon]; Invoker's Cord (215366, -1.23 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.5 spell_power points (4.01 DPS) | yes | Shadow Leggings (253993, -0.17 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.98 DPS) [crafted]; Necromancer Leggings (2277, -1.17 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.1 spell_power points (3.57 DPS) | yes | Gilded Slippers (254001, -1.21 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.22 DPS) [crafted]; Glinteye Slippers (273024, -1.22 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.25 DPS) | yes | Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Minor Channeling Ring (1449, -0.59 DPS) [quest]; Runemender's Seal (273031, -1.01 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.9 spell_power points (2.23 DPS) | yes | Minor Channeling Ring (1449, -0.57 DPS) [quest]; Ring of Power Regulation (273030, -0.60 DPS, sim-verified) [dungeon]; Runemender's Seal (273031, -0.99 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.35 DPS) [dungeon]; Spine of the Devourer (273040, -1.35 DPS) [world]; Tendonscraper Dagger (273053, -1.55 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.89 DPS) | yes | Orb of Mystic Insight (249394, -0.36 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.29 DPS) [world_drop]; Orb of Souls (249395, -1.51 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Eventide (5214, -0.32 DPS) [world_drop]; Greater Mystic Wand (217287, -0.32 DPS) [crafted]; Gravestone Scepter (7001, -0.35 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 00000000000000000-0000000000000000000-2050235103101341)

Set DPS (verified): 189.5. Weights run: 3.9s. Verify run: 3.5s. 455 eligible items had no known source.

Race sweep: human 189.5 ± 0.2, gnome 188.3 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.194 ± 0.017, crit=0.234 ± 0.006 per rating point (14 rating = 1%, 3.274 per %), hit=0.492 ± 0.023 per rating point (10 rating = 1%, 4.925 per %), spell_haste=not significant (0.640 ± 0.246), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.418 ± 0.000, fire_power=0.582 ± 0.001, spirit=not significant (0.265 ± 0.181)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.44 DPS) | yes | Miner's Hat of the Deep (9429, -2.29 DPS) [dungeon]; Papal Fez (9431, -2.94 DPS) [dungeon]; Living Cowl (5608, -3.21 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.9 spell_power points (3.96 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.52 DPS) [dungeon]; Necklace of Calisea (1714, -2.67 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.32 DPS) | yes | Mantle of Doan (7712, -0.19 DPS) [dungeon]; Green Silken Shoulders (7057, -0.25 DPS) [crafted]; Inquisitor's Shawl (19507, -0.49 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (19532, -0.51 DPS) [rep]; Long Silken Cloak (4326, -0.65 DPS) [crafted]; Guardian Cloak (5965, -0.65 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.5 spell_power points (9.84 DPS) | yes | Dreamweave Vest (10021, -0.42 DPS) [crafted]; Elemental Raiment (9434, -1.40 DPS) [world_drop]; Robe of Power (7054, -2.43 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcane Runed Bracers (4744, +0.00 DPS) [quest]; Phoenix Bindings (210781, -0.58 DPS) [crafted]; Condor Bracers (15864, -0.80 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (8.29 DPS) | yes | Black Mageweave Gloves (10003, -2.26 DPS) [crafted]; Red Mageweave Gloves (10018, -3.09 DPS) [crafted]; Fiery Handwraps (254025, -4.47 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.8 spell_power points (5.94 DPS) | yes | Star Belt (4329, -0.29 DPS) [crafted]; Fiery Cord (254041, -1.40 DPS) [crafted]; Gilded Cord (254037, -1.46 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.3 spell_power points (6.56 DPS) | yes | Gaze Dreamer Pants (6903, -0.78 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.85 DPS) [dungeon]; Black Mageweave Leggings (9999, -1.86 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.9 spell_power points (10.38 DPS) | yes | Fiery Slippers (254005, -6.37 DPS) [crafted]; Gilded Slippers (254001, -6.60 DPS) [crafted]; Wingborne Boots (15104, -6.91 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (4.41 DPS) | yes | Snake Hoop (6750, -0.71 DPS) [quest]; Underworld Band (1980, -1.03 DPS) [world_drop]; Ring of Forlorn Spirits (2043, -1.19 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Underworld Band (1980, -0.24 DPS) [world_drop]; Ring of Forlorn Spirits (2043, -0.40 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.49 DPS) [dungeon]; Hypnotic Blade (7714, -11.08 DPS) [dungeon]; Swinetusk Shank (6691, -13.20 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 455, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 301.1. Weights run: 3.7s. Verify run: 6.7s. 592 eligible items had no known source.

Race sweep: human 301.1 ± 0.2, gnome 299.0 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.428 ± 0.017, crit=0.168 ± 0.004 per rating point (14 rating = 1%, 2.353 per %), hit=0.340 ± 0.019 per rating point (10 rating = 1%, 3.401 per %), spell_haste=not significant (-0.937 ± 0.256), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.426 ± 0.000, fire_power=0.574 ± 0.001, spirit=not significant (0.548 ± 0.268)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 34.5 spell_power points (14.20 DPS) | yes | Dreamweave Circlet (10041, -2.44 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.85 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.08 DPS) [crafted] |
| neck | Talvash's Enhancing Necklace (7673) | Restoring the Necklace [quest] | 12.4 spell_power points (5.12 DPS) | yes | Mindburst Medallion (11196, -0.46 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.83 DPS) [dungeon]; Triune Amulet (7722, -1.31 DPS, sim-verified) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.9 spell_power points (9.42 DPS) | yes | Rotgrip Mantle (17732, -1.21 DPS, sim-verified) [dungeon]; Netherflame Shoulders (254053, -2.63 DPS) [crafted]; Nethergeld Shoulders (254049, -2.94 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -3.34 DPS, sim-verified) [crafted]; Featherskin Cape (10843, -3.56 DPS) [world]; Cindercloth Cloak (14044, -3.70 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.59 DPS, sim-verified) [crafted]; Vestments of the Atal'ai Prophet (10806, -0.93 DPS) [dungeon]; Embrace of the Wind Serpent (12462, -1.48 DPS) [world] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 13.2 spell_power points (5.44 DPS) | yes | Nethergeld Cuffs (254061, -0.19 DPS) [crafted]; Netherpearl Cuffs (254067, -0.80 DPS) [crafted]; Bloodband Bracers (11469, -1.12 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.5 spell_power points (9.71 DPS) | yes | Fiery Gloves (254099, -1.64 DPS) [crafted]; Gilded Gloves (254095, -2.01 DPS) [crafted]; Runecloth Gloves (13863, -2.31 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Satyrmane Sash (17755, -0.58 DPS) [dungeon]; Ghostweave Cord (254073, -1.56 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20098, -1.64 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 27.3 spell_power points (11.24 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.76 DPS) [dungeon]; Rainstrider Leggings (11123, -2.89 DPS) [quest]; Senior Designer's Pantaloons (11841, -4.01 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (11.47 DPS) | yes | Fiery Sandals (254111, -5.15 DPS) [crafted]; Vinerot Sandals (17748, -5.15 DPS) [dungeon]; Mistwalker Boots (10629, -6.83 DPS, sim-verified) [dungeon] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 18.0 spell_power points (7.43 DPS) | yes | Snake Hoop (6750, -2.14 DPS) [quest]; Philanthropist's Ring (281635, -2.42 DPS) [quest]; Cyclopean Band (11824, -5.86 DPS, sim-verified) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.07 DPS) [quest]; Philanthropist's Ring (281635, -0.35 DPS) [quest]; Cyclopean Band (11824, -1.16 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -5.07 DPS, sim-verified) [world_drop] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (5.77 DPS) | yes | Enthralled Sphere (11625, -1.03 DPS) [dungeon]; Truesilver Conduit (249455, -1.48 DPS, sim-verified) [crafted]; Twisting Essence Jar (249456, -1.49 DPS) [crafted] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.46 DPS) [crafted]; Dreambough Wand (249234, -0.87 DPS) [crafted]; Charged Lightning Rod (11860, -0.99 DPS) [quest] |

**New at 50:** head: Eye of Flame; neck: Talvash's Enhancing Necklace; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Brainlash; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Orb of the Forgotten Seer; ranged: Pyric Caduceus

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 491.7. Weights run: 4.0s. Verify run: 21.9s. 1444 eligible items had no known source.

Race sweep: human 491.7 ± 0.4, gnome 487.8 ± 0.4 (dps); the guide race human is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.157 ± 0.023, crit=0.233 ± 0.007 per rating point (14 rating = 1%, 3.268 per %), hit=0.526 ± 0.028 per rating point (10 rating = 1%, 5.261 per %), spell_haste=not significant (-0.893 ± 0.375), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.423 ± 0.001, fire_power=0.577 ± 0.001, spirit=not significant (0.497 ± 0.381)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 36.5 spell_power points (16.97 DPS) | yes | Field Marshal's Coronal (231584, -0.36 DPS) [pvp]; Eye of Flame (3075, -1.71 DPS, sim-verified) [world_drop]; Ghostweave Hood (254135, -3.06 DPS) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 15.5 spell_power points (7.21 DPS) | yes | Lady Maye's Pendant (14558, -0.71 DPS) [world_drop]; Jeweled Amulet of Cainwyn (1443, -0.80 DPS) [world_drop]; Animated Chain Necklace (18723, -7.02 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 28.5 spell_power points (13.23 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.38 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.86 DPS) [crafted]; Burial Shawl (18681, -4.89 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -2.13 DPS) [dungeon]; Shroud of the Exile (15421, -2.54 DPS) [quest]; Crystalline Threaded Cape (20697, -6.50 DPS, sim-verified) [world] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.4 spell_power points (22.02 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.41 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -8.95 DPS) [pvp]; Robes of the Exalted (13346, -16.13 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.7 spell_power points (12.42 DPS) | yes | Wyrmthalak's Shackles (13958, -5.70 DPS) [quest]; Earthenweave Cuffs (254131, -6.39 DPS) [crafted]; Runecloth Cuffs (254123, -6.58 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 28.4 spell_power points (13.18 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.97 DPS) [quest]; Dreamweave Gloves (10019, -4.54 DPS, sim-verified) [crafted] |
| waist | Stormpike Cloth Girdle (19094) | Stormpike Guard [rep] | 19.6 spell_power points (9.09 DPS) | yes | Thuzadin Sash (18740, -0.10 DPS) [dungeon]; Ban'thok Sash (11662, -0.34 DPS) [dungeon]; Ghostweave Cord (254073, -3.86 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.1 spell_power points (18.62 DPS) | yes | Marshal's Dreadweave Leggings (231587, -0.06 DPS) [pvp]; Ghostweave Pants (14144, -1.93 DPS) [crafted]; Skyshroud Leggings (13170, -3.82 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.5 spell_power points (12.76 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.70 DPS) [crafted]; Ghostweave Boots (254129, -2.33 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -1.23 DPS) [dungeon]; Songstone of Ironforge (12543, -1.38 DPS) [quest]; Naglering (11669, -10.72 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Songstone of Ironforge (12543, -0.06 DPS) [quest]; The Postmaster's Seal (13392, -0.82 DPS, sim-verified) [dungeon]; Brainlash (6440, -1.94 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, -3.72 DPS, sim-verified) [dungeon]; Serenity Field (272439, -5.11 DPS) [vendor]; Uther's Strength (11302, -5.57 DPS) [world_drop] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, -4.20 DPS, sim-verified) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, -0.12 DPS) [dungeon]; Ironbark Staff (20069, -1.54 DPS) [rep]; Blade of Eternal Darkness (17780, -19.69 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.47 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.91 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.22 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.71 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Greater Eternal Wand (249237, -0.51 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Stormpike Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 25500000033000000-0000000000000000000-0050355103101351)

Set DPS (verified): 880.0. Weights run: 4.4s. Verify run: 20.9s. 1444 eligible items had no known source.

Race sweep: human 880.0 ± 0.7, gnome 876.0 ± 0.7 (dps); the guide race human is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.340 ± 0.046, crit=0.693 ± 0.018 per rating point (14 rating = 1%, 9.703 per %), hit=1.210 ± 0.073 per rating point (10 rating = 1%, 12.099 per %), spell_haste=not significant (3.455 ± 1.016), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.367 ± 0.001, fire_power=0.633 ± 0.001, spirit=not significant (-0.182 ± 0.241)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.4 spell_power points (22.62 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Eye of Flame (3075, -1.89 DPS) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.8 spell_power points (16.13 DPS) | yes | Star of Mystaria (12103, -3.87 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.31 DPS) [world_drop]; Archlight Talisman (15856, -5.70 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.1 spell_power points (21.08 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.23 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.39 DPS) [pvp]; Burial Shawl (18681, -6.16 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.40 DPS, sim-verified) [world]; Spritecaster Cape (11623, -10.01 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -10.15 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.1 spell_power points (33.23 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.03 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -11.69 DPS) [pvp]; Alanna's Embrace (13314, -14.74 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.7 spell_power points (16.74 DPS) | yes | Runecloth Cuffs (254123, -6.99 DPS) [crafted]; Manacle Cuffs (11962, -8.55 DPS) [quest]; Funeral Cuffs (12626, -9.45 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.8 spell_power points (21.54 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Hands of Power (13253, -3.89 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Handwraps (227100, -6.40 DPS) [pvp] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.5 spell_power points (18.62 DPS) | yes | Highlander's Cloth Girdle (20047, -1.19 DPS) [rep]; Stormpike Cloth Girdle (19094, -4.13 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.8 spell_power points (37.12 DPS) | yes | Marshal's Dreadweave Leggings (231587, -7.69 DPS) [pvp]; Skyshroud Leggings (13170, -14.04 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -15.16 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.1 spell_power points (16.31 DPS) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.05 DPS) [crafted] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.94 DPS) [dungeon]; Maiden's Circle (13001, -1.16 DPS) [world_drop]; Naglering (11669, -17.77 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.74 DPS) [dungeon]; Songstone of Ironforge (12543, -0.96 DPS) [quest]; Maiden's Circle (13001, -0.96 DPS) [world_drop] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (+13.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Cannonball Runner (13382, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -4.00 DPS) [dungeon]; Draconic Infused Emblem (22268, -4.29 DPS, sim-verified) [dungeon]; Frozen Heart of the Mountain (249469, -4.82 DPS) [crafted] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.90 DPS) [pvp]; Headmaster's Charge (13937, -14.98 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -18.67 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.26 DPS) [dungeon]; Greater Eternal Wand (249237, -1.61 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 00000000000000000-0000000000000000000-0050230100000000)

Set DPS (verified): 61.9. Weights run: 3.9s. Verify run: 5.4s. 193 eligible items had no known source.

Race sweep: undead 62.5 ± 0.0, orc 61.9 ± 0.0, troll 60.9 ± 0.0 (dps); undead beats the guide race orc by 0.5 (0.87%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.106 ± 0.003, crit=0.026 ± 0.001 per rating point (14 rating = 1%, 0.365 per %), hit=0.061 ± 0.003 per rating point (10 rating = 1%, 0.605 per %), spell_haste=0.259 ± 0.026, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.800 ± 0.001, fire_power=0.200 ± 0.000, spirit=0.142 ± 0.035

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.8 spell_power points (2.60 DPS) | yes | Pristine Circlet (253949, -0.97 DPS, sim-verified) [crafted]; Flame Circlet (253953, -1.95 DPS) [crafted]; Shadow Goggles (4373, -2.19 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.4 spell_power points (0.13 DPS) | yes | Scholarly Pendant (277203, -0.04 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.0 spell_power points (1.76 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.19 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.58 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (20427, -0.28 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.46 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.67 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 11.7 spell_power points (3.47 DPS) | yes | Filigreed Shadow Gown (253907, -1.01 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.71 DPS) [crafted]; Scholarly Robes (2034, -2.08 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.89 DPS) | yes | Owlbeard Bracers (16981, -0.54 DPS, sim-verified) [quest]; Windsong Bangles (263336, -0.59 DPS) [quest]; Mindthrust Bracers (1974, -0.73 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.07 DPS) | yes | Gnoll Casting Gloves (892, -0.32 DPS, sim-verified) [world]; Shadow Gloves (253919, -0.32 DPS) [crafted]; Blight Gloves (279877, -0.59 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.0 spell_power points (1.78 DPS) | yes | Novice Arcanist's Sash (253885, -0.50 DPS) [crafted]; Pristine Sash (253925, -0.56 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.97 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.0 spell_power points (2.96 DPS) | yes | Silk-threaded Trousers (1929, -0.89 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.97 DPS, sim-verified) [crafted]; Colorful Kilt (10048, -1.27 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.4 spell_power points (2.19 DPS) | yes | Shadow Boots (253895, -0.70 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.80 DPS) [crafted]; Pristine Boots (253889, -1.21 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.48 DPS) | yes | Lavishly Jeweled Ring (1156, -1.29 DPS) [dungeon]; Loop of Sacrifice (281673, -1.32 DPS) [quest]; Furlfeather Ring (286745, -1.54 DPS, sim-verified) [world] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.06 DPS) [dungeon]; Furlfeather Ring (286745, -0.09 DPS, sim-verified) [world]; Loop of Sacrifice (281673, -0.10 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (61.9 DPS) | yes | Staff of the Blessed Seer (2271, -0.19 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -0.44 DPS) [quest]; Crescent Staff (6505, -0.64 DPS) [quest] |
| off_hand | Antipodean Rod (2879) | World drop [world_drop] | 1.2 spell_power points (0.35 DPS) | yes | Grayson's Torch (1172, -0.19 DPS) [quest]; Seer's Fine Stein (7608, -0.22 DPS) [world_drop]; Nightglow Concoction (3451, -0.23 DPS) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.30 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Evocator's Blade; off_hand: Antipodean Rod; ranged: Sizzle Stick

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade

### Band 30 (orc, 00000000000000000-0000000000000000000-1050235103001000)

Set DPS (verified): 106.2. Weights run: 3.7s. Verify run: 5.1s. 319 eligible items had no known source.

Race sweep: orc 106.2 ± 0.1, undead 106.2 ± 0.1, troll 104.4 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.156 ± 0.008, crit=0.147 ± 0.004 per rating point (14 rating = 1%, 2.061 per %), hit=0.259 ± 0.011 per rating point (10 rating = 1%, 2.593 per %), spell_haste=not significant (0.049 ± 0.129), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.809 ± 0.001, fire_power=0.191 ± 0.000, spirit=not significant (0.159 ± 0.081)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.0 spell_power points (4.12 DPS) | yes | Filigreed Shadow Circlet (253981, -0.21 DPS) [crafted]; Silk Headband (7050, -1.02 DPS) [crafted]; Shadow Circlet (253955, -1.05 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 1.3 spell_power points (0.43 DPS) | yes | Crystal Starfire Medallion (5003, -0.05 DPS) [world_drop]; Pendant of Myzrael (4614, -0.10 DPS) [dungeon]; Glowing Green Talisman (5002, -0.10 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (3.58 DPS) | yes | Invoker's Mantle (215365, -0.85 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.93 DPS) [dungeon]; Magician's Mantle (12998, -1.38 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heavy Woolen Cloak (4311, -0.34 DPS) [crafted]; Battle Healer's Cloak (19529, -0.34 DPS) [rep]; Hillman's Cloak (3719, -0.50 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (4.48 DPS) | yes | Shadow Gown (253967, -0.26 DPS) [crafted]; Beguiler Robes (7728, -0.30 DPS) [dungeon]; Necrology Robes (2292, -0.31 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.10 DPS) | yes | Phoenix Bindings (210781, -2.25 DPS) [crafted]; Tabitha's Cuffs (251486, -2.27 DPS, sim-verified) [quest]; Nightsky Wristbands (6407, -2.61 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.4 spell_power points (2.88 DPS) | yes | Hands of Darkness (7047, -0.10 DPS) [crafted]; Serpent Gloves (5970, -0.47 DPS) [dungeon]; Jutebraid Gloves (10654, -0.55 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (3.95 DPS) | yes | Warsong Sash (16975, -0.16 DPS) [quest]; Belt of Arugal (6392, -0.42 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -1.19 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 13.4 spell_power points (4.63 DPS) | yes | Shadow Leggings (253993, -0.07 DPS) [crafted]; Filigreed Shadow Leggings (253943, -1.02 DPS) [crafted]; Necromancer Leggings (2277, -1.25 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 12.2 spell_power points (4.22 DPS) | yes | Gilded Slippers (254001, -1.23 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.48 DPS) [dungeon]; Spidersilk Boots (4320, -1.59 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.2 spell_power points (2.82 DPS) | yes | Ring of Power Regulation (273030, -0.76 DPS) [dungeon]; Runemender's Seal (273031, -1.09 DPS) [world]; Solid Gold Band (279845, -1.17 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.41 DPS) | yes | Ring of Power Regulation (273030, -0.65 DPS, sim-verified) [dungeon]; Runemender's Seal (273031, -0.67 DPS) [world]; Solid Gold Band (279845, -0.76 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.60 DPS) [dungeon]; Spine of the Devourer (273040, -1.60 DPS) [world]; Tendonscraper Dagger (273053, -1.88 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (3.10 DPS) | yes | Orb of Mystic Insight (249394, -0.37 DPS) [crafted]; Orb of Mistmantle (13031, -1.21 DPS) [world_drop]; Orb of Souls (249395, -1.39 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, -0.23 DPS) [quest]; Wand of Eventide (5214, -0.34 DPS) [world_drop]; Greater Mystic Wand (217287, -0.35 DPS, sim-verified) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (orc, 00000000000000000-0000000000000000000-2050235103101341)

Set DPS (verified): 189.5. Weights run: 3.9s. Verify run: 3.9s. 448 eligible items had no known source.

Race sweep: orc 189.5 ± 0.2, undead 188.8 ± 0.2, troll 186.8 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.194 ± 0.017, crit=0.234 ± 0.006 per rating point (14 rating = 1%, 3.274 per %), hit=0.492 ± 0.023 per rating point (10 rating = 1%, 4.925 per %), spell_haste=not significant (0.640 ± 0.246), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.418 ± 0.000, fire_power=0.582 ± 0.001, spirit=not significant (0.265 ± 0.181)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.44 DPS) | yes | Miner's Hat of the Deep (9429, -2.67 DPS, sim-verified) [dungeon]; Papal Fez (9431, -2.94 DPS) [dungeon]; Living Cowl (5608, -3.21 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.9 spell_power points (3.96 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.52 DPS) [dungeon]; Necklace of Calisea (1714, -2.67 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.32 DPS) | yes | Mantle of Doan (7712, -0.19 DPS) [dungeon]; Green Silken Shoulders (7057, -0.25 DPS) [crafted]; Inquisitor's Shawl (19507, -0.49 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (189.5 DPS) | yes | Long Silken Cloak (4326, -0.65 DPS) [crafted]; Guardian Cloak (5965, -0.65 DPS) [crafted]; Battle Healer's Cloak (19528, -0.84 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.5 spell_power points (9.84 DPS) | yes | Dreamweave Vest (10021, -0.71 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.40 DPS) [world_drop]; Robe of Power (7054, -2.43 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.62 DPS) | yes | Condor Bracers (15864, -0.80 DPS) [quest]; Phoenix Bindings (210781, -0.94 DPS, sim-verified) [crafted]; Radiant Silver Bracers (4545, -1.38 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.6 spell_power points (8.29 DPS) | yes | Black Mageweave Gloves (10003, -2.75 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -3.09 DPS) [crafted]; Fiery Handwraps (254025, -4.47 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.8 spell_power points (5.94 DPS) | yes | Star Belt (4329, -0.29 DPS) [crafted]; Fiery Cord (254041, -1.40 DPS) [crafted]; Gilded Cord (254037, -1.46 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.3 spell_power points (6.56 DPS) | yes | Gaze Dreamer Pants (6903, -1.20 DPS, sim-verified) [dungeon]; Stoneweaver Leggings (9407, -1.85 DPS) [dungeon]; Black Mageweave Leggings (9999, -1.86 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.9 spell_power points (10.38 DPS) | yes | Fiery Slippers (254005, -6.02 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -6.60 DPS) [crafted]; Glinteye Slippers (273024, -6.94 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (4.41 DPS) | yes | Advisor's Ring (19520, -0.79 DPS) [rep]; Underworld Band (1980, -1.03 DPS) [world_drop]; Reedknot Ring (9622, -1.60 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (3.70 DPS) | yes | Advisor's Ring (19520, -0.09 DPS) [rep]; Underworld Band (1980, -0.32 DPS) [world_drop]; Reedknot Ring (9622, -0.89 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (189.5 DPS) | yes | Illusionary Rod (7713, -8.75 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -11.08 DPS) [dungeon]; Swinetusk Shank (6691, -13.20 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (189.5 DPS) | yes | Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.40 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 300.9. Weights run: 3.7s. Verify run: 9.7s. 585 eligible items had no known source.

Race sweep: orc 300.9 ± 0.2, undead 299.9 ± 0.2, troll 297.7 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.168 ± 0.021, crit=0.299 ± 0.008 per rating point (14 rating = 1%, 4.184 per %), hit=0.639 ± 0.028 per rating point (10 rating = 1%, 6.389 per %), spell_haste=1.781 ± 0.302, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.369 ± 0.000, fire_power=0.631 ± 0.001, spirit=not significant (0.422 ± 0.158)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 33.0 spell_power points (15.80 DPS) | yes | Dreamweave Circlet (10041, -2.08 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme Plus (15999, -2.88 DPS) [crafted]; Red Mageweave Headband (10033, -5.10 DPS) [crafted] |
| neck | Jarkal's Enhancing Necklace (7888) | Necklace Recovery, Take 3 [quest] | 12.2 spell_power points (5.82 DPS) | yes | Triune Amulet (7722, -0.58 DPS) [dungeon]; Mindburst Medallion (11196, -1.46 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -1.52 DPS) [quest] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.7 spell_power points (8.96 DPS) | yes | Rotgrip Mantle (17732, -1.29 DPS) [dungeon]; Netherflame Shoulders (254053, -1.97 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.40 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.93 DPS) [quest]; Runecloth Cloak (13860, -3.44 DPS) [crafted]; Cindercloth Cloak (14044, -3.83 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.82 DPS, sim-verified) [crafted]; Vestments of the Atal'ai Prophet (10806, -1.36 DPS) [dungeon]; Elemental Raiment (9434, -1.97 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 11.5 spell_power points (5.50 DPS) | yes | Nethergeld Cuffs (254061, -0.57 DPS) [crafted]; Spidertank Oilrag (9448, -1.19 DPS) [dungeon]; Phoenix Bindings (210781, -1.57 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.6 spell_power points (10.35 DPS) | yes | Runecloth Gloves (13863, -2.06 DPS) [crafted]; Fiery Gloves (254099, -2.20 DPS) [crafted]; Cindercloth Gloves (14043, -3.00 DPS) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 20.1 spell_power points (9.61 DPS) | yes | Satyrmane Sash (17755, -2.10 DPS) [dungeon]; Ghostweave Cord (254073, -2.19 DPS, sim-verified) [crafted]; Star Belt (4329, -2.58 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 24.7 spell_power points (11.81 DPS) | yes | Wizardweave Leggings (14132, -2.72 DPS) [crafted]; Cindercloth Leggings (12256, -3.05 DPS) [vendor]; Senior Designer's Pantaloons (11841, -3.51 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.0 spell_power points (12.90 DPS) | yes | Mistwalker Boots (10629, -5.23 DPS, sim-verified) [dungeon]; First Sergeant's Dreadweave Boots (220909, -5.29 DPS) [vendor]; Fiery Sandals (254111, -6.14 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 14.8 spell_power points (7.08 DPS) | yes | Advisor's Ring (19519, -1.34 DPS) [rep]; Cyclopean Band (11824, -1.40 DPS) [dungeon]; Philanthropist's Ring (281635, -1.89 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.22 DPS) | yes | Advisor's Ring (19519, -0.48 DPS) [rep]; Cyclopean Band (11824, -0.54 DPS) [dungeon]; Philanthropist's Ring (281635, -1.03 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.33 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.61 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.73 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.42 DPS, sim-verified) [quest]; Ankh of Life (1713, -0.45 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.85 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -6.43 DPS, sim-verified) [world_drop] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (6.70 DPS) | yes | Truesilver Conduit (249455, -1.27 DPS) [crafted]; Flask of Forest Mojo (19115, -1.75 DPS) [quest]; Beacon of Hope (9393, -2.06 DPS) [dungeon] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.70 DPS) [crafted]; Dreambough Wand (249234, -1.18 DPS) [crafted]; Twisted Nether Wand (249144, -1.66 DPS) [crafted] |

**New at 50:** head: Eye of Flame; neck: Jarkal's Enhancing Necklace; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Brainlash; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Orb of the Forgotten Seer; ranged: Pyric Caduceus

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 485.8. Weights run: 4.0s. Verify run: 22.1s. 1438 eligible items had no known source.

Race sweep: undead 486.2 ± 0.4, orc 485.8 ± 0.4, troll 479.8 ± 0.4 (dps); undead beats the guide race orc by 0.4 (0.09%) ± 0.5, inside error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.204 ± 0.034, crit=0.462 ± 0.012 per rating point (14 rating = 1%, 6.474 per %), hit=0.895 ± 0.044 per rating point (10 rating = 1%, 8.952 per %), spell_haste=not significant (-1.131 ± 0.492), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.386 ± 0.001, fire_power=0.614 ± 0.001, spirit=not significant (0.110 ± 0.307)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.1 spell_power points (17.70 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.07 DPS) [pvp]; Eye of Flame (3075, -1.65 DPS, sim-verified) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 19.6 spell_power points (10.46 DPS) | yes | Archlight Talisman (15856, -2.96 DPS) [quest]; Star of Mystaria (12103, -3.07 DPS) [dungeon]; Lady Maye's Pendant (14558, -3.30 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.1 spell_power points (13.96 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.19 DPS) [crafted]; Burial Shawl (18681, -1.53 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.18 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.72 DPS) [dungeon]; Shroud of the Exile (15421, -6.29 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.8 spell_power points (25.56 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.85 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.02 DPS) [pvp]; Alanna's Embrace (13314, -13.84 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.4 spell_power points (13.03 DPS) | yes | Runecloth Cuffs (254123, -6.07 DPS) [crafted]; Manacle Cuffs (11962, -7.66 DPS) [quest]; Netherflame Cuffs (254065, -7.72 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 25.7 spell_power points (13.74 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Gloves of Spell Mastery (14146, -0.92 DPS) [crafted]; Earth Warder's Gloves (21318, -1.98 DPS) [quest] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 23.0 spell_power points (12.28 DPS) | yes | Defiler's Cloth Girdle (20163, -0.69 DPS) [rep]; Frostwolf Cloth Belt (19090, -1.58 DPS) [rep]; Valconian Sash (12083, -3.20 DPS) [quest] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 33.0 spell_power points (17.61 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -1.24 DPS) [pvp]; Skyshroud Leggings (13170, -1.82 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.8 spell_power points (13.23 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.01 DPS) [crafted]; Omnicast Boots (11822, -1.24 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -0.14 DPS) [vendor]; Don Mauricio's Band of Domination (22433, -1.39 DPS) [dungeon]; Naglering (11669, -11.44 DPS, sim-verified) [dungeon] |
| finger2 | Eye of Orgrimmar (12545) | The Princess Saved? [quest] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Don Mauricio's Band of Domination (22433, -1.04 DPS) [dungeon]; Blessed Band of Light (272407, -1.13 DPS, sim-verified) [vendor]; Don Julio's Band (19325, -2.14 DPS) [rep] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+8.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -4.83 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -5.31 DPS) [crafted]; Second Wind (11819, -5.57 DPS, sim-verified) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20220, -0.09 DPS) [rep]; High Warlord's War Staff (234549, -2.96 DPS) [pvp]; Blade of Eternal Darkness (17780, -16.53 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (15.49 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.72 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.20 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -4.27 DPS) [quest] |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.31 DPS) [crafted]; Greater Eternal Wand (249237, -1.07 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Eye of Orgrimmar; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 25500000033000000-0000000000000000000-0050355103101351)

Set DPS (verified): 869.4. Weights run: 4.4s. Verify run: 21.2s. 1438 eligible items had no known source.

Race sweep: orc 869.4 ± 0.7, undead 865.8 ± 0.7, troll 858.5 ± 0.7 (dps); the guide race orc is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.340 ± 0.046, crit=0.693 ± 0.018 per rating point (14 rating = 1%, 9.703 per %), hit=1.210 ± 0.073 per rating point (10 rating = 1%, 12.099 per %), spell_haste=not significant (3.455 ± 1.016), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.367 ± 0.001, fire_power=0.633 ± 0.001, spirit=not significant (-0.182 ± 0.241)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.4 spell_power points (22.62 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Eye of Flame (3075, -1.89 DPS) [world_drop] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.8 spell_power points (16.13 DPS) | yes | Lady Maye's Pendant (14558, -4.31 DPS) [world_drop]; Star of Mystaria (12103, -5.03 DPS, sim-verified) [dungeon]; Archlight Talisman (15856, -5.70 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.1 spell_power points (21.08 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.23 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -3.39 DPS) [pvp]; Burial Shawl (18681, -4.37 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.48 DPS, sim-verified) [world]; Spritecaster Cape (11623, -10.01 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -10.15 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.1 spell_power points (33.23 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.03 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -11.69 DPS) [pvp]; Alanna's Embrace (13314, -16.63 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.7 spell_power points (16.74 DPS) | yes | Runecloth Cuffs (254123, -6.99 DPS) [crafted]; Manacle Cuffs (11962, -8.55 DPS) [quest]; Funeral Cuffs (12626, -9.45 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.8 spell_power points (21.54 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -6.40 DPS) [pvp]; Hands of Power (13253, -7.56 DPS, sim-verified) [dungeon] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 27.5 spell_power points (18.62 DPS) | yes | Defiler's Cloth Girdle (20163, -1.19 DPS) [rep]; Frostwolf Cloth Belt (19090, -4.13 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 34.5 spell_power points (23.34 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.26 DPS) [dungeon]; Legionnaire's Dreadweave Legguards (227097, -1.38 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.1 spell_power points (16.31 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.05 DPS) [crafted]; Dragonrider Boots (18102, -0.43 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.94 DPS) [dungeon]; Maiden's Circle (13001, -1.16 DPS) [world_drop]; Naglering (11669, -18.81 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.74 DPS) [dungeon]; Eye of Orgrimmar (12545, -0.96 DPS) [quest]; Maiden's Circle (13001, -0.96 DPS) [world_drop] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+15.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, -4.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -4.82 DPS) [crafted]; Serenity Field (272439, -7.95 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.90 DPS) [pvp]; Headmaster's Charge (13937, -13.12 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -18.67 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.26 DPS) [dungeon]; Greater Eternal Wand (249237, -1.61 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

