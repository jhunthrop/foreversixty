# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 50.9. Weights run: 1.6s. Verify run: 1.5s. 186 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.893 ± 0.024, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.410 per %), hit=0.075 ± 0.004 per rating point (10 rating = 1%, 0.755 per %), spell_haste=not significant (0.047 ± 0.070), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.787 ± 0.001, fire_power=0.213 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.7 spell_power points (2.00 DPS) | yes | Pristine Circlet (253949, -0.81 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.97 DPS) [crafted]; Flame Circlet (253953, -1.46 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.0 spell_power points (3.01 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.22 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -2.09 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (50.9 DPS) | yes | Sanguine Cape (14376, -0.10 DPS) [world_drop]; Black Whelp Cloak (7283, -0.23 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.54 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 15.5 spell_power points (3.58 DPS) | yes | Filigreed Shadow Gown (253907, -1.07 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.39 DPS) [crafted]; Gray Woolen Robe (2585, -1.62 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 4.5 spell_power points (1.03 DPS) | yes | Bright Bracers (3647, -0.52 DPS, sim-verified) [world_drop]; Mystic's Bracelets (14366, -0.62 DPS) [world_drop]; Repurposed Hair Band (281256, -0.62 DPS) [quest] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 8.2 spell_power points (1.89 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.35 DPS) [crafted]; Gnoll Casting Gloves (892, -0.50 DPS) [world] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 9.1 spell_power points (2.10 DPS) | yes | Keller's Girdle (2911, -0.45 DPS) [world_drop]; Pristine Sash (253925, -0.46 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 14.0 spell_power points (3.24 DPS) | yes | Filigreed Pristine Leggings (253937, -0.81 DPS, sim-verified) [crafted]; Filigreed Flame Leggings (253941, -1.46 DPS) [crafted]; Silk-threaded Trousers (1929, -1.62 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.6 spell_power points (2.44 DPS) | yes | Shadow Boots (253895, -0.66 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.13 DPS) [crafted]; Red Woolen Boots (4313, -1.52 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.8 spell_power points (1.57 DPS) | yes | Volcanic Rock Ring (12053, -0.95 DPS) [world_drop]; Furlfeather Ring (286745, -1.34 DPS) [world]; The 1 Ring (8350, -1.36 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 spell_power points (1.24 DPS) | yes | Volcanic Rock Ring (12053, -0.25 DPS, sim-verified) [world_drop]; Furlfeather Ring (286745, -1.01 DPS) [world]; The 1 Ring (8350, -1.03 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Witching Stave (1484) | Shadowfang Keep [dungeon] | 25.7 spell_power points (5.94 DPS) | yes | Lesser Staff of the Spire (1300, -4.70 DPS) [world]; Staff of Westfall (2042, -4.91 DPS) [quest]; Twisted Chanter's Staff (890, -5.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 96.5 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.75 DPS) [world_drop]; Deepblaze (279896, -3.78 DPS) [quest]; Sizzle Stick (8071, -4.20 DPS) [quest] |

**New at 20:** head: Shadow Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Witching Stave; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 80.4. Weights run: 1.6s. Verify run: 1.2s. 323 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.066 ± 0.004, crit=0.043 ± 0.001 per rating point (14 rating = 1%, 0.598 per %), hit=0.105 ± 0.005 per rating point (10 rating = 1%, 1.045 per %), spell_haste=-0.713 ± 0.078, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.799 ± 0.001, fire_power=0.201 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 11.2 spell_power points (2.89 DPS) | yes | Holy Shroud (2721, -0.05 DPS) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.62 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.4 spell_power points (0.10 DPS) | yes | Crystal Starfire Medallion (5003, -0.03 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.03 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.6 spell_power points (2.48 DPS) | yes | Moonlit Amice (11884, -0.67 DPS) [quest]; Death Speaker Mantle (6685, -0.74 DPS) [dungeon]; Invoker's Mantle (215365, -1.00 DPS, sim-verified) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.29 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.26 DPS) [crafted]; Black Whelp Cloak (7283, -0.52 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.35 DPS) | yes | Necrology Robes (2292, -0.38 DPS) [dungeon]; Shadow Gown (253967, -0.56 DPS) [crafted]; Green Silk Armor (7065, -0.81 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.32 DPS) | yes | Windsong Bangles (263336, -2.06 DPS) [quest]; Glowing Magical Bracelets (13106, -2.19 DPS) [world_drop]; Phoenix Bindings (210781, -2.21 DPS, sim-verified) [crafted] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.7 spell_power points (1.97 DPS) | yes | Hands of Darkness (7047, -0.12 DPS) [crafted]; Serpent Gloves (5970, -0.17 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.43 DPS) [world] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 9.2 spell_power points (2.37 DPS) | yes | Invoker's Cord (215366, -0.48 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.63 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.71 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.10 DPS) | yes | Shadow Leggings (253993, -0.09 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.73 DPS) [crafted]; Necromancer Leggings (2277, -0.85 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.8 spell_power points (2.80 DPS) | yes | Spidersilk Boots (4320, -0.92 DPS) [crafted]; Glinteye Slippers (273024, -0.99 DPS) [dungeon]; Gilded Slippers (254001, -1.02 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.3 spell_power points (1.63 DPS) | yes | Minor Channeling Ring (1449, -0.31 DPS) [quest]; Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Runemender's Seal (273031, -0.86 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.55 DPS) | yes | Minor Channeling Ring (1449, -0.47 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Runemender's Seal (273031, -0.77 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | 34.0 spell_power points (8.77 DPS) | yes | Staff of the Shade (2549, -0.95 DPS, sim-verified) [dungeon]; Witching Stave (1484, -3.12 DPS) [dungeon]; Glimmering Staff (249392, -8.59 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 127.6 spell_power points (32.93 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Gyromatic Icemaker (9489, -2.57 DPS) [dungeon]; Greater Mystic Wand (217287, -2.64 DPS) [crafted] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Staff of Soran'ruk; ranged: Starfaller

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 141.7. Weights run: 1.6s. Verify run: 1.3s. 420 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.192 ± 0.007, crit=0.064 ± 0.001 per rating point (14 rating = 1%, 0.901 per %), hit=0.166 ± 0.007 per rating point (10 rating = 1%, 1.658 per %), spell_haste=0.487 ± 0.073, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.820 ± 0.001, fire_power=0.180 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.20 DPS) | yes | Augural Shroud (2620, -2.39 DPS) [world]; Living Cowl (5608, -2.49 DPS, sim-verified) [world]; Filigreed Shadow Circlet (253981, -2.81 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.5 spell_power points (2.52 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -2.12 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.18 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (3.17 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Inquisitor's Shawl (19507, -0.36 DPS) [dungeon]; Berylline Pads (4197, -0.53 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.07 DPS) | yes | Guardian Cloak (5965, -0.01 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.59 DPS) [quest]; Long Silken Cloak (4326, -0.68 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (6.84 DPS) | yes | Dreamweave Vest (10021, -1.01 DPS) [crafted]; Elemental Raiment (9434, -1.32 DPS, sim-verified) [world_drop]; Shadoweave Robe (10004, -1.75 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (2.66 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.59 DPS) [quest]; Earthen Silk Cuffs (254019, -1.48 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.54 DPS) | yes | Shadoweave Gloves (10023, -1.14 DPS) [crafted]; Black Mageweave Gloves (10003, -1.20 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.73 DPS) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | 13.0 spell_power points (3.84 DPS) | yes | Star Belt (4329, -0.00 DPS) [crafted]; Deathmage Sash (10771, -0.93 DPS) [dungeon]; Belt of Arugal (6392, -1.02 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.2 spell_power points (5.09 DPS) | yes | Red Mageweave Pants (10009, -0.27 DPS) [crafted]; Shadow Leggings (253993, -1.30 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.54 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.09 DPS) | yes | Black Slippers (254007, -3.85 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -4.63 DPS) [crafted]; Spidersilk Boots (4320, -4.80 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.4 spell_power points (3.67 DPS) | yes | Ring of Forlorn Spirits (2043, -1.31 DPS) [quest]; Reedknot Ring (9622, -1.61 DPS) [quest]; Snake Hoop (6750, -1.62 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (3.24 DPS) | yes | Ring of Forlorn Spirits (2043, -0.87 DPS) [quest]; Reedknot Ring (9622, -1.17 DPS) [quest]; Snake Hoop (6750, -1.18 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (141.7 DPS) | yes | Staff of Soran'ruk (15109, -1.38 DPS, sim-verified) [quest]; Staff of the Shade (2549, -1.60 DPS) [dungeon]; Witching Stave (1484, -3.74 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 134.4 spell_power points (39.71 DPS) | yes | Earthen Rod (9381, +0.00 DPS, sim-verified) [dungeon]; Twisted Nether Wand (249144, -3.94 DPS) [crafted]; Umbral Wand (5216, -4.05 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Black Cord; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 420, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 207.0. Weights run: 1.6s. Verify run: 1.5s. 531 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.243 ± 0.011, crit=0.100 ± 0.002 per rating point (14 rating = 1%, 1.400 per %), hit=0.282 ± 0.012 per rating point (10 rating = 1%, 2.821 per %), spell_haste=1.187 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.768 ± 0.001, fire_power=0.232 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.78 DPS) | yes | Dreamweave Circlet (10041, -1.16 DPS) [crafted]; Red Mageweave Headband (10033, -1.68 DPS, sim-verified) [crafted]; Spellpower Goggles Xtreme (10502, -1.95 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.9 spell_power points (2.91 DPS) | yes | Triune Amulet (7722, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.48 DPS) [quest]; Horizon Choker (13085, -1.80 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.4 spell_power points (5.65 DPS) | yes | Netherpearl Shoulders (254055, -0.77 DPS) [crafted]; Shadoweave Shoulders (10028, -1.40 DPS) [crafted]; Black Mageweave Shoulders (10027, -1.69 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (5.03 DPS) | yes | Runecloth Cloak (13860, -2.05 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.10 DPS) [dungeon]; Long Silken Cloak (4326, -2.68 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.9 spell_power points (7.76 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -0.93 DPS) [world_drop]; Dreamweave Vest (10021, -1.19 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 11.7 spell_power points (3.80 DPS) | yes | Arcane Runed Bracers (4744, -0.65 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.87 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.97 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (6.17 DPS) | yes | Black Gloves (254101, -0.71 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.23 DPS) [vendor]; Black Mageweave Gloves (10003, -1.29 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 16.8 spell_power points (5.46 DPS) | yes | Satyrmane Sash (17755, -0.12 DPS) [dungeon]; Ghostweave Cord (254073, -0.91 DPS) [crafted]; Star Belt (4329, -1.23 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 22.9 spell_power points (7.44 DPS) | yes | Wizardweave Leggings (14132, -1.02 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -1.94 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.14 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.80 DPS) | yes | Black Sandals (254113, -2.83 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -3.52 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -3.57 DPS) [vendor] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.23 DPS) | yes | Philanthropist's Ring (281635, -0.58 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.63 DPS) [quest]; Snake Hoop (6750, -1.88 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 spell_power points (3.89 DPS) | yes | Philanthropist's Ring (281635, -0.25 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.29 DPS) [quest]; Snake Hoop (6750, -1.55 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (207.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (207.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (207.0 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS, sim-verified) [world]; Inventor's Focal Sword (17719, -0.95 DPS) [dungeon]; Spellshifter Rod (9527, -1.54 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 164.5 spell_power points (53.48 DPS) | yes | Noxious Shooter (17745, -1.85 DPS) [dungeon]; Wand of Allistarj (13065, -5.06 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.13 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger1: Band of the Unicorn; finger2: Underworld Band; trinket1: Frozen Heart of the Mountain; trinket2: Mark of the Chosen; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 531, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 325.5. Weights run: 1.6s. Verify run: 2.3s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.248 ± 0.020, crit=0.137 ± 0.005 per rating point (14 rating = 1%, 1.919 per %), hit=0.406 ± 0.020 per rating point (10 rating = 1%, 4.058 per %), spell_haste=not significant (0.278 ± 0.315), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.768 ± 0.001, fire_power=0.232 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.06 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Felcloth Hood (14111, -0.55 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.5 spell_power points (5.39 DPS) | yes | Glowing Eye of Mordresh (10769, -2.05 DPS) [dungeon]; Triune Amulet (7722, -2.41 DPS) [dungeon]; Mindburst Medallion (11196, -2.60 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.31 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.07 DPS) [crafted]; Felcloth Shoulders (14112, -1.87 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (5.77 DPS) | yes | Fel Cape (279269, -0.76 DPS) [crafted]; Runecloth Cloak (13860, -1.68 DPS) [crafted]; Nightfall Drape (12465, -2.42 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (17.97 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -3.83 DPS) [pvp]; Robe of Winter Night (14136, -5.42 DPS) [crafted]; Knight-Captain's Dreadweave Tunic (227096, -6.81 DPS) [pvp] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.5 spell_power points (5.02 DPS) | yes | Netherpearl Cuffs (254067, -0.66 DPS) [crafted]; Arcane Runed Bracers (4744, -1.67 DPS) [quest]; Spidertank Oilrag (9448, -1.67 DPS) [dungeon] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.20 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.00 DPS) [pvp]; Dreamweave Gloves (10019, -1.12 DPS) [crafted] |
| waist | Wisdom of the Timbermaw (19047) | Tailoring [crafted] | sim-verified (325.5 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Black Waistcord (254087, +0.00 DPS) [crafted]; Ghostweave Cord (254073, -0.17 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.69 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Felcloth Pants (14107, -1.14 DPS) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (325.5 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, +0.00 DPS) [crafted]; Knight-Lieutenant's Dreadweave Boots (17562, -0.10 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (7.35 DPS) | yes | Underworld Band (1980, -2.88 DPS) [world_drop]; Philanthropist's Ring (281635, -3.16 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.37 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.84 DPS) | yes | Underworld Band (1980, -0.38 DPS) [world_drop]; Philanthropist's Ring (281635, -0.66 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.86 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (325.5 DPS) | yes | - |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -1.08 DPS) [world]; Inventor's Focal Sword (17719, -1.11 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 212.1 spell_power points (79.00 DPS) | yes | Brilliant Wand (249385, -15.42 DPS) [crafted]; Stormrager (16997, -15.85 DPS) [quest]; Torch of Austen (13004, -22.21 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Wisdom of the Timbermaw; legs: Ghostweave Pants; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 05000000000000000-2035113011220001351-0550005000000000)

Set DPS (verified): 641.3. Weights run: 1.8s. Verify run: 2.6s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.267 ± 0.033, crit=0.383 ± 0.013 per rating point (14 rating = 1%, 5.357 per %), hit=0.617 ± 0.044 per rating point (10 rating = 1%, 6.169 per %), spell_haste=3.707 ± 0.592, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.780 ± 0.001, fire_power=0.220 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (14.92 DPS) | yes | Lieutenant Commander's Dreadweave Cowl (227093, +0.00 DPS) [pvp]; Field Marshal's Coronal (231584, +0.00 DPS) [pvp]; Felcloth Hood (14111, -0.51 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.7 spell_power points (8.11 DPS) | yes | Triune Amulet (7722, -3.69 DPS) [dungeon]; Mindburst Medallion (11196, -3.91 DPS) [quest]; Glowing Eye of Mordresh (10769, -7.54 DPS, sim-verified) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.8 spell_power points (14.27 DPS) | yes | Field Marshal's Dreadweave Shoulders (231583, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.46 DPS) [crafted]; Lieutenant Commander's Dreadweave Spaulders (227091, -2.76 DPS) [pvp] |
| back | Fel Cape (279269) | Tailoring [crafted] | 17.1 spell_power points (9.43 DPS) | yes | Spritecaster Cape (11623, -0.81 DPS) [dungeon]; Runecloth Cloak (13860, -3.27 DPS) [crafted]; Nightfall Drape (12465, -4.46 DPS) [dungeon] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.4 spell_power points (26.74 DPS) | yes | Field Marshal's Dreadweave Robe (231582, -5.52 DPS) [pvp]; Robe of Winter Night (14136, -7.05 DPS, sim-verified) [crafted]; Knight-Captain's Dreadweave Tunic (227096, -9.98 DPS) [pvp] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.7 spell_power points (7.55 DPS) | yes | Arcane Runed Bracers (4744, -2.58 DPS) [quest]; Spidertank Oilrag (9448, -2.58 DPS) [dungeon]; Netherpearl Cuffs (254067, -4.95 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.4 spell_power points (12.37 DPS) | yes | Marshal's Dreadweave Gloves (231586, +0.00 DPS) [pvp]; Knight-Lieutenant's Dreadweave Handwraps (227100, -0.18 DPS) [pvp]; Earth Warder's Gloves (21318, -5.56 DPS, sim-verified) [quest] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 17.2 spell_power points (9.52 DPS) | yes | Satyrmane Sash (17755, -0.31 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.38 DPS) [crafted]; Ghostweave Cord (254073, -1.79 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (14.37 DPS) | yes | Knight-Captain's Dreadweave Legguards (227095, +0.00 DPS) [pvp]; Marshal's Dreadweave Leggings (231587, +0.00 DPS) [pvp]; Felcloth Pants (14107, -1.38 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.26 DPS) | yes | Marshal's Dreadweave Boots (231585, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.44 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Boots (17562, -0.55 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.9 spell_power points (10.98 DPS) | yes | Underworld Band (1980, -4.20 DPS) [world_drop]; Philanthropist's Ring (281635, -4.72 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.56 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.18 DPS) | yes | Underworld Band (1980, -0.41 DPS) [world_drop]; Philanthropist's Ring (281635, -0.92 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.76 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (641.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (641.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (641.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Inventor's Focal Sword (17719, -1.77 DPS) [dungeon]; Spire of Hakkar (10844, -3.63 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 144.9 spell_power points (80.08 DPS) | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -16.65 DPS) [quest]; Greater Eternal Wand (249237, -21.77 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-2332100000000000000-0000000000000000)

Set DPS (verified): 49.6. Weights run: 1.6s. Verify run: 1.5s. 180 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.893 ± 0.024, crit=0.029 ± 0.001 per rating point (14 rating = 1%, 0.410 per %), hit=0.075 ± 0.004 per rating point (10 rating = 1%, 0.755 per %), spell_haste=not significant (0.047 ± 0.070), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.787 ± 0.001, fire_power=0.213 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 8.7 spell_power points (2.00 DPS) | yes | Pristine Circlet (253949, -0.79 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.97 DPS) [crafted]; Flame Circlet (253953, -1.46 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.0 spell_power points (3.01 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.03 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -2.09 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.10 DPS) [world_drop]; Black Whelp Cloak (7283, -0.23 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.52 DPS, sim-verified) [crafted] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 15.5 spell_power points (3.58 DPS) | yes | Filigreed Shadow Gown (253907, -0.82 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -1.39 DPS) [crafted]; Gray Woolen Robe (2585, -1.62 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 4.5 spell_power points (1.03 DPS) | yes | Bright Bracers (3647, -0.21 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.34 DPS) [quest]; Owlbeard Bracers (16981, -0.39 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Gloves (253913, -0.07 DPS) [crafted]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Shadow Gloves (253919, -0.49 DPS, sim-verified) [crafted] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 9.1 spell_power points (2.10 DPS) | yes | Keller's Girdle (2911, -0.45 DPS) [world_drop]; Pristine Sash (253925, -0.46 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.55 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 14.0 spell_power points (3.24 DPS) | yes | Filigreed Pristine Leggings (253937, -0.79 DPS, sim-verified) [crafted]; Filigreed Flame Leggings (253941, -1.46 DPS) [crafted]; Silk-threaded Trousers (1929, -1.62 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.6 spell_power points (2.44 DPS) | yes | Shadow Boots (253895, -0.49 DPS, sim-verified) [crafted]; Pristine Boots (253889, -1.13 DPS) [crafted]; Red Woolen Boots (4313, -1.52 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.4 spell_power points (1.24 DPS) | yes | Volcanic Rock Ring (12053, -0.62 DPS) [world_drop]; Furlfeather Ring (286745, -1.01 DPS) [world]; The 1 Ring (8350, -1.03 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 4.5 spell_power points (1.03 DPS) | yes | Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; Furlfeather Ring (286745, -0.80 DPS) [world]; The 1 Ring (8350, -0.83 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Witching Stave (1484) | Shadowfang Keep [dungeon] | 25.7 spell_power points (5.94 DPS) | yes | Crescent Staff (6505, -4.49 DPS) [quest]; Lesser Staff of the Spire (1300, -4.70 DPS) [world]; Twisted Chanter's Staff (890, -5.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 96.5 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.75 DPS) [world_drop]; Sizzle Stick (8071, -4.20 DPS) [quest]; Greater Magic Wand (11288, -4.35 DPS) [crafted] |

**New at 20:** head: Shadow Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Witching Stave; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 180, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (troll, 00000000000000000-2332113101220000000-0000000000000000)

Set DPS (verified): 80.4. Weights run: 1.6s. Verify run: 1.2s. 314 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.066 ± 0.004, crit=0.043 ± 0.001 per rating point (14 rating = 1%, 0.598 per %), hit=0.105 ± 0.005 per rating point (10 rating = 1%, 1.045 per %), spell_haste=-0.713 ± 0.078, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.799 ± 0.001, fire_power=0.201 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Shadow Circlet (253981) | Tailoring [crafted] | 11.2 spell_power points (2.89 DPS) | yes | Holy Shroud (2721, -0.05 DPS) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted]; Shadow Circlet (253955, -0.62 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 0.4 spell_power points (0.10 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.03 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.6 spell_power points (2.48 DPS) | yes | Death Speaker Mantle (6685, -0.74 DPS) [dungeon]; Invoker's Mantle (215365, -0.90 DPS, sim-verified) [crafted]; Magician's Mantle (12998, -1.03 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.29 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.26 DPS) [crafted]; Black Whelp Cloak (7283, -0.52 DPS) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps [dungeon] | 13.0 spell_power points (3.35 DPS) | yes | Necrology Robes (2292, -0.38 DPS) [dungeon]; Shadow Gown (253967, -0.56 DPS) [crafted]; Green Silk Armor (7065, -0.81 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.32 DPS) | yes | Tabitha's Cuffs (251486, -1.53 DPS, sim-verified) [quest]; Phoenix Bindings (210781, -1.65 DPS) [crafted]; Owlbeard Bracers (16981, -2.03 DPS) [quest] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 7.7 spell_power points (1.97 DPS) | yes | Hands of Darkness (7047, -0.12 DPS) [crafted]; Serpent Gloves (5970, -0.17 DPS) [dungeon]; Jutebraid Gloves (10654, -0.34 DPS) [quest] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.84 DPS) | yes | Belt of Arugal (6392, -0.46 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.77 DPS) [dungeon]; Invoker's Cord (215366, -0.95 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 spell_power points (3.10 DPS) | yes | Shadow Leggings (253993, -0.09 DPS) [crafted]; Filigreed Shadow Leggings (253943, -0.73 DPS) [crafted]; Necromancer Leggings (2277, -0.85 DPS) [world_drop] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 10.8 spell_power points (2.80 DPS) | yes | Spidersilk Boots (4320, -0.92 DPS) [crafted]; Glinteye Slippers (273024, -0.99 DPS) [dungeon]; Gilded Slippers (254001, -1.01 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 6.3 spell_power points (1.63 DPS) | yes | Solid Gold Band (279845, -0.60 DPS) [quest]; Electrocutioner Lagnut (9447, -0.86 DPS) [dungeon]; Runemender's Seal (273031, -0.86 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.55 DPS) | yes | Solid Gold Band (279845, -0.37 DPS, sim-verified) [quest]; Electrocutioner Lagnut (9447, -0.77 DPS) [dungeon]; Runemender's Seal (273031, -0.77 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Soran'ruk (15109) | The Orb of Soran'ruk [quest] | 34.0 spell_power points (8.77 DPS) | yes | Staff of the Shade (2549, -0.95 DPS, sim-verified) [dungeon]; Witching Stave (1484, -3.12 DPS) [dungeon]; Glimmering Staff (249392, -8.59 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 127.6 spell_power points (32.93 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Gyromatic Icemaker (9489, -2.57 DPS) [dungeon]; Greater Mystic Wand (217287, -2.64 DPS) [crafted] |

**New at 30:** head: Filigreed Shadow Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Black Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Staff of Soran'ruk; ranged: Starfaller

No-known-source sample (15 of 314, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 10000000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 142.2. Weights run: 1.6s. Verify run: 1.3s. 411 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.192 ± 0.007, crit=0.064 ± 0.001 per rating point (14 rating = 1%, 0.901 per %), hit=0.166 ± 0.007 per rating point (10 rating = 1%, 1.658 per %), spell_haste=0.487 ± 0.073, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.820 ± 0.001, fire_power=0.180 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (6.20 DPS) | yes | Augural Shroud (2620, -2.39 DPS) [world]; Living Cowl (5608, -2.50 DPS, sim-verified) [world]; Filigreed Shadow Circlet (253981, -2.81 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.5 spell_power points (2.52 DPS) | yes | Triune Amulet (7722, -0.16 DPS) [dungeon]; Necklace of Calisea (1714, -2.12 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -2.18 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.7 spell_power points (3.17 DPS) | yes | Green Silken Shoulders (7057, -0.18 DPS) [crafted]; Inquisitor's Shawl (19507, -0.36 DPS) [dungeon]; Berylline Pads (4197, -0.53 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 spell_power points (2.07 DPS) | yes | Guardian Cloak (5965, -0.01 DPS) [crafted]; Long Silken Cloak (4326, -0.55 DPS, sim-verified) [crafted]; Hillman's Cloak (3719, -0.59 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.1 spell_power points (6.84 DPS) | yes | Dreamweave Vest (10021, -1.01 DPS) [crafted]; Elemental Raiment (9434, -1.65 DPS, sim-verified) [world_drop]; Shadoweave Robe (10004, -1.75 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.66 DPS) | yes | Radiant Silver Bracers (4545, -1.02 DPS) [quest]; Condor Bracers (15864, -1.10 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -1.48 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 spell_power points (5.54 DPS) | yes | Shadoweave Gloves (10023, -1.14 DPS) [crafted]; Red Mageweave Gloves (10018, -1.73 DPS) [crafted]; Black Mageweave Gloves (10003, -2.10 DPS, sim-verified) [crafted] |
| waist | Black Cord (254043) | Tailoring [crafted] | 13.0 spell_power points (3.84 DPS) | yes | Star Belt (4329, -0.58 DPS, sim-verified) [crafted]; Warsong Sash (16975, -0.60 DPS) [quest]; Deathmage Sash (10771, -0.93 DPS) [dungeon] |
| legs | Shadoweave Pants (10002) | Tailoring [crafted] | 17.2 spell_power points (5.09 DPS) | yes | Red Mageweave Pants (10009, -0.79 DPS, sim-verified) [crafted]; Shadow Leggings (253993, -1.30 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.54 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.09 DPS) | yes | Black Slippers (254007, -4.13 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -4.63 DPS) [crafted]; Spidersilk Boots (4320, -4.80 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.4 spell_power points (3.67 DPS) | yes | Reedknot Ring (9622, -1.61 DPS) [quest]; Snake Hoop (6750, -1.62 DPS) [quest]; Ring of Power Regulation (273030, -1.90 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.0 spell_power points (3.24 DPS) | yes | Snake Hoop (6750, -1.18 DPS) [quest]; Reedknot Ring (9622, -1.38 DPS, sim-verified) [quest]; Ring of Power Regulation (273030, -1.46 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (142.2 DPS) | yes | Staff of Soran'ruk (15109, -1.34 DPS, sim-verified) [quest]; Staff of the Shade (2549, -1.60 DPS) [dungeon]; Witching Stave (1484, -3.74 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 134.4 spell_power points (39.71 DPS) | yes | Earthen Rod (9381, -3.24 DPS) [dungeon]; Twisted Nether Wand (249144, -3.94 DPS) [crafted]; Umbral Wand (5216, -4.05 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Black Cord; legs: Shadoweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; finger2: Philanthropist's Ring; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25400000000000000-2332113101220001350-0000000000000000)

Set DPS (verified): 208.6. Weights run: 1.6s. Verify run: 1.6s. 522 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.243 ± 0.011, crit=0.100 ± 0.002 per rating point (14 rating = 1%, 1.400 per %), hit=0.282 ± 0.012 per rating point (10 rating = 1%, 2.821 per %), spell_haste=1.187 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.768 ± 0.001, fire_power=0.232 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (8.78 DPS) | yes | Dreamweave Circlet (10041, -1.16 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -1.95 DPS) [crafted]; Red Mageweave Headband (10033, -2.59 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 8.9 spell_power points (2.91 DPS) | yes | Triune Amulet (7722, -0.31 DPS) [dungeon]; Mindburst Medallion (11196, -0.48 DPS) [quest]; Horizon Choker (13085, -1.80 DPS) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 17.4 spell_power points (5.65 DPS) | yes | Netherpearl Shoulders (254055, -1.17 DPS, sim-verified) [crafted]; Shadoweave Shoulders (10028, -1.40 DPS) [crafted]; Black Mageweave Shoulders (10027, -1.69 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (5.03 DPS) | yes | Deep Woodlands Cloak (19121, -1.42 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.47 DPS) [crafted]; Nightfall Drape (12465, -2.10 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 23.9 spell_power points (7.76 DPS) | yes | Robe of the Magi (1716, -0.13 DPS) [world_drop]; Elemental Raiment (9434, -0.93 DPS) [world_drop]; Dreamweave Vest (10021, -1.19 DPS) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 11.7 spell_power points (3.80 DPS) | yes | Nethergeld Cuffs (254061, -0.97 DPS) [crafted]; Spidertank Oilrag (9448, -1.01 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -1.46 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 spell_power points (6.17 DPS) | yes | Black Gloves (254101, -1.21 DPS, sim-verified) [crafted]; First Sergeant's Dreadweave Gloves (220908, -1.23 DPS) [vendor]; Black Mageweave Gloves (10003, -1.29 DPS) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 16.8 spell_power points (5.46 DPS) | yes | Satyrmane Sash (17755, -0.12 DPS) [dungeon]; Ghostweave Cord (254073, -0.91 DPS) [crafted]; Star Belt (4329, -1.23 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 22.9 spell_power points (7.44 DPS) | yes | Wizardweave Leggings (14132, -0.96 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -1.94 DPS) [crafted]; Stone Guard's Dreadweave Leggings (220906, -2.14 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (7.80 DPS) | yes | Gilded Sandals (254107, -3.52 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.57 DPS) [vendor]; Black Sandals (254113, -3.81 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.23 DPS) | yes | Philanthropist's Ring (281635, -0.58 DPS) [quest]; Snake Hoop (6750, -1.88 DPS) [quest]; Reedknot Ring (9622, -1.95 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 spell_power points (3.89 DPS) | yes | Philanthropist's Ring (281635, -0.25 DPS) [quest]; Snake Hoop (6750, -1.55 DPS) [quest]; Reedknot Ring (9622, -1.62 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (208.6 DPS) | yes | Uther's Strength (11302, -0.86 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (208.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (208.6 DPS) | yes | Spire of Hakkar (10844, -0.75 DPS) [world]; Inventor's Focal Sword (17719, -0.95 DPS) [dungeon]; Spellshifter Rod (9527, -1.54 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 164.5 spell_power points (53.48 DPS) | yes | Noxious Shooter (17745, -1.85 DPS) [dungeon]; Wand of Allistarj (13065, -5.06 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.13 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; waist: Black Waistcord; legs: Felcloth Pants; finger1: Band of the Unicorn; finger2: Underworld Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 522, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25522000000000000-2332113101220001350-0050000000000000)

Set DPS (verified): 327.2. Weights run: 1.6s. Verify run: 2.5s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.248 ± 0.020, crit=0.137 ± 0.005 per rating point (14 rating = 1%, 1.919 per %), hit=0.406 ± 0.020 per rating point (10 rating = 1%, 4.058 per %), spell_haste=not significant (0.278 ± 0.315), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.768 ± 0.001, fire_power=0.232 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.06 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Felcloth Hood (14111, -1.66 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.5 spell_power points (5.39 DPS) | yes | Glowing Eye of Mordresh (10769, -1.90 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -2.41 DPS) [dungeon]; Mindburst Medallion (11196, -2.60 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.31 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Felcloth Shoulders (14112, -1.87 DPS) [crafted]; Mantle of the Timbermaw (19050, -2.58 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (5.77 DPS) | yes | Fel Cape (279269, -0.76 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.21 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.68 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.2 spell_power points (17.97 DPS) | yes | Warlord's Dreadweave Robe (231591, -3.83 DPS) [pvp]; Legionnaire's Dreadweave Tunic (227094, -6.81 DPS) [pvp]; Robe of Winter Night (14136, -6.96 DPS, sim-verified) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.5 spell_power points (5.02 DPS) | yes | Netherpearl Cuffs (254067, -0.66 DPS) [crafted]; Spidertank Oilrag (9448, -1.67 DPS) [dungeon]; Nethergeld Cuffs (254061, -1.77 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.20 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.00 DPS) [pvp]; Dreamweave Gloves (10019, -1.24 DPS, sim-verified) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 16.8 spell_power points (6.27 DPS) | yes | Wisdom of the Timbermaw (19047, -0.88 DPS) [crafted]; Ghostweave Cord (254073, -1.05 DPS) [crafted]; Satyrmane Sash (17755, -1.55 DPS, sim-verified) [dungeon] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.69 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Felcloth Pants (14107, -3.71 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.94 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Venomspew Footpads (275606, -0.75 DPS) [crafted]; Omnicast Boots (11822, -3.07 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.7 spell_power points (7.35 DPS) | yes | Underworld Band (1980, -2.88 DPS) [world_drop]; Philanthropist's Ring (281635, -3.16 DPS) [quest]; Snake Hoop (6750, -4.65 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.84 DPS) | yes | Philanthropist's Ring (281635, -0.66 DPS) [quest]; Underworld Band (1980, -1.62 DPS, sim-verified) [world_drop]; Snake Hoop (6750, -2.15 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (327.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (327.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (327.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -1.08 DPS) [world]; Inventor's Focal Sword (17719, -1.11 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 212.1 spell_power points (79.00 DPS) | yes | Brilliant Wand (249385, -15.42 DPS) [crafted]; Stormrager (16997, -15.85 DPS) [quest]; Torch of Austen (13004, -22.21 DPS) [world_drop] |

**New at 60:** neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60, raid preset (troll, 05000000000000000-2035113011220001351-0550005000000000)

Set DPS (verified): 645.8. Weights run: 1.8s. Verify run: 2.7s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.267 ± 0.033, crit=0.383 ± 0.013 per rating point (14 rating = 1%, 5.357 per %), hit=0.617 ± 0.044 per rating point (10 rating = 1%, 6.169 per %), spell_haste=3.707 ± 0.592, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.780 ± 0.001, fire_power=0.220 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (14.92 DPS) | yes | Champion's Dreadweave Cowl (227090, +0.00 DPS) [pvp]; Warlord's Dreadweave Hood (231590, +0.00 DPS) [pvp]; Felcloth Hood (14111, -0.51 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.7 spell_power points (8.11 DPS) | yes | Glowing Eye of Mordresh (10769, -3.06 DPS) [dungeon]; Triune Amulet (7722, -3.69 DPS) [dungeon]; Mindburst Medallion (11196, -3.91 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.8 spell_power points (14.27 DPS) | yes | Warlord's Dreadweave Mantle (231592, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.46 DPS) [crafted]; Champion's Dreadweave Spaulders (227092, -2.76 DPS) [pvp] |
| back | Fel Cape (279269) | Tailoring [crafted] | 17.1 spell_power points (9.43 DPS) | yes | Spritecaster Cape (11623, -0.81 DPS) [dungeon]; Deep Woodlands Cloak (19121, -1.47 DPS) [quest]; Runecloth Cloak (13860, -3.27 DPS) [crafted] |
| chest | Robe of the Void (14153) | Tailoring [crafted] | 48.4 spell_power points (26.74 DPS) | yes | Warlord's Dreadweave Robe (231591, -5.52 DPS) [pvp]; Robe of Winter Night (14136, -9.79 DPS, sim-verified) [crafted]; Legionnaire's Dreadweave Tunic (227094, -9.98 DPS) [pvp] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.7 spell_power points (7.55 DPS) | yes | Netherpearl Cuffs (254067, -0.91 DPS) [crafted]; Spidertank Oilrag (9448, -2.58 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.65 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.4 spell_power points (12.37 DPS) | yes | General's Dreadweave Gloves (231589, +0.00 DPS) [pvp]; Blood Guard's Dreadweave Handwraps (227099, -0.18 DPS) [pvp]; Earth Warder's Gloves (21318, -4.20 DPS, sim-verified) [quest] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 17.2 spell_power points (9.52 DPS) | yes | Satyrmane Sash (17755, -0.31 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -1.38 DPS) [crafted]; Ghostweave Cord (254073, -1.79 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (14.37 DPS) | yes | Legionnaire's Dreadweave Legguards (227097, +0.00 DPS) [pvp]; General's Dreadweave Pants (231588, +0.00 DPS) [pvp]; Felcloth Pants (14107, -5.56 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.26 DPS) | yes | General's Dreadweave Boots (231593, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.44 DPS) [dungeon]; Venomspew Footpads (275606, -1.11 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.9 spell_power points (10.98 DPS) | yes | Underworld Band (1980, -4.20 DPS) [world_drop]; Philanthropist's Ring (281635, -4.72 DPS) [quest]; Snake Hoop (6750, -6.93 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.18 DPS) | yes | Underworld Band (1980, -0.41 DPS) [world_drop]; Philanthropist's Ring (281635, -0.92 DPS) [quest]; Snake Hoop (6750, -3.13 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (645.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (645.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (645.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Inventor's Focal Sword (17719, -1.77 DPS) [dungeon]; Spire of Hakkar (10844, -3.63 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 144.9 spell_power points (80.08 DPS) | yes | Brilliant Wand (249385, -14.49 DPS) [crafted]; Stormrager (16997, -16.65 DPS) [quest]; Greater Eternal Wand (249237, -21.77 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of the Void; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Band of the Unicorn; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

