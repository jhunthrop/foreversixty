# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 28.4. Weights run: 0.7s. Verify run: 0.5s. 157 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.548 ± 0.009, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.118 per %), hit=0.301 ± 0.009 per rating point (10 rating = 1%, 3.009 per %), spell_haste=not significant (0.255 ± 0.146), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.49 DPS) | yes | Shadow Goggles (4373, -1.12 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.9 spell_power points (0.82 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.49 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Sanguine Cape (14376, -0.15 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.18 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.7 spell_power points (0.64 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.31 DPS) [crafted]; Gray Woolen Robe (2585, -0.51 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.7 spell_power points (0.23 DPS) | yes | Bright Bracers (3647, -0.05 DPS) [world_drop]; Mystic's Bracelets (14366, -0.14 DPS) [world_drop]; Repurposed Hair Band (281256, -0.14 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.58 DPS) | yes | Pristine Gloves (253913, -0.11 DPS) [crafted]; Gnoll Casting Gloves (892, -0.14 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.31 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.2 spell_power points (0.51 DPS) | yes | Keller's Girdle (2911, -0.15 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.36 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 9.3 spell_power points (0.77 DPS) | yes | Colorful Kilt (10048, -0.35 DPS) [crafted]; Darkweave Breeches (12987, -0.45 DPS) [world_drop]; Silk-threaded Trousers (1929, -0.47 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.2 spell_power points (0.76 DPS) | yes | Pristine Boots (253889, -0.10 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Woolen Boots (2583, -0.50 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.1 spell_power points (0.50 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.46 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.3 spell_power points (0.27 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 5.5 spell_power points (0.45 DPS) | yes | Lesser Staff of the Spire (1300, -0.18 DPS) [world]; Witching Stave (1484, -0.23 DPS) [dungeon]; Staff of Westfall (2042, -0.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 273.8 spell_power points (22.56 DPS) | yes | Skycaller (12984, -1.00 DPS) [world_drop]; Deepblaze (279896, -4.03 DPS) [quest]; Sizzle Stick (8071, -4.50 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 157, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 58.6. Weights run: 0.8s. Verify run: 0.5s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.565 ± 0.007, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.399 per %), hit=0.223 ± 0.008 per rating point (10 rating = 1%, 2.232 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Enchanter's Cowl (4322, -1.24 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.39 DPS) | yes | Kaleidoscope Chain (13084, -0.13 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.25 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.26 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Magician's Mantle (12998, -0.46 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.57 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Prelacy Cape (7004, -0.11 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (1.87 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Death Speaker Robes (6682, -0.36 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Nightsky Wristbands (6407, -0.64 DPS) [world_drop]; Stonecloth Bindings (14416, -0.71 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.02 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (0.80 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.00 DPS) [quest]; Truefaith Gloves (7049, -0.03 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 10.7 spell_power points (1.23 DPS) | yes | Crimson Silk Belt (7055, -0.08 DPS) [crafted]; Invoker's Cord (215366, -0.10 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.31 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Silk-threaded Trousers (1929, -0.45 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.21 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (1.26 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Glinteye Slippers (273024, -0.45 DPS) [dungeon]; Acidic Walkers (9454, -1.16 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.01 DPS) | yes | Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Minor Channeling Ring (1449, -1.60 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Minor Channeling Ring (1449, -0.73 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.13 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.83 DPS) [quest]; Glimmering Staff (249392, -3.41 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 289.0 spell_power points (33.12 DPS) | yes | Wand of Mana Concentration (273043, -1.03 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.55 DPS) [crafted]; Gravestone Scepter (7001, -4.12 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 129.4. Weights run: 0.7s. Verify run: 0.6s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.092 ± 0.012, crit=0.014 ± 0.001 per rating point (14 rating = 1%, 0.198 per %), hit=0.274 ± 0.015 per rating point (10 rating = 1%, 2.743 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.6 spell_power points (5.10 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.07 DPS) [crafted]; Thinking Cap (2624, -1.54 DPS) [world]; Augural Shroud (2620, -2.23 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (3.02 DPS) | yes | Necklace of Calisea (1714, -1.55 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.76 DPS) [world]; Triune Amulet (7722, -2.66 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.2 spell_power points (4.07 DPS) | yes | Bloodmage Mantle (7684, -0.45 DPS) [dungeon]; Death Speaker Mantle (6685, -0.61 DPS) [dungeon]; Green Silken Shoulders (7057, -1.69 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.5 spell_power points (2.20 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.52 DPS) [dungeon]; Glyphed Cloak (4732, -0.73 DPS) [world_drop] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.5 spell_power points (5.48 DPS) | yes | Dreamweave Vest (10021, -0.14 DPS) [crafted]; Robe of Power (7054, -0.28 DPS) [crafted]; Green Silk Armor (7065, -1.03 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 9.8 spell_power points (1.89 DPS) | yes | Arcane Runed Bracers (4744, -0.16 DPS) [quest]; Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Mistscape Bracers (4045, -0.21 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.4 spell_power points (4.29 DPS) | yes | Red Mageweave Gloves (10018, -0.67 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -1.01 DPS) [crafted]; Gilded Handwraps (254021, -1.29 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.4 spell_power points (4.49 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Crimson Silk Belt (7055, -1.87 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -1.97 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.1 spell_power points (5.20 DPS) | yes | Crimson Silk Pantaloons (7062, -1.62 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.18 DPS) [dungeon]; Pristine Leggings (253987, -2.39 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.61 DPS) | yes | Gilded Slippers (254001, -1.22 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.97 DPS) [dungeon]; Spidersilk Boots (4320, -2.43 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.97 DPS) | yes | Ring of Forlorn Spirits (2043, -1.43 DPS) [quest]; Ogremind Ring (1993, -1.50 DPS) [world_drop]; Voodoo Band (1996, -1.50 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.5 spell_power points (2.20 DPS) | yes | Ogremind Ring (1993, -0.73 DPS) [world_drop]; Voodoo Band (1996, -0.73 DPS) [world]; Ring of Forlorn Spirits (2043, -0.80 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (129.4 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.42 DPS) [dungeon]; Spine of the Devourer (273040, -3.21 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 16.9 spell_power points (3.25 DPS) | yes | Orb of Mystic Insight (249394, -0.65 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.25 DPS, sim-verified) [dungeon]; Beacon of Hope (9393, -1.71 DPS) [dungeon] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 211.6 spell_power points (40.64 DPS) | yes | Earthen Rod (9381, -1.21 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.97 DPS) [dungeon]; Twisted Nether Wand (249144, -5.49 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 217.3. Weights run: 0.7s. Verify run: 0.8s. 509 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.472 ± 0.019, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.390 per %), hit=0.499 ± 0.026 per rating point (10 rating = 1%, 4.988 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.4 spell_power points (8.64 DPS) | yes | Dreamweave Circlet (10041, -1.98 DPS, sim-verified) [crafted]; Papal Fez (9431, -2.75 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Hat (220889, -2.92 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.6 spell_power points (3.68 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -1.03 DPS) [quest]; Necklace of Calisea (1714, -1.84 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.5 spell_power points (7.05 DPS) | yes | Inquisitor's Shawl (19507, -2.38 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.68 DPS, sim-verified) [crafted]; Green Silken Shoulders (7057, -2.73 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.8 spell_power points (4.07 DPS) | yes | Big Voodoo Cloak (8216, -0.82 DPS) [crafted]; Runecloth Cloak (13860, -1.09 DPS, sim-verified) [crafted]; Imperial Red Cloak (8248, -1.19 DPS) [world_drop] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.4 spell_power points (8.64 DPS) | yes | Runecloth Tunic (13857, -2.72 DPS) [crafted]; Robe of Power (7054, -2.99 DPS) [crafted]; Runecloth Robe (13858, -3.24 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 18.2 spell_power points (3.26 DPS) | yes | Forgotten Wraps (9433, -0.10 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.10 DPS) [quest]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted] |
| hands | Sergeant Major's Dreadweave Gloves (220890) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 26.2 spell_power points (4.68 DPS) | yes | Red Mageweave Gloves (10018, -0.09 DPS) [crafted]; Runecloth Gloves (13863, -0.18 DPS) [crafted]; Gilded Gloves (254095, -0.27 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 34.0 spell_power points (6.06 DPS) | yes | Satyrmane Sash (17755, -0.94 DPS) [dungeon]; Runecloth Belt (13856, -1.48 DPS) [crafted]; Deathmage Sash (10771, -1.98 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 31.7 spell_power points (5.65 DPS) | yes | Knight's Dreadweave Leggings (220888, -0.29 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.39 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.81 DPS) [crafted] |
| feet | Gilded Sandals (254107) | Tailoring [crafted] | sim-verified (+2.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Southsea Mojo Boots (20641, -0.01 DPS) [quest]; Earthen Silk Slippers (254013, -0.04 DPS) [crafted]; Sergeant Major's Dreadweave Boots (220891, -2.41 DPS, sim-verified) [vendor] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.4 spell_power points (3.10 DPS) | yes | Band of the Unicorn (7553, -0.78 DPS) [world_drop]; Ogremind Ring (1993, -1.26 DPS) [world_drop]; Voodoo Band (1996, -1.26 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 13.4 spell_power points (2.38 DPS) | yes | Band of the Unicorn (7553, -0.06 DPS) [world_drop]; Ogremind Ring (1993, -0.55 DPS) [world_drop]; Voodoo Band (1996, -0.55 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -3.41 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.86 DPS) [quest]; Kindling Stave (11750, -2.68 DPS) [dungeon]; Hammer of the Northern Wind (810, -35.91 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Wand of Allistarj (13065, -2.47 DPS) [world_drop]; Pyric Caduceus (11748, -3.38 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.71 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sergeant Major's Dreadweave Gloves; waist: Dawnspire Cord; feet: Gilded Sandals; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Noxious Shooter

No-known-source sample (15 of 509, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 339.5. Weights run: 0.8s. Verify run: 2.1s. 1194 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.292 ± 0.018, crit=0.027 ± 0.002 per rating point (14 rating = 1%, 0.374 per %), hit=0.581 ± 0.029 per rating point (10 rating = 1%, 5.812 per %), spell_haste=not significant (-0.005 ± 0.013), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 49.8 spell_power points (13.98 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Red Mageweave Headband (10033, -3.93 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 24.9 spell_power points (7.00 DPS) | yes | Glowing Eye of Mordresh (10769, -2.13 DPS) [dungeon]; Opaline Medallion (16623, -2.28 DPS) [quest]; Horizon Choker (13085, -3.76 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (339.5 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Elder Wizard's Mantle (13013, -5.32 DPS, sim-verified) [world_drop] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.8 spell_power points (6.11 DPS) | yes | Big Voodoo Cloak (8216, -1.44 DPS) [crafted]; Runecloth Cloak (13860, -1.84 DPS, sim-verified) [crafted]; Imperial Red Cloak (8248, -2.12 DPS) [world_drop] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 48.2 spell_power points (13.54 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Acumen Robes (17775, -4.58 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 23.9 spell_power points (6.72 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Dreadmist Bracers (16703, -1.32 DPS) [dungeon]; Magiskull Cuffs (13107, -3.43 DPS, sim-verified) [world_drop] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | 25.1 spell_power points (7.05 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -5.01 DPS, sim-verified) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (339.5 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Marshal's Satin Sash (17609, -0.03 DPS) [pvp]; Magister's Belt (16685, -6.62 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (339.5 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Magister's Leggings (16687, -7.29 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 35.5 spell_power points (9.97 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Wolfrunner Shoes (13101, -2.09 DPS, sim-verified) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 27.0 spell_power points (7.60 DPS) | yes | Glowing Crystal Ring (18402, -3.24 DPS) [quest]; Band of the Unicorn (7553, -3.95 DPS) [world_drop]; Frigid Ring (18679, -3.97 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.5 spell_power points (4.62 DPS) | yes | Band of the Unicorn (7553, -0.97 DPS) [world_drop]; Frigid Ring (18679, -0.99 DPS) [world]; Glowing Crystal Ring (18402, -3.09 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (339.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (339.5 DPS) | yes | Mark of the Chosen (17774, -4.04 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (339.5 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -5.43 DPS) [world_drop] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (5.90 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -3.45 DPS, sim-verified) [crafted] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (339.5 DPS) | yes | Stormrager (16997, -0.14 DPS) [quest]; Torch of Light (279246, -6.82 DPS, sim-verified) [crafted]; Torch of Austen (13004, -7.86 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Magister's Robes; wrist: Runecloth Cuffs; hands: Magister's Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Brilliant Wand

No-known-source sample (15 of 1194, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 507.1. Weights run: 0.9s. Verify run: 1.5s. 1194 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.013 ± 0.002, crit=0.053 ± 0.004 per rating point (14 rating = 1%, 0.738 per %), hit=0.419 ± 0.033 per rating point (10 rating = 1%, 4.191 per %), spell_haste=-8.308 ± 0.357, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.06 DPS) | yes | Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor]; Dreamweave Circlet (10041, -2.10 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.1 spell_power points (4.52 DPS) | yes | Triune Amulet (7722, -1.46 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -1.87 DPS) [dungeon]; Mindburst Medallion (11196, -2.25 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.31 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -2.61 DPS, sim-verified) [crafted]; Field Marshal's Satin Mantle (231628, -3.25 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.1 spell_power points (5.24 DPS) | yes | Runecloth Cloak (13860, -1.87 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.89 DPS) [dungeon]; Icy Cloak (4327, -2.64 DPS) [crafted] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 25.0 spell_power points (9.31 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Robe of the Magi (1716, -9.99 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.1 spell_power points (4.15 DPS) | yes | Arcane Runed Bracers (4744, -0.79 DPS) [quest]; Spidertank Oilrag (9448, -0.79 DPS) [dungeon]; Earthenweave Cuffs (254131, -1.17 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.19 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Dreamweave Gloves (10019, -1.37 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.1 spell_power points (5.26 DPS) | yes | Ghostweave Cord (254073, -0.05 DPS) [crafted]; Star Belt (4329, -0.42 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.10 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.68 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Brightcloth Pants (14104, -2.73 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.94 DPS) | yes | Marshal's Satin Sandals (17607, -0.28 DPS) [vendor]; Marshal's Satin Treads (231620, -0.28 DPS) [vendor]; Venomspew Footpads (275606, -0.74 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.1 spell_power points (6.74 DPS) | yes | Philanthropist's Ring (281635, -2.99 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.76 DPS) [quest]; Reedknot Ring (9622, -4.13 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.84 DPS) | yes | Philanthropist's Ring (281635, -1.06 DPS, sim-verified) [quest]; Ring of Forlorn Spirits (2043, -1.86 DPS) [quest]; Reedknot Ring (9622, -2.23 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, -2.75 DPS) [pvp]; Spire of Hakkar (10844, -10.49 DPS) [world] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (7.82 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -2.17 DPS) [pvp]; Orb of the Forgotten Seer (7685, -2.55 DPS, sim-verified) [dungeon] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (507.1 DPS) | yes | Stormrager (16997, -0.25 DPS) [quest]; Torch of Austen (13004, -6.18 DPS) [world_drop]; Torch of Light (279246, -11.67 DPS, sim-verified) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Truefaith Vestments; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Brilliant Wand

No-known-source sample (15 of 1194, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 33.2. Weights run: 0.7s. Verify run: 0.6s. 150 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.548 ± 0.009, crit=0.080 ± 0.002 per rating point (14 rating = 1%, 1.118 per %), hit=0.301 ± 0.009 per rating point (10 rating = 1%, 3.009 per %), spell_haste=not significant (0.255 ± 0.146), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.49 DPS) | yes | Shadow Goggles (4373, -0.83 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.9 spell_power points (0.82 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.49 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.03 DPS) [crafted]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Sanguine Cape (14376, -0.15 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.7 spell_power points (0.64 DPS) | yes | Green Woolen Robe (6243, -0.26 DPS) [crafted]; Green Woolen Vest (2582, -0.31 DPS) [crafted]; Gray Woolen Robe (2585, -0.67 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.25 DPS) | yes | Mindthrust Bracers (1974, -0.02 DPS) [dungeon]; Featherbead Bracers (15452, -0.02 DPS) [quest]; Bright Bracers (3647, -0.07 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.58 DPS) | yes | Gnoll Casting Gloves (892, -0.08 DPS) [world]; Pristine Gloves (253913, -0.11 DPS) [crafted]; Apothecary Gloves (10919, -0.25 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.2 spell_power points (0.51 DPS) | yes | Keller's Girdle (2911, -0.15 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.21 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.36 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 9.3 spell_power points (0.77 DPS) | yes | Colorful Kilt (10048, -0.35 DPS) [crafted]; Embroidered Leggings (270011, -0.44 DPS) [quest]; Silk-threaded Trousers (1929, -0.93 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.2 spell_power points (0.76 DPS) | yes | Pristine Boots (253889, -0.37 DPS) [crafted]; Red Woolen Boots (4313, -0.43 DPS) [crafted]; Woolen Boots (2583, -0.50 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.3 spell_power points (0.27 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.23 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.7 spell_power points (0.23 DPS) | yes | Volcanic Rock Ring (12053, +0.00 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 5.5 spell_power points (0.45 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Lesser Staff of the Spire (1300, -0.18 DPS) [world]; Clear Crystal Rod (16894, -0.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 273.8 spell_power points (22.56 DPS) | yes | Skycaller (12984, -1.93 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.50 DPS) [quest]; Greater Magic Wand (11288, -4.89 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 150, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 56.6. Weights run: 0.8s. Verify run: 0.6s. 280 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.565 ± 0.007, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.399 per %), hit=0.223 ± 0.008 per rating point (10 rating = 1%, 2.232 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.34 DPS) [crafted]; Enchanter's Cowl (4322, -1.12 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.39 DPS) | yes | Crystal Starfire Medallion (5003, -0.13 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.13 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (1.61 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Magician's Mantle (12998, -0.46 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.57 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.15 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (1.87 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Death Speaker Robes (6682, -0.36 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Nightsky Wristbands (6407, -0.64 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.69 DPS) [quest]; Glowing Magical Bracelets (13106, -0.83 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.8 spell_power points (1.01 DPS) | yes | Truefaith Gloves (7049, -0.24 DPS) [crafted]; Gnoll Casting Gloves (892, -0.32 DPS) [world]; Serpent Gloves (5970, -0.37 DPS, sim-verified) [dungeon] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (1.26 DPS) | yes | Belt of Arugal (6392, -0.03 DPS) [dungeon]; Crimson Silk Belt (7055, -0.12 DPS) [crafted]; Invoker's Cord (215366, -0.13 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.18 DPS) [crafted]; Silk-threaded Trousers (1929, -0.45 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.93 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (1.26 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS) [crafted]; Glinteye Slippers (273024, -0.45 DPS) [dungeon]; Acidic Walkers (9454, -1.07 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.01 DPS) | yes | Solid Gold Band (279845, -0.55 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.62 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.69 DPS) | yes | Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Solid Gold Band (279845, -0.51 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.13 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -3.41 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -3.48 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 289.0 spell_power points (33.12 DPS) | yes | Wand of Mana Concentration (273043, -1.90 DPS, sim-verified) [world]; Unstable Power Core (279847, -3.41 DPS) [quest]; Greater Mystic Wand (217287, -3.55 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Warsong Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 280, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 115.3. Weights run: 0.7s. Verify run: 0.6s. 382 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.092 ± 0.012, crit=0.014 ± 0.001 per rating point (14 rating = 1%, 0.198 per %), hit=0.274 ± 0.015 per rating point (10 rating = 1%, 2.743 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 26.6 spell_power points (5.10 DPS) | yes | Spellpower Goggles Xtreme (10502, -1.07 DPS) [crafted]; Thinking Cap (2624, -1.54 DPS) [world]; Augural Shroud (2620, -2.62 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (3.02 DPS) | yes | Necklace of Calisea (1714, -1.55 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.76 DPS) [world]; Triune Amulet (7722, -1.96 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.2 spell_power points (4.07 DPS) | yes | Bloodmage Mantle (7684, -0.45 DPS) [dungeon]; Death Speaker Mantle (6685, -0.61 DPS) [dungeon]; Green Silken Shoulders (7057, -0.92 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.5 spell_power points (2.20 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.52 DPS) [dungeon]; Glyphed Cloak (4732, -0.73 DPS) [world_drop] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.5 spell_power points (5.48 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.28 DPS) [crafted]; Green Silk Armor (7065, -1.03 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.7 spell_power points (2.44 DPS) | yes | Windchaser Cuffs (14429, -0.56 DPS) [world_drop]; Spidertank Oilrag (9448, -0.72 DPS) [dungeon]; Mistscape Bracers (4045, -0.77 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.4 spell_power points (4.29 DPS) | yes | Red Mageweave Gloves (10018, -0.09 DPS) [crafted]; Stormcloth Gloves (10011, -1.01 DPS) [crafted]; Gilded Handwraps (254021, -1.29 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 23.4 spell_power points (4.49 DPS) | yes | Gilded Cord (254037, +0.00 DPS, sim-verified) [crafted]; Crimson Silk Belt (7055, -1.87 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -1.97 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.1 spell_power points (5.20 DPS) | yes | Crimson Silk Pantaloons (7062, -1.03 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.18 DPS) [dungeon]; Pristine Leggings (253987, -2.39 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.61 DPS) | yes | Gilded Slippers (254001, -1.80 DPS) [crafted]; Acidic Walkers (9454, -1.97 DPS) [dungeon]; Spidersilk Boots (4320, -2.43 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.5 spell_power points (2.97 DPS) | yes | Ogremind Ring (1993, -1.50 DPS) [world_drop]; Voodoo Band (1996, -1.50 DPS) [world]; Mindbender Loop (5009, -1.50 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 11.5 spell_power points (2.20 DPS) | yes | Ogremind Ring (1993, -0.73 DPS) [world_drop]; Voodoo Band (1996, -0.73 DPS) [world]; Mindbender Loop (5009, -0.73 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (115.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.42 DPS) [dungeon]; Spine of the Devourer (273040, -3.21 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.69 DPS) | yes | Orb of Mystic Insight (249394, -0.09 DPS) [crafted]; Omega Orb (7749, -0.96 DPS) [quest]; Witch's Finger (16887, -1.01 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 211.6 spell_power points (40.64 DPS) | yes | Earthen Rod (9381, -2.02 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -4.97 DPS) [dungeon]; Twisted Nether Wand (249144, -5.49 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Jaina's Firestarter

No-known-source sample (15 of 382, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 197.3. Weights run: 0.7s. Verify run: 0.8s. 501 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.472 ± 0.019, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.390 per %), hit=0.499 ± 0.026 per rating point (10 rating = 1%, 4.988 per %), spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 48.4 spell_power points (8.64 DPS) | yes | Papal Fez (9431, -2.75 DPS) [dungeon]; Dreamweave Circlet (10041, -2.85 DPS, sim-verified) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.92 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 20.6 spell_power points (3.68 DPS) | yes | Glowing Eye of Mordresh (10769, -0.33 DPS) [dungeon]; Mindburst Medallion (11196, -1.03 DPS) [quest]; Necklace of Calisea (1714, -1.84 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 39.5 spell_power points (7.05 DPS) | yes | Red Mageweave Shoulders (10029, -1.82 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -2.38 DPS) [dungeon]; Green Silken Shoulders (7057, -2.73 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 25.2 spell_power points (4.51 DPS) | yes | Spritecaster Cape (11623, -0.43 DPS) [dungeon]; Runecloth Cloak (13860, -0.80 DPS) [crafted]; Big Voodoo Cloak (8216, -1.25 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 48.4 spell_power points (8.64 DPS) | yes | Runecloth Robe (13858, -2.31 DPS, sim-verified) [crafted]; Runecloth Tunic (13857, -2.72 DPS) [crafted]; Robe of Power (7054, -2.99 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 18.2 spell_power points (3.26 DPS) | yes | Forgotten Wraps (9433, -0.10 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.10 DPS) [quest]; Nethergeld Cuffs (254061, -0.17 DPS) [crafted] |
| hands | First Sergeant's Dreadweave Gloves (220908) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.2 spell_power points (4.68 DPS) | yes | Red Mageweave Gloves (10018, -0.09 DPS) [crafted]; Runecloth Gloves (13863, -0.18 DPS) [crafted]; Gilded Gloves (254095, -0.27 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 34.0 spell_power points (6.06 DPS) | yes | Satyrmane Sash (17755, -0.94 DPS) [dungeon]; Runecloth Belt (13856, -1.48 DPS) [crafted]; Deathmage Sash (10771, -2.62 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 31.7 spell_power points (5.65 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, -0.29 DPS) [vendor]; Kilt of the Atal'ai Prophet (10807, -0.39 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.81 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.2 spell_power points (4.68 DPS) | yes | Gilded Sandals (254107, -0.35 DPS) [crafted]; Southsea Mojo Boots (20641, -0.36 DPS) [quest]; Earthen Silk Slippers (254013, -0.40 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.4 spell_power points (3.10 DPS) | yes | Band of the Unicorn (7553, -0.78 DPS) [world_drop]; Ogremind Ring (1993, -1.26 DPS) [world_drop]; Voodoo Band (1996, -1.26 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 13.4 spell_power points (2.38 DPS) | yes | Band of the Unicorn (7553, -0.06 DPS) [world_drop]; Ogremind Ring (1993, -0.55 DPS) [world_drop]; Voodoo Band (1996, -0.55 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.18 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -1.51 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.86 DPS) [quest]; Kindling Stave (11750, -2.68 DPS) [dungeon]; Hammer of the Northern Wind (810, -34.99 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (197.3 DPS) | yes | Wand of Allistarj (13065, -2.47 DPS) [world_drop]; Pyric Caduceus (11748, -3.01 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.71 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: First Sergeant's Dreadweave Gloves; waist: Dawnspire Cord; feet: First Sergeant's Dreadweave Boots; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Noxious Shooter

No-known-source sample (15 of 501, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 344.6. Weights run: 0.8s. Verify run: 2.1s. 1187 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.292 ± 0.018, crit=0.027 ± 0.002 per rating point (14 rating = 1%, 0.374 per %), hit=0.581 ± 0.029 per rating point (10 rating = 1%, 5.812 per %), spell_haste=not significant (-0.005 ± 0.013), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 49.8 spell_power points (13.98 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Red Mageweave Headband (10033, -4.83 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 24.9 spell_power points (7.00 DPS) | yes | Glowing Eye of Mordresh (10769, -2.13 DPS) [dungeon]; Mindburst Medallion (11196, -3.14 DPS) [quest]; Horizon Choker (13085, -3.27 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (344.6 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Elder Wizard's Mantle (13013, -5.95 DPS, sim-verified) [world_drop] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 23.6 spell_power points (6.64 DPS) | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -1.21 DPS) [crafted]; Big Voodoo Cloak (8216, -1.97 DPS) [crafted] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | 48.2 spell_power points (13.54 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Acumen Robes (17775, -5.35 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 23.9 spell_power points (6.72 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Dreadmist Bracers (16703, -1.32 DPS) [dungeon]; Magiskull Cuffs (13107, -2.75 DPS, sim-verified) [world_drop] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | 25.1 spell_power points (7.05 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -5.78 DPS, sim-verified) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (344.6 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; General's Satin Cinch (17621, -0.03 DPS) [pvp]; Magister's Belt (16685, -7.06 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (344.6 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Magister's Leggings (16687, -7.50 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 35.5 spell_power points (9.97 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Wolfrunner Shoes (13101, -1.36 DPS, sim-verified) [world_drop] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 27.0 spell_power points (7.60 DPS) | yes | Glowing Crystal Ring (18402, -3.24 DPS) [quest]; Band of the Unicorn (7553, -3.95 DPS) [world_drop]; Frigid Ring (18679, -3.97 DPS) [world] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.5 spell_power points (4.62 DPS) | yes | Band of the Unicorn (7553, -0.97 DPS) [world_drop]; Frigid Ring (18679, -0.99 DPS) [world]; Glowing Crystal Ring (18402, -2.45 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (344.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (344.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (344.6 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -5.43 DPS) [world_drop] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (5.90 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -2.70 DPS, sim-verified) [crafted] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (344.6 DPS) | yes | Stormrager (16997, -0.14 DPS) [quest]; Torch of Light (279246, -7.22 DPS, sim-verified) [crafted]; Torch of Austen (13004, -7.86 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Magister's Robes; wrist: Runecloth Cuffs; hands: Magister's Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Brilliant Wand

No-known-source sample (15 of 1187, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 518.2. Weights run: 0.9s. Verify run: 1.6s. 1187 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.013 ± 0.002, crit=0.053 ± 0.004 per rating point (14 rating = 1%, 0.738 per %), hit=0.419 ± 0.033 per rating point (10 rating = 1%, 4.191 per %), spell_haste=-8.308 ± 0.357, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (10.06 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Dreamweave Circlet (10041, -2.11 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 12.1 spell_power points (4.52 DPS) | yes | Triune Amulet (7722, -1.56 DPS, sim-verified) [dungeon]; Glowing Eye of Mordresh (10769, -1.87 DPS) [dungeon]; Mindburst Medallion (11196, -2.25 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (9.31 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -2.25 DPS, sim-verified) [crafted]; Warlord's Satin Mantle (231631, -3.25 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.1 spell_power points (5.24 DPS) | yes | Deep Woodlands Cloak (19121, -0.73 DPS) [quest]; Runecloth Cloak (13860, -1.85 DPS) [crafted]; Nightfall Drape (12465, -1.89 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 25.0 spell_power points (9.31 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, -1.35 DPS) [vendor]; Robe of the Magi (1716, -10.18 DPS, sim-verified) [world_drop] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 11.1 spell_power points (4.15 DPS) | yes | Spidertank Oilrag (9448, -0.79 DPS) [dungeon]; Earthenweave Cuffs (254131, -1.17 DPS) [crafted]; Nethergeld Cuffs (254061, -1.50 DPS) [crafted] |
| hands | Earth Warder's Gloves (21318) | Winterfall Activity [quest] | 22.0 spell_power points (8.19 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Dreamweave Gloves (10019, -1.46 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.1 spell_power points (5.26 DPS) | yes | Ghostweave Cord (254073, -0.05 DPS) [crafted]; Star Belt (4329, -0.42 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.10 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (9.68 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; Legionnaire's Satin Leggings (227131, -1.79 DPS) [vendor]; Brightcloth Pants (14104, -2.73 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (8.94 DPS) | yes | General's Satin Boots (17618, -0.28 DPS) [vendor]; General's Satin Treads (231610, -0.28 DPS) [vendor]; Venomspew Footpads (275606, -0.74 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.1 spell_power points (6.74 DPS) | yes | Philanthropist's Ring (281635, -2.99 DPS) [quest]; Reedknot Ring (9622, -4.13 DPS) [quest]; Snake Hoop (6750, -4.48 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (4.84 DPS) | yes | Philanthropist's Ring (281635, -1.08 DPS, sim-verified) [quest]; Reedknot Ring (9622, -2.23 DPS) [quest]; Snake Hoop (6750, -2.58 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, -2.75 DPS) [pvp]; Spire of Hakkar (10844, -10.49 DPS) [world] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (7.82 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -2.17 DPS) [pvp]; Orb of the Forgotten Seer (7685, -2.56 DPS, sim-verified) [dungeon] |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | sim-verified (518.2 DPS) | yes | Stormrager (16997, -0.25 DPS) [quest]; Torch of Austen (13004, -6.18 DPS) [world_drop]; Torch of Light (279246, -12.50 DPS, sim-verified) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Argent Shoulders; back: Spritecaster Cape; chest: Truefaith Vestments; wrist: Runecloth Cuffs; hands: Earth Warder's Gloves; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Brilliant Wand

No-known-source sample (15 of 1187, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

