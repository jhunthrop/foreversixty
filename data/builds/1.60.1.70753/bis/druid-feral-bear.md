# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 18.7. Weights run: 2.8s. Verify run: 1.9s. 230 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.085 ± 0.002, defense=0.133 ± 0.027 per rating point (1 rating = 1%, 0.133 per %), dodge=0.098 ± 0.008 per rating point (12 rating = 1%, 1.176 per %), strength=0.060 ± 0.000, agility=0.146 ± 0.005, attack_power=0.030 ± 0.000, hit=0.051 ± 0.006 per rating point (10 rating = 1%, 0.508 per %), crit=0.018 ± 0.001 per rating point (14 rating = 1%, 0.257 per %), expertise=1.273 ± 0.062

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.7 stamina points (286.91 DPS) | yes | Brawler's Leather Hood (252504, -15.80 DPS) [crafted]; Defender's Leather Hood (252447, -26.42 DPS) [crafted]; Totemic Leather Hood (252448, -33.78 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.8 stamina points (89.53 DPS) | yes | Erudite's Amulet (277204, -28.19 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -43.38 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.4 stamina points (145.16 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -45.06 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -48.35 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (118.42 DPS) | yes | Sentry Cloak (2059, -5.08 DPS) [world_drop]; Firebane Cloak (12979, -16.69 DPS) [world_drop]; Cape of the Brotherhood (5193, -33.60 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (280.55 DPS) | yes | Gloomshroud Armor (1489, -18.19 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -58.68 DPS) [quest]; Bloody Apron (6226, -71.84 DPS) [dungeon] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (125.26 DPS) | yes | Drakewing Bands (12999, -13.44 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -37.15 DPS) [vendor]; Sanguine Cuffs (14375, -41.52 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (139.78 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.31 DPS) [crafted]; Defender's Leather Gloves (252496, -5.31 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.2 stamina points (172.94 DPS) | yes | Belt of the Fang (10412, -17.86 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -39.69 DPS) [crafted]; Guardsman Belt (3429, -41.00 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.1 stamina points (216.72 DPS) | yes | Smelting Pants (5199, -22.23 DPS, sim-verified) [dungeon]; Brawler's Leather Pants (252500, -22.28 DPS) [crafted]; Leggings of the Fang (10410, -28.43 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.1 stamina points (231.62 DPS) | yes | Surfer Shoes (276274, -34.26 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -46.13 DPS) [dungeon]; Trapper's Leather Boots (252440, -63.76 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (76.91 DPS) | yes | Slain Baron's Signet (279867, -28.72 DPS) [quest]; Ring of the Moon (12052, -29.84 DPS) [world_drop]; Deep Fathom Ring (6463, -30.76 DPS) [dungeon] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 4.0 stamina points (61.53 DPS) | yes | Slain Baron's Signet (279867, -10.70 DPS, sim-verified) [quest]; Ring of the Moon (12052, -14.46 DPS) [world_drop]; Deep Fathom Ring (6463, -15.38 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 15.9 stamina points (244.39 DPS) | yes | Rhahk'Zor's Hammer (5187, -2.60 DPS) [dungeon]; Staff of the Blessed Seer (2271, -5.06 DPS) [dungeon]; Twisted Chanter's Staff (890, -20.46 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Sustaining Ring; main_hand: Rakzur Club

No-known-source sample (15 of 230, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 26.2. Weights run: 2.6s. Verify run: 1.9s. 414 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.090 ± 0.005, defense=0.462 ± 0.068 per rating point (1 rating = 1%, 0.462 per %), dodge=0.110 ± 0.014 per rating point (12 rating = 1%, 1.315 per %), strength=0.052 ± 0.000, agility=0.148 ± 0.009, attack_power=0.026 ± 0.000, hit=not significant (0.041 ± 0.011) per rating point (10 rating = 1%, 0.410 per %), crit=0.019 ± 0.002 per rating point (14 rating = 1%, 0.260 per %), expertise=1.604 ± 0.122

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.0 stamina points (343.88 DPS) | yes | Defender's Leather Helm (252455, -13.26 DPS) [crafted]; Totemic Leather Helm (252456, -23.06 DPS) [crafted]; Trapper's Leather Helm (252513, -23.06 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (143.67 DPS) | yes | Master Sergeant's Insignia (18442, -3.27 DPS) [vendor]; Tarnished Locket (279870, -52.52 DPS) [quest]; Kaleidoscope Chain (13084, -68.78 DPS) [world_drop] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.4 stamina points (286.45 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -40.52 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -54.33 DPS) [crafted]; Barbaric Shoulders (5964, -70.46 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 7 · Sergeant · Alliance [vendor] | 11.3 stamina points (176.93 DPS) | yes | Tigerstrike Mantle (13108, -14.15 DPS) [world_drop]; Mourning Shawl (6751, -21.22 DPS) [quest]; Enduring Cape (14763, -30.57 DPS) [world_drop] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.8 stamina points (340.49 DPS) | yes | Spirewind Fetter (9406, -5.81 DPS) [dungeon]; Green Whelp Armor (7375, -25.58 DPS) [crafted]; Dokebi Chestguard (14581, -46.80 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.2 stamina points (175.10 DPS) | yes | Barbaric Bracers (18948, -2.98 DPS) [crafted]; Black Wolf Bracers (3230, -3.86 DPS) [dungeon]; Glowing Leather Bracers (2017, -36.29 DPS) [world] |
| hands | Brawler Gloves (720) | World drop [world_drop] | 13.0 stamina points (202.84 DPS) | yes | Wolfclaw Gloves (1978, -5.62 DPS) [dungeon]; Gloves of Insight (9698, -7.94 DPS) [quest]; Emblazoned Gloves (6397, -17.63 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.9 stamina points (247.65 DPS) | yes | Skulker's Leather Belt (252520, -24.15 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -44.35 DPS) [crafted]; Stalker's Leather Belt (252521, -44.78 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.5 stamina points (288.71 DPS) | yes | Dark Ritual Leggings (270031, -24.36 DPS, sim-verified) [quest]; Brawler's Leather Legguards (252516, -26.59 DPS) [crafted]; Trapper's Leather Legguards (252517, -32.31 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.8 stamina points (293.98 DPS) | yes | Harbinger Boots (7754, -20.83 DPS, sim-verified) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -56.20 DPS) [quest]; Emblazoned Boots (4051, -87.94 DPS) [world_drop] |
| finger1 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (124.80 DPS) | yes | Blood Ring (4998, -46.80 DPS) [world_drop]; Heart Ring (5001, -46.80 DPS) [world_drop]; Agamaggan's Clutch (6693, -46.80 DPS) [dungeon] |
| finger2 | Defias Renegade Ring (1076) | The Stockade: Defias Prisoner [dungeon] | 6.6 stamina points (102.83 DPS) | yes | Heart Ring (5001, -21.97 DPS, sim-verified) [world_drop]; Blood Ring (4998, -24.83 DPS) [world_drop]; Agamaggan's Clutch (6693, -24.83 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (2208.6 DPS) | yes | Staff of the Shade (2549, -27.54 DPS, sim-verified) [dungeon]; Slaghammer (1976, -35.95 DPS) [dungeon]; Dense Triangle Mace (3203, -41.53 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Brawler Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Plains Ring; finger2: Defias Renegade Ring; main_hand: Spine of the Devourer

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (night-elf, 0000000000000000-55230332023102000000-0000000000000000)

Set DPS (verified): 48.0. Weights run: 3.1s. Verify run: 2.0s. 549 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.018, armor=0.123 ± 0.007, defense=not significant (0.319 ± 0.092) per rating point (1 rating = 1%, 0.319 per %), dodge=0.131 ± 0.021 per rating point (12 rating = 1%, 1.572 per %), strength=0.049 ± 0.000, agility=0.157 ± 0.013, attack_power=0.025 ± 0.000, hit=0.067 ± 0.016 per rating point (10 rating = 1%, 0.670 per %), crit=not significant (0.022 ± 0.006) per rating point (14 rating = 1%, 0.311 per %), expertise=2.068 ± 0.178

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 27.0 stamina points (553.97 DPS) | yes | Brawler's Leather Helm (252512, -42.61 DPS) [crafted]; Nightscape Headband (8176, -59.48 DPS) [crafted]; Warden's Wizard Hat (14604, -60.11 DPS, sim-verified) [world_drop] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (229.57 DPS) | yes | River Pride Choker (13087, -41.01 DPS) [world_drop]; Explorers' League Commendation (7746, -45.03 DPS) [quest]; Master Sergeant's Insignia (18442, -45.03 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 27.9 stamina points (572.40 DPS) | yes | Flintrock Shoulders (7755, -130.39 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -141.28 DPS) [dungeon]; Cloudy Gustwoven Spaulders (277043, -213.96 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 14.1 stamina points (288.78 DPS) | yes | Silky Spider Cape (10776, -7.78 DPS) [dungeon]; Well Oiled Cloak (12254, -10.31 DPS) [vendor]; Heraldic Cloak (8120, -15.44 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 31.1 stamina points (637.97 DPS) | yes | Robes of the Lich (10762, -70.90 DPS) [dungeon]; Warden's Wraps (14601, -77.00 DPS, sim-verified) [world_drop]; Insignia Chestguard (4057, -110.62 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.4 stamina points (294.46 DPS) | yes | Cultist's Armguards (270032, -34.49 DPS) [quest]; Barbaric Bracers (18948, -35.56 DPS) [crafted]; Branded Leather Bracers (19508, -46.44 DPS, sim-verified) [dungeon] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 20.0 stamina points (409.08 DPS) | yes | Bonefingers (10765, -19.45 DPS) [dungeon]; Razzeric's Racing Grips (6727, -21.61 DPS) [quest]; Warden's Leather Gloves (252527, -39.19 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 19.6 stamina points (402.27 DPS) | yes | Warden's Leather Belt (252460, -32.91 DPS) [crafted]; Scorpashi Sash (14652, -50.42 DPS) [world_drop]; Kolkar Hunter's Belt (6788, -61.29 DPS, sim-verified) [quest] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 25.2 stamina points (517.59 DPS) | yes | Warden's Woolies (14605, -10.74 DPS) [world_drop]; Panther Hunter Leggings (4108, -16.86 DPS) [quest]; Imperial Leather Pants (4062, -50.07 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 23.8 stamina points (488.01 DPS) | yes | Worn Running Boots (9398, -39.77 DPS, sim-verified) [dungeon]; Gnomebot Operating Boots (9450, -49.53 DPS) [dungeon]; Skulker's Leather Shoes (252531, -69.17 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (246.05 DPS) | yes | Field Researcher's Loop (281634, -80.04 DPS) [quest]; Plains Ring (2039, -82.02 DPS) [dungeon]; Trader's Ring (15689, -82.02 DPS) [quest] |
| finger2 | Talvash's Gold Ring (9538) | Gnome Improvement [quest] | 11.0 stamina points (224.67 DPS) | yes | Plains Ring (2039, -60.63 DPS) [dungeon]; Trader's Ring (15689, -60.63 DPS) [quest]; Field Researcher's Loop (281634, -90.53 DPS, sim-verified) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (4447.2 DPS) | yes | Mograine's Might (7723, -114.50 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -161.18 DPS) [dungeon]; Huge Stone Club (1523, -218.97 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Talvash's Gold Ring; trinket1: Arena Master; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 549, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 0000000000000000-55230332023102125110-0000000000000000)

Set DPS (verified): 115.8. Weights run: 3.4s. Verify run: 2.5s. 719 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.037, armor=0.169 ± 0.014, defense=not significant (0.564 ± 0.169) per rating point (1 rating = 1%, 0.564 per %), dodge=0.187 ± 0.034 per rating point (12 rating = 1%, 2.239 per %), strength=0.056 ± 0.000, agility=0.219 ± 0.021, attack_power=0.028 ± 0.000, hit=not significant (0.065 ± 0.030) per rating point (10 rating = 1%, 0.647 per %), crit=not significant (0.027 ± 0.013) per rating point (14 rating = 1%, 0.381 per %), expertise=3.109 ± 0.331

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 44.0 stamina points (1232.59 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -110.40 DPS) [vendor]; Knight-Lieutenant's Crackling Leather Helmet (220872, -263.80 DPS) [vendor]; Scorpashi Skullcap (14658, -393.14 DPS, sim-verified) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (391.97 DPS) | yes | Gazlowe's Charm (13088, -90.47 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -133.69 DPS) [world_drop]; Medal of Courage (6723, -139.99 DPS) [quest] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | 32.8 stamina points (918.68 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -18.04 DPS) [vendor]; Warden's Leather Shoulders (252536, -62.27 DPS) [crafted]; Fleshhide Shoulders (10774, -114.41 DPS, sim-verified) [dungeon] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 20.4 stamina points (571.52 DPS) | yes | Sergeant's Cape (18441, -9.45 DPS) [vendor]; Emberscale Cape (11311, -18.90 DPS) [dungeon]; Grovekeeper's Drape (17739, -60.72 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 58.2 stamina points (1629.99 DPS) | yes | Heraldic Breastplate (8119, -425.85 DPS) [world_drop]; Jinxed Hoodoo Skin (9473, -429.81 DPS, sim-verified) [dungeon]; Knight's Leather Armor (220854, -444.10 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 26.1 stamina points (731.09 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -36.72 DPS) [vendor]; Warden's Leather Bracers (252541, -128.96 DPS) [crafted]; Serpentskin Bracers (8257, -130.70 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | 33.2 stamina points (929.01 DPS) | yes | Prowler's Leather Gauntlets (252547, -132.11 DPS) [crafted]; Wicked Leather Gauntlets (15083, -141.24 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -159.93 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 29.7 stamina points (831.41 DPS) | yes | Skulker's Leather Waistguard (252474, -26.32 DPS) [crafted]; Prowler's Leather Waistguard (252473, -35.40 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -130.90 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 39.1 stamina points (1096.11 DPS) | yes | Knight's Restored Leather Leggings (220882, -172.30 DPS) [vendor]; Knight's Crackling Leather Leggings (220864, -189.48 DPS, sim-verified) [vendor]; Scorpashi Leggings (14659, -204.58 DPS) [world_drop] |
| feet | Slitherscale Boots (10801) | Sunken Temple: Spawn of Hakkar [dungeon] | 33.8 stamina points (945.87 DPS) | yes | Skulker's Leather Boots (252469, -105.59 DPS) [crafted]; Sergeant Major's Leather Boots (220860, -113.75 DPS) [vendor]; Warden's Leather Boots (252470, -194.70 DPS, sim-verified) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (419.97 DPS) | yes | Talvash's Gold Ring (9538, -92.59 DPS) [quest]; Seafire Band (4549, -124.60 DPS) [quest]; Ring of Fortitude (10739, -139.99 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (335.98 DPS) | yes | Talvash's Gold Ring (9538, -8.59 DPS) [quest]; Seafire Band (4549, -40.60 DPS) [quest]; Ring of Fortitude (10739, -56.00 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (8723.4 DPS) | yes | Arena Master (18706, -56.00 DPS) [world]; Smotts' Compass (4130, -189.30 DPS) [quest]; Frozen Heart of the Mountain (249469, -235.67 DPS) [crafted] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (8723.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (8723.4 DPS) | yes | Blanchard's Stout (13046, -52.83 DPS) [world_drop]; Radiant Staff (249453, -94.75 DPS) [crafted]; Dreamstaff (249454, -94.75 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Penance Spaulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Slitherscale Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Mark of the Chosen

No-known-source sample (15 of 719, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 152.7. Weights run: 3.8s. Verify run: 5.0s. 1597 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.014, armor=0.184 ± 0.012, defense=0.905 ± 0.217 per rating point (1 rating = 1%, 0.905 per %), dodge=0.596 ± 0.044 per rating point (12 rating = 1%, 7.156 per %), strength=0.088 ± 0.000, agility=0.511 ± 0.028, attack_power=0.044 ± 0.000, hit=not significant (0.113 ± 0.032) per rating point (10 rating = 1%, 1.128 per %), crit=0.097 ± 0.013 per rating point (14 rating = 1%, 1.361 per %), expertise=4.253 ± 0.284

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 73.2 stamina points (5579.79 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | sim-verified (14805.5 DPS) | yes | Master Sergeant's Insignia (18443, -153.50 DPS, sim-verified) [vendor]; Sunprism Pendant (20649, -222.83 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 50.5 stamina points (3848.70 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -372.30 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 26.6 stamina points (2024.40 DPS) | yes | Sergeant's Cape (16342, -97.63 DPS) [vendor]; Cloak of Earth and Sky (279270, -103.75 DPS) [crafted]; Nightfall Drape (12465, -424.45 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 66.2 stamina points (5048.10 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -326.82 DPS, sim-verified) [dungeon] |
| wrist | Sergeant Major's Dragonhide Armsplints (18454) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 31.2 stamina points (2375.44 DPS) | yes | Arena Bracers (18710, -220.64 DPS, sim-verified) [world]; Shadowcraft Bracers (16710, -262.22 DPS) [dungeon]; Knight-Captain's Dragonhide Armsplints (16395, -270.74 DPS) [pvp] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 42.5 stamina points (3241.27 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Bileblister Claws (275610, -234.36 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 41.7 stamina points (3177.29 DPS) | yes | Elderwild Waistcord (279252, -20.24 DPS) [crafted]; Serpentine Sash (13118, -214.68 DPS) [world_drop]; Marshal's Dragonhide Waistguard (16447, -248.91 DPS) [pvp] |
| legs | Warstrife Leggings (11821) | Blackrock Depths: General Angerforge [dungeon] | 66.3 stamina points (5050.78 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Elderwild Pants (279254, +0.00 DPS) [crafted] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 46.9 stamina points (3575.97 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -289.07 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (1143.39 DPS) | yes | Ring of the Heavens (12056, -152.45 DPS) [dungeon]; Dragonscale Band (12057, -152.45 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -152.45 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1067.16 DPS) | yes | Ring of the Heavens (12056, -76.23 DPS) [dungeon]; Dragonscale Band (12057, -76.23 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -76.23 DPS) [quest] |
| trinket1 | Arena Grand Master (19024) | Arena Grandmaster [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, -11.90 DPS) [world]; Demon's Blood (10779, -269.41 DPS) [quest] |
| trinket2 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, -140.56 DPS) [quest]; Arena Master (18706, -152.45 DPS) [world]; Demon's Blood (10779, -409.97 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Greenhammer (279261, -300.64 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Sergeant Major's Dragonhide Armsplints; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Warstrife Leggings; feet: Dawn Treaders; finger2: Ring of the Dawn; trinket1: Arena Grand Master; trinket2: Guardian Talisman; main_hand: Twig of the World Tree

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (night-elf, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 320.7. Weights run: 3.9s. Verify run: 4.7s. 1597 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.014, armor=0.228 ± 0.015, defense=1.566 ± 0.285 per rating point (1 rating = 1%, 1.566 per %), dodge=0.849 ± 0.057 per rating point (12 rating = 1%, 10.187 per %), strength=0.105 ± 0.000, agility=0.758 ± 0.037, attack_power=0.048 ± 0.000, hit=not significant (0.135 ± 0.041) per rating point (10 rating = 1%, 1.350 per %), crit=0.112 ± 0.017 per rating point (14 rating = 1%, 1.571 per %), expertise=5.345 ± 0.373

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 85.8 stamina points (8481.84 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, -547.13 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 22.7 stamina points (2241.47 DPS) | yes | Master Sergeant's Insignia (18443, -501.52 DPS, sim-verified) [vendor]; Sunprism Pendant (20649, -649.91 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 59.1 stamina points (5847.44 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Shadowcraft Spaulders (16708, -393.18 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 29.0 stamina points (2865.76 DPS) | yes | Cloak of Earth and Sky (279270, -155.37 DPS) [crafted]; Sergeant's Cape (16342, -169.24 DPS) [vendor]; Sergeant Major's Cape (16337, -421.88 DPS) [pvp] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 77.8 stamina points (7697.66 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -311.52 DPS, sim-verified) [dungeon] |
| wrist | Sergeant Major's Dragonhide Armsplints (18454) | PvP rank 9 · Sergeant Major · Alliance [vendor] | 34.6 stamina points (3418.31 DPS) | yes | Arena Bracers (18710, -140.37 DPS) [world]; Deepfury Bracers (13120, -300.71 DPS) [world_drop]; Shadowcraft Bracers (16710, -403.43 DPS, sim-verified) [dungeon] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 48.8 stamina points (4826.02 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 47.5 stamina points (4699.37 DPS) | yes | Serpentine Sash (13118, -96.79 DPS) [world_drop]; Elderwild Waistcord (279252, -118.51 DPS) [crafted]; Wicked Leather Belt (15088, -402.11 DPS) [crafted] |
| legs | Warstrife Leggings (11821) | Blackrock Depths: General Angerforge [dungeon] | 82.1 stamina points (8123.83 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -544.53 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 55.5 stamina points (5494.46 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Shadowcraft Boots (16711, -302.09 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.6 stamina points (1543.60 DPS) | yes | Talvash's Gold Ring (9538, -89.76 DPS) [quest]; Ring of the Dawn (13812, -158.84 DPS) [quest]; Hardened Stone Band (18674, -204.36 DPS) [world] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (1483.67 DPS) | yes | Talvash's Gold Ring (9538, -29.83 DPS) [quest]; Ring of the Dawn (13812, -98.91 DPS) [quest]; Hardened Stone Band (18674, -144.43 DPS) [world] |
| trinket1 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-verified (28189.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Arena Master (18706, -197.82 DPS) [world] |
| trinket2 | Molten Heart of the Mountain (249470) | Enchanting [crafted] | sim-verified (28189.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS, sim-verified) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (28189.4 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Greenhammer (279261, -460.14 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: Sergeant Major's Dragonhide Armsplints; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Warstrife Leggings; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of Saviors; trinket2: Molten Heart of the Mountain; main_hand: Twig of the World Tree

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 19.3. Weights run: 2.8s. Verify run: 1.9s. 226 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.085 ± 0.002, defense=0.133 ± 0.027 per rating point (1 rating = 1%, 0.133 per %), dodge=0.098 ± 0.008 per rating point (12 rating = 1%, 1.176 per %), strength=0.060 ± 0.000, agility=0.146 ± 0.005, attack_power=0.030 ± 0.000, hit=0.051 ± 0.006 per rating point (10 rating = 1%, 0.508 per %), crit=0.018 ± 0.001 per rating point (14 rating = 1%, 0.257 per %), expertise=1.273 ± 0.062

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.7 stamina points (286.91 DPS) | yes | Brawler's Leather Hood (252504, -15.80 DPS) [crafted]; Defender's Leather Hood (252447, -26.42 DPS) [crafted]; Totemic Leather Hood (252448, -33.78 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (50.64 DPS) | yes | Scholarly Pendant (277203, -4.50 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.4 stamina points (145.16 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -45.06 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -48.35 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (118.42 DPS) | yes | Sentry Cloak (2059, -5.08 DPS) [world_drop]; Firebane Cloak (12979, -16.69 DPS) [world_drop]; Cape of the Brotherhood (5193, -33.60 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (280.55 DPS) | yes | Gloomshroud Armor (1489, -17.82 DPS, sim-verified) [dungeon]; Bloody Apron (6226, -71.84 DPS) [dungeon]; Trapper's Leather Armor (252491, -76.82 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (125.26 DPS) | yes | Spare Part Bindings (279875, -18.90 DPS, sim-verified) [quest]; Drakewing Bands (12999, -25.54 DPS) [world_drop]; Savannah Bracers (15453, -35.99 DPS) [quest] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (139.78 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.31 DPS) [crafted]; Defender's Leather Gloves (252496, -5.31 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.2 stamina points (172.94 DPS) | yes | Belt of the Fang (10412, -30.42 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -39.69 DPS) [crafted]; Guardsman Belt (3429, -41.00 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.1 stamina points (216.72 DPS) | yes | Smelting Pants (5199, -19.60 DPS, sim-verified) [dungeon]; Brawler's Leather Pants (252500, -22.28 DPS) [crafted]; Leggings of the Fang (10410, -28.43 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.1 stamina points (231.62 DPS) | yes | Surfer Shoes (276274, -37.62 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -46.13 DPS) [dungeon]; Trapper's Leather Boots (252440, -63.76 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (76.91 DPS) | yes | Slain Baron's Signet (279867, -28.72 DPS) [quest]; Ring of the Moon (12052, -29.84 DPS) [world_drop]; Deep Fathom Ring (6463, -30.76 DPS) [dungeon] |
| finger2 | Ring of Scorn (3235) | A Husband's Revenge [quest] | 4.0 stamina points (61.53 DPS) | yes | Slain Baron's Signet (279867, -9.07 DPS, sim-verified) [quest]; Ring of the Moon (12052, -14.46 DPS) [world_drop]; Deep Fathom Ring (6463, -15.38 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 15.9 stamina points (244.39 DPS) | yes | Rhahk'Zor's Hammer (5187, -2.60 DPS) [dungeon]; Staff of the Blessed Seer (2271, -5.06 DPS) [dungeon]; Twisted Chanter's Staff (890, -20.46 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Ring of Scorn; main_hand: Rakzur Club

No-known-source sample (15 of 226, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20426 Advisor's Ring; 20427 Battle Healer's Cloak; 20429 Legionnaire's Band; 20441 Scout's Blade; 20442 Scout's Medallion

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 26.2. Weights run: 2.6s. Verify run: 1.8s. 414 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.090 ± 0.005, defense=0.462 ± 0.068 per rating point (1 rating = 1%, 0.462 per %), dodge=0.110 ± 0.014 per rating point (12 rating = 1%, 1.315 per %), strength=0.052 ± 0.000, agility=0.148 ± 0.009, attack_power=0.026 ± 0.000, hit=not significant (0.041 ± 0.011) per rating point (10 rating = 1%, 0.410 per %), crit=0.019 ± 0.002 per rating point (14 rating = 1%, 0.260 per %), expertise=1.604 ± 0.122

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.0 stamina points (343.88 DPS) | yes | Defender's Leather Helm (252455, -21.69 DPS, sim-verified) [crafted]; Totemic Leather Helm (252456, -23.06 DPS) [crafted]; Trapper's Leather Helm (252513, -23.06 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (143.67 DPS) | yes | Senior Sergeant's Insignia (15200, -3.27 DPS) [vendor]; Kaleidoscope Chain (13084, -68.78 DPS) [world_drop]; Pendant of Myzrael (4614, -81.27 DPS) [dungeon] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.4 stamina points (286.45 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -50.45 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -54.33 DPS) [crafted]; Barbaric Shoulders (5964, -70.46 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 7 · Sergeant · Horde [vendor] | 11.3 stamina points (176.93 DPS) | yes | Tigerstrike Mantle (13108, -14.15 DPS) [world_drop]; Enduring Cape (14763, -30.57 DPS) [world_drop]; Grimsteel Cape (4643, -30.80 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.8 stamina points (340.49 DPS) | yes | Spirewind Fetter (9406, -5.81 DPS) [dungeon]; Green Whelp Armor (7375, -25.58 DPS) [crafted]; Dokebi Chestguard (14581, -46.80 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.2 stamina points (175.10 DPS) | yes | Barbaric Bracers (18948, -2.98 DPS) [crafted]; Black Wolf Bracers (3230, -3.86 DPS) [dungeon]; Glowing Leather Bracers (2017, -36.29 DPS) [world] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 13.8 stamina points (215.26 DPS) | yes | Wolfclaw Gloves (1978, -18.04 DPS) [dungeon]; Brawler Gloves (720, -18.06 DPS, sim-verified) [world_drop]; Emblazoned Gloves (6397, -30.05 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.9 stamina points (247.65 DPS) | yes | Larcenist's Belt (279846, -27.53 DPS) [quest]; Skulker's Leather Belt (252520, -39.88 DPS) [crafted]; Prowler's Leather Belt (252459, -44.35 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.5 stamina points (288.71 DPS) | yes | Brawler's Leather Legguards (252516, -26.59 DPS) [crafted]; Trapper's Leather Legguards (252517, -32.31 DPS) [crafted]; Dark Ritual Leggings (270031, -38.70 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.8 stamina points (293.98 DPS) | yes | Harbinger Boots (7754, -20.99 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -56.20 DPS) [quest]; Gravewalker Boots (251963, -73.49 DPS) [quest] |
| finger1 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.2 stamina points (127.25 DPS) | yes | Defias Renegade Ring (1076, -24.43 DPS) [dungeon]; Blood Ring (4998, -49.25 DPS) [world_drop]; Heart Ring (5001, -49.25 DPS) [world_drop] |
| finger2 | Plains Ring (2039) | Razorfen Kraul [dungeon] | 8.0 stamina points (124.80 DPS) | yes | Defias Renegade Ring (1076, -18.87 DPS, sim-verified) [dungeon]; Blood Ring (4998, -46.80 DPS) [world_drop]; Heart Ring (5001, -46.80 DPS) [world_drop] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-verified (2398.1 DPS) | yes | Slaghammer (1976, -35.95 DPS) [dungeon]; Staff of the Shade (2549, -40.48 DPS, sim-verified) [dungeon]; Dense Triangle Mace (3203, -41.53 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Seal of Sylvanas; finger2: Plains Ring; main_hand: Spine of the Devourer

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 0000000000000000-55230332023102000000-0000000000000000)

Set DPS (verified): 47.8. Weights run: 3.1s. Verify run: 2.1s. 547 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.018, armor=0.123 ± 0.007, defense=not significant (0.319 ± 0.092) per rating point (1 rating = 1%, 0.319 per %), dodge=0.131 ± 0.021 per rating point (12 rating = 1%, 1.572 per %), strength=0.049 ± 0.000, agility=0.157 ± 0.013, attack_power=0.025 ± 0.000, hit=0.067 ± 0.016 per rating point (10 rating = 1%, 0.670 per %), crit=not significant (0.022 ± 0.006) per rating point (14 rating = 1%, 0.311 per %), expertise=2.068 ± 0.178

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 27.0 stamina points (553.97 DPS) | yes | Warden's Wizard Hat (14604, -36.01 DPS) [world_drop]; Brawler's Leather Helm (252512, -42.61 DPS) [crafted]; Nightscape Headband (8176, -59.48 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.2 stamina points (229.57 DPS) | yes | River Pride Choker (13087, -41.01 DPS) [world_drop]; Senior Sergeant's Insignia (15200, -45.03 DPS) [vendor]; Necklace of Calisea (1714, -65.54 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 27.9 stamina points (572.40 DPS) | yes | Flintrock Shoulders (7755, -64.36 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -141.28 DPS) [dungeon]; Cloudy Gustwoven Spaulders (277043, -213.96 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 14.1 stamina points (288.78 DPS) | yes | Silky Spider Cape (10776, -7.78 DPS) [dungeon]; Well Oiled Cloak (12254, -10.31 DPS) [vendor]; Heraldic Cloak (8120, -15.44 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 31.1 stamina points (637.97 DPS) | yes | Warden's Wraps (14601, -49.54 DPS) [world_drop]; Robes of the Lich (10762, -70.90 DPS) [dungeon]; Insignia Chestguard (4057, -110.62 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.4 stamina points (294.46 DPS) | yes | Branded Leather Bracers (19508, -22.11 DPS) [dungeon]; Cultist's Armguards (270032, -34.49 DPS) [quest]; Barbaric Bracers (18948, -35.56 DPS) [crafted] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 20.0 stamina points (409.08 DPS) | yes | Warden's Leather Gloves (252527, -2.55 DPS) [crafted]; Bonefingers (10765, -19.45 DPS) [dungeon]; Razzeric's Racing Grips (6727, -21.61 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 19.6 stamina points (402.27 DPS) | yes | Kolkar Hunter's Belt (6788, -27.46 DPS) [quest]; Warden's Leather Belt (252460, -32.91 DPS) [crafted]; Scorpashi Sash (14652, -50.42 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 25.2 stamina points (517.59 DPS) | yes | Imperial Leather Pants (4062, -3.14 DPS) [dungeon]; Warden's Woolies (14605, -10.74 DPS) [world_drop]; Panther Hunter Leggings (4108, -16.86 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 23.8 stamina points (488.01 DPS) | yes | Worn Running Boots (9398, -20.94 DPS) [dungeon]; Gnomebot Operating Boots (9450, -49.53 DPS) [dungeon]; Skulker's Leather Shoes (252531, -69.17 DPS) [crafted] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (246.05 DPS) | yes | Seal of Sylvanas (6414, -79.00 DPS) [quest]; Field Researcher's Loop (281634, -80.04 DPS) [quest]; Trader's Ring (15689, -82.02 DPS) [quest] |
| finger2 | Nogg's Gold Ring (9588) | Nogg's Ring Redo [quest] | 11.0 stamina points (224.67 DPS) | yes | Seal of Sylvanas (6414, -50.76 DPS, sim-verified) [quest]; Field Researcher's Loop (281634, -58.66 DPS) [quest]; Trader's Ring (15689, -60.63 DPS) [quest] |
| trinket1 | Arena Master (18706) | Arena Treasure Chest [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| trinket2 | - | - |  |  |  |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (4771.4 DPS) | yes | Ironshod Bludgeon (9408, -63.70 DPS, sim-verified) [dungeon]; Tok'kar's Murloc Basher (9678, -156.99 DPS) [quest]; Mograine's Might (7723, -185.15 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Nogg's Gold Ring; trinket1: Arena Master; main_hand: Cragwood Maul

No-known-source sample (15 of 547, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 50 (tauren, 0000000000000000-55230332023102125110-0000000000000000)

Set DPS (verified): 115.4. Weights run: 3.4s. Verify run: 2.7s. 717 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.037, armor=0.169 ± 0.014, defense=not significant (0.564 ± 0.169) per rating point (1 rating = 1%, 0.564 per %), dodge=0.187 ± 0.034 per rating point (12 rating = 1%, 2.239 per %), strength=0.056 ± 0.000, agility=0.219 ± 0.021, attack_power=0.028 ± 0.000, hit=not significant (0.065 ± 0.030) per rating point (10 rating = 1%, 0.647 per %), crit=not significant (0.027 ± 0.013) per rating point (14 rating = 1%, 0.381 per %), expertise=3.109 ± 0.331

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 44.0 stamina points (1232.59 DPS) | yes | Blood Guard's Leather Headband (220851, -110.40 DPS) [vendor]; Blood Guard's Crackling Leather Helmet (220873, -263.80 DPS) [vendor]; Scorpashi Skullcap (14658, -281.24 DPS, sim-verified) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (391.97 DPS) | yes | Woven Ivy Necklace (19159, -47.46 DPS) [quest]; Gazlowe's Charm (13088, -77.69 DPS) [dungeon]; River Pride Choker (13087, -133.69 DPS) [world_drop] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | 32.8 stamina points (918.68 DPS) | yes | Fleshhide Shoulders (10774, -5.23 DPS) [dungeon]; Blood Guard's Leather Shoulders (220853, -18.04 DPS) [vendor]; Warden's Leather Shoulders (252536, -62.27 DPS) [crafted] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 20.4 stamina points (571.52 DPS) | yes | Sergeant's Cloak (16341, -9.45 DPS) [vendor]; Emberscale Cape (11311, -18.90 DPS) [dungeon]; Grovekeeper's Drape (17739, -60.72 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 58.2 stamina points (1629.99 DPS) | yes | Jinxed Hoodoo Skin (9473, -383.28 DPS, sim-verified) [dungeon]; Heraldic Breastplate (8119, -425.85 DPS) [world_drop]; Stone Guard's Leather Armor (220855, -444.10 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 26.1 stamina points (731.09 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -36.72 DPS) [vendor]; Warden's Leather Bracers (252541, -128.96 DPS) [crafted]; Serpentskin Bracers (8257, -130.70 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | 33.2 stamina points (929.01 DPS) | yes | Prowler's Leather Gauntlets (252547, -132.11 DPS) [crafted]; Wicked Leather Gauntlets (15083, -141.24 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -146.49 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 29.7 stamina points (831.41 DPS) | yes | Skulker's Leather Waistguard (252474, -26.32 DPS) [crafted]; Prowler's Leather Waistguard (252473, -35.40 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -130.90 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 10 · Stone Guard · Horde [vendor] | 39.1 stamina points (1096.11 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -158.07 DPS, sim-verified) [vendor]; Stone Guard's Restored Leather Leggings (220883, -172.30 DPS) [vendor]; Scorpashi Leggings (14659, -204.58 DPS) [world_drop] |
| feet | Slitherscale Boots (10801) | Sunken Temple: Spawn of Hakkar [dungeon] | 33.8 stamina points (945.87 DPS) | yes | Warden's Leather Boots (252470, -84.18 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -105.59 DPS) [crafted]; First Sergeant's Leather Boots (220861, -113.75 DPS) [vendor] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (419.97 DPS) | yes | Nogg's Gold Ring (9588, -92.59 DPS) [quest]; Seafire Band (4549, -124.60 DPS) [quest]; Sha'ni's Ring (11869, -139.99 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (335.98 DPS) | yes | Nogg's Gold Ring (9588, -8.59 DPS) [quest]; Seafire Band (4549, -40.60 DPS) [quest]; Sha'ni's Ring (11869, -56.00 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (9404.4 DPS) | yes | Arena Master (18706, -56.00 DPS) [world]; Smotts' Compass (4130, -189.30 DPS) [quest]; Rune of the Guard Captain (19120, -206.20 DPS) [quest] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (9404.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Rune of the Guard Captain (19120, +0.00 DPS) [quest] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (9404.4 DPS) | yes | Ironshod Bludgeon (9408, -123.87 DPS, sim-verified) [dungeon]; Blanchard's Stout (13046, -151.59 DPS) [world_drop]; Radiant Staff (249453, -193.50 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; shoulder: Penance Spaulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Slitherscale Boots; finger1: Ring of Saviors; finger2: Underworld Band; trinket1: Guardian Talisman; trinket2: Mark of the Chosen

No-known-source sample (15 of 717, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60 (tauren, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 154.4. Weights run: 3.8s. Verify run: 4.5s. 1597 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.014, armor=0.184 ± 0.012, defense=0.905 ± 0.217 per rating point (1 rating = 1%, 0.905 per %), dodge=0.596 ± 0.044 per rating point (12 rating = 1%, 7.156 per %), strength=0.088 ± 0.000, agility=0.511 ± 0.028, attack_power=0.044 ± 0.000, hit=not significant (0.113 ± 0.032) per rating point (10 rating = 1%, 1.128 per %), crit=0.097 ± 0.013 per rating point (14 rating = 1%, 1.361 per %), expertise=4.253 ± 0.284

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 73.2 stamina points (5579.79 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Senior Sergeant's Insignia (16335) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 17.0 stamina points (1295.84 DPS) | yes | Medallion of Grand Marshal Morris (13091, -72.09 DPS) [world_drop]; Woven Ivy Necklace (19159, -142.87 DPS) [quest]; Hunter's Insignia Medal (15704, -152.45 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 50.5 stamina points (3848.70 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -315.45 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 26.6 stamina points (2024.40 DPS) | yes | Sergeant's Cloak (18461, -97.63 DPS) [vendor]; Cloak of Earth and Sky (279270, -103.75 DPS) [crafted]; Nightfall Drape (12465, -424.45 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 66.2 stamina points (5048.10 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -366.67 DPS, sim-verified) [dungeon] |
| wrist | First Sergeant's Dragonhide Armguards (18434) | PvP rank 9 · First Sergeant · Horde [vendor] | 31.2 stamina points (2375.44 DPS) | yes | Shadowcraft Bracers (16710, -262.22 DPS) [dungeon]; Legionnaire's Dragonhide Armguards (16493, -270.74 DPS) [pvp]; Arena Bracers (18710, -273.97 DPS, sim-verified) [world] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 42.5 stamina points (3241.27 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Bileblister Claws (275610, +0.00 DPS) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 41.7 stamina points (3177.29 DPS) | yes | Elderwild Waistcord (279252, -20.24 DPS) [crafted]; Serpentine Sash (13118, -214.68 DPS) [world_drop]; General's Dragonhide Belt (16556, -248.91 DPS) [pvp] |
| legs | Warstrife Leggings (11821) | Blackrock Depths: General Angerforge [dungeon] | 66.3 stamina points (5050.78 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Elderwild Pants (279254, +0.00 DPS) [crafted] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 46.9 stamina points (3575.97 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -209.06 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (1143.39 DPS) | yes | Ring of the Heavens (12056, -152.45 DPS) [dungeon]; Dragonscale Band (12057, -152.45 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -152.45 DPS) [quest] |
| finger2 | Ring of the Dawn (13812) | <NYI> <TXT> Ring of the Dawn [quest] | 14.0 stamina points (1067.16 DPS) | yes | Ring of the Heavens (12056, -76.23 DPS) [dungeon]; Dragonscale Band (12057, -76.23 DPS) [dungeon]; Dalson Family Wedding Ring (13475, -76.23 DPS) [quest] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (15693.4 DPS) | yes | Smotts' Compass (4130, -140.56 DPS) [quest]; Arena Grand Master (19024, -140.56 DPS) [quest]; Demon's Blood (10779, -409.97 DPS) [quest] |
| trinket2 | Arena Master (18706) | Arena Treasure Chest [world] | sim-verified (15693.4 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Demon's Blood (10779, -257.52 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (15693.4 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Cragwood Maul (11265, -344.51 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Senior Sergeant's Insignia; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: First Sergeant's Dragonhide Armguards; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Warstrife Leggings; feet: Dawn Treaders; finger2: Ring of the Dawn; trinket2: Arena Master; main_hand: Twig of the World Tree

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60, raid preset (tauren, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 323.9. Weights run: 3.9s. Verify run: 4.7s. 1597 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.014, armor=0.228 ± 0.015, defense=1.566 ± 0.285 per rating point (1 rating = 1%, 1.566 per %), dodge=0.849 ± 0.057 per rating point (12 rating = 1%, 10.187 per %), strength=0.105 ± 0.000, agility=0.758 ± 0.037, attack_power=0.048 ± 0.000, hit=not significant (0.135 ± 0.041) per rating point (10 rating = 1%, 1.350 per %), crit=0.112 ± 0.017 per rating point (14 rating = 1%, 1.571 per %), expertise=5.345 ± 0.373

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 85.8 stamina points (8481.84 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 22.7 stamina points (2241.47 DPS) | yes | Woven Ivy Necklace (19159, -375.73 DPS, sim-verified) [quest]; Senior Sergeant's Insignia (16335, -559.98 DPS) [vendor]; Sunprism Pendant (20649, -649.91 DPS) [quest] |
| shoulder | Golden Mantle of the Dawn (19058) | Leatherworking [crafted] | 59.1 stamina points (5847.44 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Shadowcraft Spaulders (16708, -499.83 DPS, sim-verified) [dungeon] |
| back | Fel Cape (279269) | Tailoring [crafted] | 29.0 stamina points (2865.76 DPS) | yes | Cloak of Earth and Sky (279270, -155.37 DPS) [crafted]; Sergeant's Cloak (18461, -169.24 DPS) [vendor]; Deep River Cloak (15789, -537.11 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 77.8 stamina points (7697.66 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -577.77 DPS, sim-verified) [dungeon] |
| wrist | First Sergeant's Dragonhide Armguards (18434) | PvP rank 9 · First Sergeant · Horde [vendor] | 34.6 stamina points (3418.31 DPS) | yes | Shadowcraft Bracers (16710, -0.48 DPS) [dungeon]; Arena Bracers (18710, -140.37 DPS) [world]; Deepfury Bracers (13120, -300.71 DPS) [world_drop] |
| hands | Hivethrasher's Claws (275613) | Leatherworking [crafted] | 48.8 stamina points (4826.02 DPS) | yes | Timbermaw Brawlers (19049, +0.00 DPS) [crafted]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 47.5 stamina points (4699.37 DPS) | yes | Serpentine Sash (13118, -96.79 DPS) [world_drop]; Elderwild Waistcord (279252, -118.51 DPS) [crafted]; Wicked Leather Belt (15088, -402.11 DPS) [crafted] |
| legs | Warstrife Leggings (11821) | Blackrock Depths: General Angerforge [dungeon] | 82.1 stamina points (8123.83 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Shadowcraft Pants (16709, -955.80 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 55.5 stamina points (5494.46 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Shadowcraft Boots (16711, -296.37 DPS, sim-verified) [dungeon] |
| finger1 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.6 stamina points (1543.60 DPS) | yes | Nogg's Gold Ring (9588, -89.76 DPS) [quest]; Ring of the Dawn (13812, -158.84 DPS) [quest]; Hardened Stone Band (18674, -204.36 DPS) [world] |
| finger2 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (1483.67 DPS) | yes | Nogg's Gold Ring (9588, -29.83 DPS) [quest]; Ring of the Dawn (13812, -98.91 DPS) [quest]; Hardened Stone Band (18674, -144.43 DPS) [world] |
| trinket1 | Guardian Talisman (1490) | The Temple of Atal'Hakkar [quest] | sim-verified (29935.9 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Grand Master (19024, +0.00 DPS) [quest]; Arena Master (18706, -197.82 DPS) [world] |
| trinket2 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (29935.9 DPS) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Arena Master (18706, +0.00 DPS) [world]; Arena Grand Master (19024, +0.00 DPS) [quest] |
| main_hand | Twig of the World Tree (13047) | World drop [world_drop] | sim-verified (29935.9 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Greenhammer (279261, -413.34 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Golden Mantle of the Dawn; back: Fel Cape; wrist: First Sergeant's Dragonhide Armguards; hands: Hivethrasher's Claws; waist: Hivethrasher's Girdle; legs: Warstrife Leggings; feet: Dawn Treaders; finger1: Masons Fraternity Ring; finger2: Ring of Saviors; main_hand: Twig of the World Tree

No-known-source sample (15 of 1597, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

