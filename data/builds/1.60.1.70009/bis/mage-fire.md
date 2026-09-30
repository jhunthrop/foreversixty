# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 31.9. Weights run: 0.6s. Verify run: 0.5s. 252 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=1.617 ± 0.198, crit=2.699 ± 0.108, hit=6.702 ± 0.242, spell_haste=not significant (0.409 ± 0.385), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.45 DPS) [crafted]; Lucky Fishing Hat (19972, -0.45 DPS) [quest]; Shadow Goggles (4373, -0.79 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 11.5 | yes | Double-Stitched Woolen Shoulders (4314, +0.03 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.86 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.9 | yes | Heavy Woolen Cloak (4311, +0.03 DPS, sim-verified) [crafted]; Forest Cloak (4710, -0.27 DPS) [world_drop]; Seer's Cape (6378, -0.27 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 13.1 | yes | Seer's Robe (2981, -0.25 DPS) [world_drop]; Manaweave Robe (7509, -0.25 DPS) [quest]; Gray Woolen Robe (2585, -0.56 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 8.1 | yes | Bright Bracers (3647, -0.12 DPS) [world_drop]; Seer's Cuffs (3645, -0.48 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.55 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 11.3 | yes | Tomb Robber's Gloves (280096, -0.18 DPS, sim-verified) [quest]; Pristine Gloves (253913, -0.19 DPS) [crafted]; Adept's Gloves (4768, -0.24 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 10.5 | yes | Novice Arcanist's Sash (253885, -0.12 DPS) [crafted]; Tarantula Silk Sash (3229, -0.18 DPS) [world]; Keller's Girdle (2911, -1.01 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 21.9 | yes | Filigreed Pristine Leggings (253937, -0.00 DPS, sim-verified) [crafted]; Filigreed Silky Leggings (253939, -0.92 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.92 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 13.5 | yes | Pristine Boots (253889, -0.40 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.52 DPS) [world]; Kimbra Boots (6191, -0.52 DPS) [quest] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.7 | yes | Lorekeeper's Ring (20431, -0.35 DPS) [rep]; Black Pearl Ring (6332, -0.48 DPS) [world]; Sludge-Stained Band (286535, -0.50 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 8.2 | yes | Lorekeeper's Ring (20431, -0.00 DPS, sim-verified) [rep]; Black Pearl Ring (6332, -0.37 DPS) [world]; Sludge-Stained Band (286535, -0.39 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 16.2 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.24 DPS) [world]; Lesser Staff of the Spire (1300, -0.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Moonstone Wand (15204, -0.13 DPS) [quest]; Cookie's Stirring Rod (5198, -0.15 DPS) [dungeon]; Flaring Baton (5326, -3.56 DPS, sim-verified) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 252, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 58.3. Weights run: 0.6s. Verify run: 0.6s. 476 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.375 ± 0.506), crit=3.624 ± 0.223, hit=6.056 ± 0.383, spell_haste=5.146 ± 0.503, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 19.7 | yes | Shadow Hood (4323, -0.44 DPS) [crafted]; Nightsky Cowl (4039, -0.53 DPS, sim-verified) [world_drop]; Cloudy Stormsewn Cowl (277046, -0.70 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 15.2 | yes | Crystal Starfire Medallion (5003, -0.93 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.19 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.45 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 21.4 | yes | Death Speaker Mantle (6685, -0.15 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Batwing Mantle (6697, -0.60 DPS) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 11.0 | yes | Darkspear Raider's Cloak (272078, +0.10 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.24 DPS) [quest]; Pearl-clasped Cloak (5542, -0.46 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 26.9 | yes | Death Speaker Robes (6682, -0.47 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.98 DPS) [crafted]; Black Velvet Robes (2800, -0.99 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.07 DPS) [quest]; Mindthrust Bracers (1974, -0.20 DPS) [dungeon]; Nightsky Wristbands (6407, -0.99 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 19.1 | yes | Hotshot Pilot's Gloves (9491, -0.73 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.91 DPS) [quest]; Truefaith Gloves (7049, -0.95 DPS) [crafted] |
| waist | Crimson Silk Belt (7055) | Tailoring [crafted] | 15.6 | yes | Highlander's Cloth Girdle (20099, +0.44 DPS, sim-verified) [rep]; Invoker's Cord (215366, -0.17 DPS) [crafted]; Belt of Arugal (6392, -0.24 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 16.6 | yes | Filigreed Pristine Leggings (253937, -0.23 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.44 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.81 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 16.6 | yes | Spidersilk Boots (4320, -0.39 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -1.27 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 9.6 | yes | Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Minor Channeling Ring (1449, -0.18 DPS) [quest]; Lorekeeper's Ring (19525, -0.25 DPS) [rep] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.6 | yes | Minor Channeling Ring (1449, -0.18 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.25 DPS, sim-verified) [dungeon]; Lorekeeper's Ring (19525, -0.25 DPS) [rep] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 15.1 | yes | Gnarled Necromancer's Staff (251534, -0.13 DPS) [quest]; Twisted Chanter's Staff (890, -0.17 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.39 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Mystic Wand (11289) | Enchanting [crafted] | 5.5 | yes | Wand of Eventide (5214, +0.56 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.05 DPS) [quest]; Greater Mystic Wand (217287, -0.05 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Crimson Silk Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Lesser Mystic Wand

No-known-source sample (15 of 476, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 98.3. Weights run: 0.6s. Verify run: 0.6s. 643 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.220 ± 0.772), crit=3.709 ± 0.304, hit=6.038 ± 0.597, spell_haste=7.821 ± 0.995, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Craftsman's Monocle (4393, -0.35 DPS) [crafted]; Enchanter's Cowl (4322, -0.36 DPS) [crafted]; Augural Shroud (2620, -1.03 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.3 | yes | Necklace of Calisea (1714, -0.42 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.74 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.05 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.9 | yes | Green Silken Shoulders (7057, +0.50 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.37 DPS) [dungeon]; Death Speaker Mantle (6685, -0.44 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 13.4 | yes | Long Silken Cloak (4326, +0.86 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.17 DPS) [crafted]; Cloak of Rot (4462, -0.47 DPS) [world] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.3 | yes | Dreamweave Vest (10021, +0.26 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.09 DPS) [crafted]; Green Silk Armor (7065, -0.57 DPS) [crafted] |
| wrist | Aurora Bracers (4043) (or Mistscape Bracers (4045), Enchanted Stonecloth Bracers (4979)) | World drop [world_drop] | 9.8 | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [world_drop]; Enchanted Stonecloth Bracers (4979, +0.00 DPS) [quest]; Arcane Runed Bracers (4744, -0.10 DPS) [quest] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 23.2 | yes | Dreamweave Gloves (10019, +0.60 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.58 DPS) [crafted]; Town Clerk's Mittens (270029, -0.74 DPS) [quest] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.3 | yes | Highlander's Cloth Girdle (20098, +0.34 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.97 DPS) [crafted]; Highlander's Cloth Girdle (20099, -1.36 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.6 | yes | Crimson Silk Pantaloons (7062, -0.67 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.27 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.52 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.95 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.19 DPS) [dungeon]; Spidersilk Boots (4320, -1.55 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Voodoo Band (1996, -0.06 DPS) [world]; Mindbender Loop (5009, -0.06 DPS) [world_drop]; Black Widow Band (6199, -0.06 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 8.5 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 59.2 | yes | Windweaver Staff (7757, +0.43 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -5.88 DPS) [world_drop]; Glimmering Staff (249392, -5.88 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Burning Sliver (5249) | Crushridge Warmongers [quest] | 6.0 | yes | Twisted Nether Wand (249144, +0.00 DPS) [crafted]; Wand of Eventide (5214, -0.13 DPS) [world_drop]; Fizzle's Zippy Lighter (6729, -1.45 DPS, sim-verified) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Aurora Bracers; hands: Red Mageweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod; ranged: Burning Sliver

No-known-source sample (15 of 643, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 131.2. Weights run: 0.6s. Verify run: 0.5s. 806 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (0.248 ± 1.140), crit=4.947 ± 0.428, hit=9.366 ± 0.828, spell_haste=14.377 ± 1.359, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 86.2 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -1.54 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -7.46 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 | yes | Mindburst Medallion (11196, -0.48 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.63 DPS) [world]; Darkspear Warding Pendant (272073, -0.79 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 79.5 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -7.81 DPS) [dungeon]; Black Mageweave Shoulders (10027, -8.47 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 | yes | Nightfall Drape (12465, -0.82 DPS) [world]; Runecloth Cloak (13860, -0.91 DPS, sim-verified) [crafted]; Long Silken Cloak (4326, -1.04 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 84.0 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Acumen Robes (17775, -7.56 DPS) [quest]; Robe of the Magi (1716, -7.62 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Nethergeld Cuffs (254061, -0.03 DPS) [crafted]; Spidertank Oilrag (9448, -0.17 DPS, sim-verified) [dungeon]; Bloodband Bracers (11469, -0.22 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | First Sergeant's Dreadweave Gloves (220908, -0.47 DPS) [vendor]; Black Mageweave Gloves (10003, -0.50 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.71 DPS, sim-verified) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 79.5 | yes | Satyrmane Sash (17755, +0.12 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -8.13 DPS) [rep]; Ghostweave Cord (254073, -8.25 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 84.2 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -8.22 DPS) [crafted]; Red Mageweave Pants (10009, -8.47 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 103.9 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -10.06 DPS) [crafted]; Gilded Sandals (254107, -11.42 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 93.7 | yes | Ring of Forlorn Spirits (2043, -10.79 DPS) [quest]; Reedknot Ring (9622, -10.92 DPS) [quest]; Sea Giant's Toe Ring (274746, -11.04 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Ring of Forlorn Spirits (2043, -0.50 DPS) [quest]; Reedknot Ring (9622, -0.63 DPS) [quest]; Lorekeeper's Ring (19524, -0.75 DPS, sim-verified) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 84.3 | yes | Uther's Strength (11302, -9.86 DPS) [world]; Thunderbrew's Boot Flask (744, -10.62 DPS) [quest]; Tidal Charm (1404, -10.62 DPS) [vendor] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, -0.20 DPS, sim-verified) [world] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 72.2 | yes | Inventor's Focal Sword (17719, -0.38 DPS) [dungeon]; Illusionary Rod (7713, -1.41 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -8.38 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Burning Sliver (5249, -0.25 DPS) [quest]; Twisted Nether Wand (249144, -0.25 DPS) [crafted]; Dreambough Wand (249234, -3.79 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; main_hand: Kindling Stave; ranged: Lesser Eternal Wand

No-known-source sample (15 of 806, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 428.1. Weights run: 0.6s. Verify run: 0.6s. 1229 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.134 ± 1.662), crit=7.715 ± 0.552, hit=17.053 ± 1.684, spell_haste=not significant (-2.222 ± 2.347), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 447.6 | yes | Bloodvine Goggles (19999, -19.80 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -23.51 DPS) [quest]; Field Marshal's Coronet (16441, -62.06 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 278.5 | yes | Beads of Ogre Might (22150, -23.33 DPS) [quest]; Medallion of the Dawn (22659, -36.83 DPS) [quest]; Charm of the Shifting Sands (21504, -51.82 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 139.7 | yes | Lieutenant Commander's Silk Mantle (23319, -0.92 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.92 DPS) [pvp]; Champion's Silk Mantle (23264, -4.58 DPS, sim-verified) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 170.5 | yes | Chromatic Cloak (18509, -7.79 DPS, sim-verified) [crafted]; Drape of Vaulted Secrets (21415, -30.74 DPS) [quest]; Hide of the Wild (18510, -31.36 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 356.2 | yes | Bloodvine Vest (19682, -9.80 DPS, sim-verified) [crafted]; Zandalar Illusionist's Robe (20034, -28.38 DPS) [quest]; Enigma Robes (21343, -39.54 DPS) [quest] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 197.5 | yes | Frostfire Bindings (22503, +2.44 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -35.95 DPS) [rep]; Dryad's Wrist Bindings (19596, -36.87 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 236.4 | yes | Sorcerer's Gloves (22066, -4.15 DPS, sim-verified) [quest]; Dreadmist Wraps (16705, -12.01 DPS) [dungeon]; Frostfire Gloves (22501, -38.62 DPS) [quest] |
| waist | Frostfire Belt (22502) | Frostfire Belt [quest] | 222.3 | yes | Belt of the Archmage (18405, -9.82 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -20.20 DPS) [rep]; Highlander's Cloth Girdle (20097, -21.52 DPS) [rep] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 246.0 | yes | Sorcerer's Leggings (226933, -8.28 DPS) [quest]; Magister's Leggings (16687, -14.80 DPS) [dungeon]; Bloodvine Leggings (19683, -17.50 DPS, sim-verified) [crafted] |
| feet | Enigma Boots (21344) | Enigma Boots [quest] | 215.5 | yes | Marshal's Silk Footwraps (16437, -1.76 DPS) [vendor]; General's Silk Boots (16539, -1.76 DPS) [vendor]; Bloodvine Boots (19684, -7.43 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 278.5 | yes | Band of Earthen Might (21182, -8.63 DPS, sim-verified) [quest]; Blackstone Ring (17713, -23.33 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -23.33 DPS) [quest] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 214.3 | yes | Band of Earthen Might (21182, -8.46 DPS, sim-verified) [quest]; Blackstone Ring (17713, -9.46 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -9.46 DPS) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Uther's Strength (11302, +0.00 DPS) [world]; Thunderbrew's Boot Flask (744, -3.60 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 153.5 | yes | Uther's Strength (11302, -1.38 DPS, sim-verified) [world]; Thunderbrew's Boot Flask (744, -33.15 DPS) [quest]; Tidal Charm (1404, -33.15 DPS) [vendor] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 323.1 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Ironbark Staff (20069, -20.68 DPS) [rep]; Blade of Vaulted Secrets (21413, -29.03 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 16.9 | yes | Greater Eternal Wand (249237, -1.71 DPS) [crafted]; Lesser Eternal Wand (249232, -1.93 DPS) [crafted]; Charged Lightning Rod (11860, -19.55 DPS, sim-verified) [quest] |

**New at 60:** head: Frostfire Circlet; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Frostfire Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket1: Ankh of Life; trinket2: Frozen Heart of the Mountain; main_hand: High Warlord's War Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1229, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 29.0. Weights run: 0.6s. Verify run: 0.5s. 249 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=1.617 ± 0.198, crit=2.699 ± 0.108, hit=6.702 ± 0.242, spell_haste=not significant (0.409 ± 0.385), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.35 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.45 DPS) [crafted]; Lucky Fishing Hat (19972, -0.45 DPS) [quest] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 11.5 | yes | Double-Stitched Woolen Shoulders (4314, -0.56 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.86 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 6.9 | yes | Heavy Woolen Cloak (4311, -0.00 DPS, sim-verified) [crafted]; Forest Cloak (4710, -0.27 DPS) [world_drop]; Seer's Cape (6378, -0.27 DPS) [world_drop] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 13.1 | yes | Seer's Robe (2981, -0.25 DPS) [world_drop]; Lesser Spellfire Robes (7510, -0.25 DPS) [quest]; Gray Woolen Robe (2585, -0.79 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 9.7 | yes | Mindthrust Bracers (1974, +0.12 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.12 DPS) [quest]; Bright Bracers (3647, -0.24 DPS) [world_drop] |
| hands | Blight Gloves (279877) | The New Plague [quest] | 11.3 | yes | Pristine Gloves (253913, -0.19 DPS) [crafted]; Adept's Gloves (4768, -0.24 DPS) [world]; Tomb Robber's Gloves (280096, -0.41 DPS, sim-verified) [quest] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 12.9 | yes | Pristine Sash (253925, +0.25 DPS, sim-verified) [crafted]; Novice Arcanist's Sash (253885, -0.31 DPS) [crafted]; Tarantula Silk Sash (3229, -0.36 DPS) [world] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 21.9 | yes | Filigreed Pristine Leggings (253937, -0.10 DPS, sim-verified) [crafted]; Filigreed Silky Leggings (253939, -0.92 DPS) [crafted]; Filigreed Flame Leggings (253941, -0.92 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 13.5 | yes | Pristine Boots (253889, -0.28 DPS, sim-verified) [crafted]; Walking Boots (4660, -0.52 DPS) [world]; Woolen Boots (2583, -0.62 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 9.7 | yes | Black Pearl Ring (6332, -0.48 DPS) [world]; Sludge-Stained Band (286535, -0.50 DPS) [world]; The 1 Ring (8350, -0.61 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.15 DPS) [world]; The 1 Ring (8350, -0.25 DPS) [world]; Black Pearl Ring (6332, -0.43 DPS, sim-verified) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 16.2 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.24 DPS) [world]; Lesser Staff of the Spire (1300, -0.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Sizzle Stick (8071) | Deviate Eradication [quest] | 5.0 | yes | Flaring Baton (5326, -0.13 DPS) [quest]; Cookie's Stirring Rod (5198, -0.15 DPS) [dungeon]; Wand of Decay (5252, -2.78 DPS, sim-verified) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Blight Gloves; waist: Keller's Girdle; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Sizzle Stick

No-known-source sample (15 of 249, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 53.7. Weights run: 0.6s. Verify run: 0.6s. 473 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.375 ± 0.506), crit=3.624 ± 0.223, hit=6.056 ± 0.383, spell_haste=5.146 ± 0.503, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 19.7 | yes | Nightsky Cowl (4039, -0.30 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.44 DPS) [crafted]; Cloudy Stormsewn Cowl (277046, -0.70 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 15.2 | yes | Crystal Starfire Medallion (5003, -0.93 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.23 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.45 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 21.4 | yes | Death Speaker Mantle (6685, -0.24 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.29 DPS) [quest]; Batwing Mantle (6697, -0.60 DPS) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 11.0 | yes | Darkspear Raider's Cloak (272078, +0.27 DPS, sim-verified) [vendor]; Soft Willow Cape (16661, -0.39 DPS) [quest]; Pearl-clasped Cloak (5542, -0.46 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 26.9 | yes | Death Speaker Robes (6682, -0.47 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.98 DPS) [crafted]; Black Velvet Robes (2800, -0.99 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.07 DPS) [quest]; Mindthrust Bracers (1974, -0.20 DPS) [dungeon]; Nightsky Wristbands (6407, -0.47 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 12.9 | yes | Hotshot Pilot's Gloves (9491, -0.11 DPS, sim-verified) [dungeon]; Blight Gloves (279877, -0.31 DPS) [quest]; Truefaith Gloves (7049, -0.36 DPS) [crafted] |
| waist | Crimson Silk Belt (7055) | Tailoring [crafted] | 15.6 | yes | Defiler's Cloth Girdle (20164, +0.01 DPS, sim-verified) [rep]; Invoker's Cord (215366, -0.17 DPS) [crafted]; Belt of Arugal (6392, -0.24 DPS) [dungeon] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 16.6 | yes | Filigreed Pristine Leggings (253937, -0.23 DPS) [crafted]; Gaze Dreamer Pants (6903, -0.44 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.92 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 16.6 | yes | Spidersilk Boots (4320, -0.39 DPS) [crafted]; Frothing Slippers (254003, -0.67 DPS) [crafted]; Acidic Walkers (9454, -1.02 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 9.6 | yes | Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Advisor's Ring (19521, -0.25 DPS) [rep]; Azora's Will (4999, -0.26 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 9.6 | yes | Advisor's Ring (19521, -0.25 DPS) [rep]; Azora's Will (4999, -0.26 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.43 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 15.1 | yes | Gnarled Necromancer's Staff (251534, -0.13 DPS) [quest]; Twisted Chanter's Staff (890, -0.32 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.39 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Mystic Wand (11289) | Enchanting [crafted] | 5.5 | yes | Wand of Eventide (5214, +0.13 DPS, sim-verified) [world_drop]; Sizzle Stick (8071, -0.05 DPS) [quest]; Greater Mystic Wand (217287, -0.05 DPS) [crafted] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Crimson Silk Belt; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Lesser Mystic Wand

No-known-source sample (15 of 473, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 91.3. Weights run: 0.6s. Verify run: 0.6s. 639 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (1.220 ± 0.772), crit=3.709 ± 0.304, hit=6.038 ± 0.597, spell_haste=7.821 ± 0.995, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 23.2 | yes | Spellpower Goggles Xtreme (10502, +0.44 DPS, sim-verified) [crafted]; Craftsman's Monocle (4393, -0.63 DPS) [crafted]; Enchanter's Cowl (4322, -0.64 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.3 | yes | Necklace of Calisea (1714, -0.69 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.74 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.05 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 22.9 | yes | Green Silken Shoulders (7057, -0.10 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.37 DPS) [dungeon]; Death Speaker Mantle (6685, -0.44 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 13.4 | yes | Long Silken Cloak (4326, +0.02 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.17 DPS) [crafted]; Cloak of Rot (4462, -0.47 DPS) [world] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.3 | yes | Dreamweave Vest (10021, +0.43 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.09 DPS) [crafted]; Green Silk Armor (7065, -0.57 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 13.8 | yes | Mistscape Bracers (4045, -0.51 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.51 DPS) [quest]; Aurora Bracers (4043, -0.63 DPS, sim-verified) [world_drop] |
| hands | Red Mageweave Gloves (10018) | Tailoring [crafted] | 23.2 | yes | Dreamweave Gloves (10019, +0.83 DPS, sim-verified) [crafted]; Stormcloth Gloves (10011, -0.58 DPS) [crafted]; Gilded Handwraps (254021, -0.85 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 25.3 | yes | Defiler's Cloth Girdle (20166, -0.14 DPS, sim-verified) [rep]; Gilded Cord (254037, -0.97 DPS) [crafted]; Defiler's Cloth Girdle (20164, -1.36 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 28.6 | yes | Crimson Silk Pantaloons (7062, -0.50 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.27 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.52 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, +0.27 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.19 DPS) [dungeon]; Spidersilk Boots (4320, -1.55 DPS) [crafted] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Voodoo Band (1996, -0.06 DPS) [world]; Mindbender Loop (5009, -0.06 DPS) [world_drop]; Black Widow Band (6199, -0.06 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 8.5 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Windweaver Staff (7757) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 18.3 | yes | Staff of Jordan (873, -0.63 DPS) [world_drop]; Glimmering Staff (249392, -0.63 DPS) [crafted]; Illusionary Rod (7713, -1.07 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 6.0 | yes | Wand of Eventide (5214, -0.13 DPS) [world_drop]; Sizzle Stick (8071, -0.13 DPS) [quest]; Fizzle's Zippy Lighter (6729, -1.14 DPS, sim-verified) [quest] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Red Mageweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Windweaver Staff; ranged: Twisted Nether Wand

No-known-source sample (15 of 639, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 122.5. Weights run: 0.6s. Verify run: 0.5s. 802 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=not significant (0.248 ± 1.140), crit=4.947 ± 0.428, hit=9.366 ± 0.828, spell_haste=14.377 ± 1.359, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 86.2 | yes | Knight-Lieutenant's Dreadweave Hat (220889, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -1.54 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -7.46 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.5 | yes | Mindburst Medallion (11196, -0.43 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.63 DPS) [world]; Darkspear Warding Pendant (272073, -0.79 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 79.5 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -7.81 DPS) [dungeon]; Black Mageweave Shoulders (10027, -8.47 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.5 | yes | Runecloth Cloak (13860, -0.57 DPS) [crafted]; Nightfall Drape (12465, -0.82 DPS) [world]; Deep Woodlands Cloak (19121, -1.23 DPS, sim-verified) [quest] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 84.0 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Acumen Robes (17775, -7.56 DPS) [quest]; Robe of the Magi (1716, -7.62 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, +1.12 DPS, sim-verified) [crafted]; Bloodband Bracers (11469, -0.22 DPS) [quest]; Condor Bracers (15864, -0.25 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 19.0 | yes | First Sergeant's Dreadweave Gloves (220908, -0.47 DPS) [vendor]; Black Mageweave Gloves (10003, -0.50 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.87 DPS, sim-verified) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 79.5 | yes | Satyrmane Sash (17755, -0.27 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -8.13 DPS) [rep]; Ghostweave Cord (254073, -8.25 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 84.2 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -8.22 DPS) [crafted]; Red Mageweave Pants (10009, -8.47 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 103.9 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -10.06 DPS) [crafted]; Gilded Sandals (254107, -11.42 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 93.7 | yes | Reedknot Ring (9622, -10.92 DPS) [quest]; Sea Giant's Toe Ring (274746, -11.04 DPS) [vendor]; Electrocutioner Lagnut (9447, -11.42 DPS) [dungeon] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Reedknot Ring (9622, -0.63 DPS) [quest]; Advisor's Ring (19521, -0.63 DPS) [rep]; Advisor's Ring (19520, -1.28 DPS, sim-verified) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 6.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.20 DPS, sim-verified) [crafted]; Tidal Charm (1404, -0.76 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 72.2 | yes | Inventor's Focal Sword (17719, -0.38 DPS) [dungeon]; Illusionary Rod (7713, -1.53 DPS, sim-verified) [dungeon]; Spellshifter Rod (9527, -8.38 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Lesser Eternal Wand (249232) | Enchanting [crafted] | 8.0 | yes | Twisted Nether Wand (249144, -0.25 DPS) [crafted]; Charged Lightning Rod (11860, -0.25 DPS) [quest]; Dreambough Wand (249234, -4.86 DPS, sim-verified) [crafted] |

**New at 50:** head: Blood Guard's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Spidertank Oilrag; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Lesser Eternal Wand

No-known-source sample (15 of 802, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 400.1. Weights run: 0.6s. Verify run: 0.6s. 1225 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=not significant (1.134 ± 1.662), crit=7.715 ± 0.552, hit=17.053 ± 1.684, spell_haste=not significant (-2.222 ± 2.347), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 447.6 | yes | Bloodvine Goggles (19999, -18.23 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -23.51 DPS) [quest]; Field Marshal's Coronet (16441, -62.06 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 278.5 | yes | Beads of Ogre Might (22150, -23.33 DPS) [quest]; Medallion of the Dawn (22659, -36.83 DPS) [quest]; Charm of the Shifting Sands (21504, -51.82 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 139.7 | yes | Lieutenant Commander's Silk Mantle (23319, -0.92 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.92 DPS) [pvp]; Champion's Silk Mantle (23264, -1.51 DPS, sim-verified) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 170.5 | yes | Chromatic Cloak (18509, -2.10 DPS, sim-verified) [crafted]; Drape of Vaulted Secrets (21415, -30.74 DPS) [quest]; Hide of the Wild (18510, -31.36 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 356.2 | yes | Bloodvine Vest (19682, -12.45 DPS, sim-verified) [crafted]; Zandalar Illusionist's Robe (20034, -28.38 DPS) [quest]; Enigma Robes (21343, -39.54 DPS) [quest] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 197.5 | yes | Frostfire Bindings (22503, +0.81 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -35.95 DPS) [rep]; Dryad's Wrist Bindings (19596, -36.87 DPS) [rep] |
| hands | Sorcerer's Gloves (22066) | Just Compensation [quest] | 198.4 | yes | Dreadmist Wraps (16705, -3.82 DPS) [dungeon]; Gloves of Spell Mastery (14146, -9.82 DPS, sim-verified) [crafted]; Frostfire Gloves (22501, -30.42 DPS) [quest] |
| waist | Frostfire Belt (22502) | Frostfire Belt [quest] | 222.3 | yes | Belt of the Archmage (18405, -7.04 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20163, -20.20 DPS) [rep]; Defiler's Cloth Girdle (20165, -21.52 DPS) [rep] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 246.0 | yes | Sorcerer's Leggings (226933, -8.28 DPS) [quest]; Bloodvine Leggings (19683, -10.54 DPS, sim-verified) [crafted]; Magister's Leggings (16687, -14.80 DPS) [dungeon] |
| feet | Enigma Boots (21344) | Enigma Boots [quest] | 215.5 | yes | Marshal's Silk Footwraps (16437, -1.76 DPS) [vendor]; General's Silk Boots (16539, -1.76 DPS) [vendor]; Bloodvine Boots (19684, -10.64 DPS, sim-verified) [crafted] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 278.5 | yes | Band of Earthen Might (21182, -6.63 DPS, sim-verified) [quest]; Blackstone Ring (17713, -23.33 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -23.33 DPS) [quest] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 214.3 | yes | Blackstone Ring (17713, -9.46 DPS) [dungeon]; Master Dragonslayer's Ring (19384, -9.46 DPS) [quest]; Band of Earthen Might (21182, -11.17 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 153.5 | yes | Uther's Strength (11302, -0.75 DPS, sim-verified) [world]; Tidal Charm (1404, -33.15 DPS) [vendor]; Guardian Talisman (1490, -33.15 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 119.4 | yes | Uther's Strength (11302, +0.01 DPS, sim-verified) [world]; Tidal Charm (1404, -25.78 DPS) [vendor]; Guardian Talisman (1490, -25.78 DPS) [quest] |
| main_hand | High Warlord's War Staff (234549) | Rank 18 [pvp] | 323.1 | yes | Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Ironbark Staff (20220, -20.68 DPS) [rep]; Blade of Vaulted Secrets (21413, -29.03 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 16.9 | yes | Greater Eternal Wand (249237, -1.71 DPS) [crafted]; Lesser Eternal Wand (249232, -1.93 DPS) [crafted]; Charged Lightning Rod (11860, -17.07 DPS, sim-verified) [quest] |

**New at 60:** head: Frostfire Circlet; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Sorcerer's Gloves; waist: Frostfire Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: High Warlord's War Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1225, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

