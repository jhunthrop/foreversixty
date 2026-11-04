# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 31.6. Weights run: 1.2s. Verify run: 0.8s. 190 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.431 ± 0.010, crit=0.093 ± 0.002 per rating point (14 rating = 1%, 1.295 per %), hit=0.311 ± 0.010 per rating point (10 rating = 1%, 3.112 per %), spell_haste=0.430 ± 0.083, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.051 ± 0.092), mp5=not significant (0.311 ± 0.120)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (0.91 DPS) | yes | Pristine Circlet (253949, -0.51 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.71 DPS) [crafted]; Flying Tiger Goggles (4368, -0.89 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.1 spell_power points (0.01 DPS) | yes | Tarnished Locket (279870, -0.23 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.9 spell_power points (0.73 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted]; Slime-encrusted Pads (6461, -0.66 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.2 spell_power points (0.35 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Sanguine Cape (14376, -0.21 DPS) [world_drop] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 16.2 spell_power points (1.33 DPS) | yes | Filigreed Shadow Gown (253907, -0.49 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.73 DPS) [crafted]; Gray Woolen Robe (2585, -0.83 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 3.1 spell_power points (0.26 DPS) | yes | Windsong Bangles (263336, -0.17 DPS) [quest]; Repurposed Hair Band (281256, -0.18 DPS) [quest]; Bright Bracers (3647, -0.63 DPS, sim-verified) [world_drop] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 8.3 spell_power points (0.69 DPS) | yes | Gnoll Casting Gloves (892, -0.19 DPS) [world]; Pristine Gloves (253913, -0.25 DPS) [crafted]; Serpent Gloves (5970, -0.42 DPS, sim-verified) [dungeon] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 8.7 spell_power points (0.72 DPS) | yes | Novice Arcanist's Sash (253885, -0.28 DPS) [crafted]; Pristine Sash (253925, -0.31 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.44 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 13.8 spell_power points (1.14 DPS) | yes | Filigreed Pristine Leggings (253937, -0.51 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.56 DPS) [dungeon]; Colorful Kilt (10048, -0.71 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 spell_power points (0.72 DPS) | yes | Shadow Boots (253895, -0.32 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.37 DPS) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.9 spell_power points (0.48 DPS) | yes | Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Furlfeather Ring (286745, -0.40 DPS) [world]; The 1 Ring (8350, -0.44 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.6 spell_power points (0.21 DPS) | yes | Furlfeather Ring (286745, -0.13 DPS) [world]; The 1 Ring (8350, -0.17 DPS) [world]; Volcanic Rock Ring (12053, -0.50 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Witching Stave (1484) | Shadowfang Keep [dungeon] | 29.2 spell_power points (2.41 DPS) | yes | Lesser Staff of the Spire (1300, -2.20 DPS) [world]; Staff of Westfall (2042, -2.21 DPS) [quest]; Twisted Chanter's Staff (890, -2.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 269.9 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.63 DPS, sim-verified) [world_drop]; Deepblaze (279896, -3.78 DPS) [quest]; Sizzle Stick (8071, -4.50 DPS) [quest] |

**New at 20:** head: Shadow Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Witching Stave; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 51.6. Weights run: 1.2s. Verify run: 0.8s. 321 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.703 ± 0.012, crit=0.108 ± 0.002 per rating point (14 rating = 1%, 1.513 per %), hit=0.319 ± 0.012 per rating point (10 rating = 1%, 3.194 per %), spell_haste=3.606 ± 0.166, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.094 ± 0.100), mp5=not significant (0.464 ± 0.161)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Hood (4323) | Tailoring [crafted] | 14.7 spell_power points (1.46 DPS) | yes | Filigreed Shadow Circlet (253981, -0.07 DPS) [crafted]; Enchanter's Cowl (4322, -0.17 DPS) [crafted]; Holy Shroud (2721, -0.31 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.2 spell_power points (0.42 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.11 DPS) [world_drop]; Pendant of Myzrael (4614, -0.36 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (1.52 DPS) | yes | Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Magician's Mantle (12998, -0.40 DPS) [world_drop]; Invoker's Mantle (215365, -0.48 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.50 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Repairman's Cape (9605, -0.04 DPS) [quest]; Heavy Woolen Cloak (4311, -0.06 DPS) [crafted] |
| chest | Shadow Gown (253967) | Tailoring [crafted] | 18.3 spell_power points (1.81 DPS) | yes | Beguiler Robes (7728, +0.00 DPS, sim-verified) [dungeon]; Green Silk Armor (7065, -0.02 DPS) [crafted]; Necrology Robes (2292, -0.08 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 spell_power points (1.03 DPS) | yes | Mindthrust Bracers (1974, -0.54 DPS) [dungeon]; Nightsky Wristbands (6407, -0.58 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.06 DPS, sim-verified) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 13.9 spell_power points (1.38 DPS) | yes | Shadow Gloves (253919, -0.48 DPS) [crafted]; Hands of Darkness (7047, -0.55 DPS, sim-verified) [crafted]; Town Clerk's Mittens (270029, -0.59 DPS) [quest] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.6 spell_power points (1.15 DPS) | yes | Crimson Silk Belt (7055, -0.07 DPS) [crafted]; Invoker's Cord (215366, -0.11 DPS) [crafted]; Shadow Sash (253931, -0.18 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 19.4 spell_power points (1.92 DPS) | yes | Necromancer Leggings (2277, -0.28 DPS, sim-verified) [world_drop]; Filigreed Shadow Leggings (253943, -0.38 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.65 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 18.3 spell_power points (1.81 DPS) | yes | Acidic Walkers (9454, -0.72 DPS) [dungeon]; Gilded Slippers (254001, -0.74 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -0.84 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.94 DPS) | yes | Minor Channeling Ring (1449, -0.31 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 7.4 spell_power points (0.73 DPS) | yes | Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.37 DPS) [dungeon]; Minor Channeling Ring (1449, -0.65 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | 37.0 spell_power points (3.67 DPS) | yes | Witching Stave (1484, -0.68 DPS, sim-verified) [dungeon]; Death Speaker Scepter (2816, -0.89 DPS) [dungeon]; Glimmering Staff (249392, -2.90 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 334.2 spell_power points (33.14 DPS) | yes | Wand of Mana Concentration (273043, -1.73 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -2.78 DPS) [dungeon]; Gravestone Scepter (7001, -3.51 DPS) [quest] |

**New at 30:** head: Shadow Hood; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Shadow Gown; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Belt of Arugal; legs: Shadow Leggings; feet: Black Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Staff of the Shade; ranged: Starfaller

No-known-source sample (15 of 321, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 125.9. Weights run: 1.2s. Verify run: 0.9s. 422 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.092 ± 0.024, crit=0.243 ± 0.007 per rating point (14 rating = 1%, 3.405 per %), hit=0.403 ± 0.021 per rating point (10 rating = 1%, 4.034 per %), spell_haste=0.617 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.095 ± 0.183), mp5=4.153 ± 0.338

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.6 spell_power points (4.34 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.91 DPS) [crafted]; Shadow Hood (4323, -1.23 DPS) [crafted]; Augural Shroud (2620, -2.87 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (2.57 DPS) | yes | Necklace of Calisea (1714, -1.32 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.50 DPS) [world]; Triune Amulet (7722, -1.76 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.2 spell_power points (3.46 DPS) | yes | Green Silken Shoulders (7057, -0.19 DPS) [crafted]; Bloodmage Mantle (7684, -0.39 DPS) [dungeon]; Death Speaker Mantle (6685, -0.52 DPS) [dungeon] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.5 spell_power points (1.87 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.45 DPS) [dungeon]; Glyphed Cloak (4732, -0.62 DPS) [world_drop] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.6 spell_power points (4.67 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.24 DPS) [crafted]; Green Silk Armor (7065, -0.88 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 21.5 spell_power points (3.51 DPS) | yes | Mindthrust Bracers (1974, -0.58 DPS) [dungeon]; Windchaser Cuffs (14429, -1.90 DPS) [world_drop]; Arcane Runed Bracers (4744, -2.04 DPS) [quest] |
| hands | Shadoweave Gloves (10023) | Tailoring [crafted] | 22.5 spell_power points (3.67 DPS) | yes | Dreamweave Gloves (10019, +0.00 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.09 DPS) [crafted]; Black Handwraps (254027, -0.13 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.4 spell_power points (3.82 DPS) | yes | Black Cord (254043, -0.11 DPS) [crafted]; Gilded Cord (254037, -1.09 DPS) [crafted]; Crimson Silk Belt (7055, -1.59 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.1 spell_power points (4.43 DPS) | yes | Crimson Silk Pantaloons (7062, -0.80 DPS) [crafted]; Necromancer Leggings (2277, -0.83 DPS) [world_drop]; Leech Pants (6910, -2.50 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.92 DPS) | yes | Black Slippers (254007, -0.55 DPS) [crafted]; Gilded Slippers (254001, -1.53 DPS) [crafted]; Acidic Walkers (9454, -1.68 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.5 spell_power points (3.18 DPS) | yes | Philanthropist's Ring (281635, -0.65 DPS) [quest]; Snake Hoop (6750, -1.31 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.87 DPS) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 18.5 spell_power points (3.02 DPS) | yes | Philanthropist's Ring (281635, -0.49 DPS) [quest]; Snake Hoop (6750, -1.14 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.71 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (125.9 DPS) | yes | Witching Stave (1484, -0.74 DPS) [dungeon]; Death Speaker Scepter (2816, -1.47 DPS) [dungeon]; Staff of Jordan (873, -1.49 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 247.5 spell_power points (40.45 DPS) | yes | Earthen Rod (9381, -3.57 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.78 DPS) [dungeon]; Twisted Nether Wand (249144, -5.47 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Shadoweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; ranged: Jaina's Firestarter

No-known-source sample (15 of 422, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 194.9. Weights run: 1.2s. Verify run: 1.1s. 541 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.212 ± 0.030, crit=0.376 ± 0.010 per rating point (14 rating = 1%, 5.270 per %), hit=0.603 ± 0.031 per rating point (10 rating = 1%, 6.029 per %), spell_haste=7.434 ± 0.425, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.672 ± 0.260), mp5=not significant (0.863 ± 0.376)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 45.5 spell_power points (7.26 DPS) | yes | Red Mageweave Headband (10033, -0.35 DPS) [crafted]; Dreamweave Circlet (10041, -0.68 DPS) [crafted]; Gemburst Circlet (10751, -1.40 DPS) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 19.7 spell_power points (3.14 DPS) | yes | Mindburst Medallion (11196, -0.48 DPS) [quest]; Necklace of Calisea (1714, -1.03 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.68 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 34.8 spell_power points (5.56 DPS) | yes | Red Mageweave Shoulders (10029, -1.54 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.70 DPS) [vendor]; Netherpearl Shoulders (254055, -3.63 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.3 spell_power points (4.04 DPS) | yes | Runecloth Cloak (13860, -1.05 DPS) [crafted]; Imperial Red Cloak (8248, -1.48 DPS) [world_drop]; Big Voodoo Cloak (8216, -1.50 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 43.2 spell_power points (6.90 DPS) | yes | Robe of Power (7054, -1.49 DPS) [crafted]; Robe of the Magi (1716, -1.69 DPS) [world_drop]; Dreamweave Vest (10021, -4.75 DPS, sim-verified) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 24.8 spell_power points (3.97 DPS) | yes | Bloodband Bracers (11469, -1.10 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.32 DPS) [quest]; Nethergeld Cuffs (254061, -1.41 DPS, sim-verified) [crafted] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 36.5 spell_power points (5.83 DPS) | yes | Dreamweave Gloves (10019, -1.43 DPS) [crafted]; Gilded Gloves (254095, -1.44 DPS) [crafted]; Runecloth Gloves (13863, -2.36 DPS, sim-verified) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 29.9 spell_power points (4.77 DPS) | yes | Dawnspire Cord (12466, -0.14 DPS) [dungeon]; Black Cord (254043, -0.35 DPS) [crafted]; Satyrmane Sash (17755, -0.60 DPS) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 40.5 spell_power points (6.47 DPS) | yes | Senior Designer's Pantaloons (11841, -1.27 DPS) [dungeon]; Knight's Dreadweave Leggings (220888, -1.39 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.08 DPS, sim-verified) [dungeon] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 30.9 spell_power points (4.93 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -0.95 DPS) [vendor]; Black Slippers (254007, -1.08 DPS) [crafted]; Earthen Silk Slippers (254013, -3.19 DPS, sim-verified) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 20.1 spell_power points (3.20 DPS) | yes | Eye of Adaegus (5266, -0.75 DPS) [world_drop]; Band of the Unicorn (7553, -1.13 DPS) [world_drop]; Snake Hoop (6750, -1.28 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (2.56 DPS) | yes | Band of the Unicorn (7553, -0.49 DPS) [world_drop]; Snake Hoop (6750, -0.64 DPS) [quest]; Eye of Adaegus (5266, -1.87 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (194.9 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.01 DPS) [quest] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (194.9 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, -0.12 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.21 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (194.9 DPS) | yes | Spire of Hakkar (10844, -1.46 DPS) [world]; Zum'rah's Vexing Cane (18082, -2.52 DPS) [dungeon]; Kindling Stave (11750, -2.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 330.9 spell_power points (52.82 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.40 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.79 DPS) [crafted] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; feet: Black Sandals; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Evonice's Landin' Pilla; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 541, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 383.4. Weights run: 1.2s. Verify run: 2.0s. 1226 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.503 ± 0.035, crit=0.586 ± 0.013 per rating point (14 rating = 1%, 8.198 per %), hit=0.989 ± 0.043 per rating point (10 rating = 1%, 9.888 per %), spell_haste=1.172 ± 0.273, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=2.566 ± 0.536, mp5=3.425 ± 0.597

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 81.6 spell_power points (18.21 DPS) | yes | Dreadmist Mask (16698, +0.00 DPS) [dungeon]; Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 38.5 spell_power points (8.59 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Necklace of Calisea (1714, -2.23 DPS) [dungeon]; Mindburst Medallion (11196, -2.38 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+16.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Satin Epaulets (231621, -0.80 DPS) [vendor]; Lieutenant Commander's Satin Epaulets (227136, -0.91 DPS) [vendor]; Shroud of the Nathrezim (18720, -16.12 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 38.4 spell_power points (8.58 DPS) | yes | Indomitable Cloak (14683, -1.13 DPS) [world_drop]; Sergeant Major's Cape (16337, -1.31 DPS) [pvp]; Ethereal Mist Cape (11873, -4.30 DPS, sim-verified) [quest] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 87.8 spell_power points (19.60 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Ghostweave Vest (14141, -0.02 DPS) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 41.7 spell_power points (9.31 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Netherpearl Cuffs (254067, -1.19 DPS) [crafted]; Magister's Bindings (16683, -1.51 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (+16.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreadmist Wraps (16705, -0.96 DPS) [dungeon]; Swarmtender's Gloves (275607, -1.64 DPS) [crafted]; Magister's Gloves (16684, -16.68 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 57.2 spell_power points (12.76 DPS) | yes | Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Elunarian Belt (14465, -0.32 DPS) [world_drop]; Wisdom of the Timbermaw (19047, -0.45 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+16.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreadmist Leggings (16699, -0.59 DPS) [dungeon]; Dalewind Trousers (13008, -2.65 DPS) [world_drop]; Magister's Leggings (16687, -16.45 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | 62.7 spell_power points (13.99 DPS) | yes | Mooncloth Boots (15802, -0.06 DPS) [crafted]; Wolfrunner Shoes (13101, -1.09 DPS) [world_drop]; Elunarian Boots (14458, -2.24 DPS) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 43.9 spell_power points (9.80 DPS) | yes | Eye of Adaegus (5266, -0.92 DPS) [world_drop]; Choking Band (11868, -2.36 DPS) [quest]; Coldwater Ring (4550, -3.64 DPS) [quest] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 43.3 spell_power points (9.66 DPS) | yes | Eye of Adaegus (5266, -0.77 DPS) [world_drop]; Choking Band (11868, -2.21 DPS) [quest]; Coldwater Ring (4550, -3.49 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -2.25 DPS, sim-verified) [world_drop] |
| trinket2 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -3.05 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -12.95 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 378.1 spell_power points (84.41 DPS) | yes | Stormrager (16997, -2.84 DPS, sim-verified) [quest]; Brilliant Wand (249385, -20.46 DPS) [crafted]; Torch of Austen (13004, -27.62 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket2: Shard of the Splithooves; main_hand: Hand of Edward the Odd; ranged: Torch of Light

No-known-source sample (15 of 1226, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 707.8. Weights run: 1.3s. Verify run: 2.0s. 1226 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.174 ± 0.005, crit=0.666 ± 0.013 per rating point (14 rating = 1%, 9.323 per %), hit=0.801 ± 0.036 per rating point (10 rating = 1%, 8.006 per %), spell_haste=0.772 ± 0.189, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 31.7 spell_power points (14.54 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.68 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (6.29 DPS) | yes | Triune Amulet (7722, -2.63 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.75 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.07 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.6 spell_power points (13.10 DPS) | yes | Field Marshal's Satin Epaulets (231621, -0.05 DPS) [vendor]; Argent Shoulders (19059, -1.64 DPS) [crafted]; Felcloth Shoulders (14112, -3.48 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | 24.3 spell_power points (11.14 DPS) | yes | Spritecaster Cape (11623, -6.19 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -6.38 DPS) [crafted]; Nightfall Drape (12465, -7.02 DPS) [dungeon] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 42.1 spell_power points (19.28 DPS) | yes | Field Marshal's Satin Vestments (17605, -1.93 DPS) [vendor]; Field Marshal's Satin Robe (231618, -1.93 DPS) [vendor]; Felcloth Robe (14106, -3.36 DPS, sim-verified) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 14.2 spell_power points (6.51 DPS) | yes | Runecloth Cuffs (254123, -0.68 DPS) [crafted]; Arena Wristguards (18709, -2.24 DPS) [world]; Arcane Runed Bracers (4744, -2.39 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.4 spell_power points (13.46 DPS) | yes | Marshal's Satin Gloves (17608, -1.89 DPS) [vendor]; Marshal's Satin Grips (231617, -1.89 DPS) [vendor]; Earth Warder's Gloves (21318, -7.00 DPS, sim-verified) [quest] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 20.6 spell_power points (9.42 DPS) | yes | Satyrmane Sash (17755, -2.35 DPS, sim-verified) [dungeon]; Black Cord (254043, -2.37 DPS) [crafted]; Ghostweave Cord (254073, -3.01 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 28.1 spell_power points (12.87 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Ghostweave Pants (14144, -0.96 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.00 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Omnicast Boots (11822, -0.88 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (8.80 DPS) | yes | Band of the Unicorn (7553, -2.85 DPS) [world_drop]; Philanthropist's Ring (281635, -3.82 DPS) [quest]; Band of the Penitent (13217, -4.53 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 14.9 spell_power points (6.81 DPS) | yes | Band of the Unicorn (7553, -0.86 DPS) [world_drop]; Philanthropist's Ring (281635, -1.83 DPS) [quest]; Band of the Penitent (13217, -2.54 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (707.8 DPS) | yes | Uther's Strength (11302, -3.71 DPS, sim-verified) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (707.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (707.8 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -9.71 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 173.5 spell_power points (79.51 DPS) | yes | Brilliant Wand (249385, -15.25 DPS) [crafted]; Stormrager (16997, -16.42 DPS) [quest]; Greater Eternal Wand (249237, -22.06 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of Winter Night; hands: Gloves of Spell Mastery; feet: Earthen Silk Slippers; finger1: Maiden's Circle; finger2: Underworld Band; trinket2: Mark of the Chosen; main_hand: Hand of Edward the Odd; ranged: Torch of Light

No-known-source sample (15 of 1226, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 33.8. Weights run: 1.2s. Verify run: 0.8s. 183 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.431 ± 0.010, crit=0.093 ± 0.002 per rating point (14 rating = 1%, 1.295 per %), hit=0.311 ± 0.010 per rating point (10 rating = 1%, 3.112 per %), spell_haste=0.430 ± 0.083, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.051 ± 0.092), mp5=not significant (0.311 ± 0.120)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Circlet (253955) | Tailoring [crafted] | 11.0 spell_power points (0.91 DPS) | yes | Pristine Circlet (253949, -0.50 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.71 DPS) [crafted]; Flying Tiger Goggles (4368, -0.89 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.2 spell_power points (0.01 DPS) | yes | Scholarly Pendant (277203, +0.00 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.9 spell_power points (0.73 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.52 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.66 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.2 spell_power points (0.35 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Mystic Shawl (3449, -0.18 DPS) [quest] |
| chest | Necrology Robes (2292) | Shadowfang Keep [dungeon] | 16.2 spell_power points (1.33 DPS) | yes | Filigreed Shadow Gown (253907, -0.25 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.73 DPS) [crafted]; Gray Woolen Robe (2585, -0.83 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 3.1 spell_power points (0.26 DPS) | yes | Owlbeard Bracers (16981, -0.10 DPS) [quest]; Bright Bracers (3647, -0.11 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.46 DPS, sim-verified) [quest] |
| hands | Shadow Gloves (253919) | Tailoring [crafted] | 8.3 spell_power points (0.69 DPS) | yes | Gnoll Casting Gloves (892, -0.19 DPS) [world]; Pristine Gloves (253913, -0.25 DPS) [crafted]; Serpent Gloves (5970, -0.40 DPS, sim-verified) [dungeon] |
| waist | Shadow Sash (253931) | Tailoring [crafted] | 8.7 spell_power points (0.72 DPS) | yes | Novice Arcanist's Sash (253885, -0.28 DPS) [crafted]; Pristine Sash (253925, -0.30 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.44 DPS) [crafted] |
| legs | Filigreed Shadow Leggings (253943) | Tailoring [crafted] | 13.8 spell_power points (1.14 DPS) | yes | Filigreed Pristine Leggings (253937, -0.50 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.56 DPS) [dungeon]; Colorful Kilt (10048, -0.71 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 spell_power points (0.72 DPS) | yes | Shadow Boots (253895, -0.12 DPS) [crafted]; Pristine Boots (253889, -0.37 DPS) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.6 spell_power points (0.21 DPS) | yes | Volcanic Rock Ring (12053, -0.11 DPS) [world_drop]; Furlfeather Ring (286745, -0.13 DPS) [world]; The 1 Ring (8350, -0.17 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.2 spell_power points (0.18 DPS) | yes | Furlfeather Ring (286745, -0.10 DPS) [world]; The 1 Ring (8350, -0.14 DPS) [world]; Volcanic Rock Ring (12053, -0.17 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Witching Stave (1484) | Shadowfang Keep [dungeon] | 29.2 spell_power points (2.41 DPS) | yes | Crescent Staff (6505, -2.16 DPS) [quest]; Lesser Staff of the Spire (1300, -2.20 DPS) [world]; Twisted Chanter's Staff (890, -2.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 269.9 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.57 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.50 DPS) [quest]; Greater Magic Wand (11288, -4.64 DPS) [crafted] |

**New at 20:** head: Shadow Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Necrology Robes; wrist: Mindthrust Bracers; hands: Shadow Gloves; waist: Shadow Sash; legs: Filigreed Shadow Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Witching Stave; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 54.1. Weights run: 1.2s. Verify run: 0.8s. 313 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.703 ± 0.012, crit=0.108 ± 0.002 per rating point (14 rating = 1%, 1.513 per %), hit=0.319 ± 0.012 per rating point (10 rating = 1%, 3.194 per %), spell_haste=3.606 ± 0.166, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.094 ± 0.100), mp5=not significant (0.464 ± 0.161)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Shadow Hood (4323) | Tailoring [crafted] | 14.7 spell_power points (1.46 DPS) | yes | Filigreed Shadow Circlet (253981, -0.07 DPS) [crafted]; Enchanter's Cowl (4322, -0.17 DPS) [crafted]; Holy Shroud (2721, -0.31 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.2 spell_power points (0.42 DPS) | yes | Kaleidoscope Chain (13084, -0.10 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.11 DPS) [world_drop]; Pendant of Myzrael (4614, -0.36 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.3 spell_power points (1.52 DPS) | yes | Death Speaker Mantle (6685, -0.29 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.40 DPS) [world_drop]; Invoker's Mantle (215365, -0.48 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | 5.4 spell_power points (0.53 DPS) | yes | Heavy Woolen Cloak (4311, -0.10 DPS) [crafted]; Resilient Cape (14400, -0.11 DPS) [world_drop]; Hillman's Cloak (3719, -0.40 DPS, sim-verified) [crafted] |
| chest | Shadow Gown (253967) | Tailoring [crafted] | 18.3 spell_power points (1.81 DPS) | yes | Green Silk Armor (7065, -0.02 DPS) [crafted]; Necrology Robes (2292, -0.08 DPS) [dungeon]; Beguiler Robes (7728, -0.37 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.4 spell_power points (1.03 DPS) | yes | Mindthrust Bracers (1974, -0.54 DPS) [dungeon]; Nightsky Wristbands (6407, -0.58 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.37 DPS, sim-verified) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 13.9 spell_power points (1.38 DPS) | yes | Hands of Darkness (7047, -0.44 DPS) [crafted]; Shadow Gloves (253919, -0.48 DPS) [crafted]; Jutebraid Gloves (10654, -0.75 DPS, sim-verified) [quest] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.6 spell_power points (1.15 DPS) | yes | Warsong Sash (16975, -0.06 DPS) [quest]; Crimson Silk Belt (7055, -0.07 DPS) [crafted]; Invoker's Cord (215366, -0.11 DPS) [crafted] |
| legs | Shadow Leggings (253993) | Tailoring [crafted] | 19.4 spell_power points (1.92 DPS) | yes | Filigreed Shadow Leggings (253943, -0.38 DPS) [crafted]; Necromancer Leggings (2277, -0.50 DPS, sim-verified) [world_drop]; Gaze Dreamer Pants (6903, -0.65 DPS) [dungeon] |
| feet | Black Slippers (254007) | Tailoring [crafted] | 18.3 spell_power points (1.81 DPS) | yes | Acidic Walkers (9454, -0.72 DPS) [dungeon]; Gilded Slippers (254001, -0.73 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -0.84 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.94 DPS) | yes | Solid Gold Band (279845, -0.50 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.58 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 7.4 spell_power points (0.73 DPS) | yes | Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.37 DPS) [dungeon]; Solid Gold Band (279845, -0.76 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | 37.0 spell_power points (3.67 DPS) | yes | Death Speaker Scepter (2816, -0.89 DPS) [dungeon]; Witching Stave (1484, -1.27 DPS, sim-verified) [dungeon]; Glimmering Staff (249392, -2.90 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 334.2 spell_power points (33.14 DPS) | yes | Wand of Mana Concentration (273043, -1.93 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -2.78 DPS) [dungeon]; Unstable Power Core (279847, -3.43 DPS) [quest] |

**New at 30:** head: Shadow Hood; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Shadow Gown; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Belt of Arugal; legs: Shadow Leggings; feet: Black Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Staff of the Shade; ranged: Starfaller

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 123.2. Weights run: 1.2s. Verify run: 0.9s. 414 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.092 ± 0.024, crit=0.243 ± 0.007 per rating point (14 rating = 1%, 3.405 per %), hit=0.403 ± 0.021 per rating point (10 rating = 1%, 4.034 per %), spell_haste=0.617 ± 0.133, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (-0.095 ± 0.183), mp5=4.153 ± 0.338

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.6 spell_power points (4.34 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.91 DPS) [crafted]; Shadow Hood (4323, -1.23 DPS) [crafted]; Augural Shroud (2620, -2.79 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (2.57 DPS) | yes | Necklace of Calisea (1714, -1.32 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.50 DPS) [world]; Triune Amulet (7722, -1.78 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.2 spell_power points (3.46 DPS) | yes | Green Silken Shoulders (7057, -0.19 DPS) [crafted]; Bloodmage Mantle (7684, -0.39 DPS) [dungeon]; Death Speaker Mantle (6685, -0.52 DPS) [dungeon] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.5 spell_power points (1.87 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.45 DPS) [dungeon]; Glyphed Cloak (4732, -0.62 DPS) [world_drop] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.6 spell_power points (4.67 DPS) | yes | Dreamweave Vest (10021, -0.12 DPS) [crafted]; Robe of Power (7054, -0.24 DPS) [crafted]; Green Silk Armor (7065, -0.88 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 21.5 spell_power points (3.51 DPS) | yes | Mindthrust Bracers (1974, -1.03 DPS, sim-verified) [dungeon]; Radiant Silver Bracers (4545, -1.43 DPS) [quest]; Windchaser Cuffs (14429, -1.90 DPS) [world_drop] |
| hands | Shadoweave Gloves (10023) | Tailoring [crafted] | 22.5 spell_power points (3.67 DPS) | yes | Dreamweave Gloves (10019, -0.02 DPS) [crafted]; Red Mageweave Gloves (10018, -0.09 DPS) [crafted]; Black Handwraps (254027, -0.13 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.4 spell_power points (3.82 DPS) | yes | Black Cord (254043, -0.11 DPS) [crafted]; Gilded Cord (254037, -1.09 DPS) [crafted]; Crimson Silk Belt (7055, -1.59 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.1 spell_power points (4.43 DPS) | yes | Crimson Silk Pantaloons (7062, -0.80 DPS) [crafted]; Necromancer Leggings (2277, -0.83 DPS) [world_drop]; Leech Pants (6910, -2.78 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.92 DPS) | yes | Black Slippers (254007, -0.73 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.53 DPS) [crafted]; Acidic Walkers (9454, -1.68 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 19.5 spell_power points (3.18 DPS) | yes | Philanthropist's Ring (281635, -0.65 DPS) [quest]; Snake Hoop (6750, -1.31 DPS) [quest]; Voodoo Band (1996, -1.93 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 18.5 spell_power points (3.02 DPS) | yes | Philanthropist's Ring (281635, -0.68 DPS, sim-verified) [quest]; Snake Hoop (6750, -1.14 DPS) [quest]; Voodoo Band (1996, -1.77 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (123.2 DPS) | yes | Death Speaker Scepter (2816, -1.47 DPS) [dungeon]; Witching Stave (1484, -1.48 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -1.49 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 247.5 spell_power points (40.45 DPS) | yes | Earthen Rod (9381, -4.15 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.78 DPS) [dungeon]; Twisted Nether Wand (249144, -5.47 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Shadoweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Underworld Band; ranged: Jaina's Firestarter

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 188.8. Weights run: 1.2s. Verify run: 1.1s. 533 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.212 ± 0.030, crit=0.376 ± 0.010 per rating point (14 rating = 1%, 5.270 per %), hit=0.603 ± 0.031 per rating point (10 rating = 1%, 6.029 per %), spell_haste=7.434 ± 0.425, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=not significant (0.672 ± 0.260), mp5=not significant (0.863 ± 0.376)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bad Mojo Mask (9470) | Zul'Farrak: Shadowpriest Sezz'ziz [dungeon] | 45.5 spell_power points (7.26 DPS) | yes | Red Mageweave Headband (10033, -0.35 DPS) [crafted]; Dreamweave Circlet (10041, -0.68 DPS) [crafted]; Gemburst Circlet (10751, -1.40 DPS) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 19.7 spell_power points (3.14 DPS) | yes | Mindburst Medallion (11196, -0.48 DPS) [quest]; Necklace of Calisea (1714, -1.03 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -1.59 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 34.8 spell_power points (5.56 DPS) | yes | Red Mageweave Shoulders (10029, -1.54 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.70 DPS) [vendor]; Netherpearl Shoulders (254055, -2.27 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.3 spell_power points (4.04 DPS) | yes | Deep Woodlands Cloak (19121, -0.38 DPS) [quest]; Runecloth Cloak (13860, -1.05 DPS) [crafted]; Imperial Red Cloak (8248, -1.48 DPS) [world_drop] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 43.2 spell_power points (6.90 DPS) | yes | Robe of Power (7054, -1.49 DPS) [crafted]; Robe of the Magi (1716, -1.69 DPS) [world_drop]; Dreamweave Vest (10021, -3.28 DPS, sim-verified) [crafted] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 24.8 spell_power points (3.97 DPS) | yes | Bloodband Bracers (11469, -1.10 DPS) [quest]; Nethergeld Cuffs (254061, -1.32 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -1.32 DPS) [quest] |
| hands | Black Gloves (254101) | Tailoring [crafted] | 36.5 spell_power points (5.83 DPS) | yes | Dreamweave Gloves (10019, -1.43 DPS) [crafted]; Gilded Gloves (254095, -1.44 DPS) [crafted]; Runecloth Gloves (13863, -2.26 DPS, sim-verified) [crafted] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 29.9 spell_power points (4.77 DPS) | yes | Dawnspire Cord (12466, -0.14 DPS) [dungeon]; Black Cord (254043, -0.35 DPS) [crafted]; Satyrmane Sash (17755, -0.60 DPS) [dungeon] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 40.5 spell_power points (6.47 DPS) | yes | Senior Designer's Pantaloons (11841, -1.27 DPS) [dungeon]; Stone Guard's Dreadweave Leggings (220906, -1.39 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -1.50 DPS, sim-verified) [dungeon] |
| feet | Black Sandals (254113) | Tailoring [crafted] | 30.9 spell_power points (4.93 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -0.95 DPS) [vendor]; Black Slippers (254007, -1.08 DPS) [crafted]; Earthen Silk Slippers (254013, -3.16 DPS, sim-verified) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 20.1 spell_power points (3.20 DPS) | yes | Eye of Adaegus (5266, -0.75 DPS) [world_drop]; Band of the Unicorn (7553, -1.13 DPS) [world_drop]; Snake Hoop (6750, -1.28 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (2.56 DPS) | yes | Band of the Unicorn (7553, -0.49 DPS) [world_drop]; Snake Hoop (6750, -0.64 DPS) [quest]; Eye of Adaegus (5266, -1.31 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (188.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.67 DPS, sim-verified) [quest] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (188.8 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, -0.12 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.40 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (188.8 DPS) | yes | Kindling Stave (11750, -0.33 DPS) [dungeon]; Spire of Hakkar (10844, -1.46 DPS) [world]; Zum'rah's Vexing Cane (18082, -2.52 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 330.9 spell_power points (52.82 DPS) | yes | Noxious Shooter (17745, -2.02 DPS) [dungeon]; Wand of Allistarj (13065, -4.40 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.79 DPS) [crafted] |

**New at 50:** head: Bad Mojo Mask; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Netherpearl Cuffs; hands: Black Gloves; waist: Black Waistcord; legs: Felcloth Pants; feet: Black Sandals; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Evonice's Landin' Pilla; main_hand: Spellshifter Rod; ranged: Pyric Caduceus

No-known-source sample (15 of 533, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 375.6. Weights run: 1.2s. Verify run: 2.0s. 1219 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.503 ± 0.035, crit=0.586 ± 0.013 per rating point (14 rating = 1%, 8.198 per %), hit=0.989 ± 0.043 per rating point (10 rating = 1%, 9.888 per %), spell_haste=1.172 ± 0.273, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002, spirit=2.566 ± 0.536, mp5=3.425 ± 0.597

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | 81.6 spell_power points (18.21 DPS) | yes | Dreadmist Mask (16698, +0.00 DPS) [dungeon]; Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor] |
| neck | Lady Alizabeth's Pendant (13002) | World drop [world_drop] | 38.5 spell_power points (8.59 DPS) | yes | Horizon Choker (13085, -1.60 DPS) [world_drop]; Necklace of Calisea (1714, -2.23 DPS) [dungeon]; Choker of the High Shaman (4112, -2.33 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (+13.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Satin Epaulets (231611, -0.80 DPS) [vendor]; Champion's Satin Epaulets (227130, -0.91 DPS) [vendor]; Shroud of the Nathrezim (18720, -12.99 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 38.4 spell_power points (8.58 DPS) | yes | Indomitable Cloak (14683, -1.13 DPS) [world_drop]; Arcane Cloak (8286, -1.70 DPS) [world_drop]; Cloak of Blight (6832, -4.71 DPS, sim-verified) [quest] |
| chest | Mooncloth Vest (14138) | Tailoring [crafted] | 87.8 spell_power points (19.60 DPS) | yes | Ghostweave Vest (14141, +0.00 DPS, sim-verified) [crafted]; Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, -0.37 DPS) [vendor] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | 41.7 spell_power points (9.31 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Netherpearl Cuffs (254067, -1.19 DPS) [crafted]; Magister's Bindings (16683, -1.51 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (+14.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreadmist Wraps (16705, -0.96 DPS) [dungeon]; Swarmtender's Gloves (275607, -1.64 DPS) [crafted]; Magister's Gloves (16684, -14.80 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | 57.2 spell_power points (12.76 DPS) | yes | General's Satin Cinch (17621, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, -0.45 DPS) [crafted]; Elunarian Belt (14465, -2.36 DPS, sim-verified) [world_drop] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (+13.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreadmist Leggings (16699, -0.59 DPS) [dungeon]; Legionnaire's Satin Trousers (17611, -2.23 DPS) [pvp]; Magister's Leggings (16687, -13.85 DPS, sim-verified) [dungeon] |
| feet | Devout Sandals (16691) | Stratholme: Maleki the Pallid [dungeon] | 62.7 spell_power points (13.99 DPS) | yes | Mooncloth Boots (15802, -0.06 DPS) [crafted]; Wolfrunner Shoes (13101, -1.09 DPS) [world_drop]; Elunarian Boots (14458, -2.24 DPS) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 43.9 spell_power points (9.80 DPS) | yes | Eye of Adaegus (5266, -0.92 DPS) [world_drop]; Coldwater Ring (4550, -3.64 DPS) [quest]; Welken Ring (5011, -4.12 DPS) [world_drop] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 43.3 spell_power points (9.66 DPS) | yes | Eye of Adaegus (5266, -1.85 DPS, sim-verified) [world_drop]; Coldwater Ring (4550, -3.49 DPS) [quest]; Welken Ring (5011, -3.97 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Alchemist's Stone (13503, -2.30 DPS) [crafted]; Evonice's Landin' Pilla (18951, -3.96 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Alchemist's Stone (13503, +0.00 DPS) [crafted]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of Hale Magefire (13000, +0.00 DPS) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -10.56 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 378.1 spell_power points (84.41 DPS) | yes | Stormrager (16997, -3.49 DPS, sim-verified) [quest]; Brilliant Wand (249385, -20.46 DPS) [crafted]; Torch of Austen (13004, -27.62 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Lady Alizabeth's Pendant; shoulder: Devout Mantle; chest: Mooncloth Vest; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Devout Sandals; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket1: Shard of the Splithooves; trinket2: Frozen Heart of the Mountain; main_hand: Hand of Edward the Odd; ranged: Torch of Light

No-known-source sample (15 of 1219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 711.3. Weights run: 1.3s. Verify run: 2.0s. 1219 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.174 ± 0.005, crit=0.666 ± 0.013 per rating point (14 rating = 1%, 9.323 per %), hit=0.801 ± 0.036 per rating point (10 rating = 1%, 8.006 per %), spell_haste=0.772 ± 0.189, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001, spirit=not significant (0.000 ± 0.000), mp5=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Felcloth Hood (14111) | Tailoring [crafted] | 31.7 spell_power points (14.54 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.47 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (6.29 DPS) | yes | Triune Amulet (7722, -2.63 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.79 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.07 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.6 spell_power points (13.10 DPS) | yes | Warlord's Satin Epaulets (231611, -0.05 DPS) [vendor]; Argent Shoulders (19059, -1.64 DPS) [crafted]; Felcloth Shoulders (14112, -2.88 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | 24.3 spell_power points (11.14 DPS) | yes | Deep Woodlands Cloak (19121, -4.93 DPS) [quest]; Spritecaster Cape (11623, -5.90 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -6.38 DPS) [crafted] |
| chest | Robe of Winter Night (14136) | Tailoring [crafted] | 42.1 spell_power points (19.28 DPS) | yes | Warlord's Satin Robes (231612, -1.93 DPS) [pvp]; Felcloth Robe (14106, -2.92 DPS, sim-verified) [crafted]; Warlord's Satin Tunic (231632, -7.43 DPS) [vendor] |
| wrist | Netherpearl Cuffs (254067) | Tailoring [crafted] | 14.2 spell_power points (6.51 DPS) | yes | Runecloth Cuffs (254123, -0.68 DPS) [crafted]; Arena Wristguards (18709, -2.24 DPS) [world]; Spidertank Oilrag (9448, -2.39 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.4 spell_power points (13.46 DPS) | yes | General's Satin Gloves (17620, -1.89 DPS) [vendor]; General's Satin Grips (231613, -1.89 DPS) [vendor]; Earth Warder's Gloves (21318, -6.07 DPS, sim-verified) [quest] |
| waist | Black Waistcord (254087) | Tailoring [crafted] | 20.6 spell_power points (9.42 DPS) | yes | Black Cord (254043, -2.37 DPS) [crafted]; Satyrmane Sash (17755, -2.46 DPS, sim-verified) [dungeon]; Ghostweave Cord (254073, -3.01 DPS) [crafted] |
| legs | Felcloth Pants (14107) | Tailoring [crafted] | 28.1 spell_power points (12.87 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -0.96 DPS) [crafted]; General's Satin Legguards (231634, -2.03 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.00 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Omnicast Boots (11822, -0.88 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (8.80 DPS) | yes | Band of the Unicorn (7553, -2.85 DPS) [world_drop]; Philanthropist's Ring (281635, -3.82 DPS) [quest]; Band of the Penitent (13217, -4.53 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 14.9 spell_power points (6.81 DPS) | yes | Band of the Unicorn (7553, -0.86 DPS) [world_drop]; Philanthropist's Ring (281635, -1.83 DPS) [quest]; Band of the Penitent (13217, -2.54 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (711.3 DPS) | yes | Uther's Strength (11302, -0.55 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (711.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (711.3 DPS) | yes | Spire of Hakkar (10844, +0.00 DPS) [world]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -8.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 173.5 spell_power points (79.51 DPS) | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -16.42 DPS) [quest]; Greater Eternal Wand (249237, -22.06 DPS) [crafted] |

**New at 60:** head: Felcloth Hood; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Fel Cape; chest: Robe of Winter Night; hands: Gloves of Spell Mastery; feet: Earthen Silk Slippers; finger1: Maiden's Circle; finger2: Underworld Band; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; ranged: Torch of Light

No-known-source sample (15 of 1219, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

