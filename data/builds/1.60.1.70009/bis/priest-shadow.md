# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 22.7. Weights run: 0.6s. Verify run: 0.7s. 256 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=-3.775 ± 0.282, crit=1.211 ± 0.056, hit=1.697 ± 0.274, spell_haste=not significant (-0.927 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.54 DPS) [crafted]; Lucky Fishing Hat (19972, -0.54 DPS) [quest]; Flying Tiger Goggles (4368, -1.20 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.02 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.45 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Caretaker's Cape (20428, -0.09 DPS) [rep]; Feyscale Cloak (6632, -0.09 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.09 DPS) [crafted]; Bloody Apron (6226, -0.09 DPS) [dungeon]; Green Woolen Vest (2582, -0.42 DPS, sim-verified) [crafted] |
| wrist | Silver-lined Bracers (3224) | Timber [world] | 0.0 | yes | Seer's Cuffs (3645, +0.00 DPS) [world_drop]; Bright Bracers (3647, +0.00 DPS) [world_drop]; Mindthrust Bracers (1974, -0.44 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.27 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.45 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.0 | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.34 DPS, sim-verified) [crafted]; Lesser Belt of the Spire (1299, -0.36 DPS) [world] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Silk-threaded Trousers (1929, +0.03 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Colorful Kilt (10048, -0.36 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Feather Padded Treads (285345, -0.11 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.27 DPS) [crafted]; Pristine Boots (253889, -0.36 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.18 DPS) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.13 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | - | - |  |  |  |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Pulsating Hydra Heart (5183, -0.45 DPS) [world]; Tear of Grief (5611, -0.45 DPS) [quest]; Grayson's Torch (1172, -0.52 DPS, sim-verified) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Sable Wand (7607, -0.18 DPS) [quest]; Cookie's Stirring Rod (5198, -0.21 DPS, sim-verified) [dungeon]; Torchlight Wand (5240, -0.27 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Silver-lined Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Stout Battlehammer; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 256, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (gnome, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 51.6. Weights run: 0.5s. Verify run: 0.5s. 477 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=not significant (0.425 ± 0.239), crit=1.229 ± 0.067, hit=3.796 ± 0.202, spell_haste=not significant (0.186 ± 0.312), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, +0.22 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.18 DPS) [crafted]; Embalmed Shroud (7691, -0.27 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.5 | yes | Crystal Starfire Medallion (5003, -0.70 DPS) [world_drop]; Pendant of Myzrael (4614, -0.85 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.18 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.8 | yes | Death Speaker Mantle (6685, -0.26 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.27 DPS) [quest]; Invoker's Mantle (215365, -0.33 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Repairman's Cape (9605, +0.12 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Prelacy Cape (7004, -0.09 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.5 | yes | Death Speaker Robes (6682, -0.25 DPS) [dungeon]; Pristine Gown (253961, -0.41 DPS) [crafted]; Tree Bark Jacket (1486, -1.05 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.58 DPS) [quest]; Mindthrust Bracers (1974, -0.62 DPS) [dungeon]; Nightsky Wristbands (6407, -0.74 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.7 | yes | Shilly Mitts (9609, -0.15 DPS) [quest]; Truefaith Gloves (7049, -0.21 DPS) [crafted]; Serpent Gloves (5970, -0.43 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.3 | yes | Invoker's Cord (215366, -0.28 DPS) [crafted]; Belt of Arugal (6392, -0.29 DPS, sim-verified) [dungeon]; Crimson Silk Belt (7055, -0.30 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.4 | yes | Pristine Leggings (253987, -0.22 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.25 DPS, sim-verified) [dungeon]; Filigreed Pristine Leggings (253937, -0.34 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.14 DPS) [dungeon]; Nimbus Boots (6998, -0.36 DPS) [quest]; Spidersilk Boots (4320, -1.42 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.10 DPS) [quest]; Lorekeeper's Ring (20431, -0.18 DPS) [rep]; Electrocutioner Lagnut (9447, -0.36 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.27 DPS) [dungeon]; Sludge-Stained Band (286535, -0.27 DPS) [world]; Minor Channeling Ring (1449, -0.73 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.7 | yes | Gnarled Necromancer's Staff (251534, -0.04 DPS) [quest]; Channeler's Staff (4437, -0.11 DPS) [world]; Twisted Chanter's Staff (890, -0.36 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | World drop [world_drop] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Spellcrafter Wand (6677, -0.09 DPS) [quest]; Sizzle Stick (8071, -1.60 DPS, sim-verified) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 477, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (gnome, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 79.5. Weights run: 0.5s. Verify run: 0.5s. 647 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.442 ± 0.373), crit=1.468 ± 0.091, hit=5.629 ± 0.331, spell_haste=not significant (0.939 ± 0.654), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.36 DPS, sim-verified) [world]; Holy Shroud (2721, -0.96 DPS) [world_drop]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.7 | yes | Necklace of Calisea (1714, -0.23 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.63 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.71 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 | yes | Green Silken Shoulders (7057, +0.35 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.02 DPS) [dungeon]; Berylline Pads (4197, -0.15 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.12 DPS) [crafted]; Caretaker's Cape (19532, -0.21 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.7 | yes | Dreamweave Vest (10021, +0.36 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.51 DPS) [crafted]; Crimson Silk Vest (7058, -0.79 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.69 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.19 DPS) [quest]; Earthen Silk Cuffs (254019, -0.48 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 | yes | Red Mageweave Gloves (10018, -0.39 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.46 DPS) [crafted]; Gilded Handwraps (254021, -0.83 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 | yes | Star Belt (4329, -0.06 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.13 DPS) [rep]; Highlander's Cloth Girdle (20098, -1.21 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.3 | yes | Crimson Silk Pantaloons (7062, -0.13 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.65 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.70 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.46 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.46 DPS) [crafted]; Acidic Walkers (9454, -1.48 DPS) [dungeon] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.19 DPS) [quest]; Lorekeeper's Ring (19525, -0.19 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.29 DPS) [vendor] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 | yes | Reedknot Ring (9622, -0.11 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.19 DPS) [vendor]; Minor Channeling Ring (1449, -0.20 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Windweaver Staff (7757) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 6.6 | yes | Iron Morningstar (250606, -0.16 DPS) [crafted]; Staff of Jordan (873, -0.17 DPS) [world_drop]; Illusionary Rod (7713, -1.40 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 6.0 | yes | Fizzle's Zippy Lighter (6729, -0.06 DPS) [quest]; Wand of Eventide (5214, -0.10 DPS) [world_drop]; Burning Sliver (5249, -1.24 DPS, sim-verified) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Windweaver Staff; ranged: Twisted Nether Wand

No-known-source sample (15 of 647, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs

### Band 50 (gnome, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 116.5. Weights run: 0.5s. Verify run: 0.5s. 812 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (-0.242 ± 0.580), crit=1.866 ± 0.121, hit=8.219 ± 0.471, spell_haste=not significant (-1.234 ± 0.835), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 40.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.39 DPS) [crafted]; Eye of Theradras (17715, -1.48 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.17 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -0.74 DPS) [world_drop]; Pendant of Myzrael (4614, -0.74 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 34.1 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -2.23 DPS) [dungeon]; Black Mageweave Shoulders (10027, -2.55 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Runecloth Cloak (13860, -0.53 DPS) [crafted]; Icy Cloak (4327, -0.74 DPS) [crafted]; Nightfall Drape (12465, -3.80 DPS, sim-verified) [world] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 38.1 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -1.70 DPS) [world_drop]; Acumen Robes (17775, -2.02 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.69 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.21 DPS) [quest]; Nethergeld Cuffs (254061, -0.21 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.53 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.53 DPS) [vendor]; Black Mageweave Gloves (10003, -3.55 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 35.1 | yes | Satyrmane Sash (17755, +1.05 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -2.23 DPS) [rep]; Ghostweave Cord (254073, -2.23 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 38.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -2.02 DPS) [crafted]; Red Mageweave Pants (10009, -2.55 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 90.2 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -7.00 DPS) [crafted]; Black Mageweave Boots (10026, -8.37 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 82.2 | yes | Ring of Forlorn Spirits (2043, -7.84 DPS) [quest]; Reedknot Ring (9622, -7.95 DPS) [quest]; Sea Giant's Toe Ring (274746, -8.05 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Ring of Forlorn Spirits (2043, -0.42 DPS) [quest]; Reedknot Ring (9622, -0.53 DPS) [quest]; Lorekeeper's Ring (19524, -0.83 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 74.0 | yes | Thunderbrew's Boot Flask (744, -7.82 DPS) [quest]; Tidal Charm (1404, -7.82 DPS) [vendor]; Guardian Talisman (1490, -7.82 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Thunderbrew's Boot Flask (744, -0.63 DPS) [quest]; Tidal Charm (1404, -0.63 DPS) [vendor]; Guardian Talisman (1490, -0.63 DPS) [quest] |
| main_hand | Might of Hakkar (10838) | Avatar of Hakkar [world] | 41.1 | yes | Illusionary Rod (7713, -1.58 DPS) [dungeon]; Kindling Stave (11750, -1.58 DPS) [dungeon]; Iron Morningstar (250606, -3.82 DPS) [crafted] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 | yes | Thrash's Trash (276204, +0.22 DPS, sim-verified) [vendor]; Truesilver Conduit (249455, -0.42 DPS) [crafted]; Orb of Mystic Insight (249394, -0.74 DPS) [crafted] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Burning Sliver (5249, -0.21 DPS) [quest]; Twisted Nether Wand (249144, -0.21 DPS) [crafted]; Dreambough Wand (249234, -4.25 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Might of Hakkar; off_hand: Orb of the Forgotten Seer; ranged: Lesser Eternal Wand

No-known-source sample (15 of 812, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 60 (gnome, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 226.5. Weights run: 0.6s. Verify run: 0.6s. 1314 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.724 ± 0.748), crit=2.882 ± 0.178, hit=12.470 ± 0.766, spell_haste=not significant (1.648 ± 0.891), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tiara of the Oracle (21348) | Tiara of the Oracle [quest] | 168.6 | yes | Bloodvine Goggles (19999, -9.92 DPS, sim-verified) [crafted]; Virtuous Crown (22080, -12.06 DPS) [quest]; Virtuous Crown (226947, -12.06 DPS) [quest] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 165.1 | yes | Beads of Ogre Might (22150, -4.64 DPS) [quest]; Medallion of the Dawn (22659, -14.33 DPS) [quest]; Charm of the Shifting Sands (21504, -15.10 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 66.8 | yes | Shroud of the Nathrezim (18720, -1.22 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.37 DPS) [vendor]; Blood Guard's Dreadweave Mantle (220905, -1.37 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 124.7 | yes | Chromatic Cloak (18509, +0.10 DPS, sim-verified) [crafted]; Hide of the Wild (18510, -11.89 DPS) [crafted]; Spritecaster Cape (11623, -12.22 DPS) [dungeon] |
| chest | Vestments of the Oracle (21351) | Vestments of the Oracle [quest] | 95.2 | yes | Earthpower Vest (21183, -1.60 DPS) [quest]; Virtuous Robe (226945, -3.59 DPS) [quest]; Bloodvine Vest (19682, -5.17 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 151.7 | yes | Dryad's Wrist Bindings (19595, -0.21 DPS, sim-verified) [rep]; Dryad's Wrist Bindings (19596, -14.64 DPS) [rep]; Dryad's Wrist Bindings (19597, -15.10 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 96.9 | yes | Dreadmist Wraps (16705, -3.46 DPS, sim-verified) [dungeon]; Marshal's Satin Gloves (17608, -7.41 DPS) [vendor]; General's Satin Gloves (17620, -7.41 DPS) [vendor] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 71.9 | yes | Highlander's Cloth Girdle (20097, -2.18 DPS) [rep]; Highlander's Cloth Girdle (20047, -3.90 DPS, sim-verified) [rep]; Stormpike Cloth Girdle (19094, -5.37 DPS) [rep] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 166.0 | yes | Magister's Leggings (16687, -0.06 DPS, sim-verified) [dungeon]; Knight's Dreadweave Leggings (220888, -12.07 DPS) [vendor]; Stone Guard's Dreadweave Leggings (220906, -12.07 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 155.3 | yes | Sergeant Major's Dreadweave Boots (220891, +0.61 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Boots (220909, -1.85 DPS) [vendor]; Marshal's Satin Sandals (17607, -13.70 DPS) [vendor] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 166.0 | yes | Band of Earthen Might (21182, -0.11 DPS) [quest]; Blackstone Ring (17713, -4.75 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -4.75 DPS) [quest] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 165.1 | yes | Band of Earthen Might (21182, +0.00 DPS, sim-verified) [quest]; Blackstone Ring (17713, -4.64 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -4.64 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 112.2 | yes | Thunderbrew's Boot Flask (744, -12.90 DPS) [quest]; Tidal Charm (1404, -12.90 DPS) [vendor]; Guardian Talisman (1490, -12.90 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Thunderbrew's Boot Flask (744, -0.69 DPS) [quest]; Tidal Charm (1404, -0.69 DPS) [vendor]; Guardian Talisman (1490, -0.69 DPS) [quest] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 178.3 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Persuader (22384, -1.53 DPS) [crafted]; Atiesh, Greatstaff of the Guardian (22631, -4.38 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 14.1 | yes | Greater Eternal Wand (249237, +0.88 DPS, sim-verified) [crafted]; Lesser Eternal Wand (249232, -0.70 DPS) [crafted]; Charged Lightning Rod (11860, -0.71 DPS) [quest] |

**New at 60:** head: Tiara of the Oracle; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Vestments of the Oracle; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Belt of the Archmage; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Ring of the Fallen God; finger2: Don Julio's Band; main_hand: High Warlord's War Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1314, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 28.0. Weights run: 0.6s. Verify run: 0.6s. 255 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=-3.775 ± 0.282, crit=1.211 ± 0.056, hit=1.697 ± 0.274, spell_haste=not significant (-0.927 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.54 DPS) [crafted]; Lucky Fishing Hat (19972, -0.54 DPS) [quest]; Flying Tiger Goggles (4368, -1.11 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.40 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.45 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Battle Healer's Cloak (20427, -0.09 DPS) [rep]; Feyscale Cloak (6632, -0.71 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.09 DPS) [crafted]; Bloody Apron (6226, -0.09 DPS) [dungeon]; Green Woolen Vest (2582, -0.75 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Silver-lined Bracers (3224, +0.00 DPS) [world]; Seer's Cuffs (3645, +0.00 DPS) [world_drop]; Owlbeard Bracers (16981, -0.49 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.09 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.27 DPS) [quest]; Pristine Gloves (253913, -0.27 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.0 | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.36 DPS) [world]; Novice Arcanist's Sash (253885, -0.37 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Colorful Kilt (10048, -0.36 DPS) [crafted]; Silk-threaded Trousers (1929, -1.82 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.27 DPS) [crafted]; Pristine Boots (253889, -0.36 DPS) [crafted]; Feather Padded Treads (285345, -0.49 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world]; Ring of Scorn (3235, -0.45 DPS) [quest] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Ring of the Shadow (1462, -0.27 DPS) [world]; Ring of Scorn (3235, -0.27 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.22 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | - | - |  |  |  |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Grayson's Torch (1172, -0.19 DPS, sim-verified) [quest]; Nightglow Concoction (3451, -0.45 DPS) [quest]; Pulsating Hydra Heart (5183, -0.45 DPS) [world] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Torchlight Wand (5240, -0.27 DPS) [quest]; Greater Magic Wand (11288, -0.27 DPS) [crafted]; Cookie's Stirring Rod (5198, -0.28 DPS, sim-verified) [dungeon] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Stout Battlehammer; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 255, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak

### Band 30 (undead, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 46.3. Weights run: 0.5s. Verify run: 0.5s. 476 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=not significant (0.425 ± 0.239), crit=1.229 ± 0.067, hit=3.796 ± 0.202, spell_haste=not significant (0.186 ± 0.312), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 10.2 | yes | Silk Headband (7050, -0.11 DPS) [crafted]; Embalmed Shroud (7691, -0.20 DPS) [dungeon]; Holy Shroud (2721, -0.73 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.5 | yes | Crystal Starfire Medallion (5003, -0.70 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.78 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -0.85 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.8 | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.27 DPS) [quest]; Invoker's Mantle (215365, -0.33 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.10 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Battle Healer's Cloak (19529, -0.09 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.5 | yes | Death Speaker Robes (6682, -0.25 DPS) [dungeon]; Pristine Gown (253961, -0.41 DPS) [crafted]; Tree Bark Jacket (1486, -0.86 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.58 DPS) [quest]; Mindthrust Bracers (1974, -0.62 DPS) [dungeon]; Nightsky Wristbands (6407, -0.78 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.1 | yes | Serpent Gloves (5970, -0.10 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.17 DPS) [crafted]; Gnoll Casting Gloves (892, -0.19 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.3 | yes | Warsong Sash (16975, -0.17 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Invoker's Cord (215366, -0.28 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.4 | yes | Gaze Dreamer Pants (6903, -0.11 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.34 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 | yes | Acidic Walkers (9454, -0.14 DPS) [dungeon]; Boots of the Enchanter (4325, -0.44 DPS) [crafted]; Spidersilk Boots (4320, -0.93 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.18 DPS) [rep]; Electrocutioner Lagnut (9447, -0.36 DPS) [dungeon]; Sludge-Stained Band (286535, -0.36 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.27 DPS) [world]; Black Widow Band (6199, -0.27 DPS) [world]; Electrocutioner Lagnut (9447, -0.76 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 4.7 | yes | Gnarled Necromancer's Staff (251534, -0.04 DPS) [quest]; Twisted Chanter's Staff (890, -0.09 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.11 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | World drop [world_drop] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Cookie's Stirring Rod (5198, -0.18 DPS) [dungeon]; Sizzle Stick (8071, -1.81 DPS, sim-verified) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 476, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (undead, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 70.0. Weights run: 0.5s. Verify run: 0.5s. 645 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.442 ± 0.373), crit=1.468 ± 0.091, hit=5.629 ± 0.331, spell_haste=not significant (0.939 ± 0.654), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.25 DPS, sim-verified) [world]; Holy Shroud (2721, -0.96 DPS) [world_drop]; Enchanter's Cowl (4322, -1.02 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.7 | yes | Necklace of Calisea (1714, -0.15 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.63 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.71 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 13.0 | yes | Green Silken Shoulders (7057, +0.18 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.02 DPS) [dungeon]; Berylline Pads (4197, -0.15 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.2 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.12 DPS) [crafted]; Battle Healer's Cloak (19528, -0.21 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 24.7 | yes | Dreamweave Vest (10021, +0.62 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.51 DPS) [crafted]; Crimson Silk Vest (7058, -0.79 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Condor Bracers (15864, -0.19 DPS) [quest]; Radiant Silver Bracers (4545, -0.38 DPS, sim-verified) [quest]; Earthen Silk Cuffs (254019, -0.48 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.8 | yes | Red Mageweave Gloves (10018, -0.13 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.46 DPS) [crafted]; Gilded Handwraps (254021, -0.83 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 13.6 | yes | Star Belt (4329, -0.06 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.13 DPS) [rep]; Defiler's Cloth Girdle (20166, -0.72 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 19.3 | yes | Crimson Silk Pantaloons (7062, -0.31 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -0.65 DPS) [dungeon]; Gaze Dreamer Pants (6903, -0.70 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.54 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.46 DPS) [crafted]; Acidic Walkers (9454, -1.48 DPS) [dungeon] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.19 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.29 DPS) [vendor]; Advisor's Ring (20426, -0.38 DPS) [rep] |
| finger2 | Reedknot Ring (9622) | Jarl Needs a Blade [quest] | 7.0 | yes | Sea Giant's Toe Ring (274746, +0.50 DPS, sim-verified) [vendor]; Ogremind Ring (1993, -0.37 DPS) [world_drop]; Voodoo Band (1996, -0.37 DPS) [world] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 23.2 | yes | Windweaver Staff (7757, +0.64 DPS, sim-verified) [dungeon]; Iron Morningstar (250606, -1.75 DPS) [crafted]; Staff of Jordan (873, -1.76 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 6.0 | yes | Wand of Eventide (5214, -0.10 DPS) [world_drop]; Sizzle Stick (8071, -0.10 DPS) [quest]; Fizzle's Zippy Lighter (6729, -0.67 DPS, sim-verified) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; finger2: Reedknot Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 645, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (undead, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 104.8. Weights run: 0.5s. Verify run: 0.6s. 810 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (-0.242 ± 0.580), crit=1.866 ± 0.121, hit=8.219 ± 0.471, spell_haste=not significant (-1.234 ± 0.835), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 40.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Spellpower Goggles Xtreme Plus (15999, -1.39 DPS) [crafted]; Eye of Theradras (17715, -1.48 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, +0.37 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -0.74 DPS) [world_drop]; Choker of the High Shaman (4112, -0.74 DPS) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 34.1 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -2.23 DPS) [dungeon]; Black Mageweave Shoulders (10027, -2.55 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Deep Woodlands Cloak (19121, -0.27 DPS, sim-verified) [quest]; Nightfall Drape (12465, -0.53 DPS) [world]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 38.1 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -1.70 DPS) [world_drop]; Acumen Robes (17775, -2.02 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, -0.21 DPS) [crafted]; Bloodband Bracers (11469, -0.42 DPS) [quest]; Condor Bracers (15864, -1.02 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.53 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.53 DPS) [vendor]; Black Mageweave Gloves (10003, -1.34 DPS, sim-verified) [crafted] |
| waist | Satyrmane Sash (17755) | Maraudon: Lord Vyletongue [dungeon] | 14.0 | yes | Defiler's Cloth Girdle (20166, +0.00 DPS) [rep]; Ghostweave Cord (254073, +0.00 DPS) [crafted]; Defiler's Cloth Girdle (20165, -1.46 DPS, sim-verified) [rep] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 38.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -2.02 DPS) [crafted]; Red Mageweave Pants (10009, -2.55 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 90.2 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -7.00 DPS) [crafted]; Black Mageweave Boots (10026, -8.37 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 82.2 | yes | Reedknot Ring (9622, -7.95 DPS) [quest]; Sea Giant's Toe Ring (274746, -8.05 DPS) [vendor]; Electrocutioner Lagnut (9447, -8.37 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Reedknot Ring (9622, -0.53 DPS) [quest]; Advisor's Ring (19521, -0.53 DPS) [rep]; Advisor's Ring (19520, -0.64 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 74.0 | yes | Rune of the Guard Captain (19120, -1.74 DPS) [quest]; Uther's Strength (11302, -7.18 DPS) [world]; Tidal Charm (1404, -7.82 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world]; Rune of the Guard Captain (19120, -0.18 DPS, sim-verified) [quest] |
| main_hand | Might of Hakkar (10838) | Avatar of Hakkar [world] | 41.1 | yes | Illusionary Rod (7713, -1.58 DPS) [dungeon]; Kindling Stave (11750, -1.58 DPS) [dungeon]; Iron Morningstar (250606, -3.82 DPS) [crafted] |
| off_hand | Orb of the Forgotten Seer (7685) (or Thrash's Trash (276204)) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 14.0 | yes | Thrash's Trash (276204, +0.19 DPS, sim-verified) [vendor]; Truesilver Conduit (249455, -0.42 DPS) [crafted]; Omega Orb (7749, -0.53 DPS) [quest] |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Twisted Nether Wand (249144, -0.21 DPS) [crafted]; Wand of Eventide (5214, -0.32 DPS) [world_drop]; Dreambough Wand (249234, -4.72 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; waist: Satyrmane Sash; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; main_hand: Might of Hakkar; off_hand: Orb of the Forgotten Seer; ranged: Lesser Eternal Wand

No-known-source sample (15 of 810, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (undead, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 218.8. Weights run: 0.6s. Verify run: 0.6s. 1312 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.724 ± 0.748), crit=2.882 ± 0.178, hit=12.470 ± 0.766, spell_haste=not significant (1.648 ± 0.891), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tiara of the Oracle (21348) | Tiara of the Oracle [quest] | 168.6 | yes | Virtuous Crown (22080, -12.06 DPS) [quest]; Virtuous Crown (226947, -12.06 DPS) [quest]; Bloodvine Goggles (19999, -14.69 DPS, sim-verified) [crafted] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 165.1 | yes | Beads of Ogre Might (22150, -4.64 DPS) [quest]; Medallion of the Dawn (22659, -14.33 DPS) [quest]; Charm of the Shifting Sands (21504, -15.10 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 66.8 | yes | Shroud of the Nathrezim (18720, +1.33 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.37 DPS) [vendor]; Blood Guard's Dreadweave Mantle (220905, -1.37 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 124.7 | yes | Chromatic Cloak (18509, -0.15 DPS, sim-verified) [crafted]; Hide of the Wild (18510, -11.89 DPS) [crafted]; Deep Woodlands Cloak (19121, -12.20 DPS) [quest] |
| chest | Vestments of the Oracle (21351) | Vestments of the Oracle [quest] | 95.2 | yes | Earthpower Vest (21183, -1.60 DPS) [quest]; Virtuous Robe (226945, -3.59 DPS) [quest]; Bloodvine Vest (19682, -7.70 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 27.8 | yes | Dryad's Wrist Bindings (19596, -0.40 DPS) [rep]; Dryad's Wrist Bindings (19597, -0.86 DPS) [rep]; Rockfury Bracers (21186, -2.86 DPS, sim-verified) [quest] |
| hands | Dreadmist Wraps (16705) | Scholomance: Lorekeeper Polkelt [dungeon] | 131.2 | yes | Gloves of Spell Mastery (14146, -1.89 DPS, sim-verified) [crafted]; Marshal's Satin Gloves (17608, -11.35 DPS) [vendor]; General's Satin Gloves (17620, -11.35 DPS) [vendor] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 71.9 | yes | Defiler's Cloth Girdle (20163, -0.48 DPS, sim-verified) [rep]; Defiler's Cloth Girdle (20165, -2.18 DPS) [rep]; Frostwolf Cloth Belt (19090, -5.37 DPS) [rep] |
| legs | Magister's Leggings (16687) | Stratholme: Baron Rivendare [dungeon] | 130.1 | yes | Bloodvine Leggings (19683, -3.31 DPS, sim-verified) [crafted]; Knight's Dreadweave Leggings (220888, -7.94 DPS) [vendor]; Stone Guard's Dreadweave Leggings (220906, -7.94 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 155.3 | yes | First Sergeant's Dreadweave Boots (220909, -1.85 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -3.04 DPS, sim-verified) [vendor]; Marshal's Satin Sandals (17607, -13.70 DPS) [vendor] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 166.0 | yes | Band of Earthen Might (21182, -0.11 DPS) [quest]; Blackstone Ring (17713, -4.75 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -4.75 DPS) [quest] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 165.1 | yes | Band of Earthen Might (21182, +0.00 DPS, sim-verified) [quest]; Blackstone Ring (17713, -4.64 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -4.64 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Tidal Charm (1404, -0.69 DPS) [vendor]; Frozen Heart of the Mountain (249469, -0.74 DPS, sim-verified) [crafted] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 178.3 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Persuader (22384, -1.53 DPS) [crafted]; Atiesh, Greatstaff of the Guardian (22631, -4.38 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 14.1 | yes | Lesser Eternal Wand (249232, -0.70 DPS) [crafted]; Charged Lightning Rod (11860, -0.71 DPS) [quest]; Greater Eternal Wand (249237, -2.61 DPS, sim-verified) [crafted] |

**New at 60:** head: Tiara of the Oracle; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Vestments of the Oracle; wrist: Dryad's Wrist Bindings; hands: Dreadmist Wraps; waist: Belt of the Archmage; legs: Magister's Leggings; feet: Bloodvine Boots; finger1: Ring of the Fallen God; finger2: Don Julio's Band; trinket1: Ankh of Life; trinket2: Uther's Strength; main_hand: High Warlord's War Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1312, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

