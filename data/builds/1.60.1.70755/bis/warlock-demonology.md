# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2035100000000000000-0000000000000000)

Set DPS (verified): 54.5. Weights run: 4.8s. Verify run: 4.8s. 179 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.189 ± 0.007, crit=0.060 ± 0.001 per rating point (14 rating = 1%, 0.833 per %), hit=0.168 ± 0.008 per rating point (10 rating = 1%, 1.682 per %), spell_haste=0.936 ± 0.092, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.810 ± 0.001, fire_power=0.190 ± 0.000, spirit=not significant (0.179 ± 0.073)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.9 spell_power points (2.24 DPS) | yes | Pristine Circlet (253949, -0.73 DPS, sim-verified) [crafted]; Flame Circlet (253953, -1.71 DPS) [crafted]; Shadow Goggles (4373, -1.73 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.4 spell_power points (0.09 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 6.7 spell_power points (1.68 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.19 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.68 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Caretaker's Cape (20428, -0.29 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.43 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.54 DPS) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 12.3 spell_power points (3.08 DPS) | yes | Filigreed Shadow Gown (253907, -0.86 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.46 DPS) [crafted]; Scholarly Robes (2034, -1.79 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.05 DPS) [world_drop]; Repurposed Hair Band (281256, -0.05 DPS) [quest]; Windsong Bangles (263336, -0.14 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.76 DPS) | yes | Shadow Gloves (253919, -0.21 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.25 DPS) [world]; Pristine Gloves (253913, -0.61 DPS) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 6.4 spell_power points (1.61 DPS) | yes | Pristine Sash (253925, -0.42 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.47 DPS) [crafted]; Novice Ardent's Sash (253887, -0.83 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 10.8 spell_power points (2.70 DPS) | yes | Filigreed Pristine Leggings (253937, -0.73 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.94 DPS) [dungeon]; Colorful Kilt (10048, -1.22 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.8 spell_power points (1.95 DPS) | yes | Shadow Boots (253895, -0.60 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.72 DPS) [crafted]; Pristine Boots (253889, -1.05 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.4 spell_power points (1.35 DPS) | yes | Lavishly Jeweled Ring (1156, -1.07 DPS) [dungeon]; Band of Purification (12996, -1.08 DPS) [world_drop]; Furlfeather Ring (286745, -1.10 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (1.26 DPS) | yes | Lavishly Jeweled Ring (1156, -0.80 DPS, sim-verified) [dungeon]; Band of Purification (12996, -0.99 DPS) [world_drop]; Furlfeather Ring (286745, -1.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Evocator's Blade (2567) | Blackfathom Deeps [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Staff of the Blessed Seer (2271, -0.65 DPS, sim-verified) [dungeon]; Emberstone Staff (5201, -1.32 DPS) [dungeon]; Witching Stave (1484, -1.62 DPS) [dungeon] |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 3.2 spell_power points (0.80 DPS) | yes | Antipodean Rod (2879, -0.51 DPS) [world_drop]; Mathystra Channeling Orb (281291, -0.56 DPS) [quest]; Onu's Patience (281292, -0.57 DPS) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Spark of the People's Militia (12296, -0.25 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Evocator's Blade; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 00000000000000000-2035113101210001000-0000000000000000)

Set DPS (verified): 91.4. Weights run: 4.9s. Verify run: 4.9s. 303 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.100 ± 0.006, crit=0.042 ± 0.001 per rating point (14 rating = 1%, 0.594 per %), hit=0.112 ± 0.005 per rating point (10 rating = 1%, 1.116 per %), spell_haste=-0.538 ± 0.083, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.798 ± 0.001, fire_power=0.202 ± 0.001, spirit=not significant (0.128 ± 0.079)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.8 spell_power points (3.19 DPS) | yes | Filigreed Shadow Circlet (253981, -0.33 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.75 DPS) [crafted]; Shadow Circlet (253955, -0.81 DPS) [crafted] |
| neck | Pendant of Myzrael (4614) (or Glowing Green Talisman (5002)) | Razorfen Downs: Splinterbone Captain [dungeon] | 0.8 spell_power points (0.21 DPS) | yes | Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.07 DPS) [world_drop]; Master Sergeant's Insignia (18442, -0.07 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 spell_power points (2.44 DPS) | yes | Invoker's Mantle (215365, -0.54 DPS) [crafted]; Death Speaker Mantle (6685, -0.81 DPS) [dungeon]; Moonlit Amice (11884, -0.96 DPS, sim-verified) [quest] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Caretaker's Cape (19533, -0.13 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.52 DPS) | yes | Necrology Robes (2292, -0.41 DPS, sim-verified) [dungeon]; Shadow Gown (253967, -0.57 DPS) [crafted]; Beguiler Robes (7728, -0.81 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.44 DPS) | yes | Phoenix Bindings (210781, -2.00 DPS, sim-verified) [crafted]; Windsong Bangles (263336, -2.17 DPS) [quest]; Resilient Bands (14402, -2.30 DPS) [world_drop] |
| hands | Hands of Darkness (7047) | Tailoring [crafted] | 7.8 spell_power points (2.12 DPS) | yes | Ephemeral Grips (273034, -0.17 DPS) [world]; Serpent Gloves (5970, -0.22 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.49 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.0 spell_power points (2.98 DPS) | yes | Ghamoo-ra's Bind (6908, -0.81 DPS) [dungeon]; Belt of Arugal (6392, -0.85 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -1.08 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 13.2 spell_power points (3.56 DPS) | yes | Shadow Leggings (253993, -0.47 DPS, sim-verified) [crafted]; Filigreed Shadow Leggings (253943, -1.05 DPS) [crafted]; Necromancer Leggings (2277, -1.40 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.9 spell_power points (2.95 DPS) | yes | Gilded Slippers (254001, -0.92 DPS) [crafted]; Spidersilk Boots (4320, -1.05 DPS) [crafted]; Glinteye Slippers (273024, -1.51 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.90 DPS) | yes | Minor Channeling Ring (1449, -0.54 DPS) [quest]; Electrocutioner Lagnut (9447, -0.84 DPS) [dungeon]; Ring of Power Regulation (273030, -2.18 DPS, sim-verified) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Minor Channeling Ring (1449, -0.27 DPS) [quest]; Ring of Power Regulation (273030, -0.30 DPS, sim-verified) [dungeon]; Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.22 DPS) [dungeon]; Spine of the Devourer (273040, -1.22 DPS) [world]; Tendonscraper Dagger (273053, -1.73 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.44 DPS) | yes | Orb of Mystic Insight (249394, -0.49 DPS, sim-verified) [crafted]; Orb of Mistmantle (13031, -1.00 DPS) [world_drop]; Orb of Souls (249395, -1.15 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, -0.22 DPS) [quest]; Wand of Eventide (5214, -0.27 DPS) [world_drop]; Greater Mystic Wand (217287, -0.27 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Pendant of Myzrael; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Hands of Darkness; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 303, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 10000000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 162.7. Weights run: 4.9s. Verify run: 5.0s. 386 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.105 ± 0.006, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.902 per %), hit=0.171 ± 0.007 per rating point (10 rating = 1%, 1.706 per %), spell_haste=0.872 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.822 ± 0.001, fire_power=0.178 ± 0.000, spirit=not significant (0.167 ± 0.061)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.57 DPS) | yes | Living Cowl (5608, -2.65 DPS, sim-verified) [world]; Augural Shroud (2620, -2.80 DPS) [world]; Holy Shroud (2721, -2.82 DPS) [world_drop] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.2 spell_power points (2.87 DPS) | yes | Glowing Eye of Mordresh (10769, -0.42 DPS) [dungeon]; Necklace of Calisea (1714, -2.27 DPS) [dungeon]; Amberglow Talisman (10824, -2.35 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (3.11 DPS) | yes | Green Silken Shoulders (7057, -0.25 DPS) [crafted]; Mantle of Doan (7712, -0.31 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.50 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (162.7 DPS) | yes | Caretaker's Cape (19532, -0.37 DPS) [rep]; Long Silken Cloak (4326, -0.46 DPS) [crafted]; Guardian Cloak (5965, -0.46 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.5 spell_power points (7.34 DPS) | yes | Elemental Raiment (9434, -0.77 DPS) [world_drop]; Dreamweave Vest (10021, -0.85 DPS, sim-verified) [crafted]; Shadoweave Robe (10004, -1.57 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.82 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Condor Bracers (15864, -0.63 DPS) [quest]; Earthen Silk Cuffs (254019, -1.30 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.6 spell_power points (6.13 DPS) | yes | Shadoweave Gloves (10023, -1.59 DPS) [crafted]; Black Mageweave Gloves (10003, -2.21 DPS, sim-verified) [crafted]; Black Handwraps (254027, -2.30 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.4 spell_power points (4.51 DPS) | yes | Star Belt (4329, -0.23 DPS) [crafted]; Black Cord (254043, -0.33 DPS) [crafted]; Belt of Arugal (6392, -1.34 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.3 spell_power points (5.40 DPS) | yes | Red Mageweave Pants (10009, -0.63 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.18 DPS) [dungeon]; Shadow Leggings (253993, -1.31 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (7.88 DPS) | yes | Black Slippers (254007, -3.99 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -5.25 DPS) [crafted]; Glinteye Slippers (273024, -5.37 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 spell_power points (3.76 DPS) | yes | Lorekeeper's Ring (19524, -0.95 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.26 DPS) [quest]; Reedknot Ring (9622, -1.57 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 spell_power points (3.29 DPS) | yes | Lorekeeper's Ring (19524, -0.23 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.79 DPS) [quest]; Reedknot Ring (9622, -1.10 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (162.7 DPS) | yes | Hypnotic Blade (7714, -6.57 DPS) [dungeon]; Illusionary Rod (7713, -7.44 DPS, sim-verified) [dungeon]; Swinetusk Shank (6691, -8.03 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (162.7 DPS) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.18 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 386, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25400000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 229.6. Weights run: 5.0s. Verify run: 8.9s. 485 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.206 ± 0.011, crit=0.098 ± 0.002 per rating point (14 rating = 1%, 1.371 per %), hit=0.281 ± 0.013 per rating point (10 rating = 1%, 2.809 per %), spell_haste=not significant (0.501 ± 0.172), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.768 ± 0.001, fire_power=0.232 ± 0.001, spirit=not significant (0.291 ± 0.137)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.44 DPS) | yes | Dreamweave Circlet (10041, -0.16 DPS) [crafted]; Red Mageweave Headband (10033, -1.35 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.10 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 10.0 spell_power points (3.51 DPS) | yes | Mindburst Medallion (11196, -0.47 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.48 DPS) [dungeon]; Horizon Choker (13085, -2.09 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.7 spell_power points (5.85 DPS) | yes | Netherpearl Shoulders (254055, -0.09 DPS) [crafted]; Shadoweave Shoulders (10028, -1.28 DPS) [crafted]; Nethergeld Shoulders (254049, -1.51 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (229.6 DPS) | yes | Runecloth Cloak (13860, -2.17 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.79 DPS) [dungeon]; Caretaker's Cape (19531, -2.88 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.7 spell_power points (8.64 DPS) | yes | Acumen Robes (17775, -0.55 DPS) [quest]; Dreamweave Vest (10021, -0.91 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.29 DPS) [world_drop] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 12.9 spell_power points (4.51 DPS) | yes | Nethergeld Cuffs (254061, -1.13 DPS, sim-verified) [crafted]; Arcane Runed Bracers (4744, -1.36 DPS) [quest]; Spidertank Oilrag (9448, -1.36 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.9 spell_power points (7.30 DPS) | yes | Black Gloves (254101, -0.66 DPS) [crafted]; Runecloth Gloves (13863, -1.53 DPS) [crafted]; Black Mageweave Gloves (10003, -2.05 DPS) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | 17.2 spell_power points (6.02 DPS) | yes | Black Waistcord (254087, -0.26 DPS) [crafted]; Satyrmane Sash (17755, -0.40 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.83 DPS) [rep] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 22.4 spell_power points (7.85 DPS) | yes | Wizardweave Leggings (14132, -0.70 DPS, sim-verified) [crafted]; Senior Designer's Pantaloons (11841, -1.32 DPS) [dungeon]; Red Mageweave Pants (10009, -2.09 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 26.0 spell_power points (9.11 DPS) | yes | Black Sandals (254113, -2.72 DPS, sim-verified) [crafted]; Shadoweave Boots (10031, -4.33 DPS) [crafted]; Gilded Sandals (254107, -4.61 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.55 DPS) | yes | Underworld Band (1980, -0.43 DPS) [world_drop]; Philanthropist's Ring (281635, -0.69 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.75 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 spell_power points (4.20 DPS) | yes | Philanthropist's Ring (281635, -0.34 DPS) [quest]; Underworld Band (1980, -0.61 DPS, sim-verified) [world_drop]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (229.6 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.63 DPS, sim-verified) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (229.6 DPS) | yes | Kindling Stave (11750, -1.36 DPS, sim-verified) [dungeon]; Spire of Hakkar (10844, -1.61 DPS) [world]; Inventor's Focal Sword (17719, -3.45 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-verified (229.6 DPS) | yes | Dreambough Wand (249234, -0.35 DPS) [crafted]; Burning Sliver (5249, -0.70 DPS) [quest]; Twisted Nether Wand (249144, -0.70 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Ghostweave Cord; legs: Felcloth Pants; finger1: Band of the Unicorn; finger2: Lorekeeper's Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; ranged: Lesser Eternal Wand

No-known-source sample (15 of 485, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-2035113101220001350-0050000000000000)

Set DPS (verified): 391.0. Weights run: 9.2s. Verify run: 23.0s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.032 ± 0.013, crit=0.151 ± 0.006 per rating point (14 rating = 1%, 2.114 per %), hit=0.410 ± 0.020 per rating point (10 rating = 1%, 4.102 per %), spell_haste=1.261 ± 0.268, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.767 ± 0.001, fire_power=0.233 ± 0.001, spirit=not significant (-0.080 ± 0.240)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.86 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -1.56 DPS) [pvp]; Felcloth Hood (14111, -1.61 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.0 spell_power points (4.83 DPS) | yes | Glowing Eye of Mordresh (10769, -2.01 DPS) [dungeon]; Mindburst Medallion (11196, -2.41 DPS) [quest]; Triune Amulet (7722, -2.65 DPS, sim-verified) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.06 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.37 DPS) [crafted]; Felcloth Shoulders (14112, -2.90 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Fel Cape (279269, -0.16 DPS) [crafted]; Stormpike Sage's Cloak (19086, -1.06 DPS, sim-verified) [rep]; Nightfall Drape (12465, -2.01 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.0 spell_power points (18.51 DPS) | yes | Robe of Winter Night (14136, -5.42 DPS, sim-verified) [crafted]; Field Marshal's Dreadweave Robe (231582, -5.63 DPS) [pvp]; Felcloth Robe (14106, -7.41 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 22.0 spell_power points (8.85 DPS) | yes | Runecloth Cuffs (254123, -4.43 DPS) [crafted]; Netherpearl Cuffs (254067, -4.84 DPS) [crafted]; Arcane Runed Bracers (4744, -5.23 DPS) [quest] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.85 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.40 DPS) [pvp]; Dreamweave Gloves (10019, -1.61 DPS) [crafted] |
| waist | Stormpike Cloth Girdle (19094) | Stormpike Guard [rep] | 18.0 spell_power points (7.24 DPS) | yes | Highlander's Cloth Girdle (20047, -1.36 DPS, sim-verified) [rep]; Black Waistcord (254087, -1.38 DPS) [crafted]; Ghostweave Cord (254073, -1.61 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 36.2 spell_power points (14.58 DPS) | yes | Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -3.31 DPS) [pvp]; Ghostweave Pants (14144, -4.12 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.66 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.66 DPS, sim-verified) [crafted]; Omnicast Boots (11822, -1.61 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.0 spell_power points (7.24 DPS) | yes | Band of the Unicorn (7553, -2.01 DPS) [world_drop]; Deep Rooted Ring (19109, -2.41 DPS) [world]; Underworld Band (1980, -2.92 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19522) | Silverwing Sentinels [rep] | 14.0 spell_power points (5.63 DPS) | yes | Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Deep Rooted Ring (19109, -0.80 DPS) [world]; Underworld Band (1980, -1.32 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.98 DPS, sim-verified) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.75 DPS, sim-verified) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sageclaw (20070, -3.21 DPS) [rep]; Grand Marshal's Stave (234571, -5.18 DPS) [pvp]; Ironbark Staff (20069, -14.95 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.67 DPS) | yes | Grand Marshal's Tome of Power (234589, -1.61 DPS) [pvp]; Tome of Shadow Force (19309, -2.15 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.22 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Eternal Wand (249237, +0.00 DPS) [crafted]; Lesser Eternal Wand (249232, -0.40 DPS) [crafted]; Dreambough Wand (249234, -0.80 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Stormpike Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 05000000000000000-2035113011220001351-0550005000000000)

Set DPS (verified): 706.0. Weights run: 5.5s. Verify run: 20.3s. 1037 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.380 ± 0.045, crit=0.633 ± 0.021 per rating point (14 rating = 1%, 8.865 per %), hit=1.092 ± 0.078 per rating point (10 rating = 1%, 10.920 per %), spell_haste=6.946 ± 1.021, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.808 ± 0.002, fire_power=0.192 ± 0.000, spirit=not significant (0.626 ± 0.244)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 32.3 spell_power points (19.54 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Ghostweave Hood (254135, -0.03 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.8 spell_power points (9.55 DPS) | yes | Mindburst Medallion (11196, -2.65 DPS) [quest]; Glowing Eye of Mordresh (10769, -3.48 DPS) [dungeon]; Triune Amulet (7722, -6.34 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.8 spell_power points (18.63 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.86 DPS) [crafted]; Shroud of the Nathrezim (18720, -2.33 DPS) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-verified (706.0 DPS) | yes | Spritecaster Cape (11623, -4.13 DPS, sim-verified) [dungeon]; Caretaker's Cape (19530, -4.22 DPS) [rep]; Stormpike Sage's Cloak (19086, -4.23 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.4 spell_power points (29.88 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.02 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.17 DPS) [pvp]; Robe of Winter Night (14136, -11.06 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.4 spell_power points (17.79 DPS) | yes | Netherpearl Cuffs (254067, -7.93 DPS) [crafted]; Runecloth Cuffs (254123, -8.84 DPS) [crafted]; Earthenweave Cuffs (254131, -9.17 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.5 spell_power points (18.46 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Dreamweave Gloves (10019, -4.01 DPS) [crafted]; Black Gloves (254101, -8.42 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | 25.1 spell_power points (15.20 DPS) | yes | Stormpike Cloth Girdle (19094, -2.02 DPS) [rep]; Ghostweave Cord (254073, -2.58 DPS) [crafted]; Wisdom of the Timbermaw (19047, -4.22 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 53.5 spell_power points (32.37 DPS) | yes | Marshal's Dreadweave Leggings (231587, -5.63 DPS) [pvp]; Ghostweave Pants (14144, -9.08 DPS) [crafted]; Magister's Leggings (16687, -10.79 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.4 spell_power points (17.16 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.83 DPS) [crafted]; Knight-Lieutenant's Dreadweave Boots (17562, -2.22 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.4 spell_power points (14.76 DPS) | yes | Deep Rooted Ring (19109, -5.21 DPS) [world]; Lorekeeper's Ring (19522, -6.30 DPS) [rep]; Underworld Band (1980, -6.77 DPS) [world_drop] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 19.8 spell_power points (11.96 DPS) | yes | Deep Rooted Ring (19109, -2.41 DPS) [world]; Lorekeeper's Ring (19522, -3.50 DPS) [rep]; Underworld Band (1980, -3.97 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (706.0 DPS) | yes | Ankh of Life (1713, -1.40 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.16 DPS) [quest]; Thunderbrew's Boot Flask (744, -2.91 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (706.0 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Mark of the Chosen (17774, -3.67 DPS, sim-verified) [quest] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (706.0 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Sageclaw (20070, -6.25 DPS) [rep]; Ironbark Staff (20069, -13.14 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (17.53 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Tome of Shadow Force (19309, -2.38 DPS) [rep]; Spirit of Aquementas (11904, -4.84 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (706.0 DPS) | yes | Torch of Light (279246, +0.00 DPS) [crafted]; Greater Eternal Wand (249237, -1.61 DPS) [crafted]; Lesser Eternal Wand (249232, -2.21 DPS) [crafted] |

**New at 60:** head: Dreamweave Circlet; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Don Julio's Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1037, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-2035100000000000000-0000000000000000)

Set DPS (verified): 52.9. Weights run: 4.8s. Verify run: 4.9s. 173 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=-0.404 ± 0.007, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.452 per %), hit=0.102 ± 0.007 per rating point (10 rating = 1%, 1.021 per %), spell_haste=not significant (0.067 ± 0.111), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.770 ± 0.002, fire_power=0.230 ± 0.001, spirit=not significant (0.329 ± 0.118)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.5 spell_power points (1.86 DPS) | yes | Pristine Circlet (253949, -0.74 DPS, sim-verified) [crafted]; Flame Circlet (253953, -1.30 DPS) [crafted]; Shadow Goggles (4373, -1.42 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 1.0 spell_power points (0.22 DPS) | yes | Scholarly Pendant (277203, -0.07 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 5.0 spell_power points (1.09 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.22 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.31 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (52.9 DPS) | yes | Battle Healer's Cloak (20427, -0.29 DPS, sim-verified) [rep]; Black Whelp Cloak (7283, -0.51 DPS) [crafted]; Mystic Shawl (3449, -0.73 DPS) [quest] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 10.8 spell_power points (2.36 DPS) | yes | Filigreed Shadow Gown (253907, -0.77 DPS, sim-verified) [crafted]; Scholarly Robes (2034, -0.84 DPS) [dungeon]; Robe of the Moccasin (6465, -0.98 DPS) [dungeon] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.66 DPS) | yes | Owlbeard Bracers (16981, -0.44 DPS) [quest]; Crystalline Cuffs (14148, -0.44 DPS) [dungeon]; Windsong Bangles (263336, -0.50 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (1.53 DPS) | yes | Gnoll Casting Gloves (892, -0.25 DPS, sim-verified) [world]; Shadow Gloves (253919, -0.35 DPS) [crafted]; Blight Gloves (279877, -0.44 DPS) [quest] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 5.4 spell_power points (1.18 DPS) | yes | Novice Arcanist's Sash (253885, -0.30 DPS) [crafted]; Pristine Sash (253925, -0.43 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.53 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 9.8 spell_power points (2.14 DPS) | yes | Silk-threaded Trousers (1929, -0.61 DPS) [dungeon]; Colorful Kilt (10048, -0.69 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.74 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 spell_power points (1.53 DPS) | yes | Shadow Boots (253895, -0.52 DPS) [crafted]; Red Woolen Boots (4313, -0.68 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.88 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (1.09 DPS) | yes | Furlfeather Ring (286745, -0.88 DPS) [world]; Demon Band (12054, -0.95 DPS) [world_drop]; Horn Ring (17692, -0.95 DPS) [quest] |
| finger2 | Band of Purification (12996) | World drop [world_drop] | 2.0 spell_power points (0.43 DPS) | yes | Furlfeather Ring (286745, -0.12 DPS, sim-verified) [world]; Demon Band (12054, -0.29 DPS) [world_drop]; Horn Ring (17692, -0.29 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 31.3 spell_power points (6.85 DPS) | yes | Evocator's Blade (2567, -0.36 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.44 DPS) [quest]; Crescent Staff (6505, -1.16 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | sim-verified (52.9 DPS) | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted]; Cinder Wand (5242, -0.22 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Band of Purification; main_hand: Staff of the Blessed Seer; ranged: Sizzle Stick

No-known-source sample (15 of 173, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 248006 Militia Sword

### Band 30 (troll, 00000000000000000-2035113101210001000-0000000000000000)

Set DPS (verified): 91.3. Weights run: 4.9s. Verify run: 5.6s. 296 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.157 ± 0.008, crit=0.090 ± 0.002 per rating point (14 rating = 1%, 1.255 per %), hit=0.219 ± 0.010 per rating point (10 rating = 1%, 2.193 per %), spell_haste=1.587 ± 0.121, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.824 ± 0.001, fire_power=0.176 ± 0.000, spirit=not significant (0.204 ± 0.074)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.2 spell_power points (3.53 DPS) | yes | Filigreed Shadow Circlet (253981, -0.14 DPS, sim-verified) [crafted]; Shadow Circlet (253955, -0.91 DPS) [crafted]; Silk Headband (7050, -0.93 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 1.4 spell_power points (0.42 DPS) | yes | Crystal Starfire Medallion (5003, -0.06 DPS) [world_drop]; Pendant of Myzrael (4614, -0.06 DPS) [dungeon]; Glowing Green Talisman (5002, -0.06 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 spell_power points (3.01 DPS) | yes | Invoker's Mantle (215365, -0.76 DPS) [crafted]; Death Speaker Mantle (6685, -0.78 DPS) [dungeon]; Magician's Mantle (12998, -1.15 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (91.3 DPS) | yes | Hillman's Cloak (3719, -0.23 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.29 DPS) [crafted]; Battle Healer's Cloak (19529, -0.29 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.75 DPS) | yes | Beguiler Robes (7728, -0.14 DPS) [dungeon]; Shadow Gown (253967, -0.17 DPS, sim-verified) [crafted]; Necrology Robes (2292, -0.20 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.60 DPS) | yes | Tabitha's Cuffs (251486, -1.93 DPS, sim-verified) [quest]; Phoenix Bindings (210781, -1.94 DPS) [crafted]; Nightsky Wristbands (6407, -2.15 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.5 spell_power points (2.46 DPS) | yes | Hands of Darkness (7047, -0.02 DPS) [crafted]; Serpent Gloves (5970, -0.44 DPS) [dungeon]; Jutebraid Gloves (10654, -0.50 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (3.31 DPS) | yes | Warsong Sash (16975, -0.17 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.28 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -1.00 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 13.8 spell_power points (3.99 DPS) | yes | Shadow Leggings (253993, -0.05 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.87 DPS) [crafted]; Necromancer Leggings (2277, -1.12 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 12.6 spell_power points (3.64 DPS) | yes | Gilded Slippers (254001, -1.07 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.27 DPS) [dungeon]; Spidersilk Boots (4320, -1.44 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.02 DPS) | yes | Ring of Power Regulation (273030, -0.29 DPS) [dungeon]; Solid Gold Band (279845, -0.57 DPS) [quest]; Electrocutioner Lagnut (9447, -0.74 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (1.96 DPS) | yes | Ring of Power Regulation (273030, -0.23 DPS) [dungeon]; Solid Gold Band (279845, -0.51 DPS) [quest]; Electrocutioner Lagnut (9447, -0.68 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (91.3 DPS) | yes | Hydrocane (9452, -1.39 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.39 DPS) [world]; Soulstaff (249393, -4.27 DPS, sim-verified) [crafted] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.60 DPS) | yes | Orb of Mystic Insight (249394, -0.31 DPS) [crafted]; Orb of Mistmantle (13031, -0.94 DPS) [world_drop]; Orb of Souls (249395, -1.09 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-verified (91.3 DPS) | yes | Gravestone Scepter (7001, -0.13 DPS) [quest]; Wand of Eventide (5214, -0.29 DPS) [world_drop]; Greater Mystic Wand (217287, -0.29 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 296, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 10000000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 163.0. Weights run: 4.9s. Verify run: 3.8s. 379 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.105 ± 0.006, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.902 per %), hit=0.171 ± 0.007 per rating point (10 rating = 1%, 1.706 per %), spell_haste=0.872 ± 0.075, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.822 ± 0.001, fire_power=0.178 ± 0.000, spirit=not significant (0.167 ± 0.061)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.57 DPS) | yes | Living Cowl (5608, -2.66 DPS, sim-verified) [world]; Augural Shroud (2620, -2.80 DPS) [world]; Holy Shroud (2721, -2.82 DPS) [world_drop] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 9.2 spell_power points (2.87 DPS) | yes | Glowing Eye of Mordresh (10769, -0.42 DPS) [dungeon]; Necklace of Calisea (1714, -2.27 DPS) [dungeon]; Amberglow Talisman (10824, -2.35 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.9 spell_power points (3.11 DPS) | yes | Green Silken Shoulders (7057, -0.25 DPS) [crafted]; Mantle of Doan (7712, -0.31 DPS) [dungeon]; Inquisitor's Shawl (19507, -0.50 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-verified (163.0 DPS) | yes | Long Silken Cloak (4326, -0.46 DPS) [crafted]; Guardian Cloak (5965, -0.46 DPS) [crafted]; Battle Healer's Cloak (19528, -0.51 DPS, sim-verified) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.5 spell_power points (7.34 DPS) | yes | Dreamweave Vest (10021, -0.60 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.77 DPS) [world_drop]; Shadoweave Robe (10004, -1.57 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.82 DPS) | yes | Radiant Silver Bracers (4545, -1.30 DPS) [quest]; Earthen Silk Cuffs (254019, -1.30 DPS) [crafted]; Condor Bracers (15864, -1.34 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.6 spell_power points (6.13 DPS) | yes | Shadoweave Gloves (10023, -1.59 DPS) [crafted]; Black Mageweave Gloves (10003, -1.73 DPS, sim-verified) [crafted]; Black Handwraps (254027, -2.30 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.4 spell_power points (4.51 DPS) | yes | Star Belt (4329, -0.23 DPS) [crafted]; Black Cord (254043, -0.33 DPS) [crafted]; Warsong Sash (16975, -1.07 DPS) [quest] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.3 spell_power points (5.40 DPS) | yes | Red Mageweave Pants (10009, -0.63 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.18 DPS) [dungeon]; Shadow Leggings (253993, -1.31 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 25.2 spell_power points (7.88 DPS) | yes | Black Slippers (254007, -4.40 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -5.25 DPS) [crafted]; Glinteye Slippers (273024, -5.37 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 spell_power points (3.76 DPS) | yes | Advisor's Ring (19520, -0.95 DPS) [rep]; Reedknot Ring (9622, -1.57 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 spell_power points (3.29 DPS) | yes | Advisor's Ring (19520, -0.48 DPS) [rep]; Reedknot Ring (9622, -1.10 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (163.0 DPS) | yes | Hypnotic Blade (7714, -6.57 DPS) [dungeon]; Illusionary Rod (7713, -7.37 DPS, sim-verified) [dungeon]; Swinetusk Shank (6691, -8.03 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (163.0 DPS) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Gravestone Scepter (7001, -0.18 DPS) [quest]; Greater Mystic Wand (217287, -0.31 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 379, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25400000000000000-2035113101220001350-0000000000000000)

Set DPS (verified): 229.2. Weights run: 5.0s. Verify run: 9.0s. 478 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.157 ± 0.016, crit=0.184 ± 0.004 per rating point (14 rating = 1%, 2.577 per %), hit=0.483 ± 0.021 per rating point (10 rating = 1%, 4.829 per %), spell_haste=1.824 ± 0.272, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.781 ± 0.001, fire_power=0.219 ± 0.000, spirit=not significant (-0.144 ± 0.132)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (9.99 DPS) | yes | Dreamweave Circlet (10041, -1.38 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -1.80 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.22 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindburst Medallion (11196, -0.39 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.51 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -2.15 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 15.8 spell_power points (5.85 DPS) | yes | Netherpearl Shoulders (254055, -0.48 DPS) [crafted]; Shadoweave Shoulders (10028, -0.94 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.42 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.25 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.73 DPS) [crafted]; Nightfall Drape (12465, -2.20 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.9 spell_power points (8.49 DPS) | yes | Acumen Robes (17775, -0.63 DPS, sim-verified) [quest]; Elemental Raiment (9434, -0.72 DPS) [world_drop]; Dreamweave Vest (10021, -1.31 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 11.2 spell_power points (4.16 DPS) | yes | Spidertank Oilrag (9448, -0.78 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -1.16 DPS) [crafted]; Condor Bracers (15864, -1.57 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.6 spell_power points (6.89 DPS) | yes | Black Gloves (254101, -0.93 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.34 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.56 DPS) [vendor] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 16.2 spell_power points (6.01 DPS) | yes | Satyrmane Sash (17755, -0.25 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -0.60 DPS) [rep]; Ghostweave Cord (254073, -0.83 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 22.2 spell_power points (8.21 DPS) | yes | Wizardweave Leggings (14132, -1.65 DPS, sim-verified) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.12 DPS) [vendor]; Shadoweave Pants (10002, -2.14 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.88 DPS) | yes | Black Sandals (254113, -2.42 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.61 DPS) [vendor]; Gilded Sandals (254107, -4.29 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.81 DPS) | yes | Underworld Band (1980, -0.48 DPS) [world_drop]; Philanthropist's Ring (281635, -0.82 DPS) [quest]; Reedknot Ring (9622, -2.22 DPS) [quest] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 spell_power points (4.44 DPS) | yes | Philanthropist's Ring (281635, -0.45 DPS) [quest]; Underworld Band (1980, -0.63 DPS, sim-verified) [world_drop]; Reedknot Ring (9622, -1.85 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.36 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.91 DPS, sim-verified) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -1.09 DPS, sim-verified) [world]; Kindling Stave (11750, -2.19 DPS) [dungeon]; Inventor's Focal Sword (17719, -2.89 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.37 DPS) [crafted]; Thunderwood (13062, -0.74 DPS) [world_drop]; Twisted Nether Wand (249144, -0.74 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger1: Band of the Unicorn; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Lesser Eternal Wand

No-known-source sample (15 of 478, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-2035113101220001350-0050000000000000)

Set DPS (verified): 386.4. Weights run: 9.2s. Verify run: 22.0s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=-0.032 ± 0.013, crit=0.151 ± 0.006 per rating point (14 rating = 1%, 2.114 per %), hit=0.410 ± 0.020 per rating point (10 rating = 1%, 4.102 per %), spell_haste=1.261 ± 0.268, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.767 ± 0.001, fire_power=0.233 ± 0.001, spirit=not significant (-0.080 ± 0.240)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.86 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Felcloth Hood (14111, -1.51 DPS, sim-verified) [crafted]; Champion's Dreadweave Cowl (227090, -1.56 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.0 spell_power points (4.83 DPS) | yes | Glowing Eye of Mordresh (10769, -2.01 DPS) [dungeon]; Triune Amulet (7722, -2.23 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -2.41 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (10.06 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.37 DPS) [crafted]; Felcloth Shoulders (14112, -2.58 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Fel Cape (279269, -0.16 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.80 DPS) [quest]; Frostwolf Advisor's Cloak (19085, -0.93 DPS, sim-verified) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 46.0 spell_power points (18.51 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.63 DPS) [pvp]; Robe of Winter Night (14136, -5.89 DPS, sim-verified) [crafted]; Felcloth Robe (14106, -7.41 DPS) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 22.0 spell_power points (8.85 DPS) | yes | Runecloth Cuffs (254123, -4.43 DPS) [crafted]; Netherpearl Cuffs (254067, -4.84 DPS) [crafted]; Spidertank Oilrag (9448, -5.23 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.85 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS) [crafted]; General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Gloves (17577, -0.40 DPS) [pvp] |
| waist | Frostwolf Cloth Belt (19090) | Frostwolf Clan [rep] | 18.0 spell_power points (7.24 DPS) | yes | Defiler's Cloth Girdle (20163, -0.76 DPS) [rep]; Black Waistcord (254087, -1.38 DPS) [crafted]; Ghostweave Cord (254073, -1.61 DPS) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 28.0 spell_power points (11.27 DPS) | yes | Ghostweave Pants (14144, +0.00 DPS) [crafted]; Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (9.66 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.80 DPS) [crafted]; Omnicast Boots (11822, -1.61 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.0 spell_power points (7.24 DPS) | yes | Band of the Unicorn (7553, -2.01 DPS) [world_drop]; Deep Rooted Ring (19109, -2.41 DPS) [world]; Underworld Band (1980, -2.92 DPS) [world_drop] |
| finger2 | Advisor's Ring (19518) | Warsong Outriders [rep] | 14.0 spell_power points (5.63 DPS) | yes | Deep Rooted Ring (19109, -0.80 DPS) [world]; Band of the Unicorn (7553, -1.26 DPS, sim-verified) [world_drop]; Underworld Band (1980, -1.32 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindfang (20214, -3.21 DPS) [rep]; High Warlord's War Staff (234549, -5.18 DPS) [pvp]; Ironbark Staff (20220, -15.74 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (11.67 DPS) | yes | High Warlord's Tome of Destruction (234563, -1.61 DPS) [pvp]; Tome of Shadow Force (19309, -2.07 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.22 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Greater Eternal Wand (249237, +0.00 DPS) [crafted]; Lesser Eternal Wand (249232, -0.40 DPS) [crafted]; Dreambough Wand (249234, -0.80 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Earth Warder's Gloves; waist: Frostwolf Cloth Belt; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Advisor's Ring; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 05000000000000000-2035113011220001351-0550005000000000)

Set DPS (verified): 697.1. Weights run: 5.5s. Verify run: 19.9s. 1032 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.380 ± 0.045, crit=0.633 ± 0.021 per rating point (14 rating = 1%, 8.865 per %), hit=1.092 ± 0.078 per rating point (10 rating = 1%, 10.920 per %), spell_haste=6.946 ± 1.021, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.808 ± 0.002, fire_power=0.192 ± 0.000, spirit=not significant (0.626 ± 0.244)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 32.3 spell_power points (19.54 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Ghostweave Hood (254135, -0.03 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.8 spell_power points (9.55 DPS) | yes | Mindburst Medallion (11196, -2.65 DPS) [quest]; Glowing Eye of Mordresh (10769, -3.48 DPS) [dungeon]; Triune Amulet (7722, -4.80 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.8 spell_power points (18.63 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.86 DPS) [crafted]; Shroud of the Nathrezim (18720, -2.33 DPS) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-verified (697.1 DPS) | yes | Spritecaster Cape (11623, -0.58 DPS) [dungeon]; Deep Woodlands Cloak (19121, -3.37 DPS) [quest]; Cloak of Flames (3475, -6.63 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.4 spell_power points (29.88 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.02 DPS) [pvp]; Robe of Winter Night (14136, -6.20 DPS, sim-verified) [crafted]; Legionnaire's Dreadweave Tunic (227094, -10.17 DPS) [pvp] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 29.4 spell_power points (17.79 DPS) | yes | Netherpearl Cuffs (254067, -7.93 DPS) [crafted]; Runecloth Cuffs (254123, -8.84 DPS) [crafted]; Earthenweave Cuffs (254131, -9.17 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.5 spell_power points (18.46 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Dreamweave Gloves (10019, -4.01 DPS) [crafted]; Black Gloves (254101, -8.53 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | 25.1 spell_power points (15.20 DPS) | yes | Frostwolf Cloth Belt (19090, -2.02 DPS) [rep]; Ghostweave Cord (254073, -2.58 DPS) [crafted]; Wisdom of the Timbermaw (19047, -5.74 DPS, sim-verified) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 41.5 spell_power points (25.08 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Magister's Leggings (16687, -3.51 DPS) [dungeon]; Ghostweave Pants (14144, -4.87 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.4 spell_power points (17.16 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.83 DPS) [crafted]; Omnicast Boots (11822, -2.31 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 24.4 spell_power points (14.76 DPS) | yes | Deep Rooted Ring (19109, -5.21 DPS) [world]; Advisor's Ring (19518, -6.30 DPS) [rep]; Underworld Band (1980, -6.77 DPS) [world_drop] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 19.8 spell_power points (11.96 DPS) | yes | Advisor's Ring (19518, -3.50 DPS) [rep]; Underworld Band (1980, -3.97 DPS) [world_drop]; Deep Rooted Ring (19109, -4.27 DPS, sim-verified) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (697.1 DPS) | yes | Ankh of Life (1713, -1.40 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -2.16 DPS) [quest]; Uther's Strength (11302, -3.45 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (697.1 DPS) | yes | Ankh of Life (1713, -0.08 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.84 DPS) [quest]; Uther's Strength (11302, -0.99 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-verified (697.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Mindfang (20214, -6.25 DPS) [rep]; Ironbark Staff (20220, -15.40 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (17.53 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; Tome of Shadow Force (19309, -2.38 DPS) [rep]; Spirit of Aquementas (11904, -4.84 DPS) [quest] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (697.1 DPS) | yes | Torch of Light (279246, +0.00 DPS) [crafted]; Greater Eternal Wand (249237, -1.61 DPS) [crafted]; Lesser Eternal Wand (249232, -2.21 DPS) [crafted] |

**New at 60:** head: Dreamweave Circlet; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Don Julio's Band; trinket2: Rune of the Guard Captain; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Brilliant Wand

No-known-source sample (15 of 1032, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

