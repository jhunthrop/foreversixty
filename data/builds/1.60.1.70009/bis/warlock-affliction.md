# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 38.6. Weights run: 1.2s. Verify run: 1.0s. 253 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.127, intellect=0.422 ± -0.128, crit=-0.462 ± -0.020, hit=-1.259 ± -0.070, spell_haste=-0.173 ± -0.159, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.000 ± -0.127

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, +0.00 DPS) [crafted]; Lucky Fishing Hat (19972, +0.00 DPS) [quest]; Shadow Goggles (4373, -2.83 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.7 | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.75 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, +0.00 DPS) [dungeon]; Black Whelp Cloak (7283, +0.00 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.36 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.1 | yes | Green Woolen Vest (2582, +0.00 DPS) [crafted]; Green Woolen Robe (6243, +0.00 DPS) [crafted]; Gray Woolen Robe (2585, -1.21 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.5 | yes | Mindthrust Bracers (1974, +0.12 DPS, sim-verified) [dungeon]; Seer's Cuffs (3645, +0.00 DPS) [world_drop]; Bright Bracers (3647, +0.00 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Pristine Gloves (253913, +0.00 DPS) [crafted]; Blight Gloves (279877, +0.00 DPS) [quest]; Gnoll Casting Gloves (892, -0.26 DPS, sim-verified) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.7 | yes | Keller's Girdle (2911, +0.00 DPS) [world_drop]; Novice Ardent's Sash (253887, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.81 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.5 | yes | Silk-threaded Trousers (1929, +0.00 DPS) [dungeon]; Colorful Kilt (10048, +0.00 DPS) [crafted]; Abomination Skin Leggings (23173, -0.51 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 | yes | Red Woolen Boots (4313, +0.00 DPS) [crafted]; Pristine Boots (253889, +0.00 DPS) [crafted]; Feather Padded Treads (285345, -0.76 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Black Pearl Ring (6332, +0.00 DPS) [world]; Sludge-Stained Band (286535, +0.00 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Black Pearl Ring (6332, +0.00 DPS) [world]; Sludge-Stained Band (286535, -0.39 DPS, sim-verified) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 4.2 | yes | Lesser Staff of the Spire (1300, +0.00 DPS) [world]; Channeler's Staff (4437, +0.00 DPS) [world]; Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | The Deadmines: Cookie [dungeon] | 3.0 | yes | Torchlight Wand (5240, +0.00 DPS) [quest]; Sable Wand (7607, -0.00 DPS) [quest]; Sizzle Stick (8071, -0.46 DPS, sim-verified) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 253, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 60.5. Weights run: 1.2s. Verify run: 1.2s. 476 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.261), intellect=not significant (-0.193 ± 0.270), crit=0.624 ± 0.029, hit=1.712 ± 0.125, spell_haste=-1.617 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.261)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Embalmed Shroud (7691, -0.53 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.53 DPS) [crafted]; Silk Headband (7050, -0.80 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Glowing Green Talisman (5002, -1.23 DPS) [world_drop]; Crystal Starfire Medallion (5003, -1.23 DPS) [world_drop]; Pendant of Myzrael (4614, -2.58 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 | yes | Invoker's Mantle (215365, -0.35 DPS) [crafted]; Death Speaker Mantle (6685, -0.53 DPS) [dungeon]; Moonlit Amice (11884, -1.50 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Prelacy Cape (7004, -0.18 DPS) [quest]; Caretaker's Cape (19533, -0.18 DPS) [rep]; Heavy Woolen Cloak (4311, -0.48 DPS, sim-verified) [crafted] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, -0.77 DPS, sim-verified) [crafted]; Robes of Arcana (5770, -0.88 DPS) [crafted]; Death Speaker Robes (6682, -1.05 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Silver-lined Bracers (3224, -1.58 DPS) [world]; Seer's Cuffs (3645, -1.58 DPS) [world_drop]; Mindthrust Bracers (1974, -2.47 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) (or Shilly Mitts (9609)) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Shilly Mitts (9609, -0.03 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.18 DPS) [world]; Gloves of Meditation (4318, -0.35 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.0 | yes | Ghamoo-ra's Bind (6908, -0.53 DPS) [dungeon]; Invoker's Cord (215366, -0.70 DPS) [crafted]; Belt of Arugal (6392, -0.78 DPS, sim-verified) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -0.46 DPS, sim-verified) [dungeon]; Silk-threaded Trousers (1929, -0.88 DPS) [dungeon]; Pristine Leggings (253987, -0.88 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.0 | yes | Nimbus Boots (6998, -0.18 DPS) [quest]; Boots of the Enchanter (4325, -0.35 DPS) [crafted]; Spidersilk Boots (4320, -1.57 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.35 DPS) [quest]; Lorekeeper's Ring (20431, -0.35 DPS) [rep]; Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.53 DPS) [dungeon]; Sludge-Stained Band (286535, -0.53 DPS) [world]; Minor Channeling Ring (1449, -2.01 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | - | - |  |  |  |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.0 | yes | Eye of Paleth (2943, -0.53 DPS) [quest]; Orb of Souls (249395, -0.53 DPS) [crafted]; Dwarven Tome (279898, -1.34 DPS, sim-verified) [quest] |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | World drop [world_drop] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Spellcrafter Wand (6677, -0.18 DPS) [quest]; Sizzle Stick (8071, -1.34 DPS, sim-verified) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Notched Shortsword; off_hand: Orb of Mystic Insight; ranged: Wand of Eventide

No-known-source sample (15 of 476, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 108.5. Weights run: 1.2s. Verify run: 1.1s. 643 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.358), intellect=not significant (0.698 ± 0.495), crit=1.096 ± 0.054, hit=3.563 ± 0.219, spell_haste=-2.263 ± 0.479, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.358)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -0.89 DPS) [crafted]; Holy Shroud (2721, -1.11 DPS) [world_drop]; Augural Shroud (2620, -2.51 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Darkspear Warding Pendant (272074, -0.70 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.86 DPS) [vendor]; Necklace of Calisea (1714, -2.37 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Green Silken Shoulders (7057, +0.04 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.20 DPS) [vendor]; Icy Cloak (4327, -0.28 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 | yes | Robe of Power (7054, -0.42 DPS) [crafted]; Crimson Silk Vest (7058, -0.80 DPS) [crafted]; Dreamweave Vest (10021, -1.20 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, -0.09 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.22 DPS) [quest]; Aurora Bracers (4043, -0.38 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 | yes | Black Mageweave Gloves (10003, -0.64 DPS) [crafted]; Gilded Handwraps (254021, -0.88 DPS) [crafted]; Red Mageweave Gloves (10018, -2.15 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.8 | yes | Gilded Cord (254037, -0.36 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.41 DPS) [rep]; Deathmage Sash (10771, -1.28 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Abomination Skin Leggings (23173, -0.87 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.09 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.10 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -1.49 DPS) [dungeon]; Spidersilk Boots (4320, -1.58 DPS) [crafted]; Gilded Slippers (254001, -2.60 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.22 DPS) [quest]; Lorekeeper's Ring (19525, -0.22 DPS) [rep]; Minor Channeling Ring (1449, -0.29 DPS) [quest] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 | yes | Minor Channeling Ring (1449, -0.18 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.22 DPS) [vendor]; Reedknot Ring (9622, -0.42 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 43.3 | yes | Illusionary Rod (7713, -1.76 DPS, sim-verified) [dungeon]; Staff of Noh'Orahil (15105, -2.80 DPS) [quest]; Windweaver Staff (7757, -3.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fizzle's Zippy Lighter (6729) | News for Fizzle [quest] | 6.1 | yes | Burning Sliver (5249, +0.33 DPS, sim-verified) [quest]; Twisted Nether Wand (249144, -0.01 DPS) [crafted]; Wand of Eventide (5214, -0.12 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Dar'Orahil; ranged: Fizzle's Zippy Lighter

No-known-source sample (15 of 643, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 176.3. Weights run: 1.2s. Verify run: 1.2s. 804 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.222, intellect=not significant (0.052 ± 0.247), crit=0.455 ± 0.023, hit=2.181 ± 0.152, spell_haste=-9.055 ± 0.409, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Spellpower Goggles Xtreme (10502, -1.95 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.95 DPS) [vendor]; Dreamweave Circlet (10041, -2.61 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 | yes | Mindburst Medallion (11196, -0.57 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.14 DPS) [world]; Darkspear Warding Pendant (272073, -2.22 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 14.8 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -0.29 DPS) [dungeon]; Black Mageweave Shoulders (10027, -1.42 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 | yes | Nightfall Drape (12465, -1.72 DPS) [world]; Runecloth Cloak (13860, -2.13 DPS, sim-verified) [crafted]; Icy Cloak (4327, -2.37 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 | yes | Acumen Robes (17775, -0.84 DPS, sim-verified) [quest]; Knight's Dreadweave Vest (220886, -1.09 DPS) [vendor]; Stone Guard's Dreadweave Vest (220904, -1.09 DPS) [vendor] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.29 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.53 DPS) [crafted]; Condor Bracers (15864, -0.65 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 | yes | Black Mageweave Gloves (10003, -1.11 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.54 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.54 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 15.6 | yes | Satyrmane Sash (17755, +0.74 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.46 DPS) [rep]; Ghostweave Cord (254073, -0.53 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 | yes | Stone Guard's Dreadweave Leggings (220906, -0.00 DPS) [vendor]; Red Mageweave Pants (10009, -1.42 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.68 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 30.3 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -2.04 DPS) [crafted]; Gilded Sandals (254107, -6.11 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.8 | yes | Ring of Forlorn Spirits (2043, -4.48 DPS) [quest]; Reedknot Ring (9622, -4.81 DPS) [quest]; Sea Giant's Toe Ring (274746, -5.13 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Lorekeeper's Ring (19524, -0.83 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -1.30 DPS) [quest]; Reedknot Ring (9622, -1.62 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Frozen Heart of the Mountain (249469, -0.49 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -1.95 DPS) [quest]; Tidal Charm (1404, -1.95 DPS) [vendor] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | 12.0 | yes | Frozen Heart of the Mountain (249469, -2.97 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -3.90 DPS) [quest]; Tidal Charm (1404, -3.90 DPS) [vendor] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 22.6 | yes | Staff of Dar'Orahil (15106, +0.25 DPS, sim-verified) [quest]; Kindling Stave (11750, -5.08 DPS) [dungeon]; Illusionary Rod (7713, -5.18 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Dreambough Wand (249234) | Enchanting [crafted] | 7.0 | yes | Burning Sliver (5249, -0.32 DPS) [quest]; Twisted Nether Wand (249144, -0.32 DPS) [crafted]; Lesser Eternal Wand (249232, -2.56 DPS, sim-verified) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; waist: Highlander's Cloth Girdle; legs: Wizardweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Abyss Shard; main_hand: Soul Harvester; ranged: Dreambough Wand

No-known-source sample (15 of 804, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 361.9. Weights run: 1.3s. Verify run: 1.2s. 1239 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.837), intellect=not significant (-1.254 ± 0.970), crit=1.999 ± 0.106, hit=6.401 ± 0.516, spell_haste=not significant (0.241 ± 0.963), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.837)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | 153.0 | yes | Doomcaller's Circlet (21337, -4.06 DPS) [quest]; Deathmist Mask (226909, -10.29 DPS) [quest]; Bloodvine Goggles (19999, -16.13 DPS, sim-verified) [crafted] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 92.0 | yes | Beads of Ogre Might (22150, -4.06 DPS) [quest]; Medallion of the Dawn (22659, -9.28 DPS) [quest]; Charm of the Shifting Sands (21504, -9.72 DPS) [quest] |
| shoulder | Plagueheart Shoulderpads (22507) | Plagueheart Shoulderpads [quest] | 100.0 | yes | Doomcaller's Mantle (21335, -2.74 DPS, sim-verified) [quest]; Mantle of the Timbermaw (19050, -7.98 DPS) [crafted]; Champion's Dreadweave Spaulders (23256, -8.71 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 64.0 | yes | Chromatic Cloak (18509, -0.89 DPS, sim-verified) [crafted]; Shroud of Unspoken Names (21418, -6.67 DPS) [quest]; Spritecaster Cape (11623, -7.25 DPS) [dungeon] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 143.0 | yes | Zandalar Demoniac's Robe (20033, -7.54 DPS) [quest]; Doomcaller's Robes (21334, -10.73 DPS) [quest]; Bloodvine Vest (19682, -11.07 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 91.0 | yes | Plagueheart Bindings (22511, +0.19 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -10.01 DPS) [rep]; Dryad's Wrist Bindings (19596, -10.30 DPS) [rep] |
| hands | Deathmist Wraps (22077) (or Deathmist Wraps (226911)) | Just Compensation [quest] | 77.0 | yes | Deathmist Wraps (226911, +0.00 DPS, sim-verified) [quest]; Gloves of Spell Mastery (14146, -1.75 DPS) [crafted]; Dreadmist Wraps (16705, -1.89 DPS) [dungeon] |
| waist | Plagueheart Belt (22510) | Plagueheart Belt [quest] | 62.0 | yes | Highlander's Cloth Girdle (20047, -2.90 DPS) [rep]; Highlander's Cloth Girdle (20097, -3.63 DPS) [rep]; Belt of the Archmage (18405, -5.82 DPS, sim-verified) [crafted] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 101.0 | yes | Plagueheart Leggings (22505, -5.23 DPS) [quest]; Magister's Leggings (16687, -5.28 DPS) [dungeon]; Deathmist Leggings (226910, -5.82 DPS, sim-verified) [quest] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 83.0 | yes | First Sergeant's Dreadweave Boots (220909, -1.60 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -2.92 DPS, sim-verified) [vendor]; Plagueheart Sandals (22508, -3.34 DPS) [quest] |
| finger1 | Ring of Unspoken Names (21417) | Ring of Unspoken Names [quest] | 106.0 | yes | Don Julio's Band (19325, -2.03 DPS) [rep]; Band of Earthen Might (21182, -2.03 DPS) [quest]; Blackstone Ring (17713, -6.09 DPS) [dungeon] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 101.0 | yes | Band of Earthen Might (21182, -1.31 DPS) [quest]; Blackstone Ring (17713, -5.37 DPS) [dungeon]; Don Julio's Band (19325, -12.93 DPS, sim-verified) [rep] |
| trinket1 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Frozen Heart of the Mountain (249469, -0.38 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -0.87 DPS) [quest]; Tidal Charm (1404, -0.87 DPS) [vendor] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | 12.0 | yes | Thunderbrew's Boot Flask (744, -1.74 DPS) [quest]; Tidal Charm (1404, -1.74 DPS) [vendor]; Frozen Heart of the Mountain (249469, -3.01 DPS, sim-verified) [crafted] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 137.0 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Atiesh, Greatstaff of the Guardian (22631, -2.46 DPS) [quest]; Grand Marshal's Stave (18873, -9.57 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 9.0 | yes | Lesser Eternal Wand (249232, -0.15 DPS) [crafted]; Dreambough Wand (249234, -0.29 DPS) [crafted]; Greater Eternal Wand (249237, -6.52 DPS, sim-verified) [crafted] |

**New at 60:** head: Plagueheart Circlet; neck: Onyxia Tooth Pendant; shoulder: Plagueheart Shoulderpads; back: Earthweave Cloak; chest: Plagueheart Robe; wrist: Rockfury Bracers; hands: Deathmist Wraps; waist: Plagueheart Belt; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Ring of Unspoken Names; finger2: Ring of the Fallen God; main_hand: High Warlord's War Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1239, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.5. Weights run: 1.2s. Verify run: 1.1s. 250 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.127, intellect=0.422 ± -0.128, crit=-0.462 ± -0.020, hit=-1.259 ± -0.070, spell_haste=-0.173 ± -0.159, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.000 ± -0.127

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, +0.00 DPS) [crafted]; Lucky Fishing Hat (19972, +0.00 DPS) [quest]; Shadow Goggles (4373, -2.36 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.7 | yes | Slime-encrusted Pads (6461, +0.00 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.55 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, +0.00 DPS) [dungeon]; Black Whelp Cloak (7283, +0.00 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.41 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.1 | yes | Green Woolen Vest (2582, +0.00 DPS) [crafted]; Green Woolen Robe (6243, +0.00 DPS) [crafted]; Gray Woolen Robe (2585, -1.35 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.5 | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Owlbeard Bracers (16981, +0.00 DPS) [quest]; Mindthrust Bracers (1974, -0.07 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Apothecary Gloves (10919, +0.00 DPS) [quest]; Pristine Gloves (253913, +0.00 DPS) [crafted]; Gnoll Casting Gloves (892, -0.47 DPS, sim-verified) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.7 | yes | Keller's Girdle (2911, +0.00 DPS) [world_drop]; Novice Ardent's Sash (253887, +0.00 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.85 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 8.5 | yes | Silk-threaded Trousers (1929, +0.00 DPS) [dungeon]; Colorful Kilt (10048, +0.00 DPS) [crafted]; Abomination Skin Leggings (23173, -0.60 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.7 | yes | Red Woolen Boots (4313, +0.00 DPS) [crafted]; Pristine Boots (253889, +0.00 DPS) [crafted]; Feather Padded Treads (285345, -0.68 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS) [dungeon]; Black Pearl Ring (6332, +0.00 DPS) [world]; The 1 Ring (8350, +0.00 DPS) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Black Pearl Ring (6332, +0.00 DPS) [world]; The 1 Ring (8350, +0.00 DPS) [world]; Lavishly Jeweled Ring (1156, -0.37 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 4.2 | yes | Lesser Staff of the Spire (1300, +0.00 DPS) [world]; Channeler's Staff (4437, +0.00 DPS) [world]; Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Cookie's Stirring Rod (5198, +0.36 DPS, sim-verified) [dungeon]; Torchlight Wand (5240, +0.00 DPS) [quest]; Greater Magic Wand (11288, +0.00 DPS) [crafted] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 250, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 58.9. Weights run: 1.2s. Verify run: 1.2s. 474 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.261), intellect=not significant (-0.193 ± 0.270), crit=0.624 ± 0.029, hit=1.712 ± 0.125, spell_haste=-1.617 ± 0.248, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.261)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Silk Headband (7050, -0.40 DPS, sim-verified) [crafted]; Embalmed Shroud (7691, -0.53 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.53 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Glowing Green Talisman (5002, -1.23 DPS) [world_drop]; Crystal Starfire Medallion (5003, -1.23 DPS) [world_drop]; Pendant of Myzrael (4614, -2.29 DPS, sim-verified) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 | yes | Invoker's Mantle (215365, -0.35 DPS) [crafted]; Death Speaker Mantle (6685, -0.53 DPS) [dungeon]; Chestnut Mantle (17695, -2.09 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.02 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Battle Healer's Cloak (19529, -0.18 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, +0.45 DPS, sim-verified) [crafted]; High Robe of the Adjudicator (3461, -0.88 DPS) [quest]; Robes of Arcana (5770, -0.88 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Owlbeard Bracers (16981, -1.55 DPS, sim-verified) [quest]; Mindthrust Bracers (1974, -1.58 DPS) [dungeon]; Silver-lined Bracers (3224, -1.58 DPS) [world] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Jutebraid Gloves (10654, -0.18 DPS) [quest]; Gloves of Meditation (4318, -0.35 DPS) [crafted]; Gnoll Casting Gloves (892, -0.61 DPS, sim-verified) [world] |
| waist | Warsong Sash (16975) (or Defiler's Cloth Girdle (20164)) | Warsong Supplies [quest] | 11.0 | yes | Defiler's Cloth Girdle (20164, +0.46 DPS, sim-verified) [rep]; Belt of Arugal (6392, -0.35 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.53 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -0.19 DPS, sim-verified) [dungeon]; Silk-threaded Trousers (1929, -0.88 DPS) [dungeon]; Pristine Leggings (253987, -0.88 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 7.0 | yes | Boots of the Enchanter (4325, -0.35 DPS) [crafted]; Acidic Walkers (9454, -0.35 DPS) [dungeon]; Spidersilk Boots (4320, -2.22 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.35 DPS) [rep]; Electrocutioner Lagnut (9447, -0.70 DPS) [dungeon]; Sludge-Stained Band (286535, -0.70 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.53 DPS) [world]; Sacred Band (6669, -0.70 DPS) [quest]; Electrocutioner Lagnut (9447, -2.20 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | - | - |  |  |  |
| off_hand | Orb of Mystic Insight (249394) | Enchanting [crafted] | 7.0 | yes | Orb of Souls (249395, -0.53 DPS) [crafted]; Alliance Outrunner Healing Rod (285348, -0.53 DPS) [world]; Dwarven Tome (279898, -1.06 DPS, sim-verified) [quest] |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | World drop [world_drop] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Cookie's Stirring Rod (5198, -0.35 DPS) [dungeon]; Sizzle Stick (8071, -0.86 DPS, sim-verified) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; waist: Warsong Sash; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Notched Shortsword; off_hand: Orb of Mystic Insight; ranged: Wand of Eventide

No-known-source sample (15 of 474, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 105.8. Weights run: 1.2s. Verify run: 1.1s. 640 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.358), intellect=not significant (0.698 ± 0.495), crit=1.096 ± 0.054, hit=3.563 ± 0.219, spell_haste=-2.263 ± 0.479, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.358)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -0.89 DPS) [crafted]; Holy Shroud (2721, -1.11 DPS) [world_drop]; Augural Shroud (2620, -1.99 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 | yes | Darkspear Warding Pendant (272074, -0.70 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.86 DPS) [vendor]; Necklace of Calisea (1714, -2.18 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 | yes | Green Silken Shoulders (7057, +0.08 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.09 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.20 DPS) [vendor]; Icy Cloak (4327, -0.28 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.2 | yes | Robe of Power (7054, -0.42 DPS) [crafted]; Crimson Silk Vest (7058, -0.80 DPS) [crafted]; Dreamweave Vest (10021, -1.69 DPS, sim-verified) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 9.6 | yes | Spidertank Oilrag (9448, +0.75 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest]; Aurora Bracers (4043, -0.44 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.8 | yes | Black Mageweave Gloves (10003, -0.64 DPS) [crafted]; Gilded Handwraps (254021, -0.88 DPS) [crafted]; Red Mageweave Gloves (10018, -1.88 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 16.8 | yes | Gilded Cord (254037, -0.36 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.41 DPS) [rep]; Deathmage Sash (10771, -1.48 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 | yes | Abomination Skin Leggings (23173, -0.87 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.09 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.80 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -1.49 DPS) [dungeon]; Spidersilk Boots (4320, -1.58 DPS) [crafted]; Gilded Slippers (254001, -2.88 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.22 DPS) [rep]; Advisor's Ring (20426, -0.44 DPS) [rep]; Reedknot Ring (9622, -2.73 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Ogremind Ring (1993, -0.12 DPS) [world_drop]; Voodoo Band (1996, -0.12 DPS) [world]; Reedknot Ring (9622, -1.26 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 43.3 | yes | Illusionary Rod (7713, -0.88 DPS, sim-verified) [dungeon]; Staff of Noh'Orahil (15105, -2.80 DPS) [quest]; Windweaver Staff (7757, -3.65 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Fizzle's Zippy Lighter (6729) | News for Fizzle [quest] | 6.1 | yes | Twisted Nether Wand (249144, +0.67 DPS, sim-verified) [crafted]; Wand of Eventide (5214, -0.12 DPS) [world_drop]; Sizzle Stick (8071, -0.12 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Dar'Orahil; ranged: Fizzle's Zippy Lighter

No-known-source sample (15 of 640, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 173.2. Weights run: 1.2s. Verify run: 1.2s. 801 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.222, intellect=not significant (0.052 ± 0.247), crit=0.455 ± 0.023, hit=2.181 ± 0.152, spell_haste=-9.055 ± 0.409, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.222

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Spellpower Goggles Xtreme (10502, -1.95 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.95 DPS) [vendor]; Dreamweave Circlet (10041, -2.03 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.3 | yes | Mindburst Medallion (11196, -0.32 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.14 DPS) [world]; Darkspear Warding Pendant (272073, -2.22 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 14.8 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -0.29 DPS) [dungeon]; Black Mageweave Shoulders (10027, -1.42 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.3 | yes | Deep Woodlands Cloak (19121, -0.89 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.59 DPS) [crafted]; Nightfall Drape (12465, -1.72 DPS) [world] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.3 | yes | Acumen Robes (17775, -0.90 DPS, sim-verified) [quest]; Knight's Dreadweave Vest (220886, -1.09 DPS) [vendor]; Stone Guard's Dreadweave Vest (220904, -1.09 DPS) [vendor] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, +1.65 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.65 DPS) [quest]; Bloodband Bracers (11469, -1.15 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.2 | yes | Black Mageweave Gloves (10003, -0.52 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.54 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.54 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 15.6 | yes | Satyrmane Sash (17755, +1.32 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -0.46 DPS) [rep]; Ghostweave Cord (254073, -0.53 DPS) [crafted] |
| legs | Wizardweave Leggings (14132) | Tailoring [crafted] | 19.0 | yes | Stone Guard's Dreadweave Leggings (220906, -0.00 DPS) [vendor]; Red Mageweave Pants (10009, -1.42 DPS) [crafted]; Knight's Dreadweave Leggings (220888, -2.02 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 30.3 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -2.04 DPS) [crafted]; Gilded Sandals (254107, -6.11 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 21.8 | yes | Reedknot Ring (9622, -4.81 DPS) [quest]; Sea Giant's Toe Ring (274746, -5.13 DPS) [vendor]; Electrocutioner Lagnut (9447, -6.11 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Advisor's Ring (19520, -0.53 DPS, sim-verified) [rep]; Reedknot Ring (9622, -1.62 DPS) [quest]; Advisor's Ring (19521, -1.62 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 12.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Tidal Charm (1404, -3.90 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.36 DPS, sim-verified) [crafted]; Tidal Charm (1404, -1.95 DPS) [vendor] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 22.6 | yes | Staff of Dar'Orahil (15106, +0.50 DPS, sim-verified) [quest]; Kindling Stave (11750, -5.08 DPS) [dungeon]; Illusionary Rod (7713, -5.18 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Dreambough Wand (249234) | Enchanting [crafted] | 7.0 | yes | Twisted Nether Wand (249144, -0.32 DPS) [crafted]; Charged Lightning Rod (11860, -0.58 DPS) [quest]; Lesser Eternal Wand (249232, -1.91 DPS, sim-verified) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Wizardweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Soul Harvester; ranged: Dreambough Wand

No-known-source sample (15 of 801, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 350.7. Weights run: 1.3s. Verify run: 1.2s. 1236 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.837), intellect=not significant (-1.254 ± 0.970), crit=1.999 ± 0.106, hit=6.401 ± 0.516, spell_haste=not significant (0.241 ± 0.963), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.837)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | 153.0 | yes | Doomcaller's Circlet (21337, -4.06 DPS) [quest]; Deathmist Mask (226909, -10.29 DPS) [quest]; Bloodvine Goggles (19999, -11.72 DPS, sim-verified) [crafted] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 92.0 | yes | Beads of Ogre Might (22150, -4.06 DPS) [quest]; Medallion of the Dawn (22659, -9.28 DPS) [quest]; Charm of the Shifting Sands (21504, -9.72 DPS) [quest] |
| shoulder | Plagueheart Shoulderpads (22507) | Plagueheart Shoulderpads [quest] | 100.0 | yes | Doomcaller's Mantle (21335, -4.06 DPS, sim-verified) [quest]; Mantle of the Timbermaw (19050, -7.98 DPS) [crafted]; Champion's Dreadweave Spaulders (23256, -8.71 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 64.0 | yes | Chromatic Cloak (18509, -1.22 DPS, sim-verified) [crafted]; Shroud of Unspoken Names (21418, -6.67 DPS) [quest]; Spritecaster Cape (11623, -7.25 DPS) [dungeon] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 143.0 | yes | Bloodvine Vest (19682, -6.42 DPS, sim-verified) [crafted]; Zandalar Demoniac's Robe (20033, -7.54 DPS) [quest]; Doomcaller's Robes (21334, -10.73 DPS) [quest] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 91.0 | yes | Plagueheart Bindings (22511, -3.83 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -10.01 DPS) [rep]; Dryad's Wrist Bindings (19596, -10.30 DPS) [rep] |
| hands | Deathmist Wraps (22077) (or Deathmist Wraps (226911)) | Just Compensation [quest] | 77.0 | yes | Deathmist Wraps (226911, +0.00 DPS, sim-verified) [quest]; Gloves of Spell Mastery (14146, -1.75 DPS) [crafted]; Dreadmist Wraps (16705, -1.89 DPS) [dungeon] |
| waist | Plagueheart Belt (22510) | Plagueheart Belt [quest] | 62.0 | yes | Defiler's Cloth Girdle (20163, -2.90 DPS) [rep]; Defiler's Cloth Girdle (20165, -3.63 DPS) [rep]; Belt of the Archmage (18405, -6.93 DPS, sim-verified) [crafted] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 101.0 | yes | Plagueheart Leggings (22505, -5.23 DPS) [quest]; Magister's Leggings (16687, -5.28 DPS) [dungeon]; Deathmist Leggings (226910, -10.48 DPS, sim-verified) [quest] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 83.0 | yes | First Sergeant's Dreadweave Boots (220909, -1.60 DPS) [vendor]; Plagueheart Sandals (22508, -3.34 DPS) [quest]; Sergeant Major's Dreadweave Boots (220891, -5.43 DPS, sim-verified) [vendor] |
| finger1 | Ring of Unspoken Names (21417) | Ring of Unspoken Names [quest] | 106.0 | yes | Don Julio's Band (19325, -2.03 DPS) [rep]; Band of Earthen Might (21182, -2.03 DPS) [quest]; Blackstone Ring (17713, -6.09 DPS) [dungeon] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 101.0 | yes | Band of Earthen Might (21182, -1.31 DPS) [quest]; Blackstone Ring (17713, -5.37 DPS) [dungeon]; Don Julio's Band (19325, -15.36 DPS, sim-verified) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 12.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Tidal Charm (1404, -1.74 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Tidal Charm (1404, -0.87 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.04 DPS, sim-verified) [crafted] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 137.0 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Atiesh, Greatstaff of the Guardian (22631, -2.46 DPS) [quest]; Grand Marshal's Stave (18873, -9.57 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Greater Eternal Wand (249237) (or Brilliant Wand (249385)) | Enchanting [crafted] | 9.0 | yes | Lesser Eternal Wand (249232, -0.15 DPS) [crafted]; Dreambough Wand (249234, -0.29 DPS) [crafted]; Brilliant Wand (249385, -0.56 DPS, sim-verified) [crafted] |

**New at 60:** head: Plagueheart Circlet; neck: Onyxia Tooth Pendant; shoulder: Plagueheart Shoulderpads; back: Earthweave Cloak; chest: Plagueheart Robe; wrist: Rockfury Bracers; hands: Deathmist Wraps; waist: Plagueheart Belt; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Ring of Unspoken Names; finger2: Ring of the Fallen God; main_hand: High Warlord's War Staff; ranged: Greater Eternal Wand

No-known-source sample (15 of 1236, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

