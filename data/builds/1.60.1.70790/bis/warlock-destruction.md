# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (human, 00000000000000000-0000000000000000000-0050230100000000)

Set DPS (verified): 62.9. Weights run: 4.2s. Verify run: 5.7s. 182 eligible items had no known source.

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

No-known-source sample (15 of 182, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (human, 00000000000000000-0000000000000000000-1050235103001000)

Set DPS (verified): 104.8. Weights run: 3.9s. Verify run: 4.1s. 309 eligible items had no known source.

Race sweep: human 104.8 ± 0.1, gnome 104.4 ± 0.1 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.082 ± 0.005, crit=0.071 ± 0.002 per rating point (14 rating = 1%, 0.988 per %), hit=0.110 ± 0.005 per rating point (10 rating = 1%, 1.100 per %), spell_haste=0.250 ± 0.062, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.794 ± 0.001, fire_power=0.206 ± 0.000, spirit=not significant (0.053 ± 0.052)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.3 spell_power points (3.64 DPS) | yes | Filigreed Shadow Circlet (253981, -0.06 DPS) [crafted]; Silk Headband (7050, -0.74 DPS) [crafted]; Shadow Circlet (253955, -0.83 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 0.5 spell_power points (0.17 DPS) | yes | Archmagister's Faceted Pendant (273049, -0.02 DPS) [world]; Crystal Starfire Medallion (5003, -0.02 DPS) [world_drop]; Seal of the Expedition (271920, -0.04 DPS) [rep] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.7 spell_power points (3.13 DPS) | yes | Invoker's Mantle (215365, -0.36 DPS, sim-verified) [crafted]; Moonlit Amice (11884, -0.88 DPS) [quest]; Death Speaker Mantle (6685, -0.91 DPS) [dungeon] |
| back | Vine Pruner's Cloak (279835) | A Green Sample [quest] | sim-verified (104.8 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.25 DPS) [crafted]; Caretaker's Cape (19533, -0.25 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (4.18 DPS) | yes | Necrology Robes (2292, -0.47 DPS) [dungeon]; Shadow Gown (253967, -0.61 DPS) [crafted]; Beguiler Robes (7728, -0.84 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.89 DPS) | yes | Phoenix Bindings (210781, -2.45 DPS, sim-verified) [crafted]; Windsong Bangles (263336, -2.57 DPS) [quest]; Glowing Magical Bracelets (13106, -2.68 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.7 spell_power points (2.48 DPS) | yes | Hands of Darkness (7047, -0.10 DPS) [crafted]; Serpent Gloves (5970, -0.23 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.55 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.2 spell_power points (3.62 DPS) | yes | Belt of Arugal (6392, -0.56 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -1.04 DPS) [dungeon]; Invoker's Cord (215366, -1.23 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.5 spell_power points (4.01 DPS) | yes | Shadow Leggings (253993, -0.17 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.98 DPS) [crafted]; Necromancer Leggings (2277, -1.17 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 11.1 spell_power points (3.57 DPS) | yes | Gilded Slippers (254001, -1.20 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.22 DPS) [crafted]; Glinteye Slippers (273024, -1.22 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (2.25 DPS) | yes | Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Minor Channeling Ring (1449, -0.59 DPS) [quest]; Electrocutioner Lagnut (9447, -1.17 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.4 spell_power points (2.06 DPS) | yes | Ring of Power Regulation (273030, -0.13 DPS) [dungeon]; Minor Channeling Ring (1449, -0.40 DPS) [quest]; Electrocutioner Lagnut (9447, -0.98 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-verified (104.8 DPS) | yes | Hydrocane (9452, -1.35 DPS) [dungeon]; Spine of the Devourer (273040, -1.35 DPS) [world]; Tendonscraper Dagger (273053, -1.40 DPS, sim-verified) [world] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (2.89 DPS) | yes | Orb of Mystic Insight (249394, -0.49 DPS) [crafted]; Orb of Mistmantle (13031, -1.29 DPS) [world_drop]; Orb of Souls (249395, -1.51 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-verified (104.8 DPS) | yes | Wand of Eventide (5214, -0.32 DPS) [world_drop]; Greater Mystic Wand (217287, -0.32 DPS) [crafted]; Gravestone Scepter (7001, -0.35 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Lorekeeper's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 309, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (human, 00000000000000000-0000000000000000000-2050235103101341)

Set DPS (verified): 189.5. Weights run: 4.2s. Verify run: 7.4s. 399 eligible items had no known source.

Race sweep: human 189.5 ± 0.2, gnome 188.3 ± 0.2 (dps); the guide race human is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.190 ± 0.017, crit=0.233 ± 0.006 per rating point (14 rating = 1%, 3.259 per %), hit=0.493 ± 0.023 per rating point (10 rating = 1%, 4.932 per %), spell_haste=not significant (0.249 ± 0.235), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.432 ± 0.000, fire_power=0.568 ± 0.001, spirit=not significant (0.127 ± 0.164)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.44 DPS) | yes | Living Cowl (5608, -3.23 DPS, sim-verified) [world]; Augural Shroud (2620, -3.25 DPS) [world]; Papal Fez (9431, -3.46 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.9 spell_power points (3.57 DPS) | yes | Glowing Eye of Mordresh (10769, -0.79 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.68 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.81 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.31 DPS) | yes | Inquisitor's Shawl (19507, -0.50 DPS) [dungeon]; Mantle of Doan (7712, -0.62 DPS) [dungeon]; Green Silken Shoulders (7057, -0.93 DPS, sim-verified) [crafted] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Long Silken Cloak (4326, -0.33 DPS) [crafted]; Guardian Cloak (5965, -0.33 DPS) [crafted]; Caretaker's Cape (19532, -0.45 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (9.56 DPS) | yes | Dreamweave Vest (10021, -0.92 DPS) [crafted]; Elemental Raiment (9434, -1.12 DPS) [world_drop]; Robe of Power (7054, -2.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Arcane Runed Bracers (4744, -0.61 DPS, sim-verified) [quest]; Phoenix Bindings (210781, -0.65 DPS) [crafted]; Condor Bracers (15864, -0.80 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.7 spell_power points (7.90 DPS) | yes | Black Mageweave Gloves (10003, -1.95 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.71 DPS) [crafted]; Gilded Handwraps (254021, -4.15 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.8 spell_power points (5.93 DPS) | yes | Star Belt (4329, -0.50 DPS) [crafted]; Gilded Cord (254037, -1.80 DPS) [crafted]; Fiery Cord (254041, -1.82 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.3 spell_power points (6.54 DPS) | yes | Gaze Dreamer Pants (6903, -0.95 DPS, sim-verified) [dungeon]; Crimson Silk Pantaloons (7062, -2.33 DPS) [crafted]; Stoneweaver Leggings (9407, -2.51 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.9 spell_power points (10.00 DPS) | yes | Fiery Slippers (254005, -6.38 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -6.45 DPS) [crafted]; Glinteye Slippers (273024, -6.88 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (4.40 DPS) | yes | Ring of Forlorn Spirits (2043, -1.19 DPS) [quest]; Reedknot Ring (9622, -1.59 DPS) [quest]; Underworld Band (1980, -1.59 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (3.62 DPS) | yes | Ring of Forlorn Spirits (2043, -0.40 DPS) [quest]; Reedknot Ring (9622, -0.80 DPS) [quest]; Underworld Band (1980, -0.80 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -9.66 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -10.68 DPS) [dungeon]; Swinetusk Shank (6691, -12.85 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Burning Sliver (5249, +0.00 DPS) [quest]; Thunderwood (13062, +0.00 DPS) [world_drop]; Greater Mystic Wand (217287, -0.40 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 399, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (human, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 289.5. Weights run: 4.4s. Verify run: 6.5s. 506 eligible items had no known source.

Race sweep: human 289.5 ± 0.2, gnome 287.7 ± 0.2 (dps); the guide race human is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.428 ± 0.017, crit=0.168 ± 0.004 per rating point (14 rating = 1%, 2.353 per %), hit=0.340 ± 0.019 per rating point (10 rating = 1%, 3.401 per %), spell_haste=not significant (-0.937 ± 0.256), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.426 ± 0.000, fire_power=0.574 ± 0.001, spirit=not significant (0.548 ± 0.268)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 34.5 spell_power points (14.20 DPS) | yes | Dreamweave Circlet (10041, -2.70 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.85 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.08 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 11.8 spell_power points (4.88 DPS) | yes | Mindburst Medallion (11196, -0.22 DPS) [quest]; Glowing Eye of Mordresh (10769, -0.58 DPS) [dungeon]; Lei of Lilies (1315, -1.49 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherflame Shoulders (254053, -1.74 DPS) [crafted]; Nethergeld Shoulders (254049, -2.06 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Runecloth Cloak (13860, -2.57 DPS, sim-verified) [crafted]; Cindercloth Cloak (14044, -3.70 DPS) [crafted]; Caretaker's Cape (19531, -3.94 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 29.5 spell_power points (12.17 DPS) | yes | Acumen Robes (17775, -0.81 DPS) [quest]; Robe of the Magi (1716, -0.91 DPS) [world_drop]; Embrace of the Wind Serpent (12462, -2.39 DPS) [world] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 13.2 spell_power points (5.44 DPS) | yes | Nethergeld Cuffs (254061, -0.19 DPS) [crafted]; Netherpearl Cuffs (254067, -0.80 DPS) [crafted]; Bloodband Bracers (11469, -1.12 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.5 spell_power points (9.71 DPS) | yes | Fiery Gloves (254099, -1.64 DPS) [crafted]; Gilded Gloves (254095, -2.01 DPS) [crafted]; Runecloth Gloves (13863, -2.56 DPS, sim-verified) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, -0.72 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -1.78 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 27.3 spell_power points (11.24 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.76 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -3.21 DPS) [vendor]; Senior Designer's Pantaloons (11841, -4.54 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (11.47 DPS) | yes | Vinerot Sandals (17748, -5.15 DPS) [dungeon]; Sergeant Major's Dreadweave Boots (220891, -5.19 DPS) [vendor]; Fiery Sandals (254111, -5.49 DPS, sim-verified) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.2 spell_power points (5.85 DPS) | yes | Philanthropist's Ring (281635, -0.84 DPS) [quest]; Lorekeeper's Ring (19523, -0.90 DPS) [rep]; Brainlash (6440, -2.07 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.36 DPS) | yes | Philanthropist's Ring (281635, -0.35 DPS) [quest]; Lorekeeper's Ring (19523, -0.41 DPS) [rep]; Brainlash (6440, -1.58 DPS) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spire of Hakkar (10844, -3.45 DPS) [world]; Kindling Stave (11750, -4.03 DPS, sim-verified) [dungeon]; Glowing Brightwood Staff (812, -4.86 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.46 DPS) [crafted]; Dreambough Wand (249234, -0.87 DPS) [crafted]; Charged Lightning Rod (11860, -0.99 DPS) [quest] |

**New at 50:** head: Eye of Flame; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Dreamweave Vest; wrist: Netherflame Cuffs; waist: Ghostweave Cord; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Band of the Unicorn; trinket1: Thunderbrew's Boot Flask; trinket2: Frozen Heart of the Mountain; ranged: Pyric Caduceus

No-known-source sample (15 of 506, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (human, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 485.7. Weights run: 4.3s. Verify run: 22.3s. 1299 eligible items had no known source.

Race sweep: human 485.7 ± 0.4, gnome 483.5 ± 0.4 (dps); the guide race human is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.157 ± 0.023, crit=0.233 ± 0.007 per rating point (14 rating = 1%, 3.268 per %), hit=0.526 ± 0.028 per rating point (10 rating = 1%, 5.261 per %), spell_haste=not significant (-0.893 ± 0.375), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.423 ± 0.001, fire_power=0.577 ± 0.001, spirit=not significant (0.497 ± 0.381)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 35.2 spell_power points (16.36 DPS) | yes | Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Eye of Flame (3075, -1.81 DPS) [world_drop]; Ghostweave Hood (254135, -2.45 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Triune Amulet (7722, -0.97 DPS) [dungeon]; Star of Mystaria (12103, -1.36 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 28.5 spell_power points (13.23 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.38 DPS) [pvp]; Mantle of the Timbermaw (19050, -2.86 DPS) [crafted]; Burial Shawl (18681, -4.43 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -2.13 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -3.96 DPS) [rep]; Crystalline Threaded Cape (20697, -4.45 DPS, sim-verified) [world] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.4 spell_power points (22.02 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.41 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -8.95 DPS) [pvp]; Alanna's Embrace (13314, -10.34 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 26.7 spell_power points (12.42 DPS) | yes | Earthenweave Cuffs (254131, -6.39 DPS) [crafted]; Runecloth Cuffs (254123, -6.58 DPS) [crafted]; Netherflame Cuffs (254065, -7.27 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hands of Power (13253, +0.00 DPS) [dungeon]; Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Earth Warder's Gloves (21318, -0.05 DPS) [quest] |
| waist | Stormpike Cloth Girdle (19094) | Stormpike Guard [rep] | 19.6 spell_power points (9.09 DPS) | yes | Ban'thok Sash (11662, -0.27 DPS) [dungeon]; Thuzadin Sash (18740, -0.57 DPS) [dungeon]; Ghostweave Cord (254073, -3.25 DPS, sim-verified) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 40.1 spell_power points (18.62 DPS) | yes | Marshal's Dreadweave Leggings (231587, -0.06 DPS) [pvp]; Ghostweave Pants (14144, -1.93 DPS) [crafted]; Skyshroud Leggings (13170, -2.25 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.5 spell_power points (12.76 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.70 DPS) [crafted]; Ghostweave Boots (254129, -2.33 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.46 DPS) [quest]; Don Mauricio's Band of Domination (22433, -3.63 DPS) [dungeon]; Naglering (11669, -13.21 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, +0.00 DPS) [quest]; Don Mauricio's Band of Domination (22433, -2.31 DPS) [dungeon]; Lorekeeper's Ring (19522, -2.43 DPS) [rep] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.79 DPS) [world_drop]; Ankh of Life (1713, -2.80 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.26 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.30 DPS, sim-verified) [vendor]; Uther's Strength (11302, -7.43 DPS) [world_drop]; Ankh of Life (1713, -7.45 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sageclaw (20070, -4.15 DPS) [rep]; Grand Marshal's Stave (234571, -4.20 DPS) [pvp]; Ironbark Staff (20069, -18.39 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (13.47 DPS) | yes | Grand Marshal's Tome of Power (234589, -0.91 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.25 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.71 DPS) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.22 DPS) [dungeon]; Brilliant Wand (249385, -0.64 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; waist: Stormpike Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (human, 25500000033000000-0000000000000000000-0050355103101351)

Set DPS (verified): 868.9. Weights run: 4.6s. Verify run: 24.3s. 1299 eligible items had no known source.

Race sweep: human 868.9 ± 0.8, gnome 865.6 ± 0.7 (dps); the guide race human is best.

4 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.338 ± 0.048, crit=0.694 ± 0.017 per rating point (14 rating = 1%, 9.718 per %), hit=1.128 ± 0.072 per rating point (10 rating = 1%, 11.281 per %), spell_haste=not significant (3.155 ± 1.022), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.368 ± 0.001, fire_power=0.632 ± 0.001, spirit=not significant (-0.005 ± 0.267)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crimson Felt Hat (18727, +0.00 DPS) [dungeon]; Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Star of Mystaria (12103, -0.72 DPS) [dungeon]; Dark Advisor's Pendant (18691, -3.83 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.1 spell_power points (21.11 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, -0.25 DPS) [pvp]; Lieutenant Commander's Dreadweave Spaulders (227091, -3.39 DPS) [pvp]; Burial Shawl (18681, -3.87 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -5.86 DPS) [world]; Spritecaster Cape (11623, -9.47 DPS) [dungeon]; Cloak of Flames (3475, -9.90 DPS, sim-verified) [world_drop] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.0 spell_power points (33.27 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -6.06 DPS) [pvp]; Knight-Captain's Dreadweave Tunic (227096, -11.72 DPS) [pvp]; Alanna's Embrace (13314, -15.91 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 24.7 spell_power points (16.76 DPS) | yes | Runecloth Cuffs (254123, -7.00 DPS) [crafted]; Netherflame Cuffs (254065, -9.58 DPS) [crafted]; Arena Wristguards (18709, -10.17 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.8 spell_power points (21.58 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Hands of Power (13253, -2.57 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Handwraps (227100, -6.42 DPS) [pvp] |
| waist | Highlander's Cloth Girdle (20047) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Stormpike Cloth Girdle (19094, -2.96 DPS) [rep]; Clutch of Andros (13956, -5.46 DPS) [dungeon] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 54.8 spell_power points (37.19 DPS) | yes | Marshal's Dreadweave Leggings (231587, -7.73 DPS) [pvp]; Knight-Captain's Dreadweave Legguards (227095, -15.21 DPS) [pvp]; Skyshroud Leggings (13170, -15.89 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -1.94 DPS, sim-verified) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.38 DPS) [dungeon]; Maiden's Circle (13001, -0.62 DPS) [world_drop]; Naglering (11669, -16.80 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.19 DPS) [dungeon]; Songstone of Ironforge (12543, -0.43 DPS) [quest]; Maiden's Circle (13001, -0.43 DPS) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+9.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.25 DPS) [crafted]; Uther's Strength (11302, -4.07 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frozen Heart of the Mountain (249469, -3.18 DPS, sim-verified) [crafted]; Serenity Field (272439, -5.43 DPS) [vendor]; Uther's Strength (11302, -10.85 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, -2.91 DPS) [pvp]; Staff of Jordan (873, -14.63 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -26.28 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.25 DPS) [dungeon]; Greater Eternal Wand (249237, -1.61 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Highlander's Cloth Girdle; legs: Sentinel's Silk Leggings; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 00000000000000000-0000000000000000000-0050230100000000)

Set DPS (verified): 61.9. Weights run: 4.2s. Verify run: 6.2s. 176 eligible items had no known source.

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

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 23173 Abomination Skin Leggings

### Band 30 (orc, 00000000000000000-0000000000000000000-1050235103001000)

Set DPS (verified): 105.5. Weights run: 3.9s. Verify run: 6.4s. 302 eligible items had no known source.

Race sweep: undead 105.8 ± 0.1, orc 105.5 ± 0.1, troll 103.8 ± 0.1 (dps); undead beats the guide race orc by 0.3 (0.25%) ± 0.1, beyond error.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.166 ± 0.008, crit=0.147 ± 0.004 per rating point (14 rating = 1%, 2.053 per %), hit=0.258 ± 0.011 per rating point (10 rating = 1%, 2.578 per %), spell_haste=not significant (0.176 ± 0.132), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.810 ± 0.001, fire_power=0.190 ± 0.000, spirit=not significant (0.297 ± 0.083)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 12.8 spell_power points (4.38 DPS) | yes | Filigreed Shadow Circlet (253981, -0.49 DPS) [crafted]; Silk Headband (7050, -1.29 DPS) [crafted]; Shadow Circlet (253955, -1.32 DPS) [crafted] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 1.9 spell_power points (0.63 DPS) | yes | Pendant of Myzrael (4614, -0.02 DPS) [dungeon]; Glowing Green Talisman (5002, -0.02 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.10 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.5 spell_power points (3.60 DPS) | yes | Death Speaker Mantle (6685, -0.91 DPS) [dungeon]; Invoker's Mantle (215365, -1.09 DPS, sim-verified) [crafted]; Magician's Mantle (12998, -1.37 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glowing Thresher Cape (6901, -0.28 DPS) [dungeon]; Heavy Woolen Cloak (4311, -0.34 DPS) [crafted]; Battle Healer's Cloak (19529, -0.34 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Gown (253967, -0.04 DPS) [crafted]; Beguiler Robes (7728, -0.26 DPS, sim-verified) [dungeon]; Necrology Robes (2292, -0.28 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.08 DPS) | yes | Tabitha's Cuffs (251486, -2.03 DPS, sim-verified) [quest]; Phoenix Bindings (210781, -2.24 DPS) [crafted]; Nightsky Wristbands (6407, -2.44 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Hands of Darkness (7047, -0.31 DPS, sim-verified) [crafted]; Gloves of Meditation (4318, -0.47 DPS) [crafted]; Serpent Gloves (5970, -0.50 DPS) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 spell_power points (3.94 DPS) | yes | Warsong Sash (16975, -0.17 DPS) [quest]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -1.20 DPS) [dungeon] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gaze Dreamer Pants (6903, -0.17 DPS, sim-verified) [dungeon]; Filigreed Shadow Leggings (253943, -0.99 DPS) [crafted]; Necromancer Leggings (2277, -1.39 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 12.9 spell_power points (4.41 DPS) | yes | Gilded Slippers (254001, -1.23 DPS, sim-verified) [crafted]; Glinteye Slippers (273024, -1.40 DPS) [dungeon]; Spidersilk Boots (4320, -1.79 DPS) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (2.40 DPS) | yes | Ring of Power Regulation (273030, -0.34 DPS) [dungeon]; Solid Gold Band (279845, -0.52 DPS) [quest]; Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 6.8 spell_power points (2.34 DPS) | yes | Ring of Power Regulation (273030, -0.15 DPS, sim-verified) [dungeon]; Solid Gold Band (279845, -0.46 DPS) [quest]; Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Hydrocane (9452, -1.78 DPS) [dungeon]; Tendonscraper Dagger (273053, -1.78 DPS) [world]; Soulstaff (249393, -4.82 DPS, sim-verified) [crafted] |
| off_hand | Orb of Soran'ruk (6898) | The Orb of Soran'ruk [quest] | 9.0 spell_power points (3.08 DPS) | yes | Orb of Mystic Insight (249394, -0.34 DPS) [crafted]; Orb of Mistmantle (13031, -1.01 DPS) [world_drop]; Orb of Souls (249395, -1.10 DPS) [crafted] |
| ranged | Thunderwood (13062) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gravestone Scepter (7001, -0.22 DPS, sim-verified) [quest]; Wand of Eventide (5214, -0.34 DPS) [world_drop]; Greater Mystic Wand (217287, -0.34 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Defiler's Cloth Girdle; legs: Shadow Leggings; feet: Black Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Swinetusk Shank; off_hand: Orb of Soran'ruk; ranged: Thunderwood

No-known-source sample (15 of 302, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (orc, 00000000000000000-0000000000000000000-2050235103101341)

Set DPS (verified): 189.8. Weights run: 4.2s. Verify run: 7.2s. 392 eligible items had no known source.

Race sweep: orc 189.8 ± 0.2, undead 189.2 ± 0.2, troll 187.2 ± 0.2 (dps); the guide race orc is best.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.190 ± 0.017, crit=0.233 ± 0.006 per rating point (14 rating = 1%, 3.259 per %), hit=0.493 ± 0.023 per rating point (10 rating = 1%, 4.932 per %), spell_haste=not significant (0.249 ± 0.235), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.432 ± 0.000, fire_power=0.568 ± 0.001, spirit=not significant (0.127 ± 0.164)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (8.44 DPS) | yes | Living Cowl (5608, -3.23 DPS, sim-verified) [world]; Augural Shroud (2620, -3.25 DPS) [world]; Papal Fez (9431, -3.46 DPS) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 8.9 spell_power points (3.57 DPS) | yes | Glowing Eye of Mordresh (10769, -0.15 DPS) [dungeon]; Necklace of Calisea (1714, -2.68 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.81 DPS) [quest] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (4.31 DPS) | yes | Green Silken Shoulders (7057, -0.33 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.50 DPS) [dungeon]; Mantle of Doan (7712, -0.62 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Cloak (5965, -0.33 DPS) [crafted]; Battle Healer's Cloak (19528, -0.45 DPS) [rep]; Long Silken Cloak (4326, -0.72 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.8 spell_power points (9.56 DPS) | yes | Dreamweave Vest (10021, -0.50 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -1.12 DPS) [world_drop]; Robe of Power (7054, -2.60 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (3.62 DPS) | yes | Phoenix Bindings (210781, -0.65 DPS) [crafted]; Condor Bracers (15864, -0.80 DPS) [quest]; Radiant Silver Bracers (4545, -1.40 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.7 spell_power points (7.90 DPS) | yes | Black Mageweave Gloves (10003, -1.76 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.71 DPS) [crafted]; Gilded Handwraps (254021, -4.15 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.8 spell_power points (5.93 DPS) | yes | Star Belt (4329, -0.50 DPS) [crafted]; Warsong Sash (16975, -1.51 DPS) [quest]; Gilded Cord (254037, -1.80 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 16.3 spell_power points (6.54 DPS) | yes | Gaze Dreamer Pants (6903, -1.26 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.33 DPS) [crafted]; Stoneweaver Leggings (9407, -2.51 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.9 spell_power points (10.00 DPS) | yes | Fiery Slippers (254005, -5.86 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -6.45 DPS) [crafted]; Glinteye Slippers (273024, -6.88 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (4.40 DPS) | yes | Reedknot Ring (9622, -1.59 DPS) [quest]; Underworld Band (1980, -1.59 DPS) [world_drop]; Snake Hoop (6750, -1.61 DPS) [quest] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (3.62 DPS) | yes | Reedknot Ring (9622, -0.80 DPS) [quest]; Underworld Band (1980, -0.80 DPS) [world_drop]; Snake Hoop (6750, -0.82 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -8.63 DPS, sim-verified) [dungeon]; Hypnotic Blade (7714, -10.68 DPS) [dungeon]; Swinetusk Shank (6691, -12.85 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thunderwood (13062, +0.00 DPS) [world_drop]; Wand of Eventide (5214, -0.40 DPS) [world_drop]; Greater Mystic Wand (217287, -0.40 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Triune Amulet; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 392, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (orc, 24000000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 289.2. Weights run: 4.4s. Verify run: 8.3s. 499 eligible items had no known source.

Race sweep: orc 289.2 ± 0.2, undead 289.0 ± 0.2, troll 286.3 ± 0.2 (dps); the guide race orc is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.428 ± 0.017, crit=0.168 ± 0.004 per rating point (14 rating = 1%, 2.353 per %), hit=0.340 ± 0.019 per rating point (10 rating = 1%, 3.401 per %), spell_haste=not significant (-0.937 ± 0.256), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.426 ± 0.000, fire_power=0.574 ± 0.001, spirit=not significant (0.548 ± 0.268)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Flame (3075) | World drop [world_drop] | 34.5 spell_power points (14.20 DPS) | yes | Dreamweave Circlet (10041, -2.39 DPS, sim-verified) [crafted]; Red Mageweave Headband (10033, -2.85 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -3.08 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 11.8 spell_power points (4.88 DPS) | yes | Glowing Eye of Mordresh (10769, -0.58 DPS) [dungeon]; Mindburst Medallion (11196, -1.38 DPS, sim-verified) [quest]; Lei of Lilies (1315, -1.49 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kentic Amice (11624, +0.00 DPS) [dungeon]; Netherflame Shoulders (254053, -1.74 DPS) [crafted]; Nethergeld Shoulders (254049, -2.06 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deep Woodlands Cloak (19121, -1.12 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -3.06 DPS) [crafted]; Cindercloth Cloak (14044, -3.70 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of the Magi (1716, -0.10 DPS) [world_drop]; Dreamweave Vest (10021, -0.91 DPS, sim-verified) [crafted]; Embrace of the Wind Serpent (12462, -1.58 DPS) [world] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 13.2 spell_power points (5.44 DPS) | yes | Nethergeld Cuffs (254061, -0.19 DPS) [crafted]; Netherpearl Cuffs (254067, -0.80 DPS) [crafted]; Bloodband Bracers (11469, -1.12 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 23.5 spell_power points (9.71 DPS) | yes | Fiery Gloves (254099, -1.64 DPS) [crafted]; Gilded Gloves (254095, -2.01 DPS) [crafted]; Runecloth Gloves (13863, -2.99 DPS, sim-verified) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, -0.72 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.78 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak [dungeon] | 27.3 spell_power points (11.24 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -2.76 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -3.21 DPS) [vendor]; Senior Designer's Pantaloons (11841, -3.90 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.8 spell_power points (11.47 DPS) | yes | Vinerot Sandals (17748, -5.15 DPS) [dungeon]; First Sergeant's Dreadweave Boots (220909, -5.19 DPS) [vendor]; Fiery Sandals (254111, -6.28 DPS, sim-verified) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.2 spell_power points (5.85 DPS) | yes | Philanthropist's Ring (281635, -0.84 DPS) [quest]; Advisor's Ring (19519, -0.90 DPS) [rep]; Brainlash (6440, -2.07 DPS) [dungeon] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.36 DPS) | yes | Advisor's Ring (19519, -0.41 DPS) [rep]; Brainlash (6440, -1.58 DPS) [dungeon]; Philanthropist's Ring (281635, -2.18 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.28 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.21 DPS) [quest]; Rune of the Guard Captain (19120, -0.76 DPS, sim-verified) [quest] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Kindling Stave (11750, -1.67 DPS, sim-verified) [dungeon]; Spire of Hakkar (10844, -3.45 DPS) [world]; Glowing Brightwood Staff (812, -4.86 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Lesser Eternal Wand (249232, -0.46 DPS) [crafted]; Dreambough Wand (249234, -0.87 DPS) [crafted]; Charged Lightning Rod (11860, -0.99 DPS) [quest] |

**New at 50:** head: Eye of Flame; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherflame Cuffs; waist: Ghostweave Cord; legs: Spellshock Leggings; finger1: Cyclopean Band; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Pyric Caduceus

No-known-source sample (15 of 499, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (orc, 25522000000000000-0000000000000000000-2053235103101351)

Set DPS (verified): 485.0. Weights run: 4.3s. Verify run: 23.0s. 1293 eligible items had no known source.

Race sweep: orc 485.0 ± 0.4, undead 484.7 ± 0.4, troll 479.1 ± 0.4 (dps); the guide race orc is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.157 ± 0.023, crit=0.233 ± 0.007 per rating point (14 rating = 1%, 3.268 per %), hit=0.526 ± 0.028 per rating point (10 rating = 1%, 5.261 per %), spell_haste=not significant (-0.893 ± 0.375), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.423 ± 0.001, fire_power=0.577 ± 0.001, spirit=not significant (0.497 ± 0.381)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 35.2 spell_power points (16.36 DPS) | yes | Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Eye of Flame (3075, -2.02 DPS, sim-verified) [world_drop]; Ghostweave Hood (254135, -2.45 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Diana's Pearl Necklace (22403, +0.00 DPS) [dungeon]; Triune Amulet (7722, -0.97 DPS) [dungeon]; Star of Mystaria (12103, -1.36 DPS) [dungeon] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 28.5 spell_power points (13.23 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.38 DPS) [pvp]; Burial Shawl (18681, -2.06 DPS, sim-verified) [dungeon]; Mantle of the Timbermaw (19050, -2.86 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spritecaster Cape (11623, -2.13 DPS) [dungeon]; Crystalline Threaded Cape (20697, -3.01 DPS, sim-verified) [world]; Frostwolf Advisor's Cloak (19085, -3.96 DPS) [rep] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 47.4 spell_power points (22.02 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.41 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -8.95 DPS) [pvp]; Alanna's Embrace (13314, -12.17 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 26.7 spell_power points (12.42 DPS) | yes | Earthenweave Cuffs (254131, -6.39 DPS) [crafted]; Runecloth Cuffs (254123, -6.58 DPS) [crafted]; Netherflame Cuffs (254065, -7.27 DPS) [crafted] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 29.9 spell_power points (13.90 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Dreamweave Gloves (10019, -3.68 DPS, sim-verified) [crafted]; Earth Warder's Gloves (21318, -3.68 DPS) [quest] |
| waist | Frostwolf Cloth Belt (19090) | Frostwolf Clan [rep] | 19.6 spell_power points (9.09 DPS) | yes | Ban'thok Sash (11662, -0.27 DPS) [dungeon]; Thuzadin Sash (18740, -0.57 DPS) [dungeon]; Ghostweave Cord (254073, -2.70 DPS, sim-verified) [crafted] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | 36.0 spell_power points (16.70 DPS) | yes | General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Skyshroud Leggings (13170, -0.32 DPS) [dungeon]; Ghostweave Pants (14144, -1.75 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 27.5 spell_power points (12.76 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.70 DPS) [crafted]; Ghostweave Boots (254129, -2.33 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.46 DPS) [quest]; Don Mauricio's Band of Domination (22433, -3.63 DPS) [dungeon]; Naglering (11669, -9.97 DPS, sim-verified) [dungeon] |
| finger2 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Eye of Orgrimmar (12545, -0.70 DPS, sim-verified) [quest]; Don Mauricio's Band of Domination (22433, -2.31 DPS) [dungeon]; Advisor's Ring (19518, -2.43 DPS) [rep] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.79 DPS) [world_drop]; Ankh of Life (1713, -2.80 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -3.26 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -4.27 DPS, sim-verified) [vendor]; Uther's Strength (11302, -7.43 DPS) [world_drop]; Ankh of Life (1713, -7.45 DPS) [world_drop] |
| main_hand | Sageblade (22383) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mindfang (20214, -4.15 DPS) [rep]; High Warlord's War Staff (234549, -4.20 DPS) [pvp]; Ironbark Staff (20220, -16.07 DPS, sim-verified) [rep] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (13.47 DPS) | yes | High Warlord's Tome of Destruction (234563, -0.91 DPS) [pvp]; Tome of Fiery Arcana (19311, -3.32 DPS, sim-verified) [rep]; Spirit of Aquementas (11904, -3.71 DPS) [quest] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.22 DPS) [dungeon]; Brilliant Wand (249385, -0.64 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Frostwolf Cloth Belt; legs: Outrider's Silk Leggings; finger1: Maiden's Circle; finger2: Blessed Band of Light; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Sageblade; off_hand: Therazane's Touch; ranged: Torch of Light

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (orc, 25500000033000000-0000000000000000000-0050355103101351)

Set DPS (verified): 868.9. Weights run: 4.6s. Verify run: 24.7s. 1293 eligible items had no known source.

Race sweep: orc 868.9 ± 0.7, undead 867.6 ± 0.7, troll 858.4 ± 0.7 (dps); the guide race orc is best.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.338 ± 0.048, crit=0.694 ± 0.017 per rating point (14 rating = 1%, 9.718 per %), hit=1.128 ± 0.072 per rating point (10 rating = 1%, 11.281 per %), spell_haste=not significant (3.155 ± 1.022), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.368 ± 0.001, fire_power=0.632 ± 0.001, spirit=not significant (-0.005 ± 0.267)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crimson Felt Hat (18727) | Stratholme: Magistrate Barthilas [dungeon] | 32.7 spell_power points (22.19 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Eye of Flame (3075, -1.46 DPS) [world_drop] |
| neck | Diana's Pearl Necklace (22403) | Stratholme: Cannon Master Willey [dungeon] | 23.0 spell_power points (15.59 DPS) | yes | Star of Mystaria (12103, -5.88 DPS) [dungeon]; Archlight Talisman (15856, -7.88 DPS, sim-verified) [quest]; Dark Advisor's Pendant (18691, -8.99 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 31.1 spell_power points (21.11 DPS) | yes | Warlord's Dreadweave Mantle (231592, -0.25 DPS) [pvp]; Champion's Dreadweave Spaulders (227092, -3.39 DPS) [pvp]; Burial Shawl (18681, -4.64 DPS, sim-verified) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Crystalline Threaded Cape (20697, -6.22 DPS, sim-verified) [world]; Spritecaster Cape (11623, -9.47 DPS) [dungeon]; Fel Cape (279269, -10.00 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 49.0 spell_power points (33.27 DPS) | yes | Warlord's Dreadweave Robe (231591, -6.06 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -11.72 DPS) [pvp]; Alanna's Embrace (13314, -17.95 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 24.7 spell_power points (16.76 DPS) | yes | Runecloth Cuffs (254123, -7.00 DPS) [crafted]; Netherflame Cuffs (254065, -9.58 DPS) [crafted]; Arena Wristguards (18709, -10.17 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.8 spell_power points (21.58 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Hands of Power (13253, -2.57 DPS) [dungeon]; Blood Guard's Dreadweave Handwraps (227099, -6.42 DPS) [pvp] |
| waist | Defiler's Cloth Girdle (20163) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ban'thok Sash (11662, +0.00 DPS) [dungeon]; Frostwolf Cloth Belt (19090, -2.96 DPS) [rep]; Clutch of Andros (13956, -5.46 DPS) [dungeon] |
| legs | Outrider's Silk Leggings (22747) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skyshroud Leggings (13170, +0.00 DPS) [dungeon]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Legionnaire's Dreadweave Legguards (227097, -1.38 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.1 spell_power points (16.32 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, -0.04 DPS) [crafted]; Dragonrider Boots (18102, -0.44 DPS) [dungeon] |
| finger1 | Blessed Band of Light (272407) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.38 DPS) [dungeon]; Maiden's Circle (13001, -0.62 DPS) [world_drop]; Naglering (11669, -14.97 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Don Mauricio's Band of Domination (22433, -0.19 DPS) [dungeon]; Eye of Orgrimmar (12545, -0.43 DPS) [quest]; Maiden's Circle (13001, -0.43 DPS) [world_drop] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+11.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, -1.25 DPS) [crafted]; Rune of the Guard Captain (19120, -2.78 DPS) [quest]; Uther's Strength (11302, -4.07 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Serenity Field (272439, -5.66 DPS, sim-verified) [vendor]; Frozen Heart of the Mountain (249469, -8.04 DPS) [crafted]; Rune of the Guard Captain (19120, -9.57 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, -2.91 DPS) [pvp]; Staff of Jordan (873, -19.08 DPS, sim-verified) [world_drop]; Whiteout Staff (19101, -22.45 DPS) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bonecreeper Stylus (13938, +0.00 DPS) [dungeon]; Ritssyn's Wand of Bad Mojo (22408, -0.25 DPS) [dungeon]; Greater Eternal Wand (249237, -1.61 DPS) [crafted] |

**New at 60:** head: Crimson Felt Hat; neck: Diana's Pearl Necklace; shoulder: Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of the Void; wrist: Dryad's Wrist Bindings; hands: Gloves of Spell Mastery; waist: Defiler's Cloth Girdle; legs: Outrider's Silk Leggings; feet: Omnicast Boots; finger1: Blessed Band of Light; finger2: Don Julio's Band; trinket1: Burst of Knowledge; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1293, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

