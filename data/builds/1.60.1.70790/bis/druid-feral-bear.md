# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-51220001000000000000-0000000000000000)

Set DPS (verified): 22.7. Weights run: 6.7s. Verify run: 5.9s. 228 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.103 ± 0.000, defense=0.184 ± 0.028 per rating point (1 rating = 1%, 0.184 per %), dodge=0.112 ± 0.008 per rating point (12 rating = 1%, 1.338 per %), strength=0.096 ± 0.000, agility=0.169 ± 0.005, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.043 ± 0.004 per rating point (10 rating = 1%, 0.430 per %), crit=0.017 ± 0.001 per rating point (14 rating = 1%, 0.243 per %), expertise=1.111 ± 0.052

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 19.4 stamina points (344.37 DPS) | yes | Brawler's Leather Hood (252504, -4.69 DPS) [crafted]; Defender's Leather Hood (252447, -14.96 DPS) [crafted]; Trapper's Leather Hood (252505, -28.62 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.7 stamina points (101.31 DPS) | yes | Erudite's Amulet (277204, -44.32 DPS, sim-verified) [quest]; Sentinel's Medallion (20444, -47.89 DPS) [rep]; Scholarly Pendant (277203, -48.10 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 10.6 stamina points (187.47 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -48.72 DPS) [dungeon]; Rugged Spaulders (5254, -65.50 DPS) [dungeon] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -3.96 DPS) [world_drop]; Firebane Cloak (12979, -19.56 DPS) [world_drop]; Caretaker's Cape (20428, -35.48 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 20.0 stamina points (355.22 DPS) | yes | Gloomshroud Armor (1489, -34.29 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -66.14 DPS) [quest]; Trapper's Leather Armor (252491, -90.84 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.8 stamina points (156.05 DPS) | yes | Drakewing Bands (12999, -28.89 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -41.25 DPS) [vendor]; Sanguine Cuffs (14375, -54.15 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 10.1 stamina points (179.41 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.82 DPS) [crafted]; Defender's Leather Gloves (252496, -5.13 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 12.4 stamina points (219.36 DPS) | yes | Belt of the Fang (10412, -29.92 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -49.05 DPS) [crafted]; Guardsman Belt (3429, -50.87 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 16.0 stamina points (283.04 DPS) | yes | Brawler's Leather Pants (252500, -25.65 DPS) [crafted]; Leggings of the Fang (10410, -34.64 DPS) [dungeon]; Smelting Pants (5199, -96.72 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.7 stamina points (278.41 DPS) | yes | Footpads of the Fang (10411, -42.85 DPS, sim-verified) [dungeon]; Surfer Shoes (276274, -48.22 DPS) [vendor]; Trapper's Leather Boots (252440, -63.71 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (88.70 DPS) | yes | Slain Baron's Signet (279867, -32.22 DPS) [quest]; Ring of the Moon (12052, -33.77 DPS) [world_drop]; Protector's Band (20439, -34.43 DPS) [rep] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 4.0 stamina points (70.96 DPS) | yes | Slain Baron's Signet (279867, -14.48 DPS) [quest]; Ring of the Moon (12052, -16.03 DPS) [world_drop]; Protector's Band (20439, -16.69 DPS) [rep] |
| trinket1 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.7 stamina points (367.07 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.82 DPS) [dungeon]; Twisted Chanter's Staff (890, -16.35 DPS) [world_drop]; Staff of the Blessed Seer (2271, -20.19 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Sustaining Ring; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 228, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 0000000000000000-51230332000100100000-0000000000000000)

Set DPS (verified): 41.8. Weights run: 7.0s. Verify run: 12.1s. 400 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.081 ± 0.004, defense=0.357 ± 0.055 per rating point (1 rating = 1%, 0.357 per %), dodge=0.129 ± 0.014 per rating point (12 rating = 1%, 1.549 per %), strength=0.060 ± 0.000, agility=0.172 ± 0.009, attack_power=0.030 ± 0.000, feral_attack_power=0.030 ± 0.000, hit=0.062 ± 0.011 per rating point (10 rating = 1%, 0.621 per %), crit=0.023 ± 0.002 per rating point (14 rating = 1%, 0.329 per %), expertise=1.742 ± 0.127

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.5 stamina points (624.74 DPS) | yes | Defender's Leather Helm (252455, -28.84 DPS) [crafted]; Trapper's Leather Helm (252513, -49.89 DPS) [crafted]; Stormrider's Leather Helm (252514, -49.89 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (268.48 DPS) | yes | Master Sergeant's Insignia (18442, -7.02 DPS) [vendor]; Seal of the Expedition (271920, -36.07 DPS) [rep]; Sentinel's Medallion (19541, -83.31 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 17.6 stamina points (512.62 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -72.15 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -95.93 DPS) [crafted]; Barbaric Shoulders (5964, -126.19 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mourning Shawl (6751, -38.44 DPS) [quest]; Enduring Cape (14763, -54.00 DPS) [world_drop]; Tigerstrike Mantle (13108, -112.79 DPS, sim-verified) [world_drop] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 20.9 stamina points (607.71 DPS) | yes | Spirewind Fetter (9406, -12.96 DPS) [dungeon]; Green Whelp Armor (7375, -48.71 DPS) [crafted]; Blackened Defias Armor (10399, -86.29 DPS) [dungeon] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.9 stamina points (315.42 DPS) | yes | Barbaric Bracers (18948, -3.81 DPS) [crafted]; Black Wolf Bracers (3230, -7.61 DPS) [dungeon]; Glowing Leather Bracers (2017, -66.29 DPS) [world] |
| hands | Brawler Gloves (720) | World drop [world_drop] | 12.5 stamina points (362.93 DPS) | yes | Wolfclaw Gloves (1978, -6.73 DPS) [dungeon]; Gloves of Insight (9698, -16.38 DPS) [quest]; Emblazoned Gloves (6397, -29.88 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.5 stamina points (449.81 DPS) | yes | Prowler's Leather Belt (252459, -81.89 DPS) [crafted]; Stalker's Leather Belt (252521, -82.71 DPS) [crafted]; Skulker's Leather Belt (252520, -153.55 DPS, sim-verified) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 17.9 stamina points (519.93 DPS) | yes | Dark Ritual Leggings (270031, -51.90 DPS) [quest]; Trapper's Leather Legguards (252517, -60.41 DPS) [crafted]; Brawler's Leather Legguards (252516, -186.13 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.1 stamina points (527.01 DPS) | yes | Harbinger Boots (7754, -37.28 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -93.89 DPS) [quest]; Emblazoned Boots (4051, -158.46 DPS) [world_drop] |
| finger1 | Theramore Signet (271911) | Theramore Expeditionary Force [rep] | 8.3 stamina points (240.31 DPS) | yes | Plains Ring (2039, -7.90 DPS) [dungeon]; Defias Renegade Ring (1076, -46.04 DPS) [dungeon]; Protector's Band (19517, -83.64 DPS) [rep] |
| finger2 | Militant's Band (271916) (or Plains Ring (2039)) | Theramore Expeditionary Force [rep] | 8.0 stamina points (232.41 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Defias Renegade Ring (1076, -38.15 DPS) [dungeon]; Protector's Band (19517, -75.74 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (+134.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Shade (2549, -37.48 DPS, sim-verified) [dungeon]; Slaghammer (1976, -59.38 DPS) [dungeon]; Lorekeeper's Staff (212580, -67.56 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Brawler Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Theramore Signet; finger2: Militant's Band; trinket1: Talisman of Arathor; trinket2: Rune of Perfection; main_hand: Spine of the Devourer

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-51230332023102120010-0000000000000000)

Set DPS (verified): 83.3. Weights run: 7.2s. Verify run: 6.4s. 526 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.026, armor=0.158 ± 0.008, defense=0.445 ± 0.098 per rating point (1 rating = 1%, 0.445 per %), dodge=0.172 ± 0.024 per rating point (12 rating = 1%, 2.060 per %), strength=0.086 ± 0.000, agility=0.226 ± 0.015, attack_power=0.043 ± 0.000, feral_attack_power=0.043 ± 0.000, hit=0.108 ± 0.021 per rating point (10 rating = 1%, 1.079 per %), crit=0.050 ± 0.007 per rating point (14 rating = 1%, 0.699 per %), expertise=2.985 ± 0.240

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 30.9 stamina points (983.65 DPS) | yes | Warden's Wizard Hat (14604, -41.62 DPS) [world_drop]; Brawler's Leather Helm (252512, -74.99 DPS) [crafted]; Nightscape Headband (8176, -89.08 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shriveled Heart (9243, +0.00 DPS) [dungeon]; Seal of the Expedition (271919, -10.97 DPS) [rep]; Sentinel's Medallion (19540, -59.28 DPS) [rep] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.8 stamina points (1012.67 DPS) | yes | Flintrock Shoulders (7755, -156.80 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -251.50 DPS) [dungeon]; Nightscape Shoulders (8192, -341.54 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -18.69 DPS) [dungeon]; Heraldic Cloak (8120, -21.78 DPS) [dungeon]; Well Oiled Cloak (12254, -23.72 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.5 stamina points (1130.93 DPS) | yes | Warden's Wraps (14601, -58.56 DPS) [world_drop]; Robes of the Lich (10762, -181.94 DPS) [dungeon]; Spirewind Fetter (9406, -197.74 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.6 stamina points (527.77 DPS) | yes | Branded Leather Bracers (19508, -32.46 DPS) [dungeon]; Barbaric Bracers (18948, -60.42 DPS) [crafted]; Cultist's Armguards (270032, -69.68 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -23.66 DPS) [crafted]; Prowler's Leather Gloves (252524, -28.10 DPS) [crafted]; Stalker's Leather Gloves (252526, -404.75 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.0 stamina points (733.34 DPS) | yes | Warden's Leather Belt (252460, -70.06 DPS) [crafted]; Kolkar Hunter's Belt (6788, -74.65 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -85.32 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 30.6 stamina points (974.45 DPS) | yes | Panther Hunter Leggings (4108, -69.09 DPS) [quest]; Warden's Woolies (14605, -73.90 DPS) [world_drop]; Imperial Leather Pants (4062, -131.82 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 27.7 stamina points (882.86 DPS) | yes | Skulker's Leather Shoes (252531, -98.64 DPS) [crafted]; Prowler's Leather Shoes (252465, -116.42 DPS) [crafted]; Worn Running Boots (9398, -304.93 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (382.18 DPS) | yes | Talvash's Gold Ring (9538, -21.20 DPS) [quest]; Militant's Band (271915, -31.85 DPS) [rep]; Dragonclaw Ring (10710, -63.70 DPS) [quest] |
| finger2 | Theramore Signet (271910) | Theramore Expeditionary Force [rep] | 11.5 stamina points (366.78 DPS) | yes | Talvash's Gold Ring (9538, -5.80 DPS) [quest]; Militant's Band (271915, -16.45 DPS) [rep]; Dragonclaw Ring (10710, -48.30 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (+268.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -31.85 DPS) [world]; Enriched Seal (271905, -31.85 DPS) [rep] |
| trinket2 | Talisman of Arathor (21118) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -31.85 DPS) [world]; Volunteer's Lucky Seal (271904, -31.85 DPS) [rep]; Enriched Seal (271905, -31.85 DPS) [rep] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -156.97 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -367.17 DPS) [dungeon]; Huge Stone Club (1523, -452.57 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Theramore Signet; trinket1: Rune of Duty; trinket2: Talisman of Arathor; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 526, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0000000000000000-51230332023102125510-0000000000000000)

Set DPS (verified): 129.8. Weights run: 7.4s. Verify run: 15.8s. 690 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.058, armor=0.132 ± 0.010, defense=not significant (0.413 ± 0.116) per rating point (1 rating = 1%, 0.413 per %), dodge=0.250 ± 0.030 per rating point (12 rating = 1%, 3.004 per %), strength=0.059 ± 0.000, agility=0.246 ± 0.019, attack_power=0.030 ± 0.000, feral_attack_power=0.030 ± 0.000, hit=0.166 ± 0.029 per rating point (10 rating = 1%, 1.660 per %), crit=0.043 ± 0.009 per rating point (14 rating = 1%, 0.601 per %), expertise=3.407 ± 0.332

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 40.2 stamina points (2144.89 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -186.58 DPS) [vendor]; Embrace of the Lycan (9479, -380.86 DPS) [dungeon]; Scorpashi Skullcap (14658, -508.59 DPS) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 4 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (747.64 DPS) | yes | Shriveled Heart (9243, -53.40 DPS) [dungeon]; Lifeblood Amulet (9641, -53.40 DPS) [dungeon]; Seal of the Expedition (271918, -53.40 DPS) [rep] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fleshhide Shoulders (10774, +0.00 DPS) [dungeon]; Knight-Lieutenant's Leather Shoulders (220852, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -16.09 DPS) [dungeon] |
| back | Sergeant's Cape (18441) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Graverot Cape (11677, +0.00 DPS) [dungeon]; Nightfall Drape (12465, +0.00 DPS) [dungeon]; Emberscale Cape (11311, -14.13 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 53.0 stamina points (2829.10 DPS) | yes | Mixologist's Tunic (12793, -550.32 DPS) [dungeon]; Jinxed Hoodoo Skin (9473, -711.96 DPS) [dungeon]; Heraldic Breastplate (8119, -794.51 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.9 stamina points (1278.55 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -78.69 DPS) [vendor]; Serpentskin Bracers (8257, -241.95 DPS) [world_drop]; Warden's Leather Bracers (252541, -244.90 DPS) [crafted] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -201.44 DPS) [crafted]; Wicked Leather Gauntlets (15083, -241.15 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -251.22 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 26.5 stamina points (1414.41 DPS) | yes | Skulker's Leather Waistguard (252474, -30.01 DPS) [crafted]; Prowler's Leather Waistguard (252473, -49.92 DPS) [crafted]; Girdle of Beastial Fury (11686, -178.87 DPS) [dungeon] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 35.5 stamina points (1897.06 DPS) | yes | Knight's Crackling Leather Leggings (220864, -295.44 DPS) [vendor]; Windscale Sarong (10842, -310.27 DPS) [world]; Knight's Restored Leather Leggings (220882, -384.08 DPS) [vendor] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -43.12 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (801.04 DPS) | yes | Militant's Band (271914, -53.40 DPS) [rep]; Underworld Band (1980, -160.21 DPS) [world_drop]; Talvash's Gold Ring (9538, -200.88 DPS) [quest] |
| finger2 | Theramore Signet (271909) | Theramore Expeditionary Force [rep] | 14.4 stamina points (771.33 DPS) | yes | Militant's Band (271914, -23.69 DPS) [rep]; Underworld Band (1980, -130.50 DPS) [world_drop]; Talvash's Gold Ring (9538, -171.17 DPS) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Rune of Perfection (21565, -106.81 DPS) [rep] |
| trinket2 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Rune of Perfection (21565, -53.40 DPS) [rep] |
| main_hand | Blanchard's Stout (13046) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warden Staff (943, +0.00 DPS) [world_drop]; Ironshod Bludgeon (9408, +0.00 DPS) [dungeon]; Radiant Staff (249453, -94.89 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Penance Spaulders; back: Sergeant's Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Theramore Signet; trinket1: Enriched Seal; trinket2: Guardian Talisman; main_hand: Blanchard's Stout

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 189.7. Weights run: 7.3s. Verify run: 42.7s. 1755 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.149 ± 0.017, defense=1.375 ± 0.282 per rating point (1 rating = 1%, 1.375 per %), dodge=0.764 ± 0.049 per rating point (12 rating = 1%, 9.168 per %), strength=0.103 ± 0.000, agility=0.630 ± 0.032, attack_power=0.051 ± 0.000, feral_attack_power=0.051 ± 0.000, hit=0.173 ± 0.042 per rating point (10 rating = 1%, 1.734 per %), crit=0.151 ± 0.020 per rating point (14 rating = 1%, 2.119 per %), expertise=4.061 ± 0.392

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 71.2 stamina points (10036.51 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, -478.51 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 20.8 stamina points (2925.60 DPS) | yes | Sentinel's Medallion (19538, -463.48 DPS, sim-verified) [rep]; Talisman of Evasion (13177, -478.29 DPS) [dungeon]; Master Sergeant's Insignia (18443, -529.11 DPS) [vendor] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | sim-verified (+1188.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -1188.94 DPS, sim-verified) [rep] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phantasmal Cloak (18689, -703.67 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -1048.71 DPS) [dungeon]; Fel Cape (279269, -1225.11 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 33.9 stamina points (4777.04 DPS) | yes | Yeti Hide Bracers (19113, -377.91 DPS) [world]; Malefic Bracers (18700, -390.95 DPS) [dungeon]; Sergeant Major's Dragonhide Armsplints (18454, -767.07 DPS) [vendor] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+970.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -233.82 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, -233.82 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 43.0 stamina points (6064.75 DPS) | yes | Frostbite Girdle (14502, -491.72 DPS) [dungeon]; Hivethrasher's Girdle (275614, -533.78 DPS) [crafted]; Cloudrunner Girdle (13252, -652.91 DPS) [dungeon] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 90.1 stamina points (12705.81 DPS) | yes | Marshal's Dragonhide Leggings (231691, -299.23 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1188.77 DPS, sim-verified) [vendor]; Marshal's Dragonhide Legguards (231703, -1477.89 DPS) [pvp] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 24.3 stamina points (3426.82 DPS) | yes | Myrmidon's Signet (2246, -263.65 DPS) [world_drop]; Cutthroat's Signet (272408, -810.87 DPS) [vendor]; Band of Flesh (13373, -817.89 DPS) [dungeon] |
| finger2 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, -107.32 DPS) [world_drop]; Cutthroat's Signet (272408, -654.54 DPS) [vendor]; Band of Flesh (13373, -661.56 DPS) [dungeon] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (+832.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Enriched Seal (271903, -1110.76 DPS) [rep]; Volunteer's Lucky Seal (271906, -1110.76 DPS) [rep] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Enriched Seal (271903, -281.94 DPS) [rep]; Volunteer's Lucky Seal (271906, -281.94 DPS) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -778.04 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Dawn Treaders; finger1: Naglering; finger2: Ring of Awareness; trinket1: Defender's Grip Stabilizer; trinket2: Talisman of Arathor; main_hand: Headmaster's Charge

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 453.8. Weights run: 8.0s. Verify run: 48.9s. 1755 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.009, armor=0.196 ± 0.022, defense=1.468 ± 0.288 per rating point (1 rating = 1%, 1.468 per %), dodge=0.553 ± 0.057 per rating point (12 rating = 1%, 6.638 per %), strength=0.099 ± 0.000, agility=0.569 ± 0.037, attack_power=0.045 ± 0.000, feral_attack_power=0.045 ± 0.000, hit=0.236 ± 0.052 per rating point (10 rating = 1%, 2.360 per %), crit=0.092 ± 0.023 per rating point (14 rating = 1%, 1.294 per %), expertise=7.741 ± 0.582

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 76.3 stamina points (8439.88 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, -975.30 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 21.7 stamina points (2398.50 DPS) | yes | Sentinel's Medallion (19538, -348.17 DPS) [rep]; Master Sergeant's Insignia (18443, -517.86 DPS) [vendor]; Stormpike Soldier's Pendant (19097, -649.74 DPS) [rep] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | sim-verified (+1747.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -1747.28 DPS, sim-verified) [rep] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phantasmal Cloak (18689, -952.47 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -1644.95 DPS) [dungeon]; Fel Cape (279269, -1790.35 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 69.3 stamina points (7670.54 DPS) | yes | Breastplate of Bloodthirst (12757, +0.00 DPS) [quest]; Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 36.8 stamina points (4068.28 DPS) | yes | Malefic Bracers (18700, -294.39 DPS) [dungeon]; Sergeant Major's Dragonhide Armsplints (18454, -515.76 DPS) [vendor]; Yeti Hide Bracers (19113, -714.28 DPS, sim-verified) [world] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+978.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 46.4 stamina points (5138.44 DPS) | yes | Frostbite Girdle (14502, -352.53 DPS) [dungeon]; Hivethrasher's Girdle (275614, -355.46 DPS) [crafted]; Cloudrunner Girdle (13252, -760.17 DPS, sim-verified) [dungeon] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 96.4 stamina points (10664.32 DPS) | yes | Marshal's Dragonhide Leggings (231691, -80.81 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -873.23 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -2101.08 DPS, sim-verified) [vendor] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 27.2 stamina points (3003.95 DPS) | yes | Myrmidon's Signet (2246, -573.52 DPS) [world_drop]; Cutthroat's Signet (272408, -976.73 DPS) [vendor]; Band of Flesh (13373, -979.59 DPS) [dungeon] |
| finger2 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, -224.92 DPS) [world_drop]; Cutthroat's Signet (272408, -628.13 DPS) [vendor]; Band of Flesh (13373, -630.99 DPS) [dungeon] |
| trinket1 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-verified (+1762.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Talisman of Arathor (20071, -141.06 DPS) [rep]; Force of Will (11810, -331.69 DPS) [dungeon] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Adaptive Combat Assistant (272437, -398.65 DPS) [vendor]; Talisman of Arathor (20071, -783.85 DPS) [rep]; Force of Will (11810, -974.47 DPS) [dungeon] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -1154.16 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Dawn Treaders; finger1: Naglering; finger2: Ring of Awareness; trinket1: Stormpike Insignia Rank 6; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-51220001000000000000-0000000000000000)

Set DPS (verified): 23.1. Weights run: 6.7s. Verify run: 5.3s. 224 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.103 ± 0.000, defense=0.184 ± 0.028 per rating point (1 rating = 1%, 0.184 per %), dodge=0.112 ± 0.008 per rating point (12 rating = 1%, 1.338 per %), strength=0.096 ± 0.000, agility=0.169 ± 0.005, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.043 ± 0.004 per rating point (10 rating = 1%, 0.430 per %), crit=0.017 ± 0.001 per rating point (14 rating = 1%, 0.243 per %), expertise=1.111 ± 0.052

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 19.4 stamina points (344.37 DPS) | yes | Brawler's Leather Hood (252504, -4.69 DPS) [crafted]; Defender's Leather Hood (252447, -14.96 DPS) [crafted]; Trapper's Leather Hood (252505, -28.62 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (59.20 DPS) | yes | Scout's Medallion (20442, -5.78 DPS) [rep]; Scholarly Pendant (277203, -5.98 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 10.6 stamina points (187.47 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -48.72 DPS) [dungeon]; Rugged Spaulders (5254, -65.50 DPS) [dungeon] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -3.96 DPS) [world_drop]; Firebane Cloak (12979, -19.56 DPS) [world_drop]; Battle Healer's Cloak (20427, -35.48 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 20.0 stamina points (355.22 DPS) | yes | Gloomshroud Armor (1489, -42.13 DPS, sim-verified) [dungeon]; Trapper's Leather Armor (252491, -90.84 DPS) [crafted]; Defender's Leather Armor (252434, -90.86 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.8 stamina points (156.05 DPS) | yes | Drakewing Bands (12999, -28.20 DPS) [world_drop]; Spare Part Bindings (279875, -30.40 DPS, sim-verified) [quest]; Wolf Bracers (4794, -41.25 DPS) [vendor] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 10.1 stamina points (179.41 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.82 DPS) [crafted]; Defender's Leather Gloves (252496, -5.13 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 12.4 stamina points (219.36 DPS) | yes | Belt of the Fang (10412, -35.86 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -49.05 DPS) [crafted]; Guardsman Belt (3429, -50.87 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 16.0 stamina points (283.04 DPS) | yes | Brawler's Leather Pants (252500, -25.65 DPS) [crafted]; Leggings of the Fang (10410, -34.64 DPS) [dungeon]; Smelting Pants (5199, -112.06 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.7 stamina points (278.41 DPS) | yes | Surfer Shoes (276274, -48.22 DPS) [vendor]; Footpads of the Fang (10411, -48.55 DPS, sim-verified) [dungeon]; Trapper's Leather Boots (252440, -63.71 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (88.70 DPS) | yes | Slain Baron's Signet (279867, -32.22 DPS) [quest]; Ring of the Moon (12052, -33.77 DPS) [world_drop]; Legionnaire's Band (20429, -34.43 DPS) [rep] |
| finger2 | Ring of Scorn (3235) | A Husband's Revenge [quest] | 4.0 stamina points (70.96 DPS) | yes | Ring of the Moon (12052, -16.03 DPS) [world_drop]; Legionnaire's Band (20429, -16.69 DPS) [rep]; Slain Baron's Signet (279867, -16.78 DPS, sim-verified) [quest] |
| trinket1 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.7 stamina points (367.07 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.82 DPS) [dungeon]; Twisted Chanter's Staff (890, -16.35 DPS) [world_drop]; Staff of the Blessed Seer (2271, -20.19 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Ring of Scorn; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-51230332000100100000-0000000000000000)

Set DPS (verified): 50.2. Weights run: 7.0s. Verify run: 11.4s. 401 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.106 ± 0.003, defense=0.309 ± 0.053 per rating point (1 rating = 1%, 0.309 per %), dodge=0.150 ± 0.015 per rating point (12 rating = 1%, 1.804 per %), strength=0.096 ± 0.000, agility=0.200 ± 0.009, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.071 ± 0.011 per rating point (10 rating = 1%, 0.706 per %), crit=0.027 ± 0.001 per rating point (14 rating = 1%, 0.376 per %), expertise=1.894 ± 0.128

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 23.9 stamina points (532.90 DPS) | yes | Trapper's Leather Helm (252513, -44.70 DPS) [crafted]; Stormrider's Leather Helm (252514, -44.70 DPS) [crafted]; Defender's Leather Helm (252455, -58.32 DPS, sim-verified) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.4 stamina points (209.36 DPS) | yes | Senior Sergeant's Insignia (15200, -8.57 DPS) [vendor]; Darkspear Warding Pendant (272075, -30.88 DPS) [vendor]; Scout's Medallion (19537, -62.05 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 19.7 stamina points (440.57 DPS) | yes | Azure Gustwoven Spaulders (277051, -73.81 DPS) [crafted]; Barbaric Shoulders (5964, -94.35 DPS) [crafted]; Cloudy Gustwoven Spaulders (277043, -95.34 DPS, sim-verified) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enduring Cape (14763, -40.83 DPS) [world_drop]; Grimsteel Cape (4643, -42.76 DPS) [quest]; Tigerstrike Mantle (13108, -133.79 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | 23.5 stamina points (524.63 DPS) | yes | Raptor Hide Harness (4455, -3.40 DPS) [crafted]; Green Whelp Armor (7375, -38.58 DPS) [crafted]; Blackened Defias Armor (10399, -69.59 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.2 stamina points (271.19 DPS) | yes | Cultist's Armguards (270032, -0.51 DPS) [quest]; Black Wolf Bracers (3230, -9.29 DPS) [dungeon]; Glowing Leather Bracers (2017, -56.33 DPS) [world] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 15.2 stamina points (338.45 DPS) | yes | Wolfclaw Gloves (1978, -23.15 DPS) [dungeon]; Emblazoned Gloves (6397, -47.76 DPS) [world_drop]; Brawler Gloves (720, -49.75 DPS, sim-verified) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 17.4 stamina points (389.09 DPS) | yes | Skulker's Leather Belt (252520, -53.52 DPS) [crafted]; Prowler's Leather Belt (252459, -60.51 DPS) [crafted]; Larcenist's Belt (279846, -168.88 DPS, sim-verified) [quest] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 20.6 stamina points (459.40 DPS) | yes | Petrolspill Leggings (9509, -50.63 DPS) [dungeon]; Trapper's Leather Legguards (252517, -50.67 DPS) [crafted]; Brawler's Leather Legguards (252516, -207.02 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 20.0 stamina points (447.02 DPS) | yes | Harbinger Boots (7754, -37.77 DPS, sim-verified) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -94.37 DPS) [quest]; Gravewalker Boots (251963, -101.00 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.4 stamina points (188.12 DPS) | yes | Plains Ring (2039, -9.64 DPS) [dungeon]; Darkspear Signet (272071, -9.64 DPS) [vendor]; Defias Renegade Ring (1076, -36.38 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.3 stamina points (184.91 DPS) | yes | Plains Ring (2039, -6.43 DPS) [dungeon]; Darkspear Signet (272071, -6.43 DPS) [vendor]; Defias Renegade Ring (1076, -33.16 DPS) [dungeon] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | sim-verified (+148.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -6.45 DPS) [world]; Slaghammer (1976, -26.32 DPS) [dungeon]; Advisor's Gnarled Staff (212584, -36.31 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Defiler's Talisman; trinket2: Rune of Duty; main_hand: Viscous Hammer

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-51230332023102120010-0000000000000000)

Set DPS (verified): 82.2. Weights run: 7.2s. Verify run: 6.5s. 525 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.026, armor=0.158 ± 0.008, defense=0.445 ± 0.098 per rating point (1 rating = 1%, 0.445 per %), dodge=0.172 ± 0.024 per rating point (12 rating = 1%, 2.060 per %), strength=0.086 ± 0.000, agility=0.226 ± 0.015, attack_power=0.043 ± 0.000, feral_attack_power=0.043 ± 0.000, hit=0.108 ± 0.021 per rating point (10 rating = 1%, 1.079 per %), crit=0.050 ± 0.007 per rating point (14 rating = 1%, 0.699 per %), expertise=2.985 ± 0.240

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 30.9 stamina points (983.65 DPS) | yes | Warden's Wizard Hat (14604, -41.62 DPS) [world_drop]; Brawler's Leather Helm (252512, -74.99 DPS) [crafted]; Nightscape Headband (8176, -89.08 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shriveled Heart (9243, +0.00 DPS) [dungeon]; Dragon's Blood Necklace (10711, +0.00 DPS) [quest]; Darkspear Warding Pendant (272074, -10.97 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.8 stamina points (1012.67 DPS) | yes | Flintrock Shoulders (7755, -126.06 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -251.50 DPS) [dungeon]; Nightscape Shoulders (8192, -341.54 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -18.69 DPS) [dungeon]; Heraldic Cloak (8120, -21.78 DPS) [dungeon]; Well Oiled Cloak (12254, -23.72 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.5 stamina points (1130.93 DPS) | yes | Warden's Wraps (14601, -58.56 DPS) [world_drop]; Robes of the Lich (10762, -181.94 DPS) [dungeon]; Spirewind Fetter (9406, -197.74 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.6 stamina points (527.77 DPS) | yes | Branded Leather Bracers (19508, -32.46 DPS) [dungeon]; Barbaric Bracers (18948, -60.42 DPS) [crafted]; Cultist's Armguards (270032, -69.68 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -23.66 DPS) [crafted]; Prowler's Leather Gloves (252524, -28.10 DPS) [crafted]; Stalker's Leather Gloves (252526, -365.22 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.0 stamina points (733.34 DPS) | yes | Kolkar Hunter's Belt (6788, -62.07 DPS) [quest]; Warden's Leather Belt (252460, -70.06 DPS) [crafted]; Scorpashi Sash (14652, -85.32 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 30.6 stamina points (974.45 DPS) | yes | Panther Hunter Leggings (4108, -69.09 DPS) [quest]; Warden's Woolies (14605, -73.90 DPS) [world_drop]; Imperial Leather Pants (4062, -124.82 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 27.7 stamina points (882.86 DPS) | yes | Skulker's Leather Shoes (252531, -98.64 DPS) [crafted]; Prowler's Leather Shoes (252465, -116.42 DPS) [crafted]; Worn Running Boots (9398, -324.98 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (382.18 DPS) | yes | Nogg's Gold Ring (9588, -21.20 DPS) [quest]; Darkspear Signet (272070, -31.85 DPS) [vendor]; Dragonclaw Ring (10710, -63.70 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.5 stamina points (366.78 DPS) | yes | Nogg's Gold Ring (9588, -5.80 DPS) [quest]; Darkspear Signet (272070, -16.45 DPS) [vendor]; Dragonclaw Ring (10710, -48.30 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (+260.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -31.85 DPS) [world]; Darkspear Voodoo Seal (272059, -31.85 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21116) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -31.85 DPS) [world]; Darkspear Voodoo Seal (272059, -31.85 DPS) [vendor]; Relentless Raider's Seal (272062, -31.85 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -121.04 DPS, sim-verified) [dungeon]; Mograine's Might (7723, -264.28 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -283.65 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Insurgent's Band; trinket1: Rune of Duty; trinket2: Defiler's Talisman; main_hand: Cragwood Maul

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0000000000000000-51230332023102125510-0000000000000000)

Set DPS (verified): 120.4. Weights run: 7.4s. Verify run: 16.3s. 689 eligible items had no known source.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.058, armor=0.132 ± 0.010, defense=not significant (0.413 ± 0.116) per rating point (1 rating = 1%, 0.413 per %), dodge=0.250 ± 0.030 per rating point (12 rating = 1%, 3.004 per %), strength=0.059 ± 0.000, agility=0.246 ± 0.019, attack_power=0.030 ± 0.000, feral_attack_power=0.030 ± 0.000, hit=0.166 ± 0.029 per rating point (10 rating = 1%, 1.660 per %), crit=0.043 ± 0.009 per rating point (14 rating = 1%, 0.601 per %), expertise=3.407 ± 0.332

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 40.2 stamina points (2144.89 DPS) | yes | Blood Guard's Leather Headband (220851, -186.58 DPS) [vendor]; Embrace of the Lycan (9479, -380.86 DPS) [dungeon]; Scorpashi Skullcap (14658, -508.59 DPS) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 4 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (747.64 DPS) | yes | Shriveled Heart (9243, -53.40 DPS) [dungeon]; Lifeblood Amulet (9641, -53.40 DPS) [dungeon]; Darkspear Warding Pendant (272073, -53.40 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Blood Guard's Leather Shoulders (220853, -12.18 DPS) [vendor]; Penance Spaulders (11963, -22.51 DPS) [quest]; Phytoskin Spaulders (17749, -38.61 DPS) [dungeon] |
| back | Sergeant's Cloak (16341) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Graverot Cape (11677, +0.00 DPS) [dungeon]; Nightfall Drape (12465, +0.00 DPS) [dungeon]; Emberscale Cape (11311, -14.13 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 53.0 stamina points (2829.10 DPS) | yes | Mixologist's Tunic (12793, -550.32 DPS) [dungeon]; Jinxed Hoodoo Skin (9473, -711.96 DPS) [dungeon]; Heraldic Breastplate (8119, -794.51 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.9 stamina points (1278.55 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -78.69 DPS) [vendor]; Serpentskin Bracers (8257, -241.95 DPS) [world_drop]; Warden's Leather Bracers (252541, -244.90 DPS) [crafted] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -201.44 DPS) [crafted]; Wicked Leather Gauntlets (15083, -241.15 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -251.22 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 26.5 stamina points (1414.41 DPS) | yes | Skulker's Leather Waistguard (252474, -30.01 DPS) [crafted]; Prowler's Leather Waistguard (252473, -49.92 DPS) [crafted]; Girdle of Beastial Fury (11686, -178.87 DPS) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 35.5 stamina points (1897.06 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -295.44 DPS) [vendor]; Windscale Sarong (10842, -310.27 DPS) [world]; Stone Guard's Restored Leather Leggings (220883, -384.08 DPS) [vendor] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -43.12 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (801.04 DPS) | yes | Darkspear Signet (272069, -53.40 DPS) [vendor]; Underworld Band (1980, -160.21 DPS) [world_drop]; Nogg's Gold Ring (9588, -200.88 DPS) [quest] |
| finger2 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.4 stamina points (771.33 DPS) | yes | Darkspear Signet (272069, -23.69 DPS) [vendor]; Underworld Band (1980, -130.50 DPS) [world_drop]; Nogg's Gold Ring (9588, -171.17 DPS) [quest] |
| trinket1 | Defiler's Talisman (21115) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -53.40 DPS) [quest]; Rune of Perfection (21565, -106.81 DPS) [rep]; Rune of Duty (21567, -106.81 DPS) [rep] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -53.40 DPS) [quest]; Rune of Perfection (21565, -106.81 DPS) [rep] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -187.08 DPS) [dungeon]; Warden Staff (943, -267.83 DPS) [world_drop]; Blanchard's Stout (13046, -279.39 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Sergeant's Cloak; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal

No-known-source sample (15 of 689, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 182.6. Weights run: 7.3s. Verify run: 40.1s. 1759 eligible items had no known source.

3 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.048, armor=0.139 ± 0.022, defense=1.304 ± 0.274 per rating point (1 rating = 1%, 1.304 per %), dodge=0.491 ± 0.052 per rating point (12 rating = 1%, 5.887 per %), strength=0.086 ± 0.000, agility=0.419 ± 0.034, attack_power=0.043 ± 0.000, feral_attack_power=0.043 ± 0.000, hit=not significant (0.202 ± 0.053) per rating point (10 rating = 1%, 2.018 per %), crit=not significant (0.083 ± 0.022) per rating point (14 rating = 1%, 1.166 per %), expertise=6.399 ± 0.585

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 63.8 stamina points (4684.79 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, -376.13 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 20.0 stamina points (1471.15 DPS) | yes | Scout's Medallion (19534, -275.78 DPS) [rep]; Frostwolf Legionnaire's Pendant (19095, -312.89 DPS) [rep]; Senior Sergeant's Insignia (16335, -317.27 DPS, sim-verified) [vendor] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 46.4 stamina points (3408.24 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -1726.49 DPS, sim-verified) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phantasmal Cloak (18689, -519.48 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -552.26 DPS) [dungeon]; Fel Cape (279269, -610.27 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 57.5 stamina points (4219.41 DPS) | yes | Breastplate of Bloodthirst (12757, +0.00 DPS) [quest]; Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malefic Bracers (18700, +0.00 DPS) [dungeon]; Yeti Hide Bracers (19113, -14.77 DPS) [world]; First Sergeant's Dragonhide Armguards (18434, -84.61 DPS) [vendor] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+694.3 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 36.9 stamina points (2707.23 DPS) | yes | Shifter's Belt (272396, -10.93 DPS) [vendor]; Elderwild Waistcord (279252, -31.29 DPS) [crafted]; Frostbite Girdle (14502, -112.00 DPS) [dungeon] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 64.0 stamina points (4696.24 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Warstrife Leggings (11821, -1217.13 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Thrall's Resolve (12544) | The Princess Saved? [quest] | 28.2 stamina points (2071.62 DPS) | yes | Myrmidon's Signet (2246, -545.02 DPS) [world_drop]; Band of Flesh (13373, -766.66 DPS) [dungeon]; Naglering (11669, -2722.99 DPS, sim-verified) [dungeon] |
| finger2 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Myrmidon's Signet (2246, -119.64 DPS) [world_drop]; Band of Flesh (13373, -341.28 DPS) [dungeon] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (+1271.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Adaptive Combat Assistant (272437, -304.52 DPS) [vendor]; Frostwolf Insignia Rank 6 (17909, -379.79 DPS) [quest]; Darkspear Voodoo Seal (272061, -510.02 DPS) [vendor] |
| trinket2 | Defiler's Talisman (20072) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Insignia Rank 6 (17909, -16.63 DPS) [quest]; Relentless Raider's Seal (272060, -146.86 DPS) [vendor]; Darkspear Voodoo Seal (272061, -146.86 DPS) [vendor] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -788.23 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Hivethrasher's Girdle; legs: Outrider's Leather Pants; feet: Dawn Treaders; finger1: Thrall's Resolve; finger2: Ring of Awareness; trinket1: Defender's Grip Stabilizer; trinket2: Defiler's Talisman; main_hand: Headmaster's Charge

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 443.7. Weights run: 8.0s. Verify run: 49.6s. 1759 eligible items had no known source.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.009, armor=0.196 ± 0.022, defense=1.468 ± 0.288 per rating point (1 rating = 1%, 1.468 per %), dodge=0.553 ± 0.057 per rating point (12 rating = 1%, 6.638 per %), strength=0.099 ± 0.000, agility=0.569 ± 0.037, attack_power=0.045 ± 0.000, feral_attack_power=0.045 ± 0.000, hit=0.236 ± 0.052 per rating point (10 rating = 1%, 2.360 per %), crit=0.092 ± 0.023 per rating point (14 rating = 1%, 1.294 per %), expertise=7.741 ± 0.582

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 76.3 stamina points (8439.88 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 21.7 stamina points (2398.50 DPS) | yes | Scout's Medallion (19534, -348.17 DPS) [rep]; Senior Sergeant's Insignia (16335, -517.86 DPS) [vendor]; Frostwolf Legionnaire's Pendant (19095, -649.74 DPS) [rep] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | sim-verified (+2118.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -2117.99 DPS, sim-verified) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phantasmal Cloak (18689, -1634.05 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -1644.95 DPS) [dungeon]; Fel Cape (279269, -1790.35 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 69.3 stamina points (7670.54 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Breastplate of Bloodthirst (12757, -395.18 DPS, sim-verified) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 36.8 stamina points (4068.28 DPS) | yes | Malefic Bracers (18700, -294.39 DPS) [dungeon]; First Sergeant's Dragonhide Armguards (18434, -515.76 DPS) [vendor]; Yeti Hide Bracers (19113, -616.43 DPS, sim-verified) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+967.7 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 46.4 stamina points (5138.44 DPS) | yes | Cloudrunner Girdle (13252, -24.59 DPS) [dungeon]; Frostbite Girdle (14502, -352.53 DPS) [dungeon]; Hivethrasher's Girdle (275614, -355.46 DPS) [crafted] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 78.2 stamina points (8648.63 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Warstrife Leggings (11821, -2748.38 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Thrall's Resolve (12544) | The Princess Saved? [quest] | 36.8 stamina points (4074.98 DPS) | yes | Ring of Awareness (272409, -1419.63 DPS) [vendor]; Myrmidon's Signet (2246, -1644.55 DPS) [world_drop]; Cutthroat's Signet (272408, -2047.76 DPS) [vendor] |
| finger2 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, -573.52 DPS) [world_drop]; Ring of Awareness (272409, -782.20 DPS, sim-verified) [vendor]; Cutthroat's Signet (272408, -976.73 DPS) [vendor] |
| trinket1 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Talisman (20072, -141.06 DPS) [rep]; Force of Will (11810, -331.69 DPS) [dungeon] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Adaptive Combat Assistant (272437, -398.65 DPS) [vendor]; Defiler's Talisman (20072, -783.85 DPS) [rep]; Force of Will (11810, -974.47 DPS) [dungeon] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -1187.12 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Outrider's Leather Pants; feet: Dawn Treaders; finger1: Thrall's Resolve; finger2: Naglering; trinket1: Frostwolf Insignia Rank 6; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

