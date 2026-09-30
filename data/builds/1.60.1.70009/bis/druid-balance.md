# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 32.4. Weights run: 1.7s. Verify run: 0.9s. 285 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.463 ± 0.112, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | 0.0 | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.38 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.5 | yes | Forest Leather Mantle (4709, -0.67 DPS) [world_drop]; Rugged Spaulders (5254, -0.67 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.84 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -0.79 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.04 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.40 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 0.0 | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.89 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 0.0 | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.52 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 | yes | Abomination Skin Leggings (23173, -0.01 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Loop of Sacrifice (281673, -0.39 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest]; Sludge-Stained Band (286535, -0.32 DPS, sim-verified) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.0 | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Staff of Westfall (2042, -0.14 DPS) [quest]; Living Root (6631, -2.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Huntress (227444) (or Idol of the Raging Shambler (220915), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Rune Broker [vendor] | 0.0 | yes | Idol of the Wild (210534, +0.00 DPS) [vendor]; Idol of the Raging Shambler (220915, +0.00 DPS, sim-verified) [vendor]; Mystic Mushroom (249396, +0.00 DPS) [crafted] |

**New at 20:** head: Wisdom's Leather Hood; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Idol of the Huntress

No-known-source sample (15 of 285, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers

### Band 30 (night-elf, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 64.4. Weights run: 1.8s. Verify run: 1.3s. 589 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.971 ± 0.152, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.25 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.28 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 | yes | Death Speaker Mantle (6685, -0.31 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Repairman's Cape (9605, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Prelacy Cape (7004, -0.13 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Guardian Armor (4256, -0.60 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.86 DPS) [quest]; Technician's Bracers (270042, -0.86 DPS) [quest]; Nightsky Wristbands (6407, -1.14 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 0.0 | yes | Serpent Gloves (5970, -0.23 DPS) [dungeon]; Shilly Mitts (9609, -0.23 DPS) [quest]; Fletcher's Gloves (7348, -1.12 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Belt of Arugal (6392, -0.31 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.63 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon]; Dark Ritual Leggings (270031, -1.35 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -1.92 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.15 DPS) [quest]; Lorekeeper's Ring (20431, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.40 DPS) [dungeon]; Sludge-Stained Band (286535, -0.40 DPS) [world]; Minor Channeling Ring (1449, -1.01 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 216.7 | yes | Gnarled Ash Staff (791, -3.38 DPS) [world_drop]; Glimmering Staff (249392, -3.82 DPS) [crafted]; Cobalt Crusher (7730, -7.44 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Huntress (227444) (or Idol of the Raging Shambler (220915), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Rune Broker [vendor] | 0.0 | yes | Idol of the Wild (210534, +0.00 DPS) [vendor]; Idol of the Heckler (213594, +0.00 DPS) [world]; Idol of the Raging Shambler (220915, +0.00 DPS, sim-verified) [vendor] |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 589, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (night-elf, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 85.2. Weights run: 1.9s. Verify run: 1.3s. 836 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.985 ± 0.232, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, +0.00 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -0.13 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.98 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -0.47 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.48 DPS) [crafted]; Crimson Silk Vest (7058, -0.92 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 | yes | Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest]; Arcane Runed Bracers (4744, -0.79 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted]; Dreamweave Gloves (10019, -1.24 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 | yes | Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.01 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Kodohide Legguards (285338, -0.75 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -1.11 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 | yes | Ring of Forlorn Spirits (2043, -0.79 DPS) [quest]; Reedknot Ring (9622, -0.92 DPS) [quest]; Minor Channeling Ring (1449, -1.00 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.05 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Staff of Jordan (873, -4.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Huntress (227444) (or Idol of the Raging Shambler (220915), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Rune Broker [vendor] | 0.0 | yes | Idol of the Wild (210534, +0.00 DPS) [vendor]; Idol of the Heckler (213594, +0.00 DPS) [world]; Idol of the Raging Shambler (220915, +0.00 DPS, sim-verified) [vendor] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection

No-known-source sample (15 of 836, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (night-elf, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 96.8. Weights run: 1.9s. Verify run: 1.2s. 1087 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.782 ± 0.380, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.87 DPS) [dungeon]; Soothsayer's Headdress (17740, -4.69 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world]; Darkspear Warding Pendant (272073, -0.55 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.90 DPS) [vendor]; Blood Guard's Crackling Leather Spaulders (220871, -0.90 DPS) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.4 | yes | Big Voodoo Cloak (8216, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.93 DPS) [vendor]; Runecloth Cloak (13860, -1.32 DPS, sim-verified) [crafted] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 75.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Knight's Crackling Leather Tunic (220868, -3.12 DPS) [vendor]; Stone Guard's Crackling Leather Tunic (220869, -3.12 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Shadowskin Gloves (18238), Sergeant Major's Leather Gauntlets (220856), First Sergeant's Leather Gauntlets (220857)) | World drop [world_drop] | 27.2 | yes | Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Fletcher's Gloves (7348, -0.58 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 39.9 | yes | Highlander's Lizardhide Girdle (20103, -1.07 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20115, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) (or Stone Guard's Crackling Leather Leggings (220865)) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Knight's Leather Pants (220858, -2.22 DPS) [vendor]; Stone Guard's Leather Pants (220859, -2.22 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) (or First Sergeant's Crackling Leather Boots (220863)) | Captain Dirgehammer [vendor] | 66.2 | yes | First Sergeant's Crackling Leather Boots (220863, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.83 DPS) [crafted]; Skycaller's Leather Boots (252471, -4.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.8 | yes | Lorekeeper's Ring (19523, -4.10 DPS) [rep]; Lorekeeper's Ring (19524, -4.45 DPS) [rep]; Ring of Forlorn Spirits (2043, -4.56 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 | yes | Lorekeeper's Ring (19523, +0.00 DPS, sim-verified) [rep]; Lorekeeper's Ring (19524, -0.62 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.74 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.37 DPS, sim-verified) [world] |
| main_hand | Thorium Greatmace (250613) | Blacksmithing [crafted] | 0.0 | yes | Kindling Stave (11750, -0.00 DPS) [dungeon]; Soulkeeper (1607, -2.49 DPS) [world_drop]; Dark Iron Pulverizer (11608, -2.56 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Huntress (227444) (or Idol of the Raging Shambler (220915), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Rune Broker [vendor] | 0.0 | yes | Idol of the Wild (210534, +0.00 DPS) [vendor]; Idol of the Heckler (213594, +0.00 DPS) [world]; Idol of the Raging Shambler (220915, +0.00 DPS, sim-verified) [vendor] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Spritecaster Cape; chest: Knight's Leather Armor; wrist: Runic Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Cloth Girdle; legs: Knight's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket2: Thunderbrew's Boot Flask; main_hand: Thorium Greatmace

No-known-source sample (15 of 1087, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (night-elf, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 139.5. Weights run: 1.9s. Verify run: 1.4s. 1811 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.891 ± 0.496, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 0.0 | yes | Mask of the Unforgiven (13404, -3.04 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Leather Headband (220850, -6.51 DPS) [vendor]; Blood Guard's Leather Headband (220851, -6.51 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 0.0 | yes | Beads of Ogre Might (22150, -3.89 DPS) [quest]; Medallion of the Dawn (22659, -6.51 DPS) [quest]; Charm of the Shifting Sands (21504, -6.82 DPS) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 94.1 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -0.97 DPS) [quest]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.88 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 58.9 | yes | Chromatic Cloak (18509, -0.92 DPS, sim-verified) [crafted]; Hide of the Wild (18510, -4.28 DPS) [crafted]; Spritecaster Cape (11623, -4.55 DPS) [dungeon] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 152.9 | yes | Knight's Leather Armor (220854, -5.29 DPS, sim-verified) [vendor]; Knight-Captain's Dragonhide Chestpiece (227176, -5.61 DPS) [vendor]; Legionnaire's Dragonhide Chestpiece (227179, -5.61 DPS) [vendor] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 85.9 | yes | Primal Batskin Bracers (19687, -3.49 DPS, sim-verified) [crafted]; Dryad's Wrist Bindings (19595, -6.52 DPS) [rep]; Dryad's Wrist Bindings (19596, -6.88 DPS) [rep] |
| hands | Primal Batskin Gloves (19686) | Leatherworking [crafted] | 117.8 | yes | Stormshroud Gloves (21278, -0.69 DPS, sim-verified) [crafted]; Dreadmist Wraps (16705, -5.90 DPS) [dungeon]; Blood Guard's Dragonhide Grips (227180, -6.51 DPS) [vendor] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 65.1 | yes | Highlander's Cloth Girdle (20047, -1.35 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.97 DPS) [rep]; Belt of Preserved Heads (20216, -4.28 DPS, sim-verified) [quest] |
| legs | Knight's Crackling Leather Leggings (220864) (or Stone Guard's Crackling Leather Leggings (220865)) | Captain Dirgehammer [vendor] | 112.3 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Dragonhide Leggings (227177, -1.19 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.19 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 87.8 | yes | First Sergeant's Crackling Leather Boots (220863, -1.30 DPS) [vendor]; Sergeant Major's Crackling Leather Boots (220862, -1.76 DPS, sim-verified) [vendor]; Blood Guard's Dragonhide Treads (227181, -5.55 DPS) [vendor] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 99.6 | yes | Band of Earthen Might (21182, -0.61 DPS) [quest]; Mindtear Band (20632, -4.28 DPS) [world]; Ritssyn's Ring of Chaos (21836, -4.36 DPS) [world] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 94.1 | yes | Band of Earthen Might (21182, -0.12 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.67 DPS) [world]; Ritssyn's Ring of Chaos (21836, -3.75 DPS) [world] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.41 DPS, sim-verified) [world] |
| main_hand | Fist of Cenarius (21188) | Champion's Battlegear [quest] | 0.0 | yes | High Warlord's Destroyer (234546, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Mushgog [world] | 0.0 | yes | Idol of the Heckler (213594, +0.00 DPS) [world]; Idol of the Raging Shambler (220915, +0.00 DPS) [vendor]; Idol of the Huntress (227444, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Bloodvine Goggles; neck: Onyxia Tooth Pendant; back: Earthweave Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Primal Batskin Gloves; waist: Belt of the Archmage; feet: Bloodvine Boots; finger1: Ring of the Fallen God; finger2: Don Julio's Band; trinket1: Thunderbrew's Boot Flask; trinket2: Ankh of Life; main_hand: Fist of Cenarius; ranged: Idol of the Moon

No-known-source sample (15 of 1811, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 30.4. Weights run: 1.7s. Verify run: 0.9s. 290 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.463 ± 0.112, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | 0.0 | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.37 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.5 | yes | Double-Stitched Woolen Shoulders (4314, -0.57 DPS, sim-verified) [crafted]; Forest Leather Mantle (4709, -0.67 DPS) [world_drop]; Rugged Spaulders (5254, -0.67 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.12 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -0.41 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.3 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Owl Bracers (4796, -0.04 DPS) [vendor]; Featherbead Bracers (15452, -0.04 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 0.0 | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.87 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 0.0 | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.50 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 | yes | Abomination Skin Leggings (23173, -0.00 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 | yes | Stormrider's Leather Boots (252443, -0.05 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest]; Black Pearl Ring (6332, -0.43 DPS) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.11 DPS) [quest]; Black Pearl Ring (6332, -0.23 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.0 | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Hammerbone (270018, -0.40 DPS) [quest]; Living Root (6631, -1.38 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Huntress (227444) (or Idol of the Raging Shambler (220915), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Rune Broker [vendor] | 0.0 | yes | Idol of the Wild (210534, +0.00 DPS) [vendor]; Idol of the Raging Shambler (220915, +0.00 DPS, sim-verified) [vendor]; Mystic Mushroom (249396, +0.00 DPS) [crafted] |

**New at 20:** head: Wisdom's Leather Hood; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Idol of the Huntress

No-known-source sample (15 of 290, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers

### Band 30 (tauren, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 64.1. Weights run: 1.8s. Verify run: 1.3s. 599 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.971 ± 0.152, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.16 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.28 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 | yes | Death Speaker Mantle (6685, -0.19 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Battle Healer's Cloak (19529, -0.13 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Guardian Armor (4256, -0.61 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.68 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.86 DPS) [quest]; Technician's Bracers (270042, -0.86 DPS) [quest] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | 0.0 | yes | Jutebraid Gloves (10654, -0.23 DPS) [quest]; Serpent Gloves (5970, -0.38 DPS) [dungeon]; Fletcher's Gloves (7348, -1.47 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Warsong Sash (16975, -0.21 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.71 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Dark Ritual Leggings (270031, -0.29 DPS, sim-verified) [quest]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -1.97 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Sludge-Stained Band (286535, -0.54 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.40 DPS) [world]; Black Widow Band (6199, -0.40 DPS) [world]; Electrocutioner Lagnut (9447, -1.05 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 216.7 | yes | Gnarled Ash Staff (791, -3.38 DPS) [world_drop]; Glimmering Staff (249392, -3.82 DPS) [crafted]; Cobalt Crusher (7730, -7.33 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Huntress (227444) (or Idol of the Raging Shambler (220915), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Rune Broker [vendor] | 0.0 | yes | Idol of the Wild (210534, +0.00 DPS) [vendor]; Idol of the Heckler (213594, +0.00 DPS) [world]; Idol of the Raging Shambler (220915, +0.00 DPS, sim-verified) [vendor] |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Manual Crowd Pummeler

No-known-source sample (15 of 599, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (tauren, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 84.9. Weights run: 1.9s. Verify run: 1.4s. 845 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.985 ± 0.232, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, +0.00 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Necklace of Calisea (1714, -0.54 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.80 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.98 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest]; Green Silken Shoulders (7057, -0.30 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.48 DPS) [crafted]; Crimson Silk Vest (7058, -0.92 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 | yes | Radiant Silver Bracers (4545, -0.13 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted]; Dreamweave Gloves (10019, -1.06 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 | yes | Defiler's Cloth Girdle (20166, -0.42 DPS, sim-verified) [rep]; Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Kodohide Legguards (285338, -0.43 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -0.74 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 | yes | Reedknot Ring (9622, -0.92 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.05 DPS) [vendor]; Ogremind Ring (1993, -1.19 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Reedknot Ring (9622, -0.93 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Mograine's Might (7723, +0.00 DPS) [dungeon]; Primitive Fishing Pole (276203, +0.00 DPS) [vendor]; Staff of Jordan (873, -4.64 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Huntress (227444) (or Idol of the Raging Shambler (220915), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Rune Broker [vendor] | 0.0 | yes | Idol of the Wild (210534, +0.00 DPS) [vendor]; Idol of the Heckler (213594, +0.00 DPS) [world]; Idol of the Raging Shambler (220915, +0.00 DPS, sim-verified) [vendor] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection

No-known-source sample (15 of 845, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (tauren, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 97.5. Weights run: 1.9s. Verify run: 1.4s. 1096 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.782 ± 0.380, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.87 DPS) [dungeon]; Soothsayer's Headdress (17740, -4.69 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world]; Darkspear Warding Pendant (272073, -0.55 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 75.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -0.90 DPS) [vendor]; Blood Guard's Crackling Leather Spaulders (220871, -0.90 DPS) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 18.6 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.43 DPS) [crafted]; Big Voodoo Cloak (8216, -0.80 DPS) [crafted] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 75.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Knight's Crackling Leather Tunic (220868, -3.12 DPS) [vendor]; Stone Guard's Crackling Leather Tunic (220869, -3.12 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Shadowskin Gloves (18238), Sergeant Major's Leather Gauntlets (220856), First Sergeant's Leather Gauntlets (220857)) | World drop [world_drop] | 27.2 | yes | Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Sergeant Major's Leather Gauntlets (220856, +0.00 DPS) [vendor]; Fletcher's Gloves (7348, -0.95 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 39.9 | yes | Defiler's Lizardhide Girdle (20174, -1.44 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20193, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) (or Stone Guard's Crackling Leather Leggings (220865)) | Captain Dirgehammer [vendor] | 94.4 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Knight's Leather Pants (220858, -2.22 DPS) [vendor]; Stone Guard's Leather Pants (220859, -2.22 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) (or First Sergeant's Crackling Leather Boots (220863)) | Captain Dirgehammer [vendor] | 66.2 | yes | First Sergeant's Crackling Leather Boots (220863, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.83 DPS) [crafted]; Skycaller's Leather Boots (252471, -4.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.8 | yes | Advisor's Ring (19519, -4.10 DPS) [rep]; Advisor's Ring (19520, -4.45 DPS) [rep]; Reedknot Ring (9622, -4.67 DPS) [quest] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 | yes | Advisor's Ring (19519, +0.00 DPS, sim-verified) [rep]; Advisor's Ring (19520, -0.62 DPS) [rep]; Reedknot Ring (9622, -0.85 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, -0.52 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Frozen Heart of the Mountain (249469, -1.73 DPS, sim-verified) [crafted]; Uther's Strength (11302, -3.15 DPS) [world]; Guardian Talisman (1490, -3.83 DPS) [quest] |
| main_hand | Manual Crowd Pummeler (9449) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 0.0 | yes | Soulkeeper (1607, +0.00 DPS) [world_drop]; Kindling Stave (11750, +0.00 DPS) [dungeon]; Thorium Greatmace (250613, -1.03 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Huntress (227444) (or Idol of the Raging Shambler (220915), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Rune Broker [vendor] | 0.0 | yes | Idol of the Wild (210534, +0.00 DPS) [vendor]; Idol of the Heckler (213594, +0.00 DPS) [world]; Idol of the Raging Shambler (220915, +0.00 DPS, sim-verified) [vendor] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Deep Woodlands Cloak; chest: Knight's Leather Armor; wrist: Runic Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Cloth Girdle; legs: Knight's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket2: Rune of the Guard Captain

No-known-source sample (15 of 1096, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (tauren, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 138.2. Weights run: 1.9s. Verify run: 1.4s. 1819 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.891 ± 0.496, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bloodvine Goggles (19999) | Engineering [crafted] | 0.0 | yes | Mask of the Unforgiven (13404, -2.18 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Leather Headband (220850, -6.51 DPS) [vendor]; Blood Guard's Leather Headband (220851, -6.51 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 0.0 | yes | Beads of Ogre Might (22150, -3.89 DPS) [quest]; Medallion of the Dawn (22659, -6.51 DPS) [quest]; Charm of the Shifting Sands (21504, -6.82 DPS) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 94.1 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -0.97 DPS) [quest]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -1.88 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 58.9 | yes | Chromatic Cloak (18509, -0.81 DPS, sim-verified) [crafted]; Hide of the Wild (18510, -4.28 DPS) [crafted]; Spritecaster Cape (11623, -4.55 DPS) [dungeon] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 152.9 | yes | Knight-Captain's Dragonhide Chestpiece (227176, -5.61 DPS) [vendor]; Legionnaire's Dragonhide Chestpiece (227179, -5.61 DPS) [vendor]; Knight's Leather Armor (220854, -6.43 DPS, sim-verified) [vendor] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 85.9 | yes | Primal Batskin Bracers (19687, -3.46 DPS, sim-verified) [crafted]; Dryad's Wrist Bindings (19595, -6.52 DPS) [rep]; Dryad's Wrist Bindings (19596, -6.88 DPS) [rep] |
| hands | Primal Batskin Gloves (19686) | Leatherworking [crafted] | 117.8 | yes | Stormshroud Gloves (21278, -0.55 DPS, sim-verified) [crafted]; Dreadmist Wraps (16705, -5.90 DPS) [dungeon]; Blood Guard's Dragonhide Grips (227180, -6.51 DPS) [vendor] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 65.1 | yes | Defiler's Cloth Girdle (20163, -1.35 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.97 DPS) [rep]; Belt of Preserved Heads (20216, -4.46 DPS, sim-verified) [quest] |
| legs | Knight's Crackling Leather Leggings (220864) (or Stone Guard's Crackling Leather Leggings (220865)) | Captain Dirgehammer [vendor] | 112.3 | yes | Stone Guard's Crackling Leather Leggings (220865, +0.00 DPS, sim-verified) [vendor]; Legionnaire's Dragonhide Leggings (227177, -1.19 DPS) [vendor]; Knight-Captain's Dragonhide Leggings (227178, -1.19 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 87.8 | yes | First Sergeant's Crackling Leather Boots (220863, -1.30 DPS) [vendor]; Sergeant Major's Crackling Leather Boots (220862, -2.37 DPS, sim-verified) [vendor]; Blood Guard's Dragonhide Treads (227181, -5.55 DPS) [vendor] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 99.6 | yes | Band of Earthen Might (21182, -0.61 DPS) [quest]; Mindtear Band (20632, -4.28 DPS) [world]; Ritssyn's Ring of Chaos (21836, -4.36 DPS) [world] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 94.1 | yes | Band of Earthen Might (21182, -0.13 DPS, sim-verified) [quest]; Mindtear Band (20632, -3.67 DPS) [world]; Ritssyn's Ring of Chaos (21836, -3.75 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.66 DPS, sim-verified) [world] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.13 DPS, sim-verified) [world]; Guardian Talisman (1490, -4.56 DPS) [quest] |
| main_hand | Fist of Cenarius (21188) | Champion's Battlegear [quest] | 0.0 | yes | High Warlord's Destroyer (234546, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | Mushgog [world] | 0.0 | yes | Idol of the Heckler (213594, +0.00 DPS) [world]; Idol of the Raging Shambler (220915, +0.00 DPS) [vendor]; Idol of the Huntress (227444, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Bloodvine Goggles; neck: Onyxia Tooth Pendant; back: Earthweave Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Primal Batskin Gloves; waist: Belt of the Archmage; feet: Bloodvine Boots; finger1: Ring of the Fallen God; finger2: Don Julio's Band; main_hand: Fist of Cenarius; ranged: Idol of the Moon

No-known-source sample (15 of 1819, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

