# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 26.4. Weights run: 1.6s. Verify run: 1.1s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.723 ± 0.018, crit=0.059 ± 0.003 per rating point (14 rating = 1%, 0.822 per %), hit=0.153 ± 0.002 per rating point (10 rating = 1%, 1.532 per %), spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.02 DPS) [crafted]; Totemic Leather Hood (252448, -0.32 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 spell_power points (1.10 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.26 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.72 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.26 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 11.6 spell_power points (1.11 DPS) | yes | Wisdom's Leather Armor (252493, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Gown (253901, -0.29 DPS) [crafted]; Totemic Leather Armor (252435, -0.33 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.07 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.07 DPS) [vendor]; Owl Bracers (4796, -1.32 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.9 spell_power points (0.76 DPS) | yes | Serpent Gloves (5970, -0.08 DPS, sim-verified) [dungeon]; Windfelt Gloves (5630, -0.17 DPS) [quest]; Pristine Gloves (253913, -0.17 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.44 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.8 spell_power points (1.42 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -0.27 DPS) [quest]; Wisdom's Leather Pants (252503, -0.33 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (0.95 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.22 DPS) [crafted]; Totemic Leather Boots (252442, -0.37 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.4 spell_power points (0.62 DPS) | yes | Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; Sludge-Stained Band (286535, -0.33 DPS) [world]; Volcanic Rock Ring (12053, -0.41 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.48 DPS) | yes | Sludge-Stained Band (286535, -0.19 DPS) [world]; Lavishly Jeweled Ring (1156, -0.22 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | Westfall: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.77 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.21 DPS) [world]; Lesser Staff of the Spire (1300, -0.35 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (dwarf, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 47.7. Weights run: 1.7s. Verify run: 1.1s. 359 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.811 ± 0.026, crit=0.124 ± 0.008 per rating point (14 rating = 1%, 1.742 per %), hit=0.230 ± 0.003 per rating point (10 rating = 1%, 2.297 per %), spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 14.1 spell_power points (1.41 DPS) | yes | Enduring Cap (3020, +0.00 DPS, sim-verified) [world_drop]; Totemic Leather Helm (252456, -0.21 DPS) [crafted]; Holy Shroud (2721, -0.31 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.9 spell_power points (1.18 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.24 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.3 spell_power points (1.63 DPS) | yes | Death Speaker Mantle (6685, -0.11 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.30 DPS) [quest]; Magician's Mantle (12998, -0.40 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.5 spell_power points (0.65 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.02 DPS) [quest]; Hillman's Cloak (3719, -0.15 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 19.5 spell_power points (1.95 DPS) | yes | Guardian Armor (4256, -0.25 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.36 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.47 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.90 DPS) | yes | Nightsky Wristbands (6407, -0.41 DPS) [world_drop]; Technician's Bracers (270042, -0.41 DPS) [quest]; Glowing Magical Bracelets (13106, -1.28 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 12.9 spell_power points (1.29 DPS) | yes | Truefaith Gloves (7049, -0.55 DPS) [crafted]; Gloves of Insight (9698, -0.59 DPS) [quest]; Stormrider's Leather Gloves (252498, -1.01 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.9 spell_power points (1.48 DPS) | yes | Moss Cinch (6911, -0.29 DPS) [dungeon]; Crimson Silk Belt (7055, -0.32 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.38 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 18.9 spell_power points (1.89 DPS) | yes | Abomination Skin Leggings (23173, -0.40 DPS, sim-verified) [dungeon]; Stormrider's Leather Pants (252502, -0.40 DPS) [crafted]; Guardian Pants (5962, -0.44 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.7 spell_power points (1.27 DPS) | yes | Spidersilk Boots (4320, -0.24 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.26 DPS) [crafted]; Acidic Walkers (9454, -1.15 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.70 DPS) | yes | Black Widow Band (6199, -0.13 DPS) [world]; Snake Hoop (6750, -0.13 DPS) [quest]; Minor Channeling Ring (1449, -1.42 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (47.7 DPS) | yes | Black Widow Band (6199, -0.03 DPS) [world]; Snake Hoop (6750, -0.03 DPS) [quest]; Minor Channeling Ring (1449, -1.02 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -1.40 DPS) [dungeon]; Glimmering Staff (249392, -1.41 DPS) [crafted]; Manual Crowd Pummeler (9449, -3.53 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 359, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (dwarf, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 62.3. Weights run: 1.9s. Verify run: 1.2s. 582 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.227 ± 0.045, crit=0.214 ± 0.014 per rating point (14 rating = 1%, 2.997 per %), hit=0.380 ± 0.006 per rating point (10 rating = 1%, 3.799 per %), spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 26.2 spell_power points (2.33 DPS) | yes | Corpseshroud (10574, +0.00 DPS, sim-verified) [dungeon]; Augural Shroud (2620, -0.26 DPS) [world]; Spellpower Goggles Xtreme (10502, -0.46 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.4 spell_power points (1.28 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.42 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.51 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.51 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.9 spell_power points (2.04 DPS) | yes | Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Sheepshear Mantle (13115, -0.30 DPS) [world_drop]; Green Silken Shoulders (7057, -0.45 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 20.0 spell_power points (1.78 DPS) | yes | Long Silken Cloak (4326, -0.70 DPS) [crafted]; Guardian Cloak (5965, -0.70 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.85 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.4 spell_power points (2.61 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.06 DPS) [crafted]; Big Voodoo Robe (8200, -0.28 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 13.4 spell_power points (1.19 DPS) | yes | Windchaser Cuffs (14429, -0.21 DPS) [world_drop]; Turtle Scale Bracers (8198, -0.27 DPS, sim-verified) [crafted]; Green Whelp Bracers (7386, -0.32 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.14 DPS) | yes | Dreamweave Gloves (10019, -0.10 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.17 DPS) [crafted]; Red Mageweave Gloves (10018, -1.10 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.4 spell_power points (2.26 DPS) | yes | Highlander's Lizardhide Girdle (20104, -0.62 DPS) [rep]; Highlander's Mail Girdle (20119, -0.62 DPS) [vendor]; Highlander's Cloth Girdle (20098, -1.28 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.7 spell_power points (2.56 DPS) | yes | Kodohide Legguards (285338, -0.47 DPS) [world]; Crimson Silk Pantaloons (7062, -0.74 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.88 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.14 DPS) | yes | Skycaller's Mail Boots (252563, -0.21 DPS) [crafted]; Mender's Leather Shoes (252533, -0.66 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.68 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.4 spell_power points (1.54 DPS) | yes | Ogremind Ring (1993, -0.78 DPS) [world_drop]; Voodoo Band (1996, -0.78 DPS) [world]; Black Widow Band (6199, -0.78 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.80 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Voodoo Band (1996, -0.04 DPS) [world]; Black Widow Band (6199, -1.33 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-verified (62.3 DPS) | yes | Spellforce Rod (1664, -0.46 DPS) [world]; Mograine's Might (7723, -0.49 DPS) [dungeon]; Gut Ripper (2164, -3.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 582, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 50 (dwarf, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 86.6. Weights run: 2.0s. Verify run: 1.4s. 752 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.216 ± 0.055, crit=0.275 ± 0.018 per rating point (14 rating = 1%, 3.848 per %), hit=0.505 ± 0.007 per rating point (10 rating = 1%, 5.046 per %), spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 43.3 spell_power points (3.82 DPS) | yes | Soothsayer's Headdress (17740, +0.00 DPS, sim-verified) [dungeon]; Dreamweave Circlet (10041, -0.90 DPS) [crafted]; Chief Architect's Monocle (11839, -0.92 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.0 spell_power points (1.50 DPS) | yes | Mindburst Medallion (11196, -0.33 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.43 DPS) [quest]; Scorn's Icy Choker (23169, -2.32 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 39.3 spell_power points (3.46 DPS) | yes | Lead Surveyor's Mantle (11842, -0.64 DPS) [dungeon]; Rotgrip Mantle (17732, -0.69 DPS, sim-verified) [dungeon]; Kentic Amice (11624, -0.84 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.3 spell_power points (1.88 DPS) | yes | Mantle of Lady Falther'ess (23178, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.23 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.38 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 43.3 spell_power points (3.82 DPS) | yes | Feathered Breastplate (8349, -1.07 DPS) [crafted]; Robes of Insight (940, -1.14 DPS) [world_drop]; Runecloth Robe (13858, -2.09 DPS, sim-verified) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 22.2 spell_power points (1.95 DPS) | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Skycaller's Mail Bracers (252571, -0.32 DPS) [crafted]; Aristocratic Cuffs (12546, -0.35 DPS) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 33.8 spell_power points (2.97 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.60 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.60 DPS) [crafted]; Raider Handguards (272102, -2.24 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 30.6 spell_power points (2.70 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -0.13 DPS) [dungeon]; Satyrmane Sash (17755, -0.39 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 35.2 spell_power points (3.10 DPS) | yes | Big Voodoo Pants (8202, -0.70 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -0.91 DPS) [dungeon]; Red Mageweave Pants (10009, -2.03 DPS, sim-verified) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 28.4 spell_power points (2.50 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS, sim-verified) [crafted]; Greaves of Withering Despair (22240, -0.01 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.39 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 18.2 spell_power points (1.61 DPS) | yes | Philanthropist's Ring (281635, -0.08 DPS) [quest]; Mindseye Circle (10634, -0.32 DPS) [dungeon]; Band of the Unicorn (7553, -0.46 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.5 spell_power points (1.54 DPS) | yes | Mindseye Circle (10634, -0.26 DPS) [dungeon]; Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Philanthropist's Ring (281635, -1.20 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (86.6 DPS) | yes | Uther's Strength (11302, -0.21 DPS, sim-verified) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (86.6 DPS) | yes | Spellshifter Rod (9527, -0.64 DPS) [quest]; Mechanic's Pipehammer (9604, -0.90 DPS) [quest]; Blade of Eternal Darkness (17780, -1.73 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 752, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

### Band 60 (dwarf, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 143.6. Weights run: 1.9s. Verify run: 1.4s. 1712 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.219 ± 0.086, crit=0.399 ± 0.027 per rating point (14 rating = 1%, 5.588 per %), hit=0.656 ± 0.011 per rating point (10 rating = 1%, 6.555 per %), spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 53.5 spell_power points (4.64 DPS) | yes | Black Dragonscale Helm (252605, +0.00 DPS, sim-verified) [crafted]; Magister's Crown (16686, -0.51 DPS) [dungeon]; Living Crown (252561, -0.57 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.9 spell_power points (2.67 DPS) | yes | Beads of Ogre Mojo (22149, +0.00 DPS, sim-verified) [quest]; Archlight Talisman (15856, -0.58 DPS) [quest]; Lady Maye's Pendant (14558, -0.67 DPS) [world_drop] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 50.9 spell_power points (4.41 DPS) | yes | Darkspear Shoulderguards (272958, +0.00 DPS, sim-verified) [vendor]; Darkspear Shoulderpads (272103, -0.52 DPS) [vendor]; Darkspear Shoulders (272104, -0.52 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 32.3 spell_power points (2.80 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Crystalline Threaded Cape (20697, -0.64 DPS) [world]; Spritecaster Cape (11623, -0.95 DPS) [dungeon] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Magister's Robes (16688, -0.11 DPS) [dungeon]; Tunic of Undead Slaying (23089, -8.44 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Modest Armguards (18458, -0.66 DPS) [dungeon]; Sublime Wristguards (18497, -0.66 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -7.49 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (143.6 DPS) | yes | Raider Handguards (272102, -0.62 DPS) [vendor]; Hands of Power (13253, -0.66 DPS) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.4 spell_power points (5.06 DPS) | yes | Belt of the Archmage (18405, -1.15 DPS) [crafted]; Girdle of Insight (18504, -1.33 DPS) [crafted]; Stormseeker's Girdle (272399, -2.27 DPS, sim-verified) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 65.5 spell_power points (5.68 DPS) | yes | Red Dragonscale Leggings (252603, +0.00 DPS, sim-verified) [crafted]; Sentinel's Silk Leggings (237815, -0.88 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -0.88 DPS) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 40.5 spell_power points (3.51 DPS) | yes | Omnicast Boots (11822, -0.51 DPS) [dungeon]; Waterspout Boots (18322, -0.71 DPS) [dungeon]; Dragonrider Boots (18102, -0.74 DPS, sim-verified) [dungeon] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.65 DPS) [quest]; Maiden's Circle (13001, -1.65 DPS) [world_drop]; Naglering (11669, -9.01 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.77 DPS) [quest]; Maiden's Circle (13001, -0.77 DPS) [world_drop]; Naglering (11669, -6.02 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+10.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Briarwood Reed (12930, -2.56 DPS, sim-verified) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, -0.82 DPS) [pvp]; Hand of Edward the Odd (2243, -6.04 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Blue Dragonscale Helm; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Crackling Staff

No-known-source sample (15 of 1712, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak

## Horde

### Band 20 (orc, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 26.1. Weights run: 1.6s. Verify run: 1.0s. 207 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.723 ± 0.018, crit=0.059 ± 0.003 per rating point (14 rating = 1%, 0.822 per %), hit=0.153 ± 0.002 per rating point (10 rating = 1%, 1.532 per %), spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.02 DPS) [crafted]; Totemic Leather Hood (252448, -0.32 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Roadwatcher's Confidence (281265, -0.34 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.5 spell_power points (1.10 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.68 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.72 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.65 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.04 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.3 spell_power points (0.42 DPS) | yes | Owl Bracers (4796, +0.00 DPS, sim-verified) [vendor]; Mindthrust Bracers (1974, -0.07 DPS) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.9 spell_power points (0.76 DPS) | yes | Pristine Gloves (253913, -0.17 DPS) [crafted]; Gnoll Casting Gloves (892, -0.18 DPS) [world]; Serpent Gloves (5970, -0.33 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.45 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Pants (252503, -0.29 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.38 DPS) [crafted]; Abomination Skin Leggings (23173, -0.61 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.9 spell_power points (0.95 DPS) | yes | Stormrider's Leather Boots (252443, -0.14 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.22 DPS) [crafted]; Totemic Leather Boots (252442, -0.37 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.48 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Sludge-Stained Band (286535, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.3 spell_power points (0.42 DPS) | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -0.13 DPS) [world]; Volcanic Rock Ring (12053, -0.21 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | Westfall: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.77 DPS) | yes | Twisted Chanter's Staff (890, -0.07 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.09 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.21 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 207, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 45.1. Weights run: 1.7s. Verify run: 1.1s. 353 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.811 ± 0.026, crit=0.124 ± 0.008 per rating point (14 rating = 1%, 1.742 per %), hit=0.230 ± 0.003 per rating point (10 rating = 1%, 2.297 per %), spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 14.1 spell_power points (1.41 DPS) | yes | Enduring Cap (3020, -0.06 DPS, sim-verified) [world_drop]; Totemic Leather Helm (252456, -0.21 DPS) [crafted]; Holy Shroud (2721, -0.31 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.9 spell_power points (1.18 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.17 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.3 spell_power points (1.63 DPS) | yes | Death Speaker Mantle (6685, -0.22 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.30 DPS) [quest]; Magician's Mantle (12998, -0.40 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.5 spell_power points (0.65 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.15 DPS) [crafted]; Windsong Drape (15468, -0.15 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 19.5 spell_power points (1.95 DPS) | yes | Death Speaker Robes (6682, -0.36 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.47 DPS) [crafted]; Guardian Armor (4256, -0.50 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.90 DPS) | yes | Nightsky Wristbands (6407, -0.41 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.41 DPS) [quest]; Glowing Magical Bracelets (13106, -1.41 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.6 spell_power points (1.06 DPS) | yes | Jutebraid Gloves (10654, -0.15 DPS, sim-verified) [quest]; Stormrider's Leather Gloves (252498, -0.24 DPS) [crafted]; Truefaith Gloves (7049, -0.32 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 14.9 spell_power points (1.48 DPS) | yes | Moss Cinch (6911, -0.29 DPS) [dungeon]; Crimson Silk Belt (7055, -0.32 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.63 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 18.9 spell_power points (1.89 DPS) | yes | Stormrider's Leather Pants (252502, -0.40 DPS) [crafted]; Guardian Pants (5962, -0.44 DPS) [crafted]; Abomination Skin Leggings (23173, -0.64 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.7 spell_power points (1.27 DPS) | yes | Spidersilk Boots (4320, -0.24 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.26 DPS) [crafted]; Acidic Walkers (9454, -1.06 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.70 DPS) | yes | Black Widow Band (6199, -0.13 DPS) [world]; Snake Hoop (6750, -0.13 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.21 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.60 DPS) | yes | Snake Hoop (6750, -0.03 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Black Widow Band (6199, -1.03 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (45.1 DPS) | yes | Glimmering Staff (249392, -0.01 DPS) [crafted]; Gnarled Necromancer's Staff (251534, -0.09 DPS) [quest]; Manual Crowd Pummeler (9449, -2.03 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 11.9 spell_power points (1.18 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.66 DPS) [vendor]; Seedcloud Buckler (6630, -0.69 DPS) [dungeon]; Witch's Finger (16887, -0.82 DPS, sim-verified) [quest] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 353, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 62.7. Weights run: 1.9s. Verify run: 1.2s. 565 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.227 ± 0.045, crit=0.214 ± 0.014 per rating point (14 rating = 1%, 2.997 per %), hit=0.380 ± 0.006 per rating point (10 rating = 1%, 3.799 per %), spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 26.2 spell_power points (2.33 DPS) | yes | Corpseshroud (10574, -0.10 DPS, sim-verified) [dungeon]; Augural Shroud (2620, -0.26 DPS) [world]; Spellpower Goggles Xtreme (10502, -0.46 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.4 spell_power points (1.28 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.12 DPS, sim-verified) [quest]; Triune Amulet (7722, -0.51 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.51 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.9 spell_power points (2.04 DPS) | yes | Bloodmage Mantle (7684, -0.26 DPS) [dungeon]; Sheepshear Mantle (13115, -0.30 DPS) [world_drop]; Green Silken Shoulders (7057, -0.31 DPS, sim-verified) [crafted] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 20.0 spell_power points (1.78 DPS) | yes | Long Silken Cloak (4326, -0.70 DPS) [crafted]; Guardian Cloak (5965, -0.70 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.84 DPS, sim-verified) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (62.7 DPS) | yes | Robe of Power (7054, -0.03 DPS) [crafted]; Big Voodoo Robe (8200, -0.26 DPS) [crafted]; Robe of the Magi (1716, -0.64 DPS, sim-verified) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 13.8 spell_power points (1.23 DPS) | yes | Turtle Scale Bracers (8198, -0.15 DPS) [crafted]; Guardian Leather Bracers (4260, -0.19 DPS, sim-verified) [crafted]; Windchaser Cuffs (14429, -0.25 DPS) [world_drop] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.14 DPS) | yes | Dreamweave Gloves (10019, -0.10 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.17 DPS) [crafted]; Red Mageweave Gloves (10018, -1.08 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.4 spell_power points (2.26 DPS) | yes | Highlander's Mail Girdle (20119, -0.62 DPS) [vendor]; Defiler's Lizardhide Girdle (20173, -0.62 DPS) [rep]; Defiler's Cloth Girdle (20166, -0.78 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.7 spell_power points (2.56 DPS) | yes | Kodohide Legguards (285338, -0.47 DPS) [world]; Crimson Silk Pantaloons (7062, -0.60 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.88 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.14 DPS) | yes | Skycaller's Mail Boots (252563, -0.21 DPS) [crafted]; Mender's Leather Shoes (252533, -0.66 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.80 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.4 spell_power points (1.54 DPS) | yes | Ogremind Ring (1993, -0.78 DPS) [world_drop]; Voodoo Band (1996, -0.78 DPS) [world]; Black Widow Band (6199, -0.78 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.80 DPS) | yes | Ogremind Ring (1993, -0.04 DPS) [world_drop]; Voodoo Band (1996, -0.04 DPS) [world]; Black Widow Band (6199, -1.17 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.03 DPS) [dungeon]; Windweaver Staff (7757, -0.14 DPS) [dungeon]; Gut Ripper (2164, -3.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Radiant Silver Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 565, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 87.1. Weights run: 2.0s. Verify run: 1.4s. 720 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.216 ± 0.055, crit=0.275 ± 0.018 per rating point (14 rating = 1%, 3.848 per %), hit=0.505 ± 0.007 per rating point (10 rating = 1%, 5.046 per %), spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 43.3 spell_power points (3.82 DPS) | yes | Soothsayer's Headdress (17740, +0.00 DPS, sim-verified) [dungeon]; Dreamweave Circlet (10041, -0.90 DPS) [crafted]; Chief Architect's Monocle (11839, -0.92 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.0 spell_power points (1.50 DPS) | yes | Mindburst Medallion (11196, -0.33 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.43 DPS) [quest]; Scorn's Icy Choker (23169, -0.77 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 39.3 spell_power points (3.46 DPS) | yes | Rotgrip Mantle (17732, -0.11 DPS, sim-verified) [dungeon]; Lead Surveyor's Mantle (11842, -0.64 DPS) [dungeon]; Kentic Amice (11624, -0.84 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (87.1 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.12 DPS) [dungeon]; Runecloth Cloak (13860, -0.23 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.08 DPS, sim-verified) [quest] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 43.3 spell_power points (3.82 DPS) | yes | Runecloth Robe (13858, -1.03 DPS) [crafted]; Feathered Breastplate (8349, -1.07 DPS) [crafted]; Stone Guard's Pulsing Breastplate (220844, -1.54 DPS, sim-verified) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 22.2 spell_power points (1.95 DPS) | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Skycaller's Mail Bracers (252571, -0.32 DPS) [crafted]; Aristocratic Cuffs (12546, -0.35 DPS) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 33.8 spell_power points (2.97 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.60 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.60 DPS) [crafted]; Raider Handguards (272102, -0.93 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 30.6 spell_power points (2.70 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS, sim-verified) [crafted]; Dawnspire Cord (12466, -0.13 DPS) [dungeon]; Satyrmane Sash (17755, -0.39 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 37.6 spell_power points (3.32 DPS) | yes | Spellshock Leggings (9484, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -0.80 DPS) [crafted]; Big Voodoo Pants (8202, -0.92 DPS) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 28.4 spell_power points (2.50 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS, sim-verified) [crafted]; Greaves of Withering Despair (22240, -0.01 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.39 DPS) [crafted] |
| finger1 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 18.2 spell_power points (1.61 DPS) | yes | Philanthropist's Ring (281635, -0.08 DPS) [quest]; Mindseye Circle (10634, -0.32 DPS) [dungeon]; Band of the Unicorn (7553, -0.46 DPS) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 17.5 spell_power points (1.54 DPS) | yes | Mindseye Circle (10634, -0.26 DPS) [dungeon]; Band of the Unicorn (7553, -0.40 DPS) [world_drop]; Philanthropist's Ring (281635, -0.49 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, -0.11 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.64 DPS) [quest]; Radiant Staff (249453, -1.07 DPS) [crafted]; Blade of Eternal Darkness (17780, -1.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; feet: Skycaller's Leather Boots; finger1: Brainlash; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 720, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 145.1. Weights run: 1.9s. Verify run: 1.3s. 1634 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=1.219 ± 0.086, crit=0.399 ± 0.027 per rating point (14 rating = 1%, 5.588 per %), hit=0.656 ± 0.011 per rating point (10 rating = 1%, 6.555 per %), spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Coif of The Five Thunders (227002) | Saving the Best for Last [quest] | 57.1 spell_power points (4.95 DPS) | yes | Blue Dragonscale Helm (252604, +0.00 DPS, sim-verified) [crafted]; Warlord's Mail Helm (231663, -0.36 DPS) [pvp]; Black Dragonscale Helm (252605, -0.45 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 30.9 spell_power points (2.67 DPS) | yes | Archlight Talisman (15856, -0.58 DPS) [quest]; Lady Maye's Pendant (14558, -0.67 DPS) [world_drop]; Beads of Ogre Mojo (22149, -0.87 DPS, sim-verified) [quest] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Mail Spaulders (231659, -0.25 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.35 DPS) [vendor]; Rugged Mantle of the Timbermaw (227808, -1.63 DPS, sim-verified) [vendor] |
| back | Hide of the Wild (18510) | Leatherworking [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Crystalline Threaded Cape (20697, -0.11 DPS) [world]; Deep Woodlands Cloak (19121, -0.28 DPS) [quest]; Arcanoweave Cloak (272411, -1.76 DPS, sim-verified) [vendor] |
| chest | Warlord's Mail Breastplate (231662) | Lady Palanseer [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Legionnaire's Mail Breastplate (227165, -0.47 DPS) [vendor]; Vest of Elements (16666, -0.48 DPS) [dungeon]; Tunic of Undead Slaying (23089, -7.99 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Modest Armguards (18458, -0.66 DPS) [dungeon]; Sublime Wristguards (18497, -0.66 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -6.04 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | sim-verified (+3.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -0.26 DPS) [vendor]; Raider Handguards (272101, -2.96 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 58.4 spell_power points (5.06 DPS) | yes | Belt of the Archmage (18405, -1.15 DPS) [crafted]; Girdle of Insight (18504, -1.33 DPS) [crafted]; Stormseeker's Girdle (272399, -1.41 DPS, sim-verified) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 65.5 spell_power points (5.68 DPS) | yes | General's Mail Leggings (16579, +0.00 DPS, sim-verified) [vendor]; Red Dragonscale Leggings (252603, -0.56 DPS) [crafted] |
| feet | General's Mail Boots (16573) | Lady Palanseer [vendor] | sim-verified (+3.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Slippers of The Five Thunders (227007, -0.02 DPS) [vendor]; Dragonrider Boots (18102, -0.28 DPS) [dungeon]; General's Mail Sabatons (231661, -3.84 DPS, sim-verified) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.65 DPS) [quest]; Maiden's Circle (13001, -1.65 DPS) [world_drop]; Naglering (11669, -6.60 DPS, sim-verified) [dungeon] |
| finger2 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.77 DPS) [quest]; Maiden's Circle (13001, -0.77 DPS) [world_drop]; Naglering (11669, -4.77 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+9.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Briarwood Reed (12930, -1.31 DPS, sim-verified) [dungeon] |
| main_hand | Hammer of Divine Might (22333) | Scholomance: Kormok [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -4.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Coif of The Five Thunders; neck: Amulet of the Dawn; shoulder: Darkspear Shoulderguards; back: Hide of the Wild; chest: Warlord's Mail Breastplate; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: General's Mail Boots; finger1: Signet Ring of the Bronze Dragonflight; finger2: Elemental Focus Band; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Hammer of Divine Might

No-known-source sample (15 of 1634, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

