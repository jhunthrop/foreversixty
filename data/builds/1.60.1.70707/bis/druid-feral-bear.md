# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 35.8. Weights run: 2.3s. Verify run: 1.2s. 203 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.085 ± 0.002, defense=0.152 ± 0.029 per rating point (1 rating = 1%, 0.152 per %), dodge=0.095 ± 0.009 per rating point (12 rating = 1%, 1.140 per %), strength=0.051 ± 0.000, agility=0.149 ± 0.005, attack_power=0.025 ± 0.000, hit=0.044 ± 0.006 per rating point (10 rating = 1%, 0.442 per %), crit=0.018 ± 0.001 per rating point (14 rating = 1%, 0.246 per %), expertise=1.194 ± 0.066

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (296.01 DPS) | yes | Brawler's Leather Hood (252504, -16.10 DPS) [crafted]; Defender's Leather Hood (252447, -28.57 DPS) [crafted]; Totemic Leather Hood (252448, -35.04 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.0 stamina points (79.40 DPS) | yes | Erudite's Amulet (277204, -28.99 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -31.76 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.4 stamina points (149.56 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -46.31 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -49.74 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (122.17 DPS) | yes | Sentry Cloak (2059, -5.07 DPS) [world_drop]; Grave Shroud (279865, -8.72 DPS) [quest]; Firebane Cloak (12979, -17.22 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (288.79 DPS) | yes | Gloomshroud Armor (1489, -29.51 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -59.66 DPS) [quest]; Trapper's Leather Armor (252491, -78.54 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (129.14 DPS) | yes | Drakewing Bands (12999, -18.60 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -38.17 DPS) [vendor]; Sanguine Cuffs (14375, -42.77 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (144.24 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.34 DPS) [crafted]; Defender's Leather Gloves (252496, -6.24 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.2 stamina points (178.10 DPS) | yes | Belt of the Fang (10412, -28.49 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -40.58 DPS) [crafted]; Guardsman Belt (3429, -41.93 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.8 stamina points (266.02 DPS) | yes | Slick Deviate Leggings (6480, -44.56 DPS) [quest]; Defender's Leather Pants (252445, -49.72 DPS, sim-verified) [crafted]; Smelting Pants (5199, -52.68 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (238.96 DPS) | yes | Surfer Shoes (276274, -40.13 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -47.47 DPS) [dungeon]; Trapper's Leather Boots (252440, -65.72 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (95.28 DPS) | yes | Blood Ring (4998, -15.88 DPS) [world_drop]; Ring of the Moon (12052, -46.83 DPS) [world_drop]; Deep Fathom Ring (6463, -47.64 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.3 stamina points (84.22 DPS) | yes | Blood Ring (4998, -4.82 DPS) [world_drop]; Ring of the Moon (12052, -35.77 DPS) [world_drop]; Deep Fathom Ring (6463, -36.58 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.0 stamina points (254.16 DPS) | yes | Rakzur Club (12983, -20.57 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -22.84 DPS) [dungeon]; Twisted Chanter's Staff (890, -43.29 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 42.9. Weights run: 2.1s. Verify run: 1.3s. 387 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.097 ± 0.005, defense=0.247 ± 0.061 per rating point (1 rating = 1%, 0.247 per %), dodge=0.104 ± 0.015 per rating point (12 rating = 1%, 1.245 per %), strength=0.048 ± 0.000, agility=0.154 ± 0.009, attack_power=0.024 ± 0.000, hit=0.061 ± 0.011 per rating point (10 rating = 1%, 0.609 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.282 per %), expertise=1.645 ± 0.122

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.7 stamina points (364.52 DPS) | yes | Defender's Leather Helm (252455, -15.52 DPS) [crafted]; Totemic Leather Helm (252456, -24.80 DPS) [crafted]; Trapper's Leather Helm (252513, -24.80 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (147.89 DPS) | yes | Master Sergeant's Insignia (18442, -3.09 DPS) [vendor]; Tarnished Locket (279870, -67.45 DPS) [quest]; Kaleidoscope Chain (13084, -70.52 DPS) [world_drop] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.9 stamina points (303.65 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -42.71 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -55.53 DPS) [crafted]; Barbaric Shoulders (5964, -72.71 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 7 · Sergeant · Alliance [vendor] | 11.5 stamina points (185.21 DPS) | yes | Tigerstrike Mantle (13108, -13.89 DPS) [world_drop]; Mourning Shawl (6751, -22.31 DPS) [quest]; Enduring Cape (14763, -31.89 DPS) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | sim-verified (2492.6 DPS) | yes | Green Whelp Armor (7375, -20.96 DPS) [crafted]; Raptor Hide Harness (4455, -28.46 DPS, sim-verified) [crafted]; Dokebi Chestguard (14581, -43.26 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.5 stamina points (184.88 DPS) | yes | Barbaric Bracers (18948, -2.28 DPS) [crafted]; Black Wolf Bracers (3230, -3.87 DPS) [dungeon]; Staghide Armguards (270021, -26.75 DPS) [quest] |
| hands | Operator's Gloves (270045) | Gyrodrillmatic Excavationators [quest] | 13.8 stamina points (222.18 DPS) | yes | Brawler Gloves (720, -7.00 DPS) [world_drop]; Wolfclaw Gloves (1978, -12.08 DPS) [dungeon]; Gloves of Insight (9698, -14.74 DPS) [quest] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 16.3 stamina points (262.25 DPS) | yes | Skulker's Leather Belt (252520, -36.32 DPS, sim-verified) [crafted]; Stalker's Leather Belt (252521, -45.47 DPS) [crafted]; Prowler's Leather Belt (252459, -45.95 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 19.1 stamina points (307.47 DPS) | yes | Duty Bound Leggings (279868, -21.88 DPS) [quest]; Dark Ritual Leggings (270031, -25.88 DPS) [quest]; Brawler's Leather Legguards (252516, -27.22 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.3 stamina points (311.20 DPS) | yes | Harbinger Boots (7754, -22.31 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -62.18 DPS) [quest]; Emblazoned Boots (4051, -91.66 DPS) [world_drop] |
| finger1 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (128.71 DPS) | yes | Defias Renegade Ring (1076, -35.87 DPS) [dungeon]; Slain Baron's Signet (279867, -40.32 DPS) [quest]; Boar Signet (274078, -48.27 DPS) [quest] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (96.53 DPS) | yes | Defias Renegade Ring (1076, -3.69 DPS) [dungeon]; Slain Baron's Signet (279867, -8.14 DPS) [quest]; Boar Signet (274078, -16.09 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -24.86 DPS) [world]; Glimmering Staff (249392, -52.23 DPS) [crafted]; Manual Crowd Pummeler (9449, -182.33 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Spirewind Fetter; wrist: Cultist's Armguards; hands: Operator's Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Plains Ring; finger2: Sustaining Ring; main_hand: Staff of the Shade

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-55230332023102000000-0000000000000000)

Set DPS (verified): 58.3. Weights run: 2.5s. Verify run: 1.2s. 523 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.118 ± 0.007, defense=not significant (0.312 ± 0.079) per rating point (1 rating = 1%, 0.312 per %), dodge=0.118 ± 0.020 per rating point (12 rating = 1%, 1.416 per %), strength=0.042 ± 0.000, agility=0.155 ± 0.013, attack_power=0.021 ± 0.000, hit=not significant (0.058 ± 0.017) per rating point (10 rating = 1%, 0.580 per %), crit=not significant (0.019 ± 0.006) per rating point (14 rating = 1%, 0.264 per %), expertise=1.967 ± 0.184

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 26.5 stamina points (545.54 DPS) | yes | Warden's Wizard Hat (14604, -36.82 DPS) [world_drop]; Brawler's Leather Helm (252512, -41.72 DPS) [crafted]; Nightscape Headband (8176, -59.47 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (229.84 DPS) | yes | River Pride Choker (13087, -41.16 DPS) [world_drop]; Explorers' League Commendation (7746, -44.62 DPS) [quest]; Master Sergeant's Insignia (18442, -44.62 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 27.4 stamina points (563.31 DPS) | yes | Flintrock Shoulders (7755, -90.44 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -139.69 DPS) [dungeon]; Cloudy Gustwoven Spaulders (277043, -212.01 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.9 stamina points (286.69 DPS) | yes | Silky Spider Cape (10776, -7.89 DPS) [dungeon]; Well Oiled Cloak (12254, -10.33 DPS) [vendor]; Heraldic Cloak (8120, -15.71 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 30.5 stamina points (626.98 DPS) | yes | Warden's Wraps (14601, -49.78 DPS) [world_drop]; Robes of the Lich (10762, -64.53 DPS) [dungeon]; Insignia Chestguard (4057, -108.83 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.1 stamina points (289.72 DPS) | yes | Branded Leather Bracers (19508, -23.76 DPS) [dungeon]; Cultist's Armguards (270032, -34.28 DPS) [quest]; Barbaric Bracers (18948, -35.66 DPS) [crafted] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 19.5 stamina points (402.03 DPS) | yes | Warden's Leather Gloves (252527, -3.53 DPS) [crafted]; Bonefingers (10765, -19.73 DPS) [dungeon]; Razzeric's Racing Grips (6727, -20.98 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 19.2 stamina points (394.47 DPS) | yes | Kolkar Hunter's Belt (6788, -25.58 DPS) [quest]; Warden's Leather Belt (252460, -31.63 DPS) [crafted]; Scorpashi Sash (14652, -49.60 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 24.6 stamina points (506.56 DPS) | yes | Warden's Woolies (14605, -9.00 DPS) [world_drop]; Panther Hunter Leggings (4108, -14.71 DPS) [quest]; Imperial Leather Pants (4062, -51.63 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 23.3 stamina points (478.61 DPS) | yes | Worn Running Boots (9398, -19.67 DPS) [dungeon]; Gnomebot Operating Boots (9450, -46.74 DPS) [dungeon]; Skulker's Leather Shoes (252531, -69.56 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (246.96 DPS) | yes | Field Researcher's Loop (281634, -80.57 DPS) [quest]; Plains Ring (2039, -82.32 DPS) [dungeon]; Trader's Ring (15689, -82.32 DPS) [quest] |
| finger2 | Talvash's Gold Ring (9538) | Gnome Improvement [quest] | 10.0 stamina points (205.80 DPS) | yes | Field Researcher's Loop (281634, -39.41 DPS) [quest]; Plains Ring (2039, -41.16 DPS) [dungeon]; Trader's Ring (15689, -41.16 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (4585.2 DPS) | yes | Windweaver Staff (7757, -97.62 DPS) [dungeon]; Mograine's Might (7723, -113.00 DPS) [dungeon]; Manual Crowd Pummeler (9449, -337.25 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Talvash's Gold Ring; trinket1: Arena Master; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 523, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0000000000000000-55230332023102125110-0000000000000000)

Set DPS (verified): 135.1. Weights run: 2.7s. Verify run: 2.8s. 693 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.037, armor=0.160 ± 0.014, defense=not significant (0.415 ± 0.166) per rating point (1 rating = 1%, 0.415 per %), dodge=0.185 ± 0.033 per rating point (12 rating = 1%, 2.219 per %), strength=0.056 ± 0.000, agility=0.203 ± 0.021, attack_power=0.028 ± 0.000, hit=not significant (0.084 ± 0.031) per rating point (10 rating = 1%, 0.839 per %), crit=0.062 ± 0.013 per rating point (14 rating = 1%, 0.874 per %), expertise=3.402 ± 0.331

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.7 stamina points (1234.41 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -90.59 DPS) [vendor]; Knight-Lieutenant's Crackling Leather Helmet (220872, -268.45 DPS) [vendor]; Scorpashi Skullcap (14658, -369.23 DPS, sim-verified) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (404.37 DPS) | yes | Gazlowe's Charm (13088, -94.98 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -138.00 DPS) [world_drop]; Medal of Courage (6723, -144.42 DPS) [quest] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.7 stamina points (914.77 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -57.34 DPS) [vendor]; Penance Spaulders (11963, -124.62 DPS, sim-verified) [quest] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 20.1 stamina points (579.66 DPS) | yes | Sergeant's Cape (18441, -9.23 DPS) [vendor]; Emberscale Cape (11311, -18.45 DPS) [dungeon]; Grovekeeper's Drape (17739, -62.38 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.5 stamina points (1632.12 DPS) | yes | Heraldic Breastplate (8119, -425.26 DPS) [world_drop]; Jinxed Hoodoo Skin (9473, -425.93 DPS, sim-verified) [dungeon]; Knight's Leather Armor (220854, -429.14 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.4 stamina points (734.85 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -119.21 DPS, sim-verified) [vendor]; Serpentskin Bracers (8257, -131.95 DPS) [world_drop]; Warden's Leather Bracers (252541, -133.19 DPS) [crafted] |
| hands | Sergeant Major's Leather Gauntlets (220856) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (9014.4 DPS) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Warden's Leather Gauntlets (252549, -316.21 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.9 stamina points (834.18 DPS) | yes | Skulker's Leather Waistguard (252474, -33.28 DPS) [crafted]; Prowler's Leather Waistguard (252473, -41.82 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -134.78 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 38.6 stamina points (1115.33 DPS) | yes | Knight's Restored Leather Leggings (220882, -183.10 DPS) [vendor]; Scorpashi Leggings (14659, -230.93 DPS) [world_drop]; Knight's Crackling Leather Leggings (220864, -567.54 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (9014.4 DPS) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -438.16 DPS, sim-verified) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (404.37 DPS) | yes | Seafire Band (4549, -100.57 DPS) [quest]; Talvash's Gold Ring (9538, -115.53 DPS) [quest]; Ring of Fortitude (10739, -115.53 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (346.60 DPS) | yes | Seafire Band (4549, -42.81 DPS) [quest]; Talvash's Gold Ring (9538, -57.77 DPS) [quest]; Ring of Fortitude (10739, -57.77 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (9014.4 DPS) | yes | Arena Master (18706, -57.77 DPS) [world]; Smotts' Compass (4130, -195.86 DPS) [quest]; Frozen Heart of the Mountain (249469, -238.13 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (9014.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (9014.4 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -26.14 DPS) [world_drop]; The Hand of Antu'sul (9639, -381.93 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Sergeant Major's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Mark of the Chosen; main_hand: Radiant Staff

No-known-source sample (15 of 693, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 159.6. Weights run: 2.7s. Verify run: 3.5s. 1571 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.028, armor=0.191 ± 0.016, defense=1.033 ± 0.244 per rating point (1 rating = 1%, 1.033 per %), dodge=0.587 ± 0.048 per rating point (12 rating = 1%, 7.049 per %), strength=0.100 ± 0.000, agility=0.505 ± 0.030, attack_power=0.050 ± 0.000, hit=not significant (0.141 ± 0.039) per rating point (10 rating = 1%, 1.413 per %), crit=0.108 ± 0.017 per rating point (14 rating = 1%, 1.518 per %), expertise=5.975 ± 0.421

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 74.4 stamina points (5103.07 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 17.3 stamina points (1188.29 DPS) | yes | Master Sergeant's Insignia (18443, -22.30 DPS) [vendor]; Sunprism Pendant (20649, -292.82 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 51.5 stamina points (3529.60 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -373.95 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 27.1 stamina points (1855.84 DPS) | yes | Sergeant's Cape (16342, -100.98 DPS) [vendor]; Cloak of Earth and Sky (279270, -104.14 DPS) [crafted]; Nightfall Drape (12465, -398.34 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 67.3 stamina points (4617.81 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -208.69 DPS, sim-verified) [dungeon] |
| wrist | Sergeant Major's Dragonhide Armsplints (18454) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 31.7 stamina points (2173.62 DPS) | yes | Knight-Captain's Dragonhide Armsplints (16395, -245.02 DPS) [pvp]; Shadowcraft Bracers (16710, -245.20 DPS) [dungeon]; Arena Bracers (18710, -278.88 DPS, sim-verified) [world] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 43.2 stamina points (2961.91 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Bileblister Claws (275610, -159.99 DPS, sim-verified) [crafted] |
| waist | Elderwild Waistcord (279252) | Leatherworking [crafted] | sim-verified (14743.7 DPS) | yes | Hivethrasher's Girdle (275614, -178.16 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -182.55 DPS) [world_drop]; Marshal's Dragonhide Waistguard (16447, -205.76 DPS) [pvp] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 60.0 stamina points (4115.76 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -290.78 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 47.7 stamina points (3272.59 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -286.66 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (960.23 DPS) | yes | Ring of the Heavens (12056, -68.59 DPS) [dungeon]; Dragonscale Band (12057, -68.59 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -68.59 DPS) [quest] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (960.23 DPS) | yes | Ring of the Heavens (12056, -68.59 DPS) [dungeon]; Dragonscale Band (12057, -68.59 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -68.59 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, -133.84 DPS) [quest]; Arena Grand Master (19024, -133.84 DPS) [quest]; Arena Master (18706, -137.18 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Demon's Blood (10779, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -802.57 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Sergeant Major's Dragonhide Armsplints; hands: Hivethrasher's Claws; waist: Elderwild Waistcord; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Ring of the Dawn; finger2: Ring of Saviors; main_hand: Twig of the World Tree

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 327.1. Weights run: 2.8s. Verify run: 3.5s. 1571 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.008, armor=0.228 ± 0.020, defense=1.608 ± 0.332 per rating point (1 rating = 1%, 1.608 per %), dodge=0.852 ± 0.062 per rating point (12 rating = 1%, 10.228 per %), strength=0.119 ± 0.000, agility=0.754 ± 0.040, attack_power=0.054 ± 0.000, hit=0.305 ± 0.047 per rating point (10 rating = 1%, 3.050 per %), crit=0.118 ± 0.022 per rating point (14 rating = 1%, 1.654 per %), expertise=8.256 ± 0.536

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 85.8 stamina points (7399.48 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 23.1 stamina points (1990.83 DPS) | yes | Master Sergeant's Insignia (18443, -524.29 DPS) [vendor]; Sunprism Pendant (20649, -606.19 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 59.2 stamina points (5110.55 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -397.14 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 29.0 stamina points (2505.72 DPS) | yes | Cloak of Earth and Sky (279270, -142.69 DPS) [crafted]; Sergeant's Cape (16342, -154.70 DPS) [vendor]; Sergeant Major's Cape (16337, -367.43 DPS) [pvp] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 77.9 stamina points (6719.26 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Dawn Armor (252483, -409.06 DPS, sim-verified) [crafted] |
| wrist | Sergeant Major's Dragonhide Armsplints (18454) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 34.5 stamina points (2979.98 DPS) | yes | Arena Bracers (18710, -123.93 DPS) [world]; Deepfury Bracers (13120, -261.51 DPS) [world_drop]; Shadowcraft Bracers (16710, -429.87 DPS, sim-verified) [dungeon] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 48.7 stamina points (4205.26 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -271.48 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 47.5 stamina points (4094.98 DPS) | yes | Serpentine Sash (13118, -75.30 DPS) [world_drop]; Elderwild Waistcord (279252, -101.66 DPS) [crafted]; Wicked Leather Belt (15088, -352.61 DPS) [crafted] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | sim-verified (28225.1 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -369.45 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 55.6 stamina points (4797.16 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -367.87 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.6 stamina points (1342.24 DPS) | yes | Ring of Saviors (1447, -134.50 DPS) [world_drop]; Dragonscale Band (12057, -220.77 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -220.77 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1207.74 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Dragonscale Band (12057, -86.27 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -86.27 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (+465.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Arena Master (18706, -172.53 DPS) [world] |
| trinket2 | Demon's Blood (10779) | You Are Rakh'likh, Demon [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -1222.59 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Sergeant Major's Dragonhide Armsplints; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of the Dawn; trinket2: Demon's Blood; main_hand: Twig of the World Tree

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 36.5. Weights run: 2.3s. Verify run: 1.2s. 199 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.085 ± 0.002, defense=0.152 ± 0.029 per rating point (1 rating = 1%, 0.152 per %), dodge=0.095 ± 0.009 per rating point (12 rating = 1%, 1.140 per %), strength=0.051 ± 0.000, agility=0.149 ± 0.005, attack_power=0.025 ± 0.000, hit=0.044 ± 0.006 per rating point (10 rating = 1%, 0.442 per %), crit=0.018 ± 0.001 per rating point (14 rating = 1%, 0.246 per %), expertise=1.194 ± 0.066

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (296.01 DPS) | yes | Brawler's Leather Hood (252504, -16.10 DPS) [crafted]; Defender's Leather Hood (252447, -28.57 DPS) [crafted]; Totemic Leather Hood (252448, -35.04 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (52.37 DPS) | yes | Scholarly Pendant (277203, -13.27 DPS, sim-verified) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.4 stamina points (149.56 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -46.31 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -49.74 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (122.17 DPS) | yes | Sentry Cloak (2059, -5.07 DPS) [world_drop]; Grave Shroud (279865, -8.72 DPS) [quest]; Firebane Cloak (12979, -17.22 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (288.79 DPS) | yes | Gloomshroud Armor (1489, -32.25 DPS, sim-verified) [dungeon]; Trapper's Leather Armor (252491, -78.54 DPS) [crafted]; Camouflaged Tunic (3585, -79.88 DPS) [quest] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.8 stamina points (140.28 DPS) | yes | Savannah Bracers (15453, -12.48 DPS) [quest]; Bear Bracers (4795, -20.48 DPS, sim-verified) [vendor]; Drakewing Bands (12999, -37.52 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (144.24 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.34 DPS) [crafted]; Defender's Leather Gloves (252496, -6.24 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.2 stamina points (178.10 DPS) | yes | Belt of the Fang (10412, -31.04 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -40.58 DPS) [crafted]; Guardsman Belt (3429, -41.93 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.0 stamina points (222.26 DPS) | yes | Slick Deviate Leggings (6480, -0.80 DPS) [quest]; Smelting Pants (5199, -8.92 DPS) [dungeon]; Brawler's Leather Pants (252500, -22.35 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (238.96 DPS) | yes | Surfer Shoes (276274, -45.37 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -47.47 DPS) [dungeon]; Trapper's Leather Boots (252440, -65.72 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.3 stamina points (84.22 DPS) | yes | Ring of Scorn (3235, -20.70 DPS) [quest]; Ring of the Moon (12052, -35.77 DPS) [world_drop]; Deep Fathom Ring (6463, -36.58 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (79.40 DPS) | yes | Ring of Scorn (3235, -14.00 DPS, sim-verified) [quest]; Ring of the Moon (12052, -30.95 DPS) [world_drop]; Deep Fathom Ring (6463, -31.76 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.0 stamina points (254.16 DPS) | yes | Crescent Staff (6505, -11.40 DPS) [quest]; Hammerbone (270018, -19.33 DPS) [quest]; Rakzur Club (12983, -20.57 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 43.0. Weights run: 2.1s. Verify run: 1.2s. 387 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.097 ± 0.005, defense=0.247 ± 0.061 per rating point (1 rating = 1%, 0.247 per %), dodge=0.104 ± 0.015 per rating point (12 rating = 1%, 1.245 per %), strength=0.048 ± 0.000, agility=0.154 ± 0.009, attack_power=0.024 ± 0.000, hit=0.061 ± 0.011 per rating point (10 rating = 1%, 0.609 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.282 per %), expertise=1.645 ± 0.122

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.7 stamina points (364.52 DPS) | yes | Defender's Leather Helm (252455, -15.52 DPS) [crafted]; Totemic Leather Helm (252456, -24.80 DPS) [crafted]; Trapper's Leather Helm (252513, -24.80 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (147.89 DPS) | yes | Senior Sergeant's Insignia (15200, -3.09 DPS) [vendor]; Kaleidoscope Chain (13084, -70.52 DPS) [world_drop]; Pendant of Myzrael (4614, -83.54 DPS) [dungeon] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.9 stamina points (303.65 DPS) | yes | Tanned Shoulderpads (270023, -41.49 DPS, sim-verified) [quest]; Cloudy Gustwoven Spaulders (277043, -55.53 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -55.53 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 7 · Sergeant · Horde [vendor] | 11.5 stamina points (185.21 DPS) | yes | Tigerstrike Mantle (13108, -13.89 DPS) [world_drop]; Grimsteel Cape (4643, -31.88 DPS) [quest]; Enduring Cape (14763, -31.89 DPS) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | sim-verified (2704.9 DPS) | yes | Green Whelp Armor (7375, -20.96 DPS) [crafted]; Raptor Hide Harness (4455, -28.35 DPS, sim-verified) [crafted]; Dokebi Chestguard (14581, -43.26 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.5 stamina points (184.88 DPS) | yes | Barbaric Bracers (18948, -2.28 DPS) [crafted]; Black Wolf Bracers (3230, -3.87 DPS) [dungeon]; Technician's Bracers (270042, -27.51 DPS) [quest] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 14.3 stamina points (229.46 DPS) | yes | Brawler Gloves (720, -14.28 DPS) [world_drop]; Wolfclaw Gloves (1978, -19.36 DPS) [dungeon]; Emblazoned Gloves (6397, -31.93 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 16.3 stamina points (262.25 DPS) | yes | Larcenist's Belt (279846, -32.49 DPS, sim-verified) [quest]; Skulker's Leather Belt (252520, -40.83 DPS) [crafted]; Stalker's Leather Belt (252521, -45.47 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 19.1 stamina points (307.47 DPS) | yes | Dark Ritual Leggings (270031, -22.10 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -27.22 DPS) [crafted]; Trapper's Leather Legguards (252517, -32.63 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.3 stamina points (311.20 DPS) | yes | Harbinger Boots (7754, -24.67 DPS, sim-verified) [dungeon]; Grizzled Boots (6335, -42.29 DPS) [quest]; Nat Pagle's Extreme Anglin' Boots (19969, -62.18 DPS) [quest] |
| finger1 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.1 stamina points (131.03 DPS) | yes | Defias Renegade Ring (1076, -38.19 DPS) [dungeon]; Slain Baron's Signet (279867, -42.64 DPS) [quest]; Boar Signet (274078, -50.59 DPS) [quest] |
| finger2 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (128.71 DPS) | yes | Defias Renegade Ring (1076, -31.42 DPS, sim-verified) [dungeon]; Slain Baron's Signet (279867, -40.32 DPS) [quest]; Boar Signet (274078, -48.27 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -24.86 DPS) [world]; Glimmering Staff (249392, -52.23 DPS) [crafted]; Manual Crowd Pummeler (9449, -192.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Spirewind Fetter; wrist: Cultist's Armguards; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Seal of Sylvanas; finger2: Plains Ring; main_hand: Staff of the Shade

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-55230332023102000000-0000000000000000)

Set DPS (verified): 58.1. Weights run: 2.5s. Verify run: 1.3s. 521 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.118 ± 0.007, defense=not significant (0.312 ± 0.079) per rating point (1 rating = 1%, 0.312 per %), dodge=0.118 ± 0.020 per rating point (12 rating = 1%, 1.416 per %), strength=0.042 ± 0.000, agility=0.155 ± 0.013, attack_power=0.021 ± 0.000, hit=not significant (0.058 ± 0.017) per rating point (10 rating = 1%, 0.580 per %), crit=not significant (0.019 ± 0.006) per rating point (14 rating = 1%, 0.264 per %), expertise=1.967 ± 0.184

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 26.5 stamina points (545.54 DPS) | yes | Warden's Wizard Hat (14604, -36.82 DPS) [world_drop]; Brawler's Leather Helm (252512, -41.72 DPS) [crafted]; Nightscape Headband (8176, -59.47 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (229.84 DPS) | yes | River Pride Choker (13087, -43.86 DPS, sim-verified) [world_drop]; Senior Sergeant's Insignia (15200, -44.62 DPS) [vendor]; Necklace of Calisea (1714, -65.20 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 27.4 stamina points (563.31 DPS) | yes | Flintrock Shoulders (7755, -134.68 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -139.69 DPS) [dungeon]; Tanned Shoulderpads (270023, -200.47 DPS) [quest] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.9 stamina points (286.69 DPS) | yes | Silky Spider Cape (10776, -7.89 DPS) [dungeon]; Well Oiled Cloak (12254, -10.33 DPS) [vendor]; Heraldic Cloak (8120, -15.71 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 30.5 stamina points (626.98 DPS) | yes | Warden's Wraps (14601, -49.78 DPS) [world_drop]; Robes of the Lich (10762, -64.53 DPS) [dungeon]; Insignia Chestguard (4057, -108.83 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.1 stamina points (289.72 DPS) | yes | Branded Leather Bracers (19508, -23.76 DPS) [dungeon]; Cultist's Armguards (270032, -34.28 DPS) [quest]; Barbaric Bracers (18948, -35.66 DPS) [crafted] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 19.5 stamina points (402.03 DPS) | yes | Bonefingers (10765, -19.73 DPS) [dungeon]; Razzeric's Racing Grips (6727, -20.98 DPS) [quest]; Warden's Leather Gloves (252527, -61.62 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 19.2 stamina points (394.47 DPS) | yes | Kolkar Hunter's Belt (6788, -25.58 DPS) [quest]; Warden's Leather Belt (252460, -31.63 DPS) [crafted]; Scorpashi Sash (14652, -49.60 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 24.6 stamina points (506.56 DPS) | yes | Imperial Leather Pants (4062, -1.70 DPS) [dungeon]; Warden's Woolies (14605, -9.00 DPS) [world_drop]; Panther Hunter Leggings (4108, -14.71 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 23.3 stamina points (478.61 DPS) | yes | Worn Running Boots (9398, -19.67 DPS) [dungeon]; Gnomebot Operating Boots (9450, -46.74 DPS) [dungeon]; Skulker's Leather Shoes (252531, -69.56 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (246.96 DPS) | yes | Seal of Sylvanas (6414, -79.72 DPS) [quest]; Field Researcher's Loop (281634, -80.57 DPS) [quest]; Trader's Ring (15689, -82.32 DPS) [quest] |
| finger2 | Nogg's Gold Ring (9588) | Nogg's Ring Redo [quest] | 10.0 stamina points (205.80 DPS) | yes | Seal of Sylvanas (6414, -38.56 DPS) [quest]; Field Researcher's Loop (281634, -39.41 DPS) [quest]; Trader's Ring (15689, -41.16 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (4948.7 DPS) | yes | Ironshod Bludgeon (9408, -39.35 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -117.00 DPS) [quest]; Gut Ripper (2164, -462.66 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Nogg's Gold Ring; trinket1: Arena Master; main_hand: Cragwood Maul

No-known-source sample (15 of 521, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-55230332023102125110-0000000000000000)

Set DPS (verified): 136.7. Weights run: 2.7s. Verify run: 1.6s. 691 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.037, armor=0.160 ± 0.014, defense=not significant (0.415 ± 0.166) per rating point (1 rating = 1%, 0.415 per %), dodge=0.185 ± 0.033 per rating point (12 rating = 1%, 2.219 per %), strength=0.056 ± 0.000, agility=0.203 ± 0.021, attack_power=0.028 ± 0.000, hit=not significant (0.084 ± 0.031) per rating point (10 rating = 1%, 0.839 per %), crit=0.062 ± 0.013 per rating point (14 rating = 1%, 0.874 per %), expertise=3.402 ± 0.331

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.7 stamina points (1234.41 DPS) | yes | Blood Guard's Leather Headband (220851, -90.59 DPS) [vendor]; Blood Guard's Crackling Leather Helmet (220873, -268.45 DPS) [vendor]; Scorpashi Skullcap (14658, -327.15 DPS, sim-verified) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (404.37 DPS) | yes | Woven Ivy Necklace (19159, -53.03 DPS) [quest]; Gazlowe's Charm (13088, -80.23 DPS) [dungeon]; River Pride Choker (13087, -138.00 DPS) [world_drop] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.7 stamina points (914.77 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Penance Spaulders (11963, -1.92 DPS) [quest]; Blood Guard's Crackling Leather Spaulders (220871, -57.34 DPS) [vendor] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 20.1 stamina points (579.66 DPS) | yes | Sergeant's Cloak (16341, -9.23 DPS) [vendor]; Emberscale Cape (11311, -18.45 DPS) [dungeon]; Grovekeeper's Drape (17739, -62.38 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.5 stamina points (1632.12 DPS) | yes | Heraldic Breastplate (8119, -425.26 DPS) [world_drop]; Stone Guard's Leather Armor (220855, -429.14 DPS) [vendor]; Jinxed Hoodoo Skin (9473, -497.82 DPS, sim-verified) [dungeon] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.4 stamina points (734.85 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -114.88 DPS, sim-verified) [vendor]; Serpentskin Bracers (8257, -131.95 DPS) [world_drop]; Warden's Leather Bracers (252541, -133.19 DPS) [crafted] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | 32.1 stamina points (928.38 DPS) | yes | Skulker's Leather Gauntlets (252548, -110.06 DPS, sim-verified) [crafted]; Prowler's Leather Gauntlets (252547, -136.40 DPS) [crafted]; Wicked Leather Gauntlets (15083, -143.43 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.9 stamina points (834.18 DPS) | yes | Skulker's Leather Waistguard (252474, -33.28 DPS) [crafted]; Prowler's Leather Waistguard (252473, -41.82 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -134.78 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 38.6 stamina points (1115.33 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -125.03 DPS, sim-verified) [vendor]; Stone Guard's Restored Leather Leggings (220883, -183.10 DPS) [vendor]; Scorpashi Leggings (14659, -230.93 DPS) [world_drop] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | 30.2 stamina points (872.89 DPS) | yes | Slitherscale Boots (10801, -38.51 DPS) [dungeon]; Skulker's Leather Boots (252469, -39.16 DPS) [crafted]; First Sergeant's Leather Boots (220861, -41.32 DPS) [vendor] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (404.37 DPS) | yes | Seafire Band (4549, -100.57 DPS) [quest]; Nogg's Gold Ring (9588, -115.53 DPS) [quest]; Sha'ni's Ring (11869, -115.53 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (346.60 DPS) | yes | Seafire Band (4549, -42.81 DPS) [quest]; Nogg's Gold Ring (9588, -57.77 DPS) [quest]; Sha'ni's Ring (11869, -57.77 DPS) [quest] |
| trinket1 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (9541.8 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| trinket2 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (9541.8 DPS) | yes | Arena Master (18706, -57.77 DPS) [world]; Smotts' Compass (4130, -195.86 DPS) [quest]; Rune of the Guard Captain (19120, -209.29 DPS) [quest] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (9541.8 DPS) | yes | Radiant Staff (249453, -8.76 DPS) [crafted]; Dreamstaff (249454, -8.76 DPS) [crafted]; The Hand of Antu'sul (9639, -524.32 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Mark of the Chosen; trinket2: Guardian Talisman

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 161.1. Weights run: 2.7s. Verify run: 3.4s. 1571 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.028, armor=0.191 ± 0.016, defense=1.033 ± 0.244 per rating point (1 rating = 1%, 1.033 per %), dodge=0.587 ± 0.048 per rating point (12 rating = 1%, 7.049 per %), strength=0.100 ± 0.000, agility=0.505 ± 0.030, attack_power=0.050 ± 0.000, hit=not significant (0.141 ± 0.039) per rating point (10 rating = 1%, 1.413 per %), crit=0.108 ± 0.017 per rating point (14 rating = 1%, 1.518 per %), expertise=5.975 ± 0.421

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 74.4 stamina points (5103.07 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 17.3 stamina points (1188.29 DPS) | yes | Woven Ivy Necklace (19159, -149.82 DPS) [quest]; Senior Sergeant's Insignia (16335, -154.93 DPS, sim-verified) [vendor]; Hunter's Insignia Medal (15704, -159.47 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 51.5 stamina points (3529.60 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -413.65 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 27.1 stamina points (1855.84 DPS) | yes | Sergeant's Cloak (18461, -100.98 DPS) [vendor]; Cloak of Earth and Sky (279270, -104.14 DPS) [crafted]; Nightfall Drape (12465, -398.34 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 67.3 stamina points (4617.81 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -276.23 DPS, sim-verified) [dungeon] |
| wrist | First Sergeant's Dragonhide Armguards (18434) | PvP rank 9 · First Sergeant · Horde [vendor] | 31.7 stamina points (2173.62 DPS) | yes | Arena Bracers (18710, -168.20 DPS) [world]; Legionnaire's Dragonhide Armguards (16493, -245.02 DPS) [pvp]; Shadowcraft Bracers (16710, -245.20 DPS) [dungeon] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 43.2 stamina points (2961.91 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Bileblister Claws (275610, +0.00 DPS) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 42.3 stamina points (2899.64 DPS) | yes | Elderwild Waistcord (279252, -6.74 DPS) [crafted]; Serpentine Sash (13118, -189.29 DPS) [world_drop]; General's Dragonhide Belt (16556, -212.50 DPS) [pvp] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 60.0 stamina points (4115.76 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -515.51 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 47.7 stamina points (3272.59 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -378.42 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (960.23 DPS) | yes | Ring of the Heavens (12056, -68.59 DPS) [dungeon]; Dragonscale Band (12057, -68.59 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -68.59 DPS) [quest] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (960.23 DPS) | yes | Ring of the Heavens (12056, -68.59 DPS) [dungeon]; Dragonscale Band (12057, -68.59 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -68.59 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (15671.5 DPS) | yes | Smotts' Compass (4130, -133.84 DPS) [quest]; Arena Grand Master (19024, -133.84 DPS) [quest]; Arena Master (18706, -137.18 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (15671.5 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (15671.5 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -956.09 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: First Sergeant's Dragonhide Armguards; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Ring of the Dawn; finger2: Ring of Saviors; trinket1: Guardian Talisman; trinket2: Mark of the Chosen; main_hand: Twig of the World Tree

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 335.5. Weights run: 2.8s. Verify run: 3.5s. 1571 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.008, armor=0.228 ± 0.020, defense=1.608 ± 0.332 per rating point (1 rating = 1%, 1.608 per %), dodge=0.852 ± 0.062 per rating point (12 rating = 1%, 10.228 per %), strength=0.119 ± 0.000, agility=0.754 ± 0.040, attack_power=0.054 ± 0.000, hit=0.305 ± 0.047 per rating point (10 rating = 1%, 3.050 per %), crit=0.118 ± 0.022 per rating point (14 rating = 1%, 1.654 per %), expertise=8.256 ± 0.536

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 85.8 stamina points (7399.48 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 23.1 stamina points (1990.83 DPS) | yes | Woven Ivy Necklace (19159, -373.16 DPS, sim-verified) [quest]; Senior Sergeant's Insignia (16335, -524.29 DPS) [vendor]; Sunprism Pendant (20649, -606.19 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 59.2 stamina points (5110.55 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -536.41 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 29.0 stamina points (2505.72 DPS) | yes | Cloak of Earth and Sky (279270, -142.69 DPS) [crafted]; Sergeant's Cloak (18461, -154.70 DPS) [vendor]; Deep River Cloak (15789, -478.55 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 77.9 stamina points (6719.26 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Dawn Armor (252483, -656.62 DPS, sim-verified) [crafted] |
| wrist | First Sergeant's Dragonhide Armguards (18434) | PvP rank 9 · First Sergeant · Horde [vendor] | 34.5 stamina points (2979.98 DPS) | yes | Shadowcraft Bracers (16710, -4.63 DPS) [dungeon]; Arena Bracers (18710, -123.93 DPS) [world]; Deepfury Bracers (13120, -261.51 DPS) [world_drop] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 48.7 stamina points (4205.26 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -481.73 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 47.5 stamina points (4094.98 DPS) | yes | Elderwild Waistcord (279252, -101.66 DPS) [crafted]; Serpentine Sash (13118, -296.43 DPS, sim-verified) [world_drop]; Wicked Leather Belt (15088, -352.61 DPS) [crafted] |
| legs | Shadowcraft Pants (16709) | Stratholme: Baron Rivendare [dungeon] | 67.0 stamina points (5782.61 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Elderwild Pants (279254, +0.00 DPS) [crafted] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 55.6 stamina points (4797.16 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -664.82 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.6 stamina points (1342.24 DPS) | yes | Ring of Saviors (1447, -134.50 DPS) [world_drop]; Dragonscale Band (12057, -220.77 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -220.77 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1207.74 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Dragonscale Band (12057, -86.27 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -86.27 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (29358.1 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Arena Master (18706, -172.53 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (29358.1 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (29358.1 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -1385.49 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: First Sergeant's Dragonhide Armguards; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Shadowcraft Pants; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of the Dawn; trinket1: Guardian Talisman; trinket2: Mark of the Chosen; main_hand: Twig of the World Tree

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

