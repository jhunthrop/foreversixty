# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 19.8. Weights run: 3.5s. Verify run: 1.8s. 203 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.084 ± 0.002, defense=0.151 ± 0.028 per rating point (1 rating = 1%, 0.151 per %), dodge=0.095 ± 0.009 per rating point (12 rating = 1%, 1.135 per %), strength=0.063 ± 0.000, agility=0.149 ± 0.005, attack_power=0.032 ± 0.000, hit=0.044 ± 0.006 per rating point (10 rating = 1%, 0.439 per %), crit=0.018 ± 0.001 per rating point (14 rating = 1%, 0.246 per %), expertise=1.187 ± 0.066

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (280.39 DPS) | yes | Brawler's Leather Hood (252504, -15.48 DPS) [crafted]; Defender's Leather Hood (252447, -25.78 DPS) [crafted]; Totemic Leather Hood (252448, -33.37 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.0 stamina points (75.27 DPS) | yes | Erudite's Amulet (277204, -25.30 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -30.11 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.4 stamina points (141.47 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -43.95 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -46.97 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (115.71 DPS) | yes | Sentry Cloak (2059, -4.84 DPS) [world_drop]; Grave Shroud (279865, -7.73 DPS) [quest]; Firebane Cloak (12979, -16.32 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (274.03 DPS) | yes | Gloomshroud Armor (1489, -24.85 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -57.36 DPS) [quest]; Trapper's Leather Armor (252491, -75.19 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (122.24 DPS) | yes | Drakewing Bands (12999, -16.26 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -36.21 DPS) [vendor]; Sanguine Cuffs (14375, -40.44 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (136.45 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.27 DPS) [crafted]; Defender's Leather Gloves (252496, -5.15 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.2 stamina points (169.10 DPS) | yes | Belt of the Fang (10412, -26.14 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -39.00 DPS) [crafted]; Guardsman Belt (3429, -40.27 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.7 stamina points (251.73 DPS) | yes | Defender's Leather Pants (252445, -38.06 DPS, sim-verified) [crafted]; Slick Deviate Leggings (6480, -42.20 DPS) [quest]; Smelting Pants (5199, -49.87 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (226.34 DPS) | yes | Surfer Shoes (276274, -34.71 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -45.16 DPS) [dungeon]; Trapper's Leather Boots (252440, -62.45 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (90.32 DPS) | yes | Blood Ring (4998, -15.05 DPS) [world_drop]; Ring of the Moon (12052, -44.21 DPS) [world_drop]; Deep Fathom Ring (6463, -45.16 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.3 stamina points (79.80 DPS) | yes | Blood Ring (4998, -4.54 DPS) [world_drop]; Ring of the Moon (12052, -33.69 DPS) [world_drop]; Deep Fathom Ring (6463, -34.64 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 17.7 stamina points (266.08 DPS) | yes | Rhahk'Zor's Hammer (5187, -23.22 DPS) [dungeon]; Rakzur Club (12983, -30.57 DPS, sim-verified) [world_drop]; Twisted Chanter's Staff (890, -40.05 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 203, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde; 20428 Caretaker's Cape

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 27.7. Weights run: 3.5s. Verify run: 1.8s. 387 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.096 ± 0.005, defense=0.244 ± 0.060 per rating point (1 rating = 1%, 0.244 per %), dodge=0.103 ± 0.015 per rating point (12 rating = 1%, 1.234 per %), strength=0.057 ± 0.000, agility=0.153 ± 0.009, attack_power=0.029 ± 0.000, hit=0.059 ± 0.010 per rating point (10 rating = 1%, 0.593 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.282 per %), expertise=1.625 ± 0.120

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.6 stamina points (347.66 DPS) | yes | Defender's Leather Helm (252455, -12.98 DPS) [crafted]; Totemic Leather Helm (252456, -23.58 DPS) [crafted]; Trapper's Leather Helm (252513, -23.58 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (142.03 DPS) | yes | Master Sergeant's Insignia (18442, -3.53 DPS) [vendor]; Tarnished Locket (279870, -65.09 DPS) [quest]; Kaleidoscope Chain (13084, -67.51 DPS) [world_drop] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.8 stamina points (290.05 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -40.03 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -53.60 DPS) [crafted]; Barbaric Shoulders (5964, -69.29 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 7 · Sergeant · Alliance [vendor] | 11.5 stamina points (176.89 DPS) | yes | Tigerstrike Mantle (13108, -13.39 DPS) [world_drop]; Mourning Shawl (6751, -21.29 DPS) [quest]; Enduring Cape (14763, -30.20 DPS) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | sim-verified (2250.2 DPS) | yes | Green Whelp Armor (7375, -20.95 DPS) [crafted]; Raptor Hide Harness (4455, -27.71 DPS, sim-verified) [crafted]; Dokebi Chestguard (14581, -42.25 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.5 stamina points (177.11 DPS) | yes | Barbaric Bracers (18948, -2.41 DPS) [crafted]; Black Wolf Bracers (3230, -4.13 DPS) [dungeon]; Technician's Bracers (270042, -26.35 DPS) [quest] |
| hands | Operator's Gloves (270045) | Gyrodrillmatic Excavationators [quest] | 13.8 stamina points (212.83 DPS) | yes | Brawler Gloves (720, -6.50 DPS) [world_drop]; Wolfclaw Gloves (1978, -11.87 DPS) [dungeon]; Gloves of Insight (9698, -15.04 DPS) [quest] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 16.3 stamina points (250.97 DPS) | yes | Skulker's Leather Belt (252520, -33.10 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -43.52 DPS) [crafted]; Stalker's Leather Belt (252521, -44.39 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 19.1 stamina points (294.08 DPS) | yes | Duty Bound Leggings (279868, -21.83 DPS) [quest]; Dark Ritual Leggings (270031, -25.65 DPS) [quest]; Brawler's Leather Legguards (252516, -26.06 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.3 stamina points (296.88 DPS) | yes | Harbinger Boots (7754, -21.00 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -59.06 DPS) [quest]; Emblazoned Boots (4051, -87.59 DPS) [world_drop] |
| finger1 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (123.11 DPS) | yes | Defias Renegade Ring (1076, -34.37 DPS) [dungeon]; Slain Baron's Signet (279867, -38.65 DPS) [quest]; Boar Signet (274078, -46.17 DPS) [quest] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (92.33 DPS) | yes | Defias Renegade Ring (1076, -3.60 DPS) [dungeon]; Slain Baron's Signet (279867, -7.88 DPS) [quest]; Boar Signet (274078, -15.39 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -25.41 DPS) [world]; Glimmering Staff (249392, -47.70 DPS) [crafted]; Manual Crowd Pummeler (9449, -108.81 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Spirewind Fetter; wrist: Cultist's Armguards; hands: Operator's Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Plains Ring; finger2: Sustaining Ring; main_hand: Staff of the Shade

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 0000000000000000-55230332023102000000-0000000000000000)

Set DPS (verified): 48.7. Weights run: 4.0s. Verify run: 1.8s. 523 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.118 ± 0.007, defense=not significant (0.313 ± 0.078) per rating point (1 rating = 1%, 0.313 per %), dodge=0.118 ± 0.020 per rating point (12 rating = 1%, 1.410 per %), strength=0.048 ± 0.000, agility=0.155 ± 0.013, attack_power=0.024 ± 0.000, hit=not significant (0.058 ± 0.016) per rating point (10 rating = 1%, 0.579 per %), crit=not significant (0.020 ± 0.006) per rating point (14 rating = 1%, 0.275 per %), expertise=1.954 ± 0.182

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 26.5 stamina points (526.89 DPS) | yes | Warden's Wizard Hat (14604, -35.69 DPS) [world_drop]; Brawler's Leather Helm (252512, -40.26 DPS) [crafted]; Nightscape Headband (8176, -57.52 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (222.92 DPS) | yes | River Pride Choker (13087, -39.84 DPS) [world_drop]; Explorers' League Commendation (7746, -43.65 DPS) [quest]; Master Sergeant's Insignia (18442, -43.65 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 27.3 stamina points (544.66 DPS) | yes | Flintrock Shoulders (7755, -87.33 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -135.25 DPS) [dungeon]; Cloudy Gustwoven Spaulders (277043, -205.55 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.9 stamina points (277.14 DPS) | yes | Silky Spider Cape (10776, -7.66 DPS) [dungeon]; Well Oiled Cloak (12254, -10.01 DPS) [vendor]; Heraldic Cloak (8120, -15.23 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 30.4 stamina points (605.90 DPS) | yes | Warden's Wraps (14601, -48.67 DPS) [world_drop]; Robes of the Lich (10762, -62.26 DPS) [dungeon]; Insignia Chestguard (4057, -105.63 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.0 stamina points (279.79 DPS) | yes | Branded Leather Bracers (19508, -21.88 DPS) [dungeon]; Cultist's Armguards (270032, -32.50 DPS) [quest]; Barbaric Bracers (18948, -34.00 DPS) [crafted] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 19.5 stamina points (388.21 DPS) | yes | Warden's Leather Gloves (252527, -2.38 DPS) [crafted]; Bonefingers (10765, -19.16 DPS) [dungeon]; Razzeric's Racing Grips (6727, -20.23 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 19.2 stamina points (381.98 DPS) | yes | Kolkar Hunter's Belt (6788, -25.72 DPS) [quest]; Warden's Leather Belt (252460, -30.86 DPS) [crafted]; Scorpashi Sash (14652, -48.55 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 24.5 stamina points (488.94 DPS) | yes | Warden's Woolies (14605, -8.58 DPS) [world_drop]; Panther Hunter Leggings (4108, -14.03 DPS) [quest]; Imperial Leather Pants (4062, -49.35 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 23.2 stamina points (462.94 DPS) | yes | Worn Running Boots (9398, -19.78 DPS) [dungeon]; Gnomebot Operating Boots (9450, -45.85 DPS) [dungeon]; Skulker's Leather Shoes (252531, -67.33 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (239.03 DPS) | yes | Field Researcher's Loop (281634, -77.98 DPS) [quest]; Plains Ring (2039, -79.68 DPS) [dungeon]; Trader's Ring (15689, -79.68 DPS) [quest] |
| finger2 | Talvash's Gold Ring (9538) | Gnome Improvement [quest] | 10.0 stamina points (199.19 DPS) | yes | Field Researcher's Loop (281634, -38.15 DPS) [quest]; Plains Ring (2039, -39.84 DPS) [dungeon]; Trader's Ring (15689, -39.84 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (4401.2 DPS) | yes | Windweaver Staff (7757, -99.27 DPS) [dungeon]; Mograine's Might (7723, -107.97 DPS) [dungeon]; Gut Ripper (2164, -380.63 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Talvash's Gold Ring; trinket1: Arena Master; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 523, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb

### Band 50 (night-elf, 0000000000000000-55230332023102125110-0000000000000000)

Set DPS (verified): 124.1. Weights run: 4.4s. Verify run: 4.2s. 693 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.037, armor=0.160 ± 0.014, defense=not significant (0.412 ± 0.165) per rating point (1 rating = 1%, 0.412 per %), dodge=0.184 ± 0.033 per rating point (12 rating = 1%, 2.209 per %), strength=0.063 ± 0.000, agility=0.203 ± 0.021, attack_power=0.031 ± 0.000, hit=not significant (0.083 ± 0.031) per rating point (10 rating = 1%, 0.831 per %), crit=0.063 ± 0.013 per rating point (14 rating = 1%, 0.883 per %), expertise=3.383 ± 0.329

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.7 stamina points (1197.06 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -86.30 DPS) [vendor]; Knight-Lieutenant's Crackling Leather Helmet (220872, -260.44 DPS) [vendor]; Scorpashi Skullcap (14658, -361.48 DPS, sim-verified) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (392.26 DPS) | yes | Gazlowe's Charm (13088, -92.27 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -133.08 DPS) [world_drop]; Medal of Courage (6723, -140.09 DPS) [quest] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.7 stamina points (888.05 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -56.91 DPS) [vendor]; Penance Spaulders (11963, -123.08 DPS, sim-verified) [quest] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 20.1 stamina points (562.19 DPS) | yes | Sergeant's Cape (18441, -8.94 DPS) [vendor]; Emberscale Cape (11311, -17.89 DPS) [dungeon]; Grovekeeper's Drape (17739, -60.51 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.6 stamina points (1584.89 DPS) | yes | Heraldic Breastplate (8119, -413.82 DPS) [world_drop]; Knight's Leather Armor (220854, -416.62 DPS) [vendor]; Jinxed Hoodoo Skin (9473, -417.52 DPS, sim-verified) [dungeon] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.4 stamina points (712.64 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -117.10 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -127.83 DPS) [crafted]; Serpentskin Bracers (8257, -127.98 DPS) [world_drop] |
| hands | Sergeant Major's Leather Gauntlets (220856) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (8836.4 DPS) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Warden's Leather Gauntlets (252549, -309.81 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 29.0 stamina points (811.28 DPS) | yes | Skulker's Leather Waistguard (252474, -32.29 DPS) [crafted]; Prowler's Leather Waistguard (252473, -40.19 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -133.11 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 38.7 stamina points (1083.30 DPS) | yes | Knight's Restored Leather Leggings (220882, -179.16 DPS) [vendor]; Scorpashi Leggings (14659, -225.75 DPS) [world_drop]; Knight's Crackling Leather Leggings (220864, -557.81 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (8836.4 DPS) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -428.97 DPS, sim-verified) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (392.26 DPS) | yes | Seafire Band (4549, -97.17 DPS) [quest]; Talvash's Gold Ring (9538, -112.07 DPS) [quest]; Ring of Fortitude (10739, -112.07 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (336.22 DPS) | yes | Seafire Band (4549, -41.13 DPS) [quest]; Talvash's Gold Ring (9538, -56.04 DPS) [quest]; Ring of Fortitude (10739, -56.04 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (8836.4 DPS) | yes | Arena Master (18706, -56.04 DPS) [world]; Smotts' Compass (4130, -190.28 DPS) [quest]; Frozen Heart of the Mountain (249469, -231.21 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (8836.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (8836.4 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -14.42 DPS) [world_drop]; The Hand of Antu'sul (9639, -533.24 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Sergeant Major's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Mark of the Chosen; main_hand: Radiant Staff

No-known-source sample (15 of 693, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (night-elf, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 154.6. Weights run: 4.5s. Verify run: 5.0s. 1571 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.029, armor=0.190 ± 0.016, defense=1.026 ± 0.246 per rating point (1 rating = 1%, 1.026 per %), dodge=0.587 ± 0.048 per rating point (12 rating = 1%, 7.040 per %), strength=0.081 ± 0.000, agility=0.503 ± 0.030, attack_power=0.040 ± 0.000, hit=not significant (0.142 ± 0.040) per rating point (10 rating = 1%, 1.419 per %), crit=0.107 ± 0.017 per rating point (14 rating = 1%, 1.501 per %), expertise=6.070 ± 0.426

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 74.2 stamina points (5373.25 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 17.3 stamina points (1248.75 DPS) | yes | Master Sergeant's Insignia (18443, -18.47 DPS) [vendor]; Sunprism Pendant (20649, -305.66 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 51.1 stamina points (3699.69 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Stormshroud Shoulders (15058, -334.88 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | 27.0 stamina points (1954.97 DPS) | yes | Sergeant's Cape (16342, -105.12 DPS) [vendor]; Cloak of Earth and Sky (279270, -108.64 DPS) [crafted]; Nightfall Drape (12465, -418.60 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 67.0 stamina points (4848.49 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -208.12 DPS, sim-verified) [dungeon] |
| wrist | Sergeant Major's Dragonhide Armsplints (18454) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 31.6 stamina points (2290.44 DPS) | yes | Knight-Captain's Dragonhide Armsplints (16395, -258.41 DPS) [pvp]; Shadowcraft Bracers (16710, -260.66 DPS) [dungeon]; Arena Bracers (18710, -276.51 DPS, sim-verified) [world] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 43.1 stamina points (3120.16 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Bileblister Claws (275610, -158.67 DPS, sim-verified) [crafted] |
| waist | Elderwild Waistcord (279252) | Leatherworking [crafted] | sim-verified (14637.3 DPS) | yes | Hivethrasher's Girdle (275614, -176.75 DPS, sim-verified) [crafted]; Serpentine Sash (13118, -206.26 DPS) [world_drop]; Marshal's Dragonhide Waistguard (16447, -217.11 DPS) [pvp] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 59.9 stamina points (4335.89 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -288.63 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 47.5 stamina points (3434.16 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -285.44 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1013.17 DPS) | yes | Ring of the Heavens (12056, -72.37 DPS) [dungeon]; Dragonscale Band (12057, -72.37 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -72.37 DPS) [quest] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (1013.17 DPS) | yes | Ring of the Heavens (12056, -72.37 DPS) [dungeon]; Dragonscale Band (12057, -72.37 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -72.37 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, -141.88 DPS) [quest]; Arena Grand Master (19024, -141.88 DPS) [quest]; Arena Master (18706, -144.74 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Demon's Blood (10779, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -1008.21 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Sergeant Major's Dragonhide Armsplints; hands: Hivethrasher's Claws; waist: Elderwild Waistcord; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Ring of the Dawn; finger2: Ring of Saviors; main_hand: Twig of the World Tree

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60, raid preset (night-elf, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 319.7. Weights run: 4.7s. Verify run: 5.0s. 1571 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.008, armor=0.228 ± 0.021, defense=1.613 ± 0.335 per rating point (1 rating = 1%, 1.613 per %), dodge=0.849 ± 0.063 per rating point (12 rating = 1%, 10.192 per %), strength=0.101 ± 0.000, agility=0.753 ± 0.041, attack_power=0.046 ± 0.000, hit=0.307 ± 0.047 per rating point (10 rating = 1%, 3.074 per %), crit=0.116 ± 0.023 per rating point (14 rating = 1%, 1.622 per %), expertise=8.317 ± 0.543

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 85.8 stamina points (7726.84 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 23.1 stamina points (2083.60 DPS) | yes | Master Sergeant's Insignia (18443, -552.04 DPS) [vendor]; Sunprism Pendant (20649, -638.77 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 59.0 stamina points (5318.62 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -396.30 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 29.0 stamina points (2615.74 DPS) | yes | Cloak of Earth and Sky (279270, -146.17 DPS) [crafted]; Sergeant's Cape (16342, -158.89 DPS) [vendor]; Sergeant Major's Cape (16337, -395.46 DPS) [pvp] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 77.7 stamina points (7002.35 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Dawn Armor (252483, -408.53 DPS, sim-verified) [crafted] |
| wrist | Sergeant Major's Dragonhide Armsplints (18454) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 34.6 stamina points (3114.83 DPS) | yes | Arena Bracers (18710, -130.49 DPS) [world]; Deepfury Bracers (13120, -281.66 DPS) [world_drop]; Shadowcraft Bracers (16710, -427.19 DPS, sim-verified) [dungeon] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 48.8 stamina points (4394.66 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -267.75 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 47.5 stamina points (4279.13 DPS) | yes | Serpentine Sash (13118, -94.60 DPS) [world_drop]; Elderwild Waistcord (279252, -104.83 DPS) [crafted]; Wicked Leather Belt (15088, -369.16 DPS) [crafted] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | sim-verified (28071.8 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -367.13 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 55.5 stamina points (4997.26 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -366.85 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.5 stamina points (1400.34 DPS) | yes | Ring of Saviors (1447, -139.06 DPS) [world_drop]; Dragonscale Band (12057, -229.15 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -229.15 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1261.28 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Dragonscale Band (12057, -90.09 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -90.09 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (+462.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Arena Master (18706, -180.18 DPS) [world] |
| trinket2 | Demon's Blood (10779) | You Are Rakh'likh, Demon [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -1484.76 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Sergeant Major's Dragonhide Armsplints; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of the Dawn; trinket2: Demon's Blood; main_hand: Twig of the World Tree

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 20.7. Weights run: 3.5s. Verify run: 1.7s. 199 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.084 ± 0.002, defense=0.151 ± 0.028 per rating point (1 rating = 1%, 0.151 per %), dodge=0.095 ± 0.009 per rating point (12 rating = 1%, 1.135 per %), strength=0.063 ± 0.000, agility=0.149 ± 0.005, attack_power=0.032 ± 0.000, hit=0.044 ± 0.006 per rating point (10 rating = 1%, 0.439 per %), crit=0.018 ± 0.001 per rating point (14 rating = 1%, 0.246 per %), expertise=1.187 ± 0.066

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (280.39 DPS) | yes | Brawler's Leather Hood (252504, -15.48 DPS) [crafted]; Defender's Leather Hood (252447, -25.78 DPS) [crafted]; Totemic Leather Hood (252448, -33.37 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (49.63 DPS) | yes | Scholarly Pendant (277203, -11.26 DPS, sim-verified) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.4 stamina points (141.47 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -43.95 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -46.97 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (115.71 DPS) | yes | Sentry Cloak (2059, -4.84 DPS) [world_drop]; Grave Shroud (279865, -7.73 DPS) [quest]; Firebane Cloak (12979, -16.32 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (274.03 DPS) | yes | Gloomshroud Armor (1489, -27.53 DPS, sim-verified) [dungeon]; Trapper's Leather Armor (252491, -75.19 DPS) [crafted]; Camouflaged Tunic (3585, -76.40 DPS) [quest] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.8 stamina points (132.76 DPS) | yes | Savannah Bracers (15453, -11.79 DPS) [quest]; Bear Bracers (4795, -17.33 DPS, sim-verified) [vendor]; Drakewing Bands (12999, -35.55 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (136.45 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.27 DPS) [crafted]; Defender's Leather Gloves (252496, -5.15 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.2 stamina points (169.10 DPS) | yes | Belt of the Fang (10412, -28.11 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -39.00 DPS) [crafted]; Guardsman Belt (3429, -40.27 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.1 stamina points (211.91 DPS) | yes | Smelting Pants (5199, -10.05 DPS) [dungeon]; Slick Deviate Leggings (6480, -13.58 DPS, sim-verified) [quest]; Brawler's Leather Pants (252500, -21.77 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (226.34 DPS) | yes | Surfer Shoes (276274, -39.15 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -45.16 DPS) [dungeon]; Trapper's Leather Boots (252440, -62.45 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.3 stamina points (79.80 DPS) | yes | Ring of Scorn (3235, -19.59 DPS) [quest]; Ring of the Moon (12052, -33.69 DPS) [world_drop]; Deep Fathom Ring (6463, -34.64 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (75.27 DPS) | yes | Ring of Scorn (3235, -12.20 DPS, sim-verified) [quest]; Ring of the Moon (12052, -29.16 DPS) [world_drop]; Deep Fathom Ring (6463, -30.11 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 17.7 stamina points (266.08 DPS) | yes | Crescent Staff (6505, -11.97 DPS, sim-verified) [quest]; Hammerbone (270018, -15.51 DPS) [quest]; Rakzur Club (12983, -20.55 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 199, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion; 21566 Rune of Perfection

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 27.9. Weights run: 3.5s. Verify run: 1.8s. 387 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.096 ± 0.005, defense=0.244 ± 0.060 per rating point (1 rating = 1%, 0.244 per %), dodge=0.103 ± 0.015 per rating point (12 rating = 1%, 1.234 per %), strength=0.057 ± 0.000, agility=0.153 ± 0.009, attack_power=0.029 ± 0.000, hit=0.059 ± 0.010 per rating point (10 rating = 1%, 0.593 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.282 per %), expertise=1.625 ± 0.120

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.6 stamina points (347.66 DPS) | yes | Defender's Leather Helm (252455, -12.98 DPS) [crafted]; Totemic Leather Helm (252456, -23.58 DPS) [crafted]; Trapper's Leather Helm (252513, -23.58 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (142.03 DPS) | yes | Senior Sergeant's Insignia (15200, -3.53 DPS) [vendor]; Kaleidoscope Chain (13084, -67.51 DPS) [world_drop]; Pendant of Myzrael (4614, -80.48 DPS) [dungeon] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.8 stamina points (290.05 DPS) | yes | Tanned Shoulderpads (270023, -38.83 DPS, sim-verified) [quest]; Cloudy Gustwoven Spaulders (277043, -53.60 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -53.60 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 7 · Sergeant · Horde [vendor] | 11.5 stamina points (176.89 DPS) | yes | Tigerstrike Mantle (13108, -13.39 DPS) [world_drop]; Enduring Cape (14763, -30.20 DPS) [world_drop]; Grimsteel Cape (4643, -30.49 DPS) [quest] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | sim-verified (2443.8 DPS) | yes | Green Whelp Armor (7375, -20.95 DPS) [crafted]; Raptor Hide Harness (4455, -27.25 DPS, sim-verified) [crafted]; Dokebi Chestguard (14581, -42.25 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.5 stamina points (177.11 DPS) | yes | Barbaric Bracers (18948, -2.41 DPS) [crafted]; Black Wolf Bracers (3230, -4.13 DPS) [dungeon]; Technician's Bracers (270042, -26.35 DPS) [quest] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 14.2 stamina points (218.72 DPS) | yes | Brawler Gloves (720, -12.39 DPS) [world_drop]; Wolfclaw Gloves (1978, -17.76 DPS) [dungeon]; Emblazoned Gloves (6397, -30.44 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 16.3 stamina points (250.97 DPS) | yes | Larcenist's Belt (279846, -31.41 DPS, sim-verified) [quest]; Skulker's Leather Belt (252520, -39.09 DPS) [crafted]; Prowler's Leather Belt (252459, -43.52 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 19.1 stamina points (294.08 DPS) | yes | Dark Ritual Leggings (270031, -22.65 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -26.06 DPS) [crafted]; Trapper's Leather Legguards (252517, -32.25 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.3 stamina points (296.88 DPS) | yes | Harbinger Boots (7754, -20.79 DPS, sim-verified) [dungeon]; Grizzled Boots (6335, -39.94 DPS) [quest]; Nat Pagle's Extreme Anglin' Boots (19969, -59.06 DPS) [quest] |
| finger1 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.2 stamina points (125.76 DPS) | yes | Defias Renegade Ring (1076, -37.03 DPS) [dungeon]; Slain Baron's Signet (279867, -41.30 DPS) [quest]; Boar Signet (274078, -48.82 DPS) [quest] |
| finger2 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (123.11 DPS) | yes | Defias Renegade Ring (1076, -28.54 DPS, sim-verified) [dungeon]; Slain Baron's Signet (279867, -38.65 DPS) [quest]; Boar Signet (274078, -46.17 DPS) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -25.41 DPS) [world]; Glimmering Staff (249392, -47.70 DPS) [crafted]; Manual Crowd Pummeler (9449, -114.43 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Spirewind Fetter; wrist: Cultist's Armguards; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Seal of Sylvanas; finger2: Plains Ring; main_hand: Staff of the Shade

No-known-source sample (15 of 387, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-55230332023102000000-0000000000000000)

Set DPS (verified): 48.5. Weights run: 4.0s. Verify run: 1.8s. 521 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.118 ± 0.007, defense=not significant (0.313 ± 0.078) per rating point (1 rating = 1%, 0.313 per %), dodge=0.118 ± 0.020 per rating point (12 rating = 1%, 1.410 per %), strength=0.048 ± 0.000, agility=0.155 ± 0.013, attack_power=0.024 ± 0.000, hit=not significant (0.058 ± 0.016) per rating point (10 rating = 1%, 0.579 per %), crit=not significant (0.020 ± 0.006) per rating point (14 rating = 1%, 0.275 per %), expertise=1.954 ± 0.182

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 26.5 stamina points (526.89 DPS) | yes | Warden's Wizard Hat (14604, -35.69 DPS) [world_drop]; Brawler's Leather Helm (252512, -40.26 DPS) [crafted]; Nightscape Headband (8176, -57.52 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (222.92 DPS) | yes | River Pride Choker (13087, -42.09 DPS, sim-verified) [world_drop]; Senior Sergeant's Insignia (15200, -43.65 DPS) [vendor]; Necklace of Calisea (1714, -63.57 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 27.3 stamina points (544.66 DPS) | yes | Flintrock Shoulders (7755, -128.13 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -135.25 DPS) [dungeon]; Tanned Shoulderpads (270023, -194.35 DPS) [quest] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.9 stamina points (277.14 DPS) | yes | Silky Spider Cape (10776, -7.66 DPS) [dungeon]; Well Oiled Cloak (12254, -10.01 DPS) [vendor]; Heraldic Cloak (8120, -15.23 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 30.4 stamina points (605.90 DPS) | yes | Warden's Wraps (14601, -48.67 DPS) [world_drop]; Robes of the Lich (10762, -62.26 DPS) [dungeon]; Insignia Chestguard (4057, -105.63 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.0 stamina points (279.79 DPS) | yes | Branded Leather Bracers (19508, -21.88 DPS) [dungeon]; Cultist's Armguards (270032, -32.50 DPS) [quest]; Barbaric Bracers (18948, -34.00 DPS) [crafted] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 19.5 stamina points (388.21 DPS) | yes | Bonefingers (10765, -19.16 DPS) [dungeon]; Razzeric's Racing Grips (6727, -20.23 DPS) [quest]; Warden's Leather Gloves (252527, -56.92 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 19.2 stamina points (381.98 DPS) | yes | Kolkar Hunter's Belt (6788, -25.72 DPS) [quest]; Warden's Leather Belt (252460, -30.86 DPS) [crafted]; Scorpashi Sash (14652, -48.55 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 24.5 stamina points (488.94 DPS) | yes | Imperial Leather Pants (4062, -1.56 DPS) [dungeon]; Warden's Woolies (14605, -8.58 DPS) [world_drop]; Panther Hunter Leggings (4108, -14.03 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 23.2 stamina points (462.94 DPS) | yes | Worn Running Boots (9398, -19.78 DPS) [dungeon]; Gnomebot Operating Boots (9450, -45.85 DPS) [dungeon]; Skulker's Leather Shoes (252531, -67.33 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (239.03 DPS) | yes | Seal of Sylvanas (6414, -76.82 DPS) [quest]; Field Researcher's Loop (281634, -77.98 DPS) [quest]; Trader's Ring (15689, -79.68 DPS) [quest] |
| finger2 | Nogg's Gold Ring (9588) | Nogg's Ring Redo [quest] | 10.0 stamina points (199.19 DPS) | yes | Seal of Sylvanas (6414, -36.98 DPS) [quest]; Field Researcher's Loop (281634, -38.15 DPS) [quest]; Trader's Ring (15689, -39.84 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (4750.0 DPS) | yes | Ironshod Bludgeon (9408, -32.35 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -112.38 DPS) [quest]; Gut Ripper (2164, -478.73 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Nogg's Gold Ring; trinket1: Arena Master; main_hand: Cragwood Maul

No-known-source sample (15 of 521, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-55230332023102125110-0000000000000000)

Set DPS (verified): 126.3. Weights run: 4.4s. Verify run: 4.2s. 691 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.037, armor=0.160 ± 0.014, defense=not significant (0.412 ± 0.165) per rating point (1 rating = 1%, 0.412 per %), dodge=0.184 ± 0.033 per rating point (12 rating = 1%, 2.209 per %), strength=0.063 ± 0.000, agility=0.203 ± 0.021, attack_power=0.031 ± 0.000, hit=not significant (0.083 ± 0.031) per rating point (10 rating = 1%, 0.831 per %), crit=0.063 ± 0.013 per rating point (14 rating = 1%, 0.883 per %), expertise=3.383 ± 0.329

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.7 stamina points (1197.06 DPS) | yes | Blood Guard's Leather Headband (220851, -86.30 DPS) [vendor]; Blood Guard's Crackling Leather Helmet (220873, -260.44 DPS) [vendor]; Scorpashi Skullcap (14658, -292.24 DPS, sim-verified) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (392.26 DPS) | yes | Woven Ivy Necklace (19159, -50.27 DPS) [quest]; Gazlowe's Charm (13088, -77.05 DPS) [dungeon]; River Pride Choker (13087, -133.08 DPS) [world_drop] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.7 stamina points (888.05 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Penance Spaulders (11963, -2.91 DPS) [quest]; Blood Guard's Crackling Leather Spaulders (220871, -56.91 DPS) [vendor] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 20.1 stamina points (562.19 DPS) | yes | Sergeant's Cloak (16341, -8.94 DPS) [vendor]; Emberscale Cape (11311, -17.89 DPS) [dungeon]; Grovekeeper's Drape (17739, -60.51 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.6 stamina points (1584.89 DPS) | yes | Jinxed Hoodoo Skin (9473, -406.74 DPS, sim-verified) [dungeon]; Heraldic Breastplate (8119, -413.82 DPS) [world_drop]; Stone Guard's Leather Armor (220855, -416.62 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.4 stamina points (712.64 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -34.19 DPS) [vendor]; Warden's Leather Bracers (252541, -127.83 DPS) [crafted]; Serpentskin Bracers (8257, -127.98 DPS) [world_drop] |
| hands | First Sergeant's Leather Gauntlets (220857) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-verified (9379.8 DPS) | yes | Prowler's Leather Gauntlets (252547, +0.00 DPS) [crafted]; Skulker's Leather Gauntlets (252548, +0.00 DPS) [crafted]; Warden's Leather Gauntlets (252549, -195.53 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 29.0 stamina points (811.28 DPS) | yes | Skulker's Leather Waistguard (252474, -32.29 DPS) [crafted]; Prowler's Leather Waistguard (252473, -40.19 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -133.11 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 38.7 stamina points (1083.30 DPS) | yes | Stone Guard's Restored Leather Leggings (220883, -179.16 DPS) [vendor]; Scorpashi Leggings (14659, -225.75 DPS) [world_drop]; Stone Guard's Crackling Leather Leggings (220865, -574.43 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Leather Boots (220861) | PvP rank 9 · First Sergeant · Horde [vendor] | sim-verified (9379.8 DPS) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -368.23 DPS, sim-verified) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (392.26 DPS) | yes | Seafire Band (4549, -97.17 DPS) [quest]; Nogg's Gold Ring (9588, -112.07 DPS) [quest]; Sha'ni's Ring (11869, -112.07 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (336.22 DPS) | yes | Seafire Band (4549, -41.13 DPS) [quest]; Nogg's Gold Ring (9588, -56.04 DPS) [quest]; Sha'ni's Ring (11869, -56.04 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (9379.8 DPS) | yes | Arena Master (18706, -56.04 DPS) [world]; Smotts' Compass (4130, -147.91 DPS, sim-verified) [quest]; Rune of the Guard Captain (19120, -199.06 DPS) [quest] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (9379.8 DPS) | yes | Arena Master (18706, +0.00 DPS) [world]; Rune of the Guard Captain (19120, +0.00 DPS) [quest]; Smotts' Compass (4130, -137.88 DPS, sim-verified) [quest] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (9379.8 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Cragwood Maul (11265, -8.10 DPS) [quest]; The Hand of Antu'sul (9639, -458.99 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: First Sergeant's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: First Sergeant's Leather Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Mark of the Chosen; main_hand: Radiant Staff

No-known-source sample (15 of 691, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 156.1. Weights run: 4.5s. Verify run: 4.9s. 1571 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.029, armor=0.190 ± 0.016, defense=1.026 ± 0.246 per rating point (1 rating = 1%, 1.026 per %), dodge=0.587 ± 0.048 per rating point (12 rating = 1%, 7.040 per %), strength=0.081 ± 0.000, agility=0.503 ± 0.030, attack_power=0.040 ± 0.000, hit=not significant (0.142 ± 0.040) per rating point (10 rating = 1%, 1.419 per %), crit=0.107 ± 0.017 per rating point (14 rating = 1%, 1.501 per %), expertise=6.070 ± 0.426

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 74.2 stamina points (5373.25 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 17.3 stamina points (1248.75 DPS) | yes | Senior Sergeant's Insignia (16335, -153.08 DPS, sim-verified) [vendor]; Woven Ivy Necklace (19159, -162.57 DPS) [quest]; Hunter's Insignia Medal (15704, -163.21 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 51.1 stamina points (3699.69 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Stormshroud Shoulders (15058, -356.28 DPS, sim-verified) [crafted] |
| back | Fel Cape (279269) | Tailoring [crafted] | 27.0 stamina points (1954.97 DPS) | yes | Sergeant's Cloak (18461, -105.12 DPS) [vendor]; Cloak of Earth and Sky (279270, -108.64 DPS) [crafted]; Nightfall Drape (12465, -418.60 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 67.0 stamina points (4848.49 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -275.47 DPS, sim-verified) [dungeon] |
| wrist | First Sergeant's Dragonhide Armguards (18434) | PvP rank 9 · First Sergeant · Horde [vendor] | 31.6 stamina points (2290.44 DPS) | yes | Arena Bracers (18710, -177.84 DPS) [world]; Legionnaire's Dragonhide Armguards (16493, -258.41 DPS) [pvp]; Shadowcraft Bracers (16710, -260.66 DPS) [dungeon] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 43.1 stamina points (3120.16 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Bileblister Claws (275610, +0.00 DPS) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 42.2 stamina points (3054.84 DPS) | yes | Elderwild Waistcord (279252, -6.80 DPS) [crafted]; Serpentine Sash (13118, -213.07 DPS) [world_drop]; General's Dragonhide Belt (16556, -223.91 DPS) [pvp] |
| legs | Elderwild Pants (279254) | Leatherworking [crafted] | 59.9 stamina points (4335.89 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -511.54 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 47.5 stamina points (3434.16 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -376.40 DPS, sim-verified) [dungeon] |
| finger1 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1013.17 DPS) | yes | Ring of the Heavens (12056, -72.37 DPS) [dungeon]; Dragonscale Band (12057, -72.37 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -72.37 DPS) [quest] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 14.0 stamina points (1013.17 DPS) | yes | Ring of the Heavens (12056, -72.37 DPS) [dungeon]; Dragonscale Band (12057, -72.37 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -72.37 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (15559.0 DPS) | yes | Smotts' Compass (4130, -141.88 DPS) [quest]; Arena Grand Master (19024, -141.88 DPS) [quest]; Arena Master (18706, -144.74 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (15559.0 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (15559.0 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -1232.23 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: First Sergeant's Dragonhide Armguards; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Elderwild Pants; feet: Dawn Treaders; finger1: Ring of the Dawn; finger2: Ring of Saviors; main_hand: Twig of the World Tree

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 328.0. Weights run: 4.7s. Verify run: 5.0s. 1571 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.008, armor=0.228 ± 0.021, defense=1.613 ± 0.335 per rating point (1 rating = 1%, 1.613 per %), dodge=0.849 ± 0.063 per rating point (12 rating = 1%, 10.192 per %), strength=0.101 ± 0.000, agility=0.753 ± 0.041, attack_power=0.046 ± 0.000, hit=0.307 ± 0.047 per rating point (10 rating = 1%, 3.074 per %), crit=0.116 ± 0.023 per rating point (14 rating = 1%, 1.622 per %), expertise=8.317 ± 0.543

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 85.8 stamina points (7726.84 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 23.1 stamina points (2083.60 DPS) | yes | Woven Ivy Necklace (19159, -370.18 DPS, sim-verified) [quest]; Senior Sergeant's Insignia (16335, -552.04 DPS) [vendor]; Sunprism Pendant (20649, -638.77 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 59.0 stamina points (5318.62 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -534.51 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 29.0 stamina points (2615.74 DPS) | yes | Cloak of Earth and Sky (279270, -146.17 DPS) [crafted]; Sergeant's Cloak (18461, -158.89 DPS) [vendor]; Deep River Cloak (15789, -498.53 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 77.7 stamina points (7002.35 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Dawn Armor (252483, -654.32 DPS, sim-verified) [crafted] |
| wrist | First Sergeant's Dragonhide Armguards (18434) | PvP rank 9 · First Sergeant · Horde [vendor] | 34.6 stamina points (3114.83 DPS) | yes | Shadowcraft Bracers (16710, -6.55 DPS) [dungeon]; Arena Bracers (18710, -130.49 DPS) [world]; Deepfury Bracers (13120, -281.66 DPS) [world_drop] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 48.8 stamina points (4394.66 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Timbermaw Brawlers (19049, -476.94 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 47.5 stamina points (4279.13 DPS) | yes | Elderwild Waistcord (279252, -104.83 DPS) [crafted]; Serpentine Sash (13118, -293.70 DPS, sim-verified) [world_drop]; Wicked Leather Belt (15088, -369.16 DPS) [crafted] |
| legs | Shadowcraft Pants (16709) | Stratholme: Baron Rivendare [dungeon] | 67.1 stamina points (6041.80 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Elderwild Pants (279254, +0.00 DPS) [crafted] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 55.5 stamina points (4997.26 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -661.80 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.5 stamina points (1400.34 DPS) | yes | Ring of Saviors (1447, -139.06 DPS) [world_drop]; Dragonscale Band (12057, -229.15 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -229.15 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) (or Ring of Saviors (1447)) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1261.28 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Dragonscale Band (12057, -90.09 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -90.09 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (29201.3 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Arena Master (18706, -180.18 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (29201.3 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (29201.3 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Baron Charr's Sceptre (18671, -1783.65 DPS, sim-verified) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: First Sergeant's Dragonhide Armguards; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Shadowcraft Pants; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of the Dawn; main_hand: Twig of the World Tree

No-known-source sample (15 of 1571, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

