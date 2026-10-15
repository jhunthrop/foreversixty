# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 35.2. Weights run: 3.1s. Verify run: 3.0s. 188 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.081 ± 0.003, defense=not significant (0.118 ± 0.036) per rating point (1 rating = 1%, 0.118 per %), dodge=0.090 ± 0.012 per rating point (12 rating = 1%, 1.080 per %), strength=0.053 ± 0.000, agility=0.144 ± 0.007, attack_power=0.026 ± 0.000, hit=0.040 ± 0.007 per rating point (10 rating = 1%, 0.402 per %), crit=0.019 ± 0.002 per rating point (14 rating = 1%, 0.268 per %), expertise=1.218 ± 0.082

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.5 stamina points (354.30 DPS) | yes | Brawler's Leather Hood (252504, -22.36 DPS) [crafted]; Defender's Leather Hood (252447, -36.29 DPS) [crafted]; Totemic Leather Hood (252448, -44.35 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.0 stamina points (95.77 DPS) | yes | Erudite's Amulet (277204, -28.99 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -38.31 DPS) [quest]; Sentinel's Medallion (20444, -40.96 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.2 stamina points (176.33 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -56.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -57.65 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.6 stamina points (146.08 DPS) | yes | Sentry Cloak (2059, -6.60 DPS) [world_drop]; Grave Shroud (279865, -10.63 DPS) [quest]; Firebane Cloak (12979, -20.71 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 17.9 stamina points (342.49 DPS) | yes | Gloomshroud Armor (1489, -27.28 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -73.13 DPS) [quest]; Trapper's Leather Armor (252491, -95.03 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.0 stamina points (153.42 DPS) | yes | Drakewing Bands (12999, -18.65 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -46.46 DPS) [vendor]; Sanguine Cuffs (14375, -50.32 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 8.9 stamina points (170.19 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.56 DPS) [crafted]; Defender's Leather Gloves (252496, -6.97 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.0 stamina points (211.15 DPS) | yes | Belt of the Fang (10412, -28.35 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.75 DPS) [crafted]; Guardsman Belt (3429, -50.31 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.5 stamina points (315.08 DPS) | yes | Defender's Leather Pants (252445, -46.40 DPS, sim-verified) [crafted]; Slick Deviate Leggings (6480, -53.15 DPS) [quest]; Smelting Pants (5199, -62.59 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 14.9 stamina points (285.93 DPS) | yes | Surfer Shoes (276274, -37.67 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -59.47 DPS) [dungeon]; Trapper's Leather Boots (252440, -81.38 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (114.92 DPS) | yes | Blood Ring (4998, -19.15 DPS) [world_drop]; Ring of the Moon (12052, -56.45 DPS) [world_drop]; Deep Fathom Ring (6463, -57.46 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.2 stamina points (100.28 DPS) | yes | Blood Ring (4998, -4.51 DPS) [world_drop]; Ring of the Moon (12052, -41.81 DPS) [world_drop]; Deep Fathom Ring (6463, -42.82 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (2168.8 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.2 stamina points (311.17 DPS) | yes | Rakzur Club (12983, -25.00 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -27.83 DPS) [dungeon]; Twisted Chanter's Staff (890, -52.03 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 188, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 41.6. Weights run: 2.7s. Verify run: 3.1s. 340 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.006, defense=not significant (0.213 ± 0.068) per rating point (1 rating = 1%, 0.213 per %), dodge=0.100 ± 0.017 per rating point (12 rating = 1%, 1.205 per %), strength=0.044 ± 0.000, agility=0.143 ± 0.011, attack_power=0.022 ± 0.000, hit=not significant (0.036 ± 0.012) per rating point (10 rating = 1%, 0.363 per %), crit=0.018 ± 0.002 per rating point (14 rating = 1%, 0.247 per %), expertise=1.349 ± 0.141

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.5 stamina points (410.11 DPS) | yes | Defender's Leather Helm (252455, -17.22 DPS) [crafted]; Totemic Leather Helm (252456, -27.33 DPS) [crafted]; Trapper's Leather Helm (252513, -27.33 DPS) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (248.31 DPS) | yes | River Pride Choker (13087, -66.25 DPS, sim-verified) [world_drop]; Master Sergeant's Insignia (18442, -76.40 DPS) [vendor]; Fallen Guard's Pendant (279837, -90.45 DPS) [quest] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 17.8 stamina points (340.70 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -39.01 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -66.20 DPS) [crafted]; Darkshire Shoulderguards (270044, -85.33 DPS) [quest] |
| back | Mourning Shawl (6751) | Mortality Wanes [quest] | 14.0 stamina points (267.63 DPS) | yes | Sergeant's Cape (18440, -33.70 DPS, sim-verified) [vendor]; Tigerstrike Mantle (13108, -72.03 DPS) [world_drop]; Vine Pruner's Cloak (279835, -73.20 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.2 stamina points (405.19 DPS) | yes | Spirewind Fetter (9406, +0.00 DPS, sim-verified) [dungeon]; Green Whelp Armor (7375, -31.80 DPS) [crafted]; Dokebi Chestguard (14581, -57.30 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.9 stamina points (208.35 DPS) | yes | Black Wolf Bracers (3230, -4.13 DPS) [dungeon]; Barbaric Bracers (18948, -4.21 DPS) [crafted]; Staghide Armguards (270021, -31.95 DPS) [quest] |
| hands | Operator's Gloves (270045) | Gyrodrillmatic Excavationators [quest] | 12.9 stamina points (246.86 DPS) | yes | Brawler Gloves (720, -7.16 DPS) [world_drop]; Wolfclaw Gloves (1978, -13.99 DPS) [dungeon]; Gloves of Insight (9698, -15.50 DPS) [quest] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.4 stamina points (294.21 DPS) | yes | Skulker's Leather Belt (252520, -29.18 DPS, sim-verified) [crafted]; Stalker's Leather Belt (252521, -54.16 DPS) [crafted]; Prowler's Leather Belt (252459, -54.77 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 17.9 stamina points (341.01 DPS) | yes | Duty Bound Leggings (279868, -23.04 DPS) [quest]; Dark Ritual Leggings (270031, -28.23 DPS) [quest]; Brawler's Leather Legguards (252516, -32.73 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.4 stamina points (350.87 DPS) | yes | Harbinger Boots (7754, -25.42 DPS) [dungeon]; Highlander's Mail Greaves (20123, -57.74 DPS) [vendor]; Lancer Boots (6752, -62.07 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.2 stamina points (156.59 DPS) | yes | Plains Ring (2039, -3.79 DPS) [dungeon]; Ring of Power Regulation (273030, -3.79 DPS) [dungeon]; Boar Signet (274078, -22.89 DPS) [quest] |
| finger2 | Darkspear Signet (272071) (or Ring of Power Regulation (273030), Plains Ring (2039)) | Creeg Bothunk [vendor] | 8.0 stamina points (152.80 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Boar Signet (274078, -19.10 DPS) [quest] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (3315.4 DPS) | yes | Rune of Duty (21568, -38.20 DPS) [rep] |
| trinket2 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (3315.4 DPS) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (3315.4 DPS) | yes | Glimmering Staff (249392, -63.19 DPS) [crafted]; Soulstaff (249393, -63.84 DPS) [crafted]; Manual Crowd Pummeler (9449, -199.68 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenir Sea Shell; shoulder: Watchman Pauldrons; back: Mourning Shawl; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Operator's Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Darkspear Signet; trinket1: Talisman of Arathor; trinket2: Rune of Perfection; main_hand: Staff of the Shade

No-known-source sample (15 of 340, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 73.4. Weights run: 3.1s. Verify run: 3.2s. 445 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.028, armor=0.103 ± 0.009, defense=0.495 ± 0.121 per rating point (1 rating = 1%, 0.495 per %), dodge=0.119 ± 0.025 per rating point (12 rating = 1%, 1.424 per %), strength=0.046 ± 0.000, agility=0.136 ± 0.016, attack_power=0.023 ± 0.000, hit=not significant (0.049 ± 0.021) per rating point (10 rating = 1%, 0.491 per %), crit=not significant (0.027 ± 0.008) per rating point (14 rating = 1%, 0.375 per %), expertise=1.681 ± 0.227

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 24.9 stamina points (702.29 DPS) | yes | Warden's Wizard Hat (14604, -54.66 DPS) [world_drop]; Defender's Leather Helm (252455, -76.08 DPS) [crafted]; Brawler's Leather Helm (252512, -116.72 DPS, sim-verified) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (366.12 DPS) | yes | Gazlowe's Charm (13088, -51.14 DPS) [dungeon]; Darkspear Warding Pendant (272074, -56.33 DPS) [vendor]; River Pride Choker (13087, -107.46 DPS) [world_drop] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 25.9 stamina points (728.40 DPS) | yes | Flintrock Shoulders (7755, -165.71 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -181.83 DPS) [dungeon]; Pads of the Venom Spider (13103, -271.09 DPS) [world_drop] |
| back | Mourning Shawl (6751) | Mortality Wanes [quest] | 14.5 stamina points (407.83 DPS) | yes | Wing of the Whelpling (13121, -29.41 DPS) [world_drop]; Silky Spider Cape (10776, -38.86 DPS) [dungeon]; Well Oiled Cloak (12254, -41.77 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 28.7 stamina points (807.85 DPS) | yes | Warden's Wraps (14601, -74.74 DPS) [world_drop]; Insignia Chestguard (4057, -141.89 DPS) [world_drop]; Robes of the Lich (10762, -199.37 DPS, sim-verified) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 13.2 stamina points (371.43 DPS) | yes | Cultist's Armguards (270032, -39.72 DPS) [quest]; Barbaric Bracers (18948, -45.17 DPS) [crafted]; Branded Leather Bracers (19508, -56.90 DPS, sim-verified) [dungeon] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | 18.3 stamina points (514.55 DPS) | yes | Stalker's Leather Gloves (252526, -1.65 DPS) [crafted]; Bonefingers (10765, -25.29 DPS) [dungeon]; Razzeric's Racing Grips (6727, -26.75 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 18.0 stamina points (506.20 DPS) | yes | Warden's Leather Belt (252460, -38.64 DPS) [crafted]; Kolkar Hunter's Belt (6788, -60.86 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -67.73 DPS) [world_drop] |
| legs | Imperial Leather Pants (4062) | Uldaman: Ancient Treasure [dungeon] | 23.0 stamina points (646.53 DPS) | yes | Warden's Woolies (14605, -8.73 DPS) [world_drop]; Basilisk Hide Pants (1718, -12.07 DPS) [world_drop]; Panther Hunter Leggings (4108, -19.08 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 21.8 stamina points (613.43 DPS) | yes | Harbinger Boots (7754, -94.35 DPS) [dungeon]; Skulker's Leather Shoes (252531, -97.38 DPS) [crafted]; Gnomebot Operating Boots (9450, -164.17 DPS, sim-verified) [dungeon] |
| finger1 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.3 stamina points (318.88 DPS) | yes | Darkspear Signet (272070, -9.08 DPS) [vendor]; Underworld Band (1980, -37.25 DPS) [world_drop]; Talvash's Gold Ring (9538, -37.25 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.3 stamina points (317.58 DPS) | yes | Darkspear Signet (272070, -7.79 DPS) [vendor]; Underworld Band (1980, -35.95 DPS) [world_drop]; Talvash's Gold Ring (9538, -35.95 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | sim-verified (6144.5 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -28.16 DPS) [world]; Darkspear Voodoo Seal (272059, -28.16 DPS) [vendor] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (6144.5 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -28.16 DPS) [world]; Darkspear Voodoo Seal (272059, -28.16 DPS) [vendor] |
| main_hand | Mograine's Might (7723) | Scarlet Monastery: Scarlet Commander Mograine [dungeon] | sim-verified (6144.5 DPS) | yes | Staff of the Shade (2549, -133.18 DPS) [dungeon]; Huge Stone Club (1523, -139.45 DPS) [world]; Manual Crowd Pummeler (9449, -344.08 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; shoulder: Fleshhide Shoulders; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Imperial Leather Pants; feet: Warden's Leather Shoes; finger1: Suspicious Spare Part; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Mograine's Might

No-known-source sample (15 of 445, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 117.0. Weights run: 3.2s. Verify run: 6.6s. 580 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.043, armor=0.153 ± 0.017, defense=not significant (0.404 ± 0.219) per rating point (1 rating = 1%, 0.404 per %), dodge=0.260 ± 0.045 per rating point (12 rating = 1%, 3.115 per %), strength=0.056 ± 0.000, agility=0.242 ± 0.028, attack_power=0.028 ± 0.000, hit=not significant (0.093 ± 0.037) per rating point (10 rating = 1%, 0.928 per %), crit=not significant (0.031 ± 0.015) per rating point (14 rating = 1%, 0.438 per %), expertise=3.214 ± 0.412

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.6 stamina points (1558.02 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -265.17 DPS, sim-verified) [vendor]; Impractical Headwarmer (274756, -310.43 DPS) [vendor]; Scorpashi Skullcap (14658, -357.18 DPS) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (512.21 DPS) | yes | Darkspear Warding Pendant (272073, -36.59 DPS) [vendor]; Souvenir Sea Shell (274749, -36.59 DPS) [vendor]; Gazlowe's Charm (13088, -101.60 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | sim-verified (11735.2 DPS) | yes | Fleshhide Shoulders (10774, +0.00 DPS) [dungeon]; Warden's Leather Shoulder (252536, -56.27 DPS) [crafted]; Penance Spaulders (11963, -455.86 DPS, sim-verified) [quest] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 19.8 stamina points (725.48 DPS) | yes | Sergeant's Cape (18441, -11.22 DPS) [vendor]; Grovekeeper's Drape (17739, -78.79 DPS) [dungeon]; Mourning Shawl (6751, -151.75 DPS) [quest] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.2 stamina points (2056.68 DPS) | yes | Jinxed Hoodoo Skin (9473, -525.32 DPS, sim-verified) [dungeon]; Heraldic Breastplate (8119, -559.12 DPS) [world_drop]; Knight's Leather Armor (220854, -580.86 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.3 stamina points (924.63 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -130.10 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -168.66 DPS) [crafted]; Serpentskin Bracers (8257, -170.04 DPS) [world_drop] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 39.8 stamina points (1456.99 DPS) | yes | Warden's Leather Gauntlets (252549, -290.97 DPS) [crafted]; Raider Gloves (272100, -353.75 DPS, sim-verified) [vendor]; Feralheart Fists (226793, -354.08 DPS) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.3 stamina points (1036.19 DPS) | yes | Skulker's Leather Waistguard (252474, -22.17 DPS) [crafted]; Prowler's Leather Waistguard (252473, -35.83 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -170.81 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 37.4 stamina points (1369.19 DPS) | yes | Knight's Restored Leather Leggings (220882, -235.24 DPS) [vendor]; Scorpashi Leggings (14659, -252.92 DPS) [world_drop]; Knight's Crackling Leather Leggings (220864, -754.31 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (11735.2 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -472.41 DPS, sim-verified) [crafted] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.4 stamina points (527.51 DPS) | yes | Ring of Saviors (1447, -15.29 DPS) [world_drop]; Darkspear Signet (272069, -15.29 DPS) [vendor]; Suspicious Spare Part (274754, -110.78 DPS) [vendor] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (512.21 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Suspicious Spare Part (274754, -95.49 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (11735.2 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Guardian Talisman (1490, -36.59 DPS) [quest]; Rune of Perfection (21565, -73.17 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272060) | Creeg Bothunk [vendor] | sim-verified (11735.2 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Guardian Talisman (1490, -36.59 DPS) [quest]; Rune of Perfection (21565, -73.17 DPS) [rep] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (11735.2 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -32.73 DPS) [world_drop]; The Hand of Antu'sul (9639, -524.20 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Knight-Lieutenant's Leather Shoulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Radiant Staff

No-known-source sample (15 of 580, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 148.7. Weights run: 3.0s. Verify run: 10.8s. 1222 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.029, armor=0.177 ± 0.022, defense=not significant (1.151 ± 0.330) per rating point (1 rating = 1%, 1.151 per %), dodge=0.592 ± 0.069 per rating point (12 rating = 1%, 7.106 per %), strength=0.111 ± 0.000, agility=0.494 ± 0.043, attack_power=0.055 ± 0.000, hit=not significant (0.166 ± 0.053) per rating point (10 rating = 1%, 1.665 per %), crit=not significant (0.074 ± 0.023) per rating point (14 rating = 1%, 1.040 per %), expertise=5.890 ± 0.575

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 71.2 stamina points (4607.14 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 20.5 stamina points (1326.62 DPS) | yes | Medallion of Grand Marshal Morris (13091, -128.59 DPS) [world_drop]; Sentinel's Medallion (19538, -199.66 DPS) [rep]; Master Sergeant's Insignia (18443, -226.40 DPS) [vendor] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 67.1 stamina points (4340.94 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Highlander's Leather Shoulders (20059, -446.20 DPS, sim-verified) [rep]; Field Marshal's Dragonhide Spaulders (231699, -446.72 DPS) [pvp] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 32.0 stamina points (2069.82 DPS) | yes | Fel Cape (279269, -395.69 DPS) [crafted]; Sergeant's Cape (16342, -455.42 DPS) [vendor]; Cloak of Earth and Sky (279270, -463.00 DPS) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | 87.5 stamina points (5663.14 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Feralheart Armor (226804, -389.91 DPS, sim-verified) [vendor] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | 36.8 stamina points (2378.98 DPS) | yes | Bracers of Subterfuge (22668, -125.33 DPS) [quest]; Forest Stalker's Bracers (19587, -191.96 DPS) [rep]; Sergeant Major's Dragonhide Armsplints (18454, -398.93 DPS) [vendor] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 53.9 stamina points (3488.98 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -27.72 DPS) [vendor]; Raider Gloves (272099, -240.62 DPS, sim-verified) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 49.7 stamina points (3218.67 DPS) | yes | Shifter's Belt (272396, -481.04 DPS, sim-verified) [vendor]; Hivethrasher's Girdle (275614, -573.04 DPS) [crafted]; Belt of Preserved Heads (20216, -584.79 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 89.9 stamina points (5818.93 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -328.27 DPS) [pvp]; Dire Warbear Woolies (227804, -384.87 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 65.9 stamina points (4266.65 DPS) | yes | Marshal's Dragonhide Treads (231692, -169.27 DPS) [vendor]; Drudge Boots (21532, -552.79 DPS) [quest]; Feralheart Treads (226803, -599.36 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 33.5 stamina points (2167.99 DPS) | yes | Ring of Awareness (272409, -792.92 DPS) [vendor]; Insurgent's Band (272064, -1071.52 DPS) [vendor]; Cutthroat's Signet (272408, -1094.70 DPS) [vendor] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | 23.8 stamina points (1537.43 DPS) | yes | Ring of Awareness (272409, -162.37 DPS) [vendor]; Insurgent's Band (272064, -440.97 DPS) [vendor]; Cutthroat's Signet (272408, -464.15 DPS) [vendor] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (20857.2 DPS) | yes | Stormpike Insignia Rank 6 (17904, -48.76 DPS) [quest]; Relentless Raider's Seal (272060, -321.31 DPS) [vendor]; Darkspear Voodoo Seal (272061, -321.31 DPS) [vendor] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-verified (20857.2 DPS) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Relentless Raider's Seal (272060, -129.44 DPS) [vendor]; Darkspear Voodoo Seal (272061, -129.44 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (20857.2 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; The Lobotomizer (19324, -679.41 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Defender's Grip Stabilizer; trinket2: Talisman of Arathor; main_hand: The Unstoppable Force

No-known-source sample (15 of 1222, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 373.9. Weights run: 2.8s. Verify run: 10.9s. 1222 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.008, armor=0.200 ± 0.026, defense=not significant (1.699 ± 0.437) per rating point (1 rating = 1%, 1.699 per %), dodge=0.853 ± 0.084 per rating point (12 rating = 1%, 10.234 per %), strength=0.121 ± 0.000, agility=0.762 ± 0.055, attack_power=0.055 ± 0.000, hit=0.303 ± 0.062 per rating point (10 rating = 1%, 3.026 per %), crit=0.160 ± 0.028 per rating point (14 rating = 1%, 2.246 per %), expertise=8.111 ± 0.675

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 82.3 stamina points (7103.24 DPS) | yes | Feralheart Faceguard (226801, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 25.7 stamina points (2215.52 DPS) | yes | Medallion of Grand Marshal Morris (13091, -145.51 DPS) [world_drop]; Sentinel's Medallion (19538, -366.97 DPS) [rep] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 80.5 stamina points (6949.68 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -389.04 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -1323.35 DPS) [pvp]; Darkspear Pauldrons (272105, -1336.32 DPS, sim-verified) [vendor] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 40.8 stamina points (3517.79 DPS) | yes | Fel Cape (279269, -570.82 DPS, sim-verified) [crafted]; Cloak of the Honor Guard (20073, -1217.34 DPS) [rep]; Stalwart Cloak (272415, -1255.82 DPS) [vendor] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | 103.8 stamina points (8954.69 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Feralheart Armor (226804, -596.23 DPS, sim-verified) [vendor]; Field Marshal's Dragonhide Tunic (231702, -675.82 DPS) [vendor] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | 41.6 stamina points (3593.90 DPS) | yes | Forest Stalker's Bracers (19587, -58.26 DPS) [rep]; Bracers of Subterfuge (22668, -72.54 DPS) [quest] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 65.4 stamina points (5639.58 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Raider Gloves (272099, -336.51 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -758.57 DPS) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 58.4 stamina points (5043.38 DPS) | yes | Shifter's Belt (272396, -578.84 DPS) [vendor]; Belt of Preserved Heads (20216, -840.21 DPS) [quest]; Ferocity of the Timbermaw (227805, -965.79 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 104.2 stamina points (8991.36 DPS) | yes | Marshal's Dragonhide Leggings (231691, -253.76 DPS) [vendor]; Dire Warbear Woolies (227804, -484.67 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -1298.51 DPS) [pvp] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 80.0 stamina points (6899.22 DPS) | yes | Drudge Boots (21532, -631.55 DPS, sim-verified) [quest]; Marshal's Dragonhide Treads (231692, -797.49 DPS) [vendor]; Feralheart Treads (226803, -1442.05 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 37.5 stamina points (3232.49 DPS) | yes | Ring of Awareness (272409, -966.90 DPS) [vendor]; Cutthroat's Signet (272408, -1361.14 DPS) [vendor]; Protector's Band (19514, -1694.40 DPS) [rep] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | 26.5 stamina points (2286.19 DPS) | yes | Ring of Awareness (272409, -20.61 DPS) [vendor]; Cutthroat's Signet (272408, -414.84 DPS) [vendor]; Protector's Band (19514, -748.11 DPS) [rep] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (41472.1 DPS) | yes | Talisman of Arathor (20071, -870.33 DPS) [rep]; Smotts' Compass (4130, -1022.72 DPS) [quest]; Arena Grand Master (19024, -1022.72 DPS) [quest] |
| trinket2 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-verified (41472.1 DPS) | yes | Talisman of Arathor (20071, -494.03 DPS, sim-verified) [rep]; Smotts' Compass (4130, -883.07 DPS) [quest]; Arena Grand Master (19024, -883.07 DPS) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (41472.1 DPS) | yes | Electrified Dagger (19100, +0.00 DPS) [rep]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | Tome of the Darkspear Prophecy (272087) | Creeg Bothunk [vendor] | 17.0 stamina points (1466.90 DPS) | yes | Grand Marshal's Tome of Power (234589, +0.00 DPS) [pvp]; Grand Marshal's Tome of Restoration (234590, +0.00 DPS) [pvp]; Lei of the Lifegiver (19312, -431.44 DPS) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Defender's Grip Stabilizer; trinket2: Stormpike Insignia Rank 6; main_hand: Shadowsong's Sorrow; off_hand: Tome of the Darkspear Prophecy

No-known-source sample (15 of 1222, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 3.1s. Verify run: 2.8s. 178 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.081 ± 0.003, defense=not significant (0.118 ± 0.036) per rating point (1 rating = 1%, 0.118 per %), dodge=0.090 ± 0.012 per rating point (12 rating = 1%, 1.080 per %), strength=0.053 ± 0.000, agility=0.144 ± 0.007, attack_power=0.026 ± 0.000, hit=0.040 ± 0.007 per rating point (10 rating = 1%, 0.402 per %), crit=0.019 ± 0.002 per rating point (14 rating = 1%, 0.268 per %), expertise=1.218 ± 0.082

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.5 stamina points (354.30 DPS) | yes | Brawler's Leather Hood (252504, -22.36 DPS) [crafted]; Defender's Leather Hood (252447, -36.29 DPS) [crafted]; Totemic Leather Hood (252448, -44.35 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (62.96 DPS) | yes | Scholarly Pendant (277203, -5.50 DPS) [quest]; Scout's Medallion (20442, -8.16 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.2 stamina points (176.33 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -56.63 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -57.65 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.6 stamina points (146.08 DPS) | yes | Sentry Cloak (2059, -6.60 DPS) [world_drop]; Grave Shroud (279865, -10.63 DPS) [quest]; Firebane Cloak (12979, -20.71 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 17.9 stamina points (342.49 DPS) | yes | Gloomshroud Armor (1489, -38.44 DPS, sim-verified) [dungeon]; Trapper's Leather Armor (252491, -95.03 DPS) [crafted]; Camouflaged Tunic (3585, -95.86 DPS) [quest] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.7 stamina points (166.34 DPS) | yes | Savannah Bracers (15453, -14.48 DPS) [quest]; Bear Bracers (4795, -17.61 DPS, sim-verified) [vendor]; Drakewing Bands (12999, -45.00 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 8.9 stamina points (170.19 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.56 DPS) [crafted]; Defender's Leather Gloves (252496, -6.97 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.0 stamina points (211.15 DPS) | yes | Belt of the Fang (10412, -31.29 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.75 DPS) [crafted]; Guardsman Belt (3429, -50.31 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 13.7 stamina points (262.75 DPS) | yes | Slick Deviate Leggings (6480, -0.82 DPS) [quest]; Smelting Pants (5199, -10.26 DPS) [dungeon]; Brawler's Leather Pants (252500, -27.58 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 14.9 stamina points (285.93 DPS) | yes | Surfer Shoes (276274, -52.52 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -59.47 DPS) [dungeon]; Trapper's Leather Boots (252440, -81.38 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.2 stamina points (100.28 DPS) | yes | Ring of Scorn (3235, -23.67 DPS) [quest]; Ring of the Moon (12052, -41.81 DPS) [world_drop]; Deep Fathom Ring (6463, -42.82 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (95.77 DPS) | yes | Ring of Scorn (3235, -16.75 DPS, sim-verified) [quest]; Ring of the Moon (12052, -37.30 DPS) [world_drop]; Deep Fathom Ring (6463, -38.31 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (2229.7 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.2 stamina points (311.17 DPS) | yes | Advisor's Gnarled Staff (20425, -19.22 DPS) [pvp]; Hammerbone (270018, -22.79 DPS) [quest]; Crescent Staff (6505, -23.24 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 178, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 248005 Northshire Hammer; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263006 Scout Ranger's Tunic; 263007 Skyseer's Vest

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 42.0. Weights run: 2.7s. Verify run: 3.1s. 332 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.006, defense=not significant (0.213 ± 0.068) per rating point (1 rating = 1%, 0.213 per %), dodge=0.100 ± 0.017 per rating point (12 rating = 1%, 1.205 per %), strength=0.044 ± 0.000, agility=0.143 ± 0.011, attack_power=0.022 ± 0.000, hit=not significant (0.036 ± 0.012) per rating point (10 rating = 1%, 0.363 per %), crit=0.018 ± 0.002 per rating point (14 rating = 1%, 0.247 per %), expertise=1.349 ± 0.141

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.5 stamina points (410.11 DPS) | yes | Defender's Leather Helm (252455, -17.22 DPS) [crafted]; Totemic Leather Helm (252456, -27.33 DPS) [crafted]; Trapper's Leather Helm (252513, -27.33 DPS) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (248.31 DPS) | yes | River Pride Choker (13087, -68.82 DPS, sim-verified) [world_drop]; Senior Sergeant's Insignia (15200, -76.40 DPS) [vendor]; Darkspear Warding Pendant (272075, -95.50 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 17.8 stamina points (340.70 DPS) | yes | Tanned Shoulderpads (270023, -48.14 DPS, sim-verified) [quest]; Cloudy Gustwoven Spaulders (277043, -66.20 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -66.20 DPS) [crafted] |
| back | Grimsteel Cape (4643) | Vorrel's Revenge [quest] | 11.6 stamina points (221.73 DPS) | yes | Sergeant's Cloak (18427, -8.20 DPS) [vendor]; Tigerstrike Mantle (13108, -26.14 DPS) [world_drop]; Enduring Cape (14763, -45.65 DPS) [world_drop] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.2 stamina points (405.19 DPS) | yes | Spirewind Fetter (9406, -9.89 DPS) [dungeon]; Green Whelp Armor (7375, -31.80 DPS) [crafted]; Dokebi Chestguard (14581, -57.30 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.9 stamina points (208.35 DPS) | yes | Black Wolf Bracers (3230, -4.13 DPS) [dungeon]; Barbaric Bracers (18948, -4.21 DPS) [crafted]; Technician's Bracers (270042, -33.40 DPS) [quest] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 13.4 stamina points (255.29 DPS) | yes | Brawler Gloves (720, -15.60 DPS) [world_drop]; Wolfclaw Gloves (1978, -22.42 DPS) [dungeon]; Emblazoned Gloves (6397, -35.77 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.4 stamina points (294.21 DPS) | yes | Skulker's Leather Belt (252520, -32.18 DPS, sim-verified) [crafted]; Stalker's Leather Belt (252521, -54.16 DPS) [crafted]; Prowler's Leather Belt (252459, -54.77 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 17.9 stamina points (341.01 DPS) | yes | Dark Ritual Leggings (270031, -28.23 DPS) [quest]; Brawler's Leather Legguards (252516, -32.73 DPS) [crafted]; Trapper's Leather Legguards (252517, -38.63 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.4 stamina points (350.87 DPS) | yes | Gravewalker Boots (251963, -15.87 DPS) [quest]; Harbinger Boots (7754, -25.42 DPS) [dungeon]; Grizzled Boots (6335, -48.48 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.2 stamina points (156.59 DPS) | yes | Plains Ring (2039, -3.79 DPS) [dungeon]; Darkspear Signet (272071, -3.79 DPS) [vendor]; Ring of Power Regulation (273030, -3.79 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.1 stamina points (155.33 DPS) | yes | Plains Ring (2039, -2.53 DPS) [dungeon]; Darkspear Signet (272071, -2.53 DPS) [vendor]; Ring of Power Regulation (273030, -2.53 DPS) [dungeon] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | sim-verified (3488.1 DPS) | yes | Rune of Duty (21568, -38.20 DPS) [rep] |
| trinket2 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (3488.1 DPS) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (3488.1 DPS) | yes | Advisor's Gnarled Staff (19569, -44.27 DPS) [pvp]; Glimmering Staff (249392, -63.19 DPS) [crafted]; Manual Crowd Pummeler (9449, -217.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenir Sea Shell; shoulder: Watchman Pauldrons; back: Grimsteel Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Defiler's Talisman; trinket2: Rune of Perfection; main_hand: Staff of the Shade

No-known-source sample (15 of 332, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 73.6. Weights run: 3.1s. Verify run: 3.3s. 431 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.028, armor=0.103 ± 0.009, defense=0.495 ± 0.121 per rating point (1 rating = 1%, 0.495 per %), dodge=0.119 ± 0.025 per rating point (12 rating = 1%, 1.424 per %), strength=0.046 ± 0.000, agility=0.136 ± 0.016, attack_power=0.023 ± 0.000, hit=not significant (0.049 ± 0.021) per rating point (10 rating = 1%, 0.491 per %), crit=not significant (0.027 ± 0.008) per rating point (14 rating = 1%, 0.375 per %), expertise=1.681 ± 0.227

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 24.9 stamina points (702.29 DPS) | yes | Warden's Wizard Hat (14604, -54.66 DPS) [world_drop]; Defender's Leather Helm (252455, -76.08 DPS) [crafted]; Brawler's Leather Helm (252512, -107.87 DPS, sim-verified) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (366.12 DPS) | yes | Darkspear Warding Pendant (272074, -56.33 DPS) [vendor]; Gazlowe's Charm (13088, -58.42 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -107.46 DPS) [world_drop] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 25.9 stamina points (728.40 DPS) | yes | Flintrock Shoulders (7755, -150.35 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -181.83 DPS) [dungeon]; Tanned Shoulderpads (270023, -262.43 DPS) [quest] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.4 stamina points (378.42 DPS) | yes | Silky Spider Cape (10776, -9.46 DPS) [dungeon]; Well Oiled Cloak (12254, -12.37 DPS) [vendor]; Heraldic Cloak (8120, -22.34 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 28.7 stamina points (807.85 DPS) | yes | Warden's Wraps (14601, -74.74 DPS) [world_drop]; Insignia Chestguard (4057, -141.89 DPS) [world_drop]; Robes of the Lich (10762, -206.37 DPS, sim-verified) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 13.2 stamina points (371.43 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -29.37 DPS) [dungeon]; Cultist's Armguards (270032, -39.72 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | 18.3 stamina points (514.55 DPS) | yes | Stalker's Leather Gloves (252526, +0.00 DPS, sim-verified) [crafted]; Bonefingers (10765, -25.29 DPS) [dungeon]; Razzeric's Racing Grips (6727, -26.75 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 18.0 stamina points (506.20 DPS) | yes | Warden's Leather Belt (252460, -38.64 DPS) [crafted]; Scorpashi Sash (14652, -67.73 DPS) [world_drop]; Kolkar Hunter's Belt (6788, -92.61 DPS, sim-verified) [quest] |
| legs | Tidesoaked Leggings (279888) | The Only Good Brinescale... [quest] | 23.3 stamina points (656.91 DPS) | yes | Imperial Leather Pants (4062, -10.38 DPS) [dungeon]; Warden's Woolies (14605, -19.11 DPS) [world_drop]; Basilisk Hide Pants (1718, -22.45 DPS) [world_drop] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 21.8 stamina points (613.43 DPS) | yes | Gravewalker Boots (251963, -78.00 DPS) [quest]; Harbinger Boots (7754, -94.35 DPS) [dungeon]; Gnomebot Operating Boots (9450, -127.48 DPS, sim-verified) [dungeon] |
| finger1 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.3 stamina points (318.88 DPS) | yes | Darkspear Signet (272070, -9.08 DPS) [vendor]; Underworld Band (1980, -37.25 DPS) [world_drop]; Nogg's Gold Ring (9588, -37.25 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.3 stamina points (317.58 DPS) | yes | Darkspear Signet (272070, -7.79 DPS) [vendor]; Underworld Band (1980, -35.95 DPS) [world_drop]; Nogg's Gold Ring (9588, -35.95 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | sim-verified (6772.6 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -28.16 DPS) [world]; Darkspear Voodoo Seal (272059, -28.16 DPS) [vendor] |
| trinket2 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (6772.6 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -28.16 DPS) [world]; Darkspear Voodoo Seal (272059, -28.16 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (6772.6 DPS) | yes | Tok'kar's Murloc Basher (9678, -159.26 DPS) [quest]; Mograine's Might (7723, -201.50 DPS) [dungeon]; Manual Crowd Pummeler (9449, -586.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Tidesoaked Leggings; feet: Warden's Leather Shoes; finger1: Suspicious Spare Part; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Cragwood Maul

No-known-source sample (15 of 431, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 116.0. Weights run: 3.2s. Verify run: 3.8s. 562 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.043, armor=0.153 ± 0.017, defense=not significant (0.404 ± 0.219) per rating point (1 rating = 1%, 0.404 per %), dodge=0.260 ± 0.045 per rating point (12 rating = 1%, 3.115 per %), strength=0.056 ± 0.000, agility=0.242 ± 0.028, attack_power=0.028 ± 0.000, hit=not significant (0.093 ± 0.037) per rating point (10 rating = 1%, 0.928 per %), crit=not significant (0.031 ± 0.015) per rating point (14 rating = 1%, 0.438 per %), expertise=3.214 ± 0.412

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.6 stamina points (1558.02 DPS) | yes | Blood Guard's Leather Headband (220851, -151.36 DPS) [vendor]; Impractical Headwarmer (274756, -321.65 DPS, sim-verified) [vendor]; Scorpashi Skullcap (14658, -357.18 DPS) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (512.21 DPS) | yes | Darkspear Warding Pendant (272073, -36.59 DPS) [vendor]; Souvenir Sea Shell (274749, -36.59 DPS) [vendor]; Woven Ivy Necklace (19159, -54.28 DPS) [quest] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | 31.3 stamina points (1145.45 DPS) | yes | Fleshhide Shoulders (10774, -0.06 DPS) [dungeon]; Blood Guard's Leather Shoulders (220853, -22.59 DPS) [vendor]; Warden's Leather Shoulder (252536, -78.86 DPS) [crafted] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 19.8 stamina points (725.48 DPS) | yes | Sergeant's Cloak (16341, -11.22 DPS) [vendor]; Grovekeeper's Drape (17739, -78.79 DPS) [dungeon]; Wing of the Whelpling (13121, -166.99 DPS) [world_drop] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.2 stamina points (2056.68 DPS) | yes | Jinxed Hoodoo Skin (9473, -552.37 DPS, sim-verified) [dungeon]; Heraldic Breastplate (8119, -559.12 DPS) [world_drop]; Stone Guard's Leather Armor (220855, -580.86 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.3 stamina points (924.63 DPS) | yes | Forest Stalker's Bracers (19589, -78.45 DPS) [pvp]; First Sergeant's Dragonhide Armguards (18436, -130.14 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -168.66 DPS) [crafted] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 39.8 stamina points (1456.99 DPS) | yes | Warden's Leather Gauntlets (252549, -290.97 DPS) [crafted]; Raider Gloves (272100, -338.61 DPS, sim-verified) [vendor]; Feralheart Fists (226793, -354.08 DPS) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.3 stamina points (1036.19 DPS) | yes | Skulker's Leather Waistguard (252474, -22.17 DPS) [crafted]; Prowler's Leather Waistguard (252473, -35.83 DPS) [crafted]; Skycaller's Leather Waistguard (252476, -170.81 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 37.4 stamina points (1369.19 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -219.15 DPS, sim-verified) [vendor]; Stone Guard's Restored Leather Leggings (220883, -235.24 DPS) [vendor]; Scorpashi Leggings (14659, -252.92 DPS) [world_drop] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | 29.6 stamina points (1081.75 DPS) | yes | Skulker's Leather Boots (252469, -31.04 DPS) [crafted]; Prowler's Leather Boots (252468, -44.70 DPS) [crafted]; Slitherscale Boots (10801, -48.82 DPS) [dungeon] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.4 stamina points (527.51 DPS) | yes | Ring of Saviors (1447, -15.29 DPS) [world_drop]; Darkspear Signet (272069, -15.29 DPS) [vendor]; Suspicious Spare Part (274754, -110.78 DPS) [vendor] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (512.21 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Suspicious Spare Part (274754, -95.49 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (12157.7 DPS) | yes | Defiler's Talisman (21115, +0.00 DPS) [rep]; Guardian Talisman (1490, -36.59 DPS) [quest]; Rune of Perfection (21565, -73.17 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272060) | Creeg Bothunk [vendor] | sim-verified (12157.7 DPS) | yes | Defiler's Talisman (21115, +0.00 DPS) [rep]; Guardian Talisman (1490, -36.59 DPS) [quest]; Rune of Perfection (21565, -73.17 DPS) [rep] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (12157.7 DPS) | yes | Radiant Staff (249453, -10.51 DPS) [crafted]; Dreamstaff (249454, -10.51 DPS) [crafted]; Manual Crowd Pummeler (9449, -764.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; shoulder: Penance Spaulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Warden's Leather Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal

No-known-source sample (15 of 562, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 146.2. Weights run: 3.0s. Verify run: 9.9s. 1218 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.029, armor=0.177 ± 0.022, defense=not significant (1.151 ± 0.330) per rating point (1 rating = 1%, 1.151 per %), dodge=0.592 ± 0.069 per rating point (12 rating = 1%, 7.106 per %), strength=0.111 ± 0.000, agility=0.494 ± 0.043, attack_power=0.055 ± 0.000, hit=not significant (0.166 ± 0.053) per rating point (10 rating = 1%, 1.665 per %), crit=not significant (0.074 ± 0.023) per rating point (14 rating = 1%, 1.040 per %), expertise=5.890 ± 0.575

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 71.2 stamina points (4607.14 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 20.5 stamina points (1326.62 DPS) | yes | Medallion of Grand Marshal Morris (13091, -128.59 DPS) [world_drop]; Scout's Medallion (19534, -199.66 DPS) [rep]; Senior Sergeant's Insignia (16335, -226.40 DPS) [vendor] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 67.1 stamina points (4340.94 DPS) | yes | Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, -446.72 DPS) [vendor]; Defiler's Leather Shoulders (20194, -975.82 DPS, sim-verified) [rep] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 32.0 stamina points (2069.82 DPS) | yes | Fel Cape (279269, -282.89 DPS, sim-verified) [crafted]; Sergeant's Cloak (18461, -455.42 DPS) [vendor]; Cloak of Earth and Sky (279270, -463.00 DPS) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | 87.5 stamina points (5663.14 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Feralheart Armor (226804, -508.97 DPS, sim-verified) [vendor] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | 36.8 stamina points (2378.98 DPS) | yes | Bracers of Subterfuge (22668, -125.33 DPS) [quest]; Forest Stalker's Bracers (19587, -191.96 DPS) [rep]; First Sergeant's Dragonhide Armguards (18434, -398.93 DPS) [vendor] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 53.9 stamina points (3488.98 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, -27.72 DPS) [pvp] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 49.7 stamina points (3218.67 DPS) | yes | Shifter's Belt (272396, -309.41 DPS, sim-verified) [vendor]; Hivethrasher's Girdle (275614, -573.04 DPS) [crafted]; Belt of Preserved Heads (20216, -584.79 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 89.9 stamina points (5818.93 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; General's Dragonhide Legguards (231673, -328.27 DPS) [vendor]; Dire Warbear Woolies (227804, -384.87 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 65.9 stamina points (4266.65 DPS) | yes | General's Dragonhide Treads (231683, -169.27 DPS) [vendor]; Drudge Boots (21532, -547.59 DPS, sim-verified) [quest]; Feralheart Treads (226803, -599.36 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 33.5 stamina points (2167.99 DPS) | yes | Ring of Awareness (272409, -792.92 DPS) [vendor]; Insurgent's Band (272064, -1071.52 DPS) [vendor]; Cutthroat's Signet (272408, -1094.70 DPS) [vendor] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | 23.8 stamina points (1537.43 DPS) | yes | Ring of Awareness (272409, -162.37 DPS) [vendor]; Insurgent's Band (272064, -440.97 DPS) [vendor]; Cutthroat's Signet (272408, -464.15 DPS) [vendor] |
| trinket1 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-verified (22302.2 DPS) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Relentless Raider's Seal (272060, -272.55 DPS) [vendor]; Darkspear Voodoo Seal (272061, -272.55 DPS) [vendor] |
| trinket2 | Defiler's Talisman (20072) | The Defilers [rep] | sim-verified (22302.2 DPS) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Relentless Raider's Seal (272060, -129.44 DPS) [vendor]; Darkspear Voodoo Seal (272061, -129.44 DPS) [vendor] |
| main_hand | Whiteout Staff (19101) | Frostwolf Clan [rep] | sim-verified (22302.2 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glacial Blade (19099, -1073.58 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Frostwolf Insignia Rank 6; trinket2: Defiler's Talisman; main_hand: Whiteout Staff

No-known-source sample (15 of 1218, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 374.7. Weights run: 2.8s. Verify run: 10.6s. 1218 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.008, armor=0.200 ± 0.026, defense=not significant (1.699 ± 0.437) per rating point (1 rating = 1%, 1.699 per %), dodge=0.853 ± 0.084 per rating point (12 rating = 1%, 10.234 per %), strength=0.121 ± 0.000, agility=0.762 ± 0.055, attack_power=0.055 ± 0.000, hit=0.303 ± 0.062 per rating point (10 rating = 1%, 3.026 per %), crit=0.160 ± 0.028 per rating point (14 rating = 1%, 2.246 per %), expertise=8.111 ± 0.675

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 82.3 stamina points (7103.24 DPS) | yes | Feralheart Faceguard (226801, +0.00 DPS) [vendor]; Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 25.7 stamina points (2215.52 DPS) | yes | Scout's Medallion (19534, -366.97 DPS) [rep]; Medallion of Grand Marshal Morris (13091, -649.42 DPS, sim-verified) [world_drop]; Woven Ivy Necklace (19159, -698.76 DPS) [quest] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 80.5 stamina points (6949.68 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -389.04 DPS) [vendor]; Darkspear Pauldrons (272105, -1163.48 DPS, sim-verified) [vendor]; Warlord's Dragonhide Spaulders (231681, -1323.35 DPS) [vendor] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 40.8 stamina points (3517.79 DPS) | yes | Fel Cape (279269, -1126.36 DPS, sim-verified) [crafted]; Deathguard's Cloak (20068, -1217.34 DPS) [rep]; Stalwart Cloak (272415, -1255.82 DPS) [vendor] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | 103.8 stamina points (8954.69 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -675.82 DPS) [vendor]; Feralheart Armor (226804, -1092.95 DPS, sim-verified) [vendor] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | 41.6 stamina points (3593.90 DPS) | yes | Forest Stalker's Bracers (19587, -58.26 DPS) [rep]; Bracers of Subterfuge (22668, -72.54 DPS) [quest] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 65.4 stamina points (5639.58 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, -718.18 DPS, sim-verified) [vendor]; General's Dragonhide Gloves (231677, -758.57 DPS) [pvp] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 58.4 stamina points (5043.38 DPS) | yes | Belt of Preserved Heads (20216, -840.21 DPS) [quest]; Ferocity of the Timbermaw (227805, -965.79 DPS) [vendor]; Shifter's Belt (272396, -979.30 DPS, sim-verified) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 104.2 stamina points (8991.36 DPS) | yes | General's Dragonhide Leggings (231685, -253.76 DPS) [pvp]; Dire Warbear Woolies (227804, -652.99 DPS, sim-verified) [vendor]; General's Dragonhide Legguards (231673, -1298.51 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 80.0 stamina points (6899.22 DPS) | yes | General's Dragonhide Treads (231683, -797.49 DPS) [vendor]; Drudge Boots (21532, -882.78 DPS, sim-verified) [quest]; Feralheart Treads (226803, -1442.05 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 37.5 stamina points (3232.49 DPS) | yes | Ring of Awareness (272409, -966.90 DPS) [vendor]; Cutthroat's Signet (272408, -1361.14 DPS) [vendor]; Legionnaire's Band (19510, -1694.40 DPS) [rep] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | 26.5 stamina points (2286.19 DPS) | yes | Ring of Awareness (272409, -20.61 DPS) [vendor]; Cutthroat's Signet (272408, -414.84 DPS) [vendor]; Legionnaire's Band (19510, -748.11 DPS) [rep] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (43667.2 DPS) | yes | Defiler's Talisman (20072, -870.33 DPS) [rep]; Smotts' Compass (4130, -1022.72 DPS) [quest]; Arena Grand Master (19024, -1022.72 DPS) [quest] |
| trinket2 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-verified (43667.2 DPS) | yes | Defiler's Talisman (20072, -784.43 DPS, sim-verified) [rep]; Smotts' Compass (4130, -883.07 DPS) [quest]; Arena Grand Master (19024, -883.07 DPS) [quest] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | sim-verified (43667.2 DPS) | yes | Glacial Blade (19099, +0.00 DPS) [rep]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | Tome of the Darkspear Prophecy (272087) | Creeg Bothunk [vendor] | 17.0 stamina points (1466.90 DPS) | yes | High Warlord's Tome of Destruction (234563, +0.00 DPS) [pvp]; High Warlord's Tome of Mending (234564, +0.00 DPS) [pvp]; Lei of the Lifegiver (19312, -431.44 DPS) [rep] |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Defender's Grip Stabilizer; trinket2: Frostwolf Insignia Rank 6; main_hand: Shadowsong's Sorrow; off_hand: Tome of the Darkspear Prophecy

No-known-source sample (15 of 1218, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

