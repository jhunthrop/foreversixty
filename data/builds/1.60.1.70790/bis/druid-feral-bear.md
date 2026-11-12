# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 22.2. Weights run: 4.4s. Verify run: 4.1s. 251 eligible items had no known source.

Race sweep: night-elf 88.6 ± 0.2, high-order-skyborne 88.1 ± 0.2 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.085 ± 0.001, defense=0.218 ± 0.045 per rating point (1 rating = 1%, 0.218 per %), dodge=0.108 ± 0.014 per rating point (12 rating = 1%, 1.297 per %), strength=0.098 ± 0.000, agility=0.158 ± 0.009, attack_power=0.049 ± 0.000, feral_attack_power=0.049 ± 0.000, hit=0.043 ± 0.008 per rating point (10 rating = 1%, 0.425 per %), crit=0.016 ± 0.001 per rating point (14 rating = 1%, 0.218 per %), expertise=1.115 ± 0.097

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.7 stamina points (316.85 DPS) | yes | Defender's Leather Hood (252447, -8.69 DPS) [crafted]; Trapper's Leather Hood (252505, -22.65 DPS) [crafted]; Stormrider's Leather Hood (252506, -22.65 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.7 stamina points (102.16 DPS) | yes | Erudite's Amulet (277204, -48.36 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -48.46 DPS) [quest]; Sentinel's Medallion (20444, -49.37 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.4 stamina points (168.61 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -51.37 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -56.09 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -5.06 DPS) [world_drop]; Firebane Cloak (12979, -19.41 DPS) [world_drop]; Caretaker's Cape (20428, -35.80 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.4 stamina points (329.37 DPS) | yes | Gloomshroud Armor (1489, -34.71 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -69.27 DPS) [quest]; Bloody Apron (6226, -86.72 DPS) [dungeon] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (145.58 DPS) | yes | Drakewing Bands (12999, -31.81 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -42.37 DPS) [vendor]; Sanguine Cuffs (14375, -48.22 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (163.26 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.52 DPS) [crafted]; Defender's Leather Gloves (252496, -4.35 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.4 stamina points (204.10 DPS) | yes | Belt of the Fang (10412, -35.21 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.41 DPS) [crafted]; Guardsman Belt (3429, -49.93 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.5 stamina points (258.73 DPS) | yes | Brawler's Leather Pants (252500, -26.87 DPS) [crafted]; Leggings of the Fang (10410, -34.68 DPS) [dungeon]; Smelting Pants (5199, -79.91 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (269.36 DPS) | yes | Surfer Shoes (276274, -49.67 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -52.51 DPS) [dungeon]; Trapper's Leather Boots (252440, -73.24 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (89.49 DPS) | yes | Slain Baron's Signet (279867, -31.89 DPS) [quest]; Ring of the Moon (12052, -34.05 DPS) [world_drop]; Protector's Band (20439, -35.39 DPS) [rep] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 4.0 stamina points (71.60 DPS) | yes | Slain Baron's Signet (279867, -13.99 DPS) [quest]; Ring of the Moon (12052, -16.15 DPS) [world_drop]; Protector's Band (20439, -17.49 DPS) [rep] |
| trinket1 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.9 stamina points (373.35 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.93 DPS) [dungeon]; Twisted Chanter's Staff (890, -16.25 DPS) [world_drop]; Staff of the Blessed Seer (2271, -20.87 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Sustaining Ring; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 251, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-51100332020102100000-0000000000000000)

Set DPS (verified): 65.5. Weights run: 3.7s. Verify run: 7.2s. 424 eligible items had no known source.

Race sweep: night-elf 243.9 ± 0.5, high-order-skyborne 239.2 ± 0.4 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.103 ± 0.003, defense=0.436 ± 0.094 per rating point (1 rating = 1%, 0.436 per %), dodge=0.174 ± 0.024 per rating point (12 rating = 1%, 2.093 per %), strength=0.151 ± 0.000, agility=0.240 ± 0.016, attack_power=0.076 ± 0.000, feral_attack_power=0.076 ± 0.000, hit=not significant (0.044 ± 0.019) per rating point (10 rating = 1%, 0.443 per %), crit=0.048 ± 0.005 per rating point (14 rating = 1%, 0.668 per %), expertise=2.522 ± 0.224

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.1 stamina points (494.44 DPS) | yes | Defender's Leather Helm (252455, -12.15 DPS) [crafted]; Trapper's Leather Helm (252513, -49.35 DPS) [crafted]; Stormrider's Leather Helm (252514, -49.35 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.6 stamina points (197.14 DPS) | yes | Master Sergeant's Insignia (18442, -12.40 DPS) [vendor]; Seal of the Expedition (271920, -32.93 DPS) [rep]; Sentinel's Medallion (19541, -55.03 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 19.7 stamina points (404.85 DPS) | yes | Azure Gustwoven Spaulders (277051, -65.35 DPS) [crafted]; Barbaric Shoulders (5964, -80.24 DPS) [crafted]; Cloudy Gustwoven Spaulders (277043, -81.79 DPS, sim-verified) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 3 · Sergeant · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mourning Shawl (6751, -29.01 DPS) [quest]; Enduring Cape (14763, -33.47 DPS) [world_drop]; Tigerstrike Mantle (13108, -112.51 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | 23.6 stamina points (485.16 DPS) | yes | Raptor Hide Harness (4455, -10.36 DPS) [crafted]; Green Whelp Armor (7375, -42.92 DPS) [crafted]; Blackened Defias Armor (10399, -63.83 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.4 stamina points (255.03 DPS) | yes | Cultist's Armguards (270032, -2.48 DPS) [quest]; Black Wolf Bracers (3230, -13.90 DPS) [dungeon]; Glowing Leather Bracers (2017, -57.65 DPS) [world] |
| hands | Brawler Gloves (720) | World drop [world_drop] | 14.6 stamina points (300.05 DPS) | yes | Wolfclaw Gloves (1978, -2.34 DPS) [dungeon]; Gloves of Insight (9698, -26.92 DPS) [quest]; Emblazoned Gloves (6397, -31.26 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 17.9 stamina points (366.64 DPS) | yes | Prowler's Leather Belt (252459, -52.28 DPS) [crafted]; Stalker's Leather Belt (252521, -65.38 DPS) [crafted]; Skulker's Leather Belt (252520, -147.12 DPS, sim-verified) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 21.0 stamina points (431.97 DPS) | yes | Barbaric Leggings (5963, -47.42 DPS) [crafted]; Petrolspill Leggings (9509, -48.73 DPS) [dungeon]; Brawler's Leather Legguards (252516, -188.65 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.9 stamina points (407.59 DPS) | yes | Harbinger Boots (7754, -24.94 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -84.88 DPS) [quest]; Emblazoned Boots (4051, -113.29 DPS) [world_drop] |
| finger1 | Theramore Signet (271911) | Theramore Expeditionary Force [rep] | 8.7 stamina points (178.16 DPS) | yes | Plains Ring (2039, -13.95 DPS) [dungeon]; Defias Renegade Ring (1076, -35.26 DPS) [dungeon]; Protector's Band (19517, -47.84 DPS) [rep] |
| finger2 | Militant's Band (271916) (or Plains Ring (2039)) | Theramore Expeditionary Force [rep] | 8.0 stamina points (164.21 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Defias Renegade Ring (1076, -21.31 DPS) [dungeon]; Protector's Band (19517, -33.89 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (+132.1 DPS vs the runner-up, not corroborated against the finished set) | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slaghammer (1976, -64.44 DPS, sim-verified) [dungeon]; Lorekeeper's Staff (212580, -99.64 DPS) [vendor]; Spine of the Devourer (273040, -103.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Brawler Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Theramore Signet; finger2: Militant's Band; trinket1: Talisman of Arathor; trinket2: Rune of Duty; main_hand: Viscous Hammer

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 91.2. Weights run: 4.1s. Verify run: 6.1s. 607 eligible items had no known source.

Race sweep: night-elf 340.5 ± 0.8, high-order-skyborne 330.6 ± 0.8 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.032, armor=0.131 ± 0.006, defense=0.686 ± 0.107 per rating point (1 rating = 1%, 0.686 per %), dodge=0.194 ± 0.030 per rating point (12 rating = 1%, 2.328 per %), strength=0.080 ± 0.000, agility=0.228 ± 0.019, attack_power=0.040 ± 0.000, feral_attack_power=0.040 ± 0.000, hit=not significant (0.081 ± 0.023) per rating point (10 rating = 1%, 0.808 per %), crit=0.054 ± 0.006 per rating point (14 rating = 1%, 0.757 per %), expertise=3.026 ± 0.281

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 28.4 stamina points (987.72 DPS) | yes | Warden's Wizard Hat (14604, -49.80 DPS) [world_drop]; Brawler's Leather Helm (252512, -72.53 DPS) [crafted]; Nightscape Headband (8176, -94.27 DPS) [crafted] |
| neck | Shriveled Heart (9243) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 12.6 stamina points (438.29 DPS) | yes | Gazlowe's Charm (13088, -44.52 DPS) [dungeon]; Seal of the Expedition (271919, -55.65 DPS) [rep]; Sentinel's Medallion (19540, -107.57 DPS) [rep] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 29.2 stamina points (1017.21 DPS) | yes | Flintrock Shoulders (7755, -139.33 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -260.80 DPS) [dungeon]; Nightscape Shoulders (8192, -363.23 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -22.58 DPS) [dungeon]; Heraldic Cloak (8120, -25.65 DPS) [dungeon]; Well Oiled Cloak (12254, -27.15 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 32.4 stamina points (1125.91 DPS) | yes | Warden's Wraps (14601, -65.21 DPS) [world_drop]; Robes of the Lich (10762, -147.01 DPS) [dungeon]; Insignia Chestguard (4057, -203.23 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 15.2 stamina points (528.59 DPS) | yes | Branded Leather Bracers (19508, -40.82 DPS) [dungeon]; Barbaric Bracers (18948, -62.35 DPS) [crafted]; Cultist's Armguards (270032, -70.20 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razzeric's Racing Grips (6727, -23.81 DPS) [quest]; Skulker's Leather Gloves (252525, -25.06 DPS) [crafted]; Stalker's Leather Gloves (252526, -390.05 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 21.1 stamina points (733.79 DPS) | yes | Kolkar Hunter's Belt (6788, -60.39 DPS) [quest]; Warden's Leather Belt (252460, -68.68 DPS) [crafted]; Scorpashi Sash (14652, -89.26 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.6 stamina points (960.95 DPS) | yes | Panther Hunter Leggings (4108, -60.52 DPS) [quest]; Warden's Woolies (14605, -73.05 DPS) [world_drop]; Imperial Leather Pants (4062, -178.12 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 25.2 stamina points (877.86 DPS) | yes | Skulker's Leather Shoes (252531, -107.42 DPS) [crafted]; Gnomebot Operating Boots (9450, -113.29 DPS) [dungeon]; Worn Running Boots (9398, -312.47 DPS, sim-verified) [dungeon] |
| finger1 | Talvash's Gold Ring (9538) | Gnome Improvement [quest] | 12.1 stamina points (419.39 DPS) | yes | Theramore Signet (271910, -20.06 DPS) [rep]; Militant's Band (271915, -36.76 DPS) [rep]; Field Researcher's Loop (281634, -120.39 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (417.42 DPS) | yes | Theramore Signet (271910, -18.08 DPS) [rep]; Militant's Band (271915, -34.78 DPS) [rep]; Field Researcher's Loop (281634, -118.42 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -34.78 DPS) [world]; Enriched Seal (271905, -34.78 DPS) [rep] |
| trinket2 | Talisman of Arathor (21118) | The League of Arathor [rep] | sim-verified (+139.0 DPS vs the runner-up, not corroborated against the finished set) | yes | Arena Master (18706, -34.78 DPS) [world]; Enriched Seal (271905, -34.78 DPS) [rep]; Ankh of Life (1713, -138.95 DPS, sim-verified) [world_drop] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -121.25 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -380.06 DPS) [dungeon]; Huge Stone Club (1523, -474.11 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Talvash's Gold Ring; finger2: Underworld Band; trinket1: Rune of Duty; trinket2: Talisman of Arathor; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 140.6. Weights run: 3.9s. Verify run: 8.6s. 815 eligible items had no known source.

Race sweep: night-elf 533.9 ± 1.1, high-order-skyborne 524.9 ± 1.1 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.014, armor=0.140 ± 0.007, defense=0.958 ± 0.172 per rating point (1 rating = 1%, 0.958 per %), dodge=0.352 ± 0.045 per rating point (12 rating = 1%, 4.223 per %), strength=0.096 ± 0.000, agility=0.338 ± 0.028, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.158 ± 0.035 per rating point (10 rating = 1%, 1.577 per %), crit=0.074 ± 0.007 per rating point (14 rating = 1%, 1.039 per %), expertise=4.460 ± 0.404

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 34.7 stamina points (1815.18 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Sprightring Helm (17776, -6.69 DPS) [quest]; Scorpashi Skullcap (14658, -95.61 DPS) [world_drop] |
| neck | Lifeblood Amulet (9641) (or Master Sergeant's Insignia (18444)) | Zul'Farrak: Antu'sul [dungeon] | 14.0 stamina points (733.07 DPS) | yes | Master Sergeant's Insignia (18444, +0.00 DPS) [vendor]; Seal of the Expedition (271918, -52.36 DPS) [rep]; Shriveled Heart (9243, -77.52 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 30.8 stamina points (1612.29 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -15.26 DPS) [vendor]; Phytoskin Spaulders (17749, -62.96 DPS) [dungeon]; Warden's Leather Shoulders (252536, -105.74 DPS) [crafted] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightfall Drape (12465, -59.68 DPS) [dungeon]; Sergeant's Cape (18441, -74.33 DPS) [vendor]; Emberscale Cape (11311, -88.97 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) | PvP rank 6 · Knight · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Warbear Harness (15064, +0.00 DPS) [crafted]; Grizzled Pelt (22274, +0.00 DPS) [quest] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.0 stamina points (1307.80 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -106.15 DPS) [vendor]; Warden's Leather Bracers (252541, -226.59 DPS) [crafted]; Serpentskin Bracers (8257, -254.09 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -173.35 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -236.65 DPS) [crafted]; Wicked Leather Gauntlets (15083, -244.48 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted] |
| legs | Knight's Leather Pants (220858) | PvP rank 6 · Knight · Alliance [vendor] | 37.2 stamina points (1948.78 DPS) | yes | Knight's Crackling Leather Leggings (220864, -307.10 DPS) [vendor]; Basilisk Hide Pants (1718, -331.02 DPS) [world_drop]; Windscale Sarong (10842, -346.87 DPS) [world] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 5 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (785.44 DPS) | yes | Militant's Band (271914, -52.36 DPS) [rep]; Talvash's Gold Ring (9538, -111.27 DPS) [quest]; Underworld Band (1980, -157.09 DPS) [world_drop] |
| finger2 | Theramore Signet (271909) | Theramore Expeditionary Force [rep] | 14.7 stamina points (770.81 DPS) | yes | Militant's Band (271914, -37.74 DPS) [rep]; Talvash's Gold Ring (9538, -96.64 DPS) [quest]; Underworld Band (1980, -142.46 DPS) [world_drop] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Guardian Talisman (1490, -52.36 DPS) [quest]; Rune of Perfection (21565, -104.72 DPS) [rep] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -52.36 DPS) [quest]; Rune of Perfection (21565, -104.72 DPS) [rep]; Rune of Duty (21567, -104.72 DPS) [rep] |
| main_hand | Warden Staff (943) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -229.25 DPS) [world_drop]; Blanchard's Stout (13046, -299.91 DPS) [world_drop]; Ironshod Bludgeon (9408, -316.40 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Lifeblood Amulet; back: Graverot Cape; chest: Knight's Leather Armor; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Ring of Saviors; finger2: Theramore Signet; trinket1: Enriched Seal; trinket2: Talisman of Arathor; main_hand: Warden Staff

No-known-source sample (15 of 815, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 218.0. Weights run: 8.2s. Verify run: 41.5s. 1986 eligible items had no known source.

Race sweep: night-elf 857.9 ± 0.9, high-order-skyborne 857.4 ± 0.9 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.064, armor=0.161 ± 0.016, defense=not significant (0.837 ± 0.233) per rating point (1 rating = 1%, 0.837 per %), dodge=0.418 ± 0.043 per rating point (12 rating = 1%, 5.013 per %), strength=0.096 ± 0.000, agility=0.371 ± 0.028, attack_power=0.048 ± 0.000, feral_attack_power=0.048 ± 0.000, hit=0.201 ± 0.037 per rating point (10 rating = 1%, 2.013 per %), crit=0.071 ± 0.015 per rating point (14 rating = 1%, 0.991 per %), expertise=5.096 ± 0.418

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 64.5 stamina points (4741.50 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Master Sergeant's Insignia (18443) | PvP rank 4 · Master Sergeant · Alliance [vendor] | 17.0 stamina points (1250.51 DPS) | yes | Stormpike Soldier's Pendant (19097, -83.82 DPS) [rep]; Sentinel's Medallion (19538, -105.31 DPS) [rep]; Medallion of Grand Marshal Morris (13091, -120.09 DPS) [world_drop] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 49.0 stamina points (3605.17 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Field Marshal's Dragonhide Pauldrons (231705, +0.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, +0.00 DPS) [dungeon]; Bloodmoon Cloak (12967, +0.00 DPS) [dungeon]; Sergeant's Cape (16342, +0.00 DPS) [vendor] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 60.2 stamina points (4431.39 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | 32.6 stamina points (2401.19 DPS) | yes | Malefic Bracers (18700, -82.23 DPS) [dungeon]; Yeti Hide Bracers (19113, -189.23 DPS) [world]; Forest Stalker's Bracers (19587, -195.19 DPS) [rep] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 38.6 stamina points (2840.36 DPS) | yes | Shifter's Belt (272396, -42.28 DPS) [vendor]; Frostbite Girdle (14502, -142.42 DPS) [dungeon]; Nagmara's Whipping Belt (11866, -143.35 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 83.5 stamina points (6139.00 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -104.97 DPS) [pvp]; Marshal's Dragonhide Pants (231697, -546.98 DPS) [vendor] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 44.4 stamina points (3266.02 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, +0.00 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Theramore Signet (271908, +0.00 DPS) [rep]; Ring of Awareness (272409, +0.00 DPS) [vendor] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Ring of Awareness (272409, +0.00 DPS) [vendor]; Theramore Signet (271908, -64.34 DPS) [rep] |
| trinket1 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Enriched Seal (271903, -1.89 DPS) [rep]; Volunteer's Lucky Seal (271906, -1.89 DPS) [rep] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Grip Stabilizer (272440, -82.56 DPS) [vendor]; Enriched Seal (271903, -147.12 DPS) [rep]; Volunteer's Lucky Seal (271906, -147.12 DPS) [rep] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Warbear Helm; neck: Master Sergeant's Insignia; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Theramore Gloves; waist: Hivethrasher's Girdle; legs: Sentinel's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Stormpike Insignia Rank 6; trinket2: Talisman of Arathor; main_hand: The Unstoppable Force; ranged: Idol of Brutality

No-known-source sample (15 of 1986, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 476.6. Weights run: 9.0s. Verify run: 43.6s. 1986 eligible items had no known source.

Race sweep: night-elf 1820.5 ± 2.1, high-order-skyborne 1818.6 ± 2.0 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.170 ± 0.016, defense=1.501 ± 0.292 per rating point (1 rating = 1%, 1.501 per %), dodge=0.984 ± 0.057 per rating point (12 rating = 1%, 11.803 per %), strength=0.122 ± 0.000, agility=0.879 ± 0.038, attack_power=0.055 ± 0.000, feral_attack_power=0.055 ± 0.000, hit=0.207 ± 0.045 per rating point (10 rating = 1%, 2.069 per %), crit=0.181 ± 0.021 per rating point (14 rating = 1%, 2.533 per %), expertise=6.991 ± 0.429

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 73.3 stamina points (15303.82 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Helm (231695, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Verek's Collar (11755, +0.00 DPS) [dungeon]; Talisman of Evasion (13177, +0.00 DPS) [dungeon]; Sentinel's Medallion (19538, +0.00 DPS) [rep] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, +0.00 DPS) [rep] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, +0.00 DPS) [dungeon]; Bloodmoon Cloak (12967, +0.00 DPS) [dungeon]; Shadow Prowler's Cloak (22269, +0.00 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 71.0 stamina points (14820.04 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Armor (231696, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Malefic Bracers (18700, -292.80 DPS) [dungeon] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -1748.04 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, -1748.04 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 51.5 stamina points (10742.85 DPS) | yes | Vosh'gajin's Strand (13962, -1249.41 DPS) [quest]; Frostbite Girdle (14502, -1308.77 DPS) [dungeon]; Nagmara's Whipping Belt (11866, -1641.14 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 101.9 stamina points (21278.76 DPS) | yes | Marshal's Dragonhide Leggings (231691, -1069.52 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -3829.82 DPS) [pvp]; Marshal's Dragonhide Pants (231697, -4762.20 DPS) [vendor] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 60.8 stamina points (12686.33 DPS) | yes | Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Boots (231698, -1247.87 DPS) [pvp]; Marshal's Dragonhide Greaves (231704, -1247.87 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Cutthroat's Signet (272408, +0.00 DPS) [vendor]; Ring of Awareness (272409, +0.00 DPS) [vendor] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Cutthroat's Signet (272408, +0.00 DPS) [vendor]; Ring of Awareness (272409, +0.00 DPS) [vendor] |
| trinket1 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Mark of Tyranny (13966, -41.19 DPS) [quest]; Arena Grand Master (19024, -41.19 DPS) [quest] |
| trinket2 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defender's Grip Stabilizer (272440, -853.28 DPS) [vendor]; Mark of Tyranny (13966, -2463.72 DPS) [quest]; Arena Grand Master (19024, -2463.72 DPS) [quest] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Warbear Helm; neck: Medallion of Grand Marshal Morris; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Talisman of Arathor; trinket2: Stormpike Insignia Rank 6; main_hand: Headmaster's Charge; ranged: Idol of Brutality

No-known-source sample (15 of 1986, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 22.3. Weights run: 4.4s. Verify run: 4.1s. 247 eligible items had no known source.

Race sweep: windshaper-skyborne 89.1 ± 0.2, tauren 89.0 ± 0.2 (threat_per_second); windshaper-skyborne beats the guide race tauren by 0.1 (0.07%) ± 0.3, inside error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.085 ± 0.001, defense=0.218 ± 0.045 per rating point (1 rating = 1%, 0.218 per %), dodge=0.108 ± 0.014 per rating point (12 rating = 1%, 1.297 per %), strength=0.098 ± 0.000, agility=0.158 ± 0.009, attack_power=0.049 ± 0.000, feral_attack_power=0.049 ± 0.000, hit=0.043 ± 0.008 per rating point (10 rating = 1%, 0.425 per %), crit=0.016 ± 0.001 per rating point (14 rating = 1%, 0.218 per %), expertise=1.115 ± 0.097

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.7 stamina points (316.85 DPS) | yes | Defender's Leather Hood (252447, -8.69 DPS) [crafted]; Trapper's Leather Hood (252505, -22.65 DPS) [crafted]; Stormrider's Leather Hood (252506, -22.65 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (59.36 DPS) | yes | Scholarly Pendant (277203, -5.66 DPS) [quest]; Scout's Medallion (20442, -6.57 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.4 stamina points (168.61 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -51.37 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -56.09 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -5.06 DPS) [world_drop]; Firebane Cloak (12979, -19.41 DPS) [world_drop]; Battle Healer's Cloak (20427, -35.80 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.4 stamina points (329.37 DPS) | yes | Gloomshroud Armor (1489, -36.25 DPS, sim-verified) [dungeon]; Bloody Apron (6226, -86.72 DPS) [dungeon]; Defender's Leather Armor (252434, -90.32 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (145.58 DPS) | yes | Spare Part Bindings (279875, -28.02 DPS, sim-verified) [quest]; Drakewing Bands (12999, -29.73 DPS) [world_drop]; Savannah Bracers (15453, -41.86 DPS) [quest] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (163.26 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.52 DPS) [crafted]; Defender's Leather Gloves (252496, -4.35 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.4 stamina points (204.10 DPS) | yes | Belt of the Fang (10412, -32.58 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.41 DPS) [crafted]; Guardsman Belt (3429, -49.93 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.5 stamina points (258.73 DPS) | yes | Brawler's Leather Pants (252500, -26.87 DPS) [crafted]; Leggings of the Fang (10410, -34.68 DPS) [dungeon]; Smelting Pants (5199, -86.01 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (269.36 DPS) | yes | Surfer Shoes (276274, -47.82 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -52.51 DPS) [dungeon]; Trapper's Leather Boots (252440, -73.24 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (89.49 DPS) | yes | Slain Baron's Signet (279867, -31.89 DPS) [quest]; Ring of the Moon (12052, -34.05 DPS) [world_drop]; Legionnaire's Band (20429, -35.39 DPS) [rep] |
| finger2 | Ring of Scorn (3235) | A Husband's Revenge [quest] | 4.0 stamina points (71.60 DPS) | yes | Slain Baron's Signet (279867, -13.99 DPS) [quest]; Ring of the Moon (12052, -16.15 DPS) [world_drop]; Legionnaire's Band (20429, -17.49 DPS) [rep] |
| trinket1 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.9 stamina points (373.35 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.93 DPS) [dungeon]; Twisted Chanter's Staff (890, -16.25 DPS) [world_drop]; Staff of the Blessed Seer (2271, -20.87 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Ring of Scorn; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 8350 The 1 Ring; 14389 Durability Shoulderpads

### Band 30 (tauren, 0000000000000000-51100332020102100000-0000000000000000)

Set DPS (verified): 63.7. Weights run: 3.7s. Verify run: 6.3s. 425 eligible items had no known source.

Race sweep: windshaper-skyborne 238.2 ± 0.4, tauren 236.5 ± 0.4 (threat_per_second); windshaper-skyborne beats the guide race tauren by 1.7 (0.71%) ± 0.6, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.103 ± 0.003, defense=0.436 ± 0.094 per rating point (1 rating = 1%, 0.436 per %), dodge=0.174 ± 0.024 per rating point (12 rating = 1%, 2.093 per %), strength=0.151 ± 0.000, agility=0.240 ± 0.016, attack_power=0.076 ± 0.000, feral_attack_power=0.076 ± 0.000, hit=not significant (0.044 ± 0.019) per rating point (10 rating = 1%, 0.443 per %), crit=0.048 ± 0.005 per rating point (14 rating = 1%, 0.668 per %), expertise=2.522 ± 0.224

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 24.1 stamina points (494.44 DPS) | yes | Defender's Leather Helm (252455, -12.15 DPS) [crafted]; Trapper's Leather Helm (252513, -49.35 DPS) [crafted]; Stormrider's Leather Helm (252514, -49.35 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.6 stamina points (197.14 DPS) | yes | Senior Sergeant's Insignia (15200, -12.40 DPS) [vendor]; Darkspear Warding Pendant (272075, -32.93 DPS) [vendor]; Scout's Medallion (19537, -55.03 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 19.7 stamina points (404.85 DPS) | yes | Tanned Shoulderpads (270023, -55.40 DPS, sim-verified) [quest]; Cloudy Gustwoven Spaulders (277043, -65.35 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -65.35 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enduring Cape (14763, -33.47 DPS) [world_drop]; Grimsteel Cape (4643, -37.55 DPS) [quest]; Tigerstrike Mantle (13108, -98.21 DPS, sim-verified) [world_drop] |
| chest | Spirewind Fetter (9406) | Uldaman [dungeon] | 23.6 stamina points (485.16 DPS) | yes | Raptor Hide Harness (4455, -10.36 DPS) [crafted]; Green Whelp Armor (7375, -42.92 DPS) [crafted]; Blackened Defias Armor (10399, -63.83 DPS) [dungeon] |
| wrist | Barbaric Bracers (18948) | Leatherworking [crafted] | 12.4 stamina points (255.03 DPS) | yes | Cultist's Armguards (270032, -2.48 DPS) [quest]; Black Wolf Bracers (3230, -13.90 DPS) [dungeon]; Glowing Leather Bracers (2017, -57.65 DPS) [world] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 15.3 stamina points (314.04 DPS) | yes | Brawler Gloves (720, -13.99 DPS) [world_drop]; Wolfclaw Gloves (1978, -16.33 DPS) [dungeon]; Emblazoned Gloves (6397, -45.25 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 17.9 stamina points (366.64 DPS) | yes | Larcenist's Belt (279846, -49.74 DPS) [quest]; Prowler's Leather Belt (252459, -52.28 DPS) [crafted]; Skulker's Leather Belt (252520, -145.90 DPS, sim-verified) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 21.0 stamina points (431.97 DPS) | yes | Barbaric Leggings (5963, -47.42 DPS) [crafted]; Petrolspill Leggings (9509, -48.73 DPS) [dungeon]; Brawler's Leather Legguards (252516, -172.06 DPS, sim-verified) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 19.9 stamina points (407.59 DPS) | yes | Harbinger Boots (7754, -24.94 DPS) [dungeon]; Nat Pagle's Extreme Anglin' Boots (19969, -84.88 DPS) [quest]; Gravewalker Boots (251963, -88.57 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.7 stamina points (178.16 DPS) | yes | Plains Ring (2039, -13.95 DPS) [dungeon]; Darkspear Signet (272071, -13.95 DPS) [vendor]; Defias Renegade Ring (1076, -35.26 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.5 stamina points (173.51 DPS) | yes | Plains Ring (2039, -9.30 DPS) [dungeon]; Darkspear Signet (272071, -9.30 DPS) [vendor]; Defias Renegade Ring (1076, -30.61 DPS) [dungeon] |
| trinket1 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-verified (+90.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | Defiler's Talisman (21120) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Viscous Hammer (13045) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slaghammer (1976, -66.75 DPS, sim-verified) [dungeon]; Advisor's Gnarled Staff (212584, -99.64 DPS) [vendor]; Spine of the Devourer (273040, -103.48 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: River Pride Choker; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Spirewind Fetter; wrist: Barbaric Bracers; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket2: Defiler's Talisman; main_hand: Viscous Hammer

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 87.7. Weights run: 4.1s. Verify run: 5.9s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 326.9 ± 0.8, tauren 326.5 ± 0.8 (threat_per_second); windshaper-skyborne beats the guide race tauren by 0.4 (0.13%) ± 1.1, inside error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.032, armor=0.131 ± 0.006, defense=0.686 ± 0.107 per rating point (1 rating = 1%, 0.686 per %), dodge=0.194 ± 0.030 per rating point (12 rating = 1%, 2.328 per %), strength=0.080 ± 0.000, agility=0.228 ± 0.019, attack_power=0.040 ± 0.000, feral_attack_power=0.040 ± 0.000, hit=not significant (0.081 ± 0.023) per rating point (10 rating = 1%, 0.808 per %), crit=0.054 ± 0.006 per rating point (14 rating = 1%, 0.757 per %), expertise=3.026 ± 0.281

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 28.4 stamina points (987.72 DPS) | yes | Warden's Wizard Hat (14604, -49.80 DPS) [world_drop]; Brawler's Leather Helm (252512, -72.53 DPS) [crafted]; Nightscape Headband (8176, -94.27 DPS) [crafted] |
| neck | Shriveled Heart (9243) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 12.6 stamina points (438.29 DPS) | yes | Dragon's Blood Necklace (10711, -20.87 DPS) [quest]; Gazlowe's Charm (13088, -44.52 DPS) [dungeon]; Darkspear Warding Pendant (272074, -55.65 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 29.2 stamina points (1017.21 DPS) | yes | Flintrock Shoulders (7755, -182.39 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -260.80 DPS) [dungeon]; Tanned Shoulderpads (270023, -362.34 DPS) [quest] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -22.58 DPS) [dungeon]; Heraldic Cloak (8120, -25.65 DPS) [dungeon]; Well Oiled Cloak (12254, -27.15 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 32.4 stamina points (1125.91 DPS) | yes | Warden's Wraps (14601, -65.21 DPS) [world_drop]; Robes of the Lich (10762, -147.01 DPS) [dungeon]; Insignia Chestguard (4057, -203.23 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 15.2 stamina points (528.59 DPS) | yes | Branded Leather Bracers (19508, -40.82 DPS) [dungeon]; Barbaric Bracers (18948, -62.35 DPS) [crafted]; Cultist's Armguards (270032, -70.20 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Razzeric's Racing Grips (6727, -23.81 DPS) [quest]; Skulker's Leather Gloves (252525, -25.06 DPS) [crafted]; Stalker's Leather Gloves (252526, -251.68 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 21.1 stamina points (733.79 DPS) | yes | Kolkar Hunter's Belt (6788, -60.39 DPS) [quest]; Warden's Leather Belt (252460, -68.68 DPS) [crafted]; Scorpashi Sash (14652, -89.26 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 27.6 stamina points (960.95 DPS) | yes | Imperial Leather Pants (4062, -59.35 DPS) [dungeon]; Panther Hunter Leggings (4108, -60.52 DPS) [quest]; Warden's Woolies (14605, -73.05 DPS) [world_drop] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 25.2 stamina points (877.86 DPS) | yes | Skulker's Leather Shoes (252531, -107.42 DPS) [crafted]; Gnomebot Operating Boots (9450, -113.29 DPS) [dungeon]; Worn Running Boots (9398, -243.62 DPS, sim-verified) [dungeon] |
| finger1 | Nogg's Gold Ring (9588) | Nogg's Ring Redo [quest] | 12.1 stamina points (419.39 DPS) | yes | Insurgent's Band (272066, -20.06 DPS) [vendor]; Darkspear Signet (272070, -36.76 DPS) [vendor]; Field Researcher's Loop (281634, -120.39 DPS) [quest] |
| finger2 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (417.42 DPS) | yes | Insurgent's Band (272066, -18.08 DPS) [vendor]; Darkspear Signet (272070, -34.78 DPS) [vendor]; Field Researcher's Loop (281634, -118.42 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (+273.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -34.78 DPS) [world]; Darkspear Voodoo Seal (272059, -34.78 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21116) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -34.78 DPS) [world]; Darkspear Voodoo Seal (272059, -34.78 DPS) [vendor]; Relentless Raider's Seal (272062, -34.78 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -115.76 DPS) [dungeon]; Mograine's Might (7723, -292.83 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -302.66 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Nogg's Gold Ring; finger2: Underworld Band; trinket1: Rune of Duty; trinket2: Defiler's Talisman; main_hand: Cragwood Maul

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 136.8. Weights run: 3.9s. Verify run: 8.8s. 814 eligible items had no known source.

Race sweep: windshaper-skyborne 524.9 ± 1.1, tauren 517.4 ± 1.0 (threat_per_second); windshaper-skyborne beats the guide race tauren by 7.5 (1.46%) ± 1.5, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.046, armor=0.157 ± 0.008, defense=0.854 ± 0.181 per rating point (1 rating = 1%, 0.854 per %), dodge=0.401 ± 0.054 per rating point (12 rating = 1%, 4.814 per %), strength=0.113 ± 0.000, agility=0.390 ± 0.034, attack_power=0.057 ± 0.000, feral_attack_power=0.057 ± 0.000, hit=0.180 ± 0.039 per rating point (10 rating = 1%, 1.801 per %), crit=0.085 ± 0.009 per rating point (14 rating = 1%, 1.189 per %), expertise=5.287 ± 0.463

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 37.0 stamina points (1837.31 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Sprightring Helm (17776, -7.65 DPS) [quest]; Scorpashi Skullcap (14658, -81.55 DPS) [world_drop] |
| neck | Lifeblood Amulet (9641) | Zul'Farrak: Antu'sul [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Senior Sergeant's Insignia (18428, +0.00 DPS) [vendor]; Woven Ivy Necklace (19159, +0.00 DPS) [quest]; Darkspear Warding Pendant (272073, -49.59 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 32.8 stamina points (1626.03 DPS) | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS) [vendor]; Phytoskin Spaulders (17749, -27.33 DPS) [dungeon]; Warden's Leather Shoulders (252536, -86.63 DPS) [crafted] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nightfall Drape (12465, -57.36 DPS) [dungeon]; Sergeant's Cloak (16341, -72.90 DPS) [vendor]; Emberscale Cape (11311, -88.44 DPS) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Warbear Harness (15064, +0.00 DPS) [crafted]; Grizzled Pelt (22274, +0.00 DPS) [quest] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 26.4 stamina points (1307.51 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -116.02 DPS) [vendor]; Warden's Leather Bracers (252541, -208.62 DPS) [crafted]; Serpentskin Bracers (8257, -253.40 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gauntlets (252548, -151.28 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -219.86 DPS) [crafted]; Wicked Leather Gauntlets (15083, -242.35 DPS) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Girdle of Beastial Fury (11686, +0.00 DPS) [dungeon]; Prowler's Leather Waistguard (252473, +0.00 DPS) [crafted]; Skulker's Leather Waistguard (252474, +0.00 DPS) [crafted] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 6 · Stone Guard · Horde [vendor] | 40.0 stamina points (1983.74 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -298.56 DPS) [vendor]; Basilisk Hide Pants (1718, -303.08 DPS) [world_drop]; Keeper's Woolies (14668, -336.97 DPS) [world_drop] |
| feet | First Sergeant's Leather Boots (220861) | PvP rank 5 · First Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Shadefiend Boots (11675, +0.00 DPS) [dungeon]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (743.90 DPS) | yes | Darkspear Signet (272069, -49.59 DPS) [vendor]; Nogg's Gold Ring (9588, -120.86 DPS) [quest]; Underworld Band (1980, -148.78 DPS) [world_drop] |
| finger2 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.9 stamina points (736.47 DPS) | yes | Darkspear Signet (272069, -42.16 DPS) [vendor]; Nogg's Gold Ring (9588, -113.43 DPS) [quest]; Underworld Band (1980, -141.35 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -49.59 DPS) [quest]; Rune of Perfection (21565, -99.19 DPS) [rep] |
| trinket2 | Defiler's Talisman (21115) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -49.59 DPS) [quest]; Rune of Perfection (21565, -99.19 DPS) [rep]; Rune of Duty (21567, -99.19 DPS) [rep] |
| main_hand | Warden Staff (943) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Cragwood Maul (11265, -145.52 DPS) [quest]; Taran Icebreaker (2915, -145.90 DPS) [world_drop]; Blanchard's Stout (13046, -233.73 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Lifeblood Amulet; back: Graverot Cape; chest: Stone Guard's Leather Armor; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: First Sergeant's Leather Boots; finger1: Ring of Saviors; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Warden Staff

No-known-source sample (15 of 814, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 209.0. Weights run: 8.2s. Verify run: 43.9s. 1990 eligible items had no known source.

Race sweep: windshaper-skyborne 821.6 ± 0.8, tauren 815.6 ± 0.8 (threat_per_second); windshaper-skyborne beats the guide race tauren by 6.1 (0.74%) ± 1.2, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.149 ± 0.013, defense=1.444 ± 0.252 per rating point (1 rating = 1%, 1.444 per %), dodge=0.767 ± 0.046 per rating point (12 rating = 1%, 9.208 per %), strength=0.119 ± 0.000, agility=0.640 ± 0.029, attack_power=0.060 ± 0.000, feral_attack_power=0.060 ± 0.000, hit=0.161 ± 0.035 per rating point (10 rating = 1%, 1.606 per %), crit=0.149 ± 0.016 per rating point (14 rating = 1%, 2.087 per %), expertise=4.938 ± 0.325

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 66.7 stamina points (9737.33 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor] |
| neck | Senior Sergeant's Insignia (16335) | PvP rank 4 · Senior Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Verek's Collar (11755, +0.00 DPS) [dungeon]; Medallion of Grand Marshal Morris (13091, +0.00 DPS) [world_drop]; Talisman of Evasion (13177, +0.00 DPS) [dungeon] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, +0.00 DPS) [dungeon]; Bloodmoon Cloak (12967, +0.00 DPS) [dungeon]; Shadow Prowler's Cloak (22269, +0.00 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 63.3 stamina points (9243.17 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | 36.0 stamina points (5249.08 DPS) | yes | Forest Stalker's Bracers (19587, -248.39 DPS) [rep]; Malefic Bracers (18700, -368.94 DPS) [dungeon]; Yeti Hide Bracers (19113, -665.16 DPS) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; General's Dragonhide Gauntlets (231676, -307.69 DPS) [vendor]; General's Dragonhide Gloves (231677, -307.69 DPS) [pvp] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Nagmara's Whipping Belt (11866, +0.00 DPS) [quest]; Vosh'gajin's Strand (13962, +0.00 DPS) [quest]; Shifter's Belt (272396, +0.00 DPS) [vendor] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 72.7 stamina points (10616.01 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Pants (231680, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 52.4 stamina points (7649.29 DPS) | yes | General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; General's Dragonhide Greaves (231671, -102.12 DPS) [vendor]; General's Dragonhide Boots (231682, -102.12 DPS) [pvp] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Innervating Band (18701, +0.00 DPS) [dungeon]; Ring of Awareness (272409, +0.00 DPS) [vendor] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Ring of Awareness (272409, +0.00 DPS) [vendor]; Innervating Band (18701, -5.33 DPS) [dungeon] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Insignia Rank 6 (17909, -51.20 DPS) [quest]; Relentless Raider's Seal (272060, -1280.11 DPS) [vendor]; Darkspear Voodoo Seal (272061, -1280.11 DPS) [vendor] |
| trinket2 | Defiler's Talisman (20072) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Insignia Rank 6 (17909, +0.00 DPS) [quest]; Relentless Raider's Seal (272060, -292.02 DPS) [vendor]; Darkspear Voodoo Seal (272061, -292.02 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Warbear Helm; neck: Senior Sergeant's Insignia; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Raider Gloves; waist: Hivethrasher's Girdle; legs: Outrider's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Defender's Grip Stabilizer; trinket2: Defiler's Talisman; main_hand: The Unstoppable Force

No-known-source sample (15 of 1990, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 531.2. Weights run: 9.0s. Verify run: 46.2s. 1990 eligible items had no known source.

Race sweep: windshaper-skyborne 2053.5 ± 2.4, tauren 2028.6 ± 2.4 (threat_per_second); windshaper-skyborne beats the guide race tauren by 25.0 (1.23%) ± 3.3, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.207 ± 0.016, defense=1.191 ± 0.262 per rating point (1 rating = 1%, 1.191 per %), dodge=0.665 ± 0.054 per rating point (12 rating = 1%, 7.979 per %), strength=0.103 ± 0.000, agility=0.614 ± 0.035, attack_power=0.047 ± 0.000, feral_attack_power=0.047 ± 0.000, hit=0.204 ± 0.041 per rating point (10 rating = 1%, 2.038 per %), crit=0.109 ± 0.018 per rating point (14 rating = 1%, 1.532 per %), expertise=7.485 ± 0.449

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 74.9 stamina points (8122.64 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Helm (231678, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor] |
| neck | Talisman of Evasion (13177) | Blackrock Spire: War Master Voone [dungeon] | 20.0 stamina points (2173.82 DPS) | yes | Scout's Medallion (19534, -89.91 DPS) [rep]; Verek's Collar (11755, -94.03 DPS) [dungeon]; Medallion of Grand Marshal Morris (13091, -122.18 DPS) [world_drop] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Pauldrons (231672, +0.00 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, +0.00 DPS) [dungeon]; Bloodmoon Cloak (12967, +0.00 DPS) [dungeon]; Shadow Prowler's Cloak (22269, +0.00 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Magistrate's Cuffs (18726, -16.69 DPS) [dungeon]; Malefic Bracers (18700, -177.14 DPS) [dungeon]; Yeti Hide Bracers (19113, -328.65 DPS) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; General's Dragonhide Gauntlets (231676, -70.73 DPS) [vendor]; General's Dragonhide Gloves (231677, -70.73 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 48.7 stamina points (5281.19 DPS) | yes | Frostbite Girdle (14502, -423.62 DPS) [dungeon]; Hivethrasher's Girdle (275614, -451.51 DPS) [crafted]; Nagmara's Whipping Belt (11866, -513.92 DPS) [quest] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 81.5 stamina points (8839.61 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Pants (231680, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 56.9 stamina points (6175.44 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Boots (231682, +0.00 DPS) [pvp]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Band of Flesh (13373, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, +0.00 DPS) [vendor] |
| finger2 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Cutthroat's Signet (272408, -316.09 DPS) [vendor]; Band of Flesh (13373, -334.90 DPS) [dungeon] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Talisman (20072, -378.43 DPS) [rep]; Darkspear Voodoo Seal (272061, -595.37 DPS) [vendor] |
| trinket2 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Talisman (20072, -429.35 DPS) [rep]; Darkspear Voodoo Seal (272061, -646.29 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |

**New at 60:** head: Warbear Helm; neck: Talisman of Evasion; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Outrider's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Ring of Awareness; trinket1: Defender's Grip Stabilizer; trinket2: Frostwolf Insignia Rank 6; main_hand: The Unstoppable Force; ranged: Idol of Brutality

No-known-source sample (15 of 1990, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

