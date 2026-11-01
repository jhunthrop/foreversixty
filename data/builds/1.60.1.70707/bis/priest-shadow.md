# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 32.4. Weights run: 1.3s. Verify run: 0.8s. 163 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.506 ± 0.010, crit=0.087 ± 0.002 per rating point (14 rating = 1%, 1.214 per %), hit=0.296 ± 0.010 per rating point (10 rating = 1%, 2.957 per %), spell_haste=not significant (0.277 ± 0.112), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Shadow Goggles (4373, -1.24 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.6 spell_power points (0.87 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.46 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Pearl-clasped Cloak (5542, -0.04 DPS) [crafted]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Sanguine Cape (14376, -0.18 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.5 spell_power points (0.69 DPS) | yes | Green Woolen Robe (6243, -0.27 DPS) [crafted]; Green Woolen Vest (2582, -0.32 DPS) [crafted]; Gray Woolen Robe (2585, -0.73 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.5 spell_power points (0.23 DPS) | yes | Mystic's Bracelets (14366, -0.14 DPS) [world_drop]; Repurposed Hair Band (281256, -0.14 DPS) [quest]; Bright Bracers (3647, -0.46 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Gnoll Casting Gloves (892, -0.19 DPS, sim-verified) [world]; Tomb Robber's Gloves (280096, -0.36 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Keller's Girdle (2911, -0.18 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.48 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 9.0 spell_power points (0.82 DPS) | yes | Colorful Kilt (10048, -0.37 DPS) [crafted]; Darkweave Breeches (12987, -0.50 DPS) [world_drop]; Silk-threaded Trousers (1929, -1.04 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.0 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.16 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.46 DPS) [crafted]; Woolen Boots (2583, -0.55 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.0 spell_power points (0.55 DPS) | yes | Volcanic Rock Ring (12053, -0.41 DPS) [world_drop]; The 1 Ring (8350, -0.50 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.0 spell_power points (0.28 DPS) | yes | The 1 Ring (8350, -0.23 DPS) [world]; Volcanic Rock Ring (12053, -0.31 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 5.1 spell_power points (0.46 DPS) | yes | Witching Stave (1484, -0.23 DPS) [dungeon]; Staff of Westfall (2042, -0.23 DPS) [quest]; Lesser Staff of the Spire (1300, -0.38 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 247.7 spell_power points (22.58 DPS) | yes | Skycaller (12984, -0.54 DPS, sim-verified) [world_drop]; Deepblaze (279896, -4.05 DPS) [quest]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 163, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 18851 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (gnome, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 61.3. Weights run: 1.2s. Verify run: 0.8s. 294 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.730 ± 0.010, crit=0.102 ± 0.002 per rating point (14 rating = 1%, 1.433 per %), hit=0.295 ± 0.011 per rating point (10 rating = 1%, 2.945 per %), spell_haste=2.833 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Nightsky Cowl (4039, -0.26 DPS) [world_drop]; Enchanter's Cowl (4322, -0.95 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.4 spell_power points (0.50 DPS) | yes | Crystal Starfire Medallion (5003, -0.17 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.17 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.6 spell_power points (1.79 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Magician's Mantle (12998, -0.46 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.57 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Resilient Cape (14400, -0.07 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.09 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.5 spell_power points (2.12 DPS) | yes | Beguiler Robes (7728, -0.08 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.20 DPS) [dungeon]; Death Speaker Robes (6682, -0.40 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Nightsky Wristbands (6407, -0.53 DPS) [world_drop]; Stonecloth Bindings (14416, -0.61 DPS) [world_drop]; Glowing Magical Bracelets (13106, -0.94 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.1 spell_power points (0.93 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.13 DPS) [dungeon]; Shilly Mitts (9609, -0.13 DPS) [quest] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.2 spell_power points (1.28 DPS) | yes | Crimson Silk Belt (7055, -0.01 DPS) [crafted]; Invoker's Cord (215366, -0.06 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.37 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 12.1 spell_power points (1.39 DPS) | yes | Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Necromancer Leggings (2277, -0.47 DPS) [world_drop]; Gaze Dreamer Pants (6903, -1.17 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.1 spell_power points (1.39 DPS) | yes | Spidersilk Boots (4320, -0.25 DPS) [crafted]; Glinteye Slippers (273024, -0.59 DPS) [dungeon]; Acidic Walkers (9454, -1.02 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.6 spell_power points (1.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon]; Azora's Will (4999, -0.69 DPS) [world_drop]; Minor Channeling Ring (1449, -1.59 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; Azora's Will (4999, -0.27 DPS) [world_drop]; Minor Channeling Ring (1449, -0.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.13 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.84 DPS) [quest]; Glimmering Staff (249392, -3.21 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 289.2 spell_power points (33.19 DPS) | yes | Wand of Mana Concentration (273043, -1.12 DPS, sim-verified) [world]; Greater Mystic Wand (217287, -3.62 DPS) [crafted]; Gravestone Scepter (7001, -4.19 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 294, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 159.3. Weights run: 1.2s. Verify run: 0.9s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.547 ± 0.020, crit=0.243 ± 0.006 per rating point (14 rating = 1%, 3.401 per %), hit=0.339 ± 0.018 per rating point (10 rating = 1%, 3.389 per %), spell_haste=0.697 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 34.3 spell_power points (6.72 DPS) | yes | Thinking Cap (2624, -1.57 DPS) [world]; Craftsman's Monocle (4393, -2.17 DPS) [crafted]; Augural Shroud (2620, -4.67 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.4 spell_power points (3.80 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.98 DPS) [world]; Triune Amulet (7722, -2.23 DPS) [dungeon]; Necklace of Calisea (1714, -3.71 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 27.1 spell_power points (5.31 DPS) | yes | Death Speaker Mantle (6685, -0.80 DPS) [dungeon]; Bloodmage Mantle (7684, -0.82 DPS) [dungeon]; Green Silken Shoulders (7057, -2.74 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 13.7 spell_power points (2.69 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.27 DPS) [dungeon]; Glyphed Cloak (4732, -0.57 DPS) [world_drop] |
| chest | Robe of Power (7054) | Tailoring [crafted] | 32.6 spell_power points (6.38 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of the Magi (1716, -0.25 DPS) [world_drop]; Green Silk Armor (7065, -0.68 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 13.9 spell_power points (2.73 DPS) | yes | Aurora Bracers (4043, -0.30 DPS) [world_drop]; Mistscape Bracers (4045, -0.30 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.30 DPS) [quest] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 26.5 spell_power points (5.19 DPS) | yes | Stormcloth Gloves (10011, -0.77 DPS) [crafted]; Dreamweave Gloves (10019, -1.36 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.50 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.2 spell_power points (5.92 DPS) | yes | Gilded Cord (254037, -1.58 DPS, sim-verified) [crafted]; Razzeric's Customized Seatbelt (6726, -2.28 DPS) [quest]; Teacher's Sash (10747, -2.28 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 32.6 spell_power points (6.38 DPS) | yes | Crimson Silk Pantaloons (7062, -1.12 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.58 DPS) [dungeon]; Pristine Leggings (253987, -2.89 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.70 DPS) | yes | Gilded Slippers (254001, -1.21 DPS) [crafted]; Acidic Walkers (9454, -1.30 DPS) [dungeon]; Spidersilk Boots (4320, -2.12 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.7 spell_power points (3.47 DPS) | yes | Ogremind Ring (1993, -1.35 DPS) [world_drop]; Voodoo Band (1996, -1.35 DPS) [world]; Mindbender Loop (5009, -1.35 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 13.7 spell_power points (2.69 DPS) | yes | Ogremind Ring (1993, -0.57 DPS) [world_drop]; Voodoo Band (1996, -0.57 DPS) [world]; Mindbender Loop (5009, -0.57 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (159.3 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.61 DPS) [dungeon]; Spine of the Devourer (273040, -3.99 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 21.5 spell_power points (4.21 DPS) | yes | Orb of the Forgotten Seer (7685, -1.46 DPS) [dungeon]; Skull of Impending Doom (4984, -2.08 DPS) [quest]; Orb of Mystic Insight (249394, -4.00 DPS, sim-verified) [crafted] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 210.3 spell_power points (41.20 DPS) | yes | Earthen Rod (9381, -3.06 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -5.53 DPS) [dungeon]; Twisted Nether Wand (249144, -6.02 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of Power; wrist: Windchaser Cuffs; hands: Red Mageweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica; ranged: Jaina's Firestarter

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 50 (gnome, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 260.5. Weights run: 1.2s. Verify run: 1.1s. 515 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.766 ± 0.026, crit=0.372 ± 0.009 per rating point (14 rating = 1%, 5.205 per %), hit=0.596 ± 0.029 per rating point (10 rating = 1%, 5.958 per %), spell_haste=5.340 ± 0.309, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 54.3 spell_power points (10.67 DPS) | yes | Knight-Lieutenant's Dreadweave Hat (220889, -2.73 DPS) [vendor]; Papal Fez (9431, -3.20 DPS) [dungeon]; Dreamweave Circlet (10041, -4.84 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 24.7 spell_power points (4.86 DPS) | yes | Glowing Eye of Mordresh (10769, -0.71 DPS) [dungeon]; Mindburst Medallion (11196, -1.60 DPS) [quest]; Necklace of Calisea (1714, -2.43 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 44.8 spell_power points (8.80 DPS) | yes | Inquisitor's Shawl (19507, -2.91 DPS) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -3.08 DPS) [vendor]; Red Mageweave Shoulders (10029, -3.27 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.6 spell_power points (4.83 DPS) | yes | Runecloth Cloak (13860, -0.29 DPS) [crafted]; Big Voodoo Cloak (8216, -0.73 DPS) [crafted]; Imperial Red Cloak (8248, -1.01 DPS) [world_drop] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 54.3 spell_power points (10.67 DPS) | yes | Knight's Dreadweave Vest (220886, -3.47 DPS) [vendor]; Runecloth Tunic (13857, -3.51 DPS) [crafted]; Runecloth Robe (13858, -3.91 DPS, sim-verified) [crafted] |
| wrist | Shizzle's Nozzle Wiper (11917) (or Forgotten Wraps (9433)) | Shizzle's Flyer [quest] | 21.2 spell_power points (4.16 DPS) | yes | Forgotten Wraps (9433, +0.00 DPS) [world_drop]; Bloodband Bracers (11469, -0.06 DPS) [quest]; Imperial Red Bracers (8247, -0.35 DPS) [world_drop] |
| hands | Sergeant Major's Dreadweave Gloves (220890) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 28.9 spell_power points (5.67 DPS) | yes | Red Mageweave Gloves (10018, -0.05 DPS) [crafted]; Runecloth Gloves (13863, -0.20 DPS) [crafted]; Gilded Gloves (254095, -0.24 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 39.6 spell_power points (7.77 DPS) | yes | Satyrmane Sash (17755, -1.55 DPS) [dungeon]; Serenity Belt (13144, -1.87 DPS) [world_drop]; Deathmage Sash (10771, -2.50 DPS, sim-verified) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 38.4 spell_power points (7.54 DPS) | yes | Red Mageweave Pants (10009, -0.63 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -0.71 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.46 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 29.9 spell_power points (5.86 DPS) | yes | Southsea Mojo Boots (20641, -0.48 DPS) [quest]; Gilded Sandals (254107, -0.58 DPS) [crafted]; Earthen Silk Slippers (254013, -1.15 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 18.8 spell_power points (3.70 DPS) | yes | Band of the Unicorn (7553, -1.15 DPS) [world_drop]; Ogremind Ring (1993, -1.27 DPS) [world_drop]; Voodoo Band (1996, -1.27 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 14.8 spell_power points (2.91 DPS) | yes | Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Ogremind Ring (1993, -0.48 DPS) [world_drop]; Voodoo Band (1996, -0.48 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (260.5 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (260.5 DPS) | yes | Mark of the Chosen (17774, -3.87 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (260.5 DPS) | yes | Spellshifter Rod (9527, -1.30 DPS) [quest]; Kindling Stave (11750, -3.93 DPS) [dungeon]; Hammer of the Northern Wind (810, -49.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 267.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.18 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Shizzle's Nozzle Wiper; hands: Sergeant Major's Dreadweave Gloves; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 515, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60 (gnome, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 411.4. Weights run: 1.3s. Verify run: 3.2s. 1200 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.734 ± 0.028, crit=0.571 ± 0.012 per rating point (14 rating = 1%, 7.990 per %), hit=0.844 ± 0.037 per rating point (10 rating = 1%, 8.435 per %), spell_haste=not significant (0.559 ± 0.293), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | 63.0 spell_power points (16.57 DPS) | yes | Dreadmist Mask (16698, +0.00 DPS) [dungeon]; Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 29.3 spell_power points (7.71 DPS) | yes | Opaline Medallion (16623, -1.79 DPS) [quest]; Glowing Eye of Mordresh (10769, -2.23 DPS) [dungeon]; Horizon Choker (13085, -3.72 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (411.4 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, +0.00 DPS) [pvp]; Magister's Mantle (16689, -11.41 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.4 spell_power points (6.42 DPS) | yes | Big Voodoo Cloak (8216, -1.00 DPS) [crafted]; Imperial Red Cloak (8248, -1.40 DPS) [world_drop]; Runecloth Cloak (13860, -2.53 DPS, sim-verified) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (411.4 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Magister's Robes (16688, -11.07 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (411.4 DPS) | yes | Marshal's Satin Bracers (17606, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Magiskull Cuffs (13107, -11.56 DPS, sim-verified) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 42.3 spell_power points (11.12 DPS) | yes | Marshal's Satin Gloves (17608, +0.00 DPS) [vendor]; Marshal's Satin Grips (231617, +0.00 DPS) [vendor]; Magister's Gloves (16684, -6.77 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (411.4 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; Marshal's Satin Sash (17609, +0.00 DPS) [pvp]; Magister's Belt (16685, -11.80 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (411.4 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor]; Magister's Leggings (16687, -10.53 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 40.8 spell_power points (10.73 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 30.1 spell_power points (7.92 DPS) | yes | Philanthropist's Ring (281635, -3.01 DPS) [quest]; Band of the Hierophant (13096, -3.36 DPS) [world_drop]; Frigid Ring (18679, -3.36 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 20.8 spell_power points (5.47 DPS) | yes | Philanthropist's Ring (281635, +0.00 DPS, sim-verified) [quest]; Band of the Hierophant (13096, -0.91 DPS) [world_drop]; Frigid Ring (18679, -0.91 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (411.4 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (411.4 DPS) | yes | Mark of the Chosen (17774, -4.77 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (411.4 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -3.22 DPS) [world_drop] |
| off_hand | Elunarian Sphere (15968) | World drop [world_drop] | 24.3 spell_power points (6.38 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Truesilver Conduit (249455, +0.00 DPS, sim-verified) [crafted] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 298.0 spell_power points (78.34 DPS) | yes | Brilliant Wand (249385, -13.20 DPS) [crafted]; Stormrager (16997, -13.37 DPS) [quest]; Torch of Austen (13004, -21.55 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Glowing Crystal Ring; main_hand: Hand of Edward the Odd; off_hand: Elunarian Sphere; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

### Band 60, raid preset (gnome, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 667.7. Weights run: 1.3s. Verify run: 2.2s. 1200 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.206 ± 0.008, crit=0.665 ± 0.013 per rating point (14 rating = 1%, 9.309 per %), hit=0.822 ± 0.036 per rating point (10 rating = 1%, 8.224 per %), spell_haste=0.739 ± 0.183, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.42 DPS) | yes | Red Mageweave Headband (10033, +0.00 DPS) [crafted]; Field Marshal's Headdress (17602, +0.00 DPS) [vendor]; Field Marshal's Satin Crown (231616, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.1 spell_power points (6.47 DPS) | yes | Glowing Eye of Mordresh (10769, -2.49 DPS) [dungeon]; Triune Amulet (7722, -2.79 DPS) [dungeon]; Mindburst Medallion (11196, -3.14 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.0 spell_power points (13.33 DPS) | yes | Field Marshal's Satin Epaulets (231621, +0.00 DPS) [vendor]; Field Marshal's Satin Mantle (231628, -4.08 DPS) [pvp]; Argent Shoulders (19059, -4.12 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 spell_power points (7.01 DPS) | yes | Runecloth Cloak (13860, -2.32 DPS, sim-verified) [crafted]; Fel Cape (279269, -2.73 DPS) [crafted]; Nightfall Drape (12465, -2.87 DPS) [dungeon] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 25.0 spell_power points (11.50 DPS) | yes | Field Marshal's Satin Vestments (17605, +0.00 DPS) [vendor]; Field Marshal's Satin Robe (231618, +0.00 DPS) [vendor]; Knight's Dreadweave Vest (220886, -11.24 DPS, sim-verified) [vendor] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.1 spell_power points (6.01 DPS) | yes | Arena Wristguards (18709, -1.72 DPS) [world]; Arcane Runed Bracers (4744, -1.87 DPS) [quest]; Spidertank Oilrag (9448, -1.87 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.7 spell_power points (13.65 DPS) | yes | Marshal's Satin Gloves (17608, -1.84 DPS) [vendor]; Marshal's Satin Grips (231617, -1.84 DPS) [vendor]; Earth Warder's Gloves (21318, -5.55 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (7.38 DPS) | yes | Ghostweave Cord (254073, -0.95 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.00 DPS) [crafted]; Star Belt (4329, -1.41 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.96 DPS) | yes | Marshal's Satin Pants (17603, +0.00 DPS) [vendor]; Knight's Dreadweave Leggings (220888, +0.00 DPS) [vendor]; Marshal's Satin Leggings (231619, +0.00 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.04 DPS) | yes | Marshal's Satin Sandals (17607, +0.00 DPS) [vendor]; Marshal's Satin Treads (231620, +0.00 DPS) [vendor]; Omnicast Boots (11822, -0.70 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.4 spell_power points (8.94 DPS) | yes | Philanthropist's Ring (281635, -3.87 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.26 DPS) [quest]; Snake Hoop (6750, -5.71 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.98 DPS) | yes | Philanthropist's Ring (281635, -0.91 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.30 DPS) [quest]; Snake Hoop (6750, -2.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (667.7 DPS) | yes | Uther's Strength (11302, -0.64 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (667.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (667.7 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -13.89 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.66 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -1.43 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.38 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 172.9 spell_power points (79.52 DPS) | yes | Brilliant Wand (249385, -15.14 DPS) [crafted]; Stormrager (16997, -16.36 DPS) [quest]; Greater Eternal Wand (249237, -22.05 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Truefaith Vestments; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Mark of the Chosen; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Torch of Light

No-known-source sample (15 of 1200, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-54200000000000000)

Set DPS (verified): 34.1. Weights run: 1.3s. Verify run: 0.8s. 156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.506 ± 0.010, crit=0.087 ± 0.002 per rating point (14 rating = 1%, 1.214 per %), hit=0.296 ± 0.010 per rating point (10 rating = 1%, 2.957 per %), spell_haste=not significant (0.277 ± 0.112), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Shadow Goggles (4373, -1.28 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.6 spell_power points (0.87 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.51 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.36 DPS) | yes | Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Sanguine Cape (14376, -0.18 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.25 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.5 spell_power points (0.69 DPS) | yes | Green Woolen Robe (6243, -0.27 DPS) [crafted]; Green Woolen Vest (2582, -0.32 DPS) [crafted]; Gray Woolen Robe (2585, -0.76 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.0 spell_power points (0.27 DPS) | yes | Mindthrust Bracers (1974, -0.04 DPS) [dungeon]; Bright Bracers (3647, -0.09 DPS) [world_drop]; Featherbead Bracers (15452, -0.20 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.64 DPS) | yes | Pristine Gloves (253913, -0.14 DPS) [crafted]; Gnoll Casting Gloves (892, -0.18 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.27 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.0 spell_power points (0.55 DPS) | yes | Keller's Girdle (2911, -0.18 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.23 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.62 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 9.0 spell_power points (0.82 DPS) | yes | Colorful Kilt (10048, -0.37 DPS) [crafted]; Embroidered Leggings (270011, -0.46 DPS) [quest]; Silk-threaded Trousers (1929, -1.22 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.0 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.27 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.46 DPS) [crafted]; Woolen Boots (2583, -0.55 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.0 spell_power points (0.28 DPS) | yes | Volcanic Rock Ring (12053, -0.14 DPS) [world_drop]; The 1 Ring (8350, -0.23 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.5 spell_power points (0.23 DPS) | yes | Volcanic Rock Ring (12053, -0.09 DPS) [world_drop]; The 1 Ring (8350, -0.18 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 5.1 spell_power points (0.46 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Lesser Staff of the Spire (1300, -0.18 DPS) [world]; Clear Crystal Rod (16894, -0.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 247.7 spell_power points (22.58 DPS) | yes | Skycaller (12984, -1.16 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.48 DPS) [quest]; Greater Magic Wand (11288, -4.90 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 156, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18862 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (undead, 000000000000000000-00000000000000000-54342001200000000)

Set DPS (verified): 62.6. Weights run: 1.2s. Verify run: 0.9s. 286 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.730 ± 0.010, crit=0.102 ± 0.002 per rating point (14 rating = 1%, 1.433 per %), hit=0.295 ± 0.011 per rating point (10 rating = 1%, 2.945 per %), spell_haste=2.833 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (62.6 DPS) | yes | Silk Headband (7050, -0.23 DPS) [crafted]; Nightsky Cowl (4039, -0.26 DPS) [world_drop]; Enchanter's Cowl (4322, -2.13 DPS, sim-verified) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 4.4 spell_power points (0.50 DPS) | yes | Crystal Starfire Medallion (5003, -0.17 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.17 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 15.6 spell_power points (1.79 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Magician's Mantle (12998, -0.46 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.57 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Resilient Cape (14400, -0.07 DPS) [world_drop]; Pearl-clasped Cloak (5542, -0.09 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 18.5 spell_power points (2.12 DPS) | yes | Beguiler Robes (7728, +0.00 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.20 DPS) [dungeon]; Death Speaker Robes (6682, -0.40 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.03 DPS) | yes | Nightsky Wristbands (6407, -0.53 DPS) [world_drop]; Featherbead Bracers (15452, -0.61 DPS) [quest]; Glowing Magical Bracelets (13106, -0.64 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 9.6 spell_power points (1.11 DPS) | yes | Truefaith Gloves (7049, +0.00 DPS, sim-verified) [crafted]; Serpent Gloves (5970, -0.30 DPS) [dungeon]; Pristine Gloves (253913, -0.40 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 11.2 spell_power points (1.28 DPS) | yes | Crimson Silk Belt (7055, -0.01 DPS) [crafted]; Warsong Sash (16975, -0.02 DPS) [quest]; Invoker's Cord (215366, -0.06 DPS) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 12.1 spell_power points (1.39 DPS) | yes | Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Necromancer Leggings (2277, -0.47 DPS) [world_drop]; Gaze Dreamer Pants (6903, -0.80 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.1 spell_power points (1.39 DPS) | yes | Spidersilk Boots (4320, -0.25 DPS) [crafted]; Glinteye Slippers (273024, -0.59 DPS) [dungeon]; Acidic Walkers (9454, -1.06 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.6 spell_power points (1.11 DPS) | yes | Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon]; Solid Gold Band (279845, -0.65 DPS) [quest]; Azora's Will (4999, -0.69 DPS) [world_drop] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (0.69 DPS) | yes | Solid Gold Band (279845, -0.23 DPS) [quest]; Azora's Will (4999, -0.27 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.71 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) (or Tendonscraper Dagger (273053)) | Mana Devourer [world] | 36.0 spell_power points (4.13 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Glimmering Staff (249392, -3.21 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -3.29 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 289.2 spell_power points (33.19 DPS) | yes | Wand of Mana Concentration (273043, -1.56 DPS, sim-verified) [world]; Unstable Power Core (279847, -3.48 DPS) [quest]; Greater Mystic Wand (217287, -3.62 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Belt of Arugal; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer; ranged: Starfaller

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (undead, 000000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 148.4. Weights run: 1.2s. Verify run: 0.9s. 388 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.547 ± 0.020, crit=0.243 ± 0.006 per rating point (14 rating = 1%, 3.401 per %), hit=0.339 ± 0.018 per rating point (10 rating = 1%, 3.389 per %), spell_haste=0.697 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 34.3 spell_power points (6.72 DPS) | yes | Thinking Cap (2624, -1.57 DPS) [world]; Craftsman's Monocle (4393, -2.17 DPS) [crafted]; Augural Shroud (2620, -3.14 DPS, sim-verified) [world] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 19.4 spell_power points (3.80 DPS) | yes | Archmagister's Faceted Pendant (273049, -1.98 DPS) [world]; Triune Amulet (7722, -2.23 DPS) [dungeon]; Necklace of Calisea (1714, -3.25 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 27.1 spell_power points (5.31 DPS) | yes | Death Speaker Mantle (6685, -0.80 DPS) [dungeon]; Bloodmage Mantle (7684, -0.82 DPS) [dungeon]; Green Silken Shoulders (7057, -1.84 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 13.7 spell_power points (2.69 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Blackforge Cape (6424, -0.27 DPS) [dungeon]; Glyphed Cloak (4732, -0.57 DPS) [world_drop] |
| chest | Robe of Power (7054) | Tailoring [crafted] | 32.6 spell_power points (6.38 DPS) | yes | Dreamweave Vest (10021, -0.13 DPS) [crafted]; Robe of the Magi (1716, -0.25 DPS) [world_drop]; Green Silk Armor (7065, -0.68 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 16.4 spell_power points (3.21 DPS) | yes | Windchaser Cuffs (14429, -0.48 DPS) [world_drop]; Mistscape Bracers (4045, -0.78 DPS) [dungeon]; Enchanted Stonecloth Bracers (4979, -0.78 DPS) [quest] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 26.5 spell_power points (5.19 DPS) | yes | Dreamweave Gloves (10019, -0.45 DPS) [crafted]; Stormcloth Gloves (10011, -0.77 DPS) [crafted]; Gilded Handwraps (254021, -1.50 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 30.2 spell_power points (5.92 DPS) | yes | Gilded Cord (254037, -1.93 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -2.28 DPS) [quest]; Mistscape Sash (4736, -2.58 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 32.6 spell_power points (6.38 DPS) | yes | Crimson Silk Pantaloons (7062, -1.29 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -2.58 DPS) [dungeon]; Pristine Leggings (253987, -2.89 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.70 DPS) | yes | Gilded Slippers (254001, -1.21 DPS) [crafted]; Acidic Walkers (9454, -1.30 DPS) [dungeon]; Boots of the Maharishi (9658, -1.97 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.7 spell_power points (3.47 DPS) | yes | Ogremind Ring (1993, -1.35 DPS) [world_drop]; Voodoo Band (1996, -1.35 DPS) [world]; Mindbender Loop (5009, -1.35 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 13.7 spell_power points (2.69 DPS) | yes | Ogremind Ring (1993, -0.57 DPS) [world_drop]; Voodoo Band (1996, -0.57 DPS) [world]; Mindbender Loop (5009, -0.57 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (148.4 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.61 DPS) [dungeon]; Skullbreaker (17039, -3.65 DPS) [quest] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 16.3 spell_power points (3.19 DPS) | yes | Orb of the Forgotten Seer (7685, -0.45 DPS) [dungeon]; Witch's Finger (16887, -0.77 DPS) [quest]; Skull of Impending Doom (4984, -1.07 DPS) [quest] |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 210.3 spell_power points (41.20 DPS) | yes | Earthen Rod (9381, -1.54 DPS, sim-verified) [dungeon]; Umbral Wand (5216, -5.53 DPS) [dungeon]; Twisted Nether Wand (249144, -6.02 DPS) [crafted] |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of Power; wrist: Radiant Silver Bracers; hands: Red Mageweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Mystic Insight; ranged: Jaina's Firestarter

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16336 Sergeant Major's Cape

### Band 50 (undead, 523000000000000000-00000000000000000-54342001201300051)

Set DPS (verified): 245.8. Weights run: 1.2s. Verify run: 1.1s. 507 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.766 ± 0.026, crit=0.372 ± 0.009 per rating point (14 rating = 1%, 5.205 per %), hit=0.596 ± 0.029 per rating point (10 rating = 1%, 5.958 per %), spell_haste=5.340 ± 0.309, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 54.3 spell_power points (10.67 DPS) | yes | Blood Guard's Dreadweave Hat (220907, -2.73 DPS) [vendor]; Papal Fez (9431, -3.20 DPS) [dungeon]; Dreamweave Circlet (10041, -4.28 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 24.7 spell_power points (4.86 DPS) | yes | Glowing Eye of Mordresh (10769, -0.71 DPS) [dungeon]; Mindburst Medallion (11196, -1.60 DPS) [quest]; Necklace of Calisea (1714, -2.43 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 44.8 spell_power points (8.80 DPS) | yes | Inquisitor's Shawl (19507, -2.91 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -3.08 DPS) [vendor]; Red Mageweave Shoulders (10029, -4.00 DPS, sim-verified) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 27.9 spell_power points (5.48 DPS) | yes | Spritecaster Cape (11623, -0.65 DPS) [dungeon]; Runecloth Cloak (13860, -0.94 DPS) [crafted]; Big Voodoo Cloak (8216, -1.37 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 54.3 spell_power points (10.67 DPS) | yes | Stone Guard's Dreadweave Vest (220904, -3.47 DPS) [vendor]; Runecloth Tunic (13857, -3.51 DPS) [crafted]; Runecloth Robe (13858, -4.59 DPS, sim-verified) [crafted] |
| wrist | Shizzle's Nozzle Wiper (11917) (or Forgotten Wraps (9433)) | Shizzle's Flyer [quest] | 21.2 spell_power points (4.16 DPS) | yes | Forgotten Wraps (9433, +0.00 DPS) [world_drop]; Bloodband Bracers (11469, -0.06 DPS) [quest]; Imperial Red Bracers (8247, -0.35 DPS) [world_drop] |
| hands | First Sergeant's Dreadweave Gloves (220908) | PvP rank 9 · First Sergeant · Horde [vendor] | 28.9 spell_power points (5.67 DPS) | yes | Red Mageweave Gloves (10018, -0.05 DPS) [crafted]; Runecloth Gloves (13863, -0.20 DPS) [crafted]; Gilded Gloves (254095, -0.24 DPS) [crafted] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 39.6 spell_power points (7.77 DPS) | yes | Satyrmane Sash (17755, -1.55 DPS) [dungeon]; Serenity Belt (13144, -1.87 DPS) [world_drop]; Deathmage Sash (10771, -3.01 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | PvP rank 10 · Stone Guard · Horde [vendor] | 38.4 spell_power points (7.54 DPS) | yes | Kilt of the Atal'ai Prophet (10807, -0.71 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.46 DPS) [crafted]; Red Mageweave Pants (10009, -2.07 DPS, sim-verified) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | PvP rank 9 · First Sergeant · Horde [vendor] | 29.9 spell_power points (5.86 DPS) | yes | Gilded Sandals (254107, -0.58 DPS) [crafted]; Earthen Silk Slippers (254013, -1.15 DPS) [crafted]; Southsea Mojo Boots (20641, -1.83 DPS, sim-verified) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 18.8 spell_power points (3.70 DPS) | yes | Band of the Unicorn (7553, -1.15 DPS) [world_drop]; Ogremind Ring (1993, -1.27 DPS) [world_drop]; Voodoo Band (1996, -1.27 DPS) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 14.8 spell_power points (2.91 DPS) | yes | Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Ogremind Ring (1993, -0.48 DPS) [world_drop]; Voodoo Band (1996, -0.48 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (245.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (245.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (245.8 DPS) | yes | Spellshifter Rod (9527, -1.30 DPS) [quest]; Kindling Stave (11750, -3.93 DPS) [dungeon]; Hammer of the Northern Wind (810, -42.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 267.4 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.18 DPS) [crafted] |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Shizzle's Nozzle Wiper; hands: First Sergeant's Dreadweave Gloves; waist: Dawnspire Cord; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff; ranged: Pyric Caduceus

No-known-source sample (15 of 507, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60 (undead, 524010031130000000-00000000000000000-54342001201300051)

Set DPS (verified): 409.2. Weights run: 1.3s. Verify run: 3.4s. 1193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=1.734 ± 0.028, crit=0.571 ± 0.012 per rating point (14 rating = 1%, 7.990 per %), hit=0.844 ± 0.037 per rating point (10 rating = 1%, 8.435 per %), spell_haste=not significant (0.559 ± 0.293), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (409.2 DPS) | yes | Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor]; Magister's Crown (16686, -6.17 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 29.3 spell_power points (7.71 DPS) | yes | Glowing Eye of Mordresh (10769, -2.23 DPS) [dungeon]; Mindburst Medallion (11196, -3.40 DPS) [quest]; Horizon Choker (13085, -4.27 DPS, sim-verified) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (409.2 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Warlord's Satin Mantle (231631, +0.00 DPS) [pvp]; Magister's Mantle (16689, -6.18 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 27.6 spell_power points (7.26 DPS) | yes | Spritecaster Cape (11623, -0.84 DPS) [dungeon]; Runecloth Cloak (13860, -1.24 DPS) [crafted]; Big Voodoo Cloak (8216, -1.84 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (409.2 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Magister's Robes (16688, -5.47 DPS, sim-verified) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (409.2 DPS) | yes | General's Satin Bracers (17619, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Magiskull Cuffs (13107, -5.63 DPS, sim-verified) [world_drop] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 42.3 spell_power points (11.12 DPS) | yes | General's Satin Gloves (17620, +0.00 DPS) [vendor]; General's Satin Grips (231613, +0.00 DPS) [vendor]; Magister's Gloves (16684, -6.74 DPS, sim-verified) [dungeon] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (409.2 DPS) | yes | Dawnspire Cord (12466, +0.00 DPS) [dungeon]; General's Satin Cinch (17621, +0.00 DPS) [pvp]; Magister's Belt (16685, -6.57 DPS, sim-verified) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (409.2 DPS) | yes | General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, +0.00 DPS) [vendor]; Magister's Leggings (16687, -4.39 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 40.8 spell_power points (10.73 DPS) | yes | Mooncloth Boots (15802, +0.00 DPS) [crafted]; General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 30.1 spell_power points (7.92 DPS) | yes | Philanthropist's Ring (281635, -3.01 DPS) [quest]; Band of the Hierophant (13096, -3.36 DPS) [world_drop]; Frigid Ring (18679, -3.36 DPS) [world] |
| finger2 | Glowing Crystal Ring (18402) | Uncovering Past Secrets [quest] | 20.8 spell_power points (5.47 DPS) | yes | Philanthropist's Ring (281635, -0.56 DPS) [quest]; Band of the Hierophant (13096, -0.91 DPS) [world_drop]; Frigid Ring (18679, -0.91 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (409.2 DPS) | yes | Uther's Strength (11302, -0.42 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (409.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (409.2 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -3.22 DPS) [world_drop] |
| off_hand | Elunarian Sphere (15968) | World drop [world_drop] | 24.3 spell_power points (6.38 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -0.11 DPS) [crafted] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 298.0 spell_power points (78.34 DPS) | yes | Brilliant Wand (249385, -13.20 DPS) [crafted]; Stormrager (16997, -13.37 DPS) [quest]; Torch of Austen (13004, -21.55 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Gloves of Spell Mastery; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Glowing Crystal Ring; main_hand: Hand of Edward the Odd; off_hand: Elunarian Sphere; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

### Band 60, raid preset (undead, 325000031130000000-00000000000000000-52352101201300251)

Set DPS (verified): 669.1. Weights run: 1.3s. Verify run: 2.1s. 1193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.206 ± 0.008, crit=0.665 ± 0.013 per rating point (14 rating = 1%, 9.309 per %), hit=0.822 ± 0.036 per rating point (10 rating = 1%, 8.224 per %), spell_haste=0.739 ± 0.183, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.42 DPS) | yes | Red Mageweave Headband (10033, +0.00 DPS) [crafted]; Warlord's Satin Cowl (17623, +0.00 DPS) [vendor]; Warlord's Satin Crown (231615, +0.00 DPS) [vendor] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.1 spell_power points (6.47 DPS) | yes | Glowing Eye of Mordresh (10769, -2.49 DPS) [dungeon]; Triune Amulet (7722, -2.79 DPS) [dungeon]; Mindburst Medallion (11196, -3.14 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 29.0 spell_power points (13.33 DPS) | yes | Warlord's Satin Epaulets (231611, +0.00 DPS) [vendor]; Argent Shoulders (19059, -2.82 DPS, sim-verified) [crafted]; Warlord's Satin Mantle (231631, -4.08 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 spell_power points (7.01 DPS) | yes | Deep Woodlands Cloak (19121, -0.64 DPS) [quest]; Runecloth Cloak (13860, -2.11 DPS) [crafted]; Fel Cape (279269, -2.73 DPS) [crafted] |
| chest | Truefaith Vestments (14154) | Tailoring [crafted] | 25.0 spell_power points (11.50 DPS) | yes | Warlord's Satin Robes (231612, +0.00 DPS) [pvp]; Warlord's Satin Tunic (231632, +0.00 DPS) [vendor]; Stone Guard's Dreadweave Vest (220904, -10.84 DPS, sim-verified) [vendor] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.1 spell_power points (6.01 DPS) | yes | Arena Wristguards (18709, -1.72 DPS) [world]; Spidertank Oilrag (9448, -1.87 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.12 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.7 spell_power points (13.65 DPS) | yes | General's Satin Gloves (17620, -1.84 DPS) [vendor]; General's Satin Grips (231613, -1.84 DPS) [vendor]; Earth Warder's Gloves (21318, -5.65 DPS, sim-verified) [quest] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 16.1 spell_power points (7.38 DPS) | yes | Ghostweave Cord (254073, -0.95 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.00 DPS) [crafted]; Star Belt (4329, -1.41 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.0 spell_power points (11.96 DPS) | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS) [vendor]; General's Satin Leggings (231614, +0.00 DPS) [pvp]; General's Satin Legguards (231634, -0.77 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.04 DPS) | yes | General's Satin Boots (17618, +0.00 DPS) [vendor]; General's Satin Treads (231610, +0.00 DPS) [vendor]; Omnicast Boots (11822, -0.70 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.4 spell_power points (8.94 DPS) | yes | Philanthropist's Ring (281635, -3.87 DPS) [quest]; Snake Hoop (6750, -5.71 DPS) [quest]; Reedknot Ring (9622, -5.72 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (5.98 DPS) | yes | Philanthropist's Ring (281635, -0.91 DPS) [quest]; Snake Hoop (6750, -2.75 DPS) [quest]; Reedknot Ring (9622, -2.76 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (669.1 DPS) | yes | Uther's Strength (11302, -0.64 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (669.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (669.1 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -13.89 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.66 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -1.43 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.35 DPS, sim-verified) [dungeon] |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 172.9 spell_power points (79.52 DPS) | yes | Brilliant Wand (249385, -15.14 DPS) [crafted]; Stormrager (16997, -16.36 DPS) [quest]; Greater Eternal Wand (249237, -22.05 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Spritecaster Cape; chest: Truefaith Vestments; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Satyrmane Sash; legs: Ghostweave Pants; feet: Earthen Silk Slippers; finger1: Maiden's Circle; finger2: Band of the Unicorn; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Torch of Light

No-known-source sample (15 of 1193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10782 Hakkari Shroud

