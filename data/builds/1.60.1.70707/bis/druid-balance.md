# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 34.2. Weights run: 2.5s. Verify run: 1.3s. 203 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.735 ± 0.007, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.444 per %), hit=0.095 ± 0.007 per rating point (10 rating = 1%, 0.952 per %), spell_haste=not significant (-0.033 ± 0.098), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.025 ± 0.000, arcane_power=0.975 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.02 DPS) [crafted]; Totemic Leather Hood (252448, -0.58 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 spell_power points (1.79 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.76 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.17 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.65 DPS) | yes | Heavy Woolen Cloak (4311, -0.03 DPS) [crafted]; Black Whelp Cloak (7283, -0.19 DPS) [crafted]; Sanguine Cape (14376, -0.19 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.07 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.43 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.11 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.23 DPS) [dungeon]; Owl Bracers (4796, -0.55 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.9 spell_power points (1.22 DPS) | yes | Serpent Gloves (5970, -0.14 DPS) [dungeon]; Windfelt Gloves (5630, -0.27 DPS) [quest]; Pristine Gloves (253913, -0.27 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.15 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.80 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.4 spell_power points (2.22 DPS) | yes | Wisdom's Leather Pants (252503, -0.46 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted]; Dreamer's Leggings (270016, -1.09 DPS, sim-verified) [quest] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (1.53 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.35 DPS) [crafted]; Totemic Leather Boots (252442, -0.61 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (1.00 DPS) | yes | Volcanic Rock Ring (12053, -0.66 DPS) [world_drop]; The 1 Ring (8350, -0.88 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.4 spell_power points (0.68 DPS) | yes | Volcanic Rock Ring (12053, -0.54 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.57 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.23 DPS) | yes | Lesser Staff of the Spire (1300, -0.55 DPS) [world]; Twisted Chanter's Staff (890, -0.58 DPS, sim-verified) [world_drop]; Staff of Westfall (2042, -0.67 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 66.5. Weights run: 2.7s. Verify run: 1.6s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.522 ± 0.008, crit=0.049 ± 0.002 per rating point (14 rating = 1%, 0.691 per %), hit=0.153 ± 0.009 per rating point (10 rating = 1%, 1.534 per %), spell_haste=not significant (-0.128 ± 0.117), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.322 ± 0.001, arcane_power=0.678 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.51 DPS) | yes | Enchanter's Cowl (4322, -0.16 DPS) [crafted]; Holy Shroud (2721, -0.21 DPS) [world_drop]; Silk Headband (7050, -0.63 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.1 spell_power points (0.66 DPS) | yes | Crystal Starfire Medallion (5003, -0.22 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.7 spell_power points (2.87 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS) [dungeon]; Fairywing Mantle (9536, -0.63 DPS) [quest]; Magician's Mantle (12998, -0.84 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.05 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted]; Prelacy Cape (7004, -0.21 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.21 DPS) [dungeon]; Guardian Armor (4256, -0.22 DPS) [crafted]; Green Silk Armor (7065, -1.15 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.88 DPS) | yes | Nightsky Wristbands (6407, -1.23 DPS) [world_drop]; Technician's Bracers (270042, -1.23 DPS) [quest]; Glowing Magical Bracelets (13106, -1.90 DPS, sim-verified) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.1 spell_power points (1.48 DPS) | yes | Serpent Gloves (5970, -0.02 DPS) [dungeon]; Shilly Mitts (9609, -0.02 DPS) [quest]; Gloves of Insight (9698, -0.02 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.1 spell_power points (2.75 DPS) | yes | Belt of Arugal (6392, -0.54 DPS) [dungeon]; Crimson Silk Belt (7055, -0.73 DPS) [crafted]; Moss Cinch (6911, -1.09 DPS, sim-verified) [dungeon] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (2.93 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Guardian Pants (5962, -0.38 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.42 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.7 spell_power points (2.23 DPS) | yes | Spidersilk Boots (4320, -0.33 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.43 DPS) [crafted]; Acidic Walkers (9454, -1.74 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.6 spell_power points (1.80 DPS) | yes | Lavishly Jeweled Ring (1156, -1.15 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.17 DPS) [dungeon]; Minor Channeling Ring (1449, -2.07 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.63 DPS) [dungeon]; Minor Channeling Ring (1449, -1.01 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.35 DPS) [quest]; Manual Crowd Pummeler (9449, -11.45 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 5232221115100141-00000000000000000000-0000000000000000)

Set DPS (verified): 99.1. Weights run: 2.8s. Verify run: 1.6s. 523 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.344 ± 0.009, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.975 per %), hit=0.198 ± 0.013 per rating point (10 rating = 1%, 1.980 per %), spell_haste=not significant (0.415 ± 0.115), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.28 DPS) | yes | Augural Shroud (2620, -0.99 DPS, sim-verified) [world]; Papal Fez (9431, -1.80 DPS) [dungeon]; Big Voodoo Mask (8201, -1.81 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.7 spell_power points (2.45 DPS) | yes | Triune Amulet (7722, -1.35 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.85 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.93 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (3.04 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Inquisitor's Shawl (19507, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.42 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.7 spell_power points (1.94 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.18 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.68 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.1 spell_power points (6.05 DPS) | yes | Dreamweave Vest (10021, -0.75 DPS) [crafted]; Elemental Raiment (9434, -0.77 DPS) [world_drop]; Robe of Power (7054, -1.49 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (99.1 DPS) | yes | Guardian Leather Bracers (4260, -0.24 DPS) [crafted]; Condor Bracers (15864, -0.50 DPS) [quest]; Arcane Runed Bracers (4744, -1.20 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (6.04 DPS) | yes | Black Mageweave Gloves (10003, -2.26 DPS) [crafted]; Red Mageweave Gloves (10018, -2.41 DPS) [crafted]; Dreamweave Gloves (10019, -2.47 DPS, sim-verified) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.27 DPS) | yes | Deathmage Sash (10771, -0.21 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.24 DPS) [crafted]; Moss Cinch (6911, -0.25 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.1 spell_power points (4.56 DPS) | yes | Crimson Silk Pantaloons (7062, -1.42 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.53 DPS) [crafted]; Dark Ritual Leggings (270031, -2.92 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.04 DPS) | yes | Skycaller's Leather Shoes (252532, -2.29 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -3.42 DPS) [crafted]; Gilded Slippers (254001, -3.67 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (2.95 DPS) | yes | Snake Hoop (6750, -1.01 DPS) [quest]; Reedknot Ring (9622, -1.19 DPS) [quest]; Ring of Power Regulation (273030, -1.44 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (2.01 DPS) | yes | Snake Hoop (6750, -0.07 DPS) [quest]; Reedknot Ring (9622, -0.25 DPS) [quest]; Ring of Power Regulation (273030, -0.50 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.17 DPS) [dungeon]; Spine of the Devourer (273040, -2.70 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.52 DPS) | yes | Orb of Lorica (11262, -1.15 DPS) [quest]; Orb of Mystic Insight (249394, -1.24 DPS) [crafted]; Beacon of Hope (9393, -1.51 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Talons of Wrath

No-known-source sample (15 of 523, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 5232221115100151-00000000000000000000-5400000000000000)

Set DPS (verified): 158.1. Weights run: 3.0s. Verify run: 1.7s. 693 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.477 ± 0.013, crit=0.104 ± 0.004 per rating point (14 rating = 1%, 1.460 per %), hit=0.303 ± 0.020 per rating point (10 rating = 1%, 3.032 per %), spell_haste=not significant (0.348 ± 0.161), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.508 ± 0.001, arcane_power=0.492 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.2 spell_power points (7.14 DPS) | yes | Red Mageweave Headband (10033, -0.38 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.75 DPS) [crafted]; Dreamweave Circlet (10041, -1.04 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.8 spell_power points (2.56 DPS) | yes | Mindburst Medallion (11196, -0.46 DPS) [quest]; Triune Amulet (7722, -0.67 DPS) [dungeon]; Horizon Choker (13085, -0.98 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.5 spell_power points (5.81 DPS) | yes | Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.12 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.72 DPS) [crafted]; Rotgrip Mantle (17732, -3.03 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.9 spell_power points (3.99 DPS) | yes | Runecloth Cloak (13860, -1.75 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -1.79 DPS) [crafted]; Nightfall Drape (12465, -1.86 DPS) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.5 spell_power points (6.76 DPS) | yes | Feathered Breastplate (8349, -1.13 DPS) [crafted]; Dreamweave Vest (10021, -1.48 DPS) [crafted]; Robe of the Magi (1716, -2.09 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.8 spell_power points (3.50 DPS) | yes | Skycaller's Leather Bracers (252542, -0.34 DPS) [crafted]; Nethergeld Cuffs (254061, -1.05 DPS) [crafted]; Mender's Leather Bracers (252543, -1.29 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (5.68 DPS) | yes | Dreamweave Gloves (10019, -0.97 DPS) [crafted]; Sergeant Major's Crackling Leather Gauntlets (220866, -1.13 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -1.79 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.7 spell_power points (5.15 DPS) | yes | Satyrmane Sash (17755, -0.70 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.42 DPS) [crafted]; Dawnspire Cord (12466, -1.58 DPS) [dungeon] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 21.3 spell_power points (5.04 DPS) | yes | Red Mageweave Pants (10009, -0.36 DPS) [crafted]; Wizardweave Leggings (14132, -0.54 DPS) [crafted]; Big Voodoo Pants (8202, -2.20 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.68 DPS) | yes | Skycaller's Leather Boots (252471, -0.89 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -1.23 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.81 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.08 DPS) | yes | Snake Hoop (6750, -1.09 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.18 DPS) [quest]; Reedknot Ring (9622, -1.42 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (2.93 DPS) | yes | Snake Hoop (6750, -0.95 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.04 DPS) [quest]; Reedknot Ring (9622, -1.28 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (158.1 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (158.1 DPS) | yes | Mark of the Chosen (17774, -1.84 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (158.1 DPS) | yes | Spellshifter Rod (9527, -0.18 DPS) [quest]; Glowing Brightwood Staff (812, -0.45 DPS) [world_drop]; Hammer of the Northern Wind (810, -27.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Knight's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave

No-known-source sample (15 of 693, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 5232221115100151-00000000000000000000-5533030000000000)

Set DPS (verified): 231.9. Weights run: 2.7s. Verify run: 5.2s. 1571 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.531 ± 0.017, crit=0.130 ± 0.005 per rating point (14 rating = 1%, 1.825 per %), hit=0.397 ± 0.027 per rating point (10 rating = 1%, 3.970 per %), spell_haste=0.999 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 34.5 spell_power points (9.58 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.60 DPS) [vendor]; Soothsayer's Headdress (17740, -13.02 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.3 spell_power points (4.80 DPS) | yes | Glowing Eye of Mordresh (10769, -1.68 DPS) [dungeon]; Mindburst Medallion (11196, -2.25 DPS) [quest]; Triune Amulet (7722, -2.58 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (231.9 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Ironfeather Shoulders (15067, -8.21 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.2 spell_power points (4.76 DPS) | yes | Big Voodoo Cloak (8216, -2.05 DPS) [crafted]; Nightfall Drape (12465, -2.27 DPS) [dungeon]; Runecloth Cloak (13860, -3.08 DPS, sim-verified) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (231.9 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -7.23 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (231.9 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -11.19 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (231.9 DPS) | yes | Knight-Lieutenant's Dragonhide Gloves (227193, +0.00 DPS) [pvp]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Gloves of the Greatfather (17721, -8.78 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (231.9 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -12.23 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.4 spell_power points (13.69 DPS) | yes | Marshal's Dragonhide Pants (231697, -3.64 DPS) [vendor]; Luminary Kilt (11823, -4.70 DPS, sim-verified) [dungeon]; Knight-Captain's Dragonhide Pants (227191, -4.86 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.4 spell_power points (7.31 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.02 DPS) [vendor]; Earthen Silk Slippers (254013, -0.66 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.7 spell_power points (6.02 DPS) | yes | Philanthropist's Ring (281635, -2.51 DPS) [quest]; Snake Hoop (6750, -3.62 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.80 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.60 DPS) | yes | Philanthropist's Ring (281635, -0.10 DPS) [quest]; Snake Hoop (6750, -1.21 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.39 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (231.9 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (231.9 DPS) | yes | Mark of the Chosen (17774, -1.87 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (231.9 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, -0.01 DPS) [pvp]; Kindling Stave (11750, -8.46 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (5.82 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -1.87 DPS) [crafted] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-verified (231.9 DPS) | yes | Idol of the Moon (23197, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Band of the Unicorn; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 522.0. Weights run: 2.0s. Verify run: 4.9s. 1571 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.178 ± 0.011, crit=0.271 ± 0.011 per rating point (14 rating = 1%, 3.792 per %), hit=0.502 ± 0.030 per rating point (10 rating = 1%, 5.022 per %), spell_haste=2.953 ± 0.240, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.329 ± 0.000, arcane_power=0.671 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 28.2 spell_power points (16.41 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.12 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -16.86 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.8 spell_power points (8.02 DPS) | yes | Glowing Eye of Mordresh (10769, -3.12 DPS) [dungeon]; Triune Amulet (7722, -3.36 DPS) [dungeon]; Mindburst Medallion (11196, -3.90 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (522.0 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Argent Shoulders (19059, -8.65 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.1 spell_power points (8.77 DPS) | yes | Runecloth Cloak (13860, -2.52 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -3.53 DPS) [dungeon]; Icy Cloak (4327, -4.69 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (522.0 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -8.66 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.8 spell_power points (7.43 DPS) | yes | Runic Leather Bracers (15092, -0.58 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.89 DPS) [crafted]; Arcane Runed Bracers (4744, -2.20 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (13.96 DPS) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Knight-Lieutenant's Dragonhide Gloves (227193, -2.46 DPS) [pvp]; Earth Warder's Gloves (21318, -6.48 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 18.1 spell_power points (10.55 DPS) | yes | Satyrmane Sash (17755, -1.37 DPS) [dungeon]; Ghostweave Cord (254073, -2.40 DPS) [crafted]; Wisdom of the Timbermaw (19047, -2.70 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 45.0 spell_power points (26.18 DPS) | yes | Marshal's Dragonhide Pants (231697, -6.81 DPS) [vendor]; Knight-Captain's Dragonhide Pants (227191, -9.99 DPS) [vendor]; Ghostweave Pants (14144, -11.81 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.96 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.30 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.71 DPS) [vendor]; Omnicast Boots (11822, -1.09 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (11.20 DPS) | yes | Philanthropist's Ring (281635, -4.86 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.54 DPS) [quest]; Reedknot Ring (9622, -7.12 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.56 DPS) | yes | Philanthropist's Ring (281635, -1.23 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.91 DPS) [quest]; Reedknot Ring (9622, -3.49 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (522.0 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (522.0 DPS) | yes | Mark of the Chosen (17774, -7.29 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (522.0 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -17.56 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (12.22 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -2.04 DPS) [pvp]; Orb of the Forgotten Seer (7685, -4.25 DPS, sim-verified) [dungeon] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-verified (522.0 DPS) | yes | Idol of the Dream (220606, +0.00 DPS) [dungeon] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Idol of the Moon

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 33.1. Weights run: 2.5s. Verify run: 1.3s. 199 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.735 ± 0.007, crit=0.032 ± 0.001 per rating point (14 rating = 1%, 0.444 per %), hit=0.095 ± 0.007 per rating point (10 rating = 1%, 0.952 per %), spell_haste=not significant (-0.033 ± 0.098), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.025 ± 0.000, arcane_power=0.975 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.02 DPS) [crafted]; Totemic Leather Hood (252448, -0.57 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.6 spell_power points (1.79 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.73 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.17 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.65 DPS) | yes | Heavy Woolen Cloak (4311, -0.03 DPS) [crafted]; Black Whelp Cloak (7283, -0.19 DPS) [crafted]; Sanguine Cape (14376, -0.19 DPS) [world_drop] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.07 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.43 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.7 spell_power points (0.57 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.10 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.9 spell_power points (1.22 DPS) | yes | Serpent Gloves (5970, -0.14 DPS) [dungeon]; Pristine Gloves (253913, -0.27 DPS) [crafted]; Gnoll Casting Gloves (892, -0.30 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.15 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.79 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.15 DPS) [crafted]; Totemic Leather Pants (252446, -0.23 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.79 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (1.53 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.35 DPS) [crafted]; Totemic Leather Boots (252442, -0.61 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.4 spell_power points (0.68 DPS) | yes | Volcanic Rock Ring (12053, -0.34 DPS) [world_drop]; The 1 Ring (8350, -0.57 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 3.7 spell_power points (0.57 DPS) | yes | Volcanic Rock Ring (12053, -0.36 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.45 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.23 DPS) | yes | Twisted Chanter's Staff (890, -0.10 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.36 DPS, sim-verified) [quest]; Lesser Staff of the Spire (1300, -0.55 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 66.8. Weights run: 2.7s. Verify run: 1.4s. 387 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.522 ± 0.008, crit=0.049 ± 0.002 per rating point (14 rating = 1%, 0.691 per %), hit=0.153 ± 0.009 per rating point (10 rating = 1%, 1.534 per %), spell_haste=not significant (-0.128 ± 0.117), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.322 ± 0.001, arcane_power=0.678 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.51 DPS) | yes | Enchanter's Cowl (4322, -0.16 DPS) [crafted]; Holy Shroud (2721, -0.21 DPS) [world_drop]; Silk Headband (7050, -0.63 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.1 spell_power points (0.66 DPS) | yes | Crystal Starfire Medallion (5003, -0.22 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.22 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.7 spell_power points (2.87 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS) [dungeon]; Fairywing Mantle (9536, -0.63 DPS) [quest]; Magician's Mantle (12998, -0.84 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.05 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.21 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.30 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.21 DPS) [dungeon]; Guardian Armor (4256, -0.22 DPS) [crafted]; Green Silk Armor (7065, -1.23 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.88 DPS) | yes | Nightsky Wristbands (6407, -1.23 DPS) [world_drop]; Technician's Bracers (270042, -1.23 DPS) [quest]; Glowing Magical Bracelets (13106, -1.83 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.0 spell_power points (2.10 DPS) | yes | Jutebraid Gloves (10654, -0.30 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.62 DPS) [crafted]; Serpent Gloves (5970, -0.64 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.1 spell_power points (2.75 DPS) | yes | Warsong Sash (16975, -0.45 DPS) [quest]; Belt of Arugal (6392, -0.54 DPS) [dungeon]; Moss Cinch (6911, -1.29 DPS, sim-verified) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.20 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.24 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.67 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.7 spell_power points (2.23 DPS) | yes | Spidersilk Boots (4320, -0.33 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.43 DPS) [crafted]; Acidic Walkers (9454, -2.04 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.6 spell_power points (1.80 DPS) | yes | Solid Gold Band (279845, -0.97 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.15 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.17 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.26 DPS) | yes | Solid Gold Band (279845, -0.42 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.63 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, -5.86 DPS) [dungeon]; Manual Crowd Pummeler (9449, -11.40 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 5232221115100141-00000000000000000000-0000000000000000)

Set DPS (verified): 98.9. Weights run: 2.8s. Verify run: 1.5s. 521 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.344 ± 0.009, crit=0.070 ± 0.002 per rating point (14 rating = 1%, 0.975 per %), hit=0.198 ± 0.013 per rating point (10 rating = 1%, 1.980 per %), spell_haste=not significant (0.415 ± 0.115), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (5.28 DPS) | yes | Augural Shroud (2620, -1.30 DPS, sim-verified) [world]; Papal Fez (9431, -1.80 DPS) [dungeon]; Big Voodoo Mask (8201, -1.81 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 9.7 spell_power points (2.45 DPS) | yes | Triune Amulet (7722, -0.44 DPS) [dungeon]; Necklace of Calisea (1714, -1.85 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.93 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.1 spell_power points (3.04 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Inquisitor's Shawl (19507, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.42 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.7 spell_power points (1.94 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.18 DPS) [crafted]; Hillman's Cloak (3719, -0.68 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.1 spell_power points (6.05 DPS) | yes | Dreamweave Vest (10021, -0.75 DPS) [crafted]; Elemental Raiment (9434, -0.77 DPS) [world_drop]; Robe of Power (7054, -1.49 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.26 DPS) | yes | Condor Bracers (15864, -0.50 DPS) [quest]; Radiant Silver Bracers (4545, -0.57 DPS) [quest]; Guardian Leather Bracers (4260, -1.39 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (6.04 DPS) | yes | Black Mageweave Gloves (10003, -2.26 DPS) [crafted]; Red Mageweave Gloves (10018, -2.41 DPS) [crafted]; Dreamweave Gloves (10019, -2.80 DPS, sim-verified) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.27 DPS) | yes | Deathmage Sash (10771, -0.21 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.24 DPS) [crafted]; Moss Cinch (6911, -0.25 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.1 spell_power points (4.56 DPS) | yes | Crimson Silk Pantaloons (7062, -1.42 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.53 DPS) [crafted]; Dark Ritual Leggings (270031, -3.10 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (6.04 DPS) | yes | Skycaller's Leather Shoes (252532, -2.11 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -3.42 DPS) [crafted]; Gilded Slippers (254001, -3.67 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.7 spell_power points (2.95 DPS) | yes | Reedknot Ring (9622, -1.19 DPS) [quest]; Ring of Power Regulation (273030, -1.44 DPS) [dungeon]; Solid Gold Band (279845, -1.94 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.7 spell_power points (1.94 DPS) | yes | Reedknot Ring (9622, -0.18 DPS) [quest]; Ring of Power Regulation (273030, -0.43 DPS) [dungeon]; Solid Gold Band (279845, -0.94 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (98.9 DPS) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.17 DPS) [dungeon]; Spine of the Devourer (273040, -2.70 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.52 DPS) | yes | Orb of Mystic Insight (249394, -1.24 DPS) [crafted]; Omega Orb (7749, -1.26 DPS) [quest]; Beacon of Hope (9393, -1.51 DPS) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer; ranged: Talons of Wrath

No-known-source sample (15 of 521, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 5232221115100151-00000000000000000000-5400000000000000)

Set DPS (verified): 157.7. Weights run: 3.0s. Verify run: 1.8s. 691 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.477 ± 0.013, crit=0.104 ± 0.004 per rating point (14 rating = 1%, 1.460 per %), hit=0.303 ± 0.020 per rating point (10 rating = 1%, 3.032 per %), spell_haste=not significant (0.348 ± 0.161), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.508 ± 0.001, arcane_power=0.492 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.2 spell_power points (7.14 DPS) | yes | Red Mageweave Headband (10033, -0.38 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.75 DPS) [crafted]; Dreamweave Circlet (10041, -1.04 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.8 spell_power points (2.56 DPS) | yes | Mindburst Medallion (11196, -0.46 DPS) [quest]; Triune Amulet (7722, -0.67 DPS) [dungeon]; Horizon Choker (13085, -0.98 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.5 spell_power points (5.81 DPS) | yes | Blood Guard's Crackling Leather Spaulders (220871, -1.12 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.72 DPS) [crafted]; Rotgrip Mantle (17732, -2.01 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.9 spell_power points (3.99 DPS) | yes | Deep Woodlands Cloak (19121, -0.13 DPS) [quest]; Runecloth Cloak (13860, -0.96 DPS) [crafted]; Big Voodoo Cloak (8216, -1.79 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.5 spell_power points (6.76 DPS) | yes | Robe of the Magi (1716, -0.87 DPS) [world_drop]; Feathered Breastplate (8349, -1.13 DPS) [crafted]; Dreamweave Vest (10021, -1.48 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.8 spell_power points (3.50 DPS) | yes | Skycaller's Leather Bracers (252542, -0.34 DPS) [crafted]; Nethergeld Cuffs (254061, -1.05 DPS) [crafted]; Mender's Leather Bracers (252543, -1.29 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (5.68 DPS) | yes | Dreamweave Gloves (10019, -0.97 DPS) [crafted]; First Sergeant's Crackling Leather Gauntlets (220867, -1.13 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -1.51 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.7 spell_power points (5.15 DPS) | yes | Satyrmane Sash (17755, -0.70 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.42 DPS) [crafted]; Dawnspire Cord (12466, -1.58 DPS) [dungeon] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | PvP rank 10 · Stone Guard · Horde [vendor] | 21.3 spell_power points (5.04 DPS) | yes | Red Mageweave Pants (10009, -0.36 DPS) [crafted]; Wizardweave Leggings (14132, -0.54 DPS) [crafted]; Big Voodoo Pants (8202, -2.78 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.68 DPS) | yes | First Sergeant's Crackling Leather Boots (220863, -1.23 DPS) [vendor]; Skycaller's Leather Boots (252471, -1.74 DPS, sim-verified) [crafted]; Skycaller's Leather Shoes (252532, -1.81 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.08 DPS) | yes | Snake Hoop (6750, -1.09 DPS) [quest]; Reedknot Ring (9622, -1.42 DPS) [quest]; Ring of Power Regulation (273030, -1.66 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (2.93 DPS) | yes | Snake Hoop (6750, -0.95 DPS) [quest]; Reedknot Ring (9622, -1.28 DPS) [quest]; Ring of Power Regulation (273030, -1.51 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (157.7 DPS) | yes | Rune of the Guard Captain (19120, -0.14 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (157.7 DPS) | yes | Rune of the Guard Captain (19120, -2.19 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (157.7 DPS) | yes | Spellshifter Rod (9527, -0.18 DPS) [quest]; Glowing Brightwood Staff (812, -0.45 DPS) [world_drop]; Hammer of the Northern Wind (810, -26.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 5232221115100151-00000000000000000000-5533030000000000)

Set DPS (verified): 232.4. Weights run: 2.7s. Verify run: 5.2s. 1571 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.531 ± 0.017, crit=0.130 ± 0.005 per rating point (14 rating = 1%, 1.825 per %), hit=0.397 ± 0.027 per rating point (10 rating = 1%, 3.970 per %), spell_haste=0.999 ± 0.234, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.486 ± 0.001, arcane_power=0.514 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 34.5 spell_power points (9.58 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.60 DPS) [pvp]; Soothsayer's Headdress (17740, -13.21 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.3 spell_power points (4.80 DPS) | yes | Glowing Eye of Mordresh (10769, -1.68 DPS) [dungeon]; Mindburst Medallion (11196, -2.25 DPS) [quest]; Triune Amulet (7722, -2.58 DPS) [dungeon] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (232.4 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -8.57 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.2 spell_power points (4.76 DPS) | yes | Deep Woodlands Cloak (19121, -0.11 DPS) [quest]; Runecloth Cloak (13860, -1.09 DPS) [crafted]; Big Voodoo Cloak (8216, -2.05 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (232.4 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -7.14 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (232.4 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -11.81 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (232.4 DPS) | yes | Blood Guard's Dragonhide Gloves (227187, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Gloves of the Greatfather (17721, -10.28 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (232.4 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Wisdom of the Timbermaw (19047, +0.00 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -12.58 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.4 spell_power points (13.69 DPS) | yes | General's Dragonhide Pants (231680, -3.64 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -4.86 DPS) [vendor]; Luminary Kilt (11823, -5.92 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.4 spell_power points (7.31 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.02 DPS) [pvp]; Earthen Silk Slippers (254013, -0.66 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 21.7 spell_power points (6.02 DPS) | yes | Philanthropist's Ring (281635, -2.51 DPS) [quest]; Snake Hoop (6750, -3.62 DPS) [quest]; Reedknot Ring (9622, -4.08 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.60 DPS) | yes | Philanthropist's Ring (281635, -0.10 DPS) [quest]; Snake Hoop (6750, -1.21 DPS) [quest]; Reedknot Ring (9622, -1.66 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (232.4 DPS) | yes | Rune of the Guard Captain (19120, -0.22 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (232.4 DPS) | yes | Rune of the Guard Captain (19120, -3.30 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (232.4 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, -0.01 DPS) [pvp]; Kindling Stave (11750, -8.46 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (5.82 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -2.74 DPS, sim-verified) [crafted] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-verified (232.4 DPS) | yes | Idol of the Moon (23197, -2.16 DPS, sim-verified) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Band of the Unicorn; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5050030000000000)

Set DPS (verified): 515.2. Weights run: 2.0s. Verify run: 3.8s. 1571 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.178 ± 0.011, crit=0.271 ± 0.011 per rating point (14 rating = 1%, 3.792 per %), hit=0.502 ± 0.030 per rating point (10 rating = 1%, 5.022 per %), spell_haste=2.953 ± 0.240, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.329 ± 0.000, arcane_power=0.671 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 28.2 spell_power points (16.41 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.12 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -15.81 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.8 spell_power points (8.02 DPS) | yes | Glowing Eye of Mordresh (10769, -2.53 DPS, sim-verified) [dungeon]; Triune Amulet (7722, -3.36 DPS) [dungeon]; Mindburst Medallion (11196, -3.90 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (14.55 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Mantle of the Timbermaw (19050, -1.11 DPS) [crafted]; Champion's Dragonhide Spaulders (227184, -2.48 DPS) [pvp] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.1 spell_power points (8.77 DPS) | yes | Deep Woodlands Cloak (19121, -0.85 DPS) [quest]; Runecloth Cloak (13860, -2.70 DPS) [crafted]; Nightfall Drape (12465, -3.53 DPS) [dungeon] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | 27.4 spell_power points (15.95 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -3.10 DPS, sim-verified) [crafted] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.8 spell_power points (7.43 DPS) | yes | Runic Leather Bracers (15092, -0.58 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.89 DPS) [crafted]; Spidertank Oilrag (9448, -2.20 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (13.96 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Gloves (227187, -2.46 DPS) [vendor]; Earth Warder's Gloves (21318, -6.53 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 18.1 spell_power points (10.55 DPS) | yes | Ghostweave Cord (254073, -2.40 DPS) [crafted]; Wisdom of the Timbermaw (19047, -2.70 DPS) [crafted]; Satyrmane Sash (17755, -3.34 DPS, sim-verified) [dungeon] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 45.0 spell_power points (26.18 DPS) | yes | General's Dragonhide Pants (231680, -6.81 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -9.99 DPS) [vendor]; Ghostweave Pants (14144, -12.21 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.96 DPS) | yes | General's Dragonhide Boots (231682, -0.30 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.71 DPS) [pvp]; Omnicast Boots (11822, -1.09 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.2 spell_power points (11.20 DPS) | yes | Philanthropist's Ring (281635, -4.86 DPS) [quest]; Reedknot Ring (9622, -7.12 DPS) [quest]; Snake Hoop (6750, -7.19 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.56 DPS) | yes | Philanthropist's Ring (281635, -1.23 DPS) [quest]; Reedknot Ring (9622, -3.49 DPS) [quest]; Snake Hoop (6750, -3.56 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (515.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (515.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (515.2 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -17.56 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (12.22 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -2.04 DPS) [pvp]; Orb of the Forgotten Seer (7685, -4.28 DPS, sim-verified) [dungeon] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-verified (515.2 DPS) | yes | Talons of Wrath (249441, +0.00 DPS) [crafted] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Fernpulse Jerkin; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Idol of the Moon

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

