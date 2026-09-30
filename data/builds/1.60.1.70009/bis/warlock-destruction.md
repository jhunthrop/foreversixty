# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 43.3. Weights run: 1.5s. Verify run: 1.3s. 126 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.092, intellect=not significant (0.400 ± 0.102), crit=0.478 ± 0.022, hit=1.160 ± 0.011, spell_haste=not significant (0.104 ± 0.089), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.749 ± 0.092, fire_power=0.252 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -1.06 DPS) [crafted]; Lucky Fishing Hat (19972, -1.06 DPS) [quest]; Shadow Goggles (4373, -3.27 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 | yes | Reinforced Woolen Shoulders (4315, -0.05 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.81 DPS) [crafted]; Slime-encrusted Pads (6461, -1.52 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.18 DPS) [dungeon]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.53 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.49 DPS) [crafted]; Green Woolen Vest (2582, -0.53 DPS) [crafted]; Gray Woolen Robe (2585, -1.45 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.14 DPS) [world_drop]; Windsong Bangles (263336, -0.25 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Gnoll Casting Gloves (892, -0.27 DPS, sim-verified) [world]; Pristine Gloves (253913, -0.32 DPS) [crafted]; Blight Gloves (279877, -0.74 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Keller's Girdle (2911, -0.42 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.09 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (43.3 DPS) | yes | Silk-threaded Trousers (1929, -0.25 DPS) [dungeon]; Colorful Kilt (10048, -0.60 DPS) [crafted]; Abomination Skin Leggings (23173, -0.85 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Feather Padded Treads (285345, -0.63 DPS, sim-verified) [world]; Pristine Boots (253889, -0.78 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 | yes | Sludge-Stained Band (286535, -0.49 DPS) [world]; Lavishly Jeweled Ring (1156, -0.60 DPS) [dungeon]; Loop of Sacrifice (281673, -0.67 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Loop of Sacrifice (281673, -0.53 DPS) [quest]; Sludge-Stained Band (286535, -0.71 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 4.0 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.14 DPS) [world]; Lesser Staff of the Spire (1300, -0.28 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 129.4 | yes | Skycaller (12984, -2.36 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.55 DPS) [dungeon]; Sizzle Stick (8071, -4.31 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 126, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18852 Insignia of the Horde; 20434 Lorekeeper's Staff

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 64.9. Weights run: 1.6s. Verify run: 1.2s. 216 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.099, intellect=not significant (0.303 ± 0.115), crit=0.576 ± 0.031, hit=0.865 ± 0.009, spell_haste=not significant (0.340 ± 0.110), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.869 ± 0.099, fire_power=0.131 ± 0.000

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
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 3.3 | yes | Twisted Chanter's Staff (890, +0.00 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.10 DPS) [quest]; Channeler's Staff (4437, -0.30 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 102.4 | yes | Starfaller (13063, -0.66 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.54 DPS) [crafted]; Thunderwood (13062, -5.10 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Relentless Raider's Seal; trinket2: Rune of Perfection; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 108.3. Weights run: 1.4s. Verify run: 1.2s. 293 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.273), intellect=1.317 ± 0.307, crit=1.241 ± 0.061, hit=1.975 ± 0.022, spell_haste=not significant (0.432 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.422 ± 0.273), fire_power=0.574 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (103.6 DPS) | yes | Craftsman's Monocle (4393, -0.30 DPS) [crafted]; Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Augural Shroud (2620, -1.86 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.9 | yes | Darkspear Warding Pendant (272074, -1.39 DPS) [vendor]; Necklace of Calisea (1714, -1.71 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.03 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.1 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.80 DPS) [dungeon]; Death Speaker Mantle (6685, -0.89 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (103.0 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Cloak of Rot (4462, -0.50 DPS) [world]; Darkspear Raider's Cloak (272077, -1.23 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.9 | yes | Robe of Power (7054, -0.02 DPS) [crafted]; Dreamweave Vest (10021, -0.87 DPS, sim-verified) [crafted]; Green Silk Armor (7065, -0.92 DPS) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 11.9 | yes | Aurora Bracers (4043, -0.31 DPS, sim-verified) [world_drop]; Mistscape Bracers (4045, -0.32 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.32 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (103.2 DPS) | yes | Stormcloth Gloves (10011, -0.85 DPS) [crafted]; Town Clerk's Mittens (270029, -1.17 DPS) [quest]; Red Mageweave Gloves (10018, -1.41 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (103.4 DPS) | yes | Gilded Cord (254037, -0.18 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.85 DPS) [quest]; Deathmage Sash (10771, -1.63 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.8 | yes | Crimson Silk Pantaloons (7062, -1.67 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -2.07 DPS) [dungeon]; Gilded Slippers (254001, -2.26 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.9 | yes | Voodoo Band (1996, -2.12 DPS) [world_drop]; Mindbender Loop (5009, -2.12 DPS) [world_drop]; Black Widow Band (6199, -2.12 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 9.2 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 34.2 | yes | Illusionary Rod (7713, -0.31 DPS, sim-verified) [dungeon]; Staff of Noh'Orahil (15105, -2.84 DPS) [quest]; Windweaver Staff (7757, -3.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 168.9 | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -5.85 DPS) [crafted]; Burning Sliver (5249, -7.16 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 293, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 142.7. Weights run: 1.2s. Verify run: 1.2s. 373 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.605), intellect=not significant (-0.468 ± 0.718), crit=2.617 ± 0.140, hit=4.179 ± 0.053, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 50.6 | yes | Spellpower Goggles Xtreme Plus (15999, -3.66 DPS) [crafted]; Eye of Theradras (17715, -3.90 DPS, sim-verified) [dungeon]; Dreamweave Circlet (10041, -4.59 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.78 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.08 DPS) [world_drop]; Pendant of Myzrael (4614, -1.08 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 44.6 | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -5.37 DPS) [crafted]; Bloodmage Mantle (7684, -5.52 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Runecloth Cloak (13860, -0.77 DPS) [crafted]; Icy Cloak (4327, -1.08 DPS) [crafted]; Nightfall Drape (12465, -2.89 DPS, sim-verified) [dungeon] |
| chest | Knight's Dreadweave Vest (220886) | Captain Dirgehammer [vendor] | 48.6 | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -4.28 DPS) [world_drop]; Acumen Robes (17775, -4.59 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Condor Bracers (15864, -0.31 DPS) [quest]; Nethergeld Cuffs (254061, -0.31 DPS) [crafted]; Spidertank Oilrag (9448, -0.51 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Brightcloth Gloves (14101, -0.77 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.77 DPS) [vendor]; Black Mageweave Gloves (10003, -1.67 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 45.6 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -4.90 DPS) [rep]; Ghostweave Cord (254073, -4.90 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 48.6 | yes | Wizardweave Leggings (14132, -0.38 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -5.37 DPS) [crafted]; Gaze Dreamer Pants (6903, -5.68 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 49.8 | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -6.01 DPS) [crafted]; Gilded Sandals (254107, -6.01 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.8 | yes | Lorekeeper's Ring (19523, -4.62 DPS) [rep]; Philanthropist's Ring (281635, -4.93 DPS) [quest]; Lorekeeper's Ring (19524, -5.08 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Philanthropist's Ring (281635, -0.46 DPS) [quest]; Lorekeeper's Ring (19523, -0.54 DPS, sim-verified) [rep]; Lorekeeper's Ring (19524, -0.62 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (140.5 DPS) | yes | Uther's Strength (11302, -0.13 DPS, sim-verified) [world_drop]; Thunderbrew's Boot Flask (744, -5.83 DPS) [quest]; Tidal Charm (1404, -5.83 DPS) [vendor] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (140.5 DPS) | yes | Thunderbrew's Boot Flask (744, -1.86 DPS) [quest]; Tidal Charm (1404, -1.86 DPS) [vendor]; Uther's Strength (11302, -2.07 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (140.5 DPS) | yes | Soul Harvester (20536, +0.00 DPS) [quest]; Illusionary Rod (7713, -0.80 DPS) [dungeon]; Shortsword of Vengeance (754, -2.35 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (142.7 DPS) | yes | Wand of Allistarj (13065, -2.35 DPS) [world_drop]; Pyric Caduceus (11748, -2.41 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.79 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Arcane Runed Bracers; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Abyss Shard; ranged: Noxious Shooter

No-known-source sample (15 of 373, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 324.9. Weights run: 1.3s. Verify run: 1.2s. 769 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.964), intellect=not significant (1.610 ± 0.952), crit=5.970 ± 0.274, hit=8.863 ± 0.090, spell_haste=not significant (1.990 ± 0.697), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.383 ± 0.964), fire_power=1.387 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 343.8 | yes | Bloodvine Goggles (19999, -21.68 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -25.04 DPS) [quest]; Deathmist Mask (22074, -25.30 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (306.5 DPS) | yes | Medallion of the Dawn (22659, -0.64 DPS) [quest]; Blazefury Medallion (17111, -2.04 DPS, sim-verified) [world]; Amulet of the Dawn (22657, -6.65 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 233.1 | yes | Rugged Mantle of the Timbermaw (227808, -7.79 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -14.08 DPS) [crafted]; Heretic Mantle (240150, -14.47 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 117.5 | yes | Howler's Furs (272414, -3.64 DPS) [vendor]; Stalwart Cloak (272415, -3.64 DPS) [vendor]; Earthweave Cloak (21187, -6.27 DPS, sim-verified) [quest] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 259.5 | yes | Heretic Garb (240146, -6.18 DPS) [vendor]; Bloodvine Vest (19682, -12.21 DPS, sim-verified) [crafted]; Earthpower Vest (21183, -14.58 DPS) [quest] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 135.6 | yes | Rockfury Bracers (21186, -2.41 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -4.15 DPS) [vendor]; Dryad's Wrist Bindings (19595, -12.70 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | sim-verified (316.0 DPS) | yes | Deathmist Wraps (22077, -3.03 DPS) [quest]; Deathmist Wraps (226911, -3.03 DPS) [quest]; Gloves of Spell Mastery (14146, -8.14 DPS, sim-verified) [crafted] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 316.7 | yes | Heretic Waistguard (240151, -10.01 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -21.15 DPS) [vendor]; Belt of the Archmage (18405, -23.64 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 338.2 | yes | Sentinel's Silk Leggings (237815, -12.29 DPS, sim-verified) [vendor]; Heretic Pants (240149, -16.23 DPS) [vendor]; Bloodvine Leggings (19683, -25.60 DPS) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 147.2 | yes | Heretic Boots (240153, -4.13 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -4.55 DPS) [vendor]; Bloodvine Boots (19684, -6.38 DPS, sim-verified) [crafted] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (306.5 DPS) | yes | Mindtear Band (20632, +0.00 DPS) [world]; Band of Earthen Might (21182, +0.00 DPS) [quest]; Don Julio's Band (19325, -4.07 DPS, sim-verified) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (317.5 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Signet Ring of the Bronze Dragonflight (234028, -0.46 DPS) [vendor]; Band of Earthen Might (21182, -9.66 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (308.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Uther's Strength (11302, -1.01 DPS) [world_drop] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (305.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS, sim-verified) [vendor]; Uther's Strength (11302, -0.76 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (306.5 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -4.93 DPS, sim-verified) [world_drop]; Soul Harvester (20536, -8.69 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 614.4 | yes | Wand of Biting Cold (19108, -3.05 DPS, sim-verified) [quest]; Stormrager (16997, -13.81 DPS) [quest]; Brilliant Wand (249385, -15.38 DPS) [crafted] |

**New at 60:** head: Heretic Cowl; neck: Beads of Ogre Might; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Wrath of Cenarius; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 769, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 42.0. Weights run: 1.5s. Verify run: 1.3s. 122 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.092, intellect=not significant (0.400 ± 0.102), crit=0.478 ± 0.022, hit=1.160 ± 0.011, spell_haste=not significant (0.104 ± 0.089), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.749 ± 0.092, fire_power=0.252 ± 0.000

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -1.06 DPS) [crafted]; Lucky Fishing Hat (19972, -1.06 DPS) [quest]; Shadow Goggles (4373, -3.13 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.6 | yes | Reinforced Woolen Shoulders (4315, -0.12 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.81 DPS) [crafted]; Slime-encrusted Pads (6461, -1.52 DPS) [dungeon] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.18 DPS) [dungeon]; Black Whelp Cloak (7283, -0.18 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.32 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.0 | yes | Green Woolen Robe (6243, -0.49 DPS) [crafted]; Green Woolen Vest (2582, -0.53 DPS) [crafted]; Gray Woolen Robe (2585, -1.35 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 2.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.07 DPS) [quest]; Owlbeard Bracers (16981, -0.11 DPS) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Pristine Gloves (253913, -0.32 DPS) [crafted]; Gnoll Casting Gloves (892, -0.43 DPS, sim-verified) [world]; Apothecary Gloves (10919, -0.53 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 5.6 | yes | Keller's Girdle (2911, -0.42 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.42 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.87 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (42.0 DPS) | yes | Silk-threaded Trousers (1929, -0.25 DPS) [dungeon]; Colorful Kilt (10048, -0.60 DPS) [crafted]; Abomination Skin Leggings (23173, -0.84 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.6 | yes | Feather Padded Treads (285345, -0.39 DPS, sim-verified) [world]; Pristine Boots (253889, -0.78 DPS) [crafted]; Red Woolen Boots (4313, -0.81 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Lavishly Jeweled Ring (1156, -0.46 DPS) [dungeon]; Loop of Sacrifice (281673, -0.53 DPS) [quest]; Volcanic Rock Ring (12053, -0.67 DPS) [world_drop] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | 3.0 | yes | Loop of Sacrifice (281673, -0.18 DPS) [quest]; Volcanic Rock Ring (12053, -0.32 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.69 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 4.0 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.14 DPS) [world]; Lesser Staff of the Spire (1300, -0.28 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 129.4 | yes | Skycaller (12984, -2.31 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.55 DPS) [dungeon]; Sizzle Stick (8071, -4.31 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 122, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance; 241089 Scarlet Dagger; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 63.1. Weights run: 1.6s. Verify run: 1.2s. 212 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.099, intellect=not significant (0.303 ± 0.115), crit=0.576 ± 0.031, hit=0.865 ± 0.009, spell_haste=not significant (0.340 ± 0.110), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.869 ± 0.099, fire_power=0.131 ± 0.000

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
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (63.1 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -0.36 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 3.3 | yes | Twisted Chanter's Staff (890, -0.01 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.10 DPS) [quest]; Channeler's Staff (4437, -0.30 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 102.4 | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.54 DPS) [crafted]; Thunderwood (13062, -5.10 DPS) [world_drop] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Tree Bark Jacket; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 212, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 106.0. Weights run: 1.4s. Verify run: 1.2s. 289 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.273), intellect=1.317 ± 0.307, crit=1.241 ± 0.061, hit=1.975 ± 0.022, spell_haste=not significant (0.432 ± 0.175), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.422 ± 0.273), fire_power=0.574 ± 0.002

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | sim-verified (103.9 DPS) | yes | Craftsman's Monocle (4393, -0.30 DPS) [crafted]; Enchanter's Cowl (4322, -0.45 DPS) [crafted]; Augural Shroud (2620, -2.32 DPS, sim-verified) [world] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.9 | yes | Darkspear Warding Pendant (272074, -1.39 DPS) [vendor]; Necklace of Calisea (1714, -1.83 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272075, -2.03 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 24.1 | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.80 DPS) [dungeon]; Death Speaker Mantle (6685, -0.89 DPS) [dungeon] |
| back | Darkspear Raider's Cloak (272077) | Creeg Bothunk [vendor] | 14.5 | yes | Long Silken Cloak (4326, +0.00 DPS, sim-verified) [crafted]; Guardian Cloak (5965, -0.47 DPS) [crafted]; Cloak of Rot (4462, -0.97 DPS) [world] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.9 | yes | Robe of Power (7054, -0.02 DPS) [crafted]; Green Silk Armor (7065, -0.92 DPS) [crafted]; Dreamweave Vest (10021, -0.95 DPS, sim-verified) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 14.5 | yes | Aurora Bracers (4043, -0.98 DPS) [world_drop]; Mistscape Bracers (4045, -0.98 DPS) [world_drop]; Windchaser Cuffs (14429, -1.04 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (102.9 DPS) | yes | Stormcloth Gloves (10011, -0.85 DPS) [crafted]; Red Mageweave Gloves (10018, -1.30 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.48 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (103.0 DPS) | yes | Gilded Cord (254037, -0.18 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.85 DPS) [quest]; Deathmage Sash (10771, -1.38 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.8 | yes | Crimson Silk Pantaloons (7062, -1.46 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -2.51 DPS) [dungeon]; Stoneweaver Leggings (9407, -3.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Acidic Walkers (9454, -2.07 DPS) [dungeon]; Gilded Slippers (254001, -2.29 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -2.87 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.9 | yes | Voodoo Band (1996, -2.12 DPS) [world_drop]; Mindbender Loop (5009, -2.12 DPS) [world_drop]; Black Widow Band (6199, -2.12 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 9.2 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world_drop]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | 34.2 | yes | Illusionary Rod (7713, -0.92 DPS, sim-verified) [dungeon]; Staff of Noh'Orahil (15105, -2.84 DPS) [quest]; Windweaver Staff (7757, -3.54 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 168.9 | yes | Umbral Wand (5216, -0.01 DPS, sim-verified) [world_drop]; Twisted Nether Wand (249144, -5.85 DPS) [crafted]; Starfaller (13063, -7.16 DPS) [world_drop] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Darkspear Raider's Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Ogremind Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 289, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 10049 Diabolist's Blade; 14389 Durability Shoulderpads

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 144.4. Weights run: 1.2s. Verify run: 1.2s. 369 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.605), intellect=not significant (-0.468 ± 0.718), crit=2.617 ± 0.140, hit=4.179 ± 0.053, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 50.6 | yes | Eye of Theradras (17715, -2.84 DPS, sim-verified) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -3.66 DPS) [crafted]; Dreamweave Circlet (10041, -4.59 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 | yes | Mindburst Medallion (11196, -0.38 DPS, sim-verified) [quest]; Necklace of Calisea (1714, -1.08 DPS) [world_drop]; Choker of the High Shaman (4112, -1.08 DPS) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (138.6 DPS) | yes | Black Mageweave Shoulders (10027, -0.46 DPS) [crafted]; Bloodmage Mantle (7684, -0.62 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.92 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 | yes | Deep Woodlands Cloak (19121, -0.15 DPS, sim-verified) [quest]; Nightfall Drape (12465, -0.77 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (138.8 DPS) | yes | Elemental Raiment (9434, -0.15 DPS) [world_drop]; Acumen Robes (17775, -0.46 DPS) [quest]; Stone Guard's Dreadweave Vest (220904, -2.13 DPS, sim-verified) [vendor] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Condor Bracers (15864, +0.00 DPS, sim-verified) [quest]; Nethergeld Cuffs (254061, -0.31 DPS) [crafted]; Bloodband Bracers (11469, -0.62 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 | yes | Black Mageweave Gloves (10003, +0.00 DPS, sim-verified) [crafted]; Brightcloth Gloves (14101, -0.77 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -0.77 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 45.6 | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -4.90 DPS) [rep]; Ghostweave Cord (254073, -4.90 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 48.6 | yes | Wizardweave Leggings (14132, +0.00 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -5.37 DPS) [crafted]; Gaze Dreamer Pants (6903, -5.68 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (140.2 DPS) | yes | Black Mageweave Boots (10026, -2.01 DPS) [crafted]; Gilded Sandals (254107, -2.01 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.53 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.8 | yes | Advisor's Ring (19519, -4.62 DPS) [rep]; Philanthropist's Ring (281635, -4.93 DPS) [quest]; Advisor's Ring (19520, -5.08 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 | yes | Advisor's Ring (19519, +0.00 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.46 DPS) [quest]; Advisor's Ring (19520, -0.62 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (136.8 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.26 DPS, sim-verified) [world_drop]; Tidal Charm (1404, -1.86 DPS) [vendor] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (136.8 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Tidal Charm (1404, -4.53 DPS) [vendor] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (136.8 DPS) | yes | Soul Harvester (20536, +0.00 DPS) [quest]; Shortsword of Vengeance (754, -0.59 DPS, sim-verified) [world_drop]; Illusionary Rod (7713, -0.80 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Noxious Shooter (17745) | Maraudon: Noxxion [dungeon] | sim-verified (139.8 DPS) | yes | Wand of Allistarj (13065, -2.35 DPS) [world_drop]; Pyric Caduceus (11748, -3.07 DPS, sim-verified) [dungeon]; Lesser Eternal Wand (249232, -3.79 DPS) [crafted] |

**New at 50:** head: Blood Guard's Dreadweave Hat; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Rune of the Guard Captain; ranged: Noxious Shooter

No-known-source sample (15 of 369, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 326.0. Weights run: 1.3s. Verify run: 1.2s. 766 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.964), intellect=not significant (1.610 ± 0.952), crit=5.970 ± 0.274, hit=8.863 ± 0.090, spell_haste=not significant (1.990 ± 0.697), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.383 ± 0.964), fire_power=1.387 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 343.8 | yes | Bloodvine Goggles (19999, -19.75 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -25.04 DPS) [quest]; Deathmist Mask (22074, -25.30 DPS) [quest] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (307.2 DPS) | yes | Medallion of the Dawn (22659, -0.64 DPS) [quest]; Blazefury Medallion (17111, -0.65 DPS, sim-verified) [world]; Amulet of the Dawn (22657, -6.65 DPS) [quest] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 233.1 | yes | Rugged Mantle of the Timbermaw (227808, -6.63 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -14.08 DPS) [crafted]; Heretic Mantle (240150, -14.47 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 117.5 | yes | Howler's Furs (272414, -3.64 DPS) [vendor]; Stalwart Cloak (272415, -3.64 DPS) [vendor]; Earthweave Cloak (21187, -4.33 DPS, sim-verified) [quest] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 259.5 | yes | Heretic Garb (240146, -6.18 DPS) [vendor]; Bloodvine Vest (19682, -11.75 DPS, sim-verified) [crafted]; Earthpower Vest (21183, -14.58 DPS) [quest] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 135.6 | yes | Rockfury Bracers (21186, +0.00 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -4.15 DPS) [vendor]; Dryad's Wrist Bindings (19595, -12.70 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | sim-verified (316.5 DPS) | yes | Deathmist Wraps (22077, -3.03 DPS) [quest]; Deathmist Wraps (226911, -3.03 DPS) [quest]; Gloves of Spell Mastery (14146, -8.93 DPS, sim-verified) [crafted] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 316.7 | yes | Heretic Waistguard (240151, -8.61 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -21.15 DPS) [vendor]; Belt of the Archmage (18405, -23.64 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 338.2 | yes | Sentinel's Silk Leggings (237815, -9.25 DPS, sim-verified) [vendor]; Heretic Pants (240149, -16.23 DPS) [vendor]; Bloodvine Leggings (19683, -25.60 DPS) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 147.2 | yes | Heretic Boots (240153, -4.13 DPS) [vendor]; First Sergeant's Dreadweave Boots (220909, -4.55 DPS) [vendor]; Bloodvine Boots (19684, -5.42 DPS, sim-verified) [crafted] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (307.2 DPS) | yes | Mindtear Band (20632, +0.00 DPS) [world]; Band of Earthen Might (21182, +0.00 DPS) [quest]; Don Julio's Band (19325, -2.74 DPS, sim-verified) [rep] |
| finger2 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (316.3 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Signet Ring of the Bronze Dragonflight (234028, -0.46 DPS) [vendor]; Band of Earthen Might (21182, -8.79 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (301.0 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Abyss Shard (20534, -0.25 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (307.1 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Abyss Shard (20534, -3.17 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (307.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -3.77 DPS, sim-verified) [world_drop]; Soul Harvester (20536, -8.69 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 614.4 | yes | Wand of Biting Cold (19108, -0.33 DPS, sim-verified) [quest]; Stormrager (16997, -13.81 DPS) [quest]; Brilliant Wand (249385, -15.38 DPS) [crafted] |

**New at 60:** head: Heretic Cowl; neck: Beads of Ogre Might; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Wrath of Cenarius; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 766, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector

