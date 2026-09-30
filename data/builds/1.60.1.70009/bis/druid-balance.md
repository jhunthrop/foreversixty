# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 33.1. Weights run: 1.6s. Verify run: 1.1s. 162 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.323 ± 0.021, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (31.3 DPS) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.40 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 | yes | Double-Stitched Woolen Shoulders (4314, -0.46 DPS) [crafted]; Forest Leather Mantle (4709, -0.87 DPS) [world_drop]; Reinforced Woolen Shoulders (4315, -0.99 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.11 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -0.88 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (31.2 DPS) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.04 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.32 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (31.8 DPS) | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.91 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (31.4 DPS) | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.54 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 | yes | Abomination Skin Leggings (23173, -0.06 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 | yes | Stormrider's Leather Boots (252443, -0.09 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Loop of Sacrifice (281673, -0.39 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest]; Sludge-Stained Band (286535, -0.42 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | Westfall: Rhahk'Zor [dungeon] | 8.0 | yes | Gnarled Necromancer's Staff (251534, -0.42 DPS) [quest]; Twisted Chanter's Staff (890, -0.47 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.50 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Rhahk'Zor's Hammer; ranged: Idol of the Huntress

No-known-source sample (15 of 162, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (night-elf, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 64.4. Weights run: 1.9s. Verify run: 1.4s. 285 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.753 ± 0.028, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.25 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 | yes | Death Speaker Mantle (6685, -0.31 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Prelacy Cape (7004, -0.13 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Guardian Armor (4256, -0.60 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Glowing Magical Bracelets (13106, -0.65 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.86 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.86 DPS) [quest] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | sim-verified (64.4 DPS) | yes | Serpent Gloves (5970, -0.23 DPS) [dungeon]; Shilly Mitts (9609, -0.23 DPS) [quest]; Fletcher's Gloves (7348, -1.12 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Belt of Arugal (6392, -0.31 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.63 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon]; Dark Ritual Leggings (270031, -1.35 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -1.92 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.15 DPS) [quest]; Lorekeeper's Ring (20431, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.40 DPS) [dungeon]; Sludge-Stained Band (286535, -0.40 DPS) [world]; Minor Channeling Ring (1449, -1.01 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (59.0 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, -1.63 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (63.6 DPS) | yes | Rhahk'Zor's Hammer (5187, +0.00 DPS) [dungeon]; Glimmering Staff (249392, +0.00 DPS) [crafted]; Mechanic's Pipehammer (9604, -3.83 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 285, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 85.2. Weights run: 1.8s. Verify run: 1.3s. 399 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.733 ± 0.044, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, +0.00 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -0.13 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.98 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -0.47 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.48 DPS) [crafted]; Elemental Raiment (9434, -0.67 DPS) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 | yes | Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest]; Arcane Runed Bracers (4744, -0.79 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted]; Dreamweave Gloves (10019, -1.24 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 | yes | Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.01 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Kodohide Legguards (285338, -0.75 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -1.11 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 | yes | Ring of Forlorn Spirits (2043, -0.79 DPS) [quest]; Reedknot Ring (9622, -0.92 DPS) [quest]; Minor Channeling Ring (1449, -1.00 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.05 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (79.1 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Rune of Duty (21567, -1.00 DPS, sim-verified) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (84.6 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Mograine's Might (7723, +0.00 DPS) [dungeon]; Mechanic's Pipehammer (9604, -5.09 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life

No-known-source sample (15 of 399, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 106.6. Weights run: 1.9s. Verify run: 1.5s. 531 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.310 ± 0.074, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (103.5 DPS) | yes | Soothsayer's Headdress (17740, -0.82 DPS) [dungeon]; Red Mageweave Headband (10033, -0.86 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -2.29 DPS, sim-verified) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world_drop]; Darkspear Warding Pendant (272073, -0.55 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Crackling Leather Spaulders (220870) | Captain Dirgehammer [vendor] | sim-verified (103.1 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -1.93 DPS, sim-verified) [vendor]; Ironfeather Shoulders (15067, -3.75 DPS) [crafted]; Rotgrip Mantle (17732, -4.15 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.4 | yes | Big Voodoo Cloak (8216, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.93 DPS) [vendor]; Runecloth Cloak (13860, -1.22 DPS, sim-verified) [crafted] |
| chest | Knight's Crackling Leather Tunic (220868) | Captain Dirgehammer [vendor] | sim-verified (102.8 DPS) | yes | Knight's Leather Armor (220854, -1.59 DPS, sim-verified) [vendor]; Acumen Robes (17775, -1.61 DPS) [quest]; Robe of the Magi (1716, -2.45 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 28.6 | yes | Fletcher's Gloves (7348, -0.16 DPS) [crafted]; Shadowskin Gloves (18238, -0.16 DPS) [crafted]; Gloves of Holy Might (867, -4.85 DPS, sim-verified) [world_drop] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 39.9 | yes | Highlander's Lizardhide Girdle (20103, -1.04 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20115, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 89.7 | yes | Knight's Leather Pants (220858, -2.85 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -4.04 DPS) [crafted]; Knight's Restored Leather Leggings (220882, -6.31 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) | Captain Dirgehammer [vendor] | 61.5 | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Skycaller's Leather Boots (252471, -4.39 DPS) [crafted]; Skycaller's Leather Shoes (252532, -4.96 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.1 | yes | Band of the Unicorn (7553, -3.45 DPS) [world_drop]; Lorekeeper's Ring (19523, -3.56 DPS) [rep]; Lorekeeper's Ring (19524, -3.91 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 | yes | Lorekeeper's Ring (19523, -0.28 DPS) [rep]; Lorekeeper's Ring (19524, -0.62 DPS) [rep]; Band of the Unicorn (7553, -0.76 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (94.9 DPS) | yes | Uther's Strength (11302, -3.76 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -4.44 DPS) [quest]; Tidal Charm (1404, -4.44 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (95.2 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -0.46 DPS, sim-verified) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (101.1 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Blight (7959, -2.96 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; shoulder: Knight-Lieutenant's Crackling Leather Spaulders; back: Spritecaster Cape; chest: Knight's Crackling Leather Tunic; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Highlander's Cloth Girdle; legs: Knight's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain

No-known-source sample (15 of 531, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 191.3. Weights run: 1.9s. Verify run: 1.4s. 1101 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.526 ± 0.100, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 175.7 | yes | Bloodvine Goggles (19999, -3.32 DPS) [crafted]; Waywatcher Headpiece (240088, -6.57 DPS) [vendor]; Mask of the Unforgiven (13404, -12.08 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (192.4 DPS) | yes | Blazefury Medallion (17111, -0.56 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -2.22 DPS) [quest]; Amulet of the Dawn (22657, -3.56 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 196.4 | yes | Knight-Lieutenant's Leather Shoulders (220852, -11.05 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -12.68 DPS) [quest]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -13.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 76.2 | yes | Howler's Furs (272414, -2.32 DPS) [vendor]; Stalwart Cloak (272415, -2.32 DPS) [vendor]; Earthweave Cloak (21187, -4.30 DPS, sim-verified) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | 173.0 | yes | Bloodvine Vest (19682, -4.97 DPS, sim-verified) [crafted]; Waywatcher Tunic (240091, -5.82 DPS) [vendor]; Waywatcher Vest (240067, -8.17 DPS) [vendor] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | 99.3 | yes | Waywatcher Wristguards (240084, -4.10 DPS) [vendor]; Rockfury Bracers (21186, -4.64 DPS, sim-verified) [quest]; Primal Batskin Bracers (19687, -4.87 DPS) [crafted] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 179.4 | yes | Stormshroud Gloves (21278, -9.84 DPS) [crafted]; Primal Batskin Gloves (19686, -11.60 DPS, sim-verified) [crafted]; Waywatcher Grips (240065, -12.77 DPS) [vendor] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 168.8 | yes | Knowledge of the Timbermaw (228190, -3.27 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -10.80 DPS) [vendor]; Belt of the Archmage (18405, -11.47 DPS) [crafted] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 217.5 | yes | Waywatcher Legguards (240087, -0.01 DPS, sim-verified) [vendor]; Knight's Crackling Leather Leggings (220864, -12.03 DPS) [vendor]; Sentinel's Silk Leggings (237815, -12.04 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 167.4 | yes | Bloodvine Boots (19684, -1.85 DPS, sim-verified) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -10.50 DPS) [vendor]; Fine Dawn Treaders (227815, -12.40 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (189.2 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor]; Wrath of Cenarius (21190, -0.88 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (189.2 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor]; Wrath of Cenarius (21190, -0.68 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (188.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (189.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.88 DPS) [world_drop]; Weakness Analyzer (272438, -0.99 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (192.4 DPS) | yes | Frenzied Striker (13056, +0.00 DPS) [world_drop]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Enchanted Battlehammer (12776, -2.07 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor]; Enraged Idol (272428, +0.00 DPS) [vendor]; Idol of Synthesis (272429, +0.00 DPS) [vendor] |

**New at 60:** head: Waywatcher Hood; neck: Beads of Ogre Might; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Serenity Field; ranged: Idol of the Moon

No-known-source sample (15 of 1101, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 31.5. Weights run: 1.6s. Verify run: 1.1s. 164 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.323 ± 0.021, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (30.2 DPS) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 | yes | Double-Stitched Woolen Shoulders (4314, -0.46 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.52 DPS, sim-verified) [crafted]; Forest Leather Mantle (4709, -0.87 DPS) [world_drop] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 | yes | Totemic Leather Armor (252435, -0.27 DPS, sim-verified) [crafted]; Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.3 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Owl Bracers (4796, -0.04 DPS) [vendor]; Featherbead Bracers (15452, -0.04 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (30.6 DPS) | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.86 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (30.3 DPS) | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.51 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.11 DPS) [quest]; Volcanic Rock Ring (12053, -0.19 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | Westfall: Rhahk'Zor [dungeon] | 8.0 | yes | Twisted Chanter's Staff (890, -0.04 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.42 DPS) [quest]; Channeler's Staff (4437, -0.50 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Rhahk'Zor's Hammer; ranged: Idol of the Huntress

No-known-source sample (15 of 164, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 263006 Scout Ranger's Tunic; 263007 Skyseer's Vest; 263432 Skulker's Shiv

### Band 30 (tauren, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 64.1. Weights run: 1.9s. Verify run: 1.3s. 291 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.753 ± 0.028, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.16 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 | yes | Death Speaker Mantle (6685, -0.19 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Battle Healer's Cloak (19529, -0.13 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Guardian Armor (4256, -0.61 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Glowing Magical Bracelets (13106, -0.59 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.86 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.86 DPS) [quest] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | sim-verified (64.1 DPS) | yes | Jutebraid Gloves (10654, -0.23 DPS) [quest]; Serpent Gloves (5970, -0.38 DPS) [dungeon]; Fletcher's Gloves (7348, -1.47 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Warsong Sash (16975, -0.21 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.71 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Dark Ritual Leggings (270031, -0.29 DPS, sim-verified) [quest]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -1.97 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Sludge-Stained Band (286535, -0.54 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.40 DPS) [world]; Black Widow Band (6199, -0.40 DPS) [world]; Electrocutioner Lagnut (9447, -1.05 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (54.1 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -1.45 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (63.1 DPS) | yes | Twisted Chanter's Staff (890, +0.00 DPS) [world_drop]; Glimmering Staff (249392, +0.00 DPS) [crafted]; Rhahk'Zor's Hammer (5187, -8.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 291, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 84.9. Weights run: 1.8s. Verify run: 1.3s. 405 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.733 ± 0.044, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, +0.00 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -0.54 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.98 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -0.30 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.48 DPS) [crafted]; Elemental Raiment (9434, -0.67 DPS) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 | yes | Radiant Silver Bracers (4545, -0.13 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted]; Dreamweave Gloves (10019, -1.06 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 | yes | Defiler's Cloth Girdle (20166, -0.42 DPS, sim-verified) [rep]; Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Kodohide Legguards (285338, -0.43 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -0.74 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 | yes | Reedknot Ring (9622, -0.92 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.05 DPS) [vendor]; Ogremind Ring (1993, -1.19 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Reedknot Ring (9622, -0.93 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (78.0 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Rune of Duty (21567, -1.13 DPS, sim-verified) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (84.2 DPS) | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Windweaver Staff (7757, +0.00 DPS) [dungeon]; Illusionary Rod (7713, -6.30 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life

No-known-source sample (15 of 405, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 107.5. Weights run: 1.9s. Verify run: 1.5s. 537 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.310 ± 0.074, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (103.7 DPS) | yes | Soothsayer's Headdress (17740, -0.82 DPS) [dungeon]; Red Mageweave Headband (10033, -0.86 DPS) [crafted]; Blood Guard's Leather Headband (220851, -2.03 DPS, sim-verified) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world_drop]; Darkspear Warding Pendant (272073, -0.55 DPS) [vendor] |
| shoulder | Blood Guard's Crackling Leather Spaulders (220871) | Lady Palanseer [vendor] | sim-verified (104.5 DPS) | yes | Blood Guard's Leather Shoulders (220853, -2.83 DPS, sim-verified) [vendor]; Ironfeather Shoulders (15067, -3.75 DPS) [crafted]; Rotgrip Mantle (17732, -4.15 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.6 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.43 DPS) [crafted]; Big Voodoo Cloak (8216, -0.80 DPS) [crafted] |
| chest | Stone Guard's Crackling Leather Tunic (220869) | Lady Palanseer [vendor] | sim-verified (103.5 DPS) | yes | Acumen Robes (17775, -1.61 DPS) [quest]; Stone Guard's Leather Armor (220855, -1.83 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -2.45 DPS) [world_drop] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 28.6 | yes | Fletcher's Gloves (7348, -0.16 DPS) [crafted]; Shadowskin Gloves (18238, -0.16 DPS) [crafted]; Gloves of Holy Might (867, -4.10 DPS, sim-verified) [world_drop] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 39.9 | yes | Defiler's Lizardhide Girdle (20174, -0.30 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20193, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | Lady Palanseer [vendor] | 89.7 | yes | Stone Guard's Leather Pants (220859, -2.30 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -4.04 DPS) [crafted]; Stone Guard's Restored Leather Leggings (220883, -6.31 DPS) [vendor] |
| feet | First Sergeant's Crackling Leather Boots (220863) | Lady Palanseer [vendor] | 61.5 | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Skycaller's Leather Boots (252471, -4.39 DPS) [crafted]; Skycaller's Leather Shoes (252532, -4.96 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.1 | yes | Band of the Unicorn (7553, -3.45 DPS) [world_drop]; Advisor's Ring (19519, -3.56 DPS) [rep]; Advisor's Ring (19520, -3.91 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 | yes | Band of the Unicorn (7553, -0.04 DPS, sim-verified) [world_drop]; Advisor's Ring (19519, -0.28 DPS) [rep]; Advisor's Ring (19520, -0.62 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (95.0 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.61 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (95.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.69 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -3.46 DPS) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | sim-verified (101.5 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Blight (7959, -3.43 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; shoulder: Blood Guard's Crackling Leather Spaulders; back: Deep Woodlands Cloak; chest: Stone Guard's Crackling Leather Tunic; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Defiler's Cloth Girdle; legs: Stone Guard's Crackling Leather Leggings; feet: First Sergeant's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Rune of the Guard Captain

No-known-source sample (15 of 537, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 186.6. Weights run: 1.9s. Verify run: 1.3s. 1107 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.526 ± 0.100, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 175.7 | yes | Bloodvine Goggles (19999, -3.32 DPS) [crafted]; Waywatcher Headpiece (240088, -6.57 DPS) [vendor]; Mask of the Unforgiven (13404, -16.01 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (186.5 DPS) | yes | Medallion of the Dawn (22659, -2.22 DPS) [quest]; Blazefury Medallion (17111, -2.28 DPS, sim-verified) [world]; Amulet of the Dawn (22657, -3.56 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 196.4 | yes | Blood Guard's Leather Shoulders (220853, -10.90 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -12.68 DPS) [quest]; Blood Guard's Crackling Leather Spaulders (220871, -13.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 76.2 | yes | Howler's Furs (272414, -2.32 DPS) [vendor]; Stalwart Cloak (272415, -2.32 DPS) [vendor]; Earthweave Cloak (21187, -2.69 DPS, sim-verified) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | 173.0 | yes | Waywatcher Tunic (240091, -5.82 DPS) [vendor]; Waywatcher Vest (240067, -8.17 DPS) [vendor]; Bloodvine Vest (19682, -8.51 DPS, sim-verified) [crafted] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | 99.3 | yes | Rockfury Bracers (21186, -3.34 DPS, sim-verified) [quest]; Waywatcher Wristguards (240084, -4.10 DPS) [vendor]; Primal Batskin Bracers (19687, -4.87 DPS) [crafted] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 179.4 | yes | Stormshroud Gloves (21278, -9.84 DPS) [crafted]; Primal Batskin Gloves (19686, -12.48 DPS, sim-verified) [crafted]; Waywatcher Grips (240065, -12.77 DPS) [vendor] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 168.8 | yes | Knowledge of the Timbermaw (228190, -2.91 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -10.80 DPS) [vendor]; Belt of the Archmage (18405, -11.47 DPS) [crafted] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 217.5 | yes | Waywatcher Legguards (240087, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -12.03 DPS) [vendor]; Sentinel's Silk Leggings (237815, -12.04 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 167.4 | yes | Bloodvine Boots (19684, -3.10 DPS, sim-verified) [crafted]; First Sergeant's Crackling Leather Boots (220863, -10.50 DPS) [vendor]; Fine Dawn Treaders (227815, -12.40 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (186.5 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor]; Wrath of Cenarius (21190, -1.58 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (186.5 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor]; Wrath of Cenarius (21190, -1.46 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Blue Dragon (19288) | Darkmoon Beast Deck [quest] | sim-verified (183.1 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (186.5 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -1.02 DPS, sim-verified) [vendor] |
| main_hand | Enchanted Battlehammer (12776) | Blacksmithing [crafted] | sim-verified (186.5 DPS) | yes | Frenzied Striker (13056, +0.00 DPS) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Manual Crowd Pummeler (9449, -0.91 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor]; Enraged Idol (272428, +0.00 DPS) [vendor]; Idol of Synthesis (272429, +0.00 DPS) [vendor] |

**New at 60:** head: Waywatcher Hood; neck: Beads of Ogre Might; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Blue Dragon; trinket2: Serenity Field; main_hand: Enchanted Battlehammer; ranged: Idol of the Moon

No-known-source sample (15 of 1107, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

