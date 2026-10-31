# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 35.9. Weights run: 3.6s. Verify run: 1.7s. 196 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.082 ± 0.002, defense=0.148 ± 0.027 per rating point (1 rating = 1%, 0.148 per %), dodge=0.095 ± 0.008 per rating point (12 rating = 1%, 1.135 per %), strength=0.051 ± 0.000, agility=0.147 ± 0.005, attack_power=0.025 ± 0.000, hit=0.045 ± 0.006 per rating point (10 rating = 1%, 0.452 per %), crit=0.018 ± 0.001 per rating point (14 rating = 1%, 0.247 per %), expertise=1.181 ± 0.062

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.5 stamina points (353.74 DPS) | yes | Brawler's Leather Hood (252504, -21.03 DPS) [crafted]; Defender's Leather Hood (252447, -35.69 DPS) [crafted]; Totemic Leather Hood (252448, -43.47 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.0 stamina points (95.37 DPS) | yes | Erudite's Amulet (277204, -34.73 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -38.15 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (176.95 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -55.98 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -58.19 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.6 stamina points (145.90 DPS) | yes | Sentry Cloak (2059, -6.28 DPS) [world_drop]; Grave Shroud (279865, -10.55 DPS) [quest]; Firebane Cloak (12979, -20.65 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.0 stamina points (343.02 DPS) | yes | Gloomshroud Armor (1489, -35.83 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -72.10 DPS) [quest]; Trapper's Leather Armor (252491, -94.33 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (153.56 DPS) | yes | Drakewing Bands (12999, -22.93 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -46.00 DPS) [vendor]; Sanguine Cuffs (14375, -50.53 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.0 stamina points (170.87 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.57 DPS) [crafted]; Defender's Leather Gloves (252496, -7.33 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.1 stamina points (211.60 DPS) | yes | Belt of the Fang (10412, -33.32 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.59 DPS) [crafted]; Guardsman Belt (3429, -50.16 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.6 stamina points (315.89 DPS) | yes | Slick Deviate Leggings (6480, -53.20 DPS) [quest]; Defender's Leather Pants (252445, -59.37 DPS, sim-verified) [crafted]; Smelting Pants (5199, -62.85 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (285.51 DPS) | yes | Surfer Shoes (276274, -47.94 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -58.30 DPS) [dungeon]; Trapper's Leather Boots (252440, -80.18 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (114.45 DPS) | yes | Blood Ring (4998, -19.07 DPS) [world_drop]; Ring of the Moon (12052, -56.25 DPS) [world_drop]; Deep Fathom Ring (6463, -57.22 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.3 stamina points (101.02 DPS) | yes | Blood Ring (4998, -5.64 DPS) [world_drop]; Ring of the Moon (12052, -42.82 DPS) [world_drop]; Deep Fathom Ring (6463, -43.79 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.0 stamina points (305.58 DPS) | yes | Rakzur Club (12983, -24.72 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -27.45 DPS) [dungeon]; Twisted Chanter's Staff (890, -51.98 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 43.0. Weights run: 3.3s. Verify run: 1.8s. 380 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.091 ± 0.004, defense=0.242 ± 0.055 per rating point (1 rating = 1%, 0.242 per %), dodge=0.105 ± 0.013 per rating point (12 rating = 1%, 1.262 per %), strength=0.048 ± 0.000, agility=0.150 ± 0.008, attack_power=0.024 ± 0.000, hit=0.058 ± 0.009 per rating point (10 rating = 1%, 0.579 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.278 per %), expertise=1.586 ± 0.107

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.1 stamina points (430.70 DPS) | yes | Defender's Leather Helm (252455, -18.00 DPS) [crafted]; Totemic Leather Helm (252456, -29.20 DPS) [crafted]; Trapper's Leather Helm (252513, -29.20 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (179.10 DPS) | yes | Master Sergeant's Insignia (18442, -3.73 DPS) [vendor]; Tarnished Locket (279870, -81.67 DPS) [quest]; Kaleidoscope Chain (13084, -85.75 DPS) [world_drop] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.4 stamina points (358.28 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -51.83 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -67.36 DPS) [crafted]; Darkshire Shoulderguards (270044, -87.80 DPS) [quest] |
| back | Sergeant's Cape (18440) | PvP rank 7 · Sergeant · Alliance [vendor] | 11.4 stamina points (221.24 DPS) | yes | Tigerstrike Mantle (13108, -17.38 DPS) [world_drop]; Mourning Shawl (6751, -26.54 DPS) [quest]; Enduring Cape (14763, -38.32 DPS) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | sim-verified (3018.1 DPS) | yes | Green Whelp Armor (7375, -24.17 DPS) [crafted]; Raptor Hide Harness (4455, -31.44 DPS, sim-verified) [crafted]; Dokebi Chestguard (14581, -50.72 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.2 stamina points (218.69 DPS) | yes | Barbaric Bracers (18948, -3.45 DPS) [crafted]; Black Wolf Bracers (3230, -4.56 DPS) [dungeon]; Staghide Armguards (270021, -32.57 DPS) [quest] |
| hands | Operator's Gloves (270045) | Gyrodrillmatic Excavationators [quest] | 13.4 stamina points (261.14 DPS) | yes | Brawler Gloves (720, -7.89 DPS) [world_drop]; Wolfclaw Gloves (1978, -14.42 DPS) [dungeon]; Gloves of Insight (9698, -17.12 DPS) [quest] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.9 stamina points (309.63 DPS) | yes | Skulker's Leather Belt (252520, -42.46 DPS, sim-verified) [crafted]; Stalker's Leather Belt (252521, -55.29 DPS) [crafted]; Prowler's Leather Belt (252459, -55.66 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.5 stamina points (361.12 DPS) | yes | Duty Bound Leggings (279868, -25.33 DPS) [quest]; Dark Ritual Leggings (270031, -30.50 DPS) [quest]; Brawler's Leather Legguards (252516, -33.13 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.9 stamina points (367.91 DPS) | yes | Harbinger Boots (7754, -26.44 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -70.57 DPS) [quest]; Emblazoned Boots (4051, -109.84 DPS) [world_drop] |
| finger1 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (155.88 DPS) | yes | Defias Renegade Ring (1076, -43.86 DPS) [dungeon]; Slain Baron's Signet (279867, -49.01 DPS) [quest]; Boar Signet (274078, -58.45 DPS) [quest] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (116.91 DPS) | yes | Defias Renegade Ring (1076, -4.89 DPS) [dungeon]; Slain Baron's Signet (279867, -10.04 DPS) [quest]; Boar Signet (274078, -19.48 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -30.06 DPS) [world]; Glimmering Staff (249392, -63.31 DPS) [crafted]; Manual Crowd Pummeler (9449, -222.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Spirewind Fetter; wrist: Cultist's Armguards; hands: Operator's Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Plains Ring; finger2: Sustaining Ring; main_hand: Staff of the Shade

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 73.7. Weights run: 4.0s. Verify run: 1.8s. 516 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.106 ± 0.007, defense=0.336 ± 0.083 per rating point (1 rating = 1%, 0.336 per %), dodge=0.125 ± 0.018 per rating point (12 rating = 1%, 1.497 per %), strength=0.045 ± 0.000, agility=0.140 ± 0.012, attack_power=0.023 ± 0.000, hit=not significant (0.061 ± 0.016) per rating point (10 rating = 1%, 0.607 per %), crit=not significant (0.023 ± 0.006) per rating point (14 rating = 1%, 0.325 per %), expertise=1.851 ± 0.170

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 25.2 stamina points (702.51 DPS) | yes | Brawler's Leather Helm (252512, -53.43 DPS) [crafted]; Warden's Wizard Hat (14604, -56.19 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -77.33 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (311.80 DPS) | yes | River Pride Choker (13087, -55.20 DPS, sim-verified) [world_drop]; Explorers' League Commendation (7746, -60.84 DPS) [quest]; Master Sergeant's Insignia (18442, -60.84 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 26.1 stamina points (727.94 DPS) | yes | Flintrock Shoulders (7755, -178.21 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -181.73 DPS) [dungeon]; Pads of the Venom Spider (13103, -272.54 DPS) [world_drop] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.5 stamina points (376.98 DPS) | yes | Silky Spider Cape (10776, -9.75 DPS) [dungeon]; Well Oiled Cloak (12254, -12.70 DPS) [vendor]; Heraldic Cloak (8120, -21.99 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 29.0 stamina points (807.64 DPS) | yes | Warden's Wraps (14601, -72.57 DPS) [world_drop]; Insignia Chestguard (4057, -141.70 DPS) [world_drop]; Robes of the Lich (10762, -189.56 DPS, sim-verified) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 13.3 stamina points (371.85 DPS) | yes | Branded Leather Bracers (19508, -29.84 DPS) [dungeon]; Cultist's Armguards (270032, -40.70 DPS) [quest]; Barbaric Bracers (18948, -45.37 DPS) [crafted] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | 18.4 stamina points (514.15 DPS) | yes | Stalker's Leather Gloves (252526, -0.18 DPS) [crafted]; Bonefingers (10765, -24.57 DPS) [dungeon]; Razzeric's Racing Grips (6727, -25.68 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 18.2 stamina points (506.74 DPS) | yes | Warden's Leather Belt (252460, -39.10 DPS) [crafted]; Scorpashi Sash (14652, -67.04 DPS) [world_drop]; Kolkar Hunter's Belt (6788, -67.93 DPS, sim-verified) [quest] |
| legs | Imperial Leather Pants (4062) | Uldaman: Ancient Treasure [dungeon] | 23.2 stamina points (646.89 DPS) | yes | Basilisk Hide Pants (1718, -8.78 DPS) [world_drop]; Warden's Woolies (14605, -8.84 DPS) [world_drop]; Panther Hunter Leggings (4108, -18.23 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 22.0 stamina points (614.01 DPS) | yes | Gnomebot Operating Boots (9450, -55.50 DPS) [dungeon]; Worn Running Boots (9398, -72.66 DPS, sim-verified) [dungeon]; Harbinger Boots (7754, -95.58 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (334.62 DPS) | yes | Plains Ring (2039, -111.54 DPS) [dungeon]; Trader's Ring (15689, -111.54 DPS) [quest]; Field Researcher's Loop (281634, -112.05 DPS) [quest] |
| finger2 | Talvash's Gold Ring (9538) | Gnome Improvement [quest] | 10.0 stamina points (278.85 DPS) | yes | Trader's Ring (15689, -55.20 DPS, sim-verified) [quest]; Plains Ring (2039, -55.77 DPS) [dungeon]; Field Researcher's Loop (281634, -56.28 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (5980.6 DPS) | yes | Windweaver Staff (7757, -136.22 DPS) [dungeon]; Mograine's Might (7723, -151.96 DPS) [dungeon]; Coldrage Dagger (10761, -556.58 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Imperial Leather Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Talvash's Gold Ring; trinket1: Arena Master; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 119.5. Weights run: 4.5s. Verify run: 4.3s. 686 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.035, armor=0.159 ± 0.013, defense=not significant (0.624 ± 0.157) per rating point (1 rating = 1%, 0.624 per %), dodge=0.268 ± 0.033 per rating point (12 rating = 1%, 3.213 per %), strength=0.055 ± 0.000, agility=0.260 ± 0.020, attack_power=0.028 ± 0.000, hit=not significant (0.088 ± 0.029) per rating point (10 rating = 1%, 0.882 per %), crit=not significant (0.046 ± 0.012) per rating point (14 rating = 1%, 0.649 per %), expertise=3.543 ± 0.310

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 43.5 stamina points (1543.52 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -148.31 DPS) [vendor]; Knight-Lieutenant's Crackling Leather Helmet (220872, -360.49 DPS) [vendor]; Scorpashi Skullcap (14658, -482.90 DPS, sim-verified) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (497.26 DPS) | yes | Gazlowe's Charm (13088, -98.69 DPS) [dungeon]; River Pride Choker (13087, -169.73 DPS) [world_drop]; Medal of Courage (6723, -177.59 DPS) [quest] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | 32.1 stamina points (1140.17 DPS) | yes | Fleshhide Shoulders (10774, -6.91 DPS) [dungeon]; Knight-Lieutenant's Leather Shoulders (220852, -22.29 DPS) [vendor]; Warden's Leather Shoulders (252536, -81.18 DPS) [crafted] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 20.0 stamina points (711.39 DPS) | yes | Sergeant's Cape (18441, -11.27 DPS) [vendor]; Emberscale Cape (11311, -22.54 DPS) [dungeon]; Grovekeeper's Drape (17739, -76.67 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 57.3 stamina points (2036.95 DPS) | yes | Heraldic Breastplate (8119, -557.99 DPS) [world_drop]; Knight's Leather Armor (220854, -569.94 DPS) [vendor]; Jinxed Hoodoo Skin (9473, -581.01 DPS, sim-verified) [dungeon] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.7 stamina points (913.23 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -148.16 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -163.83 DPS) [crafted]; Serpentskin Bracers (8257, -168.03 DPS) [world_drop] |
| hands | Sergeant Major's Leather Gauntlets (220856) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (11023.1 DPS) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Warden's Leather Gauntlets (252549, -343.87 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.8 stamina points (1022.37 DPS) | yes | Skulker's Leather Waistguard (252474, -12.97 DPS) [crafted]; Prowler's Leather Waistguard (252473, -27.48 DPS) [crafted]; Pridelord Girdle (14674, -163.26 DPS) [world_drop] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 38.3 stamina points (1359.95 DPS) | yes | Knight's Restored Leather Leggings (220882, -226.62 DPS) [vendor]; Basilisk Hide Pants (1718, -245.38 DPS) [world_drop]; Knight's Crackling Leather Leggings (220864, -706.50 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (11023.1 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -458.47 DPS, sim-verified) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (497.26 DPS) | yes | Seafire Band (4549, -119.70 DPS) [quest]; Talvash's Gold Ring (9538, -142.07 DPS) [quest]; Ring of Fortitude (10739, -142.07 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (426.22 DPS) | yes | Seafire Band (4549, -48.66 DPS) [quest]; Talvash's Gold Ring (9538, -71.04 DPS) [quest]; Ring of Fortitude (10739, -71.04 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (11023.1 DPS) | yes | Mark of the Chosen (17774, -182.53 DPS, sim-verified) [quest]; Smotts' Compass (4130, -205.56 DPS) [quest]; Frozen Heart of the Mountain (249469, -291.47 DPS) [crafted] |
| trinket2 | Arena Master (18706) | Arena Treasure Chest [world] | sim-verified (11023.1 DPS) | yes | Mark of the Chosen (17774, -117.83 DPS, sim-verified) [quest]; Smotts' Compass (4130, -134.52 DPS) [quest]; Frozen Heart of the Mountain (249469, -220.43 DPS) [crafted] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (11023.1 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -32.54 DPS) [world_drop]; The Hand of Antu'sul (9639, -564.12 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Penance Spaulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Sergeant Major's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Arena Master; main_hand: Radiant Staff

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 147.7. Weights run: 4.5s. Verify run: 4.8s. 1564 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.036, armor=0.190 ± 0.017, defense=1.008 ± 0.248 per rating point (1 rating = 1%, 1.008 per %), dodge=0.606 ± 0.049 per rating point (12 rating = 1%, 7.268 per %), strength=0.103 ± 0.000, agility=0.494 ± 0.032, attack_power=0.051 ± 0.000, hit=0.190 ± 0.040 per rating point (10 rating = 1%, 1.904 per %), crit=0.073 ± 0.017 per rating point (14 rating = 1%, 1.026 per %), expertise=6.486 ± 0.441

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 73.3 stamina points (4912.05 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, -174.16 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 17.1 stamina points (1143.87 DPS) | yes | Master Sergeant's Insignia (18443, -5.26 DPS) [vendor]; Sunprism Pendant (20649, -277.91 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 51.3 stamina points (3435.82 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -355.93 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 26.5 stamina points (1777.13 DPS) | yes | Sergeant's Cape (16342, -65.41 DPS) [vendor]; Cloak of Earth and Sky (279270, -68.71 DPS) [crafted]; Nightfall Drape (12465, -355.49 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 67.1 stamina points (4491.91 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -393.13 DPS, sim-verified) [dungeon] |
| wrist | Sergeant Major's Dragonhide Armsplints (18454) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 31.6 stamina points (2119.27 DPS) | yes | Arena Bracers (18710, -173.23 DPS, sim-verified) [world]; Knight-Captain's Dragonhide Armsplints (16395, -239.14 DPS) [pvp]; Shadowcraft Bracers (16710, -249.78 DPS) [dungeon] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 43.0 stamina points (2882.99 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Bileblister Claws (275610, -245.41 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 42.1 stamina points (2822.61 DPS) | yes | Elderwild Waistcord (279252, -2.41 DPS) [crafted]; Serpentine Sash (13118, -188.08 DPS) [world_drop]; Marshal's Dragonhide Waistguard (16447, -203.34 DPS) [pvp] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 59.9 stamina points (4011.68 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -499.25 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 47.5 stamina points (3182.62 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -309.33 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (937.68 DPS) | yes | Ring of the Heavens (12056, -66.98 DPS) [dungeon]; Dragonscale Band (12057, -66.98 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -66.98 DPS) [quest] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (937.68 DPS) | yes | Ring of the Heavens (12056, -66.98 DPS) [dungeon]; Dragonscale Band (12057, -66.98 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -66.98 DPS) [quest] |
| trinket1 | Arena Grand Master (19024) | Arena Grandmaster [quest] | sim-verified (14402.5 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, -17.96 DPS) [world]; Demon's Blood (10779, -216.79 DPS) [quest] |
| trinket2 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (14402.5 DPS) | yes | Smotts' Compass (4130, -115.99 DPS) [quest]; Arena Master (18706, -133.95 DPS) [world]; Demon's Blood (10779, -332.78 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (14402.5 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -877.50 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Sergeant Major's Dragonhide Armsplints; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Ring of the Dawn; finger2: Ring of Saviors; trinket1: Arena Grand Master; trinket2: Guardian Talisman; main_hand: Twig of the World Tree

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 324.0. Weights run: 4.6s. Verify run: 4.9s. 1564 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.002, armor=0.216 ± 0.021, defense=not significant (1.261 ± 0.338) per rating point (1 rating = 1%, 1.261 per %), dodge=0.852 ± 0.062 per rating point (12 rating = 1%, 10.222 per %), strength=0.121 ± 0.000, agility=0.773 ± 0.041, attack_power=0.055 ± 0.000, hit=0.258 ± 0.049 per rating point (10 rating = 1%, 2.583 per %), crit=0.120 ± 0.022 per rating point (14 rating = 1%, 1.684 per %), expertise=7.994 ± 0.532

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 84.4 stamina points (7257.23 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, -294.16 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 19.6 stamina points (1687.00 DPS) | yes | Sunprism Pendant (20649, -287.05 DPS) [quest]; Master Sergeant's Insignia (18443, -448.92 DPS, sim-verified) [vendor] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 58.0 stamina points (4987.61 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -421.77 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 28.5 stamina points (2451.66 DPS) | yes | Cloak of Earth and Sky (279270, -144.86 DPS) [crafted]; Sergeant's Cape (16342, -151.85 DPS) [vendor]; Sergeant Major's Cape (16337, -348.11 DPS) [pvp] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 76.4 stamina points (6573.88 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Dawn Armor (252483, -437.65 DPS, sim-verified) [crafted] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | 33.9 stamina points (2920.39 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18454, -25.20 DPS) [vendor]; Arena Bracers (18710, -126.31 DPS) [world]; Deepfury Bracers (13120, -253.59 DPS) [world_drop] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 47.7 stamina points (4102.38 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -521.81 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 46.5 stamina points (4002.35 DPS) | yes | Serpentine Sash (13118, -61.63 DPS) [world_drop]; Elderwild Waistcord (279252, -130.44 DPS) [crafted]; Wicked Leather Belt (15088, -334.56 DPS) [crafted] |
| legs | Shadowcraft Pants (16709) | Stratholme: Baron Rivendare [dungeon] | 65.8 stamina points (5657.23 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Elderwild Pants (279254, +0.00 DPS) [crafted] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 54.6 stamina points (4694.18 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -305.47 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.8 stamina points (1360.85 DPS) | yes | Ring of Saviors (1447, -156.39 DPS) [world_drop]; Dragonscale Band (12057, -242.42 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -242.42 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1204.46 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Dragonscale Band (12057, -86.03 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -86.03 DPS) [quest] |
| trinket1 | Demon's Blood (10779) | You Are Rakh'likh, Demon [quest] | sim-verified (27545.9 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| trinket2 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (27545.9 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Arena Master (18706, -172.07 DPS) [world] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (27545.9 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -1236.85 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Shadowcraft Bracers; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Shadowcraft Pants; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of the Dawn; trinket1: Demon's Blood; trinket2: Guardian Talisman; main_hand: Twig of the World Tree

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 36.6. Weights run: 3.6s. Verify run: 1.7s. 192 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.082 ± 0.002, defense=0.148 ± 0.027 per rating point (1 rating = 1%, 0.148 per %), dodge=0.095 ± 0.008 per rating point (12 rating = 1%, 1.135 per %), strength=0.051 ± 0.000, agility=0.147 ± 0.005, attack_power=0.025 ± 0.000, hit=0.045 ± 0.006 per rating point (10 rating = 1%, 0.452 per %), crit=0.018 ± 0.001 per rating point (14 rating = 1%, 0.247 per %), expertise=1.181 ± 0.062

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.5 stamina points (353.74 DPS) | yes | Brawler's Leather Hood (252504, -21.03 DPS) [crafted]; Defender's Leather Hood (252447, -35.69 DPS) [crafted]; Totemic Leather Hood (252448, -43.47 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (62.83 DPS) | yes | Scholarly Pendant (277203, -15.84 DPS, sim-verified) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (176.95 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -55.98 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -58.19 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.6 stamina points (145.90 DPS) | yes | Sentry Cloak (2059, -6.28 DPS) [world_drop]; Grave Shroud (279865, -10.55 DPS) [quest]; Firebane Cloak (12979, -20.65 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.0 stamina points (343.02 DPS) | yes | Gloomshroud Armor (1489, -40.11 DPS, sim-verified) [dungeon]; Trapper's Leather Armor (252491, -94.33 DPS) [crafted]; Camouflaged Tunic (3585, -95.57 DPS) [quest] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.7 stamina points (166.70 DPS) | yes | Savannah Bracers (15453, -14.71 DPS) [quest]; Bear Bracers (4795, -24.43 DPS, sim-verified) [vendor]; Drakewing Bands (12999, -44.99 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.0 stamina points (170.87 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.57 DPS) [crafted]; Defender's Leather Gloves (252496, -7.33 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.1 stamina points (211.60 DPS) | yes | Belt of the Fang (10412, -37.89 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.59 DPS) [crafted]; Guardsman Belt (3429, -50.16 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 13.8 stamina points (263.38 DPS) | yes | Smelting Pants (5199, -10.34 DPS) [dungeon]; Slick Deviate Leggings (6480, -12.78 DPS, sim-verified) [quest]; Brawler's Leather Pants (252500, -27.04 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (285.51 DPS) | yes | Surfer Shoes (276274, -54.23 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -58.30 DPS) [dungeon]; Trapper's Leather Boots (252440, -80.18 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.3 stamina points (101.02 DPS) | yes | Ring of Scorn (3235, -24.72 DPS) [quest]; Ring of the Moon (12052, -42.82 DPS) [world_drop]; Deep Fathom Ring (6463, -43.79 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (95.37 DPS) | yes | Ring of Scorn (3235, -16.81 DPS, sim-verified) [quest]; Ring of the Moon (12052, -37.18 DPS) [world_drop]; Deep Fathom Ring (6463, -38.15 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.0 stamina points (305.58 DPS) | yes | Crescent Staff (6505, -13.95 DPS) [quest]; Hammerbone (270018, -23.18 DPS) [quest]; Rakzur Club (12983, -24.72 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 42.5. Weights run: 3.3s. Verify run: 1.7s. 380 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.091 ± 0.004, defense=0.242 ± 0.055 per rating point (1 rating = 1%, 0.242 per %), dodge=0.105 ± 0.013 per rating point (12 rating = 1%, 1.262 per %), strength=0.048 ± 0.000, agility=0.150 ± 0.008, attack_power=0.024 ± 0.000, hit=0.058 ± 0.009 per rating point (10 rating = 1%, 0.579 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.278 per %), expertise=1.586 ± 0.107

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.1 stamina points (430.70 DPS) | yes | Defender's Leather Helm (252455, -18.00 DPS) [crafted]; Totemic Leather Helm (252456, -29.20 DPS) [crafted]; Trapper's Leather Helm (252513, -29.20 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (179.10 DPS) | yes | Senior Sergeant's Insignia (15200, -3.73 DPS) [vendor]; Kaleidoscope Chain (13084, -85.75 DPS) [world_drop]; Pendant of Myzrael (4614, -101.16 DPS) [dungeon] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.4 stamina points (358.28 DPS) | yes | Tanned Shoulderpads (270023, -48.79 DPS, sim-verified) [quest]; Cloudy Gustwoven Spaulders (277043, -67.36 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -67.36 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 7 · Sergeant · Horde [vendor] | 11.4 stamina points (221.24 DPS) | yes | Tigerstrike Mantle (13108, -17.38 DPS) [world_drop]; Enduring Cape (14763, -38.32 DPS) [world_drop]; Grimsteel Cape (4643, -38.42 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.9 stamina points (426.21 DPS) | yes | Spirewind Fetter (9406, +0.00 DPS, sim-verified) [dungeon]; Green Whelp Armor (7375, -31.91 DPS) [crafted]; Dokebi Chestguard (14581, -58.45 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.2 stamina points (218.69 DPS) | yes | Barbaric Bracers (18948, -3.45 DPS) [crafted]; Black Wolf Bracers (3230, -4.56 DPS) [dungeon]; Technician's Bracers (270042, -33.68 DPS) [quest] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 13.8 stamina points (269.75 DPS) | yes | Brawler Gloves (720, -16.50 DPS) [world_drop]; Wolfclaw Gloves (1978, -23.03 DPS) [dungeon]; Emblazoned Gloves (6397, -37.67 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.9 stamina points (309.63 DPS) | yes | Larcenist's Belt (279846, -38.86 DPS, sim-verified) [quest]; Skulker's Leather Belt (252520, -49.70 DPS) [crafted]; Stalker's Leather Belt (252521, -55.29 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.5 stamina points (361.12 DPS) | yes | Dark Ritual Leggings (270031, -26.87 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -33.13 DPS) [crafted]; Trapper's Leather Legguards (252517, -39.66 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.9 stamina points (367.91 DPS) | yes | Harbinger Boots (7754, -29.42 DPS, sim-verified) [dungeon]; Grizzled Boots (6335, -50.29 DPS) [quest]; Nat Pagle's Extreme Anglin' Boots (19969, -70.57 DPS) [quest] |
| finger1 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.1 stamina points (158.68 DPS) | yes | Defias Renegade Ring (1076, -46.66 DPS) [dungeon]; Slain Baron's Signet (279867, -51.81 DPS) [quest]; Boar Signet (274078, -61.25 DPS) [quest] |
| finger2 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (155.88 DPS) | yes | Defias Renegade Ring (1076, -39.00 DPS, sim-verified) [dungeon]; Slain Baron's Signet (279867, -49.01 DPS) [quest]; Boar Signet (274078, -58.45 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (3234.2 DPS) | yes | Spine of the Devourer (273040, -30.06 DPS) [world]; Glimmering Staff (249392, -63.31 DPS) [crafted]; Manual Crowd Pummeler (9449, -235.83 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Seal of Sylvanas; finger2: Plains Ring; main_hand: Staff of the Shade

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 74.5. Weights run: 4.0s. Verify run: 1.9s. 514 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.106 ± 0.007, defense=0.336 ± 0.083 per rating point (1 rating = 1%, 0.336 per %), dodge=0.125 ± 0.018 per rating point (12 rating = 1%, 1.497 per %), strength=0.045 ± 0.000, agility=0.140 ± 0.012, attack_power=0.023 ± 0.000, hit=not significant (0.061 ± 0.016) per rating point (10 rating = 1%, 0.607 per %), crit=not significant (0.023 ± 0.006) per rating point (14 rating = 1%, 0.325 per %), expertise=1.851 ± 0.170

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 25.2 stamina points (702.51 DPS) | yes | Warden's Wizard Hat (14604, -53.28 DPS) [world_drop]; Brawler's Leather Helm (252512, -53.43 DPS) [crafted]; Defender's Leather Helm (252455, -77.33 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (311.80 DPS) | yes | River Pride Choker (13087, -55.89 DPS, sim-verified) [world_drop]; Senior Sergeant's Insignia (15200, -60.84 DPS) [vendor]; Necklace of Calisea (1714, -88.73 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 26.1 stamina points (727.94 DPS) | yes | Flintrock Shoulders (7755, -157.91 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -181.73 DPS) [dungeon]; Tanned Shoulderpads (270023, -261.93 DPS) [quest] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.5 stamina points (376.98 DPS) | yes | Silky Spider Cape (10776, -9.75 DPS) [dungeon]; Well Oiled Cloak (12254, -12.70 DPS) [vendor]; Heraldic Cloak (8120, -21.99 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 29.0 stamina points (807.64 DPS) | yes | Warden's Wraps (14601, -72.57 DPS) [world_drop]; Insignia Chestguard (4057, -141.70 DPS) [world_drop]; Robes of the Lich (10762, -145.08 DPS, sim-verified) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 13.3 stamina points (371.85 DPS) | yes | Branded Leather Bracers (19508, -29.84 DPS) [dungeon]; Cultist's Armguards (270032, -40.70 DPS) [quest]; Barbaric Bracers (18948, -45.37 DPS) [crafted] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | 18.4 stamina points (514.15 DPS) | yes | Stalker's Leather Gloves (252526, +0.00 DPS, sim-verified) [crafted]; Bonefingers (10765, -24.57 DPS) [dungeon]; Razzeric's Racing Grips (6727, -25.68 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 18.2 stamina points (506.74 DPS) | yes | Warden's Leather Belt (252460, -39.10 DPS) [crafted]; Scorpashi Sash (14652, -67.04 DPS) [world_drop]; Kolkar Hunter's Belt (6788, -68.33 DPS, sim-verified) [quest] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | sim-verified (6497.2 DPS) | yes | Warden's Woolies (14605, -0.06 DPS) [world_drop]; Panther Hunter Leggings (4108, -9.45 DPS) [quest]; Imperial Leather Pants (4062, -77.93 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 22.0 stamina points (614.01 DPS) | yes | Worn Running Boots (9398, -28.02 DPS) [dungeon]; Gnomebot Operating Boots (9450, -55.50 DPS) [dungeon]; Harbinger Boots (7754, -95.58 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (334.62 DPS) | yes | Seal of Sylvanas (6414, -107.74 DPS) [quest]; Plains Ring (2039, -111.54 DPS) [dungeon]; Trader's Ring (15689, -111.54 DPS) [quest] |
| finger2 | Nogg's Gold Ring (9588) | Nogg's Ring Redo [quest] | 10.0 stamina points (278.85 DPS) | yes | Seal of Sylvanas (6414, -49.32 DPS, sim-verified) [quest]; Plains Ring (2039, -55.77 DPS) [dungeon]; Trader's Ring (15689, -55.77 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -48.59 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -157.81 DPS) [quest]; Manual Crowd Pummeler (9449, -540.45 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Nogg's Gold Ring; trinket1: Arena Master; main_hand: Cragwood Maul

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 123.1. Weights run: 4.5s. Verify run: 4.2s. 684 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.035, armor=0.159 ± 0.013, defense=not significant (0.624 ± 0.157) per rating point (1 rating = 1%, 0.624 per %), dodge=0.268 ± 0.033 per rating point (12 rating = 1%, 3.213 per %), strength=0.055 ± 0.000, agility=0.260 ± 0.020, attack_power=0.028 ± 0.000, hit=not significant (0.088 ± 0.029) per rating point (10 rating = 1%, 0.882 per %), crit=not significant (0.046 ± 0.012) per rating point (14 rating = 1%, 0.649 per %), expertise=3.543 ± 0.310

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 43.5 stamina points (1543.52 DPS) | yes | Blood Guard's Leather Headband (220851, -148.31 DPS) [vendor]; Blood Guard's Crackling Leather Helmet (220873, -360.49 DPS) [vendor]; Scorpashi Skullcap (14658, -362.74 DPS, sim-verified) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (497.26 DPS) | yes | Woven Ivy Necklace (19159, -47.28 DPS) [quest]; Gazlowe's Charm (13088, -98.69 DPS) [dungeon]; River Pride Choker (13087, -169.73 DPS) [world_drop] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | 32.1 stamina points (1140.17 DPS) | yes | Fleshhide Shoulders (10774, -6.91 DPS) [dungeon]; Blood Guard's Leather Shoulders (220853, -22.29 DPS) [vendor]; Warden's Leather Shoulders (252536, -81.18 DPS) [crafted] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 20.0 stamina points (711.39 DPS) | yes | Sergeant's Cloak (16341, -11.27 DPS) [vendor]; Emberscale Cape (11311, -22.54 DPS) [dungeon]; Grovekeeper's Drape (17739, -76.67 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 57.3 stamina points (2036.95 DPS) | yes | Heraldic Breastplate (8119, -557.99 DPS) [world_drop]; Stone Guard's Leather Armor (220855, -569.94 DPS) [vendor]; Jinxed Hoodoo Skin (9473, -604.80 DPS, sim-verified) [dungeon] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.7 stamina points (913.23 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -55.33 DPS) [vendor]; Warden's Leather Bracers (252541, -163.83 DPS) [crafted]; Serpentskin Bracers (8257, -168.03 DPS) [world_drop] |
| hands | First Sergeant's Leather Gauntlets (220857) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-verified (11786.4 DPS) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Warden's Leather Gauntlets (252549, -328.85 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.8 stamina points (1022.37 DPS) | yes | Skulker's Leather Waistguard (252474, -12.97 DPS) [crafted]; Prowler's Leather Waistguard (252473, -27.48 DPS) [crafted]; Pridelord Girdle (14674, -163.26 DPS) [world_drop] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 38.3 stamina points (1359.95 DPS) | yes | Stone Guard's Restored Leather Leggings (220883, -226.62 DPS) [vendor]; Basilisk Hide Pants (1718, -245.38 DPS) [world_drop]; Stone Guard's Crackling Leather Leggings (220865, -744.16 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Leather Boots (220861) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-verified (11786.4 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -468.96 DPS, sim-verified) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (497.26 DPS) | yes | Seafire Band (4549, -119.70 DPS) [quest]; Nogg's Gold Ring (9588, -142.07 DPS) [quest]; Sha'ni's Ring (11869, -142.07 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (426.22 DPS) | yes | Seafire Band (4549, -48.66 DPS) [quest]; Nogg's Gold Ring (9588, -71.04 DPS) [quest]; Sha'ni's Ring (11869, -71.04 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (11786.4 DPS) | yes | Arena Master (18706, -71.04 DPS) [world]; Smotts' Compass (4130, -205.56 DPS) [quest]; Rune of the Guard Captain (19120, -256.45 DPS) [quest] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (11786.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (11786.4 DPS) | yes | Radiant Staff (249453, -11.37 DPS) [crafted]; Dreamstaff (249454, -11.37 DPS) [crafted]; The Hand of Antu'sul (9639, -631.13 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; shoulder: Penance Spaulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: First Sergeant's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: First Sergeant's Leather Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Mark of the Chosen

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 150.6. Weights run: 4.5s. Verify run: 4.8s. 1564 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.036, armor=0.190 ± 0.017, defense=1.008 ± 0.248 per rating point (1 rating = 1%, 1.008 per %), dodge=0.606 ± 0.049 per rating point (12 rating = 1%, 7.268 per %), strength=0.103 ± 0.000, agility=0.494 ± 0.032, attack_power=0.051 ± 0.000, hit=0.190 ± 0.040 per rating point (10 rating = 1%, 1.904 per %), crit=0.073 ± 0.017 per rating point (14 rating = 1%, 1.026 per %), expertise=6.486 ± 0.441

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 73.3 stamina points (4912.05 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 17.1 stamina points (1143.87 DPS) | yes | Senior Sergeant's Insignia (16335, -5.26 DPS) [vendor]; Woven Ivy Necklace (19159, -135.03 DPS) [quest]; Hunter's Insignia Medal (15704, -139.21 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 51.3 stamina points (3435.82 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -384.59 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 26.5 stamina points (1777.13 DPS) | yes | Sergeant's Cloak (18461, -65.41 DPS) [vendor]; Cloak of Earth and Sky (279270, -68.71 DPS) [crafted]; Nightfall Drape (12465, -355.49 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 67.1 stamina points (4491.91 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -159.91 DPS, sim-verified) [dungeon] |
| wrist | First Sergeant's Dragonhide Armguards (18434) | PvP rank 9 · First Sergeant · Horde [vendor] | 31.6 stamina points (2119.27 DPS) | yes | Arena Bracers (18710, -167.94 DPS) [world]; Legionnaire's Dragonhide Armguards (16493, -239.14 DPS) [pvp]; Shadowcraft Bracers (16710, -249.78 DPS) [dungeon] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 43.0 stamina points (2882.99 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Bileblister Claws (275610, -147.93 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 42.1 stamina points (2822.61 DPS) | yes | Elderwild Waistcord (279252, -2.41 DPS) [crafted]; Serpentine Sash (13118, -188.08 DPS) [world_drop]; General's Dragonhide Belt (16556, -203.34 DPS) [pvp] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 59.9 stamina points (4011.68 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -351.00 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 47.5 stamina points (3182.62 DPS) | yes | Shadowcraft Boots (16711, +0.00 DPS) [dungeon]; General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (937.68 DPS) | yes | Ring of the Heavens (12056, -66.98 DPS) [dungeon]; Dragonscale Band (12057, -66.98 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -66.98 DPS) [quest] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (937.68 DPS) | yes | Ring of the Heavens (12056, -66.98 DPS) [dungeon]; Dragonscale Band (12057, -66.98 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -66.98 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (15224.4 DPS) | yes | Smotts' Compass (4130, -115.99 DPS) [quest]; Arena Grand Master (19024, -115.99 DPS) [quest]; Arena Master (18706, -133.95 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (15224.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (15224.4 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hammer of the Northern Wind (810, -794.98 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: First Sergeant's Dragonhide Armguards; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Ring of the Dawn; finger2: Ring of Saviors; main_hand: Twig of the World Tree

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 321.1. Weights run: 4.6s. Verify run: 5.3s. 1564 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.002, armor=0.216 ± 0.021, defense=not significant (1.261 ± 0.338) per rating point (1 rating = 1%, 1.261 per %), dodge=0.852 ± 0.062 per rating point (12 rating = 1%, 10.222 per %), strength=0.121 ± 0.000, agility=0.773 ± 0.041, attack_power=0.055 ± 0.000, hit=0.258 ± 0.049 per rating point (10 rating = 1%, 2.583 per %), crit=0.120 ± 0.022 per rating point (14 rating = 1%, 1.684 per %), expertise=7.994 ± 0.532

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 84.4 stamina points (7257.23 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, -387.35 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 19.6 stamina points (1687.00 DPS) | yes | Senior Sergeant's Insignia (16335, -224.44 DPS) [vendor]; Sunprism Pendant (20649, -287.05 DPS) [quest]; Woven Ivy Necklace (19159, -466.75 DPS, sim-verified) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 58.0 stamina points (4987.61 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -393.11 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 28.5 stamina points (2451.66 DPS) | yes | Cloak of Earth and Sky (279270, -144.86 DPS) [crafted]; Sergeant's Cloak (18461, -151.85 DPS) [vendor]; Deep River Cloak (15789, -448.38 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 76.4 stamina points (6573.88 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Dawn Armor (252483, -308.32 DPS, sim-verified) [crafted] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | 33.9 stamina points (2920.39 DPS) | yes | First Sergeant's Dragonhide Armguards (18434, -25.20 DPS) [vendor]; Arena Bracers (18710, -126.31 DPS) [world]; Deepfury Bracers (13120, -253.59 DPS) [world_drop] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 47.7 stamina points (4102.38 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 46.5 stamina points (4002.35 DPS) | yes | Serpentine Sash (13118, -61.63 DPS) [world_drop]; Elderwild Waistcord (279252, -130.44 DPS) [crafted]; Wicked Leather Belt (15088, -334.56 DPS) [crafted] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | sim-verified (29140.4 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -309.09 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 54.6 stamina points (4694.18 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -366.17 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.8 stamina points (1360.85 DPS) | yes | Ring of Saviors (1447, -156.39 DPS) [world_drop]; Dragonscale Band (12057, -242.42 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -242.42 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1204.46 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Dragonscale Band (12057, -86.03 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -86.03 DPS) [quest] |
| trinket1 | Arena Grand Master (19024) | Arena Grandmaster [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, -277.18 DPS) [world]; Demon's Blood (10779, -445.50 DPS) [quest] |
| trinket2 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, -172.07 DPS) [world]; Demon's Blood (10779, -340.39 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hammer of the Northern Wind (810, -1264.90 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Shadowcraft Bracers; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of the Dawn; trinket1: Arena Grand Master; trinket2: Guardian Talisman; main_hand: Twig of the World Tree

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

