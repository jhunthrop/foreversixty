# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat. An item's crit rating reaches every table too, one point higher on melee, ranged and spell crit at once, while Agility raises melee and ranged crit only and Intellect spell crit only (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`). So a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 22.8. Weights run: 6.8s. Verify run: 6.3s. 251 eligible items had no known source.

Race sweep: high-order-skyborne 90.3 ± 0.1, night-elf 89.2 ± 0.1 (threat_per_second); high-order-skyborne beats the guide race night-elf by 1.0 (1.14%) ± 0.1, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.092 ± 0.001, defense=0.231 ± 0.025 per rating point (1 rating = 1%, 0.231 per %), dodge=0.100 ± 0.007 per rating point (12 rating = 1%, 1.203 per %), strength=0.094 ± 0.000, agility=0.153 ± 0.005, attack_power=0.047 ± 0.000, feral_attack_power=0.047 ± 0.000, hit=0.042 ± 0.004 per rating point (10 rating = 1%, 0.420 per %), crit=0.017 ± 0.001 per rating point (14 rating = 1%, 0.232 per %), expertise=1.052 ± 0.049

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 18.2 stamina points (339.97 DPS) | yes | Defender's Leather Hood (252447, -8.89 DPS) [crafted]; Trapper's Leather Hood (252505, -22.87 DPS) [crafted]; Stormrider's Leather Hood (252506, -22.87 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.7 stamina points (106.78 DPS) | yes | Erudite's Amulet (277204, -45.08 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -50.77 DPS) [quest]; Sentinel's Medallion (20444, -52.29 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.9 stamina points (184.49 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -53.53 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -63.48 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -5.52 DPS) [world_drop]; Firebane Cloak (12979, -20.39 DPS) [world_drop]; Caretaker's Cape (20428, -37.34 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 19.0 stamina points (354.97 DPS) | yes | Gloomshroud Armor (1489, -27.82 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -72.33 DPS) [quest]; Defender's Leather Armor (252434, -94.97 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.4 stamina points (156.83 DPS) | yes | Drakewing Bands (12999, -28.49 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -44.58 DPS) [vendor]; Sanguine Cuffs (14375, -52.98 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.5 stamina points (177.05 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.72 DPS) [crafted]; Defender's Leather Gloves (252496, -4.44 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.7 stamina points (219.06 DPS) | yes | Belt of the Fang (10412, -28.39 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -50.59 DPS) [crafted]; Guardsman Belt (3429, -52.31 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 15.0 stamina points (279.87 DPS) | yes | Brawler's Leather Pants (252500, -28.29 DPS) [crafted]; Leggings of the Fang (10410, -36.90 DPS) [dungeon]; Smelting Pants (5199, -88.11 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.3 stamina points (285.81 DPS) | yes | Surfer Shoes (276274, -51.08 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -51.98 DPS) [dungeon]; Trapper's Leather Boots (252440, -73.51 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (93.35 DPS) | yes | Slain Baron's Signet (279867, -33.03 DPS) [quest]; Ring of the Moon (12052, -35.59 DPS) [world_drop]; Deep Fathom Ring (6463, -37.34 DPS) [dungeon] |
| finger2 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 4.0 stamina points (74.68 DPS) | yes | Slain Baron's Signet (279867, -14.36 DPS) [quest]; Ring of the Moon (12052, -16.92 DPS) [world_drop]; Deep Fathom Ring (6463, -18.67 DPS) [dungeon] |
| trinket1 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.3 stamina points (379.84 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.93 DPS) [dungeon]; Twisted Chanter's Staff (890, -17.76 DPS) [world_drop]; Staff of the Blessed Seer (2271, -20.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Sustaining Ring; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 251, see the JSON for more): 1189 Overseer's Ring; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-51100332020102100000-0000000000000000)

Set DPS (verified): 64.8. Weights run: 6.1s. Verify run: 10.4s. 424 eligible items had no known source.

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

No-known-source sample (15 of 424, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 40 (night-elf, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 97.4. Weights run: 6.3s. Verify run: 5.8s. 607 eligible items had no known source.

Race sweep: night-elf 374.1 ± 0.4, high-order-skyborne 373.6 ± 0.4 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.025, armor=0.157 ± 0.008, defense=0.629 ± 0.127 per rating point (1 rating = 1%, 0.629 per %), dodge=0.178 ± 0.024 per rating point (12 rating = 1%, 2.136 per %), strength=0.100 ± 0.000, agility=0.229 ± 0.015, attack_power=0.050 ± 0.000, feral_attack_power=0.050 ± 0.000, hit=not significant (0.073 ± 0.019) per rating point (10 rating = 1%, 0.733 per %), crit=0.051 ± 0.008 per rating point (14 rating = 1%, 0.713 per %), expertise=2.988 ± 0.215

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 30.8 stamina points (993.56 DPS) | yes | Warden's Wizard Hat (14604, -41.95 DPS) [world_drop]; Brawler's Leather Helm (252512, -75.47 DPS) [crafted]; Nightscape Headband (8176, -89.81 DPS) [crafted] |
| neck | Shriveled Heart (9243) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 12.5 stamina points (403.19 DPS) | yes | Gazlowe's Charm (13088, -35.54 DPS) [dungeon]; Seal of the Expedition (271919, -48.41 DPS) [rep]; Sentinel's Medallion (19540, -96.27 DPS) [rep] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.8 stamina points (1024.84 DPS) | yes | Flintrock Shoulders (7755, -120.20 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -255.61 DPS) [dungeon]; Nightscape Shoulders (8192, -347.27 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -19.39 DPS) [dungeon]; Heraldic Cloak (8120, -22.13 DPS) [dungeon]; Well Oiled Cloak (12254, -24.45 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.4 stamina points (1143.06 DPS) | yes | Warden's Wraps (14601, -60.31 DPS) [world_drop]; Robes of the Lich (10762, -184.29 DPS) [dungeon]; Spirewind Fetter (9406, -199.04 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.5 stamina points (533.15 DPS) | yes | Branded Leather Bracers (19508, -29.15 DPS) [dungeon]; Barbaric Bracers (18948, -59.43 DPS) [crafted]; Cultist's Armguards (270032, -68.65 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -22.98 DPS) [crafted]; Prowler's Leather Gloves (252524, -27.14 DPS) [crafted]; Stalker's Leather Gloves (252526, -381.92 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.1 stamina points (744.89 DPS) | yes | Kolkar Hunter's Belt (6788, -66.70 DPS) [quest]; Warden's Leather Belt (252460, -72.27 DPS) [crafted]; Scorpashi Sash (14652, -88.40 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 30.5 stamina points (984.72 DPS) | yes | Imperial Leather Pants (4062, -61.33 DPS) [dungeon]; Panther Hunter Leggings (4108, -70.41 DPS) [quest]; Warden's Woolies (14605, -76.51 DPS) [world_drop] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 27.7 stamina points (894.46 DPS) | yes | Skulker's Leather Shoes (252531, -99.50 DPS) [crafted]; Prowler's Leather Shoes (252465, -116.14 DPS) [crafted]; Worn Running Boots (9398, -275.95 DPS, sim-verified) [dungeon] |
| finger1 | Talvash's Gold Ring (9538) | Gnome Improvement [quest] | 11.9 stamina points (383.36 DPS) | yes | Militant's Band (271915, -28.58 DPS) [rep]; Underworld Band (1980, -60.84 DPS) [world_drop]; Field Researcher's Loop (281634, -105.95 DPS) [quest] |
| finger2 | Theramore Signet (271910) | Theramore Expeditionary Force [rep] | 11.6 stamina points (374.09 DPS) | yes | Militant's Band (271915, -19.31 DPS) [rep]; Underworld Band (1980, -51.56 DPS) [world_drop]; Field Researcher's Loop (281634, -96.68 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (+286.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -32.25 DPS) [world]; Enriched Seal (271905, -32.25 DPS) [rep] |
| trinket2 | Talisman of Arathor (21118) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -32.25 DPS) [world]; Volunteer's Lucky Seal (271904, -32.25 DPS) [rep]; Enriched Seal (271905, -32.25 DPS) [rep] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Mograine's Might (7723, -131.95 DPS, sim-verified) [dungeon]; Windweaver Staff (7757, -415.52 DPS) [dungeon]; Huge Stone Club (1523, -500.37 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Talvash's Gold Ring; finger2: Theramore Signet; trinket1: Rune of Duty; trinket2: Talisman of Arathor; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 607, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band

### Band 50 (night-elf, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 141.8. Weights run: 6.2s. Verify run: 11.1s. 815 eligible items had no known source.

Race sweep: night-elf 545.3 ± 0.5, high-order-skyborne 542.6 ± 0.5 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.025, armor=0.119 ± 0.010, defense=not significant (0.571 ± 0.145) per rating point (1 rating = 1%, 0.571 per %), dodge=0.275 ± 0.028 per rating point (12 rating = 1%, 3.298 per %), strength=0.078 ± 0.000, agility=0.274 ± 0.018, attack_power=0.039 ± 0.000, feral_attack_power=0.039 ± 0.000, hit=0.097 ± 0.024 per rating point (10 rating = 1%, 0.975 per %), crit=0.070 ± 0.011 per rating point (14 rating = 1%, 0.980 per %), expertise=3.300 ± 0.257

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.8 stamina points (1841.06 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, +0.00 DPS) [vendor]; Sprightring Helm (17776, -9.66 DPS) [quest]; Scorpashi Skullcap (14658, -122.26 DPS) [world_drop] |
| neck | Lifeblood Amulet (9641) (or Master Sergeant's Insignia (18444)) | Zul'Farrak: Antu'sul [dungeon] | 14.0 stamina points (811.17 DPS) | yes | Master Sergeant's Insignia (18444, +0.00 DPS) [vendor]; Seal of the Expedition (271918, -57.94 DPS) [rep]; Shriveled Heart (9243, -80.57 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.3 stamina points (1640.97 DPS) | yes | Knight-Lieutenant's Leather Shoulders (220852, -59.99 DPS) [vendor]; Phytoskin Spaulders (17749, -117.21 DPS) [dungeon]; Warden's Leather Shoulders (252536, -135.82 DPS) [crafted] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant's Cape (18441, -78.59 DPS) [vendor]; Emberscale Cape (11311, -92.36 DPS) [dungeon]; Nightfall Drape (12465, -470.67 DPS, sim-verified) [dungeon] |
| chest | Knight's Leather Armor (220854) | PvP rank 6 · Knight · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jinxed Hoodoo Skin (9473, +0.00 DPS) [dungeon]; Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Warbear Harness (15064, -169.11 DPS, sim-verified) [crafted] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.2 stamina points (1347.00 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -95.28 DPS) [vendor]; Warden's Leather Bracers (252541, -258.02 DPS) [crafted]; Serpentskin Bracers (8257, -262.77 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wicked Leather Gauntlets (15083, -256.40 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -267.07 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -821.08 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Waistguard (252474, -9.43 DPS) [crafted]; Prowler's Leather Waistguard (252473, -32.14 DPS) [crafted]; Girdle of Beastial Fury (11686, -701.19 DPS, sim-verified) [dungeon] |
| legs | Knight's Leather Pants (220858) | PvP rank 6 · Knight · Alliance [vendor] | 33.6 stamina points (1945.51 DPS) | yes | Brusslehide Leggings (17751, -355.35 DPS) [dungeon]; Basilisk Hide Pants (1718, -370.64 DPS) [world_drop]; Knight's Crackling Leather Leggings (220864, -1366.35 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 5 · Sergeant Major · Alliance [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Warden's Leather Boots (252470, +0.00 DPS) [crafted]; Shadefiend Boots (11675, -816.55 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (869.11 DPS) | yes | Militant's Band (271914, -57.94 DPS) [rep]; Talvash's Gold Ring (9538, -190.53 DPS) [quest] |
| finger2 | Theramore Signet (271909) | Theramore Expeditionary Force [rep] | 14.6 stamina points (845.12 DPS) | yes | Militant's Band (271914, -33.95 DPS) [rep]; Talvash's Gold Ring (9538, -166.53 DPS) [quest] |
| trinket1 | Enriched Seal (271903) | Theramore Expeditionary Force [rep] | sim-verified (+544.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Guardian Talisman (1490, -57.94 DPS) [quest]; Rune of Perfection (21565, -115.88 DPS) [rep] |
| trinket2 | Talisman of Arathor (21117) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -57.94 DPS) [quest]; Rune of Perfection (21565, -115.88 DPS) [rep]; Rune of Duty (21567, -115.88 DPS) [rep] |
| main_hand | Warden Staff (943) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -66.25 DPS) [dungeon]; Blanchard's Stout (13046, -105.61 DPS) [world_drop]; Taran Icebreaker (2915, -2367.77 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Lifeblood Amulet; back: Graverot Cape; chest: Knight's Leather Armor; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Ring of Saviors; finger2: Theramore Signet; trinket1: Enriched Seal; trinket2: Talisman of Arathor; main_hand: Warden Staff

No-known-source sample (15 of 815, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60 (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 204.9. Weights run: 6.6s. Verify run: 39.6s. 1969 eligible items had no known source.

Race sweep: night-elf 806.3 ± 0.8, high-order-skyborne 804.5 ± 0.8 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.029, armor=0.159 ± 0.016, defense=not significant (0.750 ± 0.244) per rating point (1 rating = 1%, 0.750 per %), dodge=0.391 ± 0.044 per rating point (12 rating = 1%, 4.692 per %), strength=0.092 ± 0.000, agility=0.370 ± 0.028, attack_power=0.046 ± 0.000, feral_attack_power=0.046 ± 0.000, hit=not significant (0.128 ± 0.036) per rating point (10 rating = 1%, 1.278 per %), crit=0.077 ± 0.016 per rating point (14 rating = 1%, 1.075 per %), expertise=4.960 ± 0.434

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 64.0 stamina points (4840.76 DPS) | yes | Bone Ring Helm (14539, +0.00 DPS) [dungeon]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Master Sergeant's Insignia (18443) | PvP rank 4 · Master Sergeant · Alliance [vendor] | 17.0 stamina points (1286.00 DPS) | yes | Stormpike Soldier's Pendant (19097, -88.46 DPS) [rep]; Sentinel's Medallion (19538, -110.02 DPS) [rep]; Seal of the Expedition (271917, -151.29 DPS) [rep] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 48.6 stamina points (3673.24 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -1508.75 DPS, sim-verified) [rep] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Emperor's New Cape (11930, -786.86 DPS) [dungeon]; Sergeant's Cape (16342, -910.91 DPS) [vendor]; Bloodmoon Cloak (12967, -914.23 DPS, sim-verified) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.8 stamina points (4522.00 DPS) | yes | Breastplate of Bloodthirst (12757, +0.00 DPS) [quest]; Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | 31.9 stamina points (2409.64 DPS) | yes | Malefic Bracers (18700, -52.43 DPS) [dungeon]; Yeti Hide Bracers (19113, -152.36 DPS) [world]; Forest Stalker's Bracers (19587, -162.48 DPS) [rep] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Hivethrasher's Claws (275613, -559.62 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 38.4 stamina points (2902.31 DPS) | yes | Shifter's Belt (272396, -51.05 DPS) [vendor]; Frostbite Girdle (14502, -149.09 DPS) [dungeon]; Marshal's Dragonhide Waistguard (16447, -206.34 DPS) [pvp] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 83.1 stamina points (6284.18 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -114.63 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -841.69 DPS, sim-verified) [vendor] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 43.5 stamina points (3290.73 DPS) | yes | Dawn Treaders (19052, +0.00 DPS) [crafted]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+299.1 DPS vs the runner-up, not corroborated against the finished set) | yes | Theramore Signet (271908, +0.00 DPS) [rep]; Ring of Awareness (272409, +0.00 DPS) [vendor]; Myrmidon's Signet (2246, -299.09 DPS, sim-verified) [world_drop] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Ring of Awareness (272409, -12.87 DPS) [vendor]; Theramore Signet (271908, -66.45 DPS) [rep] |
| trinket1 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, +0.00 DPS) [rep]; Volunteer's Lucky Seal (271906, +0.00 DPS) [rep]; Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, -151.29 DPS) [rep]; Volunteer's Lucky Seal (271906, -151.29 DPS) [rep]; Defender's Grip Stabilizer (272440, -169.73 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (+413.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -413.77 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, +0.00 DPS) [world_drop] |

**New at 60:** head: Warbear Helm; neck: Master Sergeant's Insignia; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Theramore Gloves; waist: Hivethrasher's Girdle; legs: Sentinel's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Stormpike Insignia Rank 6; trinket2: Talisman of Arathor; main_hand: The Unstoppable Force; ranged: Idol of Brutality

No-known-source sample (15 of 1969, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

### Band 60, raid preset (night-elf, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 455.0. Weights run: 7.4s. Verify run: 44.2s. 1969 eligible items had no known source.

Race sweep: night-elf 1742.7 ± 2.0, high-order-skyborne 1738.7 ± 2.0 (threat_per_second); the guide race night-elf is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.020, armor=0.197 ± 0.017, defense=1.421 ± 0.294 per rating point (1 rating = 1%, 1.421 per %), dodge=0.516 ± 0.053 per rating point (12 rating = 1%, 6.196 per %), strength=0.099 ± 0.000, agility=0.531 ± 0.034, attack_power=0.045 ± 0.000, feral_attack_power=0.045 ± 0.000, hit=not significant (0.127 ± 0.041) per rating point (10 rating = 1%, 1.271 per %), crit=0.101 ± 0.019 per rating point (14 rating = 1%, 1.412 per %), expertise=6.619 ± 0.451

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 72.2 stamina points (8315.76 DPS) | yes | Bone Ring Helm (14539, +0.00 DPS) [dungeon]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Medallion of Grand Marshal Morris (13091) | World drop [world_drop] | 21.2 stamina points (2442.40 DPS) | yes | Talisman of Evasion (13177, -258.44 DPS) [dungeon]; Verek's Collar (11755, -363.87 DPS) [dungeon]; Sentinel's Medallion (19538, -374.74 DPS) [rep] |
| shoulder | Highlander's Leather Shoulders (20059) | The League of Arathor [rep] | 57.1 stamina points (6577.32 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, +0.00 DPS) [pvp]; Theramore Pauldrons (271955, -1945.44 DPS, sim-verified) [rep] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodmoon Cloak (12967, -898.50 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -1702.62 DPS) [dungeon]; Shadow Prowler's Cloak (22269, -1974.43 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 68.8 stamina points (7926.13 DPS) | yes | Breastplate of Bloodthirst (12757, +0.00 DPS) [quest]; Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | 39.4 stamina points (4532.56 DPS) | yes | Malefic Bracers (18700, -359.31 DPS) [dungeon]; Forest Stalker's Bracers (19587, -370.21 DPS) [rep]; Yeti Hide Bracers (19113, -601.43 DPS) [world] |
| hands | Theramore Gloves (271948) | Theramore Expeditionary Force [rep] | sim-verified (+636.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Marshal's Dragonhide Gauntlets (231706, +0.00 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 45.7 stamina points (5256.60 DPS) | yes | Nagmara's Whipping Belt (11866, -57.13 DPS) [quest]; Hivethrasher's Girdle (275614, -296.33 DPS) [crafted]; Frostbite Girdle (14502, -350.69 DPS) [dungeon] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 95.8 stamina points (11035.38 DPS) | yes | Marshal's Dragonhide Leggings (231691, -172.07 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -818.64 DPS) [pvp]; Sentinel's Lizardhide Pants (237817, -1682.13 DPS, sim-verified) [vendor] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 53.7 stamina points (6178.50 DPS) | yes | Dawn Treaders (19052, +0.00 DPS) [crafted]; Marshal's Dragonhide Treads (231692, +0.00 DPS) [vendor]; Marshal's Dragonhide Greaves (231704, +0.00 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Innervating Band (18701, -17.08 DPS) [dungeon]; Ring of Awareness (272409, -365.57 DPS, sim-verified) [vendor] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-verified (+958.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Ring of Awareness (272409, +0.00 DPS) [vendor]; Innervating Band (18701, -141.16 DPS) [dungeon]; Myrmidon's Signet (2246, -958.28 DPS, sim-verified) [world_drop] |
| trinket1 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Enriched Seal (271903, -230.28 DPS) [rep]; Volunteer's Lucky Seal (271906, -230.28 DPS) [rep]; Defender's Grip Stabilizer (272440, -1302.18 DPS, sim-verified) [vendor] |
| trinket2 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-verified (+771.8 DPS vs the runner-up, not corroborated against the finished set) | yes | Enriched Seal (271903, -275.46 DPS) [rep]; Volunteer's Lucky Seal (271906, -275.46 DPS) [rep]; Defender's Grip Stabilizer (272440, -771.75 DPS, sim-verified) [vendor] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | The Unstoppable Force (19323, +0.00 DPS) [rep]; Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Demolisher (234568, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, -1088.66 DPS, sim-verified) [world_drop] |

**New at 60:** head: Warbear Helm; neck: Medallion of Grand Marshal Morris; shoulder: Highlander's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Theramore Gloves; waist: Shifter's Belt; legs: Sentinel's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Talisman of Arathor; trinket2: Stormpike Insignia Rank 6; main_hand: Headmaster's Charge; ranged: Idol of Brutality

No-known-source sample (15 of 1969, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band

## Horde

### Band 20 (tauren, 0000000000000000-51100301000000000000-0000000000000000)

Set DPS (verified): 23.3. Weights run: 6.8s. Verify run: 5.9s. 247 eligible items had no known source.

Race sweep: windshaper-skyborne 91.2 ± 0.1, tauren 91.0 ± 0.1 (threat_per_second); windshaper-skyborne beats the guide race tauren by 0.1 (0.15%) ± 0.1, inside error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.092 ± 0.001, defense=0.231 ± 0.025 per rating point (1 rating = 1%, 0.231 per %), dodge=0.100 ± 0.007 per rating point (12 rating = 1%, 1.203 per %), strength=0.094 ± 0.000, agility=0.153 ± 0.005, attack_power=0.047 ± 0.000, feral_attack_power=0.047 ± 0.000, hit=0.042 ± 0.004 per rating point (10 rating = 1%, 0.420 per %), crit=0.017 ± 0.001 per rating point (14 rating = 1%, 0.232 per %), expertise=1.052 ± 0.049

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 18.2 stamina points (339.97 DPS) | yes | Defender's Leather Hood (252447, -8.89 DPS) [crafted]; Trapper's Leather Hood (252505, -22.87 DPS) [crafted]; Stormrider's Leather Hood (252506, -22.87 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (61.73 DPS) | yes | Scholarly Pendant (277203, -5.72 DPS) [quest]; Scout's Medallion (20442, -7.24 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.9 stamina points (184.49 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -53.53 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -63.48 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sentry Cloak (2059, -5.52 DPS) [world_drop]; Firebane Cloak (12979, -20.39 DPS) [world_drop]; Battle Healer's Cloak (20427, -37.34 DPS) [rep] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 19.0 stamina points (354.97 DPS) | yes | Gloomshroud Armor (1489, -34.00 DPS, sim-verified) [dungeon]; Defender's Leather Armor (252434, -94.97 DPS) [crafted]; Trapper's Leather Armor (252491, -95.77 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.4 stamina points (156.83 DPS) | yes | Spare Part Bindings (279875, -28.55 DPS, sim-verified) [quest]; Drakewing Bands (12999, -30.48 DPS) [world_drop]; Savannah Bracers (15453, -44.20 DPS) [quest] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.5 stamina points (177.05 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.72 DPS) [crafted]; Defender's Leather Gloves (252496, -4.44 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.7 stamina points (219.06 DPS) | yes | Belt of the Fang (10412, -33.02 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -50.59 DPS) [crafted]; Guardsman Belt (3429, -52.31 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 15.0 stamina points (279.87 DPS) | yes | Brawler's Leather Pants (252500, -28.29 DPS) [crafted]; Leggings of the Fang (10410, -36.90 DPS) [dungeon]; Smelting Pants (5199, -98.66 DPS, sim-verified) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.3 stamina points (285.81 DPS) | yes | Surfer Shoes (276274, -51.06 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -51.98 DPS) [dungeon]; Trapper's Leather Boots (252440, -73.51 DPS) [crafted] |
| finger1 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (93.35 DPS) | yes | Slain Baron's Signet (279867, -33.03 DPS) [quest]; Ring of the Moon (12052, -35.59 DPS) [world_drop]; Deep Fathom Ring (6463, -37.34 DPS) [dungeon] |
| finger2 | Ring of Scorn (3235) | A Husband's Revenge [quest] | 4.0 stamina points (74.68 DPS) | yes | Slain Baron's Signet (279867, -14.36 DPS) [quest]; Ring of the Moon (12052, -16.92 DPS) [world_drop]; Deep Fathom Ring (6463, -18.67 DPS) [dungeon] |
| trinket1 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Rune of Perfection (21566, +0.00 DPS) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Rakzur Club (12983) | World drop [world_drop] | 20.3 stamina points (379.84 DPS) | yes | Rhahk'Zor's Hammer (5187, -4.93 DPS) [dungeon]; Twisted Chanter's Staff (890, -17.76 DPS) [world_drop]; Staff of the Blessed Seer (2271, -20.16 DPS) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Brawler's Leather Hood; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Blood Ring; finger2: Ring of Scorn; trinket1: Rune of Duty; main_hand: Rakzur Club

No-known-source sample (15 of 247, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6360 Steelscale Crushfish; 6478 Rat Stompers; 7187 VanCleef's Boots; 8350 The 1 Ring; 14389 Durability Shoulderpads

### Band 30 (tauren, 0000000000000000-51100332020102100000-0000000000000000)

Set DPS (verified): 64.7. Weights run: 6.1s. Verify run: 9.6s. 425 eligible items had no known source.

Race sweep: tauren 249.2 ± 0.2, windshaper-skyborne 249.1 ± 0.2 (threat_per_second); the guide race tauren is best.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.104 ± 0.005, defense=0.404 ± 0.085 per rating point (1 rating = 1%, 0.404 per %), dodge=0.164 ± 0.017 per rating point (12 rating = 1%, 1.965 per %), strength=0.121 ± 0.000, agility=0.215 ± 0.011, attack_power=0.061 ± 0.000, feral_attack_power=0.061 ± 0.000, hit=0.070 ± 0.013 per rating point (10 rating = 1%, 0.702 per %), crit=0.034 ± 0.006 per rating point (14 rating = 1%, 0.471 per %), expertise=2.061 ± 0.144

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 23.9 stamina points (528.55 DPS) | yes | Defender's Leather Helm (252455, -35.03 DPS, sim-verified) [crafted]; Trapper's Leather Helm (252513, -47.72 DPS) [crafted]; Stormrider's Leather Helm (252514, -47.72 DPS) [crafted] |
| neck | River Pride Choker (13087) | World drop [world_drop] | 9.5 stamina points (210.16 DPS) | yes | Senior Sergeant's Insignia (15200, -10.76 DPS) [vendor]; Darkspear Warding Pendant (272075, -32.91 DPS) [vendor]; Scout's Medallion (19537, -61.21 DPS) [rep] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 19.7 stamina points (435.41 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -72.47 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -72.47 DPS) [crafted]; Tanned Shoulderpads (270023, -84.29 DPS, sim-verified) [quest] |
| back | Sergeant's Cloak (18427) | PvP rank 3 · Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Tigerstrike Mantle (13108, -8.43 DPS) [world_drop]; Enduring Cape (14763, -38.57 DPS) [world_drop]; Grimsteel Cape (4643, -101.23 DPS, sim-verified) [quest] |
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

No-known-source sample (15 of 425, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots

### Band 40 (tauren, 0000000000000000-51200332023102123010-0000000000000000)

Set DPS (verified): 95.9. Weights run: 6.3s. Verify run: 6.2s. 606 eligible items had no known source.

Race sweep: windshaper-skyborne 368.6 ± 0.4, tauren 368.0 ± 0.4 (threat_per_second); windshaper-skyborne beats the guide race tauren by 0.6 (0.17%) ± 0.5, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.025, armor=0.157 ± 0.008, defense=0.629 ± 0.127 per rating point (1 rating = 1%, 0.629 per %), dodge=0.178 ± 0.024 per rating point (12 rating = 1%, 2.136 per %), strength=0.100 ± 0.000, agility=0.229 ± 0.015, attack_power=0.050 ± 0.000, feral_attack_power=0.050 ± 0.000, hit=not significant (0.073 ± 0.019) per rating point (10 rating = 1%, 0.733 per %), crit=0.051 ± 0.008 per rating point (14 rating = 1%, 0.713 per %), expertise=2.988 ± 0.215

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 30.8 stamina points (993.56 DPS) | yes | Warden's Wizard Hat (14604, -41.95 DPS) [world_drop]; Brawler's Leather Helm (252512, -75.47 DPS) [crafted]; Nightscape Headband (8176, -89.81 DPS) [crafted] |
| neck | Shriveled Heart (9243) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 12.5 stamina points (403.19 DPS) | yes | Dragon's Blood Necklace (10711, -16.16 DPS) [quest]; Gazlowe's Charm (13088, -35.54 DPS) [dungeon]; Darkspear Warding Pendant (272074, -48.41 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 31.8 stamina points (1024.84 DPS) | yes | Flintrock Shoulders (7755, -128.40 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -255.61 DPS) [dungeon]; Nightscape Shoulders (8192, -347.27 DPS) [crafted] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Silky Spider Cape (10776, -19.39 DPS) [dungeon]; Heraldic Cloak (8120, -22.13 DPS) [dungeon]; Well Oiled Cloak (12254, -24.45 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 35.4 stamina points (1143.06 DPS) | yes | Warden's Wraps (14601, -68.56 DPS, sim-verified) [world_drop]; Robes of the Lich (10762, -184.29 DPS) [dungeon]; Spirewind Fetter (9406, -199.04 DPS) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 16.5 stamina points (533.15 DPS) | yes | Branded Leather Bracers (19508, -29.15 DPS) [dungeon]; Barbaric Bracers (18948, -59.43 DPS) [crafted]; Cultist's Armguards (270032, -68.65 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Gloves (252525, -22.98 DPS) [crafted]; Prowler's Leather Gloves (252524, -27.14 DPS) [crafted]; Stalker's Leather Gloves (252526, -357.58 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 23.1 stamina points (744.89 DPS) | yes | Kolkar Hunter's Belt (6788, -66.70 DPS) [quest]; Warden's Leather Belt (252460, -72.27 DPS) [crafted]; Scorpashi Sash (14652, -88.40 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 30.5 stamina points (984.72 DPS) | yes | Panther Hunter Leggings (4108, -70.41 DPS) [quest]; Warden's Woolies (14605, -76.51 DPS) [world_drop]; Imperial Leather Pants (4062, -81.04 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 27.7 stamina points (894.46 DPS) | yes | Skulker's Leather Shoes (252531, -99.50 DPS) [crafted]; Prowler's Leather Shoes (252465, -116.14 DPS) [crafted]; Worn Running Boots (9398, -333.40 DPS, sim-verified) [dungeon] |
| finger1 | Nogg's Gold Ring (9588) | Nogg's Ring Redo [quest] | 11.9 stamina points (383.36 DPS) | yes | Darkspear Signet (272070, -28.58 DPS) [vendor]; Underworld Band (1980, -60.84 DPS) [world_drop]; Field Researcher's Loop (281634, -105.95 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.6 stamina points (374.09 DPS) | yes | Darkspear Signet (272070, -19.31 DPS) [vendor]; Underworld Band (1980, -51.56 DPS) [world_drop]; Field Researcher's Loop (281634, -96.68 DPS) [quest] |
| trinket1 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (+302.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Rune of Perfection (21565, +0.00 DPS) [rep]; Arena Master (18706, -32.25 DPS) [world]; Darkspear Voodoo Seal (272059, -32.25 DPS) [vendor] |
| trinket2 | Defiler's Talisman (21116) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Arena Master (18706, -32.25 DPS) [world]; Darkspear Voodoo Seal (272059, -32.25 DPS) [vendor]; Relentless Raider's Seal (272062, -32.25 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ironshod Bludgeon (9408, -145.04 DPS, sim-verified) [dungeon]; Mograine's Might (7723, -258.92 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -302.13 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Nogg's Gold Ring; finger2: Insurgent's Band; trinket1: Rune of Duty; trinket2: Defiler's Talisman; main_hand: Cragwood Maul

No-known-source sample (15 of 606, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring

### Band 50 (tauren, 0000000000000000-51200332023102125510-0300000000000000)

Set DPS (verified): 140.5. Weights run: 6.2s. Verify run: 10.7s. 814 eligible items had no known source.

Race sweep: windshaper-skyborne 542.6 ± 0.5, tauren 539.5 ± 0.5 (threat_per_second); windshaper-skyborne beats the guide race tauren by 3.1 (0.57%) ± 0.7, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.025, armor=0.119 ± 0.010, defense=not significant (0.571 ± 0.145) per rating point (1 rating = 1%, 0.571 per %), dodge=0.275 ± 0.028 per rating point (12 rating = 1%, 3.298 per %), strength=0.078 ± 0.000, agility=0.274 ± 0.018, attack_power=0.039 ± 0.000, feral_attack_power=0.039 ± 0.000, hit=0.097 ± 0.024 per rating point (10 rating = 1%, 0.975 per %), crit=0.070 ± 0.011 per rating point (14 rating = 1%, 0.980 per %), expertise=3.300 ± 0.257

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Embrace of the Lycan (9479) | Zul'Farrak: Chief Ukorz Sandscalp [dungeon] | 31.8 stamina points (1841.06 DPS) | yes | Blood Guard's Leather Headband (220851, +0.00 DPS) [vendor]; Sprightring Helm (17776, -9.66 DPS) [quest]; Scorpashi Skullcap (14658, -122.26 DPS) [world_drop] |
| neck | Lifeblood Amulet (9641) (or Senior Sergeant's Insignia (18428)) | Zul'Farrak: Antu'sul [dungeon] | 14.0 stamina points (811.17 DPS) | yes | Senior Sergeant's Insignia (18428, +0.00 DPS) [vendor]; Darkspear Warding Pendant (272073, -57.94 DPS) [vendor]; Woven Ivy Necklace (19159, -61.68 DPS) [quest] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.3 stamina points (1640.97 DPS) | yes | Blood Guard's Leather Shoulders (220853, -59.99 DPS) [vendor]; Warden's Leather Shoulders (252536, -135.82 DPS) [crafted]; Phytoskin Spaulders (17749, -243.39 DPS, sim-verified) [dungeon] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Sergeant's Cloak (16341, -78.59 DPS) [vendor]; Emberscale Cape (11311, -92.36 DPS) [dungeon]; Nightfall Drape (12465, -540.97 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Leather Armor (220855) | PvP rank 6 · Stone Guard · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Jinxed Hoodoo Skin (9473, +0.00 DPS) [dungeon]; Mixologist's Tunic (12793, +0.00 DPS) [dungeon]; Warbear Harness (15064, -386.54 DPS, sim-verified) [crafted] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.2 stamina points (1347.00 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -95.28 DPS) [vendor]; Warden's Leather Bracers (252541, -258.02 DPS) [crafted]; Serpentskin Bracers (8257, -262.77 DPS) [world_drop] |
| hands | Warden's Leather Gauntlets (252549) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Wicked Leather Gauntlets (15083, -256.40 DPS) [crafted]; Prowler's Leather Gauntlets (252547, -267.07 DPS) [crafted]; Skulker's Leather Gauntlets (252548, -962.89 DPS, sim-verified) [crafted] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Skulker's Leather Waistguard (252474, -9.43 DPS) [crafted]; Prowler's Leather Waistguard (252473, -32.14 DPS) [crafted]; Girdle of Beastial Fury (11686, -914.16 DPS, sim-verified) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | PvP rank 6 · Stone Guard · Horde [vendor] | 33.6 stamina points (1945.51 DPS) | yes | Brusslehide Leggings (17751, -355.35 DPS) [dungeon]; Basilisk Hide Pants (1718, -370.64 DPS) [world_drop]; Stone Guard's Crackling Leather Leggings (220865, -1521.73 DPS, sim-verified) [vendor] |
| feet | First Sergeant's Leather Boots (220861) | PvP rank 5 · First Sergeant · Horde [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Slitherscale Boots (10801, +0.00 DPS) [dungeon]; Warden's Leather Boots (252470, +0.00 DPS) [crafted]; Shadefiend Boots (11675, -826.56 DPS, sim-verified) [dungeon] |
| finger1 | Ring of Saviors (1447) | World drop [world_drop] | 15.0 stamina points (869.11 DPS) | yes | Darkspear Signet (272069, -57.94 DPS) [vendor]; Nogg's Gold Ring (9588, -190.53 DPS) [quest] |
| finger2 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.6 stamina points (845.12 DPS) | yes | Darkspear Signet (272069, -33.95 DPS) [vendor]; Nogg's Gold Ring (9588, -166.53 DPS) [quest] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (+601.3 DPS vs the runner-up, not corroborated against the finished set) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -57.94 DPS) [quest]; Rune of Perfection (21565, -115.88 DPS) [rep] |
| trinket2 | Defiler's Talisman (21115) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Guardian Talisman (1490, -57.94 DPS) [quest]; Rune of Perfection (21565, -115.88 DPS) [rep]; Rune of Duty (21567, -115.88 DPS) [rep] |
| main_hand | Warden Staff (943) | World drop [world_drop] | sim-decided (no score - a real sim tournament chose this pick) | yes | Taran Icebreaker (2915, -52.89 DPS) [world_drop]; Ironshod Bludgeon (9408, -66.25 DPS) [dungeon]; Cragwood Maul (11265, -2651.54 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Embrace of the Lycan; neck: Lifeblood Amulet; back: Graverot Cape; chest: Stone Guard's Leather Armor; wrist: Arena Bracers; hands: Warden's Leather Gauntlets; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: First Sergeant's Leather Boots; finger1: Ring of Saviors; finger2: Insurgent's Band; trinket1: Darkspear Voodoo Seal; trinket2: Defiler's Talisman; main_hand: Warden Staff

No-known-source sample (15 of 814, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60 (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 194.1. Weights run: 6.6s. Verify run: 40.2s. 1973 eligible items had no known source.

Race sweep: windshaper-skyborne 757.1 ± 0.8, tauren 754.0 ± 0.8 (threat_per_second); windshaper-skyborne beats the guide race tauren by 3.1 (0.41%) ± 1.1, beyond error.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.029, armor=0.159 ± 0.016, defense=not significant (0.750 ± 0.244) per rating point (1 rating = 1%, 0.750 per %), dodge=0.391 ± 0.044 per rating point (12 rating = 1%, 4.692 per %), strength=0.092 ± 0.000, agility=0.370 ± 0.028, attack_power=0.046 ± 0.000, feral_attack_power=0.046 ± 0.000, hit=not significant (0.128 ± 0.036) per rating point (10 rating = 1%, 1.278 per %), crit=0.077 ± 0.016 per rating point (14 rating = 1%, 1.075 per %), expertise=4.960 ± 0.434

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 64.0 stamina points (4840.76 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Bone Ring Helm (14539, -670.03 DPS, sim-verified) [dungeon] |
| neck | Senior Sergeant's Insignia (16335) | PvP rank 4 · Senior Sergeant · Horde [vendor] | 17.0 stamina points (1286.00 DPS) | yes | Frostwolf Legionnaire's Pendant (19095, -88.46 DPS) [rep]; Scout's Medallion (19534, -110.02 DPS) [rep]; Hunter's Insignia Medal (15704, -151.29 DPS) [quest] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | 48.6 stamina points (3673.24 DPS) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -1591.57 DPS, sim-verified) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bloodmoon Cloak (12967, -638.81 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -786.86 DPS) [dungeon]; Sergeant's Cloak (18461, -910.91 DPS) [vendor] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 59.8 stamina points (4522.00 DPS) | yes | Breastplate of Bloodthirst (12757, +0.00 DPS) [quest]; Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor] |
| wrist | Magistrate's Cuffs (18726) | Stratholme: Magistrate Barthilas [dungeon] | 31.9 stamina points (2409.64 DPS) | yes | Yeti Hide Bracers (19113, -152.36 DPS) [world]; Forest Stalker's Bracers (19587, -162.48 DPS) [rep]; Malefic Bracers (18700, -282.59 DPS, sim-verified) [dungeon] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | General's Dragonhide Gauntlets (231676, +0.00 DPS) [vendor]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Hivethrasher's Claws (275613, -601.15 DPS, sim-verified) [crafted] |
| waist | Hivethrasher's Girdle (275614) | Leatherworking [crafted] | 38.4 stamina points (2902.31 DPS) | yes | Shifter's Belt (272396, -51.05 DPS) [vendor]; Frostbite Girdle (14502, -149.09 DPS) [dungeon]; General's Dragonhide Belt (16556, -206.34 DPS) [pvp] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 65.9 stamina points (4986.49 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Warstrife Leggings (11821, -1558.04 DPS, sim-verified) [dungeon] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 43.5 stamina points (3290.73 DPS) | yes | General's Dragonhide Greaves (231671, +0.00 DPS) [vendor]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; Dawn Treaders (19052, -602.32 DPS, sim-verified) [crafted] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+314.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Insurgent's Band (272064, +0.00 DPS) [vendor]; Ring of Awareness (272409, +0.00 DPS) [vendor]; Myrmidon's Signet (2246, -314.55 DPS, sim-verified) [world_drop] |
| finger2 | Band of Flesh (13373) | Stratholme: Ramstein the Gorger [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Ring of Awareness (272409, -12.87 DPS) [vendor]; Insurgent's Band (272064, -66.45 DPS) [vendor] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Darkspear Voodoo Seal (272061, +0.00 DPS) [vendor]; Frostwolf Insignia Rank 6 (17909, -28.13 DPS) [quest] |
| trinket2 | Defiler's Talisman (20072) | The Defilers [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, -151.29 DPS) [vendor]; Darkspear Voodoo Seal (272061, -151.29 DPS) [vendor]; Frostwolf Insignia Rank 6 (17909, -197.86 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (+132.6 DPS vs the runner-up, not corroborated against the finished set) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp]; Headmaster's Charge (13937, -132.63 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Warbear Helm; neck: Senior Sergeant's Insignia; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Magistrate's Cuffs; hands: Raider Gloves; waist: Hivethrasher's Girdle; legs: Outrider's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Band of Flesh; trinket1: Defender's Grip Stabilizer; trinket2: Defiler's Talisman; main_hand: The Unstoppable Force

No-known-source sample (15 of 1973, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

### Band 60, raid preset (tauren, 0000000000000000-51230332023102125510-0550000000000000)

Set DPS (verified): 506.5. Weights run: 7.4s. Verify run: 40.7s. 1973 eligible items had no known source.

Race sweep: windshaper-skyborne 1956.6 ± 2.3, tauren 1941.1 ± 2.2 (threat_per_second); windshaper-skyborne beats the guide race tauren by 15.6 (0.80%) ± 3.2, beyond error.

1 slot(s) kept a confirmed-stats item over one whose stats the client has not confirmed (within the sim error).

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.176 ± 0.018, defense=1.420 ± 0.325 per rating point (1 rating = 1%, 1.420 per %), dodge=0.982 ± 0.058 per rating point (12 rating = 1%, 11.782 per %), strength=0.120 ± 0.000, agility=0.869 ± 0.039, attack_power=0.054 ± 0.000, feral_attack_power=0.054 ± 0.000, hit=0.192 ± 0.047 per rating point (10 rating = 1%, 1.921 per %), crit=0.193 ± 0.023 per rating point (14 rating = 1%, 2.707 per %), expertise=6.732 ± 0.436

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warbear Helm (252485) | Leatherworking [crafted] | 74.0 stamina points (15161.70 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Bone Ring Helm (14539, -609.56 DPS, sim-verified) [dungeon] |
| neck | Talisman of Evasion (13177) | Blackrock Spire: War Master Voone [dungeon] | 27.2 stamina points (5579.49 DPS) | yes | Verek's Collar (11755, -714.35 DPS) [dungeon]; Scout's Medallion (19534, -864.18 DPS) [rep]; Medallion of Grand Marshal Morris (13091, -1239.27 DPS) [world_drop] |
| shoulder | Defiler's Leather Shoulders (20194) | The Defilers [rep] | sim-verified (+2640.6 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Spaulders (231681, +0.00 DPS) [vendor]; Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Darkspear Pauldrons (272105, -2640.60 DPS, sim-verified) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Shadow Prowler's Cloak (22269, -1331.76 DPS, sim-verified) [dungeon]; The Emperor's New Cape (11930, -1822.78 DPS) [dungeon]; Bloodmoon Cloak (12967, -2096.72 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Armor (231679, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | 40.9 stamina points (8383.74 DPS) | yes | Magistrate's Cuffs (18726, -648.95 DPS) [dungeon]; Malefic Bracers (18700, -866.57 DPS) [dungeon]; Yeti Hide Bracers (19113, -1149.99 DPS) [world] |
| hands | Raider Gloves (272099) | Creeg Bothunk [vendor] | sim-verified (+658.8 DPS vs the runner-up, not corroborated against the finished set) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; General's Dragonhide Gauntlets (231676, -1640.34 DPS) [vendor]; General's Dragonhide Gloves (231677, -1640.34 DPS) [pvp] |
| waist | Shifter's Belt (272396) | Pix Xizzix [vendor] | 51.8 stamina points (10609.00 DPS) | yes | Vosh'gajin's Strand (13962, -1248.66 DPS) [quest]; Frostbite Girdle (14502, -1267.53 DPS) [dungeon]; Hivethrasher's Girdle (275614, -1715.41 DPS) [crafted] |
| legs | Outrider's Leather Pants (22740) | Warsong Outriders [rep] | 84.5 stamina points (17291.92 DPS) | yes | General's Dragonhide Legguards (231673, +0.00 DPS) [vendor]; General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Warstrife Leggings (11821, -2229.61 DPS, sim-verified) [dungeon] |
| feet | Ash Covered Boots (18716) | Stratholme: Archivist Galford [dungeon] | 61.1 stamina points (12507.44 DPS) | yes | Dawn Treaders (19052, +0.00 DPS) [crafted]; General's Dragonhide Treads (231683, +0.00 DPS) [vendor]; General's Dragonhide Greaves (231671, -1107.55 DPS) [vendor] |
| finger1 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | sim-verified (+719.9 DPS vs the runner-up, not corroborated against the finished set) | yes | Innervating Band (18701, +0.00 DPS) [dungeon]; Cutthroat's Signet (272408, +0.00 DPS) [vendor]; Myrmidon's Signet (2246, -719.91 DPS, sim-verified) [world_drop] |
| finger2 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-verified (-775.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Myrmidon's Signet (2246, +0.00 DPS) [world_drop]; Cutthroat's Signet (272408, -338.86 DPS) [vendor]; Innervating Band (18701, -349.53 DPS) [dungeon] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Talisman (20072, -1321.89 DPS) [rep]; Mark of Tyranny (13966, -1366.43 DPS) [quest]; Arena Grand Master (19024, -1366.43 DPS) [quest] |
| trinket2 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Defiler's Talisman (20072, -2368.01 DPS) [rep]; Mark of Tyranny (13966, -2412.55 DPS) [quest]; Arena Grand Master (19024, -2412.55 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Headmaster's Charge (13937, +0.00 DPS) [dungeon]; High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's Destroyer (234546, +0.00 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | Idol of Brutality (23198) | Stratholme: Magistrate Barthilas [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Idol of the Moon (23197, +0.00 DPS) [world_drop] |

**New at 60:** head: Warbear Helm; neck: Talisman of Evasion; shoulder: Defiler's Leather Shoulders; back: Stoneshield Cloak; chest: Warbear Harness; wrist: Forest Stalker's Bracers; hands: Raider Gloves; waist: Shifter's Belt; legs: Outrider's Leather Pants; feet: Ash Covered Boots; finger1: Naglering; finger2: Ring of Awareness; trinket1: Defender's Grip Stabilizer; trinket2: Frostwolf Insignia Rank 6; main_hand: The Unstoppable Force; ranged: Idol of Brutality

No-known-source sample (15 of 1973, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2314 Toughened Leather Armor; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band

