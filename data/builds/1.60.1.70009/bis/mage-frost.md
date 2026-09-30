# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 36.7. Weights run: 0.6s. Verify run: 0.5s. 203 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.741 ± 0.174, crit=0.951 ± 0.045, hit=2.241 ± 0.102, spell_haste=not significant (-0.108 ± 0.181), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.78 DPS) [crafted]; Lucky Fishing Hat (19972, -0.78 DPS) [quest]; Shadow Goggles (4373, -1.66 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 8.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.53 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.03 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.16 DPS) [dungeon]; Black Whelp Cloak (7283, -0.16 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.7 | yes | Green Woolen Robe (6243, -0.45 DPS) [crafted]; Seer's Robe (2981, -0.55 DPS) [world_drop]; Gray Woolen Robe (2585, -0.87 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.19 DPS) [world_drop]; Repurposed Hair Band (281256, -0.38 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.13 DPS) [world]; Blight Gloves (279877, -0.24 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.0 | yes | Keller's Girdle (2911, -0.13 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.36 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.75 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.9 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.03 DPS) [dungeon]; Colorful Kilt (10048, -1.29 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 | yes | Pristine Boots (253889, -0.32 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.64 DPS) [world]; Red Woolen Boots (4313, -0.77 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 | yes | Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Loop of Sacrifice (281673, -0.36 DPS) [quest]; Sludge-Stained Band (286535, -0.45 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.55 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 165.1 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Living Root (6631, -0.19 DPS) [dungeon]; Staff of Westfall (2042, -0.42 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 175.1 | yes | Firebelcher (5243, -0.43 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.17 DPS) [quest]; Sizzle Stick (8071, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak; 9792 Ivycloth Boots

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 63.5. Weights run: 0.7s. Verify run: 0.6s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.313 ± 0.311), crit=1.967 ± 0.110, hit=3.120 ± 0.172, spell_haste=not significant (0.830 ± 0.313), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, -0.07 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 | yes | Crystal Starfire Medallion (5003, -1.15 DPS) [world_drop]; Pendant of Myzrael (4614, -1.34 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.36 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.8 | yes | Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.49 DPS) [crafted]; Death Speaker Mantle (6685, -0.64 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Prelacy Cape (7004, -0.15 DPS) [quest]; Repairman's Cape (9605, -0.27 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.1 | yes | Death Speaker Robes (6682, -0.40 DPS) [dungeon]; Pristine Gown (253961, -0.58 DPS) [crafted]; Tree Bark Jacket (1486, -1.19 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -1.07 DPS) [quest]; Mindthrust Bracers (1974, -1.12 DPS) [dungeon]; Nightsky Wristbands (6407, -1.93 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.4 | yes | Shilly Mitts (9609, -0.07 DPS) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Serpent Gloves (5970, -1.26 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 | yes | Belt of Arugal (6392, -0.30 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.51 DPS) [crafted]; Crimson Silk Belt (7055, -0.56 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.42 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.2 | yes | Acidic Walkers (9454, -0.25 DPS) [dungeon]; Nimbus Boots (6998, -0.48 DPS) [quest]; Spidersilk Boots (4320, -1.71 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.21 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Sludge-Stained Band (286535, -0.45 DPS) [world]; Minor Channeling Ring (1449, -1.43 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 0.0 | yes | Lorekeeper's Staff (212580, -0.13 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.13 DPS) [vendor]; Gnarled Ash Staff (791, -0.88 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 223.5 | yes | Greater Mystic Wand (217287, -0.27 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest]; Scorching Wand (5213, -4.81 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 94.4. Weights run: 0.7s. Verify run: 0.6s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.530 ± 0.476), crit=2.880 ± 0.172, hit=4.181 ± 0.341, spell_haste=not significant (1.576 ± 0.546), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -1.43 DPS) [crafted]; Holy Shroud (2721, -1.47 DPS) [world_drop]; Augural Shroud (2620, -1.76 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 | yes | Necklace of Calisea (1714, -0.80 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.95 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.11 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.9 | yes | Bloodmage Mantle (7684, -0.02 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest]; Green Silken Shoulders (7057, -1.52 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.7 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.24 DPS) [crafted]; Caretaker's Cape (19532, -0.39 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.2 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.71 DPS) [crafted]; Crimson Silk Vest (7058, -1.16 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest]; Aurora Bracers (4043, -0.70 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 | yes | Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Red Mageweave Gloves (10018, -1.04 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.24 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.1 | yes | Star Belt (4329, -0.46 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.52 DPS) [rep]; Deathmage Sash (10771, -1.05 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.4 | yes | Crimson Silk Pantaloons (7062, -1.02 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.05 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.23 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -1.61 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.17 DPS) [dungeon]; Spidersilk Boots (4320, -2.19 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 | yes | Ring of Forlorn Spirits (2043, -0.76 DPS) [quest]; Reedknot Ring (9622, -0.91 DPS) [quest]; Minor Channeling Ring (1449, -1.05 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.29 DPS) [quest]; Lorekeeper's Ring (19525, -0.29 DPS) [rep]; Ring of Forlorn Spirits (2043, -2.34 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 280.6 | yes | Illusionary Rod (7713, -0.90 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -4.97 DPS) [quest]; Windweaver Staff (7757, -7.68 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | 255.4 | yes | Ragefire Wand (7513, -0.19 DPS) [quest]; Icefury Wand (7514, -1.07 DPS, sim-verified) [quest]; Umbral Wand (5216, -1.89 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 121.6. Weights run: 0.6s. Verify run: 0.7s. 722 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.275 ± 0.650), crit=5.153 ± 0.298, hit=7.168 ± 0.493, spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 101.4 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.67 DPS) [dungeon]; Red Mageweave Headband (10033, -7.47 DPS) [crafted] |
| neck | Horizon Choker (13085) | Azuregos [world] | 17.8 | yes | Scorn's Icy Choker (23169, -0.08 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.55 DPS) [quest]; Darkspear Warding Pendant (272073, -0.84 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 91.6 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -7.30 DPS) [dungeon]; Red Mageweave Shoulders (10029, -8.59 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.6 | yes | Darkspear Raider's Cloak (272076, -0.50 DPS) [vendor]; Big Voodoo Cloak (8216, -0.68 DPS) [crafted]; Runecloth Cloak (13860, -1.15 DPS, sim-verified) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 98.2 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Acumen Robes (17775, -7.04 DPS) [quest]; Runecloth Robe (13858, -8.59 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 0.0 | yes | Shizzle's Nozzle Wiper (11917, -0.08 DPS) [quest]; Imperial Red Bracers (8247, -0.25 DPS) [dungeon]; Bloodband Bracers (11469, -1.75 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 101.5 | yes | Raider Handwraps (272098, -0.40 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -10.11 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -10.11 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 87.5 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.97 DPS) [dungeon]; Deathmage Sash (10771, -8.05 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 99.4 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Red Mageweave Pants (10009, -9.20 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -9.64 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 91.2 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -8.81 DPS) [crafted]; Gilded Sandals (254107, -9.01 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 71.7 | yes | Lorekeeper's Ring (19523, -7.83 DPS) [rep]; Lorekeeper's Ring (19524, -8.22 DPS) [rep]; Ogremind Ring (1993, -8.23 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.6 | yes | Lorekeeper's Ring (19523, +0.00 DPS, sim-verified) [rep]; Lorekeeper's Ring (19524, -1.13 DPS) [rep]; Ogremind Ring (1993, -1.14 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Ankh of Life (1713, -0.59 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -8.46 DPS) [quest]; Guardian Talisman (1490, -8.46 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Ankh of Life (1713, -0.50 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -0.79 DPS) [quest]; Guardian Talisman (1490, -0.79 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 | yes | Soulkeeper (1607, -8.91 DPS) [world_drop]; Radiant Staff (249453, -10.82 DPS) [crafted]; Spire of Hakkar (10844, -10.86 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world]; Lesser Eternal Wand (249232, -5.70 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 722, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 209.2. Weights run: 0.7s. Verify run: 0.7s. 1144 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.124 ± 1.054), crit=8.680 ± 0.479, hit=12.852 ± 0.837, spell_haste=not significant (1.136 ± 1.257), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 409.4 | yes | Bloodvine Goggles (19999, -5.64 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -15.41 DPS) [quest]; Field Marshal's Coronet (16441, -31.57 DPS) [vendor] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | 0.0 | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Onyxia Tooth Pendant (18404, -1.75 DPS, sim-verified) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 140.1 | yes | Champion's Silk Mantle (23264, -0.28 DPS) [vendor]; Lieutenant Commander's Silk Mantle (23319, -0.28 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.28 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 128.5 | yes | Chromatic Cloak (18509, -0.67 DPS, sim-verified) [crafted]; Drape of Vaulted Secrets (21415, -13.67 DPS) [quest]; Hide of the Wild (18510, -14.15 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 300.4 | yes | Bloodvine Vest (19682, -6.66 DPS, sim-verified) [crafted]; Enigma Robes (21343, -17.11 DPS) [quest]; Robe of the Archmage (14152, -17.16 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 155.5 | yes | Frostfire Bindings (22503, +0.00 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -16.55 DPS) [rep]; Dryad's Wrist Bindings (19596, -16.83 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 253.3 | yes | Sorcerer's Gloves (22066, -0.20 DPS, sim-verified) [quest]; Sorcerer's Gauntlets (226930, -13.87 DPS) [vendor]; Dreadmist Wraps (16705, -15.44 DPS) [dungeon] |
| waist | Frostfire Belt (22502) | Frostfire Belt [quest] | 159.1 | yes | Highlander's Cloth Girdle (20047, -2.86 DPS) [rep]; Highlander's Cloth Girdle (20097, -3.50 DPS) [rep]; Belt of the Archmage (18405, -5.36 DPS, sim-verified) [crafted] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 177.7 | yes | Enigma Leggings (21346, -2.37 DPS) [quest]; Marshal's Silk Leggings (16442, -2.97 DPS) [vendor]; Bloodvine Leggings (19683, -6.60 DPS, sim-verified) [crafted] |
| feet | Enigma Boots (21344) | Enigma Boots [quest] | 158.4 | yes | Frostfire Sandals (22500, -0.59 DPS, sim-verified) [quest]; Marshal's Silk Footwraps (16437, -0.89 DPS) [vendor]; Marshal's Silk Footwraps (231606, -0.89 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 250.0 | yes | Band of Earthen Might (21182, -3.99 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -12.93 DPS) [world]; Mindtear Band (20632, -13.21 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 0.0 | yes | Ritssyn's Ring of Chaos (21836, -2.47 DPS) [world]; Mindtear Band (20632, -2.75 DPS) [world]; Band of Earthen Might (21182, -6.88 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Ankh of Life (1713, -2.25 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -14.45 DPS) [quest]; Guardian Talisman (1490, -14.45 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Thunderbrew's Boot Flask (744, -0.75 DPS) [quest]; Guardian Talisman (1490, -0.75 DPS) [quest]; Ankh of Life (1713, -1.10 DPS, sim-verified) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | 0.0 | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; High Warlord's Quickblade (234553, -9.27 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | Korrak the Bloodrager [quest] | 512.4 | yes | Stormrager (16997, +0.00 DPS, sim-verified) [quest]; Brilliant Wand (249385, -3.19 DPS) [crafted]; Torch of Austen (13004, -7.21 DPS) [world] |

**New at 60:** head: Frostfire Circlet; neck: Jewel of Kajaro; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Frostfire Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1144, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 33.9. Weights run: 0.6s. Verify run: 0.6s. 202 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.741 ± 0.174, crit=0.951 ± 0.045, hit=2.241 ± 0.102, spell_haste=not significant (-0.108 ± 0.181), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.78 DPS) [crafted]; Lucky Fishing Hat (19972, -0.78 DPS) [quest]; Shadow Goggles (4373, -1.19 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 8.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.66 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.03 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.16 DPS) [dungeon]; Black Whelp Cloak (7283, -0.16 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.7 | yes | Green Woolen Robe (6243, -0.45 DPS) [crafted]; Seer's Robe (2981, -0.55 DPS) [world_drop]; Gray Woolen Robe (2585, -0.80 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 0.0 | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.10 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.47 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | 0.0 | yes | Gnoll Casting Gloves (892, -0.03 DPS) [world]; Blight Gloves (279877, -0.13 DPS) [quest]; Serpent Gloves (5970, -0.51 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.0 | yes | Keller's Girdle (2911, -0.13 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.36 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.82 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.9 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.03 DPS) [dungeon]; Colorful Kilt (10048, -1.29 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 | yes | Pristine Boots (253889, -0.43 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.64 DPS) [world]; Red Woolen Boots (4313, -0.77 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Black Pearl Ring (6332, -0.46 DPS) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.4 | yes | Sludge-Stained Band (286535, -0.19 DPS) [world]; Loop of Sacrifice (281673, -0.29 DPS, sim-verified) [quest]; Black Pearl Ring (6332, -0.38 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 165.1 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Living Root (6631, -0.19 DPS) [dungeon]; Crescent Staff (6505, -1.06 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 175.1 | yes | Firebelcher (5243, -0.27 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.17 DPS) [quest]; Sizzle Stick (8071, -4.40 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 202, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 57.2. Weights run: 0.7s. Verify run: 0.6s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.313 ± 0.311), crit=1.967 ± 0.110, hit=3.120 ± 0.172, spell_haste=not significant (0.830 ± 0.313), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 | yes | Crystal Starfire Medallion (5003, -1.15 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.30 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.34 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.8 | yes | Fairywing Mantle (9536, -0.45 DPS) [quest]; Death Speaker Mantle (6685, -0.49 DPS, sim-verified) [dungeon]; Invoker's Mantle (215365, -0.49 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.1 | yes | Death Speaker Robes (6682, -0.40 DPS) [dungeon]; Pristine Gown (253961, -0.58 DPS) [crafted]; Tree Bark Jacket (1486, -1.30 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -1.07 DPS) [quest]; Owlbeard Bracers (16981, -1.11 DPS) [quest]; Nightsky Wristbands (6407, -1.51 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.6 | yes | Gnoll Casting Gloves (892, -0.24 DPS) [world]; Truefaith Gloves (7049, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.70 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 | yes | Belt of Arugal (6392, -0.30 DPS) [dungeon]; Warsong Sash (16975, -0.31 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.51 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -0.02 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.42 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.2 | yes | Acidic Walkers (9454, -0.25 DPS) [dungeon]; Boots of the Enchanter (4325, -0.63 DPS) [crafted]; Spidersilk Boots (4320, -1.69 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.30 DPS) [rep]; Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon]; Sludge-Stained Band (286535, -0.60 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.45 DPS) [world]; Black Widow Band (6199, -0.57 DPS) [world]; Electrocutioner Lagnut (9447, -1.20 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Gnarled Ash Staff (791) | World drop [world_drop] | 170.1 | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Lorekeeper's Staff (212580, -0.68 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.68 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 223.5 | yes | Greater Mystic Wand (217287, -0.86 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest]; Scorching Wand (5213, -4.81 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Gnarled Ash Staff; ranged: Necrotic Wand

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 85.6. Weights run: 0.7s. Verify run: 0.7s. 567 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.530 ± 0.476), crit=2.880 ± 0.172, hit=4.181 ± 0.341, spell_haste=not significant (1.576 ± 0.546), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | 0.0 | yes | Enchanter's Cowl (4322, -0.74 DPS) [crafted]; Holy Shroud (2721, -0.78 DPS) [world_drop]; Spellpower Goggles Xtreme (10502, -0.94 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 | yes | Necklace of Calisea (1714, -0.46 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.95 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.11 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.9 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.02 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.7 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.24 DPS) [crafted]; Battle Healer's Cloak (19528, -0.39 DPS) [rep] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | 0.0 | yes | Robe of Power (7054, -0.35 DPS) [crafted]; Crimson Silk Vest (7058, -0.80 DPS) [crafted]; Robe of the Magi (1716, -1.29 DPS, sim-verified) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Radiant Silver Bracers (4545, +0.00 DPS, sim-verified) [quest]; Condor Bracers (15864, -0.29 DPS) [quest]; Aurora Bracers (4043, -0.70 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 | yes | Red Mageweave Gloves (10018, +0.00 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Gilded Handwraps (254021, -1.24 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 0.0 | yes | Star Belt (4329, -0.29 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.35 DPS) [rep]; Defiler's Cloth Girdle (20166, -1.59 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.4 | yes | Crimson Silk Pantaloons (7062, -0.64 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.05 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.23 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, +0.00 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.17 DPS) [dungeon]; Spidersilk Boots (4320, -2.19 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 | yes | Reedknot Ring (9622, -0.91 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.06 DPS) [vendor]; Ogremind Ring (1993, -1.39 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.29 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.44 DPS) [vendor]; Reedknot Ring (9622, -0.53 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 280.6 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Celestial Stave (9517, -4.97 DPS) [quest]; Windweaver Staff (7757, -7.68 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Nether Force Wand (11263) | Mage's Wand [quest] | 255.4 | yes | Icefury Wand (7514, +0.00 DPS, sim-verified) [quest]; Ragefire Wand (7513, -0.19 DPS) [quest]; Umbral Wand (5216, -1.89 DPS) [world_drop] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Staff of Jordan; ranged: Nether Force Wand

No-known-source sample (15 of 567, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 108.3. Weights run: 0.6s. Verify run: 0.7s. 721 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.275 ± 0.650), crit=5.153 ± 0.298, hit=7.168 ± 0.493, spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 101.4 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.67 DPS) [dungeon]; Red Mageweave Headband (10033, -7.47 DPS) [crafted] |
| neck | Horizon Choker (13085) | Azuregos [world] | 17.8 | yes | Scorn's Icy Choker (23169, -0.13 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.55 DPS) [quest]; Darkspear Warding Pendant (272073, -0.84 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 91.6 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -7.30 DPS) [dungeon]; Red Mageweave Shoulders (10029, -8.59 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 23.5 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.56 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.74 DPS) [vendor] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 98.2 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Acumen Robes (17775, -7.04 DPS) [quest]; Runecloth Robe (13858, -8.59 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | 0.0 | yes | Shizzle's Nozzle Wiper (11917, -0.08 DPS) [quest]; Radiant Silver Bracers (4545, -0.23 DPS) [quest]; Bloodband Bracers (11469, -1.42 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 101.5 | yes | Raider Handwraps (272098, -0.15 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -10.11 DPS) [vendor]; First Sergeant's Dreadweave Gloves (220908, -10.11 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 87.5 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.97 DPS) [dungeon]; Deathmage Sash (10771, -8.05 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 99.4 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Red Mageweave Pants (10009, -9.20 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -9.64 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 91.2 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -8.81 DPS) [crafted]; Gilded Sandals (254107, -9.01 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 71.7 | yes | Advisor's Ring (19519, -7.83 DPS) [rep]; Advisor's Ring (19520, -8.22 DPS) [rep]; Ogremind Ring (1993, -8.23 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.6 | yes | Advisor's Ring (19519, +0.00 DPS, sim-verified) [rep]; Advisor's Ring (19520, -1.13 DPS) [rep]; Ogremind Ring (1993, -1.14 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Ankh of Life (1713, -0.39 DPS, sim-verified) [world_drop]; Uther's Strength (11302, -7.68 DPS) [world]; Guardian Talisman (1490, -8.46 DPS) [quest] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Ankh of Life (1713, -0.04 DPS, sim-verified) [world_drop]; Uther's Strength (11302, -5.80 DPS) [world]; Guardian Talisman (1490, -6.58 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 | yes | Soulkeeper (1607, -8.91 DPS) [world_drop]; Radiant Staff (249453, -10.82 DPS) [crafted]; Spire of Hakkar (10844, -10.86 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world]; Lesser Eternal Wand (249232, -5.70 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Deep Woodlands Cloak; chest: Knight's Dreadweave Vest; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 721, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 188.6. Weights run: 0.7s. Verify run: 0.7s. 1143 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.124 ± 1.054), crit=8.680 ± 0.479, hit=12.852 ± 0.837, spell_haste=not significant (1.136 ± 1.257), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Frostfire Circlet (22498) | Frostfire Circlet [quest] | 409.4 | yes | Bloodvine Goggles (19999, -11.93 DPS, sim-verified) [crafted]; Enigma Circlet (21347, -15.41 DPS) [quest]; Field Marshal's Coronet (16441, -31.57 DPS) [vendor] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | 0.0 | yes | Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Onyxia Tooth Pendant (18404, -2.33 DPS, sim-verified) [quest] |
| shoulder | Mantle of the Timbermaw (19050) | Tailoring [crafted] | 140.1 | yes | Champion's Silk Mantle (23264, -0.28 DPS) [vendor]; Lieutenant Commander's Silk Mantle (23319, -0.28 DPS) [vendor]; Lieutenant Commander's Silk Mantle (227102, -0.28 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 128.5 | yes | Chromatic Cloak (18509, -0.31 DPS, sim-verified) [crafted]; Drape of Vaulted Secrets (21415, -13.67 DPS) [quest]; Hide of the Wild (18510, -14.15 DPS) [crafted] |
| chest | Frostfire Robe (22496) | Frostfire Robe [quest] | 300.4 | yes | Bloodvine Vest (19682, -7.02 DPS, sim-verified) [crafted]; Enigma Robes (21343, -17.11 DPS) [quest]; Robe of the Archmage (14152, -17.16 DPS) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 155.5 | yes | Frostfire Bindings (22503, +0.00 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -16.55 DPS) [rep]; Dryad's Wrist Bindings (19596, -16.83 DPS) [rep] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 253.3 | yes | Sorcerer's Gloves (22066, -0.20 DPS, sim-verified) [quest]; Sorcerer's Gauntlets (226930, -13.87 DPS) [vendor]; Dreadmist Wraps (16705, -15.44 DPS) [dungeon] |
| waist | Frostfire Belt (22502) | Frostfire Belt [quest] | 159.1 | yes | Defiler's Cloth Girdle (20163, -2.86 DPS) [rep]; Defiler's Cloth Girdle (20165, -3.50 DPS) [rep]; Belt of the Archmage (18405, -4.35 DPS, sim-verified) [crafted] |
| legs | Frostfire Leggings (22497) | Frostfire Leggings [quest] | 177.7 | yes | Enigma Leggings (21346, -2.37 DPS) [quest]; Marshal's Silk Leggings (16442, -2.97 DPS) [vendor]; Bloodvine Leggings (19683, -6.85 DPS, sim-verified) [crafted] |
| feet | Enigma Boots (21344) | Enigma Boots [quest] | 158.4 | yes | Marshal's Silk Footwraps (16437, -0.89 DPS) [vendor]; General's Silk Boots (16539, -0.89 DPS) [vendor]; Frostfire Sandals (22500, -0.90 DPS, sim-verified) [quest] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 250.0 | yes | Band of Earthen Might (21182, -3.12 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -12.93 DPS) [world]; Mindtear Band (20632, -13.21 DPS) [world] |
| finger2 | Ring of the Fallen God (21709) | The Savior of Kalimdor [quest] | 0.0 | yes | Ritssyn's Ring of Chaos (21836, -2.47 DPS) [world]; Mindtear Band (20632, -2.75 DPS) [world]; Band of Earthen Might (21182, -5.56 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Uther's Strength (11302, -0.47 DPS, sim-verified) [world]; Guardian Talisman (1490, -14.45 DPS) [quest]; Ankh of Life (1713, -14.45 DPS) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Uther's Strength (11302, -0.38 DPS, sim-verified) [world]; Guardian Talisman (1490, -11.24 DPS) [quest]; Ankh of Life (1713, -11.24 DPS) [world_drop] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | 0.0 | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; High Warlord's Quickblade (234553, -9.27 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Wand of Biting Cold (19108) | The Legend of Korrak [quest] | 512.4 | yes | Stormrager (16997, +0.00 DPS, sim-verified) [quest]; Brilliant Wand (249385, -3.19 DPS) [crafted]; Torch of Austen (13004, -7.21 DPS) [world] |

**New at 60:** head: Frostfire Circlet; neck: Jewel of Kajaro; shoulder: Mantle of the Timbermaw; back: Earthweave Cloak; chest: Frostfire Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Frostfire Belt; legs: Frostfire Leggings; feet: Enigma Boots; finger1: Don Julio's Band; finger2: Ring of the Fallen God; main_hand: Ironbark Staff; ranged: Wand of Biting Cold

No-known-source sample (15 of 1143, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

