# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 61.7. Weights run: 3.5s. Verify run: 4.0s. 199 eligible items had no known source.

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

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 30 (human, 25520000132001000-0000000000000000000-0000000000000000)

Set DPS (verified): 103.2. Weights run: 3.6s. Verify run: 5.2s. 326 eligible items had no known source.

Race sweep: human 103.2 ± 0.1, gnome 102.5 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.184 ± 0.005, crit=0.077 ± 0.001 per rating point (14 rating = 1%, 1.080 per %), hit=0.145 ± 0.006 per rating point (10 rating = 1%, 1.446 per %), spell_haste=0.345 ± 0.081, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.143 ± 0.091)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Circlet (253981, -0.13 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.39 DPS) | yes | Tarnished Locket (279870, -1.57 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Invoker's Mantle (215365, -0.56 DPS) [crafted]; Death Speaker Mantle (6685, -0.83 DPS) [dungeon]; Moonlit Amice (11884, -2.14 DPS, sim-verified) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Prelacy Cape (7004, -0.28 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.62 DPS) | yes | Necrology Robes (2292, -0.44 DPS) [dungeon]; Shadow Gown (253967, -0.66 DPS) [crafted]; Green Silk Armor (7065, -1.11 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Windsong Bangles (263336, -2.23 DPS) [quest]; Phoenix Bindings (210781, -2.95 DPS, sim-verified) [crafted] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent Gloves (5970, -0.10 DPS) [dungeon]; Shilly Mitts (9609, -0.10 DPS) [quest]; Hands of Darkness (7047, -0.46 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.0 spell_power points (3.06 DPS) | yes | Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.83 DPS) [dungeon]; Invoker's Cord (215366, -1.11 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gaze Dreamer Pants (6903, -0.49 DPS, sim-verified) [dungeon]; Filigreed Shadow Leggings (253943, -0.68 DPS) [crafted]; Necromancer Leggings (2277, -0.91 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.95 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.01 DPS) [crafted]; Glinteye Slippers (273024, -2.48 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.95 DPS) | yes | Minor Channeling Ring (1449, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon]; Ring of Power Regulation (273030, -2.56 DPS, sim-verified) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Minor Channeling Ring (1449, -0.28 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Ring of Power Regulation (273030, -0.99 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.11 DPS) [dungeon]; Spine of the Devourer (273040, -1.11 DPS) [world]; Tendonscraper Dagger (273053, -1.98 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.50 DPS) | yes | Orb of Mystic Insight (249394, -0.41 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.14 DPS) [world_drop]; Eye of Paleth (2943, -1.39 DPS) [quest] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.27 DPS) [crafted]; Wand of Eventide (5214, -0.28 DPS) [world_drop]; Gravestone Scepter (7001, -0.30 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 25521000132001351-0000000000000000000-0000000000000000)

Set DPS (verified): 199.6. Weights run: 3.8s. Verify run: 2.9s. 455 eligible items had no known source.

Race sweep: human 199.6 ± 0.1, gnome 198.1 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.184 ± 0.010, crit=0.250 ± 0.004 per rating point (14 rating = 1%, 3.505 per %), hit=0.298 ± 0.015 per rating point (10 rating = 1%, 2.981 per %), spell_haste=0.853 ± 0.119, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.898 ± 0.001, fire_power=0.102 ± 0.000, spirit=not significant (0.003 ± 0.084)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (9.58 DPS) | yes | Miner's Hat of the Deep (9429, -2.28 DPS, sim-verified) [dungeon]; Living Cowl (5608, -3.65 DPS) [world]; Augural Shroud (2620, -3.72 DPS) [world] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 9.0 spell_power points (4.10 DPS) | yes | Glowing Eye of Mordresh (10769, -0.24 DPS) [dungeon]; Triune Amulet (7722, -0.45 DPS) [dungeon]; Necklace of Calisea (1714, -3.51 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.86 DPS) | yes | Green Silken Shoulders (7057, -0.29 DPS) [crafted]; Inquisitor's Shawl (19507, -0.58 DPS) [dungeon]; Berylline Pads (4197, -0.83 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.05 DPS) [crafted]; Caretaker's Cape (19532, -0.46 DPS) [rep]; Long Silken Cloak (4326, -0.54 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (10.54 DPS) | yes | Elemental Raiment (9434, -0.81 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.56 DPS) [crafted]; Shadoweave Robe (10004, -1.93 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Arcane Runed Bracers (4744, -0.45 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.91 DPS) [quest]; Earthen Silk Cuffs (254019, -2.27 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (8.55 DPS) | yes | Shadoweave Gloves (10023, -1.14 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.71 DPS) [crafted]; Black Handwraps (254027, -2.23 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.7 spell_power points (6.72 DPS) | yes | Black Cord (254043, -0.30 DPS) [crafted]; Star Belt (4329, -0.79 DPS) [crafted]; Deathmage Sash (10771, -2.27 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 18.9 spell_power points (8.60 DPS) | yes | Red Mageweave Pants (10009, -0.92 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -2.27 DPS) [crafted]; Gaze Dreamer Pants (6903, -3.12 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.95 DPS) | yes | Black Slippers (254007, -5.64 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -7.17 DPS) [crafted]; Spidersilk Boots (4320, -7.43 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.5 spell_power points (6.16 DPS) | yes | Lorekeeper's Ring (19524, -2.06 DPS) [rep]; Ring of Forlorn Spirits (2043, -2.52 DPS) [quest]; Snake Hoop (6750, -2.83 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.9 spell_power points (4.98 DPS) | yes | Lorekeeper's Ring (19524, -0.45 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -1.33 DPS) [quest]; Snake Hoop (6750, -1.65 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.50 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -10.35 DPS) [dungeon]; Staff of Dar'Orahil (15106, -12.84 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Prodigious Shadowshard Pendant; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 455, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 25522300132001351-0005000000000000000-0010000000000000)

Set DPS (verified): 288.3. Weights run: 3.8s. Verify run: 7.1s. 592 eligible items had no known source.

Race sweep: human 288.3 ± 0.2, gnome 286.9 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.196 ± 0.012, crit=0.326 ± 0.005 per rating point (14 rating = 1%, 4.558 per %), hit=0.386 ± 0.019 per rating point (10 rating = 1%, 3.856 per %), spell_haste=0.765 ± 0.107, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.900 ± 0.001, fire_power=0.100 ± 0.000, spirit=not significant (0.168 ± 0.083)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.90 DPS) | yes | Dreamweave Circlet (10041, -1.02 DPS, sim-verified) [crafted]; Shadoweave Mask (10025, -1.64 DPS) [crafted]; Red Mageweave Headband (10033, -1.95 DPS) [crafted] |
| neck | Talvash's Enhancing Necklace (7673) | Restoring the Necklace [quest] | 12.2 spell_power points (5.83 DPS) | yes | Triune Amulet (7722, -1.35 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.53 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.73 DPS) [dungeon] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | 17.9 spell_power points (8.54 DPS) | yes | Kentic Amice (11624, -0.15 DPS) [dungeon]; Rotgrip Mantle (17732, -0.64 DPS) [dungeon]; Shadoweave Shoulders (10028, -1.23 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (288.3 DPS) | yes | Runecloth Cloak (13860, -2.72 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -3.43 DPS) [dungeon]; Caretaker's Cape (19531, -3.91 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (11.47 DPS) | yes | Elemental Raiment (9434, -1.44 DPS) [world_drop]; Dreamweave Vest (10021, -1.49 DPS, sim-verified) [crafted]; Shadoweave Robe (10004, -1.88 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.9 spell_power points (6.65 DPS) | yes | Nethergeld Cuffs (254061, -2.23 DPS, sim-verified) [crafted]; Arcane Runed Bracers (4744, -2.35 DPS) [quest]; Spidertank Oilrag (9448, -2.35 DPS) [dungeon] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 20.4 spell_power points (9.75 DPS) | yes | Dreamweave Gloves (10019, -0.22 DPS) [crafted]; Shadoweave Gloves (10023, -1.97 DPS) [crafted]; Runecloth Gloves (13863, -2.45 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.9 spell_power points (9.02 DPS) | yes | Ban'thok Sash (11662, -1.09 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -1.39 DPS) [dungeon]; Ghostweave Cord (254073, -1.45 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 25.8 spell_power points (12.31 DPS) | yes | Spellshock Leggings (9484, -0.38 DPS) [dungeon]; Wizardweave Leggings (14132, -3.23 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -3.27 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (12.02 DPS) | yes | Black Sandals (254113, -3.12 DPS, sim-verified) [crafted]; Shadoweave Boots (10031, -5.20 DPS) [crafted]; Black Slippers (254007, -5.46 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 14.6 spell_power points (6.97 DPS) | yes | Band of the Unicorn (7553, -0.76 DPS) [world_drop]; Lorekeeper's Ring (19523, -1.24 DPS) [rep]; Cyclopean Band (11824, -1.69 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | 13.8 spell_power points (6.59 DPS) | yes | Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.86 DPS) [rep]; Cyclopean Band (11824, -1.32 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Ankh of Life (1713, -0.70 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.86 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-verified (288.3 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -4.53 DPS, sim-verified) [world_drop] |
| off_hand | Umbral Crystal (13029) | World drop [world_drop] | 15.3 spell_power points (7.31 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.03 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.62 DPS) [dungeon]; Orb of Dar'Orahil (15108, -0.82 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-verified (288.3 DPS) | yes | Dreambough Wand (249234, -0.48 DPS) [crafted]; Burning Sliver (5249, -0.96 DPS) [quest]; Twisted Nether Wand (249144, -0.96 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Talvash's Enhancing Necklace; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Umbral Crystal; ranged: Lesser Eternal Wand

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522300132001351-2005220000000000000-0050000000000000)

Set DPS (verified): 457.0. Weights run: 3.7s. Verify run: 19.9s. 1444 eligible items had no known source.

Race sweep: human 457.0 ± 0.3, gnome 454.2 ± 0.3 (dps); the guide race human is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.234 ± 0.037, crit=0.469 ± 0.012 per rating point (14 rating = 1%, 6.567 per %), hit=0.522 ± 0.046 per rating point (10 rating = 1%, 5.217 per %), spell_haste=not significant (-0.298 ± 0.406), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.898 ± 0.001, fire_power=0.102 ± 0.000, spirit=not significant (0.168 ± 0.194)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 34.0 spell_power points (17.51 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.16 DPS) [pvp]; Felcloth Hood (14111, -1.93 DPS, sim-verified) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, -0.33 DPS) [world_drop]; Archlight Talisman (15856, -0.90 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.6 spell_power points (13.69 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.22 DPS) [crafted]; Felcloth Shoulders (14112, -1.24 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -1.11 DPS) [world]; Spritecaster Cape (11623, -3.43 DPS) [dungeon]; Shroud of the Exile (15421, -3.89 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.1 spell_power points (24.75 DPS) | yes | Robe of Winter Night (14136, -4.12 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -5.40 DPS) [pvp]; Felcloth Robe (14106, -8.12 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.1 spell_power points (12.89 DPS) | yes | Netherpearl Cuffs (254067, -5.61 DPS) [crafted]; Runecloth Cuffs (254123, -6.03 DPS) [crafted]; Manacle Cuffs (11962, -7.48 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.5 spell_power points (13.61 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Gloves of Spell Mastery (14146, -1.02 DPS) [crafted]; Earth Warder's Gloves (21318, -2.29 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 22.0 spell_power points (11.30 DPS) | yes | Stormpike Cloth Girdle (19094, -0.84 DPS) [rep]; Ban'thok Sash (11662, -1.24 DPS) [dungeon]; Black Waistcord (254087, -1.44 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 47.5 spell_power points (24.43 DPS) | yes | Marshal's Dreadweave Leggings (231587, -3.10 DPS) [pvp]; Skyshroud Leggings (13170, -7.55 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -8.45 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (12.95 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.94 DPS) [crafted]; Omnicast Boots (11822, -1.22 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.59 DPS) [dungeon]; Blessed Band of Light (272407, -2.10 DPS) [vendor]; Naglering (11669, -9.18 DPS, sim-verified) [dungeon] |
| finger2 | Songstone of Ironforge (12543) | The Princess's Surprise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.07 DPS) [dungeon]; Blessed Band of Light (272407, -1.58 DPS) [vendor]; The Postmaster's Seal (13392, -2.01 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, -3.82 DPS, sim-verified) [dungeon]; Serenity Field (272439, -5.66 DPS) [vendor]; Uther's Strength (11302, -6.17 DPS) [world_drop] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -0.57 DPS) [pvp]; Sageclaw (20070, -3.11 DPS) [rep]; Ironbark Staff (20069, -12.16 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (14.92 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.49 DPS) [pvp]; Tome of Shadow Force (19309, -0.60 DPS) [rep]; Spirit of Aquementas (11904, -4.12 DPS) [quest] |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.19 DPS) [crafted]; Greater Eternal Wand (249237, -1.03 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Songstone of Ironforge; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 05520020135201351-0050203001000000000-0500000000000000)

Set DPS (verified): 847.4. Weights run: 2.8s. Verify run: 17.4s. 1444 eligible items had no known source.

Race sweep: human 847.4 ± 0.8, gnome 843.9 ± 0.7 (dps); the guide race human is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.230 ± 0.014, crit=0.705 ± 0.018 per rating point (14 rating = 1%, 9.874 per %), hit=0.442 ± 0.061 per rating point (10 rating = 1%, 4.419 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, fire_power=not significant (0.000 ± 0.000), spirit=not significant (-0.004 ± 0.012)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) (or Crimson Felt Hat (18727)) | Tailoring [crafted] | 32.3 spell_power points (23.46 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Nacreous Shell Necklace (22403, -0.08 DPS) [dungeon]; Archlight Talisman (15856, -0.78 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.9 spell_power points (21.69 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.69 DPS) [pvp]; Felcloth Shoulders (14112, -2.81 DPS) [crafted]; Argent Shoulders (19059, -3.53 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -0.97 DPS) [world]; Spritecaster Cape (11623, -5.00 DPS) [dungeon]; Shroud of the Exile (15421, -5.78 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.1 spell_power points (34.92 DPS) | yes | Robe of Winter Night (14136, -4.09 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -7.66 DPS) [pvp]; Felcloth Robe (14106, -8.77 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.8 spell_power points (17.32 DPS) | yes | Netherpearl Cuffs (254067, -6.70 DPS) [crafted]; Runecloth Cuffs (254123, -7.66 DPS) [crafted]; Manacle Cuffs (11962, -9.73 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.1 spell_power points (22.55 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Hands of Power (13253, -4.16 DPS, sim-verified) [dungeon]; Earth Warder's Gloves (21318, -6.57 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 25.3 spell_power points (18.34 DPS) | yes | Stormpike Cloth Girdle (19094, -3.60 DPS) [rep]; Black Waistcord (254087, -3.63 DPS, sim-verified) [crafted]; Ban'thok Sash (11662, -4.75 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.1 spell_power points (39.26 DPS) | yes | Marshal's Dreadweave Leggings (231587, -9.21 DPS) [pvp]; Skyshroud Leggings (13170, -16.14 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -16.75 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (17.43 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.90 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -1.12 DPS) [pvp] |
| finger1 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.92 DPS) [quest]; Underworld Band (1980, -4.16 DPS) [world_drop]; Naglering (11669, -16.32 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Underworld Band (1980, -3.24 DPS) [world_drop]; Blessed Band of Light (272407, -3.77 DPS) [vendor] |
| trinket1 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-verified (+14.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -8.51 DPS, sim-verified) [vendor]; Uther's Strength (11302, -8.72 DPS) [world_drop]; Second Wind (11819, -9.44 DPS) [dungeon] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.75 DPS) [pvp]; Headmaster's Charge (13937, -15.48 DPS, sim-verified) [dungeon]; Trindlehaven Staff (13161, -25.75 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.28 DPS) [crafted]; Greater Eternal Wand (249237, -1.45 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Don Mauricio's Band of Domination; finger2: Maiden's Circle; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.7. Weights run: 3.5s. Verify run: 3.3s. 193 eligible items had no known source.

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

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 8350 The 1 Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade

### Band 30 (orc, 25520000132001000-0000000000000000000-0000000000000000)

Set DPS (verified): 103.3. Weights run: 3.6s. Verify run: 6.1s. 319 eligible items had no known source.

Race sweep: undead 103.7 ± 0.1, orc 103.3 ± 0.1, troll 101.9 ± 0.1 (dps); undead beats the guide race orc by 0.4 (0.35%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.184 ± 0.005, crit=0.077 ± 0.001 per rating point (14 rating = 1%, 1.080 per %), hit=0.145 ± 0.006 per rating point (10 rating = 1%, 1.446 per %), spell_haste=0.345 ± 0.081, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.143 ± 0.091)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Circlet (253981, -0.39 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.39 DPS) | yes | Roadwatcher's Confidence (281265, -1.09 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Invoker's Mantle (215365, -0.78 DPS, sim-verified) [crafted]; Death Speaker Mantle (6685, -0.83 DPS) [dungeon]; Chestnut Mantle (17695, -1.11 DPS) [quest] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Battle Healer's Cloak (19529, -0.28 DPS) [rep]; Hillman's Cloak (3719, -0.50 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Tree Bark Jacket (1486, -0.21 DPS, sim-verified) [dungeon]; Shadow Gown (253967, -0.23 DPS) [crafted]; Green Silk Armor (7065, -0.68 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.50 DPS) | yes | Phoenix Bindings (210781, -1.84 DPS) [crafted]; Windsong Bangles (263336, -2.23 DPS) [quest]; Tabitha's Cuffs (251486, -2.23 DPS, sim-verified) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent Gloves (5970, -0.10 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world]; Hands of Darkness (7047, -0.85 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Warsong Sash (16975, -0.47 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.83 DPS) [dungeon] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Leggings (253943, -0.68 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.77 DPS, sim-verified) [dungeon]; Necromancer Leggings (2277, -0.91 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.95 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.01 DPS) [crafted]; Glinteye Slippers (273024, -1.50 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.95 DPS) | yes | Solid Gold Band (279845, -0.83 DPS) [quest]; Electrocutioner Lagnut (9447, -1.11 DPS) [dungeon]; Ring of Power Regulation (273030, -1.81 DPS, sim-verified) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Solid Gold Band (279845, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -0.83 DPS) [dungeon]; Ring of Power Regulation (273030, -1.57 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.11 DPS) [dungeon]; Spine of the Devourer (273040, -1.11 DPS) [world]; Tendonscraper Dagger (273053, -1.47 DPS, sim-verified) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Soran'ruk (6898, -0.38 DPS, sim-verified) [quest]; Orb of Mistmantle (13031, -0.58 DPS) [world_drop]; Orb of Souls (249395, -0.83 DPS) [crafted] |
| ranged | Gravestone Scepter (7001) | Blackfathom Villainy [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Wand of Eventide (5214, -0.06 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Windsong Drape; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Gravestone Scepter

No-known-source sample (15 of 319, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 8350 The 1 Ring

### Band 40 (orc, 25521000132001351-0000000000000000000-0000000000000000)

Set DPS (verified): 199.8. Weights run: 3.8s. Verify run: 2.2s. 448 eligible items had no known source.

Race sweep: orc 199.8 ± 0.1, undead 199.1 ± 0.1, troll 197.2 ± 0.1 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.184 ± 0.010, crit=0.250 ± 0.004 per rating point (14 rating = 1%, 3.505 per %), hit=0.298 ± 0.015 per rating point (10 rating = 1%, 2.981 per %), spell_haste=0.853 ± 0.119, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.898 ± 0.001, fire_power=0.102 ± 0.000, spirit=not significant (0.003 ± 0.084)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (9.58 DPS) | yes | Miner's Hat of the Deep (9429, -3.57 DPS, sim-verified) [dungeon]; Living Cowl (5608, -3.65 DPS) [world]; Augural Shroud (2620, -3.72 DPS) [world] |
| neck | Prodigious Shadowshard Pendant (17773) | Shadowshard Fragments [quest] | 9.0 spell_power points (4.10 DPS) | yes | Triune Amulet (7722, -0.45 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.33 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -3.51 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.86 DPS) | yes | Inquisitor's Shawl (19507, -0.58 DPS) [dungeon]; Berylline Pads (4197, -0.83 DPS) [quest]; Green Silken Shoulders (7057, -0.96 DPS, sim-verified) [crafted] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.05 DPS) [crafted]; Battle Healer's Cloak (19528, -0.46 DPS) [rep]; Long Silken Cloak (4326, -1.32 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (10.54 DPS) | yes | Elemental Raiment (9434, -1.50 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.56 DPS) [crafted]; Shadoweave Robe (10004, -1.93 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (4.10 DPS) | yes | Radiant Silver Bracers (4545, -1.61 DPS) [quest]; Condor Bracers (15864, -1.69 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -2.27 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (8.55 DPS) | yes | Black Mageweave Gloves (10003, -1.71 DPS) [crafted]; Shadoweave Gloves (10023, -2.20 DPS, sim-verified) [crafted]; Black Handwraps (254027, -2.23 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.7 spell_power points (6.72 DPS) | yes | Star Belt (4329, -0.79 DPS) [crafted]; Black Cord (254043, -1.04 DPS, sim-verified) [crafted]; Warsong Sash (16975, -1.70 DPS) [quest] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 18.9 spell_power points (8.60 DPS) | yes | Red Mageweave Pants (10009, -1.96 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -2.27 DPS) [crafted]; Gaze Dreamer Pants (6903, -3.12 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.95 DPS) | yes | Black Slippers (254007, -6.45 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -7.17 DPS) [crafted]; Spidersilk Boots (4320, -7.43 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.5 spell_power points (6.16 DPS) | yes | Advisor's Ring (19520, -2.06 DPS) [rep]; Snake Hoop (6750, -2.83 DPS) [quest]; Reedknot Ring (9622, -2.97 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.9 spell_power points (4.98 DPS) | yes | Advisor's Ring (19520, -1.13 DPS, sim-verified) [rep]; Snake Hoop (6750, -1.65 DPS) [quest]; Reedknot Ring (9622, -1.79 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.25 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -10.35 DPS) [dungeon]; Staff of Dar'Orahil (15106, -12.84 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.27 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Prodigious Shadowshard Pendant; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 25522300132001351-0005000000000000000-0010000000000000)

Set DPS (verified): 288.9. Weights run: 3.8s. Verify run: 8.3s. 585 eligible items had no known source.

Race sweep: orc 288.9 ± 0.2, undead 288.2 ± 0.2, troll 285.1 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.144 ± 0.008, crit=0.162 ± 0.003 per rating point (14 rating = 1%, 2.265 per %), hit=0.209 ± 0.010 per rating point (10 rating = 1%, 2.086 per %), spell_haste=not significant (0.172 ± 0.066), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.892 ± 0.001, fire_power=0.108 ± 0.000, spirit=not significant (0.015 ± 0.082)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.17 DPS) | yes | Shadoweave Mask (10025, -1.17 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.97 DPS) [crafted]; Red Mageweave Headband (10033, -2.31 DPS) [crafted] |
| neck | Jarkal's Enhancing Necklace (7888) | Necklace Recovery, Take 3 [quest] | 12.1 spell_power points (5.47 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.64 DPS, sim-verified) [quest]; Glowing Eye of Mordresh (10769, -1.80 DPS) [dungeon]; Triune Amulet (7722, -1.82 DPS) [dungeon] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | 16.4 spell_power points (7.40 DPS) | yes | Kentic Amice (11624, -0.20 DPS) [dungeon]; Rotgrip Mantle (17732, -0.37 DPS) [dungeon]; Shadoweave Shoulders (10028, -0.56 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.25 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.17 DPS) [crafted]; Nightfall Drape (12465, -2.69 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (10.34 DPS) | yes | Elemental Raiment (9434, -1.54 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.55 DPS) [crafted]; Shadoweave Robe (10004, -1.85 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 12.7 spell_power points (5.72 DPS) | yes | Spidertank Oilrag (9448, -2.02 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -2.07 DPS) [crafted]; Condor Bracers (15864, -2.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.7 spell_power points (8.42 DPS) | yes | Black Gloves (254101, -0.08 DPS) [crafted]; Shadoweave Gloves (10023, -1.26 DPS) [crafted]; Black Mageweave Gloves (10003, -1.66 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 18.2 spell_power points (8.22 DPS) | yes | Ban'thok Sash (11662, -1.22 DPS) [dungeon]; Satyrmane Sash (17755, -1.26 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.65 DPS) [rep] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 24.9 spell_power points (11.23 DPS) | yes | Spellshock Leggings (9484, -0.22 DPS) [dungeon]; Wizardweave Leggings (14132, -2.67 DPS) [crafted]; Shadoweave Pants (10002, -2.79 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.1 spell_power points (10.87 DPS) | yes | Black Sandals (254113, -2.93 DPS, sim-verified) [crafted]; Black Slippers (254007, -5.16 DPS) [crafted]; Shadoweave Boots (10031, -5.17 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 13.3 spell_power points (6.00 DPS) | yes | Advisor's Ring (19519, -0.59 DPS) [rep]; Philanthropist's Ring (281635, -1.16 DPS) [quest]; Band of the Unicorn (7553, -6.85 DPS, sim-verified) [world_drop] |
| finger2 | Brainlash (6440) | Zul'Farrak [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Advisor's Ring (19519, -0.23 DPS) [rep]; Band of the Unicorn (7553, -0.53 DPS, sim-verified) [world_drop]; Philanthropist's Ring (281635, -0.81 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.19 DPS) [quest]; Ankh of Life (1713, -0.76 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.78 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -1.47 DPS, sim-verified) [quest]; Ankh of Life (1713, -2.62 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.64 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, +0.00 DPS) [dungeon]; Inventor's Focal Sword (17719, +0.00 DPS) [dungeon]; Staff of Jordan (873, -3.56 DPS, sim-verified) [world_drop] |
| off_hand | Umbral Crystal (13029) | World drop [world_drop] | 15.2 spell_power points (6.84 DPS) | yes | Mordresh's Lifeless Skull (10770, -0.37 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.52 DPS) [dungeon]; Orb of Dar'Orahil (15108, -0.88 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.45 DPS) [crafted]; Thunderwood (13062, -0.90 DPS) [world_drop]; Twisted Nether Wand (249144, -0.90 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Jarkal's Enhancing Necklace; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Umbral Crystal; ranged: Lesser Eternal Wand

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522300132001351-2005220000000000000-0050000000000000)

Set DPS (verified): 450.9. Weights run: 3.7s. Verify run: 20.1s. 1438 eligible items had no known source.

Race sweep: orc 450.9 ± 0.3, undead 449.9 ± 0.3, troll 444.8 ± 0.3 (dps); the guide race orc is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.234 ± 0.037, crit=0.469 ± 0.012 per rating point (14 rating = 1%, 6.567 per %), hit=0.522 ± 0.046 per rating point (10 rating = 1%, 5.217 per %), spell_haste=not significant (-0.298 ± 0.406), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.898 ± 0.001, fire_power=0.102 ± 0.000, spirit=not significant (0.168 ± 0.194)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 34.0 spell_power points (17.51 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -1.16 DPS) [pvp]; Felcloth Hood (14111, -3.01 DPS, sim-verified) [crafted] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Lady Maye's Pendant (14558, -0.33 DPS) [world_drop]; Archlight Talisman (15856, -0.90 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.6 spell_power points (13.69 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.22 DPS) [crafted]; Felcloth Shoulders (14112, -1.24 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -1.11 DPS) [world]; Spritecaster Cape (11623, -3.43 DPS) [dungeon]; Shroud of the Exile (15421, -3.89 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.1 spell_power points (24.75 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.40 DPS) [pvp]; Robe of Winter Night (14136, -6.37 DPS, sim-verified) [crafted]; Felcloth Robe (14106, -8.12 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.1 spell_power points (12.89 DPS) | yes | Netherpearl Cuffs (254067, -5.61 DPS) [crafted]; Runecloth Cuffs (254123, -6.03 DPS) [crafted]; Manacle Cuffs (11962, -7.48 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.5 spell_power points (13.61 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Gloves of Spell Mastery (14146, -1.02 DPS) [crafted]; Earth Warder's Gloves (21318, -2.29 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 22.0 spell_power points (11.30 DPS) | yes | Ban'thok Sash (11662, -1.24 DPS) [dungeon]; Black Waistcord (254087, -1.44 DPS) [crafted]; Frostwolf Cloth Belt (19090, -2.01 DPS, sim-verified) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 34.1 spell_power points (17.56 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.68 DPS) [dungeon]; Legionnaire's Dreadweave Legguards (227097, -1.59 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (12.95 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.94 DPS) [crafted]; Omnicast Boots (11822, -1.22 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.59 DPS) [dungeon]; Blessed Band of Light (272407, -2.10 DPS) [vendor]; Naglering (11669, -10.69 DPS, sim-verified) [dungeon] |
| finger2 | Eye of Orgrimmar (12545) | The Princess Saved? [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -1.07 DPS) [dungeon]; Blessed Band of Light (272407, -1.58 DPS) [vendor]; The Postmaster's Seal (13392, -2.01 DPS) [dungeon] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+9.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Serenity Field (272439, -5.66 DPS) [vendor]; Uther's Strength (11302, -6.17 DPS) [world_drop]; Smolderweb's Eye (13213, -6.58 DPS) [dungeon] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Draconic Infused Emblem (22268, -4.25 DPS, sim-verified) [dungeon] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -0.57 DPS) [pvp]; Mindfang (20214, -3.11 DPS) [rep]; Ironbark Staff (20220, -13.30 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (14.92 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.49 DPS) [pvp]; Tome of Shadow Force (19309, -0.60 DPS) [rep]; Spirit of Aquementas (11904, -4.12 DPS) [quest] |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.19 DPS) [crafted]; Greater Eternal Wand (249237, -1.03 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Eye of Orgrimmar; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 05520020135201351-0050203001000000000-0500000000000000)

Set DPS (verified): 833.5. Weights run: 2.8s. Verify run: 17.1s. 1438 eligible items had no known source.

Race sweep: orc 833.5 ± 0.7, undead 829.8 ± 0.7, troll 824.2 ± 0.7 (dps); the guide race orc is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.230 ± 0.014, crit=0.705 ± 0.018 per rating point (14 rating = 1%, 9.874 per %), hit=0.442 ± 0.061 per rating point (10 rating = 1%, 4.419 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, fire_power=not significant (0.000 ± 0.000), spirit=not significant (-0.004 ± 0.012)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) (or Crimson Felt Hat (18727)) | Tailoring [crafted] | 32.3 spell_power points (23.46 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp] |
| neck | Lady Maye's Pendant (14558) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dark Advisor's Pendant (18691, +0.00 DPS) [dungeon]; Nacreous Shell Necklace (22403, -0.08 DPS) [dungeon]; Archlight Talisman (15856, -0.78 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.9 spell_power points (21.69 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.69 DPS) [pvp]; Felcloth Shoulders (14112, -2.81 DPS) [crafted]; Argent Shoulders (19059, -3.53 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -0.97 DPS) [world]; Spritecaster Cape (11623, -5.00 DPS) [dungeon]; Shroud of the Exile (15421, -5.78 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.1 spell_power points (34.92 DPS) | yes | Robe of Winter Night (14136, -4.26 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -7.66 DPS) [pvp]; Felcloth Robe (14106, -8.77 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.8 spell_power points (17.32 DPS) | yes | Netherpearl Cuffs (254067, -6.70 DPS) [crafted]; Runecloth Cuffs (254123, -7.66 DPS) [crafted]; Manacle Cuffs (11962, -9.73 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.1 spell_power points (22.55 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -3.96 DPS, sim-verified) [dungeon]; Earth Warder's Gloves (21318, -6.57 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 25.3 spell_power points (18.34 DPS) | yes | Black Waistcord (254087, -3.04 DPS) [crafted]; Frostwolf Cloth Belt (19090, -3.60 DPS) [rep]; Ban'thok Sash (11662, -4.75 DPS) [dungeon] |
| legs | Skyshroud Leggings (13170) | Blackrock Spire: Highlord Omokk [dungeon] | 32.8 spell_power points (23.80 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Outrider's Silk Leggings (22747, -0.28 DPS) [rep]; Legionnaire's Dreadweave Legguards (227097, -1.29 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (17.43 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.90 DPS) [dungeon]; Black Sandals (254113, -1.40 DPS) [crafted] |
| finger1 | Don Mauricio's Band of Domination (22433) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.92 DPS) [quest]; Underworld Band (1980, -4.16 DPS) [world_drop]; Naglering (11669, -15.69 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, +0.00 DPS) [quest]; Underworld Band (1980, -3.24 DPS) [world_drop]; Blessed Band of Light (272407, -3.77 DPS) [vendor] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+14.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Second Wind (11819, +0.00 DPS) [dungeon]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -8.43 DPS, sim-verified) [vendor]; Uther's Strength (11302, -8.72 DPS) [world_drop]; Second Wind (11819, -9.44 DPS) [dungeon] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.75 DPS) [pvp]; Headmaster's Charge (13937, -16.05 DPS, sim-verified) [dungeon]; Whiteout Staff (19101, -24.35 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.28 DPS) [crafted]; Greater Eternal Wand (249237, -1.45 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Lady Maye's Pendant; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Skyshroud Leggings; finger1: Don Mauricio's Band of Domination; finger2: Maiden's Circle; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

