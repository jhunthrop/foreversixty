# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.3. Weights run: 2.5s. Verify run: 1.0s. 225 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.522 ± 0.008, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.890 per %), hit=0.145 ± 0.001 per rating point (10 rating = 1%, 1.448 per %), spell_haste=-1.485 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.940 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.15 DPS) [crafted]; Totemic Leather Hood (252448, -0.31 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.7 spell_power points (0.82 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.51 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.34 DPS) | yes | Pearl-clasped Cloak (5542, -0.04 DPS) [crafted]; Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.04 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.04 DPS) [vendor]; Owl Bracers (4796, -0.76 DPS, sim-verified) [vendor] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Gnoll Casting Gloves (892, -0.08 DPS) [world]; Windfelt Gloves (5630, -0.12 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.49 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.43 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Pants (252503, -0.25 DPS) [crafted]; Dreamer's Leggings (270016, -0.29 DPS) [quest]; Abomination Skin Leggings (23173, -0.36 DPS, sim-verified) [dungeon] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Boots (252444, -0.17 DPS) [crafted]; Totemic Leather Boots (252442, -0.22 DPS) [crafted]; Spidersilk Boots (4320, -0.64 DPS, sim-verified) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.0 spell_power points (0.51 DPS) | yes | Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.42 DPS) | yes | Lavishly Jeweled Ring (1156, -0.15 DPS, sim-verified) [dungeon]; Sludge-Stained Band (286535, -0.17 DPS) [world]; Volcanic Rock Ring (12053, -0.29 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Channeler's Staff (4437, -0.09 DPS) [world]; Lesser Staff of the Spire (1300, -0.18 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Stormrider's Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 37.8. Weights run: 2.8s. Verify run: 1.1s. 366 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.506 ± 0.014, crit=0.091 ± 0.002 per rating point (14 rating = 1%, 1.276 per %), hit=0.200 ± 0.001 per rating point (10 rating = 1%, 1.998 per %), spell_haste=-2.357 ± 0.272, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.937 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.00 DPS) [world_drop]; Silk Headband (7050, -0.17 DPS) [crafted]; Totemic Leather Helm (252456, -0.66 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.0 spell_power points (0.85 DPS) | yes | Crystal Starfire Medallion (5003, -0.68 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.68 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.01 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.6 spell_power points (1.15 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.25 DPS) [quest]; Magician's Mantle (12998, -0.34 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.0 spell_power points (0.42 DPS) | yes | Cloak of Rot (4462, -0.08 DPS) [world]; Darkspear Raider's Cloak (272078, -0.08 DPS) [vendor]; Hillman's Cloak (3719, -0.42 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.6 spell_power points (1.32 DPS) | yes | Guardian Armor (4256, -0.13 DPS) [crafted]; Stormrider's Leather Tunic (252510, -0.21 DPS) [crafted]; Tree Bark Jacket (1486, -0.22 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.76 DPS) | yes | Nightsky Wristbands (6407, -0.50 DPS) [world_drop]; Technician's Bracers (270042, -0.50 DPS) [quest]; Glowing Magical Bracelets (13106, -1.05 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 9.6 spell_power points (0.81 DPS) | yes | Shilly Mitts (9609, -0.22 DPS) [quest]; Gloves of Insight (9698, -0.22 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.55 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.0 spell_power points (1.10 DPS) | yes | Highlander's Cloth Girdle (20099, -0.04 DPS) [rep]; Moss Cinch (6911, -0.09 DPS) [dungeon]; Belt of Arugal (6392, -0.21 DPS) [dungeon] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 15.6 spell_power points (1.32 DPS) | yes | Abomination Skin Leggings (23173, -0.21 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.21 DPS) [crafted]; Dark Ritual Leggings (270031, -0.57 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.5 spell_power points (0.89 DPS) | yes | Spidersilk Boots (4320, -0.13 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.17 DPS) [crafted]; Acidic Walkers (9454, -0.96 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.59 DPS) | yes | Black Widow Band (6199, -0.29 DPS) [world]; Snake Hoop (6750, -0.29 DPS) [quest]; Minor Channeling Ring (1449, -1.37 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.21 DPS) [world]; Snake Hoop (6750, -0.21 DPS) [quest]; Minor Channeling Ring (1449, -0.79 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -1.06 DPS) [dungeon]; Rhahk'Zor's Hammer (5187, -1.14 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.88 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 4532310300103051-000000000000000000-0000000000000000)

Set DPS (verified): 57.3. Weights run: 3.1s. Verify run: 1.2s. 592 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.014 ± 0.025, crit=0.233 ± 0.008 per rating point (14 rating = 1%, 3.267 per %), hit=0.356 ± 0.003 per rating point (10 rating = 1%, 3.558 per %), spell_haste=-3.775 ± 0.284, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.929 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 23.2 spell_power points (1.83 DPS) | yes | Augural Shroud (2620, -0.16 DPS) [world]; Spellpower Goggles Xtreme (10502, -0.17 DPS) [crafted]; Corpseshroud (10574, -0.31 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.1 spell_power points (1.03 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.23 DPS) [quest]; Triune Amulet (7722, -0.47 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.47 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.2 spell_power points (1.59 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Death Speaker Mantle (6685, -0.24 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 18.1 spell_power points (1.43 DPS) | yes | Long Silken Cloak (4326, -0.56 DPS) [crafted]; Guardian Cloak (5965, -0.56 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.79 DPS, sim-verified) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (57.3 DPS) | yes | Robe of Power (7054, -0.08 DPS) [crafted]; Big Voodoo Robe (8200, -0.31 DPS) [crafted]; Robe of the Magi (1716, -0.89 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 12.1 spell_power points (0.95 DPS) | yes | Turtle Scale Bracers (8198, -0.08 DPS) [crafted]; Windchaser Cuffs (14429, -0.23 DPS) [world_drop]; Arcane Runed Bracers (4744, -0.24 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (1.89 DPS) | yes | Red Mageweave Gloves (10018, -0.23 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.31 DPS) [crafted]; Dreamweave Gloves (10019, -0.54 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.2 spell_power points (1.75 DPS) | yes | Gilded Cord (254037, -0.48 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.48 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.81 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.2 spell_power points (2.06 DPS) | yes | Crimson Silk Pantaloons (7062, -0.39 DPS) [crafted]; Kodohide Legguards (285338, -0.40 DPS) [world]; Abomination Skin Leggings (23173, -0.71 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (1.89 DPS) | yes | Skycaller's Leather Shoes (252532, -0.31 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.31 DPS) [crafted]; Mender's Leather Shoes (252533, -0.70 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (1.27 DPS) | yes | Ring of Forlorn Spirits (2043, -0.64 DPS) [quest]; Voodoo Band (1996, -0.71 DPS) [world]; Black Widow Band (6199, -0.71 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.71 DPS) | yes | Voodoo Band (1996, -0.15 DPS) [world]; Black Widow Band (6199, -0.15 DPS) [world]; Ring of Forlorn Spirits (2043, -0.69 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, -0.32 DPS) [world]; Mograine's Might (7723, -0.62 DPS) [dungeon]; Gut Ripper (2164, -3.11 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 4532310300103051-000000000000000000-5500000000000000)

Set DPS (verified): 72.0. Weights run: 3.2s. Verify run: 1.3s. 762 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.959 ± 0.033, crit=0.292 ± 0.010 per rating point (14 rating = 1%, 4.083 per %), hit=0.454 ± 0.004 per rating point (10 rating = 1%, 4.538 per %), spell_haste=2.210 ± 0.474, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.926 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 38.2 spell_power points (2.81 DPS) | yes | Dreamweave Circlet (10041, -0.56 DPS) [crafted]; Soothsayer's Headdress (17740, -0.63 DPS, sim-verified) [dungeon]; Charged Scorpid Helm (252581, -0.80 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 13.4 spell_power points (0.99 DPS) | yes | Mindburst Medallion (11196, -0.12 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.28 DPS) [quest]; Scorn's Icy Choker (23169, -1.14 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 34.2 spell_power points (2.51 DPS) | yes | Lead Surveyor's Mantle (11842, -0.42 DPS) [dungeon]; Kentic Amice (11624, -0.57 DPS) [dungeon]; Rotgrip Mantle (17732, -1.24 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.8 spell_power points (1.45 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.16 DPS) [dungeon]; Runecloth Cloak (13860, -0.23 DPS) [crafted]; Big Voodoo Cloak (8216, -0.45 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 38.2 spell_power points (2.81 DPS) | yes | Feathered Breastplate (8349, -0.70 DPS) [crafted]; Robe of the Magi (1716, -0.77 DPS) [world_drop]; Runecloth Tunic (13857, -0.78 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 19.6 spell_power points (1.44 DPS) | yes | Skycaller's Mail Bracers (252571, -0.21 DPS) [crafted]; Aristocratic Cuffs (12546, -0.38 DPS) [dungeon]; Skycaller's Leather Bracers (252542, -1.29 DPS, sim-verified) [crafted] |
| hands | Raider Handguards (272102) | Creeg Bothunk [vendor] | 29.3 spell_power points (2.15 DPS) | yes | Raider Handwraps (272098, -0.09 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -0.34 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.34 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 27.5 spell_power points (2.02 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.24 DPS) [dungeon]; Satyrmane Sash (17755, -0.29 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 32.6 spell_power points (2.40 DPS) | yes | Big Voodoo Pants (8202, -0.59 DPS) [crafted]; Turtle Scale Leggings (8185, -0.88 DPS) [crafted]; Red Mageweave Pants (10009, -2.36 DPS, sim-verified) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 25.5 spell_power points (1.88 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Greaves of Withering Despair (22240, -0.03 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.11 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.8 spell_power points (1.16 DPS) | yes | Brainlash (6440, -0.10 DPS) [dungeon]; Band of the Unicorn (7553, -0.20 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.28 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.7 spell_power points (1.16 DPS) | yes | Band of the Unicorn (7553, -0.20 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.27 DPS) [rep]; Brainlash (6440, -1.03 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (72.0 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (72.0 DPS) | yes | Smoking Heart of the Mountain (11811, -0.82 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-verified (72.0 DPS) | yes | Mechanic's Pipehammer (9604, -0.29 DPS) [quest]; Thorium Greatmace (250613, -0.35 DPS) [crafted]; Blade of Eternal Darkness (17780, -1.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handguards; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 4532310300103051-000000000000000000-5533220000000000)

Set DPS (verified): 116.5. Weights run: 3.2s. Verify run: 1.3s. 1751 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.668 ± 0.029, crit=0.345 ± 0.011 per rating point (14 rating = 1%, 4.833 per %), hit=0.535 ± 0.004 per rating point (10 rating = 1%, 5.353 per %), spell_haste=-4.011 ± 0.356, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.919 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (116.5 DPS) | yes | Crimson Felt Hat (18727, -0.13 DPS) [dungeon]; Soothsayer's Headdress (17740, -0.30 DPS) [dungeon]; Blue Dragonscale Helm (252604, -3.31 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.7 spell_power points (1.80 DPS) | yes | Orb of the Darkmoon (19426, -0.13 DPS) [quest]; Chains of the Lich (23125, -0.13 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.20 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 41.8 spell_power points (3.18 DPS) | yes | Burial Shawl (18681, -0.85 DPS) [dungeon]; Mantle of the Timbermaw (19050, -0.86 DPS) [crafted]; Darkspear Shoulderguards (272958, -1.08 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.7 spell_power points (2.03 DPS) | yes | Hide of the Wild (18510, -0.46 DPS) [crafted]; Spritecaster Cape (11623, -0.66 DPS) [dungeon]; Crystalline Threaded Cape (20697, -1.06 DPS, sim-verified) [world] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Robe of Everlasting Night (18385, -0.05 DPS) [dungeon]; Tunic of Undead Slaying (23089, -7.49 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Modest Armguards (18458, -0.66 DPS) [dungeon]; Sublime Wristguards (18497, -0.66 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -4.90 DPS, sim-verified) [world] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | 30.0 spell_power points (2.28 DPS) | yes | Hands of Power (13253, -0.00 DPS) [dungeon]; Raider Handwraps (272097, -0.30 DPS) [vendor]; Gloves of Undead Cleansing (23084, -0.41 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 45.0 spell_power points (3.42 DPS) | yes | Stormseeker's Girdle (272399, -0.99 DPS) [vendor]; Girdle of Insight (18504, -1.12 DPS) [crafted]; Belt of the Archmage (18405, -2.26 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 54.9 spell_power points (4.17 DPS) | yes | Sentinel's Lizardhide Pants (237817, -0.49 DPS) [vendor]; Red Dragonscale Leggings (252603, -1.02 DPS) [crafted]; Sentinel's Silk Leggings (237815, -1.17 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 31.7 spell_power points (2.41 DPS) | yes | Waterspout Boots (18322, -0.20 DPS) [dungeon]; Dragonrider Boots (18102, -0.23 DPS) [dungeon]; Omnicast Boots (11822, -0.28 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.51 DPS) [quest]; Maiden's Circle (13001, -0.51 DPS) [world_drop]; Naglering (11669, -4.91 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.23 DPS) [quest]; Maiden's Circle (13001, -0.23 DPS) [world_drop]; Naglering (11669, -4.77 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (+8.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.13 DPS) [world]; Hand of Edward the Odd (2243, -4.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Crackling Staff

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 22.4. Weights run: 2.5s. Verify run: 1.0s. 205 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.522 ± 0.008, crit=0.064 ± 0.002 per rating point (14 rating = 1%, 0.890 per %), hit=0.145 ± 0.001 per rating point (10 rating = 1%, 1.448 per %), spell_haste=-1.485 ± 0.118, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.940 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.15 DPS) [crafted]; Totemic Leather Hood (252448, -0.32 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 9.7 spell_power points (0.82 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.48 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.58 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.34 DPS) | yes | Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.32 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.00 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.24 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.1 spell_power points (0.26 DPS) | yes | Mindthrust Bracers (1974, -0.04 DPS) [dungeon]; Featherbead Bracers (15452, -0.04 DPS) [quest]; Owl Bracers (4796, -0.18 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 7.1 spell_power points (0.60 DPS) | yes | Serpent Gloves (5970, -0.01 DPS) [dungeon]; Gnoll Casting Gloves (892, -0.09 DPS) [world]; Pristine Gloves (253913, -0.13 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.08 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.45 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 13.2 spell_power points (1.11 DPS) | yes | Stormrider's Leather Pants (252502, -0.00 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.26 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.34 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.1 spell_power points (0.76 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.21 DPS) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.42 DPS) | yes | Sludge-Stained Band (286535, -0.17 DPS) [world]; Loop of Sacrifice (281673, -0.20 DPS) [quest]; Volcanic Rock Ring (12053, -0.29 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 3.1 spell_power points (0.26 DPS) | yes | Loop of Sacrifice (281673, -0.04 DPS) [quest]; Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; Sludge-Stained Band (286535, -0.31 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.67 DPS) | yes | Twisted Chanter's Staff (890, -0.23 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.23 DPS) [quest]; Channeler's Staff (4437, -0.32 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 35.2. Weights run: 2.8s. Verify run: 1.2s. 349 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.506 ± 0.014, crit=0.091 ± 0.002 per rating point (14 rating = 1%, 1.276 per %), hit=0.200 ± 0.001 per rating point (10 rating = 1%, 1.998 per %), spell_haste=-2.357 ± 0.272, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.937 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (1.01 DPS) | yes | Enchanter's Cowl (4322, -0.08 DPS) [crafted]; Holy Shroud (2721, -0.08 DPS) [world_drop]; Silk Headband (7050, -0.25 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.0 spell_power points (0.85 DPS) | yes | Crystal Starfire Medallion (5003, -0.68 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.68 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.98 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.6 spell_power points (1.15 DPS) | yes | Fairywing Mantle (9536, -0.25 DPS) [quest]; Magician's Mantle (12998, -0.34 DPS) [world_drop]; Death Speaker Mantle (6685, -0.46 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.42 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Cloak of Rot (4462, -0.08 DPS) [world]; Darkspear Raider's Cloak (272078, -0.08 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.6 spell_power points (1.32 DPS) | yes | Stormrider's Leather Tunic (252510, -0.21 DPS) [crafted]; Tree Bark Jacket (1486, -0.22 DPS) [dungeon]; Guardian Armor (4256, -0.55 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.76 DPS) | yes | Nightsky Wristbands (6407, -0.50 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.50 DPS) [quest]; Glowing Magical Bracelets (13106, -0.93 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.0 spell_power points (0.85 DPS) | yes | Jutebraid Gloves (10654, -0.13 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.25 DPS) [crafted]; Serpent Gloves (5970, -0.25 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.0 spell_power points (1.10 DPS) | yes | Moss Cinch (6911, -0.09 DPS) [dungeon]; Warsong Sash (16975, -0.17 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.70 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 15.6 spell_power points (1.32 DPS) | yes | Abomination Skin Leggings (23173, -0.21 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.21 DPS) [crafted]; Dark Ritual Leggings (270031, -0.96 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.5 spell_power points (0.89 DPS) | yes | Spidersilk Boots (4320, -0.13 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.17 DPS) [crafted]; Acidic Walkers (9454, -0.89 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.59 DPS) | yes | Black Widow Band (6199, -0.29 DPS) [world]; Snake Hoop (6750, -0.29 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.34 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.51 DPS) | yes | Snake Hoop (6750, -0.21 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.25 DPS) [dungeon]; Black Widow Band (6199, -0.71 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rhahk'Zor's Hammer (5187, -0.08 DPS) [dungeon]; Reef Axe (6905, -0.08 DPS) [dungeon]; Manual Crowd Pummeler (9449, -1.46 DPS, sim-verified) [dungeon] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-verified (35.2 DPS) | yes | Tome of the Darkspear Prophecy (272090, -0.08 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -0.08 DPS) [world]; Orb of Mystic Insight (249394, -0.58 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Seedcloud Buckler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 4532310300103051-000000000000000000-0000000000000000)

Set DPS (verified): 57.3. Weights run: 3.1s. Verify run: 1.1s. 555 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.014 ± 0.025, crit=0.233 ± 0.008 per rating point (14 rating = 1%, 3.267 per %), hit=0.356 ± 0.003 per rating point (10 rating = 1%, 3.558 per %), spell_haste=-3.775 ± 0.284, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.929 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 23.2 spell_power points (1.83 DPS) | yes | Augural Shroud (2620, -0.16 DPS) [world]; Spellpower Goggles Xtreme (10502, -0.17 DPS) [crafted]; Corpseshroud (10574, -0.31 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.1 spell_power points (1.03 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.23 DPS) [quest]; Triune Amulet (7722, -0.47 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.47 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.2 spell_power points (1.59 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Death Speaker Mantle (6685, -0.24 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 18.1 spell_power points (1.43 DPS) | yes | Long Silken Cloak (4326, -0.56 DPS) [crafted]; Guardian Cloak (5965, -0.56 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.70 DPS, sim-verified) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (57.3 DPS) | yes | Robe of Power (7054, -0.08 DPS) [crafted]; Big Voodoo Robe (8200, -0.31 DPS) [crafted]; Robe of the Magi (1716, -0.60 DPS, sim-verified) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 12.1 spell_power points (0.96 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Guardian Leather Bracers (4260, -0.00 DPS) [crafted]; Turtle Scale Bracers (8198, -0.08 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (1.89 DPS) | yes | Red Mageweave Gloves (10018, -0.23 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.31 DPS) [crafted]; Dreamweave Gloves (10019, -0.88 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 22.2 spell_power points (1.75 DPS) | yes | Gilded Cord (254037, -0.48 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.48 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.02 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 26.2 spell_power points (2.06 DPS) | yes | Kodohide Legguards (285338, -0.40 DPS) [world]; Crimson Silk Pantaloons (7062, -0.55 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.71 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (1.89 DPS) | yes | Skycaller's Leather Shoes (252532, -0.31 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.31 DPS) [crafted]; Mender's Leather Shoes (252533, -0.70 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (1.27 DPS) | yes | Ogremind Ring (1993, -0.71 DPS) [world_drop]; Voodoo Band (1996, -0.71 DPS) [world]; Black Widow Band (6199, -0.71 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.71 DPS) | yes | Ogremind Ring (1993, -0.15 DPS) [world_drop]; Voodoo Band (1996, -0.15 DPS) [world]; Black Widow Band (6199, -0.72 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.30 DPS) [dungeon]; Windweaver Staff (7757, -0.38 DPS) [dungeon]; Gut Ripper (2164, -3.22 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Radiant Silver Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 4532310300103051-000000000000000000-5500000000000000)

Set DPS (verified): 71.9. Weights run: 3.2s. Verify run: 1.2s. 704 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.959 ± 0.033, crit=0.292 ± 0.010 per rating point (14 rating = 1%, 4.083 per %), hit=0.454 ± 0.004 per rating point (10 rating = 1%, 4.538 per %), spell_haste=2.210 ± 0.474, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.926 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamweave Circlet (10041, -0.50 DPS) [crafted]; Blood Guard's Pulsing Helmet (220848, -0.68 DPS) [vendor]; Red Mageweave Headband (10033, -0.70 DPS, sim-verified) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 13.4 spell_power points (0.99 DPS) | yes | Mindburst Medallion (11196, -0.12 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.28 DPS) [quest]; Scorn's Icy Choker (23169, -0.86 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 34.2 spell_power points (2.51 DPS) | yes | Lead Surveyor's Mantle (11842, -0.42 DPS) [dungeon]; Kentic Amice (11624, -0.57 DPS) [dungeon]; Rotgrip Mantle (17732, -0.89 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 20.6 spell_power points (1.52 DPS) | yes | Spritecaster Cape (11623, -0.06 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.22 DPS) [dungeon]; Runecloth Cloak (13860, -0.29 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 38.2 spell_power points (2.81 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.70 DPS, sim-verified) [vendor]; Feathered Breastplate (8349, -0.70 DPS) [crafted]; Robe of the Magi (1716, -0.77 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 19.6 spell_power points (1.44 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.21 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Skycaller's Leather Gauntlets (252550, -0.25 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.25 DPS) [crafted]; Raider Handguards (272102, -0.88 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 27.5 spell_power points (2.02 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.24 DPS) [dungeon]; Satyrmane Sash (17755, -0.29 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Red Mageweave Pants (10009, -0.52 DPS) [crafted]; Big Voodoo Pants (8202, -0.59 DPS) [crafted]; Stone Guard's Pulsing Legplates (220847, -1.22 DPS, sim-verified) [vendor] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 25.5 spell_power points (1.88 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Greaves of Withering Despair (22240, -0.03 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.11 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.8 spell_power points (1.16 DPS) | yes | Brainlash (6440, -0.10 DPS) [dungeon]; Band of the Unicorn (7553, -0.20 DPS) [world_drop]; Advisor's Ring (19519, -0.28 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.7 spell_power points (1.16 DPS) | yes | Brainlash (6440, -0.10 DPS) [dungeon]; Band of the Unicorn (7553, -0.20 DPS) [world_drop]; Advisor's Ring (19519, -0.27 DPS) [rep] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Thorium Greatmace (250613, -0.35 DPS) [crafted]; Spellshifter Rod (9527, -0.42 DPS) [quest]; Blade of Eternal Darkness (17780, -0.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger2: Cyclopean Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 4532310300103051-000000000000000000-5533220000000000)

Set DPS (verified): 114.0. Weights run: 3.2s. Verify run: 1.1s. 1672 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.668 ± 0.029, crit=0.345 ± 0.011 per rating point (14 rating = 1%, 4.833 per %), hit=0.535 ± 0.004 per rating point (10 rating = 1%, 5.353 per %), spell_haste=-4.011 ± 0.356, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.919 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Coif of The Five Thunders (227002) | Saving the Best for Last [quest] | 42.5 spell_power points (3.23 DPS) | yes | Warlord's Mail Helm (231663, -0.10 DPS) [pvp]; Living Crown (252561, -0.42 DPS) [crafted]; Blue Dragonscale Helm (252604, -1.17 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.7 spell_power points (1.80 DPS) | yes | Orb of the Darkmoon (19426, -0.13 DPS) [quest]; Beads of Ogre Mojo (22149, -0.20 DPS) [quest]; Chains of the Lich (23125, -1.94 DPS, sim-verified) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 41.8 spell_power points (3.18 DPS) | yes | Warlord's Mail Spaulders (231659, -0.41 DPS) [pvp]; Pauldrons of The Five Thunders (227003, -0.72 DPS) [quest]; Darkspear Shoulderguards (272958, -1.44 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.7 spell_power points (2.03 DPS) | yes | Hide of the Wild (18510, -0.46 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.66 DPS) [quest]; Crystalline Threaded Cape (20697, -1.05 DPS, sim-verified) [world] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (114.0 DPS) | yes | Legionnaire's Mail Breastplate (227165, +0.00 DPS) [vendor]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Tunic of Undead Slaying (23089, -8.50 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-verified (114.0 DPS) | yes | Modest Armguards (18458, -0.66 DPS) [dungeon]; Sublime Wristguards (18497, -0.66 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -6.16 DPS, sim-verified) [world] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | 30.0 spell_power points (2.28 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Hands of Power (13253, -0.00 DPS) [dungeon]; General's Mail Gloves (231666, -0.09 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 45.0 spell_power points (3.42 DPS) | yes | Stormseeker's Girdle (272399, -0.99 DPS) [vendor]; Girdle of Insight (18504, -1.12 DPS) [crafted]; Belt of the Archmage (18405, -3.85 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 54.9 spell_power points (4.17 DPS) | yes | General's Mail Leggings (231664, -0.46 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -0.49 DPS) [vendor]; Sentinel's Silk Leggings (237815, -2.46 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 31.7 spell_power points (2.41 DPS) | yes | General's Mail Sabatons (231661, -0.03 DPS) [vendor]; Waterspout Boots (18322, -0.20 DPS) [dungeon]; Dragonrider Boots (18102, -0.23 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (114.0 DPS) | yes | Eye of Orgrimmar (12545, -0.51 DPS) [quest]; Maiden's Circle (13001, -0.51 DPS) [world_drop]; Naglering (11669, -6.14 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (114.0 DPS) | yes | Eye of Orgrimmar (12545, -0.23 DPS) [quest]; Maiden's Circle (13001, -0.23 DPS) [world_drop]; Naglering (11669, -6.27 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (114.0 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (114.0 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (114.0 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.21 DPS) [dungeon]; Hand of Edward the Odd (2243, -6.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Coif of The Five Thunders; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Amethyst War Staff

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

