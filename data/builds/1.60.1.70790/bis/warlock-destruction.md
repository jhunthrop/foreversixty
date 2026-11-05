# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 60.6. Weights run: 5.0s. Verify run: 2.8s. 182 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.146 ± 0.005, crit=0.057 ± 0.001 per rating point (14 rating = 1%, 0.794 per %), hit=0.167 ± 0.006 per rating point (10 rating = 1%, 1.675 per %), spell_haste=0.335 ± 0.071, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.835 ± 0.001, fire_power=0.165 ± 0.000, spirit=not significant (0.151 ± 0.054)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.2 spell_power points (2.84 DPS) | yes | Pristine Circlet (253949, -0.97 DPS, sim-verified) [crafted]; Flame Circlet (253953, -2.28 DPS) [crafted]; Shadow Goggles (4373, -2.33 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.09 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.3 spell_power points (1.95 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.71 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.36 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.49 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.67 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.4 spell_power points (3.83 DPS) | yes | Filigreed Shadow Gown (253907, -1.03 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.93 DPS) [crafted]; Scholarly Robes (2034, -2.35 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Repurposed Hair Band (281256, -0.04 DPS) [quest]; Bright Bracers (3647, -0.05 DPS) [world_drop]; Windsong Bangles (263336, -0.19 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.16 DPS) | yes | Shadow Gloves (253919, -0.22 DPS) [crafted]; Gnoll Casting Gloves (892, -0.31 DPS) [world]; Pristine Gloves (253913, -0.79 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.4 spell_power points (1.98 DPS) | yes | Pristine Sash (253925, -0.56 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.62 DPS) [crafted]; Novice Ardent's Sash (253887, -1.09 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.7 spell_power points (3.29 DPS) | yes | Filigreed Pristine Leggings (253937, -0.97 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.13 DPS) [dungeon]; Colorful Kilt (10048, -1.52 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (2.34 DPS) | yes | Shadow Boots (253895, -0.53 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.87 DPS) [crafted]; Pristine Boots (253889, -1.28 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.3 spell_power points (1.63 DPS) | yes | Furlfeather Ring (286745, -1.32 DPS) [world]; Band of Purification (12996, -1.35 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.36 DPS) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.54 DPS) | yes | Band of Purification (12996, -1.26 DPS) [world_drop]; Furlfeather Ring (286745, -1.27 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -1.27 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.91 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.56 DPS) [dungeon]; Witching Stave (1484, -1.78 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.2 spell_power points (0.97 DPS) | yes | Antipodean Rod (2879, -0.67 DPS) [world_drop]; Onu's Patience (281292, -0.74 DPS) [quest]; Mathystra Channeling Orb (281291, -0.75 DPS) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.31 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 182, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 99.5. Weights run: 5.1s. Verify run: 6.3s. 309 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.169 ± 0.008, crit=0.142 ± 0.004 per rating point (14 rating = 1%, 1.986 per %), hit=0.244 ± 0.010 per rating point (10 rating = 1%, 2.439 per %), spell_haste=not significant (0.331 ± 0.127), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.815 ± 0.001, fire_power=0.185 ± 0.000, spirit=not significant (0.061 ± 0.078)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Shadow Circlet (253981, -0.31 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.76 DPS) [crafted]; Shadow Circlet (253955, -0.77 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Seal of the Expedition (271920, -0.02 DPS) [rep]; Archmagister's Faceted Pendant (273049, -0.20 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (3.38 DPS) | yes | Invoker's Mantle (215365, -0.86 DPS) [crafted]; Death Speaker Mantle (6685, -0.87 DPS, sim-verified) [dungeon]; Moonlit Amice (11884, -1.13 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.24 DPS) [crafted]; Caretaker's Cape (19533, -0.24 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (4.18 DPS) | yes | Necrology Robes (2292, -0.24 DPS) [dungeon]; Shadow Gown (253967, -0.31 DPS) [crafted]; Beguiler Robes (7728, -0.48 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.89 DPS) | yes | Glowing Magical Bracelets (13106, -2.46 DPS) [world_drop]; Nightsky Wristbands (6407, -2.51 DPS) [world_drop]; Phoenix Bindings (210781, -2.69 DPS, sim-verified) [crafted] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.5 spell_power points (2.74 DPS) | yes | Hands of Darkness (7047, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.49 DPS) [dungeon]; Shadow Gloves (253919, -0.74 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (3.70 DPS) | yes | Belt of Arugal (6392, -0.55 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -1.13 DPS) [dungeon]; Invoker's Cord (215366, -1.18 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Leggings (253993, -0.15 DPS, sim-verified) [crafted]; Filigreed Shadow Leggings (253943, -0.75 DPS) [crafted]; Necromancer Leggings (2277, -0.81 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 12.0 spell_power points (3.87 DPS) | yes | Gilded Slippers (254001, -1.15 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.40 DPS) [crafted]; Glinteye Slippers (273024, -1.50 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.25 DPS) | yes | Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Minor Channeling Ring (1449, -0.53 DPS) [quest]; Electrocutioner Lagnut (9447, -1.15 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (2.20 DPS) | yes | Ring of Power Regulation (273030, -0.27 DPS) [dungeon]; Minor Channeling Ring (1449, -0.48 DPS) [quest]; Electrocutioner Lagnut (9447, -1.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.36 DPS) [dungeon]; Spine of the Devourer (273040, -1.36 DPS) [world]; Tendonscraper Dagger (273053, -1.56 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.89 DPS) | yes | Orb of Mystic Insight (249394, -0.36 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.24 DPS) [world_drop]; Orb of Souls (249395, -1.49 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, -0.30 DPS) [quest]; Wand of Eventide (5214, -0.32 DPS) [world_drop]; Greater Mystic Wand (217287, -0.32 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 168.7. Weights run: 5.3s. Verify run: 5.3s. 399 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.186 ± 0.011, crit=0.108 ± 0.003 per rating point (14 rating = 1%, 1.515 per %), hit=0.220 ± 0.013 per rating point (10 rating = 1%, 2.200 per %), spell_haste=not significant (-0.269 ± 0.163), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.453 ± 0.001, fire_power=0.547 ± 0.001, spirit=not significant (0.209 ± 0.168)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.72 DPS) | yes | Papal Fez (9431, -2.31 DPS, sim-verified) [dungeon]; Living Cowl (5608, -2.56 DPS) [world]; Augural Shroud (2620, -2.61 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.5 spell_power points (3.03 DPS) | yes | Glowing Eye of Mordresh (10769, -0.31 DPS) [dungeon]; Necklace of Calisea (1714, -2.14 DPS) [dungeon]; Amberglow Talisman (10824, -2.36 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (3.42 DPS) | yes | Green Silken Shoulders (7057, -0.20 DPS) [crafted]; Mantle of Doan (7712, -0.28 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.40 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (168.7 DPS) | yes | Long Silken Cloak (4326, -0.42 DPS) [crafted]; Guardian Cloak (5965, -0.42 DPS) [crafted]; Caretaker's Cape (19532, -0.74 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.2 spell_power points (7.74 DPS) | yes | Dreamweave Vest (10021, -0.50 DPS) [crafted]; Elemental Raiment (9434, -1.01 DPS) [world_drop]; Robe of Power (7054, -2.00 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.88 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Phoenix Bindings (210781, -0.61 DPS) [crafted]; Condor Bracers (15864, -0.64 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.2 spell_power points (6.47 DPS) | yes | Black Mageweave Gloves (10003, -1.90 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.35 DPS) [crafted]; Gilded Handwraps (254021, -3.49 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.7 spell_power points (4.72 DPS) | yes | Star Belt (4329, -0.29 DPS) [crafted]; Gilded Cord (254037, -1.28 DPS) [crafted]; Belt of Arugal (6392, -1.33 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.2 spell_power points (5.20 DPS) | yes | Gaze Dreamer Pants (6903, -0.75 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.68 DPS) [dungeon]; Black Mageweave Leggings (9999, -1.70 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.5 spell_power points (8.15 DPS) | yes | Gilded Slippers (254001, -5.23 DPS) [crafted]; Glinteye Slippers (273024, -5.51 DPS) [dungeon]; Fiery Slippers (254005, -6.23 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.9 spell_power points (3.50 DPS) | yes | Ring of Forlorn Spirits (2043, -0.94 DPS) [quest]; Underworld Band (1980, -1.17 DPS) [world_drop]; Reedknot Ring (9622, -1.26 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (2.88 DPS) | yes | Ring of Forlorn Spirits (2043, -0.32 DPS) [quest]; Underworld Band (1980, -0.55 DPS) [world_drop]; Reedknot Ring (9622, -0.64 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (168.7 DPS) | yes | Hypnotic Blade (7714, -7.37 DPS) [dungeon]; Illusionary Rod (7713, -7.76 DPS, sim-verified) [dungeon]; Swinetusk Shank (6691, -9.06 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (168.7 DPS) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.32 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 399, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 289.5. Weights run: 5.4s. Verify run: 8.6s. 506 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.427 ± 0.016, crit=0.166 ± 0.004 per rating point (14 rating = 1%, 2.322 per %), hit=0.341 ± 0.019 per rating point (10 rating = 1%, 3.409 per %), spell_haste=not significant (-0.920 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.425 ± 0.000, fire_power=0.575 ± 0.001, spirit=not significant (0.544 ± 0.267)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 34.4 spell_power points (14.23 DPS) | yes | Dreamweave Circlet (10041, -2.70 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.85 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.08 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 11.8 spell_power points (4.88 DPS) | yes | Mindburst Medallion (11196, -0.22 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.57 DPS) [dungeon]; Lei of Lilies (1315, -1.51 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherflame Shoulders (254053, -1.75 DPS) [crafted]; Nethergeld Shoulders (254049, -2.07 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.57 DPS, sim-verified) [crafted]; Cindercloth Cloak (14044, -3.69 DPS) [crafted]; Caretaker's Cape (19531, -3.95 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 29.5 spell_power points (12.17 DPS) | yes | Acumen Robes (17775, -0.79 DPS) [quest]; Robe of the Magi (1716, -0.90 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -2.43 DPS) [world] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 13.2 spell_power points (5.45 DPS) | yes | Nethergeld Cuffs (254061, -0.20 DPS) [crafted]; Netherpearl Cuffs (254067, -0.81 DPS) [crafted]; Bloodband Bracers (11469, -1.12 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.5 spell_power points (9.71 DPS) | yes | Fiery Gloves (254099, -1.64 DPS) [crafted]; Gilded Gloves (254095, -2.02 DPS) [crafted]; Runecloth Gloves (13863, -2.56 DPS, sim-verified) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | sim-verified (289.5 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, -0.71 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.76 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 27.3 spell_power points (11.27 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.81 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -3.23 DPS) [vendor]; Senior Designer's Pantaloons (11841, -4.54 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (11.49 DPS) | yes | Vinerot Sandals (17748, -5.17 DPS) [dungeon]; Sergeant Major's Dreadweave Boots (220891, -5.19 DPS) [vendor]; Fiery Sandals (254111, -5.49 DPS, sim-verified) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.2 spell_power points (5.85 DPS) | yes | Philanthropist's Ring (281635, -0.84 DPS) [quest]; Lorekeeper's Ring (19523, -0.89 DPS) [rep]; Brainlash (6440, -2.08 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.37 DPS) | yes | Philanthropist's Ring (281635, -0.36 DPS) [quest]; Lorekeeper's Ring (19523, -0.41 DPS) [rep]; Brainlash (6440, -1.60 DPS) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -3.47 DPS) [world]; Kindling Stave (11750, -4.05 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -4.88 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.46 DPS) [crafted]; Dreambough Wand (249234, -0.87 DPS) [crafted]; Charged Lightning Rod (11860, -0.99 DPS) [quest] |

**New at 50:** head: Eye of Flame; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Dreamweave Vest; wrist: Netherflame Cuffs; waist: Ghostweave Cord; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Band of the Unicorn; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; ranged: Pyric Caduceus

No-known-source sample (15 of 506, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 485.2. Weights run: 5.4s. Verify run: 29.2s. 1299 eligible items had no known source.

5 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.185 ± 0.034, crit=0.456 ± 0.012 per rating point (14 rating = 1%, 6.390 per %), hit=0.920 ± 0.045 per rating point (10 rating = 1%, 9.201 per %), spell_haste=not significant (-0.969 ± 0.509), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.386 ± 0.001, fire_power=0.614 ± 0.001, spirit=not significant (0.110 ± 0.281)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -1.12 DPS) [dungeon]; Dark Advisor's Pendant (18691, -2.58 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.82 DPS, sim-verified) [crafted]; Burial Shawl (18681, -1.50 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.53 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.84 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -6.78 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.7 spell_power points (25.50 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.00 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.14 DPS) [pvp]; Alanna's Embrace (13314, -11.73 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.3 spell_power points (12.97 DPS) | yes | Runecloth Cuffs (254123, -6.10 DPS) [crafted]; Netherflame Cuffs (254065, -7.72 DPS) [crafted]; Earthenweave Cuffs (254131, -8.11 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -0.87 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -0.88 DPS) [rep]; Satyrmane Sash (17755, -3.02 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.6 spell_power points (24.94 DPS) | yes | Marshal's Dreadweave Leggings (231587, -3.27 DPS) [pvp]; Skyshroud Leggings (13170, -6.27 DPS, sim-verified) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -8.68 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.8 spell_power points (13.25 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.01 DPS) [crafted]; Omnicast Boots (11822, -1.36 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.18 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.44 DPS) [dungeon]; Naglering (11669, -12.35 DPS, sim-verified) [dungeon] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.12 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.37 DPS) [dungeon]; Don Julio's Band (19325, -2.34 DPS) [rep] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+10.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.99 DPS) [crafted]; Uther's Strength (11302, -3.21 DPS) [world_drop]; Ankh of Life (1713, -5.71 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.42 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -7.34 DPS) [crafted]; Uther's Strength (11302, -8.56 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.35 DPS) [pvp]; Sageclaw (20070, -5.31 DPS) [rep]; Ironbark Staff (20069, -18.03 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (15.51 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.85 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.27 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -4.28 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (485.2 DPS) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Greater Eternal Wand (249237, -0.69 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Blessed Band of Light; finger2: Maiden's Circle; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 25500000033000000-0000000000000000000-0050355103101351)

Set DPS (verified): 865.2. Weights run: 6.0s. Verify run: 30.1s. 1299 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.174 ± 0.030, crit=0.405 ± 0.011 per rating point (14 rating = 1%, 5.671 per %), hit=0.813 ± 0.044 per rating point (10 rating = 1%, 8.133 per %), spell_haste=not significant (-0.152 ± 0.616), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.382 ± 0.001, fire_power=0.618 ± 0.001, spirit=not significant (0.394 ± 0.216)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 34.6 spell_power points (21.12 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -2.90 DPS) [pvp]; Eye of Flame (3075, -4.19 DPS, sim-verified) [world_drop] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -0.54 DPS) [dungeon]; Triune Amulet (7722, -1.82 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Burial Shawl (18681, -1.31 DPS) [dungeon]; Argent Shoulders (19059, -3.56 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -4.96 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -7.05 DPS) [rep]; Crystalline Threaded Cape (20697, -8.41 DPS, sim-verified) [world] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (29.08 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.96 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -11.66 DPS) [pvp]; Alanna's Embrace (13314, -14.63 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.2 spell_power points (15.99 DPS) | yes | Runecloth Cuffs (254123, -8.20 DPS) [crafted]; Earthenweave Cuffs (254131, -8.69 DPS) [crafted]; Netherflame Cuffs (254065, -9.13 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -0.05 DPS) [quest]; Hands of Power (13253, -2.34 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -0.60 DPS) [rep]; Ghostweave Cord (254073, -1.45 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 45.1 spell_power points (27.56 DPS) | yes | Marshal's Dreadweave Leggings (231587, -2.92 DPS) [pvp]; Ghostweave Pants (14144, -6.85 DPS) [crafted]; Skyshroud Leggings (13170, -13.81 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.8 spell_power points (16.36 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.98 DPS) [crafted]; Omnicast Boots (11822, -2.85 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.48 DPS) [quest]; Don Mauricio's Band of Domination (22433, -3.01 DPS) [dungeon]; Naglering (11669, -13.93 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.30 DPS) [quest]; Don Mauricio's Band of Domination (22433, -2.82 DPS) [dungeon]; Flaming Band (12926, -3.95 DPS) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -2.86 DPS) [crafted]; Uther's Strength (11302, -3.67 DPS) [world_drop]; Ankh of Life (1713, -4.44 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.46 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -8.97 DPS) [crafted]; Uther's Strength (11302, -9.78 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.43 DPS) [pvp]; Staff of Jordan (873, -13.22 DPS, sim-verified) [world_drop]; Spire of Hakkar (10844, -20.57 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (865.2 DPS) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.07 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 59.9. Weights run: 5.0s. Verify run: 2.1s. 176 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.146 ± 0.005, crit=0.057 ± 0.001 per rating point (14 rating = 1%, 0.794 per %), hit=0.167 ± 0.006 per rating point (10 rating = 1%, 1.675 per %), spell_haste=0.335 ± 0.071, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.835 ± 0.001, fire_power=0.165 ± 0.000, spirit=not significant (0.151 ± 0.054)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 9.2 spell_power points (2.84 DPS) | yes | Pristine Circlet (253949, -0.99 DPS, sim-verified) [crafted]; Flame Circlet (253953, -2.28 DPS) [crafted]; Shadow Goggles (4373, -2.33 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.5 spell_power points (0.14 DPS) | yes | Scholarly Pendant (277203, -0.05 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.3 spell_power points (1.95 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.42 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.71 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (59.9 DPS) | yes | Battle Healer's Cloak (20427, -0.48 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.49 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.67 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.4 spell_power points (3.83 DPS) | yes | Filigreed Shadow Gown (253907, -1.22 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.93 DPS) [crafted]; Scholarly Robes (2034, -2.35 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.93 DPS) | yes | Windsong Bangles (263336, -0.62 DPS) [quest]; Mindthrust Bracers (1974, -0.70 DPS) [dungeon]; Owlbeard Bracers (16981, -0.86 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (2.16 DPS) | yes | Gnoll Casting Gloves (892, -0.31 DPS) [world]; Shadow Gloves (253919, -0.49 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.62 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.4 spell_power points (1.98 DPS) | yes | Pristine Sash (253925, -0.57 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.62 DPS) [crafted]; Novice Ardent's Sash (253887, -1.09 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.7 spell_power points (3.29 DPS) | yes | Filigreed Pristine Leggings (253937, -0.99 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.13 DPS) [dungeon]; Colorful Kilt (10048, -1.52 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.6 spell_power points (2.34 DPS) | yes | Shadow Boots (253895, -0.78 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.87 DPS) [crafted]; Pristine Boots (253889, -1.28 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.54 DPS) | yes | Band of Purification (12996, -1.26 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.27 DPS) [dungeon]; Loop of Sacrifice (281673, -1.32 DPS) [quest] |
| finger2 | Furlfeather Ring (286745) | Dustwind Eggtender [world] | 1.0 spell_power points (0.31 DPS) | yes | Lavishly Jeweled Ring (1156, -0.04 DPS) [dungeon]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Band of Purification (12996, -0.38 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 29.5 spell_power points (9.10 DPS) | yes | Evocator's Blade (2567, -0.24 DPS, sim-verified) [dungeon]; Gnarled Necromancer's Staff (251534, -0.62 DPS) [quest]; Crescent Staff (6505, -0.77 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-verified (59.9 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.31 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Furlfeather Ring; main_hand: Staff of the Blessed Seer; ranged: Sizzle Stick

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (orc, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 100.2. Weights run: 5.1s. Verify run: 5.2s. 302 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.056 ± 0.005, crit=0.068 ± 0.002 per rating point (14 rating = 1%, 0.947 per %), hit=0.107 ± 0.005 per rating point (10 rating = 1%, 1.070 per %), spell_haste=not significant (0.209 ± 0.060), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.801 ± 0.001, fire_power=0.199 ± 0.000, spirit=not significant (0.019 ± 0.049)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (100.2 DPS) | yes | Filigreed Shadow Circlet (253981, +0.00 DPS) [crafted]; Silk Headband (7050, -0.64 DPS) [crafted]; Shadow Circlet (253955, -0.70 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | sim-verified (100.2 DPS) | yes | Archmagister's Faceted Pendant (273049, +0.00 DPS) [world]; Crystal Starfire Medallion (5003, -0.01 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.01 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.5 spell_power points (2.89 DPS) | yes | Invoker's Mantle (215365, -0.68 DPS) [crafted]; Death Speaker Mantle (6685, -0.88 DPS) [dungeon]; Magician's Mantle (12998, -1.22 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, -0.02 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.30 DPS) [crafted]; Battle Healer's Cloak (19529, -0.30 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.96 DPS) | yes | Necrology Robes (2292, -0.46 DPS) [dungeon]; Shadow Gown (253967, -0.64 DPS) [crafted]; Beguiler Robes (7728, -0.97 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.74 DPS) | yes | Tabitha's Cuffs (251486, -1.83 DPS) [quest]; Phoenix Bindings (210781, -1.95 DPS) [crafted]; Owlbeard Bracers (16981, -2.40 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.6 spell_power points (2.32 DPS) | yes | Hands of Darkness (7047, -0.09 DPS) [crafted]; Serpent Gloves (5970, -0.18 DPS) [dungeon]; Jutebraid Gloves (10654, -0.40 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.2 spell_power points (3.40 DPS) | yes | Warsong Sash (16975, -0.05 DPS) [quest]; Belt of Arugal (6392, -0.58 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.96 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.2 spell_power points (3.71 DPS) | yes | Shadow Leggings (253993, -0.14 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.90 DPS) [crafted]; Necromancer Leggings (2277, -1.08 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.9 spell_power points (3.32 DPS) | yes | Gilded Slippers (254001, -1.04 DPS) [crafted]; Spidersilk Boots (4320, -1.12 DPS) [crafted]; Glinteye Slippers (273024, -1.15 DPS) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.13 DPS) | yes | Ring of Power Regulation (273030, -0.30 DPS) [dungeon]; Solid Gold Band (279845, -0.88 DPS) [quest]; Electrocutioner Lagnut (9447, -1.18 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.3 spell_power points (1.91 DPS) | yes | Ring of Power Regulation (273030, -0.09 DPS) [dungeon]; Solid Gold Band (279845, -0.67 DPS) [quest]; Electrocutioner Lagnut (9447, -0.96 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.24 DPS) [dungeon]; Spine of the Devourer (273040, -1.24 DPS) [world]; Tendonscraper Dagger (273053, -1.24 DPS) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.74 DPS) | yes | Orb of Mystic Insight (249394, -0.51 DPS) [crafted]; Orb of Mistmantle (13031, -1.25 DPS) [world_drop]; Orb of Souls (249395, -1.49 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wand of Eventide (5214, -0.30 DPS) [world_drop]; Greater Mystic Wand (217287, -0.30 DPS) [crafted]; Gravestone Scepter (7001, -0.35 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (orc, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 169.7. Weights run: 5.3s. Verify run: 6.3s. 392 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.188 ± 0.016, crit=0.223 ± 0.006 per rating point (14 rating = 1%, 3.118 per %), hit=0.399 ± 0.022 per rating point (10 rating = 1%, 3.988 per %), spell_haste=1.124 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.425 ± 0.000, fire_power=0.575 ± 0.001, spirit=not significant (0.508 ± 0.167)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.55 DPS) | yes | Whitemane's Chapeau (7720, -2.16 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -2.49 DPS) [dungeon]; Papal Fez (9431, -3.15 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 11.6 spell_power points (4.15 DPS) | yes | Glowing Eye of Mordresh (10769, -1.02 DPS, sim-verified) [dungeon]; Amberglow Talisman (10824, -2.33 DPS) [quest]; Necklace of Calisea (1714, -2.40 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Green Silken Shoulders (7057, -0.22 DPS) [crafted]; Earthen Silk Shoulders (254033, -0.23 DPS) [crafted]; Mantle of Doan (7712, -0.48 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Battle Healer's Cloak (19528, -0.54 DPS) [rep]; Glowing Thresher Cape (6901, -1.07 DPS) [dungeon]; Windsong Drape (15468, -1.08 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Vest (10021, -0.66 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.68 DPS) [world_drop]; Robe of Power (7054, -1.92 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.23 DPS) | yes | Condor Bracers (15864, -0.72 DPS) [quest]; Earthen Silk Cuffs (254019, -0.88 DPS) [crafted]; Phoenix Bindings (210781, -1.11 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.3 spell_power points (8.02 DPS) | yes | Black Mageweave Gloves (10003, -1.82 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -3.39 DPS) [crafted]; Earthen Silk Gloves (254017, -4.40 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.30 DPS, sim-verified) [crafted]; Fiery Cord (254041, -0.77 DPS) [crafted]; Gilded Cord (254037, -0.79 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Mageweave Leggings (9999, -0.41 DPS) [crafted]; Stoneweaver Leggings (9407, -0.59 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.04 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.6 spell_power points (9.90 DPS) | yes | Fiery Slippers (254005, -6.01 DPS) [crafted]; Gilded Slippers (254001, -6.18 DPS) [crafted]; Thoughtcast Boots (10578, -6.54 DPS, sim-verified) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.9 spell_power points (3.93 DPS) | yes | Reedknot Ring (9622, -1.42 DPS) [quest]; Snake Hoop (6750, -1.44 DPS) [quest]; Underworld Band (1980, -1.46 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.23 DPS) | yes | Snake Hoop (6750, -0.74 DPS) [quest]; Underworld Band (1980, -0.76 DPS) [world_drop]; Reedknot Ring (9622, -1.35 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.09 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -9.99 DPS) [dungeon]; Swinetusk Shank (6691, -11.79 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Wand of Eventide (5214, -0.36 DPS) [world_drop]; Greater Mystic Wand (217287, -0.36 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 287.7. Weights run: 5.4s. Verify run: 9.8s. 499 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.427 ± 0.016, crit=0.166 ± 0.004 per rating point (14 rating = 1%, 2.322 per %), hit=0.341 ± 0.019 per rating point (10 rating = 1%, 3.409 per %), spell_haste=not significant (-0.920 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.425 ± 0.000, fire_power=0.575 ± 0.001, spirit=not significant (0.544 ± 0.267)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 34.4 spell_power points (14.23 DPS) | yes | Dreamweave Circlet (10041, -2.39 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.85 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.08 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 11.8 spell_power points (4.88 DPS) | yes | Glowing Eye of Mordresh (10769, -0.57 DPS) [dungeon]; Mindburst Medallion (11196, -1.38 DPS, sim-verified) [quest]; Lei of Lilies (1315, -1.51 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherflame Shoulders (254053, -1.75 DPS) [crafted]; Nethergeld Shoulders (254049, -2.07 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.12 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -3.06 DPS) [crafted]; Cindercloth Cloak (14044, -3.69 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.11 DPS) [world_drop]; Dreamweave Vest (10021, -0.91 DPS, sim-verified) [crafted]; Embrace of the Wind Serpent (12462, -1.64 DPS) [world] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 13.2 spell_power points (5.45 DPS) | yes | Nethergeld Cuffs (254061, -0.20 DPS) [crafted]; Netherpearl Cuffs (254067, -0.81 DPS) [crafted]; Bloodband Bracers (11469, -1.12 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.5 spell_power points (9.71 DPS) | yes | Fiery Gloves (254099, -1.64 DPS) [crafted]; Gilded Gloves (254095, -2.02 DPS) [crafted]; Runecloth Gloves (13863, -2.99 DPS, sim-verified) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | sim-verified (287.7 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, -0.71 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.76 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 27.3 spell_power points (11.27 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.81 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -3.23 DPS) [vendor]; Senior Designer's Pantaloons (11841, -3.90 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (11.49 DPS) | yes | Vinerot Sandals (17748, -5.17 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -5.19 DPS) [vendor]; Fiery Sandals (254111, -6.28 DPS, sim-verified) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.2 spell_power points (5.85 DPS) | yes | Philanthropist's Ring (281635, -0.84 DPS) [quest]; Advisor's Ring (19519, -0.89 DPS) [rep]; Brainlash (6440, -2.08 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.37 DPS) | yes | Advisor's Ring (19519, -0.41 DPS) [rep]; Brainlash (6440, -1.60 DPS) [dungeon]; Philanthropist's Ring (281635, -2.18 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.28 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.23 DPS) [quest]; Rune of the Guard Captain (19120, -0.76 DPS, sim-verified) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.69 DPS, sim-verified) [dungeon]; Spire of Hakkar (10844, -3.47 DPS) [world]; Glowing Brightwood Staff (812, -4.88 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.46 DPS) [crafted]; Dreambough Wand (249234, -0.87 DPS) [crafted]; Charged Lightning Rod (11860, -0.99 DPS) [quest] |

**New at 50:** head: Eye of Flame; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Ghostweave Cord; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Pyric Caduceus

No-known-source sample (15 of 499, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 486.0. Weights run: 5.4s. Verify run: 30.2s. 1293 eligible items had no known source.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.154 ± 0.023, crit=0.230 ± 0.007 per rating point (14 rating = 1%, 3.215 per %), hit=0.530 ± 0.027 per rating point (10 rating = 1%, 5.303 per %), spell_haste=not significant (-0.915 ± 0.371), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.423 ± 0.001, fire_power=0.577 ± 0.001, spirit=not significant (0.425 ± 0.378)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 34.6 spell_power points (16.12 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Eye of Flame (3075, -1.87 DPS) [world_drop]; Ghostweave Hood (254135, -2.78 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Triune Amulet (7722, -1.19 DPS) [dungeon]; Star of Mystaria (12103, -1.60 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 28.0 spell_power points (13.02 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.17 DPS) [pvp]; Burial Shawl (18681, -2.47 DPS, sim-verified) [dungeon]; Mantle of the Timbermaw (19050, -2.68 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -2.35 DPS) [dungeon]; Frostwolf Advisor's Cloak (19085, -3.97 DPS) [rep]; Crystalline Threaded Cape (20697, -3.99 DPS, sim-verified) [world] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.4 spell_power points (22.05 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.44 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -8.99 DPS) [pvp]; Alanna's Embrace (13314, -10.82 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.2 spell_power points (12.19 DPS) | yes | Runecloth Cuffs (254123, -6.36 DPS) [crafted]; Earthenweave Cuffs (254131, -6.49 DPS) [crafted]; Netherflame Cuffs (254065, -7.21 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 29.5 spell_power points (13.72 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -3.66 DPS) [pvp]; Earth Warder's Gloves (21318, -5.25 DPS, sim-verified) [quest] |
| waist | Frostwolf Cloth Belt (19090) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, -0.25 DPS) [dungeon]; Ghostweave Cord (254073, -0.40 DPS) [crafted]; Defiler's Cloth Girdle (20163, -0.65 DPS) [rep] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -0.31 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.0 spell_power points (12.55 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.73 DPS) [crafted]; Omnicast Boots (11822, -2.39 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.40 DPS) [quest]; Don Mauricio's Band of Domination (22433, -3.45 DPS) [dungeon]; Naglering (11669, -12.43 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Orgrimmar (12545, -0.69 DPS, sim-verified) [quest]; Don Mauricio's Band of Domination (22433, -2.09 DPS) [dungeon]; Advisor's Ring (19518, -2.19 DPS) [rep] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+7.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.79 DPS) [world_drop]; Ankh of Life (1713, -3.21 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.36 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.25 DPS, sim-verified) [vendor]; Uther's Strength (11302, -7.45 DPS) [world_drop]; Ankh of Life (1713, -7.86 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindfang (20214, -4.19 DPS) [rep]; High Warlord's War Staff (234549, -4.31 DPS) [pvp]; Ironbark Staff (20220, -15.77 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (13.49 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.93 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.18 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.72 DPS) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (486.0 DPS) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, +0.00 DPS) [dungeon]; Brilliant Wand (249385, -0.28 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Frostwolf Cloth Belt; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 25500000033000000-0000000000000000000-0050355103101351)

Set DPS (verified): 858.5. Weights run: 6.0s. Verify run: 32.9s. 1293 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.326 ± 0.047, crit=0.641 ± 0.018 per rating point (14 rating = 1%, 8.978 per %), hit=1.212 ± 0.074 per rating point (10 rating = 1%, 12.122 per %), spell_haste=4.205 ± 1.024, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.364 ± 0.001, fire_power=0.636 ± 0.001, spirit=not significant (0.139 ± 0.253)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 33.7 spell_power points (23.15 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Eye of Flame (3075, -1.54 DPS, sim-verified) [world_drop] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.7 spell_power points (16.29 DPS) | yes | Star of Mystaria (12103, -4.32 DPS, sim-verified) [dungeon]; Archlight Talisman (15856, -5.81 DPS) [quest]; Emberfury Talisman (12929, -9.65 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.2 spell_power points (20.74 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Burial Shawl (18681, -3.43 DPS) [dungeon]; Argent Shoulders (19059, -5.53 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.62 DPS, sim-verified) [world]; Spritecaster Cape (11623, -9.57 DPS) [dungeon]; Shroud of Arcane Mastery (22330, -10.31 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.9 spell_power points (33.59 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.25 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -11.95 DPS) [pvp]; Alanna's Embrace (13314, -15.90 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.6 spell_power points (17.56 DPS) | yes | Runecloth Cuffs (254123, -7.77 DPS) [crafted]; Netherflame Cuffs (254065, -9.84 DPS) [crafted]; Nethergeld Cuffs (254061, -10.71 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.2 spell_power points (20.74 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -0.98 DPS) [dungeon]; Blood Guard's Dreadweave Handwraps (227099, -5.43 DPS) [pvp] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -2.52 DPS) [rep]; Clutch of Andros (13956, -4.54 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -2.30 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.0 spell_power points (17.14 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.72 DPS) [dungeon]; Dragonrider Boots (18102, -1.20 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -1.45 DPS) [world_drop]; Eye of Orgrimmar (12545, -1.64 DPS) [quest]; Naglering (11669, -15.07 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Orgrimmar (12545, -0.18 DPS) [quest]; Don Mauricio's Band of Domination (22433, -0.77 DPS) [dungeon]; Maiden's Circle (13001, -1.62 DPS, sim-verified) [world_drop] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.55 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -6.86 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -7.61 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (858.5 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.01 DPS) [pvp]; Staff of Jordan (873, -14.23 DPS, sim-verified) [world_drop]; Whiteout Staff (19101, -21.71 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.20 DPS) [dungeon]; Greater Eternal Wand (249237, -1.57 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

