# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 32.7. Weights run: 0.5s. Verify run: 0.5s. 252 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=not significant (-0.108 ± 0.264), crit=2.133 ± 0.074, hit=5.698 ± 0.181, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.49 DPS) [crafted]; Lucky Fishing Hat (19972, -0.49 DPS) [quest]; Flying Tiger Goggles (4368, -1.17 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Slime-encrusted Pads (6461, -0.41 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -0.69 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Caretaker's Cape (20428, -0.08 DPS) [rep]; Feyscale Cloak (6632, -0.12 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.08 DPS) [crafted]; Bloody Apron (6226, -0.08 DPS) [dungeon]; Green Woolen Vest (2582, -1.14 DPS, sim-verified) [crafted] |
| wrist | - | - |  |  |  |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.24 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.41 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.0 | yes | Novice Ardent's Sash (253887, -0.16 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.33 DPS) [world]; Novice Arcanist's Sash (253885, -0.34 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.24 DPS) [crafted]; Colorful Kilt (10048, -0.33 DPS) [crafted]; Silk-threaded Trousers (1929, -0.85 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.24 DPS) [crafted]; Pristine Boots (253889, -0.33 DPS) [crafted]; Feather Padded Treads (285345, -0.78 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.16 DPS) [world]; Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon]; Ring of the Shadow (1462, -0.41 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.28 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.41 DPS) [dungeon]; Ring of the Shadow (1462, -0.41 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | - | - |  |  |  |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Pulsating Hydra Heart (5183, -0.41 DPS) [world]; Tear of Grief (5611, -0.41 DPS) [quest]; Grayson's Torch (1172, -0.50 DPS, sim-verified) [quest] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Sable Wand (7607, -0.16 DPS) [quest]; Torchlight Wand (5240, -0.24 DPS) [quest]; Cookie's Stirring Rod (5198, -0.36 DPS, sim-verified) [dungeon] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Notched Shortsword; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 252, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 56.8. Weights run: 0.6s. Verify run: 0.5s. 476 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.157 ± 0.128), crit=1.917 ± 0.062, hit=4.834 ± 0.199, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Embalmed Shroud (7691, -0.35 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.35 DPS) [crafted]; Silk Headband (7050, -0.73 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Pendant of Myzrael (4614, -0.92 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.24 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 | yes | Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Fairywing Mantle (9536, -0.35 DPS) [quest]; Invoker's Mantle (215365, -0.66 DPS, sim-verified) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, -0.05 DPS, sim-verified) [crafted]; Prelacy Cape (7004, -0.12 DPS) [quest]; Caretaker's Cape (19533, -0.12 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 11.0 | yes | Death Speaker Robes (6682, -0.27 DPS) [dungeon]; Pristine Gown (253961, -0.34 DPS) [crafted]; Tree Bark Jacket (1486, -0.74 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.93 DPS) [quest]; Mindthrust Bracers (1974, -0.95 DPS) [dungeon]; Nightsky Wristbands (6407, -1.11 DPS, sim-verified) [world_drop] |
| hands | Serpent Gloves (5970) (or Shilly Mitts (9609)) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Shilly Mitts (9609, +0.34 DPS, sim-verified) [quest]; Gnoll Casting Gloves (892, -0.12 DPS) [world]; Town Clerk's Mittens (270029, -0.15 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.5 | yes | Belt of Arugal (6392, -0.02 DPS, sim-verified) [dungeon]; Ghamoo-ra's Bind (6908, -0.40 DPS) [dungeon]; Invoker's Cord (215366, -0.43 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -0.05 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.45 DPS) [crafted]; Silk-threaded Trousers (1929, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Nimbus Boots (6998, -0.24 DPS) [quest]; Spidersilk Boots (4320, -1.54 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.19 DPS) [quest]; Lorekeeper's Ring (20431, -0.23 DPS) [rep]; Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.35 DPS) [dungeon]; Sludge-Stained Band (286535, -0.35 DPS) [world]; Minor Channeling Ring (1449, -0.93 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 1.7 | yes | Gnarled Necromancer's Staff (251534, -0.02 DPS) [quest]; Channeler's Staff (4437, -0.05 DPS) [world]; Twisted Chanter's Staff (890, -0.26 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | World drop [world_drop] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Spellcrafter Wand (6677, -0.12 DPS) [quest]; Sizzle Stick (8071, -0.91 DPS, sim-verified) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 476, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (gnome, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 180.7. Weights run: 0.7s. Verify run: 0.6s. 643 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.091 ± 0.015, crit=3.159 ± 0.088, hit=3.864 ± 0.105, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -3.07 DPS, sim-verified) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop]; Silk Headband (7050, -4.11 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 | yes | Darkspear Warding Pendant (272074, -2.37 DPS) [vendor]; Necklace of Calisea (1714, -2.40 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.43 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 | yes | Green Silken Shoulders (7057, -0.23 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Berylline Pads (4197, -0.65 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, -0.12 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Caretaker's Cape (19532, -0.34 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 | yes | Dreamweave Vest (10021, -1.23 DPS, sim-verified) [crafted]; Robe of Power (7054, -2.56 DPS) [crafted]; Tree Bark Jacket (1486, -3.27 DPS) [dungeon] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.69 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.26 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.4 | yes | Star Belt (4329, -0.56 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -1.06 DPS) [rep]; Belt of Arugal (6392, -1.75 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 | yes | Gaze Dreamer Pants (6903, -1.28 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.03 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -3.32 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Nimbus Boots (6998, -6.17 DPS) [quest] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.69 DPS) [quest]; Lorekeeper's Ring (19525, -0.69 DPS) [rep]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 | yes | Reedknot Ring (9622, -0.35 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.69 DPS) [vendor]; Minor Channeling Ring (1449, -0.97 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 44.8 | yes | Windweaver Staff (7757, -1.13 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -15.00 DPS) [world_drop]; Glimmering Staff (249392, -15.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Burning Sliver (5249) (or Twisted Nether Wand (249144)) | Crushridge Warmongers [quest] | 6.0 | yes | Twisted Nether Wand (249144, +0.00 DPS, sim-verified) [crafted]; Wand of Eventide (5214, -0.34 DPS) [world_drop]; Sizzle Stick (8071, -0.34 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Burning Sliver

No-known-source sample (15 of 643, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (gnome, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 255.2. Weights run: 0.7s. Verify run: 0.6s. 806 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.106 ± 0.020, crit=4.628 ± 0.134, hit=5.657 ± 0.238, spell_haste=3.000 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 80.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -4.64 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.59 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 | yes | Mindburst Medallion (11196, -0.39 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.15 DPS) [world]; Darkspear Warding Pendant (272073, -2.34 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 73.7 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -20.61 DPS) [dungeon]; Black Mageweave Shoulders (10027, -22.00 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 | yes | Runecloth Cloak (13860, -1.82 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.98 DPS) [world]; Icy Cloak (4327, -2.68 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 78.0 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -19.38 DPS) [world_drop]; Acumen Robes (17775, -19.91 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.44 DPS) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.45 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.57 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.57 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 74.3 | yes | Satyrmane Sash (17755, -0.25 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -20.98 DPS) [rep]; Ghostweave Cord (254073, -21.13 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 78.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -20.69 DPS) [crafted]; Red Mageweave Pants (10009, -22.00 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 65.5 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -14.55 DPS) [crafted]; Gilded Sandals (254107, -18.77 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 56.6 | yes | Ring of Forlorn Spirits (2043, -17.02 DPS) [quest]; Reedknot Ring (9622, -17.37 DPS) [quest]; Sea Giant's Toe Ring (274746, -17.72 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Lorekeeper's Ring (19524, -1.16 DPS, sim-verified) [rep]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest]; Reedknot Ring (9622, -1.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 50.9 | yes | Thunderbrew's Boot Flask (744, -17.84 DPS) [quest]; Tidal Charm (1404, -17.84 DPS) [vendor]; Guardian Talisman (1490, -17.84 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Thunderbrew's Boot Flask (744, -2.10 DPS) [quest]; Tidal Charm (1404, -2.10 DPS) [vendor]; Guardian Talisman (1490, -2.10 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 66.1 | yes | Illusionary Rod (7713, -0.39 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -0.45 DPS) [dungeon]; Spellshifter Rod (9527, -22.29 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Dreambough Wand (249234, -0.39 DPS, sim-verified) [crafted]; Burning Sliver (5249, -0.70 DPS) [quest]; Twisted Nether Wand (249144, -0.70 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Lesser Eternal Wand

No-known-source sample (15 of 806, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 499.9. Weights run: 0.7s. Verify run: 0.7s. 1229 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.137 ± 0.028, crit=6.853 ± 0.198, hit=8.396 ± 0.345, spell_haste=4.452 ± 0.152, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 314.0 | yes | Bloodvine Goggles (19999, -17.58 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -34.27 DPS) [quest]; Field Marshal's Coronet (16441, -64.02 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 179.9 | yes | Medallion of the Dawn (22659, -29.42 DPS) [quest]; Beads of Ogre Might (22150, -33.61 DPS) [quest]; Charm of the Shifting Sands (21504, -53.69 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 114.7 | yes | Lieutenant Commander's Silk Mantle (23319, -0.80 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.80 DPS) [pvp]; Champion's Silk Mantle (23264, -1.04 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 95.9 | yes | Earthweave Cloak (21187, -0.43 DPS, sim-verified) [quest]; Drape of Vaulted Secrets (21415, -26.88 DPS) [quest]; Hide of the Wild (18510, -28.23 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 230.6 | yes | Bloodvine Vest (19682, -10.37 DPS, sim-verified) [crafted]; Enigma Robes (21343, -32.41 DPS) [quest]; Robe of the Archmage (14152, -32.59 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 111.0 | yes | Frostfire Bindings (22503, -6.12 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -30.78 DPS) [rep]; Dryad's Wrist Bindings (19596, -31.58 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 202.3 | yes | Sorcerer's Gloves (22066, -2.95 DPS, sim-verified) [quest]; Dreadmist Wraps (16705, -41.01 DPS) [dungeon]; Frostfire Gloves (22501, -57.34 DPS) [quest] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 118.1 | yes | Frostfire Belt (22502, +3.46 DPS, sim-verified) [quest]; Highlander's Cloth Girdle (20047, -2.58 DPS) [rep]; Highlander's Cloth Girdle (20097, -4.38 DPS) [rep] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 133.5 | yes | Marshal's Silk Leggings (16442, -1.70 DPS) [vendor]; General's Silk Trousers (16534, -1.70 DPS) [vendor]; Enigma Leggings (21346, -8.76 DPS, sim-verified) [quest] |
| feet | Frostfire Sandals (22500) | Frostfire Sandals [quest] | 126.4 | yes | Enigma Boots (21344, -0.64 DPS, sim-verified) [quest]; Marshal's Silk Footwraps (16437, -6.84 DPS) [vendor]; General's Silk Boots (16539, -6.84 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 179.9 | yes | Band of Earthen Might (21182, -11.36 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -20.66 DPS) [world]; Mindtear Band (20632, -21.42 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 121.8 | yes | Ritssyn's Ring of Chaos (21836, -0.30 DPS) [world]; Mindtear Band (20632, -1.06 DPS) [world]; Band of Earthen Might (21182, -12.50 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 75.6 | yes | Thunderbrew's Boot Flask (744, -26.47 DPS) [quest]; Tidal Charm (1404, -26.47 DPS) [vendor]; Guardian Talisman (1490, -26.47 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Thunderbrew's Boot Flask (744, -2.10 DPS) [quest]; Tidal Charm (1404, -2.10 DPS) [vendor]; Guardian Talisman (1490, -2.10 DPS) [quest] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 276.0 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Ironbark Staff (20069, -29.00 DPS) [rep]; Atiesh, Greatstaff of the Guardian (22631, -53.32 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 10.0 | yes | Greater Eternal Wand (249237, -0.52 DPS, sim-verified) [crafted]; Lesser Eternal Wand (249232, -0.69 DPS) [crafted]; Dreambough Wand (249234, -1.04 DPS) [crafted] |

**New at 60:** head: Frostfire Circlet; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Chromatic Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Belt of the Archmage; legs: Frostfire Leggings; feet: Frostfire Sandals; finger1: Don Julio's Band; finger2: Ring of the Fallen God; main_hand: High Warlord's War Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1229, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 29.3. Weights run: 0.5s. Verify run: 0.5s. 249 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=not significant (-0.108 ± 0.264), crit=2.133 ± 0.074, hit=5.698 ± 0.181, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.49 DPS) [crafted]; Lucky Fishing Hat (19972, -0.49 DPS) [quest]; Flying Tiger Goggles (4368, -0.94 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.15 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.41 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Battle Healer's Cloak (20427, -0.08 DPS) [rep]; Feyscale Cloak (6632, -0.12 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.08 DPS) [crafted]; Bloody Apron (6226, -0.08 DPS) [dungeon]; Green Woolen Vest (2582, -0.90 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Silver-lined Bracers (3224, +0.00 DPS) [world]; Seer's Cuffs (3645, +0.00 DPS) [world_drop]; Owlbeard Bracers (16981, -0.56 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.08 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.24 DPS) [quest]; Pristine Gloves (253913, -0.24 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 4.0 | yes | Novice Ardent's Sash (253887, -0.16 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.30 DPS, sim-verified) [crafted]; Lesser Belt of the Spire (1299, -0.33 DPS) [world] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.24 DPS) [crafted]; Colorful Kilt (10048, -0.33 DPS) [crafted]; Silk-threaded Trousers (1929, -0.71 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Feather Padded Treads (285345, -0.22 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.24 DPS) [crafted]; Pristine Boots (253889, -0.33 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Ring of the Shadow (1462, -0.41 DPS) [world]; Ring of Scorn (3235, -0.41 DPS) [quest]; Sludge-Stained Band (286535, -0.43 DPS, sim-verified) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 0.0 | yes | Ring of the Shadow (1462, +0.00 DPS) [world]; Ring of Scorn (3235, +0.00 DPS) [quest]; Sludge-Stained Band (286535, -0.32 DPS, sim-verified) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | - | - |  |  |  |
| off_hand | Dwarven Tome (279898) | Important Heirlooms [quest] | 5.0 | yes | Grayson's Torch (1172, -0.37 DPS, sim-verified) [quest]; Nightglow Concoction (3451, -0.41 DPS) [quest]; Pulsating Hydra Heart (5183, -0.41 DPS) [world] |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Torchlight Wand (5240, -0.24 DPS) [quest]; Greater Magic Wand (11288, -0.24 DPS) [crafted]; Cookie's Stirring Rod (5198, -0.29 DPS, sim-verified) [dungeon] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Notched Shortsword; off_hand: Dwarven Tome; ranged: Sizzle Stick

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 50.0. Weights run: 0.6s. Verify run: 0.6s. 473 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.157 ± 0.128), crit=1.917 ± 0.062, hit=4.834 ± 0.199, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Embalmed Shroud (7691, -0.35 DPS) [dungeon]; Filigreed Pristine Circlet (253975, -0.35 DPS) [crafted]; Silk Headband (7050, -0.84 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.9 | yes | Crystal Starfire Medallion (5003, -0.84 DPS) [world_drop]; Pendant of Myzrael (4614, -0.92 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.02 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 10.4 | yes | Invoker's Mantle (215365, -0.30 DPS) [crafted]; Death Speaker Mantle (6685, -0.31 DPS) [dungeon]; Chestnut Mantle (17695, -0.63 DPS, sim-verified) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.08 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.12 DPS) [crafted]; Battle Healer's Cloak (19529, -0.12 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 11.0 | yes | Death Speaker Robes (6682, -0.27 DPS) [dungeon]; High Robe of the Adjudicator (3461, -0.32 DPS) [quest]; Tree Bark Jacket (1486, -0.97 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.93 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.93 DPS) [quest]; Owlbeard Bracers (16981, -1.47 DPS, sim-verified) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 6.8 | yes | Gnoll Casting Gloves (892, -0.09 DPS) [world]; Truefaith Gloves (7049, -0.15 DPS) [crafted]; Serpent Gloves (5970, -0.65 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.5 | yes | Warsong Sash (16975, -0.07 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Ghamoo-ra's Bind (6908, -0.40 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -0.09 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.45 DPS) [crafted]; Silk-threaded Trousers (1929, -0.58 DPS) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.1 | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Boots of the Enchanter (4325, -0.36 DPS) [crafted]; Spidersilk Boots (4320, -1.11 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.23 DPS) [rep]; Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.35 DPS) [world]; Sacred Band (6669, -0.46 DPS) [quest]; Electrocutioner Lagnut (9447, -0.61 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 1.7 | yes | Gnarled Necromancer's Staff (251534, -0.02 DPS) [quest]; Channeler's Staff (4437, -0.05 DPS) [world]; Twisted Chanter's Staff (890, -0.21 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Eventide (5214) (or Sizzle Stick (8071), Greater Mystic Wand (217287)) | World drop [world_drop] | 5.0 | yes | Greater Mystic Wand (217287, +0.00 DPS) [crafted]; Cookie's Stirring Rod (5198, -0.23 DPS) [dungeon]; Sizzle Stick (8071, -0.64 DPS, sim-verified) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Wand of Eventide

No-known-source sample (15 of 473, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (orc, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 173.3. Weights run: 0.7s. Verify run: 0.6s. 639 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.091 ± 0.015, crit=3.159 ± 0.088, hit=3.864 ± 0.105, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -2.88 DPS, sim-verified) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop]; Silk Headband (7050, -4.11 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 | yes | Necklace of Calisea (1714, -2.32 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -2.37 DPS) [vendor]; Darkspear Warding Pendant (272075, -2.43 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 | yes | Green Silken Shoulders (7057, -0.25 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Chestnut Mantle (17695, -0.62 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, -0.10 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Battle Healer's Cloak (19528, -0.34 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 | yes | Dreamweave Vest (10021, -1.21 DPS, sim-verified) [crafted]; Robe of Power (7054, -2.56 DPS) [crafted]; Tree Bark Jacket (1486, -3.27 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Condor Bracers (15864, -0.67 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.46 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.20 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.4 | yes | Star Belt (4329, -0.53 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -1.06 DPS) [rep]; Warsong Sash (16975, -1.15 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 | yes | Gaze Dreamer Pants (6903, -1.22 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.03 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -3.13 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Acidic Walkers (9454, -6.26 DPS) [dungeon] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.69 DPS) [rep]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor]; Advisor's Ring (20426, -1.37 DPS) [rep] |
| finger2 | Reedknot Ring (9622) | Jarl Needs a Blade [quest] | 7.0 | yes | Sea Giant's Toe Ring (274746, +1.68 DPS, sim-verified) [vendor]; Electrocutioner Lagnut (9447, -1.37 DPS) [dungeon]; Sludge-Stained Band (286535, -1.37 DPS) [world] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 44.8 | yes | Windweaver Staff (7757, -1.12 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -15.00 DPS) [world_drop]; Glimmering Staff (249392, -15.00 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 6.0 | yes | Wand of Eventide (5214, -0.34 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.34 DPS) [quest]; Greater Mystic Wand (217287, -0.34 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; finger2: Reedknot Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Twisted Nether Wand

No-known-source sample (15 of 639, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (orc, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 245.3. Weights run: 0.7s. Verify run: 0.6s. 802 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.106 ± 0.020, crit=4.628 ± 0.134, hit=5.657 ± 0.238, spell_haste=3.000 ± 0.103, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 80.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -4.64 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.59 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 | yes | Mindburst Medallion (11196, -0.37 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.15 DPS) [world]; Darkspear Warding Pendant (272073, -2.34 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 73.7 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -20.61 DPS) [dungeon]; Black Mageweave Shoulders (10027, -22.00 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 | yes | Deep Woodlands Cloak (19121, -0.62 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.68 DPS) [crafted]; Nightfall Drape (12465, -1.98 DPS) [world] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 78.0 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -19.38 DPS) [world_drop]; Acumen Robes (17775, -19.91 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, +2.16 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest]; Bloodband Bracers (11469, -1.07 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.38 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.57 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.57 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 74.3 | yes | Satyrmane Sash (17755, -0.10 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -20.98 DPS) [rep]; Ghostweave Cord (254073, -21.13 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 78.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -20.69 DPS) [crafted]; Red Mageweave Pants (10009, -22.00 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 65.5 | yes | Sergeant Major's Dreadweave Boots (220891, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -14.55 DPS) [crafted]; Gilded Sandals (254107, -18.77 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 56.6 | yes | Reedknot Ring (9622, -17.37 DPS) [quest]; Sea Giant's Toe Ring (274746, -17.72 DPS) [vendor]; Electrocutioner Lagnut (9447, -18.77 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Advisor's Ring (19520, -1.12 DPS, sim-verified) [rep]; Reedknot Ring (9622, -1.75 DPS) [quest]; Advisor's Ring (19521, -1.75 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 50.9 | yes | Uther's Strength (11302, -0.15 DPS, sim-verified) [world]; Tidal Charm (1404, -17.84 DPS) [vendor]; Guardian Talisman (1490, -17.84 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 39.6 | yes | Uther's Strength (11302, +0.36 DPS, sim-verified) [world]; Tidal Charm (1404, -13.87 DPS) [vendor]; Guardian Talisman (1490, -13.87 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 66.1 | yes | Illusionary Rod (7713, -0.37 DPS, sim-verified) [dungeon]; Inventor's Focal Sword (17719, -0.45 DPS) [dungeon]; Spellshifter Rod (9527, -22.29 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Dreambough Wand (249234, -0.37 DPS, sim-verified) [crafted]; Twisted Nether Wand (249144, -0.70 DPS) [crafted]; Charged Lightning Rod (11860, -0.90 DPS) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Lesser Eternal Wand

No-known-source sample (15 of 802, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 60 (orc, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 480.0. Weights run: 0.7s. Verify run: 0.6s. 1225 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.137 ± 0.028, crit=6.853 ± 0.198, hit=8.396 ± 0.345, spell_haste=4.452 ± 0.152, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 314.0 | yes | Bloodvine Goggles (19999, -16.19 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -34.27 DPS) [quest]; Field Marshal's Coronet (16441, -64.02 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 179.9 | yes | Medallion of the Dawn (22659, -29.42 DPS) [quest]; Beads of Ogre Might (22150, -33.61 DPS) [quest]; Charm of the Shifting Sands (21504, -53.69 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 114.7 | yes | Lieutenant Commander's Silk Mantle (23319, -0.80 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.80 DPS) [pvp]; Champion's Silk Mantle (23264, -1.01 DPS, sim-verified) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 95.9 | yes | Earthweave Cloak (21187, +0.17 DPS, sim-verified) [quest]; Drape of Vaulted Secrets (21415, -26.88 DPS) [quest]; Hide of the Wild (18510, -28.23 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 230.6 | yes | Bloodvine Vest (19682, -9.32 DPS, sim-verified) [crafted]; Enigma Robes (21343, -32.41 DPS) [quest]; Robe of the Archmage (14152, -32.59 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 111.0 | yes | Frostfire Bindings (22503, -4.87 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -30.78 DPS) [rep]; Dryad's Wrist Bindings (19596, -31.58 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 202.3 | yes | Sorcerer's Gloves (22066, -2.24 DPS, sim-verified) [quest]; Dreadmist Wraps (16705, -41.01 DPS) [dungeon]; Frostfire Gloves (22501, -57.34 DPS) [quest] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 118.1 | yes | Frostfire Belt (22502, +3.91 DPS, sim-verified) [quest]; Defiler's Cloth Girdle (20163, -2.58 DPS) [rep]; Defiler's Cloth Girdle (20165, -4.38 DPS) [rep] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 133.5 | yes | Marshal's Silk Leggings (16442, -1.70 DPS) [vendor]; General's Silk Trousers (16534, -1.70 DPS) [vendor]; Enigma Leggings (21346, -7.43 DPS, sim-verified) [quest] |
| feet | Frostfire Sandals (22500) | Frostfire Sandals [quest] | 126.4 | yes | Enigma Boots (21344, -0.11 DPS, sim-verified) [quest]; Marshal's Silk Footwraps (16437, -6.84 DPS) [vendor]; General's Silk Boots (16539, -6.84 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 179.9 | yes | Band of Earthen Might (21182, -9.95 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -20.66 DPS) [world]; Mindtear Band (20632, -21.42 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 121.8 | yes | Ritssyn's Ring of Chaos (21836, -0.30 DPS) [world]; Mindtear Band (20632, -1.06 DPS) [world]; Band of Earthen Might (21182, -11.97 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 75.6 | yes | Rune of the Guard Captain (19120, -5.88 DPS) [quest]; Tidal Charm (1404, -26.47 DPS) [vendor]; Guardian Talisman (1490, -26.47 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Rune of the Guard Captain (19120, -0.16 DPS, sim-verified) [quest]; Tidal Charm (1404, -2.10 DPS) [vendor]; Guardian Talisman (1490, -2.10 DPS) [quest] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 276.0 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Ironbark Staff (20220, -29.00 DPS) [rep]; Atiesh, Greatstaff of the Guardian (22631, -53.32 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 10.0 | yes | Greater Eternal Wand (249237, -0.50 DPS, sim-verified) [crafted]; Lesser Eternal Wand (249232, -0.69 DPS) [crafted]; Dreambough Wand (249234, -1.04 DPS) [crafted] |

**New at 60:** head: Frostfire Circlet; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Chromatic Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Belt of the Archmage; legs: Frostfire Leggings; feet: Frostfire Sandals; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket2: Uther's Strength; main_hand: High Warlord's War Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1225, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

