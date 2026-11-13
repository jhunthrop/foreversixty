# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 61.7. Weights run: 4.1s. Verify run: 3.5s. 199 eligible items had no known source.

Race sweep: human 61.7 ± 0.0, gnome 61.4 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.159 ± 0.006, crit=0.059 ± 0.001 per rating point (14 rating = 1%, 0.829 per %), hit=0.161 ± 0.006 per rating point (10 rating = 1%, 1.613 per %), spell_haste=not significant (0.301 ± 0.081), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.834 ± 0.001, fire_power=0.166 ± 0.000, spirit=0.249 ± 0.053

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.2 spell_power points (2.77 DPS) | yes | Pristine Circlet (253949, -0.96 DPS) [crafted]; Shadow Goggles (4373, -2.08 DPS) [crafted]; Flame Circlet (253953, -2.22 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.5 spell_power points (0.15 DPS) | yes | - |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.4 spell_power points (1.94 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.45 DPS) [rep]; Black Whelp Cloak (7283, -0.60 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.76 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.5 spell_power points (3.76 DPS) | yes | Filigreed Shadow Gown (253907, -1.03 DPS) [crafted]; Filigreed Pristine Gown (253901, -1.79 DPS) [crafted]; Scholarly Robes (2034, -1.95 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Windsong Bangles (263336, +0.00 DPS) [quest]; Repurposed Hair Band (281256, +0.00 DPS) [quest]; Seer's Cuffs (3645, -0.04 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.11 DPS) | yes | Shadow Gloves (253919, -0.21 DPS) [crafted]; Gnoll Casting Gloves (892, -0.30 DPS) [world]; Pristine Gloves (253913, -0.76 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.5 spell_power points (1.95 DPS) | yes | Pristine Sash (253925, -0.55 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.60 DPS) [crafted]; Novice Ardent's Sash (253887, -0.98 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 11.1 spell_power points (3.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.96 DPS) [crafted]; Silk-threaded Trousers (1929, -1.24 DPS) [dungeon]; Colorful Kilt (10048, -1.47 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (2.30 DPS) | yes | Shadow Boots (253895, -0.65 DPS) [crafted]; Red Woolen Boots (4313, -0.72 DPS) [crafted]; Pristine Boots (253889, -1.25 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.3 spell_power points (1.60 DPS) | yes | Band of Purification (12996, -1.15 DPS) [world_drop]; Furlfeather Ring (286745, -1.30 DPS) [world]; Lavishly Jeweled Ring (1156, -1.32 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.51 DPS) | yes | Band of Purification (12996, -1.06 DPS) [world_drop]; Furlfeather Ring (286745, -1.21 DPS) [world]; Lavishly Jeweled Ring (1156, -1.22 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Blessed Seer (2271, +0.00 DPS) [dungeon]; Emberstone Staff (5201, -1.68 DPS) [dungeon]; Witching Stave (1484, -1.88 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.2 spell_power points (0.98 DPS) | yes | Onu's Patience (281292, -0.60 DPS) [quest]; Grayson's Torch (1172, -0.68 DPS) [quest]; Antipodean Rod (2879, -0.68 DPS) [world_drop] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.30 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (human, 25520000132001000-0000000000000000000-0000000000000000)

Set DPS (verified): 103.2. Weights run: 4.3s. Verify run: 6.8s. 326 eligible items had no known source.

Race sweep: human 103.2 ± 0.1, gnome 102.5 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.355 ± 0.007, crit=0.079 ± 0.001 per rating point (14 rating = 1%, 1.100 per %), hit=0.155 ± 0.007 per rating point (10 rating = 1%, 1.546 per %), spell_haste=not significant (0.139 ± 0.107), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.369 ± 0.114)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Filigreed Shadow Circlet (253981, +0.00 DPS) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.39 DPS) | yes | - |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Moonlit Amice (11884, -0.56 DPS) [quest]; Invoker's Mantle (215365, -0.56 DPS) [crafted]; Death Speaker Mantle (6685, -0.83 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Prelacy Cape (7004, -0.28 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.62 DPS) | yes | Necrology Robes (2292, -0.43 DPS) [dungeon]; Shadow Gown (253967, -0.66 DPS) [crafted]; Green Silk Armor (7065, -1.11 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Phoenix Bindings (210781, -1.84 DPS) [crafted]; Windsong Bangles (263336, -2.23 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Darkness (7047, +0.00 DPS) [crafted]; Serpent Gloves (5970, -0.10 DPS) [dungeon]; Shilly Mitts (9609, -0.10 DPS) [quest] |
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

Set DPS (verified): 180.5. Weights run: 4.8s. Verify run: 4.2s. 455 eligible items had no known source.

Race sweep: human 180.5 ± 0.1, gnome 178.9 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.165 ± 0.010, crit=0.245 ± 0.003 per rating point (14 rating = 1%, 3.426 per %), hit=0.298 ± 0.014 per rating point (10 rating = 1%, 2.976 per %), spell_haste=not significant (0.346 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.885 ± 0.001, fire_power=0.115 ± 0.000, spirit=not significant (0.063 ± 0.092)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.67 DPS) | yes | Miner's Hat of the Deep (9429, -2.56 DPS) [dungeon]; Living Cowl (5608, -3.30 DPS) [world]; Augural Shroud (2620, -3.45 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (180.5 DPS) | yes | Prodigious Shadowshard Pendant (17773, +0.00 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.05 DPS) [dungeon]; Necklace of Calisea (1714, -2.83 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (4.33 DPS) | yes | Green Silken Shoulders (7057, -0.28 DPS) [crafted]; Inquisitor's Shawl (19507, -0.55 DPS) [dungeon]; Berylline Pads (4197, -0.76 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.23 DPS) [crafted]; Guardian Cloak (5965, -0.23 DPS) [crafted]; Caretaker's Cape (19532, -0.44 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (9.63 DPS) | yes | Elemental Raiment (9434, -0.95 DPS) [world_drop]; Dreamweave Vest (10021, -1.21 DPS) [crafted]; Shadoweave Robe (10004, -1.76 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (180.5 DPS) | yes | Arcane Runed Bracers (4744, +0.00 DPS) [quest]; Condor Bracers (15864, -0.83 DPS) [quest]; Earthen Silk Cuffs (254019, -1.93 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (7.89 DPS) | yes | Shadoweave Gloves (10023, -1.33 DPS) [crafted]; Black Mageweave Gloves (10003, -1.69 DPS) [crafted]; Black Handwraps (254027, -2.29 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | sim-verified (180.5 DPS) | yes | Highlander's Cloth Girdle (20098, +0.00 DPS) [rep]; Star Belt (4329, -0.35 DPS) [crafted] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 18.6 spell_power points (7.68 DPS) | yes | Red Mageweave Pants (10009, -1.08 DPS) [crafted]; Shadow Leggings (253993, -1.95 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.49 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 12.9 spell_power points (5.34 DPS) | yes | Gilded Slippers (254001, -1.86 DPS) [crafted]; Spidersilk Boots (4320, -2.17 DPS) [crafted]; Glinteye Slippers (273024, -2.29 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.2 spell_power points (5.46 DPS) | yes | Lorekeeper's Ring (19524, -1.74 DPS) [rep]; Ring of Forlorn Spirits (2043, -2.16 DPS) [quest]; Snake Hoop (6750, -2.32 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (4.47 DPS) | yes | Lorekeeper's Ring (19524, -0.75 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.17 DPS) [quest]; Snake Hoop (6750, -1.33 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.07 DPS) [dungeon]; Hypnotic Blade (7714, -9.45 DPS) [dungeon]; Swinetusk Shank (6691, -11.62 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Black Cord; legs: Shadoweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 455, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 25522300132001351-0005000000000000000-0010000000000000)

Set DPS (verified): 262.4. Weights run: 4.8s. Verify run: 5.0s. 592 eligible items had no known source.

Race sweep: human 262.4 ± 0.1, gnome 260.5 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.185 ± 0.012, crit=0.322 ± 0.004 per rating point (14 rating = 1%, 4.505 per %), hit=0.420 ± 0.020 per rating point (10 rating = 1%, 4.203 per %), spell_haste=0.731 ± 0.097, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.888 ± 0.001, fire_power=0.112 ± 0.000, spirit=not significant (0.139 ± 0.097)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.64 DPS) | yes | Dreamweave Circlet (10041, -1.07 DPS) [crafted]; Shadoweave Mask (10025, -1.66 DPS) [crafted]; Red Mageweave Headband (10033, -1.85 DPS) [crafted] |
| neck | Talvash's Enhancing Necklace (7673) | Restoring the Necklace [quest] | 12.2 spell_power points (5.25 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.37 DPS) [quest]; Triune Amulet (7722, -1.38 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.60 DPS) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Netherpearl Shoulders (254055, +0.00 DPS) [crafted]; Rotgrip Mantle (17732, -0.39 DPS) [dungeon]; Shadoweave Shoulders (10028, -0.93 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.36 DPS) [crafted]; Nightfall Drape (12465, -2.99 DPS) [dungeon]; Caretaker's Cape (19531, -3.50 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (10.26 DPS) | yes | Dreamweave Vest (10021, -0.95 DPS) [crafted]; Elemental Raiment (9434, -1.21 DPS) [world_drop]; Shadoweave Robe (10004, -1.81 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.5 spell_power points (5.83 DPS) | yes | Arcane Runed Bracers (4744, -1.95 DPS) [quest]; Spidertank Oilrag (9448, -1.95 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.96 DPS) [crafted] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 19.8 spell_power points (8.55 DPS) | yes | Dreamweave Gloves (10019, -0.05 DPS) [crafted]; Shadoweave Gloves (10023, -1.64 DPS) [crafted]; Black Mageweave Gloves (10003, -2.08 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.5 spell_power points (7.99 DPS) | yes | Ban'thok Sash (11662, -0.21 DPS) [dungeon]; Satyrmane Sash (17755, -1.16 DPS) [dungeon]; Ghostweave Cord (254073, -1.29 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.3 spell_power points (10.91 DPS) | yes | Spellshock Leggings (9484, -0.19 DPS) [dungeon]; Wizardweave Leggings (14132, -2.72 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.83 DPS) [vendor] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 19.4 spell_power points (8.37 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -2.39 DPS) [vendor]; Shadoweave Boots (10031, -2.41 DPS) [crafted]; Black Slippers (254007, -2.60 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 13.5 spell_power points (5.84 DPS) | yes | Band of the Unicorn (7553, -0.24 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.67 DPS) [rep]; Philanthropist's Ring (281635, -1.13 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 13.4 spell_power points (5.76 DPS) | yes | Band of the Unicorn (7553, -0.15 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.58 DPS) [rep]; Philanthropist's Ring (281635, -1.05 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.91 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.03 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Jordan (873, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS) [dungeon] |
| off_hand | Umbral Crystal (13029) | World drop [world_drop] | 15.1 spell_power points (6.51 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.08 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.47 DPS) [dungeon]; Orb of Dar'Orahil (15108, -0.75 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.43 DPS) [crafted]; Burning Sliver (5249, -0.86 DPS) [quest]; Twisted Nether Wand (249144, -0.86 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Talvash's Enhancing Necklace; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Underworld Band; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Umbral Crystal; ranged: Lesser Eternal Wand

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522300132001351-2005220000000000000-0050000000000000)

Set DPS (verified): 409.2. Weights run: 5.0s. Verify run: 23.8s. 1458 eligible items had no known source.

Race sweep: human 409.2 ± 0.2, gnome 407.0 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.223 ± 0.017, crit=0.470 ± 0.006 per rating point (14 rating = 1%, 6.580 per %), hit=0.490 ± 0.025 per rating point (10 rating = 1%, 4.901 per %), spell_haste=4.341 ± 0.297, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.895 ± 0.001, fire_power=0.105 ± 0.000, spirit=not significant (0.102 ± 0.112)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 39.2 spell_power points (17.86 DPS) | yes | Field Marshal's Coronal (231584, -0.85 DPS) [pvp]; Crimson Felt Hat (18727, -3.33 DPS, sim-verified) [dungeon]; Lieutenant Commander's Dreadweave Cowl (227093, -3.47 DPS) [pvp] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 19.7 spell_power points (8.97 DPS) | yes | Nacreous Shell Necklace (22403, -1.90 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -2.03 DPS) [world_drop]; Archlight Talisman (15856, -2.49 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.5 spell_power points (12.06 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -1.09 DPS, sim-verified) [crafted]; Felcloth Shoulders (14112, -1.23 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.50 DPS, sim-verified) [world]; Spritecaster Cape (11623, -3.07 DPS) [dungeon]; Shroud of the Exile (15421, -3.53 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.0 spell_power points (21.86 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.85 DPS) [pvp]; Robe of Winter Night (14136, -5.04 DPS, sim-verified) [crafted]; Felcloth Robe (14106, -7.18 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.5 spell_power points (11.15 DPS) | yes | Netherpearl Cuffs (254067, -4.91 DPS) [crafted]; Runecloth Cuffs (254123, -5.13 DPS) [crafted]; Manacle Cuffs (11962, -6.44 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 25.8 spell_power points (11.76 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Gloves of Spell Mastery (14146, -0.65 DPS) [crafted]; Earth Warder's Gloves (21318, -1.74 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 21.9 spell_power points (9.98 DPS) | yes | Stormpike Cloth Girdle (19094, -1.14 DPS, sim-verified) [rep]; Ban'thok Sash (11662, -1.27 DPS) [dungeon]; Black Waistcord (254087, -1.32 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 47.4 spell_power points (21.58 DPS) | yes | Marshal's Dreadweave Leggings (231587, -2.80 DPS) [pvp]; Skyshroud Leggings (13170, -6.70 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -7.51 DPS) [pvp] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 22.8 spell_power points (10.39 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.06 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.21 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.18 DPS) [dungeon]; Blessed Band of Light (272407, -2.03 DPS) [vendor]; Naglering (11669, -10.12 DPS, sim-verified) [dungeon] |
| finger2 | Songstone of Ironforge (12543) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.04 DPS, sim-verified) [dungeon]; Blessed Band of Light (272407, -1.75 DPS) [vendor]; The Postmaster's Seal (13392, -2.13 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of the Beast (13968, -3.37 DPS, sim-verified) [quest]; Serenity Field (272439, -5.01 DPS) [vendor]; Uther's Strength (11302, -5.46 DPS) [world_drop] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (-4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, -0.04 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -0.27 DPS) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -0.43 DPS) [pvp]; Sageclaw (20070, -2.58 DPS) [rep]; Ironbark Staff (20069, -11.47 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.20 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.50 DPS) [pvp]; Tome of Shadow Force (19309, -0.57 DPS) [rep]; Spirit of Aquementas (11904, -3.64 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.51 DPS) [dungeon]; Oblivion's Touch (18761, -0.71 DPS) [dungeon]; Brilliant Wand (249385, -0.71 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Venomspew Footpads; finger1: Maiden's Circle; finger2: Songstone of Ironforge; trinket1: Weakness Analyzer; trinket2: Second Wind; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1458, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 05520020135201351-0050203001000000000-0500000000000000)

Set DPS (verified): 759.7. Weights run: 4.5s. Verify run: 23.0s. 1458 eligible items had no known source.

Race sweep: human 759.7 ± 0.5, gnome 758.0 ± 0.5 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.210 ± 0.010, crit=0.391 ± 0.006 per rating point (14 rating = 1%, 5.470 per %), hit=0.339 ± 0.019 per rating point (10 rating = 1%, 3.389 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, fire_power=not significant (0.000 ± 0.000), spirit=not significant (0.060 ± 0.040)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 37.2 spell_power points (22.07 DPS) | yes | Field Marshal's Coronal (231584, -0.08 DPS) [pvp]; Felcloth Hood (14111, -3.01 DPS) [crafted]; Crimson Felt Hat (18727, -6.84 DPS, sim-verified) [dungeon] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 21.7 spell_power points (12.87 DPS) | yes | Lady Maye's Pendant (14558, -4.28 DPS, sim-verified) [world_drop]; Archlight Talisman (15856, -4.50 DPS) [quest]; Nacreous Shell Necklace (22403, -4.52 DPS) [dungeon] |
| shoulder | Felcloth Shoulders (14112) | Tailoring [crafted] | 26.3 spell_power points (15.61 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.52 DPS) [crafted]; Mantle of the Timbermaw (19050, -0.66 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -3.23 DPS) [dungeon]; Shroud of the Exile (15421, -3.89 DPS) [quest]; Crystalline Threaded Cape (20697, -4.36 DPS, sim-verified) [world] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.9 spell_power points (28.43 DPS) | yes | Robe of Winter Night (14136, -3.96 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -6.44 DPS) [pvp]; Felcloth Robe (14106, -7.06 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.1 spell_power points (14.31 DPS) | yes | Netherpearl Cuffs (254067, -5.54 DPS) [crafted]; Runecloth Cuffs (254123, -6.53 DPS) [crafted]; Manacle Cuffs (11962, -8.28 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -0.02 DPS) [quest]; Hands of Power (13253, -3.50 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormpike Cloth Girdle (19094, -0.38 DPS) [rep]; Ban'thok Sash (11662, -1.93 DPS) [dungeon]; Black Waistcord (254087, -2.36 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.0 spell_power points (26.74 DPS) | yes | Marshal's Dreadweave Leggings (231587, -2.41 DPS) [pvp]; Skyshroud Leggings (13170, -7.43 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -8.50 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 22.5 spell_power points (13.37 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.02 DPS) [crafted]; Knight-Lieutenant's Dreadweave Boots (17562, -0.22 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.99 DPS) [dungeon]; Underworld Band (1980, -2.84 DPS) [world_drop]; Naglering (11669, -14.22 DPS, sim-verified) [dungeon] |
| finger2 | Songstone of Ironforge (12543) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.78 DPS) [dungeon]; Underworld Band (1980, -2.62 DPS) [world_drop]; The Postmaster's Seal (13392, -3.04 DPS) [dungeon] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+6.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, -6.15 DPS, sim-verified) [dungeon]; Serenity Field (272439, -6.53 DPS) [vendor]; Uther's Strength (11302, -7.12 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.16 DPS) [pvp]; Headmaster's Charge (13937, -14.21 DPS, sim-verified) [dungeon]; Trindlehaven Staff (13161, -15.94 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.62 DPS) [dungeon]; Oblivion's Touch (18761, -0.94 DPS) [dungeon]; Brilliant Wand (249385, -0.94 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Felcloth Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Songstone of Ironforge; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1458, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.6. Weights run: 4.1s. Verify run: 3.1s. 193 eligible items had no known source.

Race sweep: undead 61.3 ± 0.0, orc 60.6 ± 0.0, troll 59.8 ± 0.0 (dps); undead beats the guide race orc by 0.7 (1.21%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.159 ± 0.006, crit=0.059 ± 0.001 per rating point (14 rating = 1%, 0.829 per %), hit=0.161 ± 0.006 per rating point (10 rating = 1%, 1.613 per %), spell_haste=not significant (0.301 ± 0.081), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.834 ± 0.001, fire_power=0.166 ± 0.000, spirit=0.249 ± 0.053

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.2 spell_power points (2.77 DPS) | yes | Pristine Circlet (253949, -0.96 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -2.08 DPS) [crafted]; Flame Circlet (253953, -2.22 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.7 spell_power points (0.23 DPS) | yes | Scholarly Pendant (277203, -0.08 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.4 spell_power points (1.94 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (20427, -0.40 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.60 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.76 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.5 spell_power points (3.76 DPS) | yes | Filigreed Shadow Gown (253907, -1.05 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.79 DPS) [crafted]; Scholarly Robes (2034, -1.95 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.90 DPS) | yes | Owlbeard Bracers (16981, -0.57 DPS, sim-verified) [quest]; Windsong Bangles (263336, -0.60 DPS) [quest]; Mindthrust Bracers (1974, -0.67 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.11 DPS) | yes | Shadow Gloves (253919, -0.23 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.30 DPS) [world]; Blight Gloves (279877, -0.60 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.5 spell_power points (1.95 DPS) | yes | Pristine Sash (253925, -0.56 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.60 DPS) [crafted]; Novice Ardent's Sash (253887, -0.98 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 11.1 spell_power points (3.35 DPS) | yes | Filigreed Pristine Leggings (253937, -0.96 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.24 DPS) [dungeon]; Colorful Kilt (10048, -1.47 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (2.30 DPS) | yes | Shadow Boots (253895, -0.58 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.72 DPS) [crafted]; Pristine Boots (253889, -1.25 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.51 DPS) | yes | Furlfeather Ring (286745, -1.21 DPS) [world]; Lavishly Jeweled Ring (1156, -1.22 DPS) [dungeon]; Loop of Sacrifice (281673, -1.27 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 1.5 spell_power points (0.45 DPS) | yes | Furlfeather Ring (286745, -0.15 DPS) [world]; Lavishly Jeweled Ring (1156, -0.16 DPS) [dungeon]; Loop of Sacrifice (281673, -0.21 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 30.5 spell_power points (9.19 DPS) | yes | Evocator's Blade (2567, -0.28 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.60 DPS) [quest]; Crescent Staff (6505, -1.02 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.30 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Staff of the Blessed Seer; ranged: Sizzle Stick

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade

### Band 30 (orc, 25520000132001000-0000000000000000000-0000000000000000)

Set DPS (verified): 104.0. Weights run: 4.3s. Verify run: 9.1s. 319 eligible items had no known source.

Race sweep: undead 104.3 ± 0.1, orc 104.0 ± 0.1, troll 102.4 ± 0.1 (dps); undead beats the guide race orc by 0.3 (0.31%) ± 0.1, beyond error.

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
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Solid Gold Band (279845, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Ring of Power Regulation (273030, -0.99 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.11 DPS) [dungeon]; Spine of the Devourer (273040, -1.11 DPS) [world]; Tendonscraper Dagger (273053, -1.85 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.50 DPS) | yes | Orb of Mystic Insight (249394, -0.56 DPS) [crafted]; Orb of Mistmantle (13031, -1.14 DPS) [world_drop]; Orb of Souls (249395, -1.39 DPS) [crafted] |
| ranged | Gravestone Scepter (7001) | Blackfathom Villainy [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Wand of Eventide (5214, -0.08 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Gravestone Scepter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (orc, 25521000132001351-0000000000000000000-0000000000000000)

Set DPS (verified): 180.7. Weights run: 4.8s. Verify run: 4.0s. 448 eligible items had no known source.

Race sweep: undead 180.8 ± 0.1, orc 180.7 ± 0.1, troll 178.3 ± 0.1 (dps); undead beats the guide race orc by 0.2 (0.08%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.165 ± 0.010, crit=0.245 ± 0.003 per rating point (14 rating = 1%, 3.426 per %), hit=0.298 ± 0.014 per rating point (10 rating = 1%, 2.976 per %), spell_haste=not significant (0.346 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.885 ± 0.001, fire_power=0.115 ± 0.000, spirit=not significant (0.063 ± 0.092)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.67 DPS) | yes | Miner's Hat of the Deep (9429, -2.66 DPS, sim-verified) [dungeon]; Living Cowl (5608, -3.30 DPS) [world]; Augural Shroud (2620, -3.45 DPS) [world] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 9.0 spell_power points (3.72 DPS) | yes | Triune Amulet (7722, -0.23 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.28 DPS) [dungeon]; Necklace of Calisea (1714, -3.06 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (4.33 DPS) | yes | Green Silken Shoulders (7057, -0.28 DPS) [crafted]; Inquisitor's Shawl (19507, -0.55 DPS) [dungeon]; Berylline Pads (4197, -0.76 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.23 DPS) [crafted]; Battle Healer's Cloak (19528, -0.44 DPS) [rep]; Long Silken Cloak (4326, -0.60 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.3 spell_power points (9.63 DPS) | yes | Dreamweave Vest (10021, -1.21 DPS) [crafted]; Elemental Raiment (9434, -1.50 DPS, sim-verified) [world_drop]; Shadoweave Robe (10004, -1.76 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.72 DPS) | yes | Radiant Silver Bracers (4545, -1.52 DPS) [quest]; Condor Bracers (15864, -1.68 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.93 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.1 spell_power points (7.89 DPS) | yes | Black Mageweave Gloves (10003, -1.69 DPS) [crafted]; Shadoweave Gloves (10023, -1.78 DPS, sim-verified) [crafted]; Black Handwraps (254027, -2.29 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.7 spell_power points (6.06 DPS) | yes | Black Cord (254043, -0.23 DPS) [crafted]; Star Belt (4329, -0.58 DPS) [crafted]; Warsong Sash (16975, -1.51 DPS) [quest] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 18.6 spell_power points (7.68 DPS) | yes | Red Mageweave Pants (10009, -1.29 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -1.95 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.49 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 12.9 spell_power points (5.34 DPS) | yes | Gilded Slippers (254001, -1.87 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.17 DPS) [crafted]; Glinteye Slippers (273024, -2.29 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.2 spell_power points (5.46 DPS) | yes | Advisor's Ring (19520, -1.74 DPS) [rep]; Snake Hoop (6750, -2.32 DPS) [quest]; Reedknot Ring (9622, -2.57 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (4.47 DPS) | yes | Advisor's Ring (19520, -0.59 DPS, sim-verified) [rep]; Snake Hoop (6750, -1.33 DPS) [quest]; Reedknot Ring (9622, -1.58 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -7.95 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -9.45 DPS) [dungeon]; Swinetusk Shank (6691, -11.62 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.21 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Prodigious Shadowshard Pendant; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Shadoweave Pants; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 25522300132001351-0005000000000000000-0010000000000000)

Set DPS (verified): 262.7. Weights run: 4.8s. Verify run: 4.5s. 585 eligible items had no known source.

Race sweep: undead 263.2 ± 0.1, orc 262.7 ± 0.1, troll 259.6 ± 0.1 (dps); undead beats the guide race orc by 0.5 (0.18%) ± 0.2, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.185 ± 0.012, crit=0.322 ± 0.004 per rating point (14 rating = 1%, 4.505 per %), hit=0.420 ± 0.020 per rating point (10 rating = 1%, 4.203 per %), spell_haste=0.731 ± 0.097, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.888 ± 0.001, fire_power=0.112 ± 0.000, spirit=not significant (0.139 ± 0.097)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.64 DPS) | yes | Dreamweave Circlet (10041, -1.02 DPS, sim-verified) [crafted]; Shadoweave Mask (10025, -1.66 DPS) [crafted]; Red Mageweave Headband (10033, -1.85 DPS) [crafted] |
| neck | Jarkal's Enhancing Necklace (7888) | Necklace Recovery, Take 3 [quest] | 12.2 spell_power points (5.25 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.27 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.38 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.60 DPS) [dungeon] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | 17.4 spell_power points (7.50 DPS) | yes | Kentic Amice (11624, -0.07 DPS) [dungeon]; Rotgrip Mantle (17732, -0.46 DPS) [dungeon]; Shadoweave Shoulders (10028, -1.00 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.72 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.36 DPS) [crafted]; Nightfall Drape (12465, -2.99 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (10.26 DPS) | yes | Elemental Raiment (9434, -1.21 DPS) [world_drop]; Shadoweave Robe (10004, -1.81 DPS) [crafted]; Dreamweave Vest (10021, -1.82 DPS, sim-verified) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.5 spell_power points (5.83 DPS) | yes | Nethergeld Cuffs (254061, -1.96 DPS) [crafted]; Spidertank Oilrag (9448, -1.99 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -2.78 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 19.8 spell_power points (8.55 DPS) | yes | Dreamweave Gloves (10019, -0.05 DPS) [crafted]; Shadoweave Gloves (10023, -1.64 DPS) [crafted]; Black Mageweave Gloves (10003, -2.08 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.5 spell_power points (7.99 DPS) | yes | Ban'thok Sash (11662, -0.21 DPS) [dungeon]; Satyrmane Sash (17755, -1.16 DPS) [dungeon]; Ghostweave Cord (254073, -1.29 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.3 spell_power points (10.91 DPS) | yes | Spellshock Leggings (9484, -0.68 DPS, sim-verified) [dungeon]; Wizardweave Leggings (14132, -2.72 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.83 DPS) [vendor] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 19.4 spell_power points (8.37 DPS) | yes | Shadoweave Boots (10031, -2.41 DPS) [crafted]; Black Slippers (254007, -2.60 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -2.93 DPS, sim-verified) [vendor] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 13.5 spell_power points (5.84 DPS) | yes | Band of the Unicorn (7553, -0.24 DPS) [world_drop]; Advisor's Ring (19519, -0.67 DPS) [rep]; Philanthropist's Ring (281635, -1.13 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 13.4 spell_power points (5.76 DPS) | yes | Band of the Unicorn (7553, -0.15 DPS) [world_drop]; Advisor's Ring (19519, -0.58 DPS) [rep]; Philanthropist's Ring (281635, -1.05 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -0.91 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.03 DPS) [quest]; Rune of the Guard Captain (19120, -1.23 DPS, sim-verified) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.32 DPS) [quest]; Ankh of Life (1713, -1.87 DPS) [world_drop]; Alchemist's Stone (13503, -2.76 DPS, sim-verified) [crafted] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS) [dungeon]; Staff of Jordan (873, -3.44 DPS, sim-verified) [world_drop] |
| off_hand | Umbral Crystal (13029) | World drop [world_drop] | 15.1 spell_power points (6.51 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.08 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.47 DPS) [dungeon]; Orb of Dar'Orahil (15108, -0.75 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.43 DPS) [crafted]; Thunderwood (13062, -0.86 DPS) [world_drop]; Twisted Nether Wand (249144, -0.86 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Jarkal's Enhancing Necklace; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; feet: Black Sandals; finger1: Brainlash; finger2: Underworld Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Umbral Crystal; ranged: Lesser Eternal Wand

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522300132001351-2005220000000000000-0050000000000000)

Set DPS (verified): 403.1. Weights run: 5.0s. Verify run: 22.2s. 1452 eligible items had no known source.

Race sweep: undead 403.2 ± 0.2, orc 403.1 ± 0.2, troll 398.2 ± 0.2 (dps); undead beats the guide race orc by 0.1 (0.02%) ± 0.3, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.241 ± 0.012, crit=0.235 ± 0.003 per rating point (14 rating = 1%, 3.286 per %), hit=0.309 ± 0.016 per rating point (10 rating = 1%, 3.093 per %), spell_haste=1.785 ± 0.164, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.893 ± 0.001, fire_power=0.107 ± 0.000, spirit=not significant (0.207 ± 0.152)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 38.8 spell_power points (16.60 DPS) | yes | Warlord's Dreadweave Hood (231590, -0.43 DPS) [pvp]; Crimson Felt Hat (18727, -3.51 DPS, sim-verified) [dungeon]; Felcloth Hood (14111, -4.10 DPS) [crafted] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 19.8 spell_power points (8.47 DPS) | yes | Archlight Talisman (15856, -2.30 DPS) [quest]; Lady Maye's Pendant (14558, -2.43 DPS, sim-verified) [world_drop]; Nacreous Shell Necklace (22403, -2.47 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 26.5 spell_power points (11.32 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Burial Shawl (18681, -1.11 DPS) [dungeon]; Felcloth Shoulders (14112, -1.42 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -0.84 DPS, sim-verified) [world]; Spritecaster Cape (11623, -1.85 DPS) [dungeon]; Shroud of the Exile (15421, -2.21 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (20.61 DPS) | yes | Robe of Winter Night (14136, -3.84 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -4.44 DPS) [pvp]; Felcloth Robe (14106, -6.85 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.4 spell_power points (10.86 DPS) | yes | Netherpearl Cuffs (254067, -4.73 DPS) [crafted]; Runecloth Cuffs (254123, -5.12 DPS) [crafted]; Manacle Cuffs (11962, -6.32 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.8 spell_power points (11.48 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -2.07 DPS) [quest]; Blood Guard's Dreadweave Handwraps (227099, -2.08 DPS) [pvp] |
| waist | Frostwolf Cloth Belt (19090) | Frostwolf Clan [rep] | 20.4 spell_power points (8.73 DPS) | yes | Black Waistcord (254087, -0.54 DPS) [crafted]; Defiler's Cloth Girdle (20163, -0.72 DPS) [rep]; Ban'thok Sash (11662, -1.24 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 34.7 spell_power points (14.83 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.75 DPS) [dungeon]; Legionnaire's Dreadweave Legguards (227097, -1.51 DPS) [pvp] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 23.7 spell_power points (10.12 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.33 DPS) [dungeon]; Dragonrider Boots (18102, -0.77 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Postmaster's Seal (13392, -2.02 DPS) [dungeon]; Blessed Band of Light (272407, -2.64 DPS) [vendor]; Naglering (11669, -8.54 DPS, sim-verified) [dungeon] |
| finger2 | Eye of Orgrimmar (12545) | The Princess Saved? [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blessed Band of Light (272407, -2.11 DPS) [vendor]; Brainlash (6440, -2.24 DPS) [dungeon]; The Postmaster's Seal (13392, -2.89 DPS, sim-verified) [dungeon] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.03 DPS, sim-verified) [vendor]; Uther's Strength (11302, -5.13 DPS) [world_drop]; Second Wind (11819, -5.56 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, -1.88 DPS) [dungeon]; High Warlord's War Staff (234549, -2.32 DPS) [pvp]; Ironbark Staff (20220, -11.88 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (12.41 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.37 DPS) [pvp]; Tome of Shadow Force (19309, -0.56 DPS) [rep]; Spirit of Aquementas (11904, -3.42 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.52 DPS) [dungeon]; Oblivion's Touch (18761, -0.65 DPS) [dungeon]; Brilliant Wand (249385, -0.65 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Frostwolf Cloth Belt; legs: Outrider's Silk Leggings; feet: Venomspew Footpads; finger1: Maiden's Circle; finger2: Eye of Orgrimmar; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1452, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 05520020135201351-0050203001000000000-0500000000000000)

Set DPS (verified): 750.3. Weights run: 4.5s. Verify run: 22.4s. 1452 eligible items had no known source.

Race sweep: undead 751.2 ± 0.5, orc 750.3 ± 0.5, troll 742.2 ± 0.5 (dps); undead beats the guide race orc by 0.8 (0.11%) ± 0.7, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.334 ± 0.016, crit=0.702 ± 0.012 per rating point (14 rating = 1%, 9.832 per %), hit=0.646 ± 0.035 per rating point (10 rating = 1%, 6.460 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, fire_power=not significant (0.000 ± 0.000), spirit=not significant (0.042 ± 0.040)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 40.0 spell_power points (25.84 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -2.05 DPS) [pvp]; Crimson Felt Hat (18727, -5.19 DPS, sim-verified) [dungeon] |
| neck | Dark Advisor's Pendant (18691) | Scholomance: Vectus [dungeon] | 22.7 spell_power points (14.65 DPS) | yes | Nacreous Shell Necklace (22403, -3.20 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -3.44 DPS) [world_drop]; Emberfury Talisman (12929, -4.42 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.2 spell_power points (20.14 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.32 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -3.23 DPS) [pvp]; Felcloth Shoulders (14112, -3.48 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.05 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.73 DPS) [dungeon]; Shroud of the Exile (15421, -6.14 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.0 spell_power points (31.66 DPS) | yes | Robe of Winter Night (14136, -3.99 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -5.81 DPS) [pvp]; Felcloth Robe (14106, -8.40 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.0 spell_power points (16.13 DPS) | yes | Netherpearl Cuffs (254067, -6.08 DPS) [crafted]; Runecloth Cuffs (254123, -6.87 DPS) [crafted]; Manacle Cuffs (11962, -8.37 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 32.0 spell_power points (20.67 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -2.62 DPS, sim-verified) [dungeon]; Black Gloves (254101, -6.03 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 25.8 spell_power points (16.69 DPS) | yes | Black Waistcord (254087, -2.57 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -2.61 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -2.90 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 34.8 spell_power points (22.46 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.49 DPS) [dungeon]; Legionnaire's Dreadweave Legguards (227097, -1.57 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.0 spell_power points (15.51 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -0.43 DPS) [dungeon]; Black Sandals (254113, -0.65 DPS) [crafted] |
| finger1 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.32 DPS) [quest]; Blessed Band of Light (272407, -2.61 DPS) [vendor]; Naglering (11669, -13.42 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.16 DPS) [quest]; Blessed Band of Light (272407, -2.45 DPS) [vendor]; Don Julio's Band (19325, -2.78 DPS) [rep] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+5.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Second Wind (11819, -5.66 DPS, sim-verified) [dungeon]; Serenity Field (272439, -7.11 DPS) [vendor]; Smolderweb's Eye (13213, -7.46 DPS) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.80 DPS) [pvp]; Headmaster's Charge (13937, -13.84 DPS, sim-verified) [dungeon]; Whiteout Staff (19101, -21.53 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Oblivion's Touch (18761, -0.86 DPS) [dungeon]; Brilliant Wand (249385, -0.86 DPS) [crafted]; Ritssyn's Wand of Bad Mojo (22408, -1.08 DPS) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Dark Advisor's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Don Mauricio's Band of Domination; finger2: Maiden's Circle; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1452, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

