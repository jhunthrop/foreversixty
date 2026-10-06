# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 28.3. Weights run: 2.7s. Verify run: 0.9s. 193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.249 ± 0.009, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), hit=0.128 ± 0.001 per rating point (10 rating = 1%, 1.284 per %), spell_haste=-1.246 ± 0.100, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.006 ± 0.000, arcane_power=0.994 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Goggles (4373, -0.29 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.31 DPS) [crafted]; Trapper's Leather Hood (252505, -1.27 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.2 spell_power points (1.69 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.65 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.27 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.7 spell_power points (0.60 DPS) | yes | Sanguine Cape (14376, -0.08 DPS) [world_drop]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Black Whelp Cloak (7283, -0.29 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Gray Woolen Robe (2585, -0.10 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.33 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.13 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.13 DPS) [vendor]; Owl Bracers (4796, -1.26 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 10.0 spell_power points (1.04 DPS) | yes | Wisdom's Leather Gloves (252499, -0.21 DPS) [crafted]; Windfelt Gloves (5630, -0.23 DPS) [quest]; Pristine Gloves (253913, -0.23 DPS) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 10.0 spell_power points (1.04 DPS) | yes | Stormrider's Leather Belt (252432, -0.10 DPS) [crafted]; Pristine Sash (253925, -0.10 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.21 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.0 spell_power points (1.97 DPS) | yes | Dreamer's Leggings (270016, -0.13 DPS) [quest]; Stormrider's Leather Pants (252502, -0.16 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.47 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.2 spell_power points (1.27 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.21 DPS) [crafted]; Black Whelp Slippers (252424, -0.44 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.5 spell_power points (0.78 DPS) | yes | Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Sludge-Stained Band (286535, -0.47 DPS) [world]; Lavishly Jeweled Ring (1156, -1.01 DPS, sim-verified) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; Sludge-Stained Band (286535, -0.21 DPS) [world]; Lavishly Jeweled Ring (1156, -1.09 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 12.5 spell_power points (1.30 DPS) | yes | Channeler's Staff (4437, -0.26 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.47 DPS) [dungeon]; Lesser Staff of the Spire (1300, -0.52 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Keller's Girdle; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 53.2. Weights run: 3.3s. Verify run: 1.1s. 322 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.878 ± 0.011, crit=0.064 ± 0.004 per rating point (14 rating = 1%, 0.899 per %), hit=0.165 ± 0.002 per rating point (10 rating = 1%, 1.649 per %), spell_haste=not significant (0.004 ± 0.065), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.419 ± 0.001, arcane_power=0.581 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | sim-verified (53.2 DPS) | yes | Totemic Leather Helm (252456, -0.29 DPS) [crafted]; Holy Shroud (2721, -0.43 DPS) [world_drop]; Enchanter's Cowl (4322, -0.65 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.3 spell_power points (1.72 DPS) | yes | Crystal Starfire Medallion (5003, -1.23 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.23 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.34 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.9 spell_power points (2.37 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.42 DPS) [quest]; Magician's Mantle (12998, -0.56 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.0 spell_power points (0.98 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.07 DPS) [quest]; Resilient Cape (14400, -0.25 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 20.4 spell_power points (2.86 DPS) | yes | Guardian Armor (4256, -0.37 DPS) [crafted]; Death Speaker Robes (6682, -0.53 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.72 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.26 DPS) | yes | Nightsky Wristbands (6407, -0.52 DPS) [world_drop]; Technician's Bracers (270042, -0.52 DPS) [quest]; Glowing Magical Bracelets (13106, -1.28 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 13.7 spell_power points (1.91 DPS) | yes | Truefaith Gloves (7049, -0.84 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.93 DPS) [dungeon]; Stormrider's Leather Gloves (252498, -1.19 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.3 spell_power points (2.14 DPS) | yes | Guardian Belt (4258, -0.44 DPS) [crafted]; Crimson Silk Belt (7055, -0.44 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.54 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 19.7 spell_power points (2.75 DPS) | yes | Abomination Skin Leggings (23173, -0.56 DPS, sim-verified) [dungeon]; Stormrider's Leather Pants (252502, -0.62 DPS) [crafted]; Guardian Pants (5962, -0.65 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.1 spell_power points (1.84 DPS) | yes | Spidersilk Boots (4320, -0.37 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -1.36 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.98 DPS) | yes | Black Widow Band (6199, -0.12 DPS) [world]; Snake Hoop (6750, -0.12 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.14 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.8 spell_power points (0.95 DPS) | yes | Black Widow Band (6199, -0.09 DPS) [world]; Snake Hoop (6750, -0.09 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.11 DPS) [vendor] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Glimmering Staff (249392, -1.92 DPS) [crafted]; Scorn's Focal Dagger (23168, -2.02 DPS) [dungeon]; Manual Crowd Pummeler (9449, -4.05 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 75.2. Weights run: 3.3s. Verify run: 1.2s. 438 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.756 ± 0.012, crit=0.097 ± 0.007 per rating point (14 rating = 1%, 1.354 per %), hit=0.257 ± 0.003 per rating point (10 rating = 1%, 2.573 per %), spell_haste=not significant (0.169 ± 0.045), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.472 ± 0.001, arcane_power=0.528 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.80 DPS) | yes | Big Voodoo Mask (8201, -0.19 DPS) [crafted]; Augural Shroud (2620, -0.33 DPS) [world]; Corpseshroud (10574, -0.89 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.5 spell_power points (1.54 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.53 DPS) [quest]; Triune Amulet (7722, -0.83 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.83 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.8 spell_power points (2.24 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.8 spell_power points (2.10 DPS) | yes | Guardian Cloak (5965, -0.80 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.00 DPS) [vendor]; Long Silken Cloak (4326, -1.64 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (75.2 DPS) | yes | Robe of Power (7054, -0.23 DPS) [crafted]; Elemental Raiment (9434, -0.51 DPS) [world_drop]; Robe of the Magi (1716, -1.48 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.5 spell_power points (1.40 DPS) | yes | Arcane Runed Bracers (4744, -0.20 DPS) [quest]; Spidertank Oilrag (9448, -0.20 DPS) [dungeon]; Condor Bracers (15864, -0.47 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (3.20 DPS) | yes | Dreamweave Gloves (10019, -0.40 DPS) [crafted]; Red Mageweave Gloves (10018, -0.73 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.83 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.3 spell_power points (2.44 DPS) | yes | Highlander's Cloth Girdle (20098, -0.17 DPS) [rep]; Skycaller's Leather Belt (252522, -0.51 DPS) [crafted]; Gilded Cord (254037, -0.57 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.1 spell_power points (3.07 DPS) | yes | Crimson Silk Pantaloons (7062, -0.70 DPS) [crafted]; Kodohide Legguards (285338, -0.74 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.20 DPS) | yes | Skycaller's Leather Shoes (252532, -0.90 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.43 DPS) [crafted]; Gilded Slippers (254001, -1.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.94 DPS) | yes | Ring of Forlorn Spirits (2043, -0.87 DPS) [quest]; Reedknot Ring (9622, -1.00 DPS) [quest]; Minor Channeling Ring (1449, -1.07 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.20 DPS) | yes | Reedknot Ring (9622, -0.27 DPS) [quest]; Minor Channeling Ring (1449, -0.33 DPS) [quest]; Ring of Forlorn Spirits (2043, -0.77 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, -0.37 DPS) [world]; Mograine's Might (7723, -1.42 DPS) [dungeon]; Manual Crowd Pummeler (9449, -3.65 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 106.1. Weights run: 3.5s. Verify run: 1.3s. 577 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.031 ± 0.017, crit=0.143 ± 0.012 per rating point (14 rating = 1%, 2.004 per %), hit=0.404 ± 0.005 per rating point (10 rating = 1%, 4.036 per %), spell_haste=0.401 ± 0.097, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.539 ± 0.001, arcane_power=0.461 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 39.6 spell_power points (4.72 DPS) | yes | Soothsayer's Headdress (17740, -0.14 DPS) [dungeon]; Dreamweave Circlet (10041, -0.99 DPS) [crafted]; Chief Architect's Monocle (11839, -1.40 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.4 spell_power points (1.72 DPS) | yes | Scorn's Icy Choker (23169, -0.15 DPS) [dungeon]; Mindburst Medallion (11196, -0.27 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.49 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.6 spell_power points (4.24 DPS) | yes | Kentic Amice (11624, -0.98 DPS) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.10 DPS) [vendor]; Rotgrip Mantle (17732, -1.87 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.2 spell_power points (2.40 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.23 DPS) [dungeon]; Runecloth Cloak (13860, -0.35 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.68 DPS) [vendor] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (106.1 DPS) | yes | Runecloth Robe (13858, -0.09 DPS) [crafted]; Runecloth Tunic (13857, -0.12 DPS) [crafted]; Acumen Robes (17775, -1.90 DPS, sim-verified) [quest] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.3 spell_power points (2.42 DPS) | yes | Skycaller's Leather Bracers (252542, -0.37 DPS) [crafted]; Aristocratic Cuffs (12546, -0.58 DPS) [dungeon]; Bloodband Bracers (11469, -0.72 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 32.4 spell_power points (3.86 DPS) | yes | Sergeant Major's Crackling Leather Gauntlets (220866, -0.84 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -0.85 DPS) [crafted]; Raider Handwraps (272098, -1.29 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 28.4 spell_power points (3.38 DPS) | yes | Dawnspire Cord (12466, -0.33 DPS) [dungeon]; Satyrmane Sash (17755, -0.48 DPS) [dungeon]; Ban'thok Sash (11662, -0.55 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 33.3 spell_power points (3.96 DPS) | yes | Red Mageweave Pants (10009, -0.83 DPS) [crafted]; Big Voodoo Pants (8202, -0.95 DPS) [crafted]; Knight's Crackling Leather Leggings (220864, -1.23 DPS, sim-verified) [vendor] |
| feet | Skycaller's Leather Boots (252471) | Leatherworking [crafted] | 26.3 spell_power points (3.14 DPS) | yes | Sergeant Major's Crackling Leather Boots (220862, -0.12 DPS) [vendor]; Earthen Silk Slippers (254013, -0.28 DPS) [crafted]; Mender's Leather Boots (252472, -0.71 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.2 spell_power points (1.93 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Mindseye Circle (10634, -0.46 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.2 spell_power points (1.93 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Mindseye Circle (10634, -0.45 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smoking Heart of the Mountain (11811, +0.00 DPS) [crafted] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mechanic's Pipehammer (9604, -0.68 DPS) [quest]; Spellshifter Rod (9527, -0.74 DPS) [quest]; Blade of Eternal Darkness (17780, -2.18 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Feathered Breastplate; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 168.3. Weights run: 3.6s. Verify run: 1.3s. 1449 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.148 ± 0.023, crit=0.186 ± 0.017 per rating point (14 rating = 1%, 2.603 per %), hit=0.515 ± 0.007 per rating point (10 rating = 1%, 5.154 per %), spell_haste=not significant (0.551 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.554 ± 0.002, arcane_power=0.446 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+3.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Magister's Crown (16686, -0.03 DPS) [dungeon]; Feralheart Cowl (226773, -3.29 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.9 spell_power points (3.54 DPS) | yes | Beads of Ogre Mojo (22149, -0.37 DPS) [quest]; Archlight Talisman (15856, -0.76 DPS) [quest]; Chains of the Lich (23125, -0.94 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.8 spell_power points (5.55 DPS) | yes | Darkspear Shoulderpads (272103, -0.45 DPS) [vendor]; Darkspear Shoulders (272104, -0.45 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -0.61 DPS) [pvp] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.3 spell_power points (3.59 DPS) | yes | Hide of the Wild (18510, -0.58 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.68 DPS) [world]; Spritecaster Cape (11623, -1.12 DPS) [dungeon] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Armor (231696, -0.05 DPS) [vendor]; Chestplate of Tranquility (18373, -0.08 DPS) [dungeon]; Tunic of Undead Slaying (23089, -8.87 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -0.91 DPS) [dungeon]; Runecloth Cuffs (254123, -1.03 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -7.20 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Hands of Power (13253, -0.12 DPS) [dungeon]; Raider Handwraps (272097, -3.57 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 55.4 spell_power points (6.56 DPS) | yes | Girdle of Insight (18504, -1.66 DPS) [crafted]; Belt of the Archmage (18405, -1.71 DPS) [crafted]; Elunite Cord (272401, -3.59 DPS, sim-verified) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 61.3 spell_power points (7.26 DPS) | yes | Sentinel's Silk Leggings (22752, -1.36 DPS) [rep]; Sentinel's Lizardhide Pants (237817, -1.49 DPS) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 40.4 spell_power points (4.78 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.39 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, -0.76 DPS) [vendor]; Dragonrider Boots (18102, -2.10 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.02 DPS) [quest]; Maiden's Circle (13001, -1.02 DPS) [world_drop]; Naglering (11669, -5.81 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.36 DPS) [quest]; Maiden's Circle (13001, -0.36 DPS) [world_drop]; Naglering (11669, -6.47 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.11 DPS) [world]; Hand of Edward the Odd (2243, -5.24 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), and 12 more) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS) [vendor] |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 27.5. Weights run: 2.7s. Verify run: 0.9s. 183 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.249 ± 0.009, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.581 per %), hit=0.128 ± 0.001 per rating point (10 rating = 1%, 1.284 per %), spell_haste=-1.246 ± 0.100, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.006 ± 0.000, arcane_power=0.994 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Shadow Goggles (4373, -0.29 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.31 DPS) [crafted]; Trapper's Leather Hood (252505, -1.35 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.2 spell_power points (1.69 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.65 DPS) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.27 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.7 spell_power points (0.60 DPS) | yes | Sanguine Cape (14376, -0.08 DPS) [world_drop]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Black Whelp Cloak (7283, -0.29 DPS) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Gray Woolen Robe (2585, -0.10 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.33 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 7.5 spell_power points (0.78 DPS) | yes | Mindthrust Bracers (1974, -0.13 DPS) [dungeon]; Owl Bracers (4796, -0.13 DPS) [vendor]; Featherbead Bracers (15452, -0.13 DPS) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Gloves (252499, -0.08 DPS) [crafted]; Pristine Gloves (253913, -0.10 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.49 DPS, sim-verified) [crafted] |
| waist | Stormrider's Leather Belt (252432) | Leatherworking [crafted] | sim-verified (+0.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Pristine Sash (253925, +0.00 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.80 DPS, sim-verified) [world_drop] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Abomination Skin Leggings (23173, -0.27 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.42 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.2 spell_power points (1.27 DPS) | yes | Spidersilk Boots (4320, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.21 DPS) [crafted]; Black Whelp Slippers (252424, -0.44 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.5 spell_power points (0.78 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Sludge-Stained Band (286535, -0.47 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (+1.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; Sludge-Stained Band (286535, -0.21 DPS) [world]; Loop of Sacrifice (281673, -1.20 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 12.5 spell_power points (1.30 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Channeler's Staff (4437, -0.26 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.47 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Blight Gloves; waist: Stormrider's Leather Belt; legs: Stormrider's Leather Pants; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; main_hand: Gnarled Necromancer's Staff

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 49.7. Weights run: 3.3s. Verify run: 1.1s. 315 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.878 ± 0.011, crit=0.064 ± 0.004 per rating point (14 rating = 1%, 0.899 per %), hit=0.165 ± 0.002 per rating point (10 rating = 1%, 1.649 per %), spell_haste=not significant (0.004 ± 0.065), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.419 ± 0.001, arcane_power=0.581 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 14.8 spell_power points (2.07 DPS) | yes | Enduring Cap (3020, +0.00 DPS, sim-verified) [world_drop]; Totemic Leather Helm (252456, -0.39 DPS) [crafted]; Holy Shroud (2721, -0.53 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.3 spell_power points (1.72 DPS) | yes | Crystal Starfire Medallion (5003, -1.23 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.23 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.42 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.9 spell_power points (2.37 DPS) | yes | Death Speaker Mantle (6685, -0.17 DPS) [dungeon]; Fairywing Mantle (9536, -0.42 DPS) [quest]; Magician's Mantle (12998, -0.56 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.0 spell_power points (0.98 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS) [vendor]; Resilient Cape (14400, -0.25 DPS) [world_drop]; Hillman's Cloak (3719, -0.28 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 20.4 spell_power points (2.86 DPS) | yes | Guardian Armor (4256, -0.46 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.53 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.72 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.26 DPS) | yes | Nightsky Wristbands (6407, -0.52 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.52 DPS) [quest]; Glowing Magical Bracelets (13106, -1.40 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.8 spell_power points (1.51 DPS) | yes | Jutebraid Gloves (10654, -0.05 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.31 DPS) [crafted]; Truefaith Gloves (7049, -0.44 DPS) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.3 spell_power points (2.14 DPS) | yes | Guardian Belt (4258, -0.44 DPS) [crafted]; Crimson Silk Belt (7055, -0.44 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.63 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 19.7 spell_power points (2.75 DPS) | yes | Stormrider's Leather Pants (252502, -0.62 DPS) [crafted]; Abomination Skin Leggings (23173, -0.65 DPS, sim-verified) [dungeon]; Guardian Pants (5962, -0.65 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.1 spell_power points (1.84 DPS) | yes | Spidersilk Boots (4320, -0.37 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -1.47 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.98 DPS) | yes | Sea Giant's Toe Ring (274746, -0.14 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Black Widow Band (6199, -1.80 DPS, sim-verified) [world] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | sim-verified (49.7 DPS) | yes | Sea Giant's Toe Ring (274746, -0.02 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.12 DPS) [dungeon]; Black Widow Band (6199, -0.70 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -0.09 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.12 DPS) [quest]; Manual Crowd Pummeler (9449, -2.06 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Snake Hoop; main_hand: Glimmering Staff

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 75.2. Weights run: 3.3s. Verify run: 1.2s. 426 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.756 ± 0.012, crit=0.097 ± 0.007 per rating point (14 rating = 1%, 1.354 per %), hit=0.257 ± 0.003 per rating point (10 rating = 1%, 2.573 per %), spell_haste=not significant (0.169 ± 0.045), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.472 ± 0.001, arcane_power=0.528 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.80 DPS) | yes | Big Voodoo Mask (8201, -0.19 DPS) [crafted]; Augural Shroud (2620, -0.33 DPS) [world]; Corpseshroud (10574, -0.89 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.5 spell_power points (1.54 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.53 DPS) [quest]; Triune Amulet (7722, -0.83 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.83 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.8 spell_power points (2.24 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS) [crafted]; Bloodmage Mantle (7684, -0.14 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 15.8 spell_power points (2.10 DPS) | yes | Guardian Cloak (5965, -0.80 DPS) [crafted]; Darkspear Raider's Cloak (272077, -1.00 DPS) [vendor]; Long Silken Cloak (4326, -1.36 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (75.2 DPS) | yes | Robe of Power (7054, -0.23 DPS) [crafted]; Elemental Raiment (9434, -0.51 DPS) [world_drop]; Robe of the Magi (1716, -1.76 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.5 spell_power points (1.40 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.07 DPS) [quest]; Spidertank Oilrag (9448, -0.20 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (3.20 DPS) | yes | Dreamweave Gloves (10019, -0.40 DPS) [crafted]; Red Mageweave Gloves (10018, -0.73 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.83 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.3 spell_power points (2.44 DPS) | yes | Skycaller's Leather Belt (252522, -0.51 DPS) [crafted]; Gilded Cord (254037, -0.57 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.83 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.1 spell_power points (3.07 DPS) | yes | Kodohide Legguards (285338, -0.63 DPS) [world]; Crimson Silk Pantaloons (7062, -0.70 DPS) [crafted]; Abomination Skin Leggings (23173, -1.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.20 DPS) | yes | Skycaller's Leather Shoes (252532, -0.68 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.43 DPS) [crafted]; Gilded Slippers (254001, -1.56 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.94 DPS) | yes | Reedknot Ring (9622, -1.00 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.14 DPS) [vendor]; Black Widow Band (6199, -1.23 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.20 DPS) | yes | Sea Giant's Toe Ring (274746, -0.40 DPS) [vendor]; Black Widow Band (6199, -0.49 DPS) [world]; Reedknot Ring (9622, -0.78 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -1.05 DPS) [dungeon]; Windweaver Staff (7757, -1.15 DPS) [dungeon]; Gut Ripper (2164, -4.31 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 104.3. Weights run: 3.5s. Verify run: 1.4s. 561 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.031 ± 0.017, crit=0.143 ± 0.012 per rating point (14 rating = 1%, 2.004 per %), hit=0.404 ± 0.005 per rating point (10 rating = 1%, 4.036 per %), spell_haste=0.401 ± 0.097, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.539 ± 0.001, arcane_power=0.461 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 39.6 spell_power points (4.72 DPS) | yes | Soothsayer's Headdress (17740, -0.14 DPS) [dungeon]; Dreamweave Circlet (10041, -0.99 DPS) [crafted]; Chief Architect's Monocle (11839, -1.40 DPS) [dungeon] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 14.4 spell_power points (1.72 DPS) | yes | Scorn's Icy Choker (23169, -0.15 DPS) [dungeon]; Mindburst Medallion (11196, -0.27 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.49 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 35.6 spell_power points (4.24 DPS) | yes | Kentic Amice (11624, -0.98 DPS) [dungeon]; Blood Guard's Crackling Leather Spaulders (220871, -1.10 DPS) [vendor]; Rotgrip Mantle (17732, -1.65 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 21.3 spell_power points (2.53 DPS) | yes | Spritecaster Cape (11623, -0.13 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.36 DPS) [dungeon]; Runecloth Cloak (13860, -0.48 DPS) [crafted] |
| chest | Feathered Breastplate (8349) | Leatherworking [crafted] | sim-verified (104.3 DPS) | yes | Runecloth Robe (13858, -0.09 DPS) [crafted]; Runecloth Tunic (13857, -0.12 DPS) [crafted]; Acumen Robes (17775, -2.20 DPS, sim-verified) [quest] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.3 spell_power points (2.42 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.37 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 32.4 spell_power points (3.86 DPS) | yes | First Sergeant's Crackling Leather Gauntlets (220867, -0.84 DPS) [vendor]; Skycaller's Leather Gauntlets (252550, -0.85 DPS) [crafted]; Raider Handwraps (272098, -1.08 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 28.4 spell_power points (3.38 DPS) | yes | Dawnspire Cord (12466, -0.33 DPS) [dungeon]; Satyrmane Sash (17755, -0.48 DPS) [dungeon]; Ban'thok Sash (11662, -0.55 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 33.3 spell_power points (3.96 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -0.59 DPS) [vendor]; Red Mageweave Pants (10009, -0.83 DPS) [crafted]; Big Voodoo Pants (8202, -0.95 DPS) [crafted] |
| feet | Skycaller's Leather Boots (252471) | Leatherworking [crafted] | 26.3 spell_power points (3.14 DPS) | yes | First Sergeant's Crackling Leather Boots (220863, -0.12 DPS) [vendor]; Earthen Silk Slippers (254013, -0.28 DPS) [crafted]; Mender's Leather Boots (252472, -0.71 DPS) [crafted] |
| finger1 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 16.2 spell_power points (1.93 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Mindseye Circle (10634, -0.46 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.2 spell_power points (1.93 DPS) | yes | Brainlash (6440, -0.09 DPS) [dungeon]; Band of the Unicorn (7553, -0.38 DPS) [world_drop]; Mindseye Circle (10634, -0.45 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellshifter Rod (9527, -0.74 DPS) [quest]; Thorium Greatmace (250613, -0.82 DPS) [crafted]; Blade of Eternal Darkness (17780, -1.59 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Feathered Breastplate; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger1: Cyclopean Band; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 168.5. Weights run: 3.6s. Verify run: 1.3s. 1446 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=1.148 ± 0.023, crit=0.186 ± 0.017 per rating point (14 rating = 1%, 2.603 per %), hit=0.515 ± 0.007 per rating point (10 rating = 1%, 5.154 per %), spell_haste=not significant (0.551 ± 0.173), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.554 ± 0.002, arcane_power=0.446 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Magister's Crown (16686, -0.03 DPS) [dungeon]; Feralheart Cowl (226773, -3.63 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 29.9 spell_power points (3.54 DPS) | yes | Beads of Ogre Mojo (22149, -0.37 DPS) [quest]; Archlight Talisman (15856, -0.76 DPS) [quest]; Chains of the Lich (23125, -0.94 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 46.8 spell_power points (5.55 DPS) | yes | Darkspear Shoulderpads (272103, -0.45 DPS) [vendor]; Darkspear Shoulders (272104, -0.45 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, -0.61 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 30.3 spell_power points (3.59 DPS) | yes | Hide of the Wild (18510, -0.58 DPS) [crafted]; Crystalline Threaded Cape (20697, -0.68 DPS) [world]; Deep Woodlands Cloak (19121, -0.95 DPS) [quest] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Armor (231679, -0.05 DPS) [vendor]; Chestplate of Tranquility (18373, -0.08 DPS) [dungeon]; Tunic of Undead Slaying (23089, -9.04 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -0.91 DPS) [dungeon]; Runecloth Cuffs (254123, -1.03 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -6.94 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (+4.0 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Hands of Power (13253, -0.12 DPS) [dungeon]; Raider Handwraps (272097, -3.98 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 55.4 spell_power points (6.56 DPS) | yes | Girdle of Insight (18504, -1.66 DPS) [crafted]; Belt of the Archmage (18405, -1.71 DPS) [crafted]; Elunite Cord (272401, -3.89 DPS, sim-verified) [vendor] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 61.3 spell_power points (7.26 DPS) | yes | Outrider's Silk Leggings (22747, -1.36 DPS) [rep]; Sentinel's Silk Leggings (237815, -1.49 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1.49 DPS) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 40.4 spell_power points (4.78 DPS) | yes | General's Dragonhide Boots (231682, -0.39 DPS) [pvp]; General's Dragonhide Greaves (231671, -0.76 DPS) [vendor]; Dragonrider Boots (18102, -1.85 DPS, sim-verified) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.02 DPS) [quest]; Maiden's Circle (13001, -1.02 DPS) [world_drop]; Naglering (11669, -6.28 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.36 DPS) [quest]; Maiden's Circle (13001, -0.36 DPS) [world_drop]; Naglering (11669, -6.32 DPS, sim-verified) [dungeon] |
| trinket1 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+7.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Draconic Infused Emblem (22268, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hammer of Divine Might (22333, -0.01 DPS) [dungeon]; Hand of Edward the Odd (2243, -5.53 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), and 12 more) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS) [vendor] |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Second Wind; trinket2: Burst of Knowledge; main_hand: Amethyst War Staff; ranged: Idol of the Moon

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

