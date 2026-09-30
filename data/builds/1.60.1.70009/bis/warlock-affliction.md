# Leveling BiS: Affliction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 37.7. Weights run: 1.3s. Verify run: 1.1s. 204 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.102, intellect=not significant (-0.303 ± 0.109), crit=0.392 ± 0.018, hit=1.206 ± 0.062, spell_haste=not significant (-0.036 ± 0.114), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.102

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.96 DPS) [crafted]; Lucky Fishing Hat (19972, -0.96 DPS) [quest]; Flying Tiger Goggles (4368, -2.40 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.67 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.80 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.16 DPS) [crafted]; Caretaker's Cape (20428, -0.16 DPS) [rep]; Feyscale Cloak (6632, -0.30 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.16 DPS) [crafted]; Bloody Apron (6226, -0.16 DPS) [dungeon]; Green Woolen Vest (2582, -1.21 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [world_drop]; Silver-lined Bracers (3224, -0.16 DPS) [world]; Seer's Cuffs (3645, -0.16 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.48 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.80 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 0.0 | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.64 DPS) [world]; Novice Arcanist's Sash (253885, -0.98 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted]; Colorful Kilt (10048, -0.64 DPS) [crafted]; Silk-threaded Trousers (1929, -0.78 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.48 DPS) [crafted]; Pristine Boots (253889, -0.64 DPS) [crafted]; Feather Padded Treads (285345, -0.74 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.32 DPS) [world]; Lavishly Jeweled Ring (1156, -0.80 DPS) [dungeon]; Ring of the Shadow (1462, -0.80 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.35 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.80 DPS) [dungeon]; Ring of the Shadow (1462, -0.80 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Staff of Westfall (2042) | The Defias Brotherhood [quest] | 0.0 | yes | Twisted Chanter's Staff (890, -0.06 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -0.14 DPS) [quest]; Living Root (6631, -0.50 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 142.1 | yes | Firebelcher (5243, -1.55 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.26 DPS) [quest]; Sizzle Stick (8071, -4.34 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Westfall; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 204, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak; 9792 Ivycloth Boots

### Band 30 (gnome, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 57.3. Weights run: 1.4s. Verify run: 1.4s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.164, intellect=0.928 ± 0.177, crit=0.391 ± 0.020, hit=1.447 ± 0.093, spell_haste=0.972 ± 0.185, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.164

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.3 | yes | Holy Shroud (2721, -0.90 DPS) [world_drop]; Shadow Hood (4323, -1.06 DPS) [crafted]; Nightsky Cowl (4039, -1.57 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.6 | yes | Darkspear Warding Pendant (272075, -1.80 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -1.85 DPS) [world_drop]; Pendant of Myzrael (4614, -2.63 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.4 | yes | Fairywing Mantle (9536, -0.63 DPS) [quest]; Death Speaker Mantle (6685, -0.70 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -1.20 DPS) [crafted] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.4 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.15 DPS) [quest]; Hillman's Cloak (3719, -0.51 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.1 | yes | Death Speaker Robes (6682, -0.70 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -1.58 DPS) [crafted]; Tree Bark Jacket (1486, -1.69 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.72 DPS) [quest]; Mindthrust Bracers (1974, -0.91 DPS) [world_drop]; Nightsky Wristbands (6407, -2.18 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 0.0 | yes | Hotshot Pilot's Gloves (9491, -0.07 DPS) [world_drop]; Serpent Gloves (5970, -0.16 DPS) [dungeon]; Town Clerk's Mittens (270029, -1.32 DPS, sim-verified) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 13.8 | yes | Belt of Arugal (6392, -0.42 DPS) [dungeon]; Invoker's Cord (215366, -0.45 DPS) [crafted]; Crimson Silk Belt (7055, -1.27 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 0.0 | yes | Gaze Dreamer Pants (6903, -0.31 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Abomination Skin Leggings (23173, -1.52 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.5 | yes | Spidersilk Boots (4320, -0.58 DPS) [crafted]; Frothing Slippers (254003, -1.46 DPS) [crafted]; Acidic Walkers (9454, -2.17 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Black Widow Band (6199, -0.10 DPS) [world]; Snake Hoop (6750, -0.10 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.21 DPS) [vendor] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 | yes | Snake Hoop (6750, -0.07 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.18 DPS) [vendor]; Black Widow Band (6199, -1.06 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 127.6 | yes | Lorekeeper's Staff (212580, +0.00 DPS, sim-verified) [vendor]; Advisor's Gnarled Staff (212584, -0.86 DPS) [vendor]; Gnarled Ash Staff (791, -1.07 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 161.8 | yes | Greater Mystic Wand (217287, -1.96 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.84 DPS) [quest]; Scorching Wand (5213, -4.99 DPS) [world_drop] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Minor Channeling Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 40 (gnome, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 109.1. Weights run: 1.1s. Verify run: 1.1s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.127, intellect=0.992 ± 0.157, crit=0.387 ± 0.020, hit=1.104 ± 0.078, spell_haste=1.008 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.127

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -1.42 DPS) [crafted]; Craftsman's Monocle (4393, -1.71 DPS) [crafted]; Augural Shroud (2620, -2.72 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.0 | yes | Darkspear Warding Pendant (272074, -1.68 DPS) [vendor]; Necklace of Calisea (1714, -1.96 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.24 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.9 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.55 DPS) [dungeon]; Berylline Pads (4197, -0.83 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.0 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Cloak of Rot (4462, -0.85 DPS) [world] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.0 | yes | Robe of Power (7054, -0.57 DPS) [crafted]; Dreamweave Vest (10021, -0.82 DPS, sim-verified) [crafted]; Crimson Silk Vest (7058, -1.69 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, -0.23 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -0.30 DPS) [world_drop]; Mistscape Bracers (4045, -0.30 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 | yes | Stormcloth Gloves (10011, -1.70 DPS) [crafted]; Black Mageweave Gloves (10003, -1.95 DPS) [crafted]; Red Mageweave Gloves (10018, -2.00 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 0.0 | yes | Gilded Cord (254037, -0.57 DPS) [crafted]; Highlander's Cloth Girdle (20099, -1.12 DPS) [rep]; Deathmage Sash (10771, -1.45 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.9 | yes | Crimson Silk Pantaloons (7062, -1.91 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -2.77 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.10 DPS) [dungeon]; Spidersilk Boots (4320, -3.65 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.0 | yes | Ring of Forlorn Spirits (2043, -2.23 DPS) [quest]; Reedknot Ring (9622, -2.51 DPS) [quest]; Minor Channeling Ring (1449, -2.51 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Ring of Forlorn Spirits (2043, -0.42 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.56 DPS) [quest]; Lorekeeper's Ring (19525, -0.56 DPS) [rep] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | 0.0 | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS, sim-verified) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 155.2 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -6.88 DPS) [world_drop]; Staff of Dar'Orahil (15106, -8.43 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 127.4 | yes | Burning Sliver (5249, -1.31 DPS) [quest]; Umbral Wand (5216, -1.36 DPS, sim-verified) [world_drop]; Necrotic Wand (7708, -1.63 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs

### Band 50 (gnome, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 180.8. Weights run: 1.2s. Verify run: 1.2s. 721 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.227, intellect=not significant (0.199 ± 0.264), crit=0.457 ± 0.026, hit=2.450 ± 0.155, spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Red Mageweave Headband (10033, -1.13 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.18 DPS) [vendor]; Dreamweave Circlet (10041, -2.43 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.2 | yes | Mindburst Medallion (11196, -0.18 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.52 DPS) [world_drop]; Darkspear Warding Pendant (272073, -1.80 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.6 | yes | Blood Guard's Dreadweave Mantle (220905, -0.11 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.34 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.42 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 | yes | Nightfall Drape (12465, -1.74 DPS) [dungeon]; Icy Cloak (4327, -2.30 DPS) [crafted]; Runecloth Cloak (13860, -2.70 DPS, sim-verified) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 | yes | Knight's Dreadweave Vest (220886, -0.73 DPS) [vendor]; Stone Guard's Dreadweave Vest (220904, -0.73 DPS) [vendor]; Acumen Robes (17775, -0.99 DPS, sim-verified) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Nethergeld Cuffs (254061, -0.17 DPS) [crafted]; Spidertank Oilrag (9448, -0.27 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 | yes | Sergeant Major's Dreadweave Gloves (220890, -1.12 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.12 DPS) [vendor]; Black Mageweave Gloves (10003, -1.61 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 16.4 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.45 DPS) [rep]; Ghostweave Cord (254073, -0.67 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 20.8 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -0.50 DPS) [crafted]; Red Mageweave Pants (10009, -1.23 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 34.3 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -2.88 DPS) [crafted]; Gilded Sandals (254107, -6.03 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 24.5 | yes | Philanthropist's Ring (281635, -3.73 DPS) [quest]; Ring of Forlorn Spirits (2043, -4.63 DPS) [quest]; Reedknot Ring (9622, -4.91 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Lorekeeper's Ring (19524, -0.84 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.12 DPS) [quest]; Philanthropist's Ring (281635, -1.44 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | 0.0 | yes | Frozen Heart of the Mountain (249469, -1.53 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -1.68 DPS) [quest]; Guardian Talisman (1490, -1.68 DPS) [quest] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, -3.36 DPS) [quest]; Guardian Talisman (1490, -3.36 DPS) [quest]; Frozen Heart of the Mountain (249469, -3.80 DPS, sim-verified) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 0.0 | yes | Kindling Stave (11750, -4.29 DPS) [dungeon]; Soulkeeper (1607, -5.42 DPS) [world_drop]; Spire of Hakkar (10844, -6.14 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | 0.0 | yes | Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted]; Pyric Caduceus (11748, -4.58 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Abyss Shard; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 721, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 315.5. Weights run: 1.2s. Verify run: 1.2s. 1140 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.837), intellect=not significant (-1.254 ± 0.970), crit=1.999 ± 0.106, hit=6.401 ± 0.516, spell_haste=not significant (0.241 ± 0.963), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.837)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | 0.0 | yes | Doomcaller's Circlet (21337, -4.06 DPS) [quest]; Deathmist Mask (226909, -10.29 DPS) [quest]; Bloodvine Goggles (19999, -13.41 DPS, sim-verified) [crafted] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 0.0 | yes | Fury of the Forgotten Swarm (21809, +0.00 DPS) [world_drop]; Beads of Ogre Might (22150, -4.06 DPS) [quest]; Medallion of the Dawn (22659, -9.28 DPS) [quest] |
| shoulder | Plagueheart Shoulderpads (22507) | Plagueheart Shoulderpads [quest] | 100.0 | yes | Doomcaller's Mantle (21335, -3.71 DPS, sim-verified) [quest]; Rugged Mantle of the Timbermaw (227808, -6.53 DPS) [vendor]; Mantle of the Timbermaw (19050, -7.98 DPS) [crafted] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 64.0 | yes | Chromatic Cloak (18509, -0.94 DPS, sim-verified) [crafted]; Shroud of Unspoken Names (21418, -6.67 DPS) [quest]; Spritecaster Cape (11623, -7.25 DPS) [dungeon] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 0.0 | yes | Zandalar Demoniac's Robe (20033, -7.54 DPS) [quest]; Bloodvine Vest (19682, -8.65 DPS, sim-verified) [crafted]; Doomcaller's Robes (21334, -10.73 DPS) [quest] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 91.0 | yes | Plagueheart Bindings (22511, -2.06 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -10.01 DPS) [rep]; Dryad's Wrist Bindings (19596, -10.30 DPS) [rep] |
| hands | Deathmist Wraps (22077) (or Deathmist Wraps (226911)) | Just Compensation [quest] | 77.0 | yes | Deathmist Wraps (226911, +0.00 DPS, sim-verified) [quest]; Gloves of Spell Mastery (14146, -1.75 DPS) [crafted]; Dreadmist Wraps (16705, -1.89 DPS) [dungeon] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 89.0 | yes | Plagueheart Belt (22510, +0.00 DPS, sim-verified) [quest]; Belt of the Archmage (18405, -5.95 DPS) [crafted]; Highlander's Cloth Girdle (20047, -6.82 DPS) [rep] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 101.0 | yes | Plagueheart Leggings (22505, -5.23 DPS) [quest]; Magister's Leggings (16687, -5.28 DPS) [dungeon]; Deathmist Leggings (226910, -7.55 DPS, sim-verified) [quest] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 83.0 | yes | First Sergeant's Dreadweave Boots (220909, -1.60 DPS) [vendor]; Argent Elite Boots (227816, -2.76 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -4.82 DPS, sim-verified) [vendor] |
| finger1 | Ring of Unspoken Names (21417) | Ring of Unspoken Names [quest] | 106.0 | yes | Don Julio's Band (19325, -2.03 DPS) [rep]; Band of Earthen Might (21182, -2.03 DPS) [quest]; Blackstone Ring (17713, -6.09 DPS) [dungeon] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 101.0 | yes | Band of Earthen Might (21182, -1.31 DPS) [quest]; Blackstone Ring (17713, -5.37 DPS) [dungeon]; Don Julio's Band (19325, -13.71 DPS, sim-verified) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | 0.0 | yes | Thunderbrew's Boot Flask (744, -0.87 DPS) [quest]; Guardian Talisman (1490, -0.87 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.31 DPS, sim-verified) [crafted] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, -1.74 DPS) [quest]; Guardian Talisman (1490, -1.74 DPS) [quest]; Frozen Heart of the Mountain (249469, -4.04 DPS, sim-verified) [crafted] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | 0.0 | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | 441.3 | yes | Brilliant Wand (249385, -3.11 DPS) [crafted]; Stormrager (16997, -3.65 DPS, sim-verified) [quest]; Torch of Austen (13004, -7.21 DPS) [world_drop] |

**New at 60:** head: Plagueheart Circlet; neck: Onyxia Tooth Pendant; shoulder: Plagueheart Shoulderpads; back: Earthweave Cloak; chest: Plagueheart Robe; wrist: Rockfury Bracers; hands: Deathmist Wraps; waist: Knowledge of the Timbermaw; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Ring of Unspoken Names; finger2: Ring of the Fallen God; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1140, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 25400000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 36.4. Weights run: 1.3s. Verify run: 1.1s. 203 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.102, intellect=not significant (-0.303 ± 0.109), crit=0.392 ± 0.018, hit=1.206 ± 0.062, spell_haste=not significant (-0.036 ± 0.114), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.102

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.96 DPS) [crafted]; Lucky Fishing Hat (19972, -0.96 DPS) [quest]; Flying Tiger Goggles (4368, -2.39 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Slime-encrusted Pads (6461, -0.80 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.28 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, +0.00 DPS, sim-verified) [dungeon]; Black Whelp Cloak (7283, -0.16 DPS) [crafted]; Battle Healer's Cloak (20427, -0.16 DPS) [rep] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.16 DPS) [crafted]; Bloody Apron (6226, -0.16 DPS) [dungeon]; Green Woolen Vest (2582, -2.15 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) (or Windsong Bangles (263336)) | Earthen Arise [quest] | 1.0 | yes | Mindthrust Bracers (1974, -0.16 DPS) [world_drop]; Silver-lined Bracers (3224, -0.16 DPS) [world]; Windsong Bangles (263336, -0.86 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.06 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.48 DPS) [quest]; Pristine Gloves (253913, -0.48 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 0.0 | yes | Novice Ardent's Sash (253887, -0.32 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.64 DPS) [world]; Novice Arcanist's Sash (253885, -0.90 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.48 DPS) [crafted]; Colorful Kilt (10048, -0.64 DPS) [crafted]; Silk-threaded Trousers (1929, -1.83 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.48 DPS) [crafted]; Pristine Boots (253889, -0.64 DPS) [crafted]; Feather Padded Treads (285345, -1.42 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.80 DPS) [dungeon]; Ring of the Shadow (1462, -0.80 DPS) [world]; Ring of Scorn (3235, -0.80 DPS) [quest] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, -0.35 DPS, sim-verified) [dungeon]; Ring of the Shadow (1462, -0.48 DPS) [world]; Ring of Scorn (3235, -0.48 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.0 | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Crescent Staff (6505, -0.10 DPS) [quest]; Living Root (6631, -0.45 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 142.1 | yes | Firebelcher (5243, -1.36 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.26 DPS) [quest]; Sizzle Stick (8071, -4.34 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (troll, 25552000110000000-0000000000000000000-0000000000000000)

Set DPS (verified): 55.3. Weights run: 1.4s. Verify run: 1.3s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.164, intellect=0.928 ± 0.177, crit=0.391 ± 0.020, hit=1.447 ± 0.093, spell_haste=0.972 ± 0.185, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.164

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 15.3 | yes | Holy Shroud (2721, -0.90 DPS) [world_drop]; Shadow Hood (4323, -1.06 DPS) [crafted]; Nightsky Cowl (4039, -1.64 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.6 | yes | Crystal Starfire Medallion (5003, -1.85 DPS) [world_drop]; Darkspear Warding Pendant (272075, -2.08 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -2.63 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 17.4 | yes | Death Speaker Mantle (6685, -0.62 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.63 DPS) [quest]; Invoker's Mantle (215365, -1.20 DPS) [crafted] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 7.4 | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Hillman's Cloak (3719, -0.51 DPS) [crafted]; Windsong Drape (15468, -0.51 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 21.1 | yes | Death Speaker Robes (6682, -0.65 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -1.58 DPS) [crafted]; Tree Bark Jacket (1486, -1.69 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.72 DPS) [quest]; Mindthrust Bracers (1974, -0.91 DPS) [world_drop]; Nightsky Wristbands (6407, -2.33 DPS, sim-verified) [world_drop] |
| hands | Truefaith Gloves (7049) | Tailoring [crafted] | 0.0 | yes | Hotshot Pilot's Gloves (9491, -0.07 DPS) [world_drop]; Serpent Gloves (5970, -0.16 DPS) [dungeon]; Jutebraid Gloves (10654, -0.80 DPS, sim-verified) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 13.8 | yes | Belt of Arugal (6392, -0.42 DPS) [dungeon]; Invoker's Cord (215366, -0.45 DPS) [crafted]; Crimson Silk Belt (7055, -1.14 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | 0.0 | yes | Gaze Dreamer Pants (6903, -0.31 DPS) [dungeon]; Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Abomination Skin Leggings (23173, -1.39 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 13.5 | yes | Spidersilk Boots (4320, -0.58 DPS) [crafted]; Frothing Slippers (254003, -1.46 DPS) [crafted]; Acidic Walkers (9454, -2.34 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Snake Hoop (6750, -0.10 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.21 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.30 DPS) [dungeon] |
| finger2 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 6.5 | yes | Snake Hoop (6750, -0.01 DPS, sim-verified) [quest]; Sea Giant's Toe Ring (274746, -0.10 DPS) [vendor]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 127.6 | yes | Lorekeeper's Staff (212580, -0.27 DPS, sim-verified) [vendor]; Advisor's Gnarled Staff (212584, -0.86 DPS) [vendor]; Gnarled Ash Staff (791, -1.07 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 161.8 | yes | Greater Mystic Wand (217287, -2.31 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.84 DPS) [quest]; Scorching Wand (5213, -4.99 DPS) [world_drop] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Truefaith Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Black Widow Band; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (troll, 25552000130201050-0000000000000000000-0000000000000000)

Set DPS (verified): 106.0. Weights run: 1.1s. Verify run: 1.1s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.127, intellect=0.992 ± 0.157, crit=0.387 ± 0.020, hit=1.104 ± 0.078, spell_haste=1.008 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.127

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -1.42 DPS) [crafted]; Craftsman's Monocle (4393, -1.71 DPS) [crafted]; Augural Shroud (2620, -2.57 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.0 | yes | Darkspear Warding Pendant (272074, -1.68 DPS) [vendor]; Necklace of Calisea (1714, -2.04 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.24 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.9 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.55 DPS) [dungeon]; Berylline Pads (4197, -0.83 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 11.0 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Cloak of Rot (4462, -0.85 DPS) [world] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.0 | yes | Robe of Power (7054, -0.57 DPS) [crafted]; Dreamweave Vest (10021, -0.73 DPS, sim-verified) [crafted]; Crimson Silk Vest (7058, -1.69 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 11.9 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Aurora Bracers (4043, -1.12 DPS) [world_drop]; Mistscape Bracers (4045, -1.12 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 | yes | Stormcloth Gloves (10011, -1.70 DPS) [crafted]; Red Mageweave Gloves (10018, -1.72 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -1.95 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 0.0 | yes | Gilded Cord (254037, -0.57 DPS) [crafted]; Deathmage Sash (10771, -1.05 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20164, -1.12 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.9 | yes | Crimson Silk Pantaloons (7062, -1.73 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.07 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -2.45 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -3.10 DPS) [dungeon]; Spidersilk Boots (4320, -3.65 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.0 | yes | Reedknot Ring (9622, -2.51 DPS) [quest]; Ogremind Ring (1993, -2.52 DPS) [world_drop]; Voodoo Band (1996, -2.52 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.56 DPS) [rep]; Ogremind Ring (1993, -0.58 DPS) [world_drop]; Reedknot Ring (9622, -0.58 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Rune of Duty (21567) | Warsong Outriders [rep] | 0.0 | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS, sim-verified) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 155.2 | yes | Illusionary Rod (7713, -0.09 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -6.88 DPS) [world_drop]; Staff of Dar'Orahil (15106, -8.43 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 127.4 | yes | Necrotic Wand (7708, -1.63 DPS) [dungeon]; Umbral Wand (5216, -1.81 DPS, sim-verified) [world_drop]; Goblin Igniter (5253, -1.82 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (troll, 25552000130201051-2340000000000000000-0000000000000000)

Set DPS (verified): 178.6. Weights run: 1.2s. Verify run: 1.2s. 721 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.227, intellect=not significant (0.199 ± 0.264), crit=0.457 ± 0.026, hit=2.450 ± 0.155, spell_haste=not significant (-0.231 ± 0.284), spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.227

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Red Mageweave Headband (10033, -1.13 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -1.18 DPS) [vendor]; Dreamweave Circlet (10041, -2.78 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.2 | yes | Mindburst Medallion (11196, -0.49 DPS, sim-verified) [quest]; Horizon Choker (13085, -1.52 DPS) [world_drop]; Darkspear Warding Pendant (272073, -1.80 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | 16.6 | yes | Blood Guard's Dreadweave Mantle (220905, -0.11 DPS) [vendor]; Black Mageweave Shoulders (10027, -1.34 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -2.51 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 15.2 | yes | Deep Woodlands Cloak (19121, -0.77 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.29 DPS) [crafted]; Nightfall Drape (12465, -1.74 DPS) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 23.2 | yes | Knight's Dreadweave Vest (220886, -0.73 DPS) [vendor]; Stone Guard's Dreadweave Vest (220904, -0.73 DPS) [vendor]; Acumen Robes (17775, -0.74 DPS, sim-verified) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.56 DPS) [quest]; Bloodband Bracers (11469, -0.62 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.8 | yes | Sergeant Major's Dreadweave Gloves (220890, -1.12 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.12 DPS) [vendor]; Black Mageweave Gloves (10003, -1.21 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 16.4 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -0.45 DPS) [rep]; Ghostweave Cord (254073, -0.67 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 20.8 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -0.50 DPS) [crafted]; Red Mageweave Pants (10009, -1.23 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 34.3 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -2.88 DPS) [crafted]; Gilded Sandals (254107, -6.03 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 24.5 | yes | Philanthropist's Ring (281635, -3.73 DPS) [quest]; Reedknot Ring (9622, -4.91 DPS) [quest]; Sea Giant's Toe Ring (274746, -5.19 DPS) [vendor] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Advisor's Ring (19520, -0.84 DPS) [rep]; Reedknot Ring (9622, -1.40 DPS) [quest]; Philanthropist's Ring (281635, -1.55 DPS, sim-verified) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Guardian Talisman (1490, -3.36 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.42 DPS, sim-verified) [crafted]; Guardian Talisman (1490, -1.68 DPS) [quest] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 0.0 | yes | Kindling Stave (11750, -4.29 DPS) [dungeon]; Soulkeeper (1607, -5.42 DPS) [world_drop]; Spire of Hakkar (10844, -6.14 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | 0.0 | yes | Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted]; Pyric Caduceus (11748, -4.57 DPS, sim-verified) [dungeon] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 721, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (troll, 25552000130201051-2355220000000000000-0000000000000000)

Set DPS (verified): 317.1. Weights run: 1.2s. Verify run: 1.2s. 1140 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.837), intellect=not significant (-1.254 ± 0.970), crit=1.999 ± 0.106, hit=6.401 ± 0.516, spell_haste=not significant (0.241 ± 0.963), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.837)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | 0.0 | yes | Doomcaller's Circlet (21337, -4.06 DPS) [quest]; Deathmist Mask (226909, -10.29 DPS) [quest]; Bloodvine Goggles (19999, -12.47 DPS, sim-verified) [crafted] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 0.0 | yes | Fury of the Forgotten Swarm (21809, +0.00 DPS) [world_drop]; Beads of Ogre Might (22150, -4.06 DPS) [quest]; Medallion of the Dawn (22659, -9.28 DPS) [quest] |
| shoulder | Plagueheart Shoulderpads (22507) | Plagueheart Shoulderpads [quest] | 100.0 | yes | Doomcaller's Mantle (21335, -2.70 DPS, sim-verified) [quest]; Rugged Mantle of the Timbermaw (227808, -6.53 DPS) [vendor]; Mantle of the Timbermaw (19050, -7.98 DPS) [crafted] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 64.0 | yes | Chromatic Cloak (18509, -0.61 DPS, sim-verified) [crafted]; Shroud of Unspoken Names (21418, -6.67 DPS) [quest]; Spritecaster Cape (11623, -7.25 DPS) [dungeon] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 0.0 | yes | Zandalar Demoniac's Robe (20033, -7.54 DPS) [quest]; Bloodvine Vest (19682, -8.69 DPS, sim-verified) [crafted]; Doomcaller's Robes (21334, -10.73 DPS) [quest] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 91.0 | yes | Plagueheart Bindings (22511, -2.90 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -10.01 DPS) [rep]; Dryad's Wrist Bindings (19596, -10.30 DPS) [rep] |
| hands | Deathmist Wraps (22077) (or Deathmist Wraps (226911)) | Just Compensation [quest] | 77.0 | yes | Deathmist Wraps (226911, +0.00 DPS, sim-verified) [quest]; Gloves of Spell Mastery (14146, -1.75 DPS) [crafted]; Dreadmist Wraps (16705, -1.89 DPS) [dungeon] |
| waist | Plagueheart Belt (22510) | Plagueheart Belt [quest] | 0.0 | yes | Belt of the Archmage (18405, -2.03 DPS) [crafted]; Defiler's Cloth Girdle (20163, -2.90 DPS) [rep]; Knowledge of the Timbermaw (228190, -3.22 DPS, sim-verified) [vendor] |
| legs | Bloodvine Leggings (19683) | Tailoring [crafted] | 101.0 | yes | Plagueheart Leggings (22505, -5.23 DPS) [quest]; Magister's Leggings (16687, -5.28 DPS) [dungeon]; Deathmist Leggings (226910, -6.91 DPS, sim-verified) [quest] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 83.0 | yes | First Sergeant's Dreadweave Boots (220909, -1.60 DPS) [vendor]; Argent Elite Boots (227816, -2.76 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -3.76 DPS, sim-verified) [vendor] |
| finger1 | Ring of Unspoken Names (21417) | Ring of Unspoken Names [quest] | 106.0 | yes | Don Julio's Band (19325, -2.03 DPS) [rep]; Band of Earthen Might (21182, -2.03 DPS) [quest]; Blackstone Ring (17713, -6.09 DPS) [dungeon] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 101.0 | yes | Band of Earthen Might (21182, -1.31 DPS) [quest]; Blackstone Ring (17713, -5.37 DPS) [dungeon]; Don Julio's Band (19325, -13.99 DPS, sim-verified) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Guardian Talisman (1490, -1.74 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Guardian Talisman (1490, -0.87 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.89 DPS, sim-verified) [crafted] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | 0.0 | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | 441.3 | yes | Brilliant Wand (249385, -3.11 DPS) [crafted]; Stormrager (16997, -3.35 DPS, sim-verified) [quest]; Torch of Austen (13004, -7.21 DPS) [world_drop] |

**New at 60:** head: Plagueheart Circlet; neck: Onyxia Tooth Pendant; shoulder: Plagueheart Shoulderpads; back: Earthweave Cloak; chest: Plagueheart Robe; wrist: Rockfury Bracers; hands: Deathmist Wraps; waist: Plagueheart Belt; legs: Bloodvine Leggings; feet: Bloodvine Boots; finger1: Ring of Unspoken Names; finger2: Ring of the Fallen God; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1140, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

