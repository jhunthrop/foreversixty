# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 33.7. Weights run: 1.1s. Verify run: 0.8s. 185 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.385 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.003 per %), hit=0.225 ± 0.007 per rating point (10 rating = 1%, 2.254 per %), spell_haste=not significant (0.144 ± 0.108), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.222 ± 0.079), mp5=not significant (0.235 ± 0.079)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (1.19 DPS) | yes | Pristine Circlet (253949, -0.78 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.84 DPS) [crafted]; Flying Tiger Goggles (4368, -1.09 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.4 spell_power points (0.05 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 spell_power points (0.91 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.57 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.84 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.9 spell_power points (0.53 DPS) | yes | Pearl-clasped Cloak (5542, -0.19 DPS) [crafted]; Black Whelp Cloak (7283, -0.20 DPS) [crafted]; Sanguine Cape (14376, -0.36 DPS) [world_drop] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 11.6 spell_power points (1.25 DPS) | yes | Gray Woolen Robe (2585, -0.61 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.62 DPS, sim-verified) [crafted]; Scholarly Robes (2034, -0.64 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.6 spell_power points (0.28 DPS) | yes | Repurposed Hair Band (281256, -0.15 DPS) [quest]; Windsong Bangles (263336, -0.18 DPS) [quest]; Bright Bracers (3647, -0.68 DPS, sim-verified) [world_drop] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 8.2 spell_power points (0.88 DPS) | yes | Serpent Gloves (5970, -0.23 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Pristine Gloves (253913, -0.32 DPS) [crafted] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 8.5 spell_power points (0.92 DPS) | yes | Novice Arcanist's Sash (253885, -0.37 DPS) [crafted]; Pristine Sash (253925, -0.47 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.51 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 14.2 spell_power points (1.53 DPS) | yes | Silk-threaded Trousers (1929, -0.78 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.78 DPS, sim-verified) [crafted]; Colorful Kilt (10048, -0.87 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.92 DPS) | yes | Silky Boots (253891, -0.33 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.47 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.62 DPS) | yes | Band of Purification (12996, -0.48 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.50 DPS) [world_drop]; Furlfeather Ring (286745, -0.52 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.3 spell_power points (0.25 DPS) | yes | Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; Furlfeather Ring (286745, -0.14 DPS) [world]; Band of Purification (12996, -0.62 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | sim-verified (33.7 DPS) | yes | Lesser Staff of the Spire (1300, -0.10 DPS) [world]; Staff of the Blessed Seer (2271, -0.11 DPS) [dungeon]; Twisted Chanter's Staff (890, -1.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 206.6 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.57 DPS, sim-verified) [world_drop]; Deepblaze (279896, -3.78 DPS) [quest]; Sizzle Stick (8071, -4.44 DPS) [quest] |

**New at 20:** head: Silky Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of Westfall; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 185, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 61.7. Weights run: 1.0s. Verify run: 0.9s. 323 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.397 ± 0.009, crit=0.112 ± 0.004 per rating point (14 rating = 1%, 1.571 per %), hit=0.233 ± 0.009 per rating point (10 rating = 1%, 2.334 per %), spell_haste=not significant (0.485 ± 0.151), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.083 ± 0.062), mp5=not significant (0.212 ± 0.091)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (2.13 DPS) | yes | Holy Shroud (2721, -0.43 DPS, sim-verified) [world_drop]; Silky Circlet (253951, -0.46 DPS) [crafted]; Azure Silk Hood (7048, -0.61 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.4 spell_power points (0.36 DPS) | yes | Kaleidoscope Chain (13084, -0.07 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Pendant of Myzrael (4614, -0.29 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (1.91 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Invoker's Mantle (215365, -0.55 DPS) [crafted]; Magician's Mantle (12998, -0.61 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | 10.2 spell_power points (1.56 DPS) | yes | Hillman's Cloak (3719, -0.80 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.90 DPS) [crafted]; Vine Pruner's Cloak (279835, -1.04 DPS, sim-verified) [quest] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 16.1 spell_power points (2.45 DPS) | yes | Beguiler Robes (7728, -0.26 DPS) [dungeon]; Green Silk Armor (7065, -0.30 DPS) [crafted]; Tree Bark Jacket (1486, -0.47 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.6 spell_power points (1.47 DPS) | yes | Nightsky Wristbands (6407, -1.07 DPS) [world_drop]; Mindthrust Bracers (1974, -1.07 DPS) [dungeon]; Glowing Magical Bracelets (13106, -1.37 DPS, sim-verified) [world_drop] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.2 spell_power points (1.56 DPS) | yes | Silky Gloves (253915, -0.31 DPS) [crafted]; Serpent Gloves (5970, -0.49 DPS) [dungeon]; Truefaith Gloves (7049, -0.62 DPS) [crafted] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 10.6 spell_power points (1.61 DPS) | yes | Crimson Silk Belt (7055, -0.28 DPS) [crafted]; Silky Sash (253927, -0.31 DPS) [crafted]; Invoker's Cord (215366, -0.47 DPS, sim-verified) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 17.2 spell_power points (2.62 DPS) | yes | Filigreed Silky Leggings (253939, -0.66 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -0.68 DPS) [dungeon]; Azure Silk Pants (7046, -0.73 DPS) [crafted] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 16.1 spell_power points (2.45 DPS) | yes | Spidersilk Boots (4320, -1.14 DPS) [crafted]; Gilded Slippers (254001, -1.15 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.16 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 spell_power points (1.21 DPS) | yes | Minor Channeling Ring (1449, -0.33 DPS) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Runemender's Seal (273031, -0.67 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.6 spell_power points (1.01 DPS) | yes | Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Runemender's Seal (273031, -0.46 DPS) [world]; Minor Channeling Ring (1449, -0.76 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | sim-verified (61.7 DPS) | yes | Twisted Chanter's Staff (890, -0.06 DPS) [world_drop]; Cryptbone Staff (2013, -0.25 DPS) [world]; Icicle Rod (2950, -2.24 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 217.6 spell_power points (33.10 DPS) | yes | Wand of Mana Concentration (273043, -1.62 DPS) [world]; Gyromatic Icemaker (9489, -2.74 DPS) [dungeon]; Greater Mystic Wand (217287, -3.34 DPS) [crafted] |

**New at 30:** head: Filigreed Silky Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Belt of Arugal; legs: Silky Leggings; feet: Frothing Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Glimmering Staff; ranged: Starfaller

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 000000000000000000-00000000000000000-253510130100301231)

Set DPS (verified): 136.0. Weights run: 1.1s. Verify run: 1.0s. 421 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.693 ± 0.021, crit=0.169 ± 0.007 per rating point (14 rating = 1%, 2.363 per %), hit=0.388 ± 0.024 per rating point (10 rating = 1%, 3.878 per %), spell_haste=not significant (0.959 ± 0.332), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.672 ± 0.256), mp5=not significant (1.093 ± 0.307)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 25.8 spell_power points (4.71 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.88 DPS) [crafted]; Electromagnetic Gigaflux Reactivator (9492, -1.14 DPS) [dungeon]; Whitemane's Chapeau (7720, -1.19 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 12.7 spell_power points (2.32 DPS) | yes | Glowing Eye of Mordresh (10769, +0.00 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -0.57 DPS) [dungeon]; Amberglow Talisman (10824, -1.09 DPS) [quest] |
| shoulder | Mantle of Doan (7712) | Scarlet Monastery: Arcanist Doan [dungeon] | 16.5 spell_power points (3.01 DPS) | yes | Inquisitor's Shawl (19507, +0.00 DPS, sim-verified) [dungeon]; Green Silken Shoulders (7057, -0.17 DPS) [crafted]; Bloodmage Mantle (7684, -0.24 DPS) [dungeon] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | 12.0 spell_power points (2.19 DPS) | yes | Icy Cloak (4327, -0.18 DPS) [crafted]; Frost Leather Cloak (7377, -0.19 DPS) [crafted]; Long Silken Cloak (4326, -0.46 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 33.6 spell_power points (6.13 DPS) | yes | Robe of the Magi (1716, -0.75 DPS) [world_drop]; Robe of Power (7054, -1.08 DPS) [crafted]; Beguiler Robes (7728, -2.00 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.3 spell_power points (2.24 DPS) | yes | Arcane Runed Bracers (4744, -0.60 DPS) [quest]; Earthen Silk Cuffs (254019, -0.90 DPS) [crafted]; Condor Bracers (15864, -0.96 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 25.5 spell_power points (4.64 DPS) | yes | Frothing Handwraps (254023, -1.23 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.38 DPS) [crafted]; Black Mageweave Gloves (10003, -1.91 DPS) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 23.6 spell_power points (4.30 DPS) | yes | Deathmage Sash (10771, -1.13 DPS) [dungeon]; Star Belt (4329, -1.44 DPS) [crafted]; Gilded Cord (254037, -1.71 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.3 spell_power points (4.07 DPS) | yes | Silky Leggings (253989, -0.02 DPS) [crafted]; Stoneweaver Leggings (9407, -0.31 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.78 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.7 spell_power points (5.23 DPS) | yes | Frothing Slippers (254003, -1.14 DPS, sim-verified) [crafted]; Furen's Boots (13100, -2.39 DPS) [world_drop]; Gilded Slippers (254001, -2.58 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (2.45 DPS) | yes | Ring of Power Regulation (273030, -0.76 DPS) [dungeon]; Welken Ring (5011, -0.97 DPS) [world_drop]; Ring of Forlorn Spirits (2043, -1.00 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (1.73 DPS) | yes | Ring of Power Regulation (273030, -0.03 DPS) [dungeon]; Welken Ring (5011, -0.24 DPS) [world_drop]; Ring of Forlorn Spirits (2043, -0.27 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Illusionary Rod (7713, -2.77 DPS) [dungeon]; Staff of the Friar (3415, -2.96 DPS) [dungeon]; Icicle Rod (2950, -16.37 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Icefury Wand (7514) | Mage's Wand [quest] | sim-verified (136.0 DPS) | yes | Nether Force Wand (11263, -1.49 DPS) [quest]; Ragefire Wand (7513, -1.69 DPS) [quest]; Jaina's Firestarter (13064, -2.12 DPS, sim-verified) [world_drop] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Mantle of Doan; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Icefury Wand

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 230005000000000000-00000000000000000-253510130100301231)

Set DPS (verified): 186.0. Weights run: 1.1s. Verify run: 1.1s. 537 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.030 ± 0.038, crit=0.279 ± 0.012 per rating point (14 rating = 1%, 3.902 per %), hit=0.696 ± 0.043 per rating point (10 rating = 1%, 6.964 per %), spell_haste=3.309 ± 0.673, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (1.094 ± 0.483), mp5=not significant (1.857 ± 0.604)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 44.4 spell_power points (7.82 DPS) | yes | Red Mageweave Headband (10033, -0.85 DPS) [crafted]; Papal Fez (9431, -1.60 DPS) [dungeon]; Gemburst Circlet (10751, -2.27 DPS, sim-verified) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 18.8 spell_power points (3.31 DPS) | yes | Mindburst Medallion (11196, -0.20 DPS) [quest]; Triune Amulet (7722, -0.55 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.63 DPS) [dungeon] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 31.8 spell_power points (5.60 DPS) | yes | Rotgrip Mantle (17732, -0.05 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.41 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.65 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.7 spell_power points (4.71 DPS) | yes | Imperial Red Cloak (8248, -1.94 DPS) [world_drop]; Runecloth Cloak (13860, -2.02 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -2.20 DPS) [crafted] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 45.7 spell_power points (8.05 DPS) | yes | Dreamweave Vest (10021, -0.55 DPS) [crafted]; Acumen Robes (17775, -1.08 DPS) [quest]; Frostweave Robe (13868, -1.71 DPS) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 25.7 spell_power points (4.52 DPS) | yes | Bloodband Bracers (11469, -1.43 DPS) [quest]; Nethergeld Cuffs (254061, -1.68 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -1.77 DPS) [quest] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 38.0 spell_power points (6.70 DPS) | yes | Dreamweave Gloves (10019, -1.46 DPS) [crafted]; Gilded Gloves (254095, -1.58 DPS) [crafted]; Runecloth Gloves (13863, -1.61 DPS, sim-verified) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 28.8 spell_power points (5.07 DPS) | yes | Frothing Waistcord (254083, -0.09 DPS) [crafted]; Ghostweave Cord (254073, -0.49 DPS) [crafted]; Dawnspire Cord (12466, -0.57 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 41.2 spell_power points (7.26 DPS) | yes | Senior Designer's Pantaloons (11841, -0.41 DPS) [dungeon]; Dalewind Trousers (13008, -0.81 DPS) [world_drop]; Venomshroud Leggings (14444, -2.23 DPS) [world_drop] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 31.7 spell_power points (5.57 DPS) | yes | Frothing Sandals (254109, -0.42 DPS) [crafted]; Vinerot Sandals (17748, -0.96 DPS) [dungeon]; Furen's Boots (13100, -1.22 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 19.3 spell_power points (3.40 DPS) | yes | Choking Band (11868, -0.90 DPS) [quest]; Band of the Unicorn (7553, -1.11 DPS) [world_drop]; Coldwater Ring (4550, -1.12 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (2.67 DPS) | yes | Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Coldwater Ring (4550, -0.39 DPS) [quest]; Choking Band (11868, -1.87 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (186.0 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.05 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (186.0 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.48 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (186.0 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS, sim-verified) [quest]; Kindling Stave (11750, -0.57 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.68 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 301.4 spell_power points (53.08 DPS) | yes | Noxious Shooter (17745, -2.20 DPS) [dungeon]; Wand of Allistarj (13065, -4.66 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.92 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Horizon Choker; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Ghostweave Vest; wrist: Netherfroth Cuffs; hands: Frothing Gloves; legs: Kilt of the Atal'ai Prophet; finger1: Eye of Adaegus; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 537, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 230005000000000000-23023000000000000-253510130100301231)

Set DPS (verified): 275.1. Weights run: 1.2s. Verify run: 2.8s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.812 ± 0.052, crit=0.425 ± 0.016 per rating point (14 rating = 1%, 5.943 per %), hit=1.177 ± 0.063 per rating point (10 rating = 1%, 11.768 per %), spell_haste=not significant (2.130 ± 0.830), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.765 ± 0.563), mp5=not significant (2.511 ± 0.707)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (275.1 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Crown of the Penitent (13216, -4.22 DPS, sim-verified) [quest] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.1 spell_power points (3.72 DPS) | yes | Mindburst Medallion (11196, -1.00 DPS) [quest]; Horizon Choker (13085, -1.05 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.22 DPS) [dungeon] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | sim-verified (275.1 DPS) | yes | Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -3.16 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.5 spell_power points (4.34 DPS) | yes | Runecloth Cloak (13860, -1.47 DPS) [crafted]; Sergeant Major's Cape (16337, -2.01 DPS) [pvp]; Big Voodoo Cloak (8216, -2.06 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.7 spell_power points (10.30 DPS) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Robe of Winter Night (14136, -1.10 DPS) [crafted]; Knight-Captain's Silk Tunic (227108, -2.05 DPS) [pvp] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (275.1 DPS) | yes | Marshal's Silk Bracers (16438, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Netherfroth Cuffs (254063, -2.82 DPS, sim-verified) [crafted] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-verified (275.1 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Frothing Gloves (254097, -3.58 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (275.1 DPS) | yes | Marshal's Silk Sash (16439, +0.00 DPS) [pvp]; Dreadmist Belt (16702, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -5.10 DPS, sim-verified) [crafted] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 45.8 spell_power points (8.48 DPS) | yes | Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, -0.24 DPS) [pvp]; Ghostweave Pants (14144, -7.82 DPS, sim-verified) [crafted] |
| feet | Rime-encrusted Boots (279268) | Tailoring [crafted] | 33.8 spell_power points (6.25 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Wolfrunner Shoes (13101, -0.63 DPS) [world_drop]; Omnicast Boots (11822, -0.75 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 28.3 spell_power points (5.23 DPS) | yes | Band of the Hierophant (13096, -2.17 DPS) [world_drop]; Philanthropist's Ring (281635, -2.63 DPS) [quest]; Eye of Adaegus (5266, -2.63 DPS) [world_drop] |
| finger2 | Frigid Ring (18679) | Princess Tempestria [world] | 22.1 spell_power points (4.09 DPS) | yes | Band of the Hierophant (13096, -1.03 DPS) [world_drop]; Philanthropist's Ring (281635, -1.49 DPS) [quest]; Eye of Adaegus (5266, -1.49 DPS) [world_drop] |
| trinket1 | Shard of the Splithooves (10659) | Heroes of Old [quest] | sim-verified (275.1 DPS) | yes | Ankh of Life (1713, -1.09 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -1.37 DPS) [quest]; Thunderbrew's Boot Flask (744, -1.65 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (275.1 DPS) | yes | Ankh of Life (1713, -0.26 DPS) [world_drop]; Evonice's Landin' Pilla (18951, -0.54 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.83 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (275.1 DPS) | yes | Grand Marshal's Stave (234571, -1.77 DPS) [pvp]; Spire of Hakkar (10844, -8.66 DPS) [world]; Kindling Stave (11750, -14.90 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 429.6 spell_power points (79.43 DPS) | yes | Stormrager (16997, -4.22 DPS, sim-verified) [quest]; Brilliant Wand (249385, -17.13 DPS) [crafted]; Torch of Austen (13004, -22.65 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Magister's Mantle; chest: Robe of the Archmage; wrist: Magister's Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Magister's Leggings; feet: Rime-encrusted Boots; finger1: Maiden's Circle; finger2: Frigid Ring; trinket1: Shard of the Splithooves; trinket2: Frozen Heart of the Mountain; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 230005200000000000-51020000000000000-253510130101301221)

Set DPS (verified): 543.1. Weights run: 1.4s. Verify run: 1.9s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.168 ± 0.010, crit=0.423 ± 0.013 per rating point (14 rating = 1%, 5.927 per %), hit=0.969 ± 0.052 per rating point (10 rating = 1%, 9.691 per %), spell_haste=6.324 ± 0.293, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.818 ± 0.002, spirit=not significant (0.004 ± 0.049), mp5=not significant (0.035 ± 0.072)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.78 DPS) | yes | Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -3.75 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (6.47 DPS) | yes | Triune Amulet (7722, -2.68 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.69 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.15 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.1 spell_power points (11.88 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.04 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, -1.10 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (7.11 DPS) | yes | Cloak of Flames (3475, -2.53 DPS) [world_drop]; Nightfall Drape (12465, -2.86 DPS) [dungeon]; Runecloth Cloak (13860, -3.06 DPS, sim-verified) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 47.9 spell_power points (22.69 DPS) | yes | Field Marshal's Silk Vestments (231603, -2.90 DPS) [pvp]; Robe of Winter Night (14136, -7.55 DPS, sim-verified) [crafted]; Knight-Captain's Silk Tunic (227108, -8.58 DPS) [pvp] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (6.00 DPS) | yes | Netherfroth Cuffs (254063, -0.40 DPS) [crafted]; Spidertank Oilrag (9448, -1.69 DPS) [dungeon]; Arcane Runed Bracers (4744, -1.74 DPS) [quest] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.5 spell_power points (10.66 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, +0.00 DPS) [quest]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor] |
| waist | Frothing Waistcord (254083) | Tailoring [crafted] | 17.1 spell_power points (8.07 DPS) | yes | Satyrmane Sash (17755, -0.65 DPS) [dungeon]; Ghostweave Cord (254073, -1.42 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.73 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.1 spell_power points (12.34 DPS) | yes | Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp]; Frostweave Pants (13871, -2.85 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.37 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.64 DPS) [pvp]; Venomspew Footpads (275606, -0.94 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (9.09 DPS) | yes | Band of the Unicorn (7553, -2.93 DPS) [world_drop]; Philanthropist's Ring (281635, -3.96 DPS) [quest]; Blackstone Ring (17713, -4.50 DPS) [dungeon] |
| finger2 | Frigid Ring (18679) | Princess Tempestria [world] | 13.1 spell_power points (6.21 DPS) | yes | Band of the Unicorn (7553, -0.06 DPS) [world_drop]; Philanthropist's Ring (281635, -1.08 DPS) [quest]; Blackstone Ring (17713, -1.63 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (543.1 DPS) | yes | Uther's Strength (11302, -3.59 DPS, sim-verified) [world_drop]; Shard of the Splithooves (10659, -4.03 DPS) [quest]; Ankh of Life (1713, -4.10 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (543.1 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (543.1 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -7.11 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -9.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 168.3 spell_power points (79.62 DPS) | yes | Brilliant Wand (249385, -15.23 DPS) [crafted]; Stormrager (16997, -16.53 DPS) [quest]; Greater Eternal Wand (249237, -22.03 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Frothing Waistcord; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Frigid Ring; trinket2: Mark of the Chosen; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-253100000000000000)

Set DPS (verified): 33.4. Weights run: 1.1s. Verify run: 0.9s. 179 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.385 ± 0.007, crit=0.072 ± 0.002 per rating point (14 rating = 1%, 1.003 per %), hit=0.225 ± 0.007 per rating point (10 rating = 1%, 2.254 per %), spell_haste=not significant (0.144 ± 0.108), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.222 ± 0.079), mp5=not significant (0.235 ± 0.079)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Silky Circlet (253951) | Tailoring [crafted] | 11.0 spell_power points (1.19 DPS) | yes | Pristine Circlet (253949, -0.70 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.84 DPS) [crafted]; Flying Tiger Goggles (4368, -1.09 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.7 spell_power points (0.07 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 spell_power points (0.91 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.41 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted]; Slime-encrusted Pads (6461, -0.84 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.9 spell_power points (0.53 DPS) | yes | Pearl-clasped Cloak (5542, -0.19 DPS) [crafted]; Black Whelp Cloak (7283, -0.20 DPS) [crafted]; Mystic Shawl (3449, -0.31 DPS) [quest] |
| chest | Filigreed Silky Gown (253903) | Tailoring [crafted] | 11.6 spell_power points (1.25 DPS) | yes | Filigreed Pristine Gown (253901, -0.56 DPS, sim-verified) [crafted]; Gray Woolen Robe (2585, -0.61 DPS) [crafted]; Scholarly Robes (2034, -0.64 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Owlbeard Bracers (16981, -0.09 DPS) [quest]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.42 DPS, sim-verified) [quest] |
| hands | Silky Gloves (253915) | Tailoring [crafted] | 8.2 spell_power points (0.88 DPS) | yes | Serpent Gloves (5970, -0.12 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.23 DPS) [world]; Pristine Gloves (253913, -0.32 DPS) [crafted] |
| waist | Silky Sash (253927) | Tailoring [crafted] | 8.5 spell_power points (0.92 DPS) | yes | Novice Arcanist's Sash (253885, -0.37 DPS) [crafted]; Pristine Sash (253925, -0.42 DPS, sim-verified) [crafted]; Novice Ardent's Sash (253887, -0.51 DPS) [crafted] |
| legs | Filigreed Silky Leggings (253939) | Tailoring [crafted] | 14.2 spell_power points (1.53 DPS) | yes | Filigreed Pristine Leggings (253937, -0.70 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.78 DPS) [dungeon]; Colorful Kilt (10048, -0.87 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.92 DPS) | yes | Silky Boots (253891, -0.35 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.37 DPS) [crafted]; Pristine Boots (253889, -0.47 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.3 spell_power points (0.25 DPS) | yes | Band of Purification (12996, -0.11 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; Furlfeather Ring (286745, -0.14 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 1.9 spell_power points (0.21 DPS) | yes | Volcanic Rock Ring (12053, -0.08 DPS) [world_drop]; Furlfeather Ring (286745, -0.10 DPS) [world]; Band of Purification (12996, -0.39 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-verified (+3.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Staff of the Spire (1300, -0.04 DPS) [world]; Staff of Orgrimmar (15444, -0.04 DPS) [quest]; Twisted Chanter's Staff (890, -3.48 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 206.6 spell_power points (22.31 DPS) | yes | Skycaller (12984, -0.60 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -4.44 DPS) [quest]; Greater Magic Wand (11288, -4.59 DPS) [crafted] |

**New at 20:** head: Silky Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Silky Gown; wrist: Mindthrust Bracers; hands: Silky Gloves; waist: Silky Sash; legs: Filigreed Silky Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Crescent Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (troll, 000000000000000000-00000000000000000-253510130100000000)

Set DPS (verified): 58.8. Weights run: 1.0s. Verify run: 0.8s. 313 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.397 ± 0.009, crit=0.112 ± 0.004 per rating point (14 rating = 1%, 1.571 per %), hit=0.233 ± 0.009 per rating point (10 rating = 1%, 2.334 per %), spell_haste=not significant (0.485 ± 0.151), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.002, spirit=not significant (0.083 ± 0.062), mp5=not significant (0.212 ± 0.091)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Filigreed Silky Circlet (253977) | Tailoring [crafted] | 14.0 spell_power points (2.13 DPS) | yes | Holy Shroud (2721, -0.38 DPS) [world_drop]; Silky Circlet (253951, -0.46 DPS) [crafted]; Azure Silk Hood (7048, -0.61 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 2.4 spell_power points (0.36 DPS) | yes | Kaleidoscope Chain (13084, -0.07 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.08 DPS) [world_drop]; Pendant of Myzrael (4614, -0.29 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (1.91 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Invoker's Mantle (215365, -0.55 DPS) [crafted]; Magician's Mantle (12998, -0.61 DPS) [world_drop] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | 10.2 spell_power points (1.56 DPS) | yes | Hillman's Cloak (3719, -0.80 DPS) [crafted]; Windsong Drape (15468, -0.87 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.90 DPS) [crafted] |
| chest | Silky Gown (253963) | Tailoring [crafted] | 16.1 spell_power points (2.45 DPS) | yes | Beguiler Robes (7728, -0.26 DPS) [dungeon]; Green Silk Armor (7065, -0.30 DPS) [crafted]; Tree Bark Jacket (1486, -0.47 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.6 spell_power points (1.47 DPS) | yes | Tabitha's Cuffs (251486, -1.01 DPS) [quest]; Nightsky Wristbands (6407, -1.07 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.58 DPS, sim-verified) [world_drop] |
| hands | Azure Silk Gloves (4319) | Tailoring [crafted] | 10.2 spell_power points (1.56 DPS) | yes | Silky Gloves (253915, -0.31 DPS) [crafted]; Jutebraid Gloves (10654, -0.34 DPS) [quest]; Serpent Gloves (5970, -0.49 DPS) [dungeon] |
| waist | Warsong Sash (16975) | Warsong Supplies [quest] | 11.0 spell_power points (1.67 DPS) | yes | Belt of Arugal (6392, +0.00 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.31 DPS) [crafted]; Crimson Silk Belt (7055, -0.34 DPS) [crafted] |
| legs | Silky Leggings (253989) | Tailoring [crafted] | 17.2 spell_power points (2.62 DPS) | yes | Filigreed Silky Leggings (253939, -0.53 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -0.68 DPS) [dungeon]; Azure Silk Pants (7046, -0.73 DPS) [crafted] |
| feet | Frothing Slippers (254003) | Tailoring [crafted] | 16.1 spell_power points (2.45 DPS) | yes | Gilded Slippers (254001, -1.09 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.14 DPS) [crafted]; Acidic Walkers (9454, -1.16 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 spell_power points (1.21 DPS) | yes | Solid Gold Band (279845, -0.54 DPS) [quest]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Runemender's Seal (273031, -0.67 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.6 spell_power points (1.01 DPS) | yes | Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Runemender's Seal (273031, -0.46 DPS) [world]; Solid Gold Band (279845, -0.98 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.4 spell_power points (0.66 DPS) | yes | Crescent Staff (6505, -0.24 DPS) [quest]; Polished Walking Staff (16889, -0.24 DPS) [quest]; Twisted Chanter's Staff (890, -6.14 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 217.6 spell_power points (33.10 DPS) | yes | Wand of Mana Concentration (273043, -1.62 DPS) [world]; Gyromatic Icemaker (9489, -2.74 DPS) [dungeon]; Greater Mystic Wand (217287, -3.34 DPS) [crafted] |

**New at 30:** head: Filigreed Silky Circlet; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Azure Silk Cloak; chest: Silky Gown; wrist: Spidertank Oilrag; hands: Azure Silk Gloves; waist: Warsong Sash; legs: Silky Leggings; feet: Frothing Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Glimmering Staff; ranged: Starfaller

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (troll, 000000000000000000-00000000000000000-253510130100301231)

Set DPS (verified): 128.7. Weights run: 1.1s. Verify run: 0.9s. 411 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.693 ± 0.021, crit=0.169 ± 0.007 per rating point (14 rating = 1%, 2.363 per %), hit=0.388 ± 0.024 per rating point (10 rating = 1%, 3.878 per %), spell_haste=not significant (0.959 ± 0.332), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.672 ± 0.256), mp5=not significant (1.093 ± 0.307)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 25.8 spell_power points (4.71 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.88 DPS) [crafted]; Electromagnetic Gigaflux Reactivator (9492, -1.14 DPS) [dungeon]; Whitemane's Chapeau (7720, -2.29 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 12.7 spell_power points (2.32 DPS) | yes | Glowing Eye of Mordresh (10769, -0.03 DPS) [dungeon]; Necklace of Calisea (1714, -0.57 DPS) [dungeon]; Amberglow Talisman (10824, -1.09 DPS) [quest] |
| shoulder | Mantle of Doan (7712) | Scarlet Monastery: Arcanist Doan [dungeon] | 16.5 spell_power points (3.01 DPS) | yes | Inquisitor's Shawl (19507, -0.10 DPS) [dungeon]; Green Silken Shoulders (7057, -0.17 DPS) [crafted]; Bloodmage Mantle (7684, -0.24 DPS) [dungeon] |
| back | Azure Silk Cloak (7053) | Tailoring [crafted] | 12.0 spell_power points (2.19 DPS) | yes | Frost Leather Cloak (7377, -0.19 DPS) [crafted]; Long Silken Cloak (4326, -0.46 DPS) [crafted]; Icy Cloak (4327, -1.01 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 33.6 spell_power points (6.13 DPS) | yes | Robe of Power (7054, -1.08 DPS) [crafted]; Robe of the Magi (1716, -1.73 DPS, sim-verified) [world_drop]; Beguiler Robes (7728, -2.00 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 12.3 spell_power points (2.24 DPS) | yes | Radiant Silver Bracers (4545, -0.50 DPS) [quest]; Earthen Silk Cuffs (254019, -0.90 DPS) [crafted]; Condor Bracers (15864, -0.96 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 25.5 spell_power points (4.64 DPS) | yes | Red Mageweave Gloves (10018, -1.38 DPS) [crafted]; Black Mageweave Gloves (10003, -1.91 DPS) [crafted]; Frothing Handwraps (254023, -2.05 DPS, sim-verified) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 23.6 spell_power points (4.30 DPS) | yes | Deathmage Sash (10771, -1.13 DPS) [dungeon]; Star Belt (4329, -1.44 DPS) [crafted]; Gilded Cord (254037, -1.63 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.3 spell_power points (4.07 DPS) | yes | Silky Leggings (253989, -0.02 DPS) [crafted]; Stoneweaver Leggings (9407, -0.31 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.78 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 28.7 spell_power points (5.23 DPS) | yes | Frothing Slippers (254003, -1.93 DPS, sim-verified) [crafted]; Furen's Boots (13100, -2.39 DPS) [world_drop]; Gilded Slippers (254001, -2.58 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.5 spell_power points (2.45 DPS) | yes | Ring of Power Regulation (273030, -0.76 DPS) [dungeon]; Welken Ring (5011, -0.97 DPS) [world_drop]; Electrocutioner Lagnut (9447, -1.05 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.5 spell_power points (1.73 DPS) | yes | Ring of Power Regulation (273030, -0.03 DPS) [dungeon]; Welken Ring (5011, -0.24 DPS) [world_drop]; Electrocutioner Lagnut (9447, -0.32 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (128.7 DPS) | yes | Staff of the Friar (3415, -2.96 DPS) [dungeon]; Gnarled Ash Staff (791, -2.98 DPS) [world_drop]; Illusionary Rod (7713, -9.98 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 222.2 spell_power points (40.50 DPS) | yes | Icefury Wand (7514, -1.29 DPS) [quest]; Nether Force Wand (11263, -2.79 DPS) [quest]; Ragefire Wand (7513, -2.98 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Mantle of Doan; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Frothing Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (troll, 230005000000000000-00000000000000000-253510130100301231)

Set DPS (verified): 179.3. Weights run: 1.1s. Verify run: 1.1s. 527 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.030 ± 0.038, crit=0.279 ± 0.012 per rating point (14 rating = 1%, 3.902 per %), hit=0.696 ± 0.043 per rating point (10 rating = 1%, 6.964 per %), spell_haste=3.309 ± 0.673, spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (1.094 ± 0.483), mp5=not significant (1.857 ± 0.604)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 44.4 spell_power points (7.82 DPS) | yes | Red Mageweave Headband (10033, -0.85 DPS) [crafted]; Papal Fez (9431, -1.60 DPS) [dungeon]; Gemburst Circlet (10751, -2.74 DPS, sim-verified) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 18.8 spell_power points (3.31 DPS) | yes | Triune Amulet (7722, -0.55 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.63 DPS) [dungeon]; Mindburst Medallion (11196, -2.18 DPS, sim-verified) [quest] |
| shoulder | Netherfroth Shoulders (254051) | Tailoring [crafted] | 31.8 spell_power points (5.60 DPS) | yes | Rotgrip Mantle (17732, -0.05 DPS) [dungeon]; Nethergeld Shoulders (254049, -1.41 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.65 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 26.7 spell_power points (4.71 DPS) | yes | Deep Woodlands Cloak (19121, -0.96 DPS) [quest]; Runecloth Cloak (13860, -1.67 DPS) [crafted]; Imperial Red Cloak (8248, -1.94 DPS) [world_drop] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 45.7 spell_power points (8.05 DPS) | yes | Dreamweave Vest (10021, -0.55 DPS) [crafted]; Acumen Robes (17775, -1.08 DPS) [quest]; Frostweave Robe (13868, -1.71 DPS) [crafted] |
| wrist | Netherfroth Cuffs (254063) | Tailoring [crafted] | 25.7 spell_power points (4.52 DPS) | yes | Bloodband Bracers (11469, -1.43 DPS) [quest]; Nethergeld Cuffs (254061, -1.58 DPS, sim-verified) [crafted]; Shizzle's Nozzle Wiper (11917, -1.77 DPS) [quest] |
| hands | Frothing Gloves (254097) | Tailoring [crafted] | 38.0 spell_power points (6.70 DPS) | yes | Dreamweave Gloves (10019, -1.46 DPS) [crafted]; Gilded Gloves (254095, -1.58 DPS) [crafted]; Runecloth Gloves (13863, -3.72 DPS, sim-verified) [crafted] |
| waist | Frothing Cord (254039) | Tailoring [crafted] | 28.8 spell_power points (5.07 DPS) | yes | Frothing Waistcord (254083, -0.09 DPS) [crafted]; Ghostweave Cord (254073, -0.49 DPS) [crafted]; Dawnspire Cord (12466, -0.57 DPS) [dungeon] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 41.2 spell_power points (7.26 DPS) | yes | Senior Designer's Pantaloons (11841, -0.41 DPS) [dungeon]; Dalewind Trousers (13008, -0.81 DPS) [world_drop]; Venomshroud Leggings (14444, -2.23 DPS) [world_drop] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 31.7 spell_power points (5.57 DPS) | yes | Frothing Sandals (254109, -0.42 DPS) [crafted]; Vinerot Sandals (17748, -0.96 DPS) [dungeon]; Furen's Boots (13100, -1.22 DPS) [world_drop] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 19.3 spell_power points (3.40 DPS) | yes | Band of the Unicorn (7553, -1.11 DPS) [world_drop]; Coldwater Ring (4550, -1.12 DPS) [quest]; Welken Ring (5011, -1.14 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (2.67 DPS) | yes | Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Coldwater Ring (4550, -0.39 DPS) [quest]; Welken Ring (5011, -0.41 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (179.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.25 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (179.3 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, -0.20 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (179.3 DPS) | yes | Spellshifter Rod (9527, -0.32 DPS) [quest]; Kindling Stave (11750, -0.57 DPS) [dungeon]; Zum'rah's Vexing Cane (18082, -2.68 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 301.4 spell_power points (53.08 DPS) | yes | Noxious Shooter (17745, -2.20 DPS) [dungeon]; Wand of Allistarj (13065, -4.66 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.92 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Horizon Choker; shoulder: Netherfroth Shoulders; back: Spritecaster Cape; chest: Ghostweave Vest; wrist: Netherfroth Cuffs; hands: Frothing Gloves; legs: Kilt of the Atal'ai Prophet; finger1: Eye of Adaegus; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 527, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (troll, 230005000000000000-23023000000000000-253510130100301231)

Set DPS (verified): 259.2. Weights run: 1.2s. Verify run: 2.7s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.812 ± 0.052, crit=0.425 ± 0.016 per rating point (14 rating = 1%, 5.943 per %), hit=1.177 ± 0.063 per rating point (10 rating = 1%, 11.768 per %), spell_haste=not significant (2.130 ± 0.830), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003, spirit=not significant (0.765 ± 0.563), mp5=not significant (2.511 ± 0.707)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (259.2 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Crown of the Penitent (13216, -7.32 DPS, sim-verified) [quest] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 20.1 spell_power points (3.72 DPS) | yes | Horizon Choker (13085, -1.05 DPS) [world_drop]; Glowing Eye of Mordresh (10769, -1.22 DPS) [dungeon]; Mindburst Medallion (11196, -2.84 DPS, sim-verified) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 33.5 spell_power points (6.19 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -0.10 DPS) [pvp]; Shroud of the Nathrezim (18720, -0.27 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 23.5 spell_power points (4.34 DPS) | yes | Deep Woodlands Cloak (19121, -0.77 DPS) [quest]; Runecloth Cloak (13860, -1.47 DPS) [crafted]; Big Voodoo Cloak (8216, -2.06 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 55.7 spell_power points (10.30 DPS) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.05 DPS) [pvp]; Robe of Winter Night (14136, -2.16 DPS, sim-verified) [crafted] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (259.2 DPS) | yes | General's Silk Cuffs (16538, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Netherfroth Cuffs (254063, -8.23 DPS, sim-verified) [crafted] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | sim-verified (259.2 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Frothing Gloves (254097, -9.51 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (259.2 DPS) | yes | General's Silk Sash (16537, +0.00 DPS) [pvp]; Dreadmist Belt (16702, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -8.91 DPS, sim-verified) [crafted] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 45.8 spell_power points (8.48 DPS) | yes | General's Silk Trousers (231595, +0.00 DPS) [pvp]; Legionnaire's Silk Legguards (227107, -0.24 DPS) [pvp]; Ghostweave Pants (14144, -10.88 DPS, sim-verified) [crafted] |
| feet | Rime-encrusted Boots (279268) | Tailoring [crafted] | 33.8 spell_power points (6.25 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Wolfrunner Shoes (13101, -0.63 DPS) [world_drop]; Omnicast Boots (11822, -0.75 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 28.3 spell_power points (5.23 DPS) | yes | Band of the Hierophant (13096, -2.17 DPS) [world_drop]; Philanthropist's Ring (281635, -2.63 DPS) [quest]; Eye of Adaegus (5266, -2.63 DPS) [world_drop] |
| finger2 | Frigid Ring (18679) | Princess Tempestria [world] | 22.1 spell_power points (4.09 DPS) | yes | Philanthropist's Ring (281635, -1.49 DPS) [quest]; Eye of Adaegus (5266, -1.49 DPS) [world_drop]; Band of the Hierophant (13096, -2.72 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (259.2 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Ankh of Life (1713, -0.26 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.44 DPS) [quest] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (259.2 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (259.2 DPS) | yes | High Warlord's War Staff (234549, -1.77 DPS) [pvp]; Spire of Hakkar (10844, -8.66 DPS) [world]; Kindling Stave (11750, -14.35 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 429.6 spell_power points (79.43 DPS) | yes | Stormrager (16997, -5.65 DPS, sim-verified) [quest]; Brilliant Wand (249385, -17.13 DPS) [crafted]; Torch of Austen (13004, -22.65 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Magister's Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Magister's Leggings; feet: Rime-encrusted Boots; finger1: Maiden's Circle; finger2: Frigid Ring; trinket2: Evonice's Landin' Pilla; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (troll, 230005200000000000-51020000000000000-253510130101301221)

Set DPS (verified): 545.2. Weights run: 1.4s. Verify run: 2.0s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.168 ± 0.010, crit=0.423 ± 0.013 per rating point (14 rating = 1%, 5.927 per %), hit=0.969 ± 0.052 per rating point (10 rating = 1%, 9.691 per %), spell_haste=6.324 ± 0.293, spell_penetration=not significant (0.000 ± 0.000), frost_power=0.818 ± 0.002, spirit=not significant (0.004 ± 0.049), mp5=not significant (0.035 ± 0.072)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 spell_power points (12.78 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Dreamweave Circlet (10041, -3.42 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.7 spell_power points (6.47 DPS) | yes | Glowing Eye of Mordresh (10769, -2.53 DPS) [dungeon]; Triune Amulet (7722, -2.68 DPS) [dungeon]; Mindburst Medallion (11196, -3.15 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 25.1 spell_power points (11.88 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Argent Shoulders (19059, -0.04 DPS) [crafted]; Champion's Silk Mantle (227104, -1.10 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.0 spell_power points (7.11 DPS) | yes | Deep Woodlands Cloak (19121, -0.72 DPS) [quest]; Runecloth Cloak (13860, -2.22 DPS) [crafted]; Cloak of Flames (3475, -2.53 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 47.9 spell_power points (22.69 DPS) | yes | Warlord's Silk Raiment (231596, -2.90 DPS) [pvp]; Robe of Winter Night (14136, -7.54 DPS, sim-verified) [crafted]; Legionnaire's Silk Tunic (227106, -8.58 DPS) [pvp] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.7 spell_power points (6.00 DPS) | yes | Netherfroth Cuffs (254063, -0.40 DPS) [crafted]; Spidertank Oilrag (9448, -1.69 DPS) [dungeon]; Nethergeld Cuffs (254061, -2.12 DPS) [crafted] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 22.5 spell_power points (10.66 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, +0.00 DPS) [quest]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor] |
| waist | Frothing Waistcord (254083) | Tailoring [crafted] | 17.1 spell_power points (8.07 DPS) | yes | Satyrmane Sash (17755, -0.65 DPS) [dungeon]; Ghostweave Cord (254073, -1.42 DPS) [crafted]; Wisdom of the Timbermaw (19047, -1.73 DPS) [crafted] |
| legs | Ghostweave Pants (14144) | Tailoring [crafted] | 26.1 spell_power points (12.34 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Frostweave Pants (13871, -3.03 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.37 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.64 DPS) [pvp]; Venomspew Footpads (275606, -0.94 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (9.09 DPS) | yes | Band of the Unicorn (7553, -2.93 DPS) [world_drop]; Philanthropist's Ring (281635, -3.96 DPS) [quest]; Blackstone Ring (17713, -4.50 DPS) [dungeon] |
| finger2 | Frigid Ring (18679) | Princess Tempestria [world] | 13.1 spell_power points (6.21 DPS) | yes | Band of the Unicorn (7553, -0.06 DPS) [world_drop]; Philanthropist's Ring (281635, -1.08 DPS) [quest]; Blackstone Ring (17713, -1.63 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (545.2 DPS) | yes | Uther's Strength (11302, -1.29 DPS) [world_drop]; Shard of the Splithooves (10659, -4.03 DPS) [quest]; Ankh of Life (1713, -4.10 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (545.2 DPS) | yes | Uther's Strength (11302, -0.37 DPS) [world_drop]; Shard of the Splithooves (10659, -3.11 DPS) [quest]; Ankh of Life (1713, -3.19 DPS) [world_drop] |
| main_hand | Elemental Mage Staff (944) | World drop [world_drop] | sim-verified (545.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -7.76 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -9.57 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 168.3 spell_power points (79.62 DPS) | yes | Brilliant Wand (249385, -15.23 DPS) [crafted]; Stormrager (16997, -16.53 DPS) [quest]; Greater Eternal Wand (249237, -22.03 DPS) [crafted] |

**New at 60:** head: Spellpower Goggles Xtreme Plus; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Frothing Waistcord; legs: Ghostweave Pants; finger1: Maiden's Circle; finger2: Frigid Ring; trinket2: Rune of the Guard Captain; main_hand: Elemental Mage Staff; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

