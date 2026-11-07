# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 22.7. Weights run: 9.1s. Verify run: 8.2s. 228 eligible items had no known source.

Race sweep: high-order-skyborne 90.0 ± 0.1, night-elf 88.9 ± 0.1 (threat_per_second); high-order-skyborne beats the guide race night-elf by 1.1 (1.19%) ± 0.1, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.098 ± 0.001, defense=0.275 ± 0.027 per rating point (1 rating = 1%, 0.275 per %), dodge=0.109 ± 0.008 per rating point (12 rating = 1%, 1.312 per %), strength=0.097 ± 0.000, agility=0.164 ± 0.005, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.043 ± 0.004 per rating point (10 rating = 1%, 0.430 per %), crit=0.017 ± 0.001 per rating point (14 rating = 1%, 0.244 per %), expertise=1.116 ± 0.051

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 19.2 stamina points (349.05 DPS) | yes | Brawler's Leather Hood (252504, -8.07 DPS) [crafted]; Defender's Leather Hood (252447, -17.90 DPS) [crafted]; Trapper's Leather Hood (252505, -31.96 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.7 stamina points (103.68 DPS) | yes | Erudite's Amulet (277204, -43.05 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -49.20 DPS) [quest]; Sentinel's Medallion (20444, -49.44 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 10.3 stamina points (186.74 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -50.57 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -65.97 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -4.43 DPS) [world_drop]; Firebane Cloak (12979, -19.94 DPS) [world_drop]; Caretaker's Cape (20428, -36.32 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 19.6 stamina points (356.26 DPS) | yes | Gloomshroud Armor (1489, -33.08 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -68.58 DPS) [quest]; Defender's Leather Armor (252434, -92.65 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.6 stamina points (156.76 DPS) | yes | Drakewing Bands (12999, -26.42 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -42.53 DPS) [vendor]; Sanguine Cuffs (14375, -53.82 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.9 stamina points (179.08 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.78 DPS) [crafted]; Defender's Leather Gloves (252496, -4.92 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 12.1 stamina points (220.09 DPS) | yes | Belt of the Fang (10412, -30.67 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -49.93 DPS) [crafted]; Guardsman Belt (3429, -51.71 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 15.6 stamina points (282.92 DPS) | yes | Brawler's Leather Pants (252500, -26.66 DPS) [crafted]; Leggings of the Fang (10410, -35.55 DPS) [dungeon]; Smelting Pants (5199, -100.32 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.5 stamina points (282.09 DPS) | yes | Footpads of the Fang (10411, -47.58 DPS, sim-verified) [dungeon]; Surfer Shoes (276274, -48.80 DPS) [vendor]; Trapper's Leather Boots (252440, -67.61 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (90.80 DPS) | yes | Slain Baron's Signet (279867, -31.33 DPS) [quest]; Ring of the Moon (12052, -34.56 DPS) [world_drop]; Protector's Band (20439, -35.50 DPS) [rep] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 4.0 stamina points (72.64 DPS) | yes | Slain Baron's Signet (279867, -13.17 DPS) [quest]; Ring of the Moon (12052, -16.40 DPS) [world_drop]; Protector's Band (20439, -17.34 DPS) [rep] |
| trinket1 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.8 stamina points (377.12 DPS) | yes | Rhahk'Zor's Hammer (5187, -5.49 DPS, sim-verified) [dungeon]; Twisted Chanter's Staff (890, -16.62 DPS) [world_drop]; Staff of the Blessed Seer (2271, -20.90 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Sustaining Ring; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 228, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots

### Band 30 (night-elf, 0000000000000000-51100332020102100000-0000000000000000)

Set DPS (verified): 64.8. Weights run: 7.3s. Verify run: 12.1s. 400 eligible items had no known source.

Race sweep: high-order-skyborne 250.8 ± 0.2, night-elf 250.0 ± 0.2 (threat_per_second); high-order-skyborne beats the guide race night-elf by 0.8 (0.34%) ± 0.3, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.104 ± 0.005, defense=0.404 ± 0.085 per rating point (1 rating = 1%, 0.404 per %), dodge=0.164 ± 0.017 per rating point (12 rating = 1%, 1.965 per %), strength=0.121 ± 0.000, agility=0.215 ± 0.011, attack_power=0.061 ± 0.000, feral_attack_power=0.061 ± 0.000, hit=0.070 ± 0.013 per rating point (10 rating = 1%, 0.702 per %), crit=0.034 ± 0.006 per rating point (14 rating = 1%, 0.471 per %), expertise=2.061 ± 0.144

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 23.9 stamina points (528.55 DPS) | yes | Defender's Leather Helm (252455, -15.45 DPS) [crafted]; Trapper's Leather Helm (252513, -47.72 DPS) [crafted]; Stormrider's Leather Helm (252514, -47.72 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.5 stamina points (210.16 DPS) | yes | Master Sergeant's Insignia (18442, -10.76 DPS) [vendor]; Seal of the Expedition (271920, -32.91 DPS) [rep]; Sentinel's Medallion (19541, -61.21 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 19.7 stamina points (435.41 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -65.07 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -72.47 DPS) [crafted]; Barbaric Shoulders (5964, -90.73 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mourning Shawl (6751, -31.34 DPS) [quest]; Enduring Cape (14763, -38.57 DPS) [world_drop]; Tigerstrike Mantle (13108, -95.46 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | 23.5 stamina points (519.61 DPS) | yes | Raptor Hide Harness (4455, -6.64 DPS) [crafted]; Green Whelp Armor (7375, -41.77 DPS) [crafted]; Blackened Defias Armor (10399, -68.70 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.2 stamina points (270.66 DPS) | yes | Cultist's Armguards (270032, -1.13 DPS) [quest]; Black Wolf Bracers (3230, -11.49 DPS) [dungeon]; Glowing Leather Bracers (2017, -58.52 DPS) [world] |
| hands | Brawler Gloves (720) | World drop [world_drop] | 14.4 stamina points (318.91 DPS) | yes | Wolfclaw Gloves (1978, -3.89 DPS) [dungeon]; Gloves of Insight (9698, -23.81 DPS) [quest]; Emblazoned Gloves (6397, -31.28 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 17.5 stamina points (388.77 DPS) | yes | Prowler's Leather Belt (252459, -58.40 DPS) [crafted]; Stalker's Leather Belt (252521, -68.29 DPS) [crafted]; Skulker's Leather Belt (252520, -147.70 DPS, sim-verified) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 20.7 stamina points (458.20 DPS) | yes | Barbaric Leggings (5963, -51.20 DPS) [crafted]; Petrolspill Leggings (9509, -51.89 DPS) [dungeon]; Brawler's Leather Legguards (252516, -164.76 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.9 stamina points (440.31 DPS) | yes | Harbinger Boots (7754, -28.25 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -91.81 DPS) [quest]; Emblazoned Boots (4051, -124.01 DPS) [world_drop] |
| finger1 | Theramore Signet (271911) | Theramore Expeditionary Force [rep] | 8.5 stamina points (189.35 DPS) | yes | Plains Ring (2039, -12.10 DPS) [dungeon]; Defias Renegade Ring (1076, -37.33 DPS) [dungeon]; Protector's Band (19517, -55.96 DPS) [rep] |
| finger2 | Militant's Band (271916) (or Plains Ring (2039)) | Theramore Expeditionary Force [rep] | 8.0 stamina points (177.25 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Defias Renegade Ring (1076, -25.23 DPS) [dungeon]; Protector's Band (19517, -43.86 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (+142.6 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -33.20 DPS, sim-verified) [world]; Slaghammer (1976, -56.46 DPS) [dungeon]; Lorekeeper's Staff (212580, -69.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Brawler Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Theramore Signet; finger2: Militant's Band; trinket1: Talisman of Arathor; trinket2: Rune of Duty; main_hand: Viscous Hammer

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 99.1. Weights run: 7.4s. Verify run: 6.9s. 526 eligible items had no known source.

Race sweep: high-order-skyborne 380.6 ± 0.4, night-elf 380.6 ± 0.4 (threat_per_second); high-order-skyborne beats the guide race night-elf by 0.0 (0.01%) ± 0.6, inside error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.157 ± 0.008, defense=not significant (0.406 ± 0.109) per rating point (1 rating = 1%, 0.406 per %), dodge=0.186 ± 0.023 per rating point (12 rating = 1%, 2.235 per %), strength=0.099 ± 0.000, agility=0.235 ± 0.015, attack_power=0.050 ± 0.000, feral_attack_power=0.050 ± 0.000, hit=not significant (0.079 ± 0.020) per rating point (10 rating = 1%, 0.787 per %), crit=0.050 ± 0.008 per rating point (14 rating = 1%, 0.698 per %), expertise=2.977 ± 0.214

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 30.9 stamina points (1001.89 DPS) | yes | Warden's Wizard Hat (14604, -74.06 DPS, sim-verified) [world_drop]; Brawler's Leather Helm (252512, -75.80 DPS) [crafted]; Nightscape Headband (8176, -89.80 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shriveled Heart (9243, +0.00 DPS) [dungeon]; Seal of the Expedition (271919, -12.86 DPS) [rep]; Sentinel's Medallion (19540, -58.92 DPS) [rep] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.8 stamina points (1032.67 DPS) | yes | Flintrock Shoulders (7755, -161.01 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -258.36 DPS) [dungeon]; Nightscape Shoulders (8192, -348.38 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -20.26 DPS) [dungeon]; Heraldic Cloak (8120, -22.26 DPS) [dungeon]; Well Oiled Cloak (12254, -25.36 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.5 stamina points (1151.19 DPS) | yes | Warden's Wraps (14601, -90.90 DPS, sim-verified) [world_drop]; Robes of the Lich (10762, -186.21 DPS) [dungeon]; Spirewind Fetter (9406, -200.92 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.6 stamina points (537.85 DPS) | yes | Branded Leather Bracers (19508, -30.68 DPS) [dungeon]; Barbaric Bracers (18948, -60.30 DPS) [crafted]; Cultist's Armguards (270032, -70.37 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -21.22 DPS) [crafted]; Prowler's Leather Gloves (252524, -25.62 DPS) [crafted]; Stalker's Leather Gloves (252526, -367.58 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.2 stamina points (751.36 DPS) | yes | Warden's Leather Belt (252460, -73.26 DPS) [crafted]; Kolkar Hunter's Belt (6788, -77.45 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -88.66 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 30.7 stamina points (995.38 DPS) | yes | Panther Hunter Leggings (4108, -73.04 DPS) [quest]; Warden's Woolies (14605, -81.04 DPS) [world_drop]; Imperial Leather Pants (4062, -90.12 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 27.8 stamina points (901.63 DPS) | yes | Skulker's Leather Shoes (252531, -99.35 DPS) [crafted]; Prowler's Leather Shoes (252465, -116.94 DPS) [crafted]; Worn Running Boots (9398, -308.82 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (389.40 DPS) | yes | Talvash's Gold Ring (9538, -25.33 DPS) [quest]; Militant's Band (271915, -32.45 DPS) [rep]; Dragonclaw Ring (10710, -64.90 DPS) [quest] |
| finger2 | Theramore Signet (271910) | Theramore Expeditionary Force [rep] | 11.6 stamina points (376.24 DPS) | yes | Talvash's Gold Ring (9538, -12.18 DPS) [quest]; Militant's Band (271915, -19.30 DPS) [rep]; Dragonclaw Ring (10710, -51.74 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (+267.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -32.45 DPS) [world]; Enriched Seal (271905, -32.45 DPS) [rep] |
| trinket2 | Talisman of Arathor (21118) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -32.45 DPS) [world]; Volunteer's Lucky Seal (271904, -32.45 DPS) [rep]; Enriched Seal (271905, -32.45 DPS) [rep] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -135.65 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -415.86 DPS) [dungeon]; Huge Stone Club (1523, -501.31 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Theramore Signet; trinket1: Rune of Duty; trinket2: Talisman of Arathor; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 526, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 139.2. Weights run: 7.4s. Verify run: 17.3s. 690 eligible items had no known source.

Race sweep: night-elf 539.4 ± 0.5, high-order-skyborne 538.3 ± 0.5 (threat_per_second); the guide race night-elf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.032, armor=0.130 ± 0.010, defense=not significant (0.469 ± 0.118) per rating point (1 rating = 1%, 0.469 per %), dodge=0.253 ± 0.029 per rating point (12 rating = 1%, 3.037 per %), strength=0.076 ± 0.000, agility=0.271 ± 0.018, attack_power=0.038 ± 0.000, feral_attack_power=0.038 ± 0.000, hit=0.102 ± 0.024 per rating point (10 rating = 1%, 1.017 per %), crit=0.058 ± 0.010 per rating point (14 rating = 1%, 0.818 per %), expertise=3.454 ± 0.276

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 40.2 stamina points (2112.98 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -220.72 DPS) [vendor]; Embrace of the Lycan (9479, -374.17 DPS) [dungeon]; Scorpashi Skullcap (14658, -497.29 DPS) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 4 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (735.10 DPS) | yes | Shriveled Heart (9243, -52.51 DPS) [dungeon]; Lifeblood Amulet (9641, -52.51 DPS) [dungeon]; Seal of the Expedition (271918, -52.51 DPS) [rep] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fleshhide Shoulders (10774, +0.00 DPS) [dungeon]; Phytoskin Spaulders (17749, -8.49 DPS) [dungeon]; Knight-Lieutenant's Leather Shoulders (220852, -24.05 DPS) [vendor] |
| back | Sergeant's Cape (18441) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Graverot Cape (11677, +0.00 DPS) [dungeon]; Nightfall Drape (12465, +0.00 DPS) [dungeon]; Emberscale Cape (11311, -13.62 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, -544.51 DPS) [dungeon]; Jinxed Hoodoo Skin (9473, -723.33 DPS) [dungeon]; Heraldic Breastplate (8119, -808.43 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.9 stamina points (1256.34 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -85.50 DPS) [vendor]; Warden's Leather Bracers (252541, -234.74 DPS) [crafted]; Serpentskin Bracers (8257, -241.12 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -191.29 DPS) [crafted]; Wicked Leather Gauntlets (15083, -239.02 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -242.68 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 26.4 stamina points (1388.70 DPS) | yes | Skulker's Leather Waistguard (252474, -10.54 DPS) [crafted]; Prowler's Leather Waistguard (252473, -31.09 DPS) [crafted]; Girdle of Beastial Fury (11686, -166.35 DPS) [dungeon] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 34.9 stamina points (1832.07 DPS) | yes | Windscale Sarong (10842, -281.30 DPS) [world]; Knight's Crackling Leather Leggings (220864, -298.27 DPS) [vendor]; Basilisk Hide Pants (1718, -343.41 DPS) [world_drop] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -24.79 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (787.60 DPS) | yes | Militant's Band (271914, -52.51 DPS) [rep]; Underworld Band (1980, -157.52 DPS) [world_drop]; Talvash's Gold Ring (9538, -188.62 DPS) [quest] |
| finger2 | Theramore Signet (271909) | Theramore Expeditionary Force [rep] | 14.6 stamina points (764.88 DPS) | yes | Militant's Band (271914, -29.78 DPS) [rep]; Underworld Band (1980, -134.79 DPS) [world_drop]; Talvash's Gold Ring (9538, -165.89 DPS) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Rune of Perfection (21565, -105.01 DPS) [rep] |
| trinket2 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Rune of Perfection (21565, -52.51 DPS) [rep] |
| main_hand | Blanchard's Stout (13046) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warden Staff (943, +0.00 DPS) [world_drop]; Ironshod Bludgeon (9408, +0.00 DPS) [dungeon]; Radiant Staff (249453, -177.78 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Penance Spaulders; back: Sergeant's Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Theramore Signet; trinket1: Enriched Seal; trinket2: Guardian Talisman; main_hand: Blanchard's Stout

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 203.6. Weights run: 8.1s. Verify run: 52.0s. 1755 eligible items had no known source.

Race sweep: night-elf 794.5 ± 0.9, high-order-skyborne 793.6 ± 0.9 (threat_per_second); the guide race night-elf is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.152 ± 0.015, defense=1.456 ± 0.283 per rating point (1 rating = 1%, 1.456 per %), dodge=0.762 ± 0.047 per rating point (12 rating = 1%, 9.148 per %), strength=0.116 ± 0.000, agility=0.622 ± 0.030, attack_power=0.058 ± 0.000, feral_attack_power=0.058 ± 0.000, hit=not significant (0.138 ± 0.038) per rating point (10 rating = 1%, 1.375 per %), crit=0.140 ± 0.018 per rating point (14 rating = 1%, 1.959 per %), expertise=4.659 ± 0.346

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 71.4 stamina points (10136.28 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, -567.27 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 21.6 stamina points (3062.10 DPS) | yes | Sentinel's Medallion (19538, -317.15 DPS) [rep]; Talisman of Evasion (13177, -614.78 DPS) [dungeon]; Master Sergeant's Insignia (18443, -647.95 DPS) [vendor] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | sim-verified (+1441.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -1441.47 DPS, sim-verified) [rep] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phantasmal Cloak (18689, -597.88 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -1142.24 DPS) [dungeon]; Fel Cape (279269, -1330.84 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 63.5 stamina points (9023.75 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Breastplate of Bloodthirst (12757, -325.63 DPS, sim-verified) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 34.2 stamina points (4856.06 DPS) | yes | Yeti Hide Bracers (19113, -380.08 DPS) [world]; Malefic Bracers (18700, -407.80 DPS) [dungeon]; Sergeant Major's Dragonhide Armsplints (18454, -775.87 DPS) [vendor] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+885.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -219.31 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, -219.31 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 43.4 stamina points (6162.55 DPS) | yes | Frostbite Girdle (14502, -446.69 DPS, sim-verified) [dungeon]; Hivethrasher's Girdle (275614, -549.87 DPS) [crafted]; Cloudrunner Girdle (13252, -604.44 DPS) [dungeon] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 90.3 stamina points (12828.51 DPS) | yes | Marshal's Dragonhide Leggings (231691, -286.10 DPS) [vendor]; Sentinel's Lizardhide Pants (237817, -1131.38 DPS, sim-verified) [vendor]; Marshal's Dragonhide Legguards (231703, -1419.59 DPS) [pvp] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Flesh (13373, -777.13 DPS) [dungeon]; Innervating Band (18701, -810.07 DPS) [dungeon] |
| finger2 | Myrmidon's Signet (2246) | World drop [world_drop] | sim-verified (+165.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Naglering (11669, +0.00 DPS) [dungeon]; Band of Flesh (13373, -561.15 DPS) [dungeon]; Innervating Band (18701, -594.09 DPS) [dungeon] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Insignia Rank 6 (17904, -383.58 DPS, sim-verified) [quest]; Force of Will (11810, -1240.82 DPS) [dungeon]; Enriched Seal (271903, -1268.36 DPS) [rep] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Force of Will (11810, -256.48 DPS) [dungeon]; Enriched Seal (271903, -284.02 DPS) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -852.29 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Dawn Treaders; finger1: Ring of Awareness; finger2: Myrmidon's Signet; trinket1: Defender's Grip Stabilizer; trinket2: Talisman of Arathor; main_hand: Headmaster's Charge

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 487.6. Weights run: 9.0s. Verify run: 53.7s. 1755 eligible items had no known source.

Race sweep: night-elf 1838.9 ± 2.1, high-order-skyborne 1830.3 ± 2.2 (threat_per_second); the guide race night-elf is best.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.027, armor=0.178 ± 0.017, defense=1.401 ± 0.295 per rating point (1 rating = 1%, 1.401 per %), dodge=0.552 ± 0.051 per rating point (12 rating = 1%, 6.627 per %), strength=0.102 ± 0.000, agility=0.530 ± 0.033, attack_power=0.046 ± 0.000, feral_attack_power=0.046 ± 0.000, hit=not significant (0.149 ± 0.040) per rating point (10 rating = 1%, 1.491 per %), crit=0.104 ± 0.018 per rating point (14 rating = 1%, 1.455 per %), expertise=6.931 ± 0.450

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.9 stamina points (8423.97 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 21.0 stamina points (2429.39 DPS) | yes | Sentinel's Medallion (19538, -353.71 DPS) [rep]; Master Sergeant's Insignia (18443, -463.71 DPS) [vendor]; Stormpike Soldier's Pendant (19097, -598.45 DPS) [rep] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 54.3 stamina points (6282.27 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phantasmal Cloak (18689, -847.78 DPS) [dungeon]; The Emperor's New Cape (11930, -1446.51 DPS) [dungeon]; Fel Cape (279269, -1562.93 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 65.8 stamina points (7613.70 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Silverwing Sentinels [rep] | 34.5 stamina points (3993.08 DPS) | yes | Yeti Hide Bracers (19113, -213.26 DPS) [world]; Malefic Bracers (18700, -229.00 DPS) [dungeon]; Sergeant Major's Dragonhide Armsplints (18454, -439.29 DPS) [vendor] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 43.7 stamina points (5053.06 DPS) | yes | Cloudrunner Girdle (13252, -152.94 DPS) [dungeon]; Hivethrasher's Girdle (275614, -277.72 DPS) [crafted]; Frostbite Girdle (14502, -319.04 DPS) [dungeon] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 92.1 stamina points (10646.50 DPS) | yes | Marshal's Dragonhide Leggings (231691, -111.06 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -797.63 DPS) [pvp]; Marshal's Dragonhide Pants (231697, -1438.77 DPS) [vendor] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.9 stamina points (2997.51 DPS) | yes | Myrmidon's Signet (2246, -484.82 DPS) [world_drop]; Band of Flesh (13373, -892.81 DPS) [dungeon]; Innervating Band (18701, -1037.47 DPS) [dungeon] |
| finger2 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, -195.95 DPS) [world_drop]; Band of Flesh (13373, -603.94 DPS) [dungeon]; Innervating Band (18701, -748.60 DPS) [dungeon] |
| trinket1 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (20071, -144.88 DPS) [rep]; Enriched Seal (271903, -376.14 DPS) [rep] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Adaptive Combat Assistant (272437, -503.10 DPS) [vendor]; Talisman of Arathor (20071, -718.45 DPS) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; wrist: Forest Stalker's Bracers; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Dawn Treaders; finger1: Naglering; finger2: Ring of Awareness; trinket1: Stormpike Insignia Rank 6; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1755, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 23.2. Weights run: 9.1s. Verify run: 7.6s. 224 eligible items had no known source.

Race sweep: windshaper-skyborne 90.9 ± 0.1, tauren 90.7 ± 0.1 (threat_per_second); windshaper-skyborne beats the guide race tauren by 0.3 (0.31%) ± 0.1, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.098 ± 0.001, defense=0.275 ± 0.027 per rating point (1 rating = 1%, 0.275 per %), dodge=0.109 ± 0.008 per rating point (12 rating = 1%, 1.312 per %), strength=0.097 ± 0.000, agility=0.164 ± 0.005, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.043 ± 0.004 per rating point (10 rating = 1%, 0.430 per %), crit=0.017 ± 0.001 per rating point (14 rating = 1%, 0.244 per %), expertise=1.116 ± 0.051

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 19.2 stamina points (349.05 DPS) | yes | Brawler's Leather Hood (252504, -14.91 DPS, sim-verified) [crafted]; Defender's Leather Hood (252447, -17.90 DPS) [crafted]; Trapper's Leather Hood (252505, -31.96 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (60.45 DPS) | yes | Scholarly Pendant (277203, -5.97 DPS) [quest]; Scout's Medallion (20442, -6.21 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 10.3 stamina points (186.74 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -50.57 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -65.97 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Firebane Cloak (12979, -19.94 DPS) [world_drop]; Sentry Cloak (2059, -21.22 DPS, sim-verified) [world_drop]; Battle Healer's Cloak (20427, -36.32 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 19.6 stamina points (356.26 DPS) | yes | Gloomshroud Armor (1489, -35.06 DPS, sim-verified) [dungeon]; Defender's Leather Armor (252434, -92.65 DPS) [crafted]; Trapper's Leather Armor (252491, -93.01 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.6 stamina points (156.76 DPS) | yes | Drakewing Bands (12999, -29.19 DPS) [world_drop]; Spare Part Bindings (279875, -34.47 DPS, sim-verified) [quest]; Wolf Bracers (4794, -42.53 DPS) [vendor] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.9 stamina points (179.08 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.78 DPS) [crafted]; Defender's Leather Gloves (252496, -4.92 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 12.1 stamina points (220.09 DPS) | yes | Belt of the Fang (10412, -37.41 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -49.93 DPS) [crafted]; Guardsman Belt (3429, -51.71 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 15.6 stamina points (282.92 DPS) | yes | Brawler's Leather Pants (252500, -26.66 DPS) [crafted]; Leggings of the Fang (10410, -35.55 DPS) [dungeon]; Smelting Pants (5199, -105.02 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.5 stamina points (282.09 DPS) | yes | Surfer Shoes (276274, -48.80 DPS) [vendor]; Footpads of the Fang (10411, -53.79 DPS, sim-verified) [dungeon]; Trapper's Leather Boots (252440, -67.61 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (90.80 DPS) | yes | Slain Baron's Signet (279867, -31.33 DPS) [quest]; Ring of the Moon (12052, -34.56 DPS) [world_drop]; Legionnaire's Band (20429, -35.50 DPS) [rep] |
| finger2 | Ring of Scorn (3235) | A Husband's Revenge [quest] | 4.0 stamina points (72.64 DPS) | yes | Slain Baron's Signet (279867, -15.06 DPS, sim-verified) [quest]; Ring of the Moon (12052, -16.40 DPS) [world_drop]; Legionnaire's Band (20429, -17.34 DPS) [rep] |
| trinket1 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.8 stamina points (377.12 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.96 DPS) [dungeon]; Twisted Chanter's Staff (890, -16.62 DPS) [world_drop]; Staff of the Blessed Seer (2271, -20.90 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Ring of Scorn; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 224, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-51100332020102100000-0000000000000000)

Set DPS (verified): 64.7. Weights run: 7.3s. Verify run: 11.5s. 401 eligible items had no known source.

Race sweep: tauren 249.2 ± 0.2, windshaper-skyborne 249.1 ± 0.2 (threat_per_second); the guide race tauren is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.104 ± 0.005, defense=0.404 ± 0.085 per rating point (1 rating = 1%, 0.404 per %), dodge=0.164 ± 0.017 per rating point (12 rating = 1%, 1.965 per %), strength=0.121 ± 0.000, agility=0.215 ± 0.011, attack_power=0.061 ± 0.000, feral_attack_power=0.061 ± 0.000, hit=0.070 ± 0.013 per rating point (10 rating = 1%, 0.702 per %), crit=0.034 ± 0.006 per rating point (14 rating = 1%, 0.471 per %), expertise=2.061 ± 0.144

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 23.9 stamina points (528.55 DPS) | yes | Defender's Leather Helm (252455, -35.03 DPS, sim-verified) [crafted]; Trapper's Leather Helm (252513, -47.72 DPS) [crafted]; Stormrider's Leather Helm (252514, -47.72 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.5 stamina points (210.16 DPS) | yes | Senior Sergeant's Insignia (15200, -10.76 DPS) [vendor]; Darkspear Warding Pendant (272075, -32.91 DPS) [vendor]; Scout's Medallion (19537, -61.21 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 19.7 stamina points (435.41 DPS) | yes | Azure Gustwoven Spaulders (277051, -72.47 DPS) [crafted]; Cloudy Gustwoven Spaulders (277043, -89.02 DPS, sim-verified) [crafted]; Barbaric Shoulders (5964, -90.73 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enduring Cape (14763, -38.57 DPS) [world_drop]; Grimsteel Cape (4643, -41.65 DPS) [quest]; Tigerstrike Mantle (13108, -115.40 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | 23.5 stamina points (519.61 DPS) | yes | Raptor Hide Harness (4455, -6.64 DPS) [crafted]; Green Whelp Armor (7375, -41.77 DPS) [crafted]; Blackened Defias Armor (10399, -68.70 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.2 stamina points (270.66 DPS) | yes | Cultist's Armguards (270032, -1.13 DPS) [quest]; Black Wolf Bracers (3230, -11.49 DPS) [dungeon]; Glowing Leather Bracers (2017, -58.52 DPS) [world] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 15.1 stamina points (335.39 DPS) | yes | Wolfclaw Gloves (1978, -20.37 DPS) [dungeon]; Brawler Gloves (720, -33.66 DPS, sim-verified) [world_drop]; Emblazoned Gloves (6397, -47.77 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 17.5 stamina points (388.77 DPS) | yes | Skulker's Leather Belt (252520, -52.15 DPS) [crafted]; Prowler's Leather Belt (252459, -58.40 DPS) [crafted]; Larcenist's Belt (279846, -160.43 DPS, sim-verified) [quest] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 20.7 stamina points (458.20 DPS) | yes | Barbaric Leggings (5963, -51.20 DPS) [crafted]; Petrolspill Leggings (9509, -51.89 DPS) [dungeon]; Brawler's Leather Legguards (252516, -210.62 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.9 stamina points (440.31 DPS) | yes | Harbinger Boots (7754, -49.35 DPS, sim-verified) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -91.81 DPS) [quest]; Gravewalker Boots (251963, -98.40 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.5 stamina points (189.35 DPS) | yes | Plains Ring (2039, -12.10 DPS) [dungeon]; Darkspear Signet (272071, -12.10 DPS) [vendor]; Defias Renegade Ring (1076, -37.33 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.4 stamina points (185.32 DPS) | yes | Plains Ring (2039, -8.07 DPS) [dungeon]; Darkspear Signet (272071, -8.07 DPS) [vendor]; Defias Renegade Ring (1076, -33.29 DPS) [dungeon] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | sim-verified (+148.3 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -34.07 DPS, sim-verified) [world]; Slaghammer (1976, -56.46 DPS) [dungeon]; Advisor's Gnarled Staff (212584, -69.00 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Defiler's Talisman; trinket2: Rune of Duty; main_hand: Viscous Hammer

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 97.7. Weights run: 7.4s. Verify run: 7.2s. 525 eligible items had no known source.

Race sweep: windshaper-skyborne 375.8 ± 0.4, tauren 375.0 ± 0.4 (threat_per_second); windshaper-skyborne beats the guide race tauren by 0.8 (0.22%) ± 0.5, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.157 ± 0.008, defense=not significant (0.406 ± 0.109) per rating point (1 rating = 1%, 0.406 per %), dodge=0.186 ± 0.023 per rating point (12 rating = 1%, 2.235 per %), strength=0.099 ± 0.000, agility=0.235 ± 0.015, attack_power=0.050 ± 0.000, feral_attack_power=0.050 ± 0.000, hit=not significant (0.079 ± 0.020) per rating point (10 rating = 1%, 0.787 per %), crit=0.050 ± 0.008 per rating point (14 rating = 1%, 0.698 per %), expertise=2.977 ± 0.214

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 30.9 stamina points (1001.89 DPS) | yes | Warden's Wizard Hat (14604, -41.41 DPS) [world_drop]; Brawler's Leather Helm (252512, -75.80 DPS) [crafted]; Nightscape Headband (8176, -89.80 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shriveled Heart (9243, +0.00 DPS) [dungeon]; Dragon's Blood Necklace (10711, +0.00 DPS) [quest]; Darkspear Warding Pendant (272074, -12.86 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.8 stamina points (1032.67 DPS) | yes | Flintrock Shoulders (7755, -192.80 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -258.36 DPS) [dungeon]; Nightscape Shoulders (8192, -348.38 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heraldic Cloak (8120, -22.26 DPS) [dungeon]; Well Oiled Cloak (12254, -25.36 DPS) [vendor]; Silky Spider Cape (10776, -78.23 DPS, sim-verified) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.5 stamina points (1151.19 DPS) | yes | Warden's Wraps (14601, -58.86 DPS) [world_drop]; Robes of the Lich (10762, -186.21 DPS) [dungeon]; Spirewind Fetter (9406, -200.92 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.6 stamina points (537.85 DPS) | yes | Barbaric Bracers (18948, -60.30 DPS) [crafted]; Cultist's Armguards (270032, -70.37 DPS) [quest]; Branded Leather Bracers (19508, -83.60 DPS, sim-verified) [dungeon] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -21.22 DPS) [crafted]; Prowler's Leather Gloves (252524, -25.62 DPS) [crafted]; Stalker's Leather Gloves (252526, -439.01 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.2 stamina points (751.36 DPS) | yes | Warden's Leather Belt (252460, -73.26 DPS) [crafted]; Kolkar Hunter's Belt (6788, -85.25 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -88.66 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 30.7 stamina points (995.38 DPS) | yes | Panther Hunter Leggings (4108, -73.04 DPS) [quest]; Warden's Woolies (14605, -81.04 DPS) [world_drop]; Imperial Leather Pants (4062, -116.10 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 27.8 stamina points (901.63 DPS) | yes | Skulker's Leather Shoes (252531, -99.35 DPS) [crafted]; Prowler's Leather Shoes (252465, -116.94 DPS) [crafted]; Worn Running Boots (9398, -337.71 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (389.40 DPS) | yes | Nogg's Gold Ring (9588, -25.33 DPS) [quest]; Darkspear Signet (272070, -32.45 DPS) [vendor]; Dragonclaw Ring (10710, -64.90 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.6 stamina points (376.24 DPS) | yes | Nogg's Gold Ring (9588, -12.18 DPS) [quest]; Darkspear Signet (272070, -19.30 DPS) [vendor]; Dragonclaw Ring (10710, -51.74 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (+301.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -32.45 DPS) [world]; Darkspear Voodoo Seal (272059, -32.45 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21116) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -32.45 DPS) [world]; Darkspear Voodoo Seal (272059, -32.45 DPS) [vendor]; Relentless Raider's Seal (272062, -32.45 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -114.04 DPS, sim-verified) [dungeon]; Mograine's Might (7723, -260.94 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -303.23 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Insurgent's Band; trinket1: Rune of Duty; trinket2: Defiler's Talisman; main_hand: Cragwood Maul

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 127.6. Weights run: 7.4s. Verify run: 18.7s. 689 eligible items had no known source.

Race sweep: windshaper-skyborne 497.2 ± 0.5, tauren 494.8 ± 0.5 (threat_per_second); windshaper-skyborne beats the guide race tauren by 2.4 (0.49%) ± 0.7, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.032, armor=0.130 ± 0.010, defense=not significant (0.469 ± 0.118) per rating point (1 rating = 1%, 0.469 per %), dodge=0.253 ± 0.029 per rating point (12 rating = 1%, 3.037 per %), strength=0.076 ± 0.000, agility=0.271 ± 0.018, attack_power=0.038 ± 0.000, feral_attack_power=0.038 ± 0.000, hit=0.102 ± 0.024 per rating point (10 rating = 1%, 1.017 per %), crit=0.058 ± 0.010 per rating point (14 rating = 1%, 0.818 per %), expertise=3.454 ± 0.276

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 40.2 stamina points (2112.98 DPS) | yes | Blood Guard's Leather Headband (220851, -220.72 DPS) [vendor]; Embrace of the Lycan (9479, -374.17 DPS) [dungeon]; Scorpashi Skullcap (14658, -497.29 DPS) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 4 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (735.10 DPS) | yes | Shriveled Heart (9243, -52.51 DPS) [dungeon]; Lifeblood Amulet (9641, -52.51 DPS) [dungeon]; Darkspear Warding Pendant (272073, -52.51 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Penance Spaulders (11963, -22.46 DPS) [quest]; Phytoskin Spaulders (17749, -30.96 DPS) [dungeon]; Blood Guard's Leather Shoulders (220853, -46.51 DPS) [vendor] |
| back | Sergeant's Cloak (16341) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Graverot Cape (11677, +0.00 DPS) [dungeon]; Nightfall Drape (12465, +0.00 DPS) [dungeon]; Emberscale Cape (11311, -13.62 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, -544.51 DPS) [dungeon]; Jinxed Hoodoo Skin (9473, -723.33 DPS) [dungeon]; Heraldic Breastplate (8119, -808.43 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.9 stamina points (1256.34 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -85.50 DPS) [vendor]; Warden's Leather Bracers (252541, -234.74 DPS) [crafted]; Serpentskin Bracers (8257, -241.12 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -191.29 DPS) [crafted]; Wicked Leather Gauntlets (15083, -239.02 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -242.68 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 26.4 stamina points (1388.70 DPS) | yes | Skulker's Leather Waistguard (252474, -10.54 DPS) [crafted]; Prowler's Leather Waistguard (252473, -31.09 DPS) [crafted]; Girdle of Beastial Fury (11686, -166.35 DPS) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 34.9 stamina points (1832.07 DPS) | yes | Windscale Sarong (10842, -281.30 DPS) [world]; Stone Guard's Crackling Leather Leggings (220865, -298.27 DPS) [vendor]; Basilisk Hide Pants (1718, -343.41 DPS) [world_drop] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -24.79 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (787.60 DPS) | yes | Darkspear Signet (272069, -52.51 DPS) [vendor]; Underworld Band (1980, -157.52 DPS) [world_drop]; Nogg's Gold Ring (9588, -188.62 DPS) [quest] |
| finger2 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.6 stamina points (764.88 DPS) | yes | Darkspear Signet (272069, -29.78 DPS) [vendor]; Underworld Band (1980, -134.79 DPS) [world_drop]; Nogg's Gold Ring (9588, -165.89 DPS) [quest] |
| trinket1 | Defiler's Talisman (21115) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -52.51 DPS) [quest]; Rune of Perfection (21565, -105.01 DPS) [rep]; Rune of Duty (21567, -105.01 DPS) [rep] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -52.51 DPS) [quest]; Rune of Perfection (21565, -105.01 DPS) [rep] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -176.67 DPS) [dungeon]; Warden Staff (943, -177.28 DPS) [world_drop]; Blanchard's Stout (13046, -219.59 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Sergeant's Cloak; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal

No-known-source sample (15 of 689, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 193.6. Weights run: 8.1s. Verify run: 53.4s. 1759 eligible items had no known source.

Race sweep: windshaper-skyborne 753.4 ± 0.8, tauren 751.9 ± 0.8 (threat_per_second); windshaper-skyborne beats the guide race tauren by 1.5 (0.20%) ± 1.1, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.027, armor=0.157 ± 0.017, defense=not significant (0.982 ± 0.272) per rating point (1 rating = 1%, 0.982 per %), dodge=0.389 ± 0.043 per rating point (12 rating = 1%, 4.665 per %), strength=0.095 ± 0.000, agility=0.376 ± 0.027, attack_power=0.047 ± 0.000, feral_attack_power=0.047 ± 0.000, hit=not significant (0.125 ± 0.039) per rating point (10 rating = 1%, 1.253 per %), crit=0.080 ± 0.017 per rating point (14 rating = 1%, 1.115 per %), expertise=5.031 ± 0.422

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 65.7 stamina points (5039.28 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, -308.76 DPS, sim-verified) [crafted] |
| neck | Senior Sergeant's Insignia (16335) | PvP rank 4 · Senior Sergeant · Horde [vendor] | 17.0 stamina points (1304.47 DPS) | yes | Medallion of Grand Marshal Morris (13091, -13.49 DPS) [world_drop]; Frostwolf Legionnaire's Pendant (19095, -87.87 DPS) [rep]; Scout's Medallion (19534, -104.16 DPS) [rep] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 48.4 stamina points (3714.91 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -1512.43 DPS, sim-verified) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phantasmal Cloak (18689, -420.20 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -800.12 DPS) [dungeon]; Fel Cape (279269, -832.32 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.6 stamina points (4573.46 DPS) | yes | Breastplate of Bloodthirst (12757, +0.00 DPS) [quest]; Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Malefic Bracers (18700, +0.00 DPS) [dungeon]; Yeti Hide Bracers (19113, +0.00 DPS) [world]; First Sergeant's Dragonhide Armguards (18434, -46.34 DPS) [vendor] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Hivethrasher's Claws (275613, -572.94 DPS, sim-verified) [crafted] |
| waist | Elderwild Waistcord (279252) | Leatherworking [crafted] | 38.4 stamina points (2947.56 DPS) | yes | Hivethrasher's Girdle (275614, -14.62 DPS) [crafted]; Shifter's Belt (272396, -56.80 DPS) [vendor]; Frostbite Girdle (14502, -147.47 DPS) [dungeon] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 65.8 stamina points (5048.09 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Warstrife Leggings (11821, -1295.15 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | 42.0 stamina points (3223.18 DPS) | yes | Ash Covered Boots (18716, +0.00 DPS) [dungeon]; General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Thrall's Resolve (12544) | The Princess Saved? [quest] | 30.9 stamina points (2371.91 DPS) | yes | Myrmidon's Signet (2246, -792.50 DPS) [world_drop]; Ring of Awareness (272409, -867.81 DPS) [vendor]; Band of Flesh (13373, -1013.85 DPS) [dungeon] |
| finger2 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, -166.72 DPS) [world_drop]; Ring of Awareness (272409, -242.03 DPS) [vendor]; Band of Flesh (13373, -388.07 DPS) [dungeon] |
| trinket1 | Defiler's Talisman (20072) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Adaptive Combat Assistant (272437, -148.76 DPS) [vendor]; Frostwolf Insignia Rank 6 (17909, -204.90 DPS) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Adaptive Combat Assistant (272437, +0.00 DPS) [vendor]; Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -464.16 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Senior Sergeant's Insignia; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Elderwild Waistcord; legs: Outrider's Leather Pants; feet: Dawn Treaders; finger1: Thrall's Resolve; finger2: Naglering; trinket1: Defiler's Talisman; main_hand: Headmaster's Charge

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 476.3. Weights run: 9.0s. Verify run: 53.3s. 1759 eligible items had no known source.

Race sweep: windshaper-skyborne 1803.3 ± 2.2, tauren 1791.3 ± 2.1 (threat_per_second); windshaper-skyborne beats the guide race tauren by 12.0 (0.67%) ± 3.0, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.027, armor=0.178 ± 0.017, defense=1.401 ± 0.295 per rating point (1 rating = 1%, 1.401 per %), dodge=0.552 ± 0.051 per rating point (12 rating = 1%, 6.627 per %), strength=0.102 ± 0.000, agility=0.530 ± 0.033, attack_power=0.046 ± 0.000, feral_attack_power=0.046 ± 0.000, hit=not significant (0.149 ± 0.040) per rating point (10 rating = 1%, 1.491 per %), crit=0.104 ± 0.018 per rating point (14 rating = 1%, 1.455 per %), expertise=6.931 ± 0.450

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.9 stamina points (8423.97 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, -609.51 DPS, sim-verified) [crafted] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 21.0 stamina points (2429.39 DPS) | yes | Scout's Medallion (19534, -353.71 DPS) [rep]; Senior Sergeant's Insignia (16335, -463.71 DPS) [vendor]; Frostwolf Legionnaire's Pendant (19095, -598.45 DPS) [rep] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 54.3 stamina points (6282.27 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -2142.92 DPS, sim-verified) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Phantasmal Cloak (18689, -891.13 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -1446.51 DPS) [dungeon]; Fel Cape (279269, -1562.93 DPS) [crafted] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 65.8 stamina points (7613.70 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Breastplate of Bloodthirst (12757, -637.86 DPS, sim-verified) [quest] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 34.5 stamina points (3993.08 DPS) | yes | Yeti Hide Bracers (19113, -213.26 DPS) [world]; Malefic Bracers (18700, -229.00 DPS) [dungeon]; First Sergeant's Dragonhide Armguards (18434, -439.29 DPS) [vendor] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+1310.9 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 43.7 stamina points (5053.06 DPS) | yes | Cloudrunner Girdle (13252, -152.94 DPS) [dungeon]; Hivethrasher's Girdle (275614, -277.72 DPS) [crafted]; Frostbite Girdle (14502, -319.04 DPS) [dungeon] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 74.2 stamina points (8574.50 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Warstrife Leggings (11821, -2332.36 DPS, sim-verified) [dungeon] |
| feet | Dawn Treaders (19052) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Thrall's Resolve (12544) | The Princess Saved? [quest] | 34.2 stamina points (3950.30 DPS) | yes | Ring of Awareness (272409, -1241.66 DPS) [vendor]; Myrmidon's Signet (2246, -1437.61 DPS) [world_drop]; Band of Flesh (13373, -1845.59 DPS) [dungeon] |
| finger2 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Awareness (272409, -288.87 DPS) [vendor]; Myrmidon's Signet (2246, -484.82 DPS) [world_drop]; Band of Flesh (13373, -892.81 DPS) [dungeon] |
| trinket1 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-verified (+2099.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Defiler's Talisman (20072, -144.88 DPS) [rep]; Darkspear Voodoo Seal (272061, -376.14 DPS) [vendor] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Adaptive Combat Assistant (272437, -503.10 DPS) [vendor]; Defiler's Talisman (20072, -599.40 DPS, sim-verified) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; The Unstoppable Force (19323, -1473.55 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Medallion of Grand Marshal Morris; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Outrider's Leather Pants; feet: Dawn Treaders; finger1: Thrall's Resolve; finger2: Naglering; trinket1: Frostwolf Insignia Rank 6; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1759, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

