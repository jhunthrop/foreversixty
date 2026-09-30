# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 42.8. Weights run: 1.4s. Verify run: 1.0s. 204 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.092, intellect=not significant (0.400 ± 0.102), crit=0.478 ± 0.022, hit=1.278 ± 0.062, spell_haste=not significant (0.104 ± 0.089), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.749 ± 0.092, fire_power=0.252 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -1.06 DPS) [crafted]; Lucky Fishing Hat (19972, -1.06 DPS) [quest]; Shadow Goggles (4373, -3.15 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.6 | yes | Double-Stitched Woolen Shoulders (4314, -0.28 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.16 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.18 DPS) [dungeon]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.40 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.49 DPS) [crafted]; Green Woolen Vest (2582, -0.53 DPS) [crafted]; Gray Woolen Robe (2585, -1.33 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.14 DPS) [world_drop]; Windsong Bangles (263336, -0.25 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.31 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.32 DPS) [crafted]; Blight Gloves (279877, -0.74 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Keller's Girdle (2911, -0.42 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.89 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 0.0 | yes | Silk-threaded Trousers (1929, -0.25 DPS) [dungeon]; Colorful Kilt (10048, -0.60 DPS) [crafted]; Abomination Skin Leggings (23173, -0.77 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Feather Padded Treads (285345, -0.50 DPS, sim-verified) [world]; Pristine Boots (253889, -0.78 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.49 DPS) [world]; Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon]; Loop of Sacrifice (281673, -0.67 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Sludge-Stained Band (286535, -0.51 DPS, sim-verified) [world]; Loop of Sacrifice (281673, -0.53 DPS) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 120.2 | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.14 DPS) [quest]; Staff of Westfall (2042, -0.36 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 129.4 | yes | Firebelcher (5243, -2.42 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -4.31 DPS) [quest]; Deepblaze (279896, -4.31 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Living Root; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 204, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak; 9792 Ivycloth Boots

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 64.9. Weights run: 1.4s. Verify run: 1.1s. 409 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.099, intellect=not significant (0.303 ± 0.115), crit=0.576 ± 0.031, hit=0.934 ± 0.062, spell_haste=not significant (0.340 ± 0.110), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.869 ± 0.099, fire_power=0.131 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, -0.63 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.67 DPS) [crafted]; Embalmed Shroud (7691, -1.00 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 | yes | Darkspear Warding Pendant (272075, -1.70 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.54 DPS) [world_drop]; Pendant of Myzrael (4614, -2.95 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 | yes | Death Speaker Mantle (6685, -0.41 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -1.00 DPS) [quest]; Invoker's Mantle (215365, -1.07 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Repairman's Cape (9605, -0.23 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.33 DPS) [crafted]; Prelacy Cape (7004, -0.33 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, -0.03 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.89 DPS) [dungeon]; Pristine Gown (253961, -1.30 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -1.96 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -2.40 DPS) [quest]; Mindthrust Bracers (1974, -2.50 DPS) [dungeon] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.3 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Shilly Mitts (9609, -0.11 DPS) [quest]; Gnoll Casting Gloves (892, -0.45 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 | yes | Belt of Arugal (6392, -0.54 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -1.13 DPS) [crafted]; Crimson Silk Belt (7055, -1.27 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -0.05 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.96 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.40 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.1 | yes | Acidic Walkers (9454, -0.57 DPS) [dungeon]; Nimbus Boots (6998, -1.04 DPS) [quest]; Spidersilk Boots (4320, -1.84 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.47 DPS) [quest]; Lorekeeper's Ring (20431, -0.67 DPS) [rep]; Electrocutioner Lagnut (9447, -1.34 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -1.00 DPS) [dungeon]; Sludge-Stained Band (286535, -1.00 DPS) [world]; Minor Channeling Ring (1449, -1.60 DPS, sim-verified) [quest] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Darkspear Voodoo Seal (272059, +0.00 DPS) [vendor] |
| trinket2 | Relentless Raider's Seal (272062) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS, sim-verified) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Darkspear Voodoo Seal (272059, +0.00 DPS) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 76.8 | yes | Lorekeeper's Staff (212580, -0.40 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.40 DPS) [vendor]; Gnarled Ash Staff (791, -0.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 102.4 | yes | Greater Mystic Wand (217287, -3.10 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -5.21 DPS) [quest]; Scorching Wand (5213, -5.36 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Talisman of Arathor; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 409, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 108.3. Weights run: 1.4s. Verify run: 1.1s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.273), intellect=1.317 ± 0.307, crit=1.241 ± 0.061, hit=2.170 ± 0.167, spell_haste=not significant (0.432 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.422 ± 0.273), fire_power=0.574 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 0.0 | yes | Craftsman's Monocle (4393, -0.30 DPS) [crafted]; Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Augural Shroud (2620, -2.55 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.9 | yes | Darkspear Warding Pendant (272074, -1.39 DPS) [vendor]; Necklace of Calisea (1714, -1.69 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.03 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.1 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.80 DPS) [dungeon]; Death Speaker Mantle (6685, -0.89 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | 0.0 | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Cloak of Rot (4462, -0.50 DPS) [world]; Darkspear Raider's Cloak (272077, -1.44 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.9 | yes | Robe of Power (7054, -0.02 DPS) [crafted]; Dreamweave Vest (10021, -0.42 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.92 DPS) [crafted] |
| wrist | Aurora Bracers (4043) (or Mistscape Bracers (4045), Enchanted Stonecloth Bracers (4979)) | World drop [world_drop] | 10.5 | yes | Mistscape Bracers (4045, +0.00 DPS, sim-verified) [world_drop]; Enchanted Stonecloth Bracers (4979, +0.00 DPS) [quest]; Arcane Runed Bracers (4744, -0.38 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 0.0 | yes | Stormcloth Gloves (10011, -0.85 DPS) [crafted]; Town Clerk's Mittens (270029, -1.17 DPS) [quest]; Red Mageweave Gloves (10018, -1.73 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 0.0 | yes | Gilded Cord (254037, -0.18 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.85 DPS) [quest]; Deathmage Sash (10771, -1.31 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.8 | yes | Crimson Silk Pantaloons (7062, -1.49 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -2.07 DPS) [dungeon]; Gilded Slippers (254001, -2.54 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.9 | yes | Voodoo Band (1996, -2.12 DPS) [world]; Mindbender Loop (5009, -2.12 DPS) [world_drop]; Black Widow Band (6199, -2.12 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 9.2 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | 0.0 | yes | Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS, sim-verified) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 179.7 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Staff of Dar'Orahil (15106, -6.21 DPS) [quest]; Windweaver Staff (7757, -6.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 0.0 | yes | Burning Sliver (5249, -1.31 DPS) [quest]; Goblin Igniter (5253, -1.43 DPS) [quest]; Umbral Wand (5216, -1.46 DPS, sim-verified) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Aurora Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots; 7473 Regal Mantle; 7475 Regal Cuffs

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 144.0. Weights run: 1.3s. Verify run: 1.1s. 721 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.605), intellect=not significant (-0.468 ± 0.718), crit=2.617 ± 0.140, hit=4.742 ± 0.386, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 50.6 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -2.17 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -3.66 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.26 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.08 DPS) [world_drop]; Pendant of Myzrael (4614, -1.08 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 44.6 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -4.90 DPS) [dungeon]; Black Mageweave Shoulders (10027, -5.37 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Runecloth Cloak (13860, -0.77 DPS) [crafted]; Nightfall Drape (12465, -0.97 DPS, sim-verified) [dungeon]; Icy Cloak (4327, -1.08 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 48.6 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -4.13 DPS) [world_drop]; Acumen Robes (17775, -4.59 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.31 DPS) [quest]; Nethergeld Cuffs (254061, -0.31 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.77 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.77 DPS) [vendor]; Black Mageweave Gloves (10003, -1.01 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 45.6 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -4.90 DPS) [rep]; Ghostweave Cord (254073, -4.90 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 48.6 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -4.59 DPS) [crafted]; Red Mageweave Pants (10009, -5.37 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 55.4 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.87 DPS) [crafted]; Black Mageweave Boots (10026, -6.88 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.4 | yes | Philanthropist's Ring (281635, -5.80 DPS) [quest]; Ring of Forlorn Spirits (2043, -6.11 DPS) [quest]; Reedknot Ring (9622, -6.26 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19523) | Silverwing Sentinels [rep] | 12.0 | yes | Philanthropist's Ring (281635, -0.21 DPS, sim-verified) [quest]; Lorekeeper's Ring (19524, -0.46 DPS) [rep]; Ring of Forlorn Spirits (2043, -0.62 DPS) [quest] |
| trinket1 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Frozen Heart of the Mountain (249469, -0.64 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -0.93 DPS) [quest]; Guardian Talisman (1490, -0.93 DPS) [quest] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, -1.86 DPS) [quest]; Guardian Talisman (1490, -1.86 DPS) [quest]; Frozen Heart of the Mountain (249469, -2.18 DPS, sim-verified) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 0.0 | yes | Kindling Stave (11750, -0.66 DPS) [dungeon]; Soulkeeper (1607, -5.62 DPS) [world_drop]; Spire of Hakkar (10844, -5.73 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | 0.0 | yes | Wand of Allistarj (13065, -2.35 DPS) [world]; Pyric Caduceus (11748, -3.76 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.79 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Arcane Runed Bracers; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Lorekeeper's Ring; trinket1: Uther's Strength; trinket2: Abyss Shard; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 721, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 270.5. Weights run: 1.3s. Verify run: 1.1s. 1147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.964), intellect=not significant (1.610 ± 0.952), crit=5.970 ± 0.274, hit=10.408 ± 0.660, spell_haste=not significant (1.990 ± 0.697), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.383 ± 0.964), fire_power=1.387 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | 344.5 | yes | Bloodvine Goggles (19999, -10.50 DPS, sim-verified) [crafted]; Doomcaller's Circlet (21337, -10.75 DPS) [quest]; Deathmist Mask (226909, -23.19 DPS) [quest] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 0.0 | yes | Beads of Ogre Might (22150, -10.55 DPS) [quest]; Medallion of the Dawn (22659, -13.13 DPS) [quest]; Charm of the Shifting Sands (21504, -18.08 DPS) [quest] |
| shoulder | Plagueheart Shoulderpads (22507) | Plagueheart Shoulderpads [quest] | 159.4 | yes | Doomcaller's Mantle (21335, -2.93 DPS, sim-verified) [quest]; Mantle of the Timbermaw (19050, -4.78 DPS) [crafted]; Champion's Dreadweave Spaulders (23256, -5.41 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 104.1 | yes | Chromatic Cloak (18509, -1.69 DPS, sim-verified) [crafted]; Shroud of Unspoken Names (21418, -9.03 DPS) [quest]; Hide of the Wild (18510, -9.33 DPS) [crafted] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 274.1 | yes | Bloodvine Vest (19682, -6.11 DPS, sim-verified) [crafted]; Doomcaller's Robes (21334, -15.41 DPS) [quest]; Earthpower Vest (21183, -16.41 DPS) [quest] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 131.1 | yes | Plagueheart Bindings (22511, -1.91 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -12.14 DPS) [rep]; Marshal's Dreadweave Cuffs (17582, -12.68 DPS) [pvp] |
| hands | Deathmist Wraps (22077) | Just Compensation [quest] | 0.0 | yes | Deathmist Wraps (226911, +0.00 DPS) [quest]; Plagueheart Gloves (22509, -0.13 DPS) [quest]; Gloves of Spell Mastery (14146, -2.81 DPS, sim-verified) [crafted] |
| waist | Plagueheart Belt (22510) | Plagueheart Belt [quest] | 136.9 | yes | Belt of the Archmage (18405, -2.90 DPS, sim-verified) [crafted]; Highlander's Cloth Girdle (20047, -3.74 DPS) [rep]; Highlander's Cloth Girdle (20097, -4.58 DPS) [rep] |
| legs | Plagueheart Leggings (22505) | Plagueheart Leggings [quest] | 160.8 | yes | Doomcaller's Trousers (21336, -1.25 DPS, sim-verified) [quest]; Bloodvine Leggings (19683, -1.27 DPS) [crafted]; Deathmist Leggings (226910, -1.61 DPS) [quest] |
| feet | Plagueheart Sandals (22508) | Plagueheart Sandals [quest] | 0.0 | yes | Sergeant Major's Dreadweave Boots (220891, -1.86 DPS) [vendor]; First Sergeant's Dreadweave Boots (220909, -1.86 DPS) [vendor]; Bloodvine Boots (19684, -4.55 DPS, sim-verified) [crafted] |
| finger1 | Ring of Unspoken Names (21417) | Ring of Unspoken Names [quest] | 201.7 | yes | Band of Earthen Might (21182, -1.77 DPS) [quest]; Ring of the Fallen God (21709, -6.42 DPS) [quest]; Mindtear Band (20632, -10.90 DPS) [world] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Stormpike Guard [rep] | 187.7 | yes | Band of Earthen Might (21182, +0.00 DPS, sim-verified) [quest]; Ring of the Fallen God (21709, -4.66 DPS) [quest]; Mindtear Band (20632, -9.14 DPS) [world] |
| trinket1 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Frozen Heart of the Mountain (249469, -0.22 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -0.76 DPS) [quest]; Guardian Talisman (1490, -0.76 DPS) [quest] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, -1.51 DPS) [quest]; Guardian Talisman (1490, -1.51 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.72 DPS, sim-verified) [crafted] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | 0.0 | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; High Warlord's Spellblade (234550, -5.33 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Stormrager (16997) | Order Must Be Restored [quest] | 0.0 | yes | Brilliant Wand (249385, -1.57 DPS) [crafted]; Wand of Biting Cold (19108, -5.11 DPS, sim-verified) [quest]; Torch of Austen (13004, -6.92 DPS) [world] |

**New at 60:** head: Plagueheart Circlet; neck: Onyxia Tooth Pendant; shoulder: Plagueheart Shoulderpads; back: Earthweave Cloak; chest: Plagueheart Robe; wrist: Rockfury Bracers; hands: Deathmist Wraps; waist: Plagueheart Belt; legs: Plagueheart Leggings; feet: Plagueheart Sandals; finger1: Ring of Unspoken Names; finger2: Don Julio's Band; main_hand: Ironbark Staff; ranged: Stormrager

No-known-source sample (15 of 1147, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 41.9. Weights run: 1.4s. Verify run: 1.1s. 203 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.092, intellect=not significant (0.400 ± 0.102), crit=0.478 ± 0.022, hit=1.278 ± 0.062, spell_haste=not significant (0.104 ± 0.089), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.749 ± 0.092, fire_power=0.252 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -1.06 DPS) [crafted]; Lucky Fishing Hat (19972, -1.06 DPS) [quest]; Shadow Goggles (4373, -2.55 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 6.6 | yes | Double-Stitched Woolen Shoulders (4314, -0.66 DPS, sim-verified) [crafted]; Slime-encrusted Pads (6461, -1.16 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Pearl-clasped Cloak (5542, -0.10 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.18 DPS) [dungeon]; Black Whelp Cloak (7283, -0.18 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.49 DPS) [crafted]; Green Woolen Vest (2582, -0.53 DPS) [crafted]; Gray Woolen Robe (2585, -1.23 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Owlbeard Bracers (16981, -0.11 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.22 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.32 DPS) [crafted]; Apothecary Gloves (10919, -0.53 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Keller's Girdle (2911, -0.42 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.84 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | 0.0 | yes | Silk-threaded Trousers (1929, -0.25 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.55 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.60 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Pristine Boots (253889, -0.78 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS) [crafted]; Feather Padded Treads (285345, -0.84 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Loop of Sacrifice (281673, -0.53 DPS) [quest]; Black Pearl Ring (6332, -0.74 DPS) [world] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Lavishly Jeweled Ring (1156, -0.09 DPS, sim-verified) [dungeon]; Loop of Sacrifice (281673, -0.18 DPS) [quest]; Black Pearl Ring (6332, -0.39 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | 0.0 | yes | Gnarled Necromancer's Staff (251534, -0.08 DPS) [quest]; Living Root (6631, -0.79 DPS, sim-verified) [dungeon]; Crescent Staff (6505, -0.81 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 129.4 | yes | Firebelcher (5243, -2.08 DPS, sim-verified) [dungeon]; Sizzle Stick (8071, -4.31 DPS) [quest]; Deepblaze (279896, -4.31 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves; 9779 Bandit Cloak; 9786 Raider's Cloak

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 63.1. Weights run: 1.4s. Verify run: 1.1s. 410 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.099, intellect=not significant (0.303 ± 0.115), crit=0.576 ± 0.031, hit=0.934 ± 0.062, spell_haste=not significant (0.340 ± 0.110), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.869 ± 0.099, fire_power=0.131 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.67 DPS) [crafted]; Embalmed Shroud (7691, -1.00 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 | yes | Darkspear Warding Pendant (272075, -1.42 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.54 DPS) [world_drop]; Pendant of Myzrael (4614, -2.95 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 | yes | Death Speaker Mantle (6685, -0.89 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -1.00 DPS) [quest]; Invoker's Mantle (215365, -1.07 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.33 DPS) [crafted]; Battle Healer's Cloak (19529, -0.33 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.89 DPS) [dungeon]; Pristine Gown (253961, -1.30 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -1.44 DPS, sim-verified) [world_drop]; Tabitha's Cuffs (251486, -2.40 DPS) [quest]; Owlbeard Bracers (16981, -2.47 DPS) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.5 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.51 DPS) [world]; Truefaith Gloves (7049, -0.54 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 | yes | Warsong Sash (16975, +0.00 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.67 DPS) [dungeon]; Invoker's Cord (215366, -1.13 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.96 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.40 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.1 | yes | Acidic Walkers (9454, -0.57 DPS) [dungeon]; Boots of the Enchanter (4325, -1.38 DPS) [crafted]; Spidersilk Boots (4320, -1.69 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.67 DPS) [rep]; Electrocutioner Lagnut (9447, -1.34 DPS) [dungeon]; Sludge-Stained Band (286535, -1.34 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -1.00 DPS) [world]; Black Widow Band (6199, -1.30 DPS) [world]; Electrocutioner Lagnut (9447, -1.69 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 76.8 | yes | Lorekeeper's Staff (212580, -0.40 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.40 DPS) [vendor]; Gnarled Ash Staff (791, -0.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 102.4 | yes | Greater Mystic Wand (217287, -2.38 DPS, sim-verified) [crafted]; Gravestone Scepter (7001, -5.21 DPS) [quest]; Scorching Wand (5213, -5.36 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 410, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 8184 Firestarter; 8186 Dire Wand; 9362 Brilliant Gold Ring; 9395 Gloves of Old; 9747 Simple Britches

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 108.0. Weights run: 1.4s. Verify run: 1.1s. 568 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.273), intellect=1.317 ± 0.307, crit=1.241 ± 0.061, hit=2.170 ± 0.167, spell_haste=not significant (0.432 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.422 ± 0.273), fire_power=0.574 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 0.0 | yes | Craftsman's Monocle (4393, -0.30 DPS) [crafted]; Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Augural Shroud (2620, -2.15 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.9 | yes | Darkspear Warding Pendant (272074, -1.39 DPS) [vendor]; Necklace of Calisea (1714, -1.43 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.03 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.1 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.80 DPS) [dungeon]; Death Speaker Mantle (6685, -0.89 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | 0.0 | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Cloak of Rot (4462, -0.50 DPS) [world]; Darkspear Raider's Cloak (272077, -1.22 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.9 | yes | Robe of Power (7054, -0.02 DPS) [crafted]; Dreamweave Vest (10021, -0.71 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.92 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 14.5 | yes | Mistscape Bracers (4045, -0.98 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.98 DPS) [quest]; Aurora Bracers (4043, -0.99 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 0.0 | yes | Stormcloth Gloves (10011, -0.85 DPS) [crafted]; Red Mageweave Gloves (10018, -1.47 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.48 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | 0.0 | yes | Gilded Cord (254037, -0.18 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.85 DPS) [quest]; Deathmage Sash (10771, -1.69 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.8 | yes | Crimson Silk Pantaloons (7062, -1.53 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -2.07 DPS) [dungeon]; Gilded Slippers (254001, -2.28 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.9 | yes | Voodoo Band (1996, -2.12 DPS) [world]; Mindbender Loop (5009, -2.12 DPS) [world_drop]; Black Widow Band (6199, -2.12 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 9.2 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 179.7 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Staff of Dar'Orahil (15106, -6.21 DPS) [quest]; Windweaver Staff (7757, -6.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Twisted Nether Wand (249144) | Enchanting [crafted] | 0.0 | yes | Umbral Wand (5216, -1.11 DPS, sim-verified) [world_drop]; Goblin Igniter (5253, -1.43 DPS) [quest]; Necrotic Wand (7708, -1.52 DPS) [dungeon] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ogremind Ring; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Staff of Jordan; ranged: Twisted Nether Wand

No-known-source sample (15 of 568, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat; 7471 Regal Gloves; 7472 Regal Boots

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 141.2. Weights run: 1.3s. Verify run: 1.1s. 721 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.605), intellect=not significant (-0.468 ± 0.718), crit=2.617 ± 0.140, hit=4.742 ± 0.386, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 50.6 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -2.17 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -3.66 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.62 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.08 DPS) [world_drop]; Choker of the High Shaman (4112, -1.08 DPS) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 44.6 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -4.90 DPS) [dungeon]; Black Mageweave Shoulders (10027, -5.37 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Deep Woodlands Cloak (19121, -0.12 DPS, sim-verified) [quest]; Nightfall Drape (12465, -0.77 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 48.6 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -4.13 DPS) [world_drop]; Acumen Robes (17775, -4.59 DPS) [quest] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, -0.31 DPS) [crafted]; Bloodband Bracers (11469, -0.62 DPS) [quest]; Condor Bracers (15864, -1.49 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.77 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.77 DPS) [vendor]; Black Mageweave Gloves (10003, -1.69 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 45.6 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -4.90 DPS) [rep]; Ghostweave Cord (254073, -4.90 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 48.6 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -4.59 DPS) [crafted]; Red Mageweave Pants (10009, -5.37 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 55.4 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.87 DPS) [crafted]; Black Mageweave Boots (10026, -6.88 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.4 | yes | Philanthropist's Ring (281635, -5.80 DPS) [quest]; Reedknot Ring (9622, -6.26 DPS) [quest]; Sea Giant's Toe Ring (274746, -6.42 DPS) [vendor] |
| finger2 | Advisor's Ring (19519) | Warsong Outriders [rep] | 12.0 | yes | Advisor's Ring (19520, -0.46 DPS) [rep]; Philanthropist's Ring (281635, -0.51 DPS, sim-verified) [quest]; Reedknot Ring (9622, -0.77 DPS) [quest] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Guardian Talisman (1490, -1.86 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.73 DPS, sim-verified) [crafted]; Guardian Talisman (1490, -0.93 DPS) [quest] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | 0.0 | yes | Kindling Stave (11750, -0.66 DPS) [dungeon]; Soulkeeper (1607, -5.62 DPS) [world_drop]; Spire of Hakkar (10844, -5.73 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | 0.0 | yes | Wand of Allistarj (13065, -2.35 DPS) [world]; Pyric Caduceus (11748, -2.60 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.79 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Advisor's Ring; trinket1: Abyss Shard; trinket2: Uther's Strength; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 721, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 263.4. Weights run: 1.3s. Verify run: 1.1s. 1147 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.964), intellect=not significant (1.610 ± 0.952), crit=5.970 ± 0.274, hit=10.408 ± 0.660, spell_haste=not significant (1.990 ± 0.697), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.383 ± 0.964), fire_power=1.387 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Plagueheart Circlet (22506) | Plagueheart Circlet [quest] | 344.5 | yes | Doomcaller's Circlet (21337, -10.75 DPS) [quest]; Bloodvine Goggles (19999, -14.10 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -23.19 DPS) [quest] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 0.0 | yes | Beads of Ogre Might (22150, -10.55 DPS) [quest]; Medallion of the Dawn (22659, -13.13 DPS) [quest]; Charm of the Shifting Sands (21504, -18.08 DPS) [quest] |
| shoulder | Plagueheart Shoulderpads (22507) | Plagueheart Shoulderpads [quest] | 159.4 | yes | Doomcaller's Mantle (21335, -4.17 DPS, sim-verified) [quest]; Mantle of the Timbermaw (19050, -4.78 DPS) [crafted]; Champion's Dreadweave Spaulders (23256, -5.41 DPS) [vendor] |
| back | Earthweave Cloak (21187) | Volunteer's Battlegear [quest] | 104.1 | yes | Chromatic Cloak (18509, -2.28 DPS, sim-verified) [crafted]; Shroud of Unspoken Names (21418, -9.03 DPS) [quest]; Hide of the Wild (18510, -9.33 DPS) [crafted] |
| chest | Plagueheart Robe (22504) | Plagueheart Robe [quest] | 274.1 | yes | Bloodvine Vest (19682, -9.33 DPS, sim-verified) [crafted]; Doomcaller's Robes (21334, -15.41 DPS) [quest]; Earthpower Vest (21183, -16.41 DPS) [quest] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 131.1 | yes | Plagueheart Bindings (22511, -3.11 DPS, sim-verified) [quest]; Dryad's Wrist Bindings (19595, -12.14 DPS) [rep]; Marshal's Dreadweave Cuffs (17582, -12.68 DPS) [pvp] |
| hands | Gloves of Spell Mastery (14146) | Tailoring [crafted] | 192.3 | yes | Deathmist Wraps (22077, -0.84 DPS, sim-verified) [quest]; Deathmist Wraps (226911, -6.85 DPS) [quest]; Plagueheart Gloves (22509, -6.98 DPS) [quest] |
| waist | Plagueheart Belt (22510) | Plagueheart Belt [quest] | 136.9 | yes | Defiler's Cloth Girdle (20163, -3.74 DPS) [rep]; Defiler's Cloth Girdle (20165, -4.58 DPS) [rep]; Belt of the Archmage (18405, -5.64 DPS, sim-verified) [crafted] |
| legs | Plagueheart Leggings (22505) | Plagueheart Leggings [quest] | 160.8 | yes | Bloodvine Leggings (19683, -1.27 DPS) [crafted]; Deathmist Leggings (226910, -1.61 DPS) [quest]; Doomcaller's Trousers (21336, -2.06 DPS, sim-verified) [quest] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 148.8 | yes | Plagueheart Sandals (22508, +0.00 DPS, sim-verified) [quest]; Sergeant Major's Dreadweave Boots (220891, -2.81 DPS) [vendor]; First Sergeant's Dreadweave Boots (220909, -2.81 DPS) [vendor] |
| finger1 | Ring of Unspoken Names (21417) | Ring of Unspoken Names [quest] | 201.7 | yes | Band of Earthen Might (21182, -1.77 DPS) [quest]; Ring of the Fallen God (21709, -6.42 DPS) [quest]; Mindtear Band (20632, -10.90 DPS) [world] |
| finger2 | Don Julio's Band (19325) (or Band of Earthen Might (21182)) | Frostwolf Clan [rep] | 187.7 | yes | Band of Earthen Might (21182, +0.00 DPS, sim-verified) [quest]; Ring of the Fallen God (21709, -4.66 DPS) [quest]; Mindtear Band (20632, -9.14 DPS) [world] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Guardian Talisman (1490, -1.51 DPS) [quest] |
| trinket2 | Uther's Strength (11302) | Azuregos [world] | 0.0 | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Guardian Talisman (1490, -0.76 DPS) [quest]; Frozen Heart of the Mountain (249469, -1.15 DPS, sim-verified) [crafted] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | 0.0 | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; High Warlord's Spellblade (234550, -5.33 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | Stormrager (16997) | The Scarlet Oracle, Demetria [quest] | 0.0 | yes | Brilliant Wand (249385, -1.57 DPS) [crafted]; Wand of Biting Cold (19108, -3.20 DPS, sim-verified) [quest]; Torch of Austen (13004, -6.92 DPS) [world] |

**New at 60:** head: Plagueheart Circlet; neck: Onyxia Tooth Pendant; shoulder: Plagueheart Shoulderpads; back: Earthweave Cloak; chest: Plagueheart Robe; wrist: Rockfury Bracers; hands: Gloves of Spell Mastery; waist: Plagueheart Belt; legs: Plagueheart Leggings; feet: Bloodvine Boots; finger1: Ring of Unspoken Names; finger2: Don Julio's Band; main_hand: Ironbark Staff; ranged: Stormrager

No-known-source sample (15 of 1147, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers

