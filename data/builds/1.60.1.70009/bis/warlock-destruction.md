# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 43.1. Weights run: 1.5s. Verify run: 1.3s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.092, intellect=not significant (0.400 ± 0.102), crit=0.478 ± 0.022, hit=1.278 ± 0.062, spell_haste=not significant (0.104 ± 0.089), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.749 ± 0.092, fire_power=0.252 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -1.06 DPS) [crafted]; Lucky Fishing Hat (19972, -1.06 DPS) [quest]; Shadow Goggles (4373, -3.18 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 | yes | Reinforced Woolen Shoulders (4315, +0.00 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.81 DPS) [crafted]; Slime-encrusted Pads (6461, -1.52 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.18 DPS) [dungeon]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.30 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.49 DPS) [crafted]; Green Woolen Vest (2582, -0.53 DPS) [crafted]; Gray Woolen Robe (2585, -1.34 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.14 DPS) [world_drop]; Windsong Bangles (263336, -0.25 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Pristine Gloves (253913, -0.32 DPS) [crafted]; Gnoll Casting Gloves (892, -0.33 DPS, sim-verified) [world]; Blight Gloves (279877, -0.74 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Keller's Girdle (2911, -0.42 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.91 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (43.1 DPS) | yes | Silk-threaded Trousers (1929, -0.25 DPS) [dungeon]; Colorful Kilt (10048, -0.60 DPS) [crafted]; Abomination Skin Leggings (23173, -1.03 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Feather Padded Treads (285345, -0.47 DPS, sim-verified) [world]; Pristine Boots (253889, -0.78 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.49 DPS) [world]; Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon]; Loop of Sacrifice (281673, -0.67 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Sludge-Stained Band (286535, -0.44 DPS, sim-verified) [world]; Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Loop of Sacrifice (281673, -0.53 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 120.2 | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.14 DPS) [quest]; Staff of Westfall (2042, -0.36 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 129.4 | yes | Skycaller (12984, -2.13 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.55 DPS) [dungeon]; Sizzle Stick (8071, -4.31 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Living Root; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 20434 Lorekeeper's Staff; 20440 Protector's Sword; 20443 Sentinel's Blade; 209615 Insignia of the Alliance

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 64.9. Weights run: 1.5s. Verify run: 1.2s. 209 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.099, intellect=not significant (0.303 ± 0.115), crit=0.576 ± 0.031, hit=0.934 ± 0.062, spell_haste=not significant (0.340 ± 0.110), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.869 ± 0.099, fire_power=0.131 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, -0.63 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.67 DPS) [crafted]; Embalmed Shroud (7691, -1.00 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 | yes | Darkspear Warding Pendant (272075, -1.70 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.54 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.54 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 | yes | Death Speaker Mantle (6685, -0.41 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -1.00 DPS) [quest]; Invoker's Mantle (215365, -1.07 DPS) [crafted] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Repairman's Cape (9605, -0.23 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.33 DPS) [crafted]; Prelacy Cape (7004, -0.33 DPS) [quest] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, -0.03 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.89 DPS) [dungeon]; Pristine Gown (253961, -1.30 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Glowing Magical Bracelets (13106, -1.72 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -2.40 DPS) [world_drop]; Tabitha's Cuffs (251486, -2.40 DPS) [quest] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.3 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Shilly Mitts (9609, -0.11 DPS) [quest]; Gnoll Casting Gloves (892, -0.45 DPS) [world] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 | yes | Belt of Arugal (6392, -0.54 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -1.13 DPS) [crafted]; Crimson Silk Belt (7055, -1.27 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, -0.05 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.96 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.40 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.1 | yes | Acidic Walkers (9454, -0.57 DPS) [dungeon]; Nimbus Boots (6998, -1.04 DPS) [quest]; Spidersilk Boots (4320, -1.84 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.47 DPS) [quest]; Lorekeeper's Ring (20431, -0.67 DPS) [rep]; Electrocutioner Lagnut (9447, -1.34 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -1.00 DPS) [dungeon]; Sludge-Stained Band (286535, -1.00 DPS) [world]; Minor Channeling Ring (1449, -1.60 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 76.8 | yes | Lorekeeper's Staff (212580, -0.40 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.40 DPS) [vendor]; Gnarled Ash Staff (791, -0.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 102.4 | yes | Starfaller (13063, -0.66 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.54 DPS) [crafted]; Thunderwood (13062, -5.10 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Talisman of Arathor; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 209, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 107.6. Weights run: 1.3s. Verify run: 1.2s. 286 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.273), intellect=1.317 ± 0.307, crit=1.241 ± 0.061, hit=2.170 ± 0.167, spell_haste=not significant (0.432 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.422 ± 0.273), fire_power=0.574 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (103.3 DPS) | yes | Craftsman's Monocle (4393, -0.30 DPS) [crafted]; Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Augural Shroud (2620, -2.54 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.9 | yes | Darkspear Warding Pendant (272074, -1.39 DPS) [vendor]; Necklace of Calisea (1714, -1.65 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.03 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.1 | yes | Green Silken Shoulders (7057, -0.11 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.80 DPS) [dungeon]; Death Speaker Mantle (6685, -0.89 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (102.4 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Cloak of Rot (4462, -0.50 DPS) [world]; Darkspear Raider's Cloak (272077, -1.70 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.9 | yes | Robe of Power (7054, -0.02 DPS) [crafted]; Dreamweave Vest (10021, -0.82 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.92 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 11.9 | yes | Mistscape Bracers (4045, -0.32 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.32 DPS) [quest]; Aurora Bracers (4043, -0.40 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (102.5 DPS) | yes | Stormcloth Gloves (10011, -0.85 DPS) [crafted]; Town Clerk's Mittens (270029, -1.17 DPS) [quest]; Red Mageweave Gloves (10018, -1.79 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (102.7 DPS) | yes | Gilded Cord (254037, -0.18 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.85 DPS) [quest]; Deathmage Sash (10771, -1.96 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.8 | yes | Crimson Silk Pantaloons (7062, -1.62 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -1.98 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.07 DPS) [dungeon]; Spidersilk Boots (4320, -2.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.9 | yes | Voodoo Band (1996, -2.12 DPS) [world_drop]; Mindbender Loop (5009, -2.12 DPS) [world_drop]; Black Widow Band (6199, -2.12 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 9.2 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 179.7 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Staff of Dar'Orahil (15106, -6.21 DPS) [quest]; Windweaver Staff (7757, -6.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Umbral Wand (5216) | World drop [world_drop] | sim-verified (101.8 DPS) | yes | Twisted Nether Wand (249144, -0.20 DPS) [crafted]; Jaina's Firestarter (13064, -1.05 DPS, sim-verified) [world_drop]; Burning Sliver (5249, -1.51 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Jordan; ranged: Umbral Wand

No-known-source sample (15 of 286, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10047 Simple Kilt; 10049 Diabolist's Blade; 14145 Cursed Felblade

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 143.7. Weights run: 1.4s. Verify run: 1.2s. 358 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.605), intellect=not significant (-0.468 ± 0.718), crit=2.617 ± 0.140, hit=4.742 ± 0.386, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 50.6 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -2.17 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -3.66 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.39 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.08 DPS) [world_drop]; Pendant of Myzrael (4614, -1.08 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 44.6 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -4.90 DPS) [dungeon]; Black Mageweave Shoulders (10027, -5.37 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Runecloth Cloak (13860, -0.77 DPS) [crafted]; Icy Cloak (4327, -1.08 DPS) [crafted]; Nightfall Drape (12465, -1.90 DPS, sim-verified) [dungeon] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 48.6 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -4.13 DPS) [world_drop]; Elemental Raiment (9434, -4.28 DPS) [world_drop] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, -0.15 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.31 DPS) [quest]; Nethergeld Cuffs (254061, -0.31 DPS) [crafted] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.77 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.77 DPS) [vendor]; Black Mageweave Gloves (10003, -0.79 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 45.6 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -4.90 DPS) [rep]; Ghostweave Cord (254073, -4.90 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 48.6 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -4.59 DPS) [crafted]; Red Mageweave Pants (10009, -5.37 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 55.4 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.87 DPS) [crafted]; Black Mageweave Boots (10026, -6.88 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.4 | yes | Lorekeeper's Ring (19523, -5.49 DPS) [rep]; Philanthropist's Ring (281635, -5.80 DPS) [quest]; Lorekeeper's Ring (19524, -5.95 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Lorekeeper's Ring (19523, -0.38 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.46 DPS) [quest]; Lorekeeper's Ring (19524, -0.62 DPS) [rep] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (140.3 DPS) | yes | Frozen Heart of the Mountain (249469, -0.78 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -0.93 DPS) [quest]; Tidal Charm (1404, -0.93 DPS) [vendor] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (140.3 DPS) | yes | Thunderbrew's Boot Flask (744, -1.86 DPS) [quest]; Tidal Charm (1404, -1.86 DPS) [vendor]; Frozen Heart of the Mountain (249469, -2.06 DPS, sim-verified) [crafted] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | sim-verified (140.3 DPS) | yes | Kindling Stave (11750, -0.66 DPS) [dungeon]; Hanzo Sword (8190, -2.49 DPS, sim-verified) [world_drop]; Soulkeeper (1607, -5.62 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (143.7 DPS) | yes | Wand of Allistarj (13065, -2.35 DPS) [world_drop]; Pyric Caduceus (11748, -3.13 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.79 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Arcane Runed Bracers; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Uther's Strength; trinket2: Abyss Shard; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 358, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 330.5. Weights run: 1.4s. Verify run: 1.2s. 722 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.964), intellect=not significant (1.610 ± 0.952), crit=5.970 ± 0.274, hit=10.408 ± 0.660, spell_haste=not significant (1.990 ± 0.697), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.383 ± 0.964), fire_power=1.387 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 374.7 | yes | Bloodvine Goggles (19999, -21.33 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -26.99 DPS) [quest]; Deathmist Mask (22074, -27.24 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (307.7 DPS) | yes | Blazefury Medallion (17111, -1.48 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -2.59 DPS) [quest]; Amulet of the Dawn (22657, -8.60 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 248.6 | yes | Rugged Mantle of the Timbermaw (227808, -6.40 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -16.03 DPS) [crafted]; Heretic Mantle (240150, -16.42 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 133.0 | yes | Howler's Furs (272414, -3.64 DPS) [vendor]; Stalwart Cloak (272415, -3.64 DPS) [vendor]; Earthweave Cloak (21187, -5.01 DPS, sim-verified) [quest] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 259.5 | yes | Heretic Garb (240146, -4.23 DPS) [vendor]; Bloodvine Vest (19682, -10.81 DPS, sim-verified) [crafted]; Earthpower Vest (21183, -14.58 DPS) [quest] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 151.0 | yes | Rockfury Bracers (21186, -0.28 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -4.15 DPS) [vendor]; Dryad's Wrist Bindings (19595, -14.65 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | sim-verified (317.8 DPS) | yes | Deathmist Wraps (22077, -3.03 DPS) [quest]; Deathmist Wraps (226911, -3.03 DPS) [quest]; Gloves of Spell Mastery (14146, -8.81 DPS, sim-verified) [crafted] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 332.2 | yes | Heretic Waistguard (240151, -9.55 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -21.15 DPS) [vendor]; Belt of the Archmage (18405, -25.59 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 353.6 | yes | Heretic Pants (240149, -14.34 DPS, sim-verified) [vendor]; Sentinel's Silk Leggings (237815, -17.45 DPS) [vendor]; Bloodvine Leggings (19683, -25.60 DPS) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 162.6 | yes | Sergeant Major's Dreadweave Boots (220891, -4.55 DPS) [vendor]; First Sergeant's Dreadweave Boots (220909, -4.55 DPS) [vendor]; Bloodvine Boots (19684, -6.07 DPS, sim-verified) [crafted] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (307.7 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Don Julio's Band (19325, -2.54 DPS, sim-verified) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (318.3 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Signet Ring of the Bronze Dragonflight (234028, -0.46 DPS) [vendor]; Band of Earthen Might (21182, -9.39 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (306.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.01 DPS) [world_drop]; Abyss Shard (20534, -6.65 DPS, sim-verified) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (306.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.02 DPS) [world_drop]; Abyss Shard (20534, -4.12 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (307.7 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Electrified Dagger (19100, -3.83 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Stormrager (16997) | Order Must Be Restored [quest] | sim-verified (313.0 DPS) | yes | Brilliant Wand (249385, -1.57 DPS) [crafted]; Wand of Biting Cold (19108, -4.09 DPS, sim-verified) [quest]; Torch of Austen (13004, -6.92 DPS) [world_drop] |

**New at 60:** head: Heretic Cowl; neck: Beads of Ogre Might; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Wrath of Cenarius; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Stormrager

No-known-source sample (15 of 722, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 41.4. Weights run: 1.5s. Verify run: 1.3s. 123 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.092, intellect=not significant (0.400 ± 0.102), crit=0.478 ± 0.022, hit=1.278 ± 0.062, spell_haste=not significant (0.104 ± 0.089), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.749 ± 0.092, fire_power=0.252 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -1.06 DPS) [crafted]; Lucky Fishing Hat (19972, -1.06 DPS) [quest]; Shadow Goggles (4373, -3.24 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 | yes | Reinforced Woolen Shoulders (4315, -0.64 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.81 DPS) [crafted]; Slime-encrusted Pads (6461, -1.52 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.18 DPS) [dungeon]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.49 DPS) [crafted]; Green Woolen Vest (2582, -0.53 DPS) [crafted]; Gray Woolen Robe (2585, -1.49 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Owlbeard Bracers (16981, -0.11 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.30 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.32 DPS) [crafted]; Apothecary Gloves (10919, -0.53 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Keller's Girdle (2911, -0.42 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.03 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (41.4 DPS) | yes | Silk-threaded Trousers (1929, -0.25 DPS) [dungeon]; Abomination Skin Leggings (23173, -0.58 DPS, sim-verified) [dungeon]; Colorful Kilt (10048, -0.60 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Feather Padded Treads (285345, -0.71 DPS, sim-verified) [world]; Pristine Boots (253889, -0.78 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Loop of Sacrifice (281673, -0.53 DPS) [quest]; Volcanic Rock Ring (12053, -0.67 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Loop of Sacrifice (281673, -0.18 DPS) [quest]; Volcanic Rock Ring (12053, -0.32 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.48 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Living Root (6631) | Wailing Caverns: Verdan the Everliving [dungeon] | 120.2 | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.14 DPS) [quest]; Crescent Staff (6505, -0.87 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 129.4 | yes | Skycaller (12984, -2.39 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.55 DPS) [dungeon]; Sizzle Stick (8071, -4.31 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Living Root; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 123, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10047 Simple Kilt; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance; 209620 Insignia of the Horde; 241089 Scarlet Dagger; 254779 A'sharahm, the Roiling Tempest

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 63.1. Weights run: 1.5s. Verify run: 1.2s. 211 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.099, intellect=not significant (0.303 ± 0.115), crit=0.576 ± 0.031, hit=0.934 ± 0.062, spell_haste=not significant (0.340 ± 0.110), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.869 ± 0.099, fire_power=0.131 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.67 DPS) [crafted]; Embalmed Shroud (7691, -1.00 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.8 | yes | Darkspear Warding Pendant (272075, -1.42 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.54 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.54 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.7 | yes | Death Speaker Mantle (6685, -0.89 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -1.00 DPS) [quest]; Invoker's Mantle (215365, -1.07 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.33 DPS) [crafted]; Battle Healer's Cloak (19529, -0.33 DPS) [rep] |
| chest | Tree Bark Jacket (1486) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 13.0 | yes | Green Silk Armor (7065, +0.00 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.89 DPS) [dungeon]; Pristine Gown (253961, -1.30 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Glowing Magical Bracelets (13106, -1.23 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -2.40 DPS) [world_drop]; Tabitha's Cuffs (251486, -2.40 DPS) [quest] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.5 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gnoll Casting Gloves (892, -0.51 DPS) [world]; Truefaith Gloves (7049, -0.54 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 | yes | Warsong Sash (16975, +0.00 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.67 DPS) [dungeon]; Invoker's Cord (215366, -1.13 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.96 DPS) [crafted]; Filigreed Pristine Leggings (253937, -1.40 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.1 | yes | Acidic Walkers (9454, -0.57 DPS) [dungeon]; Boots of the Enchanter (4325, -1.38 DPS) [crafted]; Spidersilk Boots (4320, -1.69 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.67 DPS) [rep]; Electrocutioner Lagnut (9447, -1.34 DPS) [dungeon]; Sludge-Stained Band (286535, -1.34 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -1.00 DPS) [world]; Black Widow Band (6199, -1.30 DPS) [world]; Electrocutioner Lagnut (9447, -1.69 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (63.1 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Insignia of the Horde (18852, +0.00 DPS) [vendor] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 76.8 | yes | Lorekeeper's Staff (212580, -0.40 DPS) [vendor]; Advisor's Gnarled Staff (212584, -0.40 DPS) [vendor]; Gnarled Ash Staff (791, -0.55 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 102.4 | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.54 DPS) [crafted]; Thunderwood (13062, -5.10 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 211, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 19545 Scout's Blade

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 106.1. Weights run: 1.3s. Verify run: 1.2s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.273), intellect=1.317 ± 0.307, crit=1.241 ± 0.061, hit=2.170 ± 0.167, spell_haste=not significant (0.432 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.422 ± 0.273), fire_power=0.574 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (102.7 DPS) | yes | Craftsman's Monocle (4393, -0.30 DPS) [crafted]; Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Augural Shroud (2620, -2.22 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.9 | yes | Darkspear Warding Pendant (272074, -1.39 DPS) [vendor]; Darkspear Warding Pendant (272075, -2.03 DPS) [vendor]; Necklace of Calisea (1714, -2.04 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.1 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.80 DPS) [dungeon]; Death Speaker Mantle (6685, -0.89 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (101.5 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Cloak of Rot (4462, -0.50 DPS) [world]; Darkspear Raider's Cloak (272077, -1.04 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.9 | yes | Robe of Power (7054, -0.02 DPS) [crafted]; Green Silk Armor (7065, -0.92 DPS) [crafted]; Dreamweave Vest (10021, -1.05 DPS, sim-verified) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 14.5 | yes | Aurora Bracers (4043, -0.98 DPS) [world_drop]; Mistscape Bracers (4045, -0.98 DPS) [world_drop]; Windchaser Cuffs (14429, -1.13 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (101.8 DPS) | yes | Stormcloth Gloves (10011, -0.85 DPS) [crafted]; Red Mageweave Gloves (10018, -1.33 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.48 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (102.0 DPS) | yes | Gilded Cord (254037, -0.18 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.85 DPS) [quest]; Deathmage Sash (10771, -1.48 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.8 | yes | Crimson Silk Pantaloons (7062, -1.69 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -1.83 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.07 DPS) [dungeon]; Spidersilk Boots (4320, -2.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.9 | yes | Voodoo Band (1996, -2.12 DPS) [world_drop]; Mindbender Loop (5009, -2.12 DPS) [world_drop]; Black Widow Band (6199, -2.12 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 9.2 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Jordan (873) | World drop [world_drop] | 179.7 | yes | Illusionary Rod (7713, +0.00 DPS, sim-verified) [dungeon]; Staff of Dar'Orahil (15106, -6.21 DPS) [quest]; Windweaver Staff (7757, -6.70 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 168.9 | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -5.85 DPS) [crafted]; Starfaller (13063, -7.16 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Staff of Jordan; ranged: Jaina's Firestarter

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10047 Simple Kilt; 10049 Diabolist's Blade

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 140.3. Weights run: 1.4s. Verify run: 1.2s. 360 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.605), intellect=not significant (-0.468 ± 0.718), crit=2.617 ± 0.140, hit=4.742 ± 0.386, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) (or Blood Guard's Dreadweave Hat (220907)) | Captain Dirgehammer [vendor] | 50.6 | yes | Blood Guard's Dreadweave Hat (220907, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -2.17 DPS) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -3.66 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.08 DPS) [world_drop]; Choker of the High Shaman (4112, -1.08 DPS) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) (or Blood Guard's Dreadweave Mantle (220905)) | Captain Dirgehammer [vendor] | 44.6 | yes | Blood Guard's Dreadweave Mantle (220905, +0.00 DPS, sim-verified) [vendor]; Rotgrip Mantle (17732, -4.90 DPS) [dungeon]; Black Mageweave Shoulders (10027, -5.37 DPS) [crafted] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Deep Woodlands Cloak (19121, +0.00 DPS, sim-verified) [quest]; Nightfall Drape (12465, -0.77 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted] |
| chest | Knight's Dreadweave Vest (220886) (or Stone Guard's Dreadweave Vest (220904)) | Captain Dirgehammer [vendor] | 48.6 | yes | Stone Guard's Dreadweave Vest (220904, +0.00 DPS, sim-verified) [vendor]; Robe of the Magi (1716, -4.13 DPS) [world_drop]; Elemental Raiment (9434, -4.28 DPS) [world_drop] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nethergeld Cuffs (254061, -0.31 DPS) [crafted]; Bloodband Bracers (11469, -0.62 DPS) [quest]; Condor Bracers (15864, -0.90 DPS, sim-verified) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.77 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.77 DPS) [vendor]; Black Mageweave Gloves (10003, -1.25 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 45.6 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -4.90 DPS) [rep]; Ghostweave Cord (254073, -4.90 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) (or Stone Guard's Dreadweave Leggings (220906)) | Captain Dirgehammer [vendor] | 48.6 | yes | Stone Guard's Dreadweave Leggings (220906, +0.00 DPS, sim-verified) [vendor]; Wizardweave Leggings (14132, -4.59 DPS) [crafted]; Red Mageweave Pants (10009, -5.37 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) (or First Sergeant's Dreadweave Boots (220909)) | Captain Dirgehammer [vendor] | 55.4 | yes | First Sergeant's Dreadweave Boots (220909, +0.00 DPS, sim-verified) [vendor]; Earthen Silk Slippers (254013, -4.87 DPS) [crafted]; Black Mageweave Boots (10026, -6.88 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 47.4 | yes | Advisor's Ring (19519, -5.49 DPS) [rep]; Philanthropist's Ring (281635, -5.80 DPS) [quest]; Advisor's Ring (19520, -5.95 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Advisor's Ring (19519, +0.00 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.46 DPS) [quest]; Advisor's Ring (19520, -0.62 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (137.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Uther's Strength (11302, -0.93 DPS) [world_drop]; Tidal Charm (1404, -1.86 DPS) [vendor] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (137.2 DPS) | yes | Rune of the Guard Captain (19120, -0.17 DPS, sim-verified) [quest]; Uther's Strength (11302, -5.68 DPS) [world_drop]; Tidal Charm (1404, -6.61 DPS) [vendor] |
| main_hand | Soul Harvester (20536) | Trolls of a Feather [quest] | sim-verified (137.2 DPS) | yes | Kindling Stave (11750, -0.66 DPS) [dungeon]; Hanzo Sword (8190, -1.36 DPS, sim-verified) [world_drop]; Soulkeeper (1607, -5.62 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (140.3 DPS) | yes | Wand of Allistarj (13065, -2.35 DPS) [world_drop]; Pyric Caduceus (11748, -2.59 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.79 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Frozen Heart of the Mountain; main_hand: Soul Harvester; ranged: Noxious Shooter

No-known-source sample (15 of 360, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 328.5. Weights run: 1.4s. Verify run: 1.2s. 725 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.964), intellect=not significant (1.610 ± 0.952), crit=5.970 ± 0.274, hit=10.408 ± 0.660, spell_haste=not significant (1.990 ± 0.697), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.383 ± 0.964), fire_power=1.387 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 374.7 | yes | Bloodvine Goggles (19999, -21.11 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -26.99 DPS) [quest]; Deathmist Mask (22074, -27.24 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (305.5 DPS) | yes | Blazefury Medallion (17111, -1.87 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -2.59 DPS) [quest]; Amulet of the Dawn (22657, -8.60 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 248.6 | yes | Rugged Mantle of the Timbermaw (227808, -8.87 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -16.03 DPS) [crafted]; Heretic Mantle (240150, -16.42 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 133.0 | yes | Howler's Furs (272414, -3.64 DPS) [vendor]; Stalwart Cloak (272415, -3.64 DPS) [vendor]; Earthweave Cloak (21187, -4.83 DPS, sim-verified) [quest] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 259.5 | yes | Heretic Garb (240146, -4.23 DPS) [vendor]; Bloodvine Vest (19682, -11.51 DPS, sim-verified) [crafted]; Earthpower Vest (21183, -14.58 DPS) [quest] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 151.0 | yes | Rockfury Bracers (21186, +0.00 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -4.15 DPS) [vendor]; Dryad's Wrist Bindings (19595, -14.65 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | sim-verified (317.3 DPS) | yes | Deathmist Wraps (22077, -3.03 DPS) [quest]; Deathmist Wraps (226911, -3.03 DPS) [quest]; Gloves of Spell Mastery (14146, -10.07 DPS, sim-verified) [crafted] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 332.2 | yes | Heretic Waistguard (240151, -9.37 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -21.15 DPS) [vendor]; Belt of the Archmage (18405, -25.59 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 353.6 | yes | Heretic Pants (240149, -13.23 DPS, sim-verified) [vendor]; Sentinel's Silk Leggings (237815, -17.45 DPS) [vendor]; Bloodvine Leggings (19683, -25.60 DPS) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 162.6 | yes | First Sergeant's Dreadweave Boots (220909, -4.55 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -4.55 DPS) [vendor]; Bloodvine Boots (19684, -6.52 DPS, sim-verified) [crafted] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (305.5 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Channeler's Ring (272406, +0.00 DPS) [vendor]; Don Julio's Band (19325, -3.76 DPS, sim-verified) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (315.8 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Signet Ring of the Bronze Dragonflight (234028, -0.46 DPS) [vendor]; Band of Earthen Might (21182, -8.65 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (299.1 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Abyss Shard (20534, -0.25 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (304.2 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Abyss Shard (20534, -1.26 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.41 DPS, sim-verified) [crafted] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (305.5 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Glacial Blade (19099, -3.85 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | Stormrager (16997) | The Scarlet Oracle, Demetria [quest] | sim-verified (311.1 DPS) | yes | Brilliant Wand (249385, -1.57 DPS) [crafted]; Wand of Biting Cold (19108, -3.90 DPS, sim-verified) [quest]; Torch of Austen (13004, -6.92 DPS) [world_drop] |

**New at 60:** head: Heretic Cowl; neck: Beads of Ogre Might; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Wrath of Cenarius; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Stormrager

No-known-source sample (15 of 725, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

