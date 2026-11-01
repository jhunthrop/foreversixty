# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 29.6. Weights run: 1.1s. Verify run: 0.8s. 163 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.349 ± 0.010, crit=0.092 ± 0.002 per rating point (14 rating = 1%, 1.289 per %), hit=0.309 ± 0.009 per rating point (10 rating = 1%, 3.090 per %), spell_haste=not significant (0.462 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.51 DPS) | yes | Shadow Goggles (4373, -0.36 DPS) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.69 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.15 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.35 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.34 DPS) | yes | Pearl-clasped Cloak (5542, -0.08 DPS) [crafted]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Sanguine Cape (14376, -0.22 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.7 spell_power points (0.57 DPS) | yes | Green Woolen Robe (6243, -0.23 DPS) [crafted]; Green Woolen Vest (2582, -0.23 DPS) [crafted]; Gray Woolen Robe (2585, -0.70 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.7 spell_power points (0.15 DPS) | yes | Windsong Bangles (263336, -0.06 DPS) [quest]; Repurposed Hair Band (281256, -0.09 DPS) [quest]; Bright Bracers (3647, -0.45 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.59 DPS) | yes | Gnoll Casting Gloves (892, -0.16 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.36 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.46 DPS) | yes | Novice Arcanist's Sash (253885, -0.03 DPS) [crafted]; Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.22 DPS) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.1 spell_power points (0.68 DPS) | yes | Colorful Kilt (10048, -0.26 DPS) [crafted]; Heavy Woolen Pants (4316, -0.43 DPS) [crafted]; Silk-threaded Trousers (1929, -1.56 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.71 DPS) | yes | Pristine Boots (253889, +0.00 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Woolen Boots (2583, -0.48 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.7 spell_power points (0.48 DPS) | yes | Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; The 1 Ring (8350, -0.45 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.1 spell_power points (0.18 DPS) | yes | The 1 Ring (8350, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.30 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.5 spell_power points (0.29 DPS) | yes | Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Witching Stave (1484, -0.15 DPS) [dungeon]; Staff of Westfall (2042, -0.15 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-verified (29.6 DPS) | yes | Sable Wand (7607, -0.09 DPS) [quest]; Sizzle Stick (8071, -0.39 DPS, sim-verified) [quest]; Skycaller (12984, -1.00 DPS) [world_drop] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 163, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 53.8. Weights run: 1.2s. Verify run: 0.8s. 294 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.709 ± 0.011, crit=0.108 ± 0.002 per rating point (14 rating = 1%, 1.507 per %), hit=0.315 ± 0.012 per rating point (10 rating = 1%, 3.154 per %), spell_haste=3.404 ± 0.165, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.20 DPS) [crafted]; Nightsky Cowl (4039, -0.25 DPS) [world_drop]; Enchanter's Cowl (4322, -1.20 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.3 spell_power points (0.43 DPS) | yes | Crystal Starfire Medallion (5003, +0.00 DPS, sim-verified) [world_drop]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.4 spell_power points (1.55 DPS) | yes | Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Fairywing Mantle (9536, -0.30 DPS) [quest]; Magician's Mantle (12998, -0.40 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.50 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Resilient Cape (14400, -0.07 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.09 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.2 spell_power points (1.84 DPS) | yes | Beguiler Robes (7728, +0.00 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Death Speaker Robes (6682, -0.34 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.91 DPS) | yes | Nightsky Wristbands (6407, -0.48 DPS) [world_drop]; Stonecloth Bindings (14416, -0.55 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.82 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.0 spell_power points (0.80 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.10 DPS) [dungeon]; Shilly Mitts (9609, -0.10 DPS) [quest] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.1 spell_power points (1.12 DPS) | yes | Invoker's Cord (215366, -0.06 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.32 DPS) [dungeon]; Crimson Silk Belt (7055, -0.34 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.17 DPS) [crafted]; Necromancer Leggings (2277, -0.42 DPS) [world_drop]; Gaze Dreamer Pants (6903, -0.84 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.0 spell_power points (1.21 DPS) | yes | Spidersilk Boots (4320, -0.21 DPS) [crafted]; Glinteye Slippers (273024, -0.50 DPS) [dungeon]; Acidic Walkers (9454, -0.80 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Azora's Will (4999, -0.60 DPS) [world_drop]; Minor Channeling Ring (1449, -1.35 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; Azora's Will (4999, -0.25 DPS) [world_drop]; Minor Channeling Ring (1449, -0.79 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.63 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.61 DPS) [quest]; Glimmering Staff (249392, -2.84 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 328.8 spell_power points (33.14 DPS) | yes | Wand of Mana Concentration (273043, -1.58 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -2.79 DPS) [dungeon]; Greater Mystic Wand (217287, -3.64 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 126.2. Weights run: 1.2s. Verify run: 0.9s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.083 ± 0.021, crit=0.241 ± 0.006 per rating point (14 rating = 1%, 3.376 per %), hit=0.370 ± 0.020 per rating point (10 rating = 1%, 3.695 per %), spell_haste=0.608 ± 0.119, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.4 spell_power points (4.38 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.90 DPS) [crafted]; Thinking Cap (2624, -1.33 DPS) [world]; Augural Shroud (2620, -3.56 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (2.60 DPS) | yes | Necklace of Calisea (1714, -1.34 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.52 DPS) [world]; Triune Amulet (7722, -3.02 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.1 spell_power points (3.50 DPS) | yes | Bloodmage Mantle (7684, -0.39 DPS) [dungeon]; Death Speaker Mantle (6685, -0.53 DPS) [dungeon]; Green Silken Shoulders (7057, -2.09 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.4 spell_power points (1.89 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.46 DPS) [dungeon]; Glyphed Cloak (4732, -0.64 DPS) [world_drop] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (126.2 DPS) | yes | Robe of Power (7054, -0.12 DPS) [crafted]; Green Silk Armor (7065, -0.77 DPS) [crafted]; Robe of the Magi (1716, -1.66 DPS, sim-verified) [world_drop] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 9.7 spell_power points (1.62 DPS) | yes | Spidertank Oilrag (9448, -0.12 DPS) [dungeon]; Mistscape Bracers (4045, -0.18 DPS) [dungeon]; Arcane Runed Bracers (4744, -1.56 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.3 spell_power points (3.71 DPS) | yes | Red Mageweave Gloves (10018, +0.00 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.89 DPS) [crafted]; Gilded Handwraps (254021, -1.12 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.2 spell_power points (3.86 DPS) | yes | Gilded Cord (254037, -1.09 DPS) [crafted]; Crimson Silk Belt (7055, -1.60 DPS) [crafted]; Star Belt (4329, -1.70 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.0 spell_power points (4.48 DPS) | yes | Crimson Silk Pantaloons (7062, -1.12 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.88 DPS) [dungeon]; Pristine Leggings (253987, -2.06 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.98 DPS) | yes | Gilded Slippers (254001, -1.56 DPS) [crafted]; Acidic Walkers (9454, -1.72 DPS) [dungeon]; Spidersilk Boots (4320, -2.10 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.56 DPS) | yes | Ring of Forlorn Spirits (2043, -1.23 DPS) [quest]; Ogremind Ring (1993, -1.30 DPS) [world_drop]; Voodoo Band (1996, -1.30 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.4 spell_power points (1.89 DPS) | yes | Ogremind Ring (1993, -0.64 DPS) [world_drop]; Voodoo Band (1996, -0.64 DPS) [world]; Ring of Forlorn Spirits (2043, -1.15 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.36 DPS) [dungeon]; Spine of the Devourer (273040, -2.76 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 16.8 spell_power points (2.79 DPS) | yes | Orb of Mystic Insight (249394, -0.55 DPS) [crafted]; Beacon of Hope (9393, -1.46 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -3.32 DPS, sim-verified) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 243.8 spell_power points (40.45 DPS) | yes | Earthen Rod (9381, -2.88 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.79 DPS) [dungeon]; Twisted Nether Wand (249144, -5.46 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 205.8. Weights run: 1.1s. Verify run: 1.1s. 515 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.180 ± 0.030, crit=0.375 ± 0.010 per rating point (14 rating = 1%, 5.251 per %), hit=0.613 ± 0.031 per rating point (10 rating = 1%, 6.128 per %), spell_haste=7.330 ± 0.408, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 42.6 spell_power points (6.88 DPS) | yes | Knight-Lieutenant's Dreadweave Hat (220889, -1.48 DPS) [vendor]; Dreamweave Circlet (10041, -2.01 DPS, sim-verified) [crafted]; Papal Fez (9431, -2.35 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 16.5 spell_power points (2.67 DPS) | yes | Glowing Eye of Mordresh (10769, -0.01 DPS) [dungeon]; Mindburst Medallion (11196, -0.56 DPS) [quest]; Necklace of Calisea (1714, -1.33 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 34.2 spell_power points (5.53 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -1.67 DPS) [vendor]; Inquisitor's Shawl (19507, -1.92 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.00 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.1 spell_power points (3.40 DPS) | yes | Runecloth Cloak (13860, -0.43 DPS) [crafted]; Big Voodoo Cloak (8216, -0.88 DPS) [crafted]; Imperial Red Cloak (8248, -1.31 DPS) [world_drop] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.6 spell_power points (6.88 DPS) | yes | Knight's Dreadweave Vest (220886, -2.00 DPS) [vendor]; Runecloth Tunic (13857, -2.04 DPS) [crafted]; Runecloth Robe (13858, -2.49 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 15.6 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Forgotten Wraps (9433, -0.24 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.24 DPS) [quest] |
| hands | Sergeant Major's Dreadweave Gloves (220890) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 23.6 spell_power points (3.81 DPS) | yes | Red Mageweave Gloves (10018, -0.13 DPS) [crafted]; Dreamweave Gloves (10019, -0.15 DPS) [crafted]; Runecloth Gloves (13863, -0.16 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 28.4 spell_power points (4.59 DPS) | yes | Deathmage Sash (10771, -0.60 DPS) [dungeon]; Runecloth Belt (13856, -1.01 DPS) [crafted]; Satyrmane Sash (17755, -1.80 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 31.4 spell_power points (5.07 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.16 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.30 DPS) [crafted]; Red Mageweave Pants (10009, -1.38 DPS, sim-verified) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 24.8 spell_power points (3.99 DPS) | yes | Earthen Silk Slippers (254013, -0.12 DPS) [crafted]; Gilded Sandals (254107, -0.50 DPS) [crafted]; Southsea Mojo Boots (20641, -0.61 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.9 spell_power points (2.57 DPS) | yes | Snake Hoop (6750, -0.65 DPS) [quest]; Ogremind Ring (1993, -1.23 DPS) [world_drop]; Voodoo Band (1996, -1.23 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.10 DPS) | yes | Snake Hoop (6750, -0.18 DPS) [quest]; Ogremind Ring (1993, -0.76 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.93 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.50 DPS) [quest]; Kindling Stave (11750, -1.62 DPS) [dungeon]; Hammer of the Northern Wind (810, -35.33 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (205.8 DPS) | yes | Pyric Caduceus (11748, -2.27 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.39 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.77 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sergeant Major's Dreadweave Gloves; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Noxious Shooter

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 366.2. Weights run: 1.2s. Verify run: 2.9s. 1200 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.558 ± 0.034, crit=0.586 ± 0.013 per rating point (14 rating = 1%, 8.206 per %), hit=0.969 ± 0.044 per rating point (10 rating = 1%, 9.694 per %), spell_haste=not significant (0.839 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 57.8 spell_power points (13.00 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Red Mageweave Headband (10033, -1.94 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 27.6 spell_power points (6.21 DPS) | yes | Opaline Medallion (16623, -1.65 DPS) [quest]; Horizon Choker (13085, -1.78 DPS, sim-verified) [world_drop]; Glowing Eye of Mordresh (10769, -1.83 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (366.2 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Magister's Mantle (16689, -11.54 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.4 spell_power points (5.26 DPS) | yes | Big Voodoo Cloak (8216, -0.97 DPS) [crafted]; Imperial Red Cloak (8248, -1.40 DPS) [world_drop]; Runecloth Cloak (13860, -3.25 DPS, sim-verified) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (366.2 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -9.85 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (366.2 DPS) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, -10.45 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (366.2 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Magister's Gloves (16684, -8.75 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (366.2 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Magister's Belt (16685, -11.88 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (366.2 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Magister's Leggings (16687, -9.08 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 38.7 spell_power points (8.71 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -1.93 DPS, sim-verified) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 28.9 spell_power points (6.51 DPS) | yes | Philanthropist's Ring (281635, -2.50 DPS) [quest]; Band of the Hierophant (13096, -3.00 DPS) [world_drop]; Frigid Ring (18679, -3.00 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 18.7 spell_power points (4.21 DPS) | yes | Philanthropist's Ring (281635, -0.20 DPS) [quest]; Band of the Hierophant (13096, -0.70 DPS) [world_drop]; Frigid Ring (18679, -0.70 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (366.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (366.2 DPS) | yes | Mark of the Chosen (17774, -4.23 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (366.2 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -3.39 DPS) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 22.5 spell_power points (5.06 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Elunarian Sphere (15968, -0.15 DPS) [world_drop] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 347.1 spell_power points (78.12 DPS) | yes | Stormrager (16997, -5.05 DPS, sim-verified) [quest]; Brilliant Wand (249385, -14.05 DPS) [crafted]; Torch of Austen (13004, -21.33 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Glowing Crystal Ring; main_hand: Hand of Edward the Odd; off_hand: Truesilver Conduit; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 667.7. Weights run: 1.3s. Verify run: 2.1s. 1200 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.173 ± 0.005, crit=0.666 ± 0.013 per rating point (14 rating = 1%, 9.321 per %), hit=0.803 ± 0.036 per rating point (10 rating = 1%, 8.033 per %), spell_haste=0.795 ± 0.185, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.43 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Dreamweave Circlet (10041, -3.25 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (6.32 DPS) | yes | Glowing Eye of Mordresh (10769, -2.46 DPS) [dungeon]; Triune Amulet (7722, -2.64 DPS) [dungeon]; Mindburst Medallion (11196, -3.08 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.6 spell_power points (13.15 DPS) | yes | Field Marshal's Satin Epaulets (231621, -0.05 DPS) [vendor]; Argent Shoulders (19059, -4.12 DPS, sim-verified) [crafted]; Field Marshal's Satin Mantle (231628, -4.19 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (6.92 DPS) | yes | Runecloth Cloak (13860, -2.32 DPS, sim-verified) [crafted]; Fel Cape (279269, -2.63 DPS) [crafted]; Nightfall Drape (12465, -2.78 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 25.0 spell_power points (11.51 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Knight's Dreadweave Vest (220886, -11.24 DPS, sim-verified) [vendor] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (5.86 DPS) | yes | Arena Wristguards (18709, -1.57 DPS) [world]; Arcane Runed Bracers (4744, -1.72 DPS) [quest]; Spidertank Oilrag (9448, -1.72 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.4 spell_power points (13.52 DPS) | yes | Marshal's Satin Gloves (17608, -1.90 DPS) [vendor]; Marshal's Satin Grips (231617, -1.90 DPS) [vendor]; Earth Warder's Gloves (21318, -5.55 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.7 spell_power points (7.24 DPS) | yes | Ghostweave Cord (254073, -0.79 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.06 DPS) [crafted]; Star Belt (4329, -1.25 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.97 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Knight's Dreadweave Leggings (220888, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.05 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Omnicast Boots (11822, -0.89 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (8.84 DPS) | yes | Philanthropist's Ring (281635, -3.84 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.16 DPS) [quest]; Reedknot Ring (9622, -5.62 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.98 DPS) | yes | Philanthropist's Ring (281635, -0.98 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.30 DPS) [quest]; Reedknot Ring (9622, -2.76 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (667.7 DPS) | yes | Uther's Strength (11302, -0.57 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (667.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (667.7 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -13.89 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.67 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -1.65 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.38 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 172.8 spell_power points (79.53 DPS) | yes | Brilliant Wand (249385, -15.24 DPS) [crafted]; Stormrager (16997, -16.44 DPS) [quest]; Greater Eternal Wand (249237, -22.05 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Truefaith Vestments; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; trinket2: Mark of the Chosen; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 31.6. Weights run: 1.1s. Verify run: 0.8s. 156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.349 ± 0.010, crit=0.092 ± 0.002 per rating point (14 rating = 1%, 1.289 per %), hit=0.309 ± 0.009 per rating point (10 rating = 1%, 3.090 per %), spell_haste=not significant (0.462 ± 0.119), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.51 DPS) | yes | Shadow Goggles (4373, -0.87 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.1 spell_power points (0.69 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.15 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.35 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.34 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Mystic Shawl (3449, -0.17 DPS) [quest]; Pearl-clasped Cloak (5542, -0.45 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.7 spell_power points (0.57 DPS) | yes | Green Woolen Robe (6243, -0.23 DPS) [crafted]; Green Woolen Vest (2582, -0.23 DPS) [crafted]; Gray Woolen Robe (2585, -0.51 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.25 DPS) | yes | Mindthrust Bracers (1974, -0.11 DPS) [dungeon]; Featherbead Bracers (15452, -0.11 DPS) [quest]; Owlbeard Bracers (16981, -0.11 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.59 DPS) | yes | Gnoll Casting Gloves (892, -0.08 DPS) [world]; Pristine Gloves (253913, -0.17 DPS) [crafted]; Apothecary Gloves (10919, -0.25 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.4 spell_power points (0.46 DPS) | yes | Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.22 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.72 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.1 spell_power points (0.68 DPS) | yes | Colorful Kilt (10048, -0.26 DPS) [crafted]; Embroidered Leggings (270011, -0.35 DPS) [quest]; Silk-threaded Trousers (1929, -0.40 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.4 spell_power points (0.71 DPS) | yes | Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.43 DPS, sim-verified) [crafted]; Woolen Boots (2583, -0.48 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.1 spell_power points (0.18 DPS) | yes | Volcanic Rock Ring (12053, -0.09 DPS) [world_drop]; The 1 Ring (8350, -0.15 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 1.7 spell_power points (0.15 DPS) | yes | Volcanic Rock Ring (12053, -0.06 DPS) [world_drop]; The 1 Ring (8350, -0.12 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.5 spell_power points (0.29 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Clear Crystal Rod (16894, -0.15 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | sim-verified (31.6 DPS) | yes | Sizzle Stick (8071, -0.87 DPS, sim-verified) [quest]; Skycaller (12984, -1.00 DPS) [world_drop]; Greater Magic Wand (11288, -4.89 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 156, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 55.2. Weights run: 1.2s. Verify run: 0.8s. 286 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.709 ± 0.011, crit=0.108 ± 0.002 per rating point (14 rating = 1%, 1.507 per %), hit=0.315 ± 0.012 per rating point (10 rating = 1%, 3.154 per %), spell_haste=3.404 ± 0.165, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.20 DPS) [crafted]; Nightsky Cowl (4039, -0.25 DPS) [world_drop]; Enchanter's Cowl (4322, -1.18 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.3 spell_power points (0.43 DPS) | yes | Crystal Starfire Medallion (5003, -0.14 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.14 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.4 spell_power points (1.55 DPS) | yes | Death Speaker Mantle (6685, -0.16 DPS) [dungeon]; Fairywing Mantle (9536, -0.30 DPS) [quest]; Magician's Mantle (12998, -0.40 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.50 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Resilient Cape (14400, -0.07 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.09 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.2 spell_power points (1.84 DPS) | yes | Beguiler Robes (7728, +0.00 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Death Speaker Robes (6682, -0.34 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.91 DPS) | yes | Nightsky Wristbands (6407, -0.48 DPS) [world_drop]; Featherbead Bracers (15452, -0.55 DPS) [quest]; Glowing Magical Bracelets (13106, -0.96 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.5 spell_power points (0.96 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.26 DPS) [dungeon]; Pristine Gloves (253913, -0.34 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.1 spell_power points (1.12 DPS) | yes | Warsong Sash (16975, -0.01 DPS) [quest]; Crimson Silk Belt (7055, -0.02 DPS) [crafted]; Invoker's Cord (215366, -0.06 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.17 DPS) [crafted]; Necromancer Leggings (2277, -0.42 DPS) [world_drop]; Gaze Dreamer Pants (6903, -1.06 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.0 spell_power points (1.21 DPS) | yes | Spidersilk Boots (4320, -0.21 DPS) [crafted]; Glinteye Slippers (273024, -0.50 DPS) [dungeon]; Acidic Walkers (9454, -1.17 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (0.96 DPS) | yes | Lavishly Jeweled Ring (1156, -0.53 DPS) [dungeon]; Solid Gold Band (279845, -0.56 DPS) [quest]; Azora's Will (4999, -0.60 DPS) [world_drop] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.60 DPS) | yes | Solid Gold Band (279845, -0.20 DPS) [quest]; Azora's Will (4999, -0.25 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.00 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (3.63 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -2.84 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -2.91 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 328.8 spell_power points (33.14 DPS) | yes | Wand of Mana Concentration (273043, -2.03 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -2.79 DPS) [dungeon]; Unstable Power Core (279847, -3.44 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 118.9. Weights run: 1.2s. Verify run: 1.0s. 388 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.083 ± 0.021, crit=0.241 ± 0.006 per rating point (14 rating = 1%, 3.376 per %), hit=0.370 ± 0.020 per rating point (10 rating = 1%, 3.695 per %), spell_haste=0.608 ± 0.119, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.4 spell_power points (4.38 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.90 DPS) [crafted]; Thinking Cap (2624, -1.33 DPS) [world]; Augural Shroud (2620, -2.35 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (2.60 DPS) | yes | Triune Amulet (7722, -1.19 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.34 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.52 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.1 spell_power points (3.50 DPS) | yes | Green Silken Shoulders (7057, -0.19 DPS) [crafted]; Bloodmage Mantle (7684, -0.39 DPS) [dungeon]; Death Speaker Mantle (6685, -0.53 DPS) [dungeon] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.4 spell_power points (1.89 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.46 DPS) [dungeon]; Glyphed Cloak (4732, -0.64 DPS) [world_drop] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.12 DPS) [crafted]; Green Silk Armor (7065, -0.77 DPS) [crafted]; Robe of the Magi (1716, -1.52 DPS, sim-verified) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.7 spell_power points (2.10 DPS) | yes | Windchaser Cuffs (14429, +0.00 DPS, sim-verified) [world_drop]; Spidertank Oilrag (9448, -0.61 DPS) [dungeon]; Mistscape Bracers (4045, -0.66 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.3 spell_power points (3.71 DPS) | yes | Red Mageweave Gloves (10018, -0.08 DPS) [crafted]; Stormcloth Gloves (10011, -0.89 DPS) [crafted]; Gilded Handwraps (254021, -1.12 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Crimson Silk Belt (7055, -0.51 DPS) [crafted]; Star Belt (4329, -0.61 DPS) [crafted]; Deathmage Sash (10771, -1.42 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.0 spell_power points (4.48 DPS) | yes | Crimson Silk Pantaloons (7062, -0.82 DPS) [crafted]; Stoneweaver Leggings (9407, -1.88 DPS) [dungeon]; Pristine Leggings (253987, -2.06 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.98 DPS) | yes | Gilded Slippers (254001, -1.56 DPS) [crafted]; Acidic Walkers (9454, -1.72 DPS) [dungeon]; Spidersilk Boots (4320, -2.10 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.56 DPS) | yes | Ogremind Ring (1993, -1.30 DPS) [world_drop]; Voodoo Band (1996, -1.30 DPS) [world]; Mindbender Loop (5009, -1.30 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.4 spell_power points (1.89 DPS) | yes | Ogremind Ring (1993, -0.64 DPS) [world_drop]; Voodoo Band (1996, -0.64 DPS) [world]; Mindbender Loop (5009, -0.64 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.36 DPS) [dungeon]; Spine of the Devourer (273040, -2.76 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.32 DPS) | yes | Orb of Mystic Insight (249394, +0.00 DPS, sim-verified) [crafted]; Omega Orb (7749, -0.83 DPS) [quest]; Witch's Finger (16887, -0.89 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 243.8 spell_power points (40.45 DPS) | yes | Earthen Rod (9381, -1.24 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.79 DPS) [dungeon]; Twisted Nether Wand (249144, -5.46 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 197.7. Weights run: 1.1s. Verify run: 1.1s. 507 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.180 ± 0.030, crit=0.375 ± 0.010 per rating point (14 rating = 1%, 5.251 per %), hit=0.613 ± 0.031 per rating point (10 rating = 1%, 6.128 per %), spell_haste=7.330 ± 0.408, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 42.6 spell_power points (6.88 DPS) | yes | Blood Guard's Dreadweave Hat (220907, -1.48 DPS) [vendor]; Papal Fez (9431, -2.35 DPS) [dungeon]; Dreamweave Circlet (10041, -4.15 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 16.5 spell_power points (2.67 DPS) | yes | Glowing Eye of Mordresh (10769, -0.01 DPS) [dungeon]; Mindburst Medallion (11196, -0.56 DPS) [quest]; Necklace of Calisea (1714, -1.33 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 34.2 spell_power points (5.53 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.67 DPS) [vendor]; Inquisitor's Shawl (19507, -1.92 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.38 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.6 spell_power points (3.65 DPS) | yes | Spritecaster Cape (11623, -0.25 DPS) [dungeon]; Runecloth Cloak (13860, -0.67 DPS) [crafted]; Big Voodoo Cloak (8216, -1.13 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 42.6 spell_power points (6.88 DPS) | yes | Stone Guard's Dreadweave Vest (220904, -2.00 DPS) [vendor]; Runecloth Tunic (13857, -2.04 DPS) [crafted]; Runecloth Robe (13858, -3.84 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 15.6 spell_power points (2.52 DPS) | yes | Nethergeld Cuffs (254061, -0.06 DPS) [crafted]; Forgotten Wraps (9433, -0.24 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.24 DPS) [quest] |
| hands | First Sergeant's Dreadweave Gloves (220908) | PvP rank 9 · First Sergeant · Horde [vendor] | 23.6 spell_power points (3.81 DPS) | yes | Red Mageweave Gloves (10018, -0.13 DPS) [crafted]; Dreamweave Gloves (10019, -0.15 DPS) [crafted]; Runecloth Gloves (13863, -0.16 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 28.4 spell_power points (4.59 DPS) | yes | Deathmage Sash (10771, -0.60 DPS) [dungeon]; Runecloth Belt (13856, -1.01 DPS) [crafted]; Satyrmane Sash (17755, -3.01 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 31.4 spell_power points (5.07 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -1.16 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.30 DPS) [crafted]; Red Mageweave Pants (10009, -1.78 DPS, sim-verified) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 24.8 spell_power points (3.99 DPS) | yes | Gilded Sandals (254107, -0.50 DPS) [crafted]; Southsea Mojo Boots (20641, -0.61 DPS) [quest]; Earthen Silk Slippers (254013, -2.37 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.9 spell_power points (2.57 DPS) | yes | Snake Hoop (6750, -0.65 DPS) [quest]; Ogremind Ring (1993, -1.23 DPS) [world_drop]; Voodoo Band (1996, -1.23 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.10 DPS) | yes | Snake Hoop (6750, -0.18 DPS) [quest]; Ogremind Ring (1993, -0.76 DPS) [world_drop]; Voodoo Band (1996, -0.76 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.50 DPS) [quest]; Kindling Stave (11750, -1.62 DPS) [dungeon]; Hammer of the Northern Wind (810, -33.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (197.7 DPS) | yes | Pyric Caduceus (11748, -2.33 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.39 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.77 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: First Sergeant's Dreadweave Gloves; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Noxious Shooter

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 355.0. Weights run: 1.2s. Verify run: 3.0s. 1193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.558 ± 0.034, crit=0.586 ± 0.013 per rating point (14 rating = 1%, 8.206 per %), hit=0.969 ± 0.044 per rating point (10 rating = 1%, 9.694 per %), spell_haste=not significant (0.839 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 57.8 spell_power points (13.00 DPS) | yes | Red Mageweave Headband (10033, +0.00 DPS) [crafted]; Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 27.6 spell_power points (6.21 DPS) | yes | Horizon Choker (13085, -1.30 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.83 DPS) [dungeon]; Mindburst Medallion (11196, -2.75 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (355.0 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Magister's Mantle (16689, -10.44 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 26.0 spell_power points (5.86 DPS) | yes | Spritecaster Cape (11623, -0.60 DPS) [dungeon]; Runecloth Cloak (13860, -1.03 DPS) [crafted]; Big Voodoo Cloak (8216, -1.58 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (355.0 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -8.35 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (355.0 DPS) | yes | Magiskull Cuffs (13107, +0.00 DPS) [world_drop]; General's Satin Bracers (17619, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, -9.55 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (355.0 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Magister's Gloves (16684, -7.73 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (355.0 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; General's Satin Cinch (17621, +0.00 DPS) [pvp]; Magister's Belt (16685, -10.76 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (355.0 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Magister's Leggings (16687, -7.88 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 38.7 spell_power points (8.71 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; First Sergeant's Dreadweave Boots (220909, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 28.9 spell_power points (6.51 DPS) | yes | Philanthropist's Ring (281635, -2.50 DPS) [quest]; Band of the Hierophant (13096, -3.00 DPS) [world_drop]; Frigid Ring (18679, -3.00 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 18.7 spell_power points (4.21 DPS) | yes | Philanthropist's Ring (281635, -0.20 DPS) [quest]; Band of the Hierophant (13096, -0.70 DPS) [world_drop]; Frigid Ring (18679, -0.70 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (355.0 DPS) | yes | Uther's Strength (11302, -0.61 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (355.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (355.0 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -3.39 DPS) [world_drop] |
| off_hand | Truesilver Conduit (249455) | Enchanting [crafted] | 22.5 spell_power points (5.06 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Elunarian Sphere (15968, -0.15 DPS) [world_drop] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 347.1 spell_power points (78.12 DPS) | yes | Stormrager (16997, -4.41 DPS, sim-verified) [quest]; Brilliant Wand (249385, -14.05 DPS) [crafted]; Torch of Austen (13004, -21.33 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Glowing Crystal Ring; main_hand: Hand of Edward the Odd; off_hand: Truesilver Conduit; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 669.1. Weights run: 1.3s. Verify run: 2.1s. 1193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.173 ± 0.005, crit=0.666 ± 0.013 per rating point (14 rating = 1%, 9.321 per %), hit=0.803 ± 0.036 per rating point (10 rating = 1%, 8.033 per %), spell_haste=0.795 ± 0.185, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.43 DPS) | yes | Dreamweave Circlet (10041, +0.00 DPS) [crafted]; Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (6.32 DPS) | yes | Glowing Eye of Mordresh (10769, -2.46 DPS) [dungeon]; Triune Amulet (7722, -2.64 DPS) [dungeon]; Mindburst Medallion (11196, -3.08 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 28.6 spell_power points (13.15 DPS) | yes | Warlord's Satin Epaulets (231611, -0.05 DPS) [vendor]; Argent Shoulders (19059, -2.82 DPS, sim-verified) [crafted]; Warlord's Satin Mantle (231631, -4.19 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (6.92 DPS) | yes | Deep Woodlands Cloak (19121, -0.68 DPS) [quest]; Runecloth Cloak (13860, -2.14 DPS) [crafted]; Fel Cape (279269, -2.63 DPS) [crafted] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 25.0 spell_power points (11.51 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Stone Guard's Dreadweave Vest (220904, -10.81 DPS, sim-verified) [vendor] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (5.86 DPS) | yes | Arena Wristguards (18709, -1.57 DPS) [world]; Spidertank Oilrag (9448, -1.72 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.08 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.4 spell_power points (13.52 DPS) | yes | General's Satin Gloves (17620, -1.90 DPS) [vendor]; General's Satin Grips (231613, -1.90 DPS) [vendor]; Earth Warder's Gloves (21318, -5.65 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 15.7 spell_power points (7.24 DPS) | yes | Ghostweave Cord (254073, -0.79 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.06 DPS) [crafted]; Star Belt (4329, -1.25 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.97 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS) [vendor]; General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, -1.09 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.05 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Omnicast Boots (11822, -0.89 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (8.84 DPS) | yes | Philanthropist's Ring (281635, -3.84 DPS) [quest]; Reedknot Ring (9622, -5.62 DPS) [quest]; Snake Hoop (6750, -5.68 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.98 DPS) | yes | Philanthropist's Ring (281635, -0.98 DPS) [quest]; Reedknot Ring (9622, -2.76 DPS) [quest]; Snake Hoop (6750, -2.83 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (669.1 DPS) | yes | Uther's Strength (11302, -0.57 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (669.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (669.1 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -13.89 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.67 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -1.65 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.35 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 172.8 spell_power points (79.53 DPS) | yes | Brilliant Wand (249385, -15.24 DPS) [crafted]; Stormrager (16997, -16.44 DPS) [quest]; Greater Eternal Wand (249237, -22.05 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Spritecaster Cape; chest: Truefaith Vestments; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

