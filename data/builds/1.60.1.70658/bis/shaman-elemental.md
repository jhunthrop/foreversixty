# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 25.0. Weights run: 2.6s. Verify run: 1.5s. 213 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.655 ± 0.009, crit=0.073 ± 0.002 per rating point (14 rating = 1%, 1.024 per %), hit=0.204 ± 0.006 per rating point (10 rating = 1%, 2.038 per %), spell_haste=-1.772 ± 0.076, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.32 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (0.91 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.58 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.00 DPS) [crafted]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Caretaker's Cape (20428, -0.08 DPS) [rep] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.03 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.05 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.05 DPS) [vendor]; Owl Bracers (4796, -1.41 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.6 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Windfelt Gloves (5630, -0.14 DPS) [quest]; Serpent Gloves (5970, -0.35 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.44 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.19 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -0.25 DPS) [quest]; Wisdom's Leather Pants (252503, -0.28 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (0.80 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.30 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.3 spell_power points (0.53 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop]; The 1 Ring (8350, -0.47 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.42 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; The 1 Ring (8350, -0.36 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Twisted Chanter's Staff (890, -0.12 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.34 DPS) [world]; Staff of Westfall (2042, -0.40 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 213, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 30 (dwarf, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 48.1. Weights run: 2.9s. Verify run: 1.5s. 388 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.639 ± 0.011, crit=0.089 ± 0.002 per rating point (14 rating = 1%, 1.253 per %), hit=0.254 ± 0.008 per rating point (10 rating = 1%, 2.537 per %), spell_haste=-2.308 ± 0.136, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.29 DPS) | yes | Holy Shroud (2721, -0.15 DPS) [world_drop]; Enduring Cap (3020, -0.23 DPS) [world_drop]; Totemic Leather Helm (252456, -0.68 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.13 DPS) | yes | Darkspear Warding Pendant (272075, -0.80 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.12 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (1.54 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Fairywing Mantle (9536, -0.31 DPS) [quest]; Magician's Mantle (12998, -0.42 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 5.1 spell_power points (0.53 DPS) | yes | Hillman's Cloak (3719, -0.01 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.47 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (1.81 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Guardian Armor (4256, -0.20 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.94 DPS) | yes | Nightsky Wristbands (6407, -0.54 DPS) [world_drop]; Technician's Bracers (270042, -0.54 DPS) [quest]; Glowing Magical Bracelets (13106, -0.95 DPS, sim-verified) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.6 spell_power points (0.79 DPS) | yes | Town Clerk's Mittens (270029, -0.01 DPS) [quest]; Shilly Mitts (9609, -0.06 DPS) [quest]; Gloves of Insight (9698, -0.06 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.8 spell_power points (1.44 DPS) | yes | Moss Cinch (6911, -0.19 DPS) [dungeon]; Belt of Arugal (6392, -0.30 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.38 DPS, sim-verified) [rep] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (1.47 DPS) | yes | Stormrider's Leather Pants (252502, -0.03 DPS) [crafted]; Guardian Pants (5962, -0.10 DPS) [crafted]; Dark Ritual Leggings (270031, -0.58 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.20 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.24 DPS) [crafted]; Acidic Walkers (9454, -1.32 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (0.96 DPS) | yes | Minor Channeling Ring (1449, -0.30 DPS) [quest]; Ring of Power Regulation (273030, -0.33 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.73 DPS) | yes | Ring of Power Regulation (273030, -0.10 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Minor Channeling Ring (1449, -0.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (48.1 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -1.67 DPS) [quest]; Manual Crowd Pummeler (9449, -7.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Lorekeeper's Ring; main_hand: Spine of the Devourer

No-known-source sample (15 of 388, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots

### Band 40 (dwarf, 5510331300001350-010000000000000000-0000000000000000)

Set DPS (verified): 66.2. Weights run: 3.4s. Verify run: 2.1s. 609 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.978 ± 0.026, crit=0.256 ± 0.008 per rating point (14 rating = 1%, 3.587 per %), hit=0.425 ± 0.019 per rating point (10 rating = 1%, 4.254 per %), spell_haste=not significant (-0.644 ± 0.340), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.689 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 24.6 spell_power points (2.35 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.35 DPS) [crafted]; Augural Shroud (2620, -0.37 DPS) [world]; Big Voodoo Mask (8201, -1.67 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.8 spell_power points (1.42 DPS) | yes | Triune Amulet (7722, -0.65 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.76 DPS) [vendor]; Scorn's Icy Choker (23169, -1.09 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.7 spell_power points (1.88 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest]; Green Silken Shoulders (7057, -1.04 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.9 spell_power points (1.04 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Blackforge Cape (6424, -0.29 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.9 spell_power points (2.66 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.20 DPS) [crafted]; Big Voodoo Robe (8200, -0.49 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.9 spell_power points (1.13 DPS) | yes | Arcane Runed Bracers (4744, -0.27 DPS) [quest]; Spidertank Oilrag (9448, -0.27 DPS) [dungeon]; Turtle Scale Bracers (8198, -0.57 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.29 DPS) | yes | Red Mageweave Gloves (10018, -0.31 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.40 DPS) [crafted]; Dreamweave Gloves (10019, -1.05 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.7 spell_power points (2.07 DPS) | yes | Skycaller's Leather Belt (252522, -0.55 DPS) [crafted]; Gilded Cord (254037, -0.56 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.69 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.7 spell_power points (2.46 DPS) | yes | Crimson Silk Pantaloons (7062, -0.66 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.85 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.94 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.29 DPS) | yes | Skycaller's Leather Shoes (252532, -0.40 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.40 DPS) [crafted]; Mender's Leather Shoes (252533, -0.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.9 spell_power points (1.42 DPS) | yes | Lorekeeper's Ring (19524, -0.56 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.66 DPS) [quest]; Reedknot Ring (9622, -0.75 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.9 spell_power points (1.04 DPS) | yes | Lorekeeper's Ring (19524, +0.00 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -0.28 DPS) [quest]; Reedknot Ring (9622, -0.37 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (66.2 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Spine of the Devourer (273040, -1.51 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | 15.8 spell_power points (1.51 DPS) | yes | Thrash's Trash (276204, -0.17 DPS) [vendor]; Orb of Mystic Insight (249394, -0.28 DPS) [crafted]; Orb of the Forgotten Seer (7685, -1.57 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of Lorica

No-known-source sample (15 of 609, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 50 (dwarf, 5510331300001350-030000000000000000-5030000000000000)

Set DPS (verified): 109.6. Weights run: 3.3s. Verify run: 2.4s. 777 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.796 ± 0.027, crit=0.305 ± 0.010 per rating point (14 rating = 1%, 4.272 per %), hit=0.596 ± 0.026 per rating point (10 rating = 1%, 5.955 per %), spell_haste=not significant (0.442 ± 0.321), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.761 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 34.9 spell_power points (3.55 DPS) | yes | Dreamweave Circlet (10041, -0.61 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.81 DPS) [crafted]; Red Mageweave Headband (10033, -1.49 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (1.36 DPS) | yes | Horizon Choker (13085, -0.23 DPS) [world_drop]; Mindburst Medallion (11196, -0.26 DPS) [quest]; Scorn's Icy Choker (23169, -0.93 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 30.9 spell_power points (3.15 DPS) | yes | Rockshard Pauldrons (9411, -0.87 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -1.09 DPS) [crafted]; Rotgrip Mantle (17732, -2.54 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.8 spell_power points (1.91 DPS) | yes | Big Voodoo Cloak (8216, -0.67 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.78 DPS) [vendor]; Runecloth Cloak (13860, -1.42 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 34.9 spell_power points (3.55 DPS) | yes | Feathered Breastplate (8349, +0.00 DPS, sim-verified) [crafted]; Robe of the Magi (1716, -0.83 DPS) [world_drop]; Runecloth Tunic (13857, -0.93 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.0 spell_power points (1.83 DPS) | yes | Skycaller's Leather Bracers (252542, -0.24 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.24 DPS) [crafted]; Nethergeld Cuffs (254061, -0.55 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of the Greatfather (17721, -0.05 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.14 DPS) [crafted]; Raider Handguards (272102, -1.17 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.6 spell_power points (2.60 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.37 DPS) [dungeon]; Dawnspire Cord (12466, -0.45 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.6 spell_power points (2.40 DPS) | yes | Big Voodoo Pants (8202, -0.06 DPS) [crafted]; Wizardweave Leggings (14132, -0.46 DPS) [crafted]; Turtle Scale Leggings (8185, -0.47 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.44 DPS) | yes | Skycaller's Leather Boots (252471, -0.02 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.02 DPS) [crafted]; Greaves of Withering Despair (22240, -0.51 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.42 DPS) | yes | Snake Hoop (6750, -0.41 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.61 DPS) [quest]; Band of the Unicorn (7553, -2.90 DPS, sim-verified) [world_drop] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Snake Hoop (6750, -0.21 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.41 DPS) [quest]; Band of the Unicorn (7553, -1.24 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -2.05 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.08 DPS) [quest]; Kindling Stave (11750, -0.36 DPS) [dungeon]; Hammer of the Northern Wind (810, -23.50 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60 (dwarf, 5510331300001350-030000000000000000-5033520000000000)

Set DPS (verified): 151.6. Weights run: 3.2s. Verify run: 7.7s. 1487 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.953 ± 0.027, crit=0.330 ± 0.010 per rating point (14 rating = 1%, 4.627 per %), hit=0.600 ± 0.028 per rating point (10 rating = 1%, 6.003 per %), spell_haste=not significant (0.050 ± 0.378), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.620 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (151.6 DPS) | yes | Black Dragonscale Helm (252605, -0.14 DPS) [crafted]; Magister's Crown (16686, -0.28 DPS) [dungeon]; Blue Dragonscale Helm (252604, -6.64 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 27.4 spell_power points (3.04 DPS) | yes | Beads of Ogre Mojo (22149, -0.33 DPS) [quest]; Orb of the Darkmoon (19426, -0.60 DPS) [quest]; Chains of the Lich (23125, -0.60 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (151.6 DPS) | yes | Darkspear Shoulderpads (272103, +0.00 DPS) [vendor]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -2.94 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | sim-verified (151.6 DPS) | yes | Arcanoweave Cloak (272411, +0.00 DPS) [vendor]; Spritecaster Cape (11623, -0.42 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (151.6 DPS) | yes | Magister's Robes (16688, +0.00 DPS) [dungeon]; Chain of The Five Thunders (227024, +0.00 DPS) [vendor]; Vest of Elements (16666, -5.94 DPS, sim-verified) [dungeon] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 29.6 spell_power points (3.29 DPS) | yes | Runecloth Cuffs (254123, -1.01 DPS) [crafted]; Runic Leather Bracers (15092, -1.12 DPS) [crafted]; Bands of The Five Thunders (227017, -1.34 DPS) [vendor] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (151.6 DPS) | yes | Raider Handguards (272102, -0.50 DPS) [vendor]; Gloves of Undead Cleansing (23084, -0.69 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.0 spell_power points (5.77 DPS) | yes | Stormseeker's Girdle (272399, -1.45 DPS) [vendor]; Girdle of Insight (18504, -1.67 DPS) [crafted]; Belt of the Archmage (18405, -2.11 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 59.8 spell_power points (6.64 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.00 DPS) [vendor]; Red Dragonscale Leggings (252603, -1.03 DPS) [crafted]; Sentinel's Silk Leggings (237815, -2.10 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 36.3 spell_power points (4.02 DPS) | yes | Skycaller's Leather Boots (252471, -1.20 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -1.20 DPS) [crafted]; Omnicast Boots (11822, -3.44 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 42.6 spell_power points (4.73 DPS) | yes | Maiden's Circle (13001, -1.99 DPS) [world_drop]; Deep Rooted Ring (19109, -2.34 DPS) [world]; Channeler's Ring (272406, -2.58 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 32.5 spell_power points (3.61 DPS) | yes | Maiden's Circle (13001, -0.87 DPS) [world_drop]; Deep Rooted Ring (19109, -1.22 DPS) [world]; Channeler's Ring (272406, -1.46 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (151.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (151.6 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.89 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -0.95 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (151.6 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -2.75 DPS) [world_drop] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (3.22 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, +0.00 DPS) [vendor]; Spirit of Aquementas (11904, -1.55 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; back: Hide of the Wild; chest: Ironfeather Breastplate; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch

No-known-source sample (15 of 1487, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

### Band 60, raid preset (dwarf, 5510331300001350-030000000000000000-0533520000000000)

Set DPS (verified): 461.1. Weights run: 1.8s. Verify run: 4.9s. 1487 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.371 ± 0.016, crit=0.326 ± 0.008 per rating point (14 rating = 1%, 4.565 per %), hit=0.663 ± 0.036 per rating point (10 rating = 1%, 6.629 per %), spell_haste=2.841 ± 0.436, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.769 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (461.1 DPS) | yes | Soothsayer's Headdress (17740, -1.23 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -1.85 DPS) [crafted]; Blue Dragonscale Helm (252604, -7.10 DPS, sim-verified) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.67 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.86 DPS) [quest]; Beads of Ogre Mojo (22149, -1.79 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.1 spell_power points (14.63 DPS) | yes | Darkspear Shoulderguards (272958, -4.38 DPS) [vendor]; Argent Shoulders (19059, -4.78 DPS) [crafted]; Mantle of the Timbermaw (19050, -5.06 DPS, sim-verified) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.6 spell_power points (10.09 DPS) | yes | Hide of the Wild (18510, -3.11 DPS) [crafted]; Spritecaster Cape (11623, -3.69 DPS) [dungeon]; Stormpike Sage's Cloak (19086, -4.57 DPS) [rep] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | 30.4 spell_power points (11.99 DPS) | yes | Fernpulse Jerkin (15786, -0.57 DPS) [quest]; Chain of The Five Thunders (227024, -0.61 DPS) [vendor]; Acumen Robes (17775, -1.58 DPS) [quest] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | 25.0 spell_power points (9.84 DPS) | yes | Runecloth Cuffs (254123, -4.04 DPS) [crafted]; Runic Leather Bracers (15092, -4.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -4.87 DPS) [crafted] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | 24.6 spell_power points (9.70 DPS) | yes | Gloves of the Greatfather (17721, +0.00 DPS, sim-verified) [crafted]; Raider Handguards (272101, -0.33 DPS) [vendor]; Earth Warder's Gloves (21318, -1.04 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.8 spell_power points (15.68 DPS) | yes | Belt of the Archmage (18405, -4.50 DPS, sim-verified) [crafted]; Stormseeker's Girdle (272399, -5.87 DPS) [vendor]; Girdle of Insight (18504, -6.41 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.3 spell_power points (19.41 DPS) | yes | Sentinel's Silk Leggings (237815, -1.74 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.74 DPS) [vendor]; Pristine Scorpid Leggings (252606, -4.60 DPS) [crafted] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 26.9 spell_power points (10.62 DPS) | yes | Omnicast Boots (11822, -0.98 DPS) [dungeon]; Earthen Silk Slippers (254013, -1.16 DPS) [crafted]; Venomspew Footpads (275606, -1.95 DPS) [crafted] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 38.0 spell_power points (14.96 DPS) | yes | Maiden's Circle (13001, -6.85 DPS) [world_drop]; Blessed Band of Light (272407, -8.41 DPS) [vendor]; Deep Rooted Ring (19109, -8.77 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 26.1 spell_power points (10.28 DPS) | yes | Maiden's Circle (13001, -2.16 DPS) [world_drop]; Blessed Band of Light (272407, -3.73 DPS) [vendor]; Deep Rooted Ring (19109, -4.09 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+18.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -2.90 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.15 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.17 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -11.97 DPS) [dungeon] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 29.0 spell_power points (11.43 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Barricade (235473, -0.51 DPS) [vendor]; Spirit of Aquementas (11904, -3.87 DPS, sim-verified) [quest] |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, -2.58 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Storm Gauntlets; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Totem of Thunder

No-known-source sample (15 of 1487, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4800 Mighty Chain Pants; 4816 Legionnaire's Leggings; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4833 Glorious Shoulders

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.7. Weights run: 2.6s. Verify run: 1.4s. 193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.655 ± 0.009, crit=0.073 ± 0.002 per rating point (14 rating = 1%, 1.024 per %), hit=0.204 ± 0.006 per rating point (10 rating = 1%, 2.038 per %), spell_haste=-1.772 ± 0.076, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.33 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.9 spell_power points (0.91 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.58 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.81 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Battle Healer's Cloak (20427, -0.08 DPS) [rep]; Pearl-clasped Cloak (5542, -0.40 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.03 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.25 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.3 spell_power points (0.27 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.02 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.6 spell_power points (0.64 DPS) | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Pristine Gloves (253913, -0.14 DPS) [crafted]; Serpent Gloves (5970, -0.42 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.46 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.2 spell_power points (1.19 DPS) | yes | Stormrider's Leather Pants (252502, -0.19 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.28 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.36 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.6 spell_power points (0.80 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.30 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.42 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; The 1 Ring (8350, -0.36 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.9 spell_power points (0.33 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS) [quest]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; The 1 Ring (8350, -0.27 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.12 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.34 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 206376 Bulwark Icon; 206383 Lithic Icon; 206384 Astral Icon; 206385 Clastic Icon; 206389 Brimstone Icon; 206390 Ichthyan Icon; 206391 Syzygic Icon; 241089 Scarlet Dagger

### Band 30 (orc, 5510331300000000-000000000000000000-0000000000000000)

Set DPS (verified): 48.6. Weights run: 2.9s. Verify run: 1.5s. 368 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.639 ± 0.011, crit=0.089 ± 0.002 per rating point (14 rating = 1%, 1.253 per %), hit=0.254 ± 0.008 per rating point (10 rating = 1%, 2.537 per %), spell_haste=-2.308 ± 0.136, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 12.4 spell_power points (1.29 DPS) | yes | Holy Shroud (2721, -0.15 DPS) [world_drop]; Enduring Cap (3020, -0.23 DPS) [world_drop]; Totemic Leather Helm (252456, -0.82 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (1.13 DPS) | yes | Darkspear Warding Pendant (272075, -0.80 DPS) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -1.13 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.8 spell_power points (1.54 DPS) | yes | Death Speaker Mantle (6685, -0.18 DPS) [dungeon]; Fairywing Mantle (9536, -0.31 DPS) [quest]; Magician's Mantle (12998, -0.42 DPS) [world_drop] |
| back | Darkspear Raider's Cloak (272078) | Creeg Bothunk [vendor] | 5.1 spell_power points (0.53 DPS) | yes | Windsong Drape (15468, -0.01 DPS) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Hillman's Cloak (3719, -0.33 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 17.3 spell_power points (1.81 DPS) | yes | Beguiler Robes (7728, -0.07 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.17 DPS) [dungeon]; Guardian Armor (4256, -0.20 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.94 DPS) | yes | Nightsky Wristbands (6407, -0.54 DPS) [world_drop]; Technician's Bracers (270042, -0.54 DPS) [quest]; Glowing Magical Bracelets (13106, -0.77 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.3 spell_power points (1.07 DPS) | yes | Jutebraid Gloves (10654, -0.11 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.28 DPS) [crafted]; Serpent Gloves (5970, -0.34 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.8 spell_power points (1.44 DPS) | yes | Moss Cinch (6911, -0.19 DPS) [dungeon]; Warsong Sash (16975, -0.30 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.38 DPS, sim-verified) [rep] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.1 spell_power points (1.47 DPS) | yes | Stormrider's Leather Pants (252502, -0.03 DPS) [crafted]; Guardian Pants (5962, -0.10 DPS) [crafted]; Dark Ritual Leggings (270031, -0.42 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.5 spell_power points (1.20 DPS) | yes | Spidersilk Boots (4320, -0.20 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.24 DPS) [crafted]; Acidic Walkers (9454, -1.07 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 9.2 spell_power points (0.96 DPS) | yes | Ring of Power Regulation (273030, -0.33 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor]; Solid Gold Band (279845, -0.54 DPS) [quest] |
| finger2 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.73 DPS) | yes | Sea Giant's Toe Ring (274746, +0.00 DPS, sim-verified) [vendor]; Ring of Power Regulation (273030, -0.10 DPS) [dungeon]; Solid Gold Band (279845, -0.31 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (48.6 DPS) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, -2.92 DPS) [dungeon]; Manual Crowd Pummeler (9449, -7.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Darkspear Raider's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Advisor's Ring; main_hand: Spine of the Devourer

No-known-source sample (15 of 368, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (orc, 5510331300001350-010000000000000000-0000000000000000)

Set DPS (verified): 75.7. Weights run: 3.4s. Verify run: 2.1s. 570 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.978 ± 0.026, crit=0.256 ± 0.008 per rating point (14 rating = 1%, 3.587 per %), hit=0.425 ± 0.019 per rating point (10 rating = 1%, 4.254 per %), spell_haste=not significant (-0.644 ± 0.340), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.689 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | 24.6 spell_power points (2.35 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.35 DPS) [crafted]; Augural Shroud (2620, -0.37 DPS) [world]; Big Voodoo Mask (8201, -1.68 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 14.8 spell_power points (1.42 DPS) | yes | Triune Amulet (7722, -0.65 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.76 DPS) [vendor]; Scorn's Icy Choker (23169, -1.05 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.7 spell_power points (1.88 DPS) | yes | Bloodmage Mantle (7684, -0.18 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest]; Green Silken Shoulders (7057, -1.00 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 10.9 spell_power points (1.04 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Blackforge Cape (6424, -0.29 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.9 spell_power points (2.66 DPS) | yes | Dreamweave Vest (10021, -0.10 DPS) [crafted]; Robe of Power (7054, -0.20 DPS) [crafted]; Big Voodoo Robe (8200, -0.49 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.9 spell_power points (1.13 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.00 DPS) [quest]; Turtle Scale Bracers (8198, -0.09 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.29 DPS) | yes | Red Mageweave Gloves (10018, -0.31 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.40 DPS) [crafted]; Dreamweave Gloves (10019, -0.85 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.7 spell_power points (2.07 DPS) | yes | Skycaller's Leather Belt (252522, -0.55 DPS) [crafted]; Gilded Cord (254037, -0.56 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.40 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.7 spell_power points (2.46 DPS) | yes | Abomination Skin Leggings (23173, -0.85 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.94 DPS) [crafted]; Crimson Silk Pantaloons (7062, -1.14 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.29 DPS) | yes | Skycaller's Leather Shoes (252532, -0.40 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.40 DPS) [crafted]; Mender's Leather Shoes (252533, -0.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.9 spell_power points (1.42 DPS) | yes | Advisor's Ring (19520, -0.56 DPS) [rep]; Reedknot Ring (9622, -0.75 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.9 spell_power points (1.04 DPS) | yes | Advisor's Ring (19520, -0.18 DPS) [rep]; Reedknot Ring (9622, -0.37 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.19 DPS) [dungeon]; Spine of the Devourer (273040, -1.51 DPS) [world] |
| off_hand | Thrash's Trash (276204) | Friz Frazzlespark [vendor] | sim-verified (75.7 DPS) | yes | Orb of Mystic Insight (249394, -0.11 DPS) [crafted]; Anchorhold Buckler (15865, -0.30 DPS) [quest]; Orb of the Forgotten Seer (7685, -10.53 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Thrash's Trash

No-known-source sample (15 of 570, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7748 Forcestone Buckler; 8708 Hammer of Expertise

### Band 50 (orc, 5510331300001350-030000000000000000-5030000000000000)

Set DPS (verified): 109.6. Weights run: 3.3s. Verify run: 2.3s. 720 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.796 ± 0.027, crit=0.305 ± 0.010 per rating point (14 rating = 1%, 4.272 per %), hit=0.596 ± 0.026 per rating point (10 rating = 1%, 5.955 per %), spell_haste=not significant (0.442 ± 0.321), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.761 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 34.9 spell_power points (3.55 DPS) | yes | Dreamweave Circlet (10041, -0.61 DPS) [crafted]; Blood Guard's Pulsing Helmet (220848, -0.74 DPS) [vendor]; Red Mageweave Headband (10033, -1.07 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.4 spell_power points (1.36 DPS) | yes | Horizon Choker (13085, -0.23 DPS) [world_drop]; Mindburst Medallion (11196, -0.26 DPS) [quest]; Scorn's Icy Choker (23169, -1.14 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 30.9 spell_power points (3.15 DPS) | yes | Rockshard Pauldrons (9411, -0.87 DPS) [dungeon]; Skycaller's Leather Shoulders (252537, -1.09 DPS) [crafted]; Rotgrip Mantle (17732, -2.69 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.2 spell_power points (1.95 DPS) | yes | Spritecaster Cape (11623, -0.04 DPS) [dungeon]; Runecloth Cloak (13860, -0.39 DPS) [crafted]; Big Voodoo Cloak (8216, -0.71 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 34.9 spell_power points (3.55 DPS) | yes | Feathered Breastplate (8349, -0.81 DPS) [crafted]; Robe of the Magi (1716, -0.83 DPS) [world_drop]; Stone Guard's Pulsing Breastplate (220844, -1.01 DPS, sim-verified) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.0 spell_power points (1.83 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.24 DPS) [crafted] |
| hands | Raider Handguards (272102) | Creeg Bothunk [vendor] | 26.5 spell_power points (2.70 DPS) | yes | Raider Handwraps (272098, -0.21 DPS) [vendor]; Gloves of the Greatfather (17721, -0.26 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.34 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 25.6 spell_power points (2.60 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.37 DPS) [dungeon]; Dawnspire Cord (12466, -0.45 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 33.5 spell_power points (3.41 DPS) | yes | Big Voodoo Pants (8202, -1.07 DPS) [crafted]; Wizardweave Leggings (14132, -1.48 DPS) [crafted]; Red Mageweave Pants (10009, -2.08 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.44 DPS) | yes | Skycaller's Mail Sabatons (252577, -0.02 DPS) [crafted]; Greaves of Withering Despair (22240, -0.51 DPS) [dungeon]; Skycaller's Leather Boots (252471, -0.98 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 spell_power points (1.42 DPS) | yes | Advisor's Ring (19519, -0.20 DPS) [rep]; Snake Hoop (6750, -0.41 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (1.32 DPS) | yes | Advisor's Ring (19519, -0.10 DPS) [rep]; Snake Hoop (6750, -0.31 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (109.6 DPS) | yes | Rune of the Guard Captain (19120, -0.19 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (109.6 DPS) | yes | Rune of the Guard Captain (19120, -0.12 DPS) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (109.6 DPS) | yes | Spellshifter Rod (9527, -0.08 DPS) [quest]; Kindling Stave (11750, -0.36 DPS) [dungeon]; Hammer of the Northern Wind (810, -24.31 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handguards; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 720, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60 (orc, 5510331300001350-030000000000000000-5033520000000000)

Set DPS (verified): 149.2. Weights run: 3.2s. Verify run: 7.5s. 1410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.953 ± 0.027, crit=0.330 ± 0.010 per rating point (14 rating = 1%, 4.627 per %), hit=0.600 ± 0.028 per rating point (10 rating = 1%, 6.003 per %), spell_haste=not significant (0.050 ± 0.378), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.620 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Devout Crown (16693) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (149.2 DPS) | yes | Warlord's Mail Helm (231663, +0.00 DPS) [pvp]; Blue Dragonscale Helm (252604, +0.00 DPS) [crafted]; Coif of The Five Thunders (227002, -7.86 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 27.4 spell_power points (3.04 DPS) | yes | Beads of Ogre Mojo (22149, -0.33 DPS) [quest]; Orb of the Darkmoon (19426, -0.60 DPS) [quest]; Chains of the Lich (23125, -0.60 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (149.2 DPS) | yes | Warlord's Mail Spaulders (231659, +0.00 DPS) [pvp]; Darkspear Shoulderguards (272958, +0.00 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -5.61 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | sim-verified (149.2 DPS) | yes | Deep Woodlands Cloak (19121, -0.33 DPS) [quest]; Spritecaster Cape (11623, -0.42 DPS) [dungeon]; Arcanoweave Cloak (272411, -2.09 DPS, sim-verified) [vendor] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (149.2 DPS) | yes | Vest of Elements (16666, +0.00 DPS) [dungeon]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Vest of The Five Thunders (227004, -7.78 DPS, sim-verified) [quest] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (149.2 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [rep] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (149.2 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Raider Handwraps (272097, +0.00 DPS) [vendor]; Raider Handguards (272101, -7.00 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 52.0 spell_power points (5.77 DPS) | yes | Stormseeker's Girdle (272399, -1.45 DPS) [vendor]; Girdle of Insight (18504, -1.67 DPS) [crafted]; Belt of the Archmage (18405, -2.99 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 59.8 spell_power points (6.64 DPS) | yes | General's Mail Leggings (231664, -0.48 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -1.00 DPS) [vendor]; Sentinel's Silk Leggings (237815, -2.00 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 36.3 spell_power points (4.02 DPS) | yes | General's Mail Sabatons (231661, -0.01 DPS) [vendor]; General's Mail Boots (231667, -0.78 DPS) [pvp]; Omnicast Boots (11822, -2.63 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 42.6 spell_power points (4.73 DPS) | yes | Maiden's Circle (13001, -1.99 DPS) [world_drop]; Deep Rooted Ring (19109, -2.34 DPS) [world]; Channeler's Ring (272406, -2.58 DPS) [vendor] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 32.5 spell_power points (3.61 DPS) | yes | Deep Rooted Ring (19109, -1.22 DPS) [world]; Channeler's Ring (272406, -1.46 DPS) [vendor]; Maiden's Circle (13001, -1.47 DPS, sim-verified) [world_drop] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (149.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (149.2 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.89 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -0.95 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (149.2 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glowing Brightwood Staff (812, -2.75 DPS) [world_drop] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (3.22 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, +0.00 DPS) [vendor]; Spirit of Aquementas (11904, -1.70 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 60:** head: Devout Crown; neck: Amulet of the Dawn; shoulder: Devout Mantle; back: Hide of the Wild; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch

No-known-source sample (15 of 1410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 60, raid preset (orc, 5510331300001350-030000000000000000-0533520000000000)

Set DPS (verified): 453.6. Weights run: 1.8s. Verify run: 5.1s. 1410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.371 ± 0.016, crit=0.326 ± 0.008 per rating point (14 rating = 1%, 4.565 per %), hit=0.663 ± 0.036 per rating point (10 rating = 1%, 6.629 per %), spell_haste=2.841 ± 0.436, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.769 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 38.3 spell_power points (15.10 DPS) | yes | Warlord's Mail Helm (231663, -1.31 DPS) [pvp]; Living Crown (252561, -2.61 DPS) [crafted]; Coif of The Five Thunders (227002, -3.49 DPS, sim-verified) [quest] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.67 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.86 DPS) [quest]; Beads of Ogre Mojo (22149, -1.79 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.1 spell_power points (14.63 DPS) | yes | Warlord's Mail Spaulders (231659, -2.22 DPS) [pvp]; Mantle of the Timbermaw (19050, -4.23 DPS) [crafted]; Pauldrons of The Five Thunders (227003, -5.55 DPS, sim-verified) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 25.6 spell_power points (10.09 DPS) | yes | Hide of the Wild (18510, -3.11 DPS) [crafted]; Spritecaster Cape (11623, -3.69 DPS) [dungeon]; Deep Woodlands Cloak (19121, -4.04 DPS) [quest] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | 30.4 spell_power points (11.99 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Stone Guard's Pulsing Breastplate (220844, -0.53 DPS) [vendor] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 25.0 spell_power points (9.84 DPS) | yes | Runecloth Cuffs (254123, -4.04 DPS) [crafted]; Runic Leather Bracers (15092, -4.44 DPS) [crafted]; Skycaller's Leather Bracers (252542, -4.87 DPS) [crafted] |
| hands | Storm Gauntlets (12632) | Blacksmithing [crafted] | 24.6 spell_power points (9.70 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -0.19 DPS) [vendor]; Gloves of the Greatfather (17721, -0.25 DPS) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 39.8 spell_power points (15.68 DPS) | yes | Belt of the Archmage (18405, -5.33 DPS, sim-verified) [crafted]; Stormseeker's Girdle (272399, -5.87 DPS) [vendor]; Girdle of Insight (18504, -6.41 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.3 spell_power points (19.41 DPS) | yes | Sentinel's Silk Leggings (237815, -1.74 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.74 DPS) [vendor]; General's Mail Leggings (231664, -3.06 DPS) [pvp] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 26.9 spell_power points (10.62 DPS) | yes | General's Mail Sabatons (231661, -0.25 DPS) [vendor]; Earthen Silk Slippers (254013, -1.16 DPS) [crafted]; Omnicast Boots (11822, -2.30 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 38.0 spell_power points (14.96 DPS) | yes | Maiden's Circle (13001, -6.85 DPS) [world_drop]; Blessed Band of Light (272407, -8.41 DPS) [vendor]; Deep Rooted Ring (19109, -8.77 DPS) [world] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 26.1 spell_power points (10.28 DPS) | yes | Maiden's Circle (13001, -2.36 DPS, sim-verified) [world_drop]; Blessed Band of Light (272407, -3.73 DPS) [vendor]; Deep Rooted Ring (19109, -4.09 DPS) [world] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (453.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (453.6 DPS) | yes | Weakness Analyzer (272438, -2.91 DPS, sim-verified) [vendor]; Uther's Strength (11302, -3.15 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -3.17 DPS) [crafted] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (453.6 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -11.97 DPS) [dungeon] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 29.0 spell_power points (11.43 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Barricade (235474, -0.51 DPS) [vendor]; Spirit of Aquementas (11904, -3.88 DPS, sim-verified) [quest] |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (453.6 DPS) | yes | Totem of the Storm (272432, -2.52 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Storm Gauntlets; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Hand of Edward the Odd; off_hand: Therazane's Touch; ranged: Totem of Thunder

No-known-source sample (15 of 1410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

