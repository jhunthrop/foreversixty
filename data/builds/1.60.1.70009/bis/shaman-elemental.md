# Leveling BiS: Elemental

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

## Alliance

### Band 20 (dwarf, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 32.2. Weights run: 1.6s. Verify run: 1.0s. 485 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.268 ± 0.163, crit=0.822 ± 0.037, hit=1.874 ± 0.118, spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 11.0 | yes | Totemic Leather Hood (252448, -0.19 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.38 DPS) [crafted]; Trapper's Leather Hood (252505, -1.14 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) (or Tarnished Locket (279870)) | Silverwing Sentinels [rep] | 0.0 | yes | Tarnished Locket (279870, +0.16 DPS, sim-verified) [quest] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 10.1 | yes | Double-Stitched Woolen Shoulders (4314, -0.13 DPS, sim-verified) [crafted]; Rough Bronze Shoulders (3480, -0.97 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.97 DPS) [crafted] |
| back | Heavy Woolen Cloak (4311) | Tailoring [crafted] | 4.0 | yes | Feyscale Cloak (6632, -0.10 DPS) [dungeon]; Black Whelp Cloak (7283, -0.10 DPS) [crafted]; Pearl-clasped Cloak (5542, -0.68 DPS, sim-verified) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 14.3 | yes | Acolyte's Chain Shirt (250491, +0.15 DPS, sim-verified) [crafted]; Wisdom's Leather Armor (252493, -0.29 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.29 DPS) [crafted] |
| wrist | Mindthrust Bracers (1974) | Shadowfang Keep: Son of Arugal [dungeon] | 6.3 | yes | Owl Bracers (4796, +0.00 DPS) [vendor]; Bright Bracers (3647, -0.12 DPS) [world_drop]; Tabitha's Cuffs (251486, -0.37 DPS, sim-verified) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 10.1 | yes | Acolyte's Gloves (250511, -0.07 DPS) [crafted]; Blight Gloves (279877, -0.11 DPS) [quest]; Fletcher's Gloves (7348, -0.95 DPS, sim-verified) [crafted] |
| waist | Stormrider's Leather Belt (252432) | Leatherworking [crafted] | 9.1 | yes | Pristine Sash (253925, +0.00 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.10 DPS) [crafted]; Keller's Girdle (2911, -0.51 DPS, sim-verified) [world_drop] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.1 | yes | Stormrider's Leather Pants (252502, -0.15 DPS) [crafted]; Acolyte's Chain Leggings (250496, -0.31 DPS) [crafted]; Dreamer's Leggings (270016, -0.35 DPS, sim-verified) [quest] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.3 | yes | Spidersilk Boots (4320, +0.08 DPS, sim-verified) [crafted]; Acolyte's Boots (250506, -0.19 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.19 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.6 | yes | Lorekeeper's Ring (20431, -0.25 DPS) [rep]; Sludge-Stained Band (286535, -0.44 DPS) [world]; Black Pearl Ring (6332, -0.49 DPS) [world] |
| finger2 | Minor Channeling Ring (1449) | WANTED: Chok'sul [quest] | 7.5 | yes | Lorekeeper's Ring (20431, +0.14 DPS, sim-verified) [rep]; Sludge-Stained Band (286535, -0.43 DPS) [world]; Black Pearl Ring (6332, -0.48 DPS) [world] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 12.7 | yes | Channeler's Staff (4437, -0.24 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.45 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -1.59 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Sentinel's Medallion; shoulder: Reinforced Woolen Shoulders; back: Heavy Woolen Cloak; chest: Stormrider's Leather Armor; wrist: Mindthrust Bracers; hands: Stormrider's Leather Gloves; waist: Stormrider's Leather Belt; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Minor Channeling Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 485, see the JSON for more): 1189 Overseer's Ring; 2879 Antipodean Rod; 3738 Brewing Rod; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9753 Nomad Buckler; 9756 Nomad Trousers; 9757 Nomad Tunic

### Band 30 (dwarf, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 57.7. Weights run: 1.7s. Verify run: 1.2s. 926 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.138 ± 0.340), crit=1.742 ± 0.105, hit=2.217 ± 0.197, spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 18.2 | yes | Enchanter's Cowl (4322, +0.21 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.42 DPS) [crafted]; Nightsky Cowl (4039, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 | yes | Crystal Starfire Medallion (5003, -0.93 DPS) [world_drop]; Pendant of Myzrael (4614, -1.38 DPS) [dungeon]; Darkspear Warding Pendant (272075, -1.47 DPS, sim-verified) [vendor] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.2 | yes | Fairywing Mantle (9536, -0.30 DPS) [quest]; Death Speaker Mantle (6685, -0.38 DPS, sim-verified) [dungeon]; Feline Mantle (3748, -0.61 DPS) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.1 | yes | Darkspear Raider's Cloak (272078, +0.26 DPS, sim-verified) [vendor]; Repairman's Cape (9605, -0.15 DPS) [quest]; Pearl-clasped Cloak (5542, -0.37 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.8 | yes | Death Speaker Robes (6682, -0.43 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.70 DPS) [crafted]; Guardian Armor (4256, -0.89 DPS, sim-verified) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.22 DPS) [quest]; Technician's Bracers (270042, -0.22 DPS) [quest]; Nightsky Wristbands (6407, -1.34 DPS, sim-verified) [world_drop] |
| hands | Town Clerk's Mittens (270029) | Crime and Punishment [quest] | 16.5 | yes | Stormrider's Leather Gloves (252498, -0.70 DPS) [crafted]; Hotshot Pilot's Gloves (9491, -0.74 DPS) [dungeon]; Fletcher's Gloves (7348, -1.39 DPS, sim-verified) [crafted] |
| waist | Prefect's Belt (250559) | Blacksmithing [crafted] | 15.0 | yes | Highlander's Cloth Girdle (20099, -0.05 DPS) [rep]; Guardian Belt (4258, -0.10 DPS) [crafted]; Skycaller's Leather Belt (252522, -0.73 DPS, sim-verified) [crafted] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 22.5 | yes | Guardian Pants (5962, -0.54 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.54 DPS) [crafted]; Abomination Skin Leggings (23173, -1.18 DPS, sim-verified) [dungeon] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 | yes | Stormrider's Leather Boots (252443, -0.33 DPS) [crafted]; Spidersilk Boots (4320, -0.34 DPS) [crafted]; Acidic Walkers (9454, -1.19 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.0 | yes | Minor Channeling Ring (1449, -0.07 DPS) [quest]; Lorekeeper's Ring (19525, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 | yes | Minor Channeling Ring (1449, +0.12 DPS, sim-verified) [quest]; Lorekeeper's Ring (19525, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 24.7 | yes | Twisted Chanter's Staff (890, -1.33 DPS) [world_drop]; Gnarled Necromancer's Staff (251534, -1.33 DPS) [quest]; Glimmering Staff (249392, -3.44 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Town Clerk's Mittens; waist: Prefect's Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Talisman of Arathor; main_hand: Mechanic's Pipehammer

No-known-source sample (15 of 926, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 7958 Bronze Battle Axe

### Band 40 (dwarf, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 69.3. Weights run: 1.9s. Verify run: 1.5s. 1264 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.663 ± 0.634), crit=2.997 ± 0.195, hit=4.584 ± 0.401, spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, -0.24 DPS) [crafted]; Augural Shroud (2620, -0.30 DPS) [world]; Raging Berserker's Helm (7719, -1.33 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Darkspear Warding Pendant (272074, -0.56 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.68 DPS) [vendor]; Necklace of Calisea (1714, -0.83 DPS, sim-verified) [world_drop] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 | yes | Green Silken Shoulders (7057, +0.16 DPS, sim-verified) [crafted]; Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.18 DPS) [quest] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.18 DPS) [vendor]; Icy Cloak (4327, -0.21 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 | yes | Dreamweave Vest (10021, +0.33 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.36 DPS) [crafted]; Crimson Silk Vest (7058, -0.65 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.0 | yes | Turtle Scale Bracers (8198, +0.20 DPS, sim-verified) [crafted]; Arcane Runed Bracers (4744, -0.09 DPS) [quest]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Dragonscale Gauntlets (8347), Shadowskin Gloves (18238)) | World drop [world_drop] | 42.0 | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Fletcher's Gloves (7348, -0.42 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.0 | yes | Highlander's Cloth Girdle (20098, -0.03 DPS) [rep]; Skycaller's Leather Belt (252522, -0.26 DPS) [crafted]; Highlander's Chain Girdle (20089, -1.75 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Kodohide Legguards (285338, -0.10 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -0.47 DPS) [crafted]; Abomination Skin Leggings (23173, -0.68 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Mail Boots (252563, -0.57 DPS) [crafted]; Skycaller's Leather Shoes (252532, -0.63 DPS, sim-verified) [crafted]; Mender's Leather Shoes (252533, -1.01 DPS) [crafted] |
| finger1 | Lorekeeper's Ring (19524) | Silverwing Sentinels [rep] | 9.0 | yes | Reedknot Ring (9622, -0.18 DPS) [quest]; Lorekeeper's Ring (19525, -0.18 DPS) [rep]; Minor Channeling Ring (1449, -0.24 DPS) [quest] |
| finger2 | Ring of Forlorn Spirits (2043) | The Legend of Stalvan [quest] | 8.0 | yes | Reedknot Ring (9622, -0.12 DPS, sim-verified) [quest]; Minor Channeling Ring (1449, -0.15 DPS) [quest]; Sea Giant's Toe Ring (274746, -0.18 DPS) [vendor] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS, sim-verified) [world_drop] |
| main_hand | Mechanic's Pipehammer (9604) | Data Rescue [quest] | 22.3 | yes | Mograine's Might (7723, -1.04 DPS) [dungeon]; Shadow Crescent Axe (3856, -1.10 DPS) [crafted]; Illusionary Rod (7713, -3.79 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Lorekeeper's Ring; finger2: Ring of Forlorn Spirits; trinket1: Rune of Perfection; trinket2: Rune of Duty

No-known-source sample (15 of 1264, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 50 (dwarf, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 96.4. Weights run: 1.9s. Verify run: 1.6s. 1612 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (3.110 ± 0.789), crit=3.848 ± 0.256, hit=6.081 ± 0.559, spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Pulsing Helmet (220848) | Lady Palanseer [vendor] | 108.0 | yes | Blood Guard's Inscribed Skullcap (220842, -1.76 DPS) [vendor]; Red Mageweave Headband (10033, -2.36 DPS) [crafted]; Eye of Theradras (17715, -2.98 DPS, sim-verified) [dungeon] |
| neck | Darkspear Warding Pendant (272073) | Creeg Bothunk [vendor] | 28.0 | yes | Scorn's Icy Choker (23169, -0.21 DPS) [dungeon]; Mindburst Medallion (11196, -0.29 DPS) [quest]; Horizon Choker (13085, -1.13 DPS, sim-verified) [world] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 77.2 | yes | Blood Guard's Mail Epaulets (220823, -0.69 DPS) [vendor]; Blood Guard's Inscribed Shoulder Pads (220841, -0.69 DPS) [vendor]; Blood Guard's Pulsing Shoulders (220849, -4.97 DPS, sim-verified) [vendor] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 43.5 | yes | Imperial Red Cloak (8248, +0.58 DPS, sim-verified) [dungeon]; Darkspear Raider's Cloak (272077, -0.82 DPS) [vendor]; Runecloth Cloak (13860, -0.85 DPS) [crafted] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Armor (220826, -1.97 DPS, sim-verified) [vendor]; Stone Guard's Inscribed Chestpiece (220838, -2.20 DPS) [vendor]; Knight's Mail Armor (223078, -2.20 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 41.1 | yes | Shizzle's Nozzle Wiper (11917, -0.07 DPS, sim-verified) [quest]; Imperial Red Bracers (8247, -0.61 DPS) [dungeon]; Bloodband Bracers (11469, -0.71 DPS) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 75.4 | yes | First Sergeant's Mail Gauntlets (220831, -0.53 DPS) [vendor]; First Sergeant's Inscribed Gauntlets (220843, -0.53 DPS) [vendor]; First Sergeant's Pulsing Gauntlets (220845, -3.37 DPS, sim-verified) [vendor] |
| waist | Highlander's Lizardhide Girdle (20103) (or Highlander's Mail Girdle (20118)) | The League of Arathor [rep] | 85.0 | yes | Highlander's Mail Girdle (20118, +0.00 DPS, sim-verified) [vendor]; Highlander's Cloth Girdle (20097, -0.58 DPS) [rep]; Dawnspire Cord (12466, -1.75 DPS) [world] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Legplates (220834, -2.20 DPS) [vendor]; Stone Guard's Inscribed Legplates (220839, -2.20 DPS) [vendor]; Stormshroud Pants (15057, -3.23 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 102.9 | yes | Skycaller's Leather Boots (252471, +0.06 DPS, sim-verified) [crafted]; Skycaller's Mail Sabatons (252577, -4.73 DPS) [crafted]; Mender's Leather Boots (252472, -5.26 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 60.8 | yes | Voodoo Band (1996, -3.44 DPS) [world]; Mindbender Loop (5009, -3.44 DPS) [world_drop]; Black Widow Band (6199, -3.44 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 21.8 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 54.7 | yes | Uther's Strength (11302, -4.29 DPS) [world]; Thunderbrew's Boot Flask (744, -4.82 DPS) [quest]; Tidal Charm (1404, -4.82 DPS) [vendor] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Uther's Strength (11302, +0.53 DPS) [world]; Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 0.0 | yes | Illusionary Rod (7713, +6.39 DPS) [dungeon]; Spellshifter Rod (9527, +6.30 DPS) [quest]; Kindling Stave (11750, -2.72 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Blood Guard's Pulsing Helmet; neck: Darkspear Warding Pendant; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Highlander's Lizardhide Girdle; legs: Stone Guard's Pulsing Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Ogremind Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Hammer of the Northern Wind

No-known-source sample (15 of 1612, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 60 (dwarf, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 182.8. Weights run: 1.9s. Verify run: 1.6s. 2372 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.046 ± 0.865), crit=5.588 ± 0.380, hit=6.652 ± 0.739, spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) (or Bloodvine Goggles (19999)) | Stratholme: The Unforgiven [dungeon] | 211.3 | yes | Bloodvine Goggles (19999, -0.14 DPS, sim-verified) [crafted]; Ragefury Eyepatch (11735, -4.75 DPS) [dungeon]; Bloodvine Lens (19998, -4.75 DPS) [crafted] |
| neck | Gem of Trapped Innocents (23057) | Naxxramas: Kel'Thuzad [raid] | 178.8 | yes | Stormrage's Talisman of Seething (23053, -1.94 DPS) [raid]; Onyxia Tooth Pendant (18404, -2.95 DPS) [quest]; Fury of the Forgotten Swarm (21809, -2.95 DPS) [raid] |
| shoulder | Champion's Mail Pauldrons (227154) | Rank 14 [pvp] | 150.0 | yes | Warlord's Mail Spaulders (231659, +0.78 DPS, sim-verified) [pvp]; Mantle of the Timbermaw (19050, -3.57 DPS) [crafted]; Champion's Mail Pauldrons (23260, -4.01 DPS) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 78.2 | yes | Earthweave Cloak (21187, +0.79 DPS, sim-verified) [quest]; Hide of the Wild (18510, -4.66 DPS) [crafted]; Spritecaster Cape (11623, -5.02 DPS) [dungeon] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 173.6 | yes | Bloodsoul Breastplate (19690, -1.49 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -2.50 DPS) [pvp]; Stormshroud Armor (15056, -9.21 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 93.5 | yes | Qiraji Execution Bracers (21602, -2.34 DPS) [raid]; Primal Batskin Bracers (19687, -2.58 DPS, sim-verified) [crafted]; Burrower Bracers (21611, -4.50 DPS) [raid] |
| hands | Stormshroud Gloves (21278) (or Blood Guard's Mail Vices (227159)) | Leatherworking [crafted] | 144.7 | yes | Blood Guard's Mail Vices (227159, +1.59 DPS, sim-verified) [pvp]; Primal Batskin Gloves (19686, -1.02 DPS) [crafted]; Dark Storm Gauntlets (21585, -2.21 DPS) [raid] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 115.0 | yes | Cord of The Five Thunders (227008, -0.96 DPS) [quest]; Highlander's Cloth Girdle (20047, -1.43 DPS) [rep]; Belt of Never-ending Agony (21586, -2.61 DPS, sim-verified) [raid] |
| legs | Sentinel's Chain Leggings (22748) | Silverwing Sentinels [rep] | 144.7 | yes | Legionnaire's Mail Legguards (227156, +0.00 DPS) [pvp]; General's Mail Leggings (231664, -1.16 DPS) [pvp]; Stormshroud Pants (15057, -1.98 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 102.3 | yes | Blood Guard's Mail Greaves (227158, -3.10 DPS) [pvp]; Boots of Epiphany (21600, -4.19 DPS) [raid]; Greaves of Withering Despair (22240, -5.56 DPS, sim-verified) [dungeon] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182), Band of Unnatural Forces (23038)) | Stormpike Guard [rep] | 144.7 | yes | Band of Unnatural Forces (23038, +0.00 DPS) [raid]; Ring of the Fallen God (21709, -3.03 DPS) [quest]; Mindtear Band (20632, -3.32 DPS) [world] |
| finger2 | Band of Earthen Might (21182) (or Band of Unnatural Forces (23038)) | Veteran's Battlegear [quest] | 144.7 | yes | Band of Unnatural Forces (23038, +1.20 DPS, sim-verified) [raid]; Ring of the Fallen God (21709, -3.03 DPS) [quest]; Mindtear Band (20632, -3.32 DPS) [world] |
| trinket1 | The Restrained Essence of Sapphiron (23046) | Naxxramas: Sapphiron [raid] | 40.0 | yes | Neltharion's Tear (19379, +11.88 DPS) [raid]; Eye of Diminution (23001, +10.10 DPS) [raid]; Drake Fang Talisman (19406, +8.07 DPS) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas: Maexxna [raid] | 144.7 | yes | Neltharion's Tear (19379, +2.80 DPS) [raid]; Eye of Diminution (23001, +1.02 DPS) [raid]; Drake Fang Talisman (19406, -1.02 DPS) [raid] |
| main_hand | Atiesh, Greatstaff of the Guardian (22630) | Atiesh, Greatstaff of the Guardian [quest] | 336.8 | yes | Atiesh, Greatstaff of the Guardian (22589, -1.76 DPS) [quest]; High Warlord's War Staff (234549, -6.53 DPS) [pvp]; Grand Marshal's Stave (234571, -6.53 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Mask of the Unforgiven; neck: Gem of Trapped Innocents; shoulder: Champion's Mail Pauldrons; back: Chromatic Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Stormshroud Gloves; waist: Belt of the Archmage; legs: Sentinel's Chain Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: The Restrained Essence of Sapphiron; trinket2: Kiss of the Spider; main_hand: Atiesh, Greatstaff of the Guardian

No-known-source sample (15 of 2372, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

## Horde

### Band 20 (orc, 4520000000000000-000000000000000000-0000000000000000)

Set DPS (verified): 31.6. Weights run: 1.6s. Verify run: 1.0s. 482 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=1.268 ± 0.163, crit=0.822 ± 0.037, hit=1.874 ± 0.118, spell_haste=3.471 ± 0.217, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.699 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Crusader's Silvered Chain Helm (250532) | Blacksmithing [crafted] | 11.0 | yes | Totemic Leather Hood (252448, -0.19 DPS) [crafted]; Acolyte's Silvered Chain Helm (250531, -0.38 DPS) [crafted]; Trapper's Leather Hood (252505, -1.37 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) (or Tarnished Locket (279870)) | Warsong Outriders [rep] | 0.0 | yes | Tarnished Locket (279870, -0.02 DPS, sim-verified) [quest] |
| shoulder | Reinforced Woolen Shoulders (4315) | Tailoring [crafted] | 10.1 | yes | Double-Stitched Woolen Shoulders (4314, -0.54 DPS, sim-verified) [crafted]; Rough Bronze Shoulders (3480, -0.97 DPS) [crafted]; Silvered Bronze Shoulders (3481, -0.97 DPS) [crafted] |
| back | Pearl-clasped Cloak (5542) | Tailoring [crafted] | 5.8 | yes | Heavy Woolen Cloak (4311, +0.14 DPS, sim-verified) [crafted]; Feyscale Cloak (6632, -0.27 DPS) [dungeon]; Black Whelp Cloak (7283, -0.27 DPS) [crafted] |
| chest | Stormrider's Leather Armor (252492) | Leatherworking [crafted] | 14.3 | yes | Acolyte's Chain Shirt (250491, +0.19 DPS, sim-verified) [crafted]; Wisdom's Leather Armor (252493, -0.29 DPS) [crafted]; Filigreed Pristine Gown (253901, -0.29 DPS) [crafted] |
| wrist | Tabitha's Cuffs (251486) | A Frightened Request [quest] | 7.6 | yes | Mindthrust Bracers (1974, +0.02 DPS, sim-verified) [dungeon]; Owl Bracers (4796, -0.12 DPS) [vendor]; Featherbead Bracers (15452, -0.12 DPS) [quest] |
| hands | Stormrider's Leather Gloves (252498) | Leatherworking [crafted] | 10.1 | yes | Acolyte's Gloves (250511, -0.07 DPS) [crafted]; Blight Gloves (279877, -0.11 DPS) [quest]; Fletcher's Gloves (7348, -0.76 DPS, sim-verified) [crafted] |
| waist | Keller's Girdle (2911) | World drop [world_drop] | 10.1 | yes | Stormrider's Leather Belt (252432, +0.07 DPS, sim-verified) [crafted]; Pristine Sash (253925, -0.10 DPS) [crafted]; Acolyte's Chain Belt (250516, -0.20 DPS) [crafted] |
| legs | Abomination Skin Leggings (23173) | Shadowfang Keep: Sever [dungeon] | 19.1 | yes | Stormrider's Leather Pants (252502, +0.16 DPS, sim-verified) [crafted]; Acolyte's Chain Leggings (250496, -0.31 DPS) [crafted]; Wisdom's Leather Pants (252503, -0.43 DPS) [crafted] |
| feet | Stormrider's Leather Boots (252443) | Leatherworking [crafted] | 12.3 | yes | Spidersilk Boots (4320, +0.10 DPS, sim-verified) [crafted]; Acolyte's Boots (250506, -0.19 DPS) [crafted]; Wisdom's Leather Boots (252444, -0.19 DPS) [crafted] |
| finger1 | Lavishly Jeweled Ring (1156) | The Deadmines: Gilnid [dungeon] | 7.6 | yes | Sludge-Stained Band (286535, -0.44 DPS) [world]; Black Pearl Ring (6332, -0.49 DPS) [world]; The 1 Ring (8350, -0.61 DPS) [world] |
| finger2 | Advisor's Ring (20426) | Warsong Outriders [rep] | 5.0 | yes | Black Pearl Ring (6332, -0.24 DPS) [world]; The 1 Ring (8350, -0.36 DPS) [world]; Sludge-Stained Band (286535, -0.70 DPS, sim-verified) [world] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Twisted Chanter's Staff (890) (or Gnarled Necromancer's Staff (251534)) | World drop [world_drop] | 12.7 | yes | Channeler's Staff (4437, -0.24 DPS) [world]; Rhahk'Zor's Hammer (5187, -0.45 DPS) [dungeon]; Gnarled Necromancer's Staff (251534, -1.65 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Crusader's Silvered Chain Helm; neck: Scout's Medallion; shoulder: Reinforced Woolen Shoulders; back: Pearl-clasped Cloak; chest: Stormrider's Leather Armor; wrist: Tabitha's Cuffs; hands: Stormrider's Leather Gloves; waist: Keller's Girdle; legs: Abomination Skin Leggings; feet: Stormrider's Leather Boots; finger1: Lavishly Jeweled Ring; finger2: Advisor's Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Twisted Chanter's Staff

No-known-source sample (15 of 482, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2879 Antipodean Rod; 3738 Brewing Rod; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9753 Nomad Buckler; 9756 Nomad Trousers

### Band 30 (orc, 4532310300000000-000000000000000000-0000000000000000)

Set DPS (verified): 53.2. Weights run: 1.7s. Verify run: 1.2s. 925 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.003, intellect=not significant (1.138 ± 0.340), crit=1.742 ± 0.105, hit=2.217 ± 0.197, spell_haste=2.693 ± 0.329, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.664 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Enduring Cap (3020) | World drop [world_drop] | 18.2 | yes | Enchanter's Cowl (4322, -0.14 DPS, sim-verified) [crafted]; Crusader's Chain Helm (250502, -0.42 DPS) [crafted]; Nightsky Cowl (4039, -0.45 DPS) [world_drop] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 13.8 | yes | Crystal Starfire Medallion (5003, -0.93 DPS) [world_drop]; Darkspear Warding Pendant (272075, -1.30 DPS, sim-verified) [vendor]; Pendant of Myzrael (4614, -1.38 DPS) [dungeon] |
| shoulder | Bloodmage Mantle (7684) | Scarlet Monastery: Bloodmage Thalnos [dungeon] | 19.2 | yes | Fairywing Mantle (9536, -0.30 DPS) [quest]; Death Speaker Mantle (6685, -0.41 DPS, sim-verified) [dungeon]; Feline Mantle (3748, -0.61 DPS) [dungeon] |
| back | Cloak of Rot (4462) (or Darkspear Raider's Cloak (272078)) | Lord Malathrom [world] | 9.1 | yes | Darkspear Raider's Cloak (272078, +0.21 DPS, sim-verified) [vendor]; Soft Willow Cape (16661, -0.34 DPS) [quest]; Pearl-clasped Cloak (5542, -0.37 DPS) [crafted] |
| chest | Green Silk Armor (7065) | Tailoring [crafted] | 23.8 | yes | Guardian Armor (4256, -0.12 DPS, sim-verified) [crafted]; Death Speaker Robes (6682, -0.43 DPS) [dungeon]; Stormrider's Leather Tunic (252510, -0.70 DPS) [crafted] |
| wrist | Spidertank Oilrag (9448) | Gnomeregan: Electrocutioner 6000 [dungeon] | 9.0 | yes | Tabitha's Cuffs (251486, -0.22 DPS) [quest]; Technician's Bracers (270042, -0.22 DPS) [quest]; Nightsky Wristbands (6407, -1.57 DPS, sim-verified) [world_drop] |
| hands | Jutebraid Gloves (10654) | Horde Presence [quest] | 11.7 | yes | Oilrag Handwraps (16741, -0.04 DPS) [quest]; Stormrider's Leather Gloves (252498, -0.21 DPS) [crafted]; Fletcher's Gloves (7348, -1.21 DPS, sim-verified) [crafted] |
| waist | Skycaller's Leather Belt (252522) | Leatherworking [crafted] | 16.8 | yes | Prefect's Belt (250559, +0.45 DPS, sim-verified) [crafted]; Defiler's Cloth Girdle (20164, -0.24 DPS) [rep]; Guardian Belt (4258, -0.29 DPS) [crafted] |
| legs | Kodohide Legguards (285338) | Brontus [world] | 22.5 | yes | Abomination Skin Leggings (23173, -0.41 DPS, sim-verified) [dungeon]; Guardian Pants (5962, -0.54 DPS) [crafted]; Acolyte's Silvered Chain Leggings (250526, -0.54 DPS) [crafted] |
| feet | Gilded Slippers (254001) | Tailoring [crafted] | 15.0 | yes | Stormrider's Leather Boots (252443, -0.33 DPS) [crafted]; Spidersilk Boots (4320, -0.34 DPS) [crafted]; Acidic Walkers (9454, -1.46 DPS, sim-verified) [dungeon] |
| finger1 | Black Widow Band (6199) (or Snake Hoop (6750)) | Leech Widow [world] | 8.0 | yes | Advisor's Ring (19521, -0.10 DPS) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.20 DPS) [vendor] |
| finger2 | Snake Hoop (6750) | Willix the Importer [quest] | 8.0 | yes | Advisor's Ring (19521, +0.46 DPS, sim-verified) [rep]; Lavishly Jeweled Ring (1156, -0.11 DPS) [dungeon]; Sea Giant's Toe Ring (274746, -0.20 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Glimmering Staff (249392) | Enchanting [crafted] | 12.5 | yes | Gnarled Necromancer's Staff (251534, -0.11 DPS) [quest]; Channeler's Staff (4437, -0.34 DPS) [world]; Twisted Chanter's Staff (890, -1.59 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Enduring Cap; neck: Scorn's Icy Choker; shoulder: Bloodmage Mantle; back: Cloak of Rot; chest: Green Silk Armor; wrist: Spidertank Oilrag; hands: Jutebraid Gloves; waist: Skycaller's Leather Belt; legs: Kodohide Legguards; feet: Gilded Slippers; finger1: Black Widow Band; finger2: Snake Hoop; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Glimmering Staff

No-known-source sample (15 of 925, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7956 Bronze Warhammer

### Band 40 (orc, 4532310300103031-000000000000000000-2000000000000000)

Set DPS (verified): 71.7. Weights run: 1.9s. Verify run: 1.4s. 1263 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (0.663 ± 0.634), crit=2.997 ± 0.195, hit=4.584 ± 0.401, spell_haste=not significant (-0.692 ± 0.671), spell_penetration=not significant (0.000 ± 0.000), nature_power=0.529 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Spellpower Goggles Xtreme (10502) | Engineering [crafted] | 21.0 | yes | Big Voodoo Mask (8201, -0.24 DPS) [crafted]; Augural Shroud (2620, -0.30 DPS) [world]; Raging Berserker's Helm (7719, -1.51 DPS, sim-verified) [dungeon] |
| neck | Scorn's Icy Choker (23169) | Scarlet Monastery: Scorn [dungeon] | 11.0 | yes | Necklace of Calisea (1714, -0.36 DPS, sim-verified) [world_drop]; Darkspear Warding Pendant (272074, -0.56 DPS) [vendor]; Darkspear Warding Pendant (272075, -0.68 DPS) [vendor] |
| shoulder | Inquisitor's Shawl (19507) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 15.6 | yes | Bloodmage Mantle (7684, -0.06 DPS) [dungeon]; Berylline Pads (4197, -0.18 DPS) [quest]; Green Silken Shoulders (7057, -0.26 DPS, sim-verified) [crafted] |
| back | Long Silken Cloak (4326) (or Guardian Cloak (5965)) | Tailoring [crafted] | 9.3 | yes | Guardian Cloak (5965, +0.00 DPS, sim-verified) [crafted]; Darkspear Raider's Cloak (272077, -0.18 DPS) [vendor]; Icy Cloak (4327, -0.21 DPS) [crafted] |
| chest | Robe of the Magi (1716) | World drop [world_drop] | 26.0 | yes | Dreamweave Vest (10021, +0.50 DPS, sim-verified) [crafted]; Robe of Power (7054, -0.36 DPS) [crafted]; Crimson Silk Vest (7058, -0.65 DPS) [crafted] |
| wrist | Guardian Leather Bracers (4260) | Leatherworking [crafted] | 10.0 | yes | Turtle Scale Bracers (8198, -0.05 DPS, sim-verified) [crafted]; Radiant Silver Bracers (4545, -0.06 DPS) [quest]; Spidertank Oilrag (9448, -0.09 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) (or Fletcher's Gloves (7348), Dragonscale Gauntlets (8347), Shadowskin Gloves (18238)) | World drop [world_drop] | 42.0 | yes | Dragonscale Gauntlets (8347, +0.00 DPS) [crafted]; Shadowskin Gloves (18238, +0.00 DPS) [crafted]; Fletcher's Gloves (7348, -0.41 DPS, sim-verified) [crafted] |
| waist | Deathmage Sash (10771) | Razorfen Downs: Mordresh Fire Eye [dungeon] | 17.0 | yes | Defiler's Cloth Girdle (20166, -0.03 DPS) [rep]; Skycaller's Leather Belt (252522, -0.26 DPS) [crafted]; Defiler's Chain Girdle (20153, -1.44 DPS, sim-verified) [rep] |
| legs | Red Mageweave Pants (10009) | Tailoring [crafted] | 22.0 | yes | Kodohide Legguards (285338, -0.37 DPS, sim-verified) [world]; Crimson Silk Pantaloons (7062, -0.47 DPS) [crafted]; Abomination Skin Leggings (23173, -0.68 DPS) [dungeon] |
| feet | Earthen Silk Slippers (254013) | Tailoring [crafted] | 24.0 | yes | Skycaller's Leather Shoes (252532, -0.56 DPS, sim-verified) [crafted]; Skycaller's Mail Boots (252563, -0.57 DPS) [crafted]; Mender's Leather Shoes (252533, -1.01 DPS) [crafted] |
| finger1 | Advisor's Ring (19520) | Warsong Outriders [rep] | 9.0 | yes | Advisor's Ring (19521, -0.18 DPS) [rep]; Sea Giant's Toe Ring (274746, -0.27 DPS) [vendor]; Advisor's Ring (20426, -0.36 DPS) [rep] |
| finger2 | Reedknot Ring (9622) | Jarl Needs a Blade [quest] | 7.0 | yes | Sea Giant's Toe Ring (274746, +0.60 DPS, sim-verified) [vendor]; Ogremind Ring (1993, -0.21 DPS) [world_drop]; Voodoo Band (1996, -0.21 DPS) [world] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world] |
| main_hand | Mograine's Might (7723) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | 10.6 | yes | Shadow Crescent Axe (3856, -0.05 DPS) [crafted]; Windweaver Staff (7757, -0.06 DPS) [dungeon]; Illusionary Rod (7713, -4.36 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Spellpower Goggles Xtreme; shoulder: Inquisitor's Shawl; back: Long Silken Cloak; chest: Robe of the Magi; wrist: Guardian Leather Bracers; hands: Gloves of Holy Might; waist: Deathmage Sash; legs: Red Mageweave Pants; feet: Earthen Silk Slippers; finger1: Advisor's Ring; finger2: Reedknot Ring; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Mograine's Might

No-known-source sample (15 of 1263, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape; 6189 Durable Chain Shoulders; 6478 Rat Stompers

### Band 50 (orc, 4532310300103031-000000000000000000-5520000000000000)

Set DPS (verified): 94.3. Weights run: 1.9s. Verify run: 1.5s. 1612 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (3.110 ± 0.789), crit=3.848 ± 0.256, hit=6.081 ± 0.559, spell_haste=4.425 ± 0.985, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.491 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Eye of Theradras (17715) | Maraudon: Princess Theradras [dungeon] | 113.0 | yes | Blood Guard's Pulsing Helmet (220848, +0.27 DPS, sim-verified) [vendor]; Blood Guard's Inscribed Skullcap (220842, -2.19 DPS) [vendor]; Red Mageweave Headband (10033, -2.80 DPS) [crafted] |
| neck | Horizon Choker (13085) | Azuregos [world] | 43.5 | yes | Scorn's Icy Choker (23169, -1.58 DPS) [dungeon]; Mindburst Medallion (11196, -1.66 DPS) [quest]; Darkspear Warding Pendant (272073, -2.72 DPS, sim-verified) [vendor] |
| shoulder | Ironfeather Shoulders (15067) | Leatherworking [crafted] | 77.2 | yes | Blood Guard's Mail Epaulets (220823, -0.69 DPS) [vendor]; Blood Guard's Inscribed Shoulder Pads (220841, -0.69 DPS) [vendor]; Blood Guard's Pulsing Shoulders (220849, -2.64 DPS, sim-verified) [vendor] |
| back | Darkspear Raider's Cloak (272076) | Creeg Bothunk [vendor] | 43.5 | yes | Imperial Red Cloak (8248, -0.82 DPS) [dungeon]; Darkspear Raider's Cloak (272077, -0.82 DPS) [vendor]; Deep Woodlands Cloak (19121, -1.25 DPS, sim-verified) [quest] |
| chest | Stone Guard's Pulsing Breastplate (220844) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Armor (220826, -2.14 DPS, sim-verified) [vendor]; Stone Guard's Inscribed Chestpiece (220838, -2.20 DPS) [vendor]; Knight's Mail Armor (223078, -2.20 DPS) [vendor] |
| wrist | Runic Leather Bracers (15092) | Leatherworking [crafted] | 41.1 | yes | Imperial Red Bracers (8247, -0.61 DPS) [dungeon]; Bloodband Bracers (11469, -0.71 DPS) [quest]; Shizzle's Nozzle Wiper (11917, -1.49 DPS, sim-verified) [quest] |
| hands | Raider Handwraps (272098) | Creeg Bothunk [vendor] | 75.4 | yes | First Sergeant's Mail Gauntlets (220831, -0.53 DPS) [vendor]; First Sergeant's Inscribed Gauntlets (220843, -0.53 DPS) [vendor]; First Sergeant's Pulsing Gauntlets (220845, -1.42 DPS, sim-verified) [vendor] |
| waist | Highlander's Mail Girdle (20118) (or Defiler's Lizardhide Girdle (20174), Defiler's Mail Girdle (20196)) | Samuel Hawke [vendor] | 85.0 | yes | Defiler's Lizardhide Girdle (20174, +0.00 DPS, sim-verified) [rep]; Defiler's Mail Girdle (20196, +0.00 DPS) [rep]; Defiler's Cloth Girdle (20165, -0.58 DPS) [rep] |
| legs | Stone Guard's Pulsing Legplates (220847) | Lady Palanseer [vendor] | 116.1 | yes | Stone Guard's Mail Legplates (220834, -2.20 DPS) [vendor]; Stone Guard's Inscribed Legplates (220839, -2.20 DPS) [vendor]; Stormshroud Pants (15057, -4.33 DPS, sim-verified) [crafted] |
| feet | Greaves of Withering Despair (22240) | Blackrock Depths: High Interrogator Gerstahn  [dungeon] | 102.9 | yes | Skycaller's Leather Boots (252471, -0.83 DPS, sim-verified) [crafted]; Skycaller's Mail Sabatons (252577, -4.73 DPS) [crafted]; Mender's Leather Boots (252472, -5.26 DPS) [crafted] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 60.8 | yes | Voodoo Band (1996, -3.44 DPS) [world]; Mindbender Loop (5009, -3.44 DPS) [world_drop]; Black Widow Band (6199, -3.44 DPS) [world] |
| finger2 | Ogremind Ring (1993) (or Voodoo Band (1996), Mindbender Loop (5009), Black Widow Band (6199), Snake Hoop (6750)) | World drop [world_drop] | 21.8 | yes | Voodoo Band (1996, +0.00 DPS, sim-verified) [world]; Mindbender Loop (5009, +0.00 DPS) [world_drop]; Black Widow Band (6199, +0.00 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 54.7 | yes | Rune of the Guard Captain (19120, -1.07 DPS) [quest]; Uther's Strength (11302, -4.29 DPS) [world]; Tidal Charm (1404, -4.82 DPS) [vendor] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Rune of the Guard Captain (19120, +3.75 DPS) [quest]; Uther's Strength (11302, +0.53 DPS) [world]; Tidal Charm (1404, +0.00 DPS) [vendor] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 0.0 | yes | Illusionary Rod (7713, +6.39 DPS) [dungeon]; Spellshifter Rod (9527, +6.30 DPS) [quest]; Kindling Stave (11750, -1.40 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Eye of Theradras; neck: Horizon Choker; shoulder: Ironfeather Shoulders; back: Darkspear Raider's Cloak; chest: Stone Guard's Pulsing Breastplate; wrist: Runic Leather Bracers; hands: Raider Handwraps; waist: Highlander's Mail Girdle; legs: Stone Guard's Pulsing Legplates; feet: Greaves of Withering Despair; finger1: Blackstone Ring; finger2: Ogremind Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Hammer of the Northern Wind

No-known-source sample (15 of 1612, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

### Band 60 (orc, 4532310300103031-000000000000000000-5533220000000000)

Set DPS (verified): 182.7. Weights run: 1.9s. Verify run: 1.4s. 2370 eligible items had no known source.

Stat weights (normalized to spell_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): spell_power=1.000 ± 0.004, intellect=not significant (1.046 ± 0.865), crit=5.588 ± 0.380, hit=6.652 ± 0.739, spell_haste=8.642 ± 1.178, spell_penetration=not significant (0.000 ± 0.000), nature_power=0.542 ± 0.003

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Mask of the Unforgiven (13404) (or Bloodvine Goggles (19999)) | Stratholme: The Unforgiven [dungeon] | 211.3 | yes | Bloodvine Goggles (19999, +0.52 DPS, sim-verified) [crafted]; Ragefury Eyepatch (11735, -4.75 DPS) [dungeon]; Bloodvine Lens (19998, -4.75 DPS) [crafted] |
| neck | Gem of Trapped Innocents (23057) | Naxxramas: Kel'Thuzad [raid] | 178.8 | yes | Stormrage's Talisman of Seething (23053, -1.94 DPS) [raid]; Onyxia Tooth Pendant (18404, -2.95 DPS) [quest]; Fury of the Forgotten Swarm (21809, -2.95 DPS) [raid] |
| shoulder | Champion's Mail Pauldrons (227154) | Rank 14 [pvp] | 150.0 | yes | Warlord's Mail Spaulders (231659, +0.26 DPS, sim-verified) [pvp]; Mantle of the Timbermaw (19050, -3.57 DPS) [crafted]; Champion's Mail Pauldrons (23260, -4.01 DPS) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 78.2 | yes | Earthweave Cloak (21187, +1.48 DPS, sim-verified) [quest]; Cloak of the Gathering Storm (21400, -4.39 DPS) [quest]; Hide of the Wild (18510, -4.66 DPS) [crafted] |
| chest | Bloodvine Vest (19682) | Tailoring [crafted] | 173.6 | yes | Bloodsoul Breastplate (19690, -1.49 DPS) [crafted]; Legionnaire's Mail Hauberk (227157, -2.50 DPS) [pvp]; Stormshroud Armor (15056, -8.54 DPS, sim-verified) [crafted] |
| wrist | Rockfury Bracers (21186) | Stalwart's Battlegear [quest] | 93.5 | yes | Qiraji Execution Bracers (21602, -2.34 DPS) [raid]; Primal Batskin Bracers (19687, -2.40 DPS, sim-verified) [crafted]; Burrower Bracers (21611, -4.50 DPS) [raid] |
| hands | Blood Guard's Mail Vices (227159) | Rank 11 [pvp] | 144.7 | yes | Primal Batskin Gloves (19686, -1.02 DPS) [crafted]; Stormshroud Gloves (21278, -2.10 DPS, sim-verified) [crafted]; Dark Storm Gauntlets (21585, -2.21 DPS) [raid] |
| waist | Belt of the Archmage (18405) | Tailoring [crafted] | 115.0 | yes | Cord of The Five Thunders (227008, -0.96 DPS) [quest]; Defiler's Cloth Girdle (20163, -1.43 DPS) [rep]; Belt of Never-ending Agony (21586, -2.16 DPS, sim-verified) [raid] |
| legs | Outrider's Chain Leggings (22673) | Warsong Outriders [rep] | 144.7 | yes | Legionnaire's Mail Legguards (227156, +0.00 DPS) [pvp]; General's Mail Leggings (231664, -1.16 DPS) [pvp]; Stormshroud Pants (15057, -3.00 DPS, sim-verified) [crafted] |
| feet | Bloodvine Boots (19684) | Tailoring [crafted] | 102.3 | yes | Blood Guard's Mail Greaves (227158, -3.10 DPS) [pvp]; Greaves of Withering Despair (22240, -3.38 DPS, sim-verified) [dungeon]; Boots of Epiphany (21600, -4.19 DPS) [raid] |
| finger1 | Don Julio's Band (19325) (or Band of Earthen Might (21182), Band of Unnatural Forces (23038)) | Frostwolf Clan [rep] | 144.7 | yes | Band of Unnatural Forces (23038, +0.00 DPS) [raid]; Ring of the Fallen God (21709, -3.03 DPS) [quest]; Mindtear Band (20632, -3.32 DPS) [world] |
| finger2 | Band of Earthen Might (21182) (or Band of Unnatural Forces (23038)) | Veteran's Battlegear [quest] | 144.7 | yes | Band of Unnatural Forces (23038, +1.15 DPS, sim-verified) [raid]; Ring of the Fallen God (21709, -3.03 DPS) [quest]; Mindtear Band (20632, -3.32 DPS) [world] |
| trinket1 | The Restrained Essence of Sapphiron (23046) | Naxxramas: Sapphiron [raid] | 40.0 | yes | Neltharion's Tear (19379, +11.88 DPS) [raid]; Eye of Diminution (23001, +10.10 DPS) [raid]; Drake Fang Talisman (19406, +8.07 DPS) [raid] |
| trinket2 | Kiss of the Spider (22954) | Naxxramas: Maexxna [raid] | 144.7 | yes | Neltharion's Tear (19379, +2.80 DPS) [raid]; Eye of Diminution (23001, +1.02 DPS) [raid]; Drake Fang Talisman (19406, -1.02 DPS) [raid] |
| main_hand | Atiesh, Greatstaff of the Guardian (22630) | Atiesh, Greatstaff of the Guardian [quest] | 336.8 | yes | Atiesh, Greatstaff of the Guardian (22589, -1.76 DPS) [quest]; High Warlord's War Staff (234549, -6.53 DPS) [pvp]; Grand Marshal's Stave (234571, -6.53 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Mask of the Unforgiven; neck: Gem of Trapped Innocents; shoulder: Champion's Mail Pauldrons; back: Chromatic Cloak; chest: Bloodvine Vest; wrist: Rockfury Bracers; hands: Blood Guard's Mail Vices; waist: Belt of the Archmage; legs: Outrider's Chain Leggings; feet: Bloodvine Boots; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: The Restrained Essence of Sapphiron; trinket2: Kiss of the Spider; main_hand: Atiesh, Greatstaff of the Guardian

No-known-source sample (15 of 2370, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2879 Antipodean Rod; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5822 Hedgeseed Gauntlets; 5971 Feathered Cape

