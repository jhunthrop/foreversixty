# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 30.5. Weights run: 2.0s. Verify run: 1.1s. 196 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.782 ± 0.005, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.438 per %), hit=0.102 ± 0.005 per rating point (10 rating = 1%, 1.018 per %), spell_haste=-0.851 ± 0.032, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.000, arcane_power=0.967 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Wisdom's Leather Hood (252507, -0.40 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS) [crafted]; Trapper's Leather Hood (252505, -1.51 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 spell_power points (1.60 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.05 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.07 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.3 spell_power points (0.58 DPS) | yes | Heavy Woolen Cloak (4311, -0.05 DPS) [crafted]; Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.18 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.40 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.10 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.21 DPS) [dungeon]; Owl Bracers (4796, -0.88 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (1.08 DPS) | yes | Windfelt Gloves (5630, -0.24 DPS) [quest]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.53 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.12 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.73 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.7 spell_power points (1.95 DPS) | yes | Wisdom's Leather Pants (252503, -0.40 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.53 DPS) [crafted]; Dreamer's Leggings (270016, -0.74 DPS, sim-verified) [quest] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.1 spell_power points (1.34 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.29 DPS) [crafted]; Black Whelp Slippers (252424, -0.53 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.6 spell_power points (0.87 DPS) | yes | Volcanic Rock Ring (12053, -0.56 DPS) [world_drop]; The 1 Ring (8350, -0.77 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.7 spell_power points (0.62 DPS) | yes | Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; The 1 Ring (8350, -0.52 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.06 DPS) | yes | Lesser Staff of the Spire (1300, -0.44 DPS) [world]; Staff of Westfall (2042, -0.54 DPS) [quest]; Twisted Chanter's Staff (890, -1.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 60.7. Weights run: 2.2s. Verify run: 1.3s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.548 ± 0.007, crit=0.035 ± 0.002 per rating point (14 rating = 1%, 0.497 per %), hit=0.123 ± 0.005 per rating point (10 rating = 1%, 1.225 per %), spell_haste=0.320 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.26 DPS) | yes | Holy Shroud (2721, -0.19 DPS) [world_drop]; Enchanter's Cowl (4322, -0.51 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.56 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.62 DPS) | yes | Crystal Starfire Medallion (5003, -0.21 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.21 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.9 spell_power points (2.62 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS) [dungeon]; Fairywing Mantle (9536, -0.56 DPS) [quest]; Magician's Mantle (12998, -0.75 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.94 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Prelacy Cape (7004, -0.19 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.19 DPS) [dungeon]; Guardian Armor (4256, -0.21 DPS) [crafted]; Green Silk Armor (7065, -1.06 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.69 DPS) | yes | Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Technician's Bracers (270042, -1.08 DPS) [quest]; Glowing Magical Bracelets (13106, -1.81 DPS, sim-verified) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.2 spell_power points (1.35 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Shilly Mitts (9609, -0.04 DPS) [quest]; Gloves of Insight (9698, -0.04 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (2.50 DPS) | yes | Belt of Arugal (6392, -0.50 DPS) [dungeon]; Crimson Silk Belt (7055, -0.65 DPS) [crafted]; Moss Cinch (6911, -1.07 DPS, sim-verified) [dungeon] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (2.64 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Guardian Pants (5962, -0.30 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.38 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.04 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -1.89 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.7 spell_power points (1.65 DPS) | yes | Lavishly Jeweled Ring (1156, -1.03 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.08 DPS) [dungeon]; Minor Channeling Ring (1449, -2.16 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.56 DPS) [dungeon]; Minor Channeling Ring (1449, -1.09 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.01 DPS) [quest]; Manual Crowd Pummeler (9449, -10.96 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 82.2. Weights run: 2.5s. Verify run: 1.4s. 516 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.398 ± 0.005, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.005 per %), hit=0.170 ± 0.007 per rating point (10 rating = 1%, 1.702 per %), spell_haste=not significant (-0.271 ± 0.073), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.29 DPS) | yes | Augural Shroud (2620, -1.23 DPS) [world]; Papal Fez (9431, -1.27 DPS) [dungeon]; Big Voodoo Mask (8201, -1.31 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.2 spell_power points (2.08 DPS) | yes | Necklace of Calisea (1714, -1.51 DPS) [dungeon]; Triune Amulet (7722, -1.56 DPS, sim-verified) [dungeon]; Archmagister's Faceted Pendant (273049, -1.59 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (2.57 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.0 spell_power points (1.63 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.61 DPS) [quest] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.12 DPS) [world_drop]; Robe of Power (7054, -0.57 DPS) [crafted]; Robe of the Magi (1716, -1.06 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Leather Bracers (4260, -0.13 DPS) [crafted]; Condor Bracers (15864, -0.41 DPS) [quest]; Arcane Runed Bracers (4744, -0.97 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.91 DPS) | yes | Dreamweave Gloves (10019, -1.63 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.84 DPS) [crafted]; Red Mageweave Gloves (10018, -1.84 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (2.66 DPS) | yes | Deathmage Sash (10771, -0.01 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.13 DPS) [crafted]; Moss Cinch (6911, -0.20 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.8 spell_power points (3.84 DPS) | yes | Crimson Silk Pantaloons (7062, -1.14 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.31 DPS) [crafted]; Dark Ritual Leggings (270031, -2.29 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.91 DPS) | yes | Skycaller's Leather Shoes (252532, -1.79 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.70 DPS) [crafted]; Gilded Slippers (254001, -2.91 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (2.45 DPS) | yes | Snake Hoop (6750, -0.82 DPS) [quest]; Reedknot Ring (9622, -1.02 DPS) [quest]; Ring of Power Regulation (273030, -1.22 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (1.64 DPS) | yes | Snake Hoop (6750, -0.00 DPS) [quest]; Reedknot Ring (9622, -0.20 DPS) [quest]; Ring of Power Regulation (273030, -0.41 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.16 DPS) [dungeon]; Spine of the Devourer (273040, -2.29 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Mystic Insight (249394, -0.12 DPS) [crafted]; Beacon of Hope (9393, -0.40 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.92 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of Lorica

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 134.3. Weights run: 2.6s. Verify run: 1.5s. 686 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.468 ± 0.009, crit=0.116 ± 0.006 per rating point (14 rating = 1%, 1.626 per %), hit=0.277 ± 0.013 per rating point (10 rating = 1%, 2.768 per %), spell_haste=not significant (0.045 ± 0.138), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.377 ± 0.001, arcane_power=0.623 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.0 spell_power points (5.69 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -0.82 DPS) [crafted]; Red Mageweave Headband (10033, -1.30 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.7 spell_power points (2.04 DPS) | yes | Mindburst Medallion (11196, -0.37 DPS) [quest]; Triune Amulet (7722, -0.52 DPS) [dungeon]; Horizon Choker (13085, -0.79 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.4 spell_power points (4.62 DPS) | yes | Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.93 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.35 DPS) [crafted]; Rotgrip Mantle (17732, -2.60 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (3.18 DPS) | yes | Big Voodoo Cloak (8216, -1.44 DPS) [crafted]; Nightfall Drape (12465, -1.48 DPS) [dungeon]; Runecloth Cloak (13860, -1.93 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.4 spell_power points (5.37 DPS) | yes | Feathered Breastplate (8349, -0.89 DPS) [crafted]; Dreamweave Vest (10021, -1.17 DPS) [crafted]; Robe of the Magi (1716, -3.45 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.7 spell_power points (2.78 DPS) | yes | Skycaller's Leather Bracers (252542, -0.27 DPS) [crafted]; Nethergeld Cuffs (254061, -0.83 DPS) [crafted]; Mender's Leather Bracers (252543, -1.02 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.55 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.72 DPS) [crafted]; Dreamweave Gloves (10019, -0.78 DPS) [crafted]; Sergeant Major's Crackling Leather Gauntlets (220866, -0.92 DPS) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.6 spell_power points (4.10 DPS) | yes | Satyrmane Sash (17755, -0.56 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.14 DPS) [crafted]; Dawnspire Cord (12466, -1.27 DPS) [dungeon] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 21.1 spell_power points (3.99 DPS) | yes | Red Mageweave Pants (10009, -0.28 DPS) [crafted]; Wizardweave Leggings (14132, -0.39 DPS) [crafted]; Big Voodoo Pants (8202, -1.58 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.55 DPS) | yes | Sergeant Major's Crackling Leather Boots (220862, -1.05 DPS) [vendor]; Skycaller's Leather Boots (252471, -1.17 DPS, sim-verified) [crafted]; Skycaller's Leather Shoes (252532, -1.46 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.46 DPS) | yes | Snake Hoop (6750, -0.88 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.95 DPS) [quest]; Reedknot Ring (9622, -1.14 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.3 spell_power points (2.34 DPS) | yes | Snake Hoop (6750, -0.76 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.82 DPS) [quest]; Reedknot Ring (9622, -1.01 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (134.3 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (134.3 DPS) | yes | Mark of the Chosen (17774, -1.23 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (134.3 DPS) | yes | Spellshifter Rod (9527, -0.16 DPS) [quest]; Glowing Brightwood Staff (812, -0.39 DPS) [world_drop]; Hammer of the Northern Wind (810, -24.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Knight's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 174.2. Weights run: 2.5s. Verify run: 4.7s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.666 ± 0.014, crit=0.145 ± 0.008 per rating point (14 rating = 1%, 2.037 per %), hit=0.350 ± 0.017 per rating point (10 rating = 1%, 3.500 per %), spell_haste=not significant (0.325 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.375 ± 0.001, arcane_power=0.625 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 37.0 spell_power points (7.45 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.40 DPS) [vendor]; Soothsayer's Headdress (17740, -10.09 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.7 spell_power points (3.76 DPS) | yes | Glowing Eye of Mordresh (10769, -1.28 DPS) [dungeon]; Mindburst Medallion (11196, -1.74 DPS) [quest]; Horizon Choker (13085, -1.88 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (174.2 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Ironfeather Shoulders (15067, -4.86 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.0 spell_power points (3.62 DPS) | yes | Big Voodoo Cloak (8216, -1.41 DPS) [crafted]; Long Silken Cloak (4326, -1.75 DPS) [crafted]; Runecloth Cloak (13860, -2.54 DPS, sim-verified) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (174.2 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -6.10 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (174.2 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -6.56 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (174.2 DPS) | yes | Knight-Lieutenant's Dragonhide Gloves (227193, +0.00 DPS) [pvp]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Gloves of the Greatfather (17721, -7.02 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (174.2 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, -9.37 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.0 spell_power points (10.48 DPS) | yes | Marshal's Dragonhide Pants (231697, -2.75 DPS) [vendor]; Knight-Captain's Dragonhide Pants (227191, -3.56 DPS) [vendor]; Luminary Kilt (11823, -4.07 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 28.0 spell_power points (5.64 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.07 DPS) [vendor]; Earthen Silk Slippers (254013, -1.55 DPS, sim-verified) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 22.7 spell_power points (4.56 DPS) | yes | Band of the Unicorn (7553, -1.95 DPS) [world_drop]; Snake Hoop (6750, -2.68 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.95 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (2.68 DPS) | yes | Band of the Unicorn (7553, -0.07 DPS) [world_drop]; Snake Hoop (6750, -0.81 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.07 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (174.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (174.2 DPS) | yes | Mark of the Chosen (17774, -3.70 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (174.2 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -5.91 DPS) [quest] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (4.23 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -1.14 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 4132220115501051-05000000000000000000-5053000000000000)

Set DPS (verified): 419.6. Weights run: 1.5s. Verify run: 4.2s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.233 ± 0.011, crit=0.284 ± 0.011 per rating point (14 rating = 1%, 3.969 per %), hit=0.598 ± 0.022 per rating point (10 rating = 1%, 5.982 per %), spell_haste=3.830 ± 0.105, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.194 ± 0.000, arcane_power=0.806 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 29.2 spell_power points (13.61 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.01 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -11.78 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.3 spell_power points (6.68 DPS) | yes | Glowing Eye of Mordresh (10769, -2.55 DPS) [dungeon]; Triune Amulet (7722, -2.95 DPS) [dungeon]; Mindburst Medallion (11196, -3.23 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (419.6 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Argent Shoulders (19059, -5.97 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.4 spell_power points (7.18 DPS) | yes | Runecloth Cloak (13860, -2.11 DPS) [crafted]; Nightfall Drape (12465, -2.98 DPS) [dungeon]; Long Silken Cloak (4326, -3.84 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (419.6 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -7.74 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.3 spell_power points (6.21 DPS) | yes | Runic Leather Bracers (15092, -0.47 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.79 DPS) [crafted]; Arcane Runed Bracers (4744, -2.02 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (11.19 DPS) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Knight-Lieutenant's Dragonhide Gloves (227193, -1.71 DPS) [pvp]; Earth Warder's Gloves (21318, -5.24 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 18.8 spell_power points (8.76 DPS) | yes | Satyrmane Sash (17755, -1.15 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.11 DPS) [crafted]; Ghostweave Cord (254073, -2.23 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 46.2 spell_power points (21.52 DPS) | yes | Marshal's Dragonhide Pants (231697, -5.56 DPS) [vendor]; Knight-Captain's Dragonhide Pants (227191, -8.04 DPS) [vendor]; Luminary Kilt (11823, -11.09 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.19 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.31 DPS) [vendor]; Omnicast Boots (11822, -0.56 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.6 spell_power points (9.15 DPS) | yes | Philanthropist's Ring (281635, -3.95 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.42 DPS) [quest]; Snake Hoop (6750, -5.81 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.06 DPS) | yes | Philanthropist's Ring (281635, -0.86 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.33 DPS) [quest]; Snake Hoop (6750, -2.72 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (419.6 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (419.6 DPS) | yes | Mark of the Chosen (17774, -5.89 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (419.6 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -14.10 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.79 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -1.28 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.44 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 29.2. Weights run: 2.0s. Verify run: 1.1s. 192 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.782 ± 0.005, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.438 per %), hit=0.102 ± 0.005 per rating point (10 rating = 1%, 1.018 per %), spell_haste=-0.851 ± 0.032, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.000, arcane_power=0.967 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Wisdom's Leather Hood (252507, -0.40 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS) [crafted]; Trapper's Leather Hood (252505, -1.59 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.0 spell_power points (1.60 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.01 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.07 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.3 spell_power points (0.58 DPS) | yes | Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.46 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.39 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.9 spell_power points (0.52 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.10 DPS) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (1.08 DPS) | yes | Pristine Gloves (253913, -0.24 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.27 DPS) [crafted]; Serpent Gloves (5970, -0.49 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.12 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.72 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.21 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.72 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.1 spell_power points (1.34 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.29 DPS) [crafted]; Black Whelp Slippers (252424, -0.53 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.7 spell_power points (0.62 DPS) | yes | Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; The 1 Ring (8350, -0.52 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 3.9 spell_power points (0.52 DPS) | yes | The 1 Ring (8350, -0.42 DPS) [world]; Volcanic Rock Ring (12053, -0.81 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.06 DPS) | yes | Twisted Chanter's Staff (890, -0.02 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.44 DPS) [world]; Gnarled Necromancer's Staff (251534, -1.13 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 60.8. Weights run: 2.2s. Verify run: 1.2s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.548 ± 0.007, crit=0.035 ± 0.002 per rating point (14 rating = 1%, 0.497 per %), hit=0.123 ± 0.005 per rating point (10 rating = 1%, 1.225 per %), spell_haste=0.320 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.26 DPS) | yes | Enchanter's Cowl (4322, -0.10 DPS) [crafted]; Holy Shroud (2721, -0.19 DPS) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.3 spell_power points (0.62 DPS) | yes | Kaleidoscope Chain (13084, -0.21 DPS) [world_drop]; Crystal Starfire Medallion (5003, -1.25 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.9 spell_power points (2.62 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS) [dungeon]; Fairywing Mantle (9536, -0.56 DPS) [quest]; Magician's Mantle (12998, -0.75 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.19 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.26 DPS) [crafted]; Hillman's Cloak (3719, -0.63 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.1 spell_power points (3.04 DPS) | yes | Beguiler Robes (7728, -0.10 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.29 DPS) [dungeon]; Guardian Armor (4256, -0.31 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.69 DPS) | yes | Nightsky Wristbands (6407, -1.08 DPS) [world_drop]; Technician's Bracers (270042, -1.08 DPS) [quest]; Glowing Magical Bracelets (13106, -1.80 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.1 spell_power points (1.90 DPS) | yes | Jutebraid Gloves (10654, -0.26 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.55 DPS) [crafted]; Serpent Gloves (5970, -0.58 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (2.50 DPS) | yes | Warsong Sash (16975, -0.43 DPS) [quest]; Belt of Arugal (6392, -0.50 DPS) [dungeon]; Moss Cinch (6911, -1.68 DPS, sim-verified) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.17 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.24 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.75 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.04 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -1.91 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.7 spell_power points (1.65 DPS) | yes | Solid Gold Band (279845, -0.89 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.03 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.08 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.13 DPS) | yes | Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.56 DPS) [dungeon]; Solid Gold Band (279845, -1.80 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, -5.27 DPS) [dungeon]; Manual Crowd Pummeler (9449, -10.90 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 81.2. Weights run: 2.5s. Verify run: 1.4s. 514 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.398 ± 0.005, crit=0.072 ± 0.003 per rating point (14 rating = 1%, 1.005 per %), hit=0.170 ± 0.007 per rating point (10 rating = 1%, 1.702 per %), spell_haste=not significant (-0.271 ± 0.073), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.29 DPS) | yes | Augural Shroud (2620, -1.23 DPS) [world]; Papal Fez (9431, -1.27 DPS) [dungeon]; Big Voodoo Mask (8201, -1.31 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.2 spell_power points (2.08 DPS) | yes | Triune Amulet (7722, -1.10 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.51 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.59 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (2.57 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.0 spell_power points (1.63 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Hillman's Cloak (3719, -0.61 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.12 DPS) [world_drop]; Robe of Power (7054, -0.57 DPS) [crafted]; Robe of the Magi (1716, -1.82 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.84 DPS) | yes | Radiant Silver Bracers (4545, -0.37 DPS) [quest]; Condor Bracers (15864, -0.41 DPS) [quest]; Guardian Leather Bracers (4260, -0.68 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.91 DPS) | yes | Dreamweave Gloves (10019, -1.14 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.84 DPS) [crafted]; Red Mageweave Gloves (10018, -1.84 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Leather Belt (252522, -0.12 DPS) [crafted]; Moss Cinch (6911, -0.20 DPS) [dungeon]; Star Belt (4329, -0.88 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.8 spell_power points (3.84 DPS) | yes | Crimson Silk Pantaloons (7062, -1.14 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.31 DPS) [crafted]; Dark Ritual Leggings (270031, -2.31 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.91 DPS) | yes | Skycaller's Leather Shoes (252532, -1.54 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.70 DPS) [crafted]; Gilded Slippers (254001, -2.91 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.0 spell_power points (2.45 DPS) | yes | Reedknot Ring (9622, -1.02 DPS) [quest]; Ring of Power Regulation (273030, -1.22 DPS) [dungeon]; Solid Gold Band (279845, -1.63 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 spell_power points (1.63 DPS) | yes | Reedknot Ring (9622, -0.20 DPS) [quest]; Ring of Power Regulation (273030, -0.41 DPS) [dungeon]; Solid Gold Band (279845, -0.82 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.16 DPS) [dungeon]; Spine of the Devourer (273040, -2.29 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.86 DPS) | yes | Orb of Mystic Insight (249394, -0.94 DPS) [crafted]; Omega Orb (7749, -1.02 DPS) [quest]; Beacon of Hope (9393, -1.23 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 132.8. Weights run: 2.6s. Verify run: 1.6s. 684 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.468 ± 0.009, crit=0.116 ± 0.006 per rating point (14 rating = 1%, 1.626 per %), hit=0.277 ± 0.013 per rating point (10 rating = 1%, 2.768 per %), spell_haste=not significant (0.045 ± 0.138), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.377 ± 0.001, arcane_power=0.623 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.0 spell_power points (5.69 DPS) | yes | Red Mageweave Headband (10033, -0.31 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -0.82 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.7 spell_power points (2.04 DPS) | yes | Mindburst Medallion (11196, -0.37 DPS) [quest]; Triune Amulet (7722, -0.52 DPS) [dungeon]; Horizon Choker (13085, -0.79 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.4 spell_power points (4.62 DPS) | yes | Blood Guard's Crackling Leather Spaulders (220871, -0.93 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.35 DPS) [crafted]; Rotgrip Mantle (17732, -2.57 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (3.18 DPS) | yes | Deep Woodlands Cloak (19121, -0.11 DPS) [quest]; Runecloth Cloak (13860, -0.77 DPS) [crafted]; Big Voodoo Cloak (8216, -1.44 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.4 spell_power points (5.37 DPS) | yes | Feathered Breastplate (8349, -0.89 DPS) [crafted]; Dreamweave Vest (10021, -1.17 DPS) [crafted]; Robe of the Magi (1716, -1.71 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.7 spell_power points (2.78 DPS) | yes | Skycaller's Leather Bracers (252542, -0.27 DPS) [crafted]; Nethergeld Cuffs (254061, -0.83 DPS) [crafted]; Mender's Leather Bracers (252543, -1.02 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.55 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.72 DPS) [crafted]; Dreamweave Gloves (10019, -0.78 DPS) [crafted]; First Sergeant's Crackling Leather Gauntlets (220867, -0.92 DPS) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.6 spell_power points (4.10 DPS) | yes | Satyrmane Sash (17755, -0.56 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.14 DPS) [crafted]; Dawnspire Cord (12466, -1.27 DPS) [dungeon] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | PvP rank 10 · Stone Guard · Horde [vendor] | 21.1 spell_power points (3.99 DPS) | yes | Red Mageweave Pants (10009, -0.28 DPS) [crafted]; Wizardweave Leggings (14132, -0.39 DPS) [crafted]; Big Voodoo Pants (8202, -1.38 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.55 DPS) | yes | Skycaller's Leather Boots (252471, -0.73 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -1.05 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.46 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.46 DPS) | yes | Snake Hoop (6750, -0.88 DPS) [quest]; Reedknot Ring (9622, -1.14 DPS) [quest]; Ring of Power Regulation (273030, -1.33 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.3 spell_power points (2.34 DPS) | yes | Snake Hoop (6750, -0.76 DPS) [quest]; Reedknot Ring (9622, -1.01 DPS) [quest]; Ring of Power Regulation (273030, -1.20 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (132.8 DPS) | yes | Rune of the Guard Captain (19120, -0.10 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (132.8 DPS) | yes | Rune of the Guard Captain (19120, -1.53 DPS, sim-verified) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (132.8 DPS) | yes | Spellshifter Rod (9527, -0.16 DPS) [quest]; Glowing Brightwood Staff (812, -0.39 DPS) [world_drop]; Hammer of the Northern Wind (810, -23.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 173.5. Weights run: 2.5s. Verify run: 4.7s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.666 ± 0.014, crit=0.145 ± 0.008 per rating point (14 rating = 1%, 2.037 per %), hit=0.350 ± 0.017 per rating point (10 rating = 1%, 3.500 per %), spell_haste=not significant (0.325 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.375 ± 0.001, arcane_power=0.625 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 37.0 spell_power points (7.45 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.40 DPS) [pvp]; Soothsayer's Headdress (17740, -10.77 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.7 spell_power points (3.76 DPS) | yes | Glowing Eye of Mordresh (10769, -1.28 DPS) [dungeon]; Mindburst Medallion (11196, -1.74 DPS) [quest]; Horizon Choker (13085, -1.88 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (173.5 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -4.27 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.0 spell_power points (3.62 DPS) | yes | Deep Woodlands Cloak (19121, -0.00 DPS) [quest]; Runecloth Cloak (13860, -0.74 DPS) [crafted]; Big Voodoo Cloak (8216, -1.41 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (173.5 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -5.97 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (173.5 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -6.43 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (173.5 DPS) | yes | Blood Guard's Dragonhide Gloves (227187, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Gloves of the Greatfather (17721, -7.45 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (173.5 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, -9.17 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.0 spell_power points (10.48 DPS) | yes | General's Dragonhide Pants (231680, -2.75 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -3.56 DPS) [vendor]; Luminary Kilt (11823, -4.02 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 28.0 spell_power points (5.64 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, +0.00 DPS) [crafted]; Blood Guard's Dragonhide Boots (227188, -0.07 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 22.7 spell_power points (4.56 DPS) | yes | Band of the Unicorn (7553, -1.95 DPS) [world_drop]; Snake Hoop (6750, -2.68 DPS) [quest]; Glowing Crystal Ring (18402, -2.95 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (2.68 DPS) | yes | Band of the Unicorn (7553, -0.07 DPS) [world_drop]; Snake Hoop (6750, -0.81 DPS) [quest]; Glowing Crystal Ring (18402, -1.08 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (173.5 DPS) | yes | Rune of the Guard Captain (19120, -0.71 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (173.5 DPS) | yes | Rune of the Guard Captain (19120, -0.14 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (173.5 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -5.91 DPS) [quest] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (4.23 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -1.14 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 4132220115501051-05000000000000000000-5053000000000000)

Set DPS (verified): 416.1. Weights run: 1.5s. Verify run: 4.2s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.233 ± 0.011, crit=0.284 ± 0.011 per rating point (14 rating = 1%, 3.969 per %), hit=0.598 ± 0.022 per rating point (10 rating = 1%, 5.982 per %), spell_haste=3.830 ± 0.105, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.194 ± 0.000, arcane_power=0.806 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 29.2 spell_power points (13.61 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.01 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -12.62 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 14.3 spell_power points (6.68 DPS) | yes | Glowing Eye of Mordresh (10769, -2.55 DPS) [dungeon]; Triune Amulet (7722, -2.95 DPS) [dungeon]; Mindburst Medallion (11196, -3.23 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (416.1 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Argent Shoulders (19059, -6.69 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.4 spell_power points (7.18 DPS) | yes | Deep Woodlands Cloak (19121, -0.61 DPS) [quest]; Runecloth Cloak (13860, -2.11 DPS) [crafted]; Nightfall Drape (12465, -2.98 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (416.1 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -7.79 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 13.3 spell_power points (6.21 DPS) | yes | Runic Leather Bracers (15092, -0.47 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.79 DPS) [crafted]; Spidertank Oilrag (9448, -2.02 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (11.19 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Gloves (227187, -1.71 DPS) [vendor]; Earth Warder's Gloves (21318, -5.27 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 18.8 spell_power points (8.76 DPS) | yes | Satyrmane Sash (17755, -1.15 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.11 DPS) [crafted]; Ghostweave Cord (254073, -2.23 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 46.2 spell_power points (21.52 DPS) | yes | General's Dragonhide Pants (231680, -5.56 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -8.04 DPS) [vendor]; Luminary Kilt (11823, -10.89 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (11.19 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.31 DPS) [pvp]; Omnicast Boots (11822, -0.56 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 19.6 spell_power points (9.15 DPS) | yes | Philanthropist's Ring (281635, -3.95 DPS) [quest]; Snake Hoop (6750, -5.81 DPS) [quest]; Reedknot Ring (9622, -5.89 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.06 DPS) | yes | Philanthropist's Ring (281635, -0.86 DPS) [quest]; Snake Hoop (6750, -2.72 DPS) [quest]; Reedknot Ring (9622, -2.80 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (416.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (416.1 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (416.1 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -14.10 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (9.79 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -1.28 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.46 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

