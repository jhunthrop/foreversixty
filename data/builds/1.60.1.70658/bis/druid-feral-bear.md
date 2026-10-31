# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 35.2. Weights run: 3.7s. Verify run: 1.8s. 196 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.082 ± 0.002, defense=0.140 ± 0.026 per rating point (1 rating = 1%, 0.140 per %), dodge=0.096 ± 0.008 per rating point (12 rating = 1%, 1.148 per %), strength=0.052 ± 0.000, agility=0.148 ± 0.005, attack_power=0.026 ± 0.000, hit=0.045 ± 0.005 per rating point (10 rating = 1%, 0.449 per %), crit=0.019 ± 0.001 per rating point (14 rating = 1%, 0.261 per %), expertise=1.200 ± 0.059

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.5 stamina points (352.66 DPS) | yes | Brawler's Leather Hood (252504, -21.07 DPS) [crafted]; Defender's Leather Hood (252447, -35.69 DPS) [crafted]; Totemic Leather Hood (252448, -43.63 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.0 stamina points (95.17 DPS) | yes | Erudite's Amulet (277204, -29.20 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -38.07 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (176.09 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -55.79 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -57.79 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.6 stamina points (145.44 DPS) | yes | Sentry Cloak (2059, -6.19 DPS) [world_drop]; Grave Shroud (279865, -10.42 DPS) [quest]; Firebane Cloak (12979, -20.60 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.0 stamina points (341.76 DPS) | yes | Gloomshroud Armor (1489, -27.53 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -71.89 DPS) [quest]; Trapper's Leather Armor (252491, -94.10 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.0 stamina points (152.96 DPS) | yes | Drakewing Bands (12999, -19.26 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -45.82 DPS) [vendor]; Sanguine Cuffs (14375, -50.27 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 8.9 stamina points (170.19 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.56 DPS) [crafted]; Defender's Leather Gloves (252496, -7.31 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.1 stamina points (210.93 DPS) | yes | Belt of the Fang (10412, -27.73 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.55 DPS) [crafted]; Guardsman Belt (3429, -50.11 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.5 stamina points (314.73 DPS) | yes | Defender's Leather Pants (252445, -46.53 DPS, sim-verified) [crafted]; Slick Deviate Leggings (6480, -53.08 DPS) [quest]; Smelting Pants (5199, -62.79 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (284.63 DPS) | yes | Surfer Shoes (276274, -37.40 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -58.24 DPS) [dungeon]; Trapper's Leather Boots (252440, -80.09 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (114.20 DPS) | yes | Blood Ring (4998, -19.03 DPS) [world_drop]; Ring of the Moon (12052, -56.11 DPS) [world_drop]; Deep Fathom Ring (6463, -57.10 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.3 stamina points (100.50 DPS) | yes | Blood Ring (4998, -5.33 DPS) [world_drop]; Ring of the Moon (12052, -42.40 DPS) [world_drop]; Deep Fathom Ring (6463, -43.40 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.2 stamina points (307.87 DPS) | yes | Rakzur Club (12983, -24.79 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -27.57 DPS) [dungeon]; Twisted Chanter's Staff (890, -51.75 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 196, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 41.6. Weights run: 3.5s. Verify run: 1.6s. 380 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.240 ± 0.054 per rating point (1 rating = 1%, 0.240 per %), dodge=0.106 ± 0.013 per rating point (12 rating = 1%, 1.268 per %), strength=0.049 ± 0.000, agility=0.150 ± 0.008, attack_power=0.025 ± 0.000, hit=0.047 ± 0.009 per rating point (10 rating = 1%, 0.466 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.276 per %), expertise=1.556 ± 0.106

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 stamina points (419.72 DPS) | yes | Totemic Leather Helm (252456, -29.23 DPS) [crafted]; Trapper's Leather Helm (252513, -29.23 DPS) [crafted]; Defender's Leather Helm (252455, -37.57 DPS, sim-verified) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (178.72 DPS) | yes | Master Sergeant's Insignia (18442, -3.84 DPS) [vendor]; Tarnished Locket (279870, -81.57 DPS) [quest]; Kaleidoscope Chain (13084, -85.46 DPS) [world_drop] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 17.9 stamina points (347.95 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -34.09 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -66.71 DPS) [crafted]; Darkshire Shoulderguards (270044, -87.17 DPS) [quest] |
| back | Sergeant's Cape (18440) | PvP rank 7 · Sergeant · Alliance [vendor] | 11.2 stamina points (217.56 DPS) | yes | Tigerstrike Mantle (13108, -17.12 DPS) [world_drop]; Mourning Shawl (6751, -26.00 DPS) [quest]; Enduring Cape (14763, -37.66 DPS) [world_drop] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.3 stamina points (413.47 DPS) | yes | Spirewind Fetter (9406, -9.15 DPS) [dungeon]; Green Whelp Armor (7375, -32.30 DPS) [crafted]; Dokebi Chestguard (14581, -58.29 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (213.05 DPS) | yes | Barbaric Bracers (18948, -3.78 DPS) [crafted]; Black Wolf Bracers (3230, -4.52 DPS) [dungeon]; Staghide Armguards (270021, -32.33 DPS) [quest] |
| hands | Operator's Gloves (270045) | Gyrodrillmatic Excavationators [quest] | 13.0 stamina points (252.73 DPS) | yes | Brawler Gloves (720, -7.25 DPS) [world_drop]; Wolfclaw Gloves (1978, -13.66 DPS) [dungeon]; Gloves of Insight (9698, -16.58 DPS) [quest] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.5 stamina points (301.60 DPS) | yes | Skulker's Leather Belt (252520, -33.03 DPS, sim-verified) [crafted]; Stalker's Leather Belt (252521, -55.29 DPS) [crafted]; Prowler's Leather Belt (252459, -55.41 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.0 stamina points (349.81 DPS) | yes | Dark Ritual Leggings (270031, -30.47 DPS) [quest]; Brawler's Leather Legguards (252516, -33.02 DPS) [crafted]; Duty Bound Leggings (279868, -40.28 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.4 stamina points (357.93 DPS) | yes | Harbinger Boots (7754, -28.15 DPS, sim-verified) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -65.66 DPS) [quest]; Emblazoned Boots (4051, -108.08 DPS) [world_drop] |
| finger1 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (155.45 DPS) | yes | Defias Renegade Ring (1076, -43.68 DPS) [dungeon]; Slain Baron's Signet (279867, -48.97 DPS) [quest]; Boar Signet (274078, -58.29 DPS) [quest] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (116.58 DPS) | yes | Defias Renegade Ring (1076, -4.82 DPS) [dungeon]; Slain Baron's Signet (279867, -10.11 DPS) [quest]; Boar Signet (274078, -19.43 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (2995.6 DPS) | yes | Spine of the Devourer (273040, -30.33 DPS) [world]; Glimmering Staff (249392, -62.65 DPS) [crafted]; Manual Crowd Pummeler (9449, -227.78 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Operator's Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Plains Ring; finger2: Sustaining Ring; main_hand: Staff of the Shade

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 72.4. Weights run: 4.1s. Verify run: 1.8s. 516 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.018, armor=0.107 ± 0.006, defense=0.328 ± 0.080 per rating point (1 rating = 1%, 0.328 per %), dodge=0.113 ± 0.018 per rating point (12 rating = 1%, 1.360 per %), strength=0.046 ± 0.000, agility=0.136 ± 0.011, attack_power=0.023 ± 0.000, hit=not significant (0.055 ± 0.014) per rating point (10 rating = 1%, 0.550 per %), crit=0.029 ± 0.006 per rating point (14 rating = 1%, 0.407 per %), expertise=1.721 ± 0.160

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 25.3 stamina points (713.32 DPS) | yes | Brawler's Leather Helm (252512, -54.53 DPS) [crafted]; Warden's Wizard Hat (14604, -55.16 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -77.43 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (315.83 DPS) | yes | River Pride Choker (13087, -56.97 DPS, sim-verified) [world_drop]; Explorers' League Commendation (7746, -61.68 DPS) [quest]; Master Sergeant's Insignia (18442, -61.68 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 26.2 stamina points (739.52 DPS) | yes | Flintrock Shoulders (7755, -144.93 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -183.86 DPS) [dungeon]; Pads of the Venom Spider (13103, -277.07 DPS) [world_drop] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.5 stamina points (382.18 DPS) | yes | Silky Spider Cape (10776, -9.37 DPS) [dungeon]; Well Oiled Cloak (12254, -12.38 DPS) [vendor]; Heraldic Cloak (8120, -22.21 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 29.1 stamina points (821.22 DPS) | yes | Warden's Wraps (14601, -74.47 DPS) [world_drop]; Insignia Chestguard (4057, -143.66 DPS) [world_drop]; Robes of the Lich (10762, -166.16 DPS, sim-verified) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 13.4 stamina points (377.50 DPS) | yes | Branded Leather Bracers (19508, -29.30 DPS) [dungeon]; Cultist's Armguards (270032, -40.71 DPS) [quest]; Barbaric Bracers (18948, -45.81 DPS) [crafted] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | 18.5 stamina points (523.16 DPS) | yes | Stalker's Leather Gloves (252526, -1.44 DPS) [crafted]; Bonefingers (10765, -24.87 DPS) [dungeon]; Razzeric's Racing Grips (6727, -27.22 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 18.2 stamina points (514.50 DPS) | yes | Warden's Leather Belt (252460, -39.56 DPS) [crafted]; Scorpashi Sash (14652, -68.17 DPS) [world_drop]; Kolkar Hunter's Belt (6788, -90.19 DPS, sim-verified) [quest] |
| legs | Imperial Leather Pants (4062) | Uldaman: Ancient Treasure [dungeon] | 23.3 stamina points (658.36 DPS) | yes | Warden's Woolies (14605, -9.04 DPS) [world_drop]; Basilisk Hide Pants (1718, -11.02 DPS) [world_drop]; Panther Hunter Leggings (4108, -19.88 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 22.1 stamina points (623.97 DPS) | yes | Gnomebot Operating Boots (9450, -56.03 DPS) [dungeon]; Worn Running Boots (9398, -76.00 DPS, sim-verified) [dungeon]; Harbinger Boots (7754, -96.74 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (338.87 DPS) | yes | Plains Ring (2039, -112.96 DPS) [dungeon]; Trader's Ring (15689, -112.96 DPS) [quest]; Field Researcher's Loop (281634, -114.25 DPS) [quest] |
| finger2 | Talvash's Gold Ring (9538) | Gnome Improvement [quest] | 10.0 stamina points (282.39 DPS) | yes | Plains Ring (2039, -56.48 DPS) [dungeon]; Trader's Ring (15689, -56.97 DPS, sim-verified) [quest]; Field Researcher's Loop (281634, -57.77 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (5965.4 DPS) | yes | Windweaver Staff (7757, -138.60 DPS) [dungeon]; Mograine's Might (7723, -153.69 DPS) [dungeon]; Manual Crowd Pummeler (9449, -489.03 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Imperial Leather Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Talvash's Gold Ring; trinket1: Arena Master; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 516, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 114.9. Weights run: 4.5s. Verify run: 4.5s. 686 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.032, armor=0.151 ± 0.012, defense=not significant (0.452 ± 0.148) per rating point (1 rating = 1%, 0.452 per %), dodge=0.255 ± 0.031 per rating point (12 rating = 1%, 3.057 per %), strength=0.057 ± 0.000, agility=0.252 ± 0.020, attack_power=0.029 ± 0.000, hit=not significant (0.096 ± 0.027) per rating point (10 rating = 1%, 0.959 per %), crit=not significant (0.035 ± 0.011) per rating point (14 rating = 1%, 0.488 per %), expertise=3.424 ± 0.295

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.5 stamina points (1513.01 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -150.00 DPS) [vendor]; Scorpashi Skullcap (14658, -356.82 DPS, sim-verified) [world_drop]; Knight-Lieutenant's Crackling Leather Helmet (220872, -360.26 DPS) [vendor] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (498.64 DPS) | yes | Gazlowe's Charm (13088, -98.72 DPS) [dungeon]; River Pride Choker (13087, -169.96 DPS) [world_drop]; Medal of Courage (6723, -178.08 DPS) [quest] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | sim-verified (+133.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Knight-Lieutenant's Leather Shoulders (220852, -22.88 DPS) [vendor]; Warden's Leather Shoulders (252536, -76.88 DPS) [crafted]; Fleshhide Shoulders (10774, -133.57 DPS, sim-verified) [dungeon] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 19.7 stamina points (703.40 DPS) | yes | Sergeant's Cape (18441, -10.78 DPS) [vendor]; Emberscale Cape (11311, -21.55 DPS) [dungeon]; Grovekeeper's Drape (17739, -76.62 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.1 stamina points (1996.92 DPS) | yes | Jinxed Hoodoo Skin (9473, -511.06 DPS, sim-verified) [dungeon]; Heraldic Breastplate (8119, -549.00 DPS) [world_drop]; Knight's Leather Armor (220854, -568.40 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.2 stamina points (897.34 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -106.69 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -163.86 DPS) [crafted]; Serpentskin Bracers (8257, -166.10 DPS) [world_drop] |
| hands | Sergeant Major's Leather Gauntlets (220856) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Warden's Leather Gauntlets (252549, -216.96 DPS, sim-verified) [crafted] |
| waist | Skulker's Leather Waistguard (252474) | Leatherworking [crafted] | sim-verified (+107.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Waistguard (252473, -13.88 DPS) [crafted]; Warden's Leather Waistguard (252475, -107.93 DPS, sim-verified) [crafted]; Pridelord Girdle (14674, -148.55 DPS) [world_drop] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 37.2 stamina points (1326.14 DPS) | yes | Knight's Restored Leather Leggings (220882, -230.53 DPS) [vendor]; Scorpashi Leggings (14659, -242.63 DPS) [world_drop]; Knight's Crackling Leather Leggings (220864, -697.06 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -461.79 DPS, sim-verified) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (498.64 DPS) | yes | Seafire Band (4549, -120.46 DPS) [quest]; Talvash's Gold Ring (9538, -142.47 DPS) [quest]; Ring of Fortitude (10739, -142.47 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (427.40 DPS) | yes | Seafire Band (4549, -49.22 DPS) [quest]; Talvash's Gold Ring (9538, -71.23 DPS) [quest]; Ring of Fortitude (10739, -71.23 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (+302.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Smotts' Compass (4130, -211.69 DPS) [quest]; Frozen Heart of the Mountain (249469, -289.81 DPS) [crafted] |
| trinket2 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mark of the Chosen (17774, +0.00 DPS) [quest]; Smotts' Compass (4130, -140.45 DPS) [quest]; Frozen Heart of the Mountain (249469, -218.57 DPS) [crafted] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -29.23 DPS) [world_drop]; The Hand of Antu'sul (9639, -336.57 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Penance Spaulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Sergeant Major's Leather Gauntlets; waist: Skulker's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Arena Master; main_hand: Radiant Staff

No-known-source sample (15 of 686, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 142.7. Weights run: 4.6s. Verify run: 5.0s. 1564 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.019, armor=0.185 ± 0.016, defense=0.985 ± 0.239 per rating point (1 rating = 1%, 0.985 per %), dodge=0.619 ± 0.048 per rating point (12 rating = 1%, 7.423 per %), strength=0.111 ± 0.000, agility=0.508 ± 0.030, attack_power=0.055 ± 0.000, hit=0.160 ± 0.039 per rating point (10 rating = 1%, 1.598 per %), crit=0.075 ± 0.016 per rating point (14 rating = 1%, 1.045 per %), expertise=5.806 ± 0.422

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.8 stamina points (4707.77 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, -196.94 DPS, sim-verified) [crafted] |
| neck | Master Sergeant's Insignia (18443) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 17.0 stamina points (1099.97 DPS) | yes | Medallion of Grand Marshal Morris (13091, -9.54 DPS) [world_drop]; Sunprism Pendant (20649, -252.90 DPS) [quest]; Heroic Commendation Medal (15799, -323.52 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 50.8 stamina points (3289.07 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -239.25 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 26.3 stamina points (1699.99 DPS) | yes | Sergeant's Cape (16342, -62.60 DPS) [vendor]; Cloak of Earth and Sky (279270, -67.59 DPS) [crafted]; Nightfall Drape (12465, -340.31 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 66.5 stamina points (4303.93 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -351.02 DPS, sim-verified) [dungeon] |
| wrist | Sergeant Major's Dragonhide Armsplints (18454) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 31.2 stamina points (2019.55 DPS) | yes | Arena Bracers (18710, -152.29 DPS) [world]; Shadowcraft Bracers (16710, -226.02 DPS) [dungeon]; Knight-Captain's Dragonhide Armsplints (16395, -229.94 DPS) [pvp] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 42.6 stamina points (2754.08 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -250.60 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 41.7 stamina points (2699.36 DPS) | yes | Elderwild Waistcord (279252, -14.95 DPS) [crafted]; Serpentine Sash (13118, -170.50 DPS) [world_drop]; Marshal's Dragonhide Waistguard (16447, -209.06 DPS) [pvp] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 58.9 stamina points (3813.07 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -373.11 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 47.1 stamina points (3050.41 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -320.80 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (905.86 DPS) | yes | Ring of the Heavens (12056, -64.70 DPS) [dungeon]; Dragonscale Band (12057, -64.70 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -64.70 DPS) [quest] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (905.86 DPS) | yes | Ring of the Heavens (12056, -64.70 DPS) [dungeon]; Dragonscale Band (12057, -64.70 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -64.70 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (14255.6 DPS) | yes | Smotts' Compass (4130, -102.02 DPS) [quest]; Arena Grand Master (19024, -102.02 DPS) [quest]; Arena Master (18706, -129.41 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (14255.6 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (14255.6 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -763.61 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Master Sergeant's Insignia; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Sergeant Major's Dragonhide Armsplints; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Ring of the Dawn; finger2: Ring of Saviors; trinket2: Mark of the Chosen; main_hand: Twig of the World Tree

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 324.1. Weights run: 4.7s. Verify run: 5.1s. 1564 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.015, armor=0.207 ± 0.018, defense=1.532 ± 0.310 per rating point (1 rating = 1%, 1.532 per %), dodge=0.858 ± 0.059 per rating point (12 rating = 1%, 10.293 per %), strength=0.120 ± 0.000, agility=0.762 ± 0.038, attack_power=0.055 ± 0.000, hit=0.262 ± 0.044 per rating point (10 rating = 1%, 2.623 per %), crit=0.125 ± 0.019 per rating point (14 rating = 1%, 1.755 per %), expertise=7.919 ± 0.482

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 82.7 stamina points (7173.83 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 22.3 stamina points (1936.56 DPS) | yes | Master Sergeant's Insignia (18443, -461.85 DPS) [vendor]; Sunprism Pendant (20649, -536.57 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 56.6 stamina points (4910.30 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -642.72 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 28.1 stamina points (2436.77 DPS) | yes | Cloak of Earth and Sky (279270, -152.26 DPS) [crafted]; Sergeant's Cape (16342, -155.17 DPS) [vendor]; Sergeant Major's Cape (16337, -359.83 DPS) [pvp] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 74.7 stamina points (6479.11 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -409.54 DPS, sim-verified) [dungeon] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | 33.1 stamina points (2871.27 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18454, -15.88 DPS) [vendor]; Arena Bracers (18710, -112.85 DPS) [world]; Deepfury Bracers (13120, -254.39 DPS) [world_drop] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 46.6 stamina points (4044.27 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -505.85 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 45.6 stamina points (3951.71 DPS) | yes | Elderwild Waistcord (279252, -139.69 DPS) [crafted]; Wicked Leather Belt (15088, -339.13 DPS) [crafted]; Serpentine Sash (13118, -449.89 DPS, sim-verified) [world_drop] |
| legs | Shadowcraft Pants (16709) | Stratholme: Baron Rivendare [dungeon] | 64.0 stamina points (5555.67 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Elderwild Pants (279254, +0.00 DPS) [crafted] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 53.3 stamina points (4624.31 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -478.25 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.7 stamina points (1358.62 DPS) | yes | Ring of Saviors (1447, -144.15 DPS) [world_drop]; Dragonscale Band (12057, -230.90 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -230.90 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1214.47 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Dragonscale Band (12057, -86.75 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -86.75 DPS) [quest] |
| trinket1 | Arena Grand Master (19024) | Arena Grandmaster [quest] | sim-verified (27375.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Guardian Talisman (1490, -112.15 DPS) [quest]; Arena Master (18706, -285.64 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (27375.4 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (27375.4 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hammer of the Northern Wind (810, -1307.92 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Shadowcraft Bracers; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Shadowcraft Pants; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of the Dawn; trinket1: Arena Grand Master; trinket2: Mark of the Chosen; main_hand: Twig of the World Tree

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 3.7s. Verify run: 1.7s. 192 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.082 ± 0.002, defense=0.140 ± 0.026 per rating point (1 rating = 1%, 0.140 per %), dodge=0.096 ± 0.008 per rating point (12 rating = 1%, 1.148 per %), strength=0.052 ± 0.000, agility=0.148 ± 0.005, attack_power=0.026 ± 0.000, hit=0.045 ± 0.005 per rating point (10 rating = 1%, 0.449 per %), crit=0.019 ± 0.001 per rating point (14 rating = 1%, 0.261 per %), expertise=1.200 ± 0.059

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.5 stamina points (352.66 DPS) | yes | Brawler's Leather Hood (252504, -21.07 DPS) [crafted]; Defender's Leather Hood (252447, -35.69 DPS) [crafted]; Totemic Leather Hood (252448, -43.63 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (62.74 DPS) | yes | Scholarly Pendant (277203, -5.64 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (176.09 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -55.79 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -57.79 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.6 stamina points (145.44 DPS) | yes | Sentry Cloak (2059, -6.19 DPS) [world_drop]; Grave Shroud (279865, -10.42 DPS) [quest]; Firebane Cloak (12979, -20.60 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.0 stamina points (341.76 DPS) | yes | Gloomshroud Armor (1489, -38.10 DPS, sim-verified) [dungeon]; Trapper's Leather Armor (252491, -94.10 DPS) [crafted]; Camouflaged Tunic (3585, -95.42 DPS) [quest] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.7 stamina points (166.10 DPS) | yes | Savannah Bracers (15453, -14.71 DPS) [quest]; Bear Bracers (4795, -16.66 DPS, sim-verified) [vendor]; Drakewing Bands (12999, -44.96 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 8.9 stamina points (170.19 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.56 DPS) [crafted]; Defender's Leather Gloves (252496, -7.31 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.1 stamina points (210.93 DPS) | yes | Belt of the Fang (10412, -30.43 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.55 DPS) [crafted]; Guardsman Belt (3429, -50.11 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 13.8 stamina points (262.48 DPS) | yes | Slick Deviate Leggings (6480, -0.83 DPS) [quest]; Smelting Pants (5199, -10.54 DPS) [dungeon]; Brawler's Leather Pants (252500, -26.95 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (284.63 DPS) | yes | Surfer Shoes (276274, -51.60 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -58.24 DPS) [dungeon]; Trapper's Leather Boots (252440, -80.09 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.3 stamina points (100.50 DPS) | yes | Ring of Scorn (3235, -24.36 DPS) [quest]; Ring of the Moon (12052, -42.40 DPS) [world_drop]; Deep Fathom Ring (6463, -43.40 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (95.17 DPS) | yes | Ring of Scorn (3235, -16.75 DPS, sim-verified) [quest]; Ring of the Moon (12052, -37.08 DPS) [world_drop]; Deep Fathom Ring (6463, -38.07 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.2 stamina points (307.87 DPS) | yes | Hammerbone (270018, -22.80 DPS) [quest]; Crescent Staff (6505, -23.77 DPS, sim-verified) [quest]; Rakzur Club (12983, -24.79 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 192, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 41.9. Weights run: 3.5s. Verify run: 1.7s. 380 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.240 ± 0.054 per rating point (1 rating = 1%, 0.240 per %), dodge=0.106 ± 0.013 per rating point (12 rating = 1%, 1.268 per %), strength=0.049 ± 0.000, agility=0.150 ± 0.008, attack_power=0.025 ± 0.000, hit=0.047 ± 0.009 per rating point (10 rating = 1%, 0.466 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.276 per %), expertise=1.556 ± 0.106

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 stamina points (419.72 DPS) | yes | Totemic Leather Helm (252456, -29.23 DPS) [crafted]; Trapper's Leather Helm (252513, -29.23 DPS) [crafted]; Defender's Leather Helm (252455, -32.78 DPS, sim-verified) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (178.72 DPS) | yes | Senior Sergeant's Insignia (15200, -3.84 DPS) [vendor]; Kaleidoscope Chain (13084, -85.46 DPS) [world_drop]; Pendant of Myzrael (4614, -101.00 DPS) [dungeon] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 17.9 stamina points (347.95 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -66.71 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -66.71 DPS) [crafted]; Tanned Shoulderpads (270023, -69.38 DPS, sim-verified) [quest] |
| back | Sergeant's Cloak (18427) | PvP rank 7 · Sergeant · Horde [vendor] | 11.2 stamina points (217.56 DPS) | yes | Tigerstrike Mantle (13108, -17.12 DPS) [world_drop]; Enduring Cape (14763, -37.66 DPS) [world_drop]; Grimsteel Cape (4643, -37.94 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.3 stamina points (413.47 DPS) | yes | Spirewind Fetter (9406, -9.15 DPS) [dungeon]; Green Whelp Armor (7375, -32.30 DPS) [crafted]; Dokebi Chestguard (14581, -58.29 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (213.05 DPS) | yes | Barbaric Bracers (18948, -3.78 DPS) [crafted]; Black Wolf Bracers (3230, -4.52 DPS) [dungeon]; Technician's Bracers (270042, -33.94 DPS) [quest] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 13.5 stamina points (261.54 DPS) | yes | Wolfclaw Gloves (1978, -22.47 DPS) [dungeon]; Brawler Gloves (720, -25.41 DPS, sim-verified) [world_drop]; Emblazoned Gloves (6397, -36.77 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.5 stamina points (301.60 DPS) | yes | Larcenist's Belt (279846, -42.15 DPS, sim-verified) [quest]; Skulker's Leather Belt (252520, -49.52 DPS) [crafted]; Stalker's Leather Belt (252521, -55.29 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.0 stamina points (349.81 DPS) | yes | Brawler's Leather Legguards (252516, -33.02 DPS) [crafted]; Trapper's Leather Legguards (252517, -39.74 DPS) [crafted]; Dark Ritual Leggings (270031, -43.66 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.4 stamina points (357.93 DPS) | yes | Harbinger Boots (7754, -32.36 DPS, sim-verified) [dungeon]; Grizzled Boots (6335, -49.11 DPS) [quest]; Nat Pagle's Extreme Anglin' Boots (19969, -65.66 DPS) [quest] |
| finger1 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.1 stamina points (158.33 DPS) | yes | Defias Renegade Ring (1076, -46.56 DPS) [dungeon]; Slain Baron's Signet (279867, -51.86 DPS) [quest]; Boar Signet (274078, -61.18 DPS) [quest] |
| finger2 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (155.45 DPS) | yes | Slain Baron's Signet (279867, -48.97 DPS) [quest]; Boar Signet (274078, -58.29 DPS) [quest]; Defias Renegade Ring (1076, -58.49 DPS, sim-verified) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (3238.3 DPS) | yes | Spine of the Devourer (273040, -30.33 DPS) [world]; Glimmering Staff (249392, -62.65 DPS) [crafted]; Manual Crowd Pummeler (9449, -262.20 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Seal of Sylvanas; finger2: Plains Ring; main_hand: Staff of the Shade

No-known-source sample (15 of 380, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 71.6. Weights run: 4.1s. Verify run: 1.9s. 514 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.018, armor=0.107 ± 0.006, defense=0.328 ± 0.080 per rating point (1 rating = 1%, 0.328 per %), dodge=0.113 ± 0.018 per rating point (12 rating = 1%, 1.360 per %), strength=0.046 ± 0.000, agility=0.136 ± 0.011, attack_power=0.023 ± 0.000, hit=not significant (0.055 ± 0.014) per rating point (10 rating = 1%, 0.550 per %), crit=0.029 ± 0.006 per rating point (14 rating = 1%, 0.407 per %), expertise=1.721 ± 0.160

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 25.3 stamina points (713.32 DPS) | yes | Warden's Wizard Hat (14604, -50.96 DPS, sim-verified) [world_drop]; Brawler's Leather Helm (252512, -54.53 DPS) [crafted]; Defender's Leather Helm (252455, -77.43 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (315.83 DPS) | yes | River Pride Choker (13087, -54.87 DPS, sim-verified) [world_drop]; Senior Sergeant's Insignia (15200, -61.68 DPS) [vendor]; Necklace of Calisea (1714, -89.92 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 26.2 stamina points (739.52 DPS) | yes | Flintrock Shoulders (7755, -139.95 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -183.86 DPS) [dungeon]; Tanned Shoulderpads (270023, -265.34 DPS) [quest] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.5 stamina points (382.18 DPS) | yes | Silky Spider Cape (10776, -9.37 DPS) [dungeon]; Well Oiled Cloak (12254, -12.38 DPS) [vendor]; Heraldic Cloak (8120, -22.21 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 29.1 stamina points (821.22 DPS) | yes | Warden's Wraps (14601, -74.47 DPS) [world_drop]; Insignia Chestguard (4057, -143.66 DPS) [world_drop]; Robes of the Lich (10762, -169.39 DPS, sim-verified) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 13.4 stamina points (377.50 DPS) | yes | Branded Leather Bracers (19508, -29.30 DPS) [dungeon]; Cultist's Armguards (270032, -40.71 DPS) [quest]; Barbaric Bracers (18948, -45.81 DPS) [crafted] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | sim-verified (6492.9 DPS) | yes | Bonefingers (10765, -23.43 DPS) [dungeon]; Razzeric's Racing Grips (6727, -25.78 DPS) [quest]; Warden's Leather Gloves (252527, -74.61 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 18.2 stamina points (514.50 DPS) | yes | Warden's Leather Belt (252460, -39.56 DPS) [crafted]; Kolkar Hunter's Belt (6788, -64.22 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -68.17 DPS) [world_drop] |
| legs | Imperial Leather Pants (4062) | Uldaman: Ancient Treasure [dungeon] | 23.3 stamina points (658.36 DPS) | yes | Warden's Woolies (14605, -9.04 DPS) [world_drop]; Basilisk Hide Pants (1718, -11.02 DPS) [world_drop]; Panther Hunter Leggings (4108, -19.88 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 22.1 stamina points (623.97 DPS) | yes | Worn Running Boots (9398, -29.31 DPS) [dungeon]; Gnomebot Operating Boots (9450, -56.03 DPS) [dungeon]; Harbinger Boots (7754, -96.74 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (338.87 DPS) | yes | Seal of Sylvanas (6414, -109.06 DPS) [quest]; Plains Ring (2039, -112.96 DPS) [dungeon]; Trader's Ring (15689, -112.96 DPS) [quest] |
| finger2 | Nogg's Gold Ring (9588) | Nogg's Ring Redo [quest] | 10.0 stamina points (282.39 DPS) | yes | Seal of Sylvanas (6414, -48.17 DPS, sim-verified) [quest]; Plains Ring (2039, -56.48 DPS) [dungeon]; Trader's Ring (15689, -56.48 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -48.42 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -159.70 DPS) [quest]; Gut Ripper (2164, -647.93 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Imperial Leather Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Nogg's Gold Ring; trinket1: Arena Master; main_hand: Cragwood Maul

No-known-source sample (15 of 514, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 117.6. Weights run: 4.5s. Verify run: 4.2s. 684 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.032, armor=0.151 ± 0.012, defense=not significant (0.452 ± 0.148) per rating point (1 rating = 1%, 0.452 per %), dodge=0.255 ± 0.031 per rating point (12 rating = 1%, 3.057 per %), strength=0.057 ± 0.000, agility=0.252 ± 0.020, attack_power=0.029 ± 0.000, hit=not significant (0.096 ± 0.027) per rating point (10 rating = 1%, 0.959 per %), crit=not significant (0.035 ± 0.011) per rating point (14 rating = 1%, 0.488 per %), expertise=3.424 ± 0.295

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.5 stamina points (1513.01 DPS) | yes | Blood Guard's Leather Headband (220851, -150.00 DPS) [vendor]; Scorpashi Skullcap (14658, -340.72 DPS, sim-verified) [world_drop]; Blood Guard's Crackling Leather Helmet (220873, -360.26 DPS) [vendor] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (498.64 DPS) | yes | Woven Ivy Necklace (19159, -49.51 DPS) [quest]; Gazlowe's Charm (13088, -98.72 DPS) [dungeon]; River Pride Choker (13087, -169.96 DPS) [world_drop] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | 31.2 stamina points (1110.17 DPS) | yes | Fleshhide Shoulders (10774, -0.01 DPS) [dungeon]; Blood Guard's Leather Shoulders (220853, -22.88 DPS) [vendor]; Warden's Leather Shoulders (252536, -76.88 DPS) [crafted] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 19.7 stamina points (703.40 DPS) | yes | Sergeant's Cloak (16341, -10.78 DPS) [vendor]; Emberscale Cape (11311, -21.55 DPS) [dungeon]; Grovekeeper's Drape (17739, -76.62 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.1 stamina points (1996.92 DPS) | yes | Heraldic Breastplate (8119, -549.00 DPS) [world_drop]; Jinxed Hoodoo Skin (9473, -554.20 DPS, sim-verified) [dungeon]; Stone Guard's Leather Armor (220855, -568.40 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.2 stamina points (897.34 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -120.42 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -163.86 DPS) [crafted]; Serpentskin Bracers (8257, -166.10 DPS) [world_drop] |
| hands | First Sergeant's Leather Gauntlets (220857) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-verified (11703.0 DPS) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Warden's Leather Gauntlets (252549, -405.61 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.1 stamina points (1002.60 DPS) | yes | Skulker's Leather Waistguard (252474, -16.84 DPS) [crafted]; Prowler's Leather Waistguard (252473, -30.72 DPS) [crafted]; Pridelord Girdle (14674, -165.39 DPS) [world_drop] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 37.2 stamina points (1326.14 DPS) | yes | Stone Guard's Restored Leather Leggings (220883, -230.53 DPS) [vendor]; Scorpashi Leggings (14659, -242.63 DPS) [world_drop]; Stone Guard's Crackling Leather Leggings (220865, -758.26 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Leather Boots (220861) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-verified (11703.0 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -445.98 DPS, sim-verified) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (498.64 DPS) | yes | Seafire Band (4549, -120.46 DPS) [quest]; Nogg's Gold Ring (9588, -142.47 DPS) [quest]; Sha'ni's Ring (11869, -142.47 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (427.40 DPS) | yes | Seafire Band (4549, -49.22 DPS) [quest]; Nogg's Gold Ring (9588, -71.23 DPS) [quest]; Sha'ni's Ring (11869, -71.23 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (11703.0 DPS) | yes | Arena Master (18706, -71.23 DPS) [world]; Smotts' Compass (4130, -211.69 DPS) [quest]; Rune of the Guard Captain (19120, -253.96 DPS) [quest] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (11703.0 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (11703.0 DPS) | yes | Radiant Staff (249453, -6.23 DPS) [crafted]; Dreamstaff (249454, -6.23 DPS) [crafted]; The Hand of Antu'sul (9639, -703.24 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; shoulder: Penance Spaulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: First Sergeant's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: First Sergeant's Leather Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Mark of the Chosen

No-known-source sample (15 of 684, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 143.7. Weights run: 4.6s. Verify run: 4.9s. 1564 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.019, armor=0.185 ± 0.016, defense=0.985 ± 0.239 per rating point (1 rating = 1%, 0.985 per %), dodge=0.619 ± 0.048 per rating point (12 rating = 1%, 7.423 per %), strength=0.111 ± 0.000, agility=0.508 ± 0.030, attack_power=0.055 ± 0.000, hit=0.160 ± 0.039 per rating point (10 rating = 1%, 1.598 per %), crit=0.075 ± 0.016 per rating point (14 rating = 1%, 1.045 per %), expertise=5.806 ± 0.422

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.8 stamina points (4707.77 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Senior Sergeant's Insignia (16335) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 17.0 stamina points (1099.97 DPS) | yes | Medallion of Grand Marshal Morris (13091, -9.54 DPS) [world_drop]; Woven Ivy Necklace (19159, -114.35 DPS) [quest]; Hunter's Insignia Medal (15704, -129.41 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 50.8 stamina points (3289.07 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -371.09 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 26.3 stamina points (1699.99 DPS) | yes | Sergeant's Cloak (18461, -62.60 DPS) [vendor]; Cloak of Earth and Sky (279270, -67.59 DPS) [crafted]; Nightfall Drape (12465, -340.31 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 66.5 stamina points (4303.93 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -365.70 DPS, sim-verified) [dungeon] |
| wrist | First Sergeant's Dragonhide Armguards (18434) | PvP rank 9 · First Sergeant · Horde [vendor] | 31.2 stamina points (2019.55 DPS) | yes | Arena Bracers (18710, -160.93 DPS, sim-verified) [world]; Shadowcraft Bracers (16710, -226.02 DPS) [dungeon]; Legionnaire's Dragonhide Armguards (16493, -229.94 DPS) [pvp] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 42.6 stamina points (2754.08 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -150.71 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 41.7 stamina points (2699.36 DPS) | yes | Elderwild Waistcord (279252, -14.95 DPS) [crafted]; Serpentine Sash (13118, -170.50 DPS) [world_drop]; General's Dragonhide Belt (16556, -209.06 DPS) [pvp] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 58.9 stamina points (3813.07 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -235.53 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 47.1 stamina points (3050.41 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -300.09 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (905.86 DPS) | yes | Ring of the Heavens (12056, -64.70 DPS) [dungeon]; Dragonscale Band (12057, -64.70 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -64.70 DPS) [quest] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (905.86 DPS) | yes | Ring of the Heavens (12056, -64.70 DPS) [dungeon]; Dragonscale Band (12057, -64.70 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -64.70 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (15116.9 DPS) | yes | Smotts' Compass (4130, -102.02 DPS) [quest]; Arena Grand Master (19024, -102.02 DPS) [quest]; Arena Master (18706, -129.41 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (15116.9 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (15116.9 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hammer of the Northern Wind (810, -687.88 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Senior Sergeant's Insignia; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: First Sergeant's Dragonhide Armguards; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Ring of the Dawn; finger2: Ring of Saviors; main_hand: Twig of the World Tree

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 324.8. Weights run: 4.7s. Verify run: 5.0s. 1564 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.015, armor=0.207 ± 0.018, defense=1.532 ± 0.310 per rating point (1 rating = 1%, 1.532 per %), dodge=0.858 ± 0.059 per rating point (12 rating = 1%, 10.293 per %), strength=0.120 ± 0.000, agility=0.762 ± 0.038, attack_power=0.055 ± 0.000, hit=0.262 ± 0.044 per rating point (10 rating = 1%, 2.623 per %), crit=0.125 ± 0.019 per rating point (14 rating = 1%, 1.755 per %), expertise=7.919 ± 0.482

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 82.7 stamina points (7173.83 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, -391.63 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 22.3 stamina points (1936.56 DPS) | yes | Woven Ivy Necklace (19159, -270.64 DPS, sim-verified) [quest]; Senior Sergeant's Insignia (16335, -461.85 DPS) [vendor]; Sunprism Pendant (20649, -536.57 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 56.6 stamina points (4910.30 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -462.90 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 28.1 stamina points (2436.77 DPS) | yes | Cloak of Earth and Sky (279270, -152.26 DPS) [crafted]; Sergeant's Cloak (18461, -155.17 DPS) [vendor]; Deep River Cloak (15789, -457.33 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 74.7 stamina points (6479.11 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -313.87 DPS, sim-verified) [dungeon] |
| wrist | Shadowcraft Bracers (16710) | Scholomance: Scholomance Occultist [dungeon] | 33.1 stamina points (2871.27 DPS) | yes | First Sergeant's Dragonhide Armguards (18434, +0.00 DPS, sim-verified) [vendor]; Arena Bracers (18710, -112.85 DPS) [world]; Deepfury Bracers (13120, -254.39 DPS) [world_drop] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 46.6 stamina points (4044.27 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -327.50 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 45.6 stamina points (3951.71 DPS) | yes | Serpentine Sash (13118, -68.86 DPS) [world_drop]; Elderwild Waistcord (279252, -139.69 DPS) [crafted]; Wicked Leather Belt (15088, -339.13 DPS) [crafted] |
| legs | Shadowcraft Pants (16709) | Stratholme: Baron Rivendare [dungeon] | 64.0 stamina points (5555.67 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Elderwild Pants (279254, +0.00 DPS) [crafted] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 53.3 stamina points (4624.31 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -352.83 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.7 stamina points (1358.62 DPS) | yes | Ring of Saviors (1447, -144.15 DPS) [world_drop]; Dragonscale Band (12057, -230.90 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -230.90 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1214.47 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Dragonscale Band (12057, -86.75 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -86.75 DPS) [quest] |
| trinket1 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (28963.5 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| trinket2 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (28963.5 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (28963.5 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hammer of the Northern Wind (810, -1478.24 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Shadowcraft Bracers; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Shadowcraft Pants; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of the Dawn; trinket1: Mark of the Chosen; trinket2: Guardian Talisman; main_hand: Twig of the World Tree

No-known-source sample (15 of 1564, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

