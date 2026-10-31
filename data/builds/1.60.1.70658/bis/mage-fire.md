# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 0000000000000000000-23510000000000000-000000000000000000)

Set DPS (verified): 26.3. Weights run: 0.8s. Verify run: 0.6s. 154 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.443 ± 0.009, crit=0.109 ± 0.004 per rating point (14 rating = 1%, 1.532 per %), hit=0.348 ± 0.012 per rating point (10 rating = 1%, 3.482 per %), spell_haste=-0.989 ± 0.127, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.74 DPS) | yes | Pristine Circlet (253949, -0.33 DPS) [crafted]; Shadow Goggles (4373, -0.59 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.0 spell_power points (0.60 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.33 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.34 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.27 DPS) | yes | Pearl-clasped Cloak (5542, -0.04 DPS) [crafted]; Black Whelp Cloak (7283, -0.07 DPS) [crafted]; Sanguine Cape (14376, -0.15 DPS) [world_drop] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 11.2 spell_power points (0.75 DPS) | yes | Filigreed Pristine Gown (253901, -0.27 DPS) [crafted]; Gray Woolen Robe (2585, -0.33 DPS) [crafted]; Green Woolen Robe (6243, -0.46 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.2 spell_power points (0.15 DPS) | yes | Windsong Bangles (263336, -0.08 DPS) [quest]; Repurposed Hair Band (281256, -0.09 DPS) [quest]; Bright Bracers (3647, -0.47 DPS, sim-verified) [world_drop] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.60 DPS) | yes | Flame Gloves (253917, -0.04 DPS) [crafted]; Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.20 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.8 spell_power points (0.59 DPS) | yes | Pristine Sash (253925, -0.20 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.35 DPS) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.7 spell_power points (0.91 DPS) | yes | Filigreed Pristine Leggings (253937, -0.33 DPS) [crafted]; Phoenix Pants (4317, -0.40 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.45 DPS) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.8 spell_power points (0.59 DPS) | yes | Flame Boots (253893, -0.21 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.30 DPS) [crafted]; Red Woolen Boots (4313, -0.32 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.9 spell_power points (0.39 DPS) | yes | Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; The 1 Ring (8350, -0.36 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.7 spell_power points (0.18 DPS) | yes | The 1 Ring (8350, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.28 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 4.4 spell_power points (0.30 DPS) | yes | Witching Stave (1484, -0.15 DPS) [dungeon]; Staff of Westfall (2042, -0.15 DPS) [quest]; Lesser Staff of the Spire (1300, -0.22 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 336.1 spell_power points (22.51 DPS) | yes | Skycaller (12984, -0.66 DPS, sim-verified) [world_drop]; Deepblaze (279896, -3.98 DPS) [quest]; Sizzle Stick (8071, -4.53 DPS) [quest] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Mindthrust Bracers; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 154, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 0000000000000000000-23552100120000000-000000000000000000)

Set DPS (verified): 48.4. Weights run: 0.8s. Verify run: 0.6s. 292 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.674 ± 0.012, crit=0.172 ± 0.007 per rating point (14 rating = 1%, 2.408 per %), hit=0.361 ± 0.018 per rating point (10 rating = 1%, 3.608 per %), spell_haste=not significant (-0.398 ± 0.134), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Flame Circlet (253979) | Tailoring [crafted] | 14.0 spell_power points (1.32 DPS) | yes | Enchanter's Cowl (4322, -0.12 DPS) [crafted]; Holy Shroud (2721, -0.28 DPS) [world_drop]; Flame Circlet (253953, -0.28 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.0 spell_power points (0.38 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.1 spell_power points (1.42 DPS) | yes | Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop]; Death Speaker Mantle (6685, -0.51 DPS, sim-verified) [dungeon] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.47 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Resilient Cape (14400, -0.09 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.09 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.8 spell_power points (1.68 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon]; Flame Gown (253965, -0.48 DPS, sim-verified) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.23 DPS) | yes | Spidertank Oilrag (9448, -0.38 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.72 DPS) [world_drop]; Nightsky Wristbands (6407, -0.85 DPS) [world_drop] |
| hands | Flame Gloves (253917) | Tailoring [crafted] | 9.0 spell_power points (0.85 DPS) | yes | Phoenix Gloves (4331, -0.00 DPS) [crafted]; Town Clerk's Mittens (270029, -0.12 DPS) [quest]; Truefaith Gloves (7049, -0.19 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.0 spell_power points (1.04 DPS) | yes | Crimson Silk Belt (7055, -0.03 DPS) [crafted]; Invoker's Cord (215366, -0.06 DPS) [crafted]; Flame Sash (253929, -0.13 DPS) [crafted] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.7 spell_power points (1.77 DPS) | yes | Filigreed Flame Leggings (253941, -0.47 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.54 DPS) [world]; Phoenix Pants (4317, -0.57 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.7 spell_power points (1.67 DPS) | yes | Gilded Slippers (254001, -0.57 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.4 spell_power points (0.88 DPS) | yes | Ring of Power Regulation (273030, -0.32 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon]; Azora's Will (4999, -0.57 DPS) [world_drop] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.3 spell_power points (0.60 DPS) | yes | Ring of Power Regulation (273030, +0.00 DPS, sim-verified) [dungeon]; Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Azora's Will (4999, -0.28 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.40 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.51 DPS) [quest]; Glimmering Staff (249392, -2.70 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 350.6 spell_power points (33.11 DPS) | yes | Wand of Mana Concentration (273043, -0.48 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.64 DPS) [crafted]; Scorching Wand (5213, -3.70 DPS) [world_drop] |

**New at 30:** head: Filigreed Flame Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Flame Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Minor Channeling Ring; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 292, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 0000000000000000000-23552100130125010-000000000000000000)

Set DPS (verified): 81.6. Weights run: 1.0s. Verify run: 0.7s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.852 ± 0.027, crit=0.276 ± 0.019 per rating point (14 rating = 1%, 3.869 per %), hit=0.415 ± 0.042 per rating point (10 rating = 1%, 4.151 per %), spell_haste=not significant (1.011 ± 0.322), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 22.5 spell_power points (2.59 DPS) | yes | Augural Shroud (2620, -0.34 DPS) [world]; Whitemane's Chapeau (7720, -0.80 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.23 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.8 spell_power points (1.59 DPS) | yes | Necklace of Calisea (1714, -0.90 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.00 DPS) [world]; Triune Amulet (7722, -1.67 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.1 spell_power points (2.08 DPS) | yes | Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -0.97 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.3 spell_power points (1.18 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.38 DPS) [crafted]; Blackforge Cape (6424, -0.40 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.1 spell_power points (3.12 DPS) | yes | Cindercloth Robe (10042, -0.01 DPS) [crafted]; Dreamweave Vest (10021, -0.17 DPS) [crafted]; Robe of Power (7054, -0.33 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.50 DPS) | yes | Arcane Runed Bracers (4744, -0.46 DPS) [quest]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Windchaser Cuffs (14429, -0.61 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.4 spell_power points (2.47 DPS) | yes | Fiery Handwraps (254025, -0.17 DPS) [crafted]; Red Mageweave Gloves (10018, -0.22 DPS) [crafted]; Crimson Silk Gloves (7064, -0.26 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.8 spell_power points (2.40 DPS) | yes | Deathmage Sash (10771, -0.12 DPS) [dungeon]; Gilded Cord (254037, -0.69 DPS) [crafted]; Scorching Sash (4117, -0.90 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.2 spell_power points (2.79 DPS) | yes | Flame Leggings (253991, -0.49 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.59 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.93 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Fiery Slippers (254005, -0.58 DPS) [crafted]; Gilded Slippers (254001, -1.27 DPS) [crafted]; Acidic Walkers (9454, -1.40 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.64 DPS) | yes | Ring of Forlorn Spirits (2043, -0.72 DPS) [quest]; Reedknot Ring (9622, -0.84 DPS) [quest]; Minor Channeling Ring (1449, -0.87 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 spell_power points (1.18 DPS) | yes | Ring of Forlorn Spirits (2043, -0.26 DPS) [quest]; Reedknot Ring (9622, -0.38 DPS) [quest]; Minor Channeling Ring (1449, -0.41 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.20 DPS) [dungeon]; Searing Golden Blade (12260, -0.67 DPS) [crafted] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (81.6 DPS) | yes | Orb of the Forgotten Seer (7685, -0.06 DPS) [dungeon]; Rod of Molten Fire (2565, -0.17 DPS) [world_drop]; Celestial Orb (7515, -0.93 DPS, sim-verified) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 346.9 spell_power points (39.97 DPS) | yes | Ragefire Wand (7513, -1.50 DPS) [quest]; Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 2050000000000000000-23552100130125031-000000000000000000)

Set DPS (verified): 146.9. Weights run: 0.9s. Verify run: 0.8s. 506 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.876 ± 0.034, crit=0.302 ± 0.028 per rating point (14 rating = 1%, 4.226 per %), hit=0.586 ± 0.055 per rating point (10 rating = 1%, 5.861 per %), spell_haste=not significant (0.716 ± 0.533), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.5 spell_power points (4.90 DPS) | yes | Knight-Lieutenant's Dreadweave Hat (220889, -1.04 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.28 DPS) [crafted]; Dreamweave Circlet (10041, -1.81 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.0 spell_power points (1.88 DPS) | yes | Horizon Choker (13085, -0.23 DPS) [world_drop]; Mindburst Medallion (11196, -0.37 DPS) [quest]; Triune Amulet (7722, -0.81 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.8 spell_power points (3.86 DPS) | yes | Red Mageweave Shoulders (10029, -1.16 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.16 DPS) [vendor]; Netherflame Shoulders (254053, -2.24 DPS, sim-verified) [crafted] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.0 spell_power points (2.68 DPS) | yes | Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.54 DPS) [crafted]; Cloak of Fire (14134, -0.81 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.5 spell_power points (4.90 DPS) | yes | Cindercloth Robe (10042, -1.28 DPS) [crafted]; Runecloth Tunic (13857, -1.33 DPS) [crafted]; Robe of the Magi (1716, -2.85 DPS, sim-verified) [world_drop] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.1 spell_power points (2.57 DPS) | yes | Nethergeld Cuffs (254061, -0.80 DPS) [crafted]; Phoenix Bindings (210781, -0.82 DPS) [crafted]; Bloodband Bracers (11469, -0.84 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 27.8 spell_power points (3.72 DPS) | yes | Dreamweave Gloves (10019, -0.84 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.92 DPS) [vendor]; Fiery Handwraps (254025, -1.02 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 26.9 spell_power points (3.61 DPS) | yes | Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Dawnspire Cord (12466, -0.57 DPS) [dungeon]; Fiery Cord (254041, -0.79 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 26.7 spell_power points (3.59 DPS) | yes | Red Mageweave Pants (10009, -0.30 DPS) [crafted]; Flame Leggings (253991, -0.89 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.99 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 27.9 spell_power points (3.74 DPS) | yes | Earthen Silk Slippers (254013, -0.52 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -0.82 DPS) [vendor]; Cindercloth Boots (10044, -0.92 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.93 DPS) | yes | Snake Hoop (6750, -0.54 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.86 DPS) [quest]; Reedknot Ring (9622, -0.99 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.74 DPS) | yes | Snake Hoop (6750, -0.35 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.67 DPS) [quest]; Reedknot Ring (9622, -0.80 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (146.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (146.9 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (146.9 DPS) | yes | Spellshifter Rod (9527, -0.17 DPS) [quest]; Kindling Stave (11750, -0.66 DPS) [dungeon]; Hanzo Sword (8190, -25.29 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 391.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.83 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.68 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Knight's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 506, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 2050015100000000000-23552100130125031-003000000000000000)

Set DPS (verified): 400.4. Weights run: 1.0s. Verify run: 2.5s. 1125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.820 ± 0.047, crit=0.503 ± 0.040 per rating point (14 rating = 1%, 7.044 per %), hit=1.296 ± 0.088 per rating point (10 rating = 1%, 12.960 per %), spell_haste=4.060 ± 0.754, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 35.6 spell_power points (10.99 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -0.06 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.2 spell_power points (6.24 DPS) | yes | Glowing Eye of Mordresh (10769, -2.05 DPS) [dungeon]; Horizon Choker (13085, -2.69 DPS) [world_drop]; Mindburst Medallion (11196, -2.87 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.7 spell_power points (10.72 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.12 DPS) [pvp]; Rotgrip Mantle (17732, -2.14 DPS) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (6.48 DPS) | yes | Cindercloth Cloak (14044, -0.44 DPS) [crafted]; Spritecaster Cape (11623, -0.64 DPS) [dungeon]; Runecloth Cloak (13860, -1.68 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.9 spell_power points (17.56 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.89 DPS) [pvp]; Robes of Fiery Devastation (279266, -3.46 DPS) [crafted]; Knight-Captain's Silk Tunic (227108, -4.60 DPS) [pvp] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 19.2 spell_power points (5.93 DPS) | yes | Netherflame Cuffs (254065, -0.14 DPS) [crafted]; Marshal's Silk Bracers (16438, -1.12 DPS) [pvp]; Dreadmist Bracers (16703, -1.60 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.3 spell_power points (9.66 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Fiery Gloves (254099, +0.00 DPS) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (400.4 DPS) | yes | Fiery Waistcord (254085, +0.00 DPS) [crafted]; Wisdom of the Timbermaw (19047, -0.54 DPS) [crafted]; Satyrmane Sash (17755, -0.63 DPS) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (400.4 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Cindercloth Pants (14045, -11.59 DPS, sim-verified) [crafted] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.8 spell_power points (9.21 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.20 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.46 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.7 spell_power points (7.33 DPS) | yes | Band of the Unicorn (7553, -3.32 DPS) [world_drop]; Snake Hoop (6750, -4.21 DPS) [quest]; Glowing Crystal Ring (18402, -4.29 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (4.35 DPS) | yes | Band of the Unicorn (7553, -0.34 DPS) [world_drop]; Snake Hoop (6750, -1.23 DPS) [quest]; Glowing Crystal Ring (18402, -1.31 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (400.4 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (400.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (400.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.28 DPS) [quest]; Teebu's Blazing Longsword (1728, -51.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 254.6 spell_power points (78.61 DPS) | yes | Brilliant Wand (249385, -14.48 DPS) [crafted]; Stormrager (16997, -14.66 DPS) [quest]; Torch of Austen (13004, -21.82 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; ranged: Torch of Light

No-known-source sample (15 of 1125, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 2050015100000000000-23452100130135031-003000000000000000)

Set DPS (verified): 744.5. Weights run: 1.1s. Verify run: 1.6s. 1125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.445 ± 0.027, crit=0.828 ± 0.041 per rating point (14 rating = 1%, 11.587 per %), hit=1.149 ± 0.072 per rating point (10 rating = 1%, 11.488 per %), spell_haste=3.815 ± 0.409, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 27.9 spell_power points (16.14 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, +0.00 DPS) [crafted]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.5 spell_power points (9.51 DPS) | yes | Glowing Eye of Mordresh (10769, -3.41 DPS) [dungeon]; Mindburst Medallion (11196, -4.50 DPS) [quest]; Triune Amulet (7722, -4.89 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.4 spell_power points (19.88 DPS) | yes | Field Marshal's Silk Spaulders (231602, -1.56 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.67 DPS) [pvp]; Argent Shoulders (19059, -14.07 DPS, sim-verified) [crafted] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (12.14 DPS) | yes | Spritecaster Cape (11623, -2.50 DPS) [dungeon]; Cindercloth Cloak (14044, -2.57 DPS) [crafted]; Cloak of Fire (14134, -4.05 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.9 spell_power points (32.92 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.76 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -9.70 DPS) [pvp]; Robes of Fiery Devastation (279266, -21.66 DPS, sim-verified) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 16.1 spell_power points (9.32 DPS) | yes | Runecloth Cuffs (254123, -0.38 DPS) [crafted]; Phoenix Bindings (210781, -1.80 DPS) [crafted]; Arena Wristguards (18709, -2.62 DPS) [world] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 36.6 spell_power points (21.18 DPS) | yes | Marshal's Silk Gloves (16440, -2.48 DPS) [vendor]; Marshal's Silk Gauntlets (231608, -2.48 DPS) [vendor]; Fiery Gloves (254099, -18.17 DPS, sim-verified) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 23.0 spell_power points (13.30 DPS) | yes | Fiery Cord (254041, -3.15 DPS) [crafted]; Wisdom of the Timbermaw (19047, -3.34 DPS) [crafted]; Satyrmane Sash (17755, -5.88 DPS, sim-verified) [dungeon] |
| legs | Cindercloth Pants (14045) | Tailoring [crafted] | 31.3 spell_power points (18.12 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -1.40 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.3 spell_power points (14.65 DPS) | yes | Knight-Lieutenant's Silk Walkers (227112, +0.00 DPS) [pvp]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Fiery Sandals (254111, -0.77 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.1 spell_power points (12.21 DPS) | yes | Philanthropist's Ring (281635, -5.14 DPS) [quest]; Snake Hoop (6750, -7.45 DPS) [quest]; Ring of Forlorn Spirits (2043, -7.58 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.52 DPS) | yes | Philanthropist's Ring (281635, -0.45 DPS) [quest]; Snake Hoop (6750, -2.76 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.89 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (744.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (744.5 DPS) | yes | Mark of the Chosen (17774, -7.83 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (744.5 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.64 DPS) [quest]; Teebu's Blazing Longsword (1728, -53.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 138.7 spell_power points (80.23 DPS) | yes | Brilliant Wand (249385, -13.64 DPS) [crafted]; Stormrager (16997, -16.25 DPS) [quest]; Greater Eternal Wand (249237, -21.70 DPS) [crafted] |

**New at 60:** neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; hands: Gloves of Spell Mastery; legs: Cindercloth Pants; feet: Omnicast Boots; finger1: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1125, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 0000000000000000000-23510000000000000-000000000000000000)

Set DPS (verified): 24.5. Weights run: 0.8s. Verify run: 0.6s. 148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.443 ± 0.009, crit=0.109 ± 0.004 per rating point (14 rating = 1%, 1.532 per %), hit=0.348 ± 0.012 per rating point (10 rating = 1%, 3.482 per %), spell_haste=-0.989 ± 0.127, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Flame Circlet (253953) | Tailoring [crafted] | 11.0 spell_power points (0.74 DPS) | yes | Pristine Circlet (253949, -0.33 DPS) [crafted]; Shadow Goggles (4373, -0.59 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.0 spell_power points (0.60 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.33 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.50 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.27 DPS) | yes | Black Whelp Cloak (7283, -0.07 DPS) [crafted]; Mystic Shawl (3449, -0.13 DPS) [quest]; Pearl-clasped Cloak (5542, -0.19 DPS, sim-verified) [crafted] |
| chest | Filigreed Flame Gown (253905) | Tailoring [crafted] | 11.2 spell_power points (0.75 DPS) | yes | Filigreed Pristine Gown (253901, -0.27 DPS) [crafted]; Gray Woolen Robe (2585, -0.33 DPS) [crafted]; Green Woolen Robe (6243, -0.46 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.20 DPS) | yes | Mindthrust Bracers (1974, -0.05 DPS) [dungeon]; Featherbead Bracers (15452, -0.05 DPS) [quest]; Owlbeard Bracers (16981, -0.07 DPS) [quest] |
| hands | Phoenix Gloves (4331) | Tailoring [crafted] | 9.0 spell_power points (0.60 DPS) | yes | Flame Gloves (253917, -0.04 DPS) [crafted]; Serpent Gloves (5970, -0.13 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.20 DPS) [world] |
| waist | Flame Sash (253929) | Tailoring [crafted] | 8.8 spell_power points (0.59 DPS) | yes | Pristine Sash (253925, -0.20 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.35 DPS) [world_drop] |
| legs | Filigreed Flame Leggings (253941) | Tailoring [crafted] | 13.7 spell_power points (0.91 DPS) | yes | Filigreed Pristine Leggings (253937, -0.33 DPS) [crafted]; Silk-threaded Trousers (1929, -0.45 DPS) [dungeon]; Phoenix Pants (4317, -0.53 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.8 spell_power points (0.59 DPS) | yes | Flame Boots (253893, -0.22 DPS, sim-verified) [crafted]; Pristine Boots (253889, -0.30 DPS) [crafted]; Red Woolen Boots (4313, -0.32 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.7 spell_power points (0.18 DPS) | yes | Volcanic Rock Ring (12053, -0.09 DPS) [world_drop]; The 1 Ring (8350, -0.15 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.2 spell_power points (0.15 DPS) | yes | The 1 Ring (8350, -0.12 DPS) [world]; Volcanic Rock Ring (12053, -0.41 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 4.4 spell_power points (0.30 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Clear Crystal Rod (16894, -0.15 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 336.1 spell_power points (22.51 DPS) | yes | Skycaller (12984, -0.78 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.53 DPS) [quest]; Greater Magic Wand (11288, -4.88 DPS) [crafted] |

**New at 20:** head: Flame Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Flame Gown; wrist: Tabitha's Cuffs; hands: Phoenix Gloves; waist: Flame Sash; legs: Filigreed Flame Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 0000000000000000000-23552100120000000-000000000000000000)

Set DPS (verified): 47.1. Weights run: 0.8s. Verify run: 0.6s. 282 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.674 ± 0.012, crit=0.172 ± 0.007 per rating point (14 rating = 1%, 2.408 per %), hit=0.361 ± 0.018 per rating point (10 rating = 1%, 3.608 per %), spell_haste=not significant (-0.398 ± 0.134), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (47.1 DPS) | yes | Holy Shroud (2721, -0.16 DPS) [world_drop]; Flame Circlet (253953, -0.16 DPS) [crafted]; Filigreed Flame Circlet (253979, -0.53 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.0 spell_power points (0.38 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.1 spell_power points (1.42 DPS) | yes | Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.38 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.47 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Resilient Cape (14400, -0.09 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.09 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.8 spell_power points (1.68 DPS) | yes | Flame Gown (253965, -0.00 DPS) [crafted]; Beguiler Robes (7728, -0.06 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.16 DPS) [dungeon] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.23 DPS) | yes | Spidertank Oilrag (9448, -0.38 DPS) [dungeon]; Glowing Magical Bracelets (13106, -0.72 DPS) [world_drop]; Nightsky Wristbands (6407, -0.85 DPS) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.4 spell_power points (0.88 DPS) | yes | Flame Gloves (253917, -0.03 DPS) [crafted]; Phoenix Gloves (4331, -0.03 DPS) [crafted]; Truefaith Gloves (7049, -0.22 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.0 spell_power points (1.04 DPS) | yes | Crimson Silk Belt (7055, -0.03 DPS) [crafted]; Invoker's Cord (215366, -0.06 DPS) [crafted]; Warsong Sash (16975, -0.60 DPS, sim-verified) [quest] |
| legs | Flame Leggings (253991) | Tailoring [crafted] | 18.7 spell_power points (1.77 DPS) | yes | Filigreed Flame Leggings (253941, -0.50 DPS, sim-verified) [crafted]; Smoldering Pants (3073, -0.54 DPS) [world]; Phoenix Pants (4317, -0.57 DPS) [crafted] |
| feet | Fiery Slippers (254005) | Tailoring [crafted] | 17.7 spell_power points (1.67 DPS) | yes | Gilded Slippers (254001, -0.57 DPS) [crafted]; Acidic Walkers (9454, -0.69 DPS) [dungeon]; Spidersilk Boots (4320, -0.76 DPS) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.4 spell_power points (0.88 DPS) | yes | Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon]; Solid Gold Band (279845, -0.51 DPS) [quest]; Azora's Will (4999, -0.57 DPS) [world_drop] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.57 DPS) | yes | Solid Gold Band (279845, -0.19 DPS) [quest]; Azora's Will (4999, -0.25 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.65 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.40 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -2.70 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -2.76 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 350.6 spell_power points (33.11 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Unstable Power Core (279847, -3.40 DPS) [quest]; Greater Mystic Wand (217287, -3.64 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Phoenix Bindings; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Flame Leggings; feet: Fiery Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 282, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (orc, 0000000000000000000-23552100130125010-000000000000000000)

Set DPS (verified): 76.8. Weights run: 1.0s. Verify run: 0.7s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.852 ± 0.027, crit=0.276 ± 0.019 per rating point (14 rating = 1%, 3.869 per %), hit=0.415 ± 0.042 per rating point (10 rating = 1%, 4.151 per %), spell_haste=not significant (1.011 ± 0.322), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 22.5 spell_power points (2.59 DPS) | yes | Augural Shroud (2620, -0.34 DPS) [world]; Whitemane's Chapeau (7720, -0.80 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.40 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.8 spell_power points (1.59 DPS) | yes | Necklace of Calisea (1714, -0.90 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.00 DPS) [world]; Triune Amulet (7722, -1.61 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.1 spell_power points (2.08 DPS) | yes | Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -1.13 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.3 spell_power points (1.18 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.38 DPS) [crafted]; Blackforge Cape (6424, -0.40 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.1 spell_power points (3.12 DPS) | yes | Cindercloth Robe (10042, -0.01 DPS) [crafted]; Dreamweave Vest (10021, -0.17 DPS) [crafted]; Robe of Power (7054, -0.33 DPS) [crafted] |
| wrist | Phoenix Bindings (210781) | Tailoring [crafted] | 13.0 spell_power points (1.50 DPS) | yes | Radiant Silver Bracers (4545, -0.25 DPS) [quest]; Spidertank Oilrag (9448, -0.46 DPS) [dungeon]; Windchaser Cuffs (14429, -0.61 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.4 spell_power points (2.47 DPS) | yes | Fiery Handwraps (254025, -0.17 DPS) [crafted]; Red Mageweave Gloves (10018, -0.22 DPS) [crafted]; Crimson Silk Gloves (7064, -0.26 DPS) [crafted] |
| waist | Fiery Cord (254041) | Tailoring [crafted] | 20.8 spell_power points (2.40 DPS) | yes | Deathmage Sash (10771, -0.12 DPS) [dungeon]; Gilded Cord (254037, -0.69 DPS) [crafted]; Scorching Sash (4117, -0.90 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.2 spell_power points (2.79 DPS) | yes | Crimson Silk Pantaloons (7062, -0.59 DPS) [crafted]; Flame Leggings (253991, -0.92 DPS, sim-verified) [crafted]; Filigreed Flame Leggings (253941, -0.93 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.77 DPS) | yes | Fiery Slippers (254005, -0.89 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.27 DPS) [crafted]; Acidic Walkers (9454, -1.40 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.3 spell_power points (1.64 DPS) | yes | Reedknot Ring (9622, -0.84 DPS) [quest]; Ring of Power Regulation (273030, -0.95 DPS) [dungeon]; Voodoo Band (1996, -0.96 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.3 spell_power points (1.18 DPS) | yes | Ring of Power Regulation (273030, -0.49 DPS) [dungeon]; Voodoo Band (1996, -0.50 DPS) [world]; Reedknot Ring (9622, -1.07 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (76.8 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.20 DPS) [dungeon]; Searing Golden Blade (12260, -0.67 DPS) [crafted] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.6 spell_power points (1.79 DPS) | yes | Rod of Molten Fire (2565, -0.29 DPS) [world_drop]; Orb of Mystic Insight (249394, -0.40 DPS) [crafted]; Orb of the Forgotten Seer (7685, -0.76 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 346.9 spell_power points (39.97 DPS) | yes | Ragefire Wand (7513, -1.35 DPS, sim-verified) [quest]; Nether Force Wand (11263, -2.34 DPS) [quest]; Icefury Wand (7514, -2.48 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Fiery Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 2050000000000000000-23552100130125031-000000000000000000)

Set DPS (verified): 141.5. Weights run: 0.9s. Verify run: 0.9s. 496 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.876 ± 0.034, crit=0.302 ± 0.028 per rating point (14 rating = 1%, 4.226 per %), hit=0.586 ± 0.055 per rating point (10 rating = 1%, 5.861 per %), spell_haste=not significant (0.716 ± 0.533), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.5 spell_power points (4.90 DPS) | yes | Dreamweave Circlet (10041, -0.91 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.04 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.28 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.0 spell_power points (1.88 DPS) | yes | Horizon Choker (13085, -0.23 DPS) [world_drop]; Mindburst Medallion (11196, -0.37 DPS) [quest]; Triune Amulet (7722, -0.81 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 28.8 spell_power points (3.86 DPS) | yes | Netherflame Shoulders (254053, -0.64 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.16 DPS) [crafted]; Blood Guard's Dreadweave Mantle (220905, -1.16 DPS) [vendor] |
| back | Cindercloth Cloak (14044) | Tailoring [crafted] | 20.0 spell_power points (2.68 DPS) | yes | Deep Woodlands Cloak (19121, -0.02 DPS) [quest]; Spritecaster Cape (11623, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.54 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.5 spell_power points (4.90 DPS) | yes | Robe of the Magi (1716, -1.24 DPS) [world_drop]; Cindercloth Robe (10042, -1.28 DPS) [crafted]; Runecloth Tunic (13857, -1.33 DPS) [crafted] |
| wrist | Netherflame Cuffs (254065) | Tailoring [crafted] | 19.1 spell_power points (2.57 DPS) | yes | Nethergeld Cuffs (254061, -0.80 DPS) [crafted]; Phoenix Bindings (210781, -0.82 DPS) [crafted]; Bloodband Bracers (11469, -0.84 DPS) [quest] |
| hands | Fiery Gloves (254099) | Tailoring [crafted] | 27.8 spell_power points (3.72 DPS) | yes | Dreamweave Gloves (10019, -0.84 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -0.92 DPS) [vendor]; Fiery Handwraps (254025, -1.02 DPS) [crafted] |
| waist | Fiery Waistcord (254085) | Tailoring [crafted] | 26.9 spell_power points (3.61 DPS) | yes | Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Dawnspire Cord (12466, -0.57 DPS) [dungeon]; Fiery Cord (254041, -0.79 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 26.7 spell_power points (3.59 DPS) | yes | Red Mageweave Pants (10009, -0.30 DPS) [crafted]; Flame Leggings (253991, -0.89 DPS) [crafted]; Crimson Silk Pantaloons (7062, -0.99 DPS) [crafted] |
| feet | Fiery Sandals (254111) | Tailoring [crafted] | 27.9 spell_power points (3.74 DPS) | yes | Earthen Silk Slippers (254013, -0.52 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -0.82 DPS) [vendor]; Cindercloth Boots (10044, -0.92 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.93 DPS) | yes | Snake Hoop (6750, -0.54 DPS) [quest]; Reedknot Ring (9622, -0.99 DPS) [quest]; Voodoo Band (1996, -1.11 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.74 DPS) | yes | Snake Hoop (6750, -0.35 DPS) [quest]; Reedknot Ring (9622, -0.80 DPS) [quest]; Voodoo Band (1996, -0.92 DPS) [world] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.25 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.16 DPS) [quest]; Mark of the Chosen (17774, -1.46 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.17 DPS) [quest]; Kindling Stave (11750, -0.66 DPS) [dungeon]; Hanzo Sword (8190, -23.27 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (141.5 DPS) | yes | Pyric Caduceus (11748, -1.52 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.25 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.85 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Cindercloth Cloak; chest: Acumen Robes; wrist: Netherflame Cuffs; hands: Fiery Gloves; waist: Fiery Waistcord; legs: Stone Guard's Dreadweave Leggings; feet: Fiery Sandals; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff; ranged: Noxious Shooter

No-known-source sample (15 of 496, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 2050015100000000000-23552100130125031-003000000000000000)

Set DPS (verified): 392.8. Weights run: 1.0s. Verify run: 2.6s. 1117 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.008, intellect=0.820 ± 0.047, crit=0.503 ± 0.040 per rating point (14 rating = 1%, 7.044 per %), hit=1.296 ± 0.088 per rating point (10 rating = 1%, 12.960 per %), spell_haste=4.060 ± 0.754, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.008

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (392.8 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Magister's Crown (16686, -9.36 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.2 spell_power points (6.24 DPS) | yes | Glowing Eye of Mordresh (10769, -2.05 DPS) [dungeon]; Horizon Choker (13085, -2.69 DPS) [world_drop]; Mindburst Medallion (11196, -2.87 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.7 spell_power points (10.72 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.12 DPS) [pvp]; Rotgrip Mantle (17732, -14.10 DPS, sim-verified) [dungeon] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (6.48 DPS) | yes | Cindercloth Cloak (14044, -0.44 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.50 DPS) [quest]; Spritecaster Cape (11623, -0.64 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.9 spell_power points (17.56 DPS) | yes | Warlord's Silk Raiment (231596, -0.89 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -4.60 DPS) [pvp]; Robes of Fiery Devastation (279266, -10.83 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (392.8 DPS) | yes | General's Silk Cuffs (16538, +0.00 DPS) [pvp]; Netherflame Cuffs (254065, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, +0.00 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.3 spell_power points (9.66 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Fiery Gloves (254099, -8.80 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (392.8 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Fiery Waistcord (254085, -7.64 DPS, sim-verified) [crafted] |
| legs | Cindercloth Pants (14045) | Tailoring [crafted] | 35.8 spell_power points (11.07 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -1.51 DPS) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.8 spell_power points (9.21 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.20 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -0.46 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.7 spell_power points (7.33 DPS) | yes | Band of the Unicorn (7553, -3.32 DPS) [world_drop]; Snake Hoop (6750, -4.21 DPS) [quest]; Glowing Crystal Ring (18402, -4.29 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.1 spell_power points (4.35 DPS) | yes | Band of the Unicorn (7553, -0.34 DPS) [world_drop]; Snake Hoop (6750, -1.23 DPS) [quest]; Glowing Crystal Ring (18402, -1.31 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (392.8 DPS) | yes | Uther's Strength (11302, -0.95 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (392.8 DPS) | yes | Uther's Strength (11302, -7.90 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (392.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.28 DPS) [quest]; Teebu's Blazing Longsword (1728, -61.18 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 254.6 spell_power points (78.61 DPS) | yes | Brilliant Wand (249385, -14.48 DPS) [crafted]; Stormrager (16997, -14.66 DPS) [quest]; Torch of Austen (13004, -21.82 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Cindercloth Pants; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket1: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 2050015100000000000-23452100130135031-003000000000000000)

Set DPS (verified): 730.1. Weights run: 1.1s. Verify run: 2.6s. 1117 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.445 ± 0.027, crit=0.828 ± 0.041 per rating point (14 rating = 1%, 11.587 per %), hit=1.149 ± 0.072 per rating point (10 rating = 1%, 11.488 per %), spell_haste=3.815 ± 0.409, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (730.1 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -10.00 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 16.5 spell_power points (9.51 DPS) | yes | Glowing Eye of Mordresh (10769, -3.41 DPS) [dungeon]; Mindburst Medallion (11196, -4.50 DPS) [quest]; Triune Amulet (7722, -4.89 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.4 spell_power points (19.88 DPS) | yes | Warlord's Silk Amice (231594, -1.56 DPS) [pvp]; Champion's Silk Mantle (227104, -1.67 DPS) [pvp]; Argent Shoulders (19059, -8.33 DPS, sim-verified) [crafted] |
| back | Mageflame Cloak (13007) | World drop [world_drop] | 21.0 spell_power points (12.14 DPS) | yes | Spritecaster Cape (11623, -2.50 DPS) [dungeon]; Cindercloth Cloak (14044, -2.57 DPS) [crafted]; Deep Woodlands Cloak (19121, -2.89 DPS) [quest] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.9 spell_power points (32.92 DPS) | yes | Warlord's Silk Raiment (231596, -2.76 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -9.70 DPS) [pvp]; Robes of Fiery Devastation (279266, -16.87 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (730.1 DPS) | yes | Phoenix Bindings (210781, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Netherflame Cuffs (254065, -12.17 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 36.6 spell_power points (21.18 DPS) | yes | General's Silk Handguards (16540, -2.48 DPS) [vendor]; General's Silk Gauntlets (231599, -2.48 DPS) [vendor]; Fiery Gloves (254099, -13.17 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (730.1 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Fiery Cord (254041, +0.00 DPS) [crafted]; Fiery Waistcord (254085, -8.06 DPS, sim-verified) [crafted] |
| legs | Cindercloth Pants (14045) | Tailoring [crafted] | 31.3 spell_power points (18.12 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Stone Guard's Dreadweave Leggings (220906, -1.40 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.3 spell_power points (14.65 DPS) | yes | Blood Guard's Silk Walkers (227110, +0.00 DPS) [pvp]; General's Silk Boots (231597, +0.00 DPS) [pvp]; Fiery Sandals (254111, -0.77 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.1 spell_power points (12.21 DPS) | yes | Philanthropist's Ring (281635, -5.14 DPS) [quest]; Snake Hoop (6750, -7.45 DPS) [quest]; Reedknot Ring (9622, -8.16 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.52 DPS) | yes | Philanthropist's Ring (281635, -0.45 DPS) [quest]; Snake Hoop (6750, -2.76 DPS) [quest]; Reedknot Ring (9622, -3.47 DPS) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (730.1 DPS) | yes | Uther's Strength (11302, -1.18 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (730.1 DPS) | yes | Uther's Strength (11302, -6.74 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (730.1 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.64 DPS) [quest]; Teebu's Blazing Longsword (1728, -54.76 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 138.7 spell_power points (80.23 DPS) | yes | Brilliant Wand (249385, -13.64 DPS) [crafted]; Stormrager (16997, -16.25 DPS) [quest]; Greater Eternal Wand (249237, -21.70 DPS) [crafted] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Mageflame Cloak; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Cindercloth Pants; feet: Omnicast Boots; finger1: Maiden's Circle; trinket1: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

