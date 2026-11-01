# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 32.5. Weights run: 1.0s. Verify run: 0.7s. 160 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.377 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.006 per %), hit=0.223 ± 0.007 per rating point (10 rating = 1%, 2.229 per %), spell_haste=not significant (0.176 ± 0.108), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.65 DPS) | yes | Shadow Goggles (4373, -1.26 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (0.91 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.57 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.44 DPS) | yes | Pearl-clasped Cloak (5542, -0.09 DPS) [crafted]; Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Sanguine Cape (14376, -0.27 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.75 DPS) | yes | Green Woolen Robe (6243, -0.30 DPS) [crafted]; Green Woolen Vest (2582, -0.31 DPS) [crafted]; Gray Woolen Robe (2585, -0.79 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.9 spell_power points (0.21 DPS) | yes | Windsong Bangles (263336, -0.10 DPS) [quest]; Repurposed Hair Band (281256, -0.12 DPS) [quest]; Bright Bracers (3647, -0.74 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.76 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.23 DPS, sim-verified) [world]; Heavy Woolen Gloves (4310, -0.46 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.60 DPS) | yes | Novice Ardent's Sash (253887, -0.26 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.83 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.3 spell_power points (0.90 DPS) | yes | Colorful Kilt (10048, -0.36 DPS) [crafted]; Heavy Woolen Pants (4316, -0.57 DPS) [crafted]; Silk-threaded Trousers (1929, -1.41 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.93 DPS) | yes | Pristine Boots (253889, -0.42 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.49 DPS) [crafted]; Woolen Boots (2583, -0.63 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.63 DPS) | yes | Volcanic Rock Ring (12053, -0.50 DPS) [world_drop]; The 1 Ring (8350, -0.59 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.3 spell_power points (0.25 DPS) | yes | The 1 Ring (8350, -0.21 DPS) [world]; Volcanic Rock Ring (12053, -0.37 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.8 spell_power points (0.41 DPS) | yes | Witching Stave (1484, -0.21 DPS) [dungeon]; Staff of Westfall (2042, -0.21 DPS) [quest]; Lesser Staff of the Spire (1300, -0.61 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 207.8 spell_power points (22.64 DPS) | yes | Skycaller (12984, -1.18 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.11 DPS) [quest]; Sizzle Stick (8071, -4.44 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 160, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 63.4. Weights run: 1.0s. Verify run: 0.7s. 298 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.386 ± 0.009, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.589 per %), hit=0.230 ± 0.010 per rating point (10 rating = 1%, 2.300 per %), spell_haste=not significant (0.509 ± 0.155), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.68 DPS) | yes | Silk Headband (7050, -0.31 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.46 DPS) [crafted]; Enchanter's Cowl (4322, -2.17 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.3 spell_power points (0.35 DPS) | yes | Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.12 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (1.91 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Invoker's Mantle (215365, -0.54 DPS) [crafted] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.76 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Prelacy Cape (7004, -0.15 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.0 spell_power points (2.14 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Tree Bark Jacket (1486, -0.16 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.21 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.37 DPS) | yes | Nightsky Wristbands (6407, -1.02 DPS) [world_drop]; Stonecloth Bindings (14416, -1.08 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.79 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.07 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Truefaith Gloves (7049, -0.13 DPS) [crafted]; Gnoll Casting Gloves (892, -0.15 DPS) [world] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 10.2 spell_power points (1.55 DPS) | yes | Crimson Silk Belt (7055, -0.22 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.33 DPS) [dungeon]; Invoker's Cord (215366, -0.46 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (63.4 DPS) | yes | Filigreed Pristine Leggings (253937, -0.21 DPS) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.84 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (1.48 DPS) | yes | Acidic Walkers (9454, -0.25 DPS) [dungeon]; Glinteye Slippers (273024, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -2.24 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.9 spell_power points (1.21 DPS) | yes | Minor Channeling Ring (1449, -0.33 DPS) [quest]; Electrocutioner Lagnut (9447, -0.75 DPS) [dungeon]; Runemender's Seal (273031, -0.75 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.92 DPS) | yes | Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Runemender's Seal (273031, -0.46 DPS) [world]; Minor Channeling Ring (1449, -1.05 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (5.50 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -2.44 DPS) [quest]; Glimmering Staff (249392, -4.85 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 216.8 spell_power points (33.10 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.33 DPS) [crafted]; Gravestone Scepter (7001, -4.10 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 000000000000000000-00000000000000000-253510130100301231)

Set DPS (verified): 99.2. Weights run: 1.2s. Verify run: 0.9s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.746 ± 0.024, crit=0.178 ± 0.007 per rating point (14 rating = 1%, 2.486 per %), hit=0.362 ± 0.025 per rating point (10 rating = 1%, 3.616 per %), spell_haste=not significant (1.421 ± 0.373), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+3.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.35 DPS) [world]; Whitemane's Chapeau (7720, -0.97 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -3.05 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.0 spell_power points (2.03 DPS) | yes | Necklace of Calisea (1714, -1.21 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.33 DPS) [world]; Triune Amulet (7722, -1.72 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.7 spell_power points (2.61 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.35 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.7 spell_power points (1.52 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.43 DPS) [crafted]; Blackforge Cape (6424, -0.59 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.5 spell_power points (4.14 DPS) | yes | Dreamweave Vest (10021, -0.28 DPS) [crafted]; Robe of Power (7054, -0.55 DPS) [crafted]; Elemental Raiment (9434, -0.86 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Condor Bracers (15864, -0.31 DPS) [quest]; Windchaser Cuffs (14429, -0.36 DPS) [world_drop]; Arcane Runed Bracers (4744, -1.36 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.0 spell_power points (3.28 DPS) | yes | Red Mageweave Gloves (10018, -0.39 DPS) [crafted]; Black Mageweave Gloves (10003, -0.94 DPS) [crafted]; Gilded Handwraps (254021, -1.21 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.2 spell_power points (2.84 DPS) | yes | Gilded Cord (254037, -0.66 DPS) [crafted]; Star Belt (4329, -0.81 DPS) [crafted]; Belt of Arugal (6392, -1.09 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.0 spell_power points (3.59 DPS) | yes | Crimson Silk Pantaloons (7062, -0.87 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.56 DPS) [dungeon]; Pristine Leggings (253987, -1.68 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.75 DPS) | yes | Gilded Slippers (254001, -1.84 DPS) [crafted]; Acidic Walkers (9454, -2.04 DPS) [dungeon]; Spidersilk Boots (4320, -2.19 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (2.15 DPS) | yes | Ring of Forlorn Spirits (2043, -0.90 DPS) [quest]; Reedknot Ring (9622, -1.05 DPS) [quest]; Minor Channeling Ring (1449, -1.13 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (1.52 DPS) | yes | Ring of Forlorn Spirits (2043, -0.27 DPS) [quest]; Reedknot Ring (9622, -0.43 DPS) [quest]; Minor Channeling Ring (1449, -0.51 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.23 DPS) [dungeon]; Spine of the Devourer (273040, -2.18 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.2 spell_power points (2.38 DPS) | yes | Orb of the Forgotten Seer (7685, -0.19 DPS) [dungeon]; Orb of Lorica (11262, -0.28 DPS) [quest]; Orb of Mystic Insight (249394, -0.59 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 256.3 spell_power points (40.08 DPS) | yes | Nether Force Wand (11263, -2.40 DPS) [quest]; Icefury Wand (7514, -2.54 DPS) [quest]; Ragefire Wand (7513, -2.59 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 230005000000000000-00000000000000000-253510130100301231)

Set DPS (verified): 253.2. Weights run: 1.1s. Verify run: 1.0s. 512 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=1.086 ± 0.043, crit=0.277 ± 0.011 per rating point (14 rating = 1%, 3.878 per %), hit=0.868 ± 0.056 per rating point (10 rating = 1%, 8.685 per %), spell_haste=6.633 ± 0.687, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.7 spell_power points (9.57 DPS) | yes | Dreamweave Circlet (10041, -2.08 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.31 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -3.23 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (3.69 DPS) | yes | Mindburst Medallion (11196, -0.75 DPS) [quest]; Triune Amulet (7722, -1.81 DPS) [dungeon]; Horizon Choker (13085, -3.18 DPS, sim-verified) [world_drop] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.5 spell_power points (7.65 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -2.56 DPS) [vendor]; Inquisitor's Shawl (19507, -2.69 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.84 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.5 spell_power points (4.82 DPS) | yes | Runecloth Cloak (13860, -0.66 DPS) [crafted]; Big Voodoo Cloak (8216, -1.35 DPS) [crafted]; Imperial Red Cloak (8248, -2.01 DPS) [world_drop] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.7 spell_power points (9.57 DPS) | yes | Runecloth Tunic (13857, -2.77 DPS) [crafted]; Robe of the Magi (1716, -2.87 DPS) [world_drop]; Runecloth Robe (13858, -4.56 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.8 spell_power points (3.47 DPS) | yes | Nethergeld Cuffs (254061, -0.04 DPS) [crafted]; Forgotten Wraps (9433, -0.41 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.41 DPS) [quest] |
| hands | Sergeant Major's Dreadweave Gloves (220890) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 22.8 spell_power points (5.35 DPS) | yes | Dreamweave Gloves (10019, -0.10 DPS) [crafted]; Red Mageweave Gloves (10018, -0.21 DPS) [crafted]; Runecloth Gloves (13863, -0.24 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.6 spell_power points (6.26 DPS) | yes | Satyrmane Sash (17755, -0.42 DPS) [dungeon]; Deathmage Sash (10771, -0.79 DPS) [dungeon]; Runecloth Belt (13856, -1.32 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 28.9 spell_power points (6.80 DPS) | yes | Red Mageweave Pants (10009, -0.44 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.50 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.60 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 26.5 spell_power points (6.22 DPS) | yes | Earthen Silk Slippers (254013, -0.58 DPS) [crafted]; Gilded Sandals (254107, -1.34 DPS) [crafted]; Southsea Mojo Boots (20641, -1.53 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (3.63 DPS) | yes | Snake Hoop (6750, -0.94 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.75 DPS) [quest]; Voodoo Band (1996, -1.84 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.06 DPS) | yes | Snake Hoop (6750, -0.37 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.18 DPS) [quest]; Voodoo Band (1996, -1.27 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (253.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (253.2 DPS) | yes | Mark of the Chosen (17774, -4.31 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (253.2 DPS) | yes | Spellshifter Rod (9527, -0.59 DPS) [quest]; Kindling Stave (11750, -1.99 DPS) [dungeon]; Hanzo Sword (8190, -45.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 223.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.32 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.87 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sergeant Major's Dreadweave Gloves; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 230005000000000000-23023000000000000-253510130100301231)

Set DPS (verified): 331.9. Weights run: 1.2s. Verify run: 2.9s. 1131 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.794 ± 0.049, crit=0.393 ± 0.014 per rating point (14 rating = 1%, 5.508 per %), hit=1.107 ± 0.067 per rating point (10 rating = 1%, 11.067 per %), spell_haste=6.089 ± 0.840, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (331.9 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Magister's Crown (16686, -4.71 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 19.9 spell_power points (4.91 DPS) | yes | Glowing Eye of Mordresh (10769, -1.62 DPS) [dungeon]; Horizon Choker (13085, -2.17 DPS) [world_drop]; Mindburst Medallion (11196, -2.26 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (331.9 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.8 spell_power points (4.62 DPS) | yes | Runecloth Cloak (13860, -0.84 DPS) [crafted]; Big Voodoo Cloak (8216, -1.63 DPS) [crafted]; Long Silken Cloak (4326, -2.16 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.0 spell_power points (13.54 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.74 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -3.70 DPS) [pvp]; Acumen Robes (17775, -7.88 DPS, sim-verified) [quest] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (331.9 DPS) | yes | Marshal's Silk Bracers (16438, +0.00 DPS) [pvp]; Dreadmist Bracers (16703, +0.00 DPS) [dungeon]; Runecloth Cuffs (254123, -4.82 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 28.0 spell_power points (6.88 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -4.52 DPS, sim-verified) [quest] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (331.9 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (331.9 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Magister's Leggings (16687, -4.32 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.5 spell_power points (7.26 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.14 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.82 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.6 spell_power points (5.80 DPS) | yes | Band of the Unicorn (7553, -2.60 DPS) [world_drop]; Snake Hoop (6750, -3.34 DPS) [quest]; Glowing Crystal Ring (18402, -3.45 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (3.44 DPS) | yes | Band of the Unicorn (7553, -0.24 DPS) [world_drop]; Snake Hoop (6750, -0.98 DPS) [quest]; Glowing Crystal Ring (18402, -1.09 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (331.9 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (331.9 DPS) | yes | Mark of the Chosen (17774, -4.40 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (331.9 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.19 DPS) [quest]; Teebu's Blazing Longsword (1728, -46.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 318.1 spell_power points (78.24 DPS) | yes | Stormrager (16997, -6.23 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.07 DPS) [crafted]; Torch of Austen (13004, -21.45 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; ranged: Torch of Light

No-known-source sample (15 of 1131, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 230005200000000000-51020000000000000-253510130101301221)

Set DPS (verified): 568.3. Weights run: 1.4s. Verify run: 2.0s. 1131 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.256 ± 0.016, crit=0.435 ± 0.013 per rating point (14 rating = 1%, 6.086 per %), hit=1.094 ± 0.058 per rating point (10 rating = 1%, 10.944 per %), spell_haste=6.168 ± 0.368, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.821 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.58 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -1.35 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.6 spell_power points (6.78 DPS) | yes | Glowing Eye of Mordresh (10769, -2.57 DPS) [dungeon]; Triune Amulet (7722, -3.05 DPS) [dungeon]; Mindburst Medallion (11196, -3.27 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.4 spell_power points (12.30 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.66 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, -1.17 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (7.24 DPS) | yes | Runecloth Cloak (13860, -2.09 DPS) [crafted]; Nightfall Drape (12465, -3.04 DPS) [dungeon]; Big Voodoo Cloak (8216, -3.84 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 49.2 spell_power points (22.90 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.67 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -8.26 DPS) [pvp]; Acumen Robes (17775, -13.55 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.6 spell_power points (6.31 DPS) | yes | Arcane Runed Bracers (4744, -2.12 DPS) [quest]; Spidertank Oilrag (9448, -2.12 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.22 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 23.7 spell_power points (11.05 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.6 spell_power points (7.71 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -1.19 DPS) [crafted]; Star Belt (4329, -1.66 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (12.11 DPS) | yes | Knight's Dreadweave Leggings (220888, +0.00 DPS) [vendor]; Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.18 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.21 DPS) [pvp]; Omnicast Boots (11822, -0.43 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.8 spell_power points (9.22 DPS) | yes | Philanthropist's Ring (281635, -3.96 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.49 DPS) [quest]; Snake Hoop (6750, -5.83 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.06 DPS) | yes | Philanthropist's Ring (281635, -0.80 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.33 DPS) [quest]; Snake Hoop (6750, -2.67 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (568.3 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (568.3 DPS) | yes | Mark of the Chosen (17774, -5.88 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (568.3 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.50 DPS) [world]; Teebu's Blazing Longsword (1728, -40.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 170.8 spell_power points (79.55 DPS) | yes | Brilliant Wand (249385, -14.95 DPS) [crafted]; Stormrager (16997, -16.27 DPS) [quest]; Greater Eternal Wand (249237, -22.03 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1131, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 30.1. Weights run: 1.0s. Verify run: 0.7s. 154 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.377 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.006 per %), hit=0.223 ± 0.007 per rating point (10 rating = 1%, 2.229 per %), spell_haste=not significant (0.176 ± 0.108), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.65 DPS) | yes | Shadow Goggles (4373, -1.25 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.4 spell_power points (0.91 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.36 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.44 DPS) | yes | Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Mystic Shawl (3449, -0.22 DPS) [quest]; Pearl-clasped Cloak (5542, -0.33 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.9 spell_power points (0.75 DPS) | yes | Green Woolen Robe (6243, -0.30 DPS) [crafted]; Green Woolen Vest (2582, -0.31 DPS) [crafted]; Gray Woolen Robe (2585, -0.99 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.33 DPS) | yes | Mindthrust Bracers (1974, -0.12 DPS) [dungeon]; Featherbead Bracers (15452, -0.12 DPS) [quest]; Owlbeard Bracers (16981, -0.14 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.76 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.22 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.33 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.5 spell_power points (0.60 DPS) | yes | Novice Ardent's Sash (253887, -0.26 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.55 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.3 spell_power points (0.90 DPS) | yes | Colorful Kilt (10048, -0.36 DPS) [crafted]; Embroidered Leggings (270011, -0.46 DPS) [quest]; Silk-threaded Trousers (1929, -1.20 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.93 DPS) | yes | Pristine Boots (253889, -0.18 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.49 DPS) [crafted]; Woolen Boots (2583, -0.63 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.3 spell_power points (0.25 DPS) | yes | Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; The 1 Ring (8350, -0.21 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 1.9 spell_power points (0.21 DPS) | yes | Volcanic Rock Ring (12053, -0.15 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.16 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.8 spell_power points (0.41 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.16 DPS) [world]; Clear Crystal Rod (16894, -0.21 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 207.8 spell_power points (22.64 DPS) | yes | Skycaller (12984, -1.37 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.44 DPS) [quest]; Greater Magic Wand (11288, -4.92 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 154, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (troll, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 61.0. Weights run: 1.0s. Verify run: 0.8s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.386 ± 0.009, crit=0.114 ± 0.004 per rating point (14 rating = 1%, 1.589 per %), hit=0.230 ± 0.010 per rating point (10 rating = 1%, 2.300 per %), spell_haste=not significant (0.509 ± 0.155), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 spell_power points (1.68 DPS) | yes | Silk Headband (7050, -0.31 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.46 DPS) [crafted]; Enchanter's Cowl (4322, -1.60 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.3 spell_power points (0.35 DPS) | yes | Crystal Starfire Medallion (5003, -0.12 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.12 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (1.91 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.46 DPS) [quest]; Invoker's Mantle (215365, -0.54 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.76 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.28 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.0 spell_power points (2.14 DPS) | yes | Beguiler Robes (7728, -0.06 DPS) [dungeon]; Tree Bark Jacket (1486, -0.16 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.21 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.37 DPS) | yes | Tabitha's Cuffs (251486, -0.92 DPS) [quest]; Nightsky Wristbands (6407, -1.02 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.06 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.9 spell_power points (1.21 DPS) | yes | Serpent Gloves (5970, -0.14 DPS) [dungeon]; Truefaith Gloves (7049, -0.27 DPS) [crafted]; Gnoll Casting Gloves (892, -0.29 DPS) [world] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Invoker's Cord (215366, -0.19 DPS) [crafted]; Crimson Silk Belt (7055, -0.22 DPS) [crafted]; Warsong Sash (16975, -0.69 DPS, sim-verified) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.21 DPS) [crafted]; Silk-threaded Trousers (1929, -0.41 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.43 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.7 spell_power points (1.48 DPS) | yes | Acidic Walkers (9454, -0.25 DPS) [dungeon]; Glinteye Slippers (273024, -0.41 DPS) [dungeon]; Spidersilk Boots (4320, -1.74 DPS, sim-verified) [crafted] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 7.9 spell_power points (1.21 DPS) | yes | Solid Gold Band (279845, -0.60 DPS) [quest]; Electrocutioner Lagnut (9447, -0.75 DPS) [dungeon]; Runemender's Seal (273031, -0.75 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.92 DPS) | yes | Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Runemender's Seal (273031, -0.46 DPS) [world]; Solid Gold Band (279845, -0.55 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (5.50 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -4.85 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -4.91 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Starfaller (13063, -0.62 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -1.65 DPS) [crafted]; Unstable Power Core (279847, -1.71 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (troll, 000000000000000000-00000000000000000-253510130100301231)

Set DPS (verified): 93.7. Weights run: 1.2s. Verify run: 0.9s. 386 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.746 ± 0.024, crit=0.178 ± 0.007 per rating point (14 rating = 1%, 2.486 per %), hit=0.362 ± 0.025 per rating point (10 rating = 1%, 3.616 per %), spell_haste=not significant (1.421 ± 0.373), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.35 DPS) [world]; Whitemane's Chapeau (7720, -0.97 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.10 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.0 spell_power points (2.03 DPS) | yes | Triune Amulet (7722, -0.85 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.21 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.33 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.7 spell_power points (2.61 DPS) | yes | Bloodmage Mantle (7684, -0.15 DPS) [dungeon]; Berylline Pads (4197, -0.35 DPS) [quest]; Green Silken Shoulders (7057, -1.19 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.7 spell_power points (1.52 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.43 DPS) [crafted]; Blackforge Cape (6424, -0.59 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.5 spell_power points (4.14 DPS) | yes | Dreamweave Vest (10021, -0.28 DPS) [crafted]; Robe of Power (7054, -0.55 DPS) [crafted]; Elemental Raiment (9434, -0.86 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.0 spell_power points (1.56 DPS) | yes | Spidertank Oilrag (9448, -0.15 DPS) [dungeon]; Condor Bracers (15864, -0.46 DPS) [quest]; Windchaser Cuffs (14429, -0.51 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.0 spell_power points (3.28 DPS) | yes | Black Mageweave Gloves (10003, -0.94 DPS) [crafted]; Gilded Handwraps (254021, -1.21 DPS) [crafted]; Red Mageweave Gloves (10018, -1.55 DPS, sim-verified) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Star Belt (4329, -0.15 DPS) [crafted]; Belt of Arugal (6392, -0.43 DPS) [dungeon]; Deathmage Sash (10771, -1.33 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.0 spell_power points (3.59 DPS) | yes | Crimson Silk Pantaloons (7062, -1.00 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.56 DPS) [dungeon]; Pristine Leggings (253987, -1.68 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.75 DPS) | yes | Gilded Slippers (254001, -1.42 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.04 DPS) [dungeon]; Spidersilk Boots (4320, -2.19 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (2.15 DPS) | yes | Reedknot Ring (9622, -1.05 DPS) [quest]; Ring of Power Regulation (273030, -1.21 DPS) [dungeon]; Voodoo Band (1996, -1.33 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.7 spell_power points (1.52 DPS) | yes | Reedknot Ring (9622, -0.43 DPS) [quest]; Ring of Power Regulation (273030, -0.58 DPS) [dungeon]; Voodoo Band (1996, -0.70 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.23 DPS) [dungeon]; Spine of the Devourer (273040, -2.18 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.2 spell_power points (2.38 DPS) | yes | Orb of the Forgotten Seer (7685, -0.19 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.59 DPS) [crafted]; Omega Orb (7749, -0.98 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 256.3 spell_power points (40.08 DPS) | yes | Nether Force Wand (11263, -2.40 DPS) [quest]; Icefury Wand (7514, -2.54 DPS) [quest]; Ragefire Wand (7513, -2.59 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 386, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (troll, 230005000000000000-00000000000000000-253510130100301231)

Set DPS (verified): 242.2. Weights run: 1.1s. Verify run: 1.0s. 502 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=1.086 ± 0.043, crit=0.277 ± 0.011 per rating point (14 rating = 1%, 3.878 per %), hit=0.868 ± 0.056 per rating point (10 rating = 1%, 8.685 per %), spell_haste=6.633 ± 0.687, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.7 spell_power points (9.57 DPS) | yes | Dreamweave Circlet (10041, -2.08 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -2.31 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -3.23 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.7 spell_power points (3.69 DPS) | yes | Horizon Choker (13085, -0.11 DPS) [world_drop]; Mindburst Medallion (11196, -0.75 DPS) [quest]; Triune Amulet (7722, -1.81 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.5 spell_power points (7.65 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -2.56 DPS) [vendor]; Inquisitor's Shawl (19507, -2.69 DPS) [dungeon]; Red Mageweave Shoulders (10029, -3.95 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.8 spell_power points (5.12 DPS) | yes | Spritecaster Cape (11623, -0.30 DPS) [dungeon]; Runecloth Cloak (13860, -0.96 DPS) [crafted]; Big Voodoo Cloak (8216, -1.65 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.7 spell_power points (9.57 DPS) | yes | Runecloth Tunic (13857, -2.77 DPS) [crafted]; Robe of the Magi (1716, -2.87 DPS) [world_drop]; Runecloth Robe (13858, -4.62 DPS, sim-verified) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | 14.8 spell_power points (3.47 DPS) | yes | Nethergeld Cuffs (254061, -0.04 DPS) [crafted]; Forgotten Wraps (9433, -0.41 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.41 DPS) [quest] |
| hands | First Sergeant's Dreadweave Gloves (220908) | PvP rank 9 · First Sergeant · Horde [vendor] | 22.8 spell_power points (5.35 DPS) | yes | Dreamweave Gloves (10019, -0.10 DPS) [crafted]; Red Mageweave Gloves (10018, -0.21 DPS) [crafted]; Runecloth Gloves (13863, -0.24 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.6 spell_power points (6.26 DPS) | yes | Satyrmane Sash (17755, -0.42 DPS) [dungeon]; Deathmage Sash (10771, -0.79 DPS) [dungeon]; Runecloth Belt (13856, -1.32 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 28.9 spell_power points (6.80 DPS) | yes | Red Mageweave Pants (10009, -0.44 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.50 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.60 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 26.5 spell_power points (6.22 DPS) | yes | Earthen Silk Slippers (254013, -0.58 DPS) [crafted]; Gilded Sandals (254107, -1.34 DPS) [crafted]; Southsea Mojo Boots (20641, -1.53 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (3.63 DPS) | yes | Snake Hoop (6750, -0.94 DPS) [quest]; Ogremind Ring (1993, -1.84 DPS) [world_drop]; Voodoo Band (1996, -1.84 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.06 DPS) | yes | Snake Hoop (6750, -0.37 DPS) [quest]; Ogremind Ring (1993, -1.27 DPS) [world_drop]; Voodoo Band (1996, -1.27 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (242.2 DPS) | yes | Rune of the Guard Captain (19120, -0.41 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (242.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (242.2 DPS) | yes | Spellshifter Rod (9527, -0.59 DPS) [quest]; Kindling Stave (11750, -1.99 DPS) [dungeon]; Hanzo Sword (8190, -42.71 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 223.3 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.32 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.87 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: First Sergeant's Dreadweave Gloves; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 502, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (troll, 230005000000000000-23023000000000000-253510130100301231)

Set DPS (verified): 318.2. Weights run: 1.2s. Verify run: 2.9s. 1123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.007, intellect=0.794 ± 0.049, crit=0.393 ± 0.014 per rating point (14 rating = 1%, 5.508 per %), hit=1.107 ± 0.067 per rating point (10 rating = 1%, 11.067 per %), spell_haste=6.089 ± 0.840, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.007

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (318.2 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -6.89 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 19.9 spell_power points (4.91 DPS) | yes | Glowing Eye of Mordresh (10769, -1.62 DPS) [dungeon]; Horizon Choker (13085, -2.17 DPS) [world_drop]; Mindburst Medallion (11196, -2.26 DPS) [quest] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (318.2 DPS) | yes | Champion's Silk Mantle (227104, +0.00 DPS) [pvp]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.33 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.1 spell_power points (4.71 DPS) | yes | Spritecaster Cape (11623, -0.09 DPS) [dungeon]; Runecloth Cloak (13860, -0.93 DPS) [crafted]; Big Voodoo Cloak (8216, -1.72 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.0 spell_power points (13.54 DPS) | yes | Warlord's Silk Raiment (231596, -0.74 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -3.70 DPS) [pvp]; Acumen Robes (17775, -6.96 DPS, sim-verified) [quest] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (318.2 DPS) | yes | General's Silk Cuffs (16538, +0.00 DPS) [pvp]; Dreadmist Bracers (16703, +0.00 DPS) [dungeon]; Runecloth Cuffs (254123, -5.13 DPS, sim-verified) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 28.0 spell_power points (6.88 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (318.2 DPS) | yes | Wisdom of the Timbermaw (19047, -0.38 DPS) [crafted]; Satyrmane Sash (17755, -0.43 DPS) [dungeon]; Dawnspire Cord (12466, -0.64 DPS) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 28.9 spell_power points (7.10 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Stone Guard's Dreadweave Leggings (220906, -10.36 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-verified (318.2 DPS) | yes | Omnicast Boots (11822, +0.00 DPS) [dungeon]; Blood Guard's Silk Walkers (227110, +0.00 DPS) [pvp]; General's Silk Boots (231597, +0.00 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.6 spell_power points (5.80 DPS) | yes | Band of the Unicorn (7553, -2.60 DPS) [world_drop]; Snake Hoop (6750, -3.34 DPS) [quest]; Glowing Crystal Ring (18402, -3.45 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (3.44 DPS) | yes | Band of the Unicorn (7553, -0.24 DPS) [world_drop]; Snake Hoop (6750, -0.98 DPS) [quest]; Glowing Crystal Ring (18402, -1.09 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (318.2 DPS) | yes | Uther's Strength (11302, -0.97 DPS) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (318.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.19 DPS) [quest]; Teebu's Blazing Longsword (1728, -43.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 318.1 spell_power points (78.24 DPS) | yes | Stormrager (16997, -4.89 DPS, sim-verified) [quest]; Brilliant Wand (249385, -15.07 DPS) [crafted]; Torch of Austen (13004, -21.45 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Magister's Mantle; chest: Robe of the Archmage; wrist: Magister's Bindings; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; finger1: Maiden's Circle; finger2: Philanthropist's Ring; ranged: Torch of Light

No-known-source sample (15 of 1123, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (troll, 230005200000000000-51020000000000000-253510130101301221)

Set DPS (verified): 562.3. Weights run: 1.4s. Verify run: 2.0s. 1123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.256 ± 0.016, crit=0.435 ± 0.013 per rating point (14 rating = 1%, 6.086 per %), hit=1.094 ± 0.058 per rating point (10 rating = 1%, 10.944 per %), spell_haste=6.168 ± 0.368, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.821 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.58 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -1.35 DPS) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.6 spell_power points (6.78 DPS) | yes | Glowing Eye of Mordresh (10769, -2.57 DPS) [dungeon]; Triune Amulet (7722, -3.05 DPS) [dungeon]; Mindburst Medallion (11196, -3.27 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 26.4 spell_power points (12.30 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.66 DPS) [crafted]; Champion's Silk Mantle (227104, -1.17 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 spell_power points (7.24 DPS) | yes | Deep Woodlands Cloak (19121, -0.57 DPS) [quest]; Runecloth Cloak (13860, -2.09 DPS) [crafted]; Nightfall Drape (12465, -3.04 DPS) [dungeon] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 49.2 spell_power points (22.90 DPS) | yes | Warlord's Silk Raiment (231596, -2.67 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -8.26 DPS) [pvp]; Acumen Robes (17775, -13.84 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.6 spell_power points (6.31 DPS) | yes | Spidertank Oilrag (9448, -2.12 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.22 DPS) [crafted]; Earthenweave Cuffs (254131, -2.59 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 23.7 spell_power points (11.05 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.6 spell_power points (7.71 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Ghostweave Cord (254073, -1.19 DPS) [crafted]; Star Belt (4329, -1.66 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (12.11 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS) [vendor]; Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.18 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.21 DPS) [pvp]; Omnicast Boots (11822, -0.43 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.8 spell_power points (9.22 DPS) | yes | Philanthropist's Ring (281635, -3.96 DPS) [quest]; Snake Hoop (6750, -5.83 DPS) [quest]; Reedknot Ring (9622, -5.96 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.06 DPS) | yes | Philanthropist's Ring (281635, -0.80 DPS) [quest]; Snake Hoop (6750, -2.67 DPS) [quest]; Reedknot Ring (9622, -2.79 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (562.3 DPS) | yes | Uther's Strength (11302, -1.79 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (562.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (562.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spire of Hakkar (10844, -0.50 DPS) [world]; Teebu's Blazing Longsword (1728, -40.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 170.8 spell_power points (79.55 DPS) | yes | Brilliant Wand (249385, -14.95 DPS) [crafted]; Stormrager (16997, -16.27 DPS) [quest]; Greater Eternal Wand (249237, -22.03 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Spritecaster Cape; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1123, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

