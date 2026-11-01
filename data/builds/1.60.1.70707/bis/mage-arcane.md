# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 141005000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 34.2. Weights run: 1.1s. Verify run: 1.0s. 185 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.404 ± 0.008, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.902 per %), hit=0.254 ± 0.011 per rating point (10 rating = 1%, 2.539 per %), spell_haste=not significant (0.159 ± 0.090), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.641 ± 0.006, spirit=not significant (0.136 ± 0.089), mp5=not significant (0.194 ± 0.105)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 7.1 spell_power points (0.85 DPS) | yes | Pristine Circlet (253949, -0.40 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.51 DPS) [crafted]; Flying Tiger Goggles (4368, -0.78 DPS) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.3 spell_power points (0.03 DPS) | yes | Tarnished Locket (279870, +0.00 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 spell_power points (1.03 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.56 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.65 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.96 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.5 spell_power points (0.54 DPS) | yes | Pearl-clasped Cloak (5542, -0.16 DPS) [crafted]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Sanguine Cape (14376, -0.35 DPS) [world_drop] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 8.2 spell_power points (0.98 DPS) | yes | Gray Woolen Robe (2585, -0.26 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS, sim-verified) [crafted]; Scholarly Robes (2034, -0.43 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | 2.6 spell_power points (0.31 DPS) | yes | Repurposed Hair Band (281256, -0.18 DPS) [quest]; Windsong Bangles (263336, -0.19 DPS) [quest]; Bright Bracers (3647, -0.74 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.84 DPS) | yes | Gnoll Casting Gloves (892, -0.12 DPS) [world]; Pearly Gloves (253921, -0.16 DPS) [crafted]; Pristine Gloves (253913, -0.21 DPS) [crafted] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 6.1 spell_power points (0.73 DPS) | yes | Pristine Sash (253925, -0.06 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Novice Ardent's Sash (253887, -0.30 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 10.0 spell_power points (1.20 DPS) | yes | Silk-threaded Trousers (1929, -0.36 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.40 DPS, sim-verified) [crafted]; Colorful Kilt (10048, -0.52 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 spell_power points (1.03 DPS) | yes | Red Woolen Boots (4313, -0.47 DPS) [crafted]; Pristine Boots (253889, -0.53 DPS) [crafted]; Pearly Boots (253897, -0.59 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.70 DPS) | yes | Volcanic Rock Ring (12053, -0.55 DPS) [world_drop]; Furlfeather Ring (286745, -0.58 DPS) [world]; Band of Purification (12996, -0.60 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.4 spell_power points (0.29 DPS) | yes | Furlfeather Ring (286745, -0.17 DPS) [world]; Band of Purification (12996, -0.19 DPS) [world_drop]; Volcanic Rock Ring (12053, -0.45 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | sim-verified (+2.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Staff of the Spire (1300, -0.05 DPS) [world]; Witching Stave (1484, -0.10 DPS) [dungeon]; Twisted Chanter's Staff (890, -2.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Cookie's Stirring Rod (5198, -0.53 DPS, sim-verified) [dungeon]; Deepblaze (279896, -3.49 DPS) [quest]; Sizzle Stick (8071, -4.14 DPS) [quest] |

**New at 20:** head: Pearly Circlet; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Staff of Westfall; ranged: Skycaller

No-known-source sample (15 of 185, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20428 Caretaker's Cape; 20431 Lorekeeper's Ring

### Band 30 (gnome, 141005213100300000-00000000000000000-000000000000000000)

Set DPS (verified): 75.0. Weights run: 1.3s. Verify run: 1.0s. 323 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.538 ± 0.013, crit=0.076 ± 0.002 per rating point (14 rating = 1%, 1.063 per %), hit=0.295 ± 0.014 per rating point (10 rating = 1%, 2.946 per %), spell_haste=not significant (0.274 ± 0.227), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.551 ± 0.003, spirit=not significant (0.388 ± 0.168), mp5=not significant (0.710 ± 0.198)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 13.3 spell_power points (2.28 DPS) | yes | Enchanter's Cowl (4322, -0.33 DPS) [crafted]; Silk Headband (7050, -0.74 DPS) [crafted]; Nightsky Cowl (4039, -0.84 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 3.7 spell_power points (0.64 DPS) | yes | Crystal Starfire Medallion (5003, -0.07 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Pendant of Myzrael (4614, -0.24 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.8 spell_power points (2.37 DPS) | yes | Death Speaker Mantle (6685, -0.33 DPS) [dungeon]; Magician's Mantle (12998, -0.69 DPS) [world_drop]; Invoker's Mantle (215365, -0.71 DPS) [crafted] |
| back | Glowing Thresher Cape (6901) | Blackfathom Deeps: Old Serra'kis [dungeon] | 6.1 spell_power points (1.05 DPS) | yes | Repairman's Cape (9605, -0.09 DPS) [quest]; Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.19 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 18.6 spell_power points (3.18 DPS) | yes | Green Silk Armor (7065, -0.44 DPS) [crafted]; Mechbuilder's Overalls (9508, -0.70 DPS) [dungeon]; Tree Bark Jacket (1486, -0.95 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.1 spell_power points (1.91 DPS) | yes | Mindthrust Bracers (1974, -1.03 DPS, sim-verified) [dungeon]; Nightsky Wristbands (6407, -1.15 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.17 DPS) [world_drop] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.7 spell_power points (1.50 DPS) | yes | Gloves of Meditation (4318, -0.17 DPS) [crafted]; Gloves of Old (9395, -0.24 DPS) [world_drop]; Shilly Mitts (9609, -0.24 DPS) [quest] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 12.6 spell_power points (2.15 DPS) | yes | Invoker's Cord (215366, -0.49 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.78 DPS) [dungeon]; Crimson Silk Belt (7055, -1.01 DPS, sim-verified) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 15.5 spell_power points (2.66 DPS) | yes | Pearly Leggings (253995, -0.35 DPS) [crafted]; Pristine Leggings (253987, -0.48 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.80 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 12.5 spell_power points (2.14 DPS) | yes | Gilded Slippers (254001, -0.03 DPS) [crafted]; Acidic Walkers (9454, -0.28 DPS) [dungeon]; Glinteye Slippers (273024, -0.54 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.7 spell_power points (1.49 DPS) | yes | Minor Channeling Ring (1449, -0.45 DPS) [quest]; Electrocutioner Lagnut (9447, -0.51 DPS) [dungeon]; Runemender's Seal (273031, -0.51 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.1 spell_power points (1.39 DPS) | yes | Electrocutioner Lagnut (9447, -0.41 DPS) [dungeon]; Runemender's Seal (273031, -0.41 DPS) [world]; Minor Channeling Ring (1449, -0.69 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | sim-verified (75.0 DPS) | yes | Gnarled Ash Staff (791, -0.02 DPS) [world_drop]; Twisted Chanter's Staff (890, -0.09 DPS) [world_drop]; Staff of the Friar (3415, -3.86 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 197.2 spell_power points (33.79 DPS) | yes | Wand of Mana Concentration (273043, -2.00 DPS) [world]; Gyromatic Icemaker (9489, -3.44 DPS) [dungeon]; Greater Mystic Wand (217287, -3.94 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Glowing Thresher Cape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Glimmering Staff; ranged: Starfaller

No-known-source sample (15 of 323, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 156.9. Weights run: 1.3s. Verify run: 1.1s. 421 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.739 ± 0.026, crit=0.158 ± 0.004 per rating point (14 rating = 1%, 2.213 per %), hit=0.385 ± 0.024 per rating point (10 rating = 1%, 3.847 per %), spell_haste=not significant (0.100 ± 0.325), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.590 ± 0.003, spirit=not significant (0.910 ± 0.357), mp5=not significant (0.420 ± 0.395)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 28.7 spell_power points (5.97 DPS) | yes | Whitemane's Chapeau (7720, -0.72 DPS) [dungeon]; Electromagnetic Gigaflux Reactivator (9492, -1.38 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.61 DPS) [crafted] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 14.4 spell_power points (2.98 DPS) | yes | Glowing Eye of Mordresh (10769, -0.30 DPS) [dungeon]; Necklace of Calisea (1714, -0.59 DPS) [dungeon]; Amberglow Talisman (10824, -1.10 DPS) [quest] |
| shoulder | Mantle of Doan (7712) | Scarlet Monastery: Arcanist Doan [dungeon] | 18.7 spell_power points (3.89 DPS) | yes | Inquisitor's Shawl (19507, -0.44 DPS) [dungeon]; Green Silken Shoulders (7057, -0.54 DPS) [crafted]; Bloodmage Mantle (7684, -0.64 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 12.5 spell_power points (2.59 DPS) | yes | Glowing Thresher Cape (6901, -0.45 DPS) [dungeon]; Sergeant Major's Cape (16336, -0.53 DPS) [pvp]; Long Silken Cloak (4326, -0.57 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 37.4 spell_power points (7.77 DPS) | yes | Robe of the Magi (1716, -1.33 DPS) [world_drop]; Robe of Power (7054, -1.50 DPS) [crafted]; Beguiler Robes (7728, -2.54 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.3 spell_power points (2.13 DPS) | yes | Arcane Runed Bracers (4744, -0.26 DPS) [quest]; Earthen Silk Cuffs (254019, -0.36 DPS) [crafted]; Silkstream Cuffs (16791, -0.61 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 27.3 spell_power points (5.68 DPS) | yes | Red Mageweave Gloves (10018, -2.30 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.56 DPS) [crafted]; Gloves of Old (9395, -2.64 DPS) [world_drop] |
| waist | Golden Cord (254045) | Tailoring [crafted] | 19.6 spell_power points (4.08 DPS) | yes | Gilded Cord (254037, -0.05 DPS) [crafted]; Deathmage Sash (10771, -0.32 DPS) [dungeon]; Star Belt (4329, -0.62 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | sim-verified (156.9 DPS) | yes | Black Mageweave Leggings (9999, -0.44 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.56 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.57 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.4 spell_power points (6.31 DPS) | yes | Thoughtcast Boots (10578, -2.81 DPS) [dungeon]; Golden Slippers (254009, -2.89 DPS) [crafted]; Furen's Boots (13100, -3.44 DPS, sim-verified) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (2.85 DPS) | yes | Snake Hoop (6750, -0.83 DPS) [quest]; Electrocutioner Lagnut (9447, -0.90 DPS) [dungeon]; Runemender's Seal (273031, -0.90 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 10.1 spell_power points (2.09 DPS) | yes | Snake Hoop (6750, -0.08 DPS) [quest]; Electrocutioner Lagnut (9447, -0.14 DPS) [dungeon]; Runemender's Seal (273031, -0.14 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gnarled Ash Staff (791, -3.32 DPS) [world_drop]; Staff of the Friar (3415, -3.43 DPS) [dungeon]; Windweaver Staff (7757, -19.47 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 196.7 spell_power points (40.86 DPS) | yes | Nether Force Wand (11263, +0.00 DPS, sim-verified) [quest]; Icefury Wand (7514, -3.21 DPS) [quest]; Ragefire Wand (7513, -3.26 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Mantle of Doan; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Golden Cord; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Welken Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 141005213100301531-32023000000000000-000000000000000000)

Set DPS (verified): 201.6. Weights run: 1.3s. Verify run: 1.2s. 537 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.950 ± 0.040, crit=0.256 ± 0.008 per rating point (14 rating = 1%, 3.585 per %), hit=0.702 ± 0.041 per rating point (10 rating = 1%, 7.020 per %), spell_haste=not significant (-0.549 ± 0.496), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.621 ± 0.003, spirit=not significant (1.032 ± 0.437), mp5=not significant (1.782 ± 0.511)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 42.9 spell_power points (8.25 DPS) | yes | Red Mageweave Headband (10033, -0.94 DPS) [crafted]; Papal Fez (9431, -1.82 DPS) [dungeon]; Gemburst Circlet (10751, -2.33 DPS, sim-verified) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.4 spell_power points (3.35 DPS) | yes | Mindburst Medallion (11196, -0.11 DPS) [quest]; Triune Amulet (7722, -0.43 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.54 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 30.1 spell_power points (5.79 DPS) | yes | Nethergeld Shoulders (254049, -1.41 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.70 DPS) [crafted]; Nethershine Shoulders (254057, -2.51 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.9 spell_power points (4.98 DPS) | yes | Runecloth Cloak (13860, -1.79 DPS) [crafted]; Imperial Red Cloak (8248, -2.18 DPS) [world_drop]; Big Voodoo Cloak (8216, -2.37 DPS) [crafted] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 43.9 spell_power points (8.45 DPS) | yes | Dreamweave Vest (10021, -0.56 DPS) [crafted]; Acumen Robes (17775, -1.13 DPS) [quest]; Robe of Power (7054, -1.97 DPS) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 19.9 spell_power points (3.83 DPS) | yes | Nethergeld Cuffs (254061, -0.21 DPS) [crafted]; Bloodband Bracers (11469, -0.62 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.04 DPS) [quest] |
| hands | Runecloth Gloves (13863) | Tailoring [crafted] | 29.8 spell_power points (5.74 DPS) | yes | Golden Gloves (254103, -0.06 DPS) [crafted]; Dreamweave Gloves (10019, -0.16 DPS) [crafted]; Gilded Gloves (254095, -0.40 DPS) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | 25.3 spell_power points (4.88 DPS) | yes | Dawnspire Cord (12466, -0.25 DPS) [dungeon]; Satyrmane Sash (17755, -0.36 DPS) [dungeon]; Golden Cord (254045, -0.55 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 38.7 spell_power points (7.44 DPS) | yes | Senior Designer's Pantaloons (11841, -0.28 DPS) [dungeon]; Dalewind Trousers (13008, -0.83 DPS) [world_drop]; Knight's Dreadweave Leggings (220888, -2.25 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 31.2 spell_power points (6.01 DPS) | yes | Sergeant Major's Dreadweave Boots (220891, -1.47 DPS) [vendor]; Furen's Boots (13100, -1.55 DPS) [world_drop]; Vinerot Sandals (17748, -2.30 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 18.1 spell_power points (3.48 DPS) | yes | Choking Band (11868, -0.90 DPS) [quest]; Band of the Unicorn (7553, -0.98 DPS) [world_drop]; Coldwater Ring (4550, -1.14 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (2.84 DPS) | yes | Choking Band (11868, -0.26 DPS) [quest]; Band of the Unicorn (7553, -0.34 DPS) [world_drop]; Coldwater Ring (4550, -0.50 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (201.6 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Uther's Strength (11302, -0.06 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (201.6 DPS) | yes | Evonice's Landin' Pilla (18951, -0.40 DPS) [quest]; Thunderbrew's Boot Flask (744, -0.79 DPS) [quest]; Uther's Strength (11302, -1.23 DPS) [world_drop] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (201.6 DPS) | yes | Spellshifter Rod (9527, -0.51 DPS) [quest]; Kindling Stave (11750, -0.68 DPS) [dungeon]; Inventor's Focal Sword (17719, -2.87 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 275.9 spell_power points (53.10 DPS) | yes | Noxious Shooter (17745, -2.13 DPS) [dungeon]; Wand of Allistarj (13065, -4.67 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.81 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Ghostweave Vest; wrist: Nethershine Cuffs; hands: Runecloth Gloves; waist: Ghostweave Cord; legs: Kilt of the Atal'ai Prophet; finger1: Eye of Adaegus; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Ankh of Life; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 537, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 141005213100301531-32023000000000000-053200000000000000)

Set DPS (verified): 324.6. Weights run: 1.4s. Verify run: 3.1s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.903 ± 0.043, crit=0.376 ± 0.010 per rating point (14 rating = 1%, 5.260 per %), hit=0.899 ± 0.045 per rating point (10 rating = 1%, 8.989 per %), spell_haste=not significant (-0.055 ± 0.649), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.602 ± 0.003, spirit=not significant (0.775 ± 0.404), mp5=not significant (1.370 ± 0.479)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (324.6 DPS) | yes | Dreadmist Mask (16698, +0.00 DPS) [dungeon]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 21.0 spell_power points (4.75 DPS) | yes | Horizon Choker (13085, -1.19 DPS) [world_drop]; Mindburst Medallion (11196, -1.30 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.54 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (324.6 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Lieutenant Commander's Silk Mantle (227102, +0.00 DPS) [pvp]; Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.1 spell_power points (5.43 DPS) | yes | Runecloth Cloak (13860, -1.77 DPS) [crafted]; Sergeant Major's Cape (16337, -2.40 DPS) [pvp]; Big Voodoo Cloak (8216, -2.47 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | sim-verified (324.6 DPS) | yes | Field Marshal's Silk Vestments (231603, +0.00 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -2.40 DPS) [pvp]; Magister's Robes (16688, -2.85 DPS) [dungeon] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (324.6 DPS) | yes | Marshal's Silk Bracers (16438, +0.00 DPS) [pvp]; Runecloth Cuffs (254123, +0.00 DPS) [crafted]; Dreadmist Bracers (16703, -0.04 DPS) [dungeon] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (324.6 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (324.6 DPS) | yes | Marshal's Silk Sash (16439, +0.00 DPS) [pvp]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Dreadmist Belt (16702, -0.09 DPS) [dungeon] |
| legs | Devout Skirt (16694) | Stratholme: Baron Rivendare [dungeon] | sim-verified (324.6 DPS) | yes | Magister's Leggings (16687, +0.00 DPS) [dungeon]; Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (324.6 DPS) | yes | Wolfrunner Shoes (13101, +0.00 DPS) [world_drop]; Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Mooncloth Boots (15802, -0.03 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 29.0 spell_power points (6.54 DPS) | yes | Band of the Hierophant (13096, -2.58 DPS) [world_drop]; Eye of Adaegus (5266, -3.22 DPS) [world_drop]; Band of the Unicorn (7553, -3.61 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (324.6 DPS) | yes | Eye of Adaegus (5266, +0.00 DPS) [world_drop]; Band of the Hierophant (13096, +0.00 DPS) [world_drop]; Band of the Unicorn (7553, -0.34 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, -0.08 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (324.6 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (324.6 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 354.5 spell_power points (80.04 DPS) | yes | Stormrager (16997, -16.33 DPS) [quest]; Brilliant Wand (249385, -17.00 DPS) [crafted]; Torch of Austen (13004, -23.26 DPS) [world_drop] |

**New at 60:** head: Devout Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Robe of the Archmage; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Devout Skirt; feet: Omnicast Boots; finger1: Maiden's Circle; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60, raid preset (gnome, 131205213100301531-30000000000000000-053500030000000000)

Set DPS (verified): 585.2. Weights run: 1.6s. Verify run: 3.4s. 1156 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.670 ± 0.029, crit=0.484 ± 0.011 per rating point (14 rating = 1%, 6.774 per %), hit=0.960 ± 0.049 per rating point (10 rating = 1%, 9.601 per %), spell_haste=3.033 ± 0.564, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.632 ± 0.002, spirit=not significant (-0.076 ± 0.281), mp5=not significant (-0.176 ± 0.311)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (585.2 DPS) | yes | Red Mageweave Headband (10033, +0.00 DPS) [crafted]; Lieutenant Commander's Silk Cowl (227103, +0.00 DPS) [pvp]; Field Marshal's Coronet (231604, +0.00 DPS) [pvp] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.7 spell_power points (8.46 DPS) | yes | Glowing Eye of Mordresh (10769, -2.87 DPS) [dungeon]; Mindburst Medallion (11196, -3.93 DPS) [quest]; Horizon Choker (13085, -4.22 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 32.5 spell_power points (14.70 DPS) | yes | Field Marshal's Silk Spaulders (231602, +0.00 DPS) [pvp]; Lieutenant Commander's Silk Mantle (227102, -1.51 DPS) [pvp]; Rotgrip Mantle (17732, -3.36 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.0 spell_power points (8.15 DPS) | yes | Runecloth Cloak (13860, -1.66 DPS) [crafted]; Big Voodoo Cloak (8216, -3.16 DPS) [crafted]; Cloak of Flames (3475, -3.81 DPS) [world_drop] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 54.8 spell_power points (24.80 DPS) | yes | Field Marshal's Silk Vestments (231603, -1.65 DPS) [pvp]; Knight-Captain's Silk Tunic (227108, -7.08 DPS) [pvp]; Acumen Robes (17775, -10.14 DPS) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 17.7 spell_power points (8.01 DPS) | yes | Nethershine Cuffs (254069, -2.17 DPS) [crafted]; Marshal's Silk Bracers (16438, -2.25 DPS) [pvp]; Dreadmist Bracers (16703, -2.41 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.2 spell_power points (13.23 DPS) | yes | Marshal's Silk Gloves (16440, +0.00 DPS) [vendor]; Marshal's Silk Gauntlets (231608, +0.00 DPS) [vendor]; Marshal's Silk Handwraps (231609, +0.00 DPS) [vendor] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 21.1 spell_power points (9.53 DPS) | yes | Satyrmane Sash (17755, -0.17 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -0.31 DPS) [crafted]; Dawnspire Cord (12466, -1.06 DPS) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Knight-Captain's Silk Leggings (16414, +0.00 DPS) [pvp]; Knight-Captain's Silk Legguards (227109, +0.00 DPS) [pvp]; Marshal's Silk Leggings (231605, +0.00 DPS) [pvp] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 28.0 spell_power points (12.68 DPS) | yes | Marshal's Silk Footwraps (231606, +0.00 DPS) [pvp]; Knight-Lieutenant's Silk Walkers (227112, -0.15 DPS) [pvp]; Earthen Silk Slippers (254013, -1.83 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 22.7 spell_power points (10.26 DPS) | yes | Band of the Unicorn (7553, -4.38 DPS) [world_drop]; Blackstone Ring (17713, -5.92 DPS) [dungeon]; Snake Hoop (6750, -6.03 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (6.04 DPS) | yes | Band of the Unicorn (7553, -0.16 DPS) [world_drop]; Blackstone Ring (17713, -1.70 DPS) [dungeon]; Snake Hoop (6750, -1.81 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -1.19 DPS) [world_drop] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (585.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (585.2 DPS) | yes | Spellshifter Rod (9527, +0.00 DPS) [quest]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 175.7 spell_power points (79.48 DPS) | yes | Brilliant Wand (249385, -13.70 DPS) [crafted]; Stormrager (16997, -15.27 DPS) [quest]; Greater Eternal Wand (249237, -22.07 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; trinket2: Mark of the Chosen; main_hand: Glowing Brightwood Staff; ranged: Torch of Light

No-known-source sample (15 of 1156, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (orc, 141005000000000000-00000000000000000-000000000000000000)

Set DPS (verified): 34.0. Weights run: 1.1s. Verify run: 0.9s. 179 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.005, intellect=0.404 ± 0.008, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.902 per %), hit=0.254 ± 0.011 per rating point (10 rating = 1%, 2.539 per %), spell_haste=not significant (0.159 ± 0.090), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.641 ± 0.006, spirit=not significant (0.136 ± 0.089), mp5=not significant (0.194 ± 0.105)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pearly Circlet (253957) | Tailoring [crafted] | 7.1 spell_power points (0.85 DPS) | yes | Pristine Circlet (253949, -0.29 DPS, sim-verified) [crafted]; Shadow Goggles (4373, -0.51 DPS) [crafted]; Flying Tiger Goggles (4368, -0.78 DPS) [crafted] |
| neck | Roadwatcher's Confidence (281265) | Watching the Roads [quest] | 0.4 spell_power points (0.05 DPS) | yes | Scholarly Pendant (277203, -0.02 DPS) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 spell_power points (1.03 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.56 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.66 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.96 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.5 spell_power points (0.54 DPS) | yes | Pearl-clasped Cloak (5542, -0.16 DPS) [crafted]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Mystic Shawl (3449, -0.30 DPS) [quest] |
| chest | Filigreed Pearly Gown (253909) | Tailoring [crafted] | 8.2 spell_power points (0.98 DPS) | yes | Filigreed Pristine Gown (253901, -0.09 DPS) [crafted]; Gray Woolen Robe (2585, -0.26 DPS) [crafted]; Scholarly Robes (2034, -0.43 DPS) [dungeon] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Owlbeard Bracers (16981, -0.10 DPS) [quest]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.58 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.84 DPS) | yes | Gnoll Casting Gloves (892, -0.12 DPS) [world]; Pearly Gloves (253921, -0.16 DPS) [crafted]; Pristine Gloves (253913, -0.21 DPS) [crafted] |
| waist | Pearly Sash (253933) | Tailoring [crafted] | 6.1 spell_power points (0.73 DPS) | yes | Pristine Sash (253925, -0.06 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Novice Ardent's Sash (253887, -0.30 DPS) [crafted] |
| legs | Filigreed Pearly Leggings (253945) | Tailoring [crafted] | 10.0 spell_power points (1.20 DPS) | yes | Filigreed Pristine Leggings (253937, -0.29 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.36 DPS) [dungeon]; Colorful Kilt (10048, -0.52 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 spell_power points (1.03 DPS) | yes | Pearly Boots (253897, -0.44 DPS, sim-verified) [crafted]; Red Woolen Boots (4313, -0.47 DPS) [crafted]; Pristine Boots (253889, -0.53 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 2.4 spell_power points (0.29 DPS) | yes | Volcanic Rock Ring (12053, -0.15 DPS) [world_drop]; Furlfeather Ring (286745, -0.17 DPS) [world]; Band of Purification (12996, -0.19 DPS) [world_drop] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 2.0 spell_power points (0.24 DPS) | yes | Volcanic Rock Ring (12053, -0.10 DPS) [world_drop]; Furlfeather Ring (286745, -0.12 DPS) [world]; Band of Purification (12996, -0.14 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Crescent Staff (6505) | Leaders of the Fang [quest] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Staff of the Spire (1300, -0.05 DPS) [world]; Staff of Orgrimmar (15444, -0.05 DPS) [quest]; Twisted Chanter's Staff (890, -3.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Skycaller (12984) | World drop [world_drop] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Cookie's Stirring Rod (5198, -0.44 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -4.14 DPS) [quest]; Greater Magic Wand (11288, -4.28 DPS) [crafted] |

**New at 20:** head: Pearly Circlet; neck: Roadwatcher's Confidence; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pearly Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pearly Sash; legs: Filigreed Pearly Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Crescent Staff; ranged: Skycaller

No-known-source sample (15 of 179, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20430 Legionnaire's Sword; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (orc, 141005213100300000-00000000000000000-000000000000000000)

Set DPS (verified): 72.9. Weights run: 1.3s. Verify run: 1.0s. 313 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.538 ± 0.013, crit=0.076 ± 0.002 per rating point (14 rating = 1%, 1.063 per %), hit=0.295 ± 0.014 per rating point (10 rating = 1%, 2.946 per %), spell_haste=not significant (0.274 ± 0.227), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.551 ± 0.003, spirit=not significant (0.388 ± 0.168), mp5=not significant (0.710 ± 0.198)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 13.3 spell_power points (2.28 DPS) | yes | Enchanter's Cowl (4322, -0.33 DPS) [crafted]; Silk Headband (7050, -0.74 DPS) [crafted]; Nightsky Cowl (4039, -0.84 DPS) [world_drop] |
| neck | Kaleidoscope Chain (13084) | World drop [world_drop] | 3.7 spell_power points (0.64 DPS) | yes | Crystal Starfire Medallion (5003, -0.07 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.08 DPS) [world]; Pendant of Myzrael (4614, -0.24 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.8 spell_power points (2.37 DPS) | yes | Death Speaker Mantle (6685, -0.33 DPS) [dungeon]; Magician's Mantle (12998, -0.69 DPS) [world_drop]; Invoker's Mantle (215365, -0.71 DPS) [crafted] |
| back | Windsong Drape (15468) | Free at Last [quest] | 6.6 spell_power points (1.12 DPS) | yes | Glowing Thresher Cape (6901, -0.08 DPS) [dungeon]; Heavy Woolen Cloak (4311, -0.17 DPS) [crafted]; Hillman's Cloak (3719, -0.27 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | 18.6 spell_power points (3.18 DPS) | yes | Green Silk Armor (7065, -0.61 DPS, sim-verified) [crafted]; Mechbuilder's Overalls (9508, -0.70 DPS) [dungeon]; Tree Bark Jacket (1486, -0.95 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 11.1 spell_power points (1.91 DPS) | yes | Nightsky Wristbands (6407, -1.15 DPS) [world_drop]; Glowing Magical Bracelets (13106, -1.17 DPS) [world_drop]; Mindthrust Bracers (1974, -1.52 DPS, sim-verified) [dungeon] |
| hands | Ephemeral Grips (273034) | Arcane Anomaly [world] | 8.7 spell_power points (1.50 DPS) | yes | Jutebraid Gloves (10654, -0.01 DPS) [quest]; Gloves of Meditation (4318, -0.17 DPS) [crafted]; Gloves of Old (9395, -0.24 DPS) [world_drop] |
| waist | Belt of Arugal (6392) | Shadowfang Keep: Archmage Arugal [dungeon] | 12.6 spell_power points (2.15 DPS) | yes | Warsong Sash (16975, -0.27 DPS) [quest]; Crimson Silk Belt (7055, -0.48 DPS) [crafted]; Invoker's Cord (215366, -0.49 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 15.5 spell_power points (2.66 DPS) | yes | Pearly Leggings (253995, -0.35 DPS) [crafted]; Pristine Leggings (253987, -0.48 DPS) [crafted]; Filigreed Pearly Leggings (253945, -0.80 DPS) [crafted] |
| feet | Golden Slippers (254009) | Tailoring [crafted] | 12.5 spell_power points (2.14 DPS) | yes | Gilded Slippers (254001, -0.03 DPS) [crafted]; Acidic Walkers (9454, -0.28 DPS) [dungeon]; Glinteye Slippers (273024, -0.54 DPS) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.7 spell_power points (1.49 DPS) | yes | Solid Gold Band (279845, -0.47 DPS) [quest]; Electrocutioner Lagnut (9447, -0.51 DPS) [dungeon]; Runemender's Seal (273031, -0.51 DPS) [world] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 8.1 spell_power points (1.39 DPS) | yes | Electrocutioner Lagnut (9447, -0.41 DPS) [dungeon]; Runemender's Seal (273031, -0.41 DPS) [world]; Solid Gold Band (279845, -0.60 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | sim-verified (72.9 DPS) | yes | Gnarled Ash Staff (791, -0.02 DPS) [world_drop]; Twisted Chanter's Staff (890, -0.09 DPS) [world_drop]; Staff of the Friar (3415, -4.03 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Starfaller (13063) | World drop [world_drop] | 197.2 spell_power points (33.79 DPS) | yes | Wand of Mana Concentration (273043, -0.63 DPS, sim-verified) [world]; Gyromatic Icemaker (9489, -3.44 DPS) [dungeon]; Greater Mystic Wand (217287, -3.94 DPS) [crafted] |

**New at 30:** head: Holy Shroud; neck: Kaleidoscope Chain; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Ephemeral Grips; waist: Belt of Arugal; legs: Gaze Dreamer Pants; feet: Golden Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Glimmering Staff; ranged: Starfaller

No-known-source sample (15 of 313, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19513 Legionnaire's Band

### Band 40 (orc, 141005213100301531-00000000000000000-000000000000000000)

Set DPS (verified): 152.7. Weights run: 1.3s. Verify run: 1.0s. 411 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.739 ± 0.026, crit=0.158 ± 0.004 per rating point (14 rating = 1%, 2.213 per %), hit=0.385 ± 0.024 per rating point (10 rating = 1%, 3.847 per %), spell_haste=not significant (0.100 ± 0.325), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.590 ± 0.003, spirit=not significant (0.910 ± 0.357), mp5=not significant (0.420 ± 0.395)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 28.7 spell_power points (5.97 DPS) | yes | Electromagnetic Gigaflux Reactivator (9492, -1.38 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -1.61 DPS) [crafted]; Whitemane's Chapeau (7720, -1.65 DPS, sim-verified) [dungeon] |
| neck | Triune Amulet (7722) | Scarlet Monastery: High Inquisitor Whitemane [dungeon] | 14.4 spell_power points (2.98 DPS) | yes | Glowing Eye of Mordresh (10769, -0.30 DPS) [dungeon]; Necklace of Calisea (1714, -0.59 DPS) [dungeon]; Amberglow Talisman (10824, -1.10 DPS) [quest] |
| shoulder | Mantle of Doan (7712) | Scarlet Monastery: Arcanist Doan [dungeon] | 18.7 spell_power points (3.89 DPS) | yes | Inquisitor's Shawl (19507, -0.44 DPS) [dungeon]; Green Silken Shoulders (7057, -0.54 DPS) [crafted]; Bloodmage Mantle (7684, -0.64 DPS) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 12.5 spell_power points (2.59 DPS) | yes | Cloak of Blight (6832, -0.13 DPS) [quest]; Glowing Thresher Cape (6901, -0.45 DPS) [dungeon]; Long Silken Cloak (4326, -0.57 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 37.4 spell_power points (7.77 DPS) | yes | Robe of the Magi (1716, -1.33 DPS) [world_drop]; Robe of Power (7054, -1.50 DPS) [crafted]; Beguiler Robes (7728, -2.54 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 10.3 spell_power points (2.13 DPS) | yes | Radiant Silver Bracers (4545, -0.07 DPS) [quest]; Earthen Silk Cuffs (254019, -0.36 DPS) [crafted]; Silkstream Cuffs (16791, -0.61 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 27.3 spell_power points (5.68 DPS) | yes | Red Mageweave Gloves (10018, -2.32 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.56 DPS) [crafted]; Gloves of Old (9395, -2.64 DPS) [world_drop] |
| waist | Golden Cord (254045) | Tailoring [crafted] | 19.6 spell_power points (4.08 DPS) | yes | Gilded Cord (254037, -0.05 DPS) [crafted]; Deathmage Sash (10771, -0.32 DPS) [dungeon]; Star Belt (4329, -0.62 DPS) [crafted] |
| legs | Stoneweaver Leggings (9407) | Uldaman: Ironaya [dungeon] | 23.8 spell_power points (4.95 DPS) | yes | Red Mageweave Pants (10009, -0.20 DPS) [crafted]; Black Mageweave Leggings (9999, -0.64 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.76 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 30.4 spell_power points (6.31 DPS) | yes | Thoughtcast Boots (10578, -2.81 DPS) [dungeon]; Golden Slippers (254009, -2.89 DPS) [crafted]; Furen's Boots (13100, -4.96 DPS, sim-verified) [world_drop] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.7 spell_power points (2.85 DPS) | yes | Snake Hoop (6750, -0.83 DPS) [quest]; Electrocutioner Lagnut (9447, -0.90 DPS) [dungeon]; Runemender's Seal (273031, -0.90 DPS) [world] |
| finger2 | Welken Ring (5011) | World drop [world_drop] | 10.1 spell_power points (2.09 DPS) | yes | Snake Hoop (6750, -0.08 DPS) [quest]; Electrocutioner Lagnut (9447, -0.14 DPS) [dungeon]; Runemender's Seal (273031, -0.14 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | sim-verified (152.7 DPS) | yes | Gnarled Ash Staff (791, -3.32 DPS) [world_drop]; Staff of the Friar (3415, -3.43 DPS) [dungeon]; Windweaver Staff (7757, -20.97 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 196.7 spell_power points (40.86 DPS) | yes | Nether Force Wand (11263, -1.97 DPS) [quest]; Icefury Wand (7514, -3.21 DPS) [quest]; Ragefire Wand (7513, -3.26 DPS) [quest] |

**New at 40:** head: Papal Fez; neck: Triune Amulet; shoulder: Mantle of Doan; back: Icy Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Golden Cord; legs: Stoneweaver Leggings; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Welken Ring; trinket1: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (orc, 141005213100301531-32023000000000000-000000000000000000)

Set DPS (verified): 195.4. Weights run: 1.3s. Verify run: 1.2s. 527 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.950 ± 0.040, crit=0.256 ± 0.008 per rating point (14 rating = 1%, 3.585 per %), hit=0.702 ± 0.041 per rating point (10 rating = 1%, 7.020 per %), spell_haste=not significant (-0.549 ± 0.496), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.621 ± 0.003, spirit=not significant (1.032 ± 0.437), mp5=not significant (1.782 ± 0.511)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dreamweave Circlet (10041) | Tailoring [crafted] | 42.9 spell_power points (8.25 DPS) | yes | Red Mageweave Headband (10033, -0.94 DPS) [crafted]; Papal Fez (9431, -1.82 DPS) [dungeon]; Gemburst Circlet (10751, -2.00 DPS, sim-verified) [quest] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.4 spell_power points (3.35 DPS) | yes | Mindburst Medallion (11196, -0.11 DPS) [quest]; Triune Amulet (7722, -0.43 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -0.54 DPS) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 30.1 spell_power points (5.79 DPS) | yes | Nethergeld Shoulders (254049, -1.41 DPS) [crafted]; Red Mageweave Shoulders (10029, -1.70 DPS) [crafted]; Nethershine Shoulders (254057, -2.61 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 25.9 spell_power points (4.98 DPS) | yes | Runecloth Cloak (13860, -1.79 DPS) [crafted]; Imperial Red Cloak (8248, -2.18 DPS) [world_drop]; Deep Woodlands Cloak (19121, -2.62 DPS, sim-verified) [quest] |
| chest | Ghostweave Vest (14141) | Tailoring [crafted] | 43.9 spell_power points (8.45 DPS) | yes | Dreamweave Vest (10021, -0.56 DPS) [crafted]; Acumen Robes (17775, -1.13 DPS) [quest]; Robe of Power (7054, -1.97 DPS) [crafted] |
| wrist | Nethershine Cuffs (254069) | Tailoring [crafted] | 19.9 spell_power points (3.83 DPS) | yes | Nethergeld Cuffs (254061, -0.21 DPS) [crafted]; Bloodband Bracers (11469, -0.62 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.04 DPS) [quest] |
| hands | Runecloth Gloves (13863) | Tailoring [crafted] | 29.8 spell_power points (5.74 DPS) | yes | Golden Gloves (254103, -0.06 DPS) [crafted]; Dreamweave Gloves (10019, -0.16 DPS) [crafted]; Gilded Gloves (254095, -0.40 DPS) [crafted] |
| waist | Ghostweave Cord (254073) | Tailoring [crafted] | 25.3 spell_power points (4.88 DPS) | yes | Dawnspire Cord (12466, -0.25 DPS) [dungeon]; Satyrmane Sash (17755, -0.36 DPS) [dungeon]; Golden Cord (254045, -0.55 DPS) [crafted] |
| legs | Kilt of the Atal'ai Prophet (10807) | Sunken Temple: Jammal'an the Prophet [dungeon] | 38.7 spell_power points (7.44 DPS) | yes | Senior Designer's Pantaloons (11841, -0.28 DPS) [dungeon]; Dalewind Trousers (13008, -0.83 DPS) [world_drop]; Stone Guard's Dreadweave Leggings (220906, -2.25 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 31.2 spell_power points (6.01 DPS) | yes | First Sergeant's Dreadweave Boots (220909, -1.47 DPS) [vendor]; Furen's Boots (13100, -1.55 DPS) [world_drop]; Vinerot Sandals (17748, -2.64 DPS, sim-verified) [dungeon] |
| finger1 | Eye of Adaegus (5266) | World drop [world_drop] | 18.1 spell_power points (3.48 DPS) | yes | Band of the Unicorn (7553, -0.98 DPS) [world_drop]; Coldwater Ring (4550, -1.14 DPS) [quest]; Welken Ring (5011, -1.18 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.8 spell_power points (2.84 DPS) | yes | Band of the Unicorn (7553, -0.34 DPS) [world_drop]; Coldwater Ring (4550, -0.50 DPS) [quest]; Welken Ring (5011, -0.53 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (195.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -0.27 DPS) [quest]; Uther's Strength (11302, -2.90 DPS, sim-verified) [world_drop] |
| trinket2 | Evonice's Landin' Pilla (18951) | Look at the Size of It! [quest] | sim-verified (195.4 DPS) | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Uther's Strength (11302, -0.83 DPS) [world_drop]; Rune of the Guard Captain (19120, -1.04 DPS) [quest] |
| main_hand | Spire of Hakkar (10844) | Avatar of Hakkar [world] | sim-verified (195.4 DPS) | yes | Spellshifter Rod (9527, -0.51 DPS) [quest]; Kindling Stave (11750, -0.68 DPS) [dungeon]; Inventor's Focal Sword (17719, -2.87 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 275.9 spell_power points (53.10 DPS) | yes | Noxious Shooter (17745, -2.13 DPS) [dungeon]; Wand of Allistarj (13065, -4.67 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.81 DPS) [crafted] |

**New at 50:** head: Dreamweave Circlet; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Ghostweave Vest; wrist: Nethershine Cuffs; hands: Runecloth Gloves; waist: Ghostweave Cord; legs: Kilt of the Atal'ai Prophet; finger1: Eye of Adaegus; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Evonice's Landin' Pilla; main_hand: Spire of Hakkar; ranged: Pyric Caduceus

No-known-source sample (15 of 527, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (orc, 141005213100301531-32023000000000000-053200000000000000)

Set DPS (verified): 313.7. Weights run: 1.4s. Verify run: 3.0s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.903 ± 0.043, crit=0.376 ± 0.010 per rating point (14 rating = 1%, 5.260 per %), hit=0.899 ± 0.045 per rating point (10 rating = 1%, 8.989 per %), spell_haste=not significant (-0.055 ± 0.649), spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.602 ± 0.003, spirit=not significant (0.775 ± 0.404), mp5=not significant (1.370 ± 0.479)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (313.7 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Dreadmist Mask (16698, -18.15 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 21.0 spell_power points (4.75 DPS) | yes | Horizon Choker (13085, -1.19 DPS) [world_drop]; Mindburst Medallion (11196, -1.30 DPS) [quest]; Glowing Eye of Mordresh (10769, -1.54 DPS) [dungeon] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 34.0 spell_power points (7.68 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -0.16 DPS) [pvp]; Shroud of the Nathrezim (18720, -0.25 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.1 spell_power points (5.43 DPS) | yes | Deep Woodlands Cloak (19121, -0.89 DPS) [quest]; Runecloth Cloak (13860, -1.77 DPS) [crafted]; Big Voodoo Cloak (8216, -2.47 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 56.1 spell_power points (12.66 DPS) | yes | Warlord's Silk Raiment (231596, +0.00 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -2.40 DPS) [pvp]; Magister's Robes (16688, -2.85 DPS) [dungeon] |
| wrist | Magister's Bindings (16683) | Blackrock Spire: Rage Talon Fire Tongue [dungeon] | sim-verified (313.7 DPS) | yes | General's Silk Cuffs (16538, +0.00 DPS) [pvp]; Devout Bracers (16697, +0.00 DPS) [dungeon]; Runecloth Cuffs (254123, -21.09 DPS, sim-verified) [crafted] |
| hands | Magister's Gloves (16684) | Scholomance: Doctor Theolen Krastinov [dungeon] | 30.5 spell_power points (6.88 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Gloves of Spell Mastery (14146, -23.39 DPS, sim-verified) [crafted] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | sim-verified (313.7 DPS) | yes | General's Silk Sash (16537, +0.00 DPS) [pvp]; Devout Belt (16696, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -19.29 DPS, sim-verified) [crafted] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 44.9 spell_power points (10.14 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Ghostweave Pants (14144, -25.77 DPS, sim-verified) [crafted] |
| feet | Wolfrunner Shoes (13101) | World drop [world_drop] | 31.5 spell_power points (7.10 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.14 DPS) [dungeon]; Mooncloth Boots (15802, -0.17 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 29.0 spell_power points (6.54 DPS) | yes | Eye of Adaegus (5266, -3.22 DPS) [world_drop]; Philanthropist's Ring (281635, -3.26 DPS) [quest]; Band of the Unicorn (7553, -3.61 DPS) [world_drop] |
| finger2 | Band of the Hierophant (13096) | World drop [world_drop] | 17.6 spell_power points (3.96 DPS) | yes | Eye of Adaegus (5266, -0.64 DPS) [world_drop]; Philanthropist's Ring (281635, -0.69 DPS) [quest]; Band of the Unicorn (7553, -1.03 DPS) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (313.7 DPS) | yes | Shard of the Splithooves (10659, -0.24 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.27 DPS) [crafted]; Evonice's Landin' Pilla (18951, -0.35 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (313.7 DPS) | yes | Shard of the Splithooves (10659, +0.00 DPS) [quest]; Evonice's Landin' Pilla (18951, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (313.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -0.30 DPS) [quest]; Spire of Hakkar (10844, -0.38 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 354.5 spell_power points (80.04 DPS) | yes | Stormrager (16997, -3.00 DPS, sim-verified) [quest]; Brilliant Wand (249385, -17.00 DPS) [crafted]; Torch of Austen (13004, -23.26 DPS) [world_drop] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; chest: Robe of the Archmage; wrist: Magister's Bindings; hands: Magister's Gloves; waist: Magister's Belt; legs: Magister's Leggings; feet: Wolfrunner Shoes; finger1: Maiden's Circle; finger2: Band of the Hierophant; trinket1: Ankh of Life; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60, raid preset (orc, 131205213100301531-30000000000000000-053500030000000000)

Set DPS (verified): 589.8. Weights run: 1.6s. Verify run: 3.4s. 1148 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.670 ± 0.029, crit=0.484 ± 0.011 per rating point (14 rating = 1%, 6.774 per %), hit=0.960 ± 0.049 per rating point (10 rating = 1%, 9.601 per %), spell_haste=3.033 ± 0.564, spell_penetration=not significant (0.000 ± 0.000), arcane_power=0.632 ± 0.002, spirit=not significant (-0.076 ± 0.281), mp5=not significant (-0.176 ± 0.311)

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Magister's Crown (16686) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (589.8 DPS) | yes | Champion's Silk Cowl (227105, +0.00 DPS) [pvp]; Warlord's Silk Cowl (231601, +0.00 DPS) [pvp]; Red Mageweave Headband (10033, -6.65 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.7 spell_power points (8.46 DPS) | yes | Glowing Eye of Mordresh (10769, -2.87 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.93 DPS) [quest]; Horizon Choker (13085, -4.22 DPS) [world_drop] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 32.5 spell_power points (14.70 DPS) | yes | Warlord's Silk Amice (231594, +0.00 DPS) [pvp]; Champion's Silk Mantle (227104, -1.51 DPS) [pvp]; Rotgrip Mantle (17732, -6.33 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.0 spell_power points (8.16 DPS) | yes | Spritecaster Cape (11623, -0.00 DPS) [dungeon]; Runecloth Cloak (13860, -1.66 DPS) [crafted]; Big Voodoo Cloak (8216, -3.17 DPS) [crafted] |
| chest | Robe of the Archmage (14152) | Tailoring [crafted] | 54.8 spell_power points (24.80 DPS) | yes | Warlord's Silk Raiment (231596, -1.65 DPS) [pvp]; Legionnaire's Silk Tunic (227106, -7.08 DPS) [pvp]; Acumen Robes (17775, -13.49 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 17.7 spell_power points (8.01 DPS) | yes | Nethershine Cuffs (254069, -2.17 DPS) [crafted]; General's Silk Cuffs (16538, -2.25 DPS) [pvp]; Dreadmist Bracers (16703, -2.41 DPS) [dungeon] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 29.2 spell_power points (13.23 DPS) | yes | General's Silk Handguards (16540, +0.00 DPS) [vendor]; General's Silk Gauntlets (231599, +0.00 DPS) [vendor]; Earth Warder's Gloves (21318, -4.62 DPS, sim-verified) [quest] |
| waist | Magister's Belt (16685) | Blackrock Spire: Smolderthorn Mystic [dungeon] | 21.1 spell_power points (9.53 DPS) | yes | Wisdom of the Timbermaw (19047, -0.31 DPS) [crafted]; Dawnspire Cord (12466, -1.06 DPS) [dungeon]; Satyrmane Sash (17755, -8.84 DPS, sim-verified) [dungeon] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | sim-verified (589.8 DPS) | yes | Legionnaire's Silk Legguards (227107, +0.00 DPS) [pvp]; General's Silk Trousers (231595, +0.00 DPS) [pvp]; Stone Guard's Dreadweave Leggings (220906, -13.99 DPS, sim-verified) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 28.0 spell_power points (12.68 DPS) | yes | General's Silk Boots (231597, +0.00 DPS) [pvp]; Blood Guard's Silk Walkers (227110, -0.15 DPS) [pvp]; Earthen Silk Slippers (254013, -1.83 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 22.7 spell_power points (10.26 DPS) | yes | Band of the Unicorn (7553, -4.38 DPS) [world_drop]; Blackstone Ring (17713, -5.92 DPS) [dungeon]; Snake Hoop (6750, -6.03 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.4 spell_power points (6.04 DPS) | yes | Band of the Unicorn (7553, -0.16 DPS) [world_drop]; Blackstone Ring (17713, -1.70 DPS) [dungeon]; Snake Hoop (6750, -1.81 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (589.8 DPS) | yes | Uther's Strength (11302, -1.19 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (589.8 DPS) | yes | Uther's Strength (11302, -0.33 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (589.8 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -3.19 DPS, sim-verified) [quest]; Inventor's Focal Sword (17719, -3.64 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 175.7 spell_power points (79.48 DPS) | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -15.27 DPS) [quest]; Greater Eternal Wand (249237, -22.07 DPS) [crafted] |

**New at 60:** head: Magister's Crown; neck: Archlight Talisman; shoulder: Mantle of the Timbermaw; back: Deep Woodlands Cloak; chest: Robe of the Archmage; wrist: Runecloth Cuffs; hands: Gloves of Spell Mastery; waist: Magister's Belt; legs: Magister's Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Torch of Light

No-known-source sample (15 of 1148, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

