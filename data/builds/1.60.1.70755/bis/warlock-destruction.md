# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 57.9. Weights run: 4.9s. Verify run: 4.1s. 179 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.188 ± 0.006, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.851 per %), hit=0.165 ± 0.007 per rating point (10 rating = 1%, 1.650 per %), spell_haste=not significant (0.352 ± 0.094), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.813 ± 0.001, fire_power=0.187 ± 0.000, spirit=0.300 ± 0.054

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.9 spell_power points (2.46 DPS) | yes | Pristine Circlet (253949, -0.81 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.71 DPS) [crafted]; Flame Circlet (253953, -1.90 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.6 spell_power points (0.17 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.7 spell_power points (1.84 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.22 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.74 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.48 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.61 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.73 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.3 spell_power points (3.40 DPS) | yes | Filigreed Shadow Gown (253907, -0.87 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.51 DPS) [crafted]; Scholarly Robes (2034, -1.58 DPS) [dungeon] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 spell_power points (0.28 DPS) | yes | Repurposed Hair Band (281256, -0.01 DPS) [quest]; Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Seer's Cuffs (3645, -0.06 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.93 DPS) | yes | Shadow Gloves (253919, -0.21 DPS) [crafted]; Gnoll Casting Gloves (892, -0.28 DPS) [world]; Pristine Gloves (253913, -0.67 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.4 spell_power points (1.78 DPS) | yes | Pristine Sash (253925, -0.46 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.52 DPS) [crafted]; Novice Ardent's Sash (253887, -0.82 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 11.3 spell_power points (3.11 DPS) | yes | Filigreed Pristine Leggings (253937, -0.81 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.18 DPS) [dungeon]; Colorful Kilt (10048, -1.31 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.8 spell_power points (2.14 DPS) | yes | Shadow Boots (253895, -0.64 DPS) [crafted]; Red Woolen Boots (4313, -0.65 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.15 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.4 spell_power points (1.48 DPS) | yes | Band of Purification (12996, -0.99 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.17 DPS) [dungeon]; Furlfeather Ring (286745, -1.21 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.38 DPS) | yes | Lavishly Jeweled Ring (1156, -1.07 DPS) [dungeon]; Furlfeather Ring (286745, -1.10 DPS) [world]; Band of Purification (12996, -1.20 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (57.9 DPS) | yes | Staff of the Blessed Seer (2271, -0.58 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.62 DPS) [dungeon]; Witching Stave (1484, -1.92 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.3 spell_power points (0.91 DPS) | yes | Onu's Patience (281292, -0.50 DPS) [quest]; Grayson's Torch (1172, -0.58 DPS) [quest]; Antipodean Rod (2879, -0.60 DPS) [world_drop] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.28 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 98.7. Weights run: 4.9s. Verify run: 4.3s. 303 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.167 ± 0.008, crit=0.142 ± 0.004 per rating point (14 rating = 1%, 1.985 per %), hit=0.259 ± 0.011 per rating point (10 rating = 1%, 2.585 per %), spell_haste=not significant (0.175 ± 0.126), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (0.196 ± 0.074)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.2 spell_power points (3.92 DPS) | yes | Filigreed Shadow Circlet (253981, -0.24 DPS) [crafted]; Silk Headband (7050, -1.02 DPS) [crafted]; Shadow Circlet (253955, -1.03 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 1.5 spell_power points (0.47 DPS) | yes | Crystal Starfire Medallion (5003, -0.06 DPS) [world_drop]; Pendant of Myzrael (4614, -0.09 DPS) [dungeon]; Glowing Green Talisman (5002, -0.09 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (3.38 DPS) | yes | Invoker's Mantle (215365, -0.86 DPS) [crafted]; Death Speaker Mantle (6685, -1.03 DPS, sim-verified) [dungeon]; Moonlit Amice (11884, -1.13 DPS) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-verified (98.7 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.07 DPS) [crafted]; Caretaker's Cape (19533, -0.07 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (4.18 DPS) | yes | Shadow Gown (253967, -0.14 DPS) [crafted]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Necrology Robes (2292, -0.23 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.90 DPS) | yes | Phoenix Bindings (210781, -2.01 DPS, sim-verified) [crafted]; Nightsky Wristbands (6407, -2.39 DPS) [world_drop]; Glowing Magical Bracelets (13106, -2.47 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.5 spell_power points (2.74 DPS) | yes | Hands of Darkness (7047, -0.06 DPS) [crafted]; Serpent Gloves (5970, -0.49 DPS) [dungeon]; Gloves of Meditation (4318, -0.69 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 spell_power points (3.70 DPS) | yes | Belt of Arugal (6392, -0.49 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -1.13 DPS) [dungeon]; Invoker's Cord (215366, -1.18 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 13.8 spell_power points (4.43 DPS) | yes | Shadow Leggings (253993, -0.06 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.96 DPS) [crafted]; Necromancer Leggings (2277, -1.21 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 12.6 spell_power points (4.05 DPS) | yes | Gilded Slippers (254001, -1.16 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.58 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.25 DPS) | yes | Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Minor Channeling Ring (1449, -0.54 DPS) [quest]; Electrocutioner Lagnut (9447, -0.85 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (2.20 DPS) | yes | Ring of Power Regulation (273030, -0.27 DPS) [dungeon]; Minor Channeling Ring (1449, -0.48 DPS) [quest]; Electrocutioner Lagnut (9447, -0.79 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (98.7 DPS) | yes | Hydrocane (9452, -1.54 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.54 DPS) [world]; Soulstaff (249393, -4.32 DPS, sim-verified) [crafted] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.90 DPS) | yes | Orb of Mystic Insight (249394, -0.36 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.07 DPS) [world_drop]; Orb of Souls (249395, -1.23 DPS) [crafted] |
| ranged | Gravestone Scepter (7001) | Blackfathom Villainy [quest] | sim-verified (98.7 DPS) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Wand of Eventide (5214, -0.16 DPS) [world_drop]; Greater Mystic Wand (217287, -0.16 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Gravestone Scepter

No-known-source sample (15 of 303, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 167.9. Weights run: 5.2s. Verify run: 3.7s. 386 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.204 ± 0.016, crit=0.227 ± 0.006 per rating point (14 rating = 1%, 3.176 per %), hit=0.471 ± 0.023 per rating point (10 rating = 1%, 4.712 per %), spell_haste=not significant (0.599 ± 0.243), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.438 ± 0.000, fire_power=0.562 ± 0.001, spirit=not significant (0.284 ± 0.166)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.31 DPS) | yes | Augural Shroud (2620, -2.77 DPS) [world]; Living Cowl (5608, -2.79 DPS) [world]; Papal Fez (9431, -3.03 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 10.0 spell_power points (3.48 DPS) | yes | Glowing Eye of Mordresh (10769, -1.11 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.29 DPS) [dungeon]; Amberglow Talisman (10824, -2.49 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.8 spell_power points (3.77 DPS) | yes | Green Silken Shoulders (7057, -0.21 DPS) [crafted]; Inquisitor's Shawl (19507, -0.41 DPS) [dungeon]; Mantle of Doan (7712, -0.84 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (167.9 DPS) | yes | Long Silken Cloak (4326, -0.59 DPS) [crafted]; Guardian Cloak (5965, -0.59 DPS) [crafted]; Caretaker's Cape (19532, -0.85 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.6 spell_power points (8.58 DPS) | yes | Elemental Raiment (9434, -1.27 DPS) [world_drop]; Dreamweave Vest (10021, -1.42 DPS, sim-verified) [crafted]; Robe of Power (7054, -2.06 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.13 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Phoenix Bindings (210781, -0.59 DPS) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.24 DPS) | yes | Red Mageweave Gloves (10018, -2.70 DPS) [crafted]; Black Mageweave Gloves (10003, -2.94 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -3.96 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.8 spell_power points (5.16 DPS) | yes | Star Belt (4329, -0.94 DPS, sim-verified) [crafted]; Gilded Cord (254037, -1.21 DPS) [crafted]; Fiery Cord (254041, -1.26 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.4 spell_power points (5.73 DPS) | yes | Gaze Dreamer Pants (6903, -1.23 DPS, sim-verified) [dungeon]; Stoneweaver Leggings (9407, -1.54 DPS) [dungeon]; Black Mageweave Leggings (9999, -1.56 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.0 spell_power points (9.05 DPS) | yes | Gilded Slippers (254001, -5.72 DPS) [crafted]; Wingborne Boots (15104, -5.97 DPS) [quest]; Fiery Slippers (254005, -6.25 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (3.84 DPS) | yes | Ring of Forlorn Spirits (2043, -1.05 DPS) [quest]; Underworld Band (1980, -1.35 DPS) [world_drop]; Snake Hoop (6750, -1.39 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (3.13 DPS) | yes | Underworld Band (1980, -0.64 DPS) [world_drop]; Snake Hoop (6750, -0.69 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.43 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (167.9 DPS) | yes | Illusionary Rod (7713, -8.53 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -9.58 DPS) [dungeon]; Swinetusk Shank (6691, -11.44 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (167.9 DPS) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.35 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 386, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 276.5. Weights run: 5.2s. Verify run: 10.7s. 485 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.224 ± 0.014, crit=0.167 ± 0.004 per rating point (14 rating = 1%, 2.338 per %), hit=0.337 ± 0.017 per rating point (10 rating = 1%, 3.369 per %), spell_haste=not significant (-0.676 ± 0.227), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.415 ± 0.000, fire_power=0.585 ± 0.001, spirit=not significant (0.531 ± 0.192)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -1.01 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -1.50 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.56 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 11.7 spell_power points (5.00 DPS) | yes | Glowing Eye of Mordresh (10769, -1.25 DPS) [dungeon]; Mindburst Medallion (11196, -1.77 DPS, sim-verified) [quest]; Amberglow Talisman (10824, -2.73 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.0 spell_power points (7.26 DPS) | yes | Nethergeld Shoulders (254049, -1.30 DPS) [crafted]; Flameseer Mantle (11310, -1.51 DPS) [dungeon]; Netherflame Shoulders (254053, -1.86 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.31 DPS, sim-verified) [crafted]; Icy Cloak (4327, -3.56 DPS) [crafted]; Caretaker's Cape (19531, -3.56 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Acumen Robes (17775, -1.07 DPS) [quest]; Dreamweave Vest (10021, -1.14 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -2.13 DPS) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 11.8 spell_power points (5.04 DPS) | yes | Nethergeld Cuffs (254061, -0.26 DPS) [crafted]; Netherpearl Cuffs (254067, -0.94 DPS) [crafted]; Arcane Runed Bracers (4744, -1.21 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.6 spell_power points (9.64 DPS) | yes | Fiery Gloves (254099, -2.14 DPS) [crafted]; Runecloth Gloves (13863, -2.22 DPS, sim-verified) [crafted]; Ghostweave Gloves (14142, -2.41 DPS) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | 19.8 spell_power points (8.46 DPS) | yes | Satyrmane Sash (17755, -1.53 DPS) [dungeon]; Star Belt (4329, -2.01 DPS) [crafted]; Highlander's Cloth Girdle (20098, -2.11 DPS) [rep] |
| legs | Senior Designer's Pantaloons (11841) | Blackrock Depths: General Angerforge [dungeon] | 23.1 spell_power points (9.86 DPS) | yes | Cindercloth Leggings (12256, -1.47 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -1.76 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.60 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.7 spell_power points (11.82 DPS) | yes | Fiery Sandals (254111, -5.97 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -6.11 DPS) [vendor]; Earthenweave Boots (254093, -6.22 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.54 DPS) | yes | Philanthropist's Ring (281635, -0.80 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.13 DPS) [quest]; Eye of Adaegus (5266, -2.25 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (5.12 DPS) | yes | Ring of Forlorn Spirits (2043, -1.71 DPS) [quest]; Philanthropist's Ring (281635, -1.78 DPS, sim-verified) [quest]; Eye of Adaegus (5266, -1.83 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Evonice's Landin' Pilla (18951, -0.29 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.75 DPS) [quest]; Ankh of Life (1713, -2.10 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -4.09 DPS, sim-verified) [world]; Kindling Stave (11750, -4.19 DPS) [dungeon]; Inventor's Focal Sword (17719, -5.33 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.51 DPS) [crafted]; Dreambough Wand (249234, -0.94 DPS) [crafted]; Twisted Nether Wand (249144, -1.36 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Ghostweave Cord; legs: Senior Designer's Pantaloons; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Pyric Caduceus

No-known-source sample (15 of 485, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 461.3. Weights run: 5.3s. Verify run: 22.7s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.194 ± 0.033, crit=0.442 ± 0.012 per rating point (14 rating = 1%, 6.185 per %), hit=0.981 ± 0.045 per rating point (10 rating = 1%, 9.812 per %), spell_haste=not significant (0.790 ± 0.455), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.394 ± 0.001, fire_power=0.606 ± 0.001, spirit=not significant (-0.027 ± 0.243)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.98 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.10 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.9 spell_power points (7.22 DPS) | yes | Triune Amulet (7722, -3.08 DPS) [dungeon]; Mindburst Medallion (11196, -3.51 DPS) [quest]; Glowing Eye of Mordresh (10769, -4.05 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.46 DPS, sim-verified) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.22 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Sage's Cloak (19086, -0.60 DPS) [rep]; Mageflame Cloak (13007, -1.26 DPS) [world_drop]; Fel Cape (279269, -1.59 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.7 spell_power points (24.72 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.74 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -9.77 DPS) [pvp]; Robe of the Magi (1716, -13.29 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 23.6 spell_power points (12.19 DPS) | yes | Runecloth Cuffs (254123, -5.49 DPS) [crafted]; Netherflame Cuffs (254065, -7.41 DPS) [crafted]; Arcane Runed Bracers (4744, -7.53 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.12 DPS) [pvp]; Gloves of Spell Mastery (14146, -0.76 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 21.3 spell_power points (11.05 DPS) | yes | Stormpike Cloth Girdle (19094, -2.29 DPS, sim-verified) [rep]; Satyrmane Sash (17755, -2.80 DPS) [dungeon]; Ghostweave Cord (254073, -3.80 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.3 spell_power points (23.98 DPS) | yes | Marshal's Dreadweave Leggings (231587, -2.91 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -8.17 DPS) [pvp]; Ghostweave Pants (14144, -10.51 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.43 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -1.04 DPS) [crafted]; Omnicast Boots (11822, -3.29 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.4 spell_power points (10.02 DPS) | yes | Lorekeeper's Ring (19522, -2.77 DPS) [rep]; Deep Rooted Ring (19109, -2.81 DPS) [world]; Band of the Unicorn (7553, -3.29 DPS) [world_drop] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 16.0 spell_power points (8.28 DPS) | yes | Lorekeeper's Ring (19522, -1.03 DPS) [rep]; Deep Rooted Ring (19109, -1.06 DPS) [world]; Band of the Unicorn (7553, -1.55 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -3.75 DPS, sim-verified) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.70 DPS) [pvp]; Sageclaw (20070, -5.59 DPS) [rep]; Ironbark Staff (20069, -17.35 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (15.01 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.77 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.36 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -4.14 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Eternal Wand (249237, -0.70 DPS) [crafted]; Lesser Eternal Wand (249232, -1.22 DPS) [crafted]; Pyric Caduceus (11748, -1.28 DPS) [dungeon] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Don Julio's Band; trinket2: Thunderbrew's Boot Flask; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 25500000020000000-0000000000000000000-2053345103101351)

Set DPS (verified): 764.8. Weights run: 5.7s. Verify run: 20.1s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.302 ± 0.044, crit=0.603 ± 0.017 per rating point (14 rating = 1%, 8.437 per %), hit=1.236 ± 0.070 per rating point (10 rating = 1%, 12.364 per %), spell_haste=not significant (3.350 ± 0.878), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.365 ± 0.001, fire_power=0.635 ± 0.001, spirit=not significant (-0.264 ± 0.269)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (17.57 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -3.14 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.0 spell_power points (9.78 DPS) | yes | Triune Amulet (7722, -4.57 DPS) [dungeon]; Mindburst Medallion (11196, -4.69 DPS) [quest]; Glowing Eye of Mordresh (10769, -6.31 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.4 spell_power points (19.11 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -2.84 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.25 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (764.8 DPS) | yes | Cloak of Flames (3475, +0.00 DPS) [world_drop]; Stormpike Sage's Cloak (19086, -1.18 DPS) [rep]; Fel Cape (279269, -1.24 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.7 spell_power points (31.71 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.16 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -11.50 DPS) [pvp]; Acumen Robes (17775, -17.67 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.4 spell_power points (15.89 DPS) | yes | Runecloth Cuffs (254123, -6.77 DPS) [crafted]; Netherflame Cuffs (254065, -9.14 DPS) [crafted]; Nethergeld Cuffs (254061, -9.96 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 28.9 spell_power points (18.81 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.57 DPS, sim-verified) [quest]; Knight-Lieutenant's Dreadweave Handwraps (227100, -4.35 DPS) [pvp] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 24.3 spell_power points (15.78 DPS) | yes | Stormpike Cloth Girdle (19094, -3.64 DPS, sim-verified) [rep]; Satyrmane Sash (17755, -4.70 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -5.87 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 51.9 spell_power points (33.77 DPS) | yes | Marshal's Dreadweave Leggings (231587, -5.96 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -12.99 DPS) [pvp]; Ghostweave Pants (14144, -16.85 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.62 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.24 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.30 DPS) [pvp] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 20.8 spell_power points (13.54 DPS) | yes | Deep Rooted Ring (19109, -3.76 DPS) [world]; Lorekeeper's Ring (19522, -4.43 DPS) [rep]; Band of the Unicorn (7553, -5.08 DPS) [world_drop] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 20.1 spell_power points (13.09 DPS) | yes | Lorekeeper's Ring (19522, -3.98 DPS) [rep]; Deep Rooted Ring (19109, -4.04 DPS, sim-verified) [world]; Band of the Unicorn (7553, -4.63 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (764.8 DPS) | yes | Uther's Strength (11302, -3.34 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (764.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (764.8 DPS) | yes | Grand Marshal's Stave (234571, -2.18 DPS) [pvp]; Sageclaw (20070, -7.64 DPS) [rep]; Ironbark Staff (20069, -18.28 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (18.87 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.05 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.59 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -5.21 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (764.8 DPS) | yes | Greater Eternal Wand (249237, -1.38 DPS) [crafted]; Pyric Caduceus (11748, -1.86 DPS) [dungeon]; Lesser Eternal Wand (249232, -2.03 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket2: Thunderbrew's Boot Flask; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 00000000000000000-0000000000000000000-2053100000000000)

Set DPS (verified): 57.5. Weights run: 4.9s. Verify run: 4.2s. 173 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.188 ± 0.006, crit=0.061 ± 0.002 per rating point (14 rating = 1%, 0.851 per %), hit=0.165 ± 0.007 per rating point (10 rating = 1%, 1.650 per %), spell_haste=not significant (0.352 ± 0.094), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.813 ± 0.001, fire_power=0.187 ± 0.000, spirit=0.300 ± 0.054

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.9 spell_power points (2.46 DPS) | yes | Pristine Circlet (253949, -0.83 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -1.71 DPS) [crafted]; Flame Circlet (253953, -1.90 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.9 spell_power points (0.25 DPS) | yes | Scholarly Pendant (277203, -0.08 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.7 spell_power points (1.84 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.33 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.74 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (57.5 DPS) | yes | Battle Healer's Cloak (20427, -0.24 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.61 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.73 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.3 spell_power points (3.40 DPS) | yes | Filigreed Shadow Gown (253907, -0.89 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.51 DPS) [crafted]; Scholarly Robes (2034, -1.58 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.83 DPS) | yes | Owlbeard Bracers (16981, -0.41 DPS, sim-verified) [quest]; Windsong Bangles (263336, -0.55 DPS) [quest]; Mindthrust Bracers (1974, -0.57 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.93 DPS) | yes | Gnoll Casting Gloves (892, -0.28 DPS) [world]; Shadow Gloves (253919, -0.29 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.55 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.4 spell_power points (1.78 DPS) | yes | Pristine Sash (253925, -0.47 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.52 DPS) [crafted]; Novice Ardent's Sash (253887, -0.82 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 11.3 spell_power points (3.11 DPS) | yes | Filigreed Pristine Leggings (253937, -0.83 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.18 DPS) [dungeon]; Colorful Kilt (10048, -1.31 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.8 spell_power points (2.14 DPS) | yes | Shadow Boots (253895, -0.64 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.15 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.38 DPS) | yes | Lavishly Jeweled Ring (1156, -1.07 DPS) [dungeon]; Furlfeather Ring (286745, -1.10 DPS) [world]; Loop of Sacrifice (281673, -1.12 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 1.8 spell_power points (0.50 DPS) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Furlfeather Ring (286745, -0.22 DPS) [world]; Loop of Sacrifice (281673, -0.24 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 31.0 spell_power points (8.54 DPS) | yes | Evocator's Blade (2567, -0.31 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.55 DPS) [quest]; Crescent Staff (6505, -1.01 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-verified (57.5 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.28 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Staff of the Blessed Seer; ranged: Sizzle Stick

No-known-source sample (15 of 173, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword

### Band 30 (orc, 00000000000000000-0000000000000000000-2053235100000000)

Set DPS (verified): 100.1. Weights run: 4.9s. Verify run: 4.2s. 296 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.167 ± 0.008, crit=0.142 ± 0.004 per rating point (14 rating = 1%, 1.985 per %), hit=0.259 ± 0.011 per rating point (10 rating = 1%, 2.585 per %), spell_haste=not significant (0.175 ± 0.126), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.817 ± 0.001, fire_power=0.183 ± 0.000, spirit=not significant (0.196 ± 0.074)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.2 spell_power points (3.92 DPS) | yes | Filigreed Shadow Circlet (253981, -0.43 DPS, sim-verified) [crafted]; Silk Headband (7050, -1.02 DPS) [crafted]; Shadow Circlet (253955, -1.03 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 1.5 spell_power points (0.47 DPS) | yes | Crystal Starfire Medallion (5003, -0.06 DPS) [world_drop]; Pendant of Myzrael (4614, -0.09 DPS) [dungeon]; Glowing Green Talisman (5002, -0.09 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (3.38 DPS) | yes | Invoker's Mantle (215365, -0.86 DPS) [crafted]; Death Speaker Mantle (6685, -1.20 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.29 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (100.1 DPS) | yes | Hillman's Cloak (3719, -0.25 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.32 DPS) [crafted]; Battle Healer's Cloak (19529, -0.32 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (4.18 DPS) | yes | Shadow Gown (253967, -0.14 DPS) [crafted]; Beguiler Robes (7728, -0.14 DPS) [dungeon]; Necrology Robes (2292, -0.23 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.90 DPS) | yes | Tabitha's Cuffs (251486, -1.97 DPS, sim-verified) [quest]; Phoenix Bindings (210781, -2.13 DPS) [crafted]; Nightsky Wristbands (6407, -2.39 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.5 spell_power points (2.74 DPS) | yes | Hands of Darkness (7047, -0.06 DPS) [crafted]; Serpent Gloves (5970, -0.49 DPS) [dungeon]; Jutebraid Gloves (10654, -0.54 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (3.70 DPS) | yes | Warsong Sash (16975, -0.16 DPS) [quest]; Belt of Arugal (6392, -0.33 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -1.13 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 13.8 spell_power points (4.43 DPS) | yes | Shadow Leggings (253993, -0.39 DPS, sim-verified) [crafted]; Filigreed Shadow Leggings (253943, -0.96 DPS) [crafted]; Necromancer Leggings (2277, -1.21 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 12.6 spell_power points (4.05 DPS) | yes | Gilded Slippers (254001, -1.16 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.58 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.25 DPS) | yes | Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Solid Gold Band (279845, -0.65 DPS) [quest]; Electrocutioner Lagnut (9447, -0.85 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (2.20 DPS) | yes | Ring of Power Regulation (273030, -0.27 DPS) [dungeon]; Solid Gold Band (279845, -0.60 DPS) [quest]; Electrocutioner Lagnut (9447, -0.79 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (100.1 DPS) | yes | Hydrocane (9452, -1.54 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.54 DPS) [world]; Soulstaff (249393, -4.51 DPS, sim-verified) [crafted] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.90 DPS) | yes | Orb of Mystic Insight (249394, -0.51 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.07 DPS) [world_drop]; Orb of Souls (249395, -1.23 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-verified (100.1 DPS) | yes | Gravestone Scepter (7001, -0.16 DPS) [quest]; Wand of Eventide (5214, -0.32 DPS) [world_drop]; Greater Mystic Wand (217287, -0.32 DPS, sim-verified) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 296, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (orc, 00000000000000000-0000000000000000000-2053235103101311)

Set DPS (verified): 169.7. Weights run: 5.2s. Verify run: 3.5s. 379 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.204 ± 0.016, crit=0.227 ± 0.006 per rating point (14 rating = 1%, 3.176 per %), hit=0.471 ± 0.023 per rating point (10 rating = 1%, 4.712 per %), spell_haste=not significant (0.599 ± 0.243), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.438 ± 0.000, fire_power=0.562 ± 0.001, spirit=not significant (0.284 ± 0.166)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (7.31 DPS) | yes | Augural Shroud (2620, -2.77 DPS) [world]; Living Cowl (5608, -2.79 DPS) [world]; Papal Fez (9431, -3.58 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 10.0 spell_power points (3.48 DPS) | yes | Glowing Eye of Mordresh (10769, -1.44 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.29 DPS) [dungeon]; Amberglow Talisman (10824, -2.49 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.8 spell_power points (3.77 DPS) | yes | Mantle of Doan (7712, -0.12 DPS) [dungeon]; Green Silken Shoulders (7057, -0.21 DPS) [crafted]; Inquisitor's Shawl (19507, -0.41 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (169.7 DPS) | yes | Battle Healer's Cloak (19528, -0.45 DPS) [rep]; Long Silken Cloak (4326, -0.59 DPS) [crafted]; Guardian Cloak (5965, -0.59 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.6 spell_power points (8.58 DPS) | yes | Dreamweave Vest (10021, -1.04 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.27 DPS) [world_drop]; Robe of Power (7054, -2.06 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.13 DPS) | yes | Condor Bracers (15864, -0.70 DPS) [quest]; Radiant Silver Bracers (4545, -1.17 DPS) [quest]; Phoenix Bindings (210781, -1.72 DPS, sim-verified) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 spell_power points (7.24 DPS) | yes | Red Mageweave Gloves (10018, -2.70 DPS) [crafted]; Black Mageweave Gloves (10003, -3.15 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -3.96 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.8 spell_power points (5.16 DPS) | yes | Star Belt (4329, -0.24 DPS) [crafted]; Gilded Cord (254037, -1.21 DPS) [crafted]; Fiery Cord (254041, -1.26 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.4 spell_power points (5.73 DPS) | yes | Gaze Dreamer Pants (6903, -0.66 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.54 DPS) [dungeon]; Black Mageweave Leggings (9999, -1.56 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.0 spell_power points (9.05 DPS) | yes | Gilded Slippers (254001, -5.72 DPS) [crafted]; Glinteye Slippers (273024, -6.02 DPS) [dungeon]; Fiery Slippers (254005, -6.30 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (3.84 DPS) | yes | Underworld Band (1980, -1.35 DPS) [world_drop]; Snake Hoop (6750, -1.39 DPS) [quest]; Reedknot Ring (9622, -1.40 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.13 DPS) | yes | Snake Hoop (6750, -0.69 DPS) [quest]; Reedknot Ring (9622, -0.70 DPS) [quest]; Underworld Band (1980, -1.75 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (169.7 DPS) | yes | Illusionary Rod (7713, -8.60 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -9.58 DPS) [dungeon]; Swinetusk Shank (6691, -11.44 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (169.7 DPS) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Wand of Eventide (5214, -0.35 DPS) [world_drop]; Greater Mystic Wand (217287, -0.35 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (orc, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 278.1. Weights run: 5.2s. Verify run: 9.2s. 478 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.215 ± 0.018, crit=0.294 ± 0.007 per rating point (14 rating = 1%, 4.115 per %), hit=0.647 ± 0.028 per rating point (10 rating = 1%, 6.465 per %), spell_haste=not significant (1.002 ± 0.287), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.373 ± 0.000, fire_power=0.627 ± 0.001, spirit=not significant (-0.025 ± 0.163)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.86 DPS) | yes | Red Mageweave Headband (10033, -1.43 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -1.84 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.86 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.34 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.57 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.38 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.9 spell_power points (8.03 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.34 DPS) [vendor]; Netherflame Shoulders (254053, -1.77 DPS, sim-verified) [crafted]; Black Mageweave Shoulders (10027, -2.35 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.66 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.18 DPS) [crafted]; Cindercloth Cloak (14044, -2.58 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Acumen Robes (17775, -0.74 DPS, sim-verified) [quest]; Elemental Raiment (9434, -1.09 DPS) [world_drop]; Dreamweave Vest (10021, -1.60 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 9.7 spell_power points (4.60 DPS) | yes | Nethergeld Cuffs (254061, -0.55 DPS) [crafted]; Phoenix Bindings (210781, -0.72 DPS) [crafted]; Spidertank Oilrag (9448, -1.26 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.9 spell_power points (8.99 DPS) | yes | First Sergeant's Dreadweave Gloves (220908, -1.87 DPS) [vendor]; Fiery Gloves (254099, -2.29 DPS) [crafted]; Black Mageweave Gloves (10003, -3.15 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (7.69 DPS) | yes | Defiler's Cloth Girdle (20165, -0.93 DPS) [rep]; Ghostweave Cord (254073, -1.02 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 spell_power points (9.05 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.15 DPS) [vendor]; Red Mageweave Pants (10009, -1.16 DPS) [crafted]; Senior Designer's Pantaloons (11841, -2.62 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.44 DPS) | yes | Fiery Sandals (254111, -4.54 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -4.87 DPS, sim-verified) [vendor]; Cindercloth Boots (10044, -5.16 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.19 DPS) | yes | Philanthropist's Ring (281635, -0.92 DPS) [quest]; Snake Hoop (6750, -2.82 DPS) [quest]; Reedknot Ring (9622, -2.86 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (5.72 DPS) | yes | Philanthropist's Ring (281635, -0.44 DPS) [quest]; Snake Hoop (6750, -2.35 DPS) [quest]; Reedknot Ring (9622, -2.38 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.62 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.70 DPS) [quest]; Ankh of Life (1713, -1.35 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -2.36 DPS, sim-verified) [world]; Kindling Stave (11750, -3.48 DPS) [dungeon]; Inventor's Focal Sword (17719, -4.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.07 DPS) [crafted]; Dreambough Wand (249234, -0.55 DPS) [crafted]; Twisted Nether Wand (249144, -1.02 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Satyrmane Sash; legs: Wizardweave Leggings; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Pyric Caduceus

No-known-source sample (15 of 478, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (orc, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 458.0. Weights run: 5.3s. Verify run: 22.4s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.194 ± 0.033, crit=0.442 ± 0.012 per rating point (14 rating = 1%, 6.185 per %), hit=0.981 ± 0.045 per rating point (10 rating = 1%, 9.812 per %), spell_haste=not significant (0.790 ± 0.455), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.394 ± 0.001, fire_power=0.606 ± 0.001, spirit=not significant (-0.027 ± 0.243)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (13.98 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -2.10 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.9 spell_power points (7.22 DPS) | yes | Glowing Eye of Mordresh (10769, -2.31 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.08 DPS) [dungeon]; Mindburst Medallion (11196, -3.51 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.87 DPS, sim-verified) [crafted]; Champion's Dreadweave Spaulders (227092, -2.22 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Advisor's Cloak (19085, -0.60 DPS) [rep]; Deep Woodlands Cloak (19121, -0.73 DPS) [quest]; Mageflame Cloak (13007, -1.26 DPS) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.7 spell_power points (24.72 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.74 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -9.77 DPS) [pvp]; Robe of the Magi (1716, -11.51 DPS, sim-verified) [world_drop] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 23.6 spell_power points (12.19 DPS) | yes | Runecloth Cuffs (254123, -5.49 DPS) [crafted]; Netherflame Cuffs (254065, -7.41 DPS) [crafted]; Spidertank Oilrag (9448, -7.53 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 23.3 spell_power points (12.07 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -0.68 DPS) [quest]; Blood Guard's Dreadweave Handwraps (227099, -0.79 DPS) [pvp] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 21.3 spell_power points (11.05 DPS) | yes | Frostwolf Cloth Belt (19090, -0.73 DPS) [rep]; Satyrmane Sash (17755, -2.80 DPS) [dungeon]; Ghostweave Cord (254073, -3.80 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 31.7 spell_power points (16.40 DPS) | yes | Ghostweave Pants (14144, +0.00 DPS) [crafted]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -0.60 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.43 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.87 DPS) [dungeon]; Venomspew Footpads (275606, -1.04 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.4 spell_power points (10.02 DPS) | yes | Deep Rooted Ring (19109, -2.81 DPS) [world]; Band of the Unicorn (7553, -3.29 DPS) [world_drop]; Don Julio's Band (19325, -9.72 DPS, sim-verified) [rep] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Deep Rooted Ring (19109, -0.03 DPS) [world]; Band of the Unicorn (7553, -0.52 DPS) [world_drop]; Don Julio's Band (19325, -1.06 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.47 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.45 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.70 DPS) [pvp]; Mindfang (20214, -5.59 DPS) [rep]; Ironbark Staff (20220, -16.26 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (15.01 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.77 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.39 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -4.14 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Eternal Wand (249237, -0.70 DPS) [crafted]; Lesser Eternal Wand (249232, -1.22 DPS) [crafted]; Pyric Caduceus (11748, -1.28 DPS) [dungeon] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (orc, 25500000020000000-0000000000000000000-2053345103101351)

Set DPS (verified): 761.1. Weights run: 5.7s. Verify run: 20.0s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.302 ± 0.044, crit=0.603 ± 0.017 per rating point (14 rating = 1%, 8.437 per %), hit=1.236 ± 0.070 per rating point (10 rating = 1%, 12.364 per %), spell_haste=not significant (3.350 ± 0.878), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.365 ± 0.001, fire_power=0.635 ± 0.001, spirit=not significant (-0.264 ± 0.269)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (17.57 DPS) | yes | Red Mageweave Headband (10033, +0.00 DPS) [crafted]; Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.0 spell_power points (9.78 DPS) | yes | Glowing Eye of Mordresh (10769, -4.42 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -4.57 DPS) [dungeon]; Mindburst Medallion (11196, -4.69 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.4 spell_power points (19.11 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Argent Shoulders (19059, -2.84 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -3.25 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (761.1 DPS) | yes | Deep Woodlands Cloak (19121, -0.71 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -1.18 DPS) [rep]; Fel Cape (279269, -1.24 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.7 spell_power points (31.71 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.16 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -11.50 DPS) [pvp]; Acumen Robes (17775, -15.17 DPS, sim-verified) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.4 spell_power points (15.89 DPS) | yes | Runecloth Cuffs (254123, -6.77 DPS) [crafted]; Netherflame Cuffs (254065, -9.14 DPS) [crafted]; Nethergeld Cuffs (254061, -9.96 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 28.9 spell_power points (18.81 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -3.90 DPS, sim-verified) [quest]; Blood Guard's Dreadweave Handwraps (227099, -4.35 DPS) [pvp] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 24.3 spell_power points (15.78 DPS) | yes | Frostwolf Cloth Belt (19090, -2.10 DPS) [rep]; Satyrmane Sash (17755, -4.70 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -5.87 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 33.7 spell_power points (21.96 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -1.18 DPS) [pvp]; Ghostweave Pants (14144, -6.05 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (15.62 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.24 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -0.60 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 20.8 spell_power points (13.54 DPS) | yes | Deep Rooted Ring (19109, -3.76 DPS) [world]; Advisor's Ring (19518, -4.43 DPS) [rep]; Band of the Unicorn (7553, -5.08 DPS) [world_drop] |
| finger2 | Maiden's Circle (13001) | World drop [world_drop] | 20.1 spell_power points (13.09 DPS) | yes | Deep Rooted Ring (19109, -3.03 DPS, sim-verified) [world]; Advisor's Ring (19518, -3.98 DPS) [rep]; Band of the Unicorn (7553, -4.63 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (761.1 DPS) | yes | Uther's Strength (11302, -3.34 DPS) [world_drop]; Ankh of Life (1713, -8.29 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (761.1 DPS) | yes | Uther's Strength (11302, -1.73 DPS) [world_drop]; Ankh of Life (1713, -5.20 DPS, sim-verified) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (761.1 DPS) | yes | High Warlord's War Staff (234549, -2.18 DPS) [pvp]; Mindfang (20214, -7.64 DPS) [rep]; Ironbark Staff (20220, -21.67 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (18.87 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.05 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.64 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -5.21 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (761.1 DPS) | yes | Greater Eternal Wand (249237, -1.38 DPS) [crafted]; Pyric Caduceus (11748, -1.86 DPS) [dungeon]; Lesser Eternal Wand (249232, -2.03 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Don Julio's Band; finger2: Maiden's Circle; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

