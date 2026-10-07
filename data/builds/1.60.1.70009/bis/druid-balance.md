# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5222000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 28.5. Weights run: 1.8s. Verify run: 0.8s. 193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.804 ± 0.005, crit=0.035 ± 0.002 per rating point (14 rating = 1%, 0.487 per %), hit=0.088 ± 0.001 per rating point (10 rating = 1%, 0.879 per %), spell_haste=-0.630 ± 0.029, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.036 ± 0.001, arcane_power=0.964 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.13 DPS) | yes | Wisdom's Leather Hood (252507, -0.38 DPS) [crafted]; Pristine Circlet (253949, -0.38 DPS) [crafted]; Trapper's Leather Hood (252505, -1.69 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.2 spell_power points (1.54 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.43 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.04 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.10 DPS) [world_drop]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.38 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (+0.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Bright Bracers (3647, -0.10 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor]; Owl Bracers (4796, -0.65 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.2 spell_power points (1.03 DPS) | yes | Serpent Gloves (5970, -0.15 DPS) [dungeon]; Windfelt Gloves (5630, -0.23 DPS) [quest]; Pristine Gloves (253913, -0.23 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Keller's Girdle (2911, -0.10 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.69 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.4 spell_power points (1.94 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Dreamer's Leggings (270016, -0.33 DPS) [quest]; Wisdom's Leather Pants (252503, -0.45 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.2 spell_power points (1.29 DPS) | yes | Stormrider's Leather Boots (252443, -0.02 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.28 DPS) [crafted]; Black Whelp Slippers (252424, -0.50 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.6 spell_power points (0.83 DPS) | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Sludge-Stained Band (286535, -0.45 DPS) [world]; Volcanic Rock Ring (12053, -0.53 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.63 DPS) | yes | Sludge-Stained Band (286535, -0.25 DPS) [world]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.60 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 8.0 spell_power points (1.01 DPS) | yes | Rhahk'Zor's Hammer (5187, -0.00 DPS) [dungeon]; Channeler's Staff (4437, -0.20 DPS) [world]; Lesser Staff of the Spire (1300, -0.40 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 5222211015000000-00000000000000000000-0000000000000000)

Set DPS (verified): 52.6. Weights run: 2.2s. Verify run: 0.8s. 322 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.576 ± 0.005, crit=0.053 ± 0.003 per rating point (14 rating = 1%, 0.744 per %), hit=0.102 ± 0.001 per rating point (10 rating = 1%, 1.022 per %), spell_haste=-0.530 ± 0.091, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.397 ± 0.001, arcane_power=0.603 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.05 DPS) | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Holy Shroud (2721, -0.17 DPS) [world_drop]; Enduring Cap (3020, -0.47 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.5 spell_power points (1.78 DPS) | yes | Crystal Starfire Medallion (5003, -1.39 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.39 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.70 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.2 spell_power points (2.42 DPS) | yes | Fairywing Mantle (9536, -0.51 DPS) [quest]; Magician's Mantle (12998, -0.68 DPS) [world_drop]; Death Speaker Mantle (6685, -0.75 DPS, sim-verified) [dungeon] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.3 spell_power points (0.90 DPS) | yes | Hillman's Cloak (3719, -0.05 DPS) [crafted]; Cloak of Rot (4462, -0.12 DPS) [world]; Darkspear Raider's Cloak (272078, -0.12 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.5 spell_power points (2.81 DPS) | yes | Guardian Armor (4256, -0.29 DPS) [crafted]; Stormrider's Leather Tunic (252510, -0.52 DPS) [crafted]; Death Speaker Robes (6682, -0.54 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.53 DPS) | yes | Nightsky Wristbands (6407, -0.95 DPS) [world_drop]; Technician's Bracers (270042, -0.95 DPS) [quest]; Glowing Magical Bracelets (13106, -1.81 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.3 spell_power points (1.76 DPS) | yes | Stormrider's Leather Gloves (252498, -0.52 DPS) [crafted]; Shilly Mitts (9609, -0.57 DPS) [quest]; Gloves of Insight (9698, -0.57 DPS) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.5 spell_power points (2.30 DPS) | yes | Highlander's Cloth Girdle (20099, -0.12 DPS) [rep]; Moss Cinch (6911, -0.25 DPS) [dungeon]; Belt of Arugal (6392, -0.47 DPS) [dungeon] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 16.3 spell_power points (2.79 DPS) | yes | Dark Ritual Leggings (270031, -0.40 DPS) [quest]; Abomination Skin Leggings (23173, -0.47 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.49 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (1.88 DPS) | yes | Spidersilk Boots (4320, -0.29 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.37 DPS) [crafted]; Acidic Walkers (9454, -2.10 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.19 DPS) | yes | Black Widow Band (6199, -0.51 DPS) [world]; Snake Hoop (6750, -0.51 DPS) [quest]; Minor Channeling Ring (1449, -2.15 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (52.6 DPS) | yes | Black Widow Band (6199, -0.34 DPS) [world]; Snake Hoop (6750, -0.34 DPS) [quest]; Minor Channeling Ring (1449, -1.22 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -2.20 DPS) [dungeon]; Rhahk'Zor's Hammer (5187, -2.37 DPS) [dungeon]; Manual Crowd Pummeler (9449, -4.48 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 5222211015401050-00000000000000000000-0000000000000000)

Set DPS (verified): 78.9. Weights run: 2.3s. Verify run: 0.8s. 438 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.395 ± 0.006, crit=0.070 ± 0.004 per rating point (14 rating = 1%, 0.975 per %), hit=0.140 ± 0.002 per rating point (10 rating = 1%, 1.401 per %), spell_haste=0.205 ± 0.034, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.381 ± 0.001, arcane_power=0.619 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.97 DPS) | yes | Augural Shroud (2620, -0.80 DPS, sim-verified) [world]; Big Voodoo Mask (8201, -1.22 DPS) [crafted]; Living Cowl (5608, -1.51 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (1.77 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.88 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.25 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.25 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (2.37 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.6 spell_power points (2.37 DPS) | yes | Guardian Cloak (5965, -0.87 DPS) [crafted]; Icy Cloak (4327, -1.05 DPS) [crafted]; Long Silken Cloak (4326, -2.21 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+0.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.11 DPS) [world_drop]; Robe of Power (7054, -0.53 DPS) [crafted]; Robe of the Magi (1716, -0.88 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Leather Bracers (4260, -0.12 DPS) [crafted]; Condor Bracers (15864, -0.38 DPS) [quest]; Arcane Runed Bracers (4744, -1.03 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.53 DPS) | yes | Dreamweave Gloves (10019, -1.63 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.70 DPS) [crafted]; Red Mageweave Gloves (10018, -1.71 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.6 spell_power points (2.94 DPS) | yes | Star Belt (4329, -0.49 DPS) [crafted]; Deathmage Sash (10771, -0.50 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.61 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.7 spell_power points (3.54 DPS) | yes | Dark Ritual Leggings (270031, -0.90 DPS) [quest]; Kodohide Legguards (285338, -0.93 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -1.06 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.53 DPS) | yes | Skycaller's Leather Shoes (252532, -1.85 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.50 DPS) [crafted]; Gilded Slippers (254001, -2.69 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (2.34 DPS) | yes | Ring of Forlorn Spirits (2043, -0.83 DPS) [quest]; Reedknot Ring (9622, -1.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.20 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.70 DPS) | yes | Reedknot Ring (9622, -0.38 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.57 DPS) [vendor]; Ring of Forlorn Spirits (2043, -1.33 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, -0.18 DPS) [world]; Scorn's Focal Dagger (23168, -2.26 DPS) [dungeon]; Manual Crowd Pummeler (9449, -4.80 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 5222211015401051-00000000000000000000-5400000000000000)

Set DPS (verified): 102.5. Weights run: 2.6s. Verify run: 1.0s. 577 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.558 ± 0.009, crit=0.106 ± 0.007 per rating point (14 rating = 1%, 1.478 per %), hit=0.225 ± 0.003 per rating point (10 rating = 1%, 2.247 per %), spell_haste=0.393 ± 0.069, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.465 ± 0.001, arcane_power=0.535 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.4 spell_power points (5.24 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.73 DPS) [crafted]; Dreamweave Circlet (10041, -0.80 DPS) [crafted]; Red Mageweave Headband (10033, -1.16 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.73 DPS) | yes | Mindburst Medallion (11196, -0.17 DPS) [quest]; Horizon Choker (13085, -0.42 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.80 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 26.2 spell_power points (4.37 DPS) | yes | Kentic Amice (11624, -0.82 DPS) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.06 DPS) [vendor]; Rotgrip Mantle (17732, -2.03 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.3 spell_power points (2.90 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.56 DPS) [dungeon]; Runecloth Cloak (13860, -0.65 DPS) [crafted]; Big Voodoo Cloak (8216, -1.22 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 30.2 spell_power points (5.04 DPS) | yes | Robe of the Magi (1716, -0.80 DPS) [world_drop]; Feathered Breastplate (8349, -0.93 DPS) [crafted]; Runecloth Tunic (13857, -1.17 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 15.6 spell_power points (2.60 DPS) | yes | Skycaller's Leather Bracers (252542, -0.28 DPS) [crafted]; Nethergeld Cuffs (254061, -0.78 DPS) [crafted]; Bloodband Bracers (11469, -0.93 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 26.3 spell_power points (4.38 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.87 DPS) [crafted]; Bloodfire Talons (12464, -0.91 DPS) [dungeon]; Gloves of the Greatfather (17721, -1.26 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 22.7 spell_power points (3.79 DPS) | yes | Satyrmane Sash (17755, -0.52 DPS) [dungeon]; Ban'thok Sash (11662, -0.72 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.00 DPS) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 28.6 spell_power points (4.77 DPS) | yes | Red Mageweave Pants (10009, -1.32 DPS) [crafted]; Big Voodoo Pants (8202, -1.34 DPS) [crafted]; Knight's Crackling Leather Leggings (220864, -2.56 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.01 DPS) | yes | Sergeant Major's Crackling Leather Boots (220862, -0.86 DPS) [vendor]; Skycaller's Leather Boots (252471, -0.95 DPS, sim-verified) [crafted]; Skycaller's Leather Shoes (252532, -1.18 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (2.23 DPS) | yes | Lorekeeper's Ring (19523, -0.23 DPS) [rep]; Band of the Unicorn (7553, -3.35 DPS, sim-verified) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (102.5 DPS) | yes | Lorekeeper's Ring (19523, -0.15 DPS) [rep]; Band of the Unicorn (7553, -1.78 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -1.32 DPS, sim-verified) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mechanic's Pipehammer (9604, -0.20 DPS) [quest]; Spellforce Rod (1664, -0.50 DPS) [world]; Blade of Eternal Darkness (17780, -1.37 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Thorium Greatmace

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 5222211015401051-00000000000000000000-5533300000000000)

Set DPS (verified): 180.0. Weights run: 2.6s. Verify run: 1.0s. 1449 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.571 ± 0.011, crit=0.136 ± 0.008 per rating point (14 rating = 1%, 1.911 per %), hit=0.282 ± 0.003 per rating point (10 rating = 1%, 2.821 per %), spell_haste=not significant (0.356 ± 0.101), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.458 ± 0.001, arcane_power=0.542 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+5.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -0.12 DPS) [dungeon]; Feralheart Cowl (226773, -5.34 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 22.4 spell_power points (3.92 DPS) | yes | Orb of the Darkmoon (19426, -0.07 DPS) [quest]; Chains of the Lich (23125, -0.07 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.45 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.5 spell_power points (6.55 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, -0.80 DPS) [pvp]; Burial Shawl (18681, -1.46 DPS) [dungeon]; Feralheart Spaulders (226778, -1.54 DPS) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.4 spell_power points (4.09 DPS) | yes | Crystalline Threaded Cape (20697, -0.19 DPS) [world]; Hide of the Wild (18510, -0.64 DPS) [crafted]; Amplifying Cloak (18350, -0.94 DPS) [dungeon] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Armor (231696, -0.19 DPS) [vendor]; Robe of Everlasting Night (18385, -0.32 DPS) [dungeon]; Tunic of Undead Slaying (23089, -11.08 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -1.55 DPS) [dungeon]; Runecloth Cuffs (254123, -1.72 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -8.98 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (+3.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Gloves of the Greatfather (17721, -0.42 DPS) [crafted]; Hands of Power (13253, -3.60 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 40.4 spell_power points (7.06 DPS) | yes | Elunite Cord (272401, -1.87 DPS) [vendor]; Girdle of Insight (18504, -2.14 DPS) [crafted]; Belt of the Archmage (18405, -3.74 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.2 spell_power points (8.78 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.51 DPS) [vendor]; Sentinel's Silk Leggings (22752, -1.98 DPS) [rep] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 31.1 spell_power points (5.45 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.37 DPS) [pvp]; Waterspout Boots (18322, -0.47 DPS) [dungeon]; Dragonrider Boots (18102, -0.70 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.10 DPS) [quest]; Maiden's Circle (13001, -1.10 DPS) [world_drop]; Naglering (11669, -8.20 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.52 DPS) [quest]; Maiden's Circle (13001, -0.52 DPS) [world_drop]; Naglering (11669, -8.49 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+12.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.02 DPS) [world]; Hand of Edward the Odd (2243, -7.19 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Burst of Knowledge; main_hand: Crackling Staff

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 5222000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 26.9. Weights run: 1.8s. Verify run: 0.7s. 183 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.804 ± 0.005, crit=0.035 ± 0.002 per rating point (14 rating = 1%, 0.487 per %), hit=0.088 ± 0.001 per rating point (10 rating = 1%, 0.879 per %), spell_haste=-0.630 ± 0.029, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.036 ± 0.001, arcane_power=0.964 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.13 DPS) | yes | Wisdom's Leather Hood (252507, -0.38 DPS) [crafted]; Pristine Circlet (253949, -0.38 DPS) [crafted]; Trapper's Leather Hood (252505, -1.71 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.2 spell_power points (1.54 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.33 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.04 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Sanguine Cape (14376, -0.10 DPS) [world_drop]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.38 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.08 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.38 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.8 spell_power points (0.61 DPS) | yes | Mindthrust Bracers (1974, -0.10 DPS) [dungeon]; Owl Bracers (4796, -0.10 DPS) [vendor]; Featherbead Bracers (15452, -0.10 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.2 spell_power points (1.03 DPS) | yes | Serpent Gloves (5970, -0.15 DPS) [dungeon]; Pristine Gloves (253913, -0.23 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.25 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Keller's Girdle (2911, -0.10 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.69 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Wisdom's Leather Pants (252503, -0.38 DPS) [crafted]; Abomination Skin Leggings (23173, -0.46 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.50 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.2 spell_power points (1.29 DPS) | yes | Stormrider's Leather Boots (252443, -0.02 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.28 DPS) [crafted]; Black Whelp Slippers (252424, -0.50 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.63 DPS) | yes | Loop of Sacrifice (281673, -0.12 DPS) [quest]; Sludge-Stained Band (286535, -0.25 DPS) [world]; Volcanic Rock Ring (12053, -0.33 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.8 spell_power points (0.61 DPS) | yes | Loop of Sacrifice (281673, -0.10 DPS) [quest]; Sludge-Stained Band (286535, -0.23 DPS) [world]; Volcanic Rock Ring (12053, -0.30 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Necromancer's Staff (251534) (or Twisted Chanter's Staff (890)) | The Wrath of Rath'mael [quest] | 8.0 spell_power points (1.01 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -0.00 DPS) [dungeon]; Channeler's Staff (4437, -0.20 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Gnarled Necromancer's Staff

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 5222211015000000-00000000000000000000-0000000000000000)

Set DPS (verified): 49.2. Weights run: 2.2s. Verify run: 0.9s. 315 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.576 ± 0.005, crit=0.053 ± 0.003 per rating point (14 rating = 1%, 0.744 per %), hit=0.102 ± 0.001 per rating point (10 rating = 1%, 1.022 per %), spell_haste=-0.530 ± 0.091, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.397 ± 0.001, arcane_power=0.603 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (49.2 DPS) | yes | Holy Shroud (2721, -0.13 DPS) [world_drop]; Enduring Cap (3020, -0.43 DPS) [world_drop]; Totemic Leather Helm (252456, -0.61 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.5 spell_power points (1.78 DPS) | yes | Crystal Starfire Medallion (5003, -1.39 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.39 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.72 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.2 spell_power points (2.42 DPS) | yes | Death Speaker Mantle (6685, -0.32 DPS) [dungeon]; Fairywing Mantle (9536, -0.51 DPS) [quest]; Magician's Mantle (12998, -0.68 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.85 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Cloak of Rot (4462, -0.07 DPS) [world]; Darkspear Raider's Cloak (272078, -0.07 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.5 spell_power points (2.81 DPS) | yes | Stormrider's Leather Tunic (252510, -0.52 DPS) [crafted]; Death Speaker Robes (6682, -0.54 DPS) [dungeon]; Guardian Armor (4256, -0.83 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.53 DPS) | yes | Nightsky Wristbands (6407, -0.95 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.95 DPS) [quest]; Glowing Magical Bracelets (13106, -1.13 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.2 spell_power points (1.73 DPS) | yes | Jutebraid Gloves (10654, -0.22 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.49 DPS) [crafted]; Serpent Gloves (5970, -0.54 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.5 spell_power points (2.30 DPS) | yes | Moss Cinch (6911, -0.25 DPS) [dungeon]; Warsong Sash (16975, -0.42 DPS) [quest]; Defiler's Cloth Girdle (20164, -1.03 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 16.3 spell_power points (2.79 DPS) | yes | Dark Ritual Leggings (270031, -0.40 DPS) [quest]; Abomination Skin Leggings (23173, -0.47 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.49 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 11.0 spell_power points (1.88 DPS) | yes | Spidersilk Boots (4320, -0.29 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.37 DPS) [crafted]; Acidic Walkers (9454, -1.19 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.19 DPS) | yes | Black Widow Band (6199, -0.51 DPS) [world]; Snake Hoop (6750, -0.51 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.02 DPS) | yes | Snake Hoop (6750, -0.34 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.43 DPS) [dungeon]; Black Widow Band (6199, -0.99 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rhahk'Zor's Hammer (5187, -0.17 DPS) [dungeon]; Glimmering Staff (249392, -0.45 DPS) [crafted]; Manual Crowd Pummeler (9449, -2.63 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.5 spell_power points (1.78 DPS) | yes | Witch's Finger (16887, -1.10 DPS) [quest]; Alliance Outrunner Healing Rod (285348, -1.10 DPS) [world]; Tome of the Darkspear Prophecy (272090, -1.36 DPS, sim-verified) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 5222211015401050-00000000000000000000-0000000000000000)

Set DPS (verified): 79.5. Weights run: 2.3s. Verify run: 0.8s. 426 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.395 ± 0.006, crit=0.070 ± 0.004 per rating point (14 rating = 1%, 0.975 per %), hit=0.140 ± 0.002 per rating point (10 rating = 1%, 1.401 per %), spell_haste=0.205 ± 0.034, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.381 ± 0.001, arcane_power=0.619 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (3.97 DPS) | yes | Augural Shroud (2620, -1.14 DPS) [world]; Big Voodoo Mask (8201, -1.22 DPS) [crafted]; Living Cowl (5608, -1.51 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (1.77 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.72 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.25 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.25 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 spell_power points (2.37 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.08 DPS) [dungeon]; Berylline Pads (4197, -0.30 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.6 spell_power points (2.37 DPS) | yes | Guardian Cloak (5965, -0.87 DPS) [crafted]; Icy Cloak (4327, -1.05 DPS) [crafted]; Long Silken Cloak (4326, -1.56 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (79.5 DPS) | yes | Elemental Raiment (9434, -0.11 DPS) [world_drop]; Robe of Power (7054, -0.53 DPS) [crafted]; Robe of the Magi (1716, -1.39 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.70 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Guardian Leather Bracers (4260, -0.12 DPS) [crafted]; Radiant Silver Bracers (4545, -0.35 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.53 DPS) | yes | Dreamweave Gloves (10019, -1.43 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.70 DPS) [crafted]; Red Mageweave Gloves (10018, -1.71 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.6 spell_power points (2.94 DPS) | yes | Star Belt (4329, -0.49 DPS) [crafted]; Deathmage Sash (10771, -0.50 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.61 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.7 spell_power points (3.54 DPS) | yes | Dark Ritual Leggings (270031, -0.90 DPS) [quest]; Kodohide Legguards (285338, -0.94 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -1.06 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.53 DPS) | yes | Skycaller's Leather Shoes (252532, -1.31 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.50 DPS) [crafted]; Gilded Slippers (254001, -2.69 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (2.34 DPS) | yes | Reedknot Ring (9622, -1.01 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.20 DPS) [vendor]; Sludge-Stained Band (286535, -1.77 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.70 DPS) | yes | Sea Giant's Toe Ring (274746, -0.57 DPS) [vendor]; Reedknot Ring (9622, -1.09 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -1.13 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -2.08 DPS) [dungeon]; Rhahk'Zor's Hammer (5187, -2.27 DPS) [dungeon]; Manual Crowd Pummeler (9449, -5.74 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 5222211015401051-00000000000000000000-5400000000000000)

Set DPS (verified): 101.0. Weights run: 2.6s. Verify run: 1.0s. 561 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.558 ± 0.009, crit=0.106 ± 0.007 per rating point (14 rating = 1%, 1.478 per %), hit=0.225 ± 0.003 per rating point (10 rating = 1%, 2.247 per %), spell_haste=0.393 ± 0.069, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.465 ± 0.001, arcane_power=0.535 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 31.4 spell_power points (5.24 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.73 DPS) [crafted]; Dreamweave Circlet (10041, -0.80 DPS) [crafted]; Red Mageweave Headband (10033, -1.24 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.73 DPS) | yes | Mindburst Medallion (11196, -0.17 DPS) [quest]; Horizon Choker (13085, -0.42 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.80 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 26.2 spell_power points (4.37 DPS) | yes | Kentic Amice (11624, -0.82 DPS) [dungeon]; Blood Guard's Crackling Leather Spaulders (220871, -1.06 DPS) [vendor]; Rotgrip Mantle (17732, -2.12 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.3 spell_power points (2.90 DPS) | yes | Deep Woodlands Cloak (19121, -0.05 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.56 DPS) [dungeon]; Runecloth Cloak (13860, -0.65 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 30.2 spell_power points (5.04 DPS) | yes | Feathered Breastplate (8349, -0.93 DPS) [crafted]; Runecloth Tunic (13857, -1.17 DPS) [crafted]; Robe of the Magi (1716, -1.25 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 15.6 spell_power points (2.60 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.28 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 26.3 spell_power points (4.38 DPS) | yes | Skycaller's Leather Gauntlets (252550, -0.87 DPS) [crafted]; Bloodfire Talons (12464, -0.91 DPS) [dungeon]; Gloves of the Greatfather (17721, -1.56 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 22.7 spell_power points (3.79 DPS) | yes | Ban'thok Sash (11662, -0.72 DPS) [dungeon]; Mender's Leather Waistguard (252477, -1.00 DPS) [crafted]; Satyrmane Sash (17755, -1.07 DPS, sim-verified) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 28.6 spell_power points (4.77 DPS) | yes | Red Mageweave Pants (10009, -1.32 DPS) [crafted]; Big Voodoo Pants (8202, -1.34 DPS) [crafted]; Stone Guard's Crackling Leather Leggings (220865, -2.66 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.01 DPS) | yes | First Sergeant's Crackling Leather Boots (220863, -0.86 DPS) [vendor]; Skycaller's Leather Boots (252471, -0.96 DPS, sim-verified) [crafted]; Skycaller's Leather Shoes (252532, -1.18 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.3 spell_power points (2.23 DPS) | yes | Advisor's Ring (19519, -0.23 DPS) [rep]; Band of the Unicorn (7553, -3.23 DPS, sim-verified) [world_drop] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (101.0 DPS) | yes | Advisor's Ring (19519, -0.15 DPS) [rep]; Band of the Unicorn (7553, -1.74 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, -0.50 DPS) [world]; Moonshadow Stave (22458, -0.82 DPS) [quest]; Blade of Eternal Darkness (17780, -1.12 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Thorium Greatmace

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 5222211015401051-00000000000000000000-5533300000000000)

Set DPS (verified): 179.3. Weights run: 2.6s. Verify run: 1.0s. 1446 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.571 ± 0.011, crit=0.136 ± 0.008 per rating point (14 rating = 1%, 1.911 per %), hit=0.282 ± 0.003 per rating point (10 rating = 1%, 2.821 per %), spell_haste=not significant (0.356 ± 0.101), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.458 ± 0.001, arcane_power=0.542 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+4.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -0.12 DPS) [dungeon]; Feralheart Cowl (226773, -4.78 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 22.4 spell_power points (3.92 DPS) | yes | Orb of the Darkmoon (19426, -0.07 DPS) [quest]; Chains of the Lich (23125, -0.07 DPS) [dungeon]; Beads of Ogre Mojo (22149, -0.45 DPS) [quest] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 37.5 spell_power points (6.55 DPS) | yes | Warlord's Dragonhide Spaulders (231681, -0.80 DPS) [vendor]; Burial Shawl (18681, -1.46 DPS) [dungeon]; Feralheart Spaulders (226778, -1.54 DPS) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 23.4 spell_power points (4.09 DPS) | yes | Crystalline Threaded Cape (20697, -0.19 DPS) [world]; Hide of the Wild (18510, -0.64 DPS) [crafted]; Amplifying Cloak (18350, -0.94 DPS) [dungeon] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Armor (231679, -0.19 DPS) [vendor]; Robe of Everlasting Night (18385, -0.32 DPS) [dungeon]; Tunic of Undead Slaying (23089, -9.92 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -1.55 DPS) [dungeon]; Runecloth Cuffs (254123, -1.72 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -8.23 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (+2.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Gloves of the Greatfather (17721, -0.42 DPS) [crafted]; Hands of Power (13253, -2.72 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 40.4 spell_power points (7.06 DPS) | yes | Elunite Cord (272401, -1.87 DPS) [vendor]; Girdle of Insight (18504, -2.14 DPS) [crafted]; Belt of the Archmage (18405, -2.75 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 50.2 spell_power points (8.78 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.51 DPS) [vendor]; Outrider's Silk Leggings (22747, -1.98 DPS) [rep]; Sentinel's Silk Leggings (237815, -2.03 DPS, sim-verified) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 31.1 spell_power points (5.45 DPS) | yes | General's Dragonhide Boots (231682, -0.37 DPS) [pvp]; Waterspout Boots (18322, -0.47 DPS) [dungeon]; Dragonrider Boots (18102, -0.70 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.10 DPS) [quest]; Maiden's Circle (13001, -1.10 DPS) [world_drop]; Naglering (11669, -7.29 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.52 DPS) [quest]; Maiden's Circle (13001, -0.52 DPS) [world_drop]; Naglering (11669, -7.96 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+12.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Second Wind (11819, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, -1.22 DPS) [vendor]; Serenity Field (272439, -2.62 DPS) [vendor] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.55 DPS) [dungeon]; Hand of Edward the Odd (2243, -8.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

