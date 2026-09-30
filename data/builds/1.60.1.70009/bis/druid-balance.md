# Leveling BiS: Balance

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 27.3. Weights run: 1.3s. Verify run: 0.7s. 168 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.245 ± 0.012, crit=0.582 ± 0.038, hit=1.323 ± 0.021, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | sim-verified (26.5 DPS) | yes | Shadow Goggles (4373, -0.28 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.31 DPS) [crafted]; Trapper's Leather Hood (252505, -0.94 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.2 spell_power points (1.66 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.89 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.25 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.7 spell_power points (0.59 DPS) | yes | Sanguine Cape (14376, +0.00 DPS, sim-verified) [world_drop]; Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Feyscale Cloak (6632, -0.28 DPS) [dungeon] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (25.9 DPS) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Gray Woolen Robe (2585, -0.10 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.33 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (26.4 DPS) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.13 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.84 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 10.0 spell_power points (1.02 DPS) | yes | Blight Gloves (279877, +0.00 DPS, sim-verified) [quest]; Fletcher's Gloves (7348, -0.19 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.21 DPS) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 10.0 spell_power points (1.02 DPS) | yes | Pristine Sash (253925, -0.10 DPS) [crafted]; Stormrider's Leather Belt (252432, -0.18 DPS, sim-verified) [crafted]; Wisdom's Leather Belt (252433, -0.20 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.0 spell_power points (1.95 DPS) | yes | Stormrider's Leather Pants (252502, -0.15 DPS) [crafted]; Dreamer's Leggings (270016, -0.25 DPS, sim-verified) [quest]; Wisdom's Leather Pants (252503, -0.46 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.2 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.19 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.21 DPS) [crafted]; Black Whelp Slippers (252424, -0.44 DPS) [crafted] |
| finger1 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.5 spell_power points (0.77 DPS) | yes | Loop of Sacrifice (281673, -0.13 DPS) [quest]; Lorekeeper's Ring (20431, -0.26 DPS) [rep]; Volcanic Rock Ring (12053, -0.39 DPS) [world_drop] |
| finger2 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 7.5 spell_power points (0.77 DPS) | yes | Lorekeeper's Ring (20431, -0.25 DPS) [rep]; Loop of Sacrifice (281673, -0.36 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 12.5 spell_power points (1.28 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.26 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.46 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Keller's Girdle; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Minor Channeling Ring; finger2: Lavishly Jeweled Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 168, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14148 Crystalline Cuffs

### Band 30 (night-elf, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 53.8. Weights run: 1.3s. Verify run: 0.8s. 296 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.842 ± 0.015, crit=0.886 ± 0.077, hit=1.753 ± 0.028, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enchanter's Cowl (4322) | Tailoring [crafted] | 14.4 spell_power points (1.93 DPS) | yes | Enduring Cap (3020, +0.00 DPS, sim-verified) [world_drop]; Totemic Leather Helm (252456, -0.32 DPS) [crafted]; Holy Shroud (2721, -0.46 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.61 DPS) | yes | Crystal Starfire Medallion (5003, -1.16 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.41 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.6 spell_power points (2.22 DPS) | yes | Death Speaker Mantle (6685, -0.29 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Magician's Mantle (12998, -0.54 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.7 spell_power points (0.90 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.05 DPS) [quest]; Resilient Cape (14400, -0.23 DPS) [world_drop] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 20.0 spell_power points (2.67 DPS) | yes | Guardian Armor (4256, -0.49 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.49 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.66 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.20 DPS) | yes | Nightsky Wristbands (6407, -0.53 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.53 DPS) [quest]; Glowing Magical Bracelets (13106, -1.55 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 13.3 spell_power points (1.78 DPS) | yes | Stormrider's Leather Gloves (252498, -0.66 DPS) [crafted]; Truefaith Gloves (7049, -0.77 DPS) [crafted]; Fletcher's Gloves (7348, -2.32 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.1 spell_power points (2.02 DPS) | yes | Moss Cinch (6911, -0.41 DPS) [dungeon]; Guardian Belt (4258, -0.42 DPS) [crafted]; Highlander's Cloth Girdle (20099, -0.72 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 19.3 spell_power points (2.58 DPS) | yes | Stormrider's Leather Pants (252502, -0.56 DPS) [crafted]; Guardian Pants (5962, -0.61 DPS) [crafted]; Abomination Skin Leggings (23173, -0.68 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.9 spell_power points (1.73 DPS) | yes | Spidersilk Boots (4320, -0.34 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.36 DPS) [crafted]; Acidic Walkers (9454, -1.44 DPS, sim-verified) [dungeon] |
| finger1 | Lorekeeper's Ring (19525) | Silverwing Sentinels [rep] | 7.0 spell_power points (0.94 DPS) | yes | Black Widow Band (6199, -0.15 DPS) [world]; Snake Hoop (6750, -0.15 DPS) [quest]; Minor Channeling Ring (1449, -1.90 DPS, sim-verified) [quest] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | sim-verified (53.8 DPS) | yes | Black Widow Band (6199, -0.01 DPS) [world]; Snake Hoop (6750, -0.01 DPS) [quest]; Minor Channeling Ring (1449, -0.88 DPS, sim-verified) [quest] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Talisman of Arathor (21119, -2.43 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 0.0 spell_power points (0.00 DPS) | yes | Glimmering Staff (249392, -1.87 DPS) [crafted]; Scorn's Focal Dagger (23168, -1.90 DPS) [dungeon]; Manual Crowd Pummeler (9449, -4.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enchanter's Cowl; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Lorekeeper's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 296, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (night-elf, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 72.2. Weights run: 1.4s. Verify run: 0.9s. 416 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.750 ± 0.016, crit=1.236 ± 0.144, hit=2.733 ± 0.044, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.68 DPS) | yes | Big Voodoo Mask (8201, +0.00 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.32 DPS) [world]; Corpseshroud (10574, -0.86 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.5 spell_power points (1.47 DPS) | yes | Necklace of Calisea (1714, -0.80 DPS) [world_drop]; Triune Amulet (7722, -0.80 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.51 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.7 spell_power points (2.14 DPS) | yes | Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest]; Green Silken Shoulders (7057, -0.55 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (71.1 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.19 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -1.83 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (70.5 DPS) | yes | Robe of Power (7054, -0.22 DPS) [crafted]; Elemental Raiment (9434, -0.48 DPS) [world_drop]; Robe of the Magi (1716, -1.24 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.5 spell_power points (1.34 DPS) | yes | Spidertank Oilrag (9448, -0.19 DPS) [dungeon]; Condor Bracers (15864, -0.45 DPS) [quest]; Arcane Runed Bracers (4744, -0.96 DPS, sim-verified) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Red Mageweave Gloves (10018, -0.70 DPS) [crafted]; Dreamweave Gloves (10019, -0.78 DPS, sim-verified) [crafted]; Skycaller's Leather Gloves (252529, -0.80 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.2 spell_power points (2.33 DPS) | yes | Skycaller's Leather Belt (252522, -0.48 DPS) [crafted]; Gilded Cord (254037, -0.54 DPS) [crafted]; Highlander's Cloth Girdle (20098, -0.80 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.0 spell_power points (2.93 DPS) | yes | Crimson Silk Pantaloons (7062, -0.67 DPS) [crafted]; Abomination Skin Leggings (23173, -1.02 DPS) [dungeon]; Kodohide Legguards (285338, -1.07 DPS, sim-verified) [world] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Skycaller's Leather Shoes (252532, -0.71 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.37 DPS) [crafted]; Gilded Slippers (254001, -1.50 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.85 DPS) | yes | Ring of Forlorn Spirits (2043, -0.83 DPS) [quest]; Reedknot Ring (9622, -0.96 DPS) [quest]; Minor Channeling Ring (1449, -1.02 DPS) [quest] |
| finger2 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 spell_power points (1.15 DPS) | yes | Reedknot Ring (9622, -0.26 DPS) [quest]; Lorekeeper's Ring (19525, -0.26 DPS) [rep]; Ring of Forlorn Spirits (2043, -1.14 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.12 DPS) [dungeon]; Spellforce Rod (1664, -0.35 DPS) [world_drop]; Manual Crowd Pummeler (9449, -3.88 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Lorekeeper's Ring; trinket1: Ankh of Life

No-known-source sample (15 of 416, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 50 (night-elf, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 85.8. Weights run: 1.4s. Verify run: 0.9s. 559 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.018 ± 0.022, crit=1.945 ± 0.234, hit=4.310 ± 0.074, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (75.3 DPS) | yes | Red Mageweave Headband (10033, -0.83 DPS) [crafted]; Soothsayer's Headdress (17740, -0.95 DPS) [dungeon]; Knight-Lieutenant's Leather Headband (220850, -2.51 DPS, sim-verified) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (74.7 DPS) | yes | Scorn's Icy Choker (23169, -0.13 DPS) [dungeon]; Mindburst Medallion (11196, -0.25 DPS) [quest]; Arcane Crystal Pendant (20037, -1.91 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Crackling Leather Spaulders (220870) | Captain Dirgehammer [vendor] | sim-verified (75.3 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -2.47 DPS, sim-verified) [vendor]; Ironfeather Shoulders (15067, -3.43 DPS) [crafted]; Rotgrip Mantle (17732, -3.89 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 20.1 spell_power points (2.30 DPS) | yes | Runecloth Cloak (13860, -0.34 DPS) [crafted]; Darkspear Raider's Cloak (272076, -0.67 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -4.03 DPS, sim-verified) [dungeon] |
| chest | Knight's Crackling Leather Tunic (220868) | Captain Dirgehammer [vendor] | sim-verified (74.9 DPS) | yes | Acumen Robes (17775, -1.39 DPS) [quest]; Knight's Leather Armor (220854, -2.06 DPS, sim-verified) [vendor]; Feathered Breastplate (8349, -2.55 DPS) [crafted] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.2 spell_power points (2.31 DPS) | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -0.56 DPS) [dungeon]; Bloodband Bracers (11469, -0.69 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 32.2 spell_power points (3.69 DPS) | yes | Gloves of Holy Might (867, -0.57 DPS) [world_drop]; Fletcher's Gloves (7348, -0.57 DPS) [crafted]; Raider Handwraps (272098, -1.85 DPS, sim-verified) [vendor] |
| waist | Highlander's Cloth Girdle (20097) | The League of Arathor [rep] | 41.3 spell_power points (4.73 DPS) | yes | Highlander's Lizardhide Girdle (20103, -1.33 DPS, sim-verified) [rep]; Skycaller's Leather Waistguard (252476, -1.50 DPS) [crafted]; Ban'thok Sash (11662, -1.58 DPS) [dungeon] |
| legs | Knight's Crackling Leather Leggings (220864) | Captain Dirgehammer [vendor] | 92.5 spell_power points (10.59 DPS) | yes | Knight's Leather Pants (220858, -3.59 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -4.36 DPS) [crafted]; Knight's Restored Leather Leggings (220882, -6.31 DPS) [vendor] |
| feet | Sergeant Major's Crackling Leather Boots (220862) | Captain Dirgehammer [vendor] | 64.3 spell_power points (7.36 DPS) | yes | Skycaller's Leather Boots (252471, -0.27 DPS, sim-verified) [crafted]; Earthen Silk Slippers (254013, -4.61 DPS) [crafted]; Mender's Leather Boots (252472, -5.05 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.1 spell_power points (4.94 DPS) | yes | Cyclopean Band (11824, -1.04 DPS, sim-verified) [dungeon]; Brainlash (6440, -3.19 DPS) [dungeon]; Band of the Unicorn (7553, -3.45 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (74.5 DPS) | yes | Brainlash (6440, -0.10 DPS) [dungeon]; Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Cyclopean Band (11824, -1.74 DPS, sim-verified) [dungeon] |
| trinket1 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.09 DPS, sim-verified) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.82 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.70 DPS) [dungeon]; Glowing Brightwood Staff (812, -1.14 DPS) [world_drop]; Hammer of the Northern Wind (810, -2.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Knight-Lieutenant's Crackling Leather Spaulders; back: Spritecaster Cape; chest: Knight's Crackling Leather Tunic; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Highlander's Cloth Girdle; legs: Knight's Crackling Leather Leggings; feet: Sergeant Major's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Thunderbrew's Boot Flask; trinket2: Ankh of Life; main_hand: Kindling Stave

No-known-source sample (15 of 559, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

### Band 60 (night-elf, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 172.6. Weights run: 1.4s. Verify run: 0.9s. 1299 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.134 ± 0.031, crit=2.513 ± 0.320, hit=5.526 ± 0.100, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 185.0 spell_power points (20.45 DPS) | yes | Bloodvine Goggles (19999, -4.34 DPS) [crafted]; Waywatcher Headpiece (240088, -6.68 DPS) [vendor]; Mask of the Unforgiven (13404, -14.66 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | 0.0 spell_power points (0.00 DPS) | yes | Blazefury Medallion (17111, -1.56 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -2.22 DPS) [quest]; Amulet of the Dawn (22657, -2.82 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 203.6 spell_power points (22.51 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -12.53 DPS, sim-verified) [vendor]; Feralheart Spaulders (226778, -12.79 DPS) [quest]; Rugged Mantle of the Timbermaw (227808, -13.75 DPS) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 80.3 spell_power points (8.88 DPS) | yes | Howler's Furs (272414, -2.77 DPS) [vendor]; Stalwart Cloak (272415, -2.77 DPS) [vendor]; Earthweave Cloak (21187, -4.24 DPS, sim-verified) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Bloodvine Vest (19682, -2.97 DPS) [crafted]; Waywatcher Tunic (240091, -5.30 DPS) [vendor]; Tunic of Undead Slaying (23089, -15.10 DPS, sim-verified) [world] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Rockfury Bracers (21186, -2.63 DPS) [quest]; Waywatcher Wristguards (240084, -4.10 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -11.04 DPS, sim-verified) [world] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 189.2 spell_power points (20.92 DPS) | yes | Stormshroud Gloves (21278, -10.92 DPS) [crafted]; Primal Batskin Gloves (19686, -12.13 DPS, sim-verified) [crafted]; Waywatcher Grips (240065, -13.05 DPS) [vendor] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 176.5 spell_power points (19.52 DPS) | yes | Knowledge of the Timbermaw (228190, -4.96 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -10.75 DPS) [vendor]; Belt of the Archmage (18405, -11.41 DPS) [crafted] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 227.2 spell_power points (25.13 DPS) | yes | Waywatcher Legguards (240087, -0.25 DPS, sim-verified) [vendor]; Knight's Crackling Leather Leggings (220864, -12.55 DPS) [vendor]; Sentinel's Silk Leggings (237815, -12.55 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 173.1 spell_power points (19.14 DPS) | yes | Bloodvine Boots (19684, -3.60 DPS, sim-verified) [crafted]; Sergeant Major's Crackling Leather Boots (220862, -10.56 DPS) [vendor]; Fine Dawn Treaders (227815, -13.03 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Band of Earthen Might (21182, -0.33 DPS) [quest]; Signet Ring of the Bronze Dragonflight (234028, -0.35 DPS) [vendor]; Wrath of Cenarius (21190, -6.82 DPS, sim-verified) [quest] |
| finger2 | Don Julio's Band (19325) | Stormpike Guard [rep] | 0.0 spell_power points (0.00 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Wrath of Cenarius (21190, -1.15 DPS, sim-verified) [quest]; Channeler's Ring (272406, -2.13 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (+2.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Weakness Analyzer (272438, +0.00 DPS) [vendor] |
| trinket2 | Serenity Field (272439) | Pix Xizzix [vendor] | sim-verified (172.6 DPS) | yes | Briarwood Reed (12930, +0.00 DPS) [dungeon]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Darkmoon Card: Blue Dragon (19288, -1.82 DPS, sim-verified) [quest] |
| main_hand | Enchanted Battlehammer (12776) | Blacksmithing [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Frenzied Striker (13056, +0.00 DPS) [world_drop]; Grand Marshal's Stave (234571, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -3.69 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Waywatcher Hood; neck: Beads of Ogre Might; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Talisman of Ascendance; trinket2: Serenity Field; main_hand: Enchanted Battlehammer; ranged: Idol of the Moon

No-known-source sample (15 of 1299, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe

## Horde

### Band 20 (tauren, 5222000000000000-0000000000000000000-0000000000000000)

Set DPS (verified): 27.4. Weights run: 1.3s. Verify run: 0.7s. 170 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.245 ± 0.012, crit=0.582 ± 0.038, hit=1.323 ± 0.021, spell_haste=-1.146 ± 0.145, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.007 ± 0.000, arcane_power=0.993 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Totemic Leather Hood (252448) | Leatherworking [crafted] | sim-verified (25.9 DPS) | yes | Shadow Goggles (4373, -0.28 DPS) [crafted]; Wisdom's Leather Hood (252507, -0.31 DPS) [crafted]; Trapper's Leather Hood (252505, -1.25 DPS, sim-verified) [crafted] |
| neck | - | - |  |  |  |
| shoulder | Magician's Mantle (12998) | World drop [world_drop] | 16.2 spell_power points (1.66 DPS) | yes | Reinforced Woolen Shoulders (4315, -0.04 DPS, sim-verified) [crafted]; Double-Stitched Woolen Shoulders (4314, -1.25 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.7 spell_power points (0.59 DPS) | yes | Heavy Woolen Cloak (4311, -0.18 DPS) [crafted]; Sanguine Cape (14376, -0.23 DPS, sim-verified) [world_drop]; Feyscale Cloak (6632, -0.28 DPS) [dungeon] |
| chest | Wisdom's Leather Armor (252493) | Leatherworking [crafted] | sim-verified (24.9 DPS) | yes | Filigreed Pristine Gown (253901, +0.00 DPS) [crafted]; Gray Woolen Robe (2585, -0.10 DPS) [crafted]; Stormrider's Leather Armor (252492, -0.33 DPS, sim-verified) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | sim-verified (25.0 DPS) | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Featherbead Bracers (15452, +0.00 DPS) [quest]; Tabitha's Cuffs (251486, -0.41 DPS, sim-verified) [quest] |
| hands | Blight Gloves (279877) | The New Plague [quest] | sim-verified (25.0 DPS) | yes | Fletcher's Gloves (7348, -0.06 DPS) [crafted]; Wisdom's Leather Gloves (252499, -0.08 DPS) [crafted]; Stormrider's Leather Gloves (252498, -0.36 DPS, sim-verified) [crafted] |
| waist | Stormrider's Leather Belt (252432) | Leatherworking [crafted] | sim-verified (25.3 DPS) | yes | Pristine Sash (253925, +0.00 DPS) [crafted]; Wisdom's Leather Belt (252433, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.70 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.0 spell_power points (1.95 DPS) | yes | Stormrider's Leather Pants (252502, +0.00 DPS, sim-verified) [crafted]; Wisdom's Leather Pants (252503, -0.46 DPS) [crafted]; Filigreed Pristine Leggings (253937, -0.56 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.2 spell_power points (1.25 DPS) | yes | Spidersilk Boots (4320, -0.01 DPS, sim-verified) [crafted]; Wisdom's Leather Boots (252444, -0.21 DPS) [crafted]; Black Whelp Slippers (252424, -0.44 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | Westfall: Gilnid [dungeon] | 7.5 spell_power points (0.77 DPS) | yes | Loop of Sacrifice (281673, -0.09 DPS, sim-verified) [quest]; Volcanic Rock Ring (12053, -0.38 DPS) [world_drop]; Sludge-Stained Band (286535, -0.46 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | sim-verified (25.7 DPS) | yes | Volcanic Rock Ring (12053, -0.13 DPS) [world_drop]; Sludge-Stained Band (286535, -0.21 DPS) [world]; Loop of Sacrifice (281673, -1.11 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 12.5 spell_power points (1.28 DPS) | yes | Gnarled Necromancer's Staff (251534, +0.00 DPS, sim-verified) [quest]; Channeler's Staff (4437, -0.26 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.46 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Totemic Leather Hood; shoulder: Magician's Mantle; back: Pearl-clasped Cloak; chest: Wisdom's Leather Armor; wrist: Mindthrust Bracers; hands: Blight Gloves; waist: Stormrider's Leather Belt; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 170, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff

### Band 30 (tauren, 5222211015000000-0000000000000000000-0000000000000000)

Set DPS (verified): 51.0. Weights run: 1.3s. Verify run: 0.8s. 301 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.842 ± 0.015, crit=0.886 ± 0.077, hit=1.753 ± 0.028, spell_haste=not significant (0.055 ± 0.093), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.429 ± 0.001, arcane_power=0.571 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | sim-verified (49.2 DPS) | yes | Totemic Leather Helm (252456, -0.20 DPS) [crafted]; Holy Shroud (2721, -0.33 DPS) [world_drop]; Enchanter's Cowl (4322, -0.96 DPS, sim-verified) [crafted] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 12.1 spell_power points (1.61 DPS) | yes | Crystal Starfire Medallion (5003, -1.16 DPS) [world_drop]; Kaleidoscope Chain (13084, -1.16 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.46 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 16.6 spell_power points (2.22 DPS) | yes | Death Speaker Mantle (6685, +0.00 DPS, sim-verified) [dungeon]; Fairywing Mantle (9536, -0.40 DPS) [quest]; Magician's Mantle (12998, -0.54 DPS) [world_drop] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 6.7 spell_power points (0.90 DPS) | yes | Darkspear Raider's Cloak (272078, +0.00 DPS, sim-verified) [vendor]; Resilient Cape (14400, -0.23 DPS) [world_drop]; Hillman's Cloak (3719, -0.23 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 20.0 spell_power points (2.67 DPS) | yes | Death Speaker Robes (6682, -0.49 DPS) [dungeon]; Guardian Armor (4256, -0.53 DPS, sim-verified) [crafted]; Stormrider's Leather Tunic (252510, -0.66 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 spell_power points (1.20 DPS) | yes | Nightsky Wristbands (6407, -0.53 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.53 DPS) [quest]; Glowing Magical Bracelets (13106, -0.88 DPS, sim-verified) [world_drop] |
| hands | Oilrag Handwraps (16741) | The Lost Pages [quest] | sim-verified (50.3 DPS) | yes | Jutebraid Gloves (10654, -0.06 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.31 DPS) [crafted]; Fletcher's Gloves (7348, -2.07 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 15.1 spell_power points (2.02 DPS) | yes | Moss Cinch (6911, -0.41 DPS) [dungeon]; Guardian Belt (4258, -0.42 DPS) [crafted]; Defiler's Cloth Girdle (20164, -0.51 DPS, sim-verified) [rep] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 19.3 spell_power points (2.58 DPS) | yes | Stormrider's Leather Pants (252502, -0.56 DPS) [crafted]; Guardian Pants (5962, -0.61 DPS) [crafted]; Abomination Skin Leggings (23173, -0.71 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 12.9 spell_power points (1.73 DPS) | yes | Spidersilk Boots (4320, -0.34 DPS) [crafted]; Stormrider's Leather Boots (252443, -0.36 DPS) [crafted]; Acidic Walkers (9454, -0.98 DPS, sim-verified) [dungeon] |
| finger1 | Advisor's Ring (19521) | Warsong Outriders [rep] | 7.0 spell_power points (0.94 DPS) | yes | Black Widow Band (6199, -0.15 DPS) [world]; Snake Hoop (6750, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon] |
| finger2 | Sea Giant's Toe Ring (274746) | Gezzy Gunkgear [vendor] | 6.0 spell_power points (0.80 DPS) | yes | Snake Hoop (6750, -0.01 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.13 DPS) [dungeon]; Black Widow Band (6199, -0.80 DPS, sim-verified) [world] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 spell_power points (0.00 DPS) | yes | Defiler's Talisman (21120, -2.27 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 0.0 spell_power points (0.00 DPS) | yes | Scorn's Focal Dagger (23168, -0.04 DPS) [dungeon]; Twisted Chanter's Staff (890, -0.11 DPS) [world_drop]; Manual Crowd Pummeler (9449, -0.85 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Oilrag Handwraps; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Advisor's Ring; finger2: Sea Giant's Toe Ring; trinket1: Darkspear Voodoo Seal; main_hand: Glimmering Staff

No-known-source sample (15 of 301, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 5222211015401050-0000000000000000000-0000000000000000)

Set DPS (verified): 70.0. Weights run: 1.4s. Verify run: 0.9s. 421 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=0.750 ± 0.016, crit=1.236 ± 0.144, hit=2.733 ± 0.044, spell_haste=not significant (0.153 ± 0.061), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.483 ± 0.001, arcane_power=0.517 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 spell_power points (2.68 DPS) | yes | Big Voodoo Mask (8201, +0.00 DPS, sim-verified) [crafted]; Augural Shroud (2620, -0.32 DPS) [world]; Corpseshroud (10574, -0.86 DPS) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.5 spell_power points (1.47 DPS) | yes | Necklace of Calisea (1714, -0.80 DPS) [world_drop]; Triune Amulet (7722, -0.80 DPS) [dungeon]; Prodigious Shadowshard Pendant (17773, -2.22 DPS, sim-verified) [quest] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 16.7 spell_power points (2.14 DPS) | yes | Green Silken Shoulders (7057, -0.07 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.13 DPS) [dungeon]; Berylline Pads (4197, -0.29 DPS) [quest] |
| back | Long Silken Cloak (4326) | Tailoring [crafted] | sim-verified (68.8 DPS) | yes | Guardian Cloak (5965, +0.00 DPS) [crafted]; Darkspear Raider's Cloak (272077, -0.19 DPS) [vendor]; Mantle of Lady Falther'ess (23178, -2.04 DPS, sim-verified) [dungeon] |
| chest | Dreamweave Vest (10021) | Tailoring [crafted] | sim-verified (67.8 DPS) | yes | Robe of Power (7054, -0.22 DPS) [crafted]; Elemental Raiment (9434, -0.48 DPS) [world_drop]; Robe of the Magi (1716, -1.01 DPS, sim-verified) [world_drop] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.5 spell_power points (1.34 DPS) | yes | Radiant Silver Bracers (4545, +0.00 DPS, sim-verified) [quest]; Spidertank Oilrag (9448, -0.19 DPS) [dungeon]; Condor Bracers (15864, -0.45 DPS) [quest] |
| hands | Gloves of the Greatfather (17721) | Leatherworking [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Dreamweave Gloves (10019, -0.66 DPS, sim-verified) [crafted]; Red Mageweave Gloves (10018, -0.70 DPS) [crafted]; Skycaller's Leather Gloves (252529, -0.80 DPS) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 18.2 spell_power points (2.33 DPS) | yes | Defiler's Cloth Girdle (20166, -0.44 DPS, sim-verified) [rep]; Skycaller's Leather Belt (252522, -0.48 DPS) [crafted]; Gilded Cord (254037, -0.54 DPS) [crafted] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 23.0 spell_power points (2.93 DPS) | yes | Crimson Silk Pantaloons (7062, -0.67 DPS) [crafted]; Kodohide Legguards (285338, -0.73 DPS, sim-verified) [world]; Abomination Skin Leggings (23173, -1.02 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 spell_power points (3.06 DPS) | yes | Skycaller's Leather Shoes (252532, -0.78 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.37 DPS) [crafted]; Gilded Slippers (254001, -1.50 DPS) [crafted] |
| finger1 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | 14.5 spell_power points (1.85 DPS) | yes | Reedknot Ring (9622, -0.96 DPS) [quest]; Sea Giant's Toe Ring (274746, -1.08 DPS) [vendor]; Ogremind Ring (1993, -1.18 DPS) [world_drop] |
| finger2 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 spell_power points (1.15 DPS) | yes | Advisor's Ring (19521, -0.26 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.38 DPS) [vendor]; Reedknot Ring (9622, -1.33 DPS, sim-verified) [quest] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (+1.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Illusionary Rod (7713) | Scarlet Monastery: Arcanist Doan [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Spellforce Rod (1664, -0.23 DPS) [world_drop]; Mograine's Might (7723, -1.25 DPS) [dungeon]; Manual Crowd Pummeler (9449, -2.28 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Dreamweave Vest; wrist: Guardian Leather Bracers; hands: Gloves of the Greatfather; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Philanthropist's Ring; finger2: Advisor's Ring; trinket1: Ankh of Life; main_hand: Illusionary Rod

No-known-source sample (15 of 421, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 5222211015401051-0000000000000000000-5400000000000000)

Set DPS (verified): 85.2. Weights run: 1.4s. Verify run: 1.0s. 564 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.018 ± 0.022, crit=1.945 ± 0.234, hit=4.310 ± 0.074, spell_haste=not significant (0.344 ± 0.118), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.548 ± 0.002, arcane_power=0.452 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | sim-verified (73.3 DPS) | yes | Red Mageweave Headband (10033, -0.83 DPS) [crafted]; Soothsayer's Headdress (17740, -0.95 DPS) [dungeon]; Blood Guard's Leather Headband (220851, -3.51 DPS, sim-verified) [vendor] |
| neck | Horizon Choker (13085) | World drop [world_drop] | sim-verified (72.2 DPS) | yes | Scorn's Icy Choker (23169, -0.13 DPS) [dungeon]; Mindburst Medallion (11196, -0.25 DPS) [quest]; Arcane Crystal Pendant (20037, -2.47 DPS, sim-verified) [quest] |
| shoulder | Blood Guard's Crackling Leather Spaulders (220871) | Lady Palanseer [vendor] | sim-verified (73.1 DPS) | yes | Blood Guard's Leather Shoulders (220853, -3.38 DPS, sim-verified) [vendor]; Ironfeather Shoulders (15067, -3.43 DPS) [crafted]; Rotgrip Mantle (17732, -3.89 DPS) [dungeon] |
| back | Spritecaster Cape (11623) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | sim-verified (71.1 DPS) | yes | Mantle of Lady Falther'ess (23178, -0.22 DPS) [dungeon]; Runecloth Cloak (13860, -0.34 DPS) [crafted]; Deep Woodlands Cloak (19121, -1.33 DPS, sim-verified) [quest] |
| chest | Stone Guard's Crackling Leather Tunic (220869) | Lady Palanseer [vendor] | sim-verified (72.5 DPS) | yes | Acumen Robes (17775, -1.39 DPS) [quest]; Feathered Breastplate (8349, -2.55 DPS) [crafted]; Stone Guard's Leather Armor (220855, -2.80 DPS, sim-verified) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 20.2 spell_power points (2.31 DPS) | yes | Skycaller's Leather Bracers (252542, +0.00 DPS, sim-verified) [crafted]; Aristocratic Cuffs (12546, -0.56 DPS) [dungeon]; Bloodband Bracers (11469, -0.69 DPS) [quest] |
| hands | Feralheart Hands (226777) | Mokvar [vendor] | 32.2 spell_power points (3.69 DPS) | yes | Gloves of Holy Might (867, -0.57 DPS) [world_drop]; Fletcher's Gloves (7348, -0.57 DPS) [crafted]; Raider Handwraps (272098, -1.20 DPS, sim-verified) [vendor] |
| waist | Defiler's Cloth Girdle (20165) | The Defilers [rep] | 41.3 spell_power points (4.73 DPS) | yes | Defiler's Lizardhide Girdle (20174, -0.26 DPS, sim-verified) [rep]; Skycaller's Leather Waistguard (252476, -1.50 DPS) [crafted]; Ban'thok Sash (11662, -1.58 DPS) [dungeon] |
| legs | Stone Guard's Crackling Leather Leggings (220865) | Lady Palanseer [vendor] | 92.5 spell_power points (10.59 DPS) | yes | Stone Guard's Leather Pants (220859, -3.04 DPS, sim-verified) [vendor]; Stormshroud Pants (15057, -4.36 DPS) [crafted]; Stone Guard's Restored Leather Leggings (220883, -6.31 DPS) [vendor] |
| feet | First Sergeant's Crackling Leather Boots (220863) | Lady Palanseer [vendor] | 64.3 spell_power points (7.36 DPS) | yes | Skycaller's Leather Boots (252471, +0.00 DPS, sim-verified) [crafted]; Earthen Silk Slippers (254013, -4.61 DPS) [crafted]; Mender's Leather Boots (252472, -5.05 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 43.1 spell_power points (4.94 DPS) | yes | Cyclopean Band (11824, -0.66 DPS, sim-verified) [dungeon]; Brainlash (6440, -3.19 DPS) [dungeon]; Band of the Unicorn (7553, -3.45 DPS) [world_drop] |
| finger2 | Philanthropist's Ring (281635) | Greater Friend of the Library [quest] | sim-verified (72.4 DPS) | yes | Brainlash (6440, -0.10 DPS) [dungeon]; Band of the Unicorn (7553, -0.36 DPS) [world_drop]; Cyclopean Band (11824, -2.61 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS, sim-verified) [crafted]; Uther's Strength (11302, -2.77 DPS) [world_drop] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Uther's Strength (11302, -0.81 DPS, sim-verified) [world_drop] |
| main_hand | Kindling Stave (11750) | Blackrock Depths: Pyromancer Loregrain [dungeon] | 0.0 spell_power points (0.00 DPS) | yes | Illusionary Rod (7713, -0.70 DPS) [dungeon]; Glowing Brightwood Staff (812, -1.14 DPS) [world_drop]; Hammer of the Northern Wind (810, -1.75 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Blood Guard's Crackling Leather Spaulders; back: Spritecaster Cape; chest: Stone Guard's Crackling Leather Tunic; wrist: Runic Leather Bracers; hands: Feralheart Hands; waist: Defiler's Cloth Girdle; legs: Stone Guard's Crackling Leather Leggings; feet: First Sergeant's Crackling Leather Boots; finger1: Blackstone Ring; finger2: Philanthropist's Ring; trinket1: Rune of the Guard Captain; trinket2: Ankh of Life; main_hand: Kindling Stave

No-known-source sample (15 of 564, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 5222211015401051-0000000000000000000-5533300000000000)

Set DPS (verified): 170.4. Weights run: 1.4s. Verify run: 0.9s. 1303 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=1.134 ± 0.031, crit=2.513 ± 0.320, hit=5.526 ± 0.100, spell_haste=not significant (0.737 ± 0.247), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.562 ± 0.002, arcane_power=0.438 ± 0.004

| Slot | Item | Source | Score (spell_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Waywatcher Hood (240072) | Leonid Barthalomew the Revered [vendor] | 185.0 spell_power points (20.45 DPS) | yes | Bloodvine Goggles (19999, -4.34 DPS) [crafted]; Waywatcher Headpiece (240088, -6.68 DPS) [vendor]; Mask of the Unforgiven (13404, -16.77 DPS, sim-verified) [dungeon] |
| neck | Beads of Ogre Might (22150) | Falrin's Vendetta [quest] | sim-verified (170.4 DPS) | yes | Blazefury Medallion (17111, -1.64 DPS, sim-verified) [world]; Medallion of the Dawn (22659, -2.22 DPS) [quest]; Amulet of the Dawn (22657, -2.82 DPS) [quest] |
| shoulder | Waywatcher Mantle (240070) | Leonid Barthalomew the Revered [vendor] | 203.6 spell_power points (22.51 DPS) | yes | Feralheart Spaulders (226778, -12.79 DPS) [quest]; Rugged Mantle of the Timbermaw (227808, -13.75 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, -14.93 DPS, sim-verified) [vendor] |
| back | Arcanoweave Cloak (272411) | Pix Xizzix [vendor] | 80.3 spell_power points (8.88 DPS) | yes | Howler's Furs (272414, -2.77 DPS) [vendor]; Stalwart Cloak (272415, -2.77 DPS) [vendor]; Earthweave Cloak (21187, -4.76 DPS, sim-verified) [quest] |
| chest | Waywatcher Leathers (240075) | Leonid Barthalomew the Revered [vendor] | sim-verified (170.4 DPS) | yes | Bloodvine Vest (19682, -2.97 DPS) [crafted]; Waywatcher Tunic (240091, -5.30 DPS) [vendor]; Tunic of Undead Slaying (23089, -17.88 DPS, sim-verified) [world] |
| wrist | Waywatcher Bindings (240068) | Leonid Barthalomew the Revered [vendor] | sim-verified (170.4 DPS) | yes | Rockfury Bracers (21186, -2.63 DPS) [quest]; Waywatcher Wristguards (240084, -4.10 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -13.69 DPS, sim-verified) [world] |
| hands | Waywatcher Mitts (240073) | Leonid Barthalomew the Revered [vendor] | 189.2 spell_power points (20.92 DPS) | yes | Stormshroud Gloves (21278, -10.92 DPS) [crafted]; Waywatcher Grips (240065, -13.05 DPS) [vendor]; Primal Batskin Gloves (19686, -13.48 DPS, sim-verified) [crafted] |
| waist | Waywatcher Cord (240069) | Leonid Barthalomew the Revered [vendor] | 176.5 spell_power points (19.52 DPS) | yes | Knowledge of the Timbermaw (228190, -6.39 DPS, sim-verified) [vendor]; Waywatcher Girdle (240085, -10.75 DPS) [vendor]; Belt of the Archmage (18405, -11.41 DPS) [crafted] |
| legs | Waywatcher Kilt (240071) | Leonid Barthalomew the Revered [vendor] | 227.2 spell_power points (25.13 DPS) | yes | Waywatcher Legguards (240087, -1.35 DPS, sim-verified) [vendor]; Stone Guard's Crackling Leather Leggings (220865, -12.55 DPS) [vendor]; Sentinel's Silk Leggings (237815, -12.55 DPS) [vendor] |
| feet | Waywatcher Sandals (240074) | Leonid Barthalomew the Revered [vendor] | 173.1 spell_power points (19.14 DPS) | yes | Bloodvine Boots (19684, -5.15 DPS, sim-verified) [crafted]; First Sergeant's Crackling Leather Boots (220863, -10.56 DPS) [vendor]; Fine Dawn Treaders (227815, -13.03 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234032) | Anachronos [vendor] | sim-verified (170.4 DPS) | yes | Band of Earthen Might (21182, -0.33 DPS) [quest]; Signet Ring of the Bronze Dragonflight (234028, -0.35 DPS) [vendor]; Naglering (11669, -9.73 DPS, sim-verified) [dungeon] |
| finger2 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (170.4 DPS) | yes | Band of Earthen Might (21182, +0.00 DPS) [quest]; Wrath of Cenarius (21190, -1.85 DPS, sim-verified) [quest]; Channeler's Ring (272406, -2.13 DPS) [vendor] |
| trinket1 | Talisman of Ascendance (22678) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (170.4 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Serenity Field (272439, -1.42 DPS, sim-verified) [vendor] |
| trinket2 | Darkmoon Card: Blue Dragon (19288) | Darkmoon Beast Deck [quest] | sim-verified (170.4 DPS) | yes | Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Serenity Field (272439, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Enchanted Battlehammer (12776) | Blacksmithing [crafted] | sim-verified (170.4 DPS) | yes | Frenzied Striker (13056, +0.00 DPS) [world_drop]; High Warlord's War Staff (234549, +0.00 DPS) [vendor]; Hand of Edward the Odd (2243, -3.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | Idol of the Moon (23197) (or Howling Idol (272427), Enraged Idol (272428), Idol of Synthesis (272429), Swarming Idol (272430), Idol of Swiftness (279250), Idol of the Ursine Twins (279251), Idol of the Dream (220606), Idol of the Huntress (227444), Idol of the Raging Shambler (220915), Talons of Wrath (249441), Idol of the Heckler (213594), Idol of the Wild (210534), Mystic Mushroom (249396), Ferocious Idol (208689), Idol of Ursine Rage (206954), Unbalanced Idol (210195), Lunar Idol (208414)) | World drop [world_drop] | 0.0 spell_power points (0.00 DPS) | yes | Howling Idol (272427, +0.00 DPS, sim-verified) [vendor] |

**New at 60:** head: Waywatcher Hood; neck: Beads of Ogre Might; shoulder: Waywatcher Mantle; back: Arcanoweave Cloak; chest: Waywatcher Leathers; wrist: Waywatcher Bindings; hands: Waywatcher Mitts; waist: Waywatcher Cord; legs: Waywatcher Kilt; feet: Waywatcher Sandals; finger1: Signet Ring of the Bronze Dragonflight; finger2: Don Julio's Band; trinket1: Talisman of Ascendance; trinket2: Darkmoon Card: Blue Dragon; main_hand: Enchanted Battlehammer; ranged: Idol of the Moon

No-known-source sample (15 of 1303, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

