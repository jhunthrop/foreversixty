# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 32.6. Weights run: 1.5s. Verify run: 0.6s. 225 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.690 ± 0.010, crit=0.057 ± 0.002 per rating point (14 rating = 1%, 0.795 per %), hit=0.140 ± 0.001 per rating point (10 rating = 1%, 1.405 per %), spell_haste=-1.811 ± 0.144, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.793 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.43 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.2 spell_power points (1.34 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.86 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.1 spell_power points (0.49 DPS) | yes | Heavy Woolen Cloak (4311, -0.01 DPS) [crafted]; Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.05 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.32 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.08 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.08 DPS) [vendor]; Owl Bracers (4796, -1.27 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.8 spell_power points (0.93 DPS) | yes | Serpent Gloves (5970, -0.09 DPS) [dungeon]; Windfelt Gloves (5630, -0.20 DPS) [quest]; Pristine Gloves (253913, -0.20 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.08 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.12 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.59 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamer's Leggings (270016, -0.30 DPS) [quest]; Wisdom's Leather Pants (252503, -0.36 DPS) [crafted]; Abomination Skin Leggings (23173, -0.99 DPS, sim-verified) [dungeon] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Boots (252444, -0.24 DPS) [crafted]; Spidersilk Boots (4320, -0.31 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.41 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.4 spell_power points (0.76 DPS) | yes | Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon]; Sludge-Stained Band (286535, -0.40 DPS) [world]; Volcanic Rock Ring (12053, -0.52 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.60 DPS) | yes | Sludge-Stained Band (286535, -0.24 DPS) [world]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.62 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Channeler's Staff (4437, -0.17 DPS) [world]; Lesser Staff of the Spire (1300, -0.33 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.52 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Stormrider's Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 56.2. Weights run: 1.8s. Verify run: 0.7s. 366 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.791 ± 0.017, crit=0.089 ± 0.003 per rating point (14 rating = 1%, 1.240 per %), hit=0.208 ± 0.002 per rating point (10 rating = 1%, 2.077 per %), spell_haste=-2.626 ± 0.267, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.786 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.9 spell_power points (1.68 DPS) | yes | Enduring Cap (3020, -0.15 DPS) [world_drop]; Totemic Leather Helm (252456, -0.23 DPS) [crafted]; Holy Shroud (2721, -0.35 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.7 spell_power points (1.42 DPS) | yes | Crystal Starfire Medallion (5003, -1.04 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.04 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.49 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.1 spell_power points (1.95 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.36 DPS) [quest]; Magician's Mantle (12998, -0.48 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.3 spell_power points (0.77 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Repairman's Cape (9605, -0.02 DPS) [quest]; Hillman's Cloak (3719, -0.16 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 19.3 spell_power points (2.34 DPS) | yes | Death Speaker Robes (6682, -0.43 DPS) [dungeon]; Guardian Armor (4256, -0.50 DPS, sim-verified) [crafted]; Stormrider's Leather Tunic (252510, -0.55 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.52 DPS) [world_drop]; Technician's Bracers (270042, -0.52 DPS) [quest]; Glowing Magical Bracelets (13106, -1.63 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 12.7 spell_power points (1.54 DPS) | yes | Truefaith Gloves (7049, -0.65 DPS) [crafted]; Gloves of Insight (9698, -0.69 DPS) [quest]; Stormrider's Leather Gloves (252498, -1.24 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.7 spell_power points (1.79 DPS) | yes | Moss Cinch (6911, -0.33 DPS) [dungeon]; Crimson Silk Belt (7055, -0.39 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.65 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 18.7 spell_power points (2.26 DPS) | yes | Stormrider's Leather Pants (252502, -0.48 DPS) [crafted]; Guardian Pants (5962, -0.53 DPS) [crafted]; Abomination Skin Leggings (23173, -0.69 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.5 spell_power points (1.52 DPS) | yes | Spidersilk Boots (4320, -0.29 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.31 DPS) [crafted]; Acidic Walkers (9454, -1.47 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.85 DPS) | yes | Black Widow Band (6199, -0.18 DPS) [world]; Snake Hoop (6750, -0.18 DPS) [quest]; Minor Channeling Ring (1449, -2.21 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (56.2 DPS) | yes | Black Widow Band (6199, -0.06 DPS) [world]; Snake Hoop (6750, -0.06 DPS) [quest]; Minor Channeling Ring (1449, -1.10 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -1.69 DPS) [dungeon]; Glimmering Staff (249392, -1.73 DPS) [crafted]; Manual Crowd Pummeler (9449, -4.44 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 4532310300103051-000000000000000000-0000000000000000)

Set DPS (verified): 83.5. Weights run: 1.9s. Verify run: 0.7s. 592 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.131 ± 0.030, crit=0.217 ± 0.009 per rating point (14 rating = 1%, 3.034 per %), hit=0.325 ± 0.003 per rating point (10 rating = 1%, 3.251 per %), spell_haste=-3.380 ± 0.328, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.771 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 24.8 spell_power points (2.98 DPS) | yes | Corpseshroud (10574, -0.40 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.46 DPS) [crafted]; Augural Shroud (2620, -0.85 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 spell_power points (1.65 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.30 DPS) [quest]; Triune Amulet (7722, -0.70 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.70 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.7 spell_power points (2.60 DPS) | yes | Green Silken Shoulders (7057, -0.15 DPS) [crafted]; Bloodmage Mantle (7684, -0.30 DPS) [dungeon]; Death Speaker Mantle (6685, -0.39 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 19.2 spell_power points (2.30 DPS) | yes | Long Silken Cloak (4326, -0.90 DPS) [crafted]; Guardian Cloak (5965, -0.90 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.61 DPS, sim-verified) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (83.5 DPS) | yes | Robe of Power (7054, -0.07 DPS) [crafted]; Big Voodoo Robe (8200, -0.40 DPS) [crafted]; Robe of the Magi (1716, -0.91 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 12.8 spell_power points (1.53 DPS) | yes | Turtle Scale Bracers (8198, -0.14 DPS) [crafted]; Windchaser Cuffs (14429, -0.31 DPS) [world_drop]; Green Whelp Bracers (7386, -0.45 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.88 DPS) | yes | Red Mageweave Gloves (10018, -0.20 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.34 DPS) [crafted]; Dreamweave Gloves (10019, -1.26 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.0 spell_power points (2.87 DPS) | yes | Gilded Cord (254037, -0.83 DPS) [crafted]; Highlander's Lizardhide Girdle (20104, -0.84 DPS) [rep]; Highlander's Cloth Girdle (20098, -1.30 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.6 spell_power points (3.31 DPS) | yes | Kodohide Legguards (285338, -0.62 DPS) [world]; Crimson Silk Pantaloons (7062, -1.08 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.14 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.88 DPS) | yes | Skycaller's Leather Shoes (252532, -0.37 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.37 DPS) [crafted]; Mender's Leather Shoes (252533, -0.97 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.8 spell_power points (2.01 DPS) | yes | Ring of Forlorn Spirits (2043, -1.05 DPS) [quest]; Voodoo Band (1996, -1.06 DPS) [world]; Black Widow Band (6199, -1.06 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.08 DPS) | yes | Voodoo Band (1996, -0.13 DPS) [world]; Black Widow Band (6199, -0.13 DPS) [world]; Ring of Forlorn Spirits (2043, -1.51 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, -0.56 DPS) [world]; Mograine's Might (7723, -0.79 DPS) [dungeon]; Gut Ripper (2164, -4.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 4532310300103051-000000000000000000-5500000000000000)

Set DPS (verified): 118.8. Weights run: 1.9s. Verify run: 0.8s. 762 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.364 ± 0.048, crit=0.310 ± 0.014 per rating point (14 rating = 1%, 4.347 per %), hit=0.473 ± 0.005 per rating point (10 rating = 1%, 4.730 per %), spell_haste=not significant (-1.071 ± 0.702), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.762 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 46.3 spell_power points (5.21 DPS) | yes | Soothsayer's Headdress (17740, -0.32 DPS) [dungeon]; Chief Architect's Monocle (11839, -1.06 DPS) [dungeon]; Dreamweave Circlet (10041, -1.31 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 19.1 spell_power points (2.15 DPS) | yes | Scorn's Icy Choker (23169, -0.44 DPS) [dungeon]; Mindburst Medallion (11196, -0.55 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.61 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 42.3 spell_power points (4.76 DPS) | yes | Rotgrip Mantle (17732, -0.53 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -0.92 DPS) [dungeon]; Kentic Amice (11624, -1.19 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 22.2 spell_power points (2.50 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.10 DPS) [dungeon]; Runecloth Cloak (13860, -0.26 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.35 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 46.3 spell_power points (5.21 DPS) | yes | Runecloth Robe (13858, -1.33 DPS, sim-verified) [crafted]; Robes of Insight (940, -1.37 DPS) [world_drop]; Feathered Breastplate (8349, -1.54 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 23.6 spell_power points (2.66 DPS) | yes | Aristocratic Cuffs (12546, -0.36 DPS) [dungeon]; Skycaller's Leather Bracers (252542, -0.46 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.46 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 37.0 spell_power points (4.17 DPS) | yes | Raider Handguards (272102, -0.09 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -0.98 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.98 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 32.4 spell_power points (3.65 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.05 DPS) [dungeon]; Satyrmane Sash (17755, -0.53 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 36.6 spell_power points (4.13 DPS) | yes | Big Voodoo Pants (8202, -0.90 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.02 DPS) [dungeon]; Red Mageweave Pants (10009, -3.03 DPS, sim-verified) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 30.0 spell_power points (3.38 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Greaves of Withering Despair (22240, -0.07 DPS) [dungeon]; Mender's Leather Boots (252472, -0.68 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 20.5 spell_power points (2.31 DPS) | yes | Philanthropist's Ring (281635, -0.26 DPS) [quest]; Mindseye Circle (10634, -0.46 DPS) [dungeon]; Band of the Unicorn (7553, -0.84 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.5 spell_power points (2.09 DPS) | yes | Philanthropist's Ring (281635, -0.04 DPS) [quest]; Mindseye Circle (10634, -0.25 DPS) [dungeon]; Band of the Unicorn (7553, -0.63 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (118.8 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (118.8 DPS) | yes | Smoking Heart of the Mountain (11811, -1.21 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (118.8 DPS) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -0.92 DPS) [quest]; Radiant Staff (249453, -1.54 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 4532310300103051-000000000000000000-5533220000000000)

Set DPS (verified): 193.5. Weights run: 2.0s. Verify run: 0.8s. 1751 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.254 ± 0.046, crit=0.390 ± 0.017 per rating point (14 rating = 1%, 5.464 per %), hit=0.575 ± 0.006 per rating point (10 rating = 1%, 5.752 per %), spell_haste=-4.760 ± 0.594, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.746 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 53.3 spell_power points (6.30 DPS) | yes | Black Dragonscale Helm (252605, -0.03 DPS) [crafted]; Magister's Crown (16686, -0.56 DPS) [dungeon]; Living Crown (252561, -0.68 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 31.3 spell_power points (3.70 DPS) | yes | Beads of Ogre Mojo (22149, -0.38 DPS) [quest]; Archlight Talisman (15856, -0.80 DPS) [quest]; Lady Maye's Pendant (14558, -0.88 DPS) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 51.3 spell_power points (6.06 DPS) | yes | Darkspear Shoulderguards (272958, -0.17 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.64 DPS) [vendor]; Darkspear Shoulders (272104, -0.64 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.8 spell_power points (3.75 DPS) | yes | Hide of the Wild (18510, -0.62 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.80 DPS) [world]; Spritecaster Cape (11623, -1.21 DPS) [dungeon] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Magister's Robes (16688, -0.11 DPS) [dungeon]; Tunic of Undead Slaying (23089, -11.17 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Modest Armguards (18458, -0.89 DPS) [dungeon]; Sublime Wristguards (18497, -0.89 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -7.48 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (193.5 DPS) | yes | Raider Handguards (272102, -0.89 DPS) [vendor]; Hands of Power (13253, -0.98 DPS) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.3 spell_power points (6.89 DPS) | yes | Belt of the Archmage (18405, -1.51 DPS) [crafted]; Girdle of Insight (18504, -1.71 DPS) [crafted]; Stormseeker's Girdle (272399, -3.26 DPS, sim-verified) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 66.0 spell_power points (7.80 DPS) | yes | Red Dragonscale Leggings (252603, -0.70 DPS) [crafted]; Sentinel's Silk Leggings (237815, -1.25 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.25 DPS) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 41.1 spell_power points (4.85 DPS) | yes | Dragonrider Boots (18102, -0.35 DPS) [dungeon]; Omnicast Boots (11822, -0.71 DPS) [dungeon]; Waterspout Boots (18322, -1.01 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.06 DPS) [quest]; Maiden's Circle (13001, -1.06 DPS) [world_drop]; Naglering (11669, -6.93 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.35 DPS) [quest]; Maiden's Circle (13001, -0.35 DPS) [world_drop]; Naglering (11669, -7.46 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+13.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -2.14 DPS, sim-verified) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, -1.16 DPS) [pvp]; Hand of Edward the Odd (2243, -7.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Blue Dragonscale Helm; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Crackling Staff

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 31.7. Weights run: 1.5s. Verify run: 0.6s. 205 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.690 ± 0.010, crit=0.057 ± 0.002 per rating point (14 rating = 1%, 0.795 per %), hit=0.140 ± 0.001 per rating point (10 rating = 1%, 1.405 per %), spell_haste=-1.811 ± 0.144, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.793 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.06 DPS) [crafted]; Totemic Leather Hood (252448, -0.44 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.2 spell_power points (1.34 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.86 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.89 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Feyscale Cloak (6632, -0.12 DPS) [dungeon]; Black Whelp Cloak (7283, -0.12 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.31 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.05 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.33 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) | Bernard Brubaker [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Mindthrust Bracers (1974, +0.00 DPS) [dungeon]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.80 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.8 spell_power points (0.93 DPS) | yes | Pristine Gloves (253913, -0.20 DPS) [crafted]; Gnoll Casting Gloves (892, -0.21 DPS) [world]; Serpent Gloves (5970, -0.38 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.08 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.12 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.60 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Pants (252503, -0.36 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted]; Abomination Skin Leggings (23173, -0.58 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.8 spell_power points (1.17 DPS) | yes | Stormrider's Leather Boots (252443, -0.04 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.28 DPS) [crafted]; Totemic Leather Boots (252442, -0.45 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.60 DPS) | yes | Sludge-Stained Band (286535, -0.24 DPS) [world]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.55 DPS, sim-verified) [dungeon] |
| finger2 | Loop of Sacrifice (281673) | The Offering of Blood [quest] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Sludge-Stained Band (286535, -0.05 DPS) [world]; Volcanic Rock Ring (12053, -0.17 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.78 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.96 DPS) | yes | Twisted Chanter's Staff (890, -0.13 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.13 DPS) [quest]; Channeler's Staff (4437, -0.30 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Loop of Sacrifice; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 52.8. Weights run: 1.8s. Verify run: 0.7s. 349 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.791 ± 0.017, crit=0.089 ± 0.003 per rating point (14 rating = 1%, 1.240 per %), hit=0.208 ± 0.002 per rating point (10 rating = 1%, 2.077 per %), spell_haste=-2.626 ± 0.267, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.786 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 13.9 spell_power points (1.68 DPS) | yes | Enduring Cap (3020, +0.00 DPS, sim-verified) [world_drop]; Totemic Leather Helm (252456, -0.23 DPS) [crafted]; Holy Shroud (2721, -0.35 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.7 spell_power points (1.42 DPS) | yes | Crystal Starfire Medallion (5003, -1.04 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.04 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.52 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.1 spell_power points (1.95 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.36 DPS) [quest]; Magician's Mantle (12998, -0.48 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.3 spell_power points (0.77 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Hillman's Cloak (3719, -0.16 DPS) [crafted]; Windsong Drape (15468, -0.16 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 19.3 spell_power points (2.34 DPS) | yes | Death Speaker Robes (6682, -0.43 DPS) [dungeon]; Guardian Armor (4256, -0.45 DPS, sim-verified) [crafted]; Stormrider's Leather Tunic (252510, -0.55 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.09 DPS) | yes | Nightsky Wristbands (6407, -0.52 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.52 DPS) [quest]; Glowing Magical Bracelets (13106, -1.19 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.6 spell_power points (1.28 DPS) | yes | Jutebraid Gloves (10654, -0.08 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.29 DPS) [crafted]; Truefaith Gloves (7049, -0.39 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.7 spell_power points (1.79 DPS) | yes | Moss Cinch (6911, -0.33 DPS) [dungeon]; Crimson Silk Belt (7055, -0.39 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.56 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 18.7 spell_power points (2.26 DPS) | yes | Stormrider's Leather Pants (252502, -0.48 DPS) [crafted]; Guardian Pants (5962, -0.53 DPS) [crafted]; Abomination Skin Leggings (23173, -0.63 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.5 spell_power points (1.52 DPS) | yes | Spidersilk Boots (4320, -0.29 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.31 DPS) [crafted]; Acidic Walkers (9454, -1.05 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.85 DPS) | yes | Black Widow Band (6199, -0.18 DPS) [world]; Snake Hoop (6750, -0.18 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.73 DPS) | yes | Snake Hoop (6750, -0.06 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.15 DPS) [dungeon]; Black Widow Band (6199, -0.93 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (52.8 DPS) | yes | Glimmering Staff (249392, -0.04 DPS) [crafted]; Rhahk'Zor's Hammer (5187, -0.12 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.24 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.7 spell_power points (1.42 DPS) | yes | Witch's Finger (16887, -0.74 DPS, sim-verified) [quest]; Tome of the Darkspear Prophecy (272090, -0.80 DPS) [vendor]; Seedcloud Buckler (6630, -0.82 DPS) [dungeon] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 4532310300103051-000000000000000000-0000000000000000)

Set DPS (verified): 83.7. Weights run: 1.9s. Verify run: 0.7s. 555 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.131 ± 0.030, crit=0.217 ± 0.009 per rating point (14 rating = 1%, 3.034 per %), hit=0.325 ± 0.003 per rating point (10 rating = 1%, 3.251 per %), spell_haste=-3.380 ± 0.328, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.771 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 24.8 spell_power points (2.98 DPS) | yes | Augural Shroud (2620, -0.30 DPS) [world]; Corpseshroud (10574, -0.40 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.46 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 spell_power points (1.65 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.30 DPS) [quest]; Triune Amulet (7722, -0.70 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.70 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.7 spell_power points (2.60 DPS) | yes | Green Silken Shoulders (7057, -0.15 DPS) [crafted]; Bloodmage Mantle (7684, -0.30 DPS) [dungeon]; Death Speaker Mantle (6685, -0.39 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 19.2 spell_power points (2.30 DPS) | yes | Long Silken Cloak (4326, -0.90 DPS) [crafted]; Guardian Cloak (5965, -0.90 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.18 DPS, sim-verified) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (83.7 DPS) | yes | Robe of Power (7054, -0.07 DPS) [crafted]; Big Voodoo Robe (8200, -0.40 DPS) [crafted]; Robe of the Magi (1716, -0.83 DPS, sim-verified) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 13.0 spell_power points (1.57 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Guardian Leather Bracers (4260, -0.03 DPS) [crafted]; Turtle Scale Bracers (8198, -0.17 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.88 DPS) | yes | Red Mageweave Gloves (10018, -0.20 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.34 DPS) [crafted]; Dreamweave Gloves (10019, -1.93 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.0 spell_power points (2.87 DPS) | yes | Gilded Cord (254037, -0.83 DPS) [crafted]; Highlander's Mail Girdle (20119, -0.84 DPS) [vendor]; Defiler's Cloth Girdle (20166, -1.96 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.6 spell_power points (3.31 DPS) | yes | Kodohide Legguards (285338, -0.62 DPS) [world]; Crimson Silk Pantaloons (7062, -0.72 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.14 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.88 DPS) | yes | Skycaller's Mail Boots (252563, -0.37 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.83 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -0.97 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.8 spell_power points (2.01 DPS) | yes | Ogremind Ring (1993, -1.06 DPS) [world_drop]; Voodoo Band (1996, -1.06 DPS) [world]; Black Widow Band (6199, -1.06 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.08 DPS) | yes | Ogremind Ring (1993, -0.13 DPS) [world_drop]; Voodoo Band (1996, -0.13 DPS) [world]; Black Widow Band (6199, -1.68 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.23 DPS) [dungeon]; Windweaver Staff (7757, -0.36 DPS) [dungeon]; Gut Ripper (2164, -4.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Radiant Silver Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 4532310300103051-000000000000000000-5500000000000000)

Set DPS (verified): 118.8. Weights run: 1.9s. Verify run: 0.8s. 704 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.364 ± 0.048, crit=0.310 ± 0.014 per rating point (14 rating = 1%, 4.347 per %), hit=0.473 ± 0.005 per rating point (10 rating = 1%, 4.730 per %), spell_haste=not significant (-1.071 ± 0.702), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.762 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Chief Architect's Monocle (11839, -0.75 DPS) [dungeon]; Dreamweave Circlet (10041, -0.99 DPS) [crafted]; Red Mageweave Headband (10033, -1.23 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 19.1 spell_power points (2.15 DPS) | yes | Scorn's Icy Choker (23169, -0.44 DPS) [dungeon]; Mindburst Medallion (11196, -0.55 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.61 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 42.3 spell_power points (4.76 DPS) | yes | Rotgrip Mantle (17732, -0.53 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -0.92 DPS) [dungeon]; Kentic Amice (11624, -1.19 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 24.3 spell_power points (2.74 DPS) | yes | Spritecaster Cape (11623, -0.24 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.34 DPS) [dungeon]; Runecloth Cloak (13860, -0.49 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 46.3 spell_power points (5.21 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.81 DPS) [vendor]; Runecloth Robe (13858, -1.36 DPS) [crafted]; Robes of Insight (940, -1.37 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 23.6 spell_power points (2.66 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Aristocratic Cuffs (12546, -0.36 DPS) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 37.0 spell_power points (4.17 DPS) | yes | Raider Handguards (272102, -0.09 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -0.98 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.98 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 32.4 spell_power points (3.65 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.05 DPS) [dungeon]; Satyrmane Sash (17755, -0.53 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.71 DPS) [crafted]; Big Voodoo Pants (8202, -0.90 DPS) [crafted]; Stone Guard's Pulsing Legplates (220847, -1.18 DPS, sim-verified) [vendor] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 30.0 spell_power points (3.38 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Greaves of Withering Despair (22240, -0.07 DPS) [dungeon]; Mender's Leather Boots (252472, -0.68 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 20.5 spell_power points (2.31 DPS) | yes | Philanthropist's Ring (281635, -0.26 DPS) [quest]; Mindseye Circle (10634, -0.46 DPS) [dungeon]; Band of the Unicorn (7553, -0.84 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 18.5 spell_power points (2.09 DPS) | yes | Philanthropist's Ring (281635, -0.04 DPS) [quest]; Mindseye Circle (10634, -0.25 DPS) [dungeon]; Band of the Unicorn (7553, -0.63 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Spellshifter Rod (9527, -0.92 DPS) [quest]; Radiant Staff (249453, -1.54 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 4532310300103051-000000000000000000-5533220000000000)

Set DPS (verified): 190.9. Weights run: 2.0s. Verify run: 0.7s. 1672 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.254 ± 0.046, crit=0.390 ± 0.017 per rating point (14 rating = 1%, 5.464 per %), hit=0.575 ± 0.006 per rating point (10 rating = 1%, 5.752 per %), spell_haste=-4.760 ± 0.594, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.746 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Coif of The Five Thunders (227002) | Saving the Best for Last [quest] | 57.8 spell_power points (6.83 DPS) | yes | Warlord's Mail Helm (231663, -0.50 DPS) [pvp]; Blue Dragonscale Helm (252604, -0.53 DPS) [crafted]; Black Dragonscale Helm (252605, -0.56 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 31.3 spell_power points (3.70 DPS) | yes | Beads of Ogre Mojo (22149, -0.38 DPS) [quest]; Archlight Talisman (15856, -0.80 DPS) [quest]; Lady Maye's Pendant (14558, -0.88 DPS) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 51.3 spell_power points (6.06 DPS) | yes | Darkspear Shoulderguards (272958, -0.17 DPS) [vendor]; Warlord's Mail Spaulders (231659, -0.56 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.64 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 31.8 spell_power points (3.75 DPS) | yes | Hide of the Wild (18510, -0.62 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.80 DPS) [world]; Deep Woodlands Cloak (19121, -1.00 DPS) [quest] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (190.9 DPS) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Tunic of Undead Slaying (23089, -11.33 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-verified (190.9 DPS) | yes | Modest Armguards (18458, -0.89 DPS) [dungeon]; Sublime Wristguards (18497, -0.89 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.36 DPS, sim-verified) [world] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | 42.3 spell_power points (5.00 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Raider Handwraps (272097, -0.06 DPS) [vendor]; General's Mail Gloves (231666, -0.48 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.3 spell_power points (6.89 DPS) | yes | Belt of the Archmage (18405, -1.51 DPS) [crafted]; Girdle of Insight (18504, -1.71 DPS) [crafted]; Stormseeker's Girdle (272399, -4.75 DPS, sim-verified) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 66.0 spell_power points (7.80 DPS) | yes | General's Mail Leggings (231664, -0.29 DPS) [pvp]; Red Dragonscale Leggings (252603, -0.70 DPS) [crafted]; Sentinel's Silk Leggings (237815, -1.25 DPS) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 41.1 spell_power points (4.85 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -0.35 DPS) [dungeon]; Omnicast Boots (11822, -0.71 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (190.9 DPS) | yes | Eye of Orgrimmar (12545, -1.06 DPS) [quest]; Maiden's Circle (13001, -1.06 DPS) [world_drop]; Naglering (11669, -8.21 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (190.9 DPS) | yes | Eye of Orgrimmar (12545, -0.35 DPS) [quest]; Maiden's Circle (13001, -0.35 DPS) [world_drop]; Naglering (11669, -9.17 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (190.9 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (190.9 DPS) | yes | Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Briarwood Reed (12930, -2.67 DPS, sim-verified) [dungeon] |
| main_hand | Hammer of Divine Might (22333) | Scholomance: Kormok [dungeon] | sim-verified (190.9 DPS) | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -7.86 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Coif of The Five Thunders; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Hammer of Divine Might

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

