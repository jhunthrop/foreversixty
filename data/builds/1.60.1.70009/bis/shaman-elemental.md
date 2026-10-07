# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 31.0. Weights run: 2.6s. Verify run: 1.1s. 225 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.746 ± 0.012, crit=0.058 ± 0.002 per rating point (14 rating = 1%, 0.815 per %), hit=0.146 ± 0.001 per rating point (10 rating = 1%, 1.465 per %), spell_haste=-2.107 ± 0.122, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.719 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.00 DPS) [crafted]; Totemic Leather Hood (252448, -0.39 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.7 spell_power points (1.30 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.31 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.86 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Feyscale Cloak (6632, -0.11 DPS) [dungeon]; Black Whelp Cloak (7283, -0.11 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.41 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.05 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.29 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.08 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.08 DPS) [vendor]; Owl Bracers (4796, -0.81 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.0 spell_power points (0.89 DPS) | yes | Serpent Gloves (5970, -0.11 DPS) [dungeon]; Windfelt Gloves (5630, -0.19 DPS) [quest]; Pristine Gloves (253913, -0.19 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.08 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.11 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.53 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Dreamer's Leggings (270016, -0.25 DPS) [quest]; Wisdom's Leather Pants (252503, -0.33 DPS) [crafted]; Abomination Skin Leggings (23173, -0.44 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 spell_power points (1.11 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.25 DPS) [crafted]; Totemic Leather Boots (252442, -0.44 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 spell_power points (0.72 DPS) | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Sludge-Stained Band (286535, -0.39 DPS) [world]; Volcanic Rock Ring (12053, -0.47 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.55 DPS) | yes | Sludge-Stained Band (286535, -0.22 DPS) [world]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.68 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.89 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.23 DPS) [world]; Lesser Staff of the Spire (1300, -0.39 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 225, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler; 4821 Bear Buckler; 4822 Owl's Disk; 4964 Goblin Smasher

### Band 30 (dwarf, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 51.7. Weights run: 2.9s. Verify run: 1.1s. 366 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.893 ± 0.015, crit=0.094 ± 0.003 per rating point (14 rating = 1%, 1.321 per %), hit=0.219 ± 0.002 per rating point (10 rating = 1%, 2.185 per %), spell_haste=-1.476 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.698 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 14.9 spell_power points (1.67 DPS) | yes | Enduring Cap (3020, +0.00 DPS, sim-verified) [world_drop]; Totemic Leather Helm (252456, -0.33 DPS) [crafted]; Holy Shroud (2721, -0.44 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.4 spell_power points (1.38 DPS) | yes | Crystal Starfire Medallion (5003, -0.98 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.98 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.34 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.0 spell_power points (1.91 DPS) | yes | Death Speaker Mantle (6685, -0.14 DPS) [dungeon]; Fairywing Mantle (9536, -0.34 DPS) [quest]; Magician's Mantle (12998, -0.45 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.1 spell_power points (0.80 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Repairman's Cape (9605, -0.06 DPS) [quest]; Resilient Cape (14400, -0.20 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 20.6 spell_power points (2.31 DPS) | yes | Death Speaker Robes (6682, -0.42 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.59 DPS) [crafted]; Guardian Armor (4256, -0.64 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.01 DPS) | yes | Nightsky Wristbands (6407, -0.41 DPS) [world_drop]; Technician's Bracers (270042, -0.41 DPS) [quest]; Glowing Magical Bracelets (13106, -1.34 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 13.8 spell_power points (1.55 DPS) | yes | Truefaith Gloves (7049, -0.69 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.75 DPS) [dungeon]; Stormrider's Leather Gloves (252498, -0.84 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.4 spell_power points (1.72 DPS) | yes | Guardian Belt (4258, -0.35 DPS) [crafted]; Crimson Silk Belt (7055, -0.35 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.74 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 19.8 spell_power points (2.22 DPS) | yes | Stormrider's Leather Pants (252502, -0.50 DPS) [crafted]; Guardian Pants (5962, -0.52 DPS) [crafted]; Abomination Skin Leggings (23173, -0.81 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.3 spell_power points (1.48 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.31 DPS) [crafted]; Acidic Walkers (9454, -1.25 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.78 DPS) | yes | Black Widow Band (6199, -0.08 DPS) [world]; Snake Hoop (6750, -0.08 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.11 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.8 spell_power points (0.76 DPS) | yes | Black Widow Band (6199, -0.06 DPS) [world]; Snake Hoop (6750, -0.06 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.09 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-verified (51.7 DPS) | yes | Glimmering Staff (249392, -1.53 DPS) [crafted]; Scorn's Focal Dagger (23168, -1.62 DPS) [dungeon]; Manual Crowd Pummeler (9449, -3.77 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 4532310300103051-000000000000000000-0000000000000000)

Set DPS (verified): 74.9. Weights run: 3.2s. Verify run: 1.2s. 592 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.139 ± 0.032, crit=0.212 ± 0.009 per rating point (14 rating = 1%, 2.963 per %), hit=0.337 ± 0.004 per rating point (10 rating = 1%, 3.370 per %), spell_haste=-2.948 ± 0.342, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.677 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 24.9 spell_power points (2.73 DPS) | yes | Augural Shroud (2620, -0.28 DPS) [world]; Corpseshroud (10574, -0.36 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.43 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 spell_power points (1.51 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.27 DPS) [quest]; Triune Amulet (7722, -0.64 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.64 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.8 spell_power points (2.38 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Death Speaker Mantle (6685, -0.36 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 19.3 spell_power points (2.10 DPS) | yes | Long Silken Cloak (4326, -0.83 DPS) [crafted]; Guardian Cloak (5965, -0.83 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.14 DPS, sim-verified) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Robe of Power (7054, -0.06 DPS) [crafted]; Big Voodoo Robe (8200, -0.36 DPS) [crafted]; Robe of the Magi (1716, -1.06 DPS, sim-verified) [world_drop] |
| wrist | Turtle Scale Bracers (8198) | Leatherworking [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Windchaser Cuffs (14429, -0.16 DPS) [world_drop]; Green Whelp Bracers (7386, -0.28 DPS) [crafted]; Guardian Leather Bracers (4260, -1.00 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.62 DPS) | yes | Red Mageweave Gloves (10018, -0.18 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.30 DPS) [crafted]; Dreamweave Gloves (10019, -1.02 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.1 spell_power points (2.63 DPS) | yes | Highlander's Cloth Girdle (20098, -0.60 DPS) [rep]; Gilded Cord (254037, -0.76 DPS) [crafted]; Highlander's Lizardhide Girdle (20104, -0.76 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.7 spell_power points (3.02 DPS) | yes | Kodohide Legguards (285338, -0.56 DPS) [world]; Crimson Silk Pantaloons (7062, -0.84 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.04 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.62 DPS) | yes | Skycaller's Mail Boots (252563, -0.33 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.62 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -0.88 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.8 spell_power points (1.84 DPS) | yes | Ring of Forlorn Spirits (2043, -0.97 DPS) [quest]; Voodoo Band (1996, -0.97 DPS) [world]; Black Widow Band (6199, -0.97 DPS) [world] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.98 DPS) | yes | Ring of Forlorn Spirits (2043, -0.11 DPS) [quest]; Voodoo Band (1996, -0.11 DPS) [world]; Black Widow Band (6199, -0.11 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, -0.51 DPS) [world]; Mograine's Might (7723, -0.71 DPS) [dungeon]; Gut Ripper (2164, -3.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Turtle Scale Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 4532310300103051-000000000000000000-5500000000000000)

Set DPS (verified): 106.7. Weights run: 3.2s. Verify run: 1.4s. 762 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.020 ± 0.047, crit=0.299 ± 0.014 per rating point (14 rating = 1%, 4.182 per %), hit=0.478 ± 0.006 per rating point (10 rating = 1%, 4.781 per %), spell_haste=-4.444 ± 0.357, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.645 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 39.4 spell_power points (4.12 DPS) | yes | Soothsayer's Headdress (17740, -0.11 DPS) [dungeon]; Dreamweave Circlet (10041, -0.86 DPS) [crafted]; Charged Scorpid Helm (252581, -1.16 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.3 spell_power points (1.49 DPS) | yes | Mindburst Medallion (11196, -0.23 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.43 DPS) [quest]; Scorn's Icy Choker (23169, -1.46 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.4 spell_power points (3.70 DPS) | yes | Lead Surveyor's Mantle (11842, -0.64 DPS) [dungeon]; Kentic Amice (11624, -0.85 DPS) [dungeon]; Rotgrip Mantle (17732, -1.79 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.1 spell_power points (2.10 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.20 DPS) [dungeon]; Runecloth Cloak (13860, -0.31 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.61 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 39.4 spell_power points (4.12 DPS) | yes | Feathered Breastplate (8349, -1.07 DPS) [crafted]; Runecloth Robe (13858, -1.16 DPS) [crafted]; Runecloth Tunic (13857, -1.17 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.2 spell_power points (2.11 DPS) | yes | Skycaller's Leather Bracers (252542, -0.32 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.32 DPS) [crafted]; Aristocratic Cuffs (12546, -0.51 DPS) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (106.7 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.44 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.44 DPS) [crafted]; Raider Handguards (272102, -1.63 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 28.2 spell_power points (2.95 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.30 DPS) [dungeon]; Satyrmane Sash (17755, -0.42 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 33.2 spell_power points (3.47 DPS) | yes | Big Voodoo Pants (8202, -0.84 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -1.24 DPS) [dungeon]; Red Mageweave Pants (10009, -2.75 DPS, sim-verified) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 26.2 spell_power points (2.74 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Greaves of Withering Despair (22240, -0.02 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.23 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.1 spell_power points (1.69 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.33 DPS) [world_drop]; Mindseye Circle (10634, -0.41 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (1.68 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.33 DPS) [world_drop]; Mindseye Circle (10634, -0.41 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mechanic's Pipehammer (9604, -0.57 DPS) [quest]; Spellshifter Rod (9527, -0.64 DPS) [quest]; Blade of Eternal Darkness (17780, -3.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 4532310300103051-000000000000000000-5533220000000000)

Set DPS (verified): 171.7. Weights run: 3.2s. Verify run: 1.3s. 1751 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.460 ± 0.057, crit=0.390 ± 0.018 per rating point (14 rating = 1%, 5.464 per %), hit=0.590 ± 0.007 per rating point (10 rating = 1%, 5.902 per %), spell_haste=-5.251 ± 0.522, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.647 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Black Dragonscale Helm (252605) | Leatherworking [crafted] | 61.0 spell_power points (6.37 DPS) | yes | Magister's Crown (16686, -0.64 DPS) [dungeon]; Living Crown (252561, -1.01 DPS) [crafted]; Blue Dragonscale Helm (252604, -4.10 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.0 spell_power points (3.55 DPS) | yes | Beads of Ogre Mojo (22149, -0.36 DPS) [quest]; Lady Maye's Pendant (14558, -0.65 DPS) [world_drop]; Archlight Talisman (15856, -0.77 DPS) [quest] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 55.4 spell_power points (5.79 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -0.11 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.42 DPS) [vendor]; Darkspear Shoulders (272104, -0.42 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.6 spell_power points (3.51 DPS) | yes | Hide of the Wild (18510, -0.52 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.81 DPS) [world]; Royal Tribunal Cloak (13376, -1.07 DPS) [dungeon] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (171.7 DPS) | yes | Vest of Elements (16666, -0.07 DPS) [dungeon]; Chestplate of Tranquility (18373, -0.07 DPS) [dungeon]; Tunic of Undead Slaying (23089, -11.51 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-verified (171.7 DPS) | yes | Modest Armguards (18458, -0.74 DPS) [dungeon]; Sublime Wristguards (18497, -0.74 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -8.30 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 47.4 spell_power points (4.96 DPS) | yes | Raider Handguards (272102, -1.00 DPS) [vendor]; Hands of Power (13253, -1.32 DPS) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 63.0 spell_power points (6.59 DPS) | yes | Belt of the Archmage (18405, -1.48 DPS) [crafted]; Girdle of Insight (18504, -1.51 DPS) [crafted]; Stormseeker's Girdle (272399, -4.21 DPS, sim-verified) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 69.7 spell_power points (7.29 DPS) | yes | Red Dragonscale Leggings (252603, -0.32 DPS) [crafted]; Sentinel's Silk Leggings (237815, -1.28 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.28 DPS) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 44.4 spell_power points (4.64 DPS) | yes | Dragonrider Boots (18102, -0.31 DPS) [dungeon]; Omnicast Boots (11822, -0.72 DPS) [dungeon]; Waterspout Boots (18322, -1.11 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (171.7 DPS) | yes | Songstone of Ironforge (12543, -1.03 DPS) [quest]; Maiden's Circle (13001, -1.03 DPS) [world_drop]; Naglering (11669, -7.81 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (171.7 DPS) | yes | Songstone of Ironforge (12543, -0.31 DPS) [quest]; Maiden's Circle (13001, -0.31 DPS) [world_drop]; Naglering (11669, -7.56 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (171.7 DPS) | yes | Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor]; Second Wind (11819, -5.34 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (171.7 DPS) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -0.73 DPS) [vendor]; Serenity Field (272439, -1.57 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (171.7 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, -1.13 DPS) [pvp]; Hand of Edward the Odd (2243, -8.16 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Black Dragonscale Helm; neck: Amulet of the Dawn; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Crackling Staff

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 30.2. Weights run: 2.6s. Verify run: 1.1s. 205 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.746 ± 0.012, crit=0.058 ± 0.002 per rating point (14 rating = 1%, 0.815 per %), hit=0.146 ± 0.001 per rating point (10 rating = 1%, 1.465 per %), spell_haste=-2.107 ± 0.122, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.719 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.00 DPS) [crafted]; Totemic Leather Hood (252448, -0.40 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Roadwatcher's Confidence (281265, -0.60 DPS, sim-verified) [quest] |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.7 spell_power points (1.30 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.72 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.86 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 spell_power points (0.47 DPS) | yes | Heavy Woolen Cloak (4311, -0.03 DPS) [crafted]; Feyscale Cloak (6632, -0.14 DPS) [dungeon]; Black Whelp Cloak (7283, -0.14 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.05 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.30 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.5 spell_power points (0.50 DPS) | yes | Mindthrust Bracers (1974, -0.08 DPS) [dungeon]; Featherbead Bracers (15452, -0.08 DPS) [quest]; Owl Bracers (4796, -0.54 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.0 spell_power points (0.89 DPS) | yes | Serpent Gloves (5970, -0.19 DPS, sim-verified) [dungeon]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Gnoll Casting Gloves (892, -0.22 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.08 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.11 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.54 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.0 spell_power points (1.66 DPS) | yes | Stormrider's Leather Pants (252502, -0.05 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.39 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.50 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 spell_power points (1.11 DPS) | yes | Wisdom's Leather Boots (252444, -0.25 DPS) [crafted]; Totemic Leather Boots (252442, -0.44 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.47 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.55 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Sludge-Stained Band (286535, -0.22 DPS) [world]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.5 spell_power points (0.50 DPS) | yes | Sludge-Stained Band (286535, -0.16 DPS) [world]; Volcanic Rock Ring (12053, -0.25 DPS) [world_drop]; Loop of Sacrifice (281673, -0.56 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.89 DPS) | yes | Twisted Chanter's Staff (890, -0.06 DPS) [world_drop]; Channeler's Staff (4437, -0.23 DPS) [world]; Gnarled Necromancer's Staff (251534, -0.33 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Scholarly Pendant; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 205, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 7188 Stormwind Guard Shield; 14389 Durability Shoulderpads; 14705 Brackwater Chain Shield

### Band 30 (orc, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 48.7. Weights run: 2.9s. Verify run: 1.0s. 349 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.893 ± 0.015, crit=0.094 ± 0.003 per rating point (14 rating = 1%, 1.321 per %), hit=0.219 ± 0.002 per rating point (10 rating = 1%, 2.185 per %), spell_haste=-1.476 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.698 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 14.9 spell_power points (1.67 DPS) | yes | Enduring Cap (3020, +0.00 DPS, sim-verified) [world_drop]; Totemic Leather Helm (252456, -0.33 DPS) [crafted]; Holy Shroud (2721, -0.44 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.4 spell_power points (1.38 DPS) | yes | Crystal Starfire Medallion (5003, -0.98 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.98 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.39 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.0 spell_power points (1.91 DPS) | yes | Fairywing Mantle (9536, -0.34 DPS) [quest]; Death Speaker Mantle (6685, -0.37 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.45 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.1 spell_power points (0.80 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Resilient Cape (14400, -0.20 DPS) [world_drop]; Hillman's Cloak (3719, -0.24 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 20.6 spell_power points (2.31 DPS) | yes | Death Speaker Robes (6682, -0.42 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.59 DPS) [crafted]; Guardian Armor (4256, -0.76 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.01 DPS) | yes | Nightsky Wristbands (6407, -0.41 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.41 DPS) [quest]; Glowing Magical Bracelets (13106, -1.27 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.8 spell_power points (1.21 DPS) | yes | Jutebraid Gloves (10654, -0.04 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.25 DPS) [crafted]; Truefaith Gloves (7049, -0.35 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.4 spell_power points (1.72 DPS) | yes | Guardian Belt (4258, -0.35 DPS) [crafted]; Crimson Silk Belt (7055, -0.35 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.88 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 19.8 spell_power points (2.22 DPS) | yes | Stormrider's Leather Pants (252502, -0.50 DPS) [crafted]; Guardian Pants (5962, -0.52 DPS) [crafted]; Abomination Skin Leggings (23173, -0.92 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.3 spell_power points (1.48 DPS) | yes | Spidersilk Boots (4320, -0.30 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.31 DPS) [crafted]; Acidic Walkers (9454, -1.57 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.78 DPS) | yes | Snake Hoop (6750, -0.08 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.11 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.3 spell_power points (0.70 DPS) | yes | Snake Hoop (6750, +0.00 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.03 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | sim-verified (48.7 DPS) | yes | Scorn's Focal Dagger (23168, -0.09 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.10 DPS) [quest]; Manual Crowd Pummeler (9449, -1.83 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; main_hand: Glimmering Staff

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 4532310300103051-000000000000000000-0000000000000000)

Set DPS (verified): 76.5. Weights run: 3.2s. Verify run: 1.2s. 555 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.139 ± 0.032, crit=0.212 ± 0.009 per rating point (14 rating = 1%, 2.963 per %), hit=0.337 ± 0.004 per rating point (10 rating = 1%, 3.370 per %), spell_haste=-2.948 ± 0.342, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.677 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 24.9 spell_power points (2.73 DPS) | yes | Augural Shroud (2620, -0.28 DPS) [world]; Corpseshroud (10574, -0.36 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.43 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 spell_power points (1.51 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.27 DPS) [quest]; Triune Amulet (7722, -0.64 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.64 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 21.8 spell_power points (2.38 DPS) | yes | Green Silken Shoulders (7057, -0.14 DPS) [crafted]; Bloodmage Mantle (7684, -0.28 DPS) [dungeon]; Death Speaker Mantle (6685, -0.36 DPS) [dungeon] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 19.3 spell_power points (2.10 DPS) | yes | Long Silken Cloak (4326, -0.83 DPS) [crafted]; Guardian Cloak (5965, -0.83 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.87 DPS, sim-verified) [vendor] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (76.5 DPS) | yes | Robe of Power (7054, -0.06 DPS) [crafted]; Big Voodoo Robe (8200, -0.36 DPS) [crafted]; Robe of the Magi (1716, -1.12 DPS, sim-verified) [world_drop] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 13.1 spell_power points (1.43 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Guardian Leather Bracers (4260, -0.03 DPS) [crafted]; Turtle Scale Bracers (8198, -0.15 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.62 DPS) | yes | Red Mageweave Gloves (10018, -0.18 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.30 DPS) [crafted]; Dreamweave Gloves (10019, -0.90 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 24.1 spell_power points (2.63 DPS) | yes | Gilded Cord (254037, -0.76 DPS) [crafted]; Highlander's Mail Girdle (20119, -0.76 DPS) [vendor]; Defiler's Cloth Girdle (20166, -1.36 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 27.7 spell_power points (3.02 DPS) | yes | Crimson Silk Pantaloons (7062, -0.53 DPS) [crafted]; Kodohide Legguards (285338, -0.56 DPS) [world]; Abomination Skin Leggings (23173, -1.04 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.62 DPS) | yes | Skycaller's Mail Boots (252563, -0.33 DPS) [crafted]; Mender's Leather Shoes (252533, -0.88 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.93 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.8 spell_power points (1.84 DPS) | yes | Ogremind Ring (1993, -0.97 DPS) [world_drop]; Voodoo Band (1996, -0.97 DPS) [world]; Black Widow Band (6199, -0.97 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.98 DPS) | yes | Ogremind Ring (1993, -0.11 DPS) [world_drop]; Voodoo Band (1996, -0.11 DPS) [world]; Black Widow Band (6199, -1.79 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.19 DPS) [dungeon]; Windweaver Staff (7757, -0.32 DPS) [dungeon]; Gut Ripper (2164, -4.08 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Radiant Silver Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 4532310300103051-000000000000000000-5500000000000000)

Set DPS (verified): 106.8. Weights run: 3.2s. Verify run: 1.3s. 704 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.020 ± 0.047, crit=0.299 ± 0.014 per rating point (14 rating = 1%, 4.182 per %), hit=0.478 ± 0.006 per rating point (10 rating = 1%, 4.781 per %), spell_haste=-4.444 ± 0.357, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.645 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 39.4 spell_power points (4.12 DPS) | yes | Soothsayer's Headdress (17740, -0.11 DPS) [dungeon]; Dreamweave Circlet (10041, -0.86 DPS) [crafted]; Blood Guard's Pulsing Helmet (220848, -1.09 DPS) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.3 spell_power points (1.49 DPS) | yes | Mindburst Medallion (11196, -0.23 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.43 DPS) [quest]; Scorn's Icy Choker (23169, -0.98 DPS, sim-verified) [dungeon] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.4 spell_power points (3.70 DPS) | yes | Lead Surveyor's Mantle (11842, -0.64 DPS) [dungeon]; Kentic Amice (11624, -0.85 DPS) [dungeon]; Rotgrip Mantle (17732, -1.73 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.2 spell_power points (2.21 DPS) | yes | Spritecaster Cape (11623, -0.11 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.31 DPS) [dungeon]; Runecloth Cloak (13860, -0.42 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 39.4 spell_power points (4.12 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.46 DPS) [vendor]; Feathered Breastplate (8349, -1.07 DPS) [crafted]; Runecloth Robe (13858, -1.16 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.2 spell_power points (2.11 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.32 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (106.8 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.44 DPS) [crafted]; Skycaller's Mail Gauntlets (252585, -0.44 DPS) [crafted]; Raider Handguards (272102, -2.02 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 28.2 spell_power points (2.95 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Dawnspire Cord (12466, -0.30 DPS) [dungeon]; Satyrmane Sash (17755, -0.42 DPS) [dungeon] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 35.0 spell_power points (3.66 DPS) | yes | Spellshock Leggings (9484, -0.19 DPS) [dungeon]; Red Mageweave Pants (10009, -0.92 DPS) [crafted]; Big Voodoo Pants (8202, -1.03 DPS) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 26.2 spell_power points (2.74 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Greaves of Withering Despair (22240, -0.02 DPS) [dungeon]; Earthen Silk Slippers (254013, -0.23 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.1 spell_power points (1.69 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.33 DPS) [world_drop]; Mindseye Circle (10634, -0.41 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.1 spell_power points (1.68 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.33 DPS) [world_drop]; Mindseye Circle (10634, -0.41 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of the Guard Captain (19120, -0.10 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of the Guard Captain (19120, -0.28 DPS) [quest]; Smoking Heart of the Mountain (11811, -1.06 DPS, sim-verified) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.64 DPS) [quest]; Thorium Greatmace (250613, -0.69 DPS) [crafted]; Blade of Eternal Darkness (17780, -2.64 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Stone Guard's Pulsing Legplates; feet: Skycaller's Leather Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 4532310300103051-000000000000000000-5533220000000000)

Set DPS (verified): 168.1. Weights run: 3.2s. Verify run: 1.3s. 1672 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.460 ± 0.057, crit=0.390 ± 0.018 per rating point (14 rating = 1%, 5.464 per %), hit=0.590 ± 0.007 per rating point (10 rating = 1%, 5.902 per %), spell_haste=-5.251 ± 0.522, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.647 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Coif of The Five Thunders (227002) | Saving the Best for Last [quest] | 63.0 spell_power points (6.58 DPS) | yes | Black Dragonscale Helm (252605, -0.21 DPS) [crafted]; Warlord's Mail Helm (231663, -0.55 DPS) [pvp]; Blue Dragonscale Helm (252604, -0.60 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 34.0 spell_power points (3.55 DPS) | yes | Beads of Ogre Mojo (22149, -0.36 DPS) [quest]; Lady Maye's Pendant (14558, -0.65 DPS) [world_drop]; Archlight Talisman (15856, -0.77 DPS) [quest] |
| shoulder | Darkspear Shoulderguards (272958) | Creeg Bothunk [vendor] | 55.4 spell_power points (5.79 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -0.11 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.42 DPS) [vendor]; Darkspear Shoulders (272104, -0.42 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 33.6 spell_power points (3.51 DPS) | yes | Hide of the Wild (18510, -0.52 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.81 DPS) [world]; Deep Woodlands Cloak (19121, -0.88 DPS) [quest] |
| chest | Magister's Robes (16688) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (168.1 DPS) | yes | Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Vest of The Five Thunders (227004, -0.04 DPS) [quest]; Tunic of Undead Slaying (23089, -10.38 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-verified (168.1 DPS) | yes | Modest Armguards (18458, -0.74 DPS) [dungeon]; Sublime Wristguards (18497, -0.74 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -7.36 DPS, sim-verified) [world] |
| hands | Raider Handwraps (272097) | Creeg Bothunk [vendor] | 47.4 spell_power points (4.96 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; General's Mail Gloves (231666, -0.63 DPS) [vendor]; Raider Handguards (272101, -2.55 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 63.0 spell_power points (6.59 DPS) | yes | Belt of the Archmage (18405, -1.48 DPS) [crafted]; Girdle of Insight (18504, -1.51 DPS) [crafted]; Stormseeker's Girdle (272399, -3.61 DPS, sim-verified) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 69.7 spell_power points (7.29 DPS) | yes | General's Mail Leggings (231664, -0.13 DPS) [pvp]; Red Dragonscale Leggings (252603, -0.32 DPS) [crafted]; Sentinel's Silk Leggings (237815, -1.28 DPS) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 44.4 spell_power points (4.64 DPS) | yes | General's Mail Sabatons (231661, +0.00 DPS) [vendor]; Dragonrider Boots (18102, -0.31 DPS) [dungeon]; Omnicast Boots (11822, -0.72 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (168.1 DPS) | yes | Eye of Orgrimmar (12545, -1.03 DPS) [quest]; Maiden's Circle (13001, -1.03 DPS) [world_drop]; Naglering (11669, -7.78 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (168.1 DPS) | yes | Eye of Orgrimmar (12545, -0.31 DPS) [quest]; Maiden's Circle (13001, -0.31 DPS) [world_drop]; Naglering (11669, -6.92 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (168.1 DPS) | yes | Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -4.67 DPS, sim-verified) [dungeon] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (168.1 DPS) | yes | Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, -0.63 DPS) [quest]; Weakness Analyzer (272438, -0.73 DPS) [vendor] |
| main_hand | Trindlehaven Staff (13161) | Blackrock Spire: Overlord Wyrmthalak [dungeon] | sim-verified (168.1 DPS) | yes | High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -5.84 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Coif of The Five Thunders; neck: Amulet of the Dawn; shoulder: Darkspear Shoulderguards; back: Arcanoweave Cloak; chest: Magister's Robes; wrist: Dryad's Wrist Bindings; hands: Raider Handwraps; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Briarwood Reed; main_hand: Trindlehaven Staff

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

