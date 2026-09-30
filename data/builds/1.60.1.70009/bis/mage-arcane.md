# Leveling BiS: Arcane

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 32.7. Weights run: 0.7s. Verify run: 0.6s. 203 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.395 ± 0.090, crit=1.480 ± 0.048, hit=3.723 ± 0.108, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.49 DPS) [crafted]; Lucky Fishing Hat (19972, -0.49 DPS) [quest]; Shadow Goggles (4373, -0.64 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.6 | yes | Double-Stitched Woolen Shoulders (4314, -0.19 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.54 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.23 DPS) [crafted]; Green Woolen Vest (2582, -0.24 DPS) [crafted]; Gray Woolen Robe (2585, -0.54 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Bright Bracers (3647, -0.03 DPS) [world_drop]; Windsong Bangles (263336, -0.08 DPS) [quest]; Tabitha's Cuffs (251486, -0.37 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.09 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.35 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.20 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.37 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 0.0 | yes | Silk-threaded Trousers (1929, -0.11 DPS) [dungeon]; Colorful Kilt (10048, -0.28 DPS) [crafted]; Abomination Skin Leggings (23173, -0.33 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Feather Padded Treads (285345, -0.28 DPS, sim-verified) [world]; Pristine Boots (253889, -0.36 DPS) [crafted]; Red Woolen Boots (4313, -0.38 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.23 DPS) [world]; Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Sludge-Stained Band (286535, -0.23 DPS, sim-verified) [world]; Loop of Sacrifice (281673, -0.25 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.0 | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Staff of Westfall (2042, -0.10 DPS) [quest]; Living Root (6631, -0.52 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 275.0 | yes | Firebelcher (5243, +0.00 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.03 DPS) [quest]; Sizzle Stick (8071, -4.50 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak; 9792 Ivycloth Boots

### Band 30 (gnome, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 57.8. Weights run: 0.7s. Verify run: 0.7s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.236 ± 0.191), crit=1.585 ± 0.050, hit=3.984 ± 0.329, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.34 DPS) [dungeon]; Silk Headband (7050, -0.36 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 | yes | Darkspear Warding Pendant (272075, -0.86 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Pendant of Myzrael (4614, -0.97 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 | yes | Death Speaker Mantle (6685, -0.09 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.34 DPS) [crafted]; Fairywing Mantle (9536, -0.34 DPS) [quest] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Prelacy Cape (7004, -0.11 DPS) [quest]; Caretaker's Cape (19533, -0.11 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 0.0 | yes | Death Speaker Robes (6682, -0.28 DPS) [dungeon]; Pristine Gown (253961, -0.39 DPS) [crafted]; Tree Bark Jacket (1486, -0.87 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.87 DPS) [quest]; Nightsky Wristbands (6407, -0.88 DPS, sim-verified) [world_drop]; Mindthrust Bracers (1974, -0.90 DPS) [dungeon] |
| hands | Serpent Gloves (5970) (or Shilly Mitts (9609)) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Shilly Mitts (9609, +0.00 DPS, sim-verified) [quest]; Town Clerk's Mittens (270029, -0.05 DPS) [quest]; Gnoll Casting Gloves (892, -0.11 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.7 | yes | Belt of Arugal (6392, +0.00 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.41 DPS) [crafted]; Ghamoo-ra's Bind (6908, -0.43 DPS) [dungeon] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.38 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.53 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.7 | yes | Acidic Walkers (9454, -0.20 DPS) [dungeon]; Nimbus Boots (6998, -0.31 DPS) [quest]; Spidersilk Boots (4320, -1.09 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.18 DPS) [quest]; Lorekeeper's Ring (20431, -0.23 DPS) [rep]; Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.34 DPS) [dungeon]; Sludge-Stained Band (286535, -0.34 DPS) [world]; Minor Channeling Ring (1449, -0.66 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Gnarled Ash Staff (791) | World drop [world_drop] | 222.8 | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Lorekeeper's Staff (212580, -0.80 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.80 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.8 | yes | Greater Mystic Wand (217287, -0.07 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest]; Scorching Wand (5213, -4.70 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Gnarled Ash Staff; ranged: Necrotic Wand

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 40 (gnome, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 179.7. Weights run: 0.9s. Verify run: 0.7s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.091 ± 0.015, crit=3.159 ± 0.088, hit=3.864 ± 0.105, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -3.02 DPS, sim-verified) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop]; Silk Headband (7050, -4.11 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 | yes | Darkspear Warding Pendant (272074, -2.37 DPS) [vendor]; Necklace of Calisea (1714, -2.40 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.43 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 | yes | Green Silken Shoulders (7057, -0.24 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Berylline Pads (4197, -0.65 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, -0.10 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Caretaker's Cape (19532, -0.34 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 | yes | Dreamweave Vest (10021, -1.23 DPS, sim-verified) [crafted]; Robe of Power (7054, -2.56 DPS) [crafted]; Tree Bark Jacket (1486, -3.27 DPS) [dungeon] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.69 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.21 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 14.4 | yes | Star Belt (4329, -0.51 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20099, -1.06 DPS) [rep]; Belt of Arugal (6392, -1.75 DPS) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 | yes | Gaze Dreamer Pants (6903, -1.24 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.03 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -3.28 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Nimbus Boots (6998, -6.17 DPS) [quest] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 | yes | Ring of Forlorn Spirits (2043, -0.87 DPS) [quest]; Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Ring of Forlorn Spirits (2043, -0.35 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.69 DPS) [quest]; Lorekeeper's Ring (19525, -0.69 DPS) [rep] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 146.0 | yes | Staff of Jordan (873, -1.29 DPS, sim-verified) [world_drop]; Celestial Stave (9517, -14.14 DPS) [quest]; Black Duskwood Staff (937, -16.49 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | 109.2 | yes | Icefury Wand (7514, +0.00 DPS, sim-verified) [quest]; Ragefire Wand (7513, -0.19 DPS) [quest]; Twisted Nether Wand (249144, -1.37 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Illusionary Rod; ranged: Nether Force Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs

### Band 50 (gnome, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 252.1. Weights run: 0.9s. Verify run: 0.8s. 723 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.020, crit=4.698 ± 0.137, hit=5.742 ± 0.243, spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 81.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -4.64 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.94 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 | yes | Mindburst Medallion (11196, -0.39 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.15 DPS) [world]; Darkspear Warding Pendant (272073, -2.34 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 74.7 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -20.95 DPS) [dungeon]; Black Mageweave Shoulders (10027, -22.34 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 | yes | Runecloth Cloak (13860, -1.82 DPS, sim-verified) [crafted]; Nightfall Drape (12465, -1.98 DPS) [dungeon]; Icy Cloak (4327, -2.68 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 79.0 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -19.73 DPS) [world_drop]; Acumen Robes (17775, -20.25 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Nethergeld Cuffs (254061, -0.44 DPS) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.45 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.56 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.56 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 75.3 | yes | Satyrmane Sash (17755, -0.22 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -21.33 DPS) [rep]; Ghostweave Cord (254073, -21.48 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 79.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -21.05 DPS) [crafted]; Red Mageweave Pants (10009, -22.34 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 66.4 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -14.85 DPS) [crafted]; Gilded Sandals (254107, -19.07 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 | yes | Philanthropist's Ring (281635, -16.39 DPS) [quest]; Ring of Forlorn Spirits (2043, -17.31 DPS) [quest]; Reedknot Ring (9622, -17.66 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Philanthropist's Ring (281635, -0.48 DPS, sim-verified) [quest]; Lorekeeper's Ring (19524, -1.05 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.40 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, -18.11 DPS) [quest]; Guardian Talisman (1490, -18.11 DPS) [quest]; Ankh of Life (1713, -18.11 DPS) [world_drop] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Thunderbrew's Boot Flask (744, -2.10 DPS) [quest]; Guardian Talisman (1490, -2.10 DPS) [quest]; Ankh of Life (1713, -2.10 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 | yes | Inventor's Focal Sword (17719, -11.08 DPS) [dungeon]; Illusionary Rod (7713, -11.38 DPS) [dungeon]; Soulkeeper (1607, -22.36 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 723, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 463.4. Weights run: 0.9s. Verify run: 0.8s. 1145 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.137 ± 0.028, crit=6.853 ± 0.198, hit=8.396 ± 0.345, spell_haste=4.452 ± 0.152, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 314.0 | yes | Bloodvine Goggles (19999, -17.36 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -34.27 DPS) [quest]; Field Marshal's Coronet (16441, -64.02 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 0.0 | yes | Medallion of the Dawn (22659, -29.42 DPS) [quest]; Beads of Ogre Might (22150, -33.61 DPS) [quest]; Charm of the Shifting Sands (21504, -53.69 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 114.7 | yes | Champion's Silk Mantle (23264, -0.80 DPS) [vendor]; Lieutenant Commander's Silk Mantle (23319, -0.80 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.80 DPS) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 95.9 | yes | Earthweave Cloak (21187, -0.43 DPS, sim-verified) [quest]; Drape of Vaulted Secrets (21415, -26.88 DPS) [quest]; Hide of the Wild (18510, -28.23 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 230.6 | yes | Bloodvine Vest (19682, -10.06 DPS, sim-verified) [crafted]; Enigma Robes (21343, -32.41 DPS) [quest]; Robe of the Archmage (14152, -32.59 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 111.0 | yes | Frostfire Bindings (22503, -5.64 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -30.78 DPS) [rep]; Dryad's Wrist Bindings (19596, -31.58 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 202.3 | yes | Sorcerer's Gloves (22066, -2.74 DPS, sim-verified) [quest]; Dreadmist Wraps (16705, -41.01 DPS) [dungeon]; Frostfire Gloves (22501, -57.34 DPS) [quest] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 118.1 | yes | Frostfire Belt (22502, +0.00 DPS, sim-verified) [quest]; Highlander's Cloth Girdle (20047, -2.58 DPS) [rep]; Highlander's Cloth Girdle (20097, -4.38 DPS) [rep] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 133.5 | yes | Marshal's Silk Leggings (16442, -1.70 DPS) [vendor]; General's Silk Trousers (16534, -1.70 DPS) [vendor]; Enigma Leggings (21346, -8.33 DPS, sim-verified) [quest] |
| feet | Frostfire Sandals (22500) | Frostfire Sandals [quest] | 126.4 | yes | Enigma Boots (21344, -0.61 DPS, sim-verified) [quest]; Marshal's Silk Footwraps (16437, -6.84 DPS) [vendor]; General's Silk Boots (16539, -6.84 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 179.9 | yes | Band of Earthen Might (21182, -10.35 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -20.66 DPS) [world]; Mindtear Band (20632, -21.42 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 0.0 | yes | Ritssyn's Ring of Chaos (21836, -0.30 DPS) [world]; Mindtear Band (20632, -1.06 DPS) [world]; Band of Earthen Might (21182, -12.70 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, -26.47 DPS) [quest]; Guardian Talisman (1490, -26.47 DPS) [quest]; Ankh of Life (1713, -26.47 DPS) [world_drop] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Thunderbrew's Boot Flask (744, -2.10 DPS) [quest]; Guardian Talisman (1490, -2.10 DPS) [quest]; Ankh of Life (1713, -2.10 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | 0.0 | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; High Warlord's Quickblade (234553, -28.04 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 0.0 | yes | Stormrager (16997, -0.14 DPS) [quest]; Wand of Biting Cold (19108, -4.49 DPS, sim-verified) [quest]; Torch of Austen (13004, -6.28 DPS) [world] |

**New at 60:** head: Frostfire Circlet; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Chromatic Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Belt of the Archmage; legs: Frostfire Leggings; feet: Frostfire Sandals; finger1: Don Julio's Band; finger2: Ring of the Fallen God; main_hand: Ironbark Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1145, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (orc, 253100000000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 29.9. Weights run: 0.7s. Verify run: 0.6s. 202 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.395 ± 0.090, crit=1.480 ± 0.048, hit=3.723 ± 0.108, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.35 DPS, sim-verified) [crafted]; Flying Tiger Goggles (4368, -0.49 DPS) [crafted]; Lucky Fishing Hat (19972, -0.49 DPS) [quest] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.6 | yes | Double-Stitched Woolen Shoulders (4314, -0.33 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.54 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Black Whelp Cloak (7283, -0.08 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.18 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.23 DPS) [crafted]; Green Woolen Vest (2582, -0.24 DPS) [crafted]; Gray Woolen Robe (2585, -0.44 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Owlbeard Bracers (16981, -0.02 DPS) [quest]; Tabitha's Cuffs (251486, -0.47 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.08 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.15 DPS) [crafted]; Apothecary Gloves (10919, -0.25 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Novice Ardent's Sash (253887, -0.20 DPS) [crafted]; Keller's Girdle (2911, -0.20 DPS) [world_drop]; Novice Arcanist's Sash (253885, -0.37 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 0.0 | yes | Silk-threaded Trousers (1929, -0.11 DPS) [dungeon]; Colorful Kilt (10048, -0.28 DPS) [crafted]; Abomination Skin Leggings (23173, -0.34 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Pristine Boots (253889, -0.36 DPS) [crafted]; Red Woolen Boots (4313, -0.38 DPS) [crafted]; Feather Padded Treads (285345, -0.41 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.22 DPS) [dungeon]; Loop of Sacrifice (281673, -0.25 DPS) [quest]; Black Pearl Ring (6332, -0.35 DPS) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, +0.00 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.08 DPS) [quest]; Black Pearl Ring (6332, -0.18 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.0 | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Crescent Staff (6505, -0.42 DPS) [quest]; Living Root (6631, -0.53 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 275.0 | yes | Firebelcher (5243, -0.04 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.03 DPS) [quest]; Sizzle Stick (8071, -4.50 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 202, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (orc, 253225110000000000-00000000000000000-0000000000000000000)

Set DPS (verified): 50.8. Weights run: 0.7s. Verify run: 0.7s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.236 ± 0.191), crit=1.585 ± 0.050, hit=3.984 ± 0.329, spell_haste=not significant (0.000 ± 0.000), spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.34 DPS) [dungeon]; Silk Headband (7050, -0.47 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.4 | yes | Darkspear Warding Pendant (272075, -0.78 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Pendant of Myzrael (4614, -0.97 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.1 | yes | Death Speaker Mantle (6685, -0.03 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.34 DPS) [crafted]; Fairywing Mantle (9536, -0.34 DPS) [quest] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.11 DPS) [crafted]; Battle Healer's Cloak (19529, -0.11 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.39 DPS) [dungeon]; Pristine Gown (253961, -0.50 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -0.87 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.87 DPS) [quest]; Owlbeard Bracers (16981, -1.06 DPS, sim-verified) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.2 | yes | Gnoll Casting Gloves (892, -0.14 DPS) [world]; Serpent Gloves (5970, -0.16 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.17 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.7 | yes | Warsong Sash (16975, -0.16 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.23 DPS) [dungeon]; Invoker's Cord (215366, -0.41 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.38 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.53 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 8.7 | yes | Acidic Walkers (9454, -0.20 DPS) [dungeon]; Boots of the Enchanter (4325, -0.42 DPS) [crafted]; Spidersilk Boots (4320, -1.06 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.23 DPS) [rep]; Electrocutioner Lagnut (9447, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.34 DPS) [world]; Sacred Band (6669, -0.46 DPS) [quest]; Electrocutioner Lagnut (9447, -0.63 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Gnarled Ash Staff (791) | World drop [world_drop] | 222.8 | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Lorekeeper's Staff (212580, -0.80 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.80 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 291.8 | yes | Greater Mystic Wand (217287, -0.94 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.55 DPS) [quest]; Scorching Wand (5213, -4.70 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Gnarled Ash Staff; ranged: Necrotic Wand

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (orc, 253225111100011501-00000000000000000-0000000000000000000)

Set DPS (verified): 172.8. Weights run: 0.9s. Verify run: 0.8s. 567 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.091 ± 0.015, crit=3.159 ± 0.088, hit=3.864 ± 0.105, spell_haste=2.043 ± 0.072, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -2.97 DPS, sim-verified) [world]; Holy Shroud (2721, -3.43 DPS) [world_drop]; Silk Headband (7050, -4.11 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.5 | yes | Necklace of Calisea (1714, -2.33 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -2.37 DPS) [vendor]; Darkspear Warding Pendant (272075, -2.43 DPS) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 9.8 | yes | Green Silken Shoulders (7057, -0.28 DPS, sim-verified) [crafted]; Inquisitor's Shawl (19507, -0.56 DPS) [dungeon]; Chestnut Mantle (17695, -0.62 DPS) [quest] |
| back | Icy Cloak (4327) | Tailoring [crafted] | 7.0 | yes | Long Silken Cloak (4326, -0.17 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.19 DPS) [crafted]; Battle Healer's Cloak (19528, -0.34 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 22.5 | yes | Dreamweave Vest (10021, -1.26 DPS, sim-verified) [crafted]; Robe of Power (7054, -2.56 DPS) [crafted]; Tree Bark Jacket (1486, -3.27 DPS) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Condor Bracers (15864, -0.67 DPS, sim-verified) [quest]; Radiant Silver Bracers (4545, -1.46 DPS) [quest]; Earthen Silk Cuffs (254019, -1.71 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.23 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -2.21 DPS) [crafted]; Gilded Handwraps (254021, -3.33 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 14.4 | yes | Star Belt (4329, -0.56 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -1.06 DPS) [rep]; Warsong Sash (16975, -1.15 DPS) [quest] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 15.1 | yes | Gaze Dreamer Pants (6903, -1.27 DPS, sim-verified) [dungeon]; Abomination Skin Leggings (23173, -1.84 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -2.03 DPS) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -3.23 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -5.70 DPS) [crafted]; Acidic Walkers (9454, -6.26 DPS) [dungeon] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 10.5 | yes | Reedknot Ring (9622, -1.22 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.56 DPS) [vendor]; Electrocutioner Lagnut (9447, -2.59 DPS) [dungeon] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Reedknot Ring (9622, -0.67 DPS, sim-verified) [quest]; Advisor's Ring (19521, -0.69 DPS) [rep]; Sea Giant's Toe Ring (274746, -1.03 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 146.0 | yes | Staff of Jordan (873, -1.29 DPS, sim-verified) [world_drop]; Celestial Stave (9517, -14.14 DPS) [quest]; Black Duskwood Staff (937, -16.49 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | 109.2 | yes | Icefury Wand (7514, +0.00 DPS, sim-verified) [quest]; Ragefire Wand (7513, -0.19 DPS) [quest]; Twisted Nether Wand (249144, -1.37 DPS) [crafted] |

**New at 40:** head: Spellpower Goggles Xtreme; back: Icy Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Illusionary Rod; ranged: Nether Force Wand

No-known-source sample (15 of 567, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (orc, 253225111100011501-23500000000000000-0000000000000000000)

Set DPS (verified): 242.3. Weights run: 0.9s. Verify run: 0.8s. 722 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.108 ± 0.020, crit=4.698 ± 0.137, hit=5.742 ± 0.243, spell_haste=3.045 ± 0.104, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 81.1 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -4.64 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -18.94 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.6 | yes | Mindburst Medallion (11196, -0.37 DPS, sim-verified) [quest]; Horizon Choker (13085, -2.15 DPS) [world]; Darkspear Warding Pendant (272073, -2.34 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 74.7 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -20.95 DPS) [dungeon]; Black Mageweave Shoulders (10027, -22.34 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.6 | yes | Deep Woodlands Cloak (19121, -0.62 DPS, sim-verified) [quest]; Runecloth Cloak (13860, -1.68 DPS) [crafted]; Nightfall Drape (12465, -1.98 DPS) [dungeon] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 79.0 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -19.73 DPS) [world_drop]; Acumen Robes (17775, -20.25 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, +0.00 DPS, sim-verified) [crafted]; Condor Bracers (15864, -0.70 DPS) [quest]; Bloodband Bracers (11469, -1.06 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.4 | yes | Black Mageweave Gloves (10003, -1.37 DPS, sim-verified) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.56 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -1.56 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 75.3 | yes | Satyrmane Sash (17755, -0.08 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -21.33 DPS) [rep]; Ghostweave Cord (254073, -21.48 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 79.1 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -21.05 DPS) [crafted]; Red Mageweave Pants (10009, -22.34 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 66.4 | yes | Sergeant Major's Dreadweave Boots (220891, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -14.85 DPS) [crafted]; Gilded Sandals (254107, -19.07 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 | yes | Philanthropist's Ring (281635, -16.39 DPS) [quest]; Reedknot Ring (9622, -17.66 DPS) [quest]; Sea Giant's Toe Ring (274746, -18.01 DPS) [vendor] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Philanthropist's Ring (281635, -0.40 DPS, sim-verified) [quest]; Advisor's Ring (19520, -1.05 DPS) [rep]; Reedknot Ring (9622, -1.75 DPS) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Uther's Strength (11302, -0.12 DPS, sim-verified) [world]; Guardian Talisman (1490, -18.11 DPS) [quest]; Ankh of Life (1713, -18.11 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world]; Guardian Talisman (1490, -14.08 DPS) [quest]; Ankh of Life (1713, -14.08 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 | yes | Inventor's Focal Sword (17719, -11.08 DPS) [dungeon]; Illusionary Rod (7713, -11.38 DPS) [dungeon]; Soulkeeper (1607, -22.36 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 149.9 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.95 DPS) [crafted]; Wand of Allistarj (13065, -4.08 DPS) [world] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 722, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 60 (orc, 253225111100011501-23552300000000000-0000000000000000000)

Set DPS (verified): 440.8. Weights run: 0.9s. Verify run: 0.8s. 1144 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.002, intellect=0.137 ± 0.028, crit=6.853 ± 0.198, hit=8.396 ± 0.345, spell_haste=4.452 ± 0.152, spell_penetration=not significant (0.000 ± 0.000), arcane_power=1.000 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 314.0 | yes | Bloodvine Goggles (19999, -16.14 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -34.27 DPS) [quest]; Field Marshal's Coronet (16441, -64.02 DPS) [vendor] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 0.0 | yes | Medallion of the Dawn (22659, -29.42 DPS) [quest]; Beads of Ogre Might (22150, -33.61 DPS) [quest]; Charm of the Shifting Sands (21504, -53.69 DPS) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 114.7 | yes | Champion's Silk Mantle (23264, -0.80 DPS) [vendor]; Lieutenant Commander's Silk Mantle (23319, -0.80 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.80 DPS) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 95.9 | yes | Earthweave Cloak (21187, +0.00 DPS, sim-verified) [quest]; Drape of Vaulted Secrets (21415, -26.88 DPS) [quest]; Hide of the Wild (18510, -28.23 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 230.6 | yes | Bloodvine Vest (19682, -9.20 DPS, sim-verified) [crafted]; Enigma Robes (21343, -32.41 DPS) [quest]; Robe of the Archmage (14152, -32.59 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 111.0 | yes | Frostfire Bindings (22503, -4.46 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -30.78 DPS) [rep]; Dryad's Wrist Bindings (19596, -31.58 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 202.3 | yes | Sorcerer's Gloves (22066, -1.94 DPS, sim-verified) [quest]; Dreadmist Wraps (16705, -41.01 DPS) [dungeon]; Frostfire Gloves (22501, -57.34 DPS) [quest] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 118.1 | yes | Frostfire Belt (22502, +0.00 DPS, sim-verified) [quest]; Defiler's Cloth Girdle (20163, -2.58 DPS) [rep]; Defiler's Cloth Girdle (20165, -4.38 DPS) [rep] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 133.5 | yes | Marshal's Silk Leggings (16442, -1.70 DPS) [vendor]; General's Silk Trousers (16534, -1.70 DPS) [vendor]; Enigma Leggings (21346, -7.02 DPS, sim-verified) [quest] |
| feet | Frostfire Sandals (22500) | Frostfire Sandals [quest] | 126.4 | yes | Enigma Boots (21344, +0.00 DPS, sim-verified) [quest]; Marshal's Silk Footwraps (16437, -6.84 DPS) [vendor]; General's Silk Boots (16539, -6.84 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 179.9 | yes | Band of Earthen Might (21182, -9.04 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -20.66 DPS) [world]; Mindtear Band (20632, -21.42 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 0.0 | yes | Ritssyn's Ring of Chaos (21836, -0.30 DPS) [world]; Mindtear Band (20632, -1.06 DPS) [world]; Band of Earthen Might (21182, -12.22 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, -5.88 DPS) [quest]; Guardian Talisman (1490, -26.47 DPS) [quest]; Ankh of Life (1713, -26.47 DPS) [world_drop] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Rune of the Guard Captain (19120, -0.33 DPS, sim-verified) [quest]; Guardian Talisman (1490, -2.10 DPS) [quest]; Ankh of Life (1713, -2.10 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | 0.0 | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; High Warlord's Quickblade (234553, -28.04 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | 182.7 | yes | Brilliant Wand (249385, +0.00 DPS, sim-verified) [crafted]; Stormrager (16997, -1.07 DPS) [quest]; Torch of Austen (13004, -7.21 DPS) [world] |

**New at 60:** head: Frostfire Circlet; neck: Onyxia Tooth Pendant; shoulder: Mantle of the Timbermaw; back: Chromatic Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Belt of the Archmage; legs: Frostfire Leggings; feet: Frostfire Sandals; finger1: Don Julio's Band; finger2: Ring of the Fallen God; trinket2: Uther's Strength; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1144, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

