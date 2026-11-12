# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000000-0050230100000000)

Set DPS (verified): 62.9. Weights run: 4.1s. Verify run: 5.5s. 199 eligible items had no known source.

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

Set DPS (verified): 105.5. Weights run: 4.0s. Verify run: 4.6s. 326 eligible items had no known source.

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

Set DPS (verified): 183.0. Weights run: 4.2s. Verify run: 6.0s. 455 eligible items had no known source.

Race sweep: human 183.0 ± 0.2, gnome 181.6 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.191 ± 0.016, crit=0.231 ± 0.006 per rating point (14 rating = 1%, 3.233 per %), hit=0.456 ± 0.021 per rating point (10 rating = 1%, 4.560 per %), spell_haste=not significant (0.257 ± 0.236), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.420 ± 0.000, fire_power=0.580 ± 0.001, spirit=not significant (0.121 ± 0.155)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.47 DPS) | yes | Miner's Hat of the Deep (9429, -1.78 DPS, sim-verified) [dungeon]; Living Cowl (5608, -3.23 DPS) [world]; Augural Shroud (2620, -3.26 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Eye of Mordresh (10769, -0.13 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.29 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -2.69 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.32 DPS) | yes | Green Silken Shoulders (7057, -0.25 DPS) [crafted]; Inquisitor's Shawl (19507, -0.50 DPS) [dungeon]; Mantle of Doan (7712, -0.65 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.31 DPS) [crafted]; Guardian Cloak (5965, -0.31 DPS) [crafted]; Caretaker's Cape (19532, -0.45 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (9.58 DPS) | yes | Dreamweave Vest (10021, -0.58 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.11 DPS) [world_drop]; Robe of Power (7054, -2.62 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (3.63 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Phoenix Bindings (210781, -0.59 DPS) [crafted]; Condor Bracers (15864, -0.81 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.6 spell_power points (7.91 DPS) | yes | Black Mageweave Gloves (10003, -2.54 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.70 DPS) [crafted]; Fiery Handwraps (254025, -4.09 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.8 spell_power points (5.95 DPS) | yes | Star Belt (4329, -0.52 DPS) [crafted]; Fiery Cord (254041, -1.77 DPS) [crafted]; Gilded Cord (254037, -1.82 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.3 spell_power points (6.57 DPS) | yes | Gaze Dreamer Pants (6903, -1.29 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.34 DPS) [crafted]; Flame Leggings (253991, -2.51 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 9.4 spell_power points (3.77 DPS) | yes | Gilded Slippers (254001, -0.22 DPS) [crafted]; Spidersilk Boots (4320, -0.64 DPS) [crafted]; Glinteye Slippers (273024, -0.66 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (4.42 DPS) | yes | Snake Hoop (6750, -1.12 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.19 DPS) [quest]; Reedknot Ring (9622, -1.59 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (3.63 DPS) | yes | Snake Hoop (6750, -0.33 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.40 DPS) [quest]; Reedknot Ring (9622, -0.81 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.40 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -10.47 DPS) [dungeon]; Swinetusk Shank (6691, -12.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Fiery Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 455, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 296.6. Weights run: 4.2s. Verify run: 7.9s. 592 eligible items had no known source.

Race sweep: human 296.6 ± 0.2, gnome 294.6 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.188 ± 0.019, crit=0.289 ± 0.007 per rating point (14 rating = 1%, 4.051 per %), hit=0.601 ± 0.028 per rating point (10 rating = 1%, 6.014 per %), spell_haste=2.430 ± 0.307, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.366 ± 0.000, fire_power=0.634 ± 0.001, spirit=not significant (0.250 ± 0.147)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 31.6 spell_power points (15.33 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -2.25 DPS) [crafted]; Dreamweave Circlet (10041, -2.79 DPS) [crafted]; Red Mageweave Headband (10033, -4.30 DPS) [crafted] |
| neck | Talvash's Enhancing Necklace (7673) | Restoring the Necklace [quest] | 12.2 spell_power points (5.91 DPS) | yes | Triune Amulet (7722, -1.44 DPS, sim-verified) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.55 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.79 DPS) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.9 spell_power points (8.70 DPS) | yes | Rotgrip Mantle (17732, -2.02 DPS, sim-verified) [dungeon]; Netherflame Shoulders (254053, -2.02 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.04 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cindercloth Cloak (14044, -3.34 DPS) [crafted]; Nightfall Drape (12465, -3.70 DPS) [dungeon]; Runecloth Cloak (13860, -3.74 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.4 spell_power points (11.82 DPS) | yes | Elemental Raiment (9434, -1.64 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -1.93 DPS) [world]; Dreamweave Vest (10021, -2.96 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.8 spell_power points (5.24 DPS) | yes | Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Arcane Runed Bracers (4744, -0.88 DPS) [quest]; Spidertank Oilrag (9448, -0.88 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (9.94 DPS) | yes | Fiery Gloves (254099, -2.22 DPS) [crafted]; Black Mageweave Gloves (10003, -2.67 DPS) [crafted]; Runecloth Gloves (13863, -3.04 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.9 spell_power points (9.64 DPS) | yes | Ghostweave Cord (254073, -1.82 DPS, sim-verified) [crafted]; Satyrmane Sash (17755, -1.95 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -2.49 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 24.9 spell_power points (12.06 DPS) | yes | Wizardweave Leggings (14132, -1.14 DPS, sim-verified) [crafted]; Knight's Dreadweave Leggings (220888, -3.19 DPS) [vendor]; Senior Designer's Pantaloons (11841, -3.42 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 5 · Sergeant Major · Alliance [vendor] | 15.7 spell_power points (7.61 DPS) | yes | Fiery Sandals (254111, -0.65 DPS) [crafted]; Mistwalker Boots (10629, -1.07 DPS) [dungeon]; Cindercloth Boots (10044, -1.16 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 14.1 spell_power points (6.85 DPS) | yes | Lorekeeper's Ring (19523, -1.03 DPS) [rep]; Cyclopean Band (11824, -1.36 DPS) [dungeon]; Philanthropist's Ring (281635, -1.55 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.30 DPS) | yes | Lorekeeper's Ring (19523, -0.48 DPS) [rep]; Cyclopean Band (11824, -0.82 DPS) [dungeon]; Philanthropist's Ring (281635, -1.00 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.53 DPS, sim-verified) [world_drop]; Ankh of Life (1713, -1.17 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.41 DPS) [quest] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -6.58 DPS, sim-verified) [world_drop] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (6.79 DPS) | yes | Orb of Noh'Orahil (15107, -2.03 DPS) [quest]; Beacon of Hope (9393, -2.42 DPS) [dungeon]; Truesilver Conduit (249455, -2.76 DPS, sim-verified) [crafted] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreambough Wand (249234, -0.96 DPS) [crafted]; Twisted Nether Wand (249144, -1.45 DPS) [crafted]; Lesser Eternal Wand (249232, -1.53 DPS, sim-verified) [crafted] |

**New at 50:** head: Eye of Flame; neck: Talvash's Enhancing Necklace; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Brainlash; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Blade of Eternal Darkness; off_hand: Orb of the Forgotten Seer; ranged: Pyric Caduceus

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 494.4. Weights run: 4.2s. Verify run: 21.7s. 1458 eligible items had no known source.

Race sweep: human 494.4 ± 0.4, gnome 492.2 ± 0.4 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.183 ± 0.033, crit=0.460 ± 0.012 per rating point (14 rating = 1%, 6.443 per %), hit=0.839 ± 0.042 per rating point (10 rating = 1%, 8.388 per %), spell_haste=not significant (-0.749 ± 0.470), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.386 ± 0.001, fire_power=0.614 ± 0.001, spirit=not significant (0.209 ± 0.192)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 44.1 spell_power points (23.78 DPS) | yes | Crimson Felt Hat (18727, -3.67 DPS, sim-verified) [dungeon]; Field Marshal's Coronal (231584, -4.16 DPS) [pvp]; Lieutenant Commander's Dreadweave Cowl (227093, -7.21 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 18.9 spell_power points (10.16 DPS) | yes | Lady Maye's Pendant (14558, -2.36 DPS) [world_drop]; Archlight Talisman (15856, -2.71 DPS) [quest]; Star of Mystaria (12103, -2.80 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -1.12 DPS, sim-verified) [crafted]; Burial Shawl (18681, -1.56 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -5.12 DPS) [dungeon]; Crystalline Threaded Cape (20697, -5.50 DPS, sim-verified) [world]; Shroud of the Exile (15421, -5.69 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (25.69 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.07 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -10.24 DPS) [pvp]; Alanna's Embrace (13314, -11.24 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.9 spell_power points (13.44 DPS) | yes | Runecloth Cuffs (254123, -6.52 DPS) [crafted]; Netherflame Cuffs (254065, -7.88 DPS) [crafted]; Wyrmthalak's Shackles (13958, -7.97 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.3 spell_power points (14.19 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Gloves of Spell Mastery (14146, -2.13 DPS, sim-verified) [crafted]; Earth Warder's Gloves (21318, -2.33 DPS) [quest] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 22.2 spell_power points (11.98 DPS) | yes | Stormpike Cloth Girdle (19094, -1.29 DPS) [rep]; Valconian Sash (12083, -2.81 DPS) [quest]; Highlander's Cloth Girdle (20047, -3.09 DPS, sim-verified) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 46.7 spell_power points (25.18 DPS) | yes | Marshal's Dreadweave Leggings (231587, -3.36 DPS) [pvp]; Skyshroud Leggings (13170, -7.83 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -8.81 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Boots (17562, -0.24 DPS) [pvp]; Venomspew Footpads (275606, -0.88 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.68 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.67 DPS) [dungeon]; Naglering (11669, -12.45 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.42 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.41 DPS) [dungeon]; The Postmaster's Seal (13392, -2.16 DPS) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+6.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of the Beast (13968, -3.51 DPS, sim-verified) [quest]; Smolderweb's Eye (13213, -5.18 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -5.63 DPS) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20069, +0.00 DPS) [rep]; Grand Marshal's Stave (234571, -2.91 DPS) [pvp]; Blade of Eternal Darkness (17780, -17.85 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (15.63 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.87 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.14 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -4.31 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.49 DPS) [dungeon]; Oblivion's Touch (18761, -0.88 DPS) [dungeon]; Brilliant Wand (249385, -0.88 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; legs: Sentinel's Silk Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Second Wind; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1458, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 25500000033000000-0000000000000000000-0050355103101351)

Set DPS (verified): 887.3. Weights run: 4.5s. Verify run: 24.7s. 1458 eligible items had no known source.

Race sweep: human 887.3 ± 0.7, gnome 884.2 ± 0.7 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.261 ± 0.048, crit=0.691 ± 0.017 per rating point (14 rating = 1%, 9.670 per %), hit=1.166 ± 0.075 per rating point (10 rating = 1%, 11.659 per %), spell_haste=not significant (3.702 ± 1.007), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.365 ± 0.001, fire_power=0.635 ± 0.001, spirit=not significant (0.154 ± 0.263)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 46.7 spell_power points (31.89 DPS) | yes | Field Marshal's Coronal (231584, -5.74 DPS) [pvp]; Crimson Felt Hat (18727, -6.89 DPS, sim-verified) [dungeon]; Lieutenant Commander's Dreadweave Cowl (227093, -7.72 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 22.7 spell_power points (15.54 DPS) | yes | Star of Mystaria (12103, -3.20 DPS) [dungeon]; Lady Maye's Pendant (14558, -4.64 DPS) [world_drop]; Emberfury Talisman (12929, -4.83 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.1 spell_power points (20.54 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.43 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.42 DPS) [pvp]; Argent Shoulders (19059, -4.64 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.83 DPS, sim-verified) [world]; Spritecaster Cape (11623, -9.06 DPS) [dungeon]; Shroud of the Exile (15421, -9.60 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.3 spell_power points (33.04 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.89 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -12.39 DPS) [pvp]; Alanna's Embrace (13314, -12.87 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.2 spell_power points (17.20 DPS) | yes | Runecloth Cuffs (254123, -7.90 DPS) [crafted]; Manacle Cuffs (11962, -9.74 DPS) [quest]; Netherflame Cuffs (254065, -9.79 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.9 spell_power points (21.15 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Hands of Power (13253, -4.55 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Handwraps (227100, -6.08 DPS) [pvp] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.3 spell_power points (17.95 DPS) | yes | Highlander's Cloth Girdle (20047, -0.71 DPS) [rep]; Stormpike Cloth Girdle (19094, -3.87 DPS) [rep] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 53.9 spell_power points (36.86 DPS) | yes | Marshal's Dreadweave Leggings (231587, -8.19 DPS) [pvp]; Skyshroud Leggings (13170, -14.22 DPS) [dungeon]; Knight-Captain's Dreadweave Legguards (227095, -15.41 DPS) [pvp] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 23.2 spell_power points (15.88 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.07 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.22 DPS) [pvp] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -1.46 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -1.52 DPS) [dungeon]; Naglering (11669, -15.29 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -0.39 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -0.45 DPS) [dungeon]; Songstone of Ironforge (12543, -1.03 DPS) [quest] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -4.33 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -5.13 DPS) [crafted]; Second Wind (11819, -5.51 DPS, sim-verified) [dungeon] |
| trinket2 | Eye of the Beast (13968) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -3.36 DPS) [pvp]; Headmaster's Charge (13937, -10.93 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -18.45 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Ritssyn's Wand of Bad Mojo (22408) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Oblivion's Touch (18761, -0.12 DPS) [dungeon]; Brilliant Wand (249385, -0.12 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; legs: Sentinel's Silk Leggings; feet: Venomspew Footpads; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Ironbark Staff; ranged: Ritssyn's Wand of Bad Mojo

No-known-source sample (15 of 1458, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 00000000000000000-0000000000000000000-0050230100000000)

Set DPS (verified): 61.9. Weights run: 4.1s. Verify run: 6.0s. 193 eligible items had no known source.

Race sweep: undead 62.5 ± 0.0, orc 61.9 ± 0.0, troll 60.9 ± 0.0 (dps); undead beats the guide race orc by 0.6 (1.00%) ± 0.1, beyond error.

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

Set DPS (verified): 106.1. Weights run: 4.0s. Verify run: 5.3s. 319 eligible items had no known source.

Race sweep: undead 106.2 ± 0.1, orc 106.1 ± 0.1, troll 104.4 ± 0.1 (dps); undead beats the guide race orc by 0.1 (0.06%) ± 0.1, inside error.

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

Set DPS (verified): 183.1. Weights run: 4.2s. Verify run: 6.3s. 448 eligible items had no known source.

Race sweep: orc 183.1 ± 0.2, undead 183.0 ± 0.2, troll 181.2 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.191 ± 0.016, crit=0.231 ± 0.006 per rating point (14 rating = 1%, 3.233 per %), hit=0.456 ± 0.021 per rating point (10 rating = 1%, 4.560 per %), spell_haste=not significant (0.257 ± 0.236), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.420 ± 0.000, fire_power=0.580 ± 0.001, spirit=not significant (0.121 ± 0.155)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.47 DPS) | yes | Miner's Hat of the Deep (9429, -1.51 DPS, sim-verified) [dungeon]; Living Cowl (5608, -3.23 DPS) [world]; Augural Shroud (2620, -3.26 DPS) [world] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Glowing Eye of Mordresh (10769, -0.13 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -0.57 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -2.69 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.32 DPS) | yes | Green Silken Shoulders (7057, -0.25 DPS) [crafted]; Inquisitor's Shawl (19507, -0.50 DPS) [dungeon]; Mantle of Doan (7712, -0.65 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.31 DPS) [crafted]; Guardian Cloak (5965, -0.31 DPS) [crafted]; Battle Healer's Cloak (19528, -0.45 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (9.58 DPS) | yes | Dreamweave Vest (10021, -0.38 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.11 DPS) [world_drop]; Robe of Power (7054, -2.62 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.63 DPS) | yes | Condor Bracers (15864, -0.81 DPS) [quest]; Phoenix Bindings (210781, -1.14 DPS, sim-verified) [crafted]; Radiant Silver Bracers (4545, -1.40 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.6 spell_power points (7.91 DPS) | yes | Black Mageweave Gloves (10003, -1.54 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.70 DPS) [crafted]; Fiery Handwraps (254025, -4.09 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.8 spell_power points (5.95 DPS) | yes | Star Belt (4329, -0.52 DPS) [crafted]; Warsong Sash (16975, -1.52 DPS) [quest]; Fiery Cord (254041, -1.77 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.3 spell_power points (6.57 DPS) | yes | Gaze Dreamer Pants (6903, -1.29 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.34 DPS) [crafted]; Flame Leggings (253991, -2.51 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 9.4 spell_power points (3.77 DPS) | yes | Gilded Slippers (254001, -0.22 DPS) [crafted]; Spidersilk Boots (4320, -0.64 DPS) [crafted]; Glinteye Slippers (273024, -0.66 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (4.42 DPS) | yes | Snake Hoop (6750, -1.12 DPS) [quest]; Reedknot Ring (9622, -1.59 DPS) [quest]; Underworld Band (1980, -1.66 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.63 DPS) | yes | Snake Hoop (6750, -0.33 DPS) [quest]; Reedknot Ring (9622, -0.81 DPS) [quest]; Underworld Band (1980, -0.87 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.78 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -10.47 DPS) [dungeon]; Swinetusk Shank (6691, -12.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Plaguerot Sprig (10766, +0.00 DPS) [dungeon]; Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.40 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Fiery Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 448, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 296.2. Weights run: 4.2s. Verify run: 7.8s. 585 eligible items had no known source.

Race sweep: undead 296.2 ± 0.2, orc 296.2 ± 0.2, troll 293.1 ± 0.2 (dps); undead beats the guide race orc by 0.0 (0.00%) ± 0.3, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.188 ± 0.019, crit=0.289 ± 0.007 per rating point (14 rating = 1%, 4.051 per %), hit=0.601 ± 0.028 per rating point (10 rating = 1%, 6.014 per %), spell_haste=2.430 ± 0.307, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.366 ± 0.000, fire_power=0.634 ± 0.001, spirit=not significant (0.250 ± 0.147)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 31.6 spell_power points (15.33 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -2.28 DPS, sim-verified) [crafted]; Dreamweave Circlet (10041, -2.79 DPS) [crafted]; Red Mageweave Headband (10033, -4.30 DPS) [crafted] |
| neck | Jarkal's Enhancing Necklace (7888) | Necklace Recovery, Take 3 [quest] | 12.2 spell_power points (5.91 DPS) | yes | Prodigious Shadowshard Pendant (17773, -1.55 DPS) [quest]; Triune Amulet (7722, -1.73 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -1.79 DPS) [dungeon] |
| shoulder | Kentic Amice (11624) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.9 spell_power points (8.70 DPS) | yes | Netherflame Shoulders (254053, -2.02 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -2.04 DPS) [vendor]; Rotgrip Mantle (17732, -2.61 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.81 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -2.97 DPS) [crafted]; Cindercloth Cloak (14044, -3.34 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.4 spell_power points (11.82 DPS) | yes | Elemental Raiment (9434, -1.64 DPS) [world_drop]; Dreamweave Vest (10021, -1.66 DPS, sim-verified) [crafted]; Embrace of the Wind Serpent (12462, -1.93 DPS) [world] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 10.8 spell_power points (5.24 DPS) | yes | Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Spidertank Oilrag (9448, -0.88 DPS) [dungeon]; Phoenix Bindings (210781, -1.24 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.5 spell_power points (9.94 DPS) | yes | Fiery Gloves (254099, -2.22 DPS) [crafted]; Black Mageweave Gloves (10003, -2.67 DPS) [crafted]; Runecloth Gloves (13863, -3.57 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 19.9 spell_power points (9.64 DPS) | yes | Ghostweave Cord (254073, -1.52 DPS) [crafted]; Satyrmane Sash (17755, -1.95 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -2.49 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 24.9 spell_power points (12.06 DPS) | yes | Wizardweave Leggings (14132, -2.70 DPS, sim-verified) [crafted]; Stone Guard's Dreadweave Leggings (220906, -3.19 DPS) [vendor]; Senior Designer's Pantaloons (11841, -3.42 DPS) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 5 · First Sergeant · Horde [vendor] | 15.7 spell_power points (7.61 DPS) | yes | Mistwalker Boots (10629, -1.07 DPS) [dungeon]; Cindercloth Boots (10044, -1.16 DPS) [crafted]; Fiery Sandals (254111, -1.24 DPS, sim-verified) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak [dungeon] | 14.1 spell_power points (6.85 DPS) | yes | Advisor's Ring (19519, -1.03 DPS) [rep]; Cyclopean Band (11824, -1.36 DPS) [dungeon]; Philanthropist's Ring (281635, -1.55 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.30 DPS) | yes | Advisor's Ring (19519, -0.48 DPS) [rep]; Cyclopean Band (11824, -0.82 DPS) [dungeon]; Philanthropist's Ring (281635, -1.00 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, -1.17 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.41 DPS) [quest]; Rune of the Guard Captain (19120, -3.87 DPS, sim-verified) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.68 DPS, sim-verified) [quest]; Ankh of Life (1713, -1.45 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.69 DPS) [quest] |
| main_hand | Blade of Eternal Darkness (17780) | Maraudon: Princess Theradras [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Staff of Jordan (873, -5.50 DPS, sim-verified) [world_drop] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (6.79 DPS) | yes | Flask of Forest Mojo (19115, -1.69 DPS) [quest]; Truesilver Conduit (249455, -1.90 DPS, sim-verified) [crafted]; Orb of Noh'Orahil (15107, -2.03 DPS) [quest] |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.48 DPS) [crafted]; Dreambough Wand (249234, -0.96 DPS) [crafted]; Twisted Nether Wand (249144, -1.45 DPS) [crafted] |

**New at 50:** head: Eye of Flame; neck: Jarkal's Enhancing Necklace; shoulder: Kentic Amice; back: Spritecaster Cape; wrist: Netherflame Cuffs; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Brainlash; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Blade of Eternal Darkness; off_hand: Orb of the Forgotten Seer; ranged: Pyric Caduceus

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 490.5. Weights run: 4.2s. Verify run: 21.6s. 1452 eligible items had no known source.

Race sweep: undead 490.9 ± 0.3, orc 490.5 ± 0.3, troll 484.3 ± 0.3 (dps); undead beats the guide race orc by 0.5 (0.09%) ± 0.5, inside error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.183 ± 0.033, crit=0.460 ± 0.012 per rating point (14 rating = 1%, 6.443 per %), hit=0.839 ± 0.042 per rating point (10 rating = 1%, 8.388 per %), spell_haste=not significant (-0.749 ± 0.470), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.386 ± 0.001, fire_power=0.614 ± 0.001, spirit=not significant (0.209 ± 0.192)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 44.1 spell_power points (23.78 DPS) | yes | Crimson Felt Hat (18727, -4.15 DPS, sim-verified) [dungeon]; Warlord's Dreadweave Hood (231590, -4.16 DPS) [pvp]; Champion's Dreadweave Cowl (227090, -7.21 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 18.9 spell_power points (10.16 DPS) | yes | Lady Maye's Pendant (14558, -1.58 DPS, sim-verified) [world_drop]; Archlight Talisman (15856, -2.71 DPS) [quest]; Star of Mystaria (12103, -2.80 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 26.5 spell_power points (14.27 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -0.35 DPS) [crafted]; Burial Shawl (18681, -1.91 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -3.27 DPS, sim-verified) [world]; Spritecaster Cape (11623, -5.12 DPS) [dungeon]; Shroud of the Exile (15421, -5.69 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.6 spell_power points (25.69 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.07 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -10.24 DPS) [pvp]; Alanna's Embrace (13314, -11.50 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.9 spell_power points (13.44 DPS) | yes | Runecloth Cuffs (254123, -6.52 DPS) [crafted]; Netherflame Cuffs (254065, -7.88 DPS) [crafted]; Wyrmthalak's Shackles (13958, -7.97 DPS) [quest] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 26.3 spell_power points (14.19 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Gloves of Spell Mastery (14146, -1.66 DPS, sim-verified) [crafted]; Earth Warder's Gloves (21318, -2.33 DPS) [quest] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 22.2 spell_power points (11.98 DPS) | yes | Defiler's Cloth Girdle (20163, -0.37 DPS) [rep]; Frostwolf Cloth Belt (19090, -1.29 DPS) [rep]; Valconian Sash (12083, -2.81 DPS) [quest] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 33.6 spell_power points (18.10 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -1.72 DPS) [pvp]; Skyshroud Leggings (13170, -1.87 DPS, sim-verified) [dungeon] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 23.7 spell_power points (12.76 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Dragonrider Boots (18102, -1.48 DPS) [dungeon]; Omnicast Boots (11822, -2.52 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.68 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.67 DPS) [dungeon]; Naglering (11669, -12.58 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.42 DPS) [quest]; Don Mauricio's Band of Domination (22433, -1.41 DPS) [dungeon]; The Postmaster's Seal (13392, -2.16 DPS) [dungeon] |
| trinket1 | Eye of the Beast (13968) | For The Horde! [quest] | sim-verified (+7.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, -3.34 DPS, sim-verified) [dungeon]; Smolderweb's Eye (13213, -5.18 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -5.63 DPS) [crafted] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironbark Staff (20220, +0.00 DPS) [rep]; High Warlord's War Staff (234549, -2.91 DPS) [pvp]; Blade of Eternal Darkness (17780, -17.84 DPS, sim-verified) [dungeon] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (15.63 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.87 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.20 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -4.31 DPS) [quest] |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.49 DPS) [dungeon]; Brilliant Wand (249385, -0.88 DPS) [crafted]; Oblivion's Touch (18761, -1.75 DPS, sim-verified) [dungeon] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; legs: Outrider's Silk Leggings; feet: Venomspew Footpads; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Eye of the Beast; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1452, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 25500000033000000-0000000000000000000-0050355103101351)

Set DPS (verified): 878.7. Weights run: 4.5s. Verify run: 26.3s. 1452 eligible items had no known source.

Race sweep: orc 878.7 ± 0.7, undead 875.1 ± 0.7, troll 867.6 ± 0.7 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.261 ± 0.048, crit=0.691 ± 0.017 per rating point (14 rating = 1%, 9.670 per %), hit=1.166 ± 0.075 per rating point (10 rating = 1%, 11.659 per %), spell_haste=not significant (3.702 ± 1.007), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.365 ± 0.001, fire_power=0.635 ± 0.001, spirit=not significant (0.154 ± 0.263)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellweaver's Turban (22267) | Blackrock Spire: General Drakkisath [dungeon] | 46.7 spell_power points (31.89 DPS) | yes | Warlord's Dreadweave Hood (231590, -5.74 DPS) [pvp]; Crimson Felt Hat (18727, -6.69 DPS, sim-verified) [dungeon]; Champion's Dreadweave Cowl (227090, -7.72 DPS) [pvp] |
| neck | Nacreous Shell Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 22.7 spell_power points (15.54 DPS) | yes | Star of Mystaria (12103, -3.67 DPS, sim-verified) [dungeon]; Lady Maye's Pendant (14558, -4.64 DPS) [world_drop]; Emberfury Talisman (12929, -4.83 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 30.1 spell_power points (20.54 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.43 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -3.42 DPS) [pvp]; Argent Shoulders (19059, -5.72 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.54 DPS, sim-verified) [world]; Spritecaster Cape (11623, -9.06 DPS) [dungeon]; Shroud of the Exile (15421, -9.60 DPS) [quest] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.3 spell_power points (33.04 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.89 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -12.39 DPS) [pvp]; Alanna's Embrace (13314, -15.02 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.2 spell_power points (17.20 DPS) | yes | Runecloth Cuffs (254123, -7.90 DPS) [crafted]; Manacle Cuffs (11962, -9.74 DPS) [quest]; Netherflame Cuffs (254065, -9.79 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 30.9 spell_power points (21.15 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -4.68 DPS, sim-verified) [dungeon]; Blood Guard's Dreadweave Handwraps (227099, -6.08 DPS) [pvp] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Ban'thok Sash (11662, -1.51 DPS, sim-verified) [dungeon]; Frostwolf Cloth Belt (19090, -3.16 DPS) [rep]; Valconian Sash (12083, -5.63 DPS) [quest] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 34.5 spell_power points (23.57 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.94 DPS) [dungeon]; Legionnaire's Dreadweave Legguards (227097, -2.12 DPS) [pvp] |
| feet | Venomspew Footpads (275606) | Tailoring [crafted] | 23.2 spell_power points (15.88 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.07 DPS) [dungeon]; Dragonrider Boots (18102, -0.73 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -1.46 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -1.52 DPS) [dungeon]; Naglering (11669, -16.80 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Maiden's Circle (13001, -0.39 DPS) [world_drop]; Don Mauricio's Band of Domination (22433, -0.45 DPS) [dungeon]; Eye of Orgrimmar (12545, -1.03 DPS) [quest] |
| trinket1 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Smolderweb's Eye (13213, -4.33 DPS) [dungeon]; Frozen Heart of the Mountain (249469, -5.13 DPS) [crafted]; Second Wind (11819, -5.35 DPS, sim-verified) [dungeon] |
| trinket2 | Eye of the Beast (13968) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smolderweb's Eye (13213, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -3.36 DPS) [pvp]; Headmaster's Charge (13937, -15.81 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -18.45 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Bonecreeper Stylus (13938) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ritssyn's Wand of Bad Mojo (22408, -0.89 DPS) [dungeon]; Oblivion's Touch (18761, -1.01 DPS) [dungeon]; Brilliant Wand (249385, -1.01 DPS) [crafted] |

**New at 60:** head: Spellweaver's Turban; neck: Nacreous Shell Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Venomspew Footpads; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Weakness Analyzer; trinket2: Eye of the Beast; main_hand: Ironbark Staff; ranged: Bonecreeper Stylus

No-known-source sample (15 of 1452, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

