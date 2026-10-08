# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 24.0. Weights run: 2.3s. Verify run: 0.9s. 225 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.591 ± 0.008, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.930 per %), hit=0.194 ± 0.006 per rating point (10 rating = 1%, 1.936 per %), spell_haste=-1.628 ± 0.069, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.11 DPS) [crafted]; Totemic Leather Hood (252448, -0.32 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.3 spell_power points (0.86 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.25 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.53 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Pearl-clasped Cloak (5542, -0.02 DPS) [crafted]; Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.02 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.05 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.05 DPS) [vendor]; Owl Bracers (4796, -1.36 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.4 spell_power points (0.61 DPS) | yes | Gnoll Casting Gloves (892, -0.11 DPS) [world]; Windfelt Gloves (5630, -0.13 DPS) [quest]; Serpent Gloves (5970, -0.35 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.44 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.7 spell_power points (1.14 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.26 DPS) [crafted]; Dreamer's Leggings (270016, -0.27 DPS) [quest] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.4 spell_power points (0.78 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.28 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.51 DPS) | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.37 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.42 DPS) | yes | Sludge-Stained Band (286535, -0.17 DPS) [world]; Lavishly Jeweled Ring (1156, -0.18 DPS, sim-verified) [dungeon]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.66 DPS) | yes | Twisted Chanter's Staff (890, -0.17 DPS) [world_drop]; Channeler's Staff (4437, -0.27 DPS) [world]; Lesser Staff of the Spire (1300, -0.37 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 5530311300000000-000000000000000000-0000000000000000)

Set DPS (verified): 43.6. Weights run: 2.6s. Verify run: 1.0s. 366 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.562 ± 0.009, crit=0.081 ± 0.002 per rating point (14 rating = 1%, 1.128 per %), hit=0.245 ± 0.007 per rating point (10 rating = 1%, 2.449 per %), spell_haste=-2.121 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.06 DPS) [world_drop]; Silk Headband (7050, -0.27 DPS) [crafted]; Totemic Leather Helm (252456, -0.42 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (1.08 DPS) | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.84 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.00 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (1.46 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS) [dungeon]; Fairywing Mantle (9536, -0.31 DPS) [quest]; Magician's Mantle (12998, -0.42 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.2 spell_power points (0.55 DPS) | yes | Hillman's Cloak (3719, -0.03 DPS) [crafted]; Cloak of Rot (4462, -0.08 DPS) [world]; Darkspear Raider's Cloak (272078, -0.08 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (1.69 DPS) | yes | Guardian Armor (4256, -0.18 DPS) [crafted]; Stormrider's Leather Tunic (252510, -0.31 DPS) [crafted]; Death Speaker Robes (6682, -0.32 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.93 DPS) | yes | Nightsky Wristbands (6407, -0.58 DPS) [world_drop]; Technician's Bracers (270042, -0.58 DPS) [quest]; Glowing Magical Bracelets (13106, -0.89 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.2 spell_power points (1.06 DPS) | yes | Shilly Mitts (9609, -0.33 DPS) [quest]; Gloves of Insight (9698, -0.33 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.90 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.4 spell_power points (1.39 DPS) | yes | Moss Cinch (6911, -0.14 DPS) [dungeon]; Belt of Arugal (6392, -0.28 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.38 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 16.2 spell_power points (1.68 DPS) | yes | Abomination Skin Leggings (23173, -0.28 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Dark Ritual Leggings (270031, -0.72 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.14 DPS) | yes | Spidersilk Boots (4320, -0.18 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.22 DPS) [crafted]; Acidic Walkers (9454, -1.23 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.73 DPS) | yes | Black Widow Band (6199, -0.32 DPS) [world]; Snake Hoop (6750, -0.32 DPS) [quest]; Minor Channeling Ring (1449, -1.32 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.21 DPS) [world]; Snake Hoop (6750, -0.21 DPS) [quest]; Minor Channeling Ring (1449, -0.89 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -1.33 DPS) [dungeon]; Rhahk'Zor's Hammer (5187, -1.43 DPS) [dungeon]; Manual Crowd Pummeler (9449, -3.49 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 5530311300103050-010000000000000000-0000000000000000)

Set DPS (verified): 67.0. Weights run: 2.8s. Verify run: 1.0s. 592 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.908 ± 0.024, crit=0.194 ± 0.007 per rating point (14 rating = 1%, 2.717 per %), hit=0.390 ± 0.022 per rating point (10 rating = 1%, 3.901 per %), spell_haste=-2.633 ± 0.292, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.946 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 21.7 spell_power points (2.20 DPS) | yes | Augural Shroud (2620, -0.17 DPS) [world]; Corpseshroud (10574, -0.45 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.72 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.4 spell_power points (1.26 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest]; Triune Amulet (7722, -0.62 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.62 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.8 spell_power points (1.90 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 17.2 spell_power points (1.74 DPS) | yes | Guardian Cloak (5965, -0.67 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.73 DPS) [vendor]; Long Silken Cloak (4326, -1.66 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.4 spell_power points (2.78 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.26 DPS) [crafted]; Big Voodoo Robe (8200, -0.58 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.4 spell_power points (1.16 DPS) | yes | Turtle Scale Bracers (8198, -0.09 DPS) [crafted]; Arcane Runed Bracers (4744, -0.25 DPS) [quest]; Spidertank Oilrag (9448, -0.25 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.43 DPS) | yes | Red Mageweave Gloves (10018, -0.40 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.49 DPS) [crafted]; Dreamweave Gloves (10019, -1.05 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.6 spell_power points (2.09 DPS) | yes | Skycaller's Leather Belt (252522, -0.52 DPS) [crafted]; Gilded Cord (254037, -0.54 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.94 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.9 spell_power points (2.52 DPS) | yes | Crimson Silk Pantaloons (7062, -0.52 DPS) [crafted]; Kodohide Legguards (285338, -0.68 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -0.87 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.43 DPS) | yes | Skycaller's Mail Boots (252563, -0.47 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.66 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -0.98 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (1.56 DPS) | yes | Ring of Forlorn Spirits (2043, -0.75 DPS) [quest]; Reedknot Ring (9622, -0.86 DPS) [quest]; Minor Channeling Ring (1449, -0.87 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.91 DPS) | yes | Reedknot Ring (9622, -0.20 DPS) [quest]; Minor Channeling Ring (1449, -0.22 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.29 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-verified (67.0 DPS) | yes | Spellforce Rod (1664, -0.36 DPS) [world]; Mograine's Might (7723, -0.91 DPS) [dungeon]; Gut Ripper (2164, -3.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 5530311300103050-030000000000000000-5300000000000000)

Set DPS (verified): 85.8. Weights run: 2.8s. Verify run: 1.2s. 762 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.848 ± 0.030, crit=0.247 ± 0.009 per rating point (14 rating = 1%, 3.452 per %), hit=0.503 ± 0.029 per rating point (10 rating = 1%, 5.031 per %), spell_haste=-2.533 ± 0.348, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.939 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.0 spell_power points (3.43 DPS) | yes | Soothsayer's Headdress (17740, -0.02 DPS) [dungeon]; Dreamweave Circlet (10041, -0.62 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.86 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.15 DPS) | yes | Horizon Choker (13085, -0.02 DPS) [world_drop]; Mindburst Medallion (11196, -0.10 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 32.0 spell_power points (3.05 DPS) | yes | Rotgrip Mantle (17732, -0.35 DPS) [dungeon]; Lead Surveyor's Mantle (11842, -0.49 DPS) [dungeon]; Kentic Amice (11624, -0.66 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.1 spell_power points (1.82 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.23 DPS) [dungeon]; Runecloth Cloak (13860, -0.32 DPS) [crafted]; Big Voodoo Cloak (8216, -0.62 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.0 spell_power points (3.43 DPS) | yes | Feathered Breastplate (8349, -0.81 DPS) [crafted]; Robe of the Magi (1716, -0.85 DPS) [world_drop]; Runecloth Tunic (13857, -0.92 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.5 spell_power points (1.77 DPS) | yes | Skycaller's Leather Bracers (252542, -0.24 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.24 DPS) [crafted]; Nethergeld Cuffs (254061, -0.53 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (85.8 DPS) | yes | Gloves of the Greatfather (17721, -0.16 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.19 DPS) [crafted]; Raider Handguards (272102, -1.06 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.4 spell_power points (2.52 DPS) | yes | Skycaller's Mail Belt (252589, -0.02 DPS) [crafted]; Satyrmane Sash (17755, -0.37 DPS) [dungeon]; Skycaller's Leather Waistguard (252476, -0.82 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.5 spell_power points (3.01 DPS) | yes | Big Voodoo Pants (8202, -0.76 DPS) [crafted]; Turtle Scale Leggings (8185, -1.15 DPS) [crafted]; Red Mageweave Pants (10009, -2.32 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.5 spell_power points (2.34 DPS) | yes | Skycaller's Leather Boots (252471, -0.02 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.02 DPS) [crafted]; Earthen Silk Slippers (254013, -0.05 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.44 DPS) | yes | Band of the Unicorn (7553, -0.20 DPS) [world_drop]; Brainlash (6440, -0.23 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.29 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.9 spell_power points (1.43 DPS) | yes | Brainlash (6440, -0.21 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.28 DPS) [rep]; Band of the Unicorn (7553, -1.07 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -1.00 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Mechanic's Pipehammer (9604, -0.13 DPS) [quest]; Thorium Greatmace (250613, -0.15 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Greaves of Withering Despair; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5530311300103050-030000000000000000-5533020000000000)

Set DPS (verified): 156.1. Weights run: 2.7s. Verify run: 1.2s. 1751 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.900 ± 0.027, crit=0.268 ± 0.009 per rating point (14 rating = 1%, 3.755 per %), hit=0.493 ± 0.031 per rating point (10 rating = 1%, 4.926 per %), spell_haste=not significant (-0.418 ± 0.338), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.947 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (156.1 DPS) | yes | Magister's Crown (16686, -0.36 DPS) [dungeon]; Black Dragonscale Helm (252605, -0.37 DPS) [crafted]; Blue Dragonscale Helm (252604, -5.21 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 26.7 spell_power points (3.04 DPS) | yes | Beads of Ogre Mojo (22149, -0.33 DPS) [quest]; Orb of the Darkmoon (19426, -0.54 DPS) [quest]; Chains of the Lich (23125, -0.54 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.3 spell_power points (5.05 DPS) | yes | Darkspear Shoulderguards (272958, -0.45 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.91 DPS) [vendor]; Darkspear Shoulders (272104, -0.91 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.1 spell_power points (3.21 DPS) | yes | Crystalline Threaded Cape (20697, -0.52 DPS) [world]; Hide of the Wild (18510, -0.58 DPS) [crafted]; Spritecaster Cape (11623, -1.00 DPS) [dungeon] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Robe of Everlasting Night (18385, -0.26 DPS) [dungeon]; Tunic of Undead Slaying (23089, -8.85 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Modest Armguards (18458, -0.94 DPS) [dungeon]; Sublime Wristguards (18497, -0.94 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -7.25 DPS, sim-verified) [world] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | 34.9 spell_power points (3.98 DPS) | yes | Raider Handwraps (272097, +0.00 DPS, sim-verified) [vendor]; Hands of Power (13253, -0.40 DPS) [dungeon]; Gloves of Undead Cleansing (23084, -0.90 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 49.7 spell_power points (5.67 DPS) | yes | Stormseeker's Girdle (272399, -1.38 DPS) [vendor]; Girdle of Insight (18504, -1.60 DPS) [crafted]; Belt of the Archmage (18405, -2.58 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.0 spell_power points (6.61 DPS) | yes | Red Dragonscale Leggings (252603, -1.04 DPS) [crafted]; Sentinel's Silk Leggings (237815, -1.08 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.08 DPS) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 35.4 spell_power points (4.04 DPS) | yes | Dragonrider Boots (18102, -0.34 DPS) [dungeon]; Omnicast Boots (11822, -0.52 DPS) [dungeon]; Waterspout Boots (18322, -0.57 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.87 DPS) [quest]; Maiden's Circle (13001, -0.87 DPS) [world_drop]; Naglering (11669, -6.77 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.34 DPS) [quest]; Maiden's Circle (13001, -0.34 DPS) [world_drop]; Naglering (11669, -7.14 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -1.72 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.41 DPS) [dungeon]; Hand of Edward the Odd (2243, -7.00 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Totem of Thunder

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60, raid preset (dwarf, 5530311300103051-052000000000000000-0533020000000000)

Set DPS (verified): 464.4. Weights run: 1.7s. Verify run: 0.9s. 1751 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.383 ± 0.022, crit=0.291 ± 0.008 per rating point (14 rating = 1%, 4.069 per %), hit=0.696 ± 0.046 per rating point (10 rating = 1%, 6.958 per %), spell_haste=2.787 ± 0.547, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.780 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 38.8 spell_power points (14.75 DPS) | yes | Crimson Felt Hat (18727, -2.20 DPS) [dungeon]; Living Crown (252561, -2.64 DPS) [crafted]; Soothsayer's Headdress (17740, -3.84 DPS) [dungeon] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.35 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.77 DPS) [quest]; Diana's Pearl Necklace (22403, -1.13 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.8 spell_power points (13.98 DPS) | yes | Burial Shawl (18681, -4.06 DPS) [dungeon]; Mantle of the Timbermaw (19050, -4.09 DPS) [crafted]; Darkspear Shoulderguards (272958, -5.28 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.0 spell_power points (9.88 DPS) | yes | Crystalline Threaded Cape (20697, -1.70 DPS) [world]; Amplifying Cloak (18350, -3.05 DPS) [dungeon]; Hide of the Wild (18510, -3.11 DPS) [crafted] |
| chest | Robe of Everlasting Night (18385) | Dire Maul: Immol'thar [dungeon] | sim-verified (464.4 DPS) | yes | Vest of Elements (16666, -0.50 DPS) [dungeon]; Chestplate of Tranquility (18373, -0.50 DPS) [dungeon]; Tunic of Undead Slaying (23089, -20.58 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-verified (464.4 DPS) | yes | Modest Armguards (18458, -3.51 DPS) [dungeon]; Sublime Wristguards (18497, -3.51 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -13.07 DPS, sim-verified) [world] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 28.3 spell_power points (10.74 DPS) | yes | Raider Handguards (272101, -1.62 DPS) [vendor]; Gloves of the Greatfather (17721, -1.63 DPS) [crafted]; Storm Gauntlets (12632, -3.04 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 40.4 spell_power points (15.33 DPS) | yes | Barrage Girdle (18721, -5.73 DPS) [dungeon]; Stormseeker's Girdle (272399, -5.77 DPS) [vendor]; Belt of the Archmage (18405, -5.88 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.0 spell_power points (18.59 DPS) | yes | Sentinel's Silk Leggings (237815, -1.90 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.90 DPS) [vendor]; Pristine Scorpid Leggings (252606, -4.51 DPS) [crafted] |
| feet | Waterspout Boots (18322) | Dire Maul: Hydrospawn [dungeon] | 27.3 spell_power points (10.36 DPS) | yes | Omnicast Boots (11822, -1.03 DPS) [dungeon]; Dragonrider Boots (18102, -1.20 DPS) [dungeon]; Slippers of The Five Thunders (227007, -2.88 DPS, sim-verified) [vendor] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (464.4 DPS) | yes | Rune Band of Wizardry (22339, -1.23 DPS) [dungeon]; Maiden's Circle (13001, -2.10 DPS) [world_drop]; Naglering (11669, -12.69 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (464.4 DPS) | yes | Rune Band of Wizardry (22339, -0.27 DPS) [dungeon]; Maiden's Circle (13001, -1.14 DPS) [world_drop]; Naglering (11669, -12.81 DPS, sim-verified) [dungeon] |
| trinket1 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (464.4 DPS) | yes | Weakness Analyzer (272438, -2.66 DPS) [vendor]; Serenity Field (272439, -5.70 DPS) [vendor]; Burst of Knowledge (11832, -6.45 DPS) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (464.4 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Draconic Infused Emblem (22268, -4.37 DPS, sim-verified) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (464.4 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.27 DPS) [world]; Hand of Edward the Odd (2243, -22.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (464.4 DPS) | yes | Totem of the Storm (272432, -2.60 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Everlasting Night; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Waterspout Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Briarwood Reed; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Totem of Thunder

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.0. Weights run: 2.3s. Verify run: 0.9s. 205 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.591 ± 0.008, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.930 per %), hit=0.194 ± 0.006 per rating point (10 rating = 1%, 1.936 per %), spell_haste=-1.628 ± 0.069, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.11 DPS) [crafted]; Totemic Leather Hood (252448, -0.33 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.3 spell_power points (0.86 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.53 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.93 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.33 DPS) | yes | Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.18 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.02 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.25 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.5 spell_power points (0.29 DPS) | yes | Mindthrust Bracers (1974, -0.05 DPS) [dungeon]; Featherbead Bracers (15452, -0.05 DPS) [quest]; Owl Bracers (4796, -0.23 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.4 spell_power points (0.61 DPS) | yes | Gnoll Casting Gloves (892, -0.11 DPS) [world]; Pristine Gloves (253913, -0.13 DPS) [crafted]; Serpent Gloves (5970, -0.57 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.05 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.46 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.7 spell_power points (1.14 DPS) | yes | Stormrider's Leather Pants (252502, -0.22 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.26 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.35 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.4 spell_power points (0.78 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.20 DPS) [crafted]; Totemic Leather Boots (252442, -0.28 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.42 DPS) | yes | Sludge-Stained Band (286535, -0.17 DPS) [world]; Loop of Sacrifice (281673, -0.17 DPS) [quest]; Volcanic Rock Ring (12053, -0.27 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.5 spell_power points (0.29 DPS) | yes | Loop of Sacrifice (281673, -0.05 DPS) [quest]; Volcanic Rock Ring (12053, -0.15 DPS) [world_drop]; Sludge-Stained Band (286535, -0.51 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.66 DPS) | yes | Twisted Chanter's Staff (890, -0.17 DPS) [world_drop]; Channeler's Staff (4437, -0.27 DPS) [world]; Gnarled Necromancer's Staff (251534, -0.56 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 5530311300000000-000000000000000000-0000000000000000)

Set DPS (verified): 41.4. Weights run: 2.6s. Verify run: 1.1s. 349 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.562 ± 0.009, crit=0.081 ± 0.002 per rating point (14 rating = 1%, 1.128 per %), hit=0.245 ± 0.007 per rating point (10 rating = 1%, 2.449 per %), spell_haste=-2.121 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (1.25 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Holy Shroud (2721, -0.10 DPS) [world_drop]; Silk Headband (7050, -0.31 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.4 spell_power points (1.08 DPS) | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.84 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.06 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.1 spell_power points (1.46 DPS) | yes | Death Speaker Mantle (6685, -0.19 DPS) [dungeon]; Fairywing Mantle (9536, -0.31 DPS) [quest]; Magician's Mantle (12998, -0.42 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.52 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Cloak of Rot (4462, -0.05 DPS) [world]; Darkspear Raider's Cloak (272078, -0.05 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.3 spell_power points (1.69 DPS) | yes | Guardian Armor (4256, -0.18 DPS) [crafted]; Stormrider's Leather Tunic (252510, -0.31 DPS) [crafted]; Death Speaker Robes (6682, -0.32 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.93 DPS) | yes | Nightsky Wristbands (6407, -0.58 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.58 DPS) [quest]; Glowing Magical Bracelets (13106, -0.70 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.1 spell_power points (1.05 DPS) | yes | Jutebraid Gloves (10654, -0.14 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.30 DPS) [crafted]; Serpent Gloves (5970, -0.32 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.4 spell_power points (1.39 DPS) | yes | Defiler's Cloth Girdle (20164, -0.07 DPS) [rep]; Moss Cinch (6911, -0.14 DPS) [dungeon]; Warsong Sash (16975, -0.25 DPS) [quest] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 16.2 spell_power points (1.68 DPS) | yes | Abomination Skin Leggings (23173, -0.28 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Dark Ritual Leggings (270031, -0.91 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.9 spell_power points (1.14 DPS) | yes | Spidersilk Boots (4320, -0.18 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.22 DPS) [crafted]; Acidic Walkers (9454, -1.00 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.73 DPS) | yes | Black Widow Band (6199, -0.32 DPS) [world]; Snake Hoop (6750, -0.32 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.38 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.62 DPS) | yes | Snake Hoop (6750, -0.21 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.27 DPS) [dungeon]; Black Widow Band (6199, -0.86 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rhahk'Zor's Hammer (5187, -0.10 DPS) [dungeon]; Reef Axe (6905, -0.10 DPS) [dungeon]; Manual Crowd Pummeler (9449, -1.63 DPS, sim-verified) [dungeon] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-verified (41.4 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.08 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -0.10 DPS) [world]; Orb of Mystic Insight (249394, -0.82 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Seedcloud Buckler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 5530311300103050-010000000000000000-0000000000000000)

Set DPS (verified): 68.2. Weights run: 2.8s. Verify run: 1.1s. 555 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.908 ± 0.024, crit=0.194 ± 0.007 per rating point (14 rating = 1%, 2.717 per %), hit=0.390 ± 0.022 per rating point (10 rating = 1%, 3.901 per %), spell_haste=-2.633 ± 0.292, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.946 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 21.7 spell_power points (2.20 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.07 DPS) [crafted]; Augural Shroud (2620, -0.17 DPS) [world]; Corpseshroud (10574, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.4 spell_power points (1.26 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest]; Triune Amulet (7722, -0.62 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.62 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.8 spell_power points (1.90 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.17 DPS) [dungeon]; Berylline Pads (4197, -0.28 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 17.2 spell_power points (1.74 DPS) | yes | Guardian Cloak (5965, -0.67 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.73 DPS) [vendor]; Long Silken Cloak (4326, -1.35 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (68.2 DPS) | yes | Robe of Power (7054, -0.13 DPS) [crafted]; Big Voodoo Robe (8200, -0.45 DPS) [crafted]; Robe of the Magi (1716, -1.08 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.4 spell_power points (1.16 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.02 DPS) [quest]; Turtle Scale Bracers (8198, -0.09 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.43 DPS) | yes | Dreamweave Gloves (10019, -0.24 DPS) [crafted]; Red Mageweave Gloves (10018, -0.40 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.49 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.6 spell_power points (2.09 DPS) | yes | Skycaller's Leather Belt (252522, -0.52 DPS) [crafted]; Gilded Cord (254037, -0.54 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.80 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.9 spell_power points (2.52 DPS) | yes | Kodohide Legguards (285338, -0.50 DPS) [world]; Crimson Silk Pantaloons (7062, -0.52 DPS) [crafted]; Abomination Skin Leggings (23173, -0.87 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.43 DPS) | yes | Skycaller's Leather Shoes (252532, -0.47 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.47 DPS) [crafted]; Mender's Leather Shoes (252533, -0.98 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (1.56 DPS) | yes | Reedknot Ring (9622, -0.86 DPS) [quest]; Voodoo Band (1996, -0.92 DPS) [world]; Black Widow Band (6199, -0.92 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.91 DPS) | yes | Voodoo Band (1996, -0.27 DPS) [world]; Black Widow Band (6199, -0.27 DPS) [world]; Reedknot Ring (9622, -1.09 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.55 DPS) [dungeon]; Windweaver Staff (7757, -0.65 DPS) [dungeon]; Gut Ripper (2164, -3.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5530311300103050-030000000000000000-5300000000000000)

Set DPS (verified): 86.0. Weights run: 2.8s. Verify run: 1.2s. 704 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.848 ± 0.030, crit=0.247 ± 0.009 per rating point (14 rating = 1%, 3.452 per %), hit=0.503 ± 0.029 per rating point (10 rating = 1%, 5.031 per %), spell_haste=-2.533 ± 0.348, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.939 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.0 spell_power points (3.43 DPS) | yes | Soothsayer's Headdress (17740, -0.02 DPS) [dungeon]; Dreamweave Circlet (10041, -0.62 DPS) [crafted]; Blood Guard's Pulsing Helmet (220848, -0.82 DPS) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.15 DPS) | yes | Horizon Choker (13085, -0.02 DPS) [world_drop]; Mindburst Medallion (11196, -0.10 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 32.0 spell_power points (3.05 DPS) | yes | Lead Surveyor's Mantle (11842, -0.49 DPS) [dungeon]; Kentic Amice (11624, -0.66 DPS) [dungeon]; Rotgrip Mantle (17732, -1.06 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.6 spell_power points (1.88 DPS) | yes | Spritecaster Cape (11623, -0.05 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.29 DPS) [dungeon]; Runecloth Cloak (13860, -0.37 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.0 spell_power points (3.43 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.26 DPS) [vendor]; Feathered Breastplate (8349, -0.81 DPS) [crafted]; Robe of the Magi (1716, -0.85 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.5 spell_power points (1.77 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.24 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Gloves of the Greatfather (17721, -0.16 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.19 DPS) [crafted]; Raider Handguards (272102, -1.08 DPS, sim-verified) [vendor] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 26.4 spell_power points (2.52 DPS) | yes | Skycaller's Leather Waistguard (252476, -0.02 DPS) [crafted]; Skycaller's Mail Belt (252589, -0.02 DPS) [crafted]; Satyrmane Sash (17755, -0.37 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.70 DPS) [crafted]; Big Voodoo Pants (8202, -0.76 DPS) [crafted]; Stone Guard's Pulsing Legplates (220847, -1.11 DPS, sim-verified) [vendor] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 24.5 spell_power points (2.34 DPS) | yes | Skycaller's Leather Boots (252471, -0.02 DPS) [crafted]; Skycaller's Mail Sabatons (252577, -0.02 DPS) [crafted]; Earthen Silk Slippers (254013, -0.05 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.44 DPS) | yes | Band of the Unicorn (7553, -0.20 DPS) [world_drop]; Brainlash (6440, -0.23 DPS) [dungeon]; Advisor's Ring (19519, -0.29 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 14.9 spell_power points (1.43 DPS) | yes | Brainlash (6440, -0.21 DPS) [dungeon]; Advisor's Ring (19519, -0.28 DPS) [rep]; Band of the Unicorn (7553, -0.75 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -0.15 DPS) [crafted]; Spellforce Rod (1664, -0.44 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Ban'thok Sash; legs: Spellshock Leggings; feet: Greaves of Withering Despair; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5530311300103050-030000000000000000-5533020000000000)

Set DPS (verified): 154.5. Weights run: 2.7s. Verify run: 1.1s. 1672 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.900 ± 0.027, crit=0.268 ± 0.009 per rating point (14 rating = 1%, 3.755 per %), hit=0.493 ± 0.031 per rating point (10 rating = 1%, 4.926 per %), spell_haste=not significant (-0.418 ± 0.338), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.947 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Coif of The Five Thunders (227002) | Saving the Best for Last [quest] | 47.3 spell_power points (5.39 DPS) | yes | Blue Dragonscale Helm (252604, -0.13 DPS) [crafted]; Warlord's Mail Helm (231663, -0.29 DPS) [pvp]; Living Crown (252561, -0.69 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 26.7 spell_power points (3.04 DPS) | yes | Beads of Ogre Mojo (22149, -0.33 DPS) [quest]; Orb of the Darkmoon (19426, -0.54 DPS) [quest]; Chains of the Lich (23125, -0.54 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 44.3 spell_power points (5.05 DPS) | yes | Darkspear Shoulderguards (272958, -0.45 DPS) [vendor]; Warlord's Mail Spaulders (231659, -0.58 DPS) [pvp]; Darkspear Shoulderpads (272103, -0.91 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 28.1 spell_power points (3.21 DPS) | yes | Crystalline Threaded Cape (20697, -0.52 DPS) [world]; Hide of the Wild (18510, -0.58 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.92 DPS) [quest] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (154.5 DPS) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Tunic of Undead Slaying (23089, -9.03 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-verified (154.5 DPS) | yes | Modest Armguards (18458, -0.94 DPS) [dungeon]; Sublime Wristguards (18497, -0.94 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -7.32 DPS, sim-verified) [world] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | 34.9 spell_power points (3.98 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Raider Handwraps (272097, +0.00 DPS, sim-verified) [vendor]; Hands of Power (13253, -0.40 DPS) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 49.7 spell_power points (5.67 DPS) | yes | Stormseeker's Girdle (272399, -1.38 DPS) [vendor]; Girdle of Insight (18504, -1.60 DPS) [crafted]; Belt of the Archmage (18405, -2.72 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 58.0 spell_power points (6.61 DPS) | yes | General's Mail Leggings (231664, -0.52 DPS) [pvp]; Red Dragonscale Leggings (252603, -1.04 DPS) [crafted]; Sentinel's Silk Leggings (237815, -1.08 DPS) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 35.4 spell_power points (4.04 DPS) | yes | General's Mail Sabatons (231661, -0.01 DPS) [vendor]; Dragonrider Boots (18102, -0.34 DPS) [dungeon]; Omnicast Boots (11822, -0.52 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (154.5 DPS) | yes | Eye of Orgrimmar (12545, -0.87 DPS) [quest]; Maiden's Circle (13001, -0.87 DPS) [world_drop]; Naglering (11669, -6.95 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (154.5 DPS) | yes | Eye of Orgrimmar (12545, -0.34 DPS) [quest]; Maiden's Circle (13001, -0.34 DPS) [world_drop]; Naglering (11669, -7.35 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (154.5 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (154.5 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (154.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.21 DPS) [world]; Hand of Edward the Odd (2243, -8.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (154.5 DPS) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Coif of The Five Thunders; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; main_hand: Lord Valthalak's Staff of Command; ranged: Totem of Thunder

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60, raid preset (orc, 5530311300103051-052000000000000000-0533020000000000)

Set DPS (verified): 464.6. Weights run: 1.7s. Verify run: 0.9s. 1672 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.383 ± 0.022, crit=0.291 ± 0.008 per rating point (14 rating = 1%, 4.069 per %), hit=0.696 ± 0.046 per rating point (10 rating = 1%, 6.958 per %), spell_haste=2.787 ± 0.547, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.780 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blue Dragonscale Helm (252604) | Leatherworking [crafted] | 38.8 spell_power points (14.75 DPS) | yes | Warlord's Mail Helm (231663, -1.57 DPS) [pvp]; Coif of The Five Thunders (227002, -1.60 DPS) [quest]; Crimson Felt Hat (18727, -2.20 DPS) [dungeon] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (8.35 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -0.77 DPS) [quest]; Diana's Pearl Necklace (22403, -1.13 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.8 spell_power points (13.98 DPS) | yes | Warlord's Mail Spaulders (231659, -2.13 DPS) [pvp]; Pauldrons of The Five Thunders (227003, -2.38 DPS) [quest]; Darkspear Shoulderguards (272958, -3.98 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.0 spell_power points (9.88 DPS) | yes | Crystalline Threaded Cape (20697, -1.70 DPS) [world]; Amplifying Cloak (18350, -3.05 DPS) [dungeon]; Hide of the Wild (18510, -3.11 DPS) [crafted] |
| chest | Robe of Everlasting Night (18385) | Dire Maul: Immol'thar [dungeon] | sim-verified (464.6 DPS) | yes | Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Legionnaire's Mail Breastplate (227165, -0.39 DPS) [vendor]; Tunic of Undead Slaying (23089, -20.57 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-verified (464.6 DPS) | yes | Modest Armguards (18458, -3.51 DPS) [dungeon]; Sublime Wristguards (18497, -3.51 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -12.73 DPS, sim-verified) [world] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 28.3 spell_power points (10.74 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Gauntlets of The Five Thunders (227006, -1.51 DPS) [quest]; Storm Gauntlets (12632, -3.65 DPS, sim-verified) [crafted] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 40.4 spell_power points (15.33 DPS) | yes | Belt of the Archmage (18405, -3.99 DPS, sim-verified) [crafted]; Barrage Girdle (18721, -5.73 DPS) [dungeon]; Stormseeker's Girdle (272399, -5.77 DPS) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 49.0 spell_power points (18.59 DPS) | yes | Sentinel's Silk Leggings (237815, -1.90 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.90 DPS) [vendor]; General's Mail Leggings (231664, -2.93 DPS) [pvp] |
| feet | Waterspout Boots (18322) | Dire Maul: Hydrospawn [dungeon] | 27.3 spell_power points (10.36 DPS) | yes | General's Mail Sabatons (231661, -0.30 DPS) [vendor]; Omnicast Boots (11822, -1.03 DPS) [dungeon]; Slippers of The Five Thunders (227007, -2.67 DPS, sim-verified) [vendor] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (464.6 DPS) | yes | Rune Band of Wizardry (22339, -1.23 DPS) [dungeon]; Maiden's Circle (13001, -2.10 DPS) [world_drop]; Naglering (11669, -12.84 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (464.6 DPS) | yes | Rune Band of Wizardry (22339, -0.27 DPS) [dungeon]; Maiden's Circle (13001, -1.14 DPS) [world_drop]; Naglering (11669, -12.32 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (464.6 DPS) | yes | Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (464.6 DPS) | yes | Weakness Analyzer (272438, -2.66 DPS) [vendor]; Royal Seal of Eldre'Thalas (18471, -3.75 DPS, sim-verified) [quest]; Serenity Field (272439, -5.70 DPS) [vendor] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (464.6 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.27 DPS) [world]; Hand of Edward the Odd (2243, -21.42 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (464.6 DPS) | yes | Totem of the Storm (272432, -2.66 DPS, sim-verified) [world_drop] |

**New at 60:** head: Blue Dragonscale Helm; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Robe of Everlasting Night; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Waterspout Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command; ranged: Totem of Thunder

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

