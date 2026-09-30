# Leveling BiS: Shadow

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 22.4. Weights run: 0.9s. Verify run: 0.9s. 202 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=-3.036 ± 0.220, crit=0.729 ± 0.038, hit=1.171 ± 0.150, spell_haste=not significant (-0.287 ± 0.238), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.54 DPS) [crafted]; Lucky Fishing Hat (19972, -0.54 DPS) [quest]; Flying Tiger Goggles (4368, -1.27 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, +0.00 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.45 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, +0.00 DPS, sim-verified) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Caretaker's Cape (20428, -0.09 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.09 DPS) [crafted]; Bloody Apron (6226, -0.09 DPS) [dungeon]; Green Woolen Vest (2582, -0.19 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 | yes | Silver-lined Bracers (3224, -0.09 DPS) [world]; Seer's Cuffs (3645, -0.09 DPS) [world_drop]; Mindthrust Bracers (1974, -0.47 DPS, sim-verified) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.10 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.27 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.45 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 0.0 | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.26 DPS, sim-verified) [crafted]; Lesser Belt of the Spire (1299, -0.36 DPS) [world] |
| legs | Silk-threaded Trousers (1929) | Westfall: Defias Evoker [dungeon] | 0.0 | yes | Filigreed Pristine Leggings (253937, -0.09 DPS) [crafted]; Colorful Kilt (10048, -0.18 DPS) [crafted]; Abomination Skin Leggings (23173, -0.39 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Feather Padded Treads (285345, +0.00 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.27 DPS) [crafted]; Pristine Boots (253889, -0.36 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.18 DPS) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.04 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 234.0 | yes | Staff of Westfall (2042, -0.22 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.77 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 249.2 | yes | Firebelcher (5243, -0.22 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Silk-threaded Trousers; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Living Root; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 202, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak; 9792 Ivycloth Boots

### Band 30 (gnome, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 52.2. Weights run: 0.7s. Verify run: 0.7s. 405 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.550 ± 0.331, crit=0.803 ± 0.052, hit=3.822 ± 0.217, spell_haste=not significant (-0.760 ± 0.265), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 21.5 | yes | Shadow Hood (4323, -0.40 DPS) [crafted]; Nightsky Cowl (4039, -0.58 DPS, sim-verified) [world_drop]; Cloudy Stormsewn Cowl (277046, -0.67 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 16.3 | yes | Darkspear Warding Pendant (272075, -0.71 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.90 DPS) [world_drop]; Pendant of Myzrael (4614, -1.45 DPS) [dungeon] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 23.1 | yes | Bloodmage Mantle (7684, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Batwing Mantle (6697, -0.53 DPS) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 12.4 | yes | Darkspear Raider's Cloak (272078, -0.17 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.28 DPS) [quest]; Pearl-clasped Cloak (5542, -0.51 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 29.2 | yes | Death Speaker Robes (6682, -0.47 DPS, sim-verified) [dungeon]; Black Velvet Robes (2800, -0.94 DPS) [world_drop]; Pristine Gown (253961, -1.00 DPS) [crafted] |
| wrist | Nightsky Wristbands (6407) (or Tabitha's Cuffs (251486)) | World drop [world_drop] | 9.3 | yes | Tabitha's Cuffs (251486, -0.02 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.03 DPS) [dungeon]; Mindthrust Bracers (1974, -0.14 DPS) [dungeon] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 21.1 | yes | Hotshot Pilot's Gloves (9491, -0.83 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.91 DPS) [quest]; Truefaith Gloves (7049, -1.01 DPS) [crafted] |
| waist | Crimson Silk Belt (7055) | Tailoring [crafted] | 16.9 | yes | Highlander's Cloth Girdle (20099, +0.00 DPS, sim-verified) [rep]; Invoker's Cord (215366, -0.19 DPS) [crafted]; Belt of Arugal (6392, -0.28 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 0.0 | yes | Filigreed Pristine Leggings (253937, -0.23 DPS) [crafted]; Night Watch Pantaloons (2954, -0.48 DPS) [quest]; Abomination Skin Leggings (23173, -0.76 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 17.9 | yes | Spidersilk Boots (4320, -0.41 DPS) [crafted]; Frothing Slippers (254003, -0.62 DPS) [crafted]; Acidic Walkers (9454, -1.12 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 10.9 | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Minor Channeling Ring (1449, -0.24 DPS) [quest]; Azora's Will (4999, -0.28 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.9 | yes | Lavishly Jeweled Ring (1156, -0.03 DPS, sim-verified) [dungeon]; Minor Channeling Ring (1449, -0.24 DPS) [quest]; Azora's Will (4999, -0.28 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 293.5 | yes | Gnarled Ash Staff (791, -0.33 DPS, sim-verified) [world_drop]; Lorekeeper's Staff (212580, -0.58 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.58 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 377.0 | yes | Greater Mystic Wand (217287, -0.69 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest]; Scorching Wand (5213, -4.63 DPS) [world_drop] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Death Speaker Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Nightsky Wristbands; hands: Town Clerk's Mittens; waist: Crimson Silk Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 405, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 40 (gnome, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 79.4. Weights run: 0.7s. Verify run: 0.7s. 567 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (-0.254 ± 0.524), crit=1.076 ± 0.084, hit=5.292 ± 0.396, spell_haste=not significant (0.657 ± 0.448), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -0.94 DPS) [world_drop]; Silk Headband (7050, -1.13 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Pendant of Myzrael (4614, -0.66 DPS) [dungeon]; Glowing Green Talisman (5002, -0.66 DPS) [world_drop]; Necklace of Calisea (1714, -1.27 DPS, sim-verified) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 | yes | Berylline Pads (4197, -0.19 DPS) [quest]; Moonlit Amice (11884, -0.19 DPS) [quest]; Green Silken Shoulders (7057, -0.91 DPS, sim-verified) [crafted] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Guardian Cloak (5965, -0.09 DPS) [crafted]; Caretaker's Cape (19532, -0.09 DPS) [rep]; Long Silken Cloak (4326, -0.84 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 | yes | Dreamweave Vest (10021, -0.71 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.76 DPS) [crafted]; Tree Bark Jacket (1486, -0.85 DPS) [dungeon] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Condor Bracers (15864, -0.19 DPS) [quest]; Spidertank Oilrag (9448, -0.46 DPS, sim-verified) [dungeon]; Earthen Silk Cuffs (254019, -0.47 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Red Mageweave Gloves (10018, -0.66 DPS) [crafted]; Gilded Handwraps (254021, -0.94 DPS) [crafted]; Black Mageweave Gloves (10003, -2.77 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.0 | yes | Highlander's Cloth Girdle (20099, -0.28 DPS) [rep]; Belt of Arugal (6392, -0.47 DPS) [dungeon]; Star Belt (4329, -1.86 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.0 | yes | Abomination Skin Leggings (23173, -0.47 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.57 DPS) [crafted]; Gaze Dreamer Pants (6903, -3.37 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -1.60 DPS) [crafted]; Nimbus Boots (6998, -1.70 DPS) [quest]; Spidersilk Boots (4320, -3.02 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.0 | yes | Ring of Forlorn Spirits (2043, -0.19 DPS) [quest]; Reedknot Ring (9622, -0.28 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.19 DPS) [quest]; Lorekeeper's Ring (19525, -0.19 DPS) [rep]; Ring of Forlorn Spirits (2043, -2.07 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 428.2 | yes | Illusionary Rod (7713, -2.33 DPS, sim-verified) [dungeon]; Black Duskwood Staff (937, -6.84 DPS) [world_drop]; Spiritchaser Staff (1613, -6.84 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | World drop [world_drop] | 377.9 | yes | Twisted Nether Wand (249144, +0.00 DPS, sim-verified) [crafted]; Ember Wand (5215, -2.00 DPS) [world_drop]; Necrotic Wand (7708, -2.18 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Bloodmage Mantle; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Staff of Jordan; ranged: Umbral Wand

No-known-source sample (15 of 567, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs; 7522 Gossamer Boots

### Band 50 (gnome, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 124.7. Weights run: 0.7s. Verify run: 0.7s. 722 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.962 ± 0.656), crit=1.607 ± 0.120, hit=8.647 ± 0.553, spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 60.0 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.03 DPS) [dungeon]; Red Mageweave Headband (10033, -0.19 DPS) [crafted] |
| neck | Horizon Choker (13085) | Azuregos [world] | 27.5 | yes | Scorn's Icy Choker (23169, -0.29 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -1.01 DPS) [quest]; Darkspear Warding Pendant (272073, -1.02 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 48.3 | yes | Blood Guard's Dreadweave Mantle (220905, -0.02 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.24 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.84 DPS, sim-verified) [vendor] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 27.5 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.29 DPS) [crafted]; Big Voodoo Cloak (8216, -0.50 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 58.2 | yes | Stone Guard's Dreadweave Vest (220904, -0.23 DPS) [vendor]; Runecloth Robe (13858, -1.45 DPS) [crafted]; Knight's Dreadweave Vest (220886, -2.13 DPS, sim-verified) [vendor] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 23.5 | yes | Bloodband Bracers (11469, +0.00 DPS, sim-verified) [quest]; Imperial Red Bracers (8247, -0.20 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.29 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 50.2 | yes | Virtuous Hands (226958, -0.30 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.03 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.03 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 43.3 | yes | Deathmage Sash (10771, -0.71 DPS) [dungeon]; Satyrmane Sash (17755, -1.01 DPS) [dungeon]; Highlander's Cloth Girdle (20097, -1.20 DPS, sim-verified) [rep] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 58.0 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.05 DPS) [dungeon]; Red Mageweave Pants (10009, -2.14 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 112.1 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Southsea Mojo Boots (20641, -8.60 DPS) [quest]; Gilded Sandals (254107, -8.70 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 86.5 | yes | Ogremind Ring (1993, -7.58 DPS) [world_drop]; Voodoo Band (1996, -7.58 DPS) [world]; Mindbender Loop (5009, -7.58 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 21.8 | yes | Voodoo Band (1996, -0.84 DPS) [world]; Mindbender Loop (5009, -0.84 DPS) [world_drop]; Ogremind Ring (1993, -1.17 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Ankh of Life (1713, -0.19 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -8.11 DPS) [quest]; Guardian Talisman (1490, -8.11 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Ankh of Life (1713, -0.39 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -0.63 DPS) [quest]; Guardian Talisman (1490, -0.63 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 | yes | Soulkeeper (1607, -1.83 DPS) [world_drop]; Radiant Staff (249453, -3.44 DPS) [crafted]; Spire of Hakkar (10844, -4.19 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world]; Lesser Eternal Wand (249232, -5.92 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Darkspear Raider's Cloak; chest: Acumen Robes; wrist: Shizzle's Nozzle Wiper; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 722, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 60 (gnome, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 212.3. Weights run: 0.8s. Verify run: 0.7s. 1223 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.724 ± 0.748), crit=2.882 ± 0.178, hit=12.470 ± 0.766, spell_haste=not significant (1.648 ± 0.891), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tiara of the Oracle (21348) | Tiara of the Oracle [quest] | 0.0 | yes | Virtuous Crown (22080, -12.06 DPS) [quest]; Virtuous Crown (226947, -12.06 DPS) [quest]; Bloodvine Goggles (19999, -13.12 DPS, sim-verified) [crafted] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 0.0 | yes | Beads of Ogre Might (22150, -4.64 DPS) [quest]; Medallion of the Dawn (22659, -14.33 DPS) [quest]; Charm of the Shifting Sands (21504, -15.10 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 66.8 | yes | Shroud of the Nathrezim (18720, -0.55 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.37 DPS) [vendor]; Blood Guard's Dreadweave Mantle (220905, -1.37 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 124.7 | yes | Chromatic Cloak (18509, -0.08 DPS, sim-verified) [crafted]; Hide of the Wild (18510, -11.89 DPS) [crafted]; Spritecaster Cape (11623, -12.22 DPS) [dungeon] |
| chest | Vestments of the Oracle (21351) | Vestments of the Oracle [quest] | 0.0 | yes | Earthpower Vest (21183, -1.60 DPS) [quest]; Virtuous Robe (226945, -3.59 DPS) [quest]; Bloodvine Vest (19682, -6.54 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 151.7 | yes | Dryad's Wrist Bindings (19595, +0.00 DPS, sim-verified) [rep]; Dryad's Wrist Bindings (19596, -14.64 DPS) [rep]; Dryad's Wrist Bindings (19597, -15.10 DPS) [rep] |
| hands | Dreadmist Wraps (16705) | Scholomance: Lorekeeper Polkelt [dungeon] | 131.2 | yes | Gloves of Spell Mastery (14146, -1.55 DPS, sim-verified) [crafted]; Marshal's Satin Gloves (17608, -11.35 DPS) [vendor]; General's Satin Gloves (17620, -11.35 DPS) [vendor] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 71.9 | yes | Highlander's Cloth Girdle (20097, -2.18 DPS) [rep]; Highlander's Cloth Girdle (20047, -4.11 DPS, sim-verified) [rep]; Stormpike Cloth Girdle (19094, -5.37 DPS) [rep] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 166.0 | yes | Magister's Leggings (16687, +0.00 DPS, sim-verified) [dungeon]; Knight's Dreadweave Leggings (220888, -12.07 DPS) [vendor]; Stone Guard's Dreadweave Leggings (220906, -12.07 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 155.3 | yes | Sergeant Major's Dreadweave Boots (220891, -1.85 DPS) [vendor]; First Sergeant's Dreadweave Boots (220909, -4.15 DPS, sim-verified) [vendor]; Marshal's Satin Sandals (17607, -13.70 DPS) [vendor] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 166.0 | yes | Band of Earthen Might (21182, -0.11 DPS) [quest]; Blackstone Ring (17713, -4.75 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -4.75 DPS) [quest] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 165.1 | yes | Band of Earthen Might (21182, +0.00 DPS, sim-verified) [quest]; Blackstone Ring (17713, -4.64 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -4.64 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Uther's Strength (11302, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.41 DPS, sim-verified) [world] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | 0.0 | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Stormrager (16997) | Order Must Be Restored [quest] | 0.0 | yes | Brilliant Wand (249385, -1.91 DPS) [crafted]; Wand of Biting Cold (19108, -2.10 DPS, sim-verified) [quest]; Torch of Austen (13004, -6.32 DPS) [world] |

**New at 60:** head: Tiara of the Oracle; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Vestments of the Oracle; wrist: Rockfury Bracers; hands: Dreadmist Wraps; waist: Belt of the Archmage; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Ring of the Fallen God; finger2: Don Julio's Band; trinket1: Ankh of Life; trinket2: Thunderbrew's Boot Flask; main_hand: Ironbark Staff; ranged: Stormrager

No-known-source sample (15 of 1223, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

## Horde

### Band 20 (undead, 000000000000000000-00000000000000000-443000000000000000)

Set DPS (verified): 26.9. Weights run: 0.9s. Verify run: 0.7s. 203 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=-3.036 ± 0.220, crit=0.729 ± 0.038, hit=1.171 ± 0.150, spell_haste=not significant (-0.287 ± 0.238), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.54 DPS) [crafted]; Lucky Fishing Hat (19972, -0.54 DPS) [quest]; Flying Tiger Goggles (4368, -0.80 DPS, sim-verified) [crafted] |
| neck | Scholarly Pendant (277203) | Friend of the Library [quest] | 0.0 | yes | Erudite's Amulet (277204, +0.00 DPS) [quest]; Tarnished Locket (279870, +0.00 DPS) [quest]; Scout's Medallion (20442, -0.30 DPS, sim-verified) [rep] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.26 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.45 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, +0.00 DPS, sim-verified) [dungeon]; Black Whelp Cloak (7283, -0.09 DPS) [crafted]; Battle Healer's Cloak (20427, -0.09 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.09 DPS) [crafted]; Bloody Apron (6226, -0.09 DPS) [dungeon]; Green Woolen Vest (2582, -0.52 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) (or Windsong Bangles (263336)) | Earthen Arise [quest] | 1.0 | yes | Windsong Bangles (263336, +0.00 DPS, sim-verified) [quest]; Mindthrust Bracers (1974, -0.09 DPS) [dungeon]; Silver-lined Bracers (3224, -0.09 DPS) [world] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.09 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.27 DPS) [quest]; Pristine Gloves (253913, -0.27 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 0.0 | yes | Novice Ardent's Sash (253887, -0.18 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.36 DPS) [world]; Novice Arcanist's Sash (253885, -0.62 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.27 DPS) [crafted]; Silk-threaded Trousers (1929, -0.28 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.36 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.27 DPS) [crafted]; Feather Padded Treads (285345, -0.34 DPS, sim-verified) [world]; Pristine Boots (253889, -0.36 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.45 DPS) [dungeon]; Ring of the Shadow (1462, -0.45 DPS) [world]; Ring of Scorn (3235, -0.45 DPS) [quest] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Ring of the Shadow (1462, -0.27 DPS) [world]; Ring of Scorn (3235, -0.27 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.38 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.0 | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Crescent Staff (6505, -0.10 DPS) [quest]; Living Root (6631, -0.42 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 249.2 | yes | Firebelcher (5243, -0.07 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.05 DPS) [quest]; Sizzle Stick (8071, -4.48 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scholarly Pendant; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (undead, 000000000000000000-00000000000000000-443110501200000000)

Set DPS (verified): 47.0. Weights run: 0.7s. Verify run: 0.7s. 407 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.550 ± 0.331, crit=0.803 ± 0.052, hit=3.822 ± 0.217, spell_haste=not significant (-0.760 ± 0.265), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 21.5 | yes | Nightsky Cowl (4039, -0.39 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.40 DPS) [crafted]; Cloudy Stormsewn Cowl (277046, -0.67 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 16.3 | yes | Darkspear Warding Pendant (272075, -0.82 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.90 DPS) [world_drop]; Pendant of Myzrael (4614, -1.45 DPS) [dungeon] |
| shoulder | Death Speaker Mantle (6685) | Razorfen Kraul: Death Speaker Jargba [dungeon] | 23.1 | yes | Bloodmage Mantle (7684, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Batwing Mantle (6697, -0.53 DPS) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 12.4 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Soft Willow Cape (16661, -0.41 DPS) [quest]; Pearl-clasped Cloak (5542, -0.51 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 29.2 | yes | Death Speaker Robes (6682, -0.39 DPS, sim-verified) [dungeon]; Black Velvet Robes (2800, -0.94 DPS) [world_drop]; Pristine Gown (253961, -1.00 DPS) [crafted] |
| wrist | Nightsky Wristbands (6407) (or Tabitha's Cuffs (251486)) | World drop [world_drop] | 9.3 | yes | Spidertank Oilrag (9448, -0.03 DPS) [dungeon]; Tabitha's Cuffs (251486, -0.09 DPS, sim-verified) [quest]; Mindthrust Bracers (1974, -0.14 DPS) [dungeon] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 13.8 | yes | Blight Gloves (279877, -0.26 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.33 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.36 DPS) [crafted] |
| waist | Crimson Silk Belt (7055) | Tailoring [crafted] | 16.9 | yes | Defiler's Cloth Girdle (20164, +0.00 DPS, sim-verified) [rep]; Invoker's Cord (215366, -0.19 DPS) [crafted]; Lilac Sash (6780, -0.26 DPS) [quest] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 0.0 | yes | Filigreed Pristine Leggings (253937, -0.23 DPS) [crafted]; Silver-thread Pants (4037, -0.48 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.53 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 17.9 | yes | Spidersilk Boots (4320, -0.41 DPS) [crafted]; Frothing Slippers (254003, -0.62 DPS) [crafted]; Acidic Walkers (9454, -0.81 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 10.9 | yes | Lavishly Jeweled Ring (1156, -0.14 DPS) [dungeon]; Azora's Will (4999, -0.28 DPS) [world_drop]; Loop of Sacrifice (281673, -0.28 DPS) [quest] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 10.9 | yes | Azora's Will (4999, -0.28 DPS) [world_drop]; Loop of Sacrifice (281673, -0.28 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.33 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 293.5 | yes | Gnarled Ash Staff (791, -0.58 DPS, sim-verified) [world_drop]; Lorekeeper's Staff (212580, -0.58 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.58 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 377.0 | yes | Greater Mystic Wand (217287, -0.87 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest]; Scorching Wand (5213, -4.63 DPS) [world_drop] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Death Speaker Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Nightsky Wristbands; hands: Jutebraid Gloves; waist: Crimson Silk Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 407, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (undead, 000000000000000000-00000000000000000-443110501201300240)

Set DPS (verified): 70.9. Weights run: 0.7s. Verify run: 0.7s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (-0.254 ± 0.524), crit=1.076 ± 0.084, hit=5.292 ± 0.396, spell_haste=not significant (0.657 ± 0.448), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Holy Shroud (2721, -0.94 DPS) [world_drop]; Silk Headband (7050, -1.13 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Necklace of Calisea (1714, -0.39 DPS, sim-verified) [world_drop]; Ethereal Talisman (4430, -0.66 DPS) [quest]; Pendant of Myzrael (4614, -0.66 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.0 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Chestnut Mantle (17695, -0.09 DPS) [quest]; Berylline Pads (4197, -0.19 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, +0.00 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.09 DPS) [crafted]; Battle Healer's Cloak (19528, -0.09 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.0 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.76 DPS) [crafted]; Tree Bark Jacket (1486, -0.85 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Radiant Silver Bracers (4545, -0.47 DPS) [quest]; Earthen Silk Cuffs (254019, -0.47 DPS) [crafted]; Condor Bracers (15864, -0.64 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Red Mageweave Gloves (10018, -0.66 DPS) [crafted]; Gilded Handwraps (254021, -0.94 DPS) [crafted]; Black Mageweave Gloves (10003, -1.21 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.0 | yes | Warsong Sash (16975, -0.28 DPS) [quest]; Defiler's Cloth Girdle (20164, -0.28 DPS) [rep]; Star Belt (4329, -0.28 DPS, sim-verified) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 14.0 | yes | Abomination Skin Leggings (23173, -0.47 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -0.57 DPS) [crafted]; Gaze Dreamer Pants (6903, -1.36 DPS, sim-verified) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -1.60 DPS) [crafted]; Spidersilk Boots (4320, -1.73 DPS, sim-verified) [crafted]; Boots of the Enchanter (4325, -1.79 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.0 | yes | Reedknot Ring (9622, -0.28 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Electrocutioner Lagnut (9447, -0.66 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.19 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.28 DPS) [vendor]; Reedknot Ring (9622, -0.64 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 428.2 | yes | Illusionary Rod (7713, -0.81 DPS, sim-verified) [dungeon]; Black Duskwood Staff (937, -6.84 DPS) [world_drop]; Spiritchaser Staff (1613, -6.84 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | World drop [world_drop] | 377.9 | yes | Twisted Nether Wand (249144, +0.00 DPS, sim-verified) [crafted]; Ember Wand (5215, -2.00 DPS) [world_drop]; Necrotic Wand (7708, -2.18 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Bloodmage Mantle; back: Icy Cloak; chest: Robe of the Magi; wrist: Spidertank Oilrag; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Staff of Jordan; ranged: Umbral Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs

### Band 50 (undead, 521000000000000000-00000000000000000-443110501201300251)

Set DPS (verified): 113.0. Weights run: 0.7s. Verify run: 0.7s. 723 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.962 ± 0.656), crit=1.607 ± 0.120, hit=8.647 ± 0.553, spell_haste=not significant (-0.545 ± 0.921), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 60.0 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.03 DPS) [dungeon]; Red Mageweave Headband (10033, -0.19 DPS) [crafted] |
| neck | Horizon Choker (13085) | Azuregos [world] | 27.5 | yes | Mindburst Medallion (11196, -1.01 DPS) [quest]; Darkspear Warding Pendant (272073, -1.02 DPS) [vendor]; Scorn's Icy Choker (23169, -1.76 DPS, sim-verified) [dungeon] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 48.3 | yes | Blood Guard's Dreadweave Mantle (220905, -0.02 DPS) [vendor]; Red Mageweave Shoulders (10029, -1.24 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -3.06 DPS, sim-verified) [vendor] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 29.7 | yes | Darkspear Raider's Cloak (272076, +0.00 DPS, sim-verified) [vendor]; Spritecaster Cape (11623, -0.40 DPS) [dungeon]; Runecloth Cloak (13860, -0.52 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | 58.2 | yes | Stone Guard's Dreadweave Vest (220904, -0.23 DPS) [vendor]; Runecloth Robe (13858, -1.45 DPS) [crafted]; Knight's Dreadweave Vest (220886, -3.30 DPS, sim-verified) [vendor] |
| wrist | Shizzle's Nozzle Wiper (11917) | Shizzle's Flyer [quest] | 23.5 | yes | Bloodband Bracers (11469, +0.00 DPS, sim-verified) [quest]; Imperial Red Bracers (8247, -0.20 DPS) [dungeon]; Nethergeld Cuffs (254061, -0.29 DPS) [crafted] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 50.2 | yes | Virtuous Hands (226958, -1.21 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -2.03 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -2.03 DPS) [vendor] |
| waist | Dawnspire Cord (12466) | Sunken Temple: Morphaz [dungeon] | 43.3 | yes | Deathmage Sash (10771, -0.71 DPS) [dungeon]; Satyrmane Sash (17755, -1.01 DPS) [dungeon]; Defiler's Cloth Girdle (20165, -4.26 DPS, sim-verified) [rep] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 58.0 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Kilt of the Atal'ai Prophet (10807, -2.05 DPS) [dungeon]; Red Mageweave Pants (10009, -2.14 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 112.1 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Southsea Mojo Boots (20641, -8.60 DPS) [quest]; Gilded Sandals (254107, -8.70 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 86.5 | yes | Ogremind Ring (1993, -7.58 DPS) [world_drop]; Voodoo Band (1996, -7.58 DPS) [world]; Mindbender Loop (5009, -7.58 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 21.8 | yes | Voodoo Band (1996, -0.84 DPS) [world]; Mindbender Loop (5009, -0.84 DPS) [world_drop]; Ogremind Ring (1993, -1.06 DPS, sim-verified) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Uther's Strength (11302, -0.34 DPS, sim-verified) [world]; Guardian Talisman (1490, -8.11 DPS) [quest]; Ankh of Life (1713, -8.11 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Uther's Strength (11302, -0.23 DPS, sim-verified) [world]; Guardian Talisman (1490, -6.31 DPS) [quest]; Ankh of Life (1713, -6.31 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 | yes | Soulkeeper (1607, -1.83 DPS) [world_drop]; Radiant Staff (249453, -3.44 DPS) [crafted]; Spire of Hakkar (10844, -4.19 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 503.8 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world]; Lesser Eternal Wand (249232, -5.92 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Shizzle's Nozzle Wiper; hands: Raider Handwraps; waist: Dawnspire Cord; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 723, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (undead, 524111001300000000-00000000000000000-443110501201300251)

Set DPS (verified): 199.7. Weights run: 0.8s. Verify run: 0.7s. 1224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (0.724 ± 0.748), crit=2.882 ± 0.178, hit=12.470 ± 0.766, spell_haste=not significant (1.648 ± 0.891), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Tiara of the Oracle (21348) | Tiara of the Oracle [quest] | 0.0 | yes | Bloodvine Goggles (19999, -9.29 DPS, sim-verified) [crafted]; Virtuous Crown (22080, -12.06 DPS) [quest]; Virtuous Crown (226947, -12.06 DPS) [quest] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 0.0 | yes | Beads of Ogre Might (22150, -4.64 DPS) [quest]; Medallion of the Dawn (22659, -14.33 DPS) [quest]; Charm of the Shifting Sands (21504, -15.10 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 66.8 | yes | Shroud of the Nathrezim (18720, +0.00 DPS, sim-verified) [dungeon]; Knight-Lieutenant's Dreadweave Mantle (220887, -1.37 DPS) [vendor]; Blood Guard's Dreadweave Mantle (220905, -1.37 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 124.7 | yes | Chromatic Cloak (18509, -0.17 DPS, sim-verified) [crafted]; Hide of the Wild (18510, -11.89 DPS) [crafted]; Deep Woodlands Cloak (19121, -12.20 DPS) [quest] |
| chest | Vestments of the Oracle (21351) | Vestments of the Oracle [quest] | 0.0 | yes | Earthpower Vest (21183, -1.60 DPS) [quest]; Virtuous Robe (226945, -3.59 DPS) [quest]; Bloodvine Vest (19682, -5.11 DPS, sim-verified) [crafted] |
| wrist | Dryad's Wrist Bindings (19595) | Warsong Outriders [rep] | 0.0 | yes | Dryad's Wrist Bindings (19596, -0.40 DPS) [rep]; Dryad's Wrist Bindings (19597, -0.86 DPS) [rep]; Rockfury Bracers (21186, -1.92 DPS, sim-verified) [quest] |
| hands | Dreadmist Wraps (16705) | Scholomance: Lorekeeper Polkelt [dungeon] | 131.2 | yes | Gloves of Spell Mastery (14146, -1.44 DPS, sim-verified) [crafted]; Marshal's Satin Gloves (17608, -11.35 DPS) [vendor]; General's Satin Gloves (17620, -11.35 DPS) [vendor] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 71.9 | yes | Defiler's Cloth Girdle (20165, -2.18 DPS) [rep]; Defiler's Cloth Girdle (20163, -3.09 DPS, sim-verified) [rep]; Frostwolf Cloth Belt (19090, -5.37 DPS) [rep] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 166.0 | yes | Magister's Leggings (16687, +0.00 DPS, sim-verified) [dungeon]; Knight's Dreadweave Leggings (220888, -12.07 DPS) [vendor]; Stone Guard's Dreadweave Leggings (220906, -12.07 DPS) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 155.3 | yes | First Sergeant's Dreadweave Boots (220909, -1.85 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -3.08 DPS, sim-verified) [vendor]; Marshal's Satin Sandals (17607, -13.70 DPS) [vendor] |
| finger1 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 166.0 | yes | Band of Earthen Might (21182, -0.11 DPS) [quest]; Blackstone Ring (17713, -4.75 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -4.75 DPS) [quest] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 165.1 | yes | Band of Earthen Might (21182, +0.00 DPS, sim-verified) [quest]; Blackstone Ring (17713, -4.64 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -4.64 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.60 DPS, sim-verified) [crafted]; Guardian Talisman (1490, -0.69 DPS) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | 0.0 | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | 556.9 | yes | Stormrager (16997, -0.75 DPS, sim-verified) [quest]; Brilliant Wand (249385, -2.80 DPS) [crafted]; Torch of Austen (13004, -7.21 DPS) [world] |

**New at 60:** head: Tiara of the Oracle; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Vestments of the Oracle; wrist: Dryad's Wrist Bindings; hands: Dreadmist Wraps; waist: Belt of the Archmage; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Ring of the Fallen God; finger2: Don Julio's Band; trinket1: Ankh of Life; trinket2: Uther's Strength; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1224, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

