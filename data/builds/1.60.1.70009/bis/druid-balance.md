# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 29.9. Weights run: 1.8s. Verify run: 0.7s. 193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.785 ± 0.005, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.483 per %), hit=0.088 ± 0.001 per rating point (10 rating = 1%, 0.880 per %), spell_haste=-0.856 ± 0.032, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.000, arcane_power=0.967 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.18 DPS) | yes | Wisdom's Leather Hood (252507, -0.39 DPS) [crafted]; Pristine Circlet (253949, -0.39 DPS) [crafted]; Trapper's Leather Hood (252505, -0.84 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.1 spell_power points (1.58 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -1.06 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.08 DPS, sim-verified) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.4 spell_power points (0.57 DPS) | yes | Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.52 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.07 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.39 DPS, sim-verified) [crafted] |
| wrist | Owl Bracers (4796) (or Mindthrust Bracers (1974)) | Bernard Brubaker [vendor] | 3.9 spell_power points (0.52 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.10 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (1.07 DPS) | yes | Windfelt Gloves (5630, -0.23 DPS) [quest]; Pristine Gloves (253913, -0.23 DPS) [crafted]; Serpent Gloves (5970, -0.57 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.11 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.71 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.3 spell_power points (2.00 DPS) | yes | Dreamer's Leggings (270016, -0.35 DPS) [quest]; Wisdom's Leather Pants (252503, -0.47 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.64 DPS, sim-verified) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.1 spell_power points (1.33 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.29 DPS) [crafted]; Black Whelp Slippers (252424, -0.52 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.6 spell_power points (0.86 DPS) | yes | Lavishly Jeweled Ring (1156, -0.24 DPS) [dungeon]; Sludge-Stained Band (286535, -0.47 DPS) [world]; Volcanic Rock Ring (12053, -0.55 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.66 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.61 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.05 DPS) | yes | Channeler's Staff (4437, -0.23 DPS) [world]; Lesser Staff of the Spire (1300, -0.43 DPS) [world]; Twisted Chanter's Staff (890, -1.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Owl Bracers; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 58.2. Weights run: 2.2s. Verify run: 0.8s. 322 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.549 ± 0.007, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.584 per %), hit=0.097 ± 0.001 per rating point (10 rating = 1%, 0.975 per %), spell_haste=0.318 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.337 ± 0.001, arcane_power=0.663 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.24 DPS) | yes | Holy Shroud (2721, -0.19 DPS) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted]; Enchanter's Cowl (4322, -0.83 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.92 DPS) | yes | Crystal Starfire Medallion (5003, -1.51 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.51 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.75 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.9 spell_power points (2.61 DPS) | yes | Death Speaker Mantle (6685, -0.48 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.56 DPS) [quest]; Magician's Mantle (12998, -0.75 DPS) [world_drop] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | 5.2 spell_power points (0.97 DPS) | yes | Cloak of Rot (4462, -0.15 DPS) [world]; Darkspear Raider's Cloak (272078, -0.15 DPS) [vendor]; Hillman's Cloak (3719, -0.70 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.1 spell_power points (3.02 DPS) | yes | Guardian Armor (4256, -0.31 DPS) [crafted]; Stormrider's Leather Tunic (252510, -0.53 DPS) [crafted]; Death Speaker Robes (6682, -0.58 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.68 DPS) | yes | Nightsky Wristbands (6407, -1.07 DPS) [world_drop]; Technician's Bracers (270042, -1.07 DPS) [quest]; Glowing Magical Bracelets (13106, -2.16 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 10.0 spell_power points (1.88 DPS) | yes | Shilly Mitts (9609, -0.57 DPS) [quest]; Gloves of Insight (9698, -0.57 DPS) [quest]; Stormrider's Leather Gloves (252498, -1.49 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (2.49 DPS) | yes | Moss Cinch (6911, -0.24 DPS) [dungeon]; Belt of Arugal (6392, -0.49 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.64 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 16.0 spell_power points (3.00 DPS) | yes | Abomination Skin Leggings (23173, -0.49 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.51 DPS) [crafted]; Dark Ritual Leggings (270031, -2.17 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.03 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -2.02 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (1.31 DPS) | yes | Black Widow Band (6199, -0.59 DPS) [world]; Snake Hoop (6750, -0.59 DPS) [quest]; Minor Channeling Ring (1449, -2.36 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (58.2 DPS) | yes | Black Widow Band (6199, -0.40 DPS) [world]; Snake Hoop (6750, -0.40 DPS) [quest]; Minor Channeling Ring (1449, -1.53 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -2.38 DPS) [dungeon]; Rhahk'Zor's Hammer (5187, -2.57 DPS) [dungeon]; Manual Crowd Pummeler (9449, -5.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 5232221115400030-00000000000000000000-0000000000000000)

Set DPS (verified): 85.3. Weights run: 2.2s. Verify run: 0.8s. 438 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.394 ± 0.005, crit=0.067 ± 0.003 per rating point (14 rating = 1%, 0.944 per %), hit=0.131 ± 0.001 per rating point (10 rating = 1%, 1.309 per %), spell_haste=-0.334 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.341 ± 0.001, arcane_power=0.659 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.28 DPS) | yes | Augural Shroud (2620, -1.23 DPS) [world]; Big Voodoo Mask (8201, -1.32 DPS) [crafted]; Living Cowl (5608, -1.63 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (1.91 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.88 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.35 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.35 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (2.56 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.5 spell_power points (2.56 DPS) | yes | Guardian Cloak (5965, -0.93 DPS) [crafted]; Icy Cloak (4327, -1.13 DPS) [crafted]; Long Silken Cloak (4326, -2.13 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (+1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Elemental Raiment (9434, -0.11 DPS) [world_drop]; Robe of Power (7054, -0.57 DPS) [crafted]; Robe of the Magi (1716, -0.95 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Leather Bracers (4260, -0.13 DPS) [crafted]; Condor Bracers (15864, -0.41 DPS) [quest]; Arcane Runed Bracers (4744, -1.36 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Dreamweave Gloves (10019, -1.71 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.83 DPS) [crafted]; Red Mageweave Gloves (10018, -1.85 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 15.6 spell_power points (3.17 DPS) | yes | Star Belt (4329, -0.52 DPS) [crafted]; Deathmage Sash (10771, -0.54 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.65 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.7 spell_power points (3.81 DPS) | yes | Dark Ritual Leggings (270031, -0.96 DPS) [quest]; Kodohide Legguards (285338, -1.10 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -1.14 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Skycaller's Leather Shoes (252532, -1.72 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.70 DPS) [crafted]; Gilded Slippers (254001, -2.90 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (2.52 DPS) | yes | Ring of Forlorn Spirits (2043, -0.89 DPS) [quest]; Reedknot Ring (9622, -1.09 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.30 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.83 DPS) | yes | Reedknot Ring (9622, -0.41 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.61 DPS) [vendor]; Ring of Forlorn Spirits (2043, -1.23 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spellforce Rod (1664, -0.20 DPS) [world]; Scorn's Focal Dagger (23168, -2.44 DPS) [dungeon]; Manual Crowd Pummeler (9449, -5.33 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 124.2. Weights run: 2.4s. Verify run: 0.9s. 577 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.641 ± 0.009, crit=0.115 ± 0.006 per rating point (14 rating = 1%, 1.607 per %), hit=0.210 ± 0.003 per rating point (10 rating = 1%, 2.105 per %), spell_haste=not significant (0.467 ± 0.142), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.376 ± 0.001, arcane_power=0.624 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 32.6 spell_power points (6.31 DPS) | yes | Dreamweave Circlet (10041, -1.01 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.09 DPS) [crafted]; Red Mageweave Headband (10033, -2.30 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (2.10 DPS) | yes | Mindburst Medallion (11196, -0.19 DPS) [quest]; Horizon Choker (13085, -0.36 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.86 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 27.8 spell_power points (5.38 DPS) | yes | Kentic Amice (11624, -1.06 DPS) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.41 DPS) [vendor]; Rotgrip Mantle (17732, -2.17 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.8 spell_power points (3.45 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.59 DPS) [dungeon]; Runecloth Cloak (13860, -0.72 DPS) [crafted]; Big Voodoo Cloak (8216, -1.37 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 31.8 spell_power points (6.15 DPS) | yes | Feathered Breastplate (8349, -1.24 DPS) [crafted]; Runecloth Tunic (13857, -1.50 DPS) [crafted]; Robe of the Magi (1716, -1.72 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 16.4 spell_power points (3.17 DPS) | yes | Skycaller's Leather Bracers (252542, -0.37 DPS) [crafted]; Nethergeld Cuffs (254061, -0.95 DPS) [crafted]; Bloodband Bracers (11469, -1.09 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 27.3 spell_power points (5.29 DPS) | yes | Skycaller's Leather Gauntlets (252550, -1.08 DPS) [crafted]; Bloodfire Talons (12464, -1.19 DPS) [dungeon]; Gloves of the Greatfather (17721, -1.59 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 23.7 spell_power points (4.58 DPS) | yes | Satyrmane Sash (17755, -0.63 DPS) [dungeon]; Ban'thok Sash (11662, -0.86 DPS) [dungeon]; Dawnspire Cord (12466, -1.07 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.4 spell_power points (5.69 DPS) | yes | Red Mageweave Pants (10009, -1.49 DPS) [crafted]; Big Voodoo Pants (8202, -1.55 DPS) [crafted]; Knight's Crackling Leather Leggings (220864, -2.84 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.64 DPS) | yes | Sergeant Major's Crackling Leather Boots (220862, -0.87 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.26 DPS) [crafted]; Skycaller's Leather Boots (252471, -2.44 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (2.68 DPS) | yes | Band of the Unicorn (7553, -0.16 DPS) [world_drop]; Lorekeeper's Ring (19523, -0.36 DPS) [rep]; Brainlash (6440, -0.82 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.5 spell_power points (2.61 DPS) | yes | Lorekeeper's Ring (19523, -0.29 DPS) [rep]; Brainlash (6440, -0.75 DPS) [dungeon]; Band of the Unicorn (7553, -1.90 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (124.2 DPS) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (124.2 DPS) | yes | Mark of the Chosen (17774, -1.60 DPS, sim-verified) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (124.2 DPS) | yes | Mechanic's Pipehammer (9604, -0.15 DPS) [quest]; Spellforce Rod (1664, -0.58 DPS) [world]; Blade of Eternal Darkness (17780, -2.38 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Thorium Greatmace

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 220.5. Weights run: 2.4s. Verify run: 0.9s. 1449 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.729 ± 0.014, crit=0.145 ± 0.008 per rating point (14 rating = 1%, 2.025 per %), hit=0.260 ± 0.003 per rating point (10 rating = 1%, 2.598 per %), spell_haste=not significant (0.111 ± 0.271), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.370 ± 0.001, arcane_power=0.630 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Lieutenant Commander's Dragonhide Helm (227192, -0.39 DPS) [vendor]; Feralheart Cowl (226773, -5.52 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 24.5 spell_power points (4.90 DPS) | yes | Orb of the Darkmoon (19426, -0.50 DPS) [quest]; Beads of Ogre Mojo (22149, -0.55 DPS) [quest]; Chains of the Lich (23125, -1.94 DPS, sim-verified) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 40.0 spell_power points (7.99 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, -0.95 DPS) [pvp]; Darkspear Shoulders (272104, -1.65 DPS) [vendor]; Darkspear Shoulderpads (272103, -2.42 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.4 spell_power points (4.89 DPS) | yes | Crystalline Threaded Cape (20697, -0.30 DPS) [world]; Hide of the Wild (18510, -0.63 DPS) [crafted]; Spritecaster Cape (11623, -1.21 DPS) [dungeon] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Armor (231696, -0.19 DPS) [vendor]; Chestplate of Tranquility (18373, -0.31 DPS) [dungeon]; Tunic of Undead Slaying (23089, -13.04 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -1.71 DPS) [dungeon]; Runecloth Cuffs (254123, -1.91 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -11.07 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (+3.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Raider Handwraps (272097, -0.16 DPS) [vendor]; Hands of Power (13253, -3.39 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.6 spell_power points (8.73 DPS) | yes | Elunite Cord (272401, -2.03 DPS) [vendor]; Girdle of Insight (18504, -2.37 DPS) [crafted]; Belt of the Archmage (18405, -4.43 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 53.2 spell_power points (10.63 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.96 DPS) [vendor]; Sentinel's Silk Leggings (22752, -2.26 DPS) [rep] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 33.7 spell_power points (6.73 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.49 DPS) [pvp]; Dragonrider Boots (18102, -0.80 DPS) [dungeon]; Waterspout Boots (18322, -0.86 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.38 DPS) [quest]; Maiden's Circle (13001, -1.38 DPS) [world_drop]; Naglering (11669, -10.07 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.60 DPS) [quest]; Maiden's Circle (13001, -0.60 DPS) [world_drop]; Naglering (11669, -9.93 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.40 DPS) [vendor]; Serenity Field (272439, -3.00 DPS) [vendor]; Burst of Knowledge (11832, -3.40 DPS) [dungeon] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.53 DPS) [world]; Hand of Edward the Odd (2243, -9.62 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Crackling Staff

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 29.0. Weights run: 1.8s. Verify run: 0.7s. 183 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.785 ± 0.005, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.483 per %), hit=0.088 ± 0.001 per rating point (10 rating = 1%, 0.880 per %), spell_haste=-0.856 ± 0.032, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.000, arcane_power=0.967 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | 9.0 spell_power points (1.18 DPS) | yes | Wisdom's Leather Hood (252507, -0.39 DPS) [crafted]; Pristine Circlet (253949, -0.39 DPS) [crafted]; Trapper's Leather Hood (252505, -0.83 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 12.1 spell_power points (1.58 DPS) | yes | Reinforced Woolen Shoulders (4315, -1.05 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.06 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.4 spell_power points (0.57 DPS) | yes | Sanguine Cape (14376, -0.16 DPS) [world_drop]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Heavy Woolen Cloak (4311, -0.49 DPS, sim-verified) [crafted] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (+0.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Totemic Leather Armor (252435, -0.07 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.39 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.7 spell_power points (0.62 DPS) | yes | Mindthrust Bracers (1974, -0.10 DPS) [dungeon]; Featherbead Bracers (15452, -0.10 DPS) [quest]; Owl Bracers (4796, -0.50 DPS, sim-verified) [vendor] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 8.1 spell_power points (1.07 DPS) | yes | Pristine Gloves (253913, -0.23 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.26 DPS) [crafted]; Serpent Gloves (5970, -0.56 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (+0.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Novice Arcanist's Sash (253885, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.11 DPS) [world_drop]; Stormrider's Leather Belt (252432, -0.72 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 15.3 spell_power points (2.00 DPS) | yes | Wisdom's Leather Pants (252503, -0.47 DPS) [crafted]; Stormrider's Leather Pants (252502, -0.60 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.60 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.1 spell_power points (1.33 DPS) | yes | Stormrider's Leather Boots (252443, -0.03 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.29 DPS) [crafted]; Black Whelp Slippers (252424, -0.52 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.14 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.35 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 4.7 spell_power points (0.62 DPS) | yes | Sludge-Stained Band (286535, -0.22 DPS) [world]; Volcanic Rock Ring (12053, -0.31 DPS) [world_drop]; Loop of Sacrifice (281673, -0.51 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | The Deadmines: Rhahk'Zor [dungeon] | 8.0 spell_power points (1.05 DPS) | yes | Twisted Chanter's Staff (890, -0.02 DPS) [world_drop]; Channeler's Staff (4437, -0.23 DPS) [world]; Gnarled Necromancer's Staff (251534, -1.15 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 5232221112000000-00000000000000000000-0000000000000000)

Set DPS (verified): 53.7. Weights run: 2.2s. Verify run: 0.8s. 315 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.549 ± 0.007, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.584 per %), hit=0.097 ± 0.001 per rating point (10 rating = 1%, 0.975 per %), spell_haste=0.318 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.337 ± 0.001, arcane_power=0.663 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (2.24 DPS) | yes | Holy Shroud (2721, -0.19 DPS) [world_drop]; Silk Headband (7050, -0.56 DPS) [crafted]; Enchanter's Cowl (4322, -1.21 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.3 spell_power points (1.92 DPS) | yes | Crystal Starfire Medallion (5003, -1.51 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.51 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.94 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.9 spell_power points (2.61 DPS) | yes | Fairywing Mantle (9536, -0.56 DPS) [quest]; Death Speaker Mantle (6685, -0.60 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.75 DPS) [world_drop] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.93 DPS) | yes | Windsong Drape (15468, +0.00 DPS) [quest]; Cloak of Rot (4462, -0.11 DPS) [world]; Darkspear Raider's Cloak (272078, -0.11 DPS) [vendor] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 16.1 spell_power points (3.02 DPS) | yes | Guardian Armor (4256, -0.51 DPS, sim-verified) [crafted]; Stormrider's Leather Tunic (252510, -0.53 DPS) [crafted]; Death Speaker Robes (6682, -0.58 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.68 DPS) | yes | Nightsky Wristbands (6407, -1.07 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.07 DPS) [quest]; Glowing Magical Bracelets (13106, -2.67 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 10.1 spell_power points (1.89 DPS) | yes | Stormrider's Leather Gloves (252498, -0.54 DPS) [crafted]; Serpent Gloves (5970, -0.58 DPS) [dungeon]; Jutebraid Gloves (10654, -0.60 DPS, sim-verified) [quest] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 13.3 spell_power points (2.49 DPS) | yes | Moss Cinch (6911, -0.24 DPS) [dungeon]; Warsong Sash (16975, -0.43 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.73 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 16.0 spell_power points (3.00 DPS) | yes | Abomination Skin Leggings (23173, -0.49 DPS) [dungeon]; Stormrider's Leather Pants (252502, -0.51 DPS) [crafted]; Dark Ritual Leggings (270031, -0.91 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.8 spell_power points (2.03 DPS) | yes | Spidersilk Boots (4320, -0.31 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.39 DPS) [crafted]; Acidic Walkers (9454, -2.12 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (1.31 DPS) | yes | Black Widow Band (6199, -0.59 DPS) [world]; Snake Hoop (6750, -0.59 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.69 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (1.12 DPS) | yes | Snake Hoop (6750, -0.40 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.51 DPS) [dungeon]; Black Widow Band (6199, -2.41 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Scorn's Focal Dagger (23168) | Scarlet Monastery: Scorn [dungeon] | sim-verified (53.7 DPS) | yes | Rhahk'Zor's Hammer (5187, -0.19 DPS) [dungeon]; Glimmering Staff (249392, -0.55 DPS) [crafted]; Manual Crowd Pummeler (9449, -2.32 DPS, sim-verified) [dungeon] |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 10.3 spell_power points (1.92 DPS) | yes | Orb of Souls (249395, -1.18 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -1.18 DPS) [world]; Tome of the Darkspear Prophecy (272090, -1.42 DPS, sim-verified) [vendor] |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; main_hand: Scorn's Focal Dagger; off_hand: Orb of Mystic Insight

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 5232221115400030-00000000000000000000-0000000000000000)

Set DPS (verified): 85.7. Weights run: 2.2s. Verify run: 0.8s. 426 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.394 ± 0.005, crit=0.067 ± 0.003 per rating point (14 rating = 1%, 0.944 per %), hit=0.131 ± 0.001 per rating point (10 rating = 1%, 1.309 per %), spell_haste=-0.334 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.341 ± 0.001, arcane_power=0.659 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (4.28 DPS) | yes | Augural Shroud (2620, -1.23 DPS) [world]; Big Voodoo Mask (8201, -1.32 DPS) [crafted]; Living Cowl (5608, -1.63 DPS) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 spell_power points (1.91 DPS) | yes | Prodigious Shadowshard Pendant (17773, -0.87 DPS, sim-verified) [quest]; Triune Amulet (7722, -1.35 DPS) [dungeon]; Darkspear Warding Pendant (272074, -1.35 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.5 spell_power points (2.56 DPS) | yes | Green Silken Shoulders (7057, -0.04 DPS) [crafted]; Inquisitor's Shawl (19507, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.33 DPS) [quest] |
| back | Mantle of Lady Falther'ess (23178) | Razorfen Downs: Lady Falther'ess [dungeon] | 12.5 spell_power points (2.56 DPS) | yes | Guardian Cloak (5965, -0.93 DPS) [crafted]; Icy Cloak (4327, -1.13 DPS) [crafted]; Long Silken Cloak (4326, -2.22 DPS, sim-verified) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (85.7 DPS) | yes | Elemental Raiment (9434, -0.11 DPS) [world_drop]; Robe of Power (7054, -0.57 DPS) [crafted]; Robe of the Magi (1716, -1.35 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.83 DPS) | yes | Dryad's Wrist Bindings (19597, +0.00 DPS) [pvp]; Radiant Silver Bracers (4545, -0.38 DPS) [quest]; Guardian Leather Bracers (4260, -0.81 DPS, sim-verified) [crafted] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Dreamweave Gloves (10019, -1.29 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.83 DPS) [crafted]; Red Mageweave Gloves (10018, -1.85 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 15.6 spell_power points (3.17 DPS) | yes | Star Belt (4329, -0.52 DPS) [crafted]; Deathmage Sash (10771, -0.54 DPS) [dungeon]; Skycaller's Leather Belt (252522, -0.65 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 18.7 spell_power points (3.81 DPS) | yes | Dark Ritual Leggings (270031, -0.96 DPS) [quest]; Crimson Silk Pantaloons (7062, -1.14 DPS) [crafted]; Kodohide Legguards (285338, -1.17 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.89 DPS) | yes | Skycaller's Leather Shoes (252532, -1.19 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -2.70 DPS) [crafted]; Gilded Slippers (254001, -2.90 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.4 spell_power points (2.52 DPS) | yes | Reedknot Ring (9622, -1.09 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.30 DPS) [vendor]; Sludge-Stained Band (286535, -1.91 DPS) [world] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.83 DPS) | yes | Sea Giant's Toe Ring (274746, -0.61 DPS) [vendor]; Sludge-Stained Band (286535, -1.22 DPS) [world]; Reedknot Ring (9622, -1.70 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spellforce Rod (1664) | Venture Co. Surveyor [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Scorn's Focal Dagger (23168, -2.24 DPS) [dungeon]; Rhahk'Zor's Hammer (5187, -2.44 DPS) [dungeon]; Manual Crowd Pummeler (9449, -6.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; back: Mantle of Lady Falther'ess; chest: Dreamweave Vest; hands: Gloves of the Greatfather; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; main_hand: Spellforce Rod

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 5232221115400051-05000000000000000000-2000000000000000)

Set DPS (verified): 122.3. Weights run: 2.4s. Verify run: 0.8s. 561 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.641 ± 0.009, crit=0.115 ± 0.006 per rating point (14 rating = 1%, 1.607 per %), hit=0.210 ± 0.003 per rating point (10 rating = 1%, 2.105 per %), spell_haste=not significant (0.467 ± 0.142), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.376 ± 0.001, arcane_power=0.624 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 32.6 spell_power points (6.31 DPS) | yes | Red Mageweave Headband (10033, -0.15 DPS) [crafted]; Dreamweave Circlet (10041, -1.01 DPS) [crafted]; Spellpower Goggles Xtreme Plus (15999, -1.09 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.8 spell_power points (2.10 DPS) | yes | Mindburst Medallion (11196, -0.19 DPS) [quest]; Horizon Choker (13085, -0.36 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.86 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 27.8 spell_power points (5.38 DPS) | yes | Kentic Amice (11624, -1.06 DPS) [dungeon]; Blood Guard's Crackling Leather Spaulders (220871, -1.41 DPS) [vendor]; Rotgrip Mantle (17732, -2.26 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 17.8 spell_power points (3.45 DPS) | yes | Deep Woodlands Cloak (19121, -0.01 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.59 DPS) [dungeon]; Runecloth Cloak (13860, -0.72 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 31.8 spell_power points (6.15 DPS) | yes | Feathered Breastplate (8349, -1.24 DPS) [crafted]; Robe of the Magi (1716, -1.40 DPS, sim-verified) [world_drop]; Runecloth Tunic (13857, -1.50 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 16.4 spell_power points (3.17 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.37 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 27.3 spell_power points (5.29 DPS) | yes | Skycaller's Leather Gauntlets (252550, -1.08 DPS) [crafted]; Bloodfire Talons (12464, -1.19 DPS) [dungeon]; Gloves of the Greatfather (17721, -1.49 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 23.7 spell_power points (4.58 DPS) | yes | Satyrmane Sash (17755, -0.63 DPS) [dungeon]; Ban'thok Sash (11662, -0.86 DPS) [dungeon]; Dawnspire Cord (12466, -1.07 DPS) [dungeon] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 29.4 spell_power points (5.69 DPS) | yes | Red Mageweave Pants (10009, -1.49 DPS) [crafted]; Big Voodoo Pants (8202, -1.55 DPS) [crafted]; Stone Guard's Crackling Leather Leggings (220865, -2.59 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.64 DPS) | yes | First Sergeant's Crackling Leather Boots (220863, -0.87 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.26 DPS) [crafted]; Skycaller's Leather Boots (252471, -2.00 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.8 spell_power points (2.68 DPS) | yes | Band of the Unicorn (7553, -0.16 DPS) [world_drop]; Advisor's Ring (19519, -0.36 DPS) [rep]; Brainlash (6440, -0.82 DPS) [dungeon] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | 13.5 spell_power points (2.61 DPS) | yes | Advisor's Ring (19519, -0.29 DPS) [rep]; Brainlash (6440, -0.75 DPS) [dungeon]; Band of the Unicorn (7553, -2.10 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (122.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (122.3 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (122.3 DPS) | yes | Spellforce Rod (1664, -0.58 DPS) [world]; Glowing Brightwood Staff (812, -0.85 DPS) [world_drop]; Blade of Eternal Darkness (17780, -2.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger2: Cyclopean Band; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Thorium Greatmace

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 220.8. Weights run: 2.4s. Verify run: 0.9s. 1446 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.729 ± 0.014, crit=0.145 ± 0.008 per rating point (14 rating = 1%, 2.025 per %), hit=0.260 ± 0.003 per rating point (10 rating = 1%, 2.598 per %), spell_haste=not significant (0.111 ± 0.271), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.370 ± 0.001, arcane_power=0.630 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+5.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Champion's Dragonhide Helm (227186, -0.39 DPS) [pvp]; Feralheart Cowl (226773, -5.55 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 24.5 spell_power points (4.90 DPS) | yes | Orb of the Darkmoon (19426, -0.50 DPS) [quest]; Beads of Ogre Mojo (22149, -0.55 DPS) [quest]; Chains of the Lich (23125, -1.97 DPS, sim-verified) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 40.0 spell_power points (7.99 DPS) | yes | Warlord's Dragonhide Spaulders (231681, -0.95 DPS) [vendor]; Darkspear Shoulderpads (272103, -1.65 DPS) [vendor]; Darkspear Shoulders (272104, -1.65 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.4 spell_power points (4.89 DPS) | yes | Crystalline Threaded Cape (20697, -0.30 DPS) [world]; Hide of the Wild (18510, -0.63 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.17 DPS) [quest] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Armor (231679, -0.19 DPS) [vendor]; Chestplate of Tranquility (18373, -0.31 DPS) [dungeon]; Tunic of Undead Slaying (23089, -13.29 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -1.71 DPS) [dungeon]; Runecloth Cuffs (254123, -1.91 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -10.57 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Raider Handwraps (272097, -0.16 DPS) [vendor]; Hands of Power (13253, -2.63 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.6 spell_power points (8.73 DPS) | yes | Elunite Cord (272401, -2.03 DPS) [vendor]; Girdle of Insight (18504, -2.37 DPS) [crafted]; Belt of the Archmage (18405, -4.14 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 53.2 spell_power points (10.63 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.96 DPS) [vendor]; Outrider's Silk Leggings (22747, -2.26 DPS) [rep]; Sentinel's Silk Leggings (237815, -3.06 DPS, sim-verified) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 33.7 spell_power points (6.73 DPS) | yes | General's Dragonhide Boots (231682, -0.49 DPS) [pvp]; Dragonrider Boots (18102, -0.80 DPS) [dungeon]; Waterspout Boots (18322, -0.86 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.38 DPS) [quest]; Maiden's Circle (13001, -1.38 DPS) [world_drop]; Naglering (11669, -9.61 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.60 DPS) [quest]; Maiden's Circle (13001, -0.60 DPS) [world_drop]; Naglering (11669, -9.93 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.40 DPS) [vendor]; Serenity Field (272439, -3.00 DPS) [vendor]; Burst of Knowledge (11832, -3.40 DPS) [dungeon] |
| main_hand | Amethyst War Staff (20654) | Azure Templar [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hammer of Divine Might (22333, -0.52 DPS) [dungeon]; Hand of Edward the Odd (2243, -9.72 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Amethyst War Staff

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

