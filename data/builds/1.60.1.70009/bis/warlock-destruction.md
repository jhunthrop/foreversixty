# Leveling BiS: Destruction

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 43.2. Weights run: 1.1s. Verify run: 1.0s. 129 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.152, intellect=not significant (0.589 ± 0.173), crit=0.348 ± 0.028, hit=0.848 ± 0.013, spell_haste=0.663 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.665 ± 0.152, fire_power=0.340 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -3.28 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.3 spell_power points (1.36 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.05 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.83 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.54 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.9 spell_power points (1.05 DPS) | yes | Green Woolen Robe (6243, -0.42 DPS) [crafted]; Mystic's Wrap (14369, -0.50 DPS) [world_drop]; Gray Woolen Robe (2585, -1.45 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.5 spell_power points (0.47 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Bright Bracers (3647, -0.16 DPS) [world_drop]; Mystic's Bracelets (14366, -0.31 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Pristine Gloves (253913, -0.16 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS, sim-verified) [world]; Blight Gloves (279877, -0.38 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.4 spell_power points (0.84 DPS) | yes | Keller's Girdle (2911, -0.22 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Novice Arcanist's Sash (253885, -1.09 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (43.2 DPS) | yes | Silk-threaded Trousers (1929, -0.33 DPS) [dungeon]; Colorful Kilt (10048, -0.60 DPS) [crafted]; Abomination Skin Leggings (23173, -0.85 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.4 spell_power points (1.23 DPS) | yes | Pristine Boots (253889, -0.60 DPS) [crafted]; Feather Padded Treads (285345, -0.62 DPS, sim-verified) [world]; Red Woolen Boots (4313, -0.71 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.2 spell_power points (0.81 DPS) | yes | Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Sludge-Stained Band (286535, -0.42 DPS) [world]; Loop of Sacrifice (281673, -0.43 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.66 DPS) | yes | Sludge-Stained Band (286535, -0.26 DPS) [world]; Loop of Sacrifice (281673, -0.27 DPS) [quest]; Lavishly Jeweled Ring (1156, -1.31 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 5.9 spell_power points (0.78 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.16 DPS) [world]; Lesser Staff of the Spire (1300, -0.31 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 172.2 spell_power points (22.71 DPS) | yes | Skycaller (12984, -2.36 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 129, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 10049 Diabolist's Blade; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade

### Band 30 (gnome, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 60.0. Weights run: 1.1s. Verify run: 1.0s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.118, intellect=1.226 ± 0.138, crit=0.549 ± 0.036, hit=0.668 ± 0.011, spell_haste=not significant (0.090 ± 0.086), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.832 ± 0.118, fire_power=0.168 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 18.3 spell_power points (4.73 DPS) | yes | Shadow Hood (4323, -1.24 DPS) [crafted]; Resilient Cap (14401, -1.24 DPS) [world_drop]; Nightsky Cowl (4039, -1.73 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.4 spell_power points (3.72 DPS) | yes | Darkspear Warding Pendant (272075, -1.93 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.45 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.45 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 20.0 spell_power points (5.19 DPS) | yes | Fairywing Mantle (9536, -0.78 DPS) [quest]; Death Speaker Mantle (6685, -0.80 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.04 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.8 spell_power points (2.54 DPS) | yes | Darkspear Raider's Cloak (272078, -0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.49 DPS) [quest]; Resilient Cape (14400, -0.63 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 24.9 spell_power points (6.46 DPS) | yes | Death Speaker Robes (6682, -0.74 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -1.70 DPS) [dungeon]; Pristine Gown (253961, -2.42 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (56.4 DPS) | yes | Nightsky Wristbands (6407, -0.43 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.43 DPS) [quest]; Glowing Magical Bracelets (13106, -2.55 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 17.5 spell_power points (4.53 DPS) | yes | Hotshot Pilot's Gloves (9491, -1.05 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -2.28 DPS) [crafted]; Blight Gloves (279877, -2.31 DPS) [quest] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 14.7 spell_power points (3.80 DPS) | yes | Invoker's Cord (215366, -0.40 DPS) [crafted]; Belt of Arugal (6392, -0.52 DPS) [dungeon]; Crimson Silk Belt (7055, -1.57 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (55.4 DPS) | yes | Necromancer Leggings (2277, -0.54 DPS) [world_drop]; Filigreed Pristine Leggings (253937, -0.58 DPS) [crafted]; Abomination Skin Leggings (23173, -1.52 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.6 spell_power points (4.03 DPS) | yes | Spidersilk Boots (4320, -0.95 DPS) [crafted]; Frothing Slippers (254003, -1.81 DPS) [crafted]; Acidic Walkers (9454, -2.36 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.6 spell_power points (2.22 DPS) | yes | Snake Hoop (6750, -0.20 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.41 DPS) [rep] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | sim-verified (55.3 DPS) | yes | Lavishly Jeweled Ring (1156, -0.02 DPS) [dungeon]; Lorekeeper's Ring (19525, -0.12 DPS) [rep]; Snake Hoop (6750, -1.45 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (53.9 DPS) | yes | Talisman of Arathor (21119, -0.38 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 13.5 spell_power points (3.49 DPS) | yes | Twisted Chanter's Staff (890, -0.01 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.32 DPS) [quest]; Channeler's Staff (4437, -0.95 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (54.8 DPS) | yes | Starfaller (13063, -0.87 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.69 DPS) [crafted]; Gravestone Scepter (7001, -4.99 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Minor Channeling Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 108.3. Weights run: 1.1s. Verify run: 1.0s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.396), intellect=not significant (1.267 ± 0.519), crit=1.986 ± 0.117, hit=2.861 ± 0.043, spell_haste=3.882 ± 0.364, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.107 ± 0.396), fire_power=1.100 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (103.8 DPS) | yes | Thinking Cap (2624, -0.25 DPS) [world_drop]; Miner's Hat of the Deep (9429, -0.25 DPS) [dungeon]; Corpseshroud (10574, -3.37 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.6 spell_power points (1.73 DPS) | yes | Necklace of Calisea (1714, -0.68 DPS) [world_drop]; Triune Amulet (7722, -0.68 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.55 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 23.5 spell_power points (2.78 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Death Speaker Mantle (6685, -0.42 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (102.0 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -0.11 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -1.57 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.6 spell_power points (3.51 DPS) | yes | Robe of Power (7054, -0.05 DPS) [crafted]; Green Silk Armor (7065, -0.49 DPS) [crafted]; Dreamweave Vest (10021, -0.67 DPS, sim-verified) [crafted] |
| wrist | Windchaser Cuffs (14429) | World drop [world_drop] | 11.4 spell_power points (1.35 DPS) | yes | Aurora Bracers (4043, +0.00 DPS, sim-verified) [world_drop]; Mistscape Bracers (4045, -0.15 DPS) [world_drop]; Enchanted Stonecloth Bracers (4979, -0.15 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (102.4 DPS) | yes | Stormcloth Gloves (10011, -0.46 DPS) [crafted]; Town Clerk's Mittens (270029, -0.61 DPS) [quest]; Red Mageweave Gloves (10018, -1.95 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20098) | The League of Arathor [rep] | sim-verified (102.5 DPS) | yes | Gilded Cord (254037, -0.11 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.46 DPS) [quest]; Deathmage Sash (10771, -2.01 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.2 spell_power points (3.46 DPS) | yes | Abomination Skin Leggings (23173, -1.19 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.37 DPS, sim-verified) [crafted]; Stoneweaver Leggings (9407, -1.43 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.84 DPS) | yes | Acidic Walkers (9454, -1.05 DPS) [dungeon]; Spidersilk Boots (4320, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.79 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.6 spell_power points (2.08 DPS) | yes | Ogremind Ring (1993, -1.03 DPS) [world_drop]; Voodoo Band (1996, -1.03 DPS) [world_drop]; Mindbender Loop (5009, -1.03 DPS) [world_drop] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.07 DPS) | yes | Voodoo Band (1996, -0.02 DPS) [world_drop]; Mindbender Loop (5009, -0.02 DPS) [world_drop]; Ogremind Ring (1993, -2.01 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (100.5 DPS) | yes | Gut Ripper (2164, -0.75 DPS, sim-verified) [world_drop]; Illusionary Rod (7713, -0.84 DPS) [dungeon]; Staff of Noh'Orahil (15105, -1.14 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 340.1 spell_power points (40.28 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.69 DPS) [dungeon]; Twisted Nether Wand (249144, -5.57 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Windchaser Cuffs; hands: Dreamweave Gloves; waist: Highlander's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 140.2. Weights run: 1.0s. Verify run: 0.9s. 396 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.605), intellect=not significant (-0.468 ± 0.718), crit=2.617 ± 0.140, hit=4.179 ± 0.053, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 50.6 spell_power points (7.85 DPS) | yes | Spellpower Goggles Xtreme Plus (15999, -3.66 DPS) [crafted]; Eye of Theradras (17715, -3.91 DPS, sim-verified) [dungeon]; Dreamweave Circlet (10041, -4.59 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 spell_power points (1.08 DPS) | yes | Mindburst Medallion (11196, -0.80 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Dreadweave Mantle (220887) | Captain Dirgehammer [vendor] | 44.6 spell_power points (6.92 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Black Mageweave Shoulders (10027, -5.37 DPS) [crafted]; Bloodmage Mantle (7684, -5.52 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 spell_power points (2.17 DPS) | yes | Runecloth Cloak (13860, -0.77 DPS) [crafted]; Icy Cloak (4327, -1.08 DPS) [crafted]; Nightfall Drape (12465, -2.89 DPS, sim-verified) [dungeon] |
| chest | Knight's Dreadweave Vest (220886) | Captain Dirgehammer [vendor] | 48.6 spell_power points (7.54 DPS) | yes | Robe of the Magi (1716, +0.00 DPS, sim-verified) [world_drop]; Elemental Raiment (9434, -4.28 DPS) [world_drop]; Acumen Robes (17775, -4.59 DPS) [quest] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.39 DPS) | yes | Condor Bracers (15864, -0.31 DPS) [quest]; Nethergeld Cuffs (254061, -0.31 DPS) [crafted]; Spidertank Oilrag (9448, -0.51 DPS, sim-verified) [dungeon] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 spell_power points (2.79 DPS) | yes | Brightcloth Gloves (14101, -0.77 DPS) [crafted]; Sergeant Major's Dreadweave Gloves (220890, -0.77 DPS) [vendor]; Black Mageweave Gloves (10003, -1.67 DPS, sim-verified) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 45.6 spell_power points (7.07 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Highlander's Cloth Girdle (20098, -4.90 DPS) [rep]; Ghostweave Cord (254073, -4.90 DPS) [crafted] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 48.6 spell_power points (7.54 DPS) | yes | Wizardweave Leggings (14132, -0.37 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -5.37 DPS) [crafted]; Gaze Dreamer Pants (6903, -5.68 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 49.8 spell_power points (7.71 DPS) | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Black Mageweave Boots (10026, -6.01 DPS) [crafted]; Gilded Sandals (254107, -6.01 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.8 spell_power points (6.48 DPS) | yes | Lorekeeper's Ring (19523, -4.62 DPS) [rep]; Philanthropist's Ring (281635, -4.93 DPS) [quest]; Lorekeeper's Ring (19524, -5.08 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.01 DPS) | yes | Philanthropist's Ring (281635, -0.46 DPS) [quest]; Lorekeeper's Ring (19523, -0.54 DPS, sim-verified) [rep]; Lorekeeper's Ring (19524, -0.62 DPS) [rep] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (140.2 DPS) | yes | Uther's Strength (11302, -0.13 DPS, sim-verified) [world_drop] |
| trinket2 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (140.2 DPS) | yes | Uther's Strength (11302, -2.07 DPS, sim-verified) [world_drop] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (140.2 DPS) | yes | Soul Harvester (20536, +0.00 DPS) [quest]; Illusionary Rod (7713, -0.80 DPS) [dungeon]; Shortsword of Vengeance (754, -2.37 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 338.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.73 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -6.16 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; shoulder: Knight-Lieutenant's Dreadweave Mantle; back: Spritecaster Cape; chest: Knight's Dreadweave Vest; wrist: Arcane Runed Bracers; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Frozen Heart of the Mountain; trinket2: Abyss Shard; ranged: Pyric Caduceus

No-known-source sample (15 of 396, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 336.9. Weights run: 1.0s. Verify run: 0.9s. 944 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.338), intellect=not significant (0.993 ± 0.330), crit=1.580 ± 0.087, hit=2.312 ± 0.031, spell_haste=not significant (0.018 ± 0.286), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.566 ± 0.338), fire_power=0.427 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 143.3 spell_power points (53.14 DPS) | yes | Bloodvine Goggles (19999, -21.75 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -29.05 DPS) [quest]; Deathmist Mask (22074, -29.79 DPS) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (329.5 DPS) | yes | Beads of Ogre Mojo (22149, -1.11 DPS) [quest]; Beads of Ogre Might (22150, -1.77 DPS) [quest]; Blazefury Medallion (17111, -6.04 DPS, sim-verified) [world] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 98.1 spell_power points (36.40 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -7.86 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -17.09 DPS) [vendor]; Mantle of the Timbermaw (19050, -17.11 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 47.1 spell_power points (17.46 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Earthweave Cloak (21187, -8.88 DPS) [quest]; Howler's Furs (272414, -8.88 DPS) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 126.1 spell_power points (46.77 DPS) | yes | Bloodvine Vest (19682, -11.99 DPS, sim-verified) [crafted]; Heretic Garb (240146, -18.53 DPS) [vendor]; Earthpower Vest (21183, -21.20 DPS) [quest] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 62.0 spell_power points (23.01 DPS) | yes | Rockfury Bracers (21186, -1.15 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -9.24 DPS) [vendor]; Dryad's Wrist Bindings (19595, -11.90 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 73.0 spell_power points (27.09 DPS) | yes | Deathmist Wraps (226911, -8.90 DPS) [quest]; Deathmist Wraps (22077, -8.90 DPS) [quest]; Gloves of Spell Mastery (14146, -9.10 DPS, sim-verified) [crafted] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 120.3 spell_power points (44.60 DPS) | yes | Heretic Waistguard (240151, -10.16 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -18.65 DPS) [vendor]; Belt of the Archmage (18405, -23.09 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 139.2 spell_power points (51.64 DPS) | yes | Sentinel's Silk Leggings (237815, -12.76 DPS, sim-verified) [vendor]; Heretic Pants (240149, -23.77 DPS) [vendor]; Bloodvine Leggings (19683, -27.13 DPS) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 73.0 spell_power points (27.08 DPS) | yes | Bloodvine Boots (19684, -6.45 DPS, sim-verified) [crafted]; Heretic Boots (240153, -9.26 DPS) [vendor]; Sergeant Major's Dreadweave Boots (220891, -12.23 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (329.5 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -1.11 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -1.48 DPS) [vendor]; Wrath of Cenarius (21190, -7.25 DPS, sim-verified) [quest] |
| finger2 | Mindtear Band (20632) | Taerar [world] | sim-verified (329.5 DPS) | yes | Ritssyn's Ring of Chaos (21836, -1.10 DPS) [world_drop]; Don Julio's Band (19325, -1.79 DPS) [rep]; Wrath of Cenarius (21190, -3.62 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (329.5 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Abyss Shard (20534, -0.74 DPS) [quest]; Uther's Strength (11302, -2.97 DPS) [world_drop] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (329.5 DPS) | yes | Abyss Shard (20534, -3.71 DPS) [quest]; Frozen Heart of the Mountain (249469, -5.51 DPS, sim-verified) [crafted]; Uther's Strength (11302, -5.93 DPS) [world_drop] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | sim-verified (329.5 DPS) | yes | Grand Marshal's Stave (18873, +0.00 DPS) [vendor]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -5.60 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (336.9 DPS) | yes | Cold Snap (19130, -7.35 DPS, sim-verified) [world]; Oblivion's Touch (18761, -12.43 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.34 DPS) [world] |

**New at 60:** head: Heretic Cowl; neck: Amulet of the Dawn; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Mindtear Band; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 944, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (troll, 00000000000000000-0000000000000000000-2351000000000000)

Set DPS (verified): 41.9. Weights run: 1.1s. Verify run: 1.0s. 125 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.152, intellect=not significant (0.589 ± 0.173), crit=0.348 ± 0.028, hit=0.848 ± 0.013, spell_haste=0.663 ± 0.150, spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.665 ± 0.152, fire_power=0.340 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.79 DPS) | yes | Shadow Goggles (4373, -3.01 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 10.3 spell_power points (1.36 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.08 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.83 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.53 DPS) | yes | Feyscale Cloak (6632, -0.13 DPS) [dungeon]; Black Whelp Cloak (7283, -0.13 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.47 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 7.9 spell_power points (1.05 DPS) | yes | Green Woolen Robe (6243, -0.42 DPS) [crafted]; Mystic's Wrap (14369, -0.50 DPS) [world_drop]; Gray Woolen Robe (2585, -1.44 DPS, sim-verified) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 3.5 spell_power points (0.47 DPS) | yes | Mindthrust Bracers (1974, +0.00 DPS, sim-verified) [dungeon]; Featherbead Bracers (15452, -0.08 DPS) [quest]; Bright Bracers (3647, -0.16 DPS) [world_drop] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.92 DPS) | yes | Pristine Gloves (253913, -0.16 DPS) [crafted]; Gnoll Casting Gloves (892, -0.27 DPS, sim-verified) [world]; Blight Gloves (279877, -0.38 DPS) [quest] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 6.4 spell_power points (0.84 DPS) | yes | Keller's Girdle (2911, -0.22 DPS) [world_drop]; Novice Ardent's Sash (253887, -0.34 DPS) [crafted]; Novice Arcanist's Sash (253885, -0.93 DPS, sim-verified) [crafted] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (41.2 DPS) | yes | Silk-threaded Trousers (1929, -0.33 DPS) [dungeon]; Colorful Kilt (10048, -0.60 DPS) [crafted]; Abomination Skin Leggings (23173, -0.82 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 9.4 spell_power points (1.23 DPS) | yes | Pristine Boots (253889, -0.60 DPS) [crafted]; Red Woolen Boots (4313, -0.71 DPS) [crafted]; Feather Padded Treads (285345, -0.71 DPS, sim-verified) [world] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.66 DPS) | yes | Loop of Sacrifice (281673, -0.27 DPS) [quest]; Volcanic Rock Ring (12053, -0.43 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -1.43 DPS, sim-verified) [dungeon] |
| finger2 | Sludge-Stained Band (286535) | Sludge Beast [world] | sim-verified (41.1 DPS) | yes | Loop of Sacrifice (281673, -0.01 DPS) [quest]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; Lavishly Jeweled Ring (1156, -0.69 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 5.9 spell_power points (0.78 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.16 DPS) [world]; Lesser Staff of the Spire (1300, -0.31 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 172.2 spell_power points (22.71 DPS) | yes | Skycaller (12984, -2.24 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.42 DPS) [dungeon]; Deepblaze (279896, -4.18 DPS) [quest] |

**New at 20:** head: Pristine Circlet; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Tabitha's Cuffs; hands: Serpent Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Sludge-Stained Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 125, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 10049 Diabolist's Blade; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18858 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209615 Insignia of the Alliance

### Band 30 (troll, 00000000000000000-0000000000000000000-2353224000000000)

Set DPS (verified): 57.7. Weights run: 1.1s. Verify run: 1.0s. 218 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.118, intellect=1.226 ± 0.138, crit=0.549 ± 0.036, hit=0.668 ± 0.011, spell_haste=not significant (0.090 ± 0.086), spell_penetration=not significant (0.000 ± 0.000), shadow_power=0.832 ± 0.118, fire_power=0.168 ± 0.000

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 18.3 spell_power points (4.73 DPS) | yes | Shadow Hood (4323, -1.24 DPS) [crafted]; Resilient Cap (14401, -1.24 DPS) [world_drop]; Nightsky Cowl (4039, -1.57 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.4 spell_power points (3.72 DPS) | yes | Darkspear Warding Pendant (272075, -1.88 DPS, sim-verified) [vendor]; Crystal Starfire Medallion (5003, -2.45 DPS) [world_drop]; Kaleidoscope Chain (13084, -2.45 DPS) [world_drop] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 20.0 spell_power points (5.19 DPS) | yes | Fairywing Mantle (9536, -0.78 DPS) [quest]; Death Speaker Mantle (6685, -0.83 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -1.04 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.8 spell_power points (2.54 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.63 DPS) [world_drop]; Soft Willow Cape (16661, -0.95 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 24.9 spell_power points (6.46 DPS) | yes | Death Speaker Robes (6682, -0.71 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -1.70 DPS) [dungeon]; Pristine Gown (253961, -2.42 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (55.3 DPS) | yes | Nightsky Wristbands (6407, -0.43 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.43 DPS) [quest]; Glowing Magical Bracelets (13106, -2.32 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 12.1 spell_power points (3.14 DPS) | yes | Truefaith Gloves (7049, -0.89 DPS) [crafted]; Blight Gloves (279877, -0.92 DPS) [quest]; Hotshot Pilot's Gloves (9491, -1.61 DPS, sim-verified) [dungeon] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 14.7 spell_power points (3.80 DPS) | yes | Invoker's Cord (215366, -0.40 DPS) [crafted]; Belt of Arugal (6392, -0.52 DPS) [dungeon]; Crimson Silk Belt (7055, -1.44 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (54.8 DPS) | yes | Necromancer Leggings (2277, -0.54 DPS) [world_drop]; Filigreed Pristine Leggings (253937, -0.58 DPS) [crafted]; Abomination Skin Leggings (23173, -1.72 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.6 spell_power points (4.03 DPS) | yes | Spidersilk Boots (4320, -0.95 DPS) [crafted]; Frothing Slippers (254003, -1.81 DPS) [crafted]; Acidic Walkers (9454, -2.33 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.6 spell_power points (2.22 DPS) | yes | Lavishly Jeweled Ring (1156, -0.32 DPS) [dungeon]; Advisor's Ring (19521, -0.41 DPS) [rep]; Azora's Will (4999, -0.63 DPS) [world_drop] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.6 spell_power points (2.22 DPS) | yes | Lavishly Jeweled Ring (1156, -0.14 DPS, sim-verified) [dungeon]; Advisor's Ring (19521, -0.41 DPS) [rep]; Azora's Will (4999, -0.63 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 13.5 spell_power points (3.49 DPS) | yes | Twisted Chanter's Staff (890, -0.17 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.32 DPS) [quest]; Channeler's Staff (4437, -0.95 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | sim-verified (53.7 DPS) | yes | Starfaller (13063, -0.67 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -3.69 DPS) [crafted]; Gravestone Scepter (7001, -4.99 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Relentless Raider's Seal; trinket2: Rune of Perfection; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 106.7. Weights run: 1.1s. Verify run: 1.0s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.396), intellect=not significant (1.267 ± 0.519), crit=1.986 ± 0.117, hit=2.861 ± 0.043, spell_haste=3.882 ± 0.364, spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (-0.107 ± 0.396), fire_power=1.100 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Augural Shroud (2620) | Nether Sorceress [world] | sim-verified (103.5 DPS) | yes | Thinking Cap (2624, -0.25 DPS) [world_drop]; Miner's Hat of the Deep (9429, -0.25 DPS) [dungeon]; Corpseshroud (10574, -2.95 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 14.6 spell_power points (1.73 DPS) | yes | Necklace of Calisea (1714, -0.68 DPS) [world_drop]; Triune Amulet (7722, -0.68 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -1.93 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 23.5 spell_power points (2.78 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.36 DPS) [dungeon]; Death Speaker Mantle (6685, -0.42 DPS) [dungeon] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (101.6 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Mantle of Lady Falther'ess (23178, -0.11 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -1.04 DPS, sim-verified) [vendor] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 29.6 spell_power points (3.51 DPS) | yes | Robe of Power (7054, -0.05 DPS) [crafted]; Green Silk Armor (7065, -0.49 DPS) [crafted]; Dreamweave Vest (10021, -0.87 DPS, sim-verified) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 14.1 spell_power points (1.67 DPS) | yes | Aurora Bracers (4043, -0.47 DPS) [world_drop]; Mistscape Bracers (4045, -0.47 DPS) [world_drop]; Windchaser Cuffs (14429, -0.91 DPS, sim-verified) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | sim-verified (101.7 DPS) | yes | Stormcloth Gloves (10011, -0.46 DPS) [crafted]; Gilded Handwraps (254021, -0.73 DPS) [crafted]; Red Mageweave Gloves (10018, -1.15 DPS, sim-verified) [crafted] |
| waist | Defiler's Cloth Girdle (20166) | The Defilers [rep] | sim-verified (102.0 DPS) | yes | Gilded Cord (254037, -0.11 DPS) [crafted]; Razzeric's Customized Seatbelt (6726, -0.46 DPS) [quest]; Deathmage Sash (10771, -1.37 DPS, sim-verified) [dungeon] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 29.2 spell_power points (3.46 DPS) | yes | Abomination Skin Leggings (23173, -1.19 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.43 DPS) [dungeon]; Crimson Silk Pantaloons (7062, -1.46 DPS, sim-verified) [crafted] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (2.84 DPS) | yes | Acidic Walkers (9454, -1.05 DPS) [dungeon]; Spidersilk Boots (4320, -1.41 DPS) [crafted]; Gilded Slippers (254001, -2.09 DPS, sim-verified) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 17.6 spell_power points (2.08 DPS) | yes | Ogremind Ring (1993, -1.03 DPS) [world_drop]; Voodoo Band (1996, -1.03 DPS) [world_drop]; Mindbender Loop (5009, -1.03 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.07 DPS) | yes | Voodoo Band (1996, -0.02 DPS) [world_drop]; Mindbender Loop (5009, -0.02 DPS) [world_drop]; Ogremind Ring (1993, -1.95 DPS, sim-verified) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (100.6 DPS) | yes | Illusionary Rod (7713, -0.84 DPS) [dungeon]; Staff of Noh'Orahil (15105, -1.14 DPS) [quest]; Gut Ripper (2164, -1.48 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 340.1 spell_power points (40.28 DPS) | yes | Umbral Wand (5216, +0.00 DPS, sim-verified) [world_drop]; Earthen Rod (9381, -4.69 DPS) [dungeon]; Twisted Nether Wand (249144, -5.57 DPS) [crafted] |

**New at 40:** head: Augural Shroud; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Defiler's Cloth Girdle; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of Dar'Orahil; ranged: Jaina's Firestarter

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (troll, 25300000000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 142.5. Weights run: 1.0s. Verify run: 0.9s. 390 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.605), intellect=not significant (-0.468 ± 0.718), crit=2.617 ± 0.140, hit=4.179 ± 0.053, spell_haste=not significant (-0.840 ± 0.638), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.153 ± 0.605), fire_power=0.859 ± 0.003

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 50.6 spell_power points (7.85 DPS) | yes | Eye of Theradras (17715, -2.85 DPS, sim-verified) [dungeon]; Spellpower Goggles Xtreme Plus (15999, -3.66 DPS) [crafted]; Dreamweave Circlet (10041, -4.59 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 7.0 spell_power points (1.08 DPS) | yes | Mindburst Medallion (11196, -0.37 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (138.5 DPS) | yes | Black Mageweave Shoulders (10027, -0.46 DPS) [crafted]; Bloodmage Mantle (7684, -0.62 DPS) [dungeon]; Blood Guard's Dreadweave Mantle (220905, -1.94 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 14.0 spell_power points (2.17 DPS) | yes | Deep Woodlands Cloak (19121, -0.14 DPS, sim-verified) [quest]; Nightfall Drape (12465, -0.77 DPS) [dungeon]; Runecloth Cloak (13860, -0.77 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | sim-verified (138.7 DPS) | yes | Elemental Raiment (9434, -0.15 DPS) [world_drop]; Acumen Robes (17775, -0.46 DPS) [quest]; Stone Guard's Dreadweave Vest (220904, -2.16 DPS, sim-verified) [vendor] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.39 DPS) | yes | Condor Bracers (15864, +0.00 DPS, sim-verified) [quest]; Nethergeld Cuffs (254061, -0.31 DPS) [crafted]; Bloodband Bracers (11469, -0.62 DPS) [quest] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 18.0 spell_power points (2.79 DPS) | yes | Black Mageweave Gloves (10003, +0.00 DPS, sim-verified) [crafted]; Brightcloth Gloves (14101, -0.77 DPS) [crafted]; First Sergeant's Dreadweave Gloves (220908, -0.77 DPS) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 45.6 spell_power points (7.07 DPS) | yes | Satyrmane Sash (17755, +0.00 DPS, sim-verified) [dungeon]; Defiler's Cloth Girdle (20166, -4.90 DPS) [rep]; Ghostweave Cord (254073, -4.90 DPS) [crafted] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 48.6 spell_power points (7.54 DPS) | yes | Wizardweave Leggings (14132, +0.00 DPS, sim-verified) [crafted]; Red Mageweave Pants (10009, -5.37 DPS) [crafted]; Gaze Dreamer Pants (6903, -5.68 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (140.1 DPS) | yes | Black Mageweave Boots (10026, -2.01 DPS) [crafted]; Gilded Sandals (254107, -2.01 DPS) [crafted]; First Sergeant's Dreadweave Boots (220909, -3.55 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 41.8 spell_power points (6.48 DPS) | yes | Advisor's Ring (19519, -4.62 DPS) [rep]; Philanthropist's Ring (281635, -4.93 DPS) [quest]; Advisor's Ring (19520, -5.08 DPS) [rep] |
| finger2 | Band of the Unicorn (7553) | World drop [world_drop] | 13.0 spell_power points (2.01 DPS) | yes | Advisor's Ring (19519, +0.00 DPS, sim-verified) [rep]; Philanthropist's Ring (281635, -0.46 DPS) [quest]; Advisor's Ring (19520, -0.62 DPS) [rep] |
| trinket1 | Abyss Shard (20534) | Trolls of a Feather [quest] | sim-verified (136.6 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.24 DPS, sim-verified) [world_drop] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (136.6 DPS) | yes | Uther's Strength (11302, +0.00 DPS, sim-verified) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Staff of Dar'Orahil (15106) | The Completed Orb of Dar'Orahil [quest] | sim-verified (136.6 DPS) | yes | Soul Harvester (20536, +0.00 DPS) [quest]; Shortsword of Vengeance (754, -0.59 DPS, sim-verified) [world_drop]; Illusionary Rod (7713, -0.80 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 338.8 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.73 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -5.42 DPS, sim-verified) [quest] |

**New at 50:** head: Blood Guard's Dreadweave Hat; shoulder: Rotgrip Mantle; back: Spritecaster Cape; wrist: Spidertank Oilrag; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; finger1: Blackstone Ring; finger2: Band of the Unicorn; trinket1: Abyss Shard; trinket2: Rune of the Guard Captain; ranged: Pyric Caduceus

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (troll, 25532300000000000-0000000000000000000-2353225100101051)

Set DPS (verified): 330.0. Weights run: 1.0s. Verify run: 0.9s. 938 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=not significant (1.000 ± 0.338), intellect=not significant (0.993 ± 0.330), crit=1.580 ± 0.087, hit=2.312 ± 0.031, spell_haste=not significant (0.018 ± 0.286), spell_penetration=not significant (0.000 ± 0.000), shadow_power=not significant (0.566 ± 0.338), fire_power=0.427 ± 0.001

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Heretic Cowl (240141) | Leonid Barthalomew the Revered [vendor] | 143.3 spell_power points (53.14 DPS) | yes | Bloodvine Goggles (19999, -21.89 DPS, sim-verified) [crafted]; Deathmist Mask (226909, -29.05 DPS) [quest]; Deathmist Mask (22074, -29.79 DPS) [quest] |
| neck | Amulet of the Dawn (22657) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (325.9 DPS) | yes | Beads of Ogre Mojo (22149, -1.11 DPS) [quest]; Beads of Ogre Might (22150, -1.77 DPS) [quest]; Blazefury Medallion (17111, -6.02 DPS, sim-verified) [world] |
| shoulder | Heretic Shoulderpads (240143) | Leonid Barthalomew the Revered [vendor] | 98.1 spell_power points (36.40 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -7.60 DPS, sim-verified) [vendor]; Heretic Mantle (240150, -17.09 DPS) [vendor]; Mantle of the Timbermaw (19050, -17.11 DPS) [crafted] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 47.1 spell_power points (17.46 DPS) | yes | Hide of the Wild (18510, +0.00 DPS, sim-verified) [crafted]; Earthweave Cloak (21187, -8.88 DPS) [quest]; Howler's Furs (272414, -8.88 DPS) [vendor] |
| chest | Heretic Robe (240138) | Leonid Barthalomew the Revered [vendor] | 126.1 spell_power points (46.77 DPS) | yes | Bloodvine Vest (19682, -13.56 DPS, sim-verified) [crafted]; Heretic Garb (240146, -18.53 DPS) [vendor]; Earthpower Vest (21183, -21.20 DPS) [quest] |
| wrist | Heretic Bindings (240145) | Leonid Barthalomew the Revered [vendor] | 62.0 spell_power points (23.01 DPS) | yes | Rockfury Bracers (21186, -1.34 DPS, sim-verified) [quest]; Heretic Wristguards (240152, -9.24 DPS) [vendor]; Dryad's Wrist Bindings (19595, -11.90 DPS) [rep] |
| hands | Heretic Gloves (240140) | Leonid Barthalomew the Revered [vendor] | 73.0 spell_power points (27.09 DPS) | yes | Deathmist Wraps (22077, -8.90 DPS) [quest]; Deathmist Wraps (226911, -8.90 DPS) [quest]; Gloves of Spell Mastery (14146, -9.68 DPS, sim-verified) [crafted] |
| waist | Heretic Belt (240144) | Leonid Barthalomew the Revered [vendor] | 120.3 spell_power points (44.60 DPS) | yes | Heretic Waistguard (240151, -10.09 DPS, sim-verified) [vendor]; Knowledge of the Timbermaw (228190, -18.65 DPS) [vendor]; Belt of the Archmage (18405, -23.09 DPS) [crafted] |
| legs | Heretic Leggings (240142) | Leonid Barthalomew the Revered [vendor] | 139.2 spell_power points (51.64 DPS) | yes | Sentinel's Silk Leggings (237815, -11.98 DPS, sim-verified) [vendor]; Heretic Pants (240149, -23.77 DPS) [vendor]; Bloodvine Leggings (19683, -27.13 DPS) [crafted] |
| feet | Heretic Sandals (240139) | Leonid Barthalomew the Revered [vendor] | 73.0 spell_power points (27.08 DPS) | yes | Bloodvine Boots (19684, -6.51 DPS, sim-verified) [crafted]; Heretic Boots (240153, -9.26 DPS) [vendor]; First Sergeant's Dreadweave Boots (220909, -12.23 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (325.9 DPS) | yes | Signet Ring of the Bronze Dragonflight (234028, -1.11 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234024, -1.48 DPS) [vendor]; Wrath of Cenarius (21190, -6.55 DPS, sim-verified) [quest] |
| finger2 | Mindtear Band (20632) | Taerar [world] | sim-verified (325.9 DPS) | yes | Ritssyn's Ring of Chaos (21836, -1.10 DPS) [world_drop]; Don Julio's Band (19325, -1.79 DPS) [rep]; Wrath of Cenarius (21190, -3.77 DPS, sim-verified) [quest] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (325.9 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Abyss Shard (20534, -0.74 DPS) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (325.9 DPS) | yes | Frozen Heart of the Mountain (249469, -0.44 DPS) [crafted]; Rune of the Guard Captain (19120, -2.16 DPS) [quest]; Abyss Shard (20534, -4.26 DPS, sim-verified) [quest] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | sim-verified (325.9 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -4.06 DPS, sim-verified) [world_drop]; Soul Harvester (20536, -5.62 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (330.0 DPS) | yes | Cold Snap (19130, -4.07 DPS, sim-verified) [world]; Oblivion's Touch (18761, -12.43 DPS) [dungeon]; Sparkling Crystal Wand (20672, -13.34 DPS) [world] |

**New at 60:** head: Heretic Cowl; neck: Amulet of the Dawn; shoulder: Heretic Shoulderpads; back: Arcanoweave Cloak; chest: Heretic Robe; wrist: Heretic Bindings; hands: Heretic Gloves; waist: Heretic Belt; legs: Heretic Leggings; feet: Heretic Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Mindtear Band; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 938, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3556 Dread Mage Hat; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

