# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 28.0. Weights run: 0.9s. Verify run: 0.6s. 157 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.379 ± 0.007, crit=0.086 ± 0.002 per rating point (14 rating = 1%, 1.202 per %), hit=0.281 ± 0.009 per rating point (10 rating = 1%, 2.810 per %), spell_haste=0.915 ± 0.111, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.54 DPS) | yes | Shadow Goggles (4373, -1.15 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (0.76 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.17 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.18 DPS, sim-verified) [crafted]; Sanguine Cape (14376, -0.23 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.63 DPS) | yes | Green Woolen Robe (6243, -0.25 DPS) [crafted]; Green Woolen Vest (2582, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.53 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.9 spell_power points (0.17 DPS) | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.08 DPS) [quest]; Repurposed Hair Band (281256, -0.10 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.16 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.38 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.50 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.23 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.37 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.3 spell_power points (0.75 DPS) | yes | Colorful Kilt (10048, -0.30 DPS) [crafted]; Heavy Woolen Pants (4316, -0.48 DPS) [crafted]; Silk-threaded Trousers (1929, -0.49 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.77 DPS) | yes | Pristine Boots (253889, -0.11 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.41 DPS) [crafted]; Woolen Boots (2583, -0.52 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.52 DPS) | yes | Volcanic Rock Ring (12053, -0.42 DPS) [world_drop]; The 1 Ring (8350, -0.49 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.3 spell_power points (0.21 DPS) | yes | Volcanic Rock Ring (12053, -0.10 DPS) [world_drop]; The 1 Ring (8350, -0.17 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.8 spell_power points (0.34 DPS) | yes | Lesser Staff of the Spire (1300, -0.14 DPS) [world]; Witching Stave (1484, -0.17 DPS) [dungeon]; Staff of Westfall (2042, -0.17 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 248.8 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.57 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.05 DPS) [quest]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 157, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 59.3. Weights run: 0.9s. Verify run: 0.6s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.568 ± 0.010, crit=0.096 ± 0.002 per rating point (14 rating = 1%, 1.347 per %), hit=0.247 ± 0.010 per rating point (10 rating = 1%, 2.468 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.24 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.35 DPS) [crafted]; Enchanter's Cowl (4322, -1.04 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.40 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (1.66 DPS) | yes | Death Speaker Mantle (6685, -0.22 DPS) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Magician's Mantle (12998, -0.47 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.59 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Prelacy Cape (7004, -0.12 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.4 spell_power points (1.93 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Death Speaker Robes (6682, -0.37 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.06 DPS) | yes | Nightsky Wristbands (6407, -0.66 DPS) [world_drop]; Stonecloth Bindings (14416, -0.73 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.93 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (0.82 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.00 DPS) [quest]; Truefaith Gloves (7049, -0.03 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 10.7 spell_power points (1.26 DPS) | yes | Crimson Silk Belt (7055, -0.09 DPS) [crafted]; Invoker's Cord (215366, -0.10 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.32 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Silk-threaded Trousers (1929, -0.47 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.29 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (1.29 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS) [crafted]; Glinteye Slippers (273024, -0.47 DPS) [dungeon]; Acidic Walkers (9454, -1.11 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.04 DPS) | yes | Lavishly Jeweled Ring (1156, -0.64 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon]; Minor Channeling Ring (1449, -1.71 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Minor Channeling Ring (1449, -0.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.24 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.88 DPS) [quest]; Glimmering Staff (249392, -3.50 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 281.5 spell_power points (33.13 DPS) | yes | Wand of Mana Concentration (273043, -1.12 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.54 DPS) [crafted]; Gravestone Scepter (7001, -4.13 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 136.8. Weights run: 0.9s. Verify run: 0.7s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.113 ± 0.014, crit=0.211 ± 0.005 per rating point (14 rating = 1%, 2.958 per %), hit=0.285 ± 0.016 per rating point (10 rating = 1%, 2.849 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.9 spell_power points (5.44 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.19 DPS) [crafted]; Thinking Cap (2624, -1.62 DPS) [world]; Augural Shroud (2620, -2.72 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.9 spell_power points (3.21 DPS) | yes | Necklace of Calisea (1714, -1.64 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.86 DPS) [world]; Triune Amulet (7722, -3.09 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.5 spell_power points (4.34 DPS) | yes | Bloodmage Mantle (7684, -0.50 DPS) [dungeon]; Death Speaker Mantle (6685, -0.65 DPS) [dungeon]; Green Silken Shoulders (7057, -1.98 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.6 spell_power points (2.34 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.54 DPS) [dungeon]; Glyphed Cloak (4732, -0.76 DPS) [world_drop] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.7 spell_power points (5.79 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of Power (7054, -0.27 DPS) [crafted]; Green Silk Armor (7065, -1.05 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 10.0 spell_power points (2.02 DPS) | yes | Arcane Runed Bracers (4744, -0.20 DPS) [quest]; Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Mistscape Bracers (4045, -0.22 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.5 spell_power points (4.53 DPS) | yes | Red Mageweave Gloves (10018, -0.07 DPS) [crafted]; Stormcloth Gloves (10011, -1.03 DPS) [crafted]; Gilded Handwraps (254021, -1.35 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.7 spell_power points (4.79 DPS) | yes | Gilded Cord (254037, -1.37 DPS) [crafted]; Crimson Silk Belt (7055, -2.00 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -2.09 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.4 spell_power points (5.52 DPS) | yes | Crimson Silk Pantaloons (7062, -1.60 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.31 DPS) [dungeon]; Pristine Leggings (253987, -2.54 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.85 DPS) | yes | Gilded Slippers (254001, -1.12 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.04 DPS) [dungeon]; Spidersilk Boots (4320, -2.53 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.6 spell_power points (3.14 DPS) | yes | Ring of Forlorn Spirits (2043, -1.53 DPS) [quest]; Ogremind Ring (1993, -1.57 DPS) [world_drop]; Voodoo Band (1996, -1.57 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.6 spell_power points (2.34 DPS) | yes | Ogremind Ring (1993, -0.76 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world]; Ring of Forlorn Spirits (2043, -0.91 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (136.8 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.45 DPS) [dungeon]; Spine of the Devourer (273040, -3.41 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 17.1 spell_power points (3.46 DPS) | yes | Orb of Mystic Insight (249394, -0.70 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.65 DPS, sim-verified) [dungeon]; Beacon of Hope (9393, -1.84 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 201.6 spell_power points (40.73 DPS) | yes | Earthen Rod (9381, -0.82 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -5.06 DPS) [dungeon]; Twisted Nether Wand (249144, -5.52 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 228.5. Weights run: 0.9s. Verify run: 0.8s. 509 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.455 ± 0.023, crit=0.337 ± 0.007 per rating point (14 rating = 1%, 4.724 per %), hit=0.495 ± 0.028 per rating point (10 rating = 1%, 4.953 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.1 spell_power points (9.36 DPS) | yes | Knight-Lieutenant's Dreadweave Hat (220889, -2.32 DPS) [vendor]; Dreamweave Circlet (10041, -2.37 DPS, sim-verified) [crafted]; Papal Fez (9431, -2.99 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.4 spell_power points (3.96 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -1.10 DPS) [quest]; Necklace of Calisea (1714, -1.98 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.2 spell_power points (7.63 DPS) | yes | Red Mageweave Shoulders (10029, -2.49 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -2.58 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.60 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.7 spell_power points (4.42 DPS) | yes | Runecloth Cloak (13860, -0.41 DPS) [crafted]; Big Voodoo Cloak (8216, -0.90 DPS) [crafted]; Imperial Red Cloak (8248, -1.31 DPS) [world_drop] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.1 spell_power points (9.36 DPS) | yes | Runecloth Tunic (13857, -2.94 DPS) [crafted]; Knight's Dreadweave Vest (220886, -2.99 DPS) [vendor]; Runecloth Robe (13858, -3.12 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 18.1 spell_power points (3.52 DPS) | yes | Forgotten Wraps (9433, -0.12 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.12 DPS) [quest]; Nethergeld Cuffs (254061, -0.18 DPS) [crafted] |
| hands | Sergeant Major's Dreadweave Gloves (220890) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 26.1 spell_power points (5.08 DPS) | yes | Red Mageweave Gloves (10018, -0.11 DPS) [crafted]; Runecloth Gloves (13863, -0.19 DPS) [crafted]; Gilded Gloves (254095, -0.30 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 33.6 spell_power points (6.55 DPS) | yes | Satyrmane Sash (17755, -0.99 DPS) [dungeon]; Runecloth Belt (13856, -1.59 DPS) [crafted]; Deathmage Sash (10771, -1.91 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 34.2 spell_power points (6.65 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.97 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.41 DPS) [crafted]; Red Mageweave Pants (10009, -1.47 DPS, sim-verified) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 26.0 spell_power points (5.07 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Southsea Mojo Boots (20641, -0.40 DPS) [quest]; Earthen Silk Slippers (254013, -0.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.3 spell_power points (3.36 DPS) | yes | Band of the Unicorn (7553, -0.83 DPS) [world_drop]; Ogremind Ring (1993, -1.38 DPS) [world_drop]; Voodoo Band (1996, -1.38 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 13.3 spell_power points (2.58 DPS) | yes | Band of the Unicorn (7553, +0.00 DPS, sim-verified) [world_drop]; Ogremind Ring (1993, -0.60 DPS) [world_drop]; Voodoo Band (1996, -0.60 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (228.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (228.5 DPS) | yes | Mark of the Chosen (17774, -3.79 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (228.5 DPS) | yes | Spellshifter Rod (9527, -0.92 DPS) [quest]; Kindling Stave (11750, -2.87 DPS) [dungeon]; Hammer of the Northern Wind (810, -40.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 269.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.19 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sergeant Major's Dreadweave Gloves; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 509, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 357.7. Weights run: 1.0s. Verify run: 1.6s. 1194 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=1.361 ± 0.021, crit=0.518 ± 0.010 per rating point (14 rating = 1%, 7.247 per %), hit=0.701 ± 0.031 per rating point (10 rating = 1%, 7.010 per %), spell_haste=not significant (-0.004 ± 0.004), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 51.8 spell_power points (15.36 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Red Mageweave Headband (10033, -5.77 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 25.6 spell_power points (7.59 DPS) | yes | Glowing Eye of Mordresh (10769, -2.29 DPS) [dungeon]; Opaline Medallion (16623, -2.35 DPS) [quest]; Horizon Choker (13085, -3.60 DPS, sim-verified) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 41.9 spell_power points (12.43 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Elder Wizard's Mantle (13013, -4.37 DPS, sim-verified) [world_drop] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.2 spell_power points (6.57 DPS) | yes | Big Voodoo Cloak (8216, -1.46 DPS) [crafted]; Imperial Red Cloak (8248, -2.13 DPS) [world_drop]; Runecloth Cloak (13860, -2.26 DPS, sim-verified) [crafted] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 50.1 spell_power points (14.85 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Acumen Robes (17775, -6.37 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 24.6 spell_power points (7.30 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dreadmist Bracers (16703, -1.38 DPS) [dungeon]; Magiskull Cuffs (13107, -3.19 DPS, sim-verified) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.1 spell_power points (11.00 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Magister's Gloves (16684, -5.45 DPS, sim-verified) [dungeon] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 35.6 spell_power points (10.55 DPS) | yes | Marshal's Satin Sash (17609, -1.27 DPS) [pvp]; Devout Belt (16696, -1.29 DPS) [dungeon]; Dawnspire Cord (12466, -5.99 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 35.6 spell_power points (10.55 DPS) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 36.3 spell_power points (10.77 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Wolfrunner Shoes (13101, -1.87 DPS, sim-verified) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 27.5 spell_power points (8.16 DPS) | yes | Glowing Crystal Ring (18402, -3.32 DPS) [quest]; Band of the Hierophant (13096, -4.13 DPS) [world_drop]; Frigid Ring (18679, -4.13 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.8 spell_power points (4.98 DPS) | yes | Band of the Hierophant (13096, -0.95 DPS) [world_drop]; Frigid Ring (18679, -0.95 DPS) [world]; Glowing Crystal Ring (18402, -2.82 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (357.7 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (357.7 DPS) | yes | Mark of the Chosen (17774, -4.36 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (357.7 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -5.40 DPS) [world_drop] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (6.22 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -2.90 DPS, sim-verified) [crafted] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 265.0 spell_power points (78.54 DPS) | yes | Brilliant Wand (249385, -1.11 DPS, sim-verified) [crafted]; Stormrager (16997, -13.83 DPS) [quest]; Torch of Austen (13004, -21.75 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Magister's Robes; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Torch of Light

No-known-source sample (15 of 1194, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 566.9. Weights run: 1.0s. Verify run: 1.7s. 1194 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.161 ± 0.005, crit=0.603 ± 0.011 per rating point (14 rating = 1%, 8.444 per %), hit=0.592 ± 0.026 per rating point (10 rating = 1%, 5.918 per %), spell_haste=0.254 ± 0.062, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.26 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Dreamweave Circlet (10041, -1.79 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.6 spell_power points (5.68 DPS) | yes | Glowing Eye of Mordresh (10769, -2.22 DPS) [dungeon]; Triune Amulet (7722, -2.34 DPS) [dungeon]; Mindburst Medallion (11196, -2.77 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.5 spell_power points (11.48 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Argent Shoulders (19059, -2.26 DPS, sim-verified) [crafted]; Field Marshal's Satin Mantle (231628, -3.47 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (6.24 DPS) | yes | Runecloth Cloak (13860, -2.15 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -2.49 DPS) [dungeon]; Fel Cape (279269, -2.72 DPS) [crafted] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 25.0 spell_power points (10.42 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Robe of the Magi (1716, -11.64 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.6 spell_power points (5.26 DPS) | yes | Spidertank Oilrag (9448, -1.51 DPS) [dungeon]; Arena Wristguards (18709, -1.74 DPS) [world]; Arcane Runed Bracers (4744, -1.76 DPS, sim-verified) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 27.5 spell_power points (11.47 DPS) | yes | Marshal's Satin Gloves (17608, -1.00 DPS) [vendor]; Marshal's Satin Grips (231617, -1.00 DPS) [vendor]; Earth Warder's Gloves (21318, -4.32 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (6.51 DPS) | yes | Ghostweave Cord (254073, -0.67 DPS) [crafted]; Wisdom of the Timbermaw (19047, -0.98 DPS) [crafted]; Star Belt (4329, -1.09 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (10.84 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Knight's Dreadweave Leggings (220888, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.01 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Venomspew Footpads (275606, -0.83 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.1 spell_power points (7.98 DPS) | yes | Philanthropist's Ring (281635, -3.47 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.64 DPS) [quest]; Reedknot Ring (9622, -5.06 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.42 DPS) | yes | Philanthropist's Ring (281635, -0.91 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.08 DPS) [quest]; Reedknot Ring (9622, -2.50 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (566.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (566.9 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (566.9 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -12.55 DPS) [world] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (8.76 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -1.56 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.06 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 190.1 spell_power points (79.26 DPS) | yes | Brilliant Wand (249385, -15.46 DPS) [crafted]; Stormrager (16997, -16.24 DPS) [quest]; Greater Eternal Wand (249237, -22.18 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Truefaith Vestments; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Torch of Light

No-known-source sample (15 of 1194, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 31.6. Weights run: 0.9s. Verify run: 0.6s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.379 ± 0.007, crit=0.086 ± 0.002 per rating point (14 rating = 1%, 1.202 per %), hit=0.281 ± 0.009 per rating point (10 rating = 1%, 2.810 per %), spell_haste=0.915 ± 0.111, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.54 DPS) | yes | Shadow Goggles (4373, -0.96 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (0.76 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.40 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.41 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Pearl-clasped Cloak (5542, -0.08 DPS) [crafted]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Mystic Shawl (3449, -0.18 DPS) [quest] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.63 DPS) | yes | Green Woolen Robe (6243, -0.25 DPS) [crafted]; Green Woolen Vest (2582, -0.26 DPS) [crafted]; Gray Woolen Robe (2585, -0.62 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.27 DPS) | yes | Mindthrust Bracers (1974, -0.10 DPS) [dungeon]; Featherbead Bracers (15452, -0.10 DPS) [quest]; Owlbeard Bracers (16981, -0.11 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Apothecary Gloves (10919, -0.27 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.50 DPS) | yes | Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.23 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.39 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.3 spell_power points (0.75 DPS) | yes | Colorful Kilt (10048, -0.30 DPS) [crafted]; Embroidered Leggings (270011, -0.39 DPS) [quest]; Silk-threaded Trousers (1929, -0.99 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.77 DPS) | yes | Pristine Boots (253889, -0.40 DPS) [crafted]; Red Woolen Boots (4313, -0.41 DPS) [crafted]; Woolen Boots (2583, -0.52 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.3 spell_power points (0.21 DPS) | yes | Volcanic Rock Ring (12053, -0.10 DPS) [world_drop]; The 1 Ring (8350, -0.17 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 1.9 spell_power points (0.17 DPS) | yes | The 1 Ring (8350, -0.14 DPS) [world]; Volcanic Rock Ring (12053, -0.17 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.8 spell_power points (0.34 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Lesser Staff of the Spire (1300, -0.14 DPS) [world]; Clear Crystal Rod (16894, -0.17 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 248.8 spell_power points (22.58 DPS) | yes | Skycaller (12984, -1.27 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.48 DPS) [quest]; Greater Magic Wand (11288, -4.90 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 56.1. Weights run: 0.9s. Verify run: 0.6s. 280 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.568 ± 0.010, crit=0.096 ± 0.002 per rating point (14 rating = 1%, 1.347 per %), hit=0.247 ± 0.010 per rating point (10 rating = 1%, 2.468 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.24 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.35 DPS) [crafted]; Enchanter's Cowl (4322, -0.91 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.40 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (1.66 DPS) | yes | Death Speaker Mantle (6685, -0.22 DPS) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Magician's Mantle (12998, -0.47 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.59 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.15 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.4 spell_power points (1.93 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Death Speaker Robes (6682, -0.37 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.06 DPS) | yes | Nightsky Wristbands (6407, -0.66 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.71 DPS) [quest]; Glowing Magical Bracelets (13106, -0.96 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.8 spell_power points (1.04 DPS) | yes | Truefaith Gloves (7049, -0.25 DPS) [crafted]; Gnoll Casting Gloves (892, -0.33 DPS) [world]; Serpent Gloves (5970, -0.66 DPS, sim-verified) [dungeon] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (1.29 DPS) | yes | Belt of Arugal (6392, +0.00 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.12 DPS) [crafted]; Invoker's Cord (215366, -0.14 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Silk-threaded Trousers (1929, -0.47 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.12 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (1.29 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS) [crafted]; Glinteye Slippers (273024, -0.47 DPS) [dungeon]; Acidic Walkers (9454, -1.32 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.04 DPS) | yes | Solid Gold Band (279845, -0.57 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.64 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.69 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.71 DPS) | yes | Lavishly Jeweled Ring (1156, -0.31 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Solid Gold Band (279845, -0.68 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.24 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -3.50 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -3.57 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 281.5 spell_power points (33.13 DPS) | yes | Wand of Mana Concentration (273043, -1.99 DPS, sim-verified) [world]; Unstable Power Core (279847, -3.42 DPS) [quest]; Greater Mystic Wand (217287, -3.54 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Warsong Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 280, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 120.2. Weights run: 0.9s. Verify run: 0.7s. 382 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.113 ± 0.014, crit=0.211 ± 0.005 per rating point (14 rating = 1%, 2.958 per %), hit=0.285 ± 0.016 per rating point (10 rating = 1%, 2.849 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.9 spell_power points (5.44 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.19 DPS) [crafted]; Thinking Cap (2624, -1.62 DPS) [world]; Augural Shroud (2620, -3.12 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.9 spell_power points (3.21 DPS) | yes | Necklace of Calisea (1714, -1.64 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.86 DPS) [world]; Triune Amulet (7722, -2.79 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.5 spell_power points (4.34 DPS) | yes | Bloodmage Mantle (7684, -0.50 DPS) [dungeon]; Death Speaker Mantle (6685, -0.65 DPS) [dungeon]; Green Silken Shoulders (7057, -1.14 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.6 spell_power points (2.34 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.54 DPS) [dungeon]; Glyphed Cloak (4732, -0.76 DPS) [world_drop] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.7 spell_power points (5.79 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.27 DPS) [crafted]; Green Silk Armor (7065, -1.05 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.9 spell_power points (2.61 DPS) | yes | Windchaser Cuffs (14429, -0.58 DPS) [world_drop]; Spidertank Oilrag (9448, -0.79 DPS) [dungeon]; Mistscape Bracers (4045, -0.81 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.5 spell_power points (4.53 DPS) | yes | Red Mageweave Gloves (10018, -0.07 DPS) [crafted]; Stormcloth Gloves (10011, -1.03 DPS) [crafted]; Gilded Handwraps (254021, -1.35 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.7 spell_power points (4.79 DPS) | yes | Gilded Cord (254037, -1.37 DPS) [crafted]; Crimson Silk Belt (7055, -2.00 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -2.09 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.4 spell_power points (5.52 DPS) | yes | Crimson Silk Pantaloons (7062, -0.99 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.31 DPS) [dungeon]; Pristine Leggings (253987, -2.54 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.85 DPS) | yes | Gilded Slippers (254001, -1.86 DPS) [crafted]; Acidic Walkers (9454, -2.04 DPS) [dungeon]; Spidersilk Boots (4320, -2.53 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.6 spell_power points (3.14 DPS) | yes | Ogremind Ring (1993, -1.57 DPS) [world_drop]; Voodoo Band (1996, -1.57 DPS) [world]; Mindbender Loop (5009, -1.57 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.6 spell_power points (2.34 DPS) | yes | Ogremind Ring (1993, -0.76 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world]; Mindbender Loop (5009, -0.76 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (120.2 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.45 DPS) [dungeon]; Spine of the Devourer (273040, -3.41 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.83 DPS) | yes | Orb of Mystic Insight (249394, -0.07 DPS) [crafted]; Omega Orb (7749, -1.01 DPS) [quest]; Witch's Finger (16887, -1.03 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 201.6 spell_power points (40.73 DPS) | yes | Earthen Rod (9381, -1.94 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -5.06 DPS) [dungeon]; Twisted Nether Wand (249144, -5.52 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 208.4. Weights run: 0.9s. Verify run: 0.9s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.455 ± 0.023, crit=0.337 ± 0.007 per rating point (14 rating = 1%, 4.724 per %), hit=0.495 ± 0.028 per rating point (10 rating = 1%, 4.953 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.1 spell_power points (9.36 DPS) | yes | Blood Guard's Dreadweave Hat (220907, -2.32 DPS) [vendor]; Papal Fez (9431, -2.99 DPS) [dungeon]; Dreamweave Circlet (10041, -3.23 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.4 spell_power points (3.96 DPS) | yes | Glowing Eye of Mordresh (10769, -0.34 DPS) [dungeon]; Mindburst Medallion (11196, -1.10 DPS) [quest]; Necklace of Calisea (1714, -1.98 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.2 spell_power points (7.63 DPS) | yes | Red Mageweave Shoulders (10029, -2.20 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -2.58 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -2.60 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 25.1 spell_power points (4.88 DPS) | yes | Spritecaster Cape (11623, -0.46 DPS) [dungeon]; Runecloth Cloak (13860, -0.87 DPS) [crafted]; Big Voodoo Cloak (8216, -1.36 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.1 spell_power points (9.36 DPS) | yes | Runecloth Robe (13858, -2.74 DPS, sim-verified) [crafted]; Runecloth Tunic (13857, -2.94 DPS) [crafted]; Stone Guard's Dreadweave Vest (220904, -2.99 DPS) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 18.1 spell_power points (3.52 DPS) | yes | Forgotten Wraps (9433, -0.12 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.12 DPS) [quest]; Nethergeld Cuffs (254061, -0.18 DPS) [crafted] |
| hands | First Sergeant's Dreadweave Gloves (220908) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.1 spell_power points (5.08 DPS) | yes | Red Mageweave Gloves (10018, -0.11 DPS) [crafted]; Runecloth Gloves (13863, -0.19 DPS) [crafted]; Gilded Gloves (254095, -0.30 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 33.6 spell_power points (6.55 DPS) | yes | Satyrmane Sash (17755, -0.99 DPS) [dungeon]; Runecloth Belt (13856, -1.59 DPS) [crafted]; Deathmage Sash (10771, -2.52 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 34.2 spell_power points (6.65 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.97 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.41 DPS) [crafted]; Red Mageweave Pants (10009, -1.79 DPS, sim-verified) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.0 spell_power points (5.07 DPS) | yes | Gilded Sandals (254107, +0.00 DPS, sim-verified) [crafted]; Southsea Mojo Boots (20641, -0.40 DPS) [quest]; Earthen Silk Slippers (254013, -0.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.3 spell_power points (3.36 DPS) | yes | Band of the Unicorn (7553, -0.83 DPS) [world_drop]; Ogremind Ring (1993, -1.38 DPS) [world_drop]; Voodoo Band (1996, -1.38 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 13.3 spell_power points (2.58 DPS) | yes | Band of the Unicorn (7553, -0.05 DPS) [world_drop]; Ogremind Ring (1993, -0.60 DPS) [world_drop]; Voodoo Band (1996, -0.60 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.19 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.75 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.92 DPS) [quest]; Kindling Stave (11750, -2.87 DPS) [dungeon]; Hammer of the Northern Wind (810, -37.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (208.4 DPS) | yes | Wand of Allistarj (13065, -2.55 DPS) [world_drop]; Pyric Caduceus (11748, -2.63 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.67 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: First Sergeant's Dreadweave Gloves; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Noxious Shooter

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 356.5. Weights run: 1.0s. Verify run: 1.6s. 1187 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=1.361 ± 0.021, crit=0.518 ± 0.010 per rating point (14 rating = 1%, 7.247 per %), hit=0.701 ± 0.031 per rating point (10 rating = 1%, 7.010 per %), spell_haste=not significant (-0.004 ± 0.004), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 51.8 spell_power points (15.36 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Red Mageweave Headband (10033, -6.37 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 25.6 spell_power points (7.59 DPS) | yes | Glowing Eye of Mordresh (10769, -2.29 DPS) [dungeon]; Horizon Choker (13085, -2.96 DPS, sim-verified) [world_drop]; Mindburst Medallion (11196, -3.39 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 41.9 spell_power points (12.43 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Elder Wizard's Mantle (13013, -3.71 DPS, sim-verified) [world_drop] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 24.3 spell_power points (7.19 DPS) | yes | Spritecaster Cape (11623, -0.62 DPS) [dungeon]; Runecloth Cloak (13860, -1.29 DPS) [crafted]; Big Voodoo Cloak (8216, -2.07 DPS) [crafted] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 50.1 spell_power points (14.85 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Acumen Robes (17775, -7.51 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 24.6 spell_power points (7.30 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dreadmist Bracers (16703, -1.38 DPS) [dungeon]; Magiskull Cuffs (13107, -2.39 DPS, sim-verified) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 37.1 spell_power points (11.00 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Magister's Gloves (16684, -4.36 DPS, sim-verified) [dungeon] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 35.6 spell_power points (10.55 DPS) | yes | General's Satin Cinch (17621, -1.27 DPS) [pvp]; Devout Belt (16696, -1.29 DPS) [dungeon]; Dawnspire Cord (12466, -6.39 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 35.6 spell_power points (10.55 DPS) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 36.3 spell_power points (10.77 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Wolfrunner Shoes (13101, -1.09 DPS, sim-verified) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 27.5 spell_power points (8.16 DPS) | yes | Glowing Crystal Ring (18402, -3.32 DPS) [quest]; Band of the Hierophant (13096, -4.13 DPS) [world_drop]; Frigid Ring (18679, -4.13 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.8 spell_power points (4.98 DPS) | yes | Band of the Hierophant (13096, -0.95 DPS) [world_drop]; Frigid Ring (18679, -0.95 DPS) [world]; Glowing Crystal Ring (18402, -2.06 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (356.5 DPS) | yes | Uther's Strength (11302, -0.09 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (356.5 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (356.5 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -5.40 DPS) [world_drop] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (6.22 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -2.34 DPS, sim-verified) [crafted] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 265.0 spell_power points (78.54 DPS) | yes | Brilliant Wand (249385, -1.42 DPS, sim-verified) [crafted]; Stormrager (16997, -13.83 DPS) [quest]; Torch of Austen (13004, -21.75 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Magister's Robes; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Torch of Light

No-known-source sample (15 of 1187, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 573.0. Weights run: 1.0s. Verify run: 1.7s. 1187 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.161 ± 0.005, crit=0.603 ± 0.011 per rating point (14 rating = 1%, 8.444 per %), hit=0.592 ± 0.026 per rating point (10 rating = 1%, 5.918 per %), spell_haste=0.254 ± 0.062, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (11.26 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Dreamweave Circlet (10041, -1.71 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.6 spell_power points (5.68 DPS) | yes | Glowing Eye of Mordresh (10769, -2.22 DPS) [dungeon]; Triune Amulet (7722, -2.34 DPS) [dungeon]; Mindburst Medallion (11196, -2.77 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 27.5 spell_power points (11.48 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Argent Shoulders (19059, -2.01 DPS, sim-verified) [crafted]; Warlord's Satin Mantle (231631, -3.47 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (6.24 DPS) | yes | Deep Woodlands Cloak (19121, -0.63 DPS) [quest]; Runecloth Cloak (13860, -1.95 DPS) [crafted]; Nightfall Drape (12465, -2.49 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 25.0 spell_power points (10.42 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Robe of the Magi (1716, -11.27 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.6 spell_power points (5.26 DPS) | yes | Spidertank Oilrag (9448, -1.51 DPS) [dungeon]; Arena Wristguards (18709, -1.74 DPS) [world]; Nethergeld Cuffs (254061, -1.87 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 27.5 spell_power points (11.47 DPS) | yes | General's Satin Gloves (17620, -1.00 DPS) [vendor]; General's Satin Grips (231613, -1.00 DPS) [vendor]; Earth Warder's Gloves (21318, -4.45 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.6 spell_power points (6.51 DPS) | yes | Ghostweave Cord (254073, -0.67 DPS) [crafted]; Wisdom of the Timbermaw (19047, -0.98 DPS) [crafted]; Star Belt (4329, -1.09 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (10.84 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS) [vendor]; General's Satin Leggings (231614, +0.00 DPS) [pvp]; Legionnaire's Satin Leggings (227131, -1.08 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (10.01 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Venomspew Footpads (275606, -0.83 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.1 spell_power points (7.98 DPS) | yes | Philanthropist's Ring (281635, -3.47 DPS) [quest]; Reedknot Ring (9622, -5.06 DPS) [quest]; Snake Hoop (6750, -5.14 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.42 DPS) | yes | Philanthropist's Ring (281635, -0.91 DPS) [quest]; Reedknot Ring (9622, -2.50 DPS) [quest]; Snake Hoop (6750, -2.58 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (573.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (573.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (573.0 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -12.55 DPS) [world] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (8.76 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -1.56 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.04 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 190.1 spell_power points (79.26 DPS) | yes | Brilliant Wand (249385, -15.46 DPS) [crafted]; Stormrager (16997, -16.24 DPS) [quest]; Greater Eternal Wand (249237, -22.18 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Spritecaster Cape; chest: Truefaith Vestments; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Torch of Light

No-known-source sample (15 of 1187, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

