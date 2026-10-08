# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 29.9. Weights run: 2.0s. Verify run: 0.8s. 193 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.785 ± 0.005, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.483 per %), hit=0.102 ± 0.005 per rating point (10 rating = 1%, 1.024 per %), spell_haste=-0.856 ± 0.032, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.000, arcane_power=0.967 ± 0.003

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

Set DPS (verified): 58.2. Weights run: 2.2s. Verify run: 0.9s. 322 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.549 ± 0.007, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.584 per %), hit=0.122 ± 0.005 per rating point (10 rating = 1%, 1.224 per %), spell_haste=0.318 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.337 ± 0.001, arcane_power=0.663 ± 0.002

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

Set DPS (verified): 85.3. Weights run: 2.3s. Verify run: 0.9s. 438 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.394 ± 0.005, crit=0.067 ± 0.003 per rating point (14 rating = 1%, 0.944 per %), hit=0.171 ± 0.007 per rating point (10 rating = 1%, 1.711 per %), spell_haste=-0.334 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.341 ± 0.001, arcane_power=0.659 ± 0.002

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

Set DPS (verified): 119.7. Weights run: 2.7s. Verify run: 1.0s. 577 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.469 ± 0.009, crit=0.110 ± 0.006 per rating point (14 rating = 1%, 1.536 per %), hit=0.277 ± 0.013 per rating point (10 rating = 1%, 2.770 per %), spell_haste=not significant (0.031 ± 0.137), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.379 ± 0.001, arcane_power=0.621 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.0 spell_power points (5.67 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -0.82 DPS) [crafted]; Red Mageweave Headband (10033, -1.11 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.85 DPS) | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.61 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.97 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.4 spell_power points (4.60 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.93 DPS) [vendor]; Rotgrip Mantle (17732, -1.81 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (3.17 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.68 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted]; Big Voodoo Cloak (8216, -1.43 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.4 spell_power points (5.36 DPS) | yes | Feathered Breastplate (8349, -0.89 DPS) [crafted]; Dreamweave Vest (10021, -1.16 DPS) [crafted]; Robe of the Magi (1716, -2.78 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.7 spell_power points (2.77 DPS) | yes | Skycaller's Leather Bracers (252542, -0.27 DPS) [crafted]; Nethergeld Cuffs (254061, -0.83 DPS) [crafted]; Mender's Leather Bracers (252543, -1.02 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 25.1 spell_power points (4.74 DPS) | yes | Bloodfire Talons (12464, -0.90 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -0.92 DPS) [crafted]; Gloves of the Greatfather (17721, -2.51 DPS, sim-verified) [crafted] |
| waist | Ban'thok Sash (11662) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+1.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Satyrmane Sash (17755, -0.23 DPS) [dungeon]; Highlander's Cloth Girdle (20098, -0.77 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.43 DPS, sim-verified) [crafted] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.7 spell_power points (5.23 DPS) | yes | Big Voodoo Pants (8202, -1.51 DPS) [crafted]; Red Mageweave Pants (10009, -1.52 DPS) [crafted]; Knight's Crackling Leather Leggings (220864, -2.59 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.53 DPS) | yes | Skycaller's Leather Boots (252471, -0.72 DPS) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -1.04 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.46 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.45 DPS) | yes | Lorekeeper's Ring (19523, -0.19 DPS) [rep]; Philanthropist's Ring (281635, -3.28 DPS, sim-verified) [quest] |
| finger2 | Cyclopean Band (11824) | Blackrock Depths: Ok'thor the Breaker [dungeon] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Lorekeeper's Ring (19523, -0.05 DPS) [rep]; Philanthropist's Ring (281635, -2.07 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, -1.51 DPS, sim-verified) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mechanic's Pipehammer (9604, -0.31 DPS) [quest]; Spellforce Rod (1664, -0.57 DPS) [world]; Blade of Eternal Darkness (17780, -1.42 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Ban'thok Sash; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Cyclopean Band; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Thorium Greatmace

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 216.3. Weights run: 2.6s. Verify run: 1.0s. 1449 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.666 ± 0.014, crit=0.145 ± 0.008 per rating point (14 rating = 1%, 2.037 per %), hit=0.350 ± 0.017 per rating point (10 rating = 1%, 3.500 per %), spell_haste=not significant (0.325 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.375 ± 0.001, arcane_power=0.625 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+6.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -0.33 DPS) [dungeon]; Feralheart Cowl (226773, -5.99 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.7 spell_power points (4.76 DPS) | yes | Orb of the Darkmoon (19426, -0.33 DPS) [quest]; Beads of Ogre Mojo (22149, -0.54 DPS) [quest]; Chains of the Lich (23125, -2.06 DPS, sim-verified) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 39.0 spell_power points (7.86 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, -0.94 DPS) [pvp]; Burial Shawl (18681, -1.69 DPS) [dungeon]; Feralheart Spaulders (226778, -1.72 DPS) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.8 spell_power points (5.00 DPS) | yes | Crystalline Threaded Cape (20697, -0.44 DPS) [world]; Hide of the Wild (18510, -0.84 DPS) [crafted]; Amplifying Cloak (18350, -1.37 DPS) [dungeon] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Armor (231696, -0.19 DPS) [vendor]; Chestplate of Tranquility (18373, -0.34 DPS) [dungeon]; Tunic of Undead Slaying (23089, -13.10 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -1.75 DPS) [dungeon]; Runecloth Cuffs (254123, -1.95 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -11.02 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (+2.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Raider Handwraps (272097, -0.34 DPS) [vendor]; Hands of Power (13253, -2.57 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.2 spell_power points (8.69 DPS) | yes | Elunite Cord (272401, -2.25 DPS) [vendor]; Girdle of Insight (18504, -2.58 DPS) [crafted]; Belt of the Archmage (18405, -3.84 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.0 spell_power points (10.48 DPS) | yes | Sentinel's Silk Leggings (237815, -1.87 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -2.99 DPS, sim-verified) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 32.7 spell_power points (6.58 DPS) | yes | Marshal's Dragonhide Boots (231698, -0.47 DPS) [pvp]; Waterspout Boots (18322, -0.74 DPS) [dungeon]; Dragonrider Boots (18102, -0.81 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -1.34 DPS) [quest]; Maiden's Circle (13001, -1.34 DPS) [world_drop]; Naglering (11669, -10.29 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Songstone of Ironforge (12543, -0.60 DPS) [quest]; Maiden's Circle (13001, -0.60 DPS) [world_drop]; Naglering (11669, -10.41 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Second Wind (11819) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Lord Valthalak's Staff of Command (22335, -0.30 DPS) [dungeon]; Hand of Edward the Odd (2243, -8.96 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Second Wind; main_hand: Crackling Staff

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 4132220115501051-05000000000000000000-5053000000000000)

Set DPS (verified): 505.2. Weights run: 1.5s. Verify run: 0.8s. 1449 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.356 ± 0.012, crit=0.267 ± 0.010 per rating point (14 rating = 1%, 3.739 per %), hit=0.593 ± 0.022 per rating point (10 rating = 1%, 5.928 per %), spell_haste=3.012 ± 0.110, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.196 ± 0.000, arcane_power=0.804 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cowl (226773) | Saving the Best for Last [quest] | 36.1 spell_power points (16.85 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Field Marshal's Dragonhide Helm (231695, -1.07 DPS) [vendor]; Living Crown (252561, -2.21 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.26 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.11 DPS) [quest]; Diana's Pearl Necklace (22403, -1.97 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.1 spell_power points (16.82 DPS) | yes | Field Marshal's Dragonhide Spaulders (231699, -2.03 DPS) [pvp]; Feralheart Spaulders (226778, -3.52 DPS, sim-verified) [quest]; Burial Shawl (18681, -4.84 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.8 spell_power points (11.55 DPS) | yes | Crystalline Threaded Cape (20697, -1.56 DPS) [world]; Amplifying Cloak (18350, -3.16 DPS) [dungeon]; Hide of the Wild (18510, -3.36 DPS) [crafted] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-verified (505.2 DPS) | yes | Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Robe of Everlasting Night (18385, -0.36 DPS) [dungeon]; Tunic of Undead Slaying (23089, -14.69 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Silverwing Sentinels [rep] | sim-verified (505.2 DPS) | yes | Sublime Wristguards (18497, -4.33 DPS) [dungeon]; Runecloth Cuffs (254123, -4.80 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -12.21 DPS, sim-verified) [world] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 28.1 spell_power points (13.12 DPS) | yes | Marshal's Dragonhide Gloves (231700, -0.87 DPS) [vendor]; Gloves of the Greatfather (17721, -1.93 DPS) [crafted]; Feralheart Hands (226777, -2.10 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.8 spell_power points (18.07 DPS) | yes | Belt of the Archmage (18405, -5.38 DPS, sim-verified) [crafted]; Elunite Cord (272401, -6.63 DPS) [vendor]; Girdle of Insight (18504, -7.26 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 48.1 spell_power points (22.45 DPS) | yes | Sentinel's Lizardhide Pants (237817, -2.38 DPS) [vendor]; Sentinel's Silk Leggings (237815, -2.38 DPS) [vendor]; Skyshroud Leggings (13170, -5.27 DPS) [dungeon] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 27.7 spell_power points (12.91 DPS) | yes | Waterspout Boots (18322, +0.00 DPS, sim-verified) [dungeon]; Marshal's Dragonhide Boots (231698, -0.80 DPS) [pvp]; Knight-Lieutenant's Dragonhide Boots (227194, -1.46 DPS) [vendor] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-verified (505.2 DPS) | yes | Rune Band of Wizardry (22339, -1.86 DPS) [dungeon]; Maiden's Circle (13001, -2.53 DPS) [world_drop]; Naglering (11669, -12.72 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-verified (505.2 DPS) | yes | Rune Band of Wizardry (22339, -0.73 DPS) [dungeon]; Maiden's Circle (13001, -1.40 DPS) [world_drop]; Naglering (11669, -12.09 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (505.2 DPS) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-verified (505.2 DPS) | yes | Weakness Analyzer (272438, -3.26 DPS) [vendor]; Serenity Field (272439, -5.57 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -7.93 DPS) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-verified (505.2 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.06 DPS) [world]; Hand of Edward the Odd (2243, -25.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 5231000000000000-00000000000000000000-0000000000000000)

Set DPS (verified): 29.0. Weights run: 2.0s. Verify run: 0.7s. 183 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.785 ± 0.005, crit=0.035 ± 0.001 per rating point (14 rating = 1%, 0.483 per %), hit=0.102 ± 0.005 per rating point (10 rating = 1%, 1.024 per %), spell_haste=-0.856 ± 0.032, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.033 ± 0.000, arcane_power=0.967 ± 0.003

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

Set DPS (verified): 53.7. Weights run: 2.2s. Verify run: 0.9s. 315 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.549 ± 0.007, crit=0.042 ± 0.002 per rating point (14 rating = 1%, 0.584 per %), hit=0.122 ± 0.005 per rating point (10 rating = 1%, 1.224 per %), spell_haste=0.318 ± 0.038, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.337 ± 0.001, arcane_power=0.663 ± 0.002

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

Set DPS (verified): 85.7. Weights run: 2.3s. Verify run: 0.9s. 426 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.394 ± 0.005, crit=0.067 ± 0.003 per rating point (14 rating = 1%, 0.944 per %), hit=0.171 ± 0.007 per rating point (10 rating = 1%, 1.711 per %), spell_haste=-0.334 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.341 ± 0.001, arcane_power=0.659 ± 0.002

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

Set DPS (verified): 116.7. Weights run: 2.7s. Verify run: 0.9s. 561 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.469 ± 0.009, crit=0.110 ± 0.006 per rating point (14 rating = 1%, 1.536 per %), hit=0.277 ± 0.013 per rating point (10 rating = 1%, 2.770 per %), spell_haste=not significant (0.031 ± 0.137), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.379 ± 0.001, arcane_power=0.621 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Soothsayer's Headdress (17740) | Maraudon: Celebras the Cursed [dungeon] | 30.0 spell_power points (5.67 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -0.57 DPS) [crafted]; Dreamweave Circlet (10041, -0.82 DPS) [crafted]; Red Mageweave Headband (10033, -1.40 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.8 spell_power points (1.85 DPS) | yes | Mindburst Medallion (11196, -0.19 DPS) [quest]; Horizon Choker (13085, -0.61 DPS) [world_drop]; Prodigious Shadowshard Pendant (17773, -0.97 DPS) [quest] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 24.4 spell_power points (4.60 DPS) | yes | Kentic Amice (11624, -0.81 DPS) [dungeon]; Blood Guard's Crackling Leather Spaulders (220871, -0.93 DPS) [vendor]; Rotgrip Mantle (17732, -1.82 DPS, sim-verified) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 16.8 spell_power points (3.17 DPS) | yes | Deep Woodlands Cloak (19121, -0.11 DPS) [quest]; Mantle of Lady Falther'ess (23178, -0.68 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 28.4 spell_power points (5.36 DPS) | yes | Feathered Breastplate (8349, -0.89 DPS) [crafted]; Dreamweave Vest (10021, -1.16 DPS) [crafted]; Robe of the Magi (1716, -1.77 DPS, sim-verified) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 14.7 spell_power points (2.77 DPS) | yes | Dryad's Wrist Bindings (19596, +0.00 DPS) [pvp]; Skycaller's Leather Bracers (252542, -0.27 DPS) [crafted] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 25.1 spell_power points (4.74 DPS) | yes | Bloodfire Talons (12464, -0.90 DPS) [dungeon]; Skycaller's Leather Gauntlets (252550, -0.92 DPS) [crafted]; Gloves of the Greatfather (17721, -2.41 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Waistguard (252476) | Leatherworking [crafted] | 21.6 spell_power points (4.08 DPS) | yes | Ban'thok Sash (11662, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -0.55 DPS) [dungeon]; Defiler's Cloth Girdle (20166, -1.09 DPS) [rep] |
| legs | Spellshock Leggings (9484) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 27.7 spell_power points (5.23 DPS) | yes | Big Voodoo Pants (8202, -1.51 DPS) [crafted]; Red Mageweave Pants (10009, -1.52 DPS) [crafted]; Stone Guard's Crackling Leather Leggings (220865, -2.99 DPS, sim-verified) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (4.53 DPS) | yes | Skycaller's Leather Boots (252471, -0.72 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -1.04 DPS) [vendor]; Skycaller's Leather Shoes (252532, -1.46 DPS) [crafted] |
| finger1 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.45 DPS) | yes | Cyclopean Band (11824, -0.13 DPS) [dungeon]; Advisor's Ring (19519, -0.19 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 12.8 spell_power points (2.42 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Advisor's Ring (19519, -0.15 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (116.7 DPS) | yes | Rune of the Guard Captain (19120, -0.10 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (116.7 DPS) | yes | Rune of the Guard Captain (19120, -0.77 DPS) [quest] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (116.7 DPS) | yes | Spellforce Rod (1664, -0.57 DPS) [world]; Moonshadow Stave (22458, -0.65 DPS) [quest]; Blade of Eternal Darkness (17780, -1.66 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Soothsayer's Headdress; shoulder: Ironfeather Shoulders; back: Spritecaster Cape; chest: Acumen Robes; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Skycaller's Leather Waistguard; legs: Spellshock Leggings; finger1: Band of the Unicorn; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Thorium Greatmace

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 5232221115400051-05000000000000000000-5033010000000000)

Set DPS (verified): 219.4. Weights run: 2.6s. Verify run: 1.0s. 1446 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.003, intellect=0.666 ± 0.014, crit=0.145 ± 0.008 per rating point (14 rating = 1%, 2.037 per %), hit=0.350 ± 0.017 per rating point (10 rating = 1%, 3.500 per %), spell_haste=not significant (0.325 ± 0.254), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.375 ± 0.001, arcane_power=0.625 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Living Crown (252561) | Leatherworking [crafted] | sim-verified (+5.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Crimson Felt Hat (18727, -0.33 DPS) [dungeon]; Feralheart Cowl (226773, -5.94 DPS, sim-verified) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | 23.7 spell_power points (4.76 DPS) | yes | Orb of the Darkmoon (19426, -0.33 DPS) [quest]; Beads of Ogre Mojo (22149, -0.54 DPS) [quest]; Chains of the Lich (23125, -3.91 DPS, sim-verified) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 39.0 spell_power points (7.86 DPS) | yes | Warlord's Dragonhide Spaulders (231681, -0.94 DPS) [vendor]; Burial Shawl (18681, -1.69 DPS) [dungeon]; Feralheart Spaulders (226778, -1.72 DPS) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.8 spell_power points (5.00 DPS) | yes | Crystalline Threaded Cape (20697, -0.44 DPS) [world]; Hide of the Wild (18510, -0.84 DPS) [crafted]; Amplifying Cloak (18350, -1.37 DPS) [dungeon] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Armor (231679, -0.19 DPS) [vendor]; Chestplate of Tranquility (18373, -0.34 DPS) [dungeon]; Tunic of Undead Slaying (23089, -15.15 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -1.75 DPS) [dungeon]; Runecloth Cuffs (254123, -1.95 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -10.57 DPS, sim-verified) [world] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | sim-verified (+2.3 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; Raider Handwraps (272097, -0.34 DPS) [vendor]; Hands of Power (13253, -2.31 DPS, sim-verified) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 43.2 spell_power points (8.69 DPS) | yes | Elunite Cord (272401, -2.25 DPS) [vendor]; Girdle of Insight (18504, -2.58 DPS) [crafted]; Belt of the Archmage (18405, -4.67 DPS, sim-verified) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 52.0 spell_power points (10.48 DPS) | yes | Sentinel's Lizardhide Pants (237817, -1.87 DPS) [vendor]; Outrider's Silk Leggings (22747, -2.29 DPS) [rep]; Sentinel's Silk Leggings (237815, -2.72 DPS, sim-verified) [vendor] |
| feet | Feralheart Galoshes (226774) | Mokvar [vendor] | 32.7 spell_power points (6.58 DPS) | yes | General's Dragonhide Boots (231682, -0.47 DPS) [pvp]; Waterspout Boots (18322, -0.74 DPS) [dungeon]; Dragonrider Boots (18102, -0.81 DPS) [dungeon] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -1.34 DPS) [quest]; Maiden's Circle (13001, -1.34 DPS) [world_drop]; Naglering (11669, -10.32 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Eye of Orgrimmar (12545, -0.60 DPS) [quest]; Maiden's Circle (13001, -0.60 DPS) [world_drop]; Naglering (11669, -10.36 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+13.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -1.41 DPS) [vendor]; Serenity Field (272439, -3.02 DPS) [vendor]; Burst of Knowledge (11832, -3.42 DPS) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -0.03 DPS) [world]; Hand of Edward the Odd (2243, -11.41 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Living Crown; neck: Amulet of the Dawn; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Feralheart Galoshes; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 4132220115501051-05000000000000000000-5053000000000000)

Set DPS (verified): 510.2. Weights run: 1.5s. Verify run: 0.9s. 1446 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): spell_power=1.000 ± 0.002, intellect=0.356 ± 0.012, crit=0.267 ± 0.010 per rating point (14 rating = 1%, 3.739 per %), hit=0.593 ± 0.022 per rating point (10 rating = 1%, 5.928 per %), spell_haste=3.012 ± 0.110, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.196 ± 0.000, arcane_power=0.804 ± 0.002

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Cowl (226773) | Saving the Best for Last [quest] | 36.1 spell_power points (16.85 DPS) | yes | Crimson Felt Hat (18727, +0.00 DPS, sim-verified) [dungeon]; Warlord's Dragonhide Helm (231678, -1.07 DPS) [vendor]; Living Crown (252561, -2.21 DPS) [crafted] |
| neck | Chains of the Lich (23125) (or Orb of the Darkmoon (19426)) | Stratholme: Balzaphon [dungeon] | 22.0 spell_power points (10.26 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Amulet of the Dawn (22657, -1.11 DPS) [quest]; Diana's Pearl Necklace (22403, -1.97 DPS) [dungeon] |
| shoulder | Rugged Mantle of the Timbermaw (227808) | Meilosh [vendor] | 36.1 spell_power points (16.82 DPS) | yes | Warlord's Dragonhide Spaulders (231681, -2.03 DPS) [vendor]; Feralheart Spaulders (226778, -4.33 DPS, sim-verified) [quest]; Burial Shawl (18681, -4.84 DPS) [dungeon] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 24.8 spell_power points (11.55 DPS) | yes | Crystalline Threaded Cape (20697, -1.56 DPS) [world]; Amplifying Cloak (18350, -3.16 DPS) [dungeon]; Hide of the Wild (18510, -3.36 DPS) [crafted] |
| chest | Feralheart Vest (226776) | Saving the Best for Last [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Robe of Everlasting Night (18385, -0.36 DPS) [dungeon]; Tunic of Undead Slaying (23089, -14.40 DPS, sim-verified) [world] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sublime Wristguards (18497, -4.33 DPS) [dungeon]; Runecloth Cuffs (254123, -4.80 DPS) [crafted]; Wristwraps of Undead Slaying (23093, -11.81 DPS, sim-verified) [world] |
| hands | Hands of Power (13253) | Blackrock Spire: Quartermaster Zigris [dungeon] | 28.1 spell_power points (13.12 DPS) | yes | General's Dragonhide Gloves (231677, -0.87 DPS) [pvp]; Gloves of the Greatfather (17721, -1.93 DPS) [crafted]; Feralheart Hands (226777, -2.10 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 38.8 spell_power points (18.07 DPS) | yes | Belt of the Archmage (18405, -5.31 DPS, sim-verified) [crafted]; Elunite Cord (272401, -6.63 DPS) [vendor]; Girdle of Insight (18504, -7.26 DPS) [crafted] |
| legs | Ironfeather Leggings (252486) | Leatherworking [crafted] | 48.1 spell_power points (22.45 DPS) | yes | Sentinel's Silk Leggings (237815, -2.38 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -2.38 DPS) [vendor]; Skyshroud Leggings (13170, -5.27 DPS) [dungeon] |
| feet | Waterspout Boots (18322) | Dire Maul: Hydrospawn [dungeon] | sim-verified (510.2 DPS) | yes | General's Dragonhide Boots (231682, -0.54 DPS) [pvp]; Blood Guard's Dragonhide Boots (227188, -1.20 DPS) [pvp]; Feralheart Galoshes (226774, -5.59 DPS, sim-verified) [vendor] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -1.86 DPS) [dungeon]; Maiden's Circle (13001, -2.53 DPS) [world_drop]; Naglering (11669, -11.87 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21206) | The Path of the Invoker [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune Band of Wizardry (22339, -0.73 DPS) [dungeon]; Maiden's Circle (13001, -1.40 DPS) [world_drop]; Naglering (11669, -11.33 DPS, sim-verified) [dungeon] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+24.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Burst of Knowledge (11832, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Serenity Field (272439, +0.00 DPS) [vendor] |
| trinket2 | Briarwood Reed (12930) | Blackrock Spire: Jed Runewatcher [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Weakness Analyzer (272438, -3.26 DPS) [vendor]; Serenity Field (272439, -5.57 DPS, sim-verified) [vendor]; Burst of Knowledge (11832, -7.93 DPS) [dungeon] |
| main_hand | Lord Valthalak's Staff of Command (22335) | Blackrock Spire: Lord Valthalak [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Amethyst War Staff (20654, -1.06 DPS) [world]; Hand of Edward the Odd (2243, -26.32 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Cowl; neck: Chains of the Lich; shoulder: Rugged Mantle of the Timbermaw; back: Arcanoweave Cloak; chest: Feralheart Vest; wrist: Dryad's Wrist Bindings; hands: Hands of Power; waist: Knowledge of the Timbermaw; legs: Ironfeather Leggings; feet: Waterspout Boots; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Talisman of Ascendance; trinket2: Briarwood Reed; main_hand: Lord Valthalak's Staff of Command

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

