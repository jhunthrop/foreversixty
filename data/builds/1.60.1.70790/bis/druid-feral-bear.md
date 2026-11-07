# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 22.7. Weights run: 8.8s. Verify run: 8.1s. 228 eligible items had no known source.

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

Set DPS (verified): 44.1. Weights run: 9.2s. Verify run: 17.0s. 400 eligible items had no known source.

Race sweep: high-order-skyborne 177.0 ± 0.2, night-elf 175.1 ± 0.2 (threat_per_second); high-order-skyborne beats the guide race night-elf by 1.9 (1.09%) ± 0.3, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.086 ± 0.003, defense=0.344 ± 0.056 per rating point (1 rating = 1%, 0.344 per %), dodge=0.153 ± 0.015 per rating point (12 rating = 1%, 1.833 per %), strength=0.060 ± 0.000, agility=0.187 ± 0.009, attack_power=0.030 ± 0.000, feral_attack_power=0.030 ± 0.000, hit=0.058 ± 0.011 per rating point (10 rating = 1%, 0.584 per %), crit=0.027 ± 0.003 per rating point (14 rating = 1%, 0.376 per %), expertise=1.830 ± 0.129

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 22.1 stamina points (640.17 DPS) | yes | Defender's Leather Helm (252455, -33.32 DPS) [crafted]; Trapper's Leather Helm (252513, -54.25 DPS) [crafted]; Stormrider's Leather Helm (252514, -54.25 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.2 stamina points (267.99 DPS) | yes | Master Sergeant's Insignia (18442, -6.98 DPS) [vendor]; Seal of the Expedition (271920, -35.98 DPS) [rep]; Sentinel's Medallion (19541, -79.58 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 18.0 stamina points (523.21 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -69.82 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -93.21 DPS) [crafted]; Barbaric Shoulders (5964, -124.34 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mourning Shawl (6751, -38.95 DPS) [quest]; Enduring Cape (14763, -53.61 DPS) [world_drop]; Tigerstrike Mantle (13108, -110.68 DPS, sim-verified) [world_drop] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.4 stamina points (620.74 DPS) | yes | Spirewind Fetter (9406, -10.97 DPS) [dungeon]; Green Whelp Armor (7375, -48.05 DPS) [crafted]; Blackened Defias Armor (10399, -86.13 DPS) [dungeon] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.1 stamina points (321.16 DPS) | yes | Barbaric Bracers (18948, -1.58 DPS) [crafted]; Black Wolf Bracers (3230, -7.72 DPS) [dungeon]; Glowing Leather Bracers (2017, -65.82 DPS) [world] |
| hands | Brawler Gloves (720) | World drop [world_drop] | 12.8 stamina points (371.16 DPS) | yes | Wolfclaw Gloves (1978, -4.17 DPS) [dungeon]; Gloves of Insight (9698, -16.44 DPS) [quest]; Emblazoned Gloves (6397, -28.27 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.9 stamina points (460.71 DPS) | yes | Stalker's Leather Belt (252521, -81.19 DPS) [crafted]; Prowler's Leather Belt (252459, -81.77 DPS) [crafted]; Skulker's Leather Belt (252520, -145.03 DPS, sim-verified) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.4 stamina points (535.03 DPS) | yes | Dark Ritual Leggings (270031, -55.16 DPS) [quest]; Trapper's Leather Legguards (252517, -59.36 DPS) [crafted]; Brawler's Leather Legguards (252516, -170.77 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.5 stamina points (537.03 DPS) | yes | Harbinger Boots (7754, -32.56 DPS, sim-verified) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -99.48 DPS) [quest]; Emblazoned Boots (4051, -158.57 DPS) [world_drop] |
| finger1 | Theramore Signet (271911) | Theramore Expeditionary Force [rep] | 8.3 stamina points (239.86 DPS) | yes | Plains Ring (2039, -7.85 DPS) [dungeon]; Defias Renegade Ring (1076, -44.15 DPS) [dungeon]; Protector's Band (19517, -80.84 DPS) [rep] |
| finger2 | Militant's Band (271916) (or Plains Ring (2039)) | Theramore Expeditionary Force [rep] | 8.0 stamina points (232.01 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Defias Renegade Ring (1076, -36.30 DPS) [dungeon]; Protector's Band (19517, -72.99 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (+133.5 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Spine of the Devourer (273040) | Mana Devourer [world] | sim-decided (no score - a real sim tournament chose this pick) | yes | Staff of the Shade (2549, -21.75 DPS) [dungeon]; Slaghammer (1976, -59.50 DPS) [dungeon]; Lorekeeper's Staff (212580, -67.64 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Brawler Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Theramore Signet; finger2: Militant's Band; trinket1: Talisman of Arathor; trinket2: Rune of Perfection; main_hand: Spine of the Devourer

No-known-source sample (15 of 400, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 88.6. Weights run: 9.2s. Verify run: 9.3s. 526 eligible items had no known source.

Race sweep: high-order-skyborne 344.9 ± 0.4, night-elf 343.8 ± 0.4 (threat_per_second); high-order-skyborne beats the guide race night-elf by 1.2 (0.35%) ± 0.5, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.026, armor=0.158 ± 0.008, defense=0.444 ± 0.099 per rating point (1 rating = 1%, 0.444 per %), dodge=0.171 ± 0.024 per rating point (12 rating = 1%, 2.053 per %), strength=0.087 ± 0.000, agility=0.226 ± 0.015, attack_power=0.044 ± 0.000, feral_attack_power=0.044 ± 0.000, hit=0.108 ± 0.021 per rating point (10 rating = 1%, 1.081 per %), crit=0.050 ± 0.008 per rating point (14 rating = 1%, 0.704 per %), expertise=2.987 ± 0.241

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 30.8 stamina points (997.80 DPS) | yes | Warden's Wizard Hat (14604, -42.38 DPS) [world_drop]; Brawler's Leather Helm (252512, -76.03 DPS) [crafted]; Nightscape Headband (8176, -90.46 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shriveled Heart (9243, +0.00 DPS) [dungeon]; Seal of the Expedition (271919, -11.30 DPS) [rep]; Sentinel's Medallion (19540, -60.43 DPS) [rep] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.8 stamina points (1027.48 DPS) | yes | Flintrock Shoulders (7755, -159.41 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -255.31 DPS) [dungeon]; Nightscape Shoulders (8192, -347.04 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -19.00 DPS) [dungeon]; Heraldic Cloak (8120, -22.16 DPS) [dungeon]; Well Oiled Cloak (12254, -24.09 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.5 stamina points (1147.27 DPS) | yes | Warden's Wraps (14601, -59.74 DPS) [world_drop]; Robes of the Lich (10762, -184.15 DPS) [dungeon]; Spirewind Fetter (9406, -200.67 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.5 stamina points (535.34 DPS) | yes | Branded Leather Bracers (19508, -32.60 DPS) [dungeon]; Barbaric Bracers (18948, -61.14 DPS) [crafted]; Cultist's Armguards (270032, -70.45 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -24.08 DPS) [crafted]; Prowler's Leather Gloves (252524, -28.55 DPS) [crafted]; Stalker's Leather Gloves (252526, -410.39 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.0 stamina points (744.21 DPS) | yes | Warden's Leather Belt (252460, -71.15 DPS) [crafted]; Kolkar Hunter's Belt (6788, -76.25 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -86.82 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 30.5 stamina points (988.11 DPS) | yes | Panther Hunter Leggings (4108, -69.88 DPS) [quest]; Warden's Woolies (14605, -74.82 DPS) [world_drop]; Imperial Leather Pants (4062, -134.00 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 27.7 stamina points (895.73 DPS) | yes | Skulker's Leather Shoes (252531, -100.22 DPS) [crafted]; Prowler's Leather Shoes (252465, -118.11 DPS) [crafted]; Worn Running Boots (9398, -309.14 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (388.24 DPS) | yes | Talvash's Gold Ring (9538, -21.64 DPS) [quest]; Militant's Band (271915, -32.35 DPS) [rep]; Dragonclaw Ring (10710, -64.71 DPS) [quest] |
| finger2 | Theramore Signet (271910) | Theramore Expeditionary Force [rep] | 11.5 stamina points (372.83 DPS) | yes | Talvash's Gold Ring (9538, -6.24 DPS) [quest]; Militant's Band (271915, -16.95 DPS) [rep]; Dragonclaw Ring (10710, -49.30 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (+272.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -32.35 DPS) [world]; Enriched Seal (271905, -32.35 DPS) [rep] |
| trinket2 | Talisman of Arathor (21118) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -32.35 DPS) [world]; Volunteer's Lucky Seal (271904, -32.35 DPS) [rep]; Enriched Seal (271905, -32.35 DPS) [rep] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -158.61 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -376.82 DPS) [dungeon]; Huge Stone Club (1523, -463.42 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Theramore Signet; trinket1: Rune of Duty; trinket2: Talisman of Arathor; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 526, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 129.9. Weights run: 9.3s. Verify run: 21.3s. 690 eligible items had no known source.

Race sweep: night-elf 507.5 ± 0.6, high-order-skyborne 507.2 ± 0.5 (threat_per_second); the guide race night-elf is best.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.035, armor=0.137 ± 0.010, defense=not significant (0.450 ± 0.114) per rating point (1 rating = 1%, 0.450 per %), dodge=0.256 ± 0.030 per rating point (12 rating = 1%, 3.069 per %), strength=0.064 ± 0.000, agility=0.259 ± 0.020, attack_power=0.032 ± 0.000, feral_attack_power=0.032 ± 0.000, hit=0.143 ± 0.030 per rating point (10 rating = 1%, 1.427 per %), crit=not significant (0.037 ± 0.010) per rating point (14 rating = 1%, 0.522 per %), expertise=3.864 ± 0.334

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 40.9 stamina points (2204.40 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -211.58 DPS) [vendor]; Embrace of the Lycan (9479, -389.21 DPS) [dungeon]; Scorpashi Skullcap (14658, -516.10 DPS) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 4 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (754.90 DPS) | yes | Shriveled Heart (9243, -53.92 DPS) [dungeon]; Lifeblood Amulet (9641, -53.92 DPS) [dungeon]; Seal of the Expedition (271918, -53.92 DPS) [rep] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Fleshhide Shoulders (10774, +0.00 DPS) [dungeon]; Knight-Lieutenant's Leather Shoulders (220852, -12.75 DPS) [vendor]; Phytoskin Spaulders (17749, -13.54 DPS) [dungeon] |
| back | Sergeant's Cape (18441) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Graverot Cape (11677, +0.00 DPS) [dungeon]; Nightfall Drape (12465, +0.00 DPS) [dungeon]; Emberscale Cape (11311, -14.74 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, -559.01 DPS) [dungeon]; Jinxed Hoodoo Skin (9473, -735.61 DPS) [dungeon]; Heraldic Breastplate (8119, -822.61 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 24.3 stamina points (1310.51 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -83.84 DPS) [vendor]; Warden's Leather Bracers (252541, -245.51 DPS) [crafted]; Serpentskin Bracers (8257, -247.91 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -199.74 DPS) [crafted]; Wicked Leather Gauntlets (15083, -246.66 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -252.39 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 26.9 stamina points (1452.27 DPS) | yes | Skulker's Leather Waistguard (252474, -20.06 DPS) [crafted]; Prowler's Leather Waistguard (252473, -41.12 DPS) [crafted]; Girdle of Beastial Fury (11686, -177.82 DPS) [dungeon] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 35.8 stamina points (1932.65 DPS) | yes | Windscale Sarong (10842, -293.11 DPS) [world]; Knight's Crackling Leather Leggings (220864, -300.59 DPS) [vendor]; Scorpashi Leggings (14659, -372.98 DPS) [world_drop] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -34.04 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (808.82 DPS) | yes | Militant's Band (271914, -53.92 DPS) [rep]; Underworld Band (1980, -161.76 DPS) [world_drop]; Talvash's Gold Ring (9538, -196.83 DPS) [quest] |
| finger2 | Theramore Signet (271909) | Theramore Expeditionary Force [rep] | 14.5 stamina points (780.72 DPS) | yes | Militant's Band (271914, -25.82 DPS) [rep]; Underworld Band (1980, -133.66 DPS) [world_drop]; Talvash's Gold Ring (9538, -168.73 DPS) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Rune of Perfection (21565, -107.84 DPS) [rep] |
| trinket2 | Guardian Talisman (1490) | Into The Temple of Atal'Hakkar [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Rune of Perfection (21565, -53.92 DPS) [rep] |
| main_hand | Blanchard's Stout (13046) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warden Staff (943, +0.00 DPS) [world_drop]; Ironshod Bludgeon (9408, +0.00 DPS) [dungeon]; Radiant Staff (249453, -120.54 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Penance Spaulders; back: Sergeant's Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Theramore Signet; trinket1: Enriched Seal; trinket2: Guardian Talisman; main_hand: Blanchard's Stout

No-known-source sample (15 of 690, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 189.7. Weights run: 9.4s. Verify run: 54.5s. 1755 eligible items had no known source.

Race sweep: high-order-skyborne 742.9 ± 0.9, night-elf 742.3 ± 0.9 (threat_per_second); high-order-skyborne beats the guide race night-elf by 0.5 (0.07%) ± 1.2, inside error.

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

Set DPS (verified): 453.8. Weights run: 10.3s. Verify run: 63.7s. 1755 eligible items had no known source.

Race sweep: high-order-skyborne 1742.1 ± 2.1, night-elf 1741.6 ± 2.2 (threat_per_second); high-order-skyborne beats the guide race night-elf by 0.5 (0.03%) ± 3.0, inside error.

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

### Band 20 (tauren, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 23.2. Weights run: 8.8s. Verify run: 7.5s. 224 eligible items had no known source.

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

Set DPS (verified): 53.1. Weights run: 9.2s. Verify run: 16.7s. 401 eligible items had no known source.

Race sweep: windshaper-skyborne 211.6 ± 0.2, tauren 209.4 ± 0.2 (threat_per_second); windshaper-skyborne beats the guide race tauren by 2.1 (1.02%) ± 0.3, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.107 ± 0.002, defense=0.376 ± 0.058 per rating point (1 rating = 1%, 0.376 per %), dodge=0.166 ± 0.015 per rating point (12 rating = 1%, 1.993 per %), strength=0.096 ± 0.000, agility=0.211 ± 0.009, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.078 ± 0.011 per rating point (10 rating = 1%, 0.775 per %), crit=0.026 ± 0.002 per rating point (14 rating = 1%, 0.369 per %), expertise=1.993 ± 0.136

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.1 stamina points (534.71 DPS) | yes | Defender's Leather Helm (252455, -21.32 DPS) [crafted]; Trapper's Leather Helm (252513, -46.79 DPS) [crafted]; Stormrider's Leather Helm (252514, -46.79 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.4 stamina points (207.99 DPS) | yes | Senior Sergeant's Insignia (15200, -8.49 DPS) [vendor]; Darkspear Warding Pendant (272075, -30.66 DPS) [vendor]; Scout's Medallion (19537, -59.72 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 19.9 stamina points (440.44 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -66.59 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -71.79 DPS) [crafted]; Barbaric Shoulders (5964, -92.71 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enduring Cape (14763, -40.24 DPS) [world_drop]; Grimsteel Cape (4643, -42.11 DPS) [quest]; Tigerstrike Mantle (13108, -108.81 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | 23.7 stamina points (525.03 DPS) | yes | Raptor Hide Harness (4455, -3.82 DPS) [crafted]; Green Whelp Armor (7375, -38.64 DPS) [crafted]; Blackened Defias Armor (10399, -69.19 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.3 stamina points (271.97 DPS) | yes | Cultist's Armguards (270032, -1.56 DPS) [quest]; Black Wolf Bracers (3230, -10.31 DPS) [dungeon]; Glowing Leather Bracers (2017, -56.66 DPS) [world] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 15.3 stamina points (340.11 DPS) | yes | Brawler Gloves (720, -20.53 DPS) [world_drop]; Wolfclaw Gloves (1978, -23.37 DPS) [dungeon]; Emblazoned Gloves (6397, -48.17 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 17.6 stamina points (390.12 DPS) | yes | Skulker's Leather Belt (252520, -52.46 DPS) [crafted]; Prowler's Leather Belt (252459, -60.13 DPS) [crafted]; Larcenist's Belt (279846, -162.91 DPS, sim-verified) [quest] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 20.8 stamina points (461.13 DPS) | yes | Petrolspill Leggings (9509, -48.60 DPS) [dungeon]; Trapper's Leather Legguards (252517, -49.83 DPS) [crafted]; Brawler's Leather Legguards (252516, -204.08 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 20.2 stamina points (446.73 DPS) | yes | Harbinger Boots (7754, -39.13 DPS, sim-verified) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -95.12 DPS) [quest]; Gravewalker Boots (251963, -99.33 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.4 stamina points (186.88 DPS) | yes | Plains Ring (2039, -9.55 DPS) [dungeon]; Darkspear Signet (272071, -9.55 DPS) [vendor]; Defias Renegade Ring (1076, -35.17 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.3 stamina points (183.70 DPS) | yes | Plains Ring (2039, -6.37 DPS) [dungeon]; Darkspear Signet (272071, -6.37 DPS) [vendor]; Defias Renegade Ring (1076, -31.98 DPS) [dungeon] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | sim-verified (+146.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Spine of the Devourer (273040, -5.92 DPS) [world]; Slaghammer (1976, -25.85 DPS) [dungeon]; Advisor's Gnarled Staff (212584, -35.75 DPS) [vendor] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Defiler's Talisman; trinket2: Rune of Duty; main_hand: Viscous Hammer

No-known-source sample (15 of 401, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 87.4. Weights run: 9.2s. Verify run: 9.1s. 525 eligible items had no known source.

Race sweep: windshaper-skyborne 340.9 ± 0.4, tauren 339.0 ± 0.4 (threat_per_second); windshaper-skyborne beats the guide race tauren by 1.9 (0.57%) ± 0.5, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.026, armor=0.158 ± 0.008, defense=0.444 ± 0.099 per rating point (1 rating = 1%, 0.444 per %), dodge=0.171 ± 0.024 per rating point (12 rating = 1%, 2.053 per %), strength=0.087 ± 0.000, agility=0.226 ± 0.015, attack_power=0.044 ± 0.000, feral_attack_power=0.044 ± 0.000, hit=0.108 ± 0.021 per rating point (10 rating = 1%, 1.081 per %), crit=0.050 ± 0.008 per rating point (14 rating = 1%, 0.704 per %), expertise=2.987 ± 0.241

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 30.8 stamina points (997.80 DPS) | yes | Warden's Wizard Hat (14604, -42.38 DPS) [world_drop]; Brawler's Leather Helm (252512, -76.03 DPS) [crafted]; Nightscape Headband (8176, -90.46 DPS) [crafted] |
| neck | Gazlowe's Charm (13088) | Zul'Farrak: Witch Doctor's Chest [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shriveled Heart (9243, +0.00 DPS) [dungeon]; Dragon's Blood Necklace (10711, +0.00 DPS) [quest]; Darkspear Warding Pendant (272074, -11.30 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.8 stamina points (1027.48 DPS) | yes | Flintrock Shoulders (7755, -128.02 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -255.31 DPS) [dungeon]; Nightscape Shoulders (8192, -347.04 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -19.00 DPS) [dungeon]; Heraldic Cloak (8120, -22.16 DPS) [dungeon]; Well Oiled Cloak (12254, -24.09 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.5 stamina points (1147.27 DPS) | yes | Warden's Wraps (14601, -59.74 DPS) [world_drop]; Robes of the Lich (10762, -184.15 DPS) [dungeon]; Spirewind Fetter (9406, -200.67 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.5 stamina points (535.34 DPS) | yes | Branded Leather Bracers (19508, -32.60 DPS) [dungeon]; Barbaric Bracers (18948, -61.14 DPS) [crafted]; Cultist's Armguards (270032, -70.45 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -24.08 DPS) [crafted]; Prowler's Leather Gloves (252524, -28.55 DPS) [crafted]; Stalker's Leather Gloves (252526, -370.03 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.0 stamina points (744.21 DPS) | yes | Kolkar Hunter's Belt (6788, -63.30 DPS) [quest]; Warden's Leather Belt (252460, -71.15 DPS) [crafted]; Scorpashi Sash (14652, -86.82 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 30.5 stamina points (988.11 DPS) | yes | Panther Hunter Leggings (4108, -69.88 DPS) [quest]; Warden's Woolies (14605, -74.82 DPS) [world_drop]; Imperial Leather Pants (4062, -126.72 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 27.7 stamina points (895.73 DPS) | yes | Skulker's Leather Shoes (252531, -100.22 DPS) [crafted]; Prowler's Leather Shoes (252465, -118.11 DPS) [crafted]; Worn Running Boots (9398, -329.56 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (388.24 DPS) | yes | Nogg's Gold Ring (9588, -21.64 DPS) [quest]; Darkspear Signet (272070, -32.35 DPS) [vendor]; Dragonclaw Ring (10710, -64.71 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.5 stamina points (372.83 DPS) | yes | Nogg's Gold Ring (9588, -6.24 DPS) [quest]; Darkspear Signet (272070, -16.95 DPS) [vendor]; Dragonclaw Ring (10710, -49.30 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (+264.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -32.35 DPS) [world]; Darkspear Voodoo Seal (272059, -32.35 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21116) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -32.35 DPS) [world]; Darkspear Voodoo Seal (272059, -32.35 DPS) [vendor]; Relentless Raider's Seal (272062, -32.35 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -122.83 DPS, sim-verified) [dungeon]; Mograine's Might (7723, -267.71 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -289.45 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Gazlowe's Charm; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Insurgent's Band; trinket1: Rune of Duty; trinket2: Defiler's Talisman; main_hand: Cragwood Maul

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 120.6. Weights run: 9.3s. Verify run: 19.2s. 689 eligible items had no known source.

Race sweep: windshaper-skyborne 474.5 ± 0.5, tauren 471.5 ± 0.5 (threat_per_second); windshaper-skyborne beats the guide race tauren by 3.0 (0.64%) ± 0.7, beyond error.

2 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.035, armor=0.137 ± 0.010, defense=not significant (0.450 ± 0.114) per rating point (1 rating = 1%, 0.450 per %), dodge=0.256 ± 0.030 per rating point (12 rating = 1%, 3.069 per %), strength=0.064 ± 0.000, agility=0.259 ± 0.020, attack_power=0.032 ± 0.000, feral_attack_power=0.032 ± 0.000, hit=0.143 ± 0.030 per rating point (10 rating = 1%, 1.427 per %), crit=not significant (0.037 ± 0.010) per rating point (14 rating = 1%, 0.522 per %), expertise=3.864 ± 0.334

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 40.9 stamina points (2204.40 DPS) | yes | Blood Guard's Leather Headband (220851, -211.58 DPS) [vendor]; Embrace of the Lycan (9479, -389.21 DPS) [dungeon]; Scorpashi Skullcap (14658, -516.10 DPS) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 4 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (754.90 DPS) | yes | Shriveled Heart (9243, -53.92 DPS) [dungeon]; Lifeblood Amulet (9641, -53.92 DPS) [dungeon]; Darkspear Warding Pendant (272073, -53.92 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 29.9 stamina points (1610.16 DPS) | yes | Penance Spaulders (11963, -15.60 DPS) [quest]; Blood Guard's Leather Shoulders (220853, -28.35 DPS) [vendor]; Phytoskin Spaulders (17749, -29.15 DPS) [dungeon] |
| back | Sergeant's Cloak (16341) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Graverot Cape (11677, +0.00 DPS) [dungeon]; Nightfall Drape (12465, +0.00 DPS) [dungeon]; Emberscale Cape (11311, -14.74 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 54.0 stamina points (2909.95 DPS) | yes | Mixologist's Tunic (12793, -559.01 DPS) [dungeon]; Jinxed Hoodoo Skin (9473, -735.61 DPS) [dungeon]; Heraldic Breastplate (8119, -822.61 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 24.3 stamina points (1310.51 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -83.84 DPS) [vendor]; Warden's Leather Bracers (252541, -245.51 DPS) [crafted]; Serpentskin Bracers (8257, -247.91 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -199.74 DPS) [crafted]; Wicked Leather Gauntlets (15083, -246.66 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -252.39 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 26.9 stamina points (1452.27 DPS) | yes | Skulker's Leather Waistguard (252474, -20.06 DPS) [crafted]; Prowler's Leather Waistguard (252473, -41.12 DPS) [crafted]; Girdle of Beastial Fury (11686, -177.82 DPS) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 35.8 stamina points (1932.65 DPS) | yes | Windscale Sarong (10842, -293.11 DPS) [world]; Stone Guard's Crackling Leather Leggings (220865, -300.59 DPS) [vendor]; Scorpashi Leggings (14659, -372.98 DPS) [world_drop] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, -34.04 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (808.82 DPS) | yes | Darkspear Signet (272069, -53.92 DPS) [vendor]; Underworld Band (1980, -161.76 DPS) [world_drop]; Nogg's Gold Ring (9588, -196.83 DPS) [quest] |
| finger2 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.5 stamina points (780.72 DPS) | yes | Darkspear Signet (272069, -25.82 DPS) [vendor]; Underworld Band (1980, -133.66 DPS) [world_drop]; Nogg's Gold Ring (9588, -168.73 DPS) [quest] |
| trinket1 | Defiler's Talisman (21115) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -53.92 DPS) [quest]; Rune of Perfection (21565, -107.84 DPS) [rep]; Rune of Duty (21567, -107.84 DPS) [rep] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -53.92 DPS) [quest]; Rune of Perfection (21565, -107.84 DPS) [rep] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -186.77 DPS) [dungeon]; Warden Staff (943, -233.91 DPS) [world_drop]; Blanchard's Stout (13046, -265.97 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Sergeant's Cloak; chest: Warbear Harness; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Warden's Leather Boots; finger1: Ring of Saviors; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal

No-known-source sample (15 of 689, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 182.6. Weights run: 9.4s. Verify run: 51.2s. 1759 eligible items had no known source.

Race sweep: windshaper-skyborne 720.8 ± 0.8, tauren 714.8 ± 0.8 (threat_per_second); windshaper-skyborne beats the guide race tauren by 6.0 (0.84%) ± 1.2, beyond error.

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

Set DPS (verified): 443.7. Weights run: 10.3s. Verify run: 63.4s. 1759 eligible items had no known source.

Race sweep: windshaper-skyborne 1718.7 ± 2.1, tauren 1699.4 ± 2.1 (threat_per_second); windshaper-skyborne beats the guide race tauren by 19.3 (1.14%) ± 3.0, beyond error.

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

