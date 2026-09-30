# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 29.2. Weights run: 1.2s. Verify run: 0.7s. 194 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.245 ± 0.012, crit=0.042 ± 0.003 per rating point (14 rating = 1%, 0.582 per %), hit=0.132 ± 0.002 per rating point (10 rating = 1%, 1.323 per %), spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Goggles (4373, -0.28 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.31 DPS) [crafted]; Trapper's Leather Hood (252505, -1.26 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.2 spell_power points (1.66 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.10 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.25 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.7 spell_power points (0.59 DPS) | yes | Sanguine Cape (14376, +0.00 DPS, sim-verified) [world_drop]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Caretaker's Cape (20428, -0.28 DPS) [rep] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Gray Woolen Robe (2585, -0.10 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.33 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.13 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.13 DPS) [vendor]; Owl Bracers (4796, -1.26 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 10.0 spell_power points (1.02 DPS) | yes | Wisdom's Leather Gloves (252499, +0.22 DPS, sim-verified) [crafted]; Windfelt Gloves (5630, -0.23 DPS) [quest]; Pristine Gloves (253913, -0.23 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.68 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.0 spell_power points (1.95 DPS) | yes | Dreamer's Leggings (270016, +0.15 DPS, sim-verified) [quest]; Stormrider's Leather Pants (252502, -0.15 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.46 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.2 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, +0.07 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.21 DPS) [crafted]; Black Whelp Slippers (252424, -0.44 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.5 spell_power points (0.77 DPS) | yes | Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Sludge-Stained Band (286535, -0.46 DPS) [world]; Lavishly Jeweled Ring (1156, -1.01 DPS, sim-verified) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; Sludge-Stained Band (286535, -0.21 DPS) [world]; Lavishly Jeweled Ring (1156, -1.09 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 12.5 spell_power points (1.28 DPS) | yes | Channeler's Staff (4437, -0.20 DPS, sim-verified) [world]; Rhahk'Zor's Hammer (5187, -0.46 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.51 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 194, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14148 Crystalline Cuffs

### Band 30 (night-elf, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 52.0. Weights run: 1.3s. Verify run: 0.8s. 324 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.842 ± 0.015, crit=0.063 ± 0.005 per rating point (14 rating = 1%, 0.886 per %), hit=0.175 ± 0.003 per rating point (10 rating = 1%, 1.753 per %), spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Totemic Leather Helm (252456, -0.20 DPS) [crafted]; Holy Shroud (2721, -0.33 DPS) [world_drop]; Enchanter's Cowl (4322, -0.55 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.61 DPS) | yes | Crystal Starfire Medallion (5003, -1.16 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.33 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.6 spell_power points (2.22 DPS) | yes | Death Speaker Mantle (6685, -0.33 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Magician's Mantle (12998, -0.54 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.7 spell_power points (0.90 DPS) | yes | Darkspear Raider's Cloak (272078, +0.46 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.05 DPS) [quest]; Resilient Cape (14400, -0.23 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 20.0 spell_power points (2.67 DPS) | yes | Guardian Armor (4256, -0.33 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.49 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.66 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.20 DPS) | yes | Nightsky Wristbands (6407, -0.53 DPS) [world_drop]; Technician's Bracers (270042, -0.53 DPS) [quest]; Glowing Magical Bracelets (13106, -1.26 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 13.3 spell_power points (1.78 DPS) | yes | Truefaith Gloves (7049, -0.77 DPS) [crafted]; Gloves of Insight (9698, -0.84 DPS) [quest]; Stormrider's Leather Gloves (252498, -1.14 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.1 spell_power points (2.02 DPS) | yes | Moss Cinch (6911, -0.41 DPS) [dungeon]; Crimson Silk Belt (7055, -0.42 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.57 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 19.3 spell_power points (2.58 DPS) | yes | Abomination Skin Leggings (23173, -0.50 DPS, sim-verified) [dungeon]; Stormrider's Leather Pants (252502, -0.56 DPS) [crafted]; Guardian Pants (5962, -0.61 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.9 spell_power points (1.73 DPS) | yes | Spidersilk Boots (4320, -0.34 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.36 DPS) [crafted]; Acidic Walkers (9454, -1.38 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.94 DPS) | yes | Black Widow Band (6199, -0.15 DPS) [world]; Snake Hoop (6750, -0.15 DPS) [quest]; Minor Channeling Ring (1449, -1.66 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.01 DPS) [world]; Snake Hoop (6750, -0.01 DPS) [quest]; Minor Channeling Ring (1449, -0.93 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 0.0 spell_power points (0.00 DPS) | yes | Glimmering Staff (249392, -1.87 DPS) [crafted]; Scorn's Focal Dagger (23168, -1.90 DPS) [dungeon]; Manual Crowd Pummeler (9449, -3.87 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 324, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 70.7. Weights run: 1.4s. Verify run: 0.8s. 444 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.750 ± 0.016, crit=0.088 ± 0.010 per rating point (14 rating = 1%, 1.236 per %), hit=0.273 ± 0.004 per rating point (10 rating = 1%, 2.733 per %), spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.68 DPS) | yes | Big Voodoo Mask (8201, +0.19 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.32 DPS) [world]; Corpseshroud (10574, -0.86 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.5 spell_power points (1.47 DPS) | yes | Necklace of Calisea (1714, -0.80 DPS) [world_drop]; Triune Amulet (7722, -0.80 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.31 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.7 spell_power points (2.14 DPS) | yes | Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -0.53 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.19 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.33 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.22 DPS) [crafted]; Elemental Raiment (9434, -0.48 DPS) [world_drop]; Robe of the Magi (1716, -1.03 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.5 spell_power points (1.34 DPS) | yes | Spidertank Oilrag (9448, -0.19 DPS) [dungeon]; Condor Bracers (15864, -0.45 DPS) [quest]; Arcane Runed Bracers (4744, -0.78 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Dreamweave Gloves (10019, -0.69 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.70 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.80 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.2 spell_power points (2.33 DPS) | yes | Skycaller's Leather Belt (252522, -0.48 DPS) [crafted]; Gilded Cord (254037, -0.54 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.73 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.0 spell_power points (2.93 DPS) | yes | Crimson Silk Pantaloons (7062, -0.67 DPS) [crafted]; Kodohide Legguards (285338, -0.81 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.02 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Skycaller's Leather Shoes (252532, -0.94 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.37 DPS) [crafted]; Gilded Slippers (254001, -1.50 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.85 DPS) | yes | Ring of Forlorn Spirits (2043, -0.83 DPS) [quest]; Reedknot Ring (9622, -0.96 DPS) [quest]; Minor Channeling Ring (1449, -1.02 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.15 DPS) | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.20 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 0.0 spell_power points (0.00 DPS) | yes | Spellforce Rod (1664, -0.35 DPS) [world_drop]; Mograine's Might (7723, -1.37 DPS) [dungeon]; Manual Crowd Pummeler (9449, -3.77 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (night-elf, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 91.5. Weights run: 1.4s. Verify run: 0.9s. 587 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.018 ± 0.022, crit=0.139 ± 0.017 per rating point (14 rating = 1%, 1.945 per %), hit=0.431 ± 0.007 per rating point (10 rating = 1%, 4.310 per %), spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.81 DPS) [crafted]; Red Mageweave Headband (10033, -1.20 DPS, sim-verified) [crafted]; Chief Architect's Monocle (11839, -1.23 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.13 DPS) [dungeon]; Mindburst Medallion (11196, -0.25 DPS) [quest]; Arcane Crystal Pendant (20037, -2.59 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.4 spell_power points (4.05 DPS) | yes | Kentic Amice (11624, -0.93 DPS) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.02 DPS) [vendor]; Rotgrip Mantle (17732, -1.62 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.1 spell_power points (2.30 DPS) | yes | Runecloth Cloak (13860, -0.34 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.67 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -3.59 DPS, sim-verified) [dungeon] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (+2.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Runecloth Robe (13858, -0.10 DPS) [crafted]; Runecloth Tunic (13857, -0.11 DPS) [crafted]; Acumen Robes (17775, -2.22 DPS, sim-verified) [quest] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.2 spell_power points (2.31 DPS) | yes | Skycaller's Leather Bracers (252542, +0.45 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -0.56 DPS) [dungeon]; Bloodband Bracers (11469, -0.69 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 32.2 spell_power points (3.69 DPS) | yes | Sergeant Major's Crackling Leather Gauntlets (220866, -0.81 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -0.81 DPS) [crafted]; Raider Handwraps (272098, -1.06 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 28.2 spell_power points (3.23 DPS) | yes | Satyrmane Sash (17755, -0.46 DPS) [dungeon]; Ban'thok Sash (11662, -0.53 DPS) [dungeon]; Dawnspire Cord (12466, -0.61 DPS, sim-verified) [dungeon] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | sim-verified (+4.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.25 DPS) [crafted]; Big Voodoo Pants (8202, -0.37 DPS) [crafted]; Spellshock Leggings (9484, -4.20 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Boots (252471) | Leatherworking [crafted] | 26.2 spell_power points (3.00 DPS) | yes | Earthen Silk Slippers (254013, -0.25 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -0.26 DPS, sim-verified) [vendor]; Mender's Leather Boots (252472, -0.69 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.1 spell_power points (1.85 DPS) | yes | Brainlash (6440, -0.10 DPS) [dungeon]; Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Mindseye Circle (10634, -0.45 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (1.84 DPS) | yes | Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Mindseye Circle (10634, -0.45 DPS) [dungeon]; Brainlash (6440, -2.12 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Blessed Prayer Beads (19990, -0.90 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Mechanic's Pipehammer (9604, -0.62 DPS) [quest]; Spellshifter Rod (9527, -0.70 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Knight's Crackling Leather Leggings; feet: Skycaller's Leather Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 587, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (night-elf, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 189.7. Weights run: 1.4s. Verify run: 0.9s. 1395 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.134 ± 0.031, crit=0.180 ± 0.023 per rating point (14 rating = 1%, 2.513 per %), hit=0.553 ± 0.010 per rating point (10 rating = 1%, 5.526 per %), spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 87.0 spell_power points (9.61 DPS) | yes | Waywatcher Headpiece (240088, -1.15 DPS, sim-verified) [vendor]; Feralheart Cowl (226773, -4.20 DPS) [quest]; Field Marshal's Dragonhide Helm (231695, -4.29 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.7 spell_power points (3.29 DPS) | yes | Beads of Ogre Mojo (22149, -0.48 DPS, sim-verified) [quest]; Archlight Talisman (15856, -0.71 DPS) [quest]; Arcane Crystal Pendant (20037, -0.77 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 71.4 spell_power points (7.90 DPS) | yes | Waywatcher Spaulders (240086, -2.83 DPS) [vendor]; Darkspear Shoulderpads (272103, -3.19 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -5.79 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.6 spell_power points (3.38 DPS) | yes | Hide of the Wild (18510, +0.80 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -0.67 DPS) [world_drop]; Spritecaster Cape (11623, -1.08 DPS) [dungeon] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | sim-verified (189.7 DPS) | yes | Waywatcher Tunic (240091, -1.69 DPS) [vendor]; Feralheart Vest (226776, -3.84 DPS) [quest]; Tunic of Undead Slaying (23089, -16.33 DPS, sim-verified) [world] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | sim-verified (189.7 DPS) | yes | Waywatcher Wristguards (240084, -2.21 DPS) [vendor]; Dryad's Wrist Bindings (19595, -2.79 DPS) [rep]; Wristwraps of Undead Slaying (23093, -11.95 DPS, sim-verified) [world] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 74.1 spell_power points (8.19 DPS) | yes | Waywatcher Handguards (240089, -2.41 DPS, sim-verified) [vendor]; Raider Handwraps (272097, -3.92 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -4.26 DPS) [vendor] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 77.1 spell_power points (8.52 DPS) | yes | Waywatcher Girdle (240085, -3.36 DPS) [vendor]; Elunite Cord (272401, -3.74 DPS) [vendor]; Knowledge of the Timbermaw (228190, -4.26 DPS, sim-verified) [vendor] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 95.1 spell_power points (10.52 DPS) | yes | Waywatcher Legguards (240087, -0.08 DPS, sim-verified) [vendor]; Ironfeather Leggings (252486, -3.78 DPS) [crafted]; Sentinel's Silk Leggings (22752, -5.04 DPS) [rep] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 58.0 spell_power points (6.42 DPS) | yes | Waywatcher Boots (240090, +1.25 DPS, sim-verified) [vendor]; Feralheart Galoshes (226774, -1.98 DPS) [vendor]; Marshal's Dragonhide Boots (231698, -2.34 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (189.7 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.35 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234436, -0.39 DPS) [vendor]; Naglering (11669, -9.53 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (189.7 DPS) | yes | Cauterizing Band (19140, -0.65 DPS) [world_drop]; Ritssyn's Ring of Chaos (21836, -0.77 DPS) [world_drop]; Naglering (11669, -6.88 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (189.7 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (189.7 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.22 DPS) [dungeon]; Weakness Analyzer (272438, -1.06 DPS, sim-verified) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (189.7 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Electrified Dagger (19100, -11.99 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Waywatcher Hood; neck: Amulet of the Dawn; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Crackling Staff

No-known-source sample (15 of 1395, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 27.8. Weights run: 1.2s. Verify run: 0.7s. 190 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.245 ± 0.012, crit=0.042 ± 0.003 per rating point (14 rating = 1%, 0.582 per %), hit=0.132 ± 0.002 per rating point (10 rating = 1%, 1.323 per %), spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Goggles (4373, -0.28 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.31 DPS) [crafted]; Trapper's Leather Hood (252505, -1.28 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.2 spell_power points (1.66 DPS) | yes | Reinforced Woolen Shoulders (4315, +0.16 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.25 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.7 spell_power points (0.59 DPS) | yes | Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Sanguine Cape (14376, -0.21 DPS, sim-verified) [world_drop]; Battle Healer's Cloak (20427, -0.28 DPS) [rep] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Gray Woolen Robe (2585, -0.10 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.33 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 7.5 spell_power points (0.77 DPS) | yes | Mindthrust Bracers (1974, -0.13 DPS) [dungeon]; Featherbead Bracers (15452, -0.13 DPS) [quest]; Owl Bracers (4796, -0.18 DPS, sim-verified) [vendor] |
| hands | Blight Gloves (279877) | The New Plague [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Gloves (252499, -0.08 DPS) [crafted]; Pristine Gloves (253913, -0.10 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.46 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormrider's Leather Belt (252432, +0.00 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Keller's Girdle (2911, -1.35 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.0 spell_power points (1.95 DPS) | yes | Stormrider's Leather Pants (252502, +0.22 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.46 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.56 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.2 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.21 DPS) [crafted]; Black Whelp Slippers (252424, -0.44 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 7.5 spell_power points (0.77 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; Sludge-Stained Band (286535, -0.21 DPS) [world]; Loop of Sacrifice (281673, -1.21 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 12.5 spell_power points (1.28 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.26 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.46 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Blight Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 190, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (tauren, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 48.6. Weights run: 1.3s. Verify run: 0.8s. 326 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.842 ± 0.015, crit=0.063 ± 0.005 per rating point (14 rating = 1%, 0.886 per %), hit=0.175 ± 0.003 per rating point (10 rating = 1%, 1.753 per %), spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | sim-verified (48.6 DPS) | yes | Totemic Leather Helm (252456, -0.20 DPS) [crafted]; Holy Shroud (2721, -0.33 DPS) [world_drop]; Enchanter's Cowl (4322, -0.88 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.61 DPS) | yes | Crystal Starfire Medallion (5003, -1.16 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.59 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.6 spell_power points (2.22 DPS) | yes | Death Speaker Mantle (6685, -0.21 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Magician's Mantle (12998, -0.54 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.7 spell_power points (0.90 DPS) | yes | Darkspear Raider's Cloak (272078, +0.23 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.23 DPS) [world_drop]; Hillman's Cloak (3719, -0.23 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 20.0 spell_power points (2.67 DPS) | yes | Death Speaker Robes (6682, -0.49 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.66 DPS) [crafted]; Guardian Armor (4256, -0.77 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.20 DPS) | yes | Nightsky Wristbands (6407, -0.53 DPS) [world_drop]; Technician's Bracers (270042, -0.53 DPS) [quest]; Glowing Magical Bracelets (13106, -0.78 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.7 spell_power points (1.43 DPS) | yes | Jutebraid Gloves (10654, -0.06 DPS, sim-verified) [quest]; Stormrider's Leather Gloves (252498, -0.31 DPS) [crafted]; Truefaith Gloves (7049, -0.42 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.1 spell_power points (2.02 DPS) | yes | Moss Cinch (6911, -0.41 DPS) [dungeon]; Crimson Silk Belt (7055, -0.42 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.00 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 19.3 spell_power points (2.58 DPS) | yes | Stormrider's Leather Pants (252502, -0.56 DPS) [crafted]; Guardian Pants (5962, -0.61 DPS) [crafted]; Abomination Skin Leggings (23173, -0.94 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.9 spell_power points (1.73 DPS) | yes | Spidersilk Boots (4320, -0.34 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.36 DPS) [crafted]; Acidic Walkers (9454, -1.33 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.94 DPS) | yes | Black Widow Band (6199, -0.15 DPS) [world]; Snake Hoop (6750, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.80 DPS) | yes | Black Widow Band (6199, -0.01 DPS) [world]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Snake Hoop (6750, -0.18 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Scorn's Focal Dagger (23168, -0.04 DPS) [dungeon]; Twisted Chanter's Staff (890, -0.11 DPS) [world_drop]; Manual Crowd Pummeler (9449, -0.52 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Glimmering Staff

No-known-source sample (15 of 326, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 66.0. Weights run: 1.4s. Verify run: 0.8s. 446 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.750 ± 0.016, crit=0.088 ± 0.010 per rating point (14 rating = 1%, 1.236 per %), hit=0.273 ± 0.004 per rating point (10 rating = 1%, 2.733 per %), spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Augural Shroud (2620, -0.13 DPS) [world]; Corpseshroud (10574, -0.67 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.76 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.5 spell_power points (1.47 DPS) | yes | Necklace of Calisea (1714, -0.80 DPS) [world_drop]; Triune Amulet (7722, -0.80 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.86 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.7 spell_power points (2.14 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (+1.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.19 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.94 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.22 DPS) [crafted]; Elemental Raiment (9434, -0.48 DPS) [world_drop]; Robe of the Magi (1716, -1.02 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.5 spell_power points (1.34 DPS) | yes | Radiant Silver Bracers (4545, -0.15 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.19 DPS) [dungeon]; Condor Bracers (15864, -0.45 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Dreamweave Gloves (10019, -0.63 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.70 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.80 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.2 spell_power points (2.33 DPS) | yes | Defiler's Cloth Girdle (20166, -0.19 DPS, sim-verified) [rep]; Skycaller's Leather Belt (252522, -0.48 DPS) [crafted]; Gilded Cord (254037, -0.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.0 spell_power points (2.93 DPS) | yes | Crimson Silk Pantaloons (7062, -0.67 DPS) [crafted]; Kodohide Legguards (285338, -0.69 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.02 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Skycaller's Leather Shoes (252532, -0.73 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.37 DPS) [crafted]; Gilded Slippers (254001, -1.50 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.85 DPS) | yes | Reedknot Ring (9622, -0.96 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.08 DPS) [vendor]; Voodoo Band (1996, -1.18 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.15 DPS) | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Reedknot Ring (9622, -1.06 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Manual Crowd Pummeler (9449, +0.05 DPS, sim-verified) [dungeon]; Mograine's Might (7723, -1.02 DPS) [dungeon]; Windweaver Staff (7757, -1.12 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 89.5. Weights run: 1.4s. Verify run: 0.9s. 589 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.018 ± 0.022, crit=0.139 ± 0.017 per rating point (14 rating = 1%, 1.945 per %), hit=0.431 ± 0.007 per rating point (10 rating = 1%, 4.310 per %), spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.81 DPS) [crafted]; Chief Architect's Monocle (11839, -1.23 DPS) [dungeon]; Red Mageweave Headband (10033, -1.30 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Scorn's Icy Choker (23169, -0.13 DPS) [dungeon]; Mindburst Medallion (11196, -0.25 DPS) [quest]; Arcane Crystal Pendant (20037, -2.70 DPS, sim-verified) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.4 spell_power points (4.05 DPS) | yes | Kentic Amice (11624, -0.93 DPS) [dungeon]; Blood Guard's Crackling Leather Spaulders (220871, -1.02 DPS) [vendor]; Rotgrip Mantle (17732, -1.73 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.2 spell_power points (2.42 DPS) | yes | Spritecaster Cape (11623, +0.68 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.34 DPS) [dungeon]; Runecloth Cloak (13860, -0.46 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Runecloth Robe (13858, -0.10 DPS) [crafted]; Runecloth Tunic (13857, -0.11 DPS) [crafted]; Acumen Robes (17775, -2.09 DPS, sim-verified) [quest] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.2 spell_power points (2.31 DPS) | yes | Skycaller's Leather Bracers (252542, +0.34 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -0.56 DPS) [dungeon]; Bloodband Bracers (11469, -0.69 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 32.2 spell_power points (3.69 DPS) | yes | First Sergeant's Crackling Leather Gauntlets (220867, -0.81 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -0.81 DPS) [crafted]; Raider Handwraps (272098, -1.12 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 28.2 spell_power points (3.23 DPS) | yes | Satyrmane Sash (17755, -0.46 DPS) [dungeon]; Dawnspire Cord (12466, -0.52 DPS, sim-verified) [dungeon]; Ban'thok Sash (11662, -0.53 DPS) [dungeon] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | Lady Palanseer [vendor] | sim-verified (+4.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.25 DPS) [crafted]; Big Voodoo Pants (8202, -0.37 DPS) [crafted]; Spellshock Leggings (9484, -4.14 DPS, sim-verified) [dungeon] |
| feet | Skycaller's Leather Boots (252471) | Leatherworking [crafted] | 26.2 spell_power points (3.00 DPS) | yes | First Sergeant's Crackling Leather Boots (220863, +0.14 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -0.25 DPS) [crafted]; Mender's Leather Boots (252472, -0.69 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.1 spell_power points (1.85 DPS) | yes | Brainlash (6440, -0.10 DPS) [dungeon]; Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Mindseye Circle (10634, -0.45 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (1.84 DPS) | yes | Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Mindseye Circle (10634, -0.45 DPS) [dungeon]; Brainlash (6440, -2.25 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, -0.11 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Barman Shanker (12791, +0.00 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -0.70 DPS) [quest]; Thorium Greatmace (250613, -0.75 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Feathered Breastplate; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Crackling Leather Leggings; feet: Skycaller's Leather Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Rune of the Guard Captain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 589, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 174.5. Weights run: 1.4s. Verify run: 0.9s. 1396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.134 ± 0.031, crit=0.180 ± 0.023 per rating point (14 rating = 1%, 2.513 per %), hit=0.553 ± 0.010 per rating point (10 rating = 1%, 5.526 per %), spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 87.0 spell_power points (9.61 DPS) | yes | Waywatcher Headpiece (240088, +0.81 DPS, sim-verified) [vendor]; Feralheart Cowl (226773, -4.20 DPS) [quest]; Warlord's Dragonhide Helm (231678, -4.29 DPS) [vendor] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.7 spell_power points (3.29 DPS) | yes | Beads of Ogre Mojo (22149, +1.00 DPS, sim-verified) [quest]; Archlight Talisman (15856, -0.71 DPS) [quest]; Arcane Crystal Pendant (20037, -0.77 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 71.4 spell_power points (7.90 DPS) | yes | Waywatcher Spaulders (240086, -2.83 DPS) [vendor]; Darkspear Shoulderpads (272103, -3.19 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -5.42 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.6 spell_power points (3.38 DPS) | yes | Hide of the Wild (18510, +1.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -0.67 DPS) [world_drop]; Deep Woodlands Cloak (19121, -0.93 DPS) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | sim-verified (174.5 DPS) | yes | Waywatcher Tunic (240091, -1.69 DPS) [vendor]; Feralheart Vest (226776, -3.84 DPS) [quest]; Tunic of Undead Slaying (23089, -13.98 DPS, sim-verified) [world] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | sim-verified (174.5 DPS) | yes | Waywatcher Wristguards (240084, -2.21 DPS) [vendor]; Dryad's Wrist Bindings (19595, -2.79 DPS) [rep]; Wristwraps of Undead Slaying (23093, -9.80 DPS, sim-verified) [world] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 74.1 spell_power points (8.19 DPS) | yes | Waywatcher Handguards (240089, -0.67 DPS, sim-verified) [vendor]; Raider Handwraps (272097, -3.92 DPS) [vendor]; General's Dragonhide Gloves (231677, -4.26 DPS) [pvp] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 77.1 spell_power points (8.52 DPS) | yes | Knowledge of the Timbermaw (228190, -2.78 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -3.36 DPS) [vendor]; Elunite Cord (272401, -3.74 DPS) [vendor] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 95.1 spell_power points (10.52 DPS) | yes | Waywatcher Legguards (240087, +0.58 DPS, sim-verified) [vendor]; Ironfeather Leggings (252486, -3.78 DPS) [crafted]; Outrider's Silk Leggings (22747, -5.04 DPS) [rep] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 58.0 spell_power points (6.42 DPS) | yes | Waywatcher Boots (240090, +1.74 DPS, sim-verified) [vendor]; Feralheart Galoshes (226774, -1.98 DPS) [vendor]; General's Dragonhide Boots (231682, -2.34 DPS) [pvp] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (174.5 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -0.35 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234436, -0.39 DPS) [vendor]; Naglering (11669, -7.41 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (174.5 DPS) | yes | Cauterizing Band (19140, -0.65 DPS) [world_drop]; Ritssyn's Ring of Chaos (21836, -0.77 DPS) [world_drop]; Naglering (11669, -5.41 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (174.5 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (174.5 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Burst of Knowledge (11832, -0.22 DPS) [dungeon]; Weakness Analyzer (272438, -1.00 DPS, sim-verified) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (174.5 DPS) | yes | Glacial Blade (19099, +0.40 DPS, sim-verified) [rep]; High Warlord's Destroyer (23465, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Waywatcher Hood; neck: Amulet of the Dawn; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Amethyst War Staff

No-known-source sample (15 of 1396, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

