# Leveling BiS: Demonology

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 35.1. Weights run: 1.5s. Verify run: 1.1s. 204 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.103, intellect=not significant (-0.337 ± 0.116), crit=0.435 ± 0.021, hit=1.467 ± 0.069, spell_haste=0.633 ± 0.114, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.103, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.81 DPS) [crafted]; Lucky Fishing Hat (19972, -0.81 DPS) [quest]; Flying Tiger Goggles (4368, -2.57 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Double-Stitched Woolen Shoulders (4314, -0.29 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -0.67 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Caretaker's Cape (20428, -0.13 DPS) [rep]; Feyscale Cloak (6632, -0.26 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.13 DPS) [crafted]; Bloody Apron (6226, -0.13 DPS) [dungeon]; Green Woolen Vest (2582, -1.36 DPS, sim-verified) [crafted] |
| wrist | Windsong Bangles (263336) | Breaking the Breaker [quest] | 1.0 | yes | Mindthrust Bracers (1974, -0.00 DPS, sim-verified) [dungeon]; Silver-lined Bracers (3224, -0.13 DPS) [world]; Seer's Cuffs (3645, -0.13 DPS) [dungeon] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.29 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.40 DPS) [crafted]; Heavy Woolen Gloves (4310, -0.67 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (35.1 DPS) | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.54 DPS) [world]; Novice Arcanist's Sash (253885, -0.83 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Colorful Kilt (10048, -0.54 DPS) [crafted]; Silk-threaded Trousers (1929, -0.92 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.40 DPS) [crafted]; Pristine Boots (253889, -0.54 DPS) [crafted]; Feather Padded Treads (285345, -0.63 DPS, sim-verified) [world] |
| finger1 | Minor Channeling Ring (1449) (or Lorekeeper's Ring (20431)) | WANTED: Chok'sul [quest] | 5.0 | yes | Sludge-Stained Band (286535, -0.27 DPS) [world]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Ring of the Shadow (1462, -0.67 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.40 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Ring of the Shadow (1462, -0.67 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 158.1 | yes | Staff of Westfall (2042, +0.00 DPS, sim-verified) [quest]; Twisted Chanter's Staff (890, -0.77 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -0.85 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 169.3 | yes | Firebelcher (5243, -1.48 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest]; Sizzle Stick (8071, -4.39 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Windsong Bangles; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Living Root; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 204, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak; 9792 Ivycloth Boots

### Band 30 (gnome, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 53.9. Weights run: 1.4s. Verify run: 1.2s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.288), intellect=not significant (0.400 ± 0.325), crit=0.845 ± 0.045, hit=2.738 ± 0.165, spell_haste=1.692 ± 0.302, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.288), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | Razorfen Downs: Withered Battle Boar [dungeon] | 11.0 | yes | Silk Headband (7050, -0.18 DPS) [crafted]; Embalmed Shroud (7691, -0.27 DPS) [dungeon]; Enchanter's Cowl (4322, -0.98 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 | yes | Crystal Starfire Medallion (5003, -0.71 DPS) [dungeon]; Pendant of Myzrael (4614, -0.86 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.63 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 | yes | Fairywing Mantle (9536, -0.27 DPS) [quest]; Invoker's Mantle (215365, -0.33 DPS) [crafted]; Death Speaker Mantle (6685, -0.69 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Prelacy Cape (7004, -0.09 DPS) [quest]; Repairman's Cape (9605, -0.54 DPS, sim-verified) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.2 | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.26 DPS) [dungeon]; Pristine Gown (253961, -0.40 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.60 DPS) [quest]; Mindthrust Bracers (1974, -0.64 DPS) [dungeon]; Nightsky Wristbands (6407, -2.02 DPS, sim-verified) [dungeon] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 8.4 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Shilly Mitts (9609, -0.13 DPS) [quest]; Truefaith Gloves (7049, -0.20 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 12.2 | yes | Invoker's Cord (215366, -0.29 DPS) [crafted]; Crimson Silk Belt (7055, -0.31 DPS) [crafted]; Belt of Arugal (6392, -0.44 DPS, sim-verified) [dungeon] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.35 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.8 | yes | Acidic Walkers (9454, -0.15 DPS) [dungeon]; Nimbus Boots (6998, -0.35 DPS) [quest]; Spidersilk Boots (4320, -1.62 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.11 DPS) [quest]; Lorekeeper's Ring (20431, -0.18 DPS) [rep]; Electrocutioner Lagnut (9447, -0.37 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.27 DPS) [dungeon]; Sludge-Stained Band (286535, -0.27 DPS) [world]; Minor Channeling Ring (1449, -1.47 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (53.7 DPS) | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Gnarled Ash Staff (791) | Gnomeregan: Leprous Technician [dungeon] | 279.8 | yes | Glimmering Staff (249392, +0.00 DPS, sim-verified) [crafted]; Lorekeeper's Staff (212580, -0.74 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.74 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 365.7 | yes | Greater Mystic Wand (217287, -1.96 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest]; Scorching Wand (5213, -4.63 DPS) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Gnarled Ash Staff; ranged: Necrotic Wand

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 40 (gnome, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 105.8. Weights run: 1.2s. Verify run: 1.1s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.308), intellect=not significant (0.667 ± 0.396), crit=0.836 ± 0.047, hit=3.138 ± 0.219, spell_haste=1.901 ± 0.341, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.308), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -0.91 DPS) [crafted]; Holy Shroud (2721, -1.09 DPS) [dungeon]; Augural Shroud (2620, -1.76 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Darkspear Warding Pendant (272074, -0.69 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.84 DPS) [vendor]; Necklace of Calisea (1714, -1.61 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.7 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.22 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.22 DPS) [vendor]; Icy Cloak (4327, -0.25 DPS) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 26.0 | yes | Robe of Power (7054, -0.44 DPS) [crafted]; Crimson Silk Vest (7058, -0.80 DPS) [crafted]; Dreamweave Vest (10021, -0.82 DPS, sim-verified) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.22 DPS) [quest]; Aurora Bracers (4043, -0.40 DPS) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 | yes | Black Mageweave Gloves (10003, -0.62 DPS) [crafted]; Red Mageweave Gloves (10018, -0.86 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -0.87 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (104.4 DPS) | yes | Gilded Cord (254037, -0.36 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.40 DPS) [rep]; Deathmage Sash (10771, -1.92 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Abomination Skin Leggings (23173, -0.84 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.05 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.32 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -1.49 DPS) [dungeon]; Spidersilk Boots (4320, -1.56 DPS) [crafted]; Gilded Slippers (254001, -2.11 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 | yes | Ring of Forlorn Spirits (2043, -0.65 DPS) [quest]; Reedknot Ring (9622, -0.76 DPS) [quest]; Minor Channeling Ring (1449, -0.84 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Ring of Forlorn Spirits (2043, -0.06 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.22 DPS) [quest]; Lorekeeper's Ring (19525, -0.22 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | Gnomeregan: Leprous Assistant [dungeon] | 378.4 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Black Duskwood Staff (937, -7.64 DPS) [dungeon]; Spiritchaser Staff (1613, -7.64 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (103.6 DPS) | yes | Ember Wand (5215, -0.98 DPS) [dungeon]; Necrotic Wand (7708, -1.12 DPS) [dungeon]; Umbral Wand (5216, -1.15 DPS, sim-verified) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs

### Band 50 (gnome, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 134.7. Weights run: 1.2s. Verify run: 1.1s. 721 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.208, intellect=not significant (-0.672 ± 0.236), crit=0.535 ± 0.036, hit=1.703 ± 0.134, spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Blood Guard's Dreadweave Hat (220907, -1.55 DPS) [vendor]; Dreamweave Circlet (10041, -1.68 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.47 DPS, sim-verified) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.32 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.96 DPS) [dungeon]; Pendant of Myzrael (4614, -1.96 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 15.5 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -0.70 DPS) [dungeon]; Black Mageweave Shoulders (10027, -1.54 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Runecloth Cloak (13860, -1.40 DPS) [crafted]; Nightfall Drape (12465, -1.43 DPS, sim-verified) [dungeon]; Icy Cloak (4327, -1.96 DPS) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 22.0 | yes | Stone Guard's Dreadweave Vest (220904, -0.70 DPS) [vendor]; Acumen Robes (17775, -0.84 DPS) [quest]; Knight's Dreadweave Vest (220886, -2.68 DPS, sim-verified) [vendor] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.56 DPS) [quest]; Nethergeld Cuffs (254061, -0.56 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Black Mageweave Gloves (10003, -1.09 DPS, sim-verified) [crafted]; Brightcloth Gloves (14101, -1.40 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.40 DPS) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 16.5 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -0.70 DPS) [rep]; Ghostweave Cord (254073, -0.70 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 19.5 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -0.14 DPS) [crafted]; Red Mageweave Pants (10009, -1.54 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 25.0 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -0.29 DPS) [crafted]; Black Mageweave Boots (10026, -3.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 17.0 | yes | Philanthropist's Ring (281635, -1.97 DPS) [quest]; Ring of Forlorn Spirits (2043, -2.53 DPS) [quest]; Reedknot Ring (9622, -2.81 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Philanthropist's Ring (281635, -0.71 DPS, sim-verified) [quest]; Lorekeeper's Ring (19524, -0.84 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.12 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (132.3 DPS) | yes | Frozen Heart of the Mountain (249469, -0.65 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -1.68 DPS) [quest]; Guardian Talisman (1490, -1.68 DPS) [quest] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (132.3 DPS) | yes | Frozen Heart of the Mountain (249469, -2.41 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -3.36 DPS) [quest]; Guardian Talisman (1490, -3.36 DPS) [quest] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | sim-verified (132.3 DPS) | yes | Thrash Blade (17705, -1.49 DPS, sim-verified) [quest]; Kindling Stave (11750, -1.66 DPS) [dungeon]; Soulkeeper (1607, -3.04 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (134.7 DPS) | yes | Pyric Caduceus (11748, -1.63 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -2.98 DPS) [world_drop]; Lesser Eternal Wand (249232, -3.41 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Abyss Shard; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 721, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 269.7. Weights run: 1.2s. Verify run: 1.1s. 1113 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.620, intellect=4.547 ± -0.743, crit=-1.774 ± -0.098, hit=-5.572 ± -0.332, spell_haste=-2.242 ± -0.730, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.083 ± -0.620, fire_power=-0.082 ± -0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | sim-verified (265.4 DPS) | yes | Field Marshal's Coronal (17578, +0.00 DPS) [vendor]; Doomcaller's Circlet (21337, +0.00 DPS) [quest]; Magister's Crown (16686, -9.67 DPS, sim-verified) [dungeon] |
| neck | Charm of the Shifting Sands (21504) | The Fall of Ossirian [quest] | sim-verified (253.6 DPS) | yes | Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Amulet of the Dawn (22657, +0.00 DPS) [quest]; Blazefury Medallion (17111, -6.41 DPS, sim-verified) [world] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | 108.0 | yes | Elder Wizard's Mantle (13013, +0.00 DPS, sim-verified) [world_drop]; Field Marshal's Dreadweave Shoulders (17580, +0.00 DPS) [vendor]; Warlord's Dreadweave Mantle (17590, +0.00 DPS) [vendor] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | 72.8 | yes | Hide of the Wild (18510, +0.00 DPS) [crafted]; Shroud of Unspoken Names (21418, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272076, -0.09 DPS, sim-verified) [vendor] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 151.0 | yes | Field Marshal's Dreadweave Robe (17581, +0.00 DPS) [vendor]; Warlord's Dreadweave Robe (17592, +0.00 DPS) [vendor]; Magister's Robes (16688, -17.62 DPS, sim-verified) [dungeon] |
| wrist | Plagueheart Bindings (22511) | Plagueheart Bindings [quest] | 86.7 | yes | Marshal's Dreadweave Cuffs (17582, +0.00 DPS) [pvp]; Deathmist Bracers (22071, +0.00 DPS) [quest]; Magiskull Cuffs (13107, -10.07 DPS, sim-verified) [world_drop] |
| hands | Plagueheart Gloves (22509) | Plagueheart Gloves [quest] | sim-verified (258.8 DPS) | yes | Mooncloth Gloves (18409, +0.00 DPS) [crafted]; Virtuous Gloves (22081, +0.00 DPS) [quest]; Raider Handwraps (272098, -3.09 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 125.0 | yes | Devout Belt (16696, +0.00 DPS) [dungeon]; Marshal's Dreadweave Sash (17585, +0.00 DPS) [pvp]; Magister's Belt (16685, -6.34 DPS, sim-verified) [dungeon] |
| legs | Plagueheart Leggings (22505) | Plagueheart Leggings [quest] | 150.7 | yes | Marshal's Dreadweave Leggings (17579, +0.00 DPS) [vendor]; General's Dreadweave Pants (17593, +0.00 DPS) [vendor]; Doomcaller's Trousers (21336, -5.77 DPS, sim-verified) [quest] |
| feet | Plagueheart Sandals (22508) | Plagueheart Sandals [quest] | 104.8 | yes | Knight-Lieutenant's Dreadweave Boots (17562, +0.00 DPS) [pvp]; Bloodvine Boots (19684, +0.00 DPS) [crafted]; Doomcaller's Footwraps (21338, -5.86 DPS, sim-verified) [quest] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 72.0 | yes | Ring of Entropy (18543, +0.00 DPS) [world]; Cauterizing Band (19140, +0.00 DPS) [world_drop]; Ring of the Fallen God (21709, +0.00 DPS) [quest] |
| finger2 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | 70.9 | yes | Cauterizing Band (19140, +0.00 DPS, sim-verified) [world_drop]; Signet Ring of the Bronze Dragonflight (21210, +0.00 DPS) [quest]; Signet Ring of the Bronze Dragonflight (234032, +0.00 DPS) [vendor] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (253.6 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, -2.31 DPS, sim-verified) [dungeon] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (253.6 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon]; Thunderbrew's Boot Flask (744, -1.32 DPS, sim-verified) [quest] |
| main_hand | Crackling Staff (19102) | Stormpike Guard [rep] | sim-verified (253.6 DPS) | yes | Shortsword of Vengeance (754, -7.34 DPS, sim-verified) [dungeon]; High Warlord's War Staff (234549, -8.88 DPS) [vendor]; Grand Marshal's Stave (234571, -8.88 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 40.8 | yes | Cairnstone Sliver (9654, +0.00 DPS) [quest]; Stormrager (16997, +0.00 DPS) [quest]; Charged Lightning Rod (11860, -3.67 DPS, sim-verified) [quest] |

**New at 60:** head: Plagueheart Circlet; neck: Charm of the Shifting Sands; shoulder: Magister's Mantle; back: Darkspear Raider's Cloak; chest: Plagueheart Robe; wrist: Plagueheart Bindings; hands: Plagueheart Gloves; waist: Knowledge of the Timbermaw; legs: Plagueheart Leggings; feet: Plagueheart Sandals; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Crackling Staff; ranged: Brilliant Wand

No-known-source sample (15 of 1113, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 00000000000000000-2351000000000000000-0000000000000000)

Set DPS (verified): 34.3. Weights run: 1.5s. Verify run: 1.1s. 203 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.103, intellect=not significant (-0.337 ± 0.116), crit=0.435 ± 0.021, hit=1.467 ± 0.069, spell_haste=0.633 ± 0.114, spell_penetration=not significant (0.000 ± 0.000), shadow_power=1.000 ± 0.103, fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Shadow Goggles (4373, -0.81 DPS) [crafted]; Lucky Fishing Hat (19972, -0.81 DPS) [quest]; Flying Tiger Goggles (4368, -2.49 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 5.0 | yes | Slime-encrusted Pads (6461, -0.67 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -1.03 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Battle Healer's Cloak (20427, -0.13 DPS) [rep]; Feyscale Cloak (6632, -0.41 DPS, sim-verified) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 5.0 | yes | Gray Woolen Robe (2585, -0.13 DPS) [crafted]; Bloody Apron (6226, -0.13 DPS) [dungeon]; Green Woolen Vest (2582, -2.04 DPS, sim-verified) [crafted] |
| wrist | Owlbeard Bracers (16981) (or Windsong Bangles (263336)) | Earthen Arise [quest] | 1.0 | yes | Mindthrust Bracers (1974, -0.13 DPS) [dungeon]; Silver-lined Bracers (3224, -0.13 DPS) [world]; Windsong Bangles (263336, -0.81 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.39 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.40 DPS) [quest]; Pristine Gloves (253913, -0.40 DPS) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (33.7 DPS) | yes | Novice Ardent's Sash (253887, -0.27 DPS) [crafted]; Lesser Belt of the Spire (1299, -0.54 DPS) [world]; Novice Arcanist's Sash (253885, -0.74 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 9.0 | yes | Filigreed Pristine Leggings (253937, -0.40 DPS) [crafted]; Colorful Kilt (10048, -0.54 DPS) [crafted]; Silk-threaded Trousers (1929, -2.06 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 7.0 | yes | Red Woolen Boots (4313, -0.40 DPS) [crafted]; Pristine Boots (253889, -0.54 DPS) [crafted]; Feather Padded Treads (285345, -1.40 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.67 DPS) [dungeon]; Ring of the Shadow (1462, -0.67 DPS) [world]; Ring of Scorn (3235, -0.67 DPS) [quest] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, -0.23 DPS, sim-verified) [dungeon]; Ring of the Shadow (1462, -0.40 DPS) [world]; Ring of Scorn (3235, -0.40 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | Blackfathom Deeps: Blackfathom Myrmidon [dungeon] | sim-verified (33.6 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Crescent Staff (6505, -0.10 DPS) [quest]; Living Root (6631, -0.68 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 169.3 | yes | Firebelcher (5243, -1.51 DPS, sim-verified) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest]; Sizzle Stick (8071, -4.39 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Owlbeard Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (troll, 00000000000000000-2352113101200000000-0000000000000000)

Set DPS (verified): 53.1. Weights run: 1.4s. Verify run: 1.2s. 411 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.288), intellect=not significant (0.400 ± 0.325), crit=0.845 ± 0.045, hit=2.738 ± 0.165, spell_haste=1.692 ± 0.302, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.288), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | Razorfen Downs: Withered Battle Boar [dungeon] | 11.0 | yes | Silk Headband (7050, -0.18 DPS) [crafted]; Embalmed Shroud (7691, -0.27 DPS) [dungeon]; Enchanter's Cowl (4322, -0.34 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.4 | yes | Crystal Starfire Medallion (5003, -0.71 DPS) [dungeon]; Pendant of Myzrael (4614, -0.86 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.34 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.6 | yes | Fairywing Mantle (9536, -0.27 DPS) [quest]; Invoker's Mantle (215365, -0.33 DPS) [crafted]; Death Speaker Mantle (6685, -0.71 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, -0.01 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.09 DPS) [crafted]; Battle Healer's Cloak (19529, -0.09 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.2 | yes | Tree Bark Jacket (1486, +0.00 DPS, sim-verified) [dungeon]; Death Speaker Robes (6682, -0.26 DPS) [dungeon]; Pristine Gown (253961, -0.40 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.60 DPS) [quest]; Mindthrust Bracers (1974, -0.64 DPS) [dungeon]; Nightsky Wristbands (6407, -1.44 DPS, sim-verified) [dungeon] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 8.0 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.16 DPS) [crafted]; Gnoll Casting Gloves (892, -0.18 DPS) [world] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 12.2 | yes | Belt of Arugal (6392, -0.18 DPS) [dungeon]; Warsong Sash (16975, -0.19 DPS, sim-verified) [quest]; Invoker's Cord (215366, -0.29 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 12.2 | yes | Gaze Dreamer Pants (6903, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.22 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.35 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.8 | yes | Acidic Walkers (9454, -0.15 DPS) [dungeon]; Boots of the Enchanter (4325, -0.44 DPS) [crafted]; Spidersilk Boots (4320, -1.64 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.18 DPS) [rep]; Electrocutioner Lagnut (9447, -0.37 DPS) [dungeon]; Sludge-Stained Band (286535, -0.37 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.27 DPS) [world]; Black Widow Band (6199, -0.29 DPS) [world]; Electrocutioner Lagnut (9447, -1.66 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | sim-verified (53.1 DPS) | yes | Lorekeeper's Staff (212580, -0.07 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.07 DPS) [vendor]; Gnarled Ash Staff (791, -0.74 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 365.7 | yes | Greater Mystic Wand (217287, -2.01 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -4.48 DPS) [quest]; Scorching Wand (5213, -4.63 DPS) [dungeon] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 411, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (troll, 00000000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 103.2. Weights run: 1.2s. Verify run: 1.1s. 570 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.308), intellect=not significant (0.667 ± 0.396), crit=0.836 ± 0.047, hit=3.138 ± 0.219, spell_haste=1.901 ± 0.341, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (1.000 ± 0.308), fire_power=not significant (0.000 ± 0.000)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Enchanter's Cowl (4322, -0.91 DPS) [crafted]; Holy Shroud (2721, -1.09 DPS) [dungeon]; Augural Shroud (2620, -1.94 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Darkspear Warding Pendant (272074, -0.69 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.84 DPS) [vendor]; Necklace of Calisea (1714, -1.63 DPS, sim-verified) [dungeon] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.7 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.07 DPS) [dungeon]; Berylline Pads (4197, -0.22 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.22 DPS) [vendor]; Icy Cloak (4327, -0.25 DPS) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 26.0 | yes | Robe of Power (7054, -0.44 DPS) [crafted]; Crimson Silk Vest (7058, -0.80 DPS) [crafted]; Dreamweave Vest (10021, -0.84 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (101.6 DPS) | yes | Condor Bracers (15864, -0.22 DPS) [quest]; Aurora Bracers (4043, -0.40 DPS) [dungeon]; Radiant Silver Bracers (4545, -1.76 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.7 | yes | Black Mageweave Gloves (10003, -0.62 DPS) [crafted]; Gilded Handwraps (254021, -0.87 DPS) [crafted]; Red Mageweave Gloves (10018, -1.07 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (100.9 DPS) | yes | Gilded Cord (254037, -0.36 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.40 DPS) [rep]; Deathmage Sash (10771, -1.03 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Abomination Skin Leggings (23173, -0.84 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.05 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.13 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -1.49 DPS) [dungeon]; Spidersilk Boots (4320, -1.56 DPS) [crafted]; Gilded Slippers (254001, -1.79 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.0 | yes | Reedknot Ring (9622, -0.76 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.87 DPS) [vendor]; Ogremind Ring (1993, -1.02 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.22 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.33 DPS) [vendor]; Reedknot Ring (9622, -0.64 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | Gnomeregan: Leprous Assistant [dungeon] | 378.4 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Black Duskwood Staff (937, -7.64 DPS) [dungeon]; Spiritchaser Staff (1613, -7.64 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | sim-verified (101.1 DPS) | yes | Ember Wand (5215, -0.98 DPS) [dungeon]; Necrotic Wand (7708, -1.12 DPS) [dungeon]; Umbral Wand (5216, -1.25 DPS, sim-verified) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 570, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (troll, 25300000000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 131.0. Weights run: 1.2s. Verify run: 1.0s. 723 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.208, intellect=not significant (-0.672 ± 0.236), crit=0.535 ± 0.036, hit=1.703 ± 0.134, spell_haste=not significant (0.459 ± 0.228), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.947 ± 0.209, fire_power=0.054 ± 0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme Plus (15999) | Engineering [crafted] | 27.0 | yes | Blood Guard's Dreadweave Hat (220907, -1.55 DPS) [vendor]; Dreamweave Circlet (10041, -1.68 DPS) [crafted]; Knight-Lieutenant's Dreadweave Hat (220889, -2.85 DPS, sim-verified) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.67 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.96 DPS) [dungeon]; Choker of the High Shaman (4112, -1.96 DPS) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 15.5 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -0.70 DPS) [dungeon]; Black Mageweave Shoulders (10027, -1.54 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Deep Woodlands Cloak (19121, -1.09 DPS, sim-verified) [quest]; Nightfall Drape (12465, -1.40 DPS) [dungeon]; Runecloth Cloak (13860, -1.40 DPS) [crafted] |
| chest | Robe of the Magi (1716) | Gnomeregan: Leprous Assistant [dungeon] | 22.0 | yes | Stone Guard's Dreadweave Vest (220904, -0.70 DPS) [vendor]; Acumen Robes (17775, -0.84 DPS) [quest]; Knight's Dreadweave Vest (220886, -2.83 DPS, sim-verified) [vendor] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, -0.56 DPS) [crafted]; Condor Bracers (15864, -1.08 DPS, sim-verified) [quest]; Bloodband Bracers (11469, -1.12 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Black Mageweave Gloves (10003, -1.25 DPS, sim-verified) [crafted]; Brightcloth Gloves (14101, -1.40 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -1.40 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 16.5 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -0.70 DPS) [rep]; Ghostweave Cord (254073, -0.70 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 19.5 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -0.14 DPS) [crafted]; Red Mageweave Pants (10009, -1.54 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 25.0 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -0.29 DPS) [crafted]; Black Mageweave Boots (10026, -3.93 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 17.0 | yes | Philanthropist's Ring (281635, -1.97 DPS) [quest]; Reedknot Ring (9622, -2.81 DPS) [quest]; Sea Giant's Toe Ring (274746, -3.09 DPS) [vendor] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Advisor's Ring (19520, -0.84 DPS) [rep]; Reedknot Ring (9622, -1.40 DPS) [quest]; Philanthropist's Ring (281635, -1.40 DPS, sim-verified) [quest] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (130.6 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.55 DPS, sim-verified) [crafted]; Guardian Talisman (1490, -1.68 DPS) [quest] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (130.6 DPS) | yes | Rune of the Guard Captain (19120, -0.02 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.82 DPS, sim-verified) [crafted]; Guardian Talisman (1490, -3.36 DPS) [quest] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | sim-verified (130.6 DPS) | yes | Kindling Stave (11750, -1.66 DPS) [dungeon]; Thrash Blade (17705, -2.70 DPS, sim-verified) [quest]; Soulkeeper (1607, -3.04 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 187.3 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -4.51 DPS) [crafted] |

**New at 50:** head: Spellpower Goggles Xtreme Plus; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Uther's Strength; trinket2: Abyss Shard; main_hand: Soul Harvester; ranged: Pyric Caduceus

No-known-source sample (15 of 723, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (troll, 25532300000000000-2352113101200001351-0000000000000000)

Set DPS (verified): 258.4. Weights run: 1.2s. Verify run: 1.1s. 1116 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± -0.620, intellect=4.547 ± -0.743, crit=-1.774 ± -0.098, hit=-5.572 ± -0.332, spell_haste=-2.242 ± -0.730, spell_penetration=not significant (-0.000 ± -0.000), shadow_power=1.083 ± -0.620, fire_power=-0.082 ± -0.001

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | sim-verified (254.8 DPS) | yes | Field Marshal's Coronal (17578, +0.00 DPS) [vendor]; Doomcaller's Circlet (21337, +0.00 DPS) [quest]; Magister's Crown (16686, -8.56 DPS, sim-verified) [dungeon] |
| neck | Charm of the Shifting Sands (21504) | The Fall of Ossirian [quest] | sim-verified (246.7 DPS) | yes | Beads of Ogre Mojo (22149, +0.00 DPS) [quest]; Amulet of the Dawn (22657, +0.00 DPS) [quest]; Blazefury Medallion (17111, -6.76 DPS, sim-verified) [world] |
| shoulder | Magister's Mantle (16689) | Scholomance: Ras Frostwhisper [dungeon] | 108.0 | yes | Elder Wizard's Mantle (13013, +0.00 DPS, sim-verified) [world_drop]; Field Marshal's Dreadweave Shoulders (17580, +0.00 DPS) [vendor]; Warlord's Dreadweave Mantle (17590, +0.00 DPS) [vendor] |
| back | Darkspear Raider's Cloak (272063) | Creeg Bothunk [vendor] | 72.8 | yes | Hide of the Wild (18510, +0.00 DPS) [crafted]; Shroud of Unspoken Names (21418, +0.00 DPS) [quest]; Darkspear Raider's Cloak (272076, +0.00 DPS, sim-verified) [vendor] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 151.0 | yes | Field Marshal's Dreadweave Robe (17581, +0.00 DPS) [vendor]; Warlord's Dreadweave Robe (17592, +0.00 DPS) [vendor]; Magister's Robes (16688, -16.78 DPS, sim-verified) [dungeon] |
| wrist | Plagueheart Bindings (22511) | Plagueheart Bindings [quest] | 86.7 | yes | General's Dreadweave Bracers (17587, +0.00 DPS) [pvp]; Deathmist Bracers (22071, +0.00 DPS) [quest]; Magiskull Cuffs (13107, -10.74 DPS, sim-verified) [world_drop] |
| hands | Plagueheart Gloves (22509) | Plagueheart Gloves [quest] | sim-verified (250.1 DPS) | yes | Mooncloth Gloves (18409, +0.00 DPS) [crafted]; Virtuous Gloves (22081, +0.00 DPS) [quest]; Raider Handwraps (272098, -3.81 DPS, sim-verified) [vendor] |
| waist | Knowledge of the Timbermaw (228190) | Meilosh [vendor] | 125.0 | yes | Devout Belt (16696, +0.00 DPS) [dungeon]; General's Dreadweave Belt (17589, +0.00 DPS) [pvp]; Magister's Belt (16685, -7.07 DPS, sim-verified) [dungeon] |
| legs | Plagueheart Leggings (22505) | Plagueheart Leggings [quest] | 150.7 | yes | Marshal's Dreadweave Leggings (17579, +0.00 DPS) [vendor]; General's Dreadweave Pants (17593, +0.00 DPS) [vendor]; Doomcaller's Trousers (21336, -5.25 DPS, sim-verified) [quest] |
| feet | Plagueheart Sandals (22508) | Plagueheart Sandals [quest] | 104.8 | yes | Marshal's Dreadweave Boots (17583, +0.00 DPS) [vendor]; Bloodvine Boots (19684, +0.00 DPS) [crafted]; Doomcaller's Footwraps (21338, -5.99 DPS, sim-verified) [quest] |
| finger1 | Elemental Focus Band (20682) | Prince Skaldrenox [world] | 72.0 | yes | Ring of Entropy (18543, +0.00 DPS) [world]; Cauterizing Band (19140, +0.00 DPS) [world_drop]; Ring of the Fallen God (21709, +0.00 DPS) [quest] |
| finger2 | Signet Ring of the Bronze Dragonflight (234436) | Anachronos [vendor] | 70.9 | yes | Cauterizing Band (19140, +0.00 DPS, sim-verified) [world_drop]; Signet Ring of the Bronze Dragonflight (21210, +0.00 DPS) [quest]; Signet Ring of the Bronze Dragonflight (234032, +0.00 DPS) [vendor] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (246.7 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Blazing Emblem (2802, +0.00 DPS) [dungeon]; Ankh of Life (1713, -2.85 DPS, sim-verified) [dungeon] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (246.7 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [dungeon]; Onyxia Blood Talisman (18406, -1.37 DPS, sim-verified) [quest] |
| main_hand | Staff of Hale Magefire (13000) | World drop [world_drop] | sim-verified (246.7 DPS) | yes | Shortsword of Vengeance (754, -2.94 DPS, sim-verified) [dungeon]; High Warlord's War Staff (234549, -13.61 DPS) [vendor]; Grand Marshal's Stave (234571, -13.61 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Brilliant Wand (249385) | Enchanting [crafted] | 40.8 | yes | Charged Lightning Rod (11860, +0.00 DPS) [quest]; Stormrager (16997, +0.00 DPS) [quest]; Nature's Breath (19118, -8.28 DPS, sim-verified) [quest] |

**New at 60:** head: Plagueheart Circlet; neck: Charm of the Shifting Sands; shoulder: Magister's Mantle; back: Darkspear Raider's Cloak; chest: Plagueheart Robe; wrist: Plagueheart Bindings; hands: Plagueheart Gloves; waist: Knowledge of the Timbermaw; legs: Plagueheart Leggings; feet: Plagueheart Sandals; finger1: Elemental Focus Band; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Staff of Hale Magefire; ranged: Brilliant Wand

No-known-source sample (15 of 1116, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

