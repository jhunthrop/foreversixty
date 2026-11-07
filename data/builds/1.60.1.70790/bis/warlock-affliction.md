# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 61.7. Weights run: 4.0s. Verify run: 4.4s. 182 eligible items had no known source.

Race sweep: human 61.7 ± 0.0, gnome 61.4 ± 0.0 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.184 ± 0.006, crit=0.059 ± 0.001 per rating point (14 rating = 1%, 0.829 per %), hit=0.161 ± 0.006 per rating point (10 rating = 1%, 1.613 per %), spell_haste=not significant (0.301 ± 0.081), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.834 ± 0.001, fire_power=0.166 ± 0.000, spirit=0.249 ± 0.053

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.2 spell_power points (2.77 DPS) | yes | Pristine Circlet (253949, -0.96 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -2.04 DPS) [crafted]; Flame Circlet (253953, -2.22 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.5 spell_power points (0.15 DPS) | yes | Tarnished Locket (279870, -0.32 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.7 spell_power points (2.01 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.28 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.80 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Black Whelp Cloak (7283, -0.60 DPS) [crafted]; Caretaker's Cape (20428, -0.62 DPS, sim-verified) [rep]; Pearl-clasped Cloak (5542, -0.74 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.6 spell_power points (3.80 DPS) | yes | Filigreed Shadow Gown (253907, -0.95 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.79 DPS) [crafted]; Scholarly Robes (2034, -1.99 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Repurposed Hair Band (281256, -0.02 DPS) [quest]; Bright Bracers (3647, -0.06 DPS) [world_drop]; Windsong Bangles (263336, -0.17 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.11 DPS) | yes | Shadow Gloves (253919, -0.24 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.30 DPS) [world]; Pristine Gloves (253913, -0.74 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.6 spell_power points (1.98 DPS) | yes | Pristine Sash (253925, -0.55 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.61 DPS) [crafted]; Novice Ardent's Sash (253887, -0.99 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 11.3 spell_power points (3.40 DPS) | yes | Filigreed Pristine Leggings (253937, -0.96 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.29 DPS) [dungeon]; Colorful Kilt (10048, -1.52 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.7 spell_power points (2.33 DPS) | yes | Red Woolen Boots (4313, -0.75 DPS) [crafted]; Shadow Boots (253895, -0.81 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.26 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.4 spell_power points (1.62 DPS) | yes | Band of Purification (12996, -1.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.29 DPS) [dungeon]; Furlfeather Ring (286745, -1.32 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.51 DPS) | yes | Lavishly Jeweled Ring (1156, -1.17 DPS) [dungeon]; Furlfeather Ring (286745, -1.21 DPS) [world]; Band of Purification (12996, -1.53 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.64 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.69 DPS) [dungeon]; Witching Stave (1484, -1.86 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.2 spell_power points (0.98 DPS) | yes | Onu's Patience (281292, -0.60 DPS) [quest]; Grayson's Torch (1172, -0.68 DPS) [quest]; Antipodean Rod (2879, -0.68 DPS) [world_drop] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.30 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 182, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 25520000132001000-0000000000000000000-0000000000000000)

Set DPS (verified): 102.8. Weights run: 4.1s. Verify run: 7.2s. 309 eligible items had no known source.

Race sweep: human 102.8 ± 0.1, gnome 101.9 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.184 ± 0.005, crit=0.077 ± 0.001 per rating point (14 rating = 1%, 1.080 per %), hit=0.145 ± 0.006 per rating point (10 rating = 1%, 1.446 per %), spell_haste=0.345 ± 0.081, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.143 ± 0.091)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Circlet (253981, -0.13 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.39 DPS) | yes | Tarnished Locket (279870, -1.57 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Invoker's Mantle (215365, -0.56 DPS) [crafted]; Death Speaker Mantle (6685, -0.83 DPS) [dungeon]; Moonlit Amice (11884, -2.14 DPS, sim-verified) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Caretaker's Cape (19533, -0.28 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.62 DPS) | yes | Necrology Robes (2292, -0.44 DPS) [dungeon]; Shadow Gown (253967, -0.66 DPS) [crafted]; Green Silk Armor (7065, -1.11 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Windsong Bangles (263336, -2.23 DPS) [quest]; Phoenix Bindings (210781, -2.95 DPS, sim-verified) [crafted] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent Gloves (5970, -0.10 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world]; Hands of Darkness (7047, -0.46 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.0 spell_power points (3.06 DPS) | yes | Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.83 DPS) [dungeon]; Invoker's Cord (215366, -1.11 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gaze Dreamer Pants (6903, -0.49 DPS, sim-verified) [dungeon]; Filigreed Shadow Leggings (253943, -0.68 DPS) [crafted]; Necromancer Leggings (2277, -0.91 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.95 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.01 DPS) [crafted]; Glinteye Slippers (273024, -2.48 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.95 DPS) | yes | Minor Channeling Ring (1449, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon]; Ring of Power Regulation (273030, -2.56 DPS, sim-verified) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Minor Channeling Ring (1449, -0.28 DPS) [quest]; Ring of Power Regulation (273030, -0.42 DPS, sim-verified) [dungeon]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.11 DPS) [dungeon]; Spine of the Devourer (273040, -1.11 DPS) [world]; Tendonscraper Dagger (273053, -1.98 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.50 DPS) | yes | Orb of Mystic Insight (249394, -0.41 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.14 DPS) [world_drop]; Eye of Paleth (2943, -1.39 DPS) [quest] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.27 DPS) [crafted]; Wand of Eventide (5214, -0.28 DPS) [world_drop]; Gravestone Scepter (7001, -0.30 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 25521000132001351-0000000000000000000-0000000000000000)

Set DPS (verified): 199.1. Weights run: 4.2s. Verify run: 6.5s. 399 eligible items had no known source.

Race sweep: human 199.1 ± 0.1, gnome 197.6 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.163 ± 0.005, crit=0.101 ± 0.002 per rating point (14 rating = 1%, 1.416 per %), hit=0.125 ± 0.006 per rating point (10 rating = 1%, 1.245 per %), spell_haste=0.279 ± 0.061, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.884 ± 0.001, fire_power=0.116 ± 0.000, spirit=not significant (0.023 ± 0.075)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.75 DPS) | yes | Augural Shroud (2620, -3.49 DPS) [world]; Filigreed Shadow Circlet (253981, -3.59 DPS) [crafted]; Living Cowl (5608, -3.66 DPS, sim-verified) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Eye of Mordresh (10769, -0.42 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.72 DPS) [quest]; Necklace of Calisea (1714, -2.86 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (4.36 DPS) | yes | Green Silken Shoulders (7057, -0.28 DPS) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Berylline Pads (4197, -0.77 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.13 DPS) [crafted]; Guardian Cloak (5965, -0.13 DPS) [crafted]; Caretaker's Cape (19532, -0.43 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (9.63 DPS) | yes | Elemental Raiment (9434, -0.80 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.38 DPS) [crafted]; Shadoweave Robe (10004, -1.82 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Arcane Runed Bracers (4744, -0.39 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.83 DPS) [quest]; Earthen Silk Cuffs (254019, -2.04 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (7.84 DPS) | yes | Shadoweave Gloves (10023, -1.43 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.59 DPS) [crafted]; Black Handwraps (254027, -2.20 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Highlander's Cloth Girdle (20098, -0.26 DPS, sim-verified) [rep]; Star Belt (4329, -0.30 DPS) [crafted] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 18.6 spell_power points (7.74 DPS) | yes | Red Mageweave Pants (10009, -0.64 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -2.06 DPS) [crafted]; Gaze Dreamer Pants (6903, -2.65 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.2 spell_power points (10.07 DPS) | yes | Black Slippers (254007, -4.93 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -6.64 DPS) [crafted]; Spidersilk Boots (4320, -6.88 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.2 spell_power points (5.50 DPS) | yes | Lorekeeper's Ring (19524, -1.75 DPS) [rep]; Ring of Forlorn Spirits (2043, -2.17 DPS) [quest]; Reedknot Ring (9622, -2.58 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.8 spell_power points (4.51 DPS) | yes | Lorekeeper's Ring (19524, -0.76 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.17 DPS) [quest]; Reedknot Ring (9622, -1.59 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -8.33 DPS) [dungeon]; Illusionary Rod (7713, -8.54 DPS, sim-verified) [dungeon]; Swinetusk Shank (6691, -10.53 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.26 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Black Cord; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 399, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 25522300132001351-0005000000000000000-0010000000000000)

Set DPS (verified): 282.6. Weights run: 4.3s. Verify run: 7.4s. 506 eligible items had no known source.

Race sweep: human 282.6 ± 0.2, gnome 281.3 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.204 ± 0.013, crit=0.320 ± 0.005 per rating point (14 rating = 1%, 4.481 per %), hit=0.401 ± 0.019 per rating point (10 rating = 1%, 4.011 per %), spell_haste=0.627 ± 0.106, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.900 ± 0.001, fire_power=0.100 ± 0.000, spirit=not significant (0.075 ± 0.086)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.89 DPS) | yes | Dreamweave Circlet (10041, -1.46 DPS) [crafted]; Shadoweave Mask (10025, -1.60 DPS) [crafted]; Red Mageweave Headband (10033, -1.87 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.6 spell_power points (4.12 DPS) | yes | Triune Amulet (7722, -0.05 DPS) [dungeon]; Mindburst Medallion (11196, -0.49 DPS) [quest]; Horizon Choker (13085, -2.62 DPS) [world_drop] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | 17.4 spell_power points (8.30 DPS) | yes | Kentic Amice (11624, -0.14 DPS) [dungeon]; Rotgrip Mantle (17732, -0.34 DPS) [dungeon]; Shadoweave Shoulders (10028, -0.99 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.75 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -3.19 DPS) [dungeon]; Caretaker's Cape (19531, -3.93 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.6 spell_power points (11.27 DPS) | yes | Elemental Raiment (9434, -1.24 DPS) [world_drop]; Acumen Robes (17775, -1.28 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -1.30 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.5 spell_power points (6.45 DPS) | yes | Spidertank Oilrag (9448, -2.15 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.25 DPS) [crafted]; Arcane Runed Bracers (4744, -2.93 DPS, sim-verified) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 19.7 spell_power points (9.43 DPS) | yes | Dreamweave Gloves (10019, -1.18 DPS, sim-verified) [crafted]; Shadoweave Gloves (10023, -1.63 DPS) [crafted]; Black Mageweave Gloves (10003, -2.26 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.9 spell_power points (9.04 DPS) | yes | Ban'thok Sash (11662, -1.02 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -1.38 DPS) [dungeon]; Ghostweave Cord (254073, -1.96 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.8 spell_power points (12.34 DPS) | yes | Spellshock Leggings (9484, -0.39 DPS) [dungeon]; Wizardweave Leggings (14132, -3.27 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.31 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.5 spell_power points (11.71 DPS) | yes | Black Sandals (254113, -2.38 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Boots (220891, -5.10 DPS) [vendor]; Black Slippers (254007, -5.30 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.6 spell_power points (6.50 DPS) | yes | Lorekeeper's Ring (19523, -0.77 DPS) [rep]; Philanthropist's Ring (281635, -1.24 DPS) [quest]; Cyclopean Band (11824, -1.38 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.21 DPS) | yes | Lorekeeper's Ring (19523, -0.48 DPS) [rep]; Philanthropist's Ring (281635, -0.95 DPS) [quest]; Cyclopean Band (11824, -1.09 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.37 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.44 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Thunderbrew's Boot Flask (744, -0.48 DPS, sim-verified) [quest]; Ankh of Life (1713, -2.44 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.51 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -2.01 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -3.04 DPS) [dungeon]; Spire of Hakkar (10844, -3.65 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.48 DPS) [crafted]; Burning Sliver (5249, -0.96 DPS) [quest]; Twisted Nether Wand (249144, -0.96 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Glowing Eye of Mordresh; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Lesser Eternal Wand

No-known-source sample (15 of 506, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522300132001351-2005220000000000000-0050000000000000)

Set DPS (verified): 455.9. Weights run: 4.2s. Verify run: 22.2s. 1299 eligible items had no known source.

Race sweep: human 455.9 ± 0.3, gnome 453.3 ± 0.3 (dps); the guide race human is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.241 ± 0.038, crit=0.470 ± 0.013 per rating point (14 rating = 1%, 6.578 per %), hit=0.515 ± 0.049 per rating point (10 rating = 1%, 5.145 per %), spell_haste=not significant (-0.211 ± 0.426), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.898 ± 0.001, fire_power=0.102 ± 0.000, spirit=not significant (0.582 ± 0.228)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 36.6 spell_power points (18.86 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -2.41 DPS) [pvp]; Ghostweave Hood (254135, -4.97 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Triune Amulet (7722, -1.20 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 29.1 spell_power points (14.99 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.22 DPS) [crafted]; Felcloth Shoulders (14112, -1.45 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fel Cape (279269, -1.56 DPS) [crafted]; Crystalline Threaded Cape (20697, -1.80 DPS, sim-verified) [world]; Spritecaster Cape (11623, -2.13 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (24.83 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.35 DPS) [pvp]; Alanna's Embrace (13314, -6.03 DPS) [dungeon]; Robe of Winter Night (14136, -6.18 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 28.0 spell_power points (14.43 DPS) | yes | Netherpearl Cuffs (254067, -6.05 DPS) [crafted]; Earthenweave Cuffs (254131, -7.31 DPS) [crafted]; Runecloth Cuffs (254123, -7.52 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Black Gloves (254101, -0.22 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 22.0 spell_power points (11.35 DPS) | yes | Ghostweave Cord (254073, -0.83 DPS) [crafted]; Thuzadin Sash (18740, -0.89 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -2.16 DPS, sim-verified) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 47.6 spell_power points (24.52 DPS) | yes | Marshal's Dreadweave Leggings (231587, -3.09 DPS) [pvp]; Ghostweave Pants (14144, -5.11 DPS) [crafted]; Skyshroud Leggings (13170, -6.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.1 spell_power points (14.47 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Maleki's Footwraps (18735, -0.85 DPS) [dungeon]; Venomspew Footpads (275606, -2.11 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.60 DPS) [quest]; Don Mauricio's Band of Domination (22433, -2.89 DPS) [dungeon]; Naglering (11669, -10.87 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.15 DPS) [dungeon]; Deep Rooted Ring (19109, -2.78 DPS) [world] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -2.58 DPS) [world_drop]; Uther's Strength (11302, -3.09 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.18 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.20 DPS, sim-verified) [vendor]; Ankh of Life (1713, -7.74 DPS) [world_drop]; Uther's Strength (11302, -8.25 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -0.46 DPS) [pvp]; Sageclaw (20070, -3.10 DPS) [rep]; Ironbark Staff (20069, -15.42 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (14.95 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.45 DPS) [pvp]; Tome of Shadow Force (19309, -0.59 DPS) [rep]; Spirit of Aquementas (11904, -4.12 DPS) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.23 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.72 DPS) [dungeon]; Brilliant Wand (249385, -0.89 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 05520020135201351-0050203001000000000-0500000000000000)

Set DPS (verified): 840.4. Weights run: 3.3s. Verify run: 20.1s. 1299 eligible items had no known source.

Race sweep: human 840.4 ± 0.8, gnome 836.8 ± 0.7 (dps); the guide race human is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.225 ± 0.013, crit=0.712 ± 0.018 per rating point (14 rating = 1%, 9.973 per %), hit=0.440 ± 0.061 per rating point (10 rating = 1%, 4.398 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, fire_power=not significant (0.000 ± 0.000), spirit=not significant (0.014 ± 0.013)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 32.3 spell_power points (23.45 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Crimson Felt Hat (18727, -0.25 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Anastari Heirloom (18728, -0.91 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.9 spell_power points (21.74 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.78 DPS) [pvp]; Felcloth Shoulders (14112, -2.79 DPS) [crafted]; Argent Shoulders (19059, -3.49 DPS) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arcanoweave Cloak (272411, -2.02 DPS) [vendor]; Crystalline Threaded Cape (20697, -2.96 DPS) [world]; Spritecaster Cape (11623, -6.94 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.0 spell_power points (34.92 DPS) | yes | Robe of Winter Night (14136, -4.05 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -7.72 DPS) [pvp]; Felcloth Robe (14106, -8.75 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.9 spell_power points (17.38 DPS) | yes | Netherpearl Cuffs (254067, -6.73 DPS) [crafted]; Runecloth Cuffs (254123, -7.74 DPS) [crafted]; Arena Wristguards (18709, -10.07 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.2 spell_power points (22.69 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Hands of Power (13253, -2.74 DPS) [dungeon]; Earth Warder's Gloves (21318, -6.69 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 25.3 spell_power points (18.41 DPS) | yes | Black Waistcord (254087, -3.12 DPS) [crafted]; Stormpike Cloth Girdle (19094, -3.69 DPS) [rep]; Ban'thok Sash (11662, -4.69 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.2 spell_power points (39.41 DPS) | yes | Marshal's Dreadweave Leggings (231587, -9.39 DPS) [pvp]; Skyshroud Leggings (13170, -12.65 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -16.92 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maleki's Footwraps (18735, +0.00 DPS) [dungeon]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.01 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -0.45 DPS) [vendor] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Songstone of Ironforge (12543, -0.02 DPS) [quest]; Blessed Band of Light (272407, -3.75 DPS) [vendor] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -6.15 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -7.27 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -13.12 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.61 DPS, sim-verified) [crafted]; Ankh of Life (1713, -4.24 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.78 DPS) [pvp]; Kindling Stave (11750, -25.83 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -29.15 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Greater Eternal Wand (249237, -1.15 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger2: Maiden's Circle; trinket1: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.7. Weights run: 4.0s. Verify run: 3.6s. 176 eligible items had no known source.

Race sweep: undead 61.3 ± 0.0, orc 60.7 ± 0.0, troll 59.8 ± 0.0 (dps); undead beats the guide race orc by 0.7 (1.08%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.184 ± 0.006, crit=0.059 ± 0.001 per rating point (14 rating = 1%, 0.829 per %), hit=0.161 ± 0.006 per rating point (10 rating = 1%, 1.613 per %), spell_haste=not significant (0.301 ± 0.081), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.834 ± 0.001, fire_power=0.166 ± 0.000, spirit=0.249 ± 0.053

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.2 spell_power points (2.77 DPS) | yes | Pristine Circlet (253949, -0.96 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -2.04 DPS) [crafted]; Flame Circlet (253953, -2.22 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.7 spell_power points (0.23 DPS) | yes | Scholarly Pendant (277203, -0.08 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.7 spell_power points (2.01 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.20 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.80 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (20427, -0.40 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.60 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.74 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.6 spell_power points (3.80 DPS) | yes | Filigreed Shadow Gown (253907, -1.05 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.79 DPS) [crafted]; Scholarly Robes (2034, -1.99 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.90 DPS) | yes | Owlbeard Bracers (16981, -0.57 DPS, sim-verified) [quest]; Windsong Bangles (263336, -0.60 DPS) [quest]; Mindthrust Bracers (1974, -0.63 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.11 DPS) | yes | Shadow Gloves (253919, -0.23 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.30 DPS) [world]; Blight Gloves (279877, -0.60 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.6 spell_power points (1.98 DPS) | yes | Pristine Sash (253925, -0.56 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.61 DPS) [crafted]; Novice Ardent's Sash (253887, -0.99 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 11.3 spell_power points (3.40 DPS) | yes | Filigreed Pristine Leggings (253937, -0.96 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.29 DPS) [dungeon]; Colorful Kilt (10048, -1.52 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.7 spell_power points (2.33 DPS) | yes | Shadow Boots (253895, -0.58 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.75 DPS) [crafted]; Pristine Boots (253889, -1.26 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.51 DPS) | yes | Lavishly Jeweled Ring (1156, -1.17 DPS) [dungeon]; Furlfeather Ring (286745, -1.21 DPS) [world]; Loop of Sacrifice (281673, -1.23 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 1.5 spell_power points (0.45 DPS) | yes | Lavishly Jeweled Ring (1156, -0.06 DPS, sim-verified) [dungeon]; Furlfeather Ring (286745, -0.15 DPS) [world]; Loop of Sacrifice (281673, -0.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 30.5 spell_power points (9.19 DPS) | yes | Evocator's Blade (2567, -0.27 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.60 DPS) [quest]; Crescent Staff (6505, -0.97 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.30 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Staff of the Blessed Seer; ranged: Sizzle Stick

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (orc, 25520000132001000-0000000000000000000-0000000000000000)

Set DPS (verified): 103.0. Weights run: 4.1s. Verify run: 9.2s. 302 eligible items had no known source.

Race sweep: undead 103.3 ± 0.1, orc 103.0 ± 0.1, troll 101.5 ± 0.1 (dps); undead beats the guide race orc by 0.3 (0.29%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.184 ± 0.005, crit=0.077 ± 0.001 per rating point (14 rating = 1%, 1.080 per %), hit=0.145 ± 0.006 per rating point (10 rating = 1%, 1.446 per %), spell_haste=0.345 ± 0.081, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.143 ± 0.091)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Filigreed Shadow Circlet (253981, +0.00 DPS) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.39 DPS) | yes | - |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Invoker's Mantle (215365, -0.56 DPS) [crafted]; Death Speaker Mantle (6685, -0.83 DPS) [dungeon]; Chestnut Mantle (17695, -1.11 DPS) [quest] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Battle Healer's Cloak (19529, -0.28 DPS) [rep] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tree Bark Jacket (1486, +0.00 DPS) [dungeon]; Shadow Gown (253967, -0.23 DPS) [crafted]; Green Silk Armor (7065, -0.68 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Tabitha's Cuffs (251486, -1.67 DPS) [quest]; Phoenix Bindings (210781, -1.84 DPS) [crafted]; Windsong Bangles (263336, -2.23 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Darkness (7047, +0.00 DPS) [crafted]; Serpent Gloves (5970, -0.10 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warsong Sash (16975, +0.00 DPS) [quest]; Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.83 DPS) [dungeon] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gaze Dreamer Pants (6903, +0.00 DPS) [dungeon]; Filigreed Shadow Leggings (253943, -0.68 DPS) [crafted]; Necromancer Leggings (2277, -0.91 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.95 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.01 DPS) [crafted]; Glinteye Slippers (273024, -1.01 DPS) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.95 DPS) | yes | Ring of Power Regulation (273030, -0.28 DPS) [dungeon]; Solid Gold Band (279845, -0.83 DPS) [quest]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Solid Gold Band (279845, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.11 DPS) [dungeon]; Spine of the Devourer (273040, -1.11 DPS) [world]; Tendonscraper Dagger (273053, -1.11 DPS) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Orb of Soran'ruk (6898, +0.00 DPS) [quest]; Orb of Mistmantle (13031, -0.58 DPS) [world_drop]; Orb of Souls (249395, -0.83 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.23 DPS) [crafted]; Gravestone Scepter (7001, -0.26 DPS) [quest]; Wand of Eventide (5214, -0.31 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Windsong Drape; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (orc, 25521000132001351-0000000000000000000-0000000000000000)

Set DPS (verified): 199.4. Weights run: 4.2s. Verify run: 5.7s. 392 eligible items had no known source.

Race sweep: orc 199.4 ± 0.1, undead 198.7 ± 0.1, troll 196.8 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.192 ± 0.010, crit=0.238 ± 0.004 per rating point (14 rating = 1%, 3.338 per %), hit=0.284 ± 0.015 per rating point (10 rating = 1%, 2.840 per %), spell_haste=0.747 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.899 ± 0.001, fire_power=0.101 ± 0.000, spirit=not significant (0.296 ± 0.091)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (9.64 DPS) | yes | Living Cowl (5608, -3.67 DPS) [world]; Augural Shroud (2620, -3.71 DPS) [world]; Papal Fez (9431, -3.78 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 10.1 spell_power points (4.62 DPS) | yes | Glowing Eye of Mordresh (10769, -0.71 DPS) [dungeon]; Necklace of Calisea (1714, -3.05 DPS) [dungeon]; Amberglow Talisman (10824, -3.26 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.92 DPS) | yes | Green Silken Shoulders (7057, -0.28 DPS) [crafted]; Inquisitor's Shawl (19507, -0.57 DPS) [dungeon]; Mantle of Doan (7712, -0.82 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19528, -0.51 DPS, sim-verified) [rep]; Long Silken Cloak (4326, -0.83 DPS) [crafted]; Guardian Cloak (5965, -0.83 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.6 spell_power points (11.31 DPS) | yes | Dreamweave Vest (10021, -0.53 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.67 DPS) [world_drop]; Shadoweave Robe (10004, -1.69 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (4.13 DPS) | yes | Condor Bracers (15864, -0.76 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.59 DPS) [quest]; Earthen Silk Cuffs (254019, -1.62 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (9.57 DPS) | yes | Shadoweave Gloves (10023, -1.59 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.68 DPS) [crafted]; Black Handwraps (254027, -3.17 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Cord (254043, -0.24 DPS, sim-verified) [crafted]; Star Belt (4329, -0.27 DPS) [crafted]; Gilded Cord (254037, -1.59 DPS) [crafted] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 18.9 spell_power points (8.66 DPS) | yes | Red Mageweave Pants (10009, -0.57 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -1.59 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.93 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.1 spell_power points (11.97 DPS) | yes | Black Slippers (254007, -4.96 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -7.59 DPS) [crafted]; Thoughtcast Boots (10578, -7.90 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.5 spell_power points (6.22 DPS) | yes | Advisor's Ring (19520, -2.08 DPS) [rep]; Reedknot Ring (9622, -3.00 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (5.03 DPS) | yes | Advisor's Ring (19520, -0.90 DPS) [rep]; Reedknot Ring (9622, -1.82 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.45 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -11.31 DPS) [dungeon]; Swinetusk Shank (6691, -13.71 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.46 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 25522300132001351-0005000000000000000-0010000000000000)

Set DPS (verified): 282.6. Weights run: 4.3s. Verify run: 9.2s. 499 eligible items had no known source.

Race sweep: undead 282.7 ± 0.2, orc 282.6 ± 0.2, troll 279.1 ± 0.2 (dps); undead beats the guide race orc by 0.1 (0.05%) ± 0.2, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.204 ± 0.013, crit=0.320 ± 0.005 per rating point (14 rating = 1%, 4.481 per %), hit=0.401 ± 0.019 per rating point (10 rating = 1%, 4.011 per %), spell_haste=0.627 ± 0.106, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.900 ± 0.001, fire_power=0.100 ± 0.000, spirit=not significant (0.075 ± 0.086)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.89 DPS) | yes | Dreamweave Circlet (10041, -1.46 DPS) [crafted]; Shadoweave Mask (10025, -1.60 DPS) [crafted]; Red Mageweave Headband (10033, -1.87 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.6 spell_power points (4.12 DPS) | yes | Triune Amulet (7722, -0.05 DPS) [dungeon]; Mindburst Medallion (11196, -0.49 DPS) [quest]; Horizon Choker (13085, -2.62 DPS) [world_drop] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | 17.4 spell_power points (8.30 DPS) | yes | Kentic Amice (11624, -0.14 DPS) [dungeon]; Rotgrip Mantle (17732, -0.34 DPS) [dungeon]; Shadoweave Shoulders (10028, -0.99 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -0.88 DPS) [quest]; Runecloth Cloak (13860, -2.41 DPS) [crafted]; Nightfall Drape (12465, -3.19 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.6 spell_power points (11.27 DPS) | yes | Acumen Robes (17775, -0.25 DPS) [quest]; Elemental Raiment (9434, -1.24 DPS) [world_drop]; Dreamweave Vest (10021, -1.30 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.5 spell_power points (6.45 DPS) | yes | Spidertank Oilrag (9448, -2.15 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.25 DPS) [crafted]; Bloodband Bracers (11469, -3.08 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 19.7 spell_power points (9.43 DPS) | yes | Dreamweave Gloves (10019, -0.19 DPS) [crafted]; Shadoweave Gloves (10023, -1.63 DPS) [crafted]; Black Mageweave Gloves (10003, -2.26 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.9 spell_power points (9.04 DPS) | yes | Ban'thok Sash (11662, -0.33 DPS) [dungeon]; Satyrmane Sash (17755, -1.38 DPS) [dungeon]; Ghostweave Cord (254073, -1.96 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.8 spell_power points (12.34 DPS) | yes | Spellshock Leggings (9484, -0.39 DPS) [dungeon]; Wizardweave Leggings (14132, -3.27 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.31 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.5 spell_power points (11.71 DPS) | yes | Black Sandals (254113, -2.24 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -5.10 DPS) [vendor]; Black Slippers (254007, -5.30 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.6 spell_power points (6.50 DPS) | yes | Advisor's Ring (19519, -0.77 DPS) [rep]; Philanthropist's Ring (281635, -1.24 DPS) [quest]; Cyclopean Band (11824, -1.38 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Advisor's Ring (19519, -0.48 DPS) [rep]; Philanthropist's Ring (281635, -0.95 DPS) [quest]; Cyclopean Band (11824, -1.09 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.38 DPS) [quest]; Ankh of Life (1713, -1.29 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.37 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.52 DPS) [quest]; Ankh of Life (1713, -2.44 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.51 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.87 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.04 DPS) [dungeon]; Spire of Hakkar (10844, -3.65 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.48 DPS) [crafted]; Thunderwood (13062, -0.96 DPS) [world_drop]; Twisted Nether Wand (249144, -0.96 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Glowing Eye of Mordresh; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Lesser Eternal Wand

No-known-source sample (15 of 499, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522300132001351-2005220000000000000-0050000000000000)

Set DPS (verified): 449.6. Weights run: 4.2s. Verify run: 22.2s. 1293 eligible items had no known source.

Race sweep: orc 449.6 ± 0.3, undead 449.5 ± 0.3, troll 443.8 ± 0.3 (dps); the guide race orc is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.241 ± 0.038, crit=0.470 ± 0.013 per rating point (14 rating = 1%, 6.578 per %), hit=0.515 ± 0.049 per rating point (10 rating = 1%, 5.145 per %), spell_haste=not significant (-0.211 ± 0.426), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.898 ± 0.001, fire_power=0.102 ± 0.000, spirit=not significant (0.582 ± 0.228)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 36.6 spell_power points (18.86 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -2.41 DPS) [pvp]; Ghostweave Hood (254135, -4.22 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Triune Amulet (7722, -1.20 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Felcloth Shoulders (14112, -0.23 DPS) [crafted]; Argent Shoulders (19059, -0.76 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -1.52 DPS, sim-verified) [world]; Fel Cape (279269, -1.56 DPS) [crafted]; Spritecaster Cape (11623, -2.13 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (24.83 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.35 DPS) [pvp]; Alanna's Embrace (13314, -6.03 DPS) [dungeon]; Robe of Winter Night (14136, -6.30 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 28.0 spell_power points (14.43 DPS) | yes | Netherpearl Cuffs (254067, -6.05 DPS) [crafted]; Earthenweave Cuffs (254131, -7.31 DPS) [crafted]; Runecloth Cuffs (254123, -7.52 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Black Gloves (254101, -0.22 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 22.0 spell_power points (11.35 DPS) | yes | Ghostweave Cord (254073, -0.83 DPS) [crafted]; Thuzadin Sash (18740, -0.89 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -1.70 DPS, sim-verified) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 38.4 spell_power points (19.79 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -1.27 DPS) [dungeon]; Ghostweave Pants (14144, -2.95 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.1 spell_power points (14.47 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.73 DPS) [crafted]; Maleki's Footwraps (18735, -0.85 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.60 DPS) [quest]; Don Mauricio's Band of Domination (22433, -2.89 DPS) [dungeon]; Naglering (11669, -11.40 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.15 DPS) [dungeon]; Deep Rooted Ring (19109, -2.78 DPS) [world] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+10.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, -2.58 DPS) [world_drop]; Uther's Strength (11302, -3.09 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.18 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.17 DPS, sim-verified) [vendor]; Ankh of Life (1713, -7.74 DPS) [world_drop]; Uther's Strength (11302, -8.25 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -0.46 DPS) [pvp]; Mindfang (20214, -3.10 DPS) [rep]; Ironbark Staff (20220, -14.82 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (14.95 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.45 DPS) [pvp]; Tome of Shadow Force (19309, -0.59 DPS) [rep]; Spirit of Aquementas (11904, -4.12 DPS) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, -0.23 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.72 DPS) [dungeon]; Brilliant Wand (249385, -0.89 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 05520020135201351-0050203001000000000-0500000000000000)

Set DPS (verified): 826.2. Weights run: 3.3s. Verify run: 19.7s. 1293 eligible items had no known source.

Race sweep: orc 826.2 ± 0.7, undead 822.4 ± 0.7, troll 817.1 ± 0.7 (dps); the guide race orc is best.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.225 ± 0.013, crit=0.712 ± 0.018 per rating point (14 rating = 1%, 9.973 per %), hit=0.440 ± 0.061 per rating point (10 rating = 1%, 4.398 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, fire_power=not significant (0.000 ± 0.000), spirit=not significant (0.014 ± 0.013)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 32.3 spell_power points (23.45 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Crimson Felt Hat (18727, -0.25 DPS) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Anastari Heirloom (18728, -0.91 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.9 spell_power points (21.74 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.78 DPS) [pvp]; Argent Shoulders (19059, -3.49 DPS) [crafted]; Felcloth Shoulders (14112, -3.62 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -2.96 DPS) [world]; Arcanoweave Cloak (272411, -5.03 DPS, sim-verified) [vendor]; Spritecaster Cape (11623, -6.94 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.0 spell_power points (34.92 DPS) | yes | Robe of Winter Night (14136, -3.77 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -7.72 DPS) [pvp]; Felcloth Robe (14106, -8.75 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.9 spell_power points (17.38 DPS) | yes | Netherpearl Cuffs (254067, -6.73 DPS) [crafted]; Runecloth Cuffs (254123, -7.74 DPS) [crafted]; Arena Wristguards (18709, -10.07 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.2 spell_power points (22.69 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -2.74 DPS) [dungeon]; Earth Warder's Gloves (21318, -6.69 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 25.3 spell_power points (18.41 DPS) | yes | Black Waistcord (254087, -3.12 DPS) [crafted]; Frostwolf Cloth Belt (19090, -3.69 DPS) [rep]; Ban'thok Sash (11662, -4.69 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -1.08 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maleki's Footwraps (18735, +0.00 DPS) [dungeon]; General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.01 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Blessed Band of Light (272407, -0.45 DPS) [vendor] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, +0.00 DPS) [dungeon]; Eye of Orgrimmar (12545, -0.02 DPS) [quest]; Blessed Band of Light (272407, -3.75 DPS) [vendor] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -6.07 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -7.27 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -13.12 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -1.48 DPS) [crafted]; Ankh of Life (1713, -4.25 DPS, sim-verified) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.78 DPS) [pvp]; Whiteout Staff (19101, -24.76 DPS, sim-verified) [rep]; Kindling Stave (11750, -28.10 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Greater Eternal Wand (249237, -1.15 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger2: Maiden's Circle; trinket1: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

