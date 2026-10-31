# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 30.9. Weights run: 1.8s. Verify run: 0.9s. 196 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.841 ± 0.006, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.397 per %), hit=0.092 ± 0.005 per rating point (10 rating = 1%, 0.916 per %), spell_haste=not significant (0.115 ± 0.083), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.029 ± 0.001, arcane_power=0.971 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Wisdom's Leather Hood (252507, -0.40 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS) [crafted]; Trapper's Leather Hood (252505, -1.33 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.6 spell_power points (1.66 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -1.13 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.25 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.08 DPS) [world_drop]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.40 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.09 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.41 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.11 DPS) [world_drop]; Inscribed Leather Bracers (3205, -0.22 DPS) [dungeon]; Owl Bracers (4796, -0.59 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.4 spell_power points (1.11 DPS) | yes | Windfelt Gloves (5630, -0.24 DPS) [quest]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.35 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Keller's Girdle (2911, -0.08 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.75 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 15.0 spell_power points (1.99 DPS) | yes | Wisdom's Leather Pants (252503, -0.40 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.53 DPS) [crafted]; Dreamer's Leggings (270016, -1.25 DPS, sim-verified) [quest] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.4 spell_power points (1.37 DPS) | yes | Stormrider's Leather Boots (252443, -0.02 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.29 DPS) [crafted]; Black Whelp Slippers (252424, -0.53 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.7 spell_power points (0.88 DPS) | yes | Volcanic Rock Ring (12053, -0.55 DPS) [world_drop]; The 1 Ring (8350, -0.77 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.0 spell_power points (0.67 DPS) | yes | Volcanic Rock Ring (12053, -0.37 DPS, sim-verified) [world_drop]; The 1 Ring (8350, -0.56 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Lesser Staff of the Spire (1300, -0.39 DPS) [world]; Twisted Chanter's Staff (890, -0.39 DPS, sim-verified) [world_drop]; Staff of Westfall (2042, -0.50 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 59.2. Weights run: 1.9s. Verify run: 1.0s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.561 ± 0.007, crit=0.035 ± 0.002 per rating point (14 rating = 1%, 0.489 per %), hit=0.128 ± 0.006 per rating point (10 rating = 1%, 1.280 per %), spell_haste=not significant (-0.126 ± 0.095), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.333 ± 0.001, arcane_power=0.667 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.14 DPS) | yes | Enchanter's Cowl (4322, -0.07 DPS) [crafted]; Holy Shroud (2721, -0.18 DPS) [world_drop]; Silk Headband (7050, -0.54 DPS) [crafted] |
| neck | Archmagister's Faceted Pendant (273049) | Shade of the Archmage [world] | 3.4 spell_power points (0.60 DPS) | yes | Kaleidoscope Chain (13084, -0.20 DPS) [world_drop]; Crystal Starfire Medallion (5003, -0.54 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (2.51 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.54 DPS) [quest]; Magician's Mantle (12998, -0.71 DPS) [world_drop] |
| back | Vine Pruner's Cloak (279835) (or Hillman's Cloak (3719)) | A Green Sample [quest] | 5.0 spell_power points (0.89 DPS) | yes | Hillman's Cloak (3719, +0.00 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Prelacy Cape (7004, -0.18 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (2.91 DPS) | yes | Beguiler Robes (7728, -0.10 DPS) [dungeon]; Mechbuilder's Overalls (9508, -0.28 DPS) [dungeon]; Guardian Armor (4256, -0.30 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.61 DPS) | yes | Nightsky Wristbands (6407, -1.01 DPS) [world_drop]; Technician's Bracers (270042, -1.01 DPS) [quest]; Glowing Magical Bracelets (13106, -1.90 DPS, sim-verified) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.2 spell_power points (1.29 DPS) | yes | Serpent Gloves (5970, -0.04 DPS) [dungeon]; Shilly Mitts (9609, -0.04 DPS) [quest]; Gloves of Insight (9698, -1.11 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.4 spell_power points (2.39 DPS) | yes | Belt of Arugal (6392, -0.48 DPS) [dungeon]; Crimson Silk Belt (7055, -0.61 DPS) [crafted]; Moss Cinch (6911, -2.21 DPS, sim-verified) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.16 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.24 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.63 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.95 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.38 DPS) [crafted]; Acidic Walkers (9454, -1.68 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.57 DPS) | yes | Lavishly Jeweled Ring (1156, -0.97 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.04 DPS) [dungeon]; Minor Channeling Ring (1449, -3.06 DPS, sim-verified) [quest] |
| finger2 | Ring of Power Regulation (273030) | Excavation Site: Wetlands: Relic Guardian [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.47 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Minor Channeling Ring (1449, -1.00 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Hardened Root Staff (1317, -2.86 DPS) [quest]; Manual Crowd Pummeler (9449, -10.38 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Archmagister's Faceted Pendant; shoulder: Bloodmage Mantle; back: Vine Pruner's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Skycaller's Leather Belt; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Ring of Power Regulation; main_hand: Spine of the Devourer

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 81.7. Weights run: 1.9s. Verify run: 1.1s. 516 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.437 ± 0.006, crit=0.065 ± 0.003 per rating point (14 rating = 1%, 0.912 per %), hit=0.167 ± 0.008 per rating point (10 rating = 1%, 1.666 per %), spell_haste=0.279 ± 0.029, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.01 DPS) [world]; Big Voodoo Mask (8201, -0.06 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.24 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.5 spell_power points (2.01 DPS) | yes | Necklace of Calisea (1714, -1.43 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.51 DPS) [world]; Triune Amulet (7722, -1.90 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 spell_power points (2.48 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Inquisitor's Shawl (19507, -0.05 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.2 spell_power points (1.57 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.23 DPS) [crafted]; Vine Pruner's Cloak (279835, -0.61 DPS) [quest] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.18 DPS) [world_drop]; Robe of Power (7054, -0.52 DPS) [crafted]; Robe of the Magi (1716, -0.80 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Leather Bracers (4260, -0.07 DPS) [crafted]; Condor Bracers (15864, -0.38 DPS) [quest]; Arcane Runed Bracers (4744, -1.03 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.60 DPS) | yes | Red Mageweave Gloves (10018, -1.65 DPS) [crafted]; Black Mageweave Gloves (10003, -1.73 DPS) [crafted]; Dreamweave Gloves (10019, -1.80 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 spell_power points (2.60 DPS) | yes | Skycaller's Leather Belt (252522, -0.18 DPS) [crafted]; Moss Cinch (6911, -0.30 DPS) [dungeon]; Star Belt (4329, -1.57 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.2 spell_power points (3.69 DPS) | yes | Crimson Silk Pantaloons (7062, -1.07 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.27 DPS) [crafted]; Dark Ritual Leggings (270031, -3.03 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.60 DPS) | yes | Skycaller's Leather Shoes (252532, -1.50 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.48 DPS) [crafted]; Gilded Slippers (254001, -2.67 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.2 spell_power points (2.34 DPS) | yes | Ring of Forlorn Spirits (2043, -0.80 DPS) [quest]; Reedknot Ring (9622, -0.99 DPS) [quest]; Ring of Power Regulation (273030, -1.19 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.2 spell_power points (1.57 DPS) | yes | Ring of Forlorn Spirits (2043, -0.04 DPS) [quest]; Reedknot Ring (9622, -0.23 DPS) [quest]; Ring of Power Regulation (273030, -0.42 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.17 DPS) [dungeon]; Spine of the Devourer (273040, -2.21 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.69 DPS) | yes | Orb of Lorica (11262, +0.00 DPS, sim-verified) [quest]; Orb of Mystic Insight (249394, -0.84 DPS) [crafted]; Beacon of Hope (9393, -1.15 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 130.2. Weights run: 2.1s. Verify run: 1.3s. 686 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.589 ± 0.009, crit=0.096 ± 0.005 per rating point (14 rating = 1%, 1.350 per %), hit=0.288 ± 0.015 per rating point (10 rating = 1%, 2.875 per %), spell_haste=0.566 ± 0.070, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.392 ± 0.001, arcane_power=0.608 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.8 spell_power points (5.73 DPS) | yes | Red Mageweave Headband (10033, -0.19 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.87 DPS) [crafted]; Dreamweave Circlet (10041, -0.89 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (2.11 DPS) | yes | Mindburst Medallion (11196, -0.39 DPS) [quest]; Horizon Choker (13085, -0.62 DPS) [world_drop]; Triune Amulet (7722, -0.67 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 26.8 spell_power points (4.82 DPS) | yes | Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.08 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.53 DPS) [crafted]; Rotgrip Mantle (17732, -2.54 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.5 spell_power points (3.16 DPS) | yes | Big Voodoo Cloak (8216, -1.30 DPS) [crafted]; Nightfall Drape (12465, -1.54 DPS) [dungeon]; Runecloth Cloak (13860, -1.58 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 30.8 spell_power points (5.54 DPS) | yes | Feathered Breastplate (8349, -1.06 DPS) [crafted]; Robe of the Magi (1716, -1.29 DPS, sim-verified) [world_drop]; Runecloth Tunic (13857, -1.31 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 15.9 spell_power points (2.86 DPS) | yes | Skycaller's Leather Bracers (252542, -0.32 DPS) [crafted]; Nethergeld Cuffs (254061, -0.86 DPS) [crafted]; Bloodband Bracers (11469, -1.01 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.32 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.49 DPS) [crafted]; Sergeant Major's Crackling Leather Gauntlets (220866, -0.63 DPS) [vendor]; Dreamweave Gloves (10019, -0.66 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 23.1 spell_power points (4.15 DPS) | yes | Satyrmane Sash (17755, -0.57 DPS) [dungeon]; Dawnspire Cord (12466, -1.06 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.08 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 22.1 spell_power points (3.98 DPS) | yes | Big Voodoo Pants (8202, -0.22 DPS) [crafted]; Wizardweave Leggings (14132, -0.56 DPS) [crafted]; Red Mageweave Pants (10009, -1.49 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.32 DPS) | yes | Skycaller's Leather Boots (252471, -0.45 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -0.76 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.24 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.34 DPS) | yes | Snake Hoop (6750, -0.73 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.90 DPS) [quest]; Reedknot Ring (9622, -1.08 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (2.33 DPS) | yes | Snake Hoop (6750, -0.72 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.89 DPS) [quest]; Reedknot Ring (9622, -1.07 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (130.2 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (130.2 DPS) | yes | Mark of the Chosen (17774, -1.22 DPS, sim-verified) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (130.2 DPS) | yes | Glowing Brightwood Staff (812, -0.08 DPS) [world_drop]; Kindling Stave (11750, -0.09 DPS) [dungeon]; Hammer of the Northern Wind (810, -22.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Knight's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spellshifter Rod

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 170.4. Weights run: 2.0s. Verify run: 3.8s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.650 ± 0.012, crit=0.127 ± 0.007 per rating point (14 rating = 1%, 1.777 per %), hit=0.374 ± 0.018 per rating point (10 rating = 1%, 3.735 per %), spell_haste=not significant (0.268 ± 0.074), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.390 ± 0.001, arcane_power=0.610 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 36.7 spell_power points (7.09 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.43 DPS) [vendor]; Soothsayer's Headdress (17740, -9.87 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.5 spell_power points (3.57 DPS) | yes | Glowing Eye of Mordresh (10769, -1.22 DPS) [dungeon]; Mindburst Medallion (11196, -1.66 DPS) [quest]; Horizon Choker (13085, -1.82 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (170.4 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Ironfeather Shoulders (15067, -5.76 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.9 spell_power points (3.46 DPS) | yes | Big Voodoo Cloak (8216, -1.36 DPS) [crafted]; Long Silken Cloak (4326, -1.67 DPS) [crafted]; Runecloth Cloak (13860, -2.24 DPS, sim-verified) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (170.4 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -5.15 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (170.4 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -8.18 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (170.4 DPS) | yes | Knight-Lieutenant's Dragonhide Gloves (227193, +0.00 DPS) [pvp]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Gloves of the Greatfather (17721, -7.43 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (170.4 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, -9.09 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.5 spell_power points (9.95 DPS) | yes | Marshal's Dragonhide Pants (231697, -2.63 DPS) [vendor]; Knight-Captain's Dragonhide Pants (227191, -3.41 DPS) [vendor]; Luminary Kilt (11823, -3.53 DPS, sim-verified) [dungeon] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 27.8 spell_power points (5.37 DPS) | yes | Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, +0.00 DPS) [crafted]; Knight-Lieutenant's Dragonhide Boots (227194, -0.06 DPS) [vendor] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 22.6 spell_power points (4.36 DPS) | yes | Band of the Unicorn (7553, -1.85 DPS) [world_drop]; Snake Hoop (6750, -2.57 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.81 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (2.56 DPS) | yes | Band of the Unicorn (7553, -0.05 DPS) [world_drop]; Snake Hoop (6750, -0.77 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.01 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (170.4 DPS) | yes | - |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (170.4 DPS) | yes | Mark of the Chosen (17774, -1.86 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (170.4 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -5.70 DPS) [quest] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (4.06 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -1.82 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 5232220115501351-00000000000000000000-5053000000000000)

Set DPS (verified): 479.9. Weights run: 1.4s. Verify run: 3.5s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.323 ± 0.013, crit=0.226 ± 0.011 per rating point (14 rating = 1%, 3.168 per %), hit=0.565 ± 0.031 per rating point (10 rating = 1%, 5.647 per %), spell_haste=2.298 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.322 ± 0.000, arcane_power=0.678 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 30.8 spell_power points (16.05 DPS) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.43 DPS) [vendor]; Soothsayer's Headdress (17740, -15.48 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.2 spell_power points (7.93 DPS) | yes | Glowing Eye of Mordresh (10769, -2.94 DPS) [dungeon]; Triune Amulet (7722, -3.77 DPS) [dungeon]; Mindburst Medallion (11196, -3.80 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (479.9 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Argent Shoulders (19059, -6.76 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (8.30 DPS) | yes | Nightfall Drape (12465, -3.61 DPS) [dungeon]; Runecloth Cloak (13860, -3.78 DPS, sim-verified) [crafted]; Big Voodoo Cloak (8216, -4.18 DPS) [crafted] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (479.9 DPS) | yes | Knight-Captain's Dragonhide Armor (227195, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -10.07 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 14.2 spell_power points (7.41 DPS) | yes | Runic Leather Bracers (15092, -0.52 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.03 DPS) [crafted]; Nethergeld Cuffs (254061, -2.59 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (12.50 DPS) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Knight-Lieutenant's Dragonhide Gloves (227193, -1.44 DPS) [pvp]; Earth Warder's Gloves (21318, -5.98 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 19.9 spell_power points (10.35 DPS) | yes | Satyrmane Sash (17755, -1.38 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.27 DPS) [crafted]; Ghostweave Cord (254073, -3.06 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 47.0 spell_power points (24.47 DPS) | yes | Marshal's Dragonhide Pants (231697, -6.40 DPS) [vendor]; Knight-Captain's Dragonhide Pants (227191, -9.02 DPS) [vendor]; Luminary Kilt (11823, -12.84 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.50 DPS) | yes | Knight-Lieutenant's Dragonhide Boots (227194, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Omnicast Boots (11822, -2.47 DPS, sim-verified) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.3 spell_power points (10.55 DPS) | yes | Philanthropist's Ring (281635, -4.50 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.39 DPS) [quest]; Snake Hoop (6750, -6.59 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.77 DPS) | yes | Philanthropist's Ring (281635, -0.72 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.60 DPS) [quest]; Snake Hoop (6750, -2.81 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (479.9 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (479.9 DPS) | yes | Mark of the Chosen (17774, -6.73 DPS, sim-verified) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (479.9 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Kindling Stave (11750, -15.80 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (10.94 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, -0.77 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.92 DPS, sim-verified) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-verified (479.9 DPS) | yes | Idol of the Moon (23197, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 29.2. Weights run: 1.8s. Verify run: 0.9s. 192 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.841 ± 0.006, crit=0.028 ± 0.001 per rating point (14 rating = 1%, 0.397 per %), hit=0.092 ± 0.005 per rating point (10 rating = 1%, 0.916 per %), spell_haste=not significant (0.115 ± 0.083), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.029 ± 0.001, arcane_power=0.971 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.19 DPS) | yes | Wisdom's Leather Hood (252507, -0.40 DPS) [crafted]; Pristine Circlet (253949, -0.40 DPS) [crafted]; Trapper's Leather Hood (252505, -1.25 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.6 spell_power points (1.66 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -1.13 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.16 DPS, sim-verified) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.5 spell_power points (0.60 DPS) | yes | Heavy Woolen Cloak (4311, -0.07 DPS) [crafted]; Sanguine Cape (14376, -0.15 DPS) [world_drop]; Black Whelp Cloak (7283, -0.20 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.09 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.41 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Featherbead Bracers (15452), Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 4.2 spell_power points (0.56 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.11 DPS) [world_drop] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.4 spell_power points (1.11 DPS) | yes | Pristine Gloves (253913, -0.24 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.26 DPS) [crafted]; Serpent Gloves (5970, -0.57 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Keller's Girdle (2911, -0.08 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.11 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.75 DPS, sim-verified) [crafted] |
| legs | Wisdom's Leather Pants (252503) | Leatherworking [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Leggings (253937, -0.13 DPS) [crafted]; Totemic Leather Pants (252446, -0.22 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.75 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.4 spell_power points (1.37 DPS) | yes | Stormrider's Leather Boots (252443, -0.02 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.29 DPS) [crafted]; Black Whelp Slippers (252424, -0.53 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 5.0 spell_power points (0.67 DPS) | yes | Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; The 1 Ring (8350, -0.56 DPS) [world] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | 4.2 spell_power points (0.56 DPS) | yes | The 1 Ring (8350, -0.44 DPS) [world]; Volcanic Rock Ring (12053, -0.63 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 8.4 spell_power points (1.11 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -0.05 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.44 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Wisdom's Leather Pants; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Loop of Sacrifice; main_hand: Gnarled Necromancer's Staff

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 60.9. Weights run: 1.9s. Verify run: 1.0s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.561 ± 0.007, crit=0.035 ± 0.002 per rating point (14 rating = 1%, 0.489 per %), hit=0.128 ± 0.006 per rating point (10 rating = 1%, 1.280 per %), spell_haste=not significant (-0.126 ± 0.095), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.333 ± 0.001, arcane_power=0.667 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.11 DPS) [world_drop]; Silk Headband (7050, -0.47 DPS) [crafted]; Totemic Leather Helm (252456, -1.49 DPS, sim-verified) [crafted] |
| neck | Crystal Starfire Medallion (5003) | World drop [world_drop] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Kaleidoscope Chain (13084, +0.00 DPS) [world_drop]; Archmagister's Faceted Pendant (273049, -0.71 DPS, sim-verified) [world] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (2.51 DPS) | yes | Death Speaker Mantle (6685, -0.34 DPS) [dungeon]; Fairywing Mantle (9536, -0.54 DPS) [quest]; Magician's Mantle (12998, -0.71 DPS) [world_drop] |
| back | Windsong Drape (15468) | Free at Last [quest] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.24 DPS) [crafted]; Hillman's Cloak (3719, -0.66 DPS, sim-verified) [crafted] |
| chest | Beguiler Robes (7728) | Scarlet Monastery [dungeon] | sim-verified (+1.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Mechbuilder's Overalls (9508, -0.18 DPS) [dungeon]; Guardian Armor (4256, -0.20 DPS) [crafted]; Green Silk Armor (7065, -1.51 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.61 DPS) | yes | Nightsky Wristbands (6407, -1.01 DPS) [world_drop]; Technician's Bracers (270042, -1.01 DPS) [quest]; Glowing Magical Bracelets (13106, -1.33 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.1 spell_power points (1.81 DPS) | yes | Jutebraid Gloves (10654, -0.24 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.51 DPS) [crafted]; Serpent Gloves (5970, -0.56 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.4 spell_power points (2.39 DPS) | yes | Moss Cinch (6911, -0.24 DPS) [dungeon]; Warsong Sash (16975, -0.42 DPS) [quest]; Belt of Arugal (6392, -0.48 DPS) [dungeon] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Pants (5962, -0.16 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.24 DPS) [dungeon]; Dark Ritual Leggings (270031, -1.02 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.95 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.38 DPS) [crafted]; Acidic Walkers (9454, -2.00 DPS, sim-verified) [dungeon] |
| finger1 | Snake Hoop (6750) | Willix the Importer [quest] | 8.8 spell_power points (1.57 DPS) | yes | Lavishly Jeweled Ring (1156, -0.97 DPS) [dungeon]; Electrocutioner Lagnut (9447, -1.04 DPS) [dungeon]; Ring of Power Regulation (273030, -1.38 DPS, sim-verified) [dungeon] |
| finger2 | Solid Gold Band (279845) | Opportunistic Education [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Electrocutioner Lagnut (9447, -0.18 DPS) [dungeon]; Ring of Power Regulation (273030, -0.58 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tendonscraper Dagger (273053, +0.00 DPS) [world]; Rhahk'Zor's Hammer (5187, -5.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, -10.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Crystal Starfire Medallion; shoulder: Bloodmage Mantle; back: Windsong Drape; chest: Beguiler Robes; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Stormrider's Leather Pants; feet: Gilded Slippers; finger1: Snake Hoop; finger2: Solid Gold Band; main_hand: Spine of the Devourer

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 5232221115400021-00000000000000000000-0000000000000000)

Set DPS (verified): 81.4. Weights run: 1.9s. Verify run: 1.1s. 514 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.437 ± 0.006, crit=0.065 ± 0.003 per rating point (14 rating = 1%, 0.912 per %), hit=0.167 ± 0.008 per rating point (10 rating = 1%, 1.666 per %), spell_haste=0.279 ± 0.029, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.334 ± 0.001, arcane_power=0.666 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Papal Fez (9431) | Uldaman [dungeon] | sim-verified (+2.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.01 DPS) [world]; Big Voodoo Mask (8201, -0.06 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -2.79 DPS, sim-verified) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 10.5 spell_power points (2.01 DPS) | yes | Necklace of Calisea (1714, -1.43 DPS) [dungeon]; Archmagister's Faceted Pendant (273049, -1.51 DPS) [world]; Triune Amulet (7722, -1.98 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 spell_power points (2.48 DPS) | yes | Green Silken Shoulders (7057, -0.02 DPS) [crafted]; Inquisitor's Shawl (19507, -0.05 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.2 spell_power points (1.57 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Icy Cloak (4327, -0.23 DPS) [crafted]; Hillman's Cloak (3719, -0.61 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.18 DPS) [world_drop]; Robe of Power (7054, -0.52 DPS) [crafted]; Robe of the Magi (1716, -1.25 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.73 DPS) | yes | Guardian Leather Bracers (4260, -0.07 DPS) [crafted]; Radiant Silver Bracers (4545, -0.29 DPS) [quest]; Condor Bracers (15864, -0.38 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.60 DPS) | yes | Dreamweave Gloves (10019, -1.42 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -1.65 DPS) [crafted]; Black Mageweave Gloves (10003, -1.73 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 spell_power points (2.60 DPS) | yes | Skycaller's Leather Belt (252522, -0.18 DPS) [crafted]; Moss Cinch (6911, -0.30 DPS) [dungeon]; Star Belt (4329, -0.98 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.2 spell_power points (3.69 DPS) | yes | Crimson Silk Pantaloons (7062, -1.07 DPS) [crafted]; Stormrider's Leather Pants (252502, -1.27 DPS) [crafted]; Dark Ritual Leggings (270031, -3.03 DPS, sim-verified) [quest] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.60 DPS) | yes | Skycaller's Leather Shoes (252532, -1.56 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.48 DPS) [crafted]; Gilded Slippers (254001, -2.67 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.2 spell_power points (2.34 DPS) | yes | Reedknot Ring (9622, -0.99 DPS) [quest]; Ring of Power Regulation (273030, -1.19 DPS) [dungeon]; Solid Gold Band (279845, -1.57 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.2 spell_power points (1.57 DPS) | yes | Ring of Power Regulation (273030, -0.42 DPS) [dungeon]; Solid Gold Band (279845, -0.80 DPS) [quest]; Reedknot Ring (9622, -1.50 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Hypnotic Blade (7714) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Gut Ripper (2164, +0.00 DPS) [world_drop]; Illusionary Rod (7713, -0.17 DPS) [dungeon]; Spine of the Devourer (273040, -2.21 DPS) [world] |
| off_hand | Orb of the Forgotten Seer (7685) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 spell_power points (2.69 DPS) | yes | Orb of Mystic Insight (249394, -0.84 DPS) [crafted]; Omega Orb (7749, -0.96 DPS) [quest]; Beacon of Hope (9393, -1.15 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 40:** head: Papal Fez; neck: Glowing Eye of Mordresh; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Snake Hoop; main_hand: Hypnotic Blade; off_hand: Orb of the Forgotten Seer

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 129.3. Weights run: 2.1s. Verify run: 1.3s. 684 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.589 ± 0.009, crit=0.096 ± 0.005 per rating point (14 rating = 1%, 1.350 per %), hit=0.288 ± 0.015 per rating point (10 rating = 1%, 2.875 per %), spell_haste=0.566 ± 0.070, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.392 ± 0.001, arcane_power=0.608 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.8 spell_power points (5.73 DPS) | yes | Red Mageweave Headband (10033, -0.19 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.87 DPS) [crafted]; Dreamweave Circlet (10041, -0.89 DPS) [crafted] |
| neck | Glowing Eye of Mordresh (10769) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 11.7 spell_power points (2.11 DPS) | yes | Mindburst Medallion (11196, -0.39 DPS) [quest]; Horizon Choker (13085, -0.62 DPS) [world_drop]; Triune Amulet (7722, -0.67 DPS) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 26.8 spell_power points (4.82 DPS) | yes | Blood Guard's Crackling Leather Spaulders (220871, -1.08 DPS) [vendor]; Skycaller's Leather Shoulders (252537, -1.53 DPS) [crafted]; Rotgrip Mantle (17732, -2.07 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.5 spell_power points (3.16 DPS) | yes | Deep Woodlands Cloak (19121, -0.04 DPS) [quest]; Runecloth Cloak (13860, -0.69 DPS) [crafted]; Big Voodoo Cloak (8216, -1.30 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 30.8 spell_power points (5.54 DPS) | yes | Robe of the Magi (1716, -0.94 DPS) [world_drop]; Feathered Breastplate (8349, -1.06 DPS) [crafted]; Runecloth Tunic (13857, -1.31 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 15.9 spell_power points (2.86 DPS) | yes | Skycaller's Leather Bracers (252542, -0.32 DPS) [crafted]; Nethergeld Cuffs (254061, -0.86 DPS) [crafted]; Bloodband Bracers (11469, -1.01 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.32 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.49 DPS) [crafted]; First Sergeant's Crackling Leather Gauntlets (220867, -0.63 DPS) [vendor]; Dreamweave Gloves (10019, -0.66 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 23.1 spell_power points (4.15 DPS) | yes | Satyrmane Sash (17755, -0.57 DPS) [dungeon]; Dawnspire Cord (12466, -1.06 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.08 DPS) [crafted] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | PvP rank 10 · Stone Guard · Horde [vendor] | 22.1 spell_power points (3.98 DPS) | yes | Big Voodoo Pants (8202, -0.22 DPS) [crafted]; Wizardweave Leggings (14132, -0.56 DPS) [crafted]; Red Mageweave Pants (10009, -1.26 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.32 DPS) | yes | Skycaller's Leather Boots (252471, -0.45 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -0.76 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.24 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.34 DPS) | yes | Snake Hoop (6750, -0.73 DPS) [quest]; Reedknot Ring (9622, -1.08 DPS) [quest]; Ring of Power Regulation (273030, -1.26 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.9 spell_power points (2.33 DPS) | yes | Snake Hoop (6750, -0.72 DPS) [quest]; Reedknot Ring (9622, -1.07 DPS) [quest]; Ring of Power Regulation (273030, -1.25 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (129.3 DPS) | yes | Rune of the Guard Captain (19120, -0.72 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (129.3 DPS) | yes | Rune of the Guard Captain (19120, -0.10 DPS) [quest] |
| main_hand | Spellshifter Rod (9527) | Tiara of the Deep [quest] | sim-verified (129.3 DPS) | yes | Glowing Brightwood Staff (812, -0.08 DPS) [world_drop]; Kindling Stave (11750, -0.09 DPS) [dungeon]; Hammer of the Northern Wind (810, -22.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Crackling Leather Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Frozen Heart of the Mountain; main_hand: Spellshifter Rod

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 169.2. Weights run: 2.0s. Verify run: 3.9s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.650 ± 0.012, crit=0.127 ± 0.007 per rating point (14 rating = 1%, 1.777 per %), hit=0.374 ± 0.018 per rating point (10 rating = 1%, 3.735 per %), spell_haste=not significant (0.268 ± 0.074), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.390 ± 0.001, arcane_power=0.610 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 36.7 spell_power points (7.09 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.43 DPS) [pvp]; Soothsayer's Headdress (17740, -8.85 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 18.5 spell_power points (3.57 DPS) | yes | Glowing Eye of Mordresh (10769, -1.22 DPS) [dungeon]; Mindburst Medallion (11196, -1.66 DPS) [quest]; Horizon Choker (13085, -1.82 DPS) [world_drop] |
| shoulder | Devout Mantle (16695) | Blackrock Spire: Solakar Flamewreath [dungeon] | sim-verified (169.2 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Ironfeather Shoulders (15067, -5.41 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.9 spell_power points (3.46 DPS) | yes | Deep Woodlands Cloak (19121, -0.01 DPS) [quest]; Runecloth Cloak (13860, -0.71 DPS) [crafted]; Big Voodoo Cloak (8216, -1.36 DPS) [crafted] |
| chest | Devout Robe (16690) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (169.2 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Ironfeather Breastplate (15066, -4.08 DPS, sim-verified) [crafted] |
| wrist | Devout Bracers (16697) | Stratholme: Crimson Priest [dungeon] | sim-verified (169.2 DPS) | yes | Runic Leather Bracers (15092, +0.00 DPS) [crafted]; Skycaller's Leather Bracers (252542, +0.00 DPS) [crafted]; Runecloth Cuffs (254123, -7.96 DPS, sim-verified) [crafted] |
| hands | Devout Gloves (16692) | Stratholme: Archivist Galford [dungeon] | sim-verified (169.2 DPS) | yes | Blood Guard's Dragonhide Gloves (227187, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Gloves of the Greatfather (17721, -5.82 DPS, sim-verified) [crafted] |
| waist | Devout Belt (16696) | Blackrock Spire: Blackhand Summoner [dungeon] | sim-verified (169.2 DPS) | yes | Magister's Belt (16685, +0.00 DPS) [dungeon]; Satyrmane Sash (17755, +0.00 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, -7.83 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 51.5 spell_power points (9.95 DPS) | yes | General's Dragonhide Pants (231680, -2.63 DPS) [vendor]; Luminary Kilt (11823, -3.20 DPS, sim-verified) [dungeon]; Legionnaire's Dragonhide Pants (227185, -3.41 DPS) [vendor] |
| feet | Omnicast Boots (11822) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 27.8 spell_power points (5.37 DPS) | yes | General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Earthen Silk Slippers (254013, +0.00 DPS) [crafted]; Blood Guard's Dragonhide Boots (227188, -0.06 DPS) [pvp] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 22.6 spell_power points (4.36 DPS) | yes | Band of the Unicorn (7553, -1.85 DPS) [world_drop]; Snake Hoop (6750, -2.57 DPS) [quest]; Glowing Crystal Ring (18402, -2.85 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (2.56 DPS) | yes | Band of the Unicorn (7553, -0.05 DPS) [world_drop]; Snake Hoop (6750, -0.77 DPS) [quest]; Glowing Crystal Ring (18402, -1.05 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (169.2 DPS) | yes | Rune of the Guard Captain (19120, -0.65 DPS) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (169.2 DPS) | yes | Rune of the Guard Captain (19120, -0.14 DPS) [quest] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (169.2 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Spellshifter Rod (9527, -5.70 DPS) [quest] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (4.06 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Truesilver Conduit (249455, -1.12 DPS) [crafted] |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Archlight Talisman; shoulder: Devout Mantle; chest: Devout Robe; wrist: Devout Bracers; hands: Devout Gloves; waist: Devout Belt; legs: Ironfeather Leggings; feet: Omnicast Boots; finger1: Maiden's Circle; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 5232220115501351-00000000000000000000-5053000000000000)

Set DPS (verified): 477.7. Weights run: 1.4s. Verify run: 3.4s. 1564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.323 ± 0.013, crit=0.226 ± 0.011 per rating point (14 rating = 1%, 3.168 per %), hit=0.565 ± 0.031 per rating point (10 rating = 1%, 5.647 per %), spell_haste=2.298 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.322 ± 0.000, arcane_power=0.678 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | 30.8 spell_power points (16.05 DPS) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.43 DPS) [pvp]; Soothsayer's Headdress (17740, -15.10 DPS, sim-verified) [dungeon] |
| neck | Archlight Talisman (15856) | Hidden Treasures [quest] | 15.2 spell_power points (7.93 DPS) | yes | Glowing Eye of Mordresh (10769, -2.94 DPS) [dungeon]; Triune Amulet (7722, -3.77 DPS) [dungeon]; Mindburst Medallion (11196, -3.80 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | sim-verified (477.7 DPS) | yes | Mantle of the Timbermaw (19050, +0.00 DPS) [crafted]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Argent Shoulders (19059, -7.19 DPS, sim-verified) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.9 spell_power points (8.30 DPS) | yes | Deep Woodlands Cloak (19121, -0.54 DPS) [quest]; Runecloth Cloak (13860, -2.27 DPS) [crafted]; Nightfall Drape (12465, -3.61 DPS) [dungeon] |
| chest | Ironfeather Breastplate (15066) | Leatherworking [crafted] | sim-verified (477.7 DPS) | yes | Legionnaire's Dragonhide Armor (227189, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Fernpulse Jerkin (15786, -9.70 DPS, sim-verified) [quest] |
| wrist | Runecloth Cuffs (254123) | Tailoring [crafted] | 14.2 spell_power points (7.41 DPS) | yes | Runic Leather Bracers (15092, -0.52 DPS) [crafted]; Skycaller's Leather Bracers (252542, -1.03 DPS) [crafted]; Nethergeld Cuffs (254061, -2.59 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (12.50 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Blood Guard's Dragonhide Gloves (227187, -1.44 DPS) [vendor]; Earth Warder's Gloves (21318, -6.04 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 19.9 spell_power points (10.35 DPS) | yes | Satyrmane Sash (17755, -1.38 DPS) [dungeon]; Wisdom of the Timbermaw (19047, -2.27 DPS) [crafted]; Ghostweave Cord (254073, -3.06 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 47.0 spell_power points (24.47 DPS) | yes | General's Dragonhide Pants (231680, -6.40 DPS) [vendor]; Legionnaire's Dragonhide Pants (227185, -9.02 DPS) [vendor]; Luminary Kilt (11823, -10.92 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (12.50 DPS) | yes | Blood Guard's Dragonhide Boots (227188, +0.00 DPS) [pvp]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; Omnicast Boots (11822, -0.06 DPS) [dungeon] |
| finger1 | Maiden's Circle (13001) | World drop [world_drop] | 20.3 spell_power points (10.55 DPS) | yes | Philanthropist's Ring (281635, -4.50 DPS) [quest]; Snake Hoop (6750, -6.59 DPS) [quest]; Reedknot Ring (9622, -6.91 DPS) [quest] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (6.77 DPS) | yes | Philanthropist's Ring (281635, -0.72 DPS) [quest]; Snake Hoop (6750, -2.81 DPS) [quest]; Reedknot Ring (9622, -3.13 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (477.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (477.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Hand of Edward the Odd (2243) | World drop [world_drop] | sim-verified (477.7 DPS) | yes | Masterwork Stormhammer (12794, +0.00 DPS) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Kindling Stave (11750, -15.80 DPS) [dungeon] |
| off_hand | Spirit of Aquementas (11904) | It's Dangerous to Go Alone [quest] | 21.0 spell_power points (10.94 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, -0.77 DPS) [pvp]; Orb of the Forgotten Seer (7685, -3.96 DPS, sim-verified) [dungeon] |
| ranged | Talons of Wrath (249441) | Enchanting [crafted] | sim-verified (477.7 DPS) | yes | Idol of the Moon (23197, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Archlight Talisman; chest: Ironfeather Breastplate; wrist: Runecloth Cuffs; legs: Ironfeather Leggings; finger1: Maiden's Circle; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Hand of Edward the Odd; off_hand: Spirit of Aquementas

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

