# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 22.4. Weights run: 0.9s. Verify run: 1.0s. 122 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=-3.036 ± 0.220, crit=0.729 ± 0.038, hit=1.171 ± 0.150, spell_haste=not significant (-0.287 ± 0.238), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.54 DPS) [crafted]; Lucky Fishing Hat (19972, -0.54 DPS) [quest]; Flying Tiger Goggles (4368, -1.27 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) (or Magician's Mantle (12998)) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.09 DPS) [crafted]; Magician's Mantle (12998, -0.20 DPS, sim-verified) [world_drop]; Slime-encrusted Pads (6461, -0.45 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, +0.00 DPS, sim-verified) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Caretaker's Cape (20428, -0.09 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.09 DPS) [crafted]; Bloody Apron (6226, -0.09 DPS) [dungeon]; Green Woolen Vest (2582, -0.19 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 | yes | Silver-lined Bracers (3224, -0.09 DPS) [world]; Seer's Cuffs (3645, -0.09 DPS) [world_drop]; Mindthrust Bracers (1974, -0.47 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.27 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.45 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (22.0 DPS) | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.26 DPS, sim-verified) [crafted]; Lesser Belt of the Spire (1299, -0.36 DPS) [world] |
| legs | Silk-threaded Trousers (1929) | Westfall: Defias Evoker [dungeon] | sim-verified (22.1 DPS) | yes | Filigreed Pristine Leggings (253937, -0.09 DPS) [crafted]; Colorful Kilt (10048, -0.18 DPS) [crafted]; Abomination Skin Leggings (23173, -0.39 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Feather Padded Treads (285345, +0.00 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.27 DPS) [crafted]; Pristine Boots (253889, -0.36 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.18 DPS) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.04 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 234.0 | yes | Staff of Westfall (2042, -0.22 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.77 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 249.2 | yes | Skycaller (12984, -0.28 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Silk-threaded Trousers; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Living Root; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 122, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes; 20434 Lorekeeper's Staff; 20443 Sentinel's Blade; 209613 Insignia of the Alliance; 209621 Insignia of the Horde

### Band 30 (gnome, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 51.5. Weights run: 0.7s. Verify run: 0.7s. 206 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.550 ± 0.331, crit=0.803 ± 0.052, hit=3.822 ± 0.217, spell_haste=not significant (-0.760 ± 0.265), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 21.5 | yes | Nightsky Cowl (4039, -0.34 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.40 DPS) [crafted]; Resilient Cap (14401, -0.40 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 16.3 | yes | Crystal Starfire Medallion (5003, -0.90 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.90 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.19 DPS, sim-verified) [vendor] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 23.1 | yes | Bloodmage Mantle (7684, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 12.4 | yes | Darkspear Raider's Cloak (272078, -0.13 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.28 DPS) [world_drop]; Repairman's Cape (9605, -0.28 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 29.2 | yes | Death Speaker Robes (6682, -0.34 DPS, sim-verified) [dungeon]; Black Velvet Robes (2800, -0.94 DPS) [world_drop]; Pristine Gown (253961, -1.00 DPS) [crafted] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 12.4 | yes | Nightsky Wristbands (6407, -0.06 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.28 DPS) [quest]; Spidertank Oilrag (9448, -0.30 DPS) [dungeon] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 21.1 | yes | Hotshot Pilot's Gloves (9491, -0.61 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.91 DPS) [quest]; Truefaith Gloves (7049, -1.01 DPS) [crafted] |
| waist | Crimson Silk Belt (7055) | Tailoring [crafted] | 16.9 | yes | Highlander's Cloth Girdle (20099, +0.00 DPS, sim-verified) [rep]; Invoker's Cord (215366, -0.19 DPS) [crafted]; Belt of Arugal (6392, -0.28 DPS) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 21.4 | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Necromancer Leggings (2277, -0.39 DPS) [world_drop]; Filigreed Pristine Leggings (253937, -0.54 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 17.9 | yes | Spidersilk Boots (4320, -0.41 DPS) [crafted]; Frothing Slippers (254003, -0.62 DPS) [crafted]; Acidic Walkers (9454, -1.20 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 10.9 | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Minor Channeling Ring (1449, -0.24 DPS) [quest]; Azora's Will (4999, -0.28 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.9 | yes | Minor Channeling Ring (1449, -0.24 DPS) [quest]; Azora's Will (4999, -0.28 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.35 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (51.5 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Insignia of the Horde (18851, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 293.5 | yes | Gnarled Ash Staff (791, -0.54 DPS, sim-verified) [world_drop]; Lorekeeper's Staff (212580, -0.58 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.58 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 377.0 | yes | Starfaller (13063, -0.36 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Death Speaker Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Glowing Magical Bracelets; hands: Town Clerk's Mittens; waist: Crimson Silk Belt; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 206, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 16606 Juju Hex Robes; 16607 Acolyte's Sacrificial Robes

### Band 40 (gnome, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 80.4. Weights run: 0.7s. Verify run: 0.7s. 285 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (-0.254 ± 0.524), crit=1.076 ± 0.084, hit=5.292 ± 0.396, spell_haste=not significant (0.657 ± 0.448), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (80.8 DPS) | yes | Holy Shroud (2721, +0.00 DPS) [world_drop]; Silk Headband (7050, -0.19 DPS) [crafted]; Spellpower Goggles Xtreme (10502, -0.87 DPS, sim-verified) [crafted] |
| neck | Necklace of Calisea (1714) | World drop [world_drop] | sim-verified (81.4 DPS) | yes | Pendant of Myzrael (4614, +0.00 DPS) [dungeon]; Glowing Green Talisman (5002, +0.00 DPS) [world_drop]; Scorn's Icy Choker (23169, -1.39 DPS, sim-verified) [dungeon] |
| shoulder | Green Silken Shoulders (7057) | Tailoring [crafted] | sim-verified (80.9 DPS) | yes | Berylline Pads (4197, -0.09 DPS) [quest]; Moonlit Amice (11884, -0.09 DPS) [quest]; Bloodmage Mantle (7684, -0.93 DPS, sim-verified) [dungeon] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, +0.00 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.09 DPS) [crafted]; Caretaker's Cape (19532, -0.09 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 | yes | Dreamweave Vest (10021, -0.38 DPS) [crafted]; Robe of Power (7054, -0.76 DPS) [crafted]; Elemental Raiment (9434, -0.85 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (81.2 DPS) | yes | Condor Bracers (15864, -0.19 DPS) [quest]; Earthen Silk Cuffs (254019, -0.47 DPS) [crafted]; Arcane Runed Bracers (4744, -1.21 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Red Mageweave Gloves (10018, -0.66 DPS) [crafted]; Gilded Handwraps (254021, -0.94 DPS) [crafted]; Black Mageweave Gloves (10003, -1.51 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.0 | yes | Star Belt (4329, +0.00 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -0.28 DPS) [rep]; Belt of Arugal (6392, -0.47 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.0 | yes | Abomination Skin Leggings (23173, -0.47 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.57 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.07 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Spidersilk Boots (4320, -0.63 DPS, sim-verified) [crafted]; Gilded Slippers (254001, -1.60 DPS) [crafted]; Nimbus Boots (6998, -1.70 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.0 | yes | Ring of Forlorn Spirits (2043, -0.19 DPS) [quest]; Reedknot Ring (9622, -0.28 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.19 DPS) [quest]; Lorekeeper's Ring (19525, -0.19 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.92 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 428.2 | yes | Illusionary Rod (7713, -1.13 DPS, sim-verified) [dungeon]; Black Duskwood Staff (937, -6.84 DPS) [world_drop]; Spiritchaser Staff (1613, -6.84 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 417.2 | yes | Umbral Wand (5216, -0.53 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -4.81 DPS) [crafted]; Ember Wand (5215, -5.71 DPS) [world_drop] |

**New at 40:** head: Augural Shroud; neck: Necklace of Calisea; shoulder: Green Silken Shoulders; back: Icy Cloak; chest: Robe of the Magi; wrist: Spidertank Oilrag; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 285, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 50 (gnome, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 124.3. Weights run: 0.7s. Verify run: 0.7s. 359 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.962 ± 0.656), crit=1.607 ± 0.120, hit=8.647 ± 0.553, spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 60.0 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.03 DPS) [dungeon]; Red Mageweave Headband (10033, -0.19 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 27.5 | yes | Scorn's Icy Choker (23169, -0.55 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -1.01 DPS) [quest]; Darkspear Warding Pendant (272073, -1.02 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 48.3 | yes | Blood Guard's Dreadweave Mantle (220905, -0.02 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.24 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.02 DPS, sim-verified) [vendor] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 27.5 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.29 DPS) [crafted]; Big Voodoo Cloak (8216, -0.50 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 58.2 | yes | Stone Guard's Dreadweave Vest (220904, -0.23 DPS) [vendor]; Runecloth Robe (13858, -1.45 DPS) [crafted]; Knight's Dreadweave Vest (220886, -2.31 DPS, sim-verified) [vendor] |
| wrist | Forgotten Wraps (9433) (or Shizzle's Nozzle Wiper (11917)) | World drop [world_drop] | 23.5 | yes | Shizzle's Nozzle Wiper (11917, +0.00 DPS, sim-verified) [quest]; Bloodband Bracers (11469, -0.09 DPS) [quest]; Imperial Red Bracers (8247, -0.20 DPS) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 50.2 | yes | Virtuous Hands (226958, -0.12 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.03 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.03 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 43.3 | yes | Deathmage Sash (10771, -0.71 DPS) [dungeon]; Satyrmane Sash (17755, -1.01 DPS) [dungeon]; Highlander's Cloth Girdle (20097, -1.68 DPS, sim-verified) [rep] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 58.0 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.05 DPS) [dungeon]; Red Mageweave Pants (10009, -2.14 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 112.1 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Southsea Mojo Boots (20641, -8.60 DPS) [quest]; Gilded Sandals (254107, -8.70 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 86.5 | yes | Ogremind Ring (1993, -7.58 DPS) [world_drop]; Voodoo Band (1996, -7.58 DPS) [world_drop]; Mindbender Loop (5009, -7.58 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 21.8 | yes | Voodoo Band (1996, -0.84 DPS) [world_drop]; Mindbender Loop (5009, -0.84 DPS) [world_drop]; Ogremind Ring (1993, -1.11 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (124.3 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (124.3 DPS) | yes | Frozen Heart of the Mountain (249469, -0.13 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -0.63 DPS) [quest]; Tidal Charm (1404, -0.63 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (124.3 DPS) | yes | Soulkeeper (1607, -1.83 DPS) [world_drop]; Hammer of the Northern Wind (810, -2.13 DPS, sim-verified) [world_drop]; Radiant Staff (249453, -3.44 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.92 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Acumen Robes; wrist: Forgotten Wraps; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 359, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10047 Simple Kilt

### Band 60 (gnome, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 237.6. Weights run: 0.7s. Verify run: 0.8s. 772 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.724 ± 0.748), crit=2.882 ± 0.178, hit=12.470 ± 0.766, spell_haste=not significant (1.648 ± 0.891), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Circlet of Revelation (239585) | Leonid Barthalomew the Revered [vendor] | sim-verified (208.7 DPS) | yes | Crown of Revelation (239575, -3.92 DPS) [vendor]; Virtuous Cowl (226957, -5.73 DPS) [vendor]; Bloodvine Goggles (19999, -6.99 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (201.7 DPS) | yes | Blazefury Medallion (17111, -0.34 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -9.69 DPS) [quest]; Amulet of the Dawn (22657, -11.52 DPS) [quest] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | 152.4 | yes | Rugged Mantle of the Timbermaw (227808, +0.00 DPS, sim-verified) [vendor]; Shoulderpads of Revelation (239586, -8.61 DPS) [vendor]; Mantle of the Timbermaw (19050, -9.84 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 146.5 | yes | Howler's Furs (272414, -2.50 DPS) [vendor]; Stalwart Cloak (272415, -2.50 DPS) [vendor]; Earthweave Cloak (21187, -4.83 DPS, sim-verified) [quest] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 285.8 | yes | Garb of Revelation (239565, -3.35 DPS, sim-verified) [vendor]; Earthpower Vest (21183, -23.51 DPS) [quest]; Robe of Revelation (239591, -25.24 DPS) [vendor] |
| wrist | Bindings of Revelation (239588) | Leonid Barthalomew the Revered [vendor] | sim-verified (206.2 DPS) | yes | Wrists of Revelation (239583, -2.02 DPS) [vendor]; Rockfury Bracers (21186, -4.49 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -4.87 DPS) [rep] |
| hands | Gloves of Revelation (239584) | Leonid Barthalomew the Revered [vendor] | sim-verified (206.8 DPS) | yes | Gloves of Spell Mastery (14146, -2.25 DPS) [crafted]; Dreadmist Wraps (16705, -5.18 DPS, sim-verified) [dungeon]; Hands of Revelation (239574, -6.93 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 165.6 | yes | Belt of Revelation (239590, +0.00 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -10.77 DPS) [crafted]; Highlander's Cloth Girdle (20047, -12.29 DPS) [rep] |
| legs | Leggings of Revelation (239587) | Leonid Barthalomew the Revered [vendor] | sim-verified (208.9 DPS) | yes | Magister's Leggings (16687, -0.12 DPS) [dungeon]; Sentinel's Silk Leggings (237815, -1.29 DPS) [vendor]; Bloodvine Leggings (19683, -7.19 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 155.3 | yes | First Sergeant's Dreadweave Boots (220909, -1.85 DPS) [vendor]; Argent Elite Boots (227816, -3.51 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -3.95 DPS, sim-verified) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (201.7 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Don Julio's Band (19325, -1.42 DPS, sim-verified) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (207.5 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Signet Ring of the Bronze Dragonflight (234028, -0.31 DPS) [vendor]; Band of Earthen Might (21182, -5.84 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (200.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.92 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -3.13 DPS, sim-verified) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (200.4 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.84 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -2.23 DPS, sim-verified) [quest] |
| main_hand | High Warlord's Battle Mace (234551) | Sergeant Thunderhorn [vendor] | sim-verified (201.7 DPS) | yes | Ebon Hand (19170, +0.00 DPS, sim-verified) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | Therazane's Touch (19315) | Stormpike Guard [rep] | 31.0 | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [vendor]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [vendor]; Spirit of Aquementas (11904, -1.48 DPS, sim-verified) [quest] |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | 556.9 | yes | Stormrager (16997, -0.09 DPS, sim-verified) [quest]; Brilliant Wand (249385, -2.80 DPS) [crafted]; Torch of Austen (13004, -7.21 DPS) [world_drop] |

**New at 60:** head: Circlet of Revelation; neck: Beads of Ogre Might; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Bloodvine Vest; wrist: Bindings of Revelation; hands: Gloves of Revelation; waist: Knowledge of the Timbermaw; legs: Leggings of Revelation; feet: Bloodvine Boots; finger1: Wrath of Cenarius; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: High Warlord's Battle Mace; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 772, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10047 Simple Kilt

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 26.9. Weights run: 0.9s. Verify run: 0.9s. 123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=-3.036 ± 0.220, crit=0.729 ± 0.038, hit=1.171 ± 0.150, spell_haste=not significant (-0.287 ± 0.238), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.54 DPS) [crafted]; Lucky Fishing Hat (19972, -0.54 DPS) [quest]; Flying Tiger Goggles (4368, -0.80 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | sim-verified (26.7 DPS) | yes | Erudite's Amulet (277204, +0.00 DPS) [quest]; Tarnished Locket (279870, +0.00 DPS) [quest]; Scout's Medallion (20442, -0.30 DPS, sim-verified) [rep] |
| shoulder | Reinforced Woolen Shoulders (4315) (or Magician's Mantle (12998)) | Tailoring [crafted] | 5.0 | yes | Magician's Mantle (12998, -0.08 DPS, sim-verified) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.09 DPS) [crafted]; Slime-encrusted Pads (6461, -0.45 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, +0.00 DPS, sim-verified) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Battle Healer's Cloak (20427, -0.09 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.09 DPS) [crafted]; Bloody Apron (6226, -0.09 DPS) [dungeon]; Green Woolen Vest (2582, -0.52 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) (or Windsong Bangles (263336)) | Earthen Arise [quest] | 1.0 | yes | Windsong Bangles (263336, +0.00 DPS, sim-verified) [quest]; Mindthrust Bracers (1974, -0.09 DPS) [dungeon]; Silver-lined Bracers (3224, -0.09 DPS) [world] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.09 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.27 DPS) [quest]; Pristine Gloves (253913, -0.27 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (27.1 DPS) | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.36 DPS) [world]; Novice Arcanist's Sash (253885, -0.62 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Silk-threaded Trousers (1929, -0.28 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.36 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.27 DPS) [crafted]; Feather Padded Treads (285345, -0.34 DPS, sim-verified) [world]; Pristine Boots (253889, -0.36 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world]; Ring of Scorn (3235, -0.45 DPS) [quest] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Ring of the Shadow (1462, -0.27 DPS) [world]; Ring of Scorn (3235, -0.27 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.38 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (26.9 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Crescent Staff (6505, -0.10 DPS) [quest]; Living Root (6631, -0.42 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 249.2 | yes | Skycaller (12984, -0.71 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.29 DPS) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 123, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 209613 Insignia of the Alliance; 209621 Insignia of the Horde; 241089 Scarlet Dagger; 263007 Skyseer's Vest

### Band 30 (undead, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 47.3. Weights run: 0.7s. Verify run: 0.7s. 209 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.550 ± 0.331, crit=0.803 ± 0.052, hit=3.822 ± 0.217, spell_haste=not significant (-0.760 ± 0.265), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 21.5 | yes | Shadow Hood (4323, -0.40 DPS) [crafted]; Resilient Cap (14401, -0.40 DPS) [world_drop]; Nightsky Cowl (4039, -0.44 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 16.3 | yes | Crystal Starfire Medallion (5003, -0.90 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.90 DPS) [world_drop]; Darkspear Warding Pendant (272075, -0.93 DPS, sim-verified) [vendor] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 23.1 | yes | Bloodmage Mantle (7684, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.36 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 12.4 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.28 DPS) [world_drop]; Soft Willow Cape (16661, -0.41 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 29.2 | yes | Death Speaker Robes (6682, -0.50 DPS, sim-verified) [dungeon]; Black Velvet Robes (2800, -0.94 DPS) [world_drop]; Pristine Gown (253961, -1.00 DPS) [crafted] |
| wrist | Glowing Magical Bracelets (13106) | World drop [world_drop] | 12.4 | yes | Nightsky Wristbands (6407, -0.18 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -0.28 DPS) [quest]; Spidertank Oilrag (9448, -0.30 DPS) [dungeon] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 13.8 | yes | Blight Gloves (279877, -0.26 DPS) [quest]; Truefaith Gloves (7049, -0.36 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.49 DPS, sim-verified) [dungeon] |
| waist | Crimson Silk Belt (7055) | Tailoring [crafted] | 16.9 | yes | Defiler's Cloth Girdle (20164, +0.00 DPS, sim-verified) [rep]; Invoker's Cord (215366, -0.19 DPS) [crafted]; Lilac Sash (6780, -0.26 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (47.3 DPS) | yes | Necromancer Leggings (2277, -0.07 DPS) [world_drop]; Filigreed Pristine Leggings (253937, -0.23 DPS) [crafted]; Abomination Skin Leggings (23173, -0.61 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 17.9 | yes | Spidersilk Boots (4320, -0.41 DPS) [crafted]; Frothing Slippers (254003, -0.62 DPS) [crafted]; Acidic Walkers (9454, -0.92 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 10.9 | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Azora's Will (4999, -0.28 DPS) [world_drop]; Loop of Sacrifice (281673, -0.28 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.9 | yes | Azora's Will (4999, -0.28 DPS) [world_drop]; Loop of Sacrifice (281673, -0.28 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.37 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (46.4 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Insignia of the Horde (18851, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 293.5 | yes | Gnarled Ash Staff (791, -0.56 DPS, sim-verified) [world_drop]; Lorekeeper's Staff (212580, -0.58 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.58 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 377.0 | yes | Starfaller (13063, -0.29 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Death Speaker Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Glowing Magical Bracelets; hands: Jutebraid Gloves; waist: Crimson Silk Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 16604 Moon Robes of Elune; 16605 Friar's Robes of the Light; 19545 Scout's Blade

### Band 40 (undead, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 72.4. Weights run: 0.7s. Verify run: 0.7s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (-0.254 ± 0.524), crit=1.076 ± 0.084, hit=5.292 ± 0.396, spell_haste=not significant (0.657 ± 0.448), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -0.94 DPS) [world_drop]; Silk Headband (7050, -1.13 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Necklace of Calisea (1714, +0.00 DPS, sim-verified) [world_drop]; Ethereal Talisman (4430, -0.66 DPS) [quest]; Pendant of Myzrael (4614, -0.66 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Chestnut Mantle (17695, -0.09 DPS) [quest]; Berylline Pads (4197, -0.19 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, +0.00 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.09 DPS) [crafted]; Battle Healer's Cloak (19528, -0.09 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 | yes | Dreamweave Vest (10021, -0.38 DPS) [crafted]; Robe of Power (7054, -0.76 DPS) [crafted]; Elemental Raiment (9434, -1.58 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Radiant Silver Bracers (4545, -0.47 DPS) [quest]; Earthen Silk Cuffs (254019, -0.47 DPS) [crafted]; Condor Bracers (15864, -1.58 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Red Mageweave Gloves (10018, -0.66 DPS) [crafted]; Gilded Handwraps (254021, -0.94 DPS) [crafted]; Black Mageweave Gloves (10003, -1.86 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.0 | yes | Warsong Sash (16975, -0.28 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.28 DPS) [rep]; Star Belt (4329, -0.50 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.0 | yes | Abomination Skin Leggings (23173, -0.47 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.57 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.00 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -1.60 DPS) [crafted]; Boots of the Enchanter (4325, -1.79 DPS) [crafted]; Spidersilk Boots (4320, -1.89 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.0 | yes | Reedknot Ring (9622, -0.28 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.19 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.28 DPS) [vendor]; Reedknot Ring (9622, -1.58 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (72.5 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Rune of Duty (21567, -0.89 DPS, sim-verified) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 428.2 | yes | Illusionary Rod (7713, -1.32 DPS, sim-verified) [dungeon]; Black Duskwood Staff (937, -6.84 DPS) [world_drop]; Spiritchaser Staff (1613, -6.84 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 417.2 | yes | Umbral Wand (5216, -1.57 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -4.81 DPS) [crafted]; Ember Wand (5215, -5.71 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Bloodmage Mantle; back: Icy Cloak; chest: Robe of the Magi; wrist: Spidertank Oilrag; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (undead, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 112.7. Weights run: 0.7s. Verify run: 0.7s. 362 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.962 ± 0.656), crit=1.607 ± 0.120, hit=8.647 ± 0.553, spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 60.0 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.03 DPS) [dungeon]; Red Mageweave Headband (10033, -0.19 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 27.5 | yes | Mindburst Medallion (11196, -1.01 DPS) [quest]; Darkspear Warding Pendant (272073, -1.02 DPS) [vendor]; Scorn's Icy Choker (23169, -1.95 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 48.3 | yes | Blood Guard's Dreadweave Mantle (220905, -0.02 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.24 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -3.04 DPS, sim-verified) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 29.7 | yes | Darkspear Raider's Cloak (272076, -0.01 DPS, sim-verified) [vendor]; Spritecaster Cape (11623, -0.40 DPS) [dungeon]; Runecloth Cloak (13860, -0.52 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 58.2 | yes | Stone Guard's Dreadweave Vest (220904, -0.23 DPS) [vendor]; Runecloth Robe (13858, -1.45 DPS) [crafted]; Knight's Dreadweave Vest (220886, -3.28 DPS, sim-verified) [vendor] |
| wrist | Forgotten Wraps (9433) (or Shizzle's Nozzle Wiper (11917)) | World drop [world_drop] | 23.5 | yes | Shizzle's Nozzle Wiper (11917, +0.00 DPS, sim-verified) [quest]; Bloodband Bracers (11469, -0.09 DPS) [quest]; Imperial Red Bracers (8247, -0.20 DPS) [dungeon] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 50.2 | yes | Virtuous Hands (226958, -1.27 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.03 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.03 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 43.3 | yes | Deathmage Sash (10771, -0.71 DPS) [dungeon]; Satyrmane Sash (17755, -1.01 DPS) [dungeon]; Defiler's Cloth Girdle (20165, -4.35 DPS, sim-verified) [rep] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 58.0 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.05 DPS) [dungeon]; Red Mageweave Pants (10009, -2.14 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 112.1 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Southsea Mojo Boots (20641, -8.60 DPS) [quest]; Gilded Sandals (254107, -8.70 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 86.5 | yes | Ogremind Ring (1993, -7.58 DPS) [world_drop]; Voodoo Band (1996, -7.58 DPS) [world_drop]; Mindbender Loop (5009, -7.58 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 21.8 | yes | Voodoo Band (1996, -0.84 DPS) [world_drop]; Mindbender Loop (5009, -0.84 DPS) [world_drop]; Ogremind Ring (1993, -1.01 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (112.2 DPS) | yes | Uther's Strength (11302, -0.31 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -8.11 DPS) [vendor]; Guardian Talisman (1490, -8.11 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (112.2 DPS) | yes | Uther's Strength (11302, -0.24 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -6.31 DPS) [vendor]; Guardian Talisman (1490, -6.31 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (112.2 DPS) | yes | Soulkeeper (1607, -1.83 DPS) [world_drop]; Hammer of the Northern Wind (810, -3.43 DPS, sim-verified) [world_drop]; Radiant Staff (249453, -3.44 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.92 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Forgotten Wraps; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 362, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10047 Simple Kilt

### Band 60 (undead, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 217.2. Weights run: 0.7s. Verify run: 0.8s. 777 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.724 ± 0.748), crit=2.882 ± 0.178, hit=12.470 ± 0.766, spell_haste=not significant (1.648 ± 0.891), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Circlet of Revelation (239585) | Leonid Barthalomew the Revered [vendor] | sim-verified (191.8 DPS) | yes | Crown of Revelation (239575, -3.92 DPS) [vendor]; Virtuous Cowl (226957, -5.73 DPS) [vendor]; Bloodvine Goggles (19999, -7.15 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (184.7 DPS) | yes | Blazefury Medallion (17111, -0.24 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -9.69 DPS) [quest]; Amulet of the Dawn (22657, -11.52 DPS) [quest] |
| shoulder | Virtuous Epaulets (226955) | Mokvar [vendor] | 152.4 | yes | Rugged Mantle of the Timbermaw (227808, +0.00 DPS, sim-verified) [vendor]; Shoulderpads of Revelation (239586, -8.61 DPS) [vendor]; Mantle of the Timbermaw (19050, -9.84 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 146.5 | yes | Howler's Furs (272414, -2.50 DPS) [vendor]; Stalwart Cloak (272415, -2.50 DPS) [vendor]; Earthweave Cloak (21187, -3.77 DPS, sim-verified) [quest] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 285.8 | yes | Garb of Revelation (239565, -3.06 DPS, sim-verified) [vendor]; Earthpower Vest (21183, -23.51 DPS) [quest]; Robe of Revelation (239591, -25.24 DPS) [vendor] |
| wrist | Bindings of Revelation (239588) | Leonid Barthalomew the Revered [vendor] | sim-verified (188.0 DPS) | yes | Wrists of Revelation (239583, -2.02 DPS) [vendor]; Rockfury Bracers (21186, -3.40 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -4.87 DPS) [rep] |
| hands | Gloves of Revelation (239584) | Leonid Barthalomew the Revered [vendor] | sim-verified (189.5 DPS) | yes | Gloves of Spell Mastery (14146, -2.25 DPS) [crafted]; Dreadmist Wraps (16705, -4.88 DPS, sim-verified) [dungeon]; Hands of Revelation (239574, -6.93 DPS) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 165.6 | yes | Belt of Revelation (239590, +0.00 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -10.77 DPS) [crafted]; Defiler's Cloth Girdle (20163, -12.29 DPS) [rep] |
| legs | Leggings of Revelation (239587) | Leonid Barthalomew the Revered [vendor] | sim-verified (190.5 DPS) | yes | Magister's Leggings (16687, -0.12 DPS) [dungeon]; Sentinel's Silk Leggings (237815, -1.29 DPS) [vendor]; Bloodvine Leggings (19683, -5.90 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 155.3 | yes | First Sergeant's Dreadweave Boots (220909, -1.85 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -2.54 DPS, sim-verified) [vendor]; Argent Elite Boots (227816, -3.51 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (184.7 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Don Julio's Band (19325, -1.31 DPS, sim-verified) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (190.4 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Signet Ring of the Bronze Dragonflight (234028, -0.31 DPS) [vendor]; Band of Earthen Might (21182, -5.81 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (180.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.92 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (182.9 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.09 DPS, sim-verified) [world_drop] |
| main_hand | High Warlord's Battle Mace (234551) | Sergeant Thunderhorn [vendor] | sim-verified (184.7 DPS) | yes | Ebon Hand (19170, +0.00 DPS, sim-verified) [crafted]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | Therazane's Touch (19315) | Frostwolf Clan [rep] | 31.0 | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [vendor]; Grand Marshal's Tome of Power (234589, +0.00 DPS) [vendor]; Spirit of Aquementas (11904, -1.31 DPS, sim-verified) [quest] |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | 556.9 | yes | Stormrager (16997, -0.30 DPS, sim-verified) [quest]; Brilliant Wand (249385, -2.80 DPS) [crafted]; Torch of Austen (13004, -7.21 DPS) [world_drop] |

**New at 60:** head: Circlet of Revelation; neck: Beads of Ogre Might; shoulder: Virtuous Epaulets; back: Arcanoweave Cloak; chest: Bloodvine Vest; wrist: Bindings of Revelation; hands: Gloves of Revelation; waist: Knowledge of the Timbermaw; legs: Leggings of Revelation; feet: Bloodvine Boots; finger1: Wrath of Cenarius; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: High Warlord's Battle Mace; off_hand: Therazane's Touch; ranged: Wand of Biting Cold

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 9653 Speedy Racer Goggles; 10047 Simple Kilt

