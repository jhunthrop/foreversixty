# Leveling BiS: Fire

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (gnome, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 33.0. Weights run: 0.8s. Verify run: 0.6s. 128 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.965 ± 0.012, crit=0.972 ± 0.082, hit=1.826 ± 0.028, spell_haste=-0.708 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.45 DPS) | yes | Shadow Goggles (4373, -0.55 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.7 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.40 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | sim-verified (32.2 DPS) | yes | Sanguine Cape (14376, -0.01 DPS) [world_drop]; Feyscale Cloak (6632, -0.08 DPS) [dungeon]; Pearl-clasped Cloak (5542, -0.43 DPS, sim-verified) [crafted] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.8 spell_power points (0.74 DPS) | yes | Mystic's Wrap (14369, -0.23 DPS) [world_drop]; Mystic's Robe (14371, -0.23 DPS) [world_drop]; Gray Woolen Robe (2585, -0.47 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (32.4 DPS) | yes | Bright Bracers (3647, -0.07 DPS) [world_drop]; Mystic's Bracelets (14366, -0.22 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.64 DPS, sim-verified) [quest] |
| hands | Pristine Gloves (253913) | Tailoring [crafted] | sim-verified (32.4 DPS) | yes | Blight Gloves (279877, -0.01 DPS) [quest]; Gnoll Casting Gloves (892, -0.07 DPS) [world]; Serpent Gloves (5970, -0.60 DPS, sim-verified) [dungeon] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.9 spell_power points (0.59 DPS) | yes | Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.24 DPS, sim-verified) [world_drop] |
| legs | Filigreed Pristine Leggings (253937) | Tailoring [crafted] | sim-verified (32.2 DPS) | yes | Silk-threaded Trousers (1929, -0.36 DPS) [dungeon]; Darkweave Breeches (12987, -0.38 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.44 DPS, sim-verified) [dungeon] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.18 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.44 DPS) [world]; Red Woolen Boots (4313, -0.52 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 6.9 spell_power points (0.52 DPS) | yes | Loop of Sacrifice (281673, -0.16 DPS) [quest]; Sludge-Stained Band (286535, -0.30 DPS) [world]; Lavishly Jeweled Ring (1156, -0.57 DPS, sim-verified) [dungeon] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | sim-verified (32.2 DPS) | yes | Loop of Sacrifice (281673, -0.01 DPS) [quest]; Sludge-Stained Band (286535, -0.15 DPS) [world]; Lavishly Jeweled Ring (1156, -0.44 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 9.6 spell_power points (0.73 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 299.5 spell_power points (22.54 DPS) | yes | Skycaller (12984, -0.45 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Pristine Gloves; waist: Pristine Sash; legs: Filigreed Pristine Leggings; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 128, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 6478 Rat Stompers; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (gnome, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 60.6. Weights run: 1.1s. Verify run: 0.6s. 224 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=1.146 ± 0.023, crit=2.538 ± 0.231, hit=2.838 ± 0.050, spell_haste=not significant (0.015 ± 0.231), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 17.5 spell_power points (1.61 DPS) | yes | Nightsky Cowl (4039, -0.40 DPS, sim-verified) [world_drop]; Shadow Hood (4323, -0.45 DPS) [crafted]; Resilient Cap (14401, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.9 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.17 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.3 spell_power points (1.78 DPS) | yes | Fairywing Mantle (9536, -0.28 DPS) [quest]; Death Speaker Mantle (6685, -0.33 DPS, sim-verified) [dungeon]; Magician's Mantle (12998, -0.37 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.2 spell_power points (0.84 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.15 DPS) [quest]; Resilient Cape (14400, -0.21 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.9 spell_power points (2.20 DPS) | yes | Death Speaker Robes (6682, -0.58 DPS, sim-verified) [dungeon]; Mechbuilder's Overalls (9508, -0.62 DPS) [dungeon]; Pristine Gown (253961, -0.82 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (60.6 DPS) | yes | Nightsky Wristbands (6407, -0.20 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.20 DPS) [quest]; Glowing Magical Bracelets (13106, -0.95 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 16.6 spell_power points (1.53 DPS) | yes | Truefaith Gloves (7049, -0.75 DPS) [crafted]; Blight Gloves (279877, -0.79 DPS) [quest]; Hotshot Pilot's Gloves (9491, -0.98 DPS, sim-verified) [dungeon] |
| waist | Highlander's Cloth Girdle (20099) | The League of Arathor [rep] | 14.4 spell_power points (1.33 DPS) | yes | Invoker's Cord (215366, -0.16 DPS) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Crimson Silk Belt (7055, -0.34 DPS, sim-verified) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 18.2 spell_power points (1.67 DPS) | yes | Pristine Leggings (253987, +0.00 DPS, sim-verified) [crafted]; Filigreed Pristine Leggings (253937, -0.49 DPS) [crafted]; Necromancer Leggings (2277, -0.51 DPS) [world_drop] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 spell_power points (1.38 DPS) | yes | Spidersilk Boots (4320, -0.32 DPS) [crafted]; Frothing Slippers (254003, -0.65 DPS) [crafted]; Acidic Walkers (9454, -1.25 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.0 spell_power points (0.74 DPS) | yes | Minor Channeling Ring (1449, -0.07 DPS) [quest]; Lorekeeper's Ring (19525, -0.09 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 spell_power points (0.74 DPS) | yes | Lorekeeper's Ring (19525, -0.09 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Minor Channeling Ring (1449, -0.33 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -2.14 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 12.6 spell_power points (1.16 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.11 DPS) [quest]; Twisted Chanter's Staff (890, -0.16 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.32 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 363.3 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.16 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Highlander's Cloth Girdle; legs: Abomination Skin Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers

### Band 40 (gnome, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 100.5. Weights run: 1.1s. Verify run: 0.7s. 307 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=0.993 ± 0.029, crit=2.791 ± 0.286, hit=3.869 ± 0.072, spell_haste=6.535 ± 0.548, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.64 DPS) | yes | Augural Shroud (2620, -0.13 DPS, sim-verified) [world]; Corpseshroud (10574, -0.27 DPS) [dungeon]; Thinking Cap (2624, -0.52 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.0 spell_power points (1.63 DPS) | yes | Necklace of Calisea (1714, -0.76 DPS) [world_drop]; Triune Amulet (7722, -0.76 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -3.44 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.9 spell_power points (2.51 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.25 DPS) [dungeon]; Berylline Pads (4197, -0.37 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (100.5 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.03 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.0 spell_power points (3.52 DPS) | yes | Robe of Power (7054, -0.26 DPS) [crafted]; Dreamweave Vest (10021, -0.44 DPS, sim-verified) [crafted]; Crimson Silk Vest (7058, -0.76 DPS) [crafted] |
| wrist | Arcane Runed Bracers (4744) (or Spidertank Oilrag (9448)) | Wanted! Marez Cowl [quest] | 9.0 spell_power points (1.13 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.01 DPS) [world_drop]; Aurora Bracers (4043, -0.13 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 spell_power points (2.76 DPS) | yes | Stormcloth Gloves (10011, -0.76 DPS) [crafted]; Black Mageweave Gloves (10003, -0.88 DPS) [crafted]; Red Mageweave Gloves (10018, -1.14 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.9 spell_power points (2.75 DPS) | yes | Gilded Cord (254037, -0.75 DPS) [crafted]; Highlander's Cloth Girdle (20099, -1.00 DPS) [rep]; Highlander's Cloth Girdle (20098, -1.58 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.9 spell_power points (3.26 DPS) | yes | Crimson Silk Pantaloons (7062, -0.83 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.02 DPS) | yes | Acidic Walkers (9454, -1.39 DPS) [dungeon]; Gilded Slippers (254001, -1.60 DPS, sim-verified) [crafted]; Spidersilk Boots (4320, -1.64 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.0 spell_power points (2.01 DPS) | yes | Ring of Forlorn Spirits (2043, -1.00 DPS) [quest]; Reedknot Ring (9622, -1.13 DPS) [quest]; Minor Channeling Ring (1449, -1.13 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.13 DPS) | yes | Reedknot Ring (9622, -0.25 DPS) [quest]; Lorekeeper's Ring (19525, -0.25 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.52 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, -3.06 DPS, sim-verified) [world_drop]; Spellforce Rod (1664, -3.15 DPS) [world_drop]; Windweaver Staff (7757, -3.79 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 319.0 spell_power points (40.13 DPS) | yes | Nether Force Wand (11263, -1.65 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.56 DPS) [quest]; Ragefire Wand (7513, -2.61 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Arcane Runed Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 307, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (gnome, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 139.3. Weights run: 0.6s. Verify run: 0.7s. 404 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=1.186 ± 0.069, crit=4.846 ± 0.439, hit=7.393 ± 0.114, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Dreadweave Hat (220889) | Captain Dirgehammer [vendor] | 96.1 spell_power points (12.19 DPS) | yes | Eye of Theradras (17715, +0.00 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -6.77 DPS) [crafted]; Dreamweave Circlet (10041, -8.02 DPS) [crafted] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (131.0 DPS) | yes | Scorn's Icy Choker (23169, -0.32 DPS) [dungeon]; Mindburst Medallion (11196, -0.44 DPS) [quest]; Arcane Crystal Pendant (20037, -3.18 DPS, sim-verified) [quest] |
| shoulder | Rotgrip Mantle (17732) | Maraudon: Rotgrip [dungeon] | sim-verified (130.9 DPS) | yes | Kentic Amice (11624, -0.63 DPS) [dungeon]; Red Mageweave Shoulders (10029, -1.21 DPS) [crafted]; Knight-Lieutenant's Dreadweave Mantle (220887, -3.12 DPS, sim-verified) [vendor] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 21.1 spell_power points (2.68 DPS) | yes | Runecloth Cloak (13860, -0.33 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.57 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -4.63 DPS, sim-verified) [dungeon] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (131.3 DPS) | yes | Runecloth Robe (13858, -1.47 DPS) [crafted]; Runecloth Tunic (13857, -1.61 DPS) [crafted]; Knight's Dreadweave Vest (220886, -3.46 DPS, sim-verified) [vendor] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (131.6 DPS) | yes | Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Forgotten Wraps (9433, -0.18 DPS) [world_drop]; Aristocratic Cuffs (12546, -3.80 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 102.5 spell_power points (13.01 DPS) | yes | Raider Handwraps (272098, +0.00 DPS, sim-verified) [vendor]; Sergeant Major's Dreadweave Gloves (220890, -10.01 DPS) [vendor]; Red Mageweave Gloves (10018, -10.11 DPS) [crafted] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 82.8 spell_power points (10.50 DPS) | yes | Ban'thok Sash (11662, -3.52 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -6.88 DPS) [dungeon]; Satyrmane Sash (17755, -7.22 DPS) [dungeon] |
| legs | Knight's Dreadweave Leggings (220888) | Captain Dirgehammer [vendor] | 94.1 spell_power points (11.94 DPS) | yes | Spellshock Leggings (9484, -6.70 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -8.36 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.85 DPS) [dungeon] |
| feet | Sergeant Major's Dreadweave Boots (220891) | Captain Dirgehammer [vendor] | 92.6 spell_power points (11.75 DPS) | yes | Earthen Silk Slippers (254013, -0.09 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -9.00 DPS) [crafted]; Southsea Mojo Boots (20641, -9.08 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 73.9 spell_power points (9.38 DPS) | yes | Cyclopean Band (11824, -7.19 DPS) [dungeon]; Philanthropist's Ring (281635, -7.21 DPS) [quest]; Mindseye Circle (10634, -7.58 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.8 spell_power points (2.26 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -0.08 DPS) [quest]; Mindseye Circle (10634, -0.45 DPS) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Uther's Strength (11302, -7.68 DPS) [world_drop] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 spell_power points (0.00 DPS) | yes | Uther's Strength (11302, -0.90 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.90 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.81 DPS) [dungeon]; Shortsword of Vengeance (754, -4.78 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -26.13 DPS, sim-verified) [quest] |

**New at 50:** head: Knight-Lieutenant's Dreadweave Hat; neck: Horizon Choker; shoulder: Rotgrip Mantle; back: Spritecaster Cape; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Highlander's Cloth Girdle; legs: Knight's Dreadweave Leggings; feet: Sergeant Major's Dreadweave Boots; finger1: Blackstone Ring; finger2: Brainlash; trinket1: Frozen Heart of the Mountain; trinket2: Thunderbrew's Boot Flask; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 404, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (gnome, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 394.0. Weights run: 1.0s. Verify run: 0.9s. 953 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=1.454 ± 0.084, crit=6.656 ± 0.541, hit=12.117 ± 0.288, spell_haste=13.681 ± 1.123, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (361.7 DPS) | yes | Bloodvine Goggles (19999, -4.51 DPS, sim-verified) [crafted]; Field Marshal's Coronet (16441, -6.14 DPS) [vendor]; Field Marshal's Coronet (231604, -6.14 DPS) [vendor] |
| neck | Jewel of Kajaro (19601) | The Jewel of Kajaro [quest] | 0.0 spell_power points (0.00 DPS) | yes | Amulet of the Dawn (22657, +0.00 DPS) [quest]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Beads of Ogre Might (22150, -9.97 DPS, sim-verified) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 156.0 spell_power points (32.00 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -4.72 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.52 DPS) [crafted]; Lieutenant Commander's Silk Mantle (23319, -6.53 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 148.8 spell_power points (30.53 DPS) | yes | Howler's Furs (272414, -5.67 DPS) [vendor]; Stalwart Cloak (272415, -5.67 DPS) [vendor]; Earthweave Cloak (21187, -10.35 DPS, sim-verified) [quest] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 288.2 spell_power points (59.14 DPS) | yes | Fireleaf Robe (240059, +0.00 DPS, sim-verified) [vendor]; Fireleaf Garb (240051, -20.72 DPS) [vendor]; Field Marshal's Silk Vestments (16443, -28.18 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 151.4 spell_power points (31.07 DPS) | yes | Fireleaf Wristwraps (240044, -2.63 DPS) [vendor]; Rockfury Bracers (21186, -15.71 DPS, sim-verified) [quest]; Arcanist Bindings (16799, -24.14 DPS) [world_drop] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 248.7 spell_power points (51.03 DPS) | yes | Gloves of Spell Mastery (14146, -15.65 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -17.79 DPS) [vendor]; Sorcerer's Gloves (22066, -19.53 DPS) [quest] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 186.1 spell_power points (38.18 DPS) | yes | Knowledge of the Timbermaw (228190, -5.71 DPS, sim-verified) [vendor]; Fireleaf Waistguard (240045, -5.85 DPS) [vendor]; Belt of the Archmage (18405, -10.18 DPS) [crafted] |
| legs | Sentinel's Silk Leggings (237815) | Illiyana Moonblaze [vendor] | 232.9 spell_power points (47.79 DPS) | yes | Fireleaf Leggings (240055, +0.00 DPS, sim-verified) [vendor]; Bloodvine Leggings (19683, -13.54 DPS) [crafted]; Sorcerer's Leggings (226933, -14.25 DPS) [quest] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 163.4 spell_power points (33.53 DPS) | yes | Fireleaf Boots (240050, +0.00 DPS, sim-verified) [vendor]; Marshal's Silk Footwraps (16437, -0.19 DPS) [vendor]; Marshal's Silk Footwraps (231606, -0.19 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 0.0 spell_power points (0.00 DPS) | yes | Wrath of Cenarius (21190, +0.00 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -10.69 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -11.40 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 spell_power points (0.00 DPS) | yes | Wrath of Cenarius (21190, +0.00 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -10.69 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -11.40 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -11.70 DPS, sim-verified) [crafted] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (361.4 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Weakness Analyzer (272438, +0.00 DPS) [vendor]; Frozen Heart of the Mountain (249469, -4.16 DPS, sim-verified) [crafted] |
| main_hand | Ironbark Staff (20069) | The League of Arathor [rep] | 0.0 spell_power points (0.00 DPS) | yes | Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Sageblade (22383, -14.57 DPS) [crafted]; Shortsword of Vengeance (754, -18.02 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (383.8 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.88 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.91 DPS) [dungeon]; Cold Snap (19130, -26.58 DPS, sim-verified) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Jewel of Kajaro; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Bloodvine Vest; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Sentinel's Silk Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 953, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

## Horde

### Band 20 (orc, 000000000000000000-23510000000000000-0000000000000000000)

Set DPS (verified): 29.9. Weights run: 0.8s. Verify run: 0.6s. 124 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.965 ± 0.012, crit=0.972 ± 0.082, hit=1.826 ± 0.028, spell_haste=-0.708 ± 0.159, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Pristine Circlet (253949) | Tailoring [crafted] | 6.0 spell_power points (0.45 DPS) | yes | Shadow Goggles (4373, -0.60 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 13.7 spell_power points (1.03 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.60 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -0.73 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 4.9 spell_power points (0.37 DPS) | yes | Sanguine Cape (14376, -0.08 DPS) [world_drop]; Heavy Woolen Cloak (4311, -0.10 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.14 DPS) [dungeon] |
| chest | Filigreed Pristine Gown (253901) | Tailoring [crafted] | 9.8 spell_power points (0.74 DPS) | yes | Mystic's Wrap (14369, -0.23 DPS) [world_drop]; Mystic's Robe (14371, -0.23 DPS) [world_drop]; Gray Woolen Robe (2585, -0.55 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (29.9 DPS) | yes | Featherbead Bracers (15452, +0.00 DPS) [quest]; Bright Bracers (3647, -0.07 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.30 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | 7.0 spell_power points (0.53 DPS) | yes | Pristine Gloves (253913, +0.00 DPS, sim-verified) [crafted]; Blight Gloves (279877, -0.02 DPS) [quest]; Gnoll Casting Gloves (892, -0.08 DPS) [world] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | 7.9 spell_power points (0.59 DPS) | yes | Novice Arcanist's Sash (253885, -0.07 DPS) [crafted]; Novice Ardent's Sash (253887, -0.22 DPS) [crafted]; Keller's Girdle (2911, -0.41 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 16.7 spell_power points (1.26 DPS) | yes | Filigreed Pristine Leggings (253937, +0.00 DPS, sim-verified) [crafted]; Silk-threaded Trousers (1929, -0.73 DPS) [dungeon]; Darkweave Breeches (12987, -0.75 DPS) [world_drop] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 10.9 spell_power points (0.82 DPS) | yes | Pristine Boots (253889, -0.40 DPS, sim-verified) [crafted]; Feather Padded Treads (285345, -0.44 DPS) [world]; Red Woolen Boots (4313, -0.52 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 5.8 spell_power points (0.44 DPS) | yes | Loop of Sacrifice (281673, -0.07 DPS) [quest]; Sludge-Stained Band (286535, -0.21 DPS) [world]; Volcanic Rock Ring (12053, -0.22 DPS) [world_drop] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.38 DPS) | yes | Sludge-Stained Band (286535, -0.15 DPS) [world]; Volcanic Rock Ring (12053, -0.16 DPS) [world_drop]; Loop of Sacrifice (281673, -0.35 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 9.6 spell_power points (0.73 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.15 DPS) [world]; Lesser Staff of the Spire (1300, -0.29 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Cookie's Stirring Rod (5198) | Westfall: Cookie [dungeon] | 299.5 spell_power points (22.54 DPS) | yes | Skycaller (12984, -0.58 DPS, sim-verified) [world_drop]; Firebelcher (5243, -2.25 DPS) [dungeon]; Deepblaze (279896, -4.01 DPS) [quest] |

**New at 20:** head: Pristine Circlet; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Filigreed Pristine Gown; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Abomination Skin Leggings; feet: Spidersilk Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; main_hand: Twisted Chanter's Staff; ranged: Cookie's Stirring Rod

No-known-source sample (15 of 124, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18859 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20430 Legionnaire's Sword; 20441 Scout's Blade; 209618 Insignia of the Alliance; 241089 Scarlet Dagger

### Band 30 (orc, 000000000000000000-23552110020000000-0000000000000000000)

Set DPS (verified): 56.2. Weights run: 1.1s. Verify run: 0.7s. 217 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=1.146 ± 0.023, crit=2.538 ± 0.231, hit=2.838 ± 0.050, spell_haste=not significant (0.015 ± 0.231), spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 17.5 spell_power points (1.61 DPS) | yes | Shadow Hood (4323, -0.45 DPS) [crafted]; Resilient Cap (14401, -0.45 DPS) [world_drop]; Nightsky Cowl (4039, -0.47 DPS, sim-verified) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.9 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -0.86 DPS) [world_drop]; Kaleidoscope Chain (13084, -0.86 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.20 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.3 spell_power points (1.78 DPS) | yes | Death Speaker Mantle (6685, -0.10 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.28 DPS) [quest]; Magician's Mantle (12998, -0.37 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.2 spell_power points (0.84 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.21 DPS) [world_drop]; Soft Willow Cape (16661, -0.32 DPS) [quest] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.9 spell_power points (2.20 DPS) | yes | Mechbuilder's Overalls (9508, -0.62 DPS) [dungeon]; Death Speaker Robes (6682, -0.65 DPS, sim-verified) [dungeon]; Pristine Gown (253961, -0.82 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | sim-verified (55.4 DPS) | yes | Nightsky Wristbands (6407, -0.20 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.20 DPS) [quest]; Glowing Magical Bracelets (13106, -0.66 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 11.7 spell_power points (1.08 DPS) | yes | Hotshot Pilot's Gloves (9491, -0.25 DPS, sim-verified) [dungeon]; Truefaith Gloves (7049, -0.30 DPS) [crafted]; Blight Gloves (279877, -0.34 DPS) [quest] |
| waist | Defiler's Cloth Girdle (20164) | The Defilers [rep] | 14.4 spell_power points (1.33 DPS) | yes | Invoker's Cord (215366, -0.16 DPS) [crafted]; Belt of Arugal (6392, -0.18 DPS) [dungeon]; Crimson Silk Belt (7055, -0.26 DPS, sim-verified) [crafted] |
| legs | Pristine Leggings (253987) | Tailoring [crafted] | sim-verified (55.4 DPS) | yes | Filigreed Pristine Leggings (253937, -0.20 DPS) [crafted]; Necromancer Leggings (2277, -0.22 DPS) [world_drop]; Abomination Skin Leggings (23173, -0.63 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 spell_power points (1.38 DPS) | yes | Spidersilk Boots (4320, -0.32 DPS) [crafted]; Frothing Slippers (254003, -0.65 DPS) [crafted]; Acidic Walkers (9454, -1.11 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.0 spell_power points (0.74 DPS) | yes | Advisor's Ring (19521, -0.09 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.19 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 spell_power points (0.74 DPS) | yes | Advisor's Ring (19521, +0.00 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.19 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Defiler's Talisman (21120, -2.39 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 12.6 spell_power points (1.16 DPS) | yes | Gnarled Necromancer's Staff (251534, -0.11 DPS) [quest]; Twisted Chanter's Staff (890, -0.30 DPS, sim-verified) [world_drop]; Channeler's Staff (4437, -0.32 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Necrotic Wand (7708) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 363.3 spell_power points (33.49 DPS) | yes | Starfaller (13063, -0.11 DPS, sim-verified) [world_drop]; Greater Mystic Wand (217287, -4.03 DPS) [crafted]; Gravestone Scepter (7001, -4.49 DPS) [quest] |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Defiler's Cloth Girdle; legs: Pristine Leggings; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; main_hand: Glimmering Staff; ranged: Necrotic Wand

No-known-source sample (15 of 217, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 40 (orc, 000000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 93.3. Weights run: 1.1s. Verify run: 0.6s. 300 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=0.993 ± 0.029, crit=2.791 ± 0.286, hit=3.869 ± 0.072, spell_haste=6.535 ± 0.548, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.64 DPS) | yes | Augural Shroud (2620, +0.00 DPS, sim-verified) [world]; Corpseshroud (10574, -0.27 DPS) [dungeon]; Thinking Cap (2624, -0.52 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.0 spell_power points (1.63 DPS) | yes | Necklace of Calisea (1714, -0.76 DPS) [world_drop]; Triune Amulet (7722, -0.76 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.02 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 19.9 spell_power points (2.51 DPS) | yes | Green Silken Shoulders (7057, +0.00 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.25 DPS) [dungeon]; Berylline Pads (4197, -0.37 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (93.3 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.01 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.71 DPS, sim-verified) [dungeon] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 28.0 spell_power points (3.52 DPS) | yes | Dreamweave Vest (10021, +0.00 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.26 DPS) [crafted]; Crimson Silk Vest (7058, -0.76 DPS) [crafted] |
| wrist | Radiant Silver Bracers (4545) | Foul Magics [quest] | 11.9 spell_power points (1.50 DPS) | yes | Spidertank Oilrag (9448, +0.00 DPS, sim-verified) [dungeon]; Windchaser Cuffs (14429, -0.38 DPS) [world_drop]; Aurora Bracers (4043, -0.50 DPS) [world_drop] |
| hands | Dreamweave Gloves (10019) | Tailoring [crafted] | 22.0 spell_power points (2.76 DPS) | yes | Stormcloth Gloves (10011, -0.76 DPS) [crafted]; Red Mageweave Gloves (10018, -0.78 DPS, sim-verified) [crafted]; Black Mageweave Gloves (10003, -0.88 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 21.9 spell_power points (2.75 DPS) | yes | Gilded Cord (254037, -0.75 DPS) [crafted]; Defiler's Cloth Girdle (20166, -0.90 DPS, sim-verified) [rep]; Defiler's Cloth Girdle (20164, -1.00 DPS) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 25.9 spell_power points (3.26 DPS) | yes | Crimson Silk Pantaloons (7062, -0.60 DPS, sim-verified) [crafted]; Abomination Skin Leggings (23173, -1.13 DPS) [dungeon]; Stoneweaver Leggings (9407, -1.38 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.02 DPS) | yes | Gilded Slippers (254001, -0.73 DPS, sim-verified) [crafted]; Acidic Walkers (9454, -1.39 DPS) [dungeon]; Spidersilk Boots (4320, -1.64 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 16.0 spell_power points (2.01 DPS) | yes | Reedknot Ring (9622, -1.13 DPS) [quest]; Ogremind Ring (1993, -1.13 DPS) [world_drop]; Voodoo Band (1996, -1.13 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.13 DPS) | yes | Advisor's Ring (19521, -0.25 DPS) [rep]; Ogremind Ring (1993, -0.26 DPS) [world_drop]; Reedknot Ring (9622, -1.09 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Gut Ripper (2164, -1.46 DPS, sim-verified) [world_drop]; Spellforce Rod (1664, -3.15 DPS) [world_drop]; Windweaver Staff (7757, -3.79 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Jaina's Firestarter (13064) | World drop [world_drop] | 319.0 spell_power points (40.13 DPS) | yes | Nether Force Wand (11263, -1.60 DPS, sim-verified) [quest]; Icefury Wand (7514, -2.56 DPS) [quest]; Ragefire Wand (7513, -2.61 DPS) [quest] |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Radiant Silver Bracers; hands: Dreamweave Gloves; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; main_hand: Illusionary Rod; ranged: Jaina's Firestarter

No-known-source sample (15 of 300, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring; 5742 Gemstone Dagger

### Band 50 (orc, 253000000000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 120.2. Weights run: 0.6s. Verify run: 0.7s. 397 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.005, intellect=1.186 ± 0.069, crit=4.846 ± 0.439, hit=7.393 ± 0.114, spell_haste=7.836 ± 1.430, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.005

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Dreadweave Hat (220907) | Lady Palanseer [vendor] | 96.1 spell_power points (12.19 DPS) | yes | Eye of Theradras (17715, -0.58 DPS, sim-verified) [dungeon]; Red Mageweave Headband (10033, -6.77 DPS) [crafted]; Dreamweave Circlet (10041, -8.02 DPS) [crafted] |
| neck | Arcane Crystal Pendant (20037) | Destroy Morphaz [quest] | 23.1 spell_power points (2.93 DPS) | yes | Horizon Choker (13085, +0.00 DPS, sim-verified) [world_drop]; Scorn's Icy Choker (23169, -1.14 DPS) [dungeon]; Mindburst Medallion (11196, -1.27 DPS) [quest] |
| shoulder | Blood Guard's Dreadweave Mantle (220905) | Lady Palanseer [vendor] | 86.5 spell_power points (10.98 DPS) | yes | Rotgrip Mantle (17732, +0.00 DPS, sim-verified) [dungeon]; Kentic Amice (11624, -7.25 DPS) [dungeon]; Red Mageweave Shoulders (10029, -7.83 DPS) [crafted] |
| back | Deep Woodlands Cloak (19121) | Wanted: Vile Priestess Hexx and Her Minions [quest] | 22.7 spell_power points (2.88 DPS) | yes | Spritecaster Cape (11623, -0.07 DPS, sim-verified) [dungeon]; Mantle of Lady Falther'ess (23178, -0.38 DPS) [dungeon]; Runecloth Cloak (13860, -0.53 DPS) [crafted] |
| chest | Acumen Robes (17775) | Twisted Evils [quest] | sim-verified (118.9 DPS) | yes | Stone Guard's Dreadweave Vest (220904, -1.32 DPS, sim-verified) [vendor]; Runecloth Robe (13858, -1.47 DPS) [crafted]; Runecloth Tunic (13857, -1.61 DPS) [crafted] |
| wrist | Bloodband Bracers (11469) | The Captain's Chest [quest] | sim-verified (119.0 DPS) | yes | Nethergeld Cuffs (254061, -0.05 DPS) [crafted]; Forgotten Wraps (9433, -0.18 DPS) [world_drop]; Aristocratic Cuffs (12546, -1.41 DPS, sim-verified) [dungeon] |
| hands | Sorcerer's Gauntlets (226930) | Mokvar [vendor] | 102.5 spell_power points (13.01 DPS) | yes | Raider Handwraps (272098, -0.69 DPS, sim-verified) [vendor]; First Sergeant's Dreadweave Gloves (220908, -10.01 DPS) [vendor]; Red Mageweave Gloves (10018, -10.11 DPS) [crafted] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 82.8 spell_power points (10.50 DPS) | yes | Ban'thok Sash (11662, -3.41 DPS, sim-verified) [dungeon]; Dawnspire Cord (12466, -6.88 DPS) [dungeon]; Satyrmane Sash (17755, -7.22 DPS) [dungeon] |
| legs | Stone Guard's Dreadweave Leggings (220906) | Lady Palanseer [vendor] | 94.1 spell_power points (11.94 DPS) | yes | Spellshock Leggings (9484, -5.31 DPS, sim-verified) [dungeon]; Red Mageweave Pants (10009, -8.36 DPS) [crafted]; Kilt of the Atal'ai Prophet (10807, -8.85 DPS) [dungeon] |
| feet | First Sergeant's Dreadweave Boots (220909) | Lady Palanseer [vendor] | 92.6 spell_power points (11.75 DPS) | yes | Earthen Silk Slippers (254013, -0.45 DPS, sim-verified) [crafted]; Gilded Sandals (254107, -9.00 DPS) [crafted]; Southsea Mojo Boots (20641, -9.08 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 73.9 spell_power points (9.38 DPS) | yes | Cyclopean Band (11824, -7.19 DPS) [dungeon]; Philanthropist's Ring (281635, -7.21 DPS) [quest]; Mindseye Circle (10634, -7.58 DPS) [dungeon] |
| finger2 | Brainlash (6440) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 17.8 spell_power points (2.26 DPS) | yes | Cyclopean Band (11824, +0.00 DPS, sim-verified) [dungeon]; Philanthropist's Ring (281635, -0.08 DPS) [quest]; Mindseye Circle (10634, -0.45 DPS) [dungeon] |
| trinket1 | Uther's Strength (11302) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, -0.47 DPS, sim-verified) [crafted] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.90 DPS) [dungeon]; Inventor's Focal Sword (17719, -1.81 DPS) [dungeon]; Shortsword of Vengeance (754, -3.58 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Pyric Caduceus (11748) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 413.7 spell_power points (52.50 DPS) | yes | Noxious Shooter (17745, -1.87 DPS) [dungeon]; Wand of Allistarj (13065, -4.08 DPS) [world_drop]; Woestave (20082, -26.15 DPS, sim-verified) [quest] |

**New at 50:** head: Blood Guard's Dreadweave Hat; neck: Arcane Crystal Pendant; shoulder: Blood Guard's Dreadweave Mantle; back: Deep Woodlands Cloak; chest: Acumen Robes; wrist: Bloodband Bracers; hands: Sorcerer's Gauntlets; waist: Defiler's Cloth Girdle; legs: Stone Guard's Dreadweave Leggings; feet: First Sergeant's Dreadweave Boots; finger1: Blackstone Ring; finger2: Brainlash; trinket1: Uther's Strength; trinket2: Ankh of Life; main_hand: Kindling Stave; ranged: Pyric Caduceus

No-known-source sample (15 of 397, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

### Band 60 (orc, 255115100000000000-23552110030003051-0000000000000000000)

Set DPS (verified): 365.7. Weights run: 1.0s. Verify run: 0.9s. 947 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.011, intellect=1.454 ± 0.084, crit=6.656 ± 0.541, hit=12.117 ± 0.288, spell_haste=13.681 ± 1.123, spell_penetration=not significant (0.000 ± 0.000), fire_power=1.000 ± 0.011

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Fireleaf Circlet (240056) | Leonid Barthalomew the Revered [vendor] | sim-verified (326.2 DPS) | yes | Warlord's Silk Cowl (16533, -6.14 DPS) [vendor]; Warlord's Silk Cowl (231601, -6.14 DPS) [vendor]; Bloodvine Goggles (19999, -6.19 DPS, sim-verified) [crafted] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 0.0 spell_power points (0.00 DPS) | yes | Jewel of Kajaro (19601, -0.73 DPS, sim-verified) [quest]; Medallion of the Dawn (22659, -5.74 DPS) [quest]; Amulet of the Dawn (22657, -17.91 DPS) [quest] |
| shoulder | Fireleaf Shoulderpads (240054) | Leonid Barthalomew the Revered [vendor] | 156.0 spell_power points (32.00 DPS) | yes | Rugged Mantle of the Timbermaw (227808, -4.26 DPS, sim-verified) [vendor]; Mantle of the Timbermaw (19050, -5.52 DPS) [crafted]; Champion's Silk Mantle (23264, -6.53 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 148.8 spell_power points (30.53 DPS) | yes | Howler's Furs (272414, -5.67 DPS) [vendor]; Stalwart Cloak (272415, -5.67 DPS) [vendor]; Earthweave Cloak (21187, -5.75 DPS, sim-verified) [quest] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 288.2 spell_power points (59.14 DPS) | yes | Fireleaf Robe (240059, +0.00 DPS, sim-verified) [vendor]; Fireleaf Garb (240051, -20.72 DPS) [vendor]; Warlord's Silk Raiment (16535, -28.18 DPS) [vendor] |
| wrist | Fireleaf Bindings (240052) | Leonid Barthalomew the Revered [vendor] | 151.4 spell_power points (31.07 DPS) | yes | Fireleaf Wristwraps (240044, -2.63 DPS) [vendor]; Rockfury Bracers (21186, -6.23 DPS, sim-verified) [quest]; Arcanist Bindings (16799, -24.14 DPS) [world_drop] |
| hands | Fireleaf Gloves (240057) | Leonid Barthalomew the Revered [vendor] | 248.7 spell_power points (51.03 DPS) | yes | Gloves of Spell Mastery (14146, -8.40 DPS, sim-verified) [crafted]; Fireleaf Mitts (240049, -17.79 DPS) [vendor]; Sorcerer's Gloves (22066, -19.53 DPS) [quest] |
| waist | Fireleaf Belt (240053) | Leonid Barthalomew the Revered [vendor] | 186.1 spell_power points (38.18 DPS) | yes | Knowledge of the Timbermaw (228190, -3.63 DPS, sim-verified) [vendor]; Fireleaf Waistguard (240045, -5.85 DPS) [vendor]; Belt of the Archmage (18405, -10.18 DPS) [crafted] |
| legs | Fireleaf Leggings (240055) | Leonid Barthalomew the Revered [vendor] | sim-verified (329.0 DPS) | yes | Bloodvine Leggings (19683, -2.85 DPS) [crafted]; Sorcerer's Leggings (226933, -3.56 DPS) [quest]; Sentinel's Silk Leggings (237815, -9.03 DPS, sim-verified) [vendor] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 163.4 spell_power points (33.53 DPS) | yes | General's Silk Boots (16539, -0.19 DPS) [vendor]; General's Silk Boots (231597, -0.19 DPS) [vendor]; Fireleaf Boots (240050, -0.71 DPS, sim-verified) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 0.0 spell_power points (0.00 DPS) | yes | Wrath of Cenarius (21190, -1.72 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -10.69 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -11.40 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 0.0 spell_power points (0.00 DPS) | yes | Wrath of Cenarius (21190, -1.72 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234032, -10.69 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -11.40 DPS) [vendor] |
| trinket1 | Darkmoon Card: Blue Dragon (19288) | Darkmoon Beast Deck [quest] | sim-verified (-1.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 spell_power points (0.00 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Serenity Field (272439, -1.81 DPS, sim-verified) [vendor] |
| main_hand | Ironbark Staff (20220) | The Defilers [rep] | 0.0 spell_power points (0.00 DPS) | yes | High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Shortsword of Vengeance (754, -3.27 DPS, sim-verified) [world_drop]; Sageblade (22383, -14.57 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | Torch of Light (279246) | Enchanting [crafted] | sim-verified (352.8 DPS) | yes | Ritssyn's Wand of Bad Mojo (22408, -11.88 DPS) [dungeon]; Bonecreeper Stylus (13938, -11.91 DPS) [dungeon]; Cold Snap (19130, -32.83 DPS, sim-verified) [world] |

**New at 60:** head: Fireleaf Circlet; neck: Beads of Ogre Might; shoulder: Fireleaf Shoulderpads; back: Arcanoweave Cloak; chest: Bloodvine Vest; wrist: Fireleaf Bindings; hands: Fireleaf Gloves; waist: Fireleaf Belt; legs: Fireleaf Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Blue Dragon; trinket2: Talisman of Ascendance; main_hand: Ironbark Staff; ranged: Torch of Light

No-known-source sample (15 of 947, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 3222 Wicked Dagger; 3738 Brewing Rod; 4116 Olmann Sewar; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant

