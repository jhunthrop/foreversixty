# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 35.2. Weights run: 3.5s. Verify run: 1.8s. 194 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.083 ± 0.002, defense=0.183 ± 0.035 per rating point (1 rating = 1%, 0.183 per %), dodge=0.129 ± 0.010 per rating point (12 rating = 1%, 1.544 per %), strength=0.070 ± 0.000, agility=0.172 ± 0.006, attack_power=0.035 ± 0.000, hit=0.057 ± 0.006 per rating point (10 rating = 1%, 0.572 per %), crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.321 per %), expertise=1.490 ± 0.069

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (450.67 DPS) | yes | Brawler's Leather Hood (252504, -21.25 DPS) [crafted]; Defender's Leather Hood (252447, -41.12 DPS) [crafted]; Totemic Leather Hood (252448, -54.63 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.0 stamina points (121.28 DPS) | yes | Erudite's Amulet (277204, -28.99 DPS, sim-verified) [quest]; Sentinel's Medallion (20444, -47.73 DPS) [rep]; Scholarly Pendant (277203, -48.51 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (226.27 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -68.08 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -74.72 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (185.92 DPS) | yes | Sentry Cloak (2059, -5.55 DPS) [world_drop]; Grave Shroud (279865, -10.84 DPS) [quest]; Miner's Cape (5444, -20.88 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (441.57 DPS) | yes | Gloomshroud Armor (1489, -27.28 DPS, sim-verified) [dungeon]; Loch Croc Hide Vest (6197, -77.88 DPS) [world]; Tunic of Westfall (2041, -88.59 DPS) [quest] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (196.00 DPS) | yes | Drakewing Bands (12999, -18.65 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -56.08 DPS) [vendor]; Sanguine Cuffs (14375, -64.65 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (220.75 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -2.02 DPS) [crafted]; Defender's Leather Gloves (252496, -9.93 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.3 stamina points (274.46 DPS) | yes | Belt of the Fang (10412, -28.35 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -63.81 DPS) [crafted]; Guardsman Belt (3429, -65.83 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.8 stamina points (406.92 DPS) | yes | Defender's Leather Pants (252445, -46.40 DPS, sim-verified) [crafted]; Slick Deviate Leggings (6480, -68.97 DPS) [quest]; Smelting Pants (5199, -83.65 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (363.77 DPS) | yes | Surfer Shoes (276274, -37.67 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -70.01 DPS) [dungeon]; Trapper's Leather Boots (252440, -98.44 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (145.53 DPS) | yes | Blood Ring (4998, -24.26 DPS) [world_drop]; Ring of the Moon (12052, -71.08 DPS) [world_drop]; Deep Fathom Ring (6463, -72.77 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.4 stamina points (130.18 DPS) | yes | Blood Ring (4998, -8.90 DPS) [world_drop]; Ring of the Moon (12052, -55.72 DPS) [world_drop]; Deep Fathom Ring (6463, -57.41 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (2168.8 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 18.6 stamina points (450.90 DPS) | yes | Rakzur Club (12983, -34.05 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -38.80 DPS) [dungeon]; Twisted Chanter's Staff (890, -63.66 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 194, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 41.6. Weights run: 3.4s. Verify run: 1.8s. 350 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.318 ± 0.057 per rating point (1 rating = 1%, 0.318 per %), dodge=0.139 ± 0.014 per rating point (12 rating = 1%, 1.662 per %), strength=0.065 ± 0.000, agility=0.173 ± 0.009, attack_power=0.032 ± 0.000, hit=0.072 ± 0.008 per rating point (10 rating = 1%, 0.725 per %), crit=0.024 ± 0.001 per rating point (14 rating = 1%, 0.331 per %), expertise=1.823 ± 0.100

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.7 stamina points (495.27 DPS) | yes | Totemic Leather Helm (252456, -39.42 DPS) [crafted]; Trapper's Leather Helm (252513, -39.42 DPS) [crafted]; Defender's Leather Helm (252455, -63.08 DPS, sim-verified) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (296.06 DPS) | yes | Ghostshard Talisman (7731, -63.43 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -85.19 DPS) [world_drop]; Master Sergeant's Insignia (18442, -91.09 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 17.9 stamina points (407.12 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -60.11 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -75.54 DPS) [crafted]; Barbaric Shoulders (5964, -98.82 DPS) [crafted] |
| back | Mourning Shawl (6751) | Mortality Wanes [quest] | 14.0 stamina points (318.94 DPS) | yes | Sergeant's Cape (18440, -42.54 DPS, sim-verified) [vendor]; Tigerstrike Mantle (13108, -80.43 DPS) [world_drop]; Vine Pruner's Cloak (279835, -87.29 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.2 stamina points (482.48 DPS) | yes | Spirewind Fetter (9406, -8.59 DPS) [dungeon]; Green Whelp Armor (7375, -37.94 DPS) [crafted]; Blackened Defias Armor (10399, -67.71 DPS) [dungeon] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (250.50 DPS) | yes | Barbaric Bracers (18948, -2.77 DPS) [crafted]; Black Wolf Bracers (3230, -6.33 DPS) [dungeon]; Staghide Armguards (270021, -37.02 DPS) [quest] |
| hands | Ebon Vise (7690) | Scarlet Monastery: Fallen Champion [dungeon] | 15.1 stamina points (344.91 DPS) | yes | Brawler Gloves (720, -55.74 DPS) [world_drop]; Wolfclaw Gloves (1978, -61.19 DPS) [dungeon]; Operator's Gloves (270045, -76.68 DPS, sim-verified) [quest] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.7 stamina points (357.31 DPS) | yes | Skulker's Leather Belt (252520, -43.60 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -63.89 DPS) [crafted]; Stalker's Leather Belt (252521, -65.35 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.2 stamina points (414.09 DPS) | yes | Brawler's Leather Legguards (252516, -37.66 DPS) [crafted]; Dark Ritual Leggings (270031, -41.73 DPS) [quest]; Duty Bound Leggings (279868, -47.07 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.3 stamina points (417.86 DPS) | yes | Harbinger Boots (7754, -38.27 DPS, sim-verified) [dungeon]; Highlander's Mail Greaves (20123, -63.36 DPS) [vendor]; Lancer Boots (6752, -69.21 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.3 stamina points (188.83 DPS) | yes | Plains Ring (2039, -6.64 DPS) [dungeon]; Ring of Power Regulation (273030, -6.64 DPS) [dungeon]; Boar Signet (274078, -29.41 DPS) [quest] |
| finger2 | Darkspear Signet (272071) (or Ring of Power Regulation (273030), Plains Ring (2039)) | Creeg Bothunk [vendor] | 8.0 stamina points (182.19 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Ring of Power Regulation (273030, +0.00 DPS) [dungeon]; Boar Signet (274078, -22.77 DPS) [quest] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (3392.0 DPS) | yes | Rune of Duty (21568, -45.55 DPS) [rep] |
| trinket2 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (3392.0 DPS) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul: Overlord Ramtusk [dungeon] | sim-verified (3392.0 DPS) | yes | Glimmering Staff (249392, -67.96 DPS) [crafted]; Soulstaff (249393, -69.09 DPS) [crafted]; Manual Crowd Pummeler (9449, -244.04 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenir Sea Shell; shoulder: Watchman Pauldrons; back: Mourning Shawl; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Ebon Vise; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Darkspear Signet; trinket1: Talisman of Arathor; trinket2: Rune of Perfection; main_hand: Staff of the Shade

No-known-source sample (15 of 350, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 73.7. Weights run: 3.9s. Verify run: 2.0s. 461 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.027, armor=0.125 ± 0.007, defense=0.391 ± 0.084 per rating point (1 rating = 1%, 0.391 per %), dodge=0.177 ± 0.019 per rating point (12 rating = 1%, 2.129 per %), strength=0.070 ± 0.000, agility=0.190 ± 0.012, attack_power=0.035 ± 0.000, hit=0.067 ± 0.015 per rating point (10 rating = 1%, 0.666 per %), crit=0.031 ± 0.006 per rating point (14 rating = 1%, 0.430 per %), expertise=2.367 ± 0.167

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman: Shadowforge Relic Hunter [dungeon] | 27.5 stamina points (847.35 DPS) | yes | Warden's Wizard Hat (14604, -49.71 DPS) [world_drop]; Brawler's Leather Helm (252512, -63.56 DPS) [crafted]; Nightscape Headband (8176, -86.46 DPS) [crafted] |
| neck | Shriveled Heart (9243) (or Souvenir Sea Shell (274749)) | Zul'Farrak: Sandarr Dunereaver [dungeon] | 13.0 stamina points (400.83 DPS) | yes | Souvenir Sea Shell (274749, +0.00 DPS) [vendor]; Gazlowe's Charm (13088, -52.97 DPS) [dungeon]; Darkspear Warding Pendant (272074, -61.67 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.4 stamina points (875.29 DPS) | yes | Flintrock Shoulders (7755, -143.80 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -220.82 DPS) [dungeon]; Nightscape Shoulders (8192, -324.57 DPS) [crafted] |
| back | Mourning Shawl (6751) | Mortality Wanes [quest] | 15.0 stamina points (462.63 DPS) | yes | Wing of the Whelpling (13121, -22.73 DPS) [world_drop]; Silky Spider Cape (10776, -38.51 DPS) [dungeon]; Well Oiled Cloak (12254, -42.37 DPS) [vendor] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 31.5 stamina points (971.25 DPS) | yes | Warden's Wraps (14601, -67.59 DPS) [world_drop]; Robes of the Lich (10762, -115.27 DPS) [dungeon]; Insignia Chestguard (4057, -172.85 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.7 stamina points (451.78 DPS) | yes | Branded Leather Bracers (19508, -32.76 DPS) [dungeon]; Barbaric Bracers (18948, -53.19 DPS) [crafted]; Cultist's Armguards (270032, -55.25 DPS) [quest] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 20.4 stamina points (629.58 DPS) | yes | Warden's Leather Gloves (252527, -8.34 DPS) [crafted]; Razzeric's Racing Grips (6727, -34.91 DPS) [quest]; Bonefingers (10765, -39.43 DPS) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 20.2 stamina points (623.97 DPS) | yes | Warden's Leather Belt (252460, -55.02 DPS) [crafted]; Kolkar Hunter's Belt (6788, -58.12 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -78.24 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 26.1 stamina points (806.17 DPS) | yes | Imperial Leather Pants (4062, -27.02 DPS) [dungeon]; Panther Hunter Leggings (4108, -37.72 DPS) [quest]; Warden's Woolies (14605, -38.60 DPS) [world_drop] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 24.3 stamina points (750.60 DPS) | yes | Gnomebot Operating Boots (9450, -99.46 DPS, sim-verified) [dungeon]; Skulker's Leather Shoes (252531, -99.84 DPS) [crafted]; Prowler's Leather Shoes (252465, -114.64 DPS) [crafted] |
| finger1 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.5 stamina points (354.38 DPS) | yes | Darkspear Signet (272070, -15.21 DPS) [vendor]; Underworld Band (1980, -46.05 DPS) [world_drop]; Dragonclaw Ring (10710, -46.05 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.4 stamina points (352.20 DPS) | yes | Darkspear Signet (272070, -13.04 DPS) [vendor]; Underworld Band (1980, -43.87 DPS) [world_drop]; Dragonclaw Ring (10710, -43.87 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | sim-verified (6338.5 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -30.83 DPS) [world]; Darkspear Voodoo Seal (272059, -30.83 DPS) [vendor] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (6338.5 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -30.83 DPS) [world]; Darkspear Voodoo Seal (272059, -30.83 DPS) [vendor] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (6338.5 DPS) | yes | Mograine's Might (7723, -158.68 DPS) [dungeon]; Illusionary Rod (7713, -363.60 DPS) [dungeon]; Manual Crowd Pummeler (9449, -459.67 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Suspicious Spare Part; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 461, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 120.1. Weights run: 4.5s. Verify run: 4.4s. 604 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.023, armor=0.140 ± 0.011, defense=0.573 ± 0.135 per rating point (1 rating = 1%, 0.573 per %), dodge=0.254 ± 0.030 per rating point (12 rating = 1%, 3.047 per %), strength=0.061 ± 0.000, agility=0.248 ± 0.019, attack_power=0.030 ± 0.000, hit=0.124 ± 0.024 per rating point (10 rating = 1%, 1.240 per %), crit=0.043 ± 0.010 per rating point (14 rating = 1%, 0.601 per %), expertise=3.117 ± 0.276

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | sim-verified (11856.2 DPS) | yes | Embrace of the Lycan (9479, -128.11 DPS) [dungeon]; Impractical Headwarmer (274756, -181.71 DPS) [vendor]; Sprightring Helm (17776, -441.68 DPS, sim-verified) [quest] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (548.58 DPS) | yes | Shriveled Heart (9243, -39.18 DPS) [dungeon]; Darkspear Warding Pendant (272073, -39.18 DPS) [vendor]; Souvenir Sea Shell (274749, -39.18 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | sim-verified (11856.2 DPS) | yes | Penance Spaulders (11963, +0.00 DPS) [quest]; Phytoskin Spaulders (17749, -3.58 DPS) [dungeon]; Fleshhide Shoulders (10774, -537.48 DPS, sim-verified) [dungeon] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | 20.5 stamina points (801.37 DPS) | yes | Nightfall Drape (12465, -44.66 DPS) [dungeon]; Sergeant's Cape (18441, -55.62 DPS) [vendor]; Grovekeeper's Drape (17739, -128.51 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 54.2 stamina points (2124.51 DPS) | yes | Mixologist's Tunic (12793, -426.40 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Skin (9473, -528.37 DPS) [dungeon]; Heraldic Breastplate (8119, -591.91 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 24.4 stamina points (957.46 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -58.34 DPS) [vendor]; Warden's Leather Bracers (252541, -179.30 DPS) [crafted]; Serpentskin Bracers (8257, -179.59 DPS) [world_drop] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 39.6 stamina points (1550.17 DPS) | yes | Warden's Leather Gauntlets (252549, -350.47 DPS) [crafted]; Raider Gloves (272100, -416.74 DPS, sim-verified) [vendor]; Feralheart Fists (226793, -421.62 DPS) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 27.2 stamina points (1064.54 DPS) | yes | Skulker's Leather Waistguard (252474, -20.60 DPS) [crafted]; Prowler's Leather Waistguard (252473, -35.30 DPS) [crafted]; Girdle of Beastial Fury (11686, -130.62 DPS) [dungeon] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 36.1 stamina points (1415.21 DPS) | yes | Knight's Crackling Leather Leggings (220864, -217.29 DPS) [vendor]; Knight's Restored Leather Leggings (220882, -265.88 DPS) [vendor]; Windscale Sarong (10842, -669.26 DPS, sim-verified) [world] |
| feet | Shadefiend Boots (11675) | Blackrock Depths: Anub'shiah [dungeon] | 31.6 stamina points (1237.11 DPS) | yes | Skulker's Leather Boots (252469, -162.30 DPS) [crafted]; Prowler's Leather Boots (252468, -177.00 DPS) [crafted]; Warden's Leather Boots (252470, -309.32 DPS, sim-verified) [crafted] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.5 stamina points (566.38 DPS) | yes | Ring of Saviors (1447, -17.81 DPS) [world_drop]; Darkspear Signet (272069, -17.81 DPS) [vendor]; Suspicious Spare Part (274754, -118.74 DPS) [vendor] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (548.58 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Suspicious Spare Part (274754, -100.93 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (11856.2 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Guardian Talisman (1490, -39.18 DPS) [quest]; Rune of Perfection (21565, -78.37 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272060) | Creeg Bothunk [vendor] | sim-verified (11856.2 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Guardian Talisman (1490, -39.18 DPS) [quest]; Rune of Perfection (21565, -78.37 DPS) [rep] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (11856.2 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -24.45 DPS) [world_drop]; Ragehammer (10626, -631.80 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Master Sergeant's Insignia; shoulder: Knight-Lieutenant's Leather Shoulders; back: Graverot Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Shadefiend Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Radiant Staff

No-known-source sample (15 of 604, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 140.1. Weights run: 4.4s. Verify run: 6.5s. 1450 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.025, armor=0.180 ± 0.016, defense=1.329 ± 0.266 per rating point (1 rating = 1%, 1.329 per %), dodge=0.699 ± 0.046 per rating point (12 rating = 1%, 8.389 per %), strength=0.113 ± 0.000, agility=0.554 ± 0.029, attack_power=0.057 ± 0.000, hit=0.201 ± 0.038 per rating point (10 rating = 1%, 2.006 per %), crit=0.092 ± 0.017 per rating point (14 rating = 1%, 1.288 per %), expertise=6.347 ± 0.419

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 73.3 stamina points (5046.44 DPS) | yes | Feralheart Faceguard (226801, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 21.7 stamina points (1490.37 DPS) | yes | Medallion of Grand Marshal Morris (13091, -93.76 DPS) [world_drop]; Sentinel's Medallion (19538, -230.28 DPS) [rep]; Evil Eye Pendant (18381, -278.27 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 70.3 stamina points (4840.06 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -79.04 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -645.39 DPS) [pvp]; Highlander's Leather Shoulders (20059, -936.93 DPS, sim-verified) [rep] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | 40.1 stamina points (2763.25 DPS) | yes | Stoneskin Gargoyle Cape (13397, -262.14 DPS, sim-verified) [dungeon]; Shifting Cloak (18511, -393.40 DPS) [crafted]; Phantasmal Cloak (18689, -502.56 DPS) [dungeon] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-verified (23477.3 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, -1.08 DPS) [vendor]; Tunic of Undead Slaying (23089, -1981.51 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-verified (23477.3 DPS) | yes | Bracers of Subterfuge (22668, -128.30 DPS) [quest]; Forest Stalker's Bracers (19587, -181.53 DPS) [rep]; Wristwraps of Undead Slaying (23093, -816.72 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 57.2 stamina points (3938.64 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -228.46 DPS) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 51.9 stamina points (3569.71 DPS) | yes | Shifter's Belt (272396, -538.03 DPS, sim-verified) [vendor]; Cloudrunner Girdle (13252, -600.61 DPS) [dungeon]; Belt of Preserved Heads (20216, -657.86 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 92.6 stamina points (6377.30 DPS) | yes | Marshal's Dragonhide Leggings (231691, -24.93 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -497.18 DPS) [pvp]; Dire Warbear Woolies (227804, -542.81 DPS, sim-verified) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 69.5 stamina points (4782.87 DPS) | yes | Marshal's Dragonhide Treads (231692, -333.12 DPS) [vendor]; Drudge Boots (21532, -718.54 DPS, sim-verified) [quest]; Feralheart Treads (226803, -787.61 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 34.8 stamina points (2393.51 DPS) | yes | Band of Resolution (22680, -697.11 DPS) [quest]; Ring of Awareness (272409, -819.62 DPS) [vendor]; Band of the Steadfast Hero (22331, -872.69 DPS) [dungeon] |
| finger2 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 25.6 stamina points (1764.21 DPS) | yes | Band of Resolution (22680, -67.82 DPS) [quest]; Ring of Awareness (272409, -190.32 DPS) [vendor]; Band of the Steadfast Hero (22331, -243.40 DPS) [dungeon] |
| trinket1 | Mark of Tyranny (13966) | General Drakkisath's Demise [quest] | sim-verified (23477.3 DPS) | yes | Talisman of Arathor (20071, -1338.65 DPS, sim-verified) [rep]; Defender's Grip Stabilizer (272440, -1614.78 DPS) [vendor]; Stormpike Insignia Rank 6 (17904, -1649.11 DPS) [quest] |
| trinket2 | Vigilance Charm (18370) | Dire Maul: Immol'thar [dungeon] | sim-verified (23477.3 DPS) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Talisman of Arathor (20071, -328.89 DPS) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (23477.3 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Electrified Dagger (19100, -1559.41 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Naglering; trinket1: Mark of Tyranny; trinket2: Vigilance Charm; main_hand: Headmaster's Charge

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 317.6. Weights run: 4.6s. Verify run: 8.8s. 1450 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.009, armor=0.205 ± 0.018, defense=1.860 ± 0.320 per rating point (1 rating = 1%, 1.860 per %), dodge=0.980 ± 0.057 per rating point (12 rating = 1%, 11.755 per %), strength=0.125 ± 0.000, agility=0.845 ± 0.038, attack_power=0.057 ± 0.000, hit=0.256 ± 0.045 per rating point (10 rating = 1%, 2.565 per %), crit=0.128 ± 0.020 per rating point (14 rating = 1%, 1.788 per %), expertise=7.830 ± 0.490

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 84.2 stamina points (7675.17 DPS) | yes | Feralheart Faceguard (226801, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 27.3 stamina points (2488.24 DPS) | yes | Evil Eye Pendant (18381, -146.64 DPS) [dungeon]; Medallion of Grand Marshal Morris (13091, -155.14 DPS) [world_drop]; Talisman of Evasion (13177, -416.20 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 83.5 stamina points (7606.91 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -544.12 DPS) [vendor]; Darkspear Pauldrons (272105, -717.32 DPS, sim-verified) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -1627.83 DPS) [pvp] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | 45.0 stamina points (4096.22 DPS) | yes | Stoneskin Gargoyle Cape (13397, -387.29 DPS) [dungeon]; Windshear Cape (20691, -640.22 DPS) [world]; Shifting Cloak (18511, -840.00 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-verified (43961.5 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, -942.65 DPS) [vendor]; Tunic of Undead Slaying (23089, -3196.93 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-verified (43961.5 DPS) | yes | Bracers of Subterfuge (22668, -42.93 DPS) [quest]; Forest Stalker's Bracers (19589, -468.05 DPS) [rep] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 69.2 stamina points (6306.72 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Raider Gloves (272099, -437.67 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -1088.42 DPS) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 60.9 stamina points (5550.88 DPS) | yes | Shifter's Belt (272396, -737.52 DPS, sim-verified) [vendor]; Cloudrunner Girdle (13252, -780.10 DPS) [dungeon]; Belt of Preserved Heads (20216, -982.53 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 106.6 stamina points (9713.54 DPS) | yes | Marshal's Dragonhide Leggings (231691, -358.45 DPS) [vendor]; Dire Warbear Woolies (227804, -616.05 DPS, sim-verified) [vendor]; Feralheart Legguards (226799, -1498.44 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 83.1 stamina points (7569.20 DPS) | yes | Drudge Boots (21532, -625.67 DPS, sim-verified) [quest]; Marshal's Dragonhide Treads (231692, -1015.35 DPS) [vendor]; Feralheart Treads (226803, -1642.99 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 38.7 stamina points (3521.97 DPS) | yes | Ring of Awareness (272409, -993.72 DPS) [vendor]; Band of Resolution (22680, -1034.14 DPS) [quest]; Band of the Steadfast Hero (22331, -1161.94 DPS) [dungeon] |
| finger2 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 29.6 stamina points (2692.92 DPS) | yes | Band of Resolution (22680, -205.10 DPS) [quest]; Band of the Steadfast Hero (22331, -332.90 DPS) [dungeon]; Ring of Awareness (272409, -454.69 DPS, sim-verified) [vendor] |
| trinket1 | Mark of Tyranny (13966) | General Drakkisath's Demise [quest] | sim-verified (43961.5 DPS) | yes | Defender's Grip Stabilizer (272440, -2229.96 DPS) [vendor]; Stormpike Insignia Rank 6 (17904, -2291.56 DPS) [quest]; Force of Will (11810, -3247.11 DPS) [dungeon] |
| trinket2 | Vigilance Charm (18370) | Dire Maul: Immol'thar [dungeon] | sim-verified (43961.5 DPS) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Force of Will (11810, -955.55 DPS) [dungeon] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (43961.5 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; The Lobotomizer (19324, -2102.21 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Naglering; trinket1: Mark of Tyranny; trinket2: Vigilance Charm; main_hand: Headmaster's Charge

No-known-source sample (15 of 1450, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 3.5s. Verify run: 1.8s. 184 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.083 ± 0.002, defense=0.183 ± 0.035 per rating point (1 rating = 1%, 0.183 per %), dodge=0.129 ± 0.010 per rating point (12 rating = 1%, 1.544 per %), strength=0.070 ± 0.000, agility=0.172 ± 0.006, attack_power=0.035 ± 0.000, hit=0.057 ± 0.006 per rating point (10 rating = 1%, 0.572 per %), crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.321 per %), expertise=1.490 ± 0.069

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (450.67 DPS) | yes | Brawler's Leather Hood (252504, -21.25 DPS) [crafted]; Defender's Leather Hood (252447, -41.12 DPS) [crafted]; Totemic Leather Hood (252448, -54.63 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (81.11 DPS) | yes | Scout's Medallion (20442, -7.56 DPS) [rep]; Scholarly Pendant (277203, -8.35 DPS) [quest] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (226.27 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -68.08 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -74.72 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (185.92 DPS) | yes | Sentry Cloak (2059, -5.55 DPS) [world_drop]; Grave Shroud (279865, -10.84 DPS) [quest]; Miner's Cape (5444, -20.88 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (441.57 DPS) | yes | Gloomshroud Armor (1489, -38.44 DPS, sim-verified) [dungeon]; Loch Croc Hide Vest (6197, -77.88 DPS) [world]; Trapper's Leather Armor (252491, -119.42 DPS) [crafted] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.8 stamina points (214.58 DPS) | yes | Bear Bracers (4795, -17.61 DPS, sim-verified) [vendor]; Savannah Bracers (15453, -20.60 DPS) [quest]; Drakewing Bands (12999, -59.01 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (220.75 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -2.02 DPS) [crafted]; Defender's Leather Gloves (252496, -9.93 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.3 stamina points (274.46 DPS) | yes | Belt of the Fang (10412, -31.29 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -63.81 DPS) [crafted]; Guardsman Belt (3429, -65.83 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.1 stamina points (343.03 DPS) | yes | Slick Deviate Leggings (6480, -5.08 DPS) [quest]; Smelting Pants (5199, -19.76 DPS) [dungeon]; Brawler's Leather Pants (252500, -32.72 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (363.77 DPS) | yes | Surfer Shoes (276274, -52.52 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -70.01 DPS) [dungeon]; Trapper's Leather Boots (252440, -98.44 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.4 stamina points (130.18 DPS) | yes | Ring of Scorn (3235, -33.16 DPS) [quest]; Ring of the Moon (12052, -55.72 DPS) [world_drop]; Deep Fathom Ring (6463, -57.41 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (121.28 DPS) | yes | Ring of Scorn (3235, -16.75 DPS, sim-verified) [quest]; Ring of the Moon (12052, -46.82 DPS) [world_drop]; Deep Fathom Ring (6463, -48.51 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (2229.7 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 18.6 stamina points (450.90 DPS) | yes | Hammerbone (270018, -22.51 DPS) [quest]; Crescent Staff (6505, -23.24 DPS, sim-verified) [quest]; Advisor's Gnarled Staff (20425, -24.37 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 184, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 42.1. Weights run: 3.4s. Verify run: 1.8s. 342 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.318 ± 0.057 per rating point (1 rating = 1%, 0.318 per %), dodge=0.139 ± 0.014 per rating point (12 rating = 1%, 1.662 per %), strength=0.065 ± 0.000, agility=0.173 ± 0.009, attack_power=0.032 ± 0.000, hit=0.072 ± 0.008 per rating point (10 rating = 1%, 0.725 per %), crit=0.024 ± 0.001 per rating point (14 rating = 1%, 0.331 per %), expertise=1.823 ± 0.100

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.7 stamina points (495.27 DPS) | yes | Defender's Leather Helm (252455, -25.98 DPS, sim-verified) [crafted]; Totemic Leather Helm (252456, -39.42 DPS) [crafted]; Trapper's Leather Helm (252513, -39.42 DPS) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (296.06 DPS) | yes | Ghostshard Talisman (7731, -65.37 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -85.19 DPS) [world_drop]; Senior Sergeant's Insignia (15200, -91.09 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 17.9 stamina points (407.12 DPS) | yes | Tanned Shoulderpads (270023, -57.56 DPS, sim-verified) [quest]; Cloudy Gustwoven Spaulders (277043, -75.54 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -75.54 DPS) [crafted] |
| back | Grimsteel Cape (4643) | Vorrel's Revenge [quest] | 11.7 stamina points (266.25 DPS) | yes | Sergeant's Cloak (18427, -11.83 DPS) [vendor]; Tigerstrike Mantle (13108, -27.74 DPS) [world_drop]; Enduring Cape (14763, -54.15 DPS) [world_drop] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.2 stamina points (482.48 DPS) | yes | Spirewind Fetter (9406, -8.59 DPS) [dungeon]; Green Whelp Armor (7375, -37.94 DPS) [crafted]; Blackened Defias Armor (10399, -67.71 DPS) [dungeon] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (250.50 DPS) | yes | Barbaric Bracers (18948, -2.77 DPS) [crafted]; Black Wolf Bracers (3230, -6.33 DPS) [dungeon]; Technician's Bracers (270042, -39.84 DPS) [quest] |
| hands | Ebon Vise (7690) | Scarlet Monastery: Fallen Champion [dungeon] | 15.1 stamina points (344.91 DPS) | yes | Braced Handguards (6784, -45.36 DPS, sim-verified) [quest]; Brawler Gloves (720, -55.74 DPS) [world_drop]; Wolfclaw Gloves (1978, -61.19 DPS) [dungeon] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.7 stamina points (357.31 DPS) | yes | Skulker's Leather Belt (252520, -49.94 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -63.89 DPS) [crafted]; Stalker's Leather Belt (252521, -65.35 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.2 stamina points (414.09 DPS) | yes | Brawler's Leather Legguards (252516, -32.11 DPS, sim-verified) [crafted]; Dark Ritual Leggings (270031, -41.73 DPS) [quest]; Trapper's Leather Legguards (252517, -47.99 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.3 stamina points (417.86 DPS) | yes | Gravewalker Boots (251963, -14.15 DPS) [quest]; Harbinger Boots (7754, -29.33 DPS) [dungeon]; Grizzled Boots (6335, -56.34 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.3 stamina points (188.83 DPS) | yes | Plains Ring (2039, -6.64 DPS) [dungeon]; Darkspear Signet (272071, -6.64 DPS) [vendor]; Ring of Power Regulation (273030, -6.64 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.2 stamina points (186.62 DPS) | yes | Plains Ring (2039, -4.43 DPS) [dungeon]; Darkspear Signet (272071, -4.43 DPS) [vendor]; Ring of Power Regulation (273030, -4.43 DPS) [dungeon] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | sim-verified (3533.5 DPS) | yes | Rune of Duty (21568, -45.55 DPS) [rep] |
| trinket2 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (3533.5 DPS) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul: Overlord Ramtusk [dungeon] | sim-verified (3533.5 DPS) | yes | Advisor's Gnarled Staff (19569, -34.80 DPS) [pvp]; Glimmering Staff (249392, -67.96 DPS) [crafted]; Manual Crowd Pummeler (9449, -237.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenir Sea Shell; shoulder: Watchman Pauldrons; back: Grimsteel Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Ebon Vise; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Defiler's Talisman; trinket2: Rune of Perfection; main_hand: Staff of the Shade

No-known-source sample (15 of 342, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 73.3. Weights run: 3.9s. Verify run: 2.0s. 446 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.027, armor=0.125 ± 0.007, defense=0.391 ± 0.084 per rating point (1 rating = 1%, 0.391 per %), dodge=0.177 ± 0.019 per rating point (12 rating = 1%, 2.129 per %), strength=0.070 ± 0.000, agility=0.190 ± 0.012, attack_power=0.035 ± 0.000, hit=0.067 ± 0.015 per rating point (10 rating = 1%, 0.666 per %), crit=0.031 ± 0.006 per rating point (14 rating = 1%, 0.430 per %), expertise=2.367 ± 0.167

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman: Shadowforge Relic Hunter [dungeon] | 27.5 stamina points (847.35 DPS) | yes | Warden's Wizard Hat (14604, -49.71 DPS) [world_drop]; Brawler's Leather Helm (252512, -63.56 DPS) [crafted]; Nightscape Headband (8176, -86.46 DPS) [crafted] |
| neck | Shriveled Heart (9243) (or Souvenir Sea Shell (274749)) | Zul'Farrak: Sandarr Dunereaver [dungeon] | 13.0 stamina points (400.83 DPS) | yes | Souvenir Sea Shell (274749, +0.00 DPS) [vendor]; Dragon's Blood Necklace (10711, -30.83 DPS) [quest]; Gazlowe's Charm (13088, -52.97 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.4 stamina points (875.29 DPS) | yes | Flintrock Shoulders (7755, -120.70 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -220.82 DPS) [dungeon]; Tanned Shoulderpads (270023, -316.02 DPS) [quest] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 14.3 stamina points (439.90 DPS) | yes | Silky Spider Cape (10776, -15.77 DPS) [dungeon]; Well Oiled Cloak (12254, -19.63 DPS) [vendor]; Heraldic Cloak (8120, -23.11 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 31.5 stamina points (971.25 DPS) | yes | Warden's Wraps (14601, -67.59 DPS) [world_drop]; Robes of the Lich (10762, -115.27 DPS) [dungeon]; Cultist's Chestguard (270054, -156.57 DPS) [quest] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.7 stamina points (451.78 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -32.76 DPS) [dungeon]; Barbaric Bracers (18948, -53.19 DPS) [crafted] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 20.4 stamina points (629.58 DPS) | yes | Warden's Leather Gloves (252527, -8.34 DPS) [crafted]; Razzeric's Racing Grips (6727, -34.91 DPS) [quest]; Bonefingers (10765, -39.43 DPS) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 20.2 stamina points (623.97 DPS) | yes | Warden's Leather Belt (252460, -55.02 DPS) [crafted]; Kolkar Hunter's Belt (6788, -58.59 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -78.24 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 26.1 stamina points (806.17 DPS) | yes | Tidesoaked Leggings (279888, -9.64 DPS) [quest]; Imperial Leather Pants (4062, -27.02 DPS) [dungeon]; Panther Hunter Leggings (4108, -37.72 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 24.3 stamina points (750.60 DPS) | yes | Skulker's Leather Shoes (252531, -99.84 DPS) [crafted]; Gravewalker Boots (251963, -100.09 DPS) [quest]; Gnomebot Operating Boots (9450, -126.51 DPS, sim-verified) [dungeon] |
| finger1 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.5 stamina points (354.38 DPS) | yes | Darkspear Signet (272070, -15.21 DPS) [vendor]; Underworld Band (1980, -46.05 DPS) [world_drop]; Dragonclaw Ring (10710, -46.05 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.4 stamina points (352.20 DPS) | yes | Darkspear Signet (272070, -13.04 DPS) [vendor]; Underworld Band (1980, -43.87 DPS) [world_drop]; Dragonclaw Ring (10710, -43.87 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | sim-verified (6835.2 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -30.83 DPS) [world]; Darkspear Voodoo Seal (272059, -30.83 DPS) [vendor] |
| trinket2 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (6835.2 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -30.83 DPS) [world]; Darkspear Voodoo Seal (272059, -30.83 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (6835.2 DPS) | yes | Ironshod Bludgeon (9408, -15.42 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -168.72 DPS) [quest]; Manual Crowd Pummeler (9449, -620.60 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Suspicious Spare Part; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Cragwood Maul

No-known-source sample (15 of 446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 117.6. Weights run: 4.5s. Verify run: 2.3s. 585 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.023, armor=0.140 ± 0.011, defense=0.573 ± 0.135 per rating point (1 rating = 1%, 0.573 per %), dodge=0.254 ± 0.030 per rating point (12 rating = 1%, 3.047 per %), strength=0.061 ± 0.000, agility=0.248 ± 0.019, attack_power=0.030 ± 0.000, hit=0.124 ± 0.024 per rating point (10 rating = 1%, 1.240 per %), crit=0.043 ± 0.010 per rating point (14 rating = 1%, 0.601 per %), expertise=3.117 ± 0.276

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 41.1 stamina points (1609.55 DPS) | yes | Blood Guard's Leather Headband (220851, -151.21 DPS) [vendor]; Embrace of the Lycan (9479, -321.44 DPS, sim-verified) [dungeon]; Impractical Headwarmer (274756, -332.92 DPS) [vendor] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (548.58 DPS) | yes | Shriveled Heart (9243, -39.18 DPS) [dungeon]; Darkspear Warding Pendant (272073, -39.18 DPS) [vendor]; Souvenir Sea Shell (274749, -39.18 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 30.1 stamina points (1178.31 DPS) | yes | Penance Spaulders (11963, -10.44 DPS) [quest]; Blood Guard's Leather Shoulders (220853, -19.34 DPS) [vendor]; Phytoskin Spaulders (17749, -22.92 DPS) [dungeon] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | 20.5 stamina points (801.37 DPS) | yes | Nightfall Drape (12465, -44.66 DPS) [dungeon]; Sergeant's Cloak (16341, -55.62 DPS) [vendor]; Grovekeeper's Drape (17739, -128.51 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 54.2 stamina points (2124.51 DPS) | yes | Mixologist's Tunic (12793, -410.68 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Skin (9473, -528.37 DPS) [dungeon]; Heraldic Breastplate (8119, -591.91 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 24.4 stamina points (957.46 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -58.34 DPS) [vendor]; Forest Stalker's Bracers (19589, -85.71 DPS) [pvp]; Warden's Leather Bracers (252541, -179.30 DPS) [crafted] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 39.6 stamina points (1550.17 DPS) | yes | Raider Gloves (272100, -270.74 DPS, sim-verified) [vendor]; Warden's Leather Gauntlets (252549, -350.47 DPS) [crafted]; Feralheart Fists (226793, -421.62 DPS) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 27.2 stamina points (1064.54 DPS) | yes | Skulker's Leather Waistguard (252474, -20.60 DPS) [crafted]; Prowler's Leather Waistguard (252473, -35.30 DPS) [crafted]; Girdle of Beastial Fury (11686, -130.62 DPS) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 36.1 stamina points (1415.21 DPS) | yes | Windscale Sarong (10842, -205.80 DPS, sim-verified) [world]; Stone Guard's Crackling Leather Leggings (220865, -217.29 DPS) [vendor]; Stone Guard's Restored Leather Leggings (220883, -265.88 DPS) [vendor] |
| feet | Shadefiend Boots (11675) | Blackrock Depths: Anub'shiah [dungeon] | 31.6 stamina points (1237.11 DPS) | yes | Warden's Leather Boots (252470, -142.92 DPS, sim-verified) [crafted]; Skulker's Leather Boots (252469, -162.30 DPS) [crafted]; Prowler's Leather Boots (252468, -177.00 DPS) [crafted] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.5 stamina points (566.38 DPS) | yes | Ring of Saviors (1447, -17.81 DPS) [world_drop]; Darkspear Signet (272069, -17.81 DPS) [vendor]; Suspicious Spare Part (274754, -118.74 DPS) [vendor] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (548.58 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Suspicious Spare Part (274754, -100.93 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (12187.4 DPS) | yes | Defiler's Talisman (21115, +0.00 DPS) [rep]; Guardian Talisman (1490, -39.18 DPS) [quest]; Rune of Perfection (21565, -78.37 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272060) | Creeg Bothunk [vendor] | sim-verified (12187.4 DPS) | yes | Defiler's Talisman (21115, +0.00 DPS) [rep]; Guardian Talisman (1490, -39.18 DPS) [quest]; Rune of Perfection (21565, -78.37 DPS) [rep] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (12187.4 DPS) | yes | Advisor's Gnarled Staff (19567, +0.00 DPS) [pvp]; Dreamstaff (249454, +0.00 DPS) [crafted]; Shadowblade (2163, -682.03 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Graverot Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Shadefiend Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Radiant Staff

No-known-source sample (15 of 585, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 141.6. Weights run: 4.4s. Verify run: 6.4s. 1444 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.025, armor=0.180 ± 0.016, defense=1.329 ± 0.266 per rating point (1 rating = 1%, 1.329 per %), dodge=0.699 ± 0.046 per rating point (12 rating = 1%, 8.389 per %), strength=0.113 ± 0.000, agility=0.554 ± 0.029, attack_power=0.057 ± 0.000, hit=0.201 ± 0.038 per rating point (10 rating = 1%, 2.006 per %), crit=0.092 ± 0.017 per rating point (14 rating = 1%, 1.288 per %), expertise=6.347 ± 0.419

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 73.3 stamina points (5046.44 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Feralheart Faceguard (226801, -341.45 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 21.7 stamina points (1490.37 DPS) | yes | Scout's Medallion (19534, -230.28 DPS) [rep]; Evil Eye Pendant (18381, -278.27 DPS) [dungeon]; Medallion of Grand Marshal Morris (13091, -325.97 DPS, sim-verified) [world_drop] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 70.3 stamina points (4840.06 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -79.04 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, -645.39 DPS) [vendor]; Defiler's Leather Shoulders (20194, -891.39 DPS, sim-verified) [rep] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | 40.1 stamina points (2763.25 DPS) | yes | Shifting Cloak (18511, -393.40 DPS) [crafted]; Phantasmal Cloak (18689, -502.56 DPS) [dungeon]; Stoneskin Gargoyle Cape (13397, -558.30 DPS, sim-verified) [dungeon] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-verified (25248.2 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -1.08 DPS) [vendor]; Tunic of Undead Slaying (23089, -2199.71 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-verified (25248.2 DPS) | yes | Bracers of Subterfuge (22668, -128.30 DPS) [quest]; Forest Stalker's Bracers (19587, -181.53 DPS) [rep]; Wristwraps of Undead Slaying (23093, -827.35 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 57.2 stamina points (3938.64 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, -228.46 DPS) [pvp]; Raider Gloves (272099, -440.14 DPS, sim-verified) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 51.9 stamina points (3569.71 DPS) | yes | Cloudrunner Girdle (13252, -600.61 DPS) [dungeon]; Shifter's Belt (272396, -613.61 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -657.86 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 92.6 stamina points (6377.30 DPS) | yes | General's Dragonhide Leggings (231685, -24.93 DPS) [pvp]; General's Dragonhide Legguards (231673, -497.18 DPS) [vendor]; Dire Warbear Woolies (227804, -590.10 DPS, sim-verified) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 69.5 stamina points (4782.87 DPS) | yes | General's Dragonhide Treads (231683, -333.12 DPS) [vendor]; Drudge Boots (21532, -732.00 DPS, sim-verified) [quest]; Feralheart Treads (226803, -787.61 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | sim-verified (25248.2 DPS) | yes | Naglering (11669, -641.44 DPS, sim-verified) [dungeon]; Band of Resolution (22680, -697.11 DPS) [quest]; Ring of Awareness (272409, -819.62 DPS) [vendor] |
| finger2 | Thrall's Resolve (12544) | The Princess Saved? [quest] | sim-verified (25248.2 DPS) | yes | Naglering (11669, -637.04 DPS, sim-verified) [dungeon]; Band of Resolution (22680, -672.03 DPS) [quest]; Ring of Awareness (272409, -794.53 DPS) [vendor] |
| trinket1 | Mark of Tyranny (13966) | For The Horde! [quest] | sim-verified (25248.2 DPS) | yes | Defender's Grip Stabilizer (272440, -1614.78 DPS) [vendor]; Frostwolf Insignia Rank 6 (17909, -1649.11 DPS) [quest]; Defiler's Talisman (20072, -1978.00 DPS) [rep] |
| trinket2 | Vigilance Charm (18370) | Dire Maul: Immol'thar [dungeon] | sim-verified (25248.2 DPS) | yes | Frostwolf Insignia Rank 6 (17909, +0.00 DPS) [quest]; Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Defiler's Talisman (20072, -328.89 DPS) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (25248.2 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glacial Blade (19099, -1764.21 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Thrall's Resolve; trinket1: Mark of Tyranny; trinket2: Vigilance Charm; main_hand: Headmaster's Charge

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 324.1. Weights run: 4.6s. Verify run: 6.8s. 1444 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.009, armor=0.205 ± 0.018, defense=1.860 ± 0.320 per rating point (1 rating = 1%, 1.860 per %), dodge=0.980 ± 0.057 per rating point (12 rating = 1%, 11.755 per %), strength=0.125 ± 0.000, agility=0.845 ± 0.038, attack_power=0.057 ± 0.000, hit=0.256 ± 0.045 per rating point (10 rating = 1%, 2.565 per %), crit=0.128 ± 0.020 per rating point (14 rating = 1%, 1.788 per %), expertise=7.830 ± 0.490

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 84.2 stamina points (7675.17 DPS) | yes | Feralheart Faceguard (226801, +0.00 DPS) [vendor]; Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 27.3 stamina points (2488.24 DPS) | yes | Evil Eye Pendant (18381, -146.64 DPS) [dungeon]; Medallion of Grand Marshal Morris (13091, -155.14 DPS) [world_drop]; Talisman of Evasion (13177, -416.20 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 83.5 stamina points (7606.91 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -544.12 DPS) [vendor]; Darkspear Pauldrons (272105, -1152.43 DPS, sim-verified) [vendor]; Warlord's Dragonhide Spaulders (231681, -1627.83 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | 45.0 stamina points (4096.22 DPS) | yes | Shifting Cloak (18511, -90.51 DPS) [crafted]; Stoneskin Gargoyle Cape (13397, -387.29 DPS) [dungeon]; Windshear Cape (20691, -640.22 DPS) [world] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -942.65 DPS) [vendor]; Tunic of Undead Slaying (23089, -2597.71 DPS, sim-verified) [world] |
| wrist | Forest Stalker's Bracers (19587) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | Feralheart Wristguards (226796, -5.38 DPS) [vendor]; Bracers of Subterfuge (22668, -48.31 DPS) [quest]; Wristwraps of Undead Slaying (23093, -932.84 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 69.2 stamina points (6306.72 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, -437.67 DPS) [vendor]; General's Dragonhide Gloves (231677, -1088.42 DPS) [pvp] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 60.9 stamina points (5550.88 DPS) | yes | Shifter's Belt (272396, -598.01 DPS) [vendor]; Cloudrunner Girdle (13252, -780.10 DPS) [dungeon]; Belt of Preserved Heads (20216, -982.53 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 106.6 stamina points (9713.54 DPS) | yes | General's Dragonhide Leggings (231685, -358.45 DPS) [pvp]; Dire Warbear Woolies (227804, -436.66 DPS) [vendor]; Feralheart Legguards (226799, -1498.44 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 83.1 stamina points (7569.20 DPS) | yes | Drudge Boots (21532, -456.73 DPS, sim-verified) [quest]; General's Dragonhide Treads (231683, -1015.35 DPS) [vendor]; Feralheart Treads (226803, -1642.99 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, -829.04 DPS) [dungeon]; Ring of Awareness (272409, -993.72 DPS) [vendor]; Band of Resolution (22680, -1034.14 DPS) [quest] |
| finger2 | Thrall's Resolve (12544) | The Princess Saved? [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, -792.81 DPS) [dungeon]; Ring of Awareness (272409, -957.49 DPS) [vendor]; Band of Resolution (22680, -997.91 DPS) [quest] |
| trinket1 | Mark of Tyranny (13966) | For The Horde! [quest] | sim-verified (+1817.5 DPS vs the runner-up, not corroborated against the finished set) | yes | Frostwolf Insignia Rank 6 (17909, -2291.56 DPS) [quest]; Vigilance Charm (18370, -2291.56 DPS) [dungeon]; Force of Will (11810, -3247.11 DPS) [dungeon] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (46760.2 DPS) | yes | Frostwolf Insignia Rank 6 (17909, -61.61 DPS) [quest]; Vigilance Charm (18370, -505.72 DPS, sim-verified) [dungeon]; Force of Will (11810, -1017.15 DPS) [dungeon] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glacial Blade (19099, -1216.15 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Forest Stalker's Bracers; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Thrall's Resolve; trinket1: Mark of Tyranny; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1444, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

