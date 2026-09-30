# Leveling BiS: Frost

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 37.1. Weights run: 0.6s. Verify run: 0.7s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.741 ± 0.174, crit=0.951 ± 0.045, hit=1.870 ± 0.020, spell_haste=not significant (-0.108 ± 0.181), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.78 DPS) [crafted]; Lucky Fishing Hat (19972, -0.78 DPS) [quest]; Shadow Goggles (4373, -1.56 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.7 | yes | Reinforced Woolen Shoulders (4315, -0.43 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.99 DPS) [crafted]; Slime-encrusted Pads (6461, -1.51 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.16 DPS) [dungeon]; Black Whelp Cloak (7283, -0.16 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.7 | yes | Green Woolen Robe (6243, -0.45 DPS) [crafted]; Mystic's Wrap (14369, -0.46 DPS) [world_drop]; Gray Woolen Robe (2585, -1.06 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 4.4 | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.19 DPS) [world_drop]; Mystic's Bracelets (14366, -0.38 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Gnoll Casting Gloves (892, -0.13 DPS) [world]; Blight Gloves (279877, -0.24 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.0 | yes | Keller's Girdle (2911, -0.13 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.36 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.81 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.9 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.03 DPS) [dungeon]; Darkweave Breeches (12987, -1.26 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 | yes | Pristine Boots (253889, -0.37 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.64 DPS) [world]; Red Woolen Boots (4313, -0.77 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.5 | yes | Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; Loop of Sacrifice (281673, -0.36 DPS) [quest]; Sludge-Stained Band (286535, -0.45 DPS) [world] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Lavishly Jeweled Ring (1156, -0.85 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.4 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.19 DPS) [world]; Lesser Staff of the Spire (1300, -0.38 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 175.1 | yes | Skycaller (12984, -1.00 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.41 DPS) [dungeon]; Deepblaze (279896, -4.17 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18850 Insignia of the Horde; 20434 Lorekeeper's Staff; 20440 Protector's Sword

### Band 30 (gnome, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 63.5. Weights run: 0.7s. Verify run: 0.6s. 216 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.313 ± 0.311), crit=1.967 ± 0.110, hit=2.360 ± 0.031, spell_haste=not significant (0.830 ± 0.313), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 | yes | Crystal Starfire Medallion (5003, -1.15 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.15 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.61 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.8 | yes | Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.49 DPS) [crafted]; Death Speaker Mantle (6685, -0.62 DPS, sim-verified) [dungeon] |
| back | Hillman's Cloak (3719) | Leatherworking [crafted] | 5.0 | yes | Repairman's Cape (9605, -0.02 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Prelacy Cape (7004, -0.15 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.1 | yes | Death Speaker Robes (6682, -0.40 DPS) [dungeon]; Pristine Gown (253961, -0.58 DPS) [crafted]; Tree Bark Jacket (1486, -0.74 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -1.07 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.07 DPS) [quest]; Glowing Magical Bracelets (13106, -1.57 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 7.4 | yes | Shilly Mitts (9609, -0.07 DPS) [quest]; Gnoll Casting Gloves (892, -0.22 DPS) [world]; Serpent Gloves (5970, -0.46 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 11.9 | yes | Belt of Arugal (6392, -0.42 DPS, sim-verified) [dungeon]; Invoker's Cord (215366, -0.51 DPS) [crafted]; Crimson Silk Belt (7055, -0.56 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.42 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.2 | yes | Acidic Walkers (9454, -0.25 DPS) [dungeon]; Nimbus Boots (6998, -0.48 DPS) [quest]; Spidersilk Boots (4320, -1.84 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 | yes | Minor Channeling Ring (1449, -0.21 DPS) [quest]; Lorekeeper's Ring (20431, -0.30 DPS) [rep]; Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Electrocutioner Lagnut (9447, -0.45 DPS) [dungeon]; Sludge-Stained Band (286535, -0.45 DPS) [world]; Minor Channeling Ring (1449, -1.43 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (62.8 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Talisman of Arathor (21119, -2.12 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 3.4 | yes | Gnarled Necromancer's Staff (251534, -0.05 DPS) [quest]; Channeler's Staff (4437, -0.14 DPS) [world]; Twisted Chanter's Staff (890, -0.20 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 223.5 | yes | Starfaller (13063, -0.24 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 216, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 40 (gnome, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 94.1. Weights run: 0.7s. Verify run: 0.6s. 293 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.530 ± 0.476), crit=2.880 ± 0.172, hit=3.716 ± 0.052, spell_haste=not significant (1.576 ± 0.546), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, -1.03 DPS, sim-verified) [world]; Enchanter's Cowl (4322, -1.43 DPS) [crafted]; Holy Shroud (2721, -1.47 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 | yes | Darkspear Warding Pendant (272074, -0.95 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.11 DPS) [vendor]; Necklace of Calisea (1714, -1.17 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.9 | yes | Bloodmage Mantle (7684, -0.02 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest]; Green Silken Shoulders (7057, -0.31 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.7 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.24 DPS) [crafted]; Caretaker's Cape (19532, -0.39 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.2 | yes | Dreamweave Vest (10021, -0.25 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.62 DPS) [world_drop]; Robe of Power (7054, -0.71 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Condor Bracers (15864, -0.29 DPS) [quest]; Windchaser Cuffs (14429, -0.62 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 | yes | Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Red Mageweave Gloves (10018, -0.84 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.24 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | 16.1 | yes | Deathmage Sash (10771, -0.34 DPS, sim-verified) [dungeon]; Star Belt (4329, -0.46 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.52 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.4 | yes | Abomination Skin Leggings (23173, -1.05 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.05 DPS, sim-verified) [crafted]; Gaze Dreamer Pants (6903, -1.23 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -1.68 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.17 DPS) [dungeon]; Spidersilk Boots (4320, -2.19 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 | yes | Ring of Forlorn Spirits (2043, -0.76 DPS) [quest]; Reedknot Ring (9622, -0.91 DPS) [quest]; Minor Channeling Ring (1449, -1.05 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.29 DPS) [quest]; Lorekeeper's Ring (19525, -0.29 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.16 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 43.5 | yes | Windweaver Staff (7757, -0.25 DPS, sim-verified) [dungeon]; Staff of Jordan (873, -5.54 DPS) [world_drop]; Glimmering Staff (249392, -5.54 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 270.9 | yes | Nether Force Wand (11263, -1.70 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.42 DPS) [quest]; Ragefire Wand (7513, -2.47 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 293, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes

### Band 50 (gnome, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 122.9. Weights run: 0.7s. Verify run: 0.7s. 380 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.275 ± 0.650), crit=5.153 ± 0.298, hit=5.744 ± 0.087, spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 101.4 | yes | Eye of Theradras (17715, -1.23 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -7.47 DPS) [crafted]; Dreamweave Circlet (10041, -8.88 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.8 | yes | Scorn's Icy Choker (23169, -0.33 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.55 DPS) [quest]; Darkspear Warding Pendant (272073, -0.84 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 91.6 | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Shoulders (10029, -8.59 DPS) [crafted]; Inquisitor's Shawl (19507, -8.93 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.6 | yes | Darkspear Raider's Cloak (272076, -0.50 DPS) [vendor]; Big Voodoo Cloak (8216, -0.68 DPS) [crafted]; Runecloth Cloak (13860, -1.15 DPS, sim-verified) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (119.9 DPS) | yes | Knight's Dreadweave Vest (220886, -1.27 DPS, sim-verified) [vendor]; Runecloth Robe (13858, -1.55 DPS) [crafted]; Runecloth Tunic (13857, -1.77 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (120.1 DPS) | yes | Forgotten Wraps (9433, -0.08 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.08 DPS) [quest]; Bloodband Bracers (11469, -1.50 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 87.3 | yes | Raider Handwraps (272098, -0.78 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -8.24 DPS) [vendor]; Red Mageweave Gloves (10018, -8.34 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 87.5 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.97 DPS) [dungeon]; Deathmage Sash (10771, -8.05 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 99.4 | yes | Red Mageweave Pants (10009, -0.67 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -9.64 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -9.82 DPS) [crafted] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 76.9 | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -7.14 DPS) [crafted]; Southsea Mojo Boots (20641, -7.20 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 | yes | Band of the Unicorn (7553, -5.83 DPS) [world_drop]; Lorekeeper's Ring (19523, -5.96 DPS) [rep]; Lorekeeper's Ring (19524, -6.36 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.6 | yes | Band of the Unicorn (7553, -0.61 DPS, sim-verified) [world_drop]; Lorekeeper's Ring (19523, -0.74 DPS) [rep]; Lorekeeper's Ring (19524, -1.13 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (119.8 DPS) | yes | Guardian Talisman (1490, -1.23 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -6.78 DPS) [quest]; Tidal Charm (1404, -6.78 DPS) [vendor] |
| trinket2 | Uther's Strength (11302) | World drop [world_drop] | sim-verified (118.0 DPS) | yes | Thunderbrew's Boot Flask (744, -0.79 DPS) [quest]; Tidal Charm (1404, -0.79 DPS) [vendor]; Guardian Talisman (1490, -0.79 DPS) [quest] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (118.0 DPS) | yes | Illusionary Rod (7713, -1.00 DPS) [dungeon]; Inventor's Focal Sword (17719, -2.01 DPS) [dungeon]; Shortsword of Vengeance (754, -3.07 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.70 DPS) [crafted] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Uther's Strength; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 60 (gnome, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 229.2. Weights run: 0.7s. Verify run: 0.7s. 777 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.124 ± 1.054), crit=8.680 ± 0.479, hit=10.097 ± 0.133, spell_haste=not significant (1.136 ± 1.257), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (222.0 DPS) | yes | Field Marshal's Coronet (16441, -3.40 DPS) [vendor]; Field Marshal's Coronet (231604, -3.40 DPS) [vendor]; Bloodvine Goggles (19999, -4.38 DPS, sim-verified) [crafted] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (217.7 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -4.12 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 164.4 | yes | Rugged Mantle of the Timbermaw (227808, -2.62 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -3.03 DPS) [crafted]; Lieutenant Commander's Silk Mantle (23319, -3.31 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 121.5 | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Arcanoweave Cloak (272411, -0.44 DPS) [vendor]; Earthweave Cloak (21187, -2.57 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 308.4 | yes | Bloodvine Vest (19682, -6.14 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -14.91 DPS) [vendor]; Robe of the Archmage (14152, -18.16 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 158.5 | yes | Fireleaf Wristwraps (240044, -1.37 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -3.81 DPS) [quest]; Dryad's Wrist Bindings (19595, -16.93 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 286.8 | yes | Gloves of Spell Mastery (14146, -6.24 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -14.54 DPS) [vendor]; Sorcerer's Gloves (22066, -21.49 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 167.2 | yes | Fireleaf Belt (240053, +0.00 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -2.97 DPS) [crafted]; Highlander's Cloth Girdle (20047, -3.87 DPS) [rep] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (222.6 DPS) | yes | Marshal's Silk Leggings (16442, -3.73 DPS) [vendor]; Marshal's Silk Leggings (231605, -3.73 DPS) [vendor]; Sentinel's Silk Leggings (237815, -5.05 DPS, sim-verified) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 166.2 | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -5.31 DPS) [vendor]; Marshal's Silk Footwraps (231606, -5.31 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (220.7 DPS) | yes | Wrath of Cenarius (21190, -3.06 DPS, sim-verified) [quest]; Ritssyn's Ring of Chaos (21836, -9.49 DPS) [world_drop]; Mindtear Band (20632, -9.77 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 222.5 | yes | Ritssyn's Ring of Chaos (21836, +0.00 DPS, sim-verified) [world_drop]; Mindtear Band (20632, -9.77 DPS) [world]; Signet Ring of the Bronze Dragonflight (234032, -11.54 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (213.7 DPS) | yes | Uther's Strength (11302, -1.00 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -1.75 DPS) [quest]; Frozen Heart of the Mountain (249469, -3.17 DPS, sim-verified) [crafted] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (213.7 DPS) | yes | Uther's Strength (11302, -2.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -2.03 DPS, sim-verified) [crafted]; Thunderbrew's Boot Flask (744, -2.75 DPS) [quest] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (214.8 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -3.80 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -15.15 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 620.5 | yes | Wand of Biting Cold (19108, -4.82 DPS, sim-verified) [quest]; Stormrager (16997, -14.74 DPS) [quest]; Brilliant Wand (249385, -16.70 DPS) [crafted] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Chromatic Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 777, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

## Horde

### Band 20 (troll, 000000000000000000-00000000000000000-2531000000000000000)

Set DPS (verified): 34.4. Weights run: 0.6s. Verify run: 0.7s. 121 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=0.741 ± 0.174, crit=0.951 ± 0.045, hit=1.870 ± 0.020, spell_haste=not significant (-0.108 ± 0.181), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 | yes | Flying Tiger Goggles (4368, -0.78 DPS) [crafted]; Lucky Fishing Hat (19972, -0.78 DPS) [quest]; Shadow Goggles (4373, -1.11 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 11.7 | yes | Reinforced Woolen Shoulders (4315, -0.51 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.99 DPS) [crafted]; Slime-encrusted Pads (6461, -1.51 DPS) [dungeon] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.2 | yes | Heavy Woolen Cloak (4311, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.16 DPS) [dungeon]; Black Whelp Cloak (7283, -0.16 DPS) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 8.7 | yes | Green Woolen Robe (6243, -0.45 DPS) [crafted]; Mystic's Wrap (14369, -0.46 DPS) [world_drop]; Gray Woolen Robe (2585, -0.83 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (34.0 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.10 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.34 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (34.0 DPS) | yes | Gnoll Casting Gloves (892, -0.03 DPS) [world]; Blight Gloves (279877, -0.13 DPS) [quest]; Serpent Gloves (5970, -0.38 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.0 | yes | Keller's Girdle (2911, -0.13 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.36 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.47 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 14.9 | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -1.03 DPS) [dungeon]; Darkweave Breeches (12987, -1.26 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.0 | yes | Pristine Boots (253889, -0.07 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.64 DPS) [world]; Red Woolen Boots (4313, -0.77 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Loop of Sacrifice (281673, -0.17 DPS) [quest]; Sludge-Stained Band (286535, -0.26 DPS) [world]; Volcanic Rock Ring (12053, -0.36 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 4.4 | yes | Loop of Sacrifice (281673, +0.00 DPS, sim-verified) [quest]; Sludge-Stained Band (286535, -0.19 DPS) [world]; Volcanic Rock Ring (12053, -0.29 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 7.4 | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.19 DPS) [world]; Lesser Staff of the Spire (1300, -0.38 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 175.1 | yes | Skycaller (12984, -1.10 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.41 DPS) [dungeon]; Deepblaze (279896, -4.17 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 121, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger; 254779 A'sharahm, the Roiling Tempest; 263007 Skyseer's Vest; 263432 Skulker's Shiv

### Band 30 (troll, 000000000000000000-00000000000000000-2535111300000000000)

Set DPS (verified): 57.7. Weights run: 0.7s. Verify run: 0.7s. 211 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.313 ± 0.311), crit=1.967 ± 0.110, hit=2.360 ± 0.031, spell_haste=not significant (0.830 ± 0.313), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | 11.0 | yes | Enchanter's Cowl (4322, +0.00 DPS, sim-verified) [crafted]; Silk Headband (7050, -0.30 DPS) [crafted]; Embalmed Shroud (7691, -0.45 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 8.9 | yes | Crystal Starfire Medallion (5003, -1.15 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.15 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.31 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 11.8 | yes | Death Speaker Mantle (6685, -0.32 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.45 DPS) [quest]; Invoker's Mantle (215365, -0.49 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.15 DPS) [crafted]; Battle Healer's Cloak (19529, -0.15 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 13.1 | yes | Death Speaker Robes (6682, -0.40 DPS) [dungeon]; Pristine Gown (253961, -0.58 DPS) [crafted]; Tree Bark Jacket (1486, -0.82 DPS, sim-verified) [dungeon] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Nightsky Wristbands (6407, -1.07 DPS) [world_drop]; Tabitha's Cuffs (251486, -1.07 DPS) [quest]; Glowing Magical Bracelets (13106, -1.37 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 7.6 | yes | Gnoll Casting Gloves (892, -0.24 DPS) [world]; Truefaith Gloves (7049, -0.24 DPS) [crafted]; Serpent Gloves (5970, -0.40 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 11.9 | yes | Warsong Sash (16975, -0.17 DPS, sim-verified) [quest]; Belt of Arugal (6392, -0.30 DPS) [dungeon]; Invoker's Cord (215366, -0.51 DPS) [crafted] |
| legs | Gaze Dreamer Pants (6903) | Blackfathom Deeps: Twilight Lord Kelris [dungeon] | 12.0 | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Pristine Leggings (253987, -0.42 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.62 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 9.2 | yes | Acidic Walkers (9454, -0.25 DPS) [dungeon]; Boots of the Enchanter (4325, -0.63 DPS) [crafted]; Spidersilk Boots (4320, -1.53 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 | yes | Advisor's Ring (20426, -0.30 DPS) [rep]; Electrocutioner Lagnut (9447, -0.60 DPS) [dungeon]; Sludge-Stained Band (286535, -0.60 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 | yes | Sludge-Stained Band (286535, -0.45 DPS) [world]; Black Widow Band (6199, -0.57 DPS) [world]; Electrocutioner Lagnut (9447, -1.14 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (57.2 DPS) | yes | Minor Recombobulator (4381, +0.00 DPS) [crafted]; Gnomish Universal Remote (7506, +0.00 DPS) [crafted]; Defiler's Talisman (21120, -2.19 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 3.4 | yes | Gnarled Necromancer's Staff (251534, -0.05 DPS) [quest]; Twisted Chanter's Staff (890, -0.08 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.14 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 223.5 | yes | Starfaller (13063, +0.00 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.91 DPS) [crafted]; Gravestone Scepter (7001, -4.66 DPS) [quest] |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Gaze Dreamer Pants; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 211, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 5000 Coral Band; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia; 18859 Insignia of the Alliance; 19545 Scout's Blade

### Band 40 (troll, 000000000000000000-00000000000000000-2535111300000301051)

Set DPS (verified): 86.0. Weights run: 0.7s. Verify run: 0.7s. 288 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.530 ± 0.476), crit=2.880 ± 0.172, hit=3.716 ± 0.052, spell_haste=not significant (1.576 ± 0.546), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Enchanter's Cowl (4322, -1.43 DPS) [crafted]; Holy Shroud (2721, -1.47 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 10.2 | yes | Necklace of Calisea (1714, -0.55 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.95 DPS) [vendor]; Darkspear Warding Pendant (272075, -1.11 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 13.9 | yes | Bloodmage Mantle (7684, -0.02 DPS) [dungeon]; Berylline Pads (4197, -0.23 DPS) [quest]; Green Silken Shoulders (7057, -0.91 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 8.7 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Icy Cloak (4327, -0.24 DPS) [crafted]; Battle Healer's Cloak (19528, -0.39 DPS) [rep] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 25.2 | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Elemental Raiment (9434, -0.62 DPS) [world_drop]; Robe of Power (7054, -0.71 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Condor Bracers (15864, -0.29 DPS) [quest]; Radiant Silver Bracers (4545, -0.53 DPS, sim-verified) [quest]; Windchaser Cuffs (14429, -0.62 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 20.1 | yes | Black Mageweave Gloves (10003, -0.75 DPS) [crafted]; Red Mageweave Gloves (10018, -0.90 DPS, sim-verified) [crafted]; Gilded Handwraps (254021, -1.24 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | sim-verified (86.2 DPS) | yes | Star Belt (4329, -0.29 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.35 DPS) [rep]; Defiler's Cloth Girdle (20166, -1.34 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 20.4 | yes | Crimson Silk Pantaloons (7062, -1.05 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.05 DPS) [dungeon]; Gaze Dreamer Pants (6903, -1.23 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Gilded Slippers (254001, -0.92 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -2.17 DPS) [dungeon]; Spidersilk Boots (4320, -2.19 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 13.2 | yes | Reedknot Ring (9622, -0.91 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.06 DPS) [vendor]; Ogremind Ring (1993, -1.39 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.29 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.44 DPS) [vendor]; Reedknot Ring (9622, -1.40 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Windweaver Staff (7757) | Scarlet Monastery: Houndmaster Loksey [dungeon] | sim-verified (85.9 DPS) | yes | Staff of Jordan (873, -0.31 DPS) [world_drop]; Glimmering Staff (249392, -0.31 DPS) [crafted]; Illusionary Rod (7713, -1.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 270.9 | yes | Nether Force Wand (11263, -0.90 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.42 DPS) [quest]; Ragefire Wand (7513, -2.47 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Windweaver Staff; ranged: Jaina's Firestarter

No-known-source sample (15 of 288, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9718 Reforged Blade of Heroes; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape

### Band 50 (troll, 000000000000000000-23500000000000000-2535111300000301051)

Set DPS (verified): 110.8. Weights run: 0.7s. Verify run: 0.7s. 375 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.275 ± 0.650), crit=5.153 ± 0.298, hit=5.744 ± 0.087, spell_haste=not significant (-1.194 ± 0.805), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 101.4 | yes | Eye of Theradras (17715, -0.30 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -7.47 DPS) [crafted]; Dreamweave Circlet (10041, -8.88 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | 17.8 | yes | Scorn's Icy Choker (23169, -0.13 DPS, sim-verified) [dungeon]; Mindburst Medallion (11196, -0.55 DPS) [quest]; Darkspear Warding Pendant (272073, -0.84 DPS) [vendor] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (108.1 DPS) | yes | Blood Guard's Dreadweave Mantle (220905, -1.21 DPS, sim-verified) [vendor]; Red Mageweave Shoulders (10029, -1.29 DPS) [crafted]; Inquisitor's Shawl (19507, -1.62 DPS) [dungeon] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 23.5 | yes | Spritecaster Cape (11623, +0.00 DPS, sim-verified) [dungeon]; Runecloth Cloak (13860, -0.56 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.74 DPS) [vendor] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (108.4 DPS) | yes | Stone Guard's Dreadweave Vest (220904, -1.54 DPS, sim-verified) [vendor]; Runecloth Robe (13858, -1.55 DPS) [crafted]; Runecloth Tunic (13857, -1.77 DPS) [crafted] |
| wrist | Nethergeld Cuffs (254061) | Tailoring [crafted] | sim-verified (108.3 DPS) | yes | Forgotten Wraps (9433, -0.08 DPS) [world_drop]; Shizzle's Nozzle Wiper (11917, -0.08 DPS) [quest]; Bloodband Bracers (11469, -1.42 DPS, sim-verified) [quest] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 87.3 | yes | Raider Handwraps (272098, -0.15 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -8.24 DPS) [vendor]; Red Mageweave Gloves (10018, -8.34 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 87.5 | yes | Dawnspire Cord (12466, +0.00 DPS, sim-verified) [dungeon]; Satyrmane Sash (17755, -7.97 DPS) [dungeon]; Deathmage Sash (10771, -8.05 DPS) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 99.4 | yes | Red Mageweave Pants (10009, -0.56 DPS, sim-verified) [crafted]; Kilt of the Atal'ai Prophet (10807, -9.64 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -9.82 DPS) [crafted] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 76.9 | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -7.14 DPS) [crafted]; Southsea Mojo Boots (20641, -7.20 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 57.4 | yes | Band of the Unicorn (7553, -5.83 DPS) [world_drop]; Advisor's Ring (19519, -5.96 DPS) [rep]; Advisor's Ring (19520, -6.36 DPS) [rep] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.6 | yes | Band of the Unicorn (7553, -0.68 DPS, sim-verified) [world_drop]; Advisor's Ring (19519, -0.74 DPS) [rep]; Advisor's Ring (19520, -1.13 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (106.0 DPS) | yes | Ankh of Life (1713, -0.39 DPS, sim-verified) [world_drop]; Uther's Strength (11302, -6.00 DPS) [world_drop]; Tidal Charm (1404, -6.78 DPS) [vendor] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (106.0 DPS) | yes | Ankh of Life (1713, -0.04 DPS, sim-verified) [world_drop]; Uther's Strength (11302, -4.49 DPS) [world_drop]; Tidal Charm (1404, -5.28 DPS) [vendor] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (106.0 DPS) | yes | Illusionary Rod (7713, -1.00 DPS) [dungeon]; Inventor's Focal Sword (17719, -2.01 DPS) [dungeon]; Shortsword of Vengeance (754, -2.38 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 400.2 | yes | Noxious Shooter (17745, +0.00 DPS, sim-verified) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Lesser Eternal Wand (249232, -5.70 DPS) [crafted] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Nethergeld Cuffs; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Frozen Heart of the Mountain; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 375, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

### Band 60 (troll, 000000000000000000-23552300000000000-2535111300000301051)

Set DPS (verified): 205.8. Weights run: 0.7s. Verify run: 0.7s. 774 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.124 ± 1.054), crit=8.680 ± 0.479, hit=10.097 ± 0.133, spell_haste=not significant (1.136 ± 1.257), spell_penetration=not significant (0.000 ± 0.000), frost_power=1.000 ± 0.004

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (200.6 DPS) | yes | Warlord's Silk Cowl (16533, -3.40 DPS) [vendor]; Warlord's Silk Cowl (231601, -3.40 DPS) [vendor]; Bloodvine Goggles (19999, -9.03 DPS, sim-verified) [crafted] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | sim-verified (190.8 DPS) | yes | Orb of the Darkmoon (19426, +0.00 DPS) [quest]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -3.35 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 164.4 | yes | Rugged Mantle of the Timbermaw (227808, -2.28 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -3.03 DPS) [crafted]; Champion's Silk Mantle (23264, -3.31 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 121.5 | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Arcanoweave Cloak (272411, -0.44 DPS) [vendor]; Earthweave Cloak (21187, -2.57 DPS) [quest] |
| chest | Fireleaf Robe (240059) | Leonid Barthalomew the Revered [vendor] | 308.4 | yes | Bloodvine Vest (19682, -4.61 DPS, sim-verified) [crafted]; Fireleaf Garb (240051, -14.91 DPS) [vendor]; Robe of the Archmage (14152, -18.16 DPS) [crafted] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 158.5 | yes | Fireleaf Wristwraps (240044, -0.56 DPS, sim-verified) [vendor]; Rockfury Bracers (21186, -3.81 DPS) [quest]; Dryad's Wrist Bindings (19595, -16.93 DPS) [rep] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 286.8 | yes | Gloves of Spell Mastery (14146, -4.81 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -14.54 DPS) [vendor]; Sorcerer's Gloves (22066, -21.49 DPS) [quest] |
| waist | Fireleaf Waistguard (240045) | Leonid Barthalomew the Revered [vendor] | 167.2 | yes | Fireleaf Belt (240053, +0.00 DPS, sim-verified) [vendor]; Belt of the Archmage (18405, -2.97 DPS) [crafted]; Defiler's Cloth Girdle (20163, -3.87 DPS) [rep] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (198.1 DPS) | yes | General's Silk Trousers (16534, -3.73 DPS) [vendor]; General's Silk Trousers (231595, -3.73 DPS) [vendor]; Sentinel's Silk Leggings (237815, -6.58 DPS, sim-verified) [vendor] |
| feet | Fireleaf Boots (240050) (or Fireleaf Sandals (240058)) | Leonid Barthalomew the Revered [vendor] | 166.2 | yes | Fireleaf Sandals (240058, +0.00 DPS, sim-verified) [vendor]; General's Silk Boots (16539, -5.31 DPS) [vendor]; General's Silk Boots (231597, -5.31 DPS) [vendor] |
| finger1 | Wrath of Cenarius (21190) | Champion's Battlegear [quest] | sim-verified (190.2 DPS) | yes | Don Julio's Band (19325, +0.00 DPS, sim-verified) [rep]; Mindtear Band (20632, +0.00 DPS) [world]; Band of Earthen Might (21182, +0.00 DPS) [quest] |
| finger2 | Ritssyn's Ring of Chaos (21836) | World drop [world_drop] | sim-verified (193.6 DPS) | yes | Don Julio's Band (19325, +0.00 DPS) [rep]; Mindtear Band (20632, -0.28 DPS) [world]; Band of Earthen Might (21182, -2.03 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (190.0 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.00 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (193.9 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -2.00 DPS) [world_drop]; Rune of the Guard Captain (19120, -2.39 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (190.2 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -4.12 DPS, sim-verified) [world_drop]; Kindling Stave (11750, -15.15 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | 620.5 | yes | Wand of Biting Cold (19108, -4.66 DPS, sim-verified) [quest]; Stormrager (16997, -14.74 DPS) [quest]; Brilliant Wand (249385, -16.70 DPS) [crafted] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Chromatic Cloak; chest: Fireleaf Robe; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Waistguard; legs: Fireleaf Leggings; feet: Fireleaf Boots; finger1: Wrath of Cenarius; finger2: Ritssyn's Ring of Chaos; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 774, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4116 Olmann Sewar; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 9380 Jang'thraze the Protector; 9653 Speedy Racer Goggles

