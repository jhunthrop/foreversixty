# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 141005000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 29.8. Weights run: 1.1s. Verify run: 0.7s. 160 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.313 ± 0.008, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.034 per %), hit=0.261 ± 0.009 per rating point (10 rating = 1%, 2.606 per %), spell_haste=0.437 ± 0.093, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.491 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -1.47 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.8 spell_power points (0.77 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.38 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.50 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.40 DPS) | yes | Pearl-clasped Cloak (5542, -0.11 DPS) [crafted]; Sanguine Cape (14376, -0.27 DPS) [world_drop]; Black Whelp Cloak (7283, -0.42 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.6 spell_power points (0.65 DPS) | yes | Green Woolen Vest (2582, -0.25 DPS) [crafted]; Bloody Apron (6226, -0.25 DPS) [dungeon]; Gray Woolen Robe (2585, -0.96 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 1.6 spell_power points (0.15 DPS) | yes | Windsong Bangles (263336, -0.06 DPS) [quest]; Repurposed Hair Band (281256, -0.09 DPS) [quest]; Bright Bracers (3647, -0.99 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Gnoll Casting Gloves (892, -0.20 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.20 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.43 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.3 spell_power points (0.52 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.81 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 7.9 spell_power points (0.78 DPS) | yes | Colorful Kilt (10048, -0.29 DPS) [crafted]; Heavy Woolen Pants (4316, -0.48 DPS) [crafted]; Silk-threaded Trousers (1929, -1.38 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.3 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.43 DPS) [crafted]; Woolen Boots (2583, -0.56 DPS) [crafted]; Red Woolen Boots (4313, -0.64 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.6 spell_power points (0.56 DPS) | yes | Volcanic Rock Ring (12053, -0.46 DPS) [world_drop]; The 1 Ring (8350, -0.53 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.9 spell_power points (0.19 DPS) | yes | The 1 Ring (8350, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.49 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 3.1 spell_power points (0.31 DPS) | yes | Witching Stave (1484, -0.15 DPS) [dungeon]; Staff of Westfall (2042, -0.15 DPS) [quest]; Lesser Staff of the Spire (1300, -0.61 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 228.2 spell_power points (22.61 DPS) | yes | Skycaller (12984, -0.77 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.08 DPS) [quest]; Sizzle Stick (8071, -4.46 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 160, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 141005213100300000-00000000000000000-000000000000000000)

Set DPS (verified): 92.9. Weights run: 1.2s. Verify run: 0.8s. 298 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.561 ± 0.019, crit=0.073 ± 0.002 per rating point (14 rating = 1%, 1.024 per %), hit=0.290 ± 0.029 per rating point (10 rating = 1%, 2.899 per %), spell_haste=not significant (0.511 ± 0.288), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.589 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.36 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.55 DPS) [crafted]; Enchanter's Cowl (4322, -3.04 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.61 DPS) | yes | Crystal Starfire Medallion (5003, -0.20 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.20 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (2.56 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.55 DPS) [quest]; Magician's Mantle (12998, -0.73 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.91 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Prelacy Cape (7004, -0.18 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Death Speaker Robes (6682, -0.47 DPS) [dungeon]; Green Silk Armor (7065, -1.58 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.64 DPS) | yes | Nightsky Wristbands (6407, -1.03 DPS) [world_drop]; Stonecloth Bindings (14416, -1.13 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.53 DPS, sim-verified) [world_drop] |
| hands | Shilly Mitts (9609) (or Serpent Gloves (5970)) | Gyrodrillmatic Excavationators [quest] | 7.0 spell_power points (1.28 DPS) | yes | Serpent Gloves (5970, +0.00 DPS) [dungeon]; Town Clerk's Mittens (270029, -0.01 DPS) [quest]; Truefaith Gloves (7049, -0.06 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 10.7 spell_power points (1.95 DPS) | yes | Crimson Silk Belt (7055, -0.14 DPS) [crafted]; Invoker's Cord (215366, -0.16 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.49 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.28 DPS) [crafted]; Silk-threaded Trousers (1929, -0.72 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.63 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.99 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Glinteye Slippers (273024, -0.72 DPS) [dungeon]; Acidic Walkers (9454, -1.90 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.61 DPS) | yes | Ring of Power Regulation (273030, -0.51 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.99 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.1 spell_power points (1.12 DPS) | yes | Ring of Power Regulation (273030, -0.02 DPS) [dungeon]; Lavishly Jeweled Ring (1156, -0.50 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.57 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (6.56 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -2.92 DPS) [quest]; Glimmering Staff (249392, -5.44 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Mana Concentration (273043) | Arcanic Enigma [world] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Starfaller (13063, -1.04 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -1.68 DPS) [crafted]; Gravestone Scepter (7001, -2.59 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Shilly Mitts; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Minor Channeling Ring; main_hand: Spine of the Devourer; ranged: Wand of Mana Concentration

No-known-source sample (15 of 298, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 141.5. Weights run: 1.4s. Verify run: 1.0s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.833 ± 0.037, crit=0.166 ± 0.004 per rating point (14 rating = 1%, 2.328 per %), hit=0.426 ± 0.050 per rating point (10 rating = 1%, 4.257 per %), spell_haste=not significant (1.078 ± 0.566), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.646 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 22.2 spell_power points (4.32 DPS) | yes | Augural Shroud (2620, -0.55 DPS) [world]; Whitemane's Chapeau (7720, -1.33 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -3.07 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.7 spell_power points (2.66 DPS) | yes | Triune Amulet (7722, -1.10 DPS) [dungeon]; Necklace of Calisea (1714, -1.53 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.69 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.8 spell_power points (3.48 DPS) | yes | Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest]; Green Silken Shoulders (7057, -2.46 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.2 spell_power points (1.98 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.62 DPS) [crafted]; Blackforge Cape (6424, -0.68 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.0 spell_power points (5.26 DPS) | yes | Dreamweave Vest (10021, -0.29 DPS) [crafted]; Robe of Power (7054, -0.59 DPS) [crafted]; Elemental Raiment (9434, -1.17 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.75 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS) [dungeon]; Windchaser Cuffs (14429, -0.29 DPS) [world_drop]; Condor Bracers (15864, -0.39 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (4.16 DPS) | yes | Black Mageweave Gloves (10003, -1.23 DPS) [crafted]; Stormcloth Gloves (10011, -1.43 DPS) [crafted]; Red Mageweave Gloves (10018, -2.09 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.5 spell_power points (3.80 DPS) | yes | Gilded Cord (254037, -0.94 DPS) [crafted]; Star Belt (4329, -1.27 DPS) [crafted]; Crimson Silk Belt (7055, -1.49 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.0 spell_power points (4.68 DPS) | yes | Crimson Silk Pantaloons (7062, -1.80 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.01 DPS) [dungeon]; Pristine Leggings (253987, -2.18 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.68 DPS) | yes | Gilded Slippers (254001, -2.15 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.41 DPS) [dungeon]; Spidersilk Boots (4320, -2.67 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (2.76 DPS) | yes | Ring of Forlorn Spirits (2043, -1.20 DPS) [quest]; Reedknot Ring (9622, -1.40 DPS) [quest]; Minor Channeling Ring (1449, -1.46 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.2 spell_power points (1.98 DPS) | yes | Ring of Forlorn Spirits (2043, -0.42 DPS) [quest]; Reedknot Ring (9622, -0.62 DPS) [quest]; Minor Channeling Ring (1449, -0.68 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (141.5 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.32 DPS) [dungeon]; Spine of the Devourer (273040, -2.86 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.5 spell_power points (3.02 DPS) | yes | Orb of Lorica (11262, -0.23 DPS) [quest]; Orb of the Forgotten Seer (7685, -0.29 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.68 DPS) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 207.0 spell_power points (40.35 DPS) | yes | Nether Force Wand (11263, -2.54 DPS) [quest]; Icefury Wand (7514, -2.68 DPS) [quest]; Ragefire Wand (7513, -2.73 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 141005213100301531-32023000000000000-000000000000000000)

Set DPS (verified): 239.1. Weights run: 1.3s. Verify run: 1.2s. 512 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.073 ± 0.058, crit=0.261 ± 0.007 per rating point (14 rating = 1%, 3.659 per %), hit=0.644 ± 0.081 per rating point (10 rating = 1%, 6.445 per %), spell_haste=not significant (1.716 ± 0.866), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.662 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.5 spell_power points (7.60 DPS) | yes | Dreamweave Circlet (10041, -1.64 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.86 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.53 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.6 spell_power points (2.93 DPS) | yes | Horizon Choker (13085, -0.11 DPS) [world_drop]; Mindburst Medallion (11196, -0.59 DPS) [quest]; Triune Amulet (7722, -1.42 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.3 spell_power points (6.07 DPS) | yes | Knight-Lieutenant's Dreadweave Mantle (220887, -2.07 DPS) [vendor]; Inquisitor's Shawl (19507, -2.13 DPS) [dungeon]; Red Mageweave Shoulders (10029, -2.90 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.4 spell_power points (3.84 DPS) | yes | Runecloth Cloak (13860, -0.54 DPS) [crafted]; Big Voodoo Cloak (8216, -1.09 DPS) [crafted]; Imperial Red Cloak (8248, -1.62 DPS) [world_drop] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.5 spell_power points (7.60 DPS) | yes | Runecloth Tunic (13857, -2.19 DPS) [crafted]; Robe of the Magi (1716, -2.26 DPS) [world_drop]; Runecloth Robe (13858, -3.51 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Forgotten Wraps (9433, -0.31 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.31 DPS) [quest]; Bloodband Bracers (11469, -3.41 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Gloves (10018, -0.11 DPS) [crafted]; Runecloth Gloves (13863, -0.12 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -2.98 DPS, sim-verified) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.4 spell_power points (4.96 DPS) | yes | Satyrmane Sash (17755, -0.31 DPS) [dungeon]; Deathmage Sash (10771, -0.62 DPS) [dungeon]; Runecloth Belt (13856, -1.04 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 28.5 spell_power points (5.36 DPS) | yes | Red Mageweave Pants (10009, -0.31 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.17 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.24 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 24.1 spell_power points (4.53 DPS) | yes | Earthen Silk Slippers (254013, -0.02 DPS) [crafted]; Gilded Sandals (254107, -0.65 DPS) [crafted]; Southsea Mojo Boots (20641, -0.81 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.89 DPS) | yes | Snake Hoop (6750, -0.75 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.38 DPS) [quest]; Voodoo Band (1996, -1.47 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.44 DPS) | yes | Snake Hoop (6750, -0.31 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.94 DPS) [quest]; Voodoo Band (1996, -1.03 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -3.67 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.46 DPS) [quest]; Kindling Stave (11750, -1.55 DPS) [dungeon]; Hanzo Sword (8190, -41.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 279.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.56 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.25 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 512, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 141005213100301531-32023000000000000-053200000000000000)

Set DPS (verified): 398.8. Weights run: 1.4s. Verify run: 3.2s. 1131 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.149 ± 0.066, crit=0.411 ± 0.011 per rating point (14 rating = 1%, 5.756 per %), hit=1.115 ± 0.107 per rating point (10 rating = 1%, 11.154 per %), spell_haste=not significant (2.372 ± 1.180), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.638 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (398.8 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Magister's Crown (16686, -6.46 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 23.5 spell_power points (5.22 DPS) | yes | Glowing Eye of Mordresh (10769, -1.62 DPS) [dungeon]; Horizon Choker (13085, -1.64 DPS) [world_drop]; Opaline Medallion (16623, -1.90 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (398.8 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Magister's Mantle (16689, -5.46 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.9 spell_power points (4.64 DPS) | yes | Runecloth Cloak (13860, -0.60 DPS) [crafted]; Big Voodoo Cloak (8216, -1.23 DPS) [crafted]; Imperial Red Cloak (8248, -1.83 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 59.5 spell_power points (13.23 DPS) | yes | Field Marshal's Silk Vestments (231603, -0.28 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -2.94 DPS) [pvp]; Magister's Robes (16688, -10.48 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (398.8 DPS) | yes | Marshal's Silk Bracers (16438, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Magister's Bindings (16683, -8.15 DPS, sim-verified) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (398.8 DPS) | yes | Gloves of Spell Mastery (14146, +0.00 DPS) [crafted]; Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (398.8 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Magister's Belt (16685, -5.84 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (398.8 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Magister's Leggings (16687, -6.32 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 33.8 spell_power points (7.51 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.29 DPS) [pvp]; Sergeant Major's Dreadweave Boots (220891, -0.95 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 26.0 spell_power points (5.79 DPS) | yes | Glowing Crystal Ring (18402, -2.72 DPS) [quest]; Band of the Unicorn (7553, -2.90 DPS) [world_drop]; Snake Hoop (6750, -3.18 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.7 spell_power points (3.50 DPS) | yes | Glowing Crystal Ring (18402, -0.43 DPS) [quest]; Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Snake Hoop (6750, -0.89 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (398.8 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (398.8 DPS) | yes | Mark of the Chosen (17774, -4.62 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (398.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.64 DPS) [quest]; Teebu's Blazing Longsword (1728, -51.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 351.5 spell_power points (78.09 DPS) | yes | Stormrager (16997, -4.65 DPS, sim-verified) [quest]; Brilliant Wand (249385, -14.73 DPS) [crafted]; Torch of Austen (13004, -21.30 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; ranged: Torch of Light

No-known-source sample (15 of 1131, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 131205213100301531-30000000000000000-053500030000000000)

Set DPS (verified): 663.4. Weights run: 1.6s. Verify run: 2.2s. 1131 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.784 ± 0.042, crit=0.515 ± 0.011 per rating point (14 rating = 1%, 7.204 per %), hit=1.208 ± 0.085 per rating point (10 rating = 1%, 12.082 per %), spell_haste=3.538 ± 0.795, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.670 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (663.4 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -8.69 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 19.8 spell_power points (8.81 DPS) | yes | Glowing Eye of Mordresh (10769, -3.30 DPS, sim-verified) [dungeon]; Horizon Choker (13085, -3.94 DPS) [world_drop]; Mindburst Medallion (11196, -4.06 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.4 spell_power points (15.27 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.58 DPS) [pvp]; Rotgrip Mantle (17732, -5.22 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.7 spell_power points (8.31 DPS) | yes | Runecloth Cloak (13860, -1.52 DPS) [crafted]; Big Voodoo Cloak (8216, -2.95 DPS) [crafted]; Long Silken Cloak (4326, -3.90 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.6 spell_power points (25.14 DPS) | yes | Field Marshal's Silk Vestments (231603, -1.37 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -6.70 DPS) [pvp]; Acumen Robes (17775, -11.74 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 18.8 spell_power points (8.37 DPS) | yes | Marshal's Silk Bracers (16438, -1.75 DPS) [pvp]; Dreadmist Bracers (16703, -2.32 DPS) [dungeon]; Magiskull Cuffs (13107, -2.45 DPS) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.2 spell_power points (13.87 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -9.86 DPS, sim-verified) [quest] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 23.5 spell_power points (10.42 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Satyrmane Sash (17755, -0.72 DPS) [dungeon]; Dawnspire Cord (12466, -1.14 DPS) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 29.6 spell_power points (13.16 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight's Dreadweave Leggings (220888, -3.20 DPS, sim-verified) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.4 spell_power points (13.06 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.25 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.5 spell_power points (10.43 DPS) | yes | Band of the Unicorn (7553, -4.66 DPS) [world_drop]; Snake Hoop (6750, -6.02 DPS) [quest]; Glowing Crystal Ring (18402, -6.25 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (6.18 DPS) | yes | Band of the Unicorn (7553, -0.41 DPS) [world_drop]; Snake Hoop (6750, -1.78 DPS) [quest]; Glowing Crystal Ring (18402, -2.00 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -6.35 DPS, sim-verified) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.31 DPS) [quest]; Teebu's Blazing Longsword (1728, -60.30 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 178.9 spell_power points (79.42 DPS) | yes | Brilliant Wand (249385, -13.41 DPS) [crafted]; Stormrager (16997, -14.99 DPS) [quest]; Greater Eternal Wand (249237, -22.10 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Mark of the Chosen; ranged: Torch of Light

No-known-source sample (15 of 1131, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 141005000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 27.9. Weights run: 1.1s. Verify run: 0.7s. 154 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.313 ± 0.008, crit=0.074 ± 0.002 per rating point (14 rating = 1%, 1.034 per %), hit=0.261 ± 0.009 per rating point (10 rating = 1%, 2.606 per %), spell_haste=0.437 ± 0.093, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.491 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.59 DPS) | yes | Shadow Goggles (4373, -1.41 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 7.8 spell_power points (0.77 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.38 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.41 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.40 DPS) | yes | Pearl-clasped Cloak (5542, -0.11 DPS) [crafted]; Mystic Shawl (3449, -0.20 DPS) [quest]; Black Whelp Cloak (7283, -0.47 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 6.6 spell_power points (0.65 DPS) | yes | Green Woolen Vest (2582, -0.25 DPS) [crafted]; Bloody Apron (6226, -0.25 DPS) [dungeon]; Gray Woolen Robe (2585, -0.73 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.30 DPS) | yes | Owlbeard Bracers (16981, -0.14 DPS) [quest]; Mindthrust Bracers (1974, -0.14 DPS) [dungeon]; Featherbead Bracers (15452, -0.14 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.69 DPS) | yes | Gnoll Casting Gloves (892, -0.17 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.20 DPS) [crafted]; Apothecary Gloves (10919, -0.30 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.3 spell_power points (0.52 DPS) | yes | Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Keller's Girdle (2911, -0.27 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.56 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 7.9 spell_power points (0.78 DPS) | yes | Colorful Kilt (10048, -0.29 DPS) [crafted]; Embroidered Leggings (270011, -0.38 DPS) [quest]; Silk-threaded Trousers (1929, -1.36 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.3 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.43 DPS) [crafted]; Woolen Boots (2583, -0.56 DPS) [crafted]; Red Woolen Boots (4313, -0.65 DPS, sim-verified) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 1.9 spell_power points (0.19 DPS) | yes | Volcanic Rock Ring (12053, -0.09 DPS) [world_drop]; The 1 Ring (8350, -0.15 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 1.6 spell_power points (0.15 DPS) | yes | The 1 Ring (8350, -0.12 DPS) [world]; Volcanic Rock Ring (12053, -0.31 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 3.1 spell_power points (0.31 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.12 DPS) [world]; Clear Crystal Rod (16894, -0.15 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 228.2 spell_power points (22.61 DPS) | yes | Skycaller (12984, -0.50 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.46 DPS) [quest]; Greater Magic Wand (11288, -4.91 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 154, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 141005213100300000-00000000000000000-000000000000000000)

Set DPS (verified): 90.3. Weights run: 1.2s. Verify run: 0.9s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.561 ± 0.019, crit=0.073 ± 0.002 per rating point (14 rating = 1%, 1.024 per %), hit=0.290 ± 0.029 per rating point (10 rating = 1%, 2.899 per %), spell_haste=not significant (0.511 ± 0.288), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.589 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.36 DPS) [crafted]; Filigreed Pristine Circlet (253975, -0.55 DPS) [crafted]; Enchanter's Cowl (4322, -2.61 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.61 DPS) | yes | Crystal Starfire Medallion (5003, -0.20 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.20 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (2.56 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.55 DPS) [quest]; Magician's Mantle (12998, -0.73 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.91 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.24 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Death Speaker Robes (6682, -0.47 DPS) [dungeon]; Green Silk Armor (7065, -1.37 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.64 DPS) | yes | Nightsky Wristbands (6407, -1.03 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.09 DPS) [quest]; Glowing Magical Bracelets (13106, -1.66 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.8 spell_power points (1.61 DPS) | yes | Serpent Gloves (5970, -0.33 DPS) [dungeon]; Truefaith Gloves (7049, -0.39 DPS) [crafted]; Gnoll Casting Gloves (892, -0.51 DPS) [world] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (2.00 DPS) | yes | Belt of Arugal (6392, -0.06 DPS) [dungeon]; Crimson Silk Belt (7055, -0.20 DPS) [crafted]; Invoker's Cord (215366, -0.22 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.28 DPS) [crafted]; Silk-threaded Trousers (1929, -0.72 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.39 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.99 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Glinteye Slippers (273024, -0.72 DPS) [dungeon]; Acidic Walkers (9454, -2.13 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.61 DPS) | yes | Solid Gold Band (279845, -0.88 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.99 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.06 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.09 DPS) | yes | Solid Gold Band (279845, -0.36 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.48 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.55 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (6.56 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -5.44 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -5.54 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 182.5 spell_power points (33.27 DPS) | yes | Wand of Mana Concentration (273043, -1.68 DPS) [world]; Greater Mystic Wand (217287, -3.36 DPS) [crafted]; Unstable Power Core (279847, -3.56 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Warsong Sash; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (orc, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 140.4. Weights run: 1.4s. Verify run: 1.0s. 386 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.833 ± 0.037, crit=0.166 ± 0.004 per rating point (14 rating = 1%, 2.328 per %), hit=0.426 ± 0.050 per rating point (10 rating = 1%, 4.257 per %), spell_haste=not significant (1.078 ± 0.566), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.646 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 22.2 spell_power points (4.32 DPS) | yes | Augural Shroud (2620, -0.55 DPS) [world]; Whitemane's Chapeau (7720, -1.33 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -2.95 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.7 spell_power points (2.66 DPS) | yes | Necklace of Calisea (1714, -1.53 DPS) [dungeon]; Triune Amulet (7722, -1.62 DPS, sim-verified) [dungeon]; Archmagister's Faceted Pendant (273049, -1.69 DPS) [world] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 17.8 spell_power points (3.48 DPS) | yes | Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Berylline Pads (4197, -0.49 DPS) [quest]; Green Silken Shoulders (7057, -1.71 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.2 spell_power points (1.98 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.62 DPS) [crafted]; Blackforge Cape (6424, -0.68 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.0 spell_power points (5.26 DPS) | yes | Dreamweave Vest (10021, -0.29 DPS) [crafted]; Robe of Power (7054, -0.59 DPS) [crafted]; Elemental Raiment (9434, -1.17 DPS) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 10.7 spell_power points (2.08 DPS) | yes | Spidertank Oilrag (9448, -0.32 DPS) [dungeon]; Windchaser Cuffs (14429, -0.62 DPS) [world_drop]; Condor Bracers (15864, -0.71 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 21.3 spell_power points (4.16 DPS) | yes | Red Mageweave Gloves (10018, -0.39 DPS) [crafted]; Black Mageweave Gloves (10003, -1.23 DPS) [crafted]; Stormcloth Gloves (10011, -1.43 DPS) [crafted] |
| waist | Gilded Cord (254037) | Tailoring [crafted] | sim-verified (140.4 DPS) | yes | Star Belt (4329, -0.32 DPS) [crafted]; Crimson Silk Belt (7055, -0.55 DPS) [crafted]; Deathmage Sash (10771, -1.79 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.0 spell_power points (4.68 DPS) | yes | Crimson Silk Pantaloons (7062, -1.01 DPS) [crafted]; Stoneweaver Leggings (9407, -2.01 DPS) [dungeon]; Pristine Leggings (253987, -2.18 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.68 DPS) | yes | Gilded Slippers (254001, -2.18 DPS) [crafted]; Acidic Walkers (9454, -2.41 DPS) [dungeon]; Spidersilk Boots (4320, -2.67 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (2.76 DPS) | yes | Reedknot Ring (9622, -1.40 DPS) [quest]; Ring of Power Regulation (273030, -1.59 DPS) [dungeon]; Voodoo Band (1996, -1.63 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.2 spell_power points (1.98 DPS) | yes | Reedknot Ring (9622, -0.62 DPS) [quest]; Ring of Power Regulation (273030, -0.81 DPS) [dungeon]; Voodoo Band (1996, -0.85 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.32 DPS) [dungeon]; Spine of the Devourer (273040, -2.86 DPS) [world] |
| off_hand | Celestial Orb (7515) | Celestial Power [quest] | 15.5 spell_power points (3.02 DPS) | yes | Orb of the Forgotten Seer (7685, -0.29 DPS) [dungeon]; Orb of Mystic Insight (249394, -0.68 DPS) [crafted]; Omega Orb (7749, -1.27 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 207.0 spell_power points (40.35 DPS) | yes | Nether Force Wand (11263, -2.54 DPS) [quest]; Icefury Wand (7514, -2.68 DPS) [quest]; Ragefire Wand (7513, -2.73 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Gilded Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Celestial Orb; ranged: Jaina's Firestarter

No-known-source sample (15 of 386, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 141005213100301531-32023000000000000-000000000000000000)

Set DPS (verified): 233.7. Weights run: 1.3s. Verify run: 1.2s. 502 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.073 ± 0.058, crit=0.261 ± 0.007 per rating point (14 rating = 1%, 3.659 per %), hit=0.644 ± 0.081 per rating point (10 rating = 1%, 6.445 per %), spell_haste=not significant (1.716 ± 0.866), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.662 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 40.5 spell_power points (7.60 DPS) | yes | Dreamweave Circlet (10041, -1.64 DPS) [crafted]; Blood Guard's Dreadweave Hat (220907, -1.86 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -2.53 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 15.6 spell_power points (2.93 DPS) | yes | Horizon Choker (13085, -0.11 DPS) [world_drop]; Mindburst Medallion (11196, -0.59 DPS) [quest]; Triune Amulet (7722, -1.42 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 32.3 spell_power points (6.07 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -2.07 DPS) [vendor]; Inquisitor's Shawl (19507, -2.13 DPS) [dungeon]; Red Mageweave Shoulders (10029, -4.58 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.7 spell_power points (4.07 DPS) | yes | Spritecaster Cape (11623, -0.23 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted]; Big Voodoo Cloak (8216, -1.31 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 40.5 spell_power points (7.60 DPS) | yes | Runecloth Tunic (13857, -2.19 DPS) [crafted]; Robe of the Magi (1716, -2.26 DPS) [world_drop]; Runecloth Robe (13858, -5.18 DPS, sim-verified) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (233.7 DPS) | yes | Forgotten Wraps (9433, -0.31 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.31 DPS) [quest]; Bloodband Bracers (11469, -2.91 DPS, sim-verified) [quest] |
| hands | First Sergeant's Dreadweave Gloves (220908) | PvP rank 9 · First Sergeant · Horde [vendor] | 22.7 spell_power points (4.26 DPS) | yes | Dreamweave Gloves (10019, -0.07 DPS) [crafted]; Red Mageweave Gloves (10018, -0.17 DPS) [crafted]; Runecloth Gloves (13863, -0.19 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 26.4 spell_power points (4.96 DPS) | yes | Deathmage Sash (10771, -0.62 DPS) [dungeon]; Runecloth Belt (13856, -1.04 DPS) [crafted]; Satyrmane Sash (17755, -3.66 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 28.5 spell_power points (5.36 DPS) | yes | Red Mageweave Pants (10009, -0.31 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.17 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.24 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 24.1 spell_power points (4.53 DPS) | yes | Earthen Silk Slippers (254013, -0.02 DPS) [crafted]; Gilded Sandals (254107, -0.65 DPS) [crafted]; Southsea Mojo Boots (20641, -0.81 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (2.89 DPS) | yes | Snake Hoop (6750, -0.75 DPS) [quest]; Ogremind Ring (1993, -1.47 DPS) [world_drop]; Voodoo Band (1996, -1.47 DPS) [world] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.44 DPS) | yes | Snake Hoop (6750, -0.31 DPS) [quest]; Ogremind Ring (1993, -1.03 DPS) [world_drop]; Voodoo Band (1996, -1.03 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+4.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.24 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.28 DPS) [quest]; Mark of the Chosen (17774, -3.65 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.46 DPS) [quest]; Kindling Stave (11750, -1.55 DPS) [dungeon]; Hanzo Sword (8190, -41.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 279.5 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.56 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.25 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: First Sergeant's Dreadweave Gloves; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 502, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 141005213100301531-32023000000000000-053200000000000000)

Set DPS (verified): 389.7. Weights run: 1.4s. Verify run: 3.1s. 1123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.149 ± 0.066, crit=0.411 ± 0.011 per rating point (14 rating = 1%, 5.756 per %), hit=1.115 ± 0.107 per rating point (10 rating = 1%, 11.154 per %), spell_haste=not significant (2.372 ± 1.180), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.638 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (389.7 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Magister's Crown (16686, -18.84 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 23.5 spell_power points (5.22 DPS) | yes | Glowing Eye of Mordresh (10769, -1.62 DPS) [dungeon]; Horizon Choker (13085, -1.64 DPS) [world_drop]; Mindburst Medallion (11196, -2.35 DPS) [quest] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (389.7 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Magister's Mantle (16689, -16.91 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.3 spell_power points (4.96 DPS) | yes | Spritecaster Cape (11623, -0.32 DPS) [dungeon]; Runecloth Cloak (13860, -0.92 DPS) [crafted]; Big Voodoo Cloak (8216, -1.56 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 59.5 spell_power points (13.23 DPS) | yes | Warlord's Silk Raiment (231596, -0.28 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.94 DPS) [pvp]; Magister's Robes (16688, -8.13 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (389.7 DPS) | yes | General's Silk Cuffs (16538, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Magister's Bindings (16683, -17.95 DPS, sim-verified) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 32.0 spell_power points (7.11 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -5.83 DPS, sim-verified) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (389.7 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Magister's Belt (16685, -17.28 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (389.7 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Magister's Leggings (16687, -14.40 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 33.8 spell_power points (7.51 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.29 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -0.95 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 26.0 spell_power points (5.79 DPS) | yes | Glowing Crystal Ring (18402, -2.72 DPS) [quest]; Band of the Unicorn (7553, -2.90 DPS) [world_drop]; Snake Hoop (6750, -3.18 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.7 spell_power points (3.50 DPS) | yes | Glowing Crystal Ring (18402, -0.43 DPS) [quest]; Band of the Unicorn (7553, -0.61 DPS) [world_drop]; Snake Hoop (6750, -0.89 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (389.7 DPS) | yes | Uther's Strength (11302, -0.90 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (389.7 DPS) | yes | Uther's Strength (11302, -0.40 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (389.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.64 DPS) [quest]; Teebu's Blazing Longsword (1728, -49.54 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 351.5 spell_power points (78.09 DPS) | yes | Stormrager (16997, -14.13 DPS) [quest]; Brilliant Wand (249385, -14.73 DPS) [crafted]; Torch of Austen (13004, -21.30 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1123, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 131205213100301531-30000000000000000-053500030000000000)

Set DPS (verified): 667.5. Weights run: 1.6s. Verify run: 2.3s. 1123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.784 ± 0.042, crit=0.515 ± 0.011 per rating point (14 rating = 1%, 7.204 per %), hit=1.208 ± 0.085 per rating point (10 rating = 1%, 12.082 per %), spell_haste=3.538 ± 0.795, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.670 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (667.5 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -7.69 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 19.8 spell_power points (8.81 DPS) | yes | Glowing Eye of Mordresh (10769, -2.92 DPS) [dungeon]; Horizon Choker (13085, -3.94 DPS) [world_drop]; Mindburst Medallion (11196, -4.06 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.4 spell_power points (15.27 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.58 DPS) [pvp]; Rotgrip Mantle (17732, -6.17 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.1 spell_power points (8.46 DPS) | yes | Spritecaster Cape (11623, -0.16 DPS) [dungeon]; Runecloth Cloak (13860, -1.68 DPS) [crafted]; Big Voodoo Cloak (8216, -3.11 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.6 spell_power points (25.14 DPS) | yes | Warlord's Silk Raiment (231596, -1.37 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -6.70 DPS) [pvp]; Acumen Robes (17775, -13.51 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 18.8 spell_power points (8.37 DPS) | yes | General's Silk Cuffs (16538, -1.75 DPS) [pvp]; Magiskull Cuffs (13107, -2.45 DPS) [world_drop]; Dreadmist Bracers (16703, -3.51 DPS, sim-verified) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 31.2 spell_power points (13.87 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -8.75 DPS, sim-verified) [quest] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 23.5 spell_power points (10.42 DPS) | yes | Wisdom of the Timbermaw (19047, +0.00 DPS, sim-verified) [crafted]; Satyrmane Sash (17755, -0.72 DPS) [dungeon]; Dawnspire Cord (12466, -1.14 DPS) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 29.6 spell_power points (13.16 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS) [vendor]; Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.4 spell_power points (13.06 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.25 DPS) [pvp]; First Sergeant's Dreadweave Boots (220909, -1.01 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 23.5 spell_power points (10.43 DPS) | yes | Band of the Unicorn (7553, -4.66 DPS) [world_drop]; Snake Hoop (6750, -6.02 DPS) [quest]; Glowing Crystal Ring (18402, -6.25 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.9 spell_power points (6.18 DPS) | yes | Band of the Unicorn (7553, -0.41 DPS) [world_drop]; Snake Hoop (6750, -1.78 DPS) [quest]; Glowing Crystal Ring (18402, -2.00 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+8.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -2.16 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.09 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.31 DPS) [quest]; Teebu's Blazing Longsword (1728, -64.52 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 178.9 spell_power points (79.42 DPS) | yes | Brilliant Wand (249385, -13.41 DPS) [crafted]; Stormrager (16997, -14.99 DPS) [quest]; Greater Eternal Wand (249237, -22.10 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Philanthropist's Ring; trinket2: Rune of the Guard Captain; ranged: Torch of Light

No-known-source sample (15 of 1123, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

