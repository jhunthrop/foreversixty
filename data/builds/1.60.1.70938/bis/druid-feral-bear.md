# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 22.9. Weights run: 4.5s. Verify run: 5.3s. 251 eligible items had no known source.

Race sweep: night-elf 93.0 ± 0.2, high-order-skyborne 92.5 ± 0.2 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.091 ± 0.000, defense=0.222 ± 0.043 per rating point (1 rating = 1%, 0.222 per %), dodge=0.107 ± 0.014 per rating point (12 rating = 1%, 1.289 per %), strength=0.097 ± 0.000, agility=0.160 ± 0.009, attack_power=0.049 ± 0.000, feral_attack_power=0.049 ± 0.000, hit=0.049 ± 0.008 per rating point (10 rating = 1%, 0.489 per %), crit=0.015 ± 0.001 per rating point (14 rating = 1%, 0.213 per %), expertise=1.150 ± 0.097

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 18.2 stamina points (329.61 DPS) | yes | Defender's Leather Hood (252447, -9.09 DPS) [crafted]; Trapper's Leather Hood (252505, -23.19 DPS) [crafted]; Stormrider's Leather Hood (252506, -23.19 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.7 stamina points (103.23 DPS) | yes | Erudite's Amulet (277204, -37.70 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -48.97 DPS) [quest]; Sentinel's Medallion (20444, -49.66 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.8 stamina points (178.07 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -51.25 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -61.12 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -4.84 DPS) [world_drop]; Firebane Cloak (12979, -19.74 DPS) [world_drop]; Caretaker's Cape (20428, -36.17 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 19.0 stamina points (343.64 DPS) | yes | Gloomshroud Armor (1489, -34.85 DPS) [dungeon]; Tunic of Westfall (2041, -69.34 DPS) [quest]; Defender's Leather Armor (252434, -91.75 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.4 stamina points (151.55 DPS) | yes | Drakewing Bands (12999, -29.60 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -42.67 DPS) [vendor]; Sanguine Cuffs (14375, -51.12 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.5 stamina points (171.49 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.65 DPS) [crafted]; Defender's Leather Gloves (252496, -4.54 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.8 stamina points (212.55 DPS) | yes | Belt of the Fang (10412, -32.48 DPS) [dungeon]; Dark Leather Belt (4249, -49.32 DPS) [crafted]; Guardsman Belt (3429, -50.97 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 15.0 stamina points (271.44 DPS) | yes | Brawler's Leather Pants (252500, -26.97 DPS) [crafted]; Leggings of the Fang (10410, -35.34 DPS) [dungeon]; Smelting Pants (5199, -81.28 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.3 stamina points (276.51 DPS) | yes | Surfer Shoes (276274, -48.21 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -49.83 DPS) [dungeon]; Trapper's Leather Boots (252440, -70.82 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (90.44 DPS) | yes | Slain Baron's Signet (279867, -32.15 DPS) [quest]; Ring of the Moon (12052, -34.41 DPS) [world_drop]; Protector's Band (20439, -35.61 DPS) [rep] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 4.0 stamina points (72.35 DPS) | yes | Slain Baron's Signet (279867, -14.06 DPS) [quest]; Ring of the Moon (12052, -16.32 DPS) [world_drop]; Protector's Band (20439, -17.53 DPS) [rep] |
| trinket1 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.9 stamina points (377.23 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.98 DPS) [dungeon]; Twisted Chanter's Staff (890, -16.42 DPS) [world_drop]; Staff of the Blessed Seer (2271, -21.09 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Sustaining Ring; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 251, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-51100332020102100000-0000000000000000)

Set DPS (verified): 66.8. Weights run: 3.9s. Verify run: 7.6s. 424 eligible items had no known source.

Race sweep: night-elf 252.4 ± 0.5, high-order-skyborne 249.7 ± 0.4 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.107 ± 0.000, defense=0.424 ± 0.088 per rating point (1 rating = 1%, 0.424 per %), dodge=0.178 ± 0.024 per rating point (12 rating = 1%, 2.131 per %), strength=0.150 ± 0.000, agility=0.245 ± 0.015, attack_power=0.075 ± 0.000, feral_attack_power=0.075 ± 0.000, hit=not significant (0.062 ± 0.018) per rating point (10 rating = 1%, 0.619 per %), crit=0.045 ± 0.004 per rating point (14 rating = 1%, 0.632 per %), expertise=2.483 ± 0.213

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.4 stamina points (506.33 DPS) | yes | Defender's Leather Helm (252455, -13.37 DPS) [crafted]; Trapper's Leather Helm (252513, -50.79 DPS) [crafted]; Stormrider's Leather Helm (252514, -50.79 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.6 stamina points (199.11 DPS) | yes | Master Sergeant's Insignia (18442, -12.47 DPS) [vendor]; Seal of the Expedition (271920, -33.21 DPS) [rep]; Sentinel's Medallion (19541, -54.79 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 20.0 stamina points (414.56 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -62.50 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -65.61 DPS) [crafted]; Barbaric Shoulders (5964, -80.91 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mourning Shawl (6751, -29.59 DPS) [quest]; Enduring Cape (14763, -33.93 DPS) [world_drop]; Tigerstrike Mantle (13108, -96.24 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | 24.0 stamina points (497.87 DPS) | yes | Raptor Hide Harness (4455, -11.34 DPS) [crafted]; Green Whelp Armor (7375, -43.96 DPS) [crafted]; Blackened Defias Armor (10399, -65.77 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.6 stamina points (261.26 DPS) | yes | Cultist's Armguards (270032, -3.10 DPS) [quest]; Black Wolf Bracers (3230, -14.67 DPS) [dungeon]; Glowing Leather Bracers (2017, -58.86 DPS) [world] |
| hands | Brawler Gloves (720) | World drop [world_drop] | 14.8 stamina points (307.36 DPS) | yes | Wolfclaw Gloves (1978, -1.83 DPS) [dungeon]; Gloves of Insight (9698, -27.16 DPS) [quest]; Emblazoned Gloves (6397, -31.36 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 18.1 stamina points (375.28 DPS) | yes | Prowler's Leather Belt (252459, -52.86 DPS) [crafted]; Stalker's Leather Belt (252521, -65.68 DPS) [crafted]; Skulker's Leather Belt (252520, -140.87 DPS, sim-verified) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 21.4 stamina points (443.25 DPS) | yes | Barbaric Leggings (5963, -48.12 DPS) [crafted]; Petrolspill Leggings (9509, -48.49 DPS) [dungeon]; Brawler's Leather Legguards (252516, -172.27 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 20.1 stamina points (417.09 DPS) | yes | Harbinger Boots (7754, -25.57 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -88.55 DPS) [quest]; Emblazoned Boots (4051, -115.01 DPS) [world_drop] |
| finger1 | Theramore Signet (271911) | Theramore Expeditionary Force [rep] | 8.7 stamina points (179.93 DPS) | yes | Plains Ring (2039, -14.03 DPS) [dungeon]; Defias Renegade Ring (1076, -35.19 DPS) [dungeon]; Protector's Band (19517, -47.80 DPS) [rep] |
| finger2 | Militant's Band (271916) (or Plains Ring (2039)) | Theramore Expeditionary Force [rep] | 8.0 stamina points (165.90 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Defias Renegade Ring (1076, -21.16 DPS) [dungeon]; Protector's Band (19517, -33.77 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (+132.8 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slaghammer (1976, -68.07 DPS, sim-verified) [dungeon]; Lorekeeper's Staff (212580, -99.84 DPS) [vendor]; Spine of the Devourer (273040, -103.35 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Brawler Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Theramore Signet; finger2: Militant's Band; trinket1: Talisman of Arathor; trinket2: Rune of Duty; main_hand: Viscous Hammer

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 96.9. Weights run: 4.1s. Verify run: 5.3s. 607 eligible items had no known source.

Race sweep: night-elf 368.5 ± 0.8, high-order-skyborne 358.2 ± 0.8 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.028, armor=0.178 ± 0.005, defense=not significant (0.486 ± 0.123) per rating point (1 rating = 1%, 0.486 per %), dodge=0.279 ± 0.038 per rating point (12 rating = 1%, 3.343 per %), strength=0.118 ± 0.000, agility=0.300 ± 0.023, attack_power=0.059 ± 0.000, feral_attack_power=0.059 ± 0.000, hit=not significant (0.103 ± 0.027) per rating point (10 rating = 1%, 1.031 per %), crit=0.057 ± 0.007 per rating point (14 rating = 1%, 0.801 per %), expertise=3.768 ± 0.326

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 33.5 stamina points (1026.63 DPS) | yes | Warden's Wizard Hat (14604, -27.90 DPS) [world_drop]; Brawler's Leather Helm (252512, -76.18 DPS) [crafted]; Nightscape Headband (8176, -80.86 DPS) [crafted] |
| neck | Shriveled Heart (9243) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 12.4 stamina points (380.87 DPS) | yes | Gazlowe's Charm (13088, -28.79 DPS) [dungeon]; Seal of the Expedition (271919, -43.28 DPS) [rep]; Sentinel's Medallion (19540, -64.77 DPS) [rep] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 34.3 stamina points (1053.27 DPS) | yes | Flintrock Shoulders (7755, -138.56 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -267.28 DPS) [dungeon]; Nightscape Shoulders (8192, -328.20 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heraldic Cloak (8120, -19.75 DPS) [dungeon]; Silky Spider Cape (10776, -25.88 DPS) [dungeon]; Well Oiled Cloak (12254, -31.35 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 38.2 stamina points (1173.05 DPS) | yes | Warden's Wraps (14601, -37.98 DPS) [world_drop]; Spirewind Fetter (9406, -197.55 DPS) [dungeon]; Insignia Chestguard (4057, -210.74 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 18.1 stamina points (554.47 DPS) | yes | Branded Leather Bracers (19508, -33.31 DPS) [dungeon]; Barbaric Bracers (18948, -61.97 DPS) [crafted]; Cultist's Armguards (270032, -80.88 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -0.01 DPS) [crafted]; Prowler's Leather Gloves (252524, -5.60 DPS) [crafted]; Stalker's Leather Gloves (252526, -304.54 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 25.4 stamina points (779.97 DPS) | yes | Kolkar Hunter's Belt (6788, -74.61 DPS) [quest]; Warden's Leather Belt (252460, -82.23 DPS) [crafted]; Scorpashi Sash (14652, -86.69 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 34.4 stamina points (1056.87 DPS) | yes | Panther Hunter Leggings (4108, -102.17 DPS) [quest]; Imperial Leather Pants (4062, -108.85 DPS) [dungeon]; Warden's Woolies (14605, -125.25 DPS) [world_drop] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 30.3 stamina points (930.55 DPS) | yes | Skulker's Leather Shoes (252531, -85.94 DPS) [crafted]; Prowler's Leather Shoes (252465, -108.28 DPS) [crafted]; Worn Running Boots (9398, -237.29 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (368.28 DPS) | yes | Talvash's Gold Ring (9538, -16.61 DPS) [quest]; Militant's Band (271915, -30.69 DPS) [rep]; Field Researcher's Loop (281634, -89.01 DPS) [quest] |
| finger2 | Theramore Signet (271910) | Theramore Expeditionary Force [rep] | 11.7 stamina points (359.32 DPS) | yes | Talvash's Gold Ring (9538, -7.64 DPS) [quest]; Militant's Band (271915, -21.72 DPS) [rep]; Field Researcher's Loop (281634, -80.05 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (+251.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -30.69 DPS) [world]; Enriched Seal (271905, -30.69 DPS) [rep] |
| trinket2 | Talisman of Arathor (21118) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -30.69 DPS) [world]; Volunteer's Lucky Seal (271904, -30.69 DPS) [rep]; Enriched Seal (271905, -30.69 DPS) [rep] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -138.45 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -450.69 DPS) [dungeon]; Huge Stone Club (1523, -529.37 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Theramore Signet; trinket1: Rune of Duty; trinket2: Talisman of Arathor; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 152.7. Weights run: 4.1s. Verify run: 7.9s. 817 eligible items had no known source.

Race sweep: night-elf 595.3 ± 1.1, high-order-skyborne 585.1 ± 1.1 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.022, armor=0.146 ± 0.003, defense=0.918 ± 0.157 per rating point (1 rating = 1%, 0.918 per %), dodge=0.357 ± 0.046 per rating point (12 rating = 1%, 4.280 per %), strength=0.096 ± 0.000, agility=0.340 ± 0.029, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.144 ± 0.033 per rating point (10 rating = 1%, 1.438 per %), crit=0.079 ± 0.007 per rating point (14 rating = 1%, 1.110 per %), expertise=4.661 ± 0.393

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 35.4 stamina points (1902.12 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Sprightring Helm (17776, -9.25 DPS) [quest]; Scorpashi Skullcap (14658, -100.18 DPS) [world_drop] |
| neck | Lifeblood Amulet (9641) (or Master Sergeant's Insignia (18444)) | Zul'Farrak: Antu'sul [dungeon] | 14.0 stamina points (751.57 DPS) | yes | Master Sergeant's Insignia (18444, +0.00 DPS) [vendor]; Seal of the Expedition (271918, -53.68 DPS) [rep]; Shriveled Heart (9243, -79.38 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.4 stamina points (1687.24 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -11.69 DPS) [vendor]; Phytoskin Spaulders (17749, -57.90 DPS) [dungeon]; Warden's Leather Shoulders (252536, -105.19 DPS) [crafted] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightfall Drape (12465, -61.55 DPS) [dungeon]; Sergeant's Cape (18441, -77.27 DPS) [vendor]; Emberscale Cape (11311, -92.99 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) | PvP rank 6 · Knight · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Warbear Harness (15064, +0.00 DPS) [crafted]; Grizzled Pelt (22274, +0.00 DPS) [quest] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.4 stamina points (1364.14 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -109.41 DPS) [vendor]; Warden's Leather Bracers (252541, -232.44 DPS) [crafted]; Serpentskin Bracers (8257, -262.93 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -177.25 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -242.72 DPS) [crafted]; Wicked Leather Gauntlets (15083, -254.81 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | PvP rank 6 · Knight · Alliance [vendor] | 38.0 stamina points (2041.77 DPS) | yes | Knight's Crackling Leather Leggings (220864, -314.67 DPS) [vendor]; Basilisk Hide Pants (1718, -340.98 DPS) [world_drop]; Windscale Sarong (10842, -358.53 DPS) [world] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 5 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (805.26 DPS) | yes | Militant's Band (271914, -53.68 DPS) [rep]; Talvash's Gold Ring (9538, -120.54 DPS) [quest]; Underworld Band (1980, -161.05 DPS) [world_drop] |
| finger2 | Theramore Signet (271909) | Theramore Expeditionary Force [rep] | 14.7 stamina points (790.12 DPS) | yes | Militant's Band (271914, -38.55 DPS) [rep]; Talvash's Gold Ring (9538, -105.40 DPS) [quest]; Underworld Band (1980, -145.91 DPS) [world_drop] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Guardian Talisman (1490, -53.68 DPS) [quest]; Rune of Perfection (21565, -107.37 DPS) [rep] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -53.68 DPS) [quest]; Rune of Perfection (21565, -107.37 DPS) [rep]; Rune of Duty (21567, -107.37 DPS) [rep] |
| main_hand | Warden Staff (943) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -213.94 DPS) [world_drop]; Blanchard's Stout (13046, -285.90 DPS) [world_drop]; Ironshod Bludgeon (9408, -301.75 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Lifeblood Amulet; back: Graverot Cape; chest: Knight's Leather Armor; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Ring of Saviors; finger2: Theramore Signet; trinket1: Enriched Seal; trinket2: Talisman of Arathor; main_hand: Warden Staff

No-known-source sample (15 of 817, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 237.7. Weights run: 7.9s. Verify run: 42.2s. 1988 eligible items had no known source.

Race sweep: high-order-skyborne 963.8 ± 0.9, night-elf 956.4 ± 0.9 (threat_per_second); high-order-skyborne beats the guide race night-elf by 7.4 (0.78%) ± 1.3, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.153 ± 0.000, defense=1.519 ± 0.192 per rating point (1 rating = 1%, 1.519 per %), dodge=0.723 ± 0.043 per rating point (12 rating = 1%, 8.674 per %), strength=0.117 ± 0.000, agility=0.598 ± 0.027, attack_power=0.058 ± 0.000, feral_attack_power=0.058 ± 0.000, hit=0.150 ± 0.032 per rating point (10 rating = 1%, 1.503 per %), crit=0.137 ± 0.015 per rating point (14 rating = 1%, 1.913 per %), expertise=4.683 ± 0.293

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 66.7 stamina points (10063.73 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Master Sergeant's Insignia (18443) | PvP rank 4 · Master Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Verek's Collar (11755, +0.00 DPS) [dungeon]; Medallion of Grand Marshal Morris (13091, +0.00 DPS) [world_drop]; Talisman of Evasion (13177, +0.00 DPS) [dungeon] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, +0.00 DPS) [dungeon]; Bloodmoon Cloak (12967, +0.00 DPS) [dungeon]; Shadow Prowler's Cloak (22269, +0.00 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 63.2 stamina points (9533.99 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | 36.8 stamina points (5549.93 DPS) | yes | Forest Stalker's Bracers (19587, -451.97 DPS) [rep]; Malefic Bracers (18700, -476.46 DPS) [dungeon]; Yeti Hide Bracers (19113, -817.35 DPS) [world] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -137.13 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, -137.13 DPS) [pvp] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nagmara's Whipping Belt (11866, +0.00 DPS) [quest]; Vosh'gajin's Strand (13962, +0.00 DPS) [quest]; Shifter's Belt (272396, +0.00 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 89.7 stamina points (13531.22 DPS) | yes | Marshal's Dragonhide Leggings (231691, -232.09 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -1398.17 DPS) [pvp]; Marshal's Dragonhide Pants (231697, -2165.86 DPS) [vendor] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 52.0 stamina points (7850.63 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Innervating Band (18701, +0.00 DPS) [dungeon]; Ring of Awareness (272409, +0.00 DPS) [vendor] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Ring of Awareness (272409, +0.00 DPS) [vendor]; Innervating Band (18701, -79.17 DPS) [dungeon] |
| trinket1 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; The Lion Horn of Stormwind (14557, -784.56 DPS) [world_drop]; Enriched Seal (271903, -1108.88 DPS) [rep] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Lion Horn of Stormwind (14557, +0.00 DPS) [world_drop]; Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Enriched Seal (271903, -301.82 DPS) [rep] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Warbear Helm; neck: Master Sergeant's Insignia; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Theramore Gloves; waist: Hivethrasher's Girdle; legs: Sentinel's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Stormpike Insignia Rank 6; trinket2: Talisman of Arathor; main_hand: The Unstoppable Force; ranged: Idol of Brutality

No-known-source sample (15 of 1988, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 497.9. Weights run: 9.0s. Verify run: 46.0s. 1988 eligible items had no known source.

Race sweep: high-order-skyborne 1940.3 ± 2.1, night-elf 1938.7 ± 2.2 (threat_per_second); high-order-skyborne beats the guide race night-elf by 1.5 (0.08%) ± 3.0, inside error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.015, armor=0.213 ± 0.002, defense=1.269 ± 0.184 per rating point (1 rating = 1%, 1.269 per %), dodge=0.576 ± 0.055 per rating point (12 rating = 1%, 6.910 per %), strength=0.103 ± 0.000, agility=0.546 ± 0.035, attack_power=0.047 ± 0.000, feral_attack_power=0.047 ± 0.000, hit=0.212 ± 0.042 per rating point (10 rating = 1%, 2.118 per %), crit=0.098 ± 0.017 per rating point (14 rating = 1%, 1.369 per %), expertise=7.826 ± 0.463

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 74.9 stamina points (9958.49 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 19.7 stamina points (2617.22 DPS) | yes | Talisman of Evasion (13177, -106.01 DPS) [dungeon]; Verek's Collar (11755, -180.43 DPS) [dungeon]; Sentinel's Medallion (19538, -199.49 DPS) [rep] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 59.8 stamina points (7949.43 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, +0.00 DPS) [dungeon]; Bloodmoon Cloak (12967, +0.00 DPS) [dungeon]; Shadow Prowler's Cloak (22269, +0.00 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 71.7 stamina points (9525.80 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | 39.5 stamina points (5244.36 DPS) | yes | Forest Stalker's Bracers (19587, -212.65 DPS) [rep]; Malefic Bracers (18700, -303.43 DPS) [dungeon]; Yeti Hide Bracers (19113, -523.78 DPS) [world] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 47.7 stamina points (6344.86 DPS) | yes | Nagmara's Whipping Belt (11866, -332.71 DPS) [quest]; Hivethrasher's Girdle (275614, -408.77 DPS) [crafted]; Frostbite Girdle (14502, -442.52 DPS) [dungeon] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 99.4 stamina points (13206.46 DPS) | yes | Marshal's Dragonhide Leggings (231691, -125.69 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -1021.63 DPS) [pvp]; Marshal's Dragonhide Pants (231697, -1769.90 DPS) [vendor] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 56.0 stamina points (7441.85 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Cutthroat's Signet (272408, +0.00 DPS) [vendor]; Ring of Awareness (272409, +0.00 DPS) [vendor] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Ring of Awareness (272409, +0.00 DPS) [vendor]; Cutthroat's Signet (272408, -65.30 DPS) [vendor] |
| trinket1 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; The Lion Horn of Stormwind (14557, -245.08 DPS) [world_drop]; Enriched Seal (271903, -265.78 DPS) [rep] |
| trinket2 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; The Lion Horn of Stormwind (14557, -487.00 DPS) [world_drop]; Enriched Seal (271903, -507.70 DPS) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Warbear Helm; neck: Medallion of Grand Marshal Morris; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Talisman of Arathor; trinket2: Stormpike Insignia Rank 6; main_hand: Headmaster's Charge; ranged: Idol of Brutality

No-known-source sample (15 of 1988, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 23.0. Weights run: 4.5s. Verify run: 5.2s. 247 eligible items had no known source.

Race sweep: windshaper-skyborne 93.5 ± 0.2, tauren 93.5 ± 0.2 (threat_per_second); windshaper-skyborne beats the guide race tauren by 0.0 (0.03%) ± 0.3, inside error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.091 ± 0.000, defense=0.222 ± 0.043 per rating point (1 rating = 1%, 0.222 per %), dodge=0.107 ± 0.014 per rating point (12 rating = 1%, 1.289 per %), strength=0.097 ± 0.000, agility=0.160 ± 0.009, attack_power=0.049 ± 0.000, feral_attack_power=0.049 ± 0.000, hit=0.049 ± 0.008 per rating point (10 rating = 1%, 0.489 per %), crit=0.015 ± 0.001 per rating point (14 rating = 1%, 0.213 per %), expertise=1.150 ± 0.097

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 18.2 stamina points (329.61 DPS) | yes | Defender's Leather Hood (252447, -11.38 DPS, sim-verified) [crafted]; Trapper's Leather Hood (252505, -23.19 DPS) [crafted]; Stormrider's Leather Hood (252506, -23.19 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (60.06 DPS) | yes | Scholarly Pendant (277203, -5.80 DPS) [quest]; Scout's Medallion (20442, -6.49 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.8 stamina points (178.07 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -51.25 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -61.12 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -4.84 DPS) [world_drop]; Firebane Cloak (12979, -19.74 DPS) [world_drop]; Battle Healer's Cloak (20427, -36.17 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 19.0 stamina points (343.64 DPS) | yes | Gloomshroud Armor (1489, -33.62 DPS, sim-verified) [dungeon]; Defender's Leather Armor (252434, -91.75 DPS) [crafted]; Trapper's Leather Armor (252491, -92.50 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.4 stamina points (151.55 DPS) | yes | Spare Part Bindings (279875, -27.48 DPS) [quest]; Drakewing Bands (12999, -29.57 DPS) [world_drop]; Wolf Bracers (4794, -42.67 DPS) [vendor] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.5 stamina points (171.49 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.65 DPS) [crafted]; Defender's Leather Gloves (252496, -4.54 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.8 stamina points (212.55 DPS) | yes | Belt of the Fang (10412, -31.35 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -49.32 DPS) [crafted]; Guardsman Belt (3429, -50.97 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 15.0 stamina points (271.44 DPS) | yes | Brawler's Leather Pants (252500, -26.97 DPS) [crafted]; Leggings of the Fang (10410, -35.34 DPS) [dungeon]; Smelting Pants (5199, -91.48 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.3 stamina points (276.51 DPS) | yes | Surfer Shoes (276274, -49.83 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -49.83 DPS) [dungeon]; Trapper's Leather Boots (252440, -70.82 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (90.44 DPS) | yes | Slain Baron's Signet (279867, -32.15 DPS) [quest]; Ring of the Moon (12052, -34.41 DPS) [world_drop]; Legionnaire's Band (20429, -35.61 DPS) [rep] |
| finger2 | Ring of Scorn (3235) | A Husband's Revenge [quest] | 4.0 stamina points (72.35 DPS) | yes | Slain Baron's Signet (279867, -14.06 DPS) [quest]; Ring of the Moon (12052, -16.32 DPS) [world_drop]; Legionnaire's Band (20429, -17.53 DPS) [rep] |
| trinket1 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.9 stamina points (377.23 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.98 DPS) [dungeon]; Twisted Chanter's Staff (890, -16.42 DPS) [world_drop]; Staff of the Blessed Seer (2271, -21.09 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Ring of Scorn; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 8350 The 1 Ring; 14389 Durability Shoulderpads

### Band 30 (tauren, 0000000000000000-51100332020102100000-0000000000000000)

Set DPS (verified): 65.3. Weights run: 3.9s. Verify run: 7.4s. 425 eligible items had no known source.

Race sweep: windshaper-skyborne 248.4 ± 0.4, tauren 245.9 ± 0.5 (threat_per_second); windshaper-skyborne beats the guide race tauren by 2.5 (1.01%) ± 0.6, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.107 ± 0.000, defense=0.424 ± 0.088 per rating point (1 rating = 1%, 0.424 per %), dodge=0.178 ± 0.024 per rating point (12 rating = 1%, 2.131 per %), strength=0.150 ± 0.000, agility=0.245 ± 0.015, attack_power=0.075 ± 0.000, feral_attack_power=0.075 ± 0.000, hit=not significant (0.062 ± 0.018) per rating point (10 rating = 1%, 0.619 per %), crit=0.045 ± 0.004 per rating point (14 rating = 1%, 0.632 per %), expertise=2.483 ± 0.213

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.4 stamina points (506.33 DPS) | yes | Defender's Leather Helm (252455, -13.37 DPS) [crafted]; Trapper's Leather Helm (252513, -50.79 DPS) [crafted]; Stormrider's Leather Helm (252514, -50.79 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.6 stamina points (199.11 DPS) | yes | Senior Sergeant's Insignia (15200, -12.47 DPS) [vendor]; Darkspear Warding Pendant (272075, -33.21 DPS) [vendor]; Scout's Medallion (19537, -54.79 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 20.0 stamina points (414.56 DPS) | yes | Azure Gustwoven Spaulders (277051, -65.61 DPS) [crafted]; Barbaric Shoulders (5964, -80.91 DPS) [crafted]; Cloudy Gustwoven Spaulders (277043, -85.73 DPS, sim-verified) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enduring Cape (14763, -33.93 DPS) [world_drop]; Grimsteel Cape (4643, -37.96 DPS) [quest]; Tigerstrike Mantle (13108, -121.23 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | 24.0 stamina points (497.87 DPS) | yes | Raptor Hide Harness (4455, -11.34 DPS) [crafted]; Green Whelp Armor (7375, -43.96 DPS) [crafted]; Blackened Defias Armor (10399, -65.77 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.6 stamina points (261.26 DPS) | yes | Cultist's Armguards (270032, -3.10 DPS) [quest]; Black Wolf Bracers (3230, -14.67 DPS) [dungeon]; Glowing Leather Bracers (2017, -58.86 DPS) [world] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 15.5 stamina points (322.39 DPS) | yes | Brawler Gloves (720, -15.03 DPS) [world_drop]; Wolfclaw Gloves (1978, -16.87 DPS) [dungeon]; Emblazoned Gloves (6397, -46.39 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 18.1 stamina points (375.28 DPS) | yes | Larcenist's Belt (279846, -50.52 DPS) [quest]; Prowler's Leather Belt (252459, -52.86 DPS) [crafted]; Skulker's Leather Belt (252520, -150.66 DPS, sim-verified) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 21.4 stamina points (443.25 DPS) | yes | Barbaric Leggings (5963, -48.12 DPS) [crafted]; Petrolspill Leggings (9509, -48.49 DPS) [dungeon]; Brawler's Leather Legguards (252516, -171.25 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 20.1 stamina points (417.09 DPS) | yes | Harbinger Boots (7754, -25.57 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -88.55 DPS) [quest]; Gravewalker Boots (251963, -89.36 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.7 stamina points (179.93 DPS) | yes | Plains Ring (2039, -14.03 DPS) [dungeon]; Darkspear Signet (272071, -14.03 DPS) [vendor]; Defias Renegade Ring (1076, -35.19 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.5 stamina points (175.25 DPS) | yes | Plains Ring (2039, -9.35 DPS) [dungeon]; Darkspear Signet (272071, -9.35 DPS) [vendor]; Defias Renegade Ring (1076, -30.51 DPS) [dungeon] |
| trinket1 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-verified (+91.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slaghammer (1976, -75.75 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (212584, -99.84 DPS) [vendor]; Spine of the Devourer (273040, -103.35 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket2: Defiler's Talisman; main_hand: Viscous Hammer

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 93.3. Weights run: 4.1s. Verify run: 5.0s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 353.9 ± 0.8, tauren 353.8 ± 0.8 (threat_per_second); windshaper-skyborne beats the guide race tauren by 0.1 (0.03%) ± 1.1, inside error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.028, armor=0.178 ± 0.005, defense=not significant (0.486 ± 0.123) per rating point (1 rating = 1%, 0.486 per %), dodge=0.279 ± 0.038 per rating point (12 rating = 1%, 3.343 per %), strength=0.118 ± 0.000, agility=0.300 ± 0.023, attack_power=0.059 ± 0.000, feral_attack_power=0.059 ± 0.000, hit=not significant (0.103 ± 0.027) per rating point (10 rating = 1%, 1.031 per %), crit=0.057 ± 0.007 per rating point (14 rating = 1%, 0.801 per %), expertise=3.768 ± 0.326

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 33.5 stamina points (1026.63 DPS) | yes | Warden's Wizard Hat (14604, -27.90 DPS) [world_drop]; Brawler's Leather Helm (252512, -76.18 DPS) [crafted]; Nightscape Headband (8176, -80.86 DPS) [crafted] |
| neck | Shriveled Heart (9243) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 12.4 stamina points (380.87 DPS) | yes | Dragon's Blood Necklace (10711, -12.59 DPS) [quest]; Gazlowe's Charm (13088, -28.79 DPS) [dungeon]; Darkspear Warding Pendant (272074, -43.28 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 34.3 stamina points (1053.27 DPS) | yes | Flintrock Shoulders (7755, -190.47 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -267.28 DPS) [dungeon]; Nightscape Shoulders (8192, -328.20 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Heraldic Cloak (8120, -19.75 DPS) [dungeon]; Silky Spider Cape (10776, -25.88 DPS) [dungeon]; Well Oiled Cloak (12254, -31.35 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 38.2 stamina points (1173.05 DPS) | yes | Warden's Wraps (14601, -37.98 DPS) [world_drop]; Spirewind Fetter (9406, -197.55 DPS) [dungeon]; Insignia Chestguard (4057, -210.74 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 18.1 stamina points (554.47 DPS) | yes | Branded Leather Bracers (19508, -33.31 DPS) [dungeon]; Barbaric Bracers (18948, -61.97 DPS) [crafted]; Cultist's Armguards (270032, -80.88 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -0.01 DPS) [crafted]; Prowler's Leather Gloves (252524, -5.60 DPS) [crafted]; Stalker's Leather Gloves (252526, -412.90 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 25.4 stamina points (779.97 DPS) | yes | Kolkar Hunter's Belt (6788, -74.61 DPS) [quest]; Warden's Leather Belt (252460, -82.23 DPS) [crafted]; Scorpashi Sash (14652, -86.69 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 34.4 stamina points (1056.87 DPS) | yes | Imperial Leather Pants (4062, -108.85 DPS) [dungeon]; Warden's Woolies (14605, -125.25 DPS) [world_drop]; Panther Hunter Leggings (4108, -136.06 DPS, sim-verified) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 30.3 stamina points (930.55 DPS) | yes | Skulker's Leather Shoes (252531, -85.94 DPS) [crafted]; Prowler's Leather Shoes (252465, -108.28 DPS) [crafted]; Worn Running Boots (9398, -298.60 DPS, sim-verified) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (368.28 DPS) | yes | Nogg's Gold Ring (9588, -16.61 DPS) [quest]; Darkspear Signet (272070, -30.69 DPS) [vendor]; Field Researcher's Loop (281634, -89.01 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.7 stamina points (359.32 DPS) | yes | Nogg's Gold Ring (9588, -7.64 DPS) [quest]; Darkspear Signet (272070, -21.72 DPS) [vendor]; Field Researcher's Loop (281634, -80.05 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (+259.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -30.69 DPS) [world]; Darkspear Voodoo Seal (272059, -30.69 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21116) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -30.69 DPS) [world]; Darkspear Voodoo Seal (272059, -30.69 DPS) [vendor]; Relentless Raider's Seal (272062, -30.69 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -92.35 DPS) [dungeon]; Mograine's Might (7723, -235.34 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -306.34 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Insurgent's Band; trinket1: Rune of Duty; trinket2: Defiler's Talisman; main_hand: Cragwood Maul

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 148.7. Weights run: 4.1s. Verify run: 11.0s. 816 eligible items had no known source.

Race sweep: windshaper-skyborne 585.1 ± 1.1, tauren 578.3 ± 1.1 (threat_per_second); windshaper-skyborne beats the guide race tauren by 6.8 (1.18%) ± 1.5, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.036, armor=0.159 ± 0.004, defense=0.856 ± 0.171 per rating point (1 rating = 1%, 0.856 per %), dodge=0.399 ± 0.052 per rating point (12 rating = 1%, 4.785 per %), strength=0.110 ± 0.000, agility=0.378 ± 0.032, attack_power=0.055 ± 0.000, feral_attack_power=0.055 ± 0.000, hit=0.157 ± 0.037 per rating point (10 rating = 1%, 1.566 per %), crit=0.079 ± 0.008 per rating point (14 rating = 1%, 1.110 per %), expertise=5.213 ± 0.441

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.3 stamina points (1906.49 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Sprightring Helm (17776, -12.14 DPS) [quest]; Scorpashi Skullcap (14658, -90.61 DPS) [world_drop] |
| neck | Lifeblood Amulet (9641) | Zul'Farrak: Antu'sul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Senior Sergeant's Insignia (18428, +0.00 DPS) [vendor]; Woven Ivy Necklace (19159, +0.00 DPS) [quest]; Darkspear Warding Pendant (272073, -51.16 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 32.9 stamina points (1685.10 DPS) | yes | Blood Guard's Leather Shoulders (220853, -2.31 DPS) [vendor]; Phytoskin Spaulders (17749, -30.74 DPS) [dungeon]; Warden's Leather Shoulders (252536, -89.43 DPS) [crafted] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightfall Drape (12465, -59.30 DPS) [dungeon]; Sergeant's Cloak (16341, -75.58 DPS) [vendor]; Emberscale Cape (11311, -91.86 DPS) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Warbear Harness (15064, +0.00 DPS) [crafted]; Grizzled Pelt (22274, +0.00 DPS) [quest] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 26.5 stamina points (1353.35 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -116.13 DPS) [vendor]; Warden's Leather Bracers (252541, -216.27 DPS) [crafted]; Serpentskin Bracers (8257, -260.39 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -159.02 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -227.56 DPS) [crafted]; Wicked Leather Gauntlets (15083, -252.05 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 6 · Stone Guard · Horde [vendor] | 40.0 stamina points (2045.90 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -306.61 DPS) [vendor]; Basilisk Hide Pants (1718, -310.28 DPS) [world_drop]; Keeper's Woolies (14668, -343.17 DPS) [world_drop] |
| feet | First Sergeant's Leather Boots (220861) | PvP rank 5 · First Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (767.38 DPS) | yes | Darkspear Signet (272069, -51.16 DPS) [vendor]; Nogg's Gold Ring (9588, -124.48 DPS) [quest]; Underworld Band (1980, -153.48 DPS) [world_drop] |
| finger2 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.8 stamina points (758.57 DPS) | yes | Darkspear Signet (272069, -42.34 DPS) [vendor]; Nogg's Gold Ring (9588, -115.67 DPS) [quest]; Underworld Band (1980, -144.66 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -51.16 DPS) [quest]; Rune of Perfection (21565, -102.32 DPS) [rep] |
| trinket2 | Defiler's Talisman (21115) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -51.16 DPS) [quest]; Rune of Perfection (21565, -102.32 DPS) [rep]; Rune of Duty (21567, -102.32 DPS) [rep] |
| main_hand | Warden Staff (943) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cragwood Maul (11265, -140.81 DPS) [quest]; Taran Icebreaker (2915, -154.66 DPS) [world_drop]; Blanchard's Stout (13046, -241.54 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Lifeblood Amulet; back: Graverot Cape; chest: Stone Guard's Leather Armor; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: First Sergeant's Leather Boots; finger1: Ring of Saviors; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Warden Staff

No-known-source sample (15 of 816, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 229.7. Weights run: 7.9s. Verify run: 43.8s. 1992 eligible items had no known source.

Race sweep: windshaper-skyborne 921.8 ± 0.9, tauren 917.4 ± 0.9 (threat_per_second); windshaper-skyborne beats the guide race tauren by 4.5 (0.49%) ± 1.2, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.046, armor=0.181 ± 0.006, defense=0.874 ± 0.158 per rating point (1 rating = 1%, 0.874 per %), dodge=0.400 ± 0.043 per rating point (12 rating = 1%, 4.795 per %), strength=0.096 ± 0.000, agility=0.374 ± 0.026, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.196 ± 0.034 per rating point (10 rating = 1%, 1.964 per %), crit=0.068 ± 0.012 per rating point (14 rating = 1%, 0.950 per %), expertise=5.316 ± 0.384

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 67.5 stamina points (5017.23 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor] |
| neck | Senior Sergeant's Insignia (16335) | PvP rank 4 · Senior Sergeant · Horde [vendor] | 17.0 stamina points (1263.31 DPS) | yes | Frostwolf Legionnaire's Pendant (19095, -84.47 DPS) [rep]; Medallion of Grand Marshal Morris (13091, -93.76 DPS) [world_drop]; Scout's Medallion (19534, -103.35 DPS) [rep] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 52.0 stamina points (3864.12 DPS) | yes | Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, +0.00 DPS) [dungeon]; Bloodmoon Cloak (12967, +0.00 DPS) [dungeon]; Sergeant's Cloak (18461, +0.00 DPS) [vendor] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 63.4 stamina points (4713.43 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | 34.3 stamina points (2552.60 DPS) | yes | Malefic Bracers (18700, -97.10 DPS) [dungeon]; Forest Stalker's Bracers (19587, -193.21 DPS) [rep]; Yeti Hide Bracers (19113, -202.77 DPS) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 40.5 stamina points (3007.86 DPS) | yes | Shifter's Belt (272396, -22.72 DPS) [vendor]; Nagmara's Whipping Belt (11866, -127.14 DPS) [quest]; Frostbite Girdle (14502, -134.82 DPS) [dungeon] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 69.8 stamina points (5185.58 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Pants (231680, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 46.7 stamina points (3467.97 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Insurgent's Band (272064, +0.00 DPS) [vendor]; Ring of Awareness (272409, +0.00 DPS) [vendor] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Ring of Awareness (272409, +0.00 DPS) [vendor]; Insurgent's Band (272064, -65.55 DPS) [vendor] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, -101.05 DPS) [vendor]; Darkspear Voodoo Seal (272061, -101.05 DPS) [vendor]; Frostwolf Insignia Rank 6 (17909, -131.58 DPS) [quest] |
| trinket2 | Defiler's Talisman (20072) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, -148.63 DPS) [vendor]; Darkspear Voodoo Seal (272061, -148.63 DPS) [vendor]; Frostwolf Insignia Rank 6 (17909, -179.16 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Warbear Helm; neck: Senior Sergeant's Insignia; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Raider Gloves; waist: Hivethrasher's Girdle; legs: Outrider's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Defender's Grip Stabilizer; trinket2: Defiler's Talisman; main_hand: The Unstoppable Force

No-known-source sample (15 of 1992, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 552.6. Weights run: 9.0s. Verify run: 46.6s. 1992 eligible items had no known source.

Race sweep: windshaper-skyborne 2167.0 ± 2.4, tauren 2149.1 ± 2.4 (threat_per_second); windshaper-skyborne beats the guide race tauren by 17.9 (0.83%) ± 3.4, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.015, armor=0.213 ± 0.002, defense=1.269 ± 0.184 per rating point (1 rating = 1%, 1.269 per %), dodge=0.576 ± 0.055 per rating point (12 rating = 1%, 6.910 per %), strength=0.103 ± 0.000, agility=0.546 ± 0.035, attack_power=0.047 ± 0.000, feral_attack_power=0.047 ± 0.000, hit=0.212 ± 0.042 per rating point (10 rating = 1%, 2.118 per %), crit=0.098 ± 0.017 per rating point (14 rating = 1%, 1.369 per %), expertise=7.826 ± 0.463

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 74.9 stamina points (9958.49 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor] |
| neck | Talisman of Evasion (13177) | Blackrock Spire: War Master Voone [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Medallion of Grand Marshal Morris (13091, +0.00 DPS) [world_drop]; Verek's Collar (11755, -74.42 DPS) [dungeon]; Scout's Medallion (19534, -93.48 DPS) [rep] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 59.8 stamina points (7949.43 DPS) | yes | Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, +0.00 DPS) [dungeon]; Bloodmoon Cloak (12967, +0.00 DPS) [dungeon]; Shadow Prowler's Cloak (22269, +0.00 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 71.7 stamina points (9525.80 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magistrate's Cuffs (18726, +0.00 DPS) [dungeon]; Malefic Bracers (18700, -90.79 DPS) [dungeon]; Yeti Hide Bracers (19113, -311.14 DPS) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 47.7 stamina points (6344.86 DPS) | yes | Nagmara's Whipping Belt (11866, -332.71 DPS) [quest]; Hivethrasher's Girdle (275614, -408.77 DPS) [crafted]; Frostbite Girdle (14502, -442.52 DPS) [dungeon] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 80.6 stamina points (10708.39 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Pants (231680, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 56.0 stamina points (7441.85 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Band of Flesh (13373, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, +0.00 DPS) [vendor] |
| finger2 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, -38.86 DPS) [world_drop]; Band of Flesh (13373, -530.58 DPS) [dungeon]; Cutthroat's Signet (272408, -595.88 DPS) [vendor] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Talisman (20072, -598.42 DPS) [rep]; The Lion Horn of Stormwind (14557, -843.50 DPS) [world_drop] |
| trinket2 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Talisman (20072, -241.93 DPS) [rep]; The Lion Horn of Stormwind (14557, -487.00 DPS) [world_drop] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Warbear Helm; neck: Talisman of Evasion; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Outrider's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Ring of Awareness; trinket1: Defender's Grip Stabilizer; trinket2: Frostwolf Insignia Rank 6; main_hand: The Unstoppable Force; ranged: Idol of Brutality

No-known-source sample (15 of 1992, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

