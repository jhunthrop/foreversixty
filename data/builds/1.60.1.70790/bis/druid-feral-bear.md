# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55000001000000000000-0000000000000000)

Set DPS (verified): 20.5. Weights run: 10.1s. Verify run: 7.0s. 222 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.104 ± 0.003, defense=0.190 ± 0.034 per rating point (1 rating = 1%, 0.190 per %), dodge=0.114 ± 0.009 per rating point (12 rating = 1%, 1.365 per %), strength=0.099 ± 0.000, agility=0.170 ± 0.006, attack_power=0.049 ± 0.000, feral_attack_power=0.049 ± 0.000, hit=0.047 ± 0.006 per rating point (10 rating = 1%, 0.466 per %), crit=0.019 ± 0.001 per rating point (14 rating = 1%, 0.261 per %), expertise=1.201 ± 0.064

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 19.5 stamina points (297.12 DPS) | yes | Brawler's Leather Hood (252504, -3.10 DPS) [crafted]; Defender's Leather Hood (252447, -11.84 DPS) [crafted]; Trapper's Leather Hood (252505, -23.89 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.7 stamina points (87.02 DPS) | yes | Erudite's Amulet (277204, -35.20 DPS, sim-verified) [quest]; Sentinel's Medallion (20444, -40.91 DPS) [rep]; Scholarly Pendant (277203, -41.25 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 10.7 stamina points (162.63 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -41.68 DPS) [dungeon]; Rugged Spaulders (5254, -56.26 DPS) [dungeon] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-verified (1965.8 DPS) | yes | Sentry Cloak (2059, -3.27 DPS) [world_drop]; Firebane Cloak (12979, -16.84 DPS) [world_drop]; Caretaker's Cape (20428, -30.51 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 20.2 stamina points (307.69 DPS) | yes | Gloomshroud Armor (1489, -19.14 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -56.75 DPS) [quest]; Defender's Leather Armor (252434, -78.11 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.9 stamina points (135.02 DPS) | yes | Drakewing Bands (12999, -19.51 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -35.37 DPS) [vendor]; Sanguine Cuffs (14375, -47.01 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 10.2 stamina points (155.56 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.59 DPS) [crafted]; Defender's Leather Gloves (252496, -4.37 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 12.5 stamina points (190.02 DPS) | yes | Belt of the Fang (10412, -17.97 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -42.40 DPS) [crafted]; Guardsman Belt (3429, -43.98 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 16.1 stamina points (245.68 DPS) | yes | Brawler's Leather Pants (252500, -22.04 DPS) [crafted]; Leggings of the Fang (10410, -29.89 DPS) [dungeon]; Smelting Pants (5199, -81.62 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.7 stamina points (240.23 DPS) | yes | Footpads of the Fang (10411, -33.81 DPS, sim-verified) [dungeon]; Surfer Shoes (276274, -41.63 DPS) [vendor]; Trapper's Leather Boots (252440, -54.11 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (76.28 DPS) | yes | Slain Baron's Signet (279867, -27.61 DPS) [quest]; Ring of the Moon (12052, -29.01 DPS) [world_drop]; Protector's Band (20439, -29.35 DPS) [rep] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 4.0 stamina points (61.03 DPS) | yes | Slain Baron's Signet (279867, -12.60 DPS, sim-verified) [quest]; Ring of the Moon (12052, -13.75 DPS) [world_drop]; Protector's Band (20439, -14.09 DPS) [rep] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (1965.8 DPS) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 21.0 stamina points (320.80 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.25 DPS) [dungeon]; Twisted Chanter's Staff (890, -13.63 DPS) [world_drop]; Staff of the Blessed Seer (2271, -18.22 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Sustaining Ring; trinket1: Rune of Perfection; main_hand: Rakzur Club

No-known-source sample (15 of 222, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 3738 Brewing Rod; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 18853 Insignia of the Horde

### Band 30 (night-elf, 0000000000000000-55230202000100100000-0000000000000000)

Set DPS (verified): 38.2. Weights run: 10.4s. Verify run: 17.6s. 390 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.082 ± 0.005, defense=not significant (0.318 ± 0.080) per rating point (1 rating = 1%, 0.318 per %), dodge=0.112 ± 0.015 per rating point (12 rating = 1%, 1.344 per %), strength=0.051 ± 0.000, agility=0.158 ± 0.010, attack_power=0.025 ± 0.000, feral_attack_power=0.025 ± 0.000, hit=0.060 ± 0.013 per rating point (10 rating = 1%, 0.605 per %), crit=0.019 ± 0.002 per rating point (14 rating = 1%, 0.262 per %), expertise=1.643 ± 0.143

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.5 stamina points (463.76 DPS) | yes | Defender's Leather Helm (252455, -27.15 DPS, sim-verified) [crafted]; Trapper's Leather Helm (252513, -34.08 DPS) [crafted]; Stormrider's Leather Helm (252514, -34.08 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (198.61 DPS) | yes | Master Sergeant's Insignia (18442, -4.38 DPS) [vendor]; Seal of the Expedition (271920, -25.96 DPS) [rep]; Sentinel's Medallion (19541, -63.45 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 17.7 stamina points (382.71 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -65.82 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -72.86 DPS) [crafted]; Barbaric Shoulders (5964, -95.78 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mourning Shawl (6751, -28.68 DPS) [quest]; Enduring Cape (14763, -41.26 DPS) [world_drop]; Tigerstrike Mantle (13108, -91.21 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | sim-verified (+13.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Raptor Hide Harness (4455, -13.30 DPS, sim-verified) [crafted]; Green Whelp Armor (7375, -25.43 DPS) [crafted]; Dokebi Chestguard (14581, -54.11 DPS) [world_drop] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | sim-verified (+14.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Black Wolf Bracers (3230, -1.34 DPS) [dungeon]; Cultist's Armguards (270032, -14.50 DPS, sim-verified) [quest]; Glowing Leather Bracers (2017, -45.21 DPS) [world] |
| hands | Brawler Gloves (720) | World drop [world_drop] | 12.5 stamina points (269.91 DPS) | yes | Wolfclaw Gloves (1978, -6.20 DPS) [dungeon]; Gloves of Insight (9698, -10.54 DPS) [quest]; Emblazoned Gloves (6397, -22.19 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.4 stamina points (333.10 DPS) | yes | Stalker's Leather Belt (252521, -61.10 DPS) [crafted]; Prowler's Leather Belt (252459, -61.45 DPS) [crafted]; Skulker's Leather Belt (252520, -138.91 DPS, sim-verified) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 17.9 stamina points (385.54 DPS) | yes | Brawler's Leather Legguards (252516, -36.35 DPS) [crafted]; Trapper's Leather Legguards (252517, -44.02 DPS) [crafted]; Dark Ritual Leggings (270031, -180.30 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.3 stamina points (393.90 DPS) | yes | Harbinger Boots (7754, -28.27 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -71.02 DPS) [quest]; Emblazoned Boots (4051, -118.99 DPS) [world_drop] |
| finger1 | Theramore Signet (271911) | Theramore Expeditionary Force [rep] | 8.2 stamina points (177.58 DPS) | yes | Plains Ring (2039, -4.93 DPS) [dungeon]; Defias Renegade Ring (1076, -34.46 DPS) [dungeon]; Protector's Band (19517, -64.23 DPS) [rep] |
| finger2 | Militant's Band (271916) (or Plains Ring (2039)) | Theramore Expeditionary Force [rep] | 8.0 stamina points (172.65 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Defias Renegade Ring (1076, -29.53 DPS) [dungeon]; Protector's Band (19517, -59.30 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (+118.2 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Shade (2549, -10.30 DPS) [dungeon]; Slaghammer (1976, -50.82 DPS) [dungeon]; Manual Crowd Pummeler (9449, -63.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Brawler Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Theramore Signet; finger2: Militant's Band; trinket1: Talisman of Arathor; trinket2: Rune of Perfection; main_hand: Spine of the Devourer

No-known-source sample (15 of 390, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14148 Crystalline Cuffs

### Band 40 (night-elf, 0000000000000000-55230332023100100010-0000000000000000)

Set DPS (verified): 74.1. Weights run: 10.6s. Verify run: 14.6s. 509 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.031, armor=0.161 ± 0.010, defense=0.577 ± 0.122 per rating point (1 rating = 1%, 0.577 per %), dodge=0.220 ± 0.025 per rating point (12 rating = 1%, 2.638 per %), strength=0.091 ± 0.000, agility=0.234 ± 0.016, attack_power=0.045 ± 0.000, feral_attack_power=0.045 ± 0.000, hit=not significant (0.084 ± 0.023) per rating point (10 rating = 1%, 0.843 per %), crit=0.039 ± 0.006 per rating point (14 rating = 1%, 0.551 per %), expertise=3.095 ± 0.245

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 31.2 stamina points (878.36 DPS) | yes | Warden's Wizard Hat (14604, -60.14 DPS, sim-verified) [world_drop]; Brawler's Leather Helm (252512, -66.79 DPS) [crafted]; Nightscape Headband (8176, -78.20 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.4 stamina points (319.61 DPS) | yes | Seal of the Expedition (271919, -10.21 DPS) [rep]; Sentinel's Medallion (19540, -50.28 DPS) [rep]; River Pride Choker (13087, -56.26 DPS) [world_drop] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 32.1 stamina points (904.01 DPS) | yes | Flintrock Shoulders (7755, -154.76 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -224.99 DPS) [dungeon]; Nightscape Shoulders (8192, -301.83 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-verified (7284.1 DPS) | yes | Silky Spider Cape (10776, -17.29 DPS) [dungeon]; Heraldic Cloak (8120, -19.08 DPS) [dungeon]; Well Oiled Cloak (12254, -21.82 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.9 stamina points (1009.37 DPS) | yes | Warden's Wraps (14601, -102.42 DPS, sim-verified) [world_drop]; Robes of the Lich (10762, -166.30 DPS) [dungeon]; Spirewind Fetter (9406, -175.36 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.8 stamina points (471.68 DPS) | yes | Branded Leather Bracers (19508, -28.55 DPS) [dungeon]; Barbaric Bracers (18948, -53.71 DPS) [crafted]; Cultist's Armguards (270032, -62.95 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-verified (7284.1 DPS) | yes | Skulker's Leather Gloves (252525, -18.53 DPS) [crafted]; Prowler's Leather Gloves (252524, -22.56 DPS) [crafted]; Stalker's Leather Gloves (252526, -332.31 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.3 stamina points (656.62 DPS) | yes | Warden's Leather Belt (252460, -63.61 DPS) [crafted]; Kolkar Hunter's Belt (6788, -73.68 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -76.00 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 31.1 stamina points (874.58 DPS) | yes | Panther Hunter Leggings (4108, -64.97 DPS) [quest]; Warden's Woolies (14605, -71.03 DPS) [world_drop]; Imperial Leather Pants (4062, -85.84 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 28.1 stamina points (789.61 DPS) | yes | Skulker's Leather Shoes (252531, -86.17 DPS) [crafted]; Prowler's Leather Shoes (252465, -102.30 DPS) [crafted]; Worn Running Boots (9398, -273.28 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (337.53 DPS) | yes | Theramore Signet (271910, -12.82 DPS) [rep]; Militant's Band (271915, -28.13 DPS) [rep]; Field Researcher's Loop (281634, -94.54 DPS) [quest] |
| finger2 | Talvash's Gold Ring (9538) | Gnome Improvement [quest] | 11.7 stamina points (329.97 DPS) | yes | Theramore Signet (271910, -5.25 DPS) [rep]; Militant's Band (271915, -20.56 DPS) [rep]; Field Researcher's Loop (281634, -86.98 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | sim-verified (7284.1 DPS) | yes | Rune of Duty (21567, +0.00 DPS) [rep]; Arena Master (18706, -28.13 DPS) [world]; Enriched Seal (271905, -28.13 DPS) [rep] |
| trinket2 | Talisman of Arathor (21118) | The League of Arathor [rep] | sim-verified (7284.1 DPS) | yes | Arena Master (18706, -28.13 DPS) [world]; Volunteer's Lucky Seal (271904, -28.13 DPS) [rep]; Enriched Seal (271905, -28.13 DPS) [rep] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (7284.1 DPS) | yes | Mograine's Might (7723, -130.47 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -337.10 DPS) [dungeon]; Huge Stone Club (1523, -412.04 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Talvash's Gold Ring; trinket1: Rune of Perfection; trinket2: Talisman of Arathor; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 509, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (night-elf, 0000000000000000-55230332023102125110-0000000000000000)

Set DPS (verified): 117.8. Weights run: 10.7s. Verify run: 20.8s. 665 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.053, armor=0.190 ± 0.018, defense=not significant (0.961 ± 0.252) per rating point (1 rating = 1%, 0.961 per %), dodge=0.261 ± 0.039 per rating point (12 rating = 1%, 3.136 per %), strength=0.095 ± 0.000, agility=0.294 ± 0.026, attack_power=0.047 ± 0.000, feral_attack_power=0.047 ± 0.000, hit=not significant (0.144 ± 0.039) per rating point (10 rating = 1%, 1.440 per %), crit=0.078 ± 0.017 per rating point (14 rating = 1%, 1.098 per %), expertise=4.751 ± 0.420

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 47.6 stamina points (1807.18 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -114.57 DPS) [vendor]; Scorpashi Skullcap (14658, -334.73 DPS, sim-verified) [world_drop]; Helm of Fire (8348, -387.77 DPS) [crafted] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 4 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (531.07 DPS) | yes | Seal of the Expedition (271918, -37.93 DPS) [rep]; Sentinel's Medallion (19539, -93.83 DPS) [rep]; Gazlowe's Charm (13088, -99.39 DPS) [dungeon] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | 36.1 stamina points (1368.84 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Fleshhide Shoulders (10774, -30.16 DPS) [dungeon]; Knight-Lieutenant's Crackling Leather Spaulders (220870, -84.31 DPS) [vendor] |
| back | Sergeant's Cape (18441) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emberscale Cape (11311, -14.42 DPS) [dungeon]; Grovekeeper's Drape (17739, -68.66 DPS) [dungeon]; Nightfall Drape (12465, -350.17 DPS, sim-verified) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 63.4 stamina points (2403.38 DPS) | yes | Jinxed Hoodoo Skin (9473, -558.92 DPS, sim-verified) [dungeon]; Knight's Leather Armor (220854, -602.83 DPS) [vendor]; Heraldic Breastplate (8119, -657.20 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 27.9 stamina points (1059.28 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -86.97 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -164.46 DPS) [crafted]; Serpentskin Bracers (8257, -190.49 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Gauntlets (252547, -171.66 DPS) [crafted]; Wicked Leather Gauntlets (15083, -206.48 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -493.97 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | sim-verified (+691.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Prowler's Leather Waistguard (252473, -10.76 DPS) [crafted]; Pridelord Girdle (14674, -180.31 DPS) [world_drop]; Skulker's Leather Waistguard (252474, -691.03 DPS, sim-verified) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 43.9 stamina points (1663.59 DPS) | yes | Knight's Restored Leather Leggings (220882, -276.71 DPS) [vendor]; Basilisk Hide Pants (1718, -311.49 DPS) [world_drop]; Knight's Crackling Leather Leggings (220864, -329.79 DPS, sim-verified) [vendor] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | sim-verified (+624.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Skulker's Leather Boots (252469, -6.81 DPS) [crafted]; Prowler's Leather Boots (252468, -21.91 DPS) [crafted]; Slitherscale Boots (10801, -624.54 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (569.00 DPS) | yes | Militant's Band (271914, -37.93 DPS) [rep]; Talvash's Gold Ring (9538, -80.34 DPS) [quest]; Underworld Band (1980, -113.80 DPS) [world_drop] |
| finger2 | Theramore Signet (271909) | Theramore Expeditionary Force [rep] | 14.7 stamina points (558.08 DPS) | yes | Militant's Band (271914, -27.01 DPS) [rep]; Talvash's Gold Ring (9538, -69.42 DPS) [quest]; Underworld Band (1980, -102.88 DPS) [world_drop] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Rune of Perfection (21565, -75.87 DPS) [rep] |
| trinket2 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Rune of Perfection (21565, -37.93 DPS) [rep] |
| main_hand | Blanchard's Stout (13046) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -9.54 DPS) [dungeon]; Mograine's Might (7723, -196.21 DPS) [dungeon]; Radiant Staff (249453, -200.03 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Penance Spaulders; back: Sergeant's Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Theramore Signet; trinket1: Enriched Seal; trinket2: Guardian Talisman; main_hand: Blanchard's Stout

No-known-source sample (15 of 665, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 165.5. Weights run: 11.2s. Verify run: 51.3s. 1443 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.136 ± 0.017, defense=1.037 ± 0.248 per rating point (1 rating = 1%, 1.037 per %), dodge=0.646 ± 0.048 per rating point (12 rating = 1%, 7.755 per %), strength=0.100 ± 0.000, agility=0.549 ± 0.031, attack_power=0.050 ± 0.000, feral_attack_power=0.050 ± 0.000, hit=0.209 ± 0.040 per rating point (10 rating = 1%, 2.090 per %), crit=0.121 ± 0.018 per rating point (14 rating = 1%, 1.688 per %), expertise=3.852 ± 0.364

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 66.9 stamina points (7795.73 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | sim-verified (+189.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Master Sergeant's Insignia (18443, -42.72 DPS) [vendor]; Stormpike Soldier's Pendant (19097, -171.30 DPS) [rep]; Sentinel's Medallion (19538, -189.58 DPS, sim-verified) [rep] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | sim-verified (+1356.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -1356.13 DPS, sim-verified) [rep] |
| back | Sergeant's Cape (16342) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Earth and Sky (279270, -37.16 DPS) [crafted]; Cloak of the Honor Guard (20073, -102.40 DPS) [rep]; Fel Cape (279269, -734.78 DPS, sim-verified) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.5 stamina points (6938.11 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -510.14 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 31.2 stamina points (3643.19 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18454, -438.00 DPS) [vendor]; Arena Bracers (18710, -610.12 DPS) [world]; Shadowcraft Bracers (16710, -738.82 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+432.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 37.5 stamina points (4373.89 DPS) | yes | Elderwild Waistcord (279252, -162.24 DPS) [crafted]; Serpentine Sash (13118, -279.23 DPS) [world_drop]; Stormpike Leather Girdle (19093, -347.84 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 84.6 stamina points (9862.32 DPS) | yes | Marshal's Dragonhide Leggings (231691, -27.75 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -674.00 DPS, sim-verified) [vendor]; Marshal's Dragonhide Legguards (231703, -845.01 DPS) [pvp] |
| feet | Highlander's Leather Boots (20052) | The League of Arathor [rep] | sim-verified (+617.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor]; Dawn Treaders (19052, -617.58 DPS, sim-verified) [crafted] |
| finger1 | Theramore Signet (271908) | Theramore Expeditionary Force [rep] | 16.8 stamina points (1964.21 DPS) | yes | Protector's Band (19514, -187.75 DPS) [rep]; Ring of Saviors (1447, -215.38 DPS) [world_drop]; Don Julio's Band (19325, -276.41 DPS) [rep] |
| finger2 | Militant's Band (271913) | Theramore Expeditionary Force [rep] | 16.0 stamina points (1865.42 DPS) | yes | Protector's Band (19514, -88.96 DPS) [rep]; Ring of Saviors (1447, -116.59 DPS) [world_drop]; Don Julio's Band (19325, -177.63 DPS) [rep] |
| trinket1 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Guardian Talisman (1490, -349.77 DPS) [quest]; Rune of Perfection (21565, -466.35 DPS) [rep] |
| trinket2 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Thunderbrew's Boot Flask (744, -431.52 DPS, sim-verified) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Electrified Dagger (19100, -549.63 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Highlander's Leather Shoulders; back: Sergeant's Cape; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Hivethrasher's Girdle; legs: Sentinel's Leather Pants; feet: Highlander's Leather Boots; finger1: Theramore Signet; finger2: Militant's Band; trinket1: Talisman of Arathor; trinket2: Enriched Seal; main_hand: The Unstoppable Force

No-known-source sample (15 of 1443, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60, raid preset (night-elf, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 396.4. Weights run: 12.3s. Verify run: 54.1s. 1443 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.042, armor=0.233 ± 0.027, defense=not significant (0.891 ± 0.298) per rating point (1 rating = 1%, 0.891 per %), dodge=0.594 ± 0.060 per rating point (12 rating = 1%, 7.123 per %), strength=0.098 ± 0.000, agility=0.602 ± 0.039, attack_power=0.045 ± 0.000, feral_attack_power=0.045 ± 0.000, hit=0.371 ± 0.059 per rating point (10 rating = 1%, 3.714 per %), crit=not significant (0.051 ± 0.026) per rating point (14 rating = 1%, 0.714 per %), expertise=9.095 ± 0.655

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 82.0 stamina points (7168.50 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Sentinel's Medallion (19538) | Silverwing Sentinels [rep] | 19.0 stamina points (1663.33 DPS) | yes | Medallion of Grand Marshal Morris (13091, -271.90 DPS) [world_drop]; Stormpike Soldier's Pendant (19097, -281.68 DPS) [rep]; Master Sergeant's Insignia (18443, -564.18 DPS, sim-verified) [vendor] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 63.7 stamina points (5565.44 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -2560.66 DPS, sim-verified) [rep] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Earth and Sky (279270, -62.39 DPS) [crafted]; Sergeant's Cape (16342, -76.83 DPS) [vendor]; Cloak of the Honor Guard (20073, -104.10 DPS) [rep] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 75.7 stamina points (6620.92 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -1058.35 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 40.5 stamina points (3545.33 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18454, -490.13 DPS) [vendor]; Shadowcraft Bracers (16710, -697.73 DPS) [dungeon]; Arena Bracers (18710, -701.70 DPS) [world] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+1107.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 46.9 stamina points (4099.12 DPS) | yes | Elderwild Waistcord (279252, -1.46 DPS) [crafted]; Serpentine Sash (13118, -187.08 DPS) [world_drop]; Stormpike Leather Girdle (19093, -231.77 DPS) [rep] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 103.5 stamina points (9050.11 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -713.92 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -1675.74 DPS, sim-verified) [vendor] |
| feet | Highlander's Leather Boots (20052) | The League of Arathor [rep] | sim-verified (+631.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor]; Dawn Treaders (19052, -631.69 DPS, sim-verified) [crafted] |
| finger1 | Theramore Signet (271908) | Theramore Expeditionary Force [rep] | 16.8 stamina points (1465.19 DPS) | yes | Protector's Band (19514, -84.26 DPS) [rep]; Don Julio's Band (19325, -149.15 DPS) [rep]; Ring of Saviors (1447, -153.67 DPS) [world_drop] |
| finger2 | Militant's Band (271913) | Theramore Expeditionary Force [rep] | 16.0 stamina points (1398.95 DPS) | yes | Protector's Band (19514, -18.03 DPS) [rep]; Don Julio's Band (19325, -82.91 DPS) [rep]; Ring of Saviors (1447, -87.43 DPS) [world_drop] |
| trinket1 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, -174.87 DPS) [rep]; Volunteer's Lucky Seal (271906, -174.87 DPS) [rep]; Guardian Talisman (1490, -262.30 DPS) [quest] |
| trinket2 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-verified (+518.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Volunteer's Lucky Seal (271906, -371.27 DPS) [rep]; Guardian Talisman (1490, -458.71 DPS) [quest]; Enriched Seal (271903, -518.04 DPS, sim-verified) [rep] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Lorekeeper's Staff (19570, -470.25 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Sentinel's Medallion; shoulder: Highlander's Leather Shoulders; back: Fel Cape; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Hivethrasher's Girdle; legs: Sentinel's Leather Pants; feet: Highlander's Leather Boots; finger1: Theramore Signet; finger2: Militant's Band; trinket1: Talisman of Arathor; trinket2: Stormpike Insignia Rank 6; main_hand: The Unstoppable Force

No-known-source sample (15 of 1443, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (tauren, 0000000000000000-55000001000000000000-0000000000000000)

Set DPS (verified): 21.0. Weights run: 10.1s. Verify run: 7.6s. 218 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.104 ± 0.003, defense=0.190 ± 0.034 per rating point (1 rating = 1%, 0.190 per %), dodge=0.114 ± 0.009 per rating point (12 rating = 1%, 1.365 per %), strength=0.099 ± 0.000, agility=0.170 ± 0.006, attack_power=0.049 ± 0.000, feral_attack_power=0.049 ± 0.000, hit=0.047 ± 0.006 per rating point (10 rating = 1%, 0.466 per %), crit=0.019 ± 0.001 per rating point (14 rating = 1%, 0.261 per %), expertise=1.201 ± 0.064

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 19.5 stamina points (297.12 DPS) | yes | Brawler's Leather Hood (252504, -3.10 DPS) [crafted]; Defender's Leather Hood (252447, -11.84 DPS) [crafted]; Trapper's Leather Hood (252505, -23.89 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (50.97 DPS) | yes | Scout's Medallion (20442, -4.86 DPS) [rep]; Scholarly Pendant (277203, -5.20 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 10.7 stamina points (162.63 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -41.68 DPS) [dungeon]; Rugged Spaulders (5254, -56.26 DPS) [dungeon] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-verified (2059.4 DPS) | yes | Sentry Cloak (2059, -3.27 DPS) [world_drop]; Firebane Cloak (12979, -16.84 DPS) [world_drop]; Battle Healer's Cloak (20427, -30.51 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 20.2 stamina points (307.69 DPS) | yes | Gloomshroud Armor (1489, -28.98 DPS, sim-verified) [dungeon]; Defender's Leather Armor (252434, -78.11 DPS) [crafted]; Trapper's Leather Armor (252491, -78.26 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.9 stamina points (135.02 DPS) | yes | Spare Part Bindings (279875, -18.56 DPS, sim-verified) [quest]; Drakewing Bands (12999, -24.16 DPS) [world_drop]; Wolf Bracers (4794, -35.37 DPS) [vendor] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 10.2 stamina points (155.56 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.59 DPS) [crafted]; Defender's Leather Gloves (252496, -4.37 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 12.5 stamina points (190.02 DPS) | yes | Belt of the Fang (10412, -25.87 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -42.40 DPS) [crafted]; Guardsman Belt (3429, -43.98 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 16.1 stamina points (245.68 DPS) | yes | Brawler's Leather Pants (252500, -22.04 DPS) [crafted]; Leggings of the Fang (10410, -29.89 DPS) [dungeon]; Smelting Pants (5199, -88.35 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.7 stamina points (240.23 DPS) | yes | Surfer Shoes (276274, -41.63 DPS) [vendor]; Footpads of the Fang (10411, -41.80 DPS, sim-verified) [dungeon]; Trapper's Leather Boots (252440, -54.11 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (76.28 DPS) | yes | Slain Baron's Signet (279867, -27.61 DPS) [quest]; Ring of the Moon (12052, -29.01 DPS) [world_drop]; Legionnaire's Band (20429, -29.35 DPS) [rep] |
| finger2 | Ring of Scorn (3235) | A Husband's Revenge [quest] | 4.0 stamina points (61.03 DPS) | yes | Slain Baron's Signet (279867, -12.18 DPS, sim-verified) [quest]; Ring of the Moon (12052, -13.75 DPS) [world_drop]; Legionnaire's Band (20429, -14.09 DPS) [rep] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (2059.4 DPS) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 21.0 stamina points (320.80 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.25 DPS) [dungeon]; Twisted Chanter's Staff (890, -13.63 DPS) [world_drop]; Staff of the Blessed Seer (2271, -18.22 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Ring of Scorn; trinket1: Rune of Perfection; main_hand: Rakzur Club

No-known-source sample (15 of 218, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 20425 Advisor's Gnarled Staff; 20441 Scout's Blade; 23173 Abomination Skin Leggings; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 248005 Northshire Hammer

### Band 30 (tauren, 0000000000000000-55230202000100100000-0000000000000000)

Set DPS (verified): 45.2. Weights run: 10.4s. Verify run: 17.5s. 391 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.104 ± 0.006, defense=not significant (0.194 ± 0.073) per rating point (1 rating = 1%, 0.194 per %), dodge=0.153 ± 0.018 per rating point (12 rating = 1%, 1.839 per %), strength=0.098 ± 0.000, agility=0.206 ± 0.011, attack_power=0.049 ± 0.000, feral_attack_power=0.049 ± 0.000, hit=0.072 ± 0.014 per rating point (10 rating = 1%, 0.723 per %), crit=0.026 ± 0.003 per rating point (14 rating = 1%, 0.370 per %), expertise=2.154 ± 0.152

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 23.8 stamina points (466.73 DPS) | yes | Defender's Leather Helm (252455, -17.46 DPS) [crafted]; Trapper's Leather Helm (252513, -40.42 DPS) [crafted]; Stormrider's Leather Helm (252514, -40.42 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.4 stamina points (184.17 DPS) | yes | Senior Sergeant's Insignia (15200, -7.65 DPS) [vendor]; Darkspear Warding Pendant (272075, -27.27 DPS) [vendor]; Scout's Medallion (19537, -53.77 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 19.6 stamina points (384.66 DPS) | yes | Azure Gustwoven Spaulders (277051, -64.06 DPS) [crafted]; Cloudy Gustwoven Spaulders (277043, -73.62 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -82.18 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enduring Cape (14763, -35.47 DPS) [world_drop]; Grimsteel Cape (4643, -37.26 DPS) [quest]; Tigerstrike Mantle (13108, -94.61 DPS, sim-verified) [world_drop] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | sim-verified (+11.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Spirewind Fetter (9406, -11.02 DPS, sim-verified) [dungeon]; Green Whelp Armor (7375, -31.07 DPS) [crafted]; Blackened Defias Armor (10399, -57.41 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.1 stamina points (237.37 DPS) | yes | Cultist's Armguards (270032, -0.76 DPS) [quest]; Black Wolf Bracers (3230, -8.54 DPS) [dungeon]; Glowing Leather Bracers (2017, -49.63 DPS) [world] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 15.1 stamina points (296.13 DPS) | yes | Brawler Gloves (720, -17.06 DPS) [world_drop]; Wolfclaw Gloves (1978, -20.21 DPS) [dungeon]; Emblazoned Gloves (6397, -41.98 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 17.4 stamina points (340.76 DPS) | yes | Skulker's Leather Belt (252520, -46.71 DPS) [crafted]; Prowler's Leather Belt (252459, -53.10 DPS) [crafted]; Larcenist's Belt (279846, -128.70 DPS, sim-verified) [quest] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 20.5 stamina points (401.79 DPS) | yes | Petrolspill Leggings (9509, -43.94 DPS) [dungeon]; Trapper's Leather Legguards (252517, -44.54 DPS) [crafted]; Brawler's Leather Legguards (252516, -171.43 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.9 stamina points (390.38 DPS) | yes | Harbinger Boots (7754, -25.98 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -81.59 DPS) [quest]; Gravewalker Boots (251963, -88.05 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.4 stamina points (165.51 DPS) | yes | Plains Ring (2039, -8.61 DPS) [dungeon]; Darkspear Signet (272071, -8.61 DPS) [vendor]; Defias Renegade Ring (1076, -31.67 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.3 stamina points (162.64 DPS) | yes | Plains Ring (2039, -5.74 DPS) [dungeon]; Darkspear Signet (272071, -5.74 DPS) [vendor]; Defias Renegade Ring (1076, -28.80 DPS) [dungeon] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | sim-verified (+129.4 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -8.34 DPS) [world]; Slaghammer (1976, -24.81 DPS) [dungeon]; Advisor's Gnarled Staff (212584, -33.74 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Raptor Hide Harness; wrist: Barbaric Bracers; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Defiler's Talisman; trinket2: Rune of Perfection; main_hand: Viscous Hammer

No-known-source sample (15 of 391, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape

### Band 40 (tauren, 0000000000000000-55230332023100100010-0000000000000000)

Set DPS (verified): 73.2. Weights run: 10.6s. Verify run: 13.3s. 508 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.031, armor=0.161 ± 0.010, defense=0.577 ± 0.122 per rating point (1 rating = 1%, 0.577 per %), dodge=0.220 ± 0.025 per rating point (12 rating = 1%, 2.638 per %), strength=0.091 ± 0.000, agility=0.234 ± 0.016, attack_power=0.045 ± 0.000, feral_attack_power=0.045 ± 0.000, hit=not significant (0.084 ± 0.023) per rating point (10 rating = 1%, 0.843 per %), crit=0.039 ± 0.006 per rating point (14 rating = 1%, 0.551 per %), expertise=3.095 ± 0.245

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 31.2 stamina points (878.36 DPS) | yes | Warden's Wizard Hat (14604, -35.42 DPS) [world_drop]; Brawler's Leather Helm (252512, -66.79 DPS) [crafted]; Nightscape Headband (8176, -78.20 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | 11.4 stamina points (319.61 DPS) | yes | Darkspear Warding Pendant (272074, -10.21 DPS) [vendor]; Scout's Medallion (19536, -50.28 DPS) [rep]; River Pride Choker (13087, -56.26 DPS) [world_drop] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 32.1 stamina points (904.01 DPS) | yes | Flintrock Shoulders (7755, -163.40 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -224.99 DPS) [dungeon]; Nightscape Shoulders (8192, -301.83 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-verified (7818.7 DPS) | yes | Silky Spider Cape (10776, -17.29 DPS) [dungeon]; Heraldic Cloak (8120, -19.08 DPS) [dungeon]; Well Oiled Cloak (12254, -21.82 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.9 stamina points (1009.37 DPS) | yes | Warden's Wraps (14601, -49.88 DPS) [world_drop]; Robes of the Lich (10762, -166.30 DPS) [dungeon]; Spirewind Fetter (9406, -175.36 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.8 stamina points (471.68 DPS) | yes | Barbaric Bracers (18948, -53.71 DPS) [crafted]; Branded Leather Bracers (19508, -59.13 DPS, sim-verified) [dungeon]; Cultist's Armguards (270032, -62.95 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-verified (7818.7 DPS) | yes | Skulker's Leather Gloves (252525, -18.53 DPS) [crafted]; Prowler's Leather Gloves (252524, -22.56 DPS) [crafted]; Stalker's Leather Gloves (252526, -318.64 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.3 stamina points (656.62 DPS) | yes | Warden's Leather Belt (252460, -63.61 DPS) [crafted]; Kolkar Hunter's Belt (6788, -69.00 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -76.00 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 31.1 stamina points (874.58 DPS) | yes | Panther Hunter Leggings (4108, -64.97 DPS) [quest]; Warden's Woolies (14605, -71.03 DPS) [world_drop]; Imperial Leather Pants (4062, -127.70 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 28.1 stamina points (789.61 DPS) | yes | Skulker's Leather Shoes (252531, -86.17 DPS) [crafted]; Prowler's Leather Shoes (252465, -102.30 DPS) [crafted]; Worn Running Boots (9398, -286.95 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (337.53 DPS) | yes | Insurgent's Band (272066, -12.82 DPS) [vendor]; Darkspear Signet (272070, -28.13 DPS) [vendor]; Field Researcher's Loop (281634, -94.54 DPS) [quest] |
| finger2 | Nogg's Gold Ring (9588) | Nogg's Ring Redo [quest] | 11.7 stamina points (329.97 DPS) | yes | Insurgent's Band (272066, -5.25 DPS) [vendor]; Darkspear Signet (272070, -20.56 DPS) [vendor]; Field Researcher's Loop (281634, -86.98 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | sim-verified (7818.7 DPS) | yes | Rune of Duty (21567, +0.00 DPS) [rep]; Arena Master (18706, -28.13 DPS) [world]; Darkspear Voodoo Seal (272059, -28.13 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21116) | The Defilers [rep] | sim-verified (7818.7 DPS) | yes | Arena Master (18706, -28.13 DPS) [world]; Darkspear Voodoo Seal (272059, -28.13 DPS) [vendor]; Relentless Raider's Seal (272062, -28.13 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (7818.7 DPS) | yes | Ironshod Bludgeon (9408, -89.16 DPS, sim-verified) [dungeon]; Mograine's Might (7723, -230.85 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -254.88 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Nogg's Gold Ring; trinket1: Rune of Perfection; trinket2: Defiler's Talisman; main_hand: Cragwood Maul

No-known-source sample (15 of 508, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads

### Band 50 (tauren, 0000000000000000-55230332023102125110-0000000000000000)

Set DPS (verified): 109.2. Weights run: 10.7s. Verify run: 12.0s. 664 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.045, armor=0.142 ± 0.016, defense=not significant (0.480 ± 0.195) per rating point (1 rating = 1%, 0.480 per %), dodge=0.206 ± 0.035 per rating point (12 rating = 1%, 2.468 per %), strength=0.068 ± 0.000, agility=0.236 ± 0.022, attack_power=0.034 ± 0.000, feral_attack_power=0.034 ± 0.000, hit=not significant (0.115 ± 0.035) per rating point (10 rating = 1%, 1.151 per %), crit=not significant (0.057 ± 0.015) per rating point (14 rating = 1%, 0.799 per %), expertise=3.446 ± 0.371

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 41.2 stamina points (1660.27 DPS) | yes | Blood Guard's Leather Headband (220851, -140.60 DPS) [vendor]; Scorpashi Skullcap (14658, -394.44 DPS, sim-verified) [world_drop]; Blood Guard's Restored Leather Helm (220875, -402.75 DPS) [vendor] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 4 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (564.84 DPS) | yes | Darkspear Warding Pendant (272073, -40.35 DPS) [vendor]; Woven Ivy Necklace (19159, -59.02 DPS) [quest]; Gazlowe's Charm (13088, -109.99 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 30.2 stamina points (1220.23 DPS) | yes | Blood Guard's Leather Shoulders (220853, -10.94 DPS) [vendor]; Penance Spaulders (11963, -12.99 DPS) [quest]; Warden's Leather Shoulders (252536, -89.26 DPS) [crafted] |
| back | Sergeant's Cloak (16341) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Emberscale Cape (11311, -11.45 DPS) [dungeon]; Grovekeeper's Drape (17739, -74.96 DPS) [dungeon]; Nightfall Drape (12465, -354.04 DPS, sim-verified) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 54.4 stamina points (2196.11 DPS) | yes | Heraldic Breastplate (8119, -605.04 DPS) [world_drop]; Stone Guard's Leather Armor (220855, -611.20 DPS) [vendor]; Jinxed Hoodoo Skin (9473, -650.38 DPS, sim-verified) [dungeon] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 24.5 stamina points (988.54 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -57.19 DPS) [vendor]; Warden's Leather Bracers (252541, -182.40 DPS) [crafted]; Serpentskin Bracers (8257, -183.99 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Prowler's Leather Gauntlets (252547, -187.92 DPS) [crafted]; Wicked Leather Gauntlets (15083, -191.74 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -642.08 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 27.5 stamina points (1107.66 DPS) | yes | Prowler's Leather Waistguard (252473, -41.48 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -194.52 DPS) [crafted]; Skulker's Leather Waistguard (252474, -655.43 DPS, sim-verified) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 36.6 stamina points (1476.10 DPS) | yes | Stone Guard's Restored Leather Leggings (220883, -273.01 DPS) [vendor]; Scorpashi Leggings (14659, -299.72 DPS) [world_drop]; Stone Guard's Crackling Leather Leggings (220865, -358.01 DPS, sim-verified) [vendor] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | sim-verified (+502.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Skulker's Leather Boots (252469, -37.47 DPS) [crafted]; Prowler's Leather Boots (252468, -51.01 DPS) [crafted]; Slitherscale Boots (10801, -502.88 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (605.18 DPS) | yes | Darkspear Signet (272069, -40.35 DPS) [vendor]; Underworld Band (1980, -121.04 DPS) [world_drop]; Nogg's Gold Ring (9588, -143.59 DPS) [quest] |
| finger2 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.5 stamina points (585.55 DPS) | yes | Darkspear Signet (272069, -20.71 DPS) [vendor]; Underworld Band (1980, -101.40 DPS) [world_drop]; Nogg's Gold Ring (9588, -123.95 DPS) [quest] |
| trinket1 | Defiler's Talisman (21115) | The Defilers [rep] | sim-verified (+256.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Guardian Talisman (1490, -40.35 DPS) [quest]; Rune of Perfection (21565, -80.69 DPS) [rep]; Mark of the Chosen (17774, -256.73 DPS, sim-verified) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -40.35 DPS) [quest]; Rune of Perfection (21565, -80.69 DPS) [rep] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -130.51 DPS, sim-verified) [dungeon]; Blanchard's Stout (13046, -187.16 DPS) [world_drop]; Radiant Staff (249453, -295.51 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Sergeant's Cloak; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal

No-known-source sample (15 of 664, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60 (tauren, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 163.9. Weights run: 11.2s. Verify run: 51.5s. 1448 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.136 ± 0.017, defense=1.037 ± 0.248 per rating point (1 rating = 1%, 1.037 per %), dodge=0.646 ± 0.048 per rating point (12 rating = 1%, 7.755 per %), strength=0.100 ± 0.000, agility=0.549 ± 0.031, attack_power=0.050 ± 0.000, feral_attack_power=0.050 ± 0.000, hit=0.209 ± 0.040 per rating point (10 rating = 1%, 2.090 per %), crit=0.121 ± 0.018 per rating point (14 rating = 1%, 1.688 per %), expertise=3.852 ± 0.364

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 66.9 stamina points (7795.73 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | sim-verified (+94.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Senior Sergeant's Insignia (16335, -42.72 DPS) [vendor]; Scout's Medallion (19534, -94.68 DPS, sim-verified) [rep]; Frostwolf Legionnaire's Pendant (19095, -171.30 DPS) [rep] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | sim-verified (+1302.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -1302.57 DPS, sim-verified) [vendor] |
| back | Sergeant's Cloak (18461) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cloak of Earth and Sky (279270, -37.16 DPS) [crafted]; Deathguard's Cloak (20068, -102.40 DPS) [rep]; Fel Cape (279269, -583.54 DPS, sim-verified) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.5 stamina points (6938.11 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Shadowcraft Tunic (16721, -403.02 DPS, sim-verified) [dungeon] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 31.2 stamina points (3643.19 DPS) | yes | First Sergeant's Dragonhide Armguards (18434, -438.00 DPS) [vendor]; Arena Bracers (18710, -610.12 DPS) [world]; Shadowcraft Bracers (16710, -738.82 DPS) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+525.2 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 37.5 stamina points (4373.89 DPS) | yes | Elderwild Waistcord (279252, -162.24 DPS) [crafted]; Serpentine Sash (13118, -279.23 DPS) [world_drop]; Frostwolf Leather Belt (19089, -347.84 DPS) [rep] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 67.6 stamina points (7885.64 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Warstrife Leggings (11821, -1165.46 DPS, sim-verified) [dungeon] |
| feet | Defiler's Leather Boots (20186) | The Defilers [rep] | sim-verified (+583.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Dawn Treaders (19052, -583.51 DPS, sim-verified) [crafted] |
| finger1 | Insurgent's Band (272064) | Creeg Bothunk [vendor] | 16.8 stamina points (1964.21 DPS) | yes | Legionnaire's Band (19510, -187.75 DPS) [rep]; Ring of Saviors (1447, -215.38 DPS) [world_drop]; Don Julio's Band (19325, -276.41 DPS) [rep] |
| finger2 | Darkspear Signet (272068) | Creeg Bothunk [vendor] | 16.0 stamina points (1865.42 DPS) | yes | Legionnaire's Band (19510, -88.96 DPS) [rep]; Ring of Saviors (1447, -116.59 DPS) [world_drop]; Don Julio's Band (19325, -177.63 DPS) [rep] |
| trinket1 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (+256.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Frostwolf Insignia Rank 6 (17909, +0.00 DPS) [quest]; Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor] |
| trinket2 | Defiler's Talisman (20072) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Insignia Rank 6 (17909, +0.00 DPS) [quest]; Relentless Raider's Seal (272060, -233.18 DPS) [vendor]; Darkspear Voodoo Seal (272061, -233.18 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Advisor's Gnarled Staff (19566, +0.00 DPS) [rep]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Defiler's Leather Shoulders; back: Sergeant's Cloak; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Hivethrasher's Girdle; legs: Outrider's Leather Pants; feet: Defiler's Leather Boots; finger1: Insurgent's Band; finger2: Darkspear Signet; trinket1: Mark of the Chosen; trinket2: Defiler's Talisman; main_hand: The Unstoppable Force

No-known-source sample (15 of 1448, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

### Band 60, raid preset (tauren, 0000000000000000-55230332023102125510-0510000000000000)

Set DPS (verified): 393.7. Weights run: 12.3s. Verify run: 54.0s. 1448 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.179 ± 0.020, defense=1.370 ± 0.313 per rating point (1 rating = 1%, 1.370 per %), dodge=0.911 ± 0.062 per rating point (12 rating = 1%, 10.928 per %), strength=0.114 ± 0.000, agility=0.825 ± 0.041, attack_power=0.052 ± 0.000, feral_attack_power=0.052 ± 0.000, hit=0.210 ± 0.048 per rating point (10 rating = 1%, 2.104 per %), crit=0.156 ± 0.024 per rating point (14 rating = 1%, 2.182 per %), expertise=5.826 ± 0.455

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 80.2 stamina points (13093.73 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, -722.14 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (19534) | Warsong Outriders [rep] | 22.4 stamina points (3653.42 DPS) | yes | Medallion of Grand Marshal Morris (13091, -274.35 DPS) [world_drop]; Woven Ivy Necklace (19159, -696.75 DPS) [quest]; Senior Sergeant's Insignia (16335, -877.74 DPS) [vendor] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | sim-verified (+2280.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -2280.25 DPS, sim-verified) [vendor] |
| back | Fel Cape (279269) | Tailoring [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Deathguard's Cloak (20068, -211.81 DPS) [rep]; Sergeant's Cloak (18461, -338.86 DPS) [vendor]; Cloak of Earth and Sky (279270, -356.23 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 71.3 stamina points (11648.23 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Dawn Armor (252483, -820.94 DPS, sim-verified) [crafted] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 40.3 stamina points (6579.71 DPS) | yes | Shadowcraft Bracers (16710, -1344.28 DPS) [dungeon]; First Sergeant's Dragonhide Armguards (18434, -1557.11 DPS) [vendor]; Arena Bracers (18710, -1618.02 DPS) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+1254.5 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; General's Dragonhide Gauntlets (231676, -1100.72 DPS) [vendor]; General's Dragonhide Gloves (231677, -1100.72 DPS) [pvp] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 43.4 stamina points (7085.56 DPS) | yes | Serpentine Sash (13118, -61.85 DPS) [world_drop]; Frostwolf Leather Belt (19089, -312.71 DPS) [rep]; Elderwild Waistcord (279252, -417.72 DPS) [crafted] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 83.2 stamina points (13584.84 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Warstrife Leggings (11821, -2061.55 DPS, sim-verified) [dungeon] |
| feet | Defiler's Leather Boots (20186) | The Defilers [rep] | sim-verified (+926.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Dawn Treaders (19052, -926.75 DPS, sim-verified) [crafted] |
| finger1 | Legionnaire's Band (19510) | Warsong Outriders [rep] | 18.4 stamina points (3011.05 DPS) | yes | Masons Fraternity Ring (9533, -308.72 DPS) [quest]; Darkspear Signet (272068, -398.66 DPS) [vendor]; Don Julio's Band (19325, -560.33 DPS) [rep] |
| finger2 | Insurgent's Band (272064) | Creeg Bothunk [vendor] | 16.9 stamina points (2756.02 DPS) | yes | Darkspear Signet (272068, -143.62 DPS) [vendor]; Don Julio's Band (19325, -305.29 DPS) [rep]; Masons Fraternity Ring (9533, -587.06 DPS, sim-verified) [quest] |
| trinket1 | Arena Grand Master (19024) | Arena Grandmaster [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, +0.00 DPS) [quest]; Frostwolf Insignia Rank 3 (17906, +0.00 DPS) [quest]; Defiler's Talisman (20072, +0.00 DPS) [rep] |
| trinket2 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Talisman (20072, -1609.12 DPS) [rep]; Smotts' Compass (4130, -1784.21 DPS) [quest]; Frostwolf Insignia Rank 3 (17906, -1784.21 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Advisor's Gnarled Staff (19566, -751.14 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Scout's Medallion; shoulder: Defiler's Leather Shoulders; back: Fel Cape; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Hivethrasher's Girdle; legs: Outrider's Leather Pants; feet: Defiler's Leather Boots; finger1: Legionnaire's Band; finger2: Insurgent's Band; trinket1: Arena Grand Master; trinket2: Frostwolf Insignia Rank 6; main_hand: The Unstoppable Force

No-known-source sample (15 of 1448, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise

