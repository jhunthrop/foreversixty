# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 29.1. Weights run: 1.1s. Verify run: 0.8s. 168 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.323 ± 0.021, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (27.2 DPS) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.39 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 spell_power points (0.87 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.46 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -1.16 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.41 DPS) | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.11 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 spell_power points (1.02 DPS) | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -1.14 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (27.2 DPS) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.04 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.42 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (27.7 DPS) | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.96 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (27.3 DPS) | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.54 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 spell_power points (1.26 DPS) | yes | Abomination Skin Leggings (23173, -0.03 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.88 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 5.8 spell_power points (0.59 DPS) | yes | Sludge-Stained Band (286535, -0.28 DPS) [world]; Lavishly Jeweled Ring (1156, -0.35 DPS) [dungeon]; Loop of Sacrifice (281673, -0.39 DPS) [quest] |
| finger2 | Lorekeeper's Ring (20431) | Silverwing Sentinels [rep] | 5.0 spell_power points (0.51 DPS) | yes | Lavishly Jeweled Ring (1156, -0.28 DPS) [dungeon]; Loop of Sacrifice (281673, -0.31 DPS) [quest]; Sludge-Stained Band (286535, -0.39 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | Westfall: Rhahk'Zor [dungeon] | 8.0 spell_power points (0.82 DPS) | yes | Twisted Chanter's Staff (890, -0.12 DPS, sim-verified) [world_drop]; Gnarled Necromancer's Staff (251534, -0.42 DPS) [quest]; Channeler's Staff (4437, -0.50 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Sentinel's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Minor Channeling Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Rhahk'Zor's Hammer; ranged: Idol of the Huntress

No-known-source sample (15 of 168, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14148 Crystalline Cuffs

### Band 30 (night-elf, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 53.9. Weights run: 1.4s. Verify run: 0.9s. 296 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.753 ± 0.028, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Holy Shroud (2721) | World drop [world_drop] | sim-verified (52.1 DPS) | yes | Enchanter's Cowl (4322, -0.10 DPS) [crafted]; Silk Headband (7050, -0.27 DPS) [crafted]; Totemic Leather Helm (252456, -0.73 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.12 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 spell_power points (1.72 DPS) | yes | Death Speaker Mantle (6685, -0.01 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Repairman's Cape (9605) | Data Rescue [quest] | sim-verified (52.5 DPS) | yes | Heavy Woolen Cloak (4311, -0.10 DPS) [crafted]; Prelacy Cape (7004, -0.10 DPS) [quest]; Hillman's Cloak (3719, -1.16 DPS, sim-verified) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 spell_power points (1.95 DPS) | yes | Guardian Armor (4256, -0.09 DPS, sim-verified) [crafted]; Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.20 DPS) | yes | Nightsky Wristbands (6407, -0.86 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.86 DPS) [quest]; Glowing Magical Bracelets (13106, -1.18 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | sim-verified (52.8 DPS) | yes | Serpent Gloves (5970, -0.23 DPS) [dungeon]; Shilly Mitts (9609, -0.23 DPS) [quest]; Fletcher's Gloves (7348, -1.42 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 spell_power points (1.68 DPS) | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Belt of Arugal (6392, -0.31 DPS) [dungeon]; Highlander's Cloth Girdle (20099, -0.53 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 spell_power points (1.97 DPS) | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon]; Dark Ritual Leggings (270031, -1.27 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 spell_power points (1.34 DPS) | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -1.84 DPS, sim-verified) [crafted] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.94 DPS) | yes | Minor Channeling Ring (1449, -0.15 DPS) [quest]; Lorekeeper's Ring (20431, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.80 DPS) | yes | Electrocutioner Lagnut (9447, -0.40 DPS) [dungeon]; Sludge-Stained Band (286535, -0.40 DPS) [world]; Minor Channeling Ring (1449, -0.71 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (51.3 DPS) | yes | Talisman of Arathor (21119, -2.15 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-verified (51.3 DPS) | yes | Rhahk'Zor's Hammer (5187, -1.76 DPS) [dungeon]; Glimmering Staff (249392, -2.20 DPS) [crafted]; Manual Crowd Pummeler (9449, -3.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Holy Shroud; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Repairman's Cape; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 296, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 72.2. Weights run: 1.4s. Verify run: 0.9s. 416 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.733 ± 0.044, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.68 DPS) | yes | Big Voodoo Mask (8201, -0.06 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Corpseshroud (10574, -0.98 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 spell_power points (1.43 DPS) | yes | Necklace of Calisea (1714, -0.80 DPS) [world_drop]; Triune Amulet (7722, -0.80 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.42 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 spell_power points (2.06 DPS) | yes | Green Silken Shoulders (7057, -0.05 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (1.21 DPS) | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (72.2 DPS) | yes | Robe of Power (7054, -0.24 DPS) [crafted]; Elemental Raiment (9434, -0.42 DPS) [world_drop]; Robe of the Magi (1716, -1.09 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 spell_power points (1.30 DPS) | yes | Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest]; Arcane Runed Bracers (4744, -0.84 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Dreamweave Gloves (10019, -0.82 DPS, sim-verified) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 spell_power points (2.24 DPS) | yes | Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted]; Highlander's Cloth Girdle (20098, -1.23 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 spell_power points (2.86 DPS) | yes | Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Kodohide Legguards (285338, -0.73 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Skycaller's Leather Shoes (252532, -1.00 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (1.81 DPS) | yes | Ring of Forlorn Spirits (2043, -0.79 DPS) [quest]; Reedknot Ring (9622, -0.92 DPS) [quest]; Minor Channeling Ring (1449, -1.00 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.15 DPS) | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.24 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (71.1 DPS) | yes | Rune of Duty (21567, -2.11 DPS, sim-verified) [rep] |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | sim-verified (71.1 DPS) | yes | Illusionary Rod (7713, -0.13 DPS) [dungeon]; Mograine's Might (7723, -1.44 DPS) [dungeon]; Gut Ripper (2164, -3.89 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (night-elf, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 85.4. Weights run: 1.4s. Verify run: 1.0s. 558 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.310 ± 0.074, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (80.1 DPS) | yes | Soothsayer's Headdress (17740, -0.82 DPS) [dungeon]; Red Mageweave Headband (10033, -0.86 DPS) [crafted]; Knight-Lieutenant's Leather Headband (220850, -3.40 DPS, sim-verified) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 spell_power points (1.31 DPS) | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world_drop]; Gemshard Heart (17707, -0.46 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Crackling Leather Spaulders (220870) | Captain Dirgehammer [vendor] | sim-verified (79.8 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -3.09 DPS, sim-verified) [vendor]; Ironfeather Shoulders (15067, -3.75 DPS) [crafted]; Rotgrip Mantle (17732, -4.15 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 18.4 spell_power points (2.11 DPS) | yes | Big Voodoo Cloak (8216, -0.78 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.93 DPS) [vendor]; Runecloth Cloak (13860, -1.41 DPS, sim-verified) [crafted] |
| chest | Knight's Crackling Leather Tunic (220868) | Captain Dirgehammer [vendor] | sim-verified (79.7 DPS) | yes | Acumen Robes (17775, -1.61 DPS) [quest]; Robe of the Magi (1716, -2.45 DPS) [world_drop]; Knight's Leather Armor (220854, -3.04 DPS, sim-verified) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 spell_power points (1.99 DPS) | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 28.6 spell_power points (3.27 DPS) | yes | Fletcher's Gloves (7348, -0.16 DPS) [crafted]; Shadowskin Gloves (18238, -0.16 DPS) [crafted]; Gloves of Holy Might (867, -5.63 DPS, sim-verified) [world_drop] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 39.9 spell_power points (4.57 DPS) | yes | Highlander's Lizardhide Girdle (20103, -0.75 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20115, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 89.7 spell_power points (10.27 DPS) | yes | Knight's Leather Pants (220858, -2.84 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -4.04 DPS) [crafted]; Knight's Restored Leather Leggings (220882, -6.31 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) | Captain Dirgehammer [vendor] | 61.5 spell_power points (7.04 DPS) | yes | Earthen Silk Slippers (254013, +0.00 DPS, sim-verified) [crafted]; Skycaller's Leather Boots (252471, -4.39 DPS) [crafted]; Skycaller's Leather Shoes (252532, -4.96 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.1 spell_power points (4.94 DPS) | yes | Band of the Unicorn (7553, -3.45 DPS) [world_drop]; Lorekeeper's Ring (19523, -3.56 DPS) [rep]; Brainlash (6440, -3.67 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.65 DPS) | yes | Lorekeeper's Ring (19523, -0.28 DPS) [rep]; Brainlash (6440, -0.39 DPS) [dungeon]; Band of the Unicorn (7553, -0.61 DPS, sim-verified) [world_drop] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (76.7 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (76.7 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.73 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (76.7 DPS) | yes | Illusionary Rod (7713, -0.51 DPS) [dungeon]; Thorium Greatmace (250613, -1.50 DPS) [crafted]; Hammer of the Northern Wind (810, -1.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; shoulder: Knight-Lieutenant's Crackling Leather Spaulders; back: Spritecaster Cape; chest: Knight's Crackling Leather Tunic; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Highlander's Cloth Girdle; legs: Knight's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Thunderbrew's Boot Flask; main_hand: Kindling Stave

No-known-source sample (15 of 558, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (night-elf, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 164.0. Weights run: 1.4s. Verify run: 0.9s. 1298 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.526 ± 0.100, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 175.7 spell_power points (19.43 DPS) | yes | Bloodvine Goggles (19999, -3.32 DPS) [crafted]; Waywatcher Headpiece (240088, -6.57 DPS) [vendor]; Mask of the Unforgiven (13404, -13.80 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (162.1 DPS) | yes | Blazefury Medallion (17111, -1.24 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -2.22 DPS) [quest]; Amulet of the Dawn (22657, -3.56 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 196.4 spell_power points (21.71 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -11.92 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -12.68 DPS) [quest]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -13.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 76.2 spell_power points (8.43 DPS) | yes | Howler's Furs (272414, -2.32 DPS) [vendor]; Stalwart Cloak (272415, -2.32 DPS) [vendor]; Earthweave Cloak (21187, -4.80 DPS, sim-verified) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | sim-verified (162.1 DPS) | yes | Bloodvine Vest (19682, -3.03 DPS) [crafted]; Waywatcher Tunic (240091, -5.82 DPS) [vendor]; Tunic of Undead Slaying (23089, -15.76 DPS, sim-verified) [world] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | sim-verified (162.1 DPS) | yes | Rockfury Bracers (21186, -1.89 DPS) [quest]; Waywatcher Wristguards (240084, -4.10 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -10.83 DPS, sim-verified) [world] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 179.4 spell_power points (19.84 DPS) | yes | Stormshroud Gloves (21278, -9.84 DPS) [crafted]; Primal Batskin Gloves (19686, -11.42 DPS, sim-verified) [crafted]; Waywatcher Grips (240065, -12.77 DPS) [vendor] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 168.8 spell_power points (18.66 DPS) | yes | Knowledge of the Timbermaw (228190, -2.72 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -10.80 DPS) [vendor]; Belt of the Archmage (18405, -11.47 DPS) [crafted] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 217.5 spell_power points (24.05 DPS) | yes | Waywatcher Legguards (240087, -0.48 DPS, sim-verified) [vendor]; Knight's Crackling Leather Leggings (220864, -12.03 DPS) [vendor]; Sentinel's Silk Leggings (237815, -12.04 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 167.4 spell_power points (18.51 DPS) | yes | Bloodvine Boots (19684, -1.60 DPS, sim-verified) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -10.50 DPS) [vendor]; Fine Dawn Treaders (227815, -12.40 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (162.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Wrath of Cenarius (21190, -0.26 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (162.1 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Wrath of Cenarius (21190, -0.26 DPS, sim-verified) [quest]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor] |
| trinket1 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (162.1 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.88 DPS) [world_drop]; Thunderbrew's Boot Flask (744, -1.20 DPS, sim-verified) [quest] |
| trinket2 | Weakness Analyzer (272438) | Pix Xizzix [vendor] | sim-verified (164.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -1.77 DPS) [world_drop]; Darkmoon Card: Blue Dragon (19288, -1.92 DPS, sim-verified) [quest] |
| main_hand | Enchanted Battlehammer (12776) | Blacksmithing [crafted] | sim-verified (162.1 DPS) | yes | Frenzied Striker (13056, +0.00 DPS) [world_drop]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -3.85 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Waywatcher Hood; neck: Beads of Ogre Might; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Serenity Field; trinket2: Weakness Analyzer; main_hand: Enchanted Battlehammer; ranged: Idol of the Moon

No-known-source sample (15 of 1298, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 28.9. Weights run: 1.1s. Verify run: 0.8s. 170 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.386 ± 0.087, crit=0.582 ± 0.038, hit=1.323 ± 0.021, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Wisdom's Leather Hood (252507) | Leatherworking [crafted] | sim-verified (25.5 DPS) | yes | Pristine Circlet (253949, +0.00 DPS) [crafted]; Trapper's Leather Hood (252505, -0.30 DPS) [crafted]; Totemic Leather Hood (252448, -0.38 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 8.5 spell_power points (0.87 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.46 DPS) [crafted]; Reinforced Woolen Shoulders (4315, -0.74 DPS, sim-verified) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 spell_power points (0.41 DPS) | yes | Pearl-clasped Cloak (5542, +0.00 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 9.9 spell_power points (1.02 DPS) | yes | Wisdom's Leather Armor (252493, -0.31 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.31 DPS) [crafted]; Totemic Leather Armor (252435, -0.70 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (25.5 DPS) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.34 DPS, sim-verified) [quest] |
| hands | Serpent Gloves (5970) | Wailing Caverns: Lord Serpentis [dungeon] | sim-verified (26.1 DPS) | yes | Stormrider's Leather Gloves (252498, -0.05 DPS) [crafted]; Gnoll Casting Gloves (892, -0.10 DPS) [world]; Fletcher's Gloves (7348, -0.92 DPS, sim-verified) [crafted] |
| waist | Pristine Sash (253925) | Tailoring [crafted] | sim-verified (25.7 DPS) | yes | Novice Arcanist's Sash (253885, -0.04 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.52 DPS, sim-verified) [crafted] |
| legs | Stormrider's Leather Pants (252502) | Leatherworking [crafted] | 12.3 spell_power points (1.26 DPS) | yes | Abomination Skin Leggings (23173, +0.00 DPS, sim-verified) [dungeon]; Wisdom's Leather Pants (252503, -0.31 DPS) [crafted]; Totemic Leather Pants (252446, -0.39 DPS) [crafted] |
| feet | Spidersilk Boots (4320) | Tailoring [crafted] | 8.5 spell_power points (0.88 DPS) | yes | Stormrider's Leather Boots (252443, +0.00 DPS, sim-verified) [crafted]; Totemic Leather Boots (252442, -0.26 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.27 DPS) [crafted] |
| finger1 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 spell_power points (0.51 DPS) | yes | Loop of Sacrifice (281673, -0.31 DPS) [quest]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop]; Sludge-Stained Band (286535, -1.00 DPS, sim-verified) [world] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | sim-verified (25.7 DPS) | yes | Loop of Sacrifice (281673, -0.04 DPS) [quest]; Volcanic Rock Ring (12053, -0.12 DPS) [world_drop]; Sludge-Stained Band (286535, -0.54 DPS, sim-verified) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) | World drop [world_drop] | sim-verified (25.8 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS) [quest]; Channeler's Staff (4437, -0.08 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.68 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Wisdom's Leather Hood; neck: Scout's Medallion; shoulder: Magician's Mantle; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Serpent Gloves; waist: Pristine Sash; legs: Stormrider's Leather Pants; feet: Spidersilk Boots; finger1: Advisor's Ring; finger2: Lavishly Jeweled Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff; ranged: Idol of the Huntress

No-known-source sample (15 of 170, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (tauren, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 49.3. Weights run: 1.4s. Verify run: 0.9s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.428 ± 0.116), crit=0.886 ± 0.077, hit=1.753 ± 0.028, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Helm (252456) | Leatherworking [crafted] | 12.0 spell_power points (1.61 DPS) | yes | Holy Shroud (2721, +0.00 DPS, sim-verified) [world_drop]; Enchanter's Cowl (4322, -0.23 DPS) [crafted]; Silk Headband (7050, -0.40 DPS) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 9.6 spell_power points (1.28 DPS) | yes | Crystal Starfire Medallion (5003, -1.05 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.05 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.27 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 12.9 spell_power points (1.72 DPS) | yes | Death Speaker Mantle (6685, -0.36 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Invoker's Mantle (215365, -0.50 DPS) [crafted] |
| back | Hillman's Cloak (3719) (or Windsong Drape (15468)) | Leatherworking [crafted] | 5.0 spell_power points (0.67 DPS) | yes | Windsong Drape (15468, +0.00 DPS, sim-verified) [quest]; Heavy Woolen Cloak (4311, -0.13 DPS) [crafted]; Battle Healer's Cloak (19529, -0.13 DPS) [rep] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 14.6 spell_power points (1.95 DPS) | yes | Tree Bark Jacket (1486, -0.21 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.27 DPS) [crafted]; Guardian Armor (4256, -1.18 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.20 DPS) | yes | Glowing Magical Bracelets (13106, -0.72 DPS, sim-verified) [world_drop]; Nightsky Wristbands (6407, -0.86 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.86 DPS) [quest] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | sim-verified (49.3 DPS) | yes | Jutebraid Gloves (10654, -0.23 DPS) [quest]; Serpent Gloves (5970, -0.38 DPS) [dungeon]; Fletcher's Gloves (7348, -1.42 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 12.6 spell_power points (1.68 DPS) | yes | Moss Cinch (6911, -0.08 DPS) [dungeon]; Warsong Sash (16975, -0.21 DPS) [quest]; Defiler's Cloth Girdle (20164, -1.25 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 14.7 spell_power points (1.97 DPS) | yes | Stormrider's Leather Pants (252502, -0.29 DPS) [crafted]; Abomination Skin Leggings (23173, -0.31 DPS) [dungeon]; Dark Ritual Leggings (270031, -0.50 DPS, sim-verified) [quest] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 10.0 spell_power points (1.34 DPS) | yes | Acidic Walkers (9454, -0.21 DPS) [dungeon]; Stormrider's Leather Boots (252443, -0.25 DPS) [crafted]; Spidersilk Boots (4320, -2.50 DPS, sim-verified) [crafted] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.94 DPS) | yes | Advisor's Ring (20426, -0.27 DPS) [rep]; Electrocutioner Lagnut (9447, -0.54 DPS) [dungeon]; Sludge-Stained Band (286535, -0.54 DPS) [world] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.80 DPS) | yes | Sludge-Stained Band (286535, -0.40 DPS) [world]; Black Widow Band (6199, -0.40 DPS) [world]; Electrocutioner Lagnut (9447, -0.71 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (47.9 DPS) | yes | Defiler's Talisman (21120, -2.88 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rhahk'Zor's Hammer (5187) | Westfall: Rhahk'Zor [dungeon] | sim-verified (47.9 DPS) | yes | Glimmering Staff (249392, -0.44 DPS) [crafted]; Twisted Chanter's Staff (890, -0.50 DPS) [world_drop]; Manual Crowd Pummeler (9449, -1.28 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Totemic Leather Helm; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Hillman's Cloak; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Rhahk'Zor's Hammer

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 70.0. Weights run: 1.4s. Verify run: 0.9s. 421 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.703 ± 0.150, crit=1.236 ± 0.144, hit=2.733 ± 0.044, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.68 DPS) | yes | Big Voodoo Mask (8201, -0.08 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.38 DPS) [world]; Corpseshroud (10574, -0.98 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.2 spell_power points (1.43 DPS) | yes | Necklace of Calisea (1714, -0.80 DPS) [world_drop]; Triune Amulet (7722, -0.80 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.29 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.1 spell_power points (2.06 DPS) | yes | Bloodmage Mantle (7684, -0.10 DPS) [dungeon]; Green Silken Shoulders (7057, -0.24 DPS, sim-verified) [crafted]; Berylline Pads (4197, -0.27 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.5 spell_power points (1.21 DPS) | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.23 DPS) [vendor]; Icy Cloak (4327, -0.32 DPS) [crafted] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (70.0 DPS) | yes | Robe of Power (7054, -0.24 DPS) [crafted]; Elemental Raiment (9434, -0.42 DPS) [world_drop]; Robe of the Magi (1716, -1.20 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.2 spell_power points (1.30 DPS) | yes | Radiant Silver Bracers (4545, -0.01 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.16 DPS) [dungeon]; Condor Bracers (15864, -0.41 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Dreamweave Gloves (10019, -0.74 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.76 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.85 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.5 spell_power points (2.24 DPS) | yes | Skycaller's Leather Belt (252522, -0.42 DPS) [crafted]; Gilded Cord (254037, -0.50 DPS) [crafted]; Defiler's Cloth Girdle (20166, -1.10 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.4 spell_power points (2.86 DPS) | yes | Crimson Silk Pantaloons (7062, -0.68 DPS) [crafted]; Kodohide Legguards (285338, -0.69 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.00 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Skycaller's Leather Shoes (252532, -1.05 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.41 DPS) [crafted]; Gilded Slippers (254001, -1.54 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.2 spell_power points (1.81 DPS) | yes | Reedknot Ring (9622, -0.92 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.05 DPS) [vendor]; Ogremind Ring (1993, -1.19 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.15 DPS) | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Reedknot Ring (9622, -1.05 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (68.8 DPS) | yes | Rune of Duty (21567, -1.82 DPS, sim-verified) [rep] |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | sim-verified (68.8 DPS) | yes | Mograine's Might (7723, -1.31 DPS) [dungeon]; Windweaver Staff (7757, -1.40 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.57 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Illusionary Rod

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 84.5. Weights run: 1.4s. Verify run: 1.0s. 563 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.737 ± 0.179, crit=1.945 ± 0.234, hit=4.310 ± 0.074, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (78.1 DPS) | yes | Soothsayer's Headdress (17740, -0.82 DPS) [dungeon]; Red Mageweave Headband (10033, -0.86 DPS) [crafted]; Blood Guard's Leather Headband (220851, -3.84 DPS, sim-verified) [vendor] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.4 spell_power points (1.31 DPS) | yes | Mindburst Medallion (11196, +0.00 DPS, sim-verified) [quest]; Horizon Choker (13085, -0.13 DPS) [world_drop]; Gemshard Heart (17707, -0.46 DPS) [dungeon] |
| shoulder | Blood Guard's Crackling Leather Spaulders (220871) | Lady Palanseer [vendor] | sim-verified (77.3 DPS) | yes | Blood Guard's Leather Shoulders (220853, -3.11 DPS, sim-verified) [vendor]; Ironfeather Shoulders (15067, -3.75 DPS) [crafted]; Rotgrip Mantle (17732, -4.15 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (75.1 DPS) | yes | Runecloth Cloak (13860, -0.40 DPS) [crafted]; Big Voodoo Cloak (8216, -0.78 DPS) [crafted]; Deep Woodlands Cloak (19121, -0.91 DPS, sim-verified) [quest] |
| chest | Stone Guard's Crackling Leather Tunic (220869) | Lady Palanseer [vendor] | sim-verified (77.2 DPS) | yes | Acumen Robes (17775, -1.61 DPS) [quest]; Robe of the Magi (1716, -2.45 DPS) [world_drop]; Stone Guard's Leather Armor (220855, -2.98 DPS, sim-verified) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 17.4 spell_power points (1.99 DPS) | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Nethergeld Cuffs (254061, -0.60 DPS) [crafted]; Bloodband Bracers (11469, -0.66 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 28.6 spell_power points (3.27 DPS) | yes | Fletcher's Gloves (7348, -0.16 DPS) [crafted]; Shadowskin Gloves (18238, -0.16 DPS) [crafted]; Gloves of Holy Might (867, -6.11 DPS, sim-verified) [world_drop] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 39.9 spell_power points (4.57 DPS) | yes | Defiler's Lizardhide Girdle (20174, -0.69 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20193, -1.45 DPS) [rep]; Skycaller's Leather Waistguard (252476, -1.73 DPS) [crafted] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | Lady Palanseer [vendor] | 89.7 spell_power points (10.27 DPS) | yes | Stone Guard's Leather Pants (220859, -3.17 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -4.04 DPS) [crafted]; Stone Guard's Restored Leather Leggings (220883, -6.31 DPS) [vendor] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | sim-verified (75.1 DPS) | yes | Skycaller's Leather Boots (252471, -0.10 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.67 DPS) [crafted]; First Sergeant's Crackling Leather Boots (220863, -0.85 DPS, sim-verified) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.1 spell_power points (4.94 DPS) | yes | Band of the Unicorn (7553, -3.45 DPS) [world_drop]; Advisor's Ring (19519, -3.56 DPS) [rep]; Brainlash (6440, -3.67 DPS) [dungeon] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.4 spell_power points (1.65 DPS) | yes | Advisor's Ring (19519, -0.28 DPS) [rep]; Band of the Unicorn (7553, -0.33 DPS, sim-verified) [world_drop]; Brainlash (6440, -0.39 DPS) [dungeon] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (74.2 DPS) | yes | Uther's Strength (11302, +0.00 DPS) [world_drop]; Frozen Heart of the Mountain (249469, -1.02 DPS, sim-verified) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (74.2 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted]; Uther's Strength (11302, -2.77 DPS) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | sim-verified (74.2 DPS) | yes | Illusionary Rod (7713, -0.51 DPS) [dungeon]; Thorium Greatmace (250613, -1.50 DPS) [crafted]; Hammer of the Northern Wind (810, -2.13 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; shoulder: Blood Guard's Crackling Leather Spaulders; back: Spritecaster Cape; chest: Stone Guard's Crackling Leather Tunic; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Defiler's Cloth Girdle; legs: Stone Guard's Crackling Leather Leggings; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Ankh of Life; trinket2: Rune of the Guard Captain; main_hand: Kindling Stave

No-known-source sample (15 of 563, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 161.8. Weights run: 1.4s. Verify run: 0.9s. 1302 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.620 ± 0.359), crit=2.513 ± 0.320, hit=5.526 ± 0.100, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 175.7 spell_power points (19.43 DPS) | yes | Bloodvine Goggles (19999, -3.32 DPS) [crafted]; Waywatcher Headpiece (240088, -6.57 DPS) [vendor]; Mask of the Unforgiven (13404, -17.15 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (161.8 DPS) | yes | Blazefury Medallion (17111, -1.80 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -2.22 DPS) [quest]; Amulet of the Dawn (22657, -3.56 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 196.4 spell_power points (21.71 DPS) | yes | Feralheart Spaulders (226778, -12.68 DPS) [quest]; Blood Guard's Leather Shoulders (220853, -13.35 DPS, sim-verified) [vendor]; Blood Guard's Crackling Leather Spaulders (220871, -13.59 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 76.2 spell_power points (8.43 DPS) | yes | Howler's Furs (272414, -2.32 DPS) [vendor]; Stalwart Cloak (272415, -2.32 DPS) [vendor]; Earthweave Cloak (21187, -4.30 DPS, sim-verified) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | sim-verified (161.8 DPS) | yes | Bloodvine Vest (19682, -3.03 DPS) [crafted]; Waywatcher Tunic (240091, -5.82 DPS) [vendor]; Tunic of Undead Slaying (23089, -15.88 DPS, sim-verified) [world] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | sim-verified (161.8 DPS) | yes | Rockfury Bracers (21186, -1.89 DPS) [quest]; Waywatcher Wristguards (240084, -4.10 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -11.91 DPS, sim-verified) [world] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 179.4 spell_power points (19.84 DPS) | yes | Stormshroud Gloves (21278, -9.84 DPS) [crafted]; Waywatcher Grips (240065, -12.77 DPS) [vendor]; Primal Batskin Gloves (19686, -14.15 DPS, sim-verified) [crafted] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 168.8 spell_power points (18.66 DPS) | yes | Knowledge of the Timbermaw (228190, -3.63 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -10.80 DPS) [vendor]; Belt of the Archmage (18405, -11.47 DPS) [crafted] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 217.5 spell_power points (24.05 DPS) | yes | Waywatcher Legguards (240087, -0.31 DPS, sim-verified) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -12.03 DPS) [vendor]; Sentinel's Silk Leggings (237815, -12.04 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 167.4 spell_power points (18.51 DPS) | yes | Bloodvine Boots (19684, -2.98 DPS, sim-verified) [crafted]; First Sergeant's Crackling Leather Boots (220863, -10.50 DPS) [vendor]; Fine Dawn Treaders (227815, -12.40 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (161.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor]; Naglering (11669, -2.95 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (161.8 DPS) | yes | Signet Ring of the Bronze Dragonflight (234032, -0.18 DPS) [vendor]; Signet Ring of the Bronze Dragonflight (234028, -0.47 DPS) [vendor]; Naglering (11669, -2.95 DPS, sim-verified) [dungeon] |
| trinket1 | Darkmoon Card: Blue Dragon (19288) | Darkmoon Beast Deck [quest] | sim-verified (161.8 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (161.8 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, -1.03 DPS, sim-verified) [vendor] |
| main_hand | Enchanted Battlehammer (12776) | Blacksmithing [crafted] | sim-verified (161.8 DPS) | yes | Frenzied Striker (13056, +0.00 DPS) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -4.12 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Waywatcher Hood; neck: Beads of Ogre Might; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Blue Dragon; trinket2: Serenity Field; main_hand: Enchanted Battlehammer; ranged: Idol of the Moon

No-known-source sample (15 of 1302, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

