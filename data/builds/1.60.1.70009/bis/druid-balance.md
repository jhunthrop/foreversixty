# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 32.4. Weights run: 1.6s. Verify run: 1.0s. 285 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.463 ± 0.112, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (28.4 DPS) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.38 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.5 | yes | Forest Leather Mantle (4709, -0.67 DPS) [dungeon]; Rugged Spaulders (5254, -0.67 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.84 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -0.79 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (28.5 DPS) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.04 DPS) [dungeon]; Tabitha's Cuffs (251486, -0.40 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (28.9 DPS) | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.89 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (28.6 DPS) | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.52 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 | yes | Abomination Skin Leggings (23173, -0.01 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Loop of Sacrifice (281673, -0.39 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest]; Sludge-Stained Band (286535, -0.32 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | sim-verified (30.2 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Staff of Westfall (2042, -0.14 DPS) [quest]; Living Root (6631, -2.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Idol of the Huntress

No-known-source sample (15 of 285, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers

### Band 30 (night-elf, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 64.4. Weights run: 1.7s. Verify run: 1.5s. 589 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.971 ± 0.152, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [dungeon]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.25 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.28 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 | yes | Death Speaker Mantle (6685, -0.31 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Prelacy Cape (7004, -0.13 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Guardian Armor (4256, -0.60 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.86 DPS) [quest]; Technician's Bracers (270042, -0.86 DPS) [quest]; Nightsky Wristbands (6407, -1.14 DPS, sim-verified) [dungeon] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | sim-verified (64.4 DPS) | yes | Serpent Gloves (5970, -0.23 DPS) [dungeon]; Shilly Mitts (9609, -0.23 DPS) [quest]; Fletcher's Gloves (7348, -1.12 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Belt of Arugal (6392, -0.31 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.63 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon]; Dark Ritual Leggings (270031, -1.35 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -1.92 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.15 DPS) [quest]; Lorekeeper's Ring (20431, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.40 DPS) [dungeon]; Sludge-Stained Band (286535, -0.40 DPS) [world]; Minor Channeling Ring (1449, -1.01 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (63.6 DPS) | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 216.7 | yes | Gnarled Ash Staff (791, -3.38 DPS) [dungeon]; Glimmering Staff (249392, -3.82 DPS) [crafted]; Cobalt Crusher (7730, -7.44 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 589, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (night-elf, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 85.2. Weights run: 1.8s. Verify run: 1.4s. 835 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.985 ± 0.232, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, +0.00 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -0.13 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.98 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -0.47 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 26.2 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.48 DPS) [crafted]; Crimson Silk Vest (7058, -0.92 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 | yes | Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest]; Arcane Runed Bracers (4744, -0.79 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted]; Dreamweave Gloves (10019, -1.24 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 | yes | Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.01 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Kodohide Legguards (285338, -0.75 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -1.11 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 | yes | Ring of Forlorn Spirits (2043, -0.79 DPS) [quest]; Reedknot Ring (9622, -0.92 DPS) [quest]; Minor Channeling Ring (1449, -1.00 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.05 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (79.6 DPS) | yes | Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (84.6 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Staff of Jordan (873, -4.95 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection

No-known-source sample (15 of 835, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (night-elf, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 101.5. Weights run: 1.9s. Verify run: 1.3s. 1082 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.782 ± 0.380, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.87 DPS) [dungeon]; Soothsayer's Headdress (17740, -4.69 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world_drop]; Darkspear Warding Pendant (272073, -0.55 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.90 DPS) [vendor]; Blood Guard's Crackling Leather Spaulders (220871, -0.90 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.4 | yes | Big Voodoo Cloak (8216, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.93 DPS) [vendor]; Runecloth Cloak (13860, -1.36 DPS, sim-verified) [crafted] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 75.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Knight's Crackling Leather Tunic (220868, -3.12 DPS) [vendor]; Stone Guard's Crackling Leather Tunic (220869, -3.12 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 | yes | Skycaller's Leather Bracers (252542, -0.15 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 28.6 | yes | Fletcher's Gloves (7348, -0.16 DPS) [crafted]; Shadowskin Gloves (18238, -0.16 DPS) [crafted]; Gloves of Holy Might (867, -4.70 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 39.9 | yes | Highlander's Lizardhide Girdle (20103, -0.93 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20115, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) (or Stone Guard's Crackling Leather Leggings (220865)) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Knight's Leather Pants (220858, -2.22 DPS) [vendor]; Stone Guard's Leather Pants (220859, -2.22 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) (or First Sergeant's Crackling Leather Boots (220863)) | Captain Dirgehammer [vendor] | 66.2 | yes | First Sergeant's Crackling Leather Boots (220863, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.83 DPS) [crafted]; Skycaller's Leather Boots (252471, -4.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.8 | yes | Lorekeeper's Ring (19523, -4.10 DPS) [rep]; Lorekeeper's Ring (19524, -4.45 DPS) [rep]; Ring of Forlorn Spirits (2043, -4.56 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 | yes | Lorekeeper's Ring (19523, -0.04 DPS, sim-verified) [rep]; Lorekeeper's Ring (19524, -0.62 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.74 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (102.3 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (102.5 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.15 DPS, sim-verified) [world_drop] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (102.5 DPS) | yes | Kindling Stave (11750, -0.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, -0.32 DPS, sim-verified) [dungeon]; Soulkeeper (1607, -2.49 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Spritecaster Cape; chest: Knight's Leather Armor; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Highlander's Cloth Girdle; legs: Knight's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 1082, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (night-elf, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 212.7. Weights run: 1.9s. Verify run: 1.3s. 1649 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.891 ± 0.496, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 175.7 | yes | Bloodvine Goggles (19999, -2.51 DPS) [crafted]; Waywatcher Headpiece (240088, -6.57 DPS) [vendor]; Mask of the Unforgiven (13404, -14.28 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | sim-verified (211.9 DPS) | yes | Blazefury Medallion (17111, -2.33 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -3.89 DPS) [quest]; Medallion of the Dawn (22659, -6.51 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 203.7 | yes | Blood Guard's Leather Shoulders (220853, -12.12 DPS) [vendor]; Knight-Lieutenant's Leather Shoulders (220852, -12.20 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -13.08 DPS) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 79.9 | yes | Howler's Furs (272414, -2.32 DPS) [vendor]; Stalwart Cloak (272415, -2.32 DPS) [vendor]; Earthweave Cloak (21187, -5.04 DPS, sim-verified) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | 173.0 | yes | Bloodvine Vest (19682, -5.43 DPS, sim-verified) [crafted]; Waywatcher Tunic (240091, -5.82 DPS) [vendor]; Waywatcher Vest (240067, -7.76 DPS) [vendor] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | 103.0 | yes | Waywatcher Wristguards (240084, -4.50 DPS) [vendor]; Primal Batskin Bracers (19687, -4.87 DPS) [crafted]; Rockfury Bracers (21186, -5.05 DPS, sim-verified) [quest] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 183.1 | yes | Stormshroud Gloves (21278, -9.84 DPS) [crafted]; Primal Batskin Gloves (19686, -11.68 DPS, sim-verified) [crafted]; Waywatcher Grips (240065, -12.77 DPS) [vendor] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 176.1 | yes | Knowledge of the Timbermaw (228190, -5.36 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -11.61 DPS) [vendor]; Belt of the Archmage (18405, -12.27 DPS) [crafted] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 224.8 | yes | Waywatcher Legguards (240087, -0.74 DPS, sim-verified) [vendor]; Knight's Crackling Leather Leggings (220864, -12.44 DPS) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -12.44 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 171.1 | yes | Bloodvine Boots (19684, -4.24 DPS, sim-verified) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -10.50 DPS) [vendor]; First Sergeant's Crackling Leather Boots (220863, -10.50 DPS) [vendor] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 99.6 | yes | Band of Earthen Might (21182, -0.61 DPS) [quest]; Signet Ring of the Bronze Dragonflight (234032, -0.79 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -1.08 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 94.1 | yes | Band of Earthen Might (21182, -0.13 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (211.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.88 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -3.10 DPS, sim-verified) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (211.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.77 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -2.06 DPS, sim-verified) [quest] |
| main_hand | High Warlord's Pig Poker (234548) | Sergeant Thunderhorn [vendor] | sim-verified (211.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -9.09 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Mushgog [world] | 0.0 | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor]; Enraged Idol (272428, +0.00 DPS) [vendor]; Idol of Synthesis (272429, +0.00 DPS) [vendor] |

**New at 60:** head: Waywatcher Hood; neck: Onyxia Tooth Pendant; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Ring of the Fallen God; finger2: Don Julio's Band; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: High Warlord's Pig Poker; ranged: Idol of the Moon

No-known-source sample (15 of 1649, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 30.4. Weights run: 1.6s. Verify run: 1.0s. 290 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.463 ± 0.112, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (27.6 DPS) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.5 | yes | Double-Stitched Woolen Shoulders (4314, -0.57 DPS, sim-verified) [crafted]; Forest Leather Mantle (4709, -0.67 DPS) [dungeon]; Rugged Spaulders (5254, -0.67 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.12 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -0.41 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.3 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Owl Bracers (4796, -0.04 DPS) [vendor]; Featherbead Bracers (15452, -0.04 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (28.1 DPS) | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.87 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (27.7 DPS) | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.50 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 | yes | Abomination Skin Leggings (23173, -0.00 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 | yes | Stormrider's Leather Boots (252443, -0.05 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest]; Black Pearl Ring (6332, -0.43 DPS) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.11 DPS) [quest]; Black Pearl Ring (6332, -0.23 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | sim-verified (28.6 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Hammerbone (270018, -0.40 DPS) [quest]; Living Root (6631, -1.38 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Idol of the Huntress

No-known-source sample (15 of 290, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers

### Band 30 (tauren, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 64.1. Weights run: 1.7s. Verify run: 1.4s. 600 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.971 ± 0.152, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [dungeon]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.16 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.28 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 | yes | Death Speaker Mantle (6685, -0.19 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Battle Healer's Cloak (19529, -0.13 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Guardian Armor (4256, -0.61 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.68 DPS, sim-verified) [dungeon]; Tabitha's Cuffs (251486, -0.86 DPS) [quest]; Technician's Bracers (270042, -0.86 DPS) [quest] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | sim-verified (64.1 DPS) | yes | Jutebraid Gloves (10654, -0.23 DPS) [quest]; Serpent Gloves (5970, -0.38 DPS) [dungeon]; Fletcher's Gloves (7348, -1.47 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Warsong Sash (16975, -0.21 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.71 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Dark Ritual Leggings (270031, -0.29 DPS, sim-verified) [quest]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -1.97 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Sludge-Stained Band (286535, -0.54 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.40 DPS) [world]; Black Widow Band (6199, -0.40 DPS) [world]; Electrocutioner Lagnut (9447, -1.05 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (63.1 DPS) | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 216.7 | yes | Gnarled Ash Staff (791, -3.38 DPS) [dungeon]; Glimmering Staff (249392, -3.82 DPS) [crafted]; Cobalt Crusher (7730, -7.33 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 600, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (tauren, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 84.9. Weights run: 1.8s. Verify run: 1.4s. 846 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.985 ± 0.232, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, +0.00 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -0.54 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.98 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -0.30 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 26.2 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.48 DPS) [crafted]; Crimson Silk Vest (7058, -0.92 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 | yes | Radiant Silver Bracers (4545, -0.13 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted]; Dreamweave Gloves (10019, -1.06 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 | yes | Defiler's Cloth Girdle (20166, -0.42 DPS, sim-verified) [rep]; Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Kodohide Legguards (285338, -0.43 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -0.74 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 | yes | Reedknot Ring (9622, -0.92 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.05 DPS) [vendor]; Ogremind Ring (1993, -1.19 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Reedknot Ring (9622, -0.93 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | Maraudon: Theradrim Shardling [dungeon] | sim-verified (79.5 DPS) | yes | Blazing Emblem (2802, +0.00 DPS) [dungeon]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (84.2 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Staff of Jordan (873, -4.64 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection

No-known-source sample (15 of 846, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (tauren, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 100.8. Weights run: 1.9s. Verify run: 1.3s. 1093 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.782 ± 0.380, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.87 DPS) [dungeon]; Soothsayer's Headdress (17740, -4.69 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world_drop]; Darkspear Warding Pendant (272073, -0.55 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.90 DPS) [vendor]; Blood Guard's Crackling Leather Spaulders (220871, -0.90 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.6 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.43 DPS) [crafted]; Big Voodoo Cloak (8216, -0.80 DPS) [crafted] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 75.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Knight's Crackling Leather Tunic (220868, -3.12 DPS) [vendor]; Stone Guard's Crackling Leather Tunic (220869, -3.12 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 28.6 | yes | Fletcher's Gloves (7348, -0.16 DPS) [crafted]; Shadowskin Gloves (18238, -0.16 DPS) [crafted]; Gloves of Holy Might (867, -4.94 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 39.9 | yes | Defiler's Lizardhide Girdle (20174, -0.49 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20193, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) (or Stone Guard's Crackling Leather Leggings (220865)) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Knight's Leather Pants (220858, -2.22 DPS) [vendor]; Stone Guard's Leather Pants (220859, -2.22 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) (or First Sergeant's Crackling Leather Boots (220863)) | Captain Dirgehammer [vendor] | 66.2 | yes | First Sergeant's Crackling Leather Boots (220863, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.83 DPS) [crafted]; Skycaller's Leather Boots (252471, -4.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.8 | yes | Advisor's Ring (19519, -4.10 DPS) [rep]; Advisor's Ring (19520, -4.45 DPS) [rep]; Reedknot Ring (9622, -4.67 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 | yes | Advisor's Ring (19519, +0.00 DPS, sim-verified) [rep]; Advisor's Ring (19520, -0.62 DPS) [rep]; Reedknot Ring (9622, -0.85 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (101.0 DPS) | yes | Ankh of Life (1713, +0.00 DPS, sim-verified) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Guardian Talisman (1490, -0.69 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (101.0 DPS) | yes | Frozen Heart of the Mountain (249469, -1.05 DPS, sim-verified) [crafted]; Guardian Talisman (1490, -3.83 DPS) [quest]; Ankh of Life (1713, -3.83 DPS) [dungeon] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | sim-verified (101.0 DPS) | yes | Manual Crowd Pummeler (9449, +0.00 DPS, sim-verified) [dungeon]; Kindling Stave (11750, -0.00 DPS) [dungeon]; Soulkeeper (1607, -2.49 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Deep Woodlands Cloak; chest: Knight's Leather Armor; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Defiler's Cloth Girdle; legs: Knight's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Uther's Strength; trinket2: Rune of the Guard Captain; main_hand: Thorium Greatmace

No-known-source sample (15 of 1093, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (tauren, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 210.9. Weights run: 1.9s. Verify run: 1.3s. 1658 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.891 ± 0.496, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 175.7 | yes | Bloodvine Goggles (19999, -2.51 DPS) [crafted]; Waywatcher Headpiece (240088, -6.57 DPS) [vendor]; Mask of the Unforgiven (13404, -13.79 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | sim-verified (210.3 DPS) | yes | Blazefury Medallion (17111, -2.45 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -3.89 DPS) [quest]; Medallion of the Dawn (22659, -6.51 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 203.7 | yes | Knight-Lieutenant's Leather Shoulders (220852, -11.50 DPS, sim-verified) [vendor]; Blood Guard's Leather Shoulders (220853, -12.12 DPS) [vendor]; Feralheart Spaulders (226778, -13.08 DPS) [quest] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 79.9 | yes | Howler's Furs (272414, -2.32 DPS) [vendor]; Stalwart Cloak (272415, -2.32 DPS) [vendor]; Earthweave Cloak (21187, -3.99 DPS, sim-verified) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | 173.0 | yes | Bloodvine Vest (19682, -4.95 DPS, sim-verified) [crafted]; Waywatcher Tunic (240091, -5.82 DPS) [vendor]; Waywatcher Vest (240067, -7.76 DPS) [vendor] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | 103.0 | yes | Waywatcher Wristguards (240084, -4.50 DPS) [vendor]; Primal Batskin Bracers (19687, -4.87 DPS) [crafted]; Rockfury Bracers (21186, -4.91 DPS, sim-verified) [quest] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 183.1 | yes | Stormshroud Gloves (21278, -9.84 DPS) [crafted]; Primal Batskin Gloves (19686, -12.48 DPS, sim-verified) [crafted]; Waywatcher Grips (240065, -12.77 DPS) [vendor] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 176.1 | yes | Knowledge of the Timbermaw (228190, -4.47 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -11.61 DPS) [vendor]; Belt of the Archmage (18405, -12.27 DPS) [crafted] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 224.8 | yes | Waywatcher Legguards (240087, -0.33 DPS, sim-verified) [vendor]; Knight's Crackling Leather Leggings (220864, -12.44 DPS) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -12.44 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 171.1 | yes | Bloodvine Boots (19684, -3.32 DPS, sim-verified) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -10.50 DPS) [vendor]; First Sergeant's Crackling Leather Boots (220863, -10.50 DPS) [vendor] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 99.6 | yes | Band of Earthen Might (21182, -0.61 DPS) [quest]; Signet Ring of the Bronze Dragonflight (234032, -0.79 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -1.08 DPS) [vendor] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 94.1 | yes | Band of Earthen Might (21182, -0.13 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (208.3 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.88 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (210.3 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Rune of the Guard Captain (19120, -1.62 DPS, sim-verified) [quest]; Uther's Strength (11302, -1.77 DPS) [world_drop] |
| main_hand | High Warlord's Pig Poker (234548) | Sergeant Thunderhorn [vendor] | sim-verified (210.3 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Sulfuron Hammer (17193, -9.47 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Mushgog [world] | 0.0 | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor]; Enraged Idol (272428, +0.00 DPS) [vendor]; Idol of Synthesis (272429, +0.00 DPS) [vendor] |

**New at 60:** head: Waywatcher Hood; neck: Onyxia Tooth Pendant; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Ring of the Fallen God; finger2: Don Julio's Band; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: High Warlord's Pig Poker; ranged: Idol of the Moon

No-known-source sample (15 of 1658, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

