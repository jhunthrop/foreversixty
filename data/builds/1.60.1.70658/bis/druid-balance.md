# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 36.5. Weights run: 2.5s. Verify run: 1.2s. 196 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.754 ± 0.009, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.435 per %), hit=0.121 ± 0.010 per rating point (10 rating = 1%, 1.208 per %), spell_haste=not significant (-0.363 ± 0.140), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.022 ± 0.000, arcane_power=0.978 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.49 DPS) | yes | Wisdom's Leather Hood (252507, -0.50 DPS) [crafted]; Pristine Circlet (253949, -0.50 DPS) [crafted]; Trapper's Leather Hood (252505, -1.21 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.8 spell_power points (1.95 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.47 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.29 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.17 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.47 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.12 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.25 DPS) [dungeon]; Owl Bracers (4796, -0.61 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.0 spell_power points (1.33 DPS) | yes | Serpent Gloves (5970, -0.17 DPS) [dungeon]; Windfelt Gloves (5630, -0.29 DPS) [quest]; Pristine Gloves (253913, -0.29 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.16 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.86 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 14.5 spell_power points (2.40 DPS) | yes | Wisdom's Leather Pants (252503, -0.50 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.66 DPS) [crafted]; Dreamer's Leggings (270016, -1.02 DPS, sim-verified) [quest] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 spell_power points (1.66 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.37 DPS) [crafted]; Black Whelp Slippers (252424, -0.66 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (1.08 DPS) | yes | Volcanic Rock Ring (12053, -0.70 DPS) [world_drop]; The 1 Ring (8350, -0.95 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.5 spell_power points (0.75 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.62 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.32 DPS) | yes | Lesser Staff of the Spire (1300, -0.58 DPS) [world]; Staff of Westfall (2042, -0.70 DPS) [quest]; Twisted Chanter's Staff (890, -0.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 71.9. Weights run: 2.8s. Verify run: 1.4s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.572 ± 0.011, crit=0.050 ± 0.002 per rating point (14 rating = 1%, 0.698 per %), hit=0.114 ± 0.012 per rating point (10 rating = 1%, 1.136 per %), spell_haste=not significant (-0.377 ± 0.160), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.304 ± 0.001, arcane_power=0.696 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.66 DPS) | yes | Enchanter's Cowl (4322, -0.06 DPS) [crafted]; Holy Shroud (2721, -0.22 DPS) [world_drop]; Enduring Cap (3020, -0.63 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.76 DPS) | yes | Crystal Starfire Medallion (5003, -0.25 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.25 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (3.14 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS) [dungeon]; Fairywing Mantle (9536, -0.67 DPS) [quest]; Magician's Mantle (12998, -0.89 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (1.11 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Prelacy Cape (7004, -0.22 DPS) [quest] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.22 DPS) [dungeon]; Guardian Armor (4256, -0.25 DPS) [crafted]; Green Silk Armor (7065, -1.67 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.00 DPS) | yes | Nightsky Wristbands (6407, -1.24 DPS) [world_drop]; Technician's Bracers (270042, -1.24 DPS) [quest]; Glowing Magical Bracelets (13106, -2.25 DPS, sim-verified) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.3 spell_power points (1.62 DPS) | yes | Town Clerk's Mittens (270029, -0.06 DPS) [quest]; Shilly Mitts (9609, -0.06 DPS) [quest]; Gloves of Insight (9698, -0.06 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.4 spell_power points (2.98 DPS) | yes | Belt of Arugal (6392, -0.60 DPS) [dungeon]; Crimson Silk Belt (7055, -0.76 DPS) [crafted]; Moss Cinch (6911, -1.58 DPS, sim-verified) [dungeon] |
| legs | Dark Ritual Leggings (270031) | Blackfathom Villainy [quest] | 14.0 spell_power points (3.11 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Guardian Pants (5962, -0.32 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.44 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (2.44 DPS) | yes | Spidersilk Boots (4320, -0.38 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.48 DPS) [crafted]; Acidic Walkers (9454, -2.24 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.9 spell_power points (1.97 DPS) | yes | Lavishly Jeweled Ring (1156, -1.21 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.30 DPS) [dungeon]; Minor Channeling Ring (1449, -2.75 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon]; Minor Channeling Ring (1449, -0.95 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -3.55 DPS) [quest]; Manual Crowd Pummeler (9449, -12.36 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; legs: Dark Ritual Leggings; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 92.6. Weights run: 3.0s. Verify run: 1.6s. 516 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.369 ± 0.011, crit=0.083 ± 0.003 per rating point (14 rating = 1%, 1.161 per %), hit=0.239 ± 0.018 per rating point (10 rating = 1%, 2.385 per %), spell_haste=not significant (0.569 ± 0.156), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.313 ± 0.001, arcane_power=0.687 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.91 DPS) | yes | Augural Shroud (2620, -1.48 DPS) [world]; Papal Fez (9431, -1.57 DPS) [dungeon]; Big Voodoo Mask (8201, -1.60 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.0 spell_power points (2.33 DPS) | yes | Triune Amulet (7722, -1.57 DPS, sim-verified) [dungeon]; Necklace of Calisea (1714, -1.72 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.81 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.3 spell_power points (2.88 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Inquisitor's Shawl (19507, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.38 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.8 spell_power points (1.83 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.67 DPS) [quest] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.2 spell_power points (5.66 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.75 DPS) [world_drop]; Robe of Power (7054, -1.35 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Leather Bracers (4260, -0.18 DPS) [crafted]; Condor Bracers (15864, -0.47 DPS) [quest]; Arcane Runed Bracers (4744, -1.16 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (5.61 DPS) | yes | Dreamweave Gloves (10019, -1.86 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.11 DPS) [crafted]; Red Mageweave Gloves (10018, -2.18 DPS) [crafted] |
| waist | Star Belt (4329) | Tailoring [crafted] | 13.0 spell_power points (3.04 DPS) | yes | Deathmage Sash (10771, -0.11 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.18 DPS) [crafted]; Moss Cinch (6911, -0.23 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.4 spell_power points (4.31 DPS) | yes | Crimson Silk Pantaloons (7062, -1.32 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.45 DPS) [crafted]; Dark Ritual Leggings (270031, -2.83 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.61 DPS) | yes | Skycaller's Leather Shoes (252532, -1.61 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -3.14 DPS) [crafted]; Gilded Slippers (254001, -3.37 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.8 spell_power points (2.77 DPS) | yes | Snake Hoop (6750, -0.94 DPS) [quest]; Reedknot Ring (9622, -1.13 DPS) [quest]; Ring of Power Regulation (273030, -1.37 DPS) [dungeon] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 spell_power points (1.87 DPS) | yes | Snake Hoop (6750, -0.04 DPS) [quest]; Reedknot Ring (9622, -0.23 DPS) [quest]; Ring of Power Regulation (273030, -0.47 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.17 DPS) [dungeon]; Spine of the Devourer (273040, -2.56 DPS) [world] |
| off_hand | Orb of Lorica (11262) | In the Name of the Light [quest] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Orb of Mystic Insight (249394, -0.11 DPS) [crafted]; Beacon of Hope (9393, -0.39 DPS) [dungeon]; Orb of the Forgotten Seer (7685, -0.96 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Robe of the Magi; hands: Gloves of the Greatfather; waist: Star Belt; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ring of Forlorn Spirits; main_hand: Hypnotic Blade; off_hand: Orb of Lorica

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 153.4. Weights run: 3.3s. Verify run: 1.8s. 686 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.539 ± 0.016, crit=0.125 ± 0.005 per rating point (14 rating = 1%, 1.755 per %), hit=0.276 ± 0.024 per rating point (10 rating = 1%, 2.757 per %), spell_haste=not significant (0.667 ± 0.197), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.352 ± 0.001, arcane_power=0.648 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.1 spell_power points (6.96 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.91 DPS) [crafted]; Dreamweave Circlet (10041, -1.05 DPS) [crafted]; Red Mageweave Headband (10033, -1.71 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.3 spell_power points (2.53 DPS) | yes | Mindburst Medallion (11196, -0.47 DPS) [quest]; Triune Amulet (7722, -0.74 DPS) [dungeon]; Horizon Choker (13085, -0.84 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 25.8 spell_power points (5.77 DPS) | yes | Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.26 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.77 DPS) [crafted]; Rotgrip Mantle (17732, -2.75 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.2 spell_power points (3.86 DPS) | yes | Big Voodoo Cloak (8216, -1.65 DPS) [crafted]; Nightfall Drape (12465, -1.84 DPS) [dungeon]; Runecloth Cloak (13860, -2.06 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 29.8 spell_power points (6.66 DPS) | yes | Feathered Breastplate (8349, -1.21 DPS) [crafted]; Runecloth Tunic (13857, -1.53 DPS) [crafted]; Robe of the Magi (1716, -2.34 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 15.4 spell_power points (3.44 DPS) | yes | Skycaller's Leather Bracers (252542, -0.36 DPS) [crafted]; Nethergeld Cuffs (254061, -1.03 DPS) [crafted]; Bloodband Bracers (11469, -1.24 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (5.37 DPS) | yes | Dreamweave Gloves (10019, -0.86 DPS) [crafted]; Sergeant Major's Crackling Leather Gauntlets (220866, -0.91 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -1.62 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 22.5 spell_power points (5.03 DPS) | yes | Satyrmane Sash (17755, -0.69 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.34 DPS) [crafted]; Dawnspire Cord (12466, -1.39 DPS) [dungeon] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 21.9 spell_power points (4.90 DPS) | yes | Big Voodoo Pants (8202, -0.34 DPS) [crafted]; Wizardweave Leggings (14132, -0.65 DPS) [crafted]; Red Mageweave Pants (10009, -1.52 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.37 DPS) | yes | Sergeant Major's Crackling Leather Boots (220862, -1.09 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.62 DPS) [crafted]; Skycaller's Leather Boots (252471, -1.83 DPS, sim-verified) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.91 DPS) | yes | Snake Hoop (6750, -0.96 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.12 DPS) [quest]; Reedknot Ring (9622, -1.34 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (2.84 DPS) | yes | Snake Hoop (6750, -0.90 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.05 DPS) [quest]; Reedknot Ring (9622, -1.27 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (153.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (153.4 DPS) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (153.4 DPS) | yes | Spellshifter Rod (9527, -0.02 DPS) [quest]; Glowing Brightwood Staff (812, -0.19 DPS) [world_drop]; Hammer of the Northern Wind (810, -25.70 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Knight's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 201.3. Weights run: 3.1s. Verify run: 5.4s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.592 ± 0.024, crit=0.164 ± 0.006 per rating point (14 rating = 1%, 2.293 per %), hit=0.433 ± 0.033 per rating point (10 rating = 1%, 4.332 per %), spell_haste=not significant (0.693 ± 0.317), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.347 ± 0.001, arcane_power=0.653 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 35.7 spell_power points (8.49 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.41 DPS) [vendor]; Soothsayer's Headdress (17740, -11.60 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.9 spell_power points (4.27 DPS) | yes | Glowing Eye of Mordresh (10769, -1.47 DPS) [dungeon]; Mindburst Medallion (11196, -1.99 DPS) [quest]; Horizon Choker (13085, -2.29 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (201.3 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Ironfeather Shoulders (15067, -3.89 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.6 spell_power points (4.18 DPS) | yes | Big Voodoo Cloak (8216, -1.72 DPS) [crafted]; Nightfall Drape (12465, -2.04 DPS) [dungeon]; Runecloth Cloak (13860, -2.58 DPS, sim-verified) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (201.3 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -5.41 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (201.3 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -7.43 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (201.3 DPS) | yes | Knight-Lieutenant's Dragonhide Gloves (227193, +0.00 DPS) [pvp]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Gloves of the Greatfather (17721, -7.25 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (201.3 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, -8.69 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.0 spell_power points (12.13 DPS) | yes | Marshal's Dragonhide Pants (231697, -3.18 DPS) [vendor]; Knight-Captain's Dragonhide Pants (227191, -4.19 DPS) [vendor]; Luminary Kilt (11823, -4.53 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 27.1 spell_power points (6.46 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -0.04 DPS) [vendor]; Earthen Silk Slippers (254013, -0.74 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 22.1 spell_power points (5.27 DPS) | yes | Philanthropist's Ring (281635, -2.19 DPS) [quest]; Snake Hoop (6750, -3.14 DPS) [quest]; Ring of Forlorn Spirits (2043, -3.37 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.10 DPS) | yes | Philanthropist's Ring (281635, -0.01 DPS) [quest]; Snake Hoop (6750, -0.96 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.19 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (201.3 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (201.3 DPS) | yes | Mark of the Chosen (17774, -4.02 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (201.3 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -7.16 DPS) [quest] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (5.00 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -1.49 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5053000000000000)

Set DPS (verified): 494.3. Weights run: 2.1s. Verify run: 3.8s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.108 ± 0.010, crit=0.279 ± 0.012 per rating point (14 rating = 1%, 3.907 per %), hit=0.593 ± 0.031 per rating point (10 rating = 1%, 5.930 per %), spell_haste=3.574 ± 0.241, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.326 ± 0.000, arcane_power=0.674 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (494.3 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.05 DPS) [vendor]; Spellpower Goggles Xtreme Plus (15999, -15.56 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.1 spell_power points (7.26 DPS) | yes | Triune Amulet (7722, -2.82 DPS) [dungeon]; Glowing Eye of Mordresh (10769, -2.90 DPS) [dungeon]; Mindburst Medallion (11196, -3.57 DPS) [quest] |
| shoulder | Argent Shoulders (19059) | Tailoring [crafted] | 25.0 spell_power points (13.87 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Mantle of the Timbermaw (19050, -1.49 DPS) [crafted]; Lieutenant Commander's Dragonhide Spaulders (227190, -2.72 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (8.13 DPS) | yes | Nightfall Drape (12465, -3.14 DPS) [dungeon]; Runecloth Cloak (13860, -3.60 DPS, sim-verified) [crafted]; Icy Cloak (4327, -4.25 DPS) [crafted] |
| chest | Fernpulse Jerkin (15786) | High Chief Winterfall [quest] | 26.9 spell_power points (14.91 DPS) | yes | Robe of the Magi (1716, +0.00 DPS) [world_drop]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Knight-Captain's Dragonhide Armor (227195, -0.07 DPS) [vendor] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.1 spell_power points (6.71 DPS) | yes | Runic Leather Bracers (15092, -0.55 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.74 DPS) [crafted]; Arcane Runed Bracers (4744, -1.71 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (13.32 DPS) | yes | Marshal's Dragonhide Gloves (231700, -0.39 DPS) [vendor]; Knight-Lieutenant's Dragonhide Gloves (227193, -2.73 DPS) [pvp]; Earth Warder's Gloves (21318, -6.14 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 17.3 spell_power points (9.60 DPS) | yes | Satyrmane Sash (17755, -1.23 DPS) [dungeon]; Ghostweave Cord (254073, -1.83 DPS) [crafted]; Star Belt (4329, -2.39 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 43.9 spell_power points (24.34 DPS) | yes | Marshal's Dragonhide Pants (231697, -6.35 DPS) [vendor]; Knight-Captain's Dragonhide Pants (227191, -9.49 DPS) [vendor]; Ghostweave Pants (14144, -11.75 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.32 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.82 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -1.06 DPS) [vendor]; Venomspew Footpads (275606, -1.11 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.8 spell_power points (10.41 DPS) | yes | Philanthropist's Ring (281635, -4.56 DPS) [quest]; Ring of Forlorn Spirits (2043, -5.97 DPS) [quest]; Reedknot Ring (9622, -6.53 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.21 DPS) | yes | Philanthropist's Ring (281635, -1.36 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.77 DPS) [quest]; Reedknot Ring (9622, -3.33 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -6.91 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, -0.06 DPS) [pvp]; Spire of Hakkar (10844, -16.32 DPS) [world] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (11.65 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -2.49 DPS) [pvp]; Orb of the Forgotten Seer (7685, -4.03 DPS, sim-verified) [dungeon] |
| ranged | Idol of the Moon (23197) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talons of Wrath (249441, -3.76 DPS, sim-verified) [crafted] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Argent Shoulders; chest: Fernpulse Jerkin; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas; ranged: Idol of the Moon

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 35.5. Weights run: 2.5s. Verify run: 1.4s. 192 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.754 ± 0.009, crit=0.031 ± 0.001 per rating point (14 rating = 1%, 0.435 per %), hit=0.121 ± 0.010 per rating point (10 rating = 1%, 1.208 per %), spell_haste=not significant (-0.363 ± 0.140), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.022 ± 0.000, arcane_power=0.978 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.49 DPS) | yes | Wisdom's Leather Hood (252507, -0.50 DPS) [crafted]; Pristine Circlet (253949, -0.50 DPS) [crafted]; Trapper's Leather Hood (252505, -1.23 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.8 spell_power points (1.95 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.92 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.29 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.3 spell_power points (0.71 DPS) | yes | Heavy Woolen Cloak (4311, -0.04 DPS) [crafted]; Sanguine Cape (14376, -0.21 DPS) [world_drop]; Black Whelp Cloak (7283, -0.21 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.47 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.8 spell_power points (0.62 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.12 DPS) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.0 spell_power points (1.33 DPS) | yes | Serpent Gloves (5970, -0.17 DPS) [dungeon]; Pristine Gloves (253913, -0.29 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.33 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.12 DPS) [crafted]; Keller's Girdle (2911, -0.16 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.86 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.17 DPS) [crafted]; Totemic Leather Pants (252446, -0.25 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.86 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 spell_power points (1.66 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.37 DPS) [crafted]; Black Whelp Slippers (252424, -0.66 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.5 spell_power points (0.75 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS) [world_drop]; The 1 Ring (8350, -0.62 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 3.8 spell_power points (0.62 DPS) | yes | Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; The 1 Ring (8350, -0.50 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.32 DPS) | yes | Twisted Chanter's Staff (890, -0.08 DPS) [world_drop]; Lesser Staff of the Spire (1300, -0.58 DPS) [world]; Gnarled Necromancer's Staff (251534, -0.81 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 72.2. Weights run: 2.8s. Verify run: 1.4s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.572 ± 0.011, crit=0.050 ± 0.002 per rating point (14 rating = 1%, 0.698 per %), hit=0.114 ± 0.012 per rating point (10 rating = 1%, 1.136 per %), spell_haste=not significant (-0.377 ± 0.160), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.304 ± 0.001, arcane_power=0.696 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.66 DPS) | yes | Enchanter's Cowl (4322, -0.06 DPS) [crafted]; Holy Shroud (2721, -0.22 DPS) [world_drop]; Enduring Cap (3020, -0.63 DPS) [world_drop] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.76 DPS) | yes | Crystal Starfire Medallion (5003, -0.25 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.25 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (3.14 DPS) | yes | Death Speaker Mantle (6685, -0.41 DPS) [dungeon]; Fairywing Mantle (9536, -0.67 DPS) [quest]; Magician's Mantle (12998, -0.89 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (1.11 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.22 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.29 DPS) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.22 DPS) [dungeon]; Guardian Armor (4256, -0.25 DPS) [crafted]; Green Silk Armor (7065, -1.40 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.00 DPS) | yes | Nightsky Wristbands (6407, -1.24 DPS) [world_drop]; Technician's Bracers (270042, -1.24 DPS) [quest]; Glowing Magical Bracelets (13106, -1.47 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.1 spell_power points (2.25 DPS) | yes | Jutebraid Gloves (10654, -0.29 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.63 DPS) [crafted]; Serpent Gloves (5970, -0.70 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.4 spell_power points (2.98 DPS) | yes | Warsong Sash (16975, -0.54 DPS) [quest]; Belt of Arugal (6392, -0.60 DPS) [dungeon]; Moss Cinch (6911, -1.57 DPS, sim-verified) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.19 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.32 DPS) [dungeon]; Dark Ritual Leggings (270031, -1.36 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (2.44 DPS) | yes | Spidersilk Boots (4320, -0.38 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.48 DPS) [crafted]; Acidic Walkers (9454, -2.18 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.9 spell_power points (1.97 DPS) | yes | Solid Gold Band (279845, -1.08 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.21 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.30 DPS) [dungeon] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | 6.0 spell_power points (1.33 DPS) | yes | Solid Gold Band (279845, -0.44 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.57 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.67 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, -6.22 DPS) [dungeon]; Manual Crowd Pummeler (9449, -12.21 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 94.0. Weights run: 3.0s. Verify run: 1.6s. 514 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.369 ± 0.011, crit=0.083 ± 0.003 per rating point (14 rating = 1%, 1.161 per %), hit=0.239 ± 0.018 per rating point (10 rating = 1%, 2.385 per %), spell_haste=not significant (0.569 ± 0.156), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.313 ± 0.001, arcane_power=0.687 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.91 DPS) | yes | Augural Shroud (2620, -1.48 DPS) [world]; Papal Fez (9431, -1.57 DPS) [dungeon]; Big Voodoo Mask (8201, -1.60 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.0 spell_power points (2.33 DPS) | yes | Triune Amulet (7722, -0.46 DPS) [dungeon]; Necklace of Calisea (1714, -1.72 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.81 DPS) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.3 spell_power points (2.88 DPS) | yes | Green Silken Shoulders (7057, -0.06 DPS) [crafted]; Inquisitor's Shawl (19507, -0.12 DPS) [dungeon]; Berylline Pads (4197, -0.38 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 7.8 spell_power points (1.83 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.20 DPS) [crafted]; Hillman's Cloak (3719, -0.67 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.07 DPS) [world_drop]; Robe of Power (7054, -0.68 DPS) [crafted]; Robe of the Magi (1716, -1.03 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (2.11 DPS) | yes | Guardian Leather Bracers (4260, -0.18 DPS) [crafted]; Condor Bracers (15864, -0.47 DPS) [quest]; Radiant Silver Bracers (4545, -0.48 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (5.61 DPS) | yes | Dreamweave Gloves (10019, -1.86 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -2.11 DPS) [crafted]; Red Mageweave Gloves (10018, -2.18 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Leather Belt (252522, -0.07 DPS) [crafted]; Moss Cinch (6911, -0.12 DPS) [dungeon]; Star Belt (4329, -1.29 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.4 spell_power points (4.31 DPS) | yes | Crimson Silk Pantaloons (7062, -1.32 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.45 DPS) [crafted]; Dark Ritual Leggings (270031, -1.99 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.61 DPS) | yes | Skycaller's Leather Shoes (252532, -1.38 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -3.14 DPS) [crafted]; Gilded Slippers (254001, -3.37 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 11.8 spell_power points (2.77 DPS) | yes | Reedknot Ring (9622, -1.13 DPS) [quest]; Ring of Power Regulation (273030, -1.37 DPS) [dungeon]; Solid Gold Band (279845, -1.83 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 7.8 spell_power points (1.83 DPS) | yes | Reedknot Ring (9622, -0.20 DPS) [quest]; Ring of Power Regulation (273030, -0.43 DPS) [dungeon]; Solid Gold Band (279845, -0.90 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.17 DPS) [dungeon]; Spine of the Devourer (273040, -2.56 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (3.27 DPS) | yes | Orb of Mystic Insight (249394, -1.12 DPS) [crafted]; Omega Orb (7749, -1.17 DPS) [quest]; Beacon of Hope (9393, -1.40 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 151.8. Weights run: 3.3s. Verify run: 1.8s. 684 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.539 ± 0.016, crit=0.125 ± 0.005 per rating point (14 rating = 1%, 1.755 per %), hit=0.276 ± 0.024 per rating point (10 rating = 1%, 2.757 per %), spell_haste=not significant (0.667 ± 0.197), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.352 ± 0.001, arcane_power=0.648 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.1 spell_power points (6.96 DPS) | yes | Red Mageweave Headband (10033, -0.29 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.91 DPS) [crafted]; Dreamweave Circlet (10041, -1.05 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.3 spell_power points (2.53 DPS) | yes | Mindburst Medallion (11196, -0.47 DPS) [quest]; Triune Amulet (7722, -0.74 DPS) [dungeon]; Horizon Choker (13085, -0.84 DPS) [world_drop] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 25.8 spell_power points (5.77 DPS) | yes | Blood Guard's Crackling Leather Spaulders (220871, -1.26 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.77 DPS) [crafted]; Rotgrip Mantle (17732, -2.88 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.2 spell_power points (3.86 DPS) | yes | Deep Woodlands Cloak (19121, -0.09 DPS) [quest]; Runecloth Cloak (13860, -0.88 DPS) [crafted]; Big Voodoo Cloak (8216, -1.65 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 29.8 spell_power points (6.66 DPS) | yes | Feathered Breastplate (8349, -1.21 DPS) [crafted]; Runecloth Tunic (13857, -1.53 DPS) [crafted]; Robe of the Magi (1716, -2.14 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 15.4 spell_power points (3.44 DPS) | yes | Skycaller's Leather Bracers (252542, -0.36 DPS) [crafted]; Nethergeld Cuffs (254061, -1.03 DPS) [crafted]; Bloodband Bracers (11469, -1.24 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (5.37 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.70 DPS) [crafted]; Dreamweave Gloves (10019, -0.86 DPS) [crafted]; First Sergeant's Crackling Leather Gauntlets (220867, -0.91 DPS) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 22.5 spell_power points (5.03 DPS) | yes | Satyrmane Sash (17755, -0.69 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.34 DPS) [crafted]; Dawnspire Cord (12466, -1.39 DPS) [dungeon] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | PvP rank 10 · Stone Guard · Horde [vendor] | 21.9 spell_power points (4.90 DPS) | yes | Red Mageweave Pants (10009, -0.32 DPS) [crafted]; Big Voodoo Pants (8202, -0.34 DPS) [crafted]; Wizardweave Leggings (14132, -0.65 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (5.37 DPS) | yes | First Sergeant's Crackling Leather Boots (220863, -1.09 DPS) [vendor]; Skycaller's Leather Boots (252471, -1.56 DPS, sim-verified) [crafted]; Skycaller's Leather Shoes (252532, -1.62 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.91 DPS) | yes | Snake Hoop (6750, -0.96 DPS) [quest]; Reedknot Ring (9622, -1.34 DPS) [quest]; Ring of Power Regulation (273030, -1.57 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.7 spell_power points (2.84 DPS) | yes | Snake Hoop (6750, -0.90 DPS) [quest]; Reedknot Ring (9622, -1.27 DPS) [quest]; Ring of Power Regulation (273030, -1.50 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (151.8 DPS) | yes | Rune of the Guard Captain (19120, -0.91 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (151.8 DPS) | yes | Rune of the Guard Captain (19120, -0.12 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (151.8 DPS) | yes | Spellshifter Rod (9527, -0.02 DPS) [quest]; Glowing Brightwood Staff (812, -0.19 DPS) [world_drop]; Hammer of the Northern Wind (810, -25.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Kindling Stave

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 200.7. Weights run: 3.1s. Verify run: 5.4s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.592 ± 0.024, crit=0.164 ± 0.006 per rating point (14 rating = 1%, 2.293 per %), hit=0.433 ± 0.033 per rating point (10 rating = 1%, 4.332 per %), spell_haste=not significant (0.693 ± 0.317), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.347 ± 0.001, arcane_power=0.653 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 35.7 spell_power points (8.49 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.41 DPS) [pvp]; Soothsayer's Headdress (17740, -11.22 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 17.9 spell_power points (4.27 DPS) | yes | Glowing Eye of Mordresh (10769, -1.47 DPS) [dungeon]; Mindburst Medallion (11196, -1.99 DPS) [quest]; Horizon Choker (13085, -2.29 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (200.7 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -5.93 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.6 spell_power points (4.18 DPS) | yes | Runecloth Cloak (13860, -0.91 DPS) [crafted]; Big Voodoo Cloak (8216, -1.72 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.85 DPS, sim-verified) [quest] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (200.7 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -5.56 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (200.7 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -7.41 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (200.7 DPS) | yes | Blood Guard's Dragonhide Gloves (227187, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Gloves of the Greatfather (17721, -6.67 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (200.7 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, -9.01 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.0 spell_power points (12.13 DPS) | yes | General's Dragonhide Pants (231680, -3.18 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -4.19 DPS) [vendor]; Luminary Kilt (11823, -4.39 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 27.1 spell_power points (6.46 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -0.04 DPS) [pvp]; Earthen Silk Slippers (254013, -0.74 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 22.1 spell_power points (5.27 DPS) | yes | Philanthropist's Ring (281635, -2.19 DPS) [quest]; Snake Hoop (6750, -3.14 DPS) [quest]; Glowing Crystal Ring (18402, -3.58 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (3.10 DPS) | yes | Philanthropist's Ring (281635, -0.01 DPS) [quest]; Snake Hoop (6750, -0.96 DPS) [quest]; Glowing Crystal Ring (18402, -1.40 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (200.7 DPS) | yes | Rune of the Guard Captain (19120, -0.21 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (200.7 DPS) | yes | Rune of the Guard Captain (19120, -3.00 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (200.7 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -7.16 DPS) [quest] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (5.00 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -1.49 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5053000000000000)

Set DPS (verified): 492.0. Weights run: 2.1s. Verify run: 4.9s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.001, intellect=0.108 ± 0.010, crit=0.279 ± 0.012 per rating point (14 rating = 1%, 3.907 per %), hit=0.593 ± 0.031 per rating point (10 rating = 1%, 5.930 per %), spell_haste=3.574 ± 0.241, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.326 ± 0.000, arcane_power=0.674 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (492.0 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.05 DPS) [pvp]; Spellpower Goggles Xtreme Plus (15999, -15.82 DPS, sim-verified) [crafted] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 13.1 spell_power points (7.26 DPS) | yes | Glowing Eye of Mordresh (10769, -2.90 DPS) [dungeon]; Triune Amulet (7722, -3.49 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -3.57 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (492.0 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Argent Shoulders (19059, -8.55 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 spell_power points (8.13 DPS) | yes | Deep Woodlands Cloak (19121, -0.93 DPS) [quest]; Runecloth Cloak (13860, -2.65 DPS) [crafted]; Nightfall Drape (12465, -3.14 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (492.0 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -9.18 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 12.1 spell_power points (6.71 DPS) | yes | Runic Leather Bracers (15092, -0.55 DPS) [crafted]; Skycaller's Leather Bracers (252542, -0.74 DPS) [crafted]; Spidertank Oilrag (9448, -1.71 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (13.32 DPS) | yes | General's Dragonhide Gloves (231677, -0.39 DPS) [pvp]; Blood Guard's Dragonhide Gloves (227187, -2.73 DPS) [vendor]; Earth Warder's Gloves (21318, -6.22 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 17.3 spell_power points (9.60 DPS) | yes | Satyrmane Sash (17755, -1.23 DPS) [dungeon]; Ghostweave Cord (254073, -1.83 DPS) [crafted]; Star Belt (4329, -2.39 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 43.9 spell_power points (24.34 DPS) | yes | General's Dragonhide Pants (231680, -6.35 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -9.49 DPS) [vendor]; Ghostweave Pants (14144, -14.44 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (13.32 DPS) | yes | General's Dragonhide Boots (231682, -0.82 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -1.06 DPS) [pvp]; Venomspew Footpads (275606, -1.11 DPS) [crafted] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 18.8 spell_power points (10.41 DPS) | yes | Philanthropist's Ring (281635, -4.56 DPS) [quest]; Reedknot Ring (9622, -6.53 DPS) [quest]; Snake Hoop (6750, -6.78 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (7.21 DPS) | yes | Philanthropist's Ring (281635, -2.81 DPS, sim-verified) [quest]; Reedknot Ring (9622, -3.33 DPS) [quest]; Snake Hoop (6750, -3.58 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (492.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (492.0 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (492.0 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, -0.06 DPS) [pvp]; Spire of Hakkar (10844, -16.32 DPS) [world] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (11.65 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -2.49 DPS) [pvp]; Orb of the Forgotten Seer (7685, -4.08 DPS, sim-verified) [dungeon] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

