# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 24.0. Weights run: 2.4s. Verify run: 1.0s. 225 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.591 ± 0.008, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.930 per %), hit=0.149 ± 0.001 per rating point (10 rating = 1%, 1.490 per %), spell_haste=-1.628 ± 0.069, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.003

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

### Band 30 (dwarf, 5532311100000000-000000000000000000-0000000000000000)

Set DPS (verified): 42.7. Weights run: 2.5s. Verify run: 1.0s. 366 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.528 ± 0.009, crit=0.082 ± 0.002 per rating point (14 rating = 1%, 1.147 per %), hit=0.193 ± 0.001 per rating point (10 rating = 1%, 1.927 per %), spell_haste=not significant (0.089 ± 0.078), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.943 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.03 DPS) [world_drop]; Silk Headband (7050, -0.23 DPS) [crafted]; Totemic Leather Helm (252456, -0.41 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 spell_power points (1.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.82 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.82 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.21 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.7 spell_power points (1.39 DPS) | yes | Death Speaker Mantle (6685, -0.20 DPS) [dungeon]; Fairywing Mantle (9536, -0.30 DPS) [quest]; Magician's Mantle (12998, -0.41 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.1 spell_power points (0.52 DPS) | yes | Cloak of Rot (4462, -0.09 DPS) [world]; Darkspear Raider's Cloak (272078, -0.09 DPS) [vendor]; Hillman's Cloak (3719, -0.35 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.9 spell_power points (1.61 DPS) | yes | Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Tree Bark Jacket (1486, -0.29 DPS) [dungeon]; Guardian Armor (4256, -0.38 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.91 DPS) | yes | Nightsky Wristbands (6407, -0.59 DPS) [world_drop]; Technician's Bracers (270042, -0.59 DPS) [quest]; Glowing Magical Bracelets (13106, -1.14 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 9.8 spell_power points (0.99 DPS) | yes | Shilly Mitts (9609, -0.28 DPS) [quest]; Gloves of Insight (9698, -0.28 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.73 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.2 spell_power points (1.33 DPS) | yes | Moss Cinch (6911, -0.12 DPS) [dungeon]; Belt of Arugal (6392, -0.26 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.31 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 15.8 spell_power points (1.60 DPS) | yes | Abomination Skin Leggings (23173, -0.26 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.27 DPS) [crafted]; Dark Ritual Leggings (270031, -0.67 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.7 spell_power points (1.08 DPS) | yes | Spidersilk Boots (4320, -0.16 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.21 DPS) [crafted]; Acidic Walkers (9454, -1.23 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.71 DPS) | yes | Black Widow Band (6199, -0.33 DPS) [world]; Snake Hoop (6750, -0.33 DPS) [quest]; Minor Channeling Ring (1449, -1.50 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Widow Band (6199, -0.23 DPS) [world]; Snake Hoop (6750, -0.23 DPS) [quest]; Minor Channeling Ring (1449, -0.74 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -1.28 DPS) [dungeon]; Rhahk'Zor's Hammer (5187, -1.38 DPS) [dungeon]; Manual Crowd Pummeler (9449, -3.35 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 366, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4820 Guardian Buckler

### Band 40 (dwarf, 5532311300103040-000000000000000000-0000000000000000)

Set DPS (verified): 65.9. Weights run: 3.0s. Verify run: 1.1s. 592 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.894 ± 0.023, crit=0.189 ± 0.006 per rating point (14 rating = 1%, 2.646 per %), hit=0.326 ± 0.003 per rating point (10 rating = 1%, 3.259 per %), spell_haste=-2.658 ± 0.280, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 21.5 spell_power points (2.14 DPS) | yes | Augural Shroud (2620, -0.16 DPS) [world]; Corpseshroud (10574, -0.45 DPS) [dungeon]; Spellpower Goggles Xtreme (10502, -0.67 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.4 spell_power points (1.23 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest]; Triune Amulet (7722, -0.61 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.61 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.6 spell_power points (1.85 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 17.0 spell_power points (1.70 DPS) | yes | Guardian Cloak (5965, -0.65 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.72 DPS) [vendor]; Long Silken Cloak (4326, -1.58 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 27.4 spell_power points (2.73 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.26 DPS) [crafted]; Big Voodoo Robe (8200, -0.58 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.4 spell_power points (1.13 DPS) | yes | Turtle Scale Bracers (8198, -0.09 DPS) [crafted]; Arcane Runed Bracers (4744, -0.24 DPS) [quest]; Spidertank Oilrag (9448, -0.24 DPS) [dungeon] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.39 DPS) | yes | Red Mageweave Gloves (10018, -0.40 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.49 DPS) [crafted]; Dreamweave Gloves (10019, -1.08 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.4 spell_power points (2.03 DPS) | yes | Skycaller's Leather Belt (252522, -0.50 DPS) [crafted]; Gilded Cord (254037, -0.52 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.89 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.7 spell_power points (2.46 DPS) | yes | Crimson Silk Pantaloons (7062, -0.51 DPS) [crafted]; Kodohide Legguards (285338, -0.63 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -0.85 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.39 DPS) | yes | Skycaller's Mail Boots (252563, -0.47 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.69 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -0.97 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (1.53 DPS) | yes | Ring of Forlorn Spirits (2043, -0.73 DPS) [quest]; Reedknot Ring (9622, -0.83 DPS) [quest]; Minor Channeling Ring (1449, -0.85 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (0.90 DPS) | yes | Reedknot Ring (9622, -0.20 DPS) [quest]; Minor Channeling Ring (1449, -0.22 DPS) [quest]; Ring of Forlorn Spirits (2043, -1.26 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-verified (65.9 DPS) | yes | Spellforce Rod (1664, -0.35 DPS) [world]; Mograine's Might (7723, -0.91 DPS) [dungeon]; Gut Ripper (2164, -3.83 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 592, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 50 (dwarf, 5532311300103050-010000000000000000-5300000000000000)

Set DPS (verified): 83.8. Weights run: 3.1s. Verify run: 1.3s. 762 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.855 ± 0.029, crit=0.260 ± 0.009 per rating point (14 rating = 1%, 3.643 per %), hit=0.415 ± 0.004 per rating point (10 rating = 1%, 4.152 per %), spell_haste=-2.521 ± 0.343, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.939 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.1 spell_power points (3.39 DPS) | yes | Soothsayer's Headdress (17740, -0.03 DPS) [dungeon]; Dreamweave Circlet (10041, -0.61 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.85 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.14 DPS) | yes | Horizon Choker (13085, -0.01 DPS) [world_drop]; Mindburst Medallion (11196, -0.09 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 32.1 spell_power points (3.01 DPS) | yes | Lead Surveyor's Mantle (11842, -0.48 DPS) [dungeon]; Kentic Amice (11624, -0.66 DPS) [dungeon]; Rotgrip Mantle (17732, -1.01 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 19.1 spell_power points (1.79 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.23 DPS) [dungeon]; Runecloth Cloak (13860, -0.31 DPS) [crafted]; Big Voodoo Cloak (8216, -0.60 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.1 spell_power points (3.39 DPS) | yes | Feathered Breastplate (8349, -0.80 DPS) [crafted]; Robe of the Magi (1716, -0.84 DPS) [world_drop]; Runecloth Tunic (13857, -0.91 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.6 spell_power points (1.74 DPS) | yes | Skycaller's Leather Bracers (252542, -0.24 DPS) [crafted]; Skycaller's Mail Bracers (252571, -0.24 DPS) [crafted]; Nethergeld Cuffs (254061, -0.52 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | sim-verified (83.8 DPS) | yes | Gloves of the Greatfather (17721, -0.17 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.20 DPS) [crafted]; Raider Handguards (272102, -1.07 DPS, sim-verified) [vendor] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 26.3 spell_power points (2.46 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.35 DPS) [dungeon]; Dawnspire Cord (12466, -0.38 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.6 spell_power points (2.96 DPS) | yes | Big Voodoo Pants (8202, -0.75 DPS) [crafted]; Turtle Scale Leggings (8185, -1.13 DPS) [crafted]; Red Mageweave Pants (10009, -2.41 DPS, sim-verified) [crafted] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 24.4 spell_power points (2.29 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -0.04 DPS) [crafted]; Greaves of Withering Despair (22240, -0.07 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.42 DPS) | yes | Band of the Unicorn (7553, -0.20 DPS) [world_drop]; Brainlash (6440, -0.22 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.29 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.0 spell_power points (1.41 DPS) | yes | Brainlash (6440, -0.20 DPS) [dungeon]; Lorekeeper's Ring (19523, -0.28 DPS) [rep]; Band of the Unicorn (7553, -1.35 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -0.97 DPS, sim-verified) [quest] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Mechanic's Pipehammer (9604, -0.14 DPS) [quest]; Thorium Greatmace (250613, -0.17 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 762, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

### Band 60 (dwarf, 5532311300103050-010000000000000000-5533020000000000)

Set DPS (verified): 137.8. Weights run: 3.0s. Verify run: 1.3s. 1751 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.764 ± 0.031, crit=0.300 ± 0.010 per rating point (14 rating = 1%, 4.200 per %), hit=0.482 ± 0.004 per rating point (10 rating = 1%, 4.816 per %), spell_haste=-2.370 ± 0.348, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.938 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (137.8 DPS) | yes | Crimson Felt Hat (18727, -0.26 DPS) [dungeon]; Soothsayer's Headdress (17740, -0.42 DPS) [dungeon]; Blue Dragonscale Helm (252604, -3.46 DPS, sim-verified) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 24.9 spell_power points (2.47 DPS) | yes | Beads of Ogre Mojo (22149, -0.27 DPS) [quest]; Orb of the Darkmoon (19426, -0.29 DPS) [quest]; Chains of the Lich (23125, -0.29 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 42.7 spell_power points (4.22 DPS) | yes | Darkspear Shoulderguards (272958, -0.60 DPS) [vendor]; Darkspear Shoulderpads (272103, -0.99 DPS) [vendor]; Darkspear Shoulders (272104, -0.99 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.9 spell_power points (2.66 DPS) | yes | Hide of the Wild (18510, -0.52 DPS) [crafted]; Spritecaster Cape (11623, -0.83 DPS) [dungeon]; Crystalline Threaded Cape (20697, -1.50 DPS, sim-verified) [world] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Robe of Everlasting Night (18385, -0.13 DPS) [dungeon]; Tunic of Undead Slaying (23089, -9.65 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Modest Armguards (18458, -0.84 DPS) [dungeon]; Sublime Wristguards (18497, -0.84 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -6.79 DPS, sim-verified) [world] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | 32.1 spell_power points (3.17 DPS) | yes | Hands of Power (13253, -0.15 DPS) [dungeon]; Raider Handwraps (272097, -0.34 DPS) [vendor]; Gloves of Undead Cleansing (23084, -0.63 DPS) [quest] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 46.6 spell_power points (4.61 DPS) | yes | Stormseeker's Girdle (272399, -1.22 DPS) [vendor]; Girdle of Insight (18504, -1.39 DPS) [crafted]; Belt of the Archmage (18405, -3.47 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 56.0 spell_power points (5.53 DPS) | yes | Sentinel's Lizardhide Pants (237817, -0.78 DPS) [vendor]; Red Dragonscale Leggings (252603, -1.14 DPS) [crafted]; Sentinel's Silk Leggings (237815, -2.09 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 33.2 spell_power points (3.29 DPS) | yes | Dragonrider Boots (18102, -0.30 DPS) [dungeon]; Waterspout Boots (18322, -0.36 DPS) [dungeon]; Omnicast Boots (11822, -0.40 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.70 DPS) [quest]; Maiden's Circle (13001, -0.70 DPS) [world_drop]; Naglering (11669, -6.70 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.30 DPS) [quest]; Maiden's Circle (13001, -0.30 DPS) [world_drop]; Naglering (11669, -6.95 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Second Wind (11819, -1.67 DPS, sim-verified) [dungeon] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.32 DPS) [world]; Hand of Edward the Odd (2243, -7.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Talisman of Ascendance; main_hand: Crackling Staff; ranged: Totem of Thunder

No-known-source sample (15 of 1751, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul

## Horde

### Band 20 (orc, 5510000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 23.0. Weights run: 2.4s. Verify run: 0.9s. 205 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.591 ± 0.008, crit=0.066 ± 0.002 per rating point (14 rating = 1%, 0.930 per %), hit=0.149 ± 0.001 per rating point (10 rating = 1%, 1.490 per %), spell_haste=-1.628 ± 0.069, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.935 ± 0.003

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

### Band 30 (orc, 5532311100000000-000000000000000000-0000000000000000)

Set DPS (verified): 40.8. Weights run: 2.5s. Verify run: 1.1s. 349 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.528 ± 0.009, crit=0.082 ± 0.002 per rating point (14 rating = 1%, 1.147 per %), hit=0.193 ± 0.001 per rating point (10 rating = 1%, 1.927 per %), spell_haste=not significant (0.089 ± 0.078), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.943 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Holy Shroud (2721, -0.03 DPS) [world_drop]; Silk Headband (7050, -0.23 DPS) [crafted]; Totemic Leather Helm (252456, -0.41 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 spell_power points (1.03 DPS) | yes | Crystal Starfire Medallion (5003, -0.82 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.82 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.24 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.7 spell_power points (1.39 DPS) | yes | Death Speaker Mantle (6685, -0.28 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.30 DPS) [quest]; Magician's Mantle (12998, -0.41 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.51 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Cloak of Rot (4462, -0.08 DPS) [world]; Darkspear Raider's Cloak (272078, -0.08 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 15.9 spell_power points (1.61 DPS) | yes | Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Tree Bark Jacket (1486, -0.29 DPS) [dungeon]; Guardian Armor (4256, -0.38 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (0.91 DPS) | yes | Nightsky Wristbands (6407, -0.59 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.59 DPS) [quest]; Glowing Magical Bracelets (13106, -0.87 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.1 spell_power points (1.02 DPS) | yes | Jutebraid Gloves (10654, -0.14 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.30 DPS) [crafted]; Serpent Gloves (5970, -0.31 DPS) [dungeon] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.2 spell_power points (1.33 DPS) | yes | Moss Cinch (6911, -0.12 DPS) [dungeon]; Warsong Sash (16975, -0.22 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.49 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 15.8 spell_power points (1.60 DPS) | yes | Abomination Skin Leggings (23173, -0.26 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.27 DPS) [crafted]; Dark Ritual Leggings (270031, -0.94 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.7 spell_power points (1.08 DPS) | yes | Spidersilk Boots (4320, -0.16 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.21 DPS) [crafted]; Acidic Walkers (9454, -1.27 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.71 DPS) | yes | Black Widow Band (6199, -0.33 DPS) [world]; Snake Hoop (6750, -0.33 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.39 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.61 DPS) | yes | Snake Hoop (6750, -0.23 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.29 DPS) [dungeon]; Black Widow Band (6199, -0.91 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rhahk'Zor's Hammer (5187, -0.10 DPS) [dungeon]; Reef Axe (6905, -0.10 DPS) [dungeon]; Manual Crowd Pummeler (9449, -1.84 DPS, sim-verified) [dungeon] |
| off_hand | Seedcloud Buckler (6630) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-verified (+0.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Tome of the Darkspear Prophecy (272090, -0.09 DPS) [vendor]; Alliance Outrunner Healing Rod (285348, -0.10 DPS) [world]; Orb of Mystic Insight (249394, -0.54 DPS, sim-verified) [crafted] |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Seedcloud Buckler

No-known-source sample (15 of 349, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5255 Quilboar Tomahawk; 5821 Darkstalker Boots

### Band 40 (orc, 5532311300103040-000000000000000000-0000000000000000)

Set DPS (verified): 67.3. Weights run: 3.0s. Verify run: 1.1s. 555 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.894 ± 0.023, crit=0.189 ± 0.006 per rating point (14 rating = 1%, 2.646 per %), hit=0.326 ± 0.003 per rating point (10 rating = 1%, 3.259 per %), spell_haste=-2.658 ± 0.280, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.945 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Big Voodoo Mask (8201) | Leatherworking [crafted] | 21.5 spell_power points (2.14 DPS) | yes | Spellpower Goggles Xtreme (10502, -0.05 DPS) [crafted]; Augural Shroud (2620, -0.16 DPS) [world]; Corpseshroud (10574, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.4 spell_power points (1.23 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest]; Triune Amulet (7722, -0.61 DPS) [dungeon]; Darkspear Warding Pendant (272074, -0.61 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 18.6 spell_power points (1.85 DPS) | yes | Green Silken Shoulders (7057, -0.08 DPS) [crafted]; Bloodmage Mantle (7684, -0.16 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 17.0 spell_power points (1.70 DPS) | yes | Guardian Cloak (5965, -0.65 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.72 DPS) [vendor]; Long Silken Cloak (4326, -1.29 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (67.3 DPS) | yes | Robe of Power (7054, -0.13 DPS) [crafted]; Big Voodoo Robe (8200, -0.45 DPS) [crafted]; Robe of the Magi (1716, -1.11 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 11.4 spell_power points (1.13 DPS) | yes | Radiant Silver Bracers (4545, +0.00 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Turtle Scale Bracers (8198, -0.09 DPS) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (2.39 DPS) | yes | Dreamweave Gloves (10019, -0.24 DPS) [crafted]; Red Mageweave Gloves (10018, -0.40 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.49 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 20.4 spell_power points (2.03 DPS) | yes | Skycaller's Leather Belt (252522, -0.50 DPS) [crafted]; Gilded Cord (254037, -0.52 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.76 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 24.7 spell_power points (2.46 DPS) | yes | Kodohide Legguards (285338, -0.49 DPS) [world]; Crimson Silk Pantaloons (7062, -0.51 DPS) [crafted]; Abomination Skin Leggings (23173, -0.85 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.39 DPS) | yes | Skycaller's Leather Shoes (252532, -0.47 DPS) [crafted]; Skycaller's Mail Boots (252563, -0.47 DPS) [crafted]; Mender's Leather Shoes (252533, -0.97 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.4 spell_power points (1.53 DPS) | yes | Reedknot Ring (9622, -0.83 DPS) [quest]; Voodoo Band (1996, -0.91 DPS) [world]; Black Widow Band (6199, -0.91 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (0.90 DPS) | yes | Voodoo Band (1996, -0.27 DPS) [world]; Black Widow Band (6199, -0.27 DPS) [world]; Reedknot Ring (9622, -1.03 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -0.57 DPS) [dungeon]; Windweaver Staff (7757, -0.66 DPS) [dungeon]; Gut Ripper (2164, -3.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Big Voodoo Mask; shoulder: Inquisitor's Shawl; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 555, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 50 (orc, 5532311300103050-010000000000000000-5300000000000000)

Set DPS (verified): 82.6. Weights run: 3.1s. Verify run: 1.2s. 704 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.004, intellect=0.855 ± 0.029, crit=0.260 ± 0.009 per rating point (14 rating = 1%, 3.643 per %), hit=0.415 ± 0.004 per rating point (10 rating = 1%, 4.152 per %), spell_haste=-2.521 ± 0.343, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.939 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Red Mageweave Headband (10033) | Tailoring [crafted] | 36.1 spell_power points (3.39 DPS) | yes | Soothsayer's Headdress (17740, -0.03 DPS) [dungeon]; Dreamweave Circlet (10041, -0.61 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -0.85 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.14 DPS) | yes | Horizon Choker (13085, -0.01 DPS) [world_drop]; Mindburst Medallion (11196, -0.09 DPS) [quest]; Prodigious Shadowshard Pendant (17773, -0.34 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 32.1 spell_power points (3.01 DPS) | yes | Lead Surveyor's Mantle (11842, -0.48 DPS) [dungeon]; Kentic Amice (11624, -0.66 DPS) [dungeon]; Rotgrip Mantle (17732, -1.63 DPS, sim-verified) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 19.7 spell_power points (1.85 DPS) | yes | Spritecaster Cape (11623, -0.05 DPS) [dungeon]; Mantle of Lady Falther'ess (23178, -0.28 DPS) [dungeon]; Runecloth Cloak (13860, -0.36 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 36.1 spell_power points (3.39 DPS) | yes | Stone Guard's Pulsing Breastplate (220844, -0.35 DPS) [vendor]; Feathered Breastplate (8349, -0.80 DPS) [crafted]; Robe of the Magi (1716, -0.84 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 18.6 spell_power points (1.74 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.24 DPS) [crafted] |
| hands | Raider Handguards (272102) | Creeg Bothunk [vendor] | 27.5 spell_power points (2.58 DPS) | yes | Raider Handwraps (272098, -0.16 DPS) [vendor]; Gloves of the Greatfather (17721, -0.33 DPS) [crafted]; Skycaller's Leather Gauntlets (252550, -0.36 DPS) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) (or Skycaller's Mail Belt (252589)) | Leatherworking [crafted] | 26.3 spell_power points (2.46 DPS) | yes | Skycaller's Mail Belt (252589, +0.00 DPS) [crafted]; Satyrmane Sash (17755, -0.35 DPS) [dungeon]; Dawnspire Cord (12466, -0.38 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | sim-verified (82.6 DPS) | yes | Red Mageweave Pants (10009, -0.68 DPS) [crafted]; Big Voodoo Pants (8202, -0.75 DPS) [crafted]; Stone Guard's Pulsing Legplates (220847, -0.88 DPS, sim-verified) [vendor] |
| feet | Skycaller's Leather Boots (252471) (or Skycaller's Mail Sabatons (252577)) | Leatherworking [crafted] | 24.4 spell_power points (2.29 DPS) | yes | Skycaller's Mail Sabatons (252577, +0.00 DPS) [crafted]; Earthen Silk Slippers (254013, -0.04 DPS) [crafted]; Greaves of Withering Despair (22240, -0.07 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 15.1 spell_power points (1.42 DPS) | yes | Band of the Unicorn (7553, -0.20 DPS) [world_drop]; Brainlash (6440, -0.22 DPS) [dungeon]; Advisor's Ring (19519, -0.29 DPS) [rep] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 15.0 spell_power points (1.41 DPS) | yes | Brainlash (6440, -0.20 DPS) [dungeon]; Advisor's Ring (19519, -0.28 DPS) [rep]; Band of the Unicorn (7553, -1.41 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Glowing Brightwood Staff (812) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blade of Eternal Darkness (17780, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -0.17 DPS) [crafted]; Spellforce Rod (1664, -0.45 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Red Mageweave Headband; shoulder: Ironfeather Shoulders; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Raider Handguards; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; feet: Skycaller's Leather Boots; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Glowing Brightwood Staff

No-known-source sample (15 of 704, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

### Band 60 (orc, 5532311300103050-010000000000000000-5533020000000000)

Set DPS (verified): 135.3. Weights run: 3.0s. Verify run: 1.2s. 1672 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.764 ± 0.031, crit=0.300 ± 0.010 per rating point (14 rating = 1%, 4.200 per %), hit=0.482 ± 0.004 per rating point (10 rating = 1%, 4.816 per %), spell_haste=-2.370 ± 0.348, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.938 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Coif of The Five Thunders (227002) | Saving the Best for Last [quest] | 44.3 spell_power points (4.38 DPS) | yes | Blue Dragonscale Helm (252604, -0.07 DPS) [crafted]; Warlord's Mail Helm (231663, -0.18 DPS) [pvp]; Living Crown (252561, -0.55 DPS) [crafted] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 24.9 spell_power points (2.47 DPS) | yes | Beads of Ogre Mojo (22149, -0.27 DPS) [quest]; Orb of the Darkmoon (19426, -0.29 DPS) [quest]; Chains of the Lich (23125, -0.29 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 42.7 spell_power points (4.22 DPS) | yes | Warlord's Mail Spaulders (231659, -0.52 DPS) [pvp]; Pauldrons of The Five Thunders (227003, -0.96 DPS) [quest]; Darkspear Shoulderguards (272958, -1.49 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 26.9 spell_power points (2.66 DPS) | yes | Crystalline Threaded Cape (20697, -0.38 DPS) [world]; Hide of the Wild (18510, -0.52 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.80 DPS) [quest] |
| chest | Vest of Elements (16666) | Blackrock Spire: General Drakkisath [dungeon] | sim-verified (135.3 DPS) | yes | Chestplate of Tranquility (18373, +0.00 DPS) [dungeon]; Warlord's Mail Breastplate (231662, +0.00 DPS) [vendor]; Tunic of Undead Slaying (23089, -8.43 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-verified (135.3 DPS) | yes | Modest Armguards (18458, -0.84 DPS) [dungeon]; Sublime Wristguards (18497, -0.84 DPS) [dungeon]; Wristwraps of Undead Slaying (23093, -6.86 DPS, sim-verified) [world] |
| hands | Raider Handguards (272101) | Creeg Bothunk [vendor] | 32.1 spell_power points (3.17 DPS) | yes | General's Mail Gauntlets (231660, +0.00 DPS) [pvp]; Hands of Power (13253, -0.15 DPS) [dungeon]; General's Mail Gloves (231666, -0.24 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 46.6 spell_power points (4.61 DPS) | yes | Stormseeker's Girdle (272399, -1.22 DPS) [vendor]; Girdle of Insight (18504, -1.39 DPS) [crafted]; Belt of the Archmage (18405, -2.94 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 56.0 spell_power points (5.53 DPS) | yes | General's Mail Leggings (231664, -0.54 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -0.78 DPS) [vendor]; Sentinel's Silk Leggings (237815, -2.47 DPS, sim-verified) [vendor] |
| feet | Slippers of The Five Thunders (227007) | Mokvar [vendor] | 33.2 spell_power points (3.29 DPS) | yes | General's Mail Sabatons (231661, -0.02 DPS) [vendor]; Dragonrider Boots (18102, -0.30 DPS) [dungeon]; Waterspout Boots (18322, -0.36 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (135.3 DPS) | yes | Eye of Orgrimmar (12545, -0.70 DPS) [quest]; Maiden's Circle (13001, -0.70 DPS) [world_drop]; Naglering (11669, -7.44 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (135.3 DPS) | yes | Eye of Orgrimmar (12545, -0.30 DPS) [quest]; Maiden's Circle (13001, -0.30 DPS) [world_drop]; Naglering (11669, -7.01 DPS, sim-verified) [dungeon] |
| trinket1 | Burst of Knowledge (11832) | Blackrock Depths: Ambassador Flamelash [dungeon] | sim-verified (135.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (135.3 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Royal Seal of Eldre'Thalas (18471, +0.00 DPS) [quest]; Talisman of Ascendance (22678, +0.00 DPS) [quest] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-verified (135.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hammer of Divine Might (22333, -0.24 DPS) [dungeon]; Hand of Edward the Odd (2243, -8.40 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Totem of Thunder (228176) | Pix Xizzix [vendor] | sim-verified (135.3 DPS) | yes | Totem of the Storm (272432, +0.00 DPS) [world_drop] |

**New at 60:** head: Coif of The Five Thunders; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Vest of Elements; wrist: Dryad's Wrist Bindings; hands: Raider Handguards; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Slippers of The Five Thunders; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Burst of Knowledge; trinket2: Second Wind; main_hand: Amethyst War Staff; ranged: Totem of Thunder

No-known-source sample (15 of 1672, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4196 Feathered Mantle; 4642 Star of Xil'yeh

