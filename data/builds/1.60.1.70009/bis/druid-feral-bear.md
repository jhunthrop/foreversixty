# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 34.1. Weights run: 3.5s. Verify run: 1.3s. 193 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.083 ± 0.002, defense=0.169 ± 0.034 per rating point (1 rating = 1%, 0.169 per %), dodge=0.129 ± 0.010 per rating point (12 rating = 1%, 1.547 per %), strength=0.070 ± 0.000, agility=0.152 ± 0.006, attack_power=0.035 ± 0.000, hit=0.066 ± 0.006 per rating point (10 rating = 1%, 0.659 per %), crit=0.024 ± 0.002 per rating point (14 rating = 1%, 0.340 per %), expertise=1.471 ± 0.067

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (432.17 DPS) | yes | Brawler's Leather Hood (252504, -24.18 DPS) [crafted]; Defender's Leather Hood (252447, -39.55 DPS) [crafted]; Totemic Leather Hood (252448, -52.51 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.6 stamina points (153.77 DPS) | yes | Scholarly Pendant (277203, -14.17 DPS) [quest]; Tarnished Locket (279870, -37.43 DPS) [quest]; Sentinel's Medallion (20444, -85.98 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (216.85 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -67.62 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -71.56 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (178.28 DPS) | yes | Sentry Cloak (2059, -7.17 DPS) [world_drop]; Grave Shroud (279865, -11.32 DPS) [quest]; Miner's Cape (5444, -20.03 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.1 stamina points (421.90 DPS) | yes | Gloomshroud Armor (1489, -22.29 DPS, sim-verified) [dungeon]; Loch Croc Hide Vest (6197, -73.31 DPS) [world]; Tunic of Westfall (2041, -88.68 DPS) [quest] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (187.89 DPS) | yes | Drakewing Bands (12999, -38.80 DPS) [world_drop]; Wolf Bracers (4794, -55.63 DPS) [vendor]; Sanguine Cuffs (14375, -61.95 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.0 stamina points (209.74 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.93 DPS) [crafted]; Defender's Leather Gloves (252496, -7.69 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.2 stamina points (260.81 DPS) | yes | Belt of the Fang (10412, -28.28 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -60.74 DPS) [crafted]; Guardsman Belt (3429, -62.67 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.6 stamina points (387.31 DPS) | yes | Defender's Leather Pants (252445, -46.72 DPS, sim-verified) [crafted]; Slick Deviate Leggings (6480, -65.22 DPS) [quest]; Smelting Pants (5199, -77.45 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (348.83 DPS) | yes | Surfer Shoes (276274, -36.27 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -70.00 DPS) [dungeon]; Trapper's Leather Boots (252440, -96.81 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (139.60 DPS) | yes | Blood Ring (4998, -23.27 DPS) [world_drop]; Ring of the Moon (12052, -68.18 DPS) [world_drop]; Deep Fathom Ring (6463, -69.80 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.3 stamina points (124.20 DPS) | yes | Blood Ring (4998, -7.87 DPS) [world_drop]; Ring of the Moon (12052, -52.78 DPS) [world_drop]; Deep Fathom Ring (6463, -54.40 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (2100.1 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 18.6 stamina points (432.47 DPS) | yes | Rakzur Club (12983, -32.66 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -37.22 DPS) [dungeon]; Twisted Chanter's Staff (890, -61.07 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 40.5. Weights run: 3.5s. Verify run: 1.3s. 322 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.317 ± 0.056 per rating point (1 rating = 1%, 0.317 per %), dodge=0.139 ± 0.014 per rating point (12 rating = 1%, 1.667 per %), strength=0.065 ± 0.000, agility=0.161 ± 0.008, attack_power=0.032 ± 0.000, hit=0.073 ± 0.008 per rating point (10 rating = 1%, 0.732 per %), crit=0.024 ± 0.001 per rating point (14 rating = 1%, 0.337 per %), expertise=1.835 ± 0.100

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 stamina points (473.17 DPS) | yes | Totemic Leather Helm (252456, -35.10 DPS) [crafted]; Trapper's Leather Helm (252513, -35.10 DPS) [crafted]; Defender's Leather Helm (252455, -43.29 DPS, sim-verified) [crafted] |
| neck | Souvenier Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (284.23 DPS) | yes | Ghostshard Talisman (7731, -60.80 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -81.79 DPS) [world_drop]; Master Sergeant's Insignia (18442, -87.46 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 17.9 stamina points (391.27 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -65.97 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -74.47 DPS) [crafted]; Barbaric Shoulders (5964, -96.26 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 7 · Sergeant · Alliance [vendor] | 11.2 stamina points (244.39 DPS) | yes | Tigerstrike Mantle (13108, -25.76 DPS, sim-verified) [world_drop]; Enduring Cape (14763, -41.20 DPS) [world_drop]; Watch Master's Cloak (2953, -49.22 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.2 stamina points (463.71 DPS) | yes | Spirewind Fetter (9406, -8.16 DPS) [dungeon]; Green Whelp Armor (7375, -36.40 DPS) [crafted]; Dokebi Chestguard (14581, -65.59 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (240.72 DPS) | yes | Barbaric Bracers (18948, -3.75 DPS) [crafted]; Black Wolf Bracers (3230, -6.08 DPS) [dungeon]; Staghide Armguards (270021, -36.93 DPS) [quest] |
| hands | Ebon Vise (7690) | Scarlet Monastery: Fallen Champion [dungeon] | 15.1 stamina points (329.84 DPS) | yes | Brawler Gloves (720, -51.90 DPS) [world_drop]; Operator's Gloves (270045, -57.73 DPS, sim-verified) [quest]; Wolfclaw Gloves (1978, -58.79 DPS) [dungeon] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.6 stamina points (341.72 DPS) | yes | Prowler's Leather Belt (252459, -61.34 DPS) [crafted]; Skulker's Leather Belt (252520, -63.55 DPS, sim-verified) [crafted]; Stalker's Leather Belt (252521, -63.57 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.1 stamina points (396.10 DPS) | yes | Brawler's Leather Legguards (252516, -36.71 DPS) [crafted]; Dark Ritual Leggings (270031, -38.15 DPS) [quest]; Duty Bound Leggings (279868, -58.08 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.4 stamina points (401.56 DPS) | yes | Harbinger Boots (7754, -29.32 DPS, sim-verified) [dungeon]; Highlander's Mail Greaves (20123, -63.04 DPS) [vendor]; Lancer Boots (6752, -68.38 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.3 stamina points (181.29 DPS) | yes | Plains Ring (2039, -6.38 DPS) [dungeon]; Sustaining Ring (6743, -50.11 DPS) [quest]; Defias Renegade Ring (1076, -54.42 DPS) [dungeon] |
| finger2 | Darkspear Signet (272071) (or Plains Ring (2039)) | Creeg Bothunk [vendor] | 8.0 stamina points (174.91 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Sustaining Ring (6743, -43.73 DPS) [quest]; Defias Renegade Ring (1076, -48.04 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (3239.3 DPS) | yes | Talisman of Arathor (21119, +0.00 DPS) [rep]; Rune of Perfection (21566, -43.73 DPS) [rep]; Rune of Duty (21568, -43.73 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272062) | Creeg Bothunk [vendor] | sim-verified (3239.3 DPS) | yes | Talisman of Arathor (21119, +0.00 DPS) [rep]; Rune of Perfection (21566, -43.73 DPS) [rep]; Rune of Duty (21568, -43.73 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul: Overlord Ramtusk [dungeon] | sim-verified (3239.3 DPS) | yes | Glimmering Staff (249392, -65.23 DPS) [crafted]; Soulstaff (249393, -66.32 DPS) [crafted]; Manual Crowd Pummeler (9449, -234.91 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenier Sea Shell; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Ebon Vise; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Staff of the Shade

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 71.5. Weights run: 4.0s. Verify run: 1.4s. 438 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.025, armor=0.128 ± 0.007, defense=0.466 ± 0.082 per rating point (1 rating = 1%, 0.466 per %), dodge=0.176 ± 0.020 per rating point (12 rating = 1%, 2.117 per %), strength=0.072 ± 0.000, agility=0.183 ± 0.012, attack_power=0.036 ± 0.000, hit=0.072 ± 0.015 per rating point (10 rating = 1%, 0.722 per %), crit=0.031 ± 0.006 per rating point (14 rating = 1%, 0.430 per %), expertise=2.407 ± 0.170

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman: Shadowforge Relic Hunter [dungeon] | 27.6 stamina points (801.63 DPS) | yes | Warden's Wizard Hat (14604, -47.31 DPS) [world_drop]; Brawler's Leather Helm (252512, -60.73 DPS) [crafted]; Nightscape Headband (8176, -82.22 DPS) [crafted] |
| neck | Shriveled Heart (9243) (or Souvenier Sea Shell (274749)) | Zul'Farrak: Sandarr Dunereaver [dungeon] | 13.0 stamina points (376.98 DPS) | yes | Souvenier Sea Shell (274749, +0.00 DPS) [vendor]; Gazlowe's Charm (13088, -49.64 DPS) [dungeon]; Darkspear Warding Pendant (272074, -58.00 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.6 stamina points (828.89 DPS) | yes | Flintrock Shoulders (7755, -96.33 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -207.47 DPS) [dungeon]; Nightscape Shoulders (8192, -307.28 DPS) [crafted] |
| back | Silky Spider Cape (10776) | Razorfen Downs: Tuten'kash [dungeon] | 14.8 stamina points (430.05 DPS) | yes | Wing of the Whelpling (13121, -15.23 DPS) [world_drop]; Well Oiled Cloak (12254, -32.70 DPS) [vendor]; Heraldic Cloak (8120, -36.82 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 31.8 stamina points (921.36 DPS) | yes | Warden's Wraps (14601, -65.60 DPS) [world_drop]; Robes of the Lich (10762, -111.86 DPS) [dungeon]; Insignia Chestguard (4057, -163.06 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.7 stamina points (427.27 DPS) | yes | Branded Leather Bracers (19508, -28.77 DPS) [dungeon]; Barbaric Bracers (18948, -49.74 DPS) [crafted]; Cultist's Armguards (270032, -50.94 DPS) [quest] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 20.5 stamina points (595.29 DPS) | yes | Warden's Leather Gloves (252527, -5.15 DPS) [crafted]; Razzeric's Racing Grips (6727, -32.80 DPS) [quest]; Bonefingers (10765, -34.43 DPS) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 20.4 stamina points (590.28 DPS) | yes | Kolkar Hunter's Belt (6788, -46.30 DPS) [quest]; Warden's Leather Belt (252460, -51.76 DPS) [crafted]; Scorpashi Sash (14652, -74.25 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 26.3 stamina points (761.50 DPS) | yes | Imperial Leather Pants (4062, -21.09 DPS) [dungeon]; Warden's Woolies (14605, -32.20 DPS) [world_drop]; Panther Hunter Leggings (4108, -34.17 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 24.5 stamina points (711.26 DPS) | yes | Gnomebot Operating Boots (9450, -91.15 DPS, sim-verified) [dungeon]; Skulker's Leather Shoes (252531, -94.82 DPS) [crafted]; Prowler's Leather Shoes (252465, -107.64 DPS) [crafted] |
| finger1 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.5 stamina points (333.61 DPS) | yes | Darkspear Signet (272070, -14.62 DPS) [vendor]; Underworld Band (1980, -43.62 DPS) [world_drop]; Dragonclaw Ring (10710, -43.62 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.4 stamina points (331.52 DPS) | yes | Darkspear Signet (272070, -12.53 DPS) [vendor]; Underworld Band (1980, -41.53 DPS) [world_drop]; Dragonclaw Ring (10710, -41.53 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | sim-verified (6069.2 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -29.00 DPS) [world] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (6069.2 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -29.00 DPS) [world] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (6069.2 DPS) | yes | Mograine's Might (7723, -148.70 DPS) [dungeon]; Illusionary Rod (7713, -330.63 DPS) [dungeon]; Manual Crowd Pummeler (9449, -412.76 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Silky Spider Cape; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Suspicious Spare Part; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 115.4. Weights run: 4.5s. Verify run: 1.6s. 577 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.035, armor=0.121 ± 0.010, defense=0.504 ± 0.123 per rating point (1 rating = 1%, 0.504 per %), dodge=0.239 ± 0.028 per rating point (12 rating = 1%, 2.864 per %), strength=0.059 ± 0.000, agility=0.249 ± 0.018, attack_power=0.029 ± 0.000, hit=not significant (0.071 ± 0.023) per rating point (10 rating = 1%, 0.708 per %), crit=0.043 ± 0.010 per rating point (14 rating = 1%, 0.605 per %), expertise=3.011 ± 0.259

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 38.9 stamina points (1546.69 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -183.16 DPS) [vendor]; Embrace of the Lycan (9479, -285.94 DPS) [dungeon]; Impractical Headwarmer (274756, -337.51 DPS) [vendor] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (556.08 DPS) | yes | Shriveled Heart (9243, -39.72 DPS) [dungeon]; Darkspear Warding Pendant (272073, -39.72 DPS) [vendor]; Souvenier Sea Shell (274749, -39.72 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.3 stamina points (1125.01 DPS) | yes | Penance Spaulders (11963, -24.68 DPS) [quest]; Phytoskin Spaulders (17749, -34.30 DPS) [dungeon]; Knight-Lieutenant's Leather Shoulders (220852, -57.37 DPS) [vendor] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | 19.7 stamina points (783.93 DPS) | yes | Sergeant's Cape (18441, -54.19 DPS) [vendor]; Nightfall Drape (12465, -114.97 DPS, sim-verified) [dungeon]; Grovekeeper's Drape (17739, -128.81 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 51.3 stamina points (2038.09 DPS) | yes | Mixologist's Tunic (12793, -455.05 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Skin (9473, -525.71 DPS) [dungeon]; Heraldic Breastplate (8119, -583.12 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.3 stamina points (924.08 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -59.27 DPS) [vendor]; Serpentskin Bracers (8257, -177.74 DPS) [world_drop]; Warden's Leather Bracers (252541, -182.27 DPS) [crafted] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 36.9 stamina points (1465.51 DPS) | yes | Warden's Leather Gauntlets (252549, -321.96 DPS) [crafted]; Feralheart Fists (226793, -400.79 DPS) [vendor]; Raider Gloves (272100, -419.99 DPS, sim-verified) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 25.5 stamina points (1013.41 DPS) | yes | Skulker's Leather Waistguard (252474, -20.59 DPS) [crafted]; Prowler's Leather Waistguard (252473, -35.68 DPS) [crafted]; Girdle of Beastial Fury (11686, -133.21 DPS) [dungeon] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 33.1 stamina points (1315.37 DPS) | yes | Knight's Crackling Leather Leggings (220864, -219.60 DPS) [vendor]; Scorpashi Leggings (14659, -239.37 DPS) [world_drop]; Windscale Sarong (10842, -562.62 DPS, sim-verified) [world] |
| feet | Shadefiend Boots (11675) | Blackrock Depths: Anub'shiah [dungeon] | 29.6 stamina points (1174.16 DPS) | yes | Slitherscale Boots (10801, -43.99 DPS) [dungeon]; Warden's Leather Boots (252470, -130.44 DPS) [crafted]; Skulker's Leather Boots (252469, -160.91 DPS) [crafted] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.4 stamina points (573.58 DPS) | yes | Ring of Saviors (1447, -17.50 DPS) [world_drop]; Darkspear Signet (272069, -17.50 DPS) [vendor]; Suspicious Spare Part (274754, -120.33 DPS) [vendor] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (556.08 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Suspicious Spare Part (274754, -102.83 DPS) [vendor] |
| trinket1 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (11237.9 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Smotts' Compass (4130, -160.52 DPS, sim-verified) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (11237.9 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -39.72 DPS) [quest] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (11237.9 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -28.87 DPS) [world_drop]; Ragehammer (10626, -693.37 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; back: Graverot Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Shadefiend Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Mark of the Chosen; trinket2: Darkspear Voodoo Seal; main_hand: Radiant Staff

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 137.6. Weights run: 4.5s. Verify run: 1.7s. 1449 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.140 ± 0.015, defense=1.303 ± 0.250 per rating point (1 rating = 1%, 1.303 per %), dodge=0.708 ± 0.045 per rating point (12 rating = 1%, 8.495 per %), strength=0.119 ± 0.000, agility=0.547 ± 0.029, attack_power=0.059 ± 0.000, hit=0.184 ± 0.036 per rating point (10 rating = 1%, 1.839 per %), crit=0.086 ± 0.016 per rating point (14 rating = 1%, 1.198 per %), expertise=6.074 ± 0.393

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Faceguard (226801) | Mokvar [vendor] | sim-verified (24571.2 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -274.02 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 21.6 stamina points (1703.59 DPS) | yes | Medallion of Grand Marshal Morris (13091, -121.73 DPS) [world_drop]; Sentinel's Medallion (19538, -266.53 DPS) [rep]; Evil Eye Pendant (18381, -335.71 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 64.0 stamina points (5053.97 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -111.60 DPS) [vendor]; Feralheart Pauldrons (226798, -675.09 DPS, sim-verified) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -761.00 DPS) [pvp] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | 32.7 stamina points (2580.58 DPS) | yes | Stoneskin Gargoyle Cape (13397, -32.90 DPS) [dungeon]; Redoubt Cloak (18495, -176.82 DPS) [dungeon]; Shifting Cloak (18511, -582.74 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, -25.59 DPS) [vendor]; Tunic of Undead Slaying (23089, -3182.78 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Subterfuge (22668, -151.97 DPS) [quest]; Forest Stalker's Bracers (19587, -230.31 DPS) [rep]; Wristwraps of Undead Slaying (23093, -773.23 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 52.9 stamina points (4178.53 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -335.01 DPS) [vendor]; Raider Gloves (272099, -352.67 DPS, sim-verified) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 47.7 stamina points (3764.43 DPS) | yes | Shifter's Belt (272396, -551.46 DPS, sim-verified) [vendor]; Hivethrasher's Girdle (275614, -772.92 DPS) [crafted]; Belt of Preserved Heads (20216, -774.19 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 84.4 stamina points (6662.85 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Dire Warbear Woolies (227804, -408.82 DPS, sim-verified) [vendor]; Marshal's Dragonhide Legguards (231703, -494.56 DPS) [pvp] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 63.6 stamina points (5021.20 DPS) | yes | Marshal's Dragonhide Treads (231692, -375.84 DPS) [vendor]; Feralheart Treads (226803, -567.56 DPS, sim-verified) [vendor]; Drudge Boots (21532, -856.38 DPS) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (21196) | The Path of the Protector [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Awareness (272409, -238.78 DPS) [vendor]; Myrmidon's Signet (2246, -291.38 DPS) [world_drop]; Naglering (11669, -376.69 DPS, sim-verified) [dungeon] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Awareness (272409, -144.91 DPS) [vendor]; Myrmidon's Signet (2246, -197.50 DPS) [world_drop]; Naglering (11669, -320.83 DPS, sim-verified) [dungeon] |
| trinket1 | Mark of Tyranny (13966) | General Drakkisath's Demise [quest] | sim-verified (+1969.2 DPS vs the runner-up, not corroborated against the finished set) | yes | Stormpike Insignia Rank 6 (17904, -1325.18 DPS) [quest]; Vigilance Charm (18370, -1325.18 DPS) [dungeon]; Talisman of Arathor (20071, -1719.23 DPS) [rep] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Vigilance Charm (18370, -324.99 DPS, sim-verified) [dungeon]; Talisman of Arathor (20071, -390.40 DPS) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Electrified Dagger (19100, -1587.94 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Faceguard; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Mark of Tyranny; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60, raid preset (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 286.8. Weights run: 4.6s. Verify run: 1.7s. 1449 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.164 ± 0.018, defense=1.818 ± 0.327 per rating point (1 rating = 1%, 1.818 per %), dodge=0.995 ± 0.059 per rating point (12 rating = 1%, 11.936 per %), strength=0.137 ± 0.000, agility=0.770 ± 0.037, attack_power=0.069 ± 0.000, hit=0.257 ± 0.046 per rating point (10 rating = 1%, 2.574 per %), crit=0.130 ± 0.021 per rating point (14 rating = 1%, 1.814 per %), expertise=8.076 ± 0.505

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Faceguard (226801) | Mokvar [vendor] | sim-verified (42548.2 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -468.72 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 26.0 stamina points (2506.11 DPS) | yes | Medallion of Grand Marshal Morris (13091, -79.53 DPS) [world_drop]; Evil Eye Pendant (18381, -166.56 DPS) [dungeon]; Talisman of Evasion (13177, -391.07 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 76.0 stamina points (7321.85 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -573.56 DPS) [vendor]; Darkspear Pauldrons (272105, -1503.80 DPS, sim-verified) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -1649.45 DPS) [pvp] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 40.9 stamina points (3941.09 DPS) | yes | Stoneshield Cloak (12551, -361.50 DPS) [dungeon]; Stoneskin Gargoyle Cape (13397, -436.44 DPS) [dungeon]; Redoubt Cloak (18495, -558.28 DPS) [dungeon] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, -959.45 DPS) [vendor]; Tunic of Undead Slaying (23089, -5101.13 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -107.32 DPS) [rep]; Bracers of Subterfuge (22668, -110.90 DPS) [quest]; Wristwraps of Undead Slaying (23093, -847.41 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 64.3 stamina points (6193.43 DPS) | yes | Marshal's Dragonhide Grips (231694, -85.81 DPS) [vendor]; Raider Gloves (272099, -655.54 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -1201.34 DPS) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 55.9 stamina points (5385.87 DPS) | yes | Shifter's Belt (272396, -530.02 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -1073.18 DPS) [quest]; Ferocity of the Timbermaw (227805, -1159.27 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 96.4 stamina points (9286.92 DPS) | yes | Marshal's Dragonhide Leggings (231691, -211.21 DPS) [vendor]; Dire Warbear Woolies (227804, -450.74 DPS, sim-verified) [vendor]; Feralheart Legguards (226799, -1315.42 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 76.2 stamina points (7338.89 DPS) | yes | Marshal's Dragonhide Treads (231692, -1039.15 DPS) [vendor]; Drudge Boots (21532, -1262.61 DPS, sim-verified) [quest]; Feralheart Treads (226803, -1623.11 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (21196) | The Path of the Protector [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, -113.54 DPS) [dungeon]; Band of Resolution (22680, -132.41 DPS) [quest]; Band of the Steadfast Hero (22331, -267.55 DPS) [dungeon] |
| finger2 | Ring of Awareness (272409) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Naglering (11669, -16.89 DPS) [dungeon]; Band of Resolution (22680, -35.75 DPS) [quest]; Band of the Steadfast Hero (22331, -170.90 DPS) [dungeon] |
| trinket1 | Mark of Tyranny (13966) | General Drakkisath's Demise [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Insignia Rank 6 (17904, -1693.10 DPS) [quest]; Defender's Grip Stabilizer (272440, -1716.12 DPS) [vendor]; Arena Grand Master (19024, -2819.29 DPS, sim-verified) [quest] |
| trinket2 | Vigilance Charm (18370) | Dire Maul: Immol'thar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Defender's Grip Stabilizer (272440, -23.03 DPS) [vendor]; Force of Will (11810, -1074.22 DPS) [dungeon] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Electrified Dagger (19100, -2024.52 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Faceguard; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Ring of Awareness; trinket1: Mark of Tyranny; trinket2: Vigilance Charm; main_hand: Headmaster's Charge

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 35.0. Weights run: 3.5s. Verify run: 1.3s. 183 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.083 ± 0.002, defense=0.169 ± 0.034 per rating point (1 rating = 1%, 0.169 per %), dodge=0.129 ± 0.010 per rating point (12 rating = 1%, 1.547 per %), strength=0.070 ± 0.000, agility=0.152 ± 0.006, attack_power=0.035 ± 0.000, hit=0.066 ± 0.006 per rating point (10 rating = 1%, 0.659 per %), crit=0.024 ± 0.002 per rating point (14 rating = 1%, 0.340 per %), expertise=1.471 ± 0.067

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (432.17 DPS) | yes | Brawler's Leather Hood (252504, -24.18 DPS) [crafted]; Defender's Leather Hood (252447, -39.55 DPS) [crafted]; Totemic Leather Hood (252448, -52.51 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.6 stamina points (153.77 DPS) | yes | Scholarly Pendant (277203, -14.17 DPS) [quest]; Scout's Medallion (20442, -85.98 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (216.85 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -67.62 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -71.56 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (178.28 DPS) | yes | Sentry Cloak (2059, -7.17 DPS) [world_drop]; Grave Shroud (279865, -11.32 DPS) [quest]; Miner's Cape (5444, -20.03 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.1 stamina points (421.90 DPS) | yes | Gloomshroud Armor (1489, -25.62 DPS, sim-verified) [dungeon]; Loch Croc Hide Vest (6197, -73.31 DPS) [world]; Trapper's Leather Armor (252491, -116.38 DPS) [crafted] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.8 stamina points (204.32 DPS) | yes | Bear Bracers (4795, -16.43 DPS) [vendor]; Savannah Bracers (15453, -18.36 DPS) [quest]; Drakewing Bands (12999, -55.22 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.0 stamina points (209.74 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.93 DPS) [crafted]; Defender's Leather Gloves (252496, -7.69 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.2 stamina points (260.81 DPS) | yes | Belt of the Fang (10412, -22.62 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -60.74 DPS) [crafted]; Guardsman Belt (3429, -62.67 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.1 stamina points (326.94 DPS) | yes | Slick Deviate Leggings (6480, -4.85 DPS) [quest]; Smelting Pants (5199, -17.08 DPS) [dungeon]; Brawler's Leather Pants (252500, -33.69 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (348.83 DPS) | yes | Surfer Shoes (276274, -43.62 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -70.00 DPS) [dungeon]; Trapper's Leather Boots (252440, -96.81 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.3 stamina points (124.20 DPS) | yes | Ring of Scorn (3235, -31.13 DPS) [quest]; Ring of the Moon (12052, -52.78 DPS) [world_drop]; Deep Fathom Ring (6463, -54.40 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (116.33 DPS) | yes | Ring of Scorn (3235, -16.04 DPS, sim-verified) [quest]; Ring of the Moon (12052, -44.91 DPS) [world_drop]; Deep Fathom Ring (6463, -46.53 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (2182.6 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 18.6 stamina points (432.47 DPS) | yes | Crescent Staff (6505, -13.29 DPS, sim-verified) [quest]; Hammerbone (270018, -21.60 DPS) [quest]; Advisor's Gnarled Staff (20425, -23.38 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 41.1. Weights run: 3.5s. Verify run: 1.3s. 315 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.317 ± 0.056 per rating point (1 rating = 1%, 0.317 per %), dodge=0.139 ± 0.014 per rating point (12 rating = 1%, 1.667 per %), strength=0.065 ± 0.000, agility=0.161 ± 0.008, attack_power=0.032 ± 0.000, hit=0.073 ± 0.008 per rating point (10 rating = 1%, 0.732 per %), crit=0.024 ± 0.001 per rating point (14 rating = 1%, 0.337 per %), expertise=1.835 ± 0.100

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 stamina points (473.17 DPS) | yes | Defender's Leather Helm (252455, -26.62 DPS, sim-verified) [crafted]; Totemic Leather Helm (252456, -35.10 DPS) [crafted]; Trapper's Leather Helm (252513, -35.10 DPS) [crafted] |
| neck | Souvenier Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (284.23 DPS) | yes | Ghostshard Talisman (7731, -62.38 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -81.79 DPS) [world_drop]; Senior Sergeant's Insignia (15200, -87.46 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 17.9 stamina points (391.27 DPS) | yes | Tanned Shoulderpads (270023, -54.00 DPS, sim-verified) [quest]; Cloudy Gustwoven Spaulders (277043, -74.47 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -74.47 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 7 · Sergeant · Horde [vendor] | 11.2 stamina points (244.39 DPS) | yes | Tigerstrike Mantle (13108, -17.48 DPS) [world_drop]; Enduring Cape (14763, -41.20 DPS) [world_drop]; Grimsteel Cape (4643, -64.07 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.2 stamina points (463.71 DPS) | yes | Spirewind Fetter (9406, -8.16 DPS) [dungeon]; Green Whelp Armor (7375, -36.40 DPS) [crafted]; Dokebi Chestguard (14581, -65.59 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (240.72 DPS) | yes | Barbaric Bracers (18948, -3.75 DPS) [crafted]; Black Wolf Bracers (3230, -6.08 DPS) [dungeon]; Technician's Bracers (270042, -38.23 DPS) [quest] |
| hands | Ebon Vise (7690) | Scarlet Monastery: Fallen Champion [dungeon] | 15.1 stamina points (329.84 DPS) | yes | Braced Handguards (6784, -24.13 DPS, sim-verified) [quest]; Brawler Gloves (720, -51.90 DPS) [world_drop]; Wolfclaw Gloves (1978, -58.79 DPS) [dungeon] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.6 stamina points (341.72 DPS) | yes | Skulker's Leather Belt (252520, -44.73 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -61.34 DPS) [crafted]; Stalker's Leather Belt (252521, -63.57 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.1 stamina points (396.10 DPS) | yes | Brawler's Leather Legguards (252516, -36.71 DPS) [crafted]; Dark Ritual Leggings (270031, -38.15 DPS) [quest]; Trapper's Leather Legguards (252517, -46.63 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.4 stamina points (401.56 DPS) | yes | Harbinger Boots (7754, -28.19 DPS) [dungeon]; Grizzled Boots (6335, -54.13 DPS) [quest]; Highlander's Mail Greaves (20123, -63.04 DPS) [vendor] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.3 stamina points (181.29 DPS) | yes | Plains Ring (2039, -6.38 DPS) [dungeon]; Darkspear Signet (272071, -6.38 DPS) [vendor]; Defias Renegade Ring (1076, -54.42 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.2 stamina points (179.16 DPS) | yes | Plains Ring (2039, -4.25 DPS) [dungeon]; Darkspear Signet (272071, -4.25 DPS) [vendor]; Defias Renegade Ring (1076, -52.30 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (3410.4 DPS) | yes | Defiler's Talisman (21120, +0.00 DPS) [rep]; Rune of Perfection (21566, -43.73 DPS) [rep]; Rune of Duty (21568, -43.73 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272062) | Creeg Bothunk [vendor] | sim-verified (3410.4 DPS) | yes | Defiler's Talisman (21120, +0.00 DPS) [rep]; Rune of Perfection (21566, -43.73 DPS) [rep]; Rune of Duty (21568, -43.73 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul: Overlord Ramtusk [dungeon] | sim-verified (3410.4 DPS) | yes | Advisor's Gnarled Staff (19569, -33.38 DPS) [pvp]; Glimmering Staff (249392, -65.23 DPS) [crafted]; Manual Crowd Pummeler (9449, -234.09 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenier Sea Shell; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Ebon Vise; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Staff of the Shade

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 71.4. Weights run: 4.0s. Verify run: 1.4s. 426 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.025, armor=0.128 ± 0.007, defense=0.466 ± 0.082 per rating point (1 rating = 1%, 0.466 per %), dodge=0.176 ± 0.020 per rating point (12 rating = 1%, 2.117 per %), strength=0.072 ± 0.000, agility=0.183 ± 0.012, attack_power=0.036 ± 0.000, hit=0.072 ± 0.015 per rating point (10 rating = 1%, 0.722 per %), crit=0.031 ± 0.006 per rating point (14 rating = 1%, 0.430 per %), expertise=2.407 ± 0.170

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman: Shadowforge Relic Hunter [dungeon] | 27.6 stamina points (801.63 DPS) | yes | Warden's Wizard Hat (14604, -47.31 DPS) [world_drop]; Brawler's Leather Helm (252512, -60.73 DPS) [crafted]; Nightscape Headband (8176, -82.22 DPS) [crafted] |
| neck | Shriveled Heart (9243) (or Souvenier Sea Shell (274749)) | Zul'Farrak: Sandarr Dunereaver [dungeon] | 13.0 stamina points (376.98 DPS) | yes | Souvenier Sea Shell (274749, +0.00 DPS) [vendor]; Dragon's Blood Necklace (10711, -29.00 DPS) [quest]; Gazlowe's Charm (13088, -49.64 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.6 stamina points (828.89 DPS) | yes | Flintrock Shoulders (7755, -128.16 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -207.47 DPS) [dungeon]; Tanned Shoulderpads (270023, -297.65 DPS) [quest] |
| back | Silky Spider Cape (10776) | Razorfen Downs: Tuten'kash [dungeon] | 14.8 stamina points (430.05 DPS) | yes | Wing of the Whelpling (13121, -15.23 DPS) [world_drop]; Well Oiled Cloak (12254, -32.70 DPS) [vendor]; Heraldic Cloak (8120, -36.82 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 31.8 stamina points (921.36 DPS) | yes | Warden's Wraps (14601, -65.60 DPS) [world_drop]; Robes of the Lich (10762, -111.86 DPS) [dungeon]; Insignia Chestguard (4057, -163.06 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.7 stamina points (427.27 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -28.77 DPS) [dungeon]; Barbaric Bracers (18948, -49.74 DPS) [crafted] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 20.5 stamina points (595.29 DPS) | yes | Warden's Leather Gloves (252527, -5.15 DPS) [crafted]; Razzeric's Racing Grips (6727, -32.80 DPS) [quest]; Bonefingers (10765, -34.43 DPS) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 20.4 stamina points (590.28 DPS) | yes | Warden's Leather Belt (252460, -51.76 DPS) [crafted]; Kolkar Hunter's Belt (6788, -52.37 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -74.25 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 26.3 stamina points (761.50 DPS) | yes | Warden's Woolies (14605, -32.20 DPS) [world_drop]; Panther Hunter Leggings (4108, -34.17 DPS) [quest]; Imperial Leather Pants (4062, -50.13 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 24.5 stamina points (711.26 DPS) | yes | Gnomebot Operating Boots (9450, -59.66 DPS, sim-verified) [dungeon]; Skulker's Leather Shoes (252531, -94.82 DPS) [crafted]; Prowler's Leather Shoes (252465, -107.64 DPS) [crafted] |
| finger1 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.5 stamina points (333.61 DPS) | yes | Darkspear Signet (272070, -14.62 DPS) [vendor]; Underworld Band (1980, -43.62 DPS) [world_drop]; Dragonclaw Ring (10710, -43.62 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.4 stamina points (331.52 DPS) | yes | Darkspear Signet (272070, -12.53 DPS) [vendor]; Underworld Band (1980, -41.53 DPS) [world_drop]; Dragonclaw Ring (10710, -41.53 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | sim-verified (6559.1 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -29.00 DPS) [world] |
| trinket2 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (6559.1 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -29.00 DPS) [world] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (6559.1 DPS) | yes | Ironshod Bludgeon (9408, -12.31 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -158.35 DPS) [quest]; Manual Crowd Pummeler (9449, -568.74 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Silky Spider Cape; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Suspicious Spare Part; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Cragwood Maul

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 115.1. Weights run: 4.5s. Verify run: 1.6s. 561 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.035, armor=0.121 ± 0.010, defense=0.504 ± 0.123 per rating point (1 rating = 1%, 0.504 per %), dodge=0.239 ± 0.028 per rating point (12 rating = 1%, 2.864 per %), strength=0.059 ± 0.000, agility=0.249 ± 0.018, attack_power=0.029 ± 0.000, hit=not significant (0.071 ± 0.023) per rating point (10 rating = 1%, 0.708 per %), crit=0.043 ± 0.010 per rating point (14 rating = 1%, 0.605 per %), expertise=3.011 ± 0.259

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 38.9 stamina points (1546.69 DPS) | yes | Blood Guard's Leather Headband (220851, -183.16 DPS) [vendor]; Embrace of the Lycan (9479, -279.89 DPS, sim-verified) [dungeon]; Impractical Headwarmer (274756, -337.51 DPS) [vendor] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (556.08 DPS) | yes | Shriveled Heart (9243, -39.72 DPS) [dungeon]; Darkspear Warding Pendant (272073, -39.72 DPS) [vendor]; Souvenier Sea Shell (274749, -39.72 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.3 stamina points (1125.01 DPS) | yes | Penance Spaulders (11963, -24.68 DPS) [quest]; Phytoskin Spaulders (17749, -34.30 DPS) [dungeon]; Blood Guard's Leather Shoulders (220853, -57.37 DPS) [vendor] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | 19.7 stamina points (783.93 DPS) | yes | Nightfall Drape (12465, -44.54 DPS) [dungeon]; Sergeant's Cloak (16341, -54.19 DPS) [vendor]; Grovekeeper's Drape (17739, -128.81 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 51.3 stamina points (2038.09 DPS) | yes | Mixologist's Tunic (12793, -390.59 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Skin (9473, -525.71 DPS) [dungeon]; Heraldic Breastplate (8119, -583.12 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.3 stamina points (924.08 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -59.27 DPS) [vendor]; Forest Stalker's Bracers (19589, -95.32 DPS) [pvp]; Serpentskin Bracers (8257, -177.74 DPS) [world_drop] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 36.9 stamina points (1465.51 DPS) | yes | Raider Gloves (272100, -310.49 DPS, sim-verified) [vendor]; Warden's Leather Gauntlets (252549, -321.96 DPS) [crafted]; Feralheart Fists (226793, -400.79 DPS) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 25.5 stamina points (1013.41 DPS) | yes | Skulker's Leather Waistguard (252474, -20.59 DPS) [crafted]; Prowler's Leather Waistguard (252473, -35.68 DPS) [crafted]; Girdle of Beastial Fury (11686, -133.21 DPS) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 33.1 stamina points (1315.37 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -219.60 DPS) [vendor]; Scorpashi Leggings (14659, -239.37 DPS) [world_drop]; Windscale Sarong (10842, -431.87 DPS, sim-verified) [world] |
| feet | Shadefiend Boots (11675) | Blackrock Depths: Anub'shiah [dungeon] | 29.6 stamina points (1174.16 DPS) | yes | Slitherscale Boots (10801, -43.99 DPS) [dungeon]; Warden's Leather Boots (252470, -130.44 DPS) [crafted]; Skulker's Leather Boots (252469, -160.91 DPS) [crafted] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.4 stamina points (573.58 DPS) | yes | Ring of Saviors (1447, -17.50 DPS) [world_drop]; Darkspear Signet (272069, -17.50 DPS) [vendor]; Suspicious Spare Part (274754, -120.33 DPS) [vendor] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (556.08 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Suspicious Spare Part (274754, -102.83 DPS) [vendor] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (12179.0 DPS) | yes | Defiler's Talisman (21115, +0.00 DPS) [rep]; Guardian Talisman (1490, -39.72 DPS) [quest] |
| trinket2 | Relentless Raider's Seal (272060) | Creeg Bothunk [vendor] | sim-verified (12179.0 DPS) | yes | Defiler's Talisman (21115, +0.00 DPS) [rep]; Guardian Talisman (1490, -39.72 DPS) [quest] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (12179.0 DPS) | yes | Advisor's Gnarled Staff (19567, +0.00 DPS) [pvp]; Radiant Staff (249453, -1.30 DPS) [crafted]; Shadowblade (2163, -827.24 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Graverot Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Shadefiend Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 139.9. Weights run: 4.5s. Verify run: 1.7s. 1446 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.140 ± 0.015, defense=1.303 ± 0.250 per rating point (1 rating = 1%, 1.303 per %), dodge=0.708 ± 0.045 per rating point (12 rating = 1%, 8.495 per %), strength=0.119 ± 0.000, agility=0.547 ± 0.029, attack_power=0.059 ± 0.000, hit=0.184 ± 0.036 per rating point (10 rating = 1%, 1.839 per %), crit=0.086 ± 0.016 per rating point (14 rating = 1%, 1.198 per %), expertise=6.074 ± 0.393

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Faceguard (226801) | Mokvar [vendor] | sim-verified (26086.6 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -268.94 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 21.6 stamina points (1703.59 DPS) | yes | Medallion of Grand Marshal Morris (13091, -121.73 DPS) [world_drop]; Scout's Medallion (19534, -266.53 DPS) [rep]; Evil Eye Pendant (18381, -335.71 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 64.0 stamina points (5053.97 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -111.60 DPS) [vendor]; Feralheart Pauldrons (226798, -657.19 DPS, sim-verified) [vendor]; Warlord's Dragonhide Spaulders (231681, -761.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | 32.7 stamina points (2580.58 DPS) | yes | Stoneskin Gargoyle Cape (13397, -32.90 DPS) [dungeon]; Redoubt Cloak (18495, -176.82 DPS) [dungeon]; Shifting Cloak (18511, -755.00 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -25.59 DPS) [vendor]; Tunic of Undead Slaying (23089, -3149.23 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Subterfuge (22668, -151.97 DPS) [quest]; Forest Stalker's Bracers (19587, -230.31 DPS) [rep]; Wristwraps of Undead Slaying (23093, -661.32 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 52.9 stamina points (4178.53 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, -335.01 DPS) [pvp] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 47.7 stamina points (3764.43 DPS) | yes | Shifter's Belt (272396, -595.06 DPS, sim-verified) [vendor]; Hivethrasher's Girdle (275614, -772.92 DPS) [crafted]; Belt of Preserved Heads (20216, -774.19 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 84.4 stamina points (6662.85 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Dire Warbear Woolies (227804, -417.87 DPS, sim-verified) [vendor]; General's Dragonhide Legguards (231673, -494.56 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 63.6 stamina points (5021.20 DPS) | yes | General's Dragonhide Treads (231683, -375.84 DPS) [vendor]; Feralheart Treads (226803, -639.90 DPS, sim-verified) [vendor]; Drudge Boots (21532, -856.38 DPS) [quest] |
| finger1 | Thrall's Resolve (12544) | The Princess Saved? [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Resolution (22680, -317.73 DPS) [quest]; Ring of Awareness (272409, -462.64 DPS) [vendor]; Naglering (11669, -590.81 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21196) | The Path of the Protector [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Band of Resolution (22680, -93.87 DPS) [quest]; Naglering (11669, -171.10 DPS) [dungeon]; Ring of Awareness (272409, -238.78 DPS) [vendor] |
| trinket1 | Mark of Tyranny (13966) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Insignia Rank 6 (17909, -1325.18 DPS) [quest]; Defender's Grip Stabilizer (272440, -1389.45 DPS, sim-verified) [vendor]; Defiler's Talisman (20072, -1719.23 DPS) [rep] |
| trinket2 | Vigilance Charm (18370) | Dire Maul: Immol'thar [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Insignia Rank 6 (17909, +0.00 DPS) [quest]; Defender's Grip Stabilizer (272440, -3.65 DPS) [vendor]; Defiler's Talisman (20072, -394.05 DPS) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glacial Blade (19099, -1686.78 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Faceguard; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Thrall's Resolve; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Mark of Tyranny; trinket2: Vigilance Charm; main_hand: Headmaster's Charge

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60, raid preset (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 285.9. Weights run: 4.6s. Verify run: 1.7s. 1446 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.164 ± 0.018, defense=1.818 ± 0.327 per rating point (1 rating = 1%, 1.818 per %), dodge=0.995 ± 0.059 per rating point (12 rating = 1%, 11.936 per %), strength=0.137 ± 0.000, agility=0.770 ± 0.037, attack_power=0.069 ± 0.000, hit=0.257 ± 0.046 per rating point (10 rating = 1%, 2.574 per %), crit=0.130 ± 0.021 per rating point (14 rating = 1%, 1.814 per %), expertise=8.076 ± 0.505

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Faceguard (226801) | Mokvar [vendor] | sim-verified (+986.7 DPS vs the runner-up, not corroborated against the finished set) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -986.68 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 26.0 stamina points (2506.11 DPS) | yes | Medallion of Grand Marshal Morris (13091, -79.53 DPS) [world_drop]; Evil Eye Pendant (18381, -166.56 DPS) [dungeon]; Talisman of Evasion (13177, -391.07 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 76.0 stamina points (7321.85 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -573.56 DPS) [vendor]; Darkspear Pauldrons (272105, -1475.03 DPS, sim-verified) [vendor]; Warlord's Dragonhide Spaulders (231681, -1649.45 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (+955.4 DPS vs the runner-up, not corroborated against the finished set) | yes | Stoneskin Gargoyle Cape (13397, -74.94 DPS) [dungeon]; Redoubt Cloak (18495, -196.79 DPS) [dungeon]; Shifting Cloak (18511, -955.40 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -959.45 DPS) [vendor]; Tunic of Undead Slaying (23089, -4875.57 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Forest Stalker's Bracers (19587, -107.32 DPS) [rep]; Bracers of Subterfuge (22668, -110.90 DPS) [quest]; Wristwraps of Undead Slaying (23093, -1599.78 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 64.3 stamina points (6193.43 DPS) | yes | General's Dragonhide Grips (231688, -85.81 DPS) [vendor]; Raider Gloves (272099, -582.35 DPS, sim-verified) [vendor]; General's Dragonhide Gloves (231677, -1201.34 DPS) [pvp] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 55.9 stamina points (5385.87 DPS) | yes | Shifter's Belt (272396, -830.87 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -1073.18 DPS) [quest]; Ferocity of the Timbermaw (227805, -1159.27 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 96.4 stamina points (9286.92 DPS) | yes | General's Dragonhide Leggings (231685, -211.21 DPS) [pvp]; Dire Warbear Woolies (227804, -281.31 DPS) [vendor]; Feralheart Legguards (226799, -1315.42 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 76.2 stamina points (7338.89 DPS) | yes | General's Dragonhide Treads (231683, -1039.15 DPS) [vendor]; Drudge Boots (21532, -1072.05 DPS, sim-verified) [quest]; Feralheart Treads (226803, -1623.11 DPS) [vendor] |
| finger1 | Thrall's Resolve (12544) | The Princess Saved? [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Awareness (272409, -450.64 DPS) [vendor]; Band of Resolution (22680, -486.39 DPS) [quest]; Naglering (11669, -808.05 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21196) | The Path of the Protector [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Ring of Awareness (272409, -96.65 DPS) [vendor]; Naglering (11669, -113.54 DPS) [dungeon]; Band of Resolution (22680, -132.41 DPS) [quest] |
| trinket1 | Mark of Tyranny (13966) | For The Horde! [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Insignia Rank 6 (17909, -1693.10 DPS) [quest]; Vigilance Charm (18370, -1926.38 DPS, sim-verified) [dungeon]; Force of Will (11810, -2767.32 DPS) [dungeon] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Frostwolf Insignia Rank 6 (17909, +0.00 DPS) [quest]; Vigilance Charm (18370, +0.00 DPS) [dungeon]; Defiler's Talisman (20072, -527.88 DPS, sim-verified) [rep] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; The Lobotomizer (19324, -2173.33 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Faceguard; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Thrall's Resolve; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Mark of Tyranny; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

