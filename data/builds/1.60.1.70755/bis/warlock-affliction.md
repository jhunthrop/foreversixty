# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 59.7. Weights run: 5.0s. Verify run: 3.1s. 179 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.176 ± 0.006, crit=0.059 ± 0.001 per rating point (14 rating = 1%, 0.833 per %), hit=0.174 ± 0.007 per rating point (10 rating = 1%, 1.740 per %), spell_haste=0.550 ± 0.088, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.821 ± 0.001, fire_power=0.179 ± 0.000, spirit=0.320 ± 0.063

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.0 spell_power points (2.56 DPS) | yes | Pristine Circlet (253949, -0.85 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.77 DPS) [crafted]; Flame Circlet (253953, -2.00 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.6 spell_power points (0.18 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.6 spell_power points (1.86 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.37 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.65 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.78 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.4 spell_power points (3.50 DPS) | yes | Filigreed Shadow Gown (253907, -1.22 DPS, sim-verified) [crafted]; Scholarly Robes (2034, -1.57 DPS) [dungeon]; Filigreed Pristine Gown (253901, -1.57 DPS) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.28 DPS) | yes | Repurposed Hair Band (281256, -0.00 DPS) [quest]; Mindthrust Bracers (1974, -0.03 DPS) [dungeon]; Seer's Cuffs (3645, -0.05 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.98 DPS) | yes | Shadow Gloves (253919, -0.24 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.28 DPS) [world]; Pristine Gloves (253913, -0.70 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.5 spell_power points (1.83 DPS) | yes | Pristine Sash (253925, -0.49 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted]; Novice Ardent's Sash (253887, -0.84 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 11.4 spell_power points (3.22 DPS) | yes | Filigreed Pristine Leggings (253937, -0.85 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.24 DPS) [dungeon]; Colorful Kilt (10048, -1.35 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.7 spell_power points (2.18 DPS) | yes | Shadow Boots (253895, -0.64 DPS) [crafted]; Red Woolen Boots (4313, -0.91 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.18 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.4 spell_power points (1.52 DPS) | yes | Band of Purification (12996, -0.97 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.22 DPS) [dungeon]; Furlfeather Ring (286745, -1.23 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.42 DPS) | yes | Lavishly Jeweled Ring (1156, -1.12 DPS) [dungeon]; Furlfeather Ring (286745, -1.13 DPS) [world]; Band of Purification (12996, -1.35 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (59.7 DPS) | yes | Staff of the Blessed Seer (2271, -0.91 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.69 DPS) [dungeon]; Witching Stave (1484, -1.95 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.3 spell_power points (0.94 DPS) | yes | Onu's Patience (281292, -0.49 DPS) [quest]; Grayson's Torch (1172, -0.58 DPS) [quest]; Antipodean Rod (2879, -0.64 DPS) [world_drop] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.28 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 25522300100001000-0000000000000000000-0000000000000000)

Set DPS (verified): 100.2. Weights run: 5.1s. Verify run: 6.6s. 303 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.209 ± 0.005, crit=0.049 ± 0.001 per rating point (14 rating = 1%, 0.681 per %), hit=0.136 ± 0.005 per rating point (10 rating = 1%, 1.355 per %), spell_haste=0.383 ± 0.077, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.046 ± 0.085)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Circlet (253981, -0.12 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.40 DPS) | yes | Tarnished Locket (279870, -1.55 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.51 DPS) | yes | Invoker's Mantle (215365, -0.56 DPS) [crafted]; Death Speaker Mantle (6685, -0.84 DPS) [dungeon]; Moonlit Amice (11884, -1.98 DPS, sim-verified) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Caretaker's Cape (19533, -0.28 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.63 DPS) | yes | Necrology Robes (2292, -0.43 DPS) [dungeon]; Shadow Gown (253967, -0.66 DPS) [crafted]; Green Silk Armor (7065, -1.12 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.51 DPS) | yes | Windsong Bangles (263336, -2.23 DPS) [quest]; Phoenix Bindings (210781, -2.51 DPS, sim-verified) [crafted] |
| hands | Hands of Darkness (7047) (or Ephemeral Grips (273034)) | Tailoring [crafted] | 7.4 spell_power points (2.05 DPS) | yes | Ephemeral Grips (273034, +0.00 DPS) [world]; Serpent Gloves (5970, -0.10 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.0 spell_power points (3.07 DPS) | yes | Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.84 DPS) [dungeon]; Invoker's Cord (215366, -1.12 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.35 DPS) | yes | Shadow Leggings (253993, -0.16 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.84 DPS) [crafted]; Necromancer Leggings (2277, -1.07 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.97 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.01 DPS) [crafted]; Glinteye Slippers (273024, -1.72 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.95 DPS) | yes | Snake Hoop (6750, -0.28 DPS) [quest]; Minor Channeling Ring (1449, -0.56 DPS) [quest]; Electrocutioner Lagnut (9447, -1.12 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.68 DPS) | yes | Snake Hoop (6750, -0.26 DPS, sim-verified) [quest]; Minor Channeling Ring (1449, -0.28 DPS) [quest]; Electrocutioner Lagnut (9447, -0.84 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.12 DPS) [dungeon]; Spine of the Devourer (273040, -1.12 DPS) [world]; Tendonscraper Dagger (273053, -1.46 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.51 DPS) | yes | Orb of Mystic Insight (249394, -0.44 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.14 DPS) [world_drop]; Eye of Paleth (2943, -1.40 DPS) [quest] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Eventide (5214, -0.28 DPS) [world_drop]; Gravestone Scepter (7001, -0.31 DPS) [quest]; Greater Mystic Wand (217287, -0.31 DPS, sim-verified) [crafted] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Hands of Darkness; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Power Regulation; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 303, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 153.4. Weights run: 5.2s. Verify run: 5.2s. 386 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.041 ± 0.005, crit=0.090 ± 0.002 per rating point (14 rating = 1%, 1.256 per %), hit=0.158 ± 0.006 per rating point (10 rating = 1%, 1.576 per %), spell_haste=0.675 ± 0.081, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.830 ± 0.001, fire_power=0.170 ± 0.000, spirit=not significant (0.002 ± 0.085)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.48 DPS) | yes | Living Cowl (5608, -2.69 DPS, sim-verified) [world]; Filigreed Shadow Circlet (253981, -2.89 DPS) [crafted]; Holy Shroud (2721, -3.08 DPS) [world_drop] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Triune Amulet (7722, -0.23 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -2.15 DPS) [quest]; Explorers' League Commendation (7746, -2.15 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.78 DPS) | yes | Green Silken Shoulders (7057, -0.31 DPS) [crafted]; Mantle of Doan (7712, -0.61 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.62 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.31 DPS) [crafted]; Guardian Cloak (5965, -0.31 DPS) [crafted]; Caretaker's Cape (19532, -0.51 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 spell_power points (6.79 DPS) | yes | Elemental Raiment (9434, -0.66 DPS, sim-verified) [world_drop]; Dreamweave Vest (10021, -1.23 DPS) [crafted]; Shadoweave Robe (10004, -1.41 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Arcane Runed Bracers (4744, -0.15 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.62 DPS) [quest]; Earthen Silk Cuffs (254019, -1.54 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 spell_power points (5.55 DPS) | yes | Shadoweave Gloves (10023, -1.20 DPS) [crafted]; Black Mageweave Gloves (10003, -1.41 DPS, sim-verified) [crafted]; Black Handwraps (254027, -1.97 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.0 spell_power points (4.32 DPS) | yes | Star Belt (4329, -0.31 DPS) [crafted]; Black Cord (254043, -0.73 DPS) [crafted]; Belt of Arugal (6392, -1.54 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.4 spell_power points (5.38 DPS) | yes | Red Mageweave Pants (10009, -0.65 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.67 DPS) [dungeon]; Shadow Leggings (253993, -1.79 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.40 DPS) | yes | Black Slippers (254007, -3.79 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -5.24 DPS) [dungeon]; Gilded Slippers (254001, -5.24 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 11.6 spell_power points (3.58 DPS) | yes | Lorekeeper's Ring (19524, -0.81 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.12 DPS) [quest]; Reedknot Ring (9622, -1.43 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.0 spell_power points (3.08 DPS) | yes | Lorekeeper's Ring (19524, -0.46 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.62 DPS) [quest]; Reedknot Ring (9622, -0.93 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hypnotic Blade (7714, -5.67 DPS) [dungeon]; Illusionary Rod (7713, -6.38 DPS, sim-verified) [dungeon]; Swinetusk Shank (6691, -6.90 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.31 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 386, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 213.3. Weights run: 5.3s. Verify run: 8.4s. 485 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.175 ± 0.007, crit=0.137 ± 0.002 per rating point (14 rating = 1%, 1.924 per %), hit=0.228 ± 0.009 per rating point (10 rating = 1%, 2.283 per %), spell_haste=1.026 ± 0.113, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.845 ± 0.001, fire_power=0.155 ± 0.000, spirit=not significant (0.008 ± 0.115)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.08 DPS) | yes | Dreamweave Circlet (10041, -0.64 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -1.51 DPS) [crafted]; Shadoweave Mask (10025, -1.67 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.32 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.54 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -1.87 DPS) [world_drop] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadoweave Shoulders (10028, -0.49 DPS) [crafted]; Rotgrip Mantle (17732, -0.58 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -1.42 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -1.71 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.05 DPS) [dungeon]; Caretaker's Cape (19531, -2.71 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (7.76 DPS) | yes | Elemental Raiment (9434, -0.70 DPS) [world_drop]; Acumen Robes (17775, -0.80 DPS, sim-verified) [quest]; Dreamweave Vest (10021, -1.14 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 12.2 spell_power points (4.12 DPS) | yes | Arcane Runed Bracers (4744, -0.85 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -1.09 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.34 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (6.31 DPS) | yes | Black Gloves (254101, -0.30 DPS) [crafted]; Shadoweave Gloves (10023, -1.19 DPS) [crafted]; Black Mageweave Gloves (10003, -1.26 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 17.6 spell_power points (5.93 DPS) | yes | Satyrmane Sash (17755, -0.54 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.98 DPS) [rep]; Ghostweave Cord (254073, -1.19 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 24.1 spell_power points (8.09 DPS) | yes | Wizardweave Leggings (14132, -0.69 DPS, sim-verified) [crafted]; Shadoweave Pants (10002, -2.13 DPS) [crafted]; Red Mageweave Pants (10009, -2.68 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.1 spell_power points (8.09 DPS) | yes | Black Sandals (254113, -1.90 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -3.86 DPS) [crafted]; Black Slippers (254007, -3.97 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.37 DPS) | yes | Lorekeeper's Ring (19523, -0.34 DPS) [rep]; Philanthropist's Ring (281635, -0.71 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.7 spell_power points (4.27 DPS) | yes | Lorekeeper's Ring (19523, -0.24 DPS) [rep]; Philanthropist's Ring (281635, -0.61 DPS) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -1.12 DPS, sim-verified) [world_drop] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -0.25 DPS, sim-verified) [crafted]; Ankh of Life (1713, -1.98 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.99 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -0.54 DPS, sim-verified) [world]; Kindling Stave (11750, -1.00 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.71 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.34 DPS) [crafted]; Burning Sliver (5249, -0.67 DPS) [quest]; Twisted Nether Wand (249144, -0.67 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; neck: Triune Amulet; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger1: Band of the Unicorn; finger2: Underworld Band; trinket1: Thunderbrew's Boot Flask; trinket2: Uther's Strength; ranged: Lesser Eternal Wand

No-known-source sample (15 of 485, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 345.5. Weights run: 5.5s. Verify run: 21.2s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.399 ± 0.016, crit=0.199 ± 0.003 per rating point (14 rating = 1%, 2.791 per %), hit=0.386 ± 0.016 per rating point (10 rating = 1%, 3.865 per %), spell_haste=not significant (-0.003 ± 0.235), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.844 ± 0.001, fire_power=0.156 ± 0.000, spirit=not significant (-0.558 ± 0.244)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.50 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Felcloth Hood (14111, -0.59 DPS) [crafted]; Lieutenant Commander's Dreadweave Cowl (227093, -1.13 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.0 spell_power points (4.22 DPS) | yes | Glowing Eye of Mordresh (10769, -1.76 DPS) [dungeon]; Mindburst Medallion (11196, -2.11 DPS) [quest]; Triune Amulet (7722, -2.17 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (8.80 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Felcloth Shoulders (14112, -1.40 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -1.83 DPS) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-verified (345.5 DPS) | yes | Spritecaster Cape (11623, -0.51 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -1.75 DPS, sim-verified) [rep]; Nightfall Drape (12465, -2.27 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.0 spell_power points (16.19 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -4.93 DPS) [pvp]; Robe of Winter Night (14136, -5.25 DPS, sim-verified) [crafted]; Felcloth Robe (14106, -5.50 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.0 spell_power points (7.74 DPS) | yes | Runecloth Cuffs (254123, -3.87 DPS) [crafted]; Netherpearl Cuffs (254067, -3.88 DPS) [crafted]; Arcane Runed Bracers (4744, -4.57 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (7.74 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.35 DPS) [pvp]; Dreamweave Gloves (10019, -1.27 DPS, sim-verified) [crafted] |
| waist | Stormpike Cloth Girdle (19094) | Stormpike Guard [rep] | 18.0 spell_power points (6.33 DPS) | yes | Highlander's Cloth Girdle (20047, -0.43 DPS) [rep]; Black Waistcord (254087, -0.69 DPS) [crafted]; Ghostweave Cord (254073, -1.41 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 37.6 spell_power points (13.23 DPS) | yes | Marshal's Dreadweave Leggings (231587, -0.20 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -3.37 DPS) [pvp]; Ghostweave Pants (14144, -4.08 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.45 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.12 DPS, sim-verified) [crafted]; Omnicast Boots (11822, -1.41 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.0 spell_power points (6.33 DPS) | yes | Band of the Unicorn (7553, -1.76 DPS) [world_drop]; Deep Rooted Ring (19109, -2.11 DPS) [world]; Underworld Band (1980, -2.18 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 14.0 spell_power points (4.93 DPS) | yes | Deep Rooted Ring (19109, -0.70 DPS) [world]; Underworld Band (1980, -0.77 DPS) [world_drop]; Band of the Unicorn (7553, -1.64 DPS, sim-verified) [world_drop] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (345.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (345.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (345.5 DPS) | yes | Sageclaw (20070, -2.49 DPS) [rep]; Grand Marshal's Stave (234571, -3.97 DPS) [pvp]; Ironbark Staff (20069, -12.66 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (10.21 DPS) | yes | Tome of Shadow Force (19309, -0.89 DPS, sim-verified) [rep]; Grand Marshal's Tome of Power (234589, -1.41 DPS) [pvp]; Spirit of Aquementas (11904, -2.82 DPS) [quest] |
| ranged | Greater Eternal Wand (249237) | Enchanting [crafted] | sim-verified (345.5 DPS) | yes | Brilliant Wand (249385, +0.00 DPS) [crafted]; Lesser Eternal Wand (249232, -0.35 DPS) [crafted]; Dreambough Wand (249234, -0.70 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; back: Fel Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Stormpike Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; trinket2: Frozen Heart of the Mountain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Greater Eternal Wand

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25520010135001050-0005200000000000000-0550004000000000)

Set DPS (verified): 640.9. Weights run: 5.9s. Verify run: 22.3s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.254 ± 0.032, crit=0.677 ± 0.017 per rating point (14 rating = 1%, 9.478 per %), hit=1.053 ± 0.066 per rating point (10 rating = 1%, 10.528 per %), spell_haste=5.601 ± 0.686, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.884 ± 0.001, fire_power=0.116 ± 0.000, spirit=not significant (-0.019 ± 0.180)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 29.1 spell_power points (16.37 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -1.16 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.5 spell_power points (8.19 DPS) | yes | Triune Amulet (7722, -3.68 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -3.81 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.95 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.8 spell_power points (16.78 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.26 DPS) [pvp]; Argent Shoulders (19059, -2.69 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.82 DPS) [pvp] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-verified (640.9 DPS) | yes | Stormpike Sage's Cloak (19086, -4.92 DPS) [rep]; Spritecaster Cape (11623, -4.96 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -6.59 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.3 spell_power points (27.20 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.74 DPS) [pvp]; Robe of Winter Night (14136, -5.98 DPS, sim-verified) [crafted]; Felcloth Robe (14106, -9.28 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.0 spell_power points (13.54 DPS) | yes | Runecloth Cuffs (254123, -5.91 DPS) [crafted]; Netherpearl Cuffs (254067, -6.06 DPS) [crafted]; Arena Wristguards (18709, -8.20 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.5 spell_power points (17.18 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -4.19 DPS, sim-verified) [quest]; Knight-Lieutenant's Dreadweave Handwraps (227100, -4.78 DPS) [pvp] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 25.0 spell_power points (14.09 DPS) | yes | Stormpike Cloth Girdle (19094, -2.51 DPS) [rep]; Black Waistcord (254087, -3.34 DPS) [crafted]; Satyrmane Sash (17755, -4.77 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 53.5 spell_power points (30.14 DPS) | yes | Marshal's Dreadweave Leggings (231587, -6.58 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -12.51 DPS) [pvp]; Felcloth Pants (14107, -15.48 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.52 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.67 DPS) [pvp]; Omnicast Boots (11822, -3.02 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 20.0 spell_power points (11.27 DPS) | yes | Deep Rooted Ring (19109, -3.08 DPS) [world]; Lorekeeper's Ring (19522, -3.38 DPS) [rep]; Underworld Band (1980, -3.59 DPS) [world_drop] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 19.8 spell_power points (11.14 DPS) | yes | Deep Rooted Ring (19109, -3.04 DPS, sim-verified) [world]; Lorekeeper's Ring (19522, -3.25 DPS) [rep]; Underworld Band (1980, -3.46 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (640.9 DPS) | yes | Uther's Strength (11302, -4.33 DPS, sim-verified) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (640.9 DPS) | yes | Uther's Strength (11302, -2.94 DPS, sim-verified) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (640.9 DPS) | yes | Grand Marshal's Stave (234571, -0.14 DPS) [pvp]; Sageclaw (20070, -4.83 DPS) [rep]; Ironbark Staff (20069, -15.20 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (16.34 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.39 DPS) [pvp]; Tome of Shadow Force (19309, -0.90 DPS) [rep]; Spirit of Aquementas (11904, -4.51 DPS) [quest] |
| ranged | Greater Eternal Wand (249237) | Enchanting [crafted] | sim-verified (640.9 DPS) | yes | Brilliant Wand (249385, +0.00 DPS) [crafted]; Lesser Eternal Wand (249232, -0.56 DPS) [crafted]; Dreambough Wand (249234, -1.13 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Greater Eternal Wand

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 58.9. Weights run: 5.0s. Verify run: 2.4s. 173 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.176 ± 0.006, crit=0.059 ± 0.001 per rating point (14 rating = 1%, 0.833 per %), hit=0.174 ± 0.007 per rating point (10 rating = 1%, 1.740 per %), spell_haste=0.550 ± 0.088, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.821 ± 0.001, fire_power=0.179 ± 0.000, spirit=0.320 ± 0.063

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.0 spell_power points (2.56 DPS) | yes | Pristine Circlet (253949, -0.87 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.77 DPS) [crafted]; Flame Circlet (253953, -2.00 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.0 spell_power points (0.27 DPS) | yes | Scholarly Pendant (277203, -0.09 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.6 spell_power points (1.86 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.77 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (58.9 DPS) | yes | Battle Healer's Cloak (20427, -0.41 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.65 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.78 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.4 spell_power points (3.50 DPS) | yes | Filigreed Shadow Gown (253907, -1.09 DPS, sim-verified) [crafted]; Scholarly Robes (2034, -1.57 DPS) [dungeon]; Filigreed Pristine Gown (253901, -1.57 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.85 DPS) | yes | Windsong Bangles (263336, -0.57 DPS) [quest]; Crystalline Cuffs (14148, -0.58 DPS) [dungeon]; Owlbeard Bracers (16981, -0.66 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.98 DPS) | yes | Shadow Gloves (253919, -0.27 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.28 DPS) [world]; Blight Gloves (279877, -0.57 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.5 spell_power points (1.83 DPS) | yes | Pristine Sash (253925, -0.50 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted]; Novice Ardent's Sash (253887, -0.84 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 11.4 spell_power points (3.22 DPS) | yes | Filigreed Pristine Leggings (253937, -0.87 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.24 DPS) [dungeon]; Colorful Kilt (10048, -1.35 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.7 spell_power points (2.18 DPS) | yes | Shadow Boots (253895, -0.64 DPS) [crafted]; Red Woolen Boots (4313, -0.92 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.18 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.42 DPS) | yes | Lavishly Jeweled Ring (1156, -1.12 DPS) [dungeon]; Furlfeather Ring (286745, -1.13 DPS) [world]; Loop of Sacrifice (281673, -1.17 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 1.9 spell_power points (0.54 DPS) | yes | Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Furlfeather Ring (286745, -0.26 DPS) [world]; Loop of Sacrifice (281673, -0.30 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 31.2 spell_power points (8.83 DPS) | yes | Evocator's Blade (2567, -0.35 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.57 DPS) [quest]; Crescent Staff (6505, -1.12 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-verified (58.9 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.28 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Staff of the Blessed Seer; ranged: Sizzle Stick

No-known-source sample (15 of 173, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword

### Band 30 (orc, 25522300100001000-0000000000000000000-0000000000000000)

Set DPS (verified): 100.3. Weights run: 5.1s. Verify run: 9.3s. 296 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.209 ± 0.005, crit=0.049 ± 0.001 per rating point (14 rating = 1%, 0.681 per %), hit=0.136 ± 0.005 per rating point (10 rating = 1%, 1.355 per %), spell_haste=0.383 ± 0.077, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (-0.046 ± 0.085)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Circlet (253981, -0.42 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.56 DPS) [crafted] |
| neck | Stygian Bone Amulet (6695) | Razorfen Kraul: Blind Hunter (rare) [dungeon] | 5.0 spell_power points (1.40 DPS) | yes | Roadwatcher's Confidence (281265, -1.53 DPS, sim-verified) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.51 DPS) | yes | Death Speaker Mantle (6685, -0.84 DPS) [dungeon]; Chestnut Mantle (17695, -1.12 DPS) [quest]; Invoker's Mantle (215365, -1.66 DPS, sim-verified) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.28 DPS) [crafted]; Battle Healer's Cloak (19529, -0.28 DPS) [rep]; Hillman's Cloak (3719, -0.36 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Gown (253967, -0.23 DPS) [crafted]; Tree Bark Jacket (1486, -0.34 DPS, sim-verified) [dungeon]; Green Silk Armor (7065, -0.68 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.51 DPS) | yes | Phoenix Bindings (210781, -1.85 DPS) [crafted]; Windsong Bangles (263336, -2.23 DPS) [quest]; Tabitha's Cuffs (251486, -2.49 DPS, sim-verified) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Serpent Gloves (5970, -0.10 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.38 DPS) [world]; Hands of Darkness (7047, -0.84 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warsong Sash (16975, -0.23 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.56 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.84 DPS) [dungeon] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Leggings (253943, -0.68 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.76 DPS, sim-verified) [dungeon]; Necromancer Leggings (2277, -0.91 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.6 spell_power points (2.97 DPS) | yes | Spidersilk Boots (4320, -1.01 DPS) [crafted]; Gilded Slippers (254001, -1.01 DPS) [crafted]; Glinteye Slippers (273024, -2.18 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.95 DPS) | yes | Solid Gold Band (279845, -0.84 DPS) [quest]; Electrocutioner Lagnut (9447, -1.12 DPS) [dungeon]; Ring of Power Regulation (273030, -2.73 DPS, sim-verified) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Solid Gold Band (279845, -0.56 DPS) [quest]; Ring of Power Regulation (273030, -0.61 DPS, sim-verified) [dungeon]; Electrocutioner Lagnut (9447, -0.84 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.12 DPS) [dungeon]; Spine of the Devourer (273040, -1.12 DPS) [world]; Tendonscraper Dagger (273053, -2.01 DPS, sim-verified) [world] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Soran'ruk (6898, -0.35 DPS, sim-verified) [quest]; Orb of Mistmantle (13031, -0.59 DPS) [world_drop]; Orb of Souls (249395, -0.84 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Mystic Wand (217287, -0.23 DPS) [crafted]; Gravestone Scepter (7001, -0.26 DPS) [quest]; Wand of Eventide (5214, -0.31 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Stygian Bone Amulet; shoulder: Bloodmage Mantle; back: Windsong Drape; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Mystic Insight; ranged: Thunderwood

No-known-source sample (15 of 296, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (orc, 25522300130201311-0000000000000000000-0000000000000000)

Set DPS (verified): 155.1. Weights run: 5.2s. Verify run: 5.7s. 379 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.141 ± 0.010, crit=0.194 ± 0.003 per rating point (14 rating = 1%, 2.715 per %), hit=0.367 ± 0.014 per rating point (10 rating = 1%, 3.668 per %), spell_haste=2.284 ± 0.147, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.850 ± 0.001, fire_power=0.150 ± 0.000, spirit=not significant (-0.000 ± 0.083)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.12 DPS) | yes | Living Cowl (5608, -2.73 DPS, sim-verified) [world]; Augural Shroud (2620, -2.91 DPS) [world]; Filigreed Shadow Circlet (253981, -3.08 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Eye of Mordresh (10769, -0.39 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.38 DPS) [dungeon]; Darkspear Warding Pendant (272074, -2.38 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.3 spell_power points (3.48 DPS) | yes | Green Silken Shoulders (7057, -0.24 DPS) [crafted]; Inquisitor's Shawl (19507, -0.49 DPS) [dungeon]; Berylline Pads (4197, -0.63 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.10 DPS) [crafted]; Battle Healer's Cloak (19528, -0.34 DPS) [rep]; Long Silken Cloak (4326, -0.65 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.8 spell_power points (7.74 DPS) | yes | Elemental Raiment (9434, -0.62 DPS) [world_drop]; Dreamweave Vest (10021, -1.21 DPS) [crafted]; Shadoweave Robe (10004, -1.69 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.05 DPS) | yes | Condor Bracers (15864, -0.72 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.31 DPS) [quest]; Earthen Silk Cuffs (254019, -1.69 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (6.29 DPS) | yes | Shadoweave Gloves (10023, -0.97 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.21 DPS) [crafted]; Black Handwraps (254027, -1.92 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.6 spell_power points (4.93 DPS) | yes | Black Cord (254043, -0.52 DPS) [crafted]; Star Belt (4329, -0.53 DPS) [crafted]; Warsong Sash (16975, -1.21 DPS) [quest] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.8 spell_power points (6.05 DPS) | yes | Red Mageweave Pants (10009, -0.73 DPS) [crafted]; Shadow Leggings (253993, -1.68 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.98 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.13 DPS) | yes | Black Slippers (254007, -4.15 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -5.43 DPS) [crafted]; Spidersilk Boots (4320, -5.57 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.6 spell_power points (4.27 DPS) | yes | Advisor's Ring (19520, -1.22 DPS) [rep]; Reedknot Ring (9622, -1.90 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.7 spell_power points (3.63 DPS) | yes | Advisor's Ring (19520, -0.31 DPS, sim-verified) [rep]; Reedknot Ring (9622, -1.25 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -6.88 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -7.86 DPS) [dungeon]; Swinetusk Shank (6691, -9.59 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.31 DPS) [quest]; Greater Mystic Wand (217287, -0.34 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (orc, 25522300130201351-2004000000000000000-0000000000000000)

Set DPS (verified): 215.6. Weights run: 5.3s. Verify run: 8.6s. 478 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.147 ± 0.012, crit=0.250 ± 0.004 per rating point (14 rating = 1%, 3.497 per %), hit=0.449 ± 0.017 per rating point (10 rating = 1%, 4.491 per %), spell_haste=2.795 ± 0.188, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.859 ± 0.001, fire_power=0.141 ± 0.000, spirit=not significant (0.230 ± 0.084)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.99 DPS) | yes | Shadoweave Mask (10025, -1.82 DPS) [crafted]; Red Mageweave Headband (10033, -1.87 DPS) [crafted]; Dreamweave Circlet (10041, -2.36 DPS, sim-verified) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.6 spell_power points (3.55 DPS) | yes | Mindburst Medallion (11196, -0.58 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.27 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.45 DPS) [world_drop] |
| shoulder | Netherpearl Shoulders (254055) | Tailoring [crafted] | 17.2 spell_power points (6.35 DPS) | yes | Rotgrip Mantle (17732, -0.79 DPS, sim-verified) [dungeon]; Shadoweave Shoulders (10028, -0.95 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.60 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.48 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.25 DPS) [crafted]; Nightfall Drape (12465, -2.69 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.0 spell_power points (8.89 DPS) | yes | Acumen Robes (17775, -0.77 DPS) [quest]; Elemental Raiment (9434, -1.12 DPS) [world_drop]; Dreamweave Vest (10021, -1.81 DPS, sim-verified) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 13.3 spell_power points (4.94 DPS) | yes | Nethergeld Cuffs (254061, -1.58 DPS, sim-verified) [crafted]; Spidertank Oilrag (9448, -1.61 DPS) [dungeon]; Bloodband Bracers (11469, -2.34 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.2 spell_power points (7.47 DPS) | yes | Black Gloves (254101, -1.15 DPS, sim-verified) [crafted]; Runecloth Gloves (13863, -1.78 DPS) [crafted]; Shadoweave Gloves (10023, -1.80 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 17.6 spell_power points (6.53 DPS) | yes | Satyrmane Sash (17755, -0.80 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.13 DPS) [rep]; Ghostweave Cord (254073, -1.61 DPS, sim-verified) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 24.1 spell_power points (8.91 DPS) | yes | Wizardweave Leggings (14132, -2.11 DPS, sim-verified) [crafted]; Shadoweave Pants (10002, -2.24 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.53 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.6 spell_power points (9.47 DPS) | yes | Black Sandals (254113, -3.33 DPS, sim-verified) [crafted]; Shadoweave Boots (10031, -4.18 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.36 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.81 DPS) | yes | Advisor's Ring (19519, -0.37 DPS) [rep]; Philanthropist's Ring (281635, -0.84 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.8 spell_power points (4.72 DPS) | yes | Advisor's Ring (19519, -0.28 DPS) [rep]; Philanthropist's Ring (281635, -0.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.33 DPS) [quest]; Ankh of Life (1713, -0.48 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.65 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.97 DPS, sim-verified) [quest]; Ankh of Life (1713, -1.20 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.37 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -1.64 DPS, sim-verified) [world]; Kindling Stave (11750, -2.48 DPS) [dungeon]; Inventor's Focal Sword (17719, -3.14 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.37 DPS) [crafted]; Thunderwood (13062, -0.74 DPS) [world_drop]; Twisted Nether Wand (249144, -0.74 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Netherpearl Shoulders; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger1: Band of the Unicorn; finger2: Underworld Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Lesser Eternal Wand

No-known-source sample (15 of 478, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (orc, 25522300130201351-2005220000000000000-0050000000000000)

Set DPS (verified): 345.2. Weights run: 5.5s. Verify run: 22.9s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.399 ± 0.016, crit=0.199 ± 0.003 per rating point (14 rating = 1%, 2.791 per %), hit=0.386 ± 0.016 per rating point (10 rating = 1%, 3.865 per %), spell_haste=not significant (-0.003 ± 0.235), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.844 ± 0.001, fire_power=0.156 ± 0.000, spirit=not significant (-0.558 ± 0.244)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.50 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Felcloth Hood (14111, -0.59 DPS) [crafted]; Champion's Dreadweave Cowl (227090, -1.13 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.0 spell_power points (4.22 DPS) | yes | Glowing Eye of Mordresh (10769, -1.76 DPS) [dungeon]; Triune Amulet (7722, -1.99 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -2.11 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (8.80 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Felcloth Shoulders (14112, -1.62 DPS, sim-verified) [crafted]; Mantle of the Timbermaw (19050, -1.83 DPS) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -0.51 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.21 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -1.79 DPS, sim-verified) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.0 spell_power points (16.19 DPS) | yes | Robe of Winter Night (14136, -4.88 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -4.93 DPS) [pvp]; Felcloth Robe (14106, -5.50 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.0 spell_power points (7.74 DPS) | yes | Runecloth Cuffs (254123, -3.87 DPS) [crafted]; Netherpearl Cuffs (254067, -3.88 DPS) [crafted]; Spidertank Oilrag (9448, -4.57 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (7.74 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Gloves (17577, -0.35 DPS) [pvp]; Dreamweave Gloves (10019, -1.29 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Waistcord (254087, -0.27 DPS) [crafted]; Frostwolf Cloth Belt (19090, -0.38 DPS, sim-verified) [rep]; Ghostweave Cord (254073, -0.98 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 28.0 spell_power points (9.85 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -4.41 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.45 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.70 DPS) [crafted]; Omnicast Boots (11822, -1.41 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.0 spell_power points (6.33 DPS) | yes | Band of the Unicorn (7553, -1.76 DPS) [world_drop]; Deep Rooted Ring (19109, -2.11 DPS) [world]; Underworld Band (1980, -2.18 DPS) [world_drop] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 14.0 spell_power points (4.93 DPS) | yes | Band of the Unicorn (7553, -0.35 DPS) [world_drop]; Deep Rooted Ring (19109, -0.70 DPS) [world]; Underworld Band (1980, -0.77 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -3.10 DPS, sim-verified) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.46 DPS, sim-verified) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindfang (20214, -2.49 DPS) [rep]; High Warlord's War Staff (234549, -3.97 DPS) [pvp]; Ironbark Staff (20220, -13.18 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (10.21 DPS) | yes | Tome of Shadow Force (19309, -0.89 DPS, sim-verified) [rep]; High Warlord's Tome of Destruction (234563, -1.41 DPS) [pvp]; Spirit of Aquementas (11904, -2.82 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Eternal Wand (249237, +0.00 DPS) [crafted]; Lesser Eternal Wand (249232, -0.35 DPS) [crafted]; Dreambough Wand (249234, -0.70 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; back: Fel Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Advisor's Ring; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (orc, 25520010135001050-0005200000000000000-0550004000000000)

Set DPS (verified): 635.3. Weights run: 5.9s. Verify run: 21.4s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.254 ± 0.032, crit=0.677 ± 0.017 per rating point (14 rating = 1%, 9.478 per %), hit=1.053 ± 0.066 per rating point (10 rating = 1%, 10.528 per %), spell_haste=5.601 ± 0.686, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.884 ± 0.001, fire_power=0.116 ± 0.000, spirit=not significant (-0.019 ± 0.180)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 29.1 spell_power points (16.37 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -4.25 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.5 spell_power points (8.19 DPS) | yes | Triune Amulet (7722, -3.68 DPS) [dungeon]; Mindburst Medallion (11196, -3.95 DPS) [quest]; Glowing Eye of Mordresh (10769, -4.07 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.8 spell_power points (16.78 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.26 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -2.82 DPS) [pvp]; Argent Shoulders (19059, -3.82 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-verified (635.3 DPS) | yes | Spritecaster Cape (11623, -2.99 DPS, sim-verified) [dungeon]; Deep Woodlands Cloak (19121, -4.76 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -4.92 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.3 spell_power points (27.20 DPS) | yes | Robe of Winter Night (14136, -4.37 DPS, sim-verified) [crafted]; Warlord's Dreadweave Robe (231591, -5.74 DPS) [pvp]; Felcloth Robe (14106, -9.28 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.0 spell_power points (13.54 DPS) | yes | Runecloth Cuffs (254123, -5.91 DPS) [crafted]; Netherpearl Cuffs (254067, -6.06 DPS) [crafted]; Arena Wristguards (18709, -8.20 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.5 spell_power points (17.18 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -4.78 DPS) [pvp]; Earth Warder's Gloves (21318, -8.08 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 25.0 spell_power points (14.09 DPS) | yes | Black Waistcord (254087, -3.34 DPS) [crafted]; Frostwolf Cloth Belt (19090, -4.13 DPS, sim-verified) [rep]; Satyrmane Sash (17755, -4.77 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 32.8 spell_power points (18.49 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -0.86 DPS) [pvp]; Felcloth Pants (14107, -6.70 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.52 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.13 DPS) [crafted]; Omnicast Boots (11822, -2.84 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 20.0 spell_power points (11.27 DPS) | yes | Deep Rooted Ring (19109, -3.08 DPS) [world]; Advisor's Ring (19518, -3.38 DPS) [rep]; Underworld Band (1980, -3.59 DPS) [world_drop] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 19.8 spell_power points (11.14 DPS) | yes | Advisor's Ring (19518, -3.25 DPS) [rep]; Underworld Band (1980, -3.46 DPS) [world_drop]; Deep Rooted Ring (19109, -4.27 DPS, sim-verified) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (635.3 DPS) | yes | Uther's Strength (11302, -5.07 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (635.3 DPS) | yes | Uther's Strength (11302, -0.77 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (635.3 DPS) | yes | High Warlord's War Staff (234549, -0.14 DPS) [pvp]; Mindfang (20214, -4.83 DPS) [rep]; Ironbark Staff (20220, -14.27 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (16.34 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.39 DPS) [pvp]; Tome of Shadow Force (19309, -0.90 DPS) [rep]; Spirit of Aquementas (11904, -4.51 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (635.3 DPS) | yes | Lesser Eternal Wand (249232, -1.56 DPS) [crafted]; Dreambough Wand (249234, -2.13 DPS) [crafted]; Greater Eternal Wand (249237, -2.77 DPS, sim-verified) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

