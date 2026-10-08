# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 35.3. Weights run: 2.2s. Verify run: 1.2s. 193 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.083 ± 0.002, defense=0.183 ± 0.035 per rating point (1 rating = 1%, 0.183 per %), dodge=0.129 ± 0.010 per rating point (12 rating = 1%, 1.544 per %), strength=0.070 ± 0.000, agility=0.172 ± 0.006, attack_power=0.035 ± 0.000, hit=0.057 ± 0.006 per rating point (10 rating = 1%, 0.572 per %), crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.321 per %), expertise=1.490 ± 0.069

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (434.84 DPS) | yes | Brawler's Leather Hood (252504, -20.51 DPS) [crafted]; Defender's Leather Hood (252447, -39.67 DPS) [crafted]; Totemic Leather Hood (252448, -52.71 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.7 stamina points (156.52 DPS) | yes | Scholarly Pendant (277203, -16.10 DPS) [quest]; Tarnished Locket (279870, -39.51 DPS) [quest]; Sentinel's Medallion (20444, -85.56 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (218.32 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -65.69 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -72.10 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (179.39 DPS) | yes | Sentry Cloak (2059, -5.35 DPS) [world_drop]; Grave Shroud (279865, -10.46 DPS) [quest]; Miner's Cape (5444, -20.14 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (426.05 DPS) | yes | Gloomshroud Armor (1489, -23.13 DPS, sim-verified) [dungeon]; Loch Croc Hide Vest (6197, -75.15 DPS) [world]; Tunic of Westfall (2041, -85.48 DPS) [quest] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.1 stamina points (189.11 DPS) | yes | Drakewing Bands (12999, -39.01 DPS) [world_drop]; Wolf Bracers (4794, -54.11 DPS) [vendor]; Sanguine Cuffs (14375, -62.37 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (212.99 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.95 DPS) [crafted]; Defender's Leather Gloves (252496, -9.58 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.3 stamina points (264.82 DPS) | yes | Belt of the Fang (10412, -30.03 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -61.57 DPS) [crafted]; Guardsman Belt (3429, -63.52 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.8 stamina points (392.62 DPS) | yes | Defender's Leather Pants (252445, -45.95 DPS, sim-verified) [crafted]; Slick Deviate Leggings (6480, -66.55 DPS) [quest]; Smelting Pants (5199, -80.71 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (350.99 DPS) | yes | Surfer Shoes (276274, -42.82 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -67.55 DPS) [dungeon]; Trapper's Leather Boots (252440, -94.98 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (140.42 DPS) | yes | Blood Ring (4998, -23.40 DPS) [world_drop]; Ring of the Moon (12052, -68.58 DPS) [world_drop]; Deep Fathom Ring (6463, -70.21 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.4 stamina points (125.60 DPS) | yes | Blood Ring (4998, -8.59 DPS) [world_drop]; Ring of the Moon (12052, -53.76 DPS) [world_drop]; Deep Fathom Ring (6463, -55.39 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (2119.2 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 18.6 stamina points (435.06 DPS) | yes | Rakzur Club (12983, -32.86 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -37.44 DPS) [dungeon]; Twisted Chanter's Staff (890, -61.42 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 193, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 41.6. Weights run: 2.1s. Verify run: 1.4s. 322 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.318 ± 0.057 per rating point (1 rating = 1%, 0.318 per %), dodge=0.139 ± 0.014 per rating point (12 rating = 1%, 1.662 per %), strength=0.065 ± 0.000, agility=0.173 ± 0.009, attack_power=0.032 ± 0.000, hit=0.072 ± 0.008 per rating point (10 rating = 1%, 0.724 per %), crit=0.024 ± 0.001 per rating point (14 rating = 1%, 0.331 per %), expertise=1.823 ± 0.100

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.7 stamina points (477.85 DPS) | yes | Totemic Leather Helm (252456, -38.03 DPS) [crafted]; Trapper's Leather Helm (252513, -38.03 DPS) [crafted]; Defender's Leather Helm (252455, -50.56 DPS, sim-verified) [crafted] |
| neck | Souvenier Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (285.66 DPS) | yes | Ghostshard Talisman (7731, -61.28 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -82.20 DPS) [world_drop]; Master Sergeant's Insignia (18442, -87.89 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 17.9 stamina points (392.80 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -66.78 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -72.89 DPS) [crafted]; Barbaric Shoulders (5964, -95.35 DPS) [crafted] |
| back | Sergeant's Cape (18440) | PvP rank 7 · Sergeant · Alliance [vendor] | 11.2 stamina points (245.48 DPS) | yes | Tigerstrike Mantle (13108, -24.95 DPS, sim-verified) [world_drop]; Enduring Cape (14763, -40.83 DPS) [world_drop]; Watch Master's Cloak (2953, -49.45 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.2 stamina points (465.51 DPS) | yes | Spirewind Fetter (9406, -24.53 DPS, sim-verified) [dungeon]; Green Whelp Armor (7375, -36.61 DPS) [crafted]; Blackened Defias Armor (10399, -65.33 DPS) [dungeon] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (241.68 DPS) | yes | Barbaric Bracers (18948, -2.68 DPS) [crafted]; Black Wolf Bracers (3230, -6.11 DPS) [dungeon]; Staghide Armguards (270021, -35.72 DPS) [quest] |
| hands | Ebon Vise (7690) | Scarlet Monastery: Fallen Champion [dungeon] | 15.1 stamina points (332.77 DPS) | yes | Brawler Gloves (720, -53.78 DPS) [world_drop]; Wolfclaw Gloves (1978, -59.04 DPS) [dungeon]; Operator's Gloves (270045, -65.91 DPS, sim-verified) [quest] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.7 stamina points (344.74 DPS) | yes | Skulker's Leather Belt (252520, -61.30 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -61.65 DPS) [crafted]; Stalker's Leather Belt (252521, -63.06 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.2 stamina points (399.52 DPS) | yes | Brawler's Leather Legguards (252516, -36.34 DPS) [crafted]; Dark Ritual Leggings (270031, -40.26 DPS) [quest]; Duty Bound Leggings (279868, -62.25 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.3 stamina points (403.16 DPS) | yes | Harbinger Boots (7754, -34.11 DPS, sim-verified) [dungeon]; Highlander's Mail Greaves (20123, -61.14 DPS) [vendor]; Lancer Boots (6752, -66.78 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.3 stamina points (182.20 DPS) | yes | Plains Ring (2039, -6.41 DPS) [dungeon]; Sustaining Ring (6743, -50.35 DPS) [quest]; Defias Renegade Ring (1076, -53.31 DPS) [dungeon] |
| finger2 | Darkspear Signet (272071) (or Plains Ring (2039)) | Creeg Bothunk [vendor] | 8.0 stamina points (175.79 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Sustaining Ring (6743, -43.95 DPS) [quest]; Defias Renegade Ring (1076, -46.91 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (3266.8 DPS) | yes | Talisman of Arathor (21119, +0.00 DPS) [rep]; Rune of Perfection (21566, -43.95 DPS) [rep]; Rune of Duty (21568, -43.95 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272062) | Creeg Bothunk [vendor] | sim-verified (3266.8 DPS) | yes | Talisman of Arathor (21119, +0.00 DPS) [rep]; Rune of Perfection (21566, -43.95 DPS) [rep]; Rune of Duty (21568, -43.95 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul: Overlord Ramtusk [dungeon] | sim-verified (3266.8 DPS) | yes | Glimmering Staff (249392, -65.57 DPS) [crafted]; Soulstaff (249393, -66.66 DPS) [crafted]; Manual Crowd Pummeler (9449, -233.70 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenier Sea Shell; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Ebon Vise; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Darkspear Signet; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Staff of the Shade

No-known-source sample (15 of 322, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 40 (night-elf, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 73.5. Weights run: 2.6s. Verify run: 1.5s. 438 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.027, armor=0.125 ± 0.007, defense=0.391 ± 0.084 per rating point (1 rating = 1%, 0.391 per %), dodge=0.177 ± 0.019 per rating point (12 rating = 1%, 2.129 per %), strength=0.070 ± 0.000, agility=0.191 ± 0.012, attack_power=0.035 ± 0.000, hit=0.066 ± 0.015 per rating point (10 rating = 1%, 0.665 per %), crit=0.031 ± 0.006 per rating point (14 rating = 1%, 0.434 per %), expertise=2.366 ± 0.167

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman: Shadowforge Relic Hunter [dungeon] | 27.5 stamina points (818.13 DPS) | yes | Warden's Wizard Hat (14604, -47.95 DPS) [world_drop]; Brawler's Leather Helm (252512, -61.35 DPS) [crafted]; Nightscape Headband (8176, -83.44 DPS) [crafted] |
| neck | Shriveled Heart (9243) (or Souvenier Sea Shell (274749)) | Zul'Farrak: Sandarr Dunereaver [dungeon] | 13.0 stamina points (386.96 DPS) | yes | Souvenier Sea Shell (274749, +0.00 DPS) [vendor]; Gazlowe's Charm (13088, -51.16 DPS) [dungeon]; Darkspear Warding Pendant (272074, -59.53 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.4 stamina points (845.05 DPS) | yes | Flintrock Shoulders (7755, -160.56 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -213.23 DPS) [dungeon]; Nightscape Shoulders (8192, -313.27 DPS) [crafted] |
| back | Silky Spider Cape (10776) | Razorfen Downs: Tuten'kash [dungeon] | 14.8 stamina points (439.22 DPS) | yes | Wing of the Whelpling (13121, -14.50 DPS) [world_drop]; Well Oiled Cloak (12254, -33.49 DPS) [vendor]; Heraldic Cloak (8120, -36.81 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 31.5 stamina points (937.66 DPS) | yes | Warden's Wraps (14601, -65.14 DPS) [world_drop]; Robes of the Lich (10762, -111.30 DPS) [dungeon]; Insignia Chestguard (4057, -166.88 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.7 stamina points (436.21 DPS) | yes | Branded Leather Bracers (19508, -31.73 DPS) [dungeon]; Barbaric Bracers (18948, -51.38 DPS) [crafted]; Cultist's Armguards (270032, -53.42 DPS) [quest] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 20.4 stamina points (607.90 DPS) | yes | Razzeric's Racing Grips (6727, -33.72 DPS) [quest]; Bonefingers (10765, -38.17 DPS) [dungeon]; Warden's Leather Gloves (252527, -56.87 DPS, sim-verified) [crafted] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 20.2 stamina points (602.43 DPS) | yes | Warden's Leather Belt (252460, -53.13 DPS) [crafted]; Kolkar Hunter's Belt (6788, -74.22 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -75.50 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 26.2 stamina points (778.49 DPS) | yes | Panther Hunter Leggings (4108, -36.53 DPS) [quest]; Warden's Woolies (14605, -37.48 DPS) [world_drop]; Imperial Leather Pants (4062, -129.38 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 24.3 stamina points (724.67 DPS) | yes | Skulker's Leather Shoes (252531, -96.35 DPS) [crafted]; Prowler's Leather Shoes (252465, -110.69 DPS) [crafted]; Gnomebot Operating Boots (9450, -131.94 DPS, sim-verified) [dungeon] |
| finger1 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.5 stamina points (342.09 DPS) | yes | Darkspear Signet (272070, -14.66 DPS) [vendor]; Underworld Band (1980, -44.42 DPS) [world_drop]; Dragonclaw Ring (10710, -44.42 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.4 stamina points (339.99 DPS) | yes | Darkspear Signet (272070, -12.56 DPS) [vendor]; Underworld Band (1980, -42.33 DPS) [world_drop]; Dragonclaw Ring (10710, -42.33 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | sim-verified (6159.1 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -29.77 DPS) [world] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (6159.1 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -29.77 DPS) [world] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (6159.1 DPS) | yes | Mograine's Might (7723, -153.25 DPS) [dungeon]; Illusionary Rod (7713, -337.97 DPS) [dungeon]; Manual Crowd Pummeler (9449, -468.75 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Silky Spider Cape; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Suspicious Spare Part; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 438, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 50 (night-elf, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 120.4. Weights run: 2.9s. Verify run: 3.2s. 577 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.035, armor=0.127 ± 0.011, defense=not significant (0.491 ± 0.126) per rating point (1 rating = 1%, 0.491 per %), dodge=0.246 ± 0.029 per rating point (12 rating = 1%, 2.951 per %), strength=0.061 ± 0.000, agility=0.252 ± 0.018, attack_power=0.030 ± 0.000, hit=not significant (0.069 ± 0.023) per rating point (10 rating = 1%, 0.694 per %), crit=not significant (0.035 ± 0.010) per rating point (14 rating = 1%, 0.493 per %), expertise=3.041 ± 0.263

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) | Captain Dirgehammer [vendor] | sim-verified (12134.7 DPS) | yes | Embrace of the Lycan (9479, -96.16 DPS) [dungeon]; Impractical Headwarmer (274756, -146.22 DPS) [vendor]; Sprightring Helm (17776, -860.44 DPS, sim-verified) [quest] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (540.40 DPS) | yes | Shriveled Heart (9243, -38.60 DPS) [dungeon]; Darkspear Warding Pendant (272073, -38.60 DPS) [vendor]; Souvenier Sea Shell (274749, -38.60 DPS) [vendor] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | sim-verified (12134.7 DPS) | yes | Penance Spaulders (11963, +0.00 DPS) [quest]; Phytoskin Spaulders (17749, +0.00 DPS) [dungeon]; Fleshhide Shoulders (10774, -866.76 DPS, sim-verified) [dungeon] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | 20.0 stamina points (770.57 DPS) | yes | Nightfall Drape (12465, -43.51 DPS) [dungeon]; Sergeant's Cape (18441, -53.34 DPS) [vendor]; Grovekeeper's Drape (17739, -125.62 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 52.3 stamina points (2019.43 DPS) | yes | Mixologist's Tunic (12793, -319.17 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Skin (9473, -516.74 DPS) [dungeon]; Heraldic Breastplate (8119, -574.96 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.7 stamina points (913.25 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -58.48 DPS) [vendor]; Serpentskin Bracers (8257, -174.51 DPS) [world_drop]; Warden's Leather Bracers (252541, -176.65 DPS) [crafted] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 37.6 stamina points (1449.92 DPS) | yes | Raider Gloves (272100, -276.37 DPS, sim-verified) [vendor]; Warden's Leather Gauntlets (252549, -314.49 DPS) [crafted]; Feralheart Fists (226793, -388.43 DPS) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 26.1 stamina points (1005.60 DPS) | yes | Skulker's Leather Waistguard (252474, -17.95 DPS) [crafted]; Prowler's Leather Waistguard (252473, -32.77 DPS) [crafted]; Girdle of Beastial Fury (11686, -128.71 DPS) [dungeon] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 33.8 stamina points (1304.04 DPS) | yes | Knight's Crackling Leather Leggings (220864, -214.02 DPS) [vendor]; Scorpashi Leggings (14659, -231.52 DPS) [world_drop]; Windscale Sarong (10842, -1100.72 DPS, sim-verified) [world] |
| feet | Shadefiend Boots (11675) | Blackrock Depths: Anub'shiah [dungeon] | 30.3 stamina points (1167.78 DPS) | yes | Slitherscale Boots (10801, -49.44 DPS) [dungeon]; Warden's Leather Boots (252470, -129.43 DPS) [crafted]; Skulker's Leather Boots (252469, -157.13 DPS) [crafted] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.5 stamina points (557.92 DPS) | yes | Ring of Saviors (1447, -17.52 DPS) [world_drop]; Darkspear Signet (272069, -17.52 DPS) [vendor]; Suspicious Spare Part (274754, -116.97 DPS) [vendor] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (540.40 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Suspicious Spare Part (274754, -99.45 DPS) [vendor] |
| trinket1 | Mark of the Chosen (17774) | The Pariah's Instructions [quest] | sim-verified (12134.7 DPS) | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Talisman of Arathor (21117, +0.00 DPS) [rep]; Relentless Raider's Seal (272060, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (12134.7 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -38.60 DPS) [quest] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (12134.7 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -24.26 DPS) [world_drop]; Ragehammer (10626, -698.23 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Master Sergeant's Insignia; shoulder: Knight-Lieutenant's Leather Shoulders; back: Graverot Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Shadefiend Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Mark of the Chosen; trinket2: Darkspear Voodoo Seal; main_hand: Radiant Staff

No-known-source sample (15 of 577, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

### Band 60 (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 138.1. Weights run: 3.0s. Verify run: 4.8s. 1449 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.140 ± 0.015, defense=1.303 ± 0.250 per rating point (1 rating = 1%, 1.303 per %), dodge=0.708 ± 0.045 per rating point (12 rating = 1%, 8.495 per %), strength=0.119 ± 0.000, agility=0.547 ± 0.029, attack_power=0.059 ± 0.000, hit=0.184 ± 0.036 per rating point (10 rating = 1%, 1.839 per %), crit=0.086 ± 0.016 per rating point (14 rating = 1%, 1.198 per %), expertise=6.074 ± 0.393

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Faceguard (226801) | Mokvar [vendor] | sim-verified (24764.8 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -467.68 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 21.6 stamina points (1703.59 DPS) | yes | Medallion of Grand Marshal Morris (13091, -121.73 DPS) [world_drop]; Sentinel's Medallion (19538, -266.53 DPS) [rep]; Evil Eye Pendant (18381, -335.71 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 64.0 stamina points (5053.97 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -111.60 DPS) [vendor]; Feralheart Pauldrons (226798, -474.38 DPS, sim-verified) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -761.00 DPS) [pvp] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | 32.7 stamina points (2580.58 DPS) | yes | Stoneskin Gargoyle Cape (13397, -32.90 DPS) [dungeon]; Redoubt Cloak (18495, -176.82 DPS) [dungeon]; Shifting Cloak (18511, -582.74 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, -25.59 DPS) [vendor]; Tunic of Undead Slaying (23089, -3182.78 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Subterfuge (22668, -151.97 DPS) [quest]; Forest Stalker's Bracers (19587, -230.31 DPS) [rep]; Wristwraps of Undead Slaying (23093, -773.23 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 52.9 stamina points (4178.53 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -335.01 DPS) [vendor]; Raider Gloves (272099, -352.67 DPS, sim-verified) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 47.7 stamina points (3764.43 DPS) | yes | Shifter's Belt (272396, -551.46 DPS, sim-verified) [vendor]; Hivethrasher's Girdle (275614, -772.92 DPS) [crafted]; Belt of Preserved Heads (20216, -774.19 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 84.4 stamina points (6662.85 DPS) | yes | Marshal's Dragonhide Leggings (231691, +0.00 DPS) [vendor]; Dire Warbear Woolies (227804, -408.82 DPS, sim-verified) [vendor]; Marshal's Dragonhide Legguards (231703, -494.56 DPS) [pvp] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 63.6 stamina points (5021.20 DPS) | yes | Marshal's Dragonhide Treads (231692, -375.84 DPS) [vendor]; Drudge Boots (21532, -856.38 DPS) [quest]; Feralheart Treads (226803, -915.45 DPS, sim-verified) [vendor] |
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

Set DPS (verified): 314.0. Weights run: 3.0s. Verify run: 6.4s. 1449 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.163 ± 0.017, defense=1.681 ± 0.313 per rating point (1 rating = 1%, 1.681 per %), dodge=0.990 ± 0.058 per rating point (12 rating = 1%, 11.874 per %), strength=0.130 ± 0.000, agility=0.838 ± 0.037, attack_power=0.059 ± 0.000, hit=0.260 ± 0.044 per rating point (10 rating = 1%, 2.602 per %), crit=0.123 ± 0.020 per rating point (14 rating = 1%, 1.716 per %), expertise=7.837 ± 0.475

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Faceguard (226801) | Mokvar [vendor] | sim-verified (46078.4 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -1119.55 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 27.2 stamina points (2835.28 DPS) | yes | Evil Eye Pendant (18381, -301.07 DPS) [dungeon]; Medallion of Grand Marshal Morris (13091, -355.58 DPS) [world_drop]; Talisman of Evasion (13177, -464.07 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 75.6 stamina points (7874.55 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -529.99 DPS) [vendor]; Darkspear Pauldrons (272105, -1146.31 DPS, sim-verified) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -1767.15 DPS) [pvp] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (46078.4 DPS) | yes | Stoneskin Gargoyle Cape (13397, -21.84 DPS) [dungeon]; Windshear Cape (20691, -128.13 DPS) [world]; Shifting Cloak (18511, -491.25 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-verified (46078.4 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, -911.89 DPS) [vendor]; Tunic of Undead Slaying (23089, -5028.31 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-verified (46078.4 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, -8.13 DPS) [quest] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 63.6 stamina points (6620.05 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Raider Gloves (272099, -462.29 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -1235.15 DPS) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 55.4 stamina points (5768.21 DPS) | yes | Shifter's Belt (272396, -570.99 DPS) [vendor]; Belt of Preserved Heads (20216, -1017.30 DPS) [quest]; Ferocity of the Timbermaw (227805, -1108.22 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 97.8 stamina points (10191.25 DPS) | yes | Marshal's Dragonhide Leggings (231691, -313.13 DPS) [vendor]; Dire Warbear Woolies (227804, -548.20 DPS, sim-verified) [vendor]; Marshal's Dragonhide Legguards (231703, -1614.13 DPS) [pvp] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 75.2 stamina points (7829.88 DPS) | yes | Marshal's Dragonhide Treads (231692, -982.68 DPS) [vendor]; Drudge Boots (21532, -1075.68 DPS, sim-verified) [quest]; Feralheart Treads (226803, -1664.70 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (21196) | The Path of the Protector [quest] | 27.7 stamina points (2886.02 DPS) | yes | Band of Resolution (22680, -135.78 DPS) [quest]; Ring of Awareness (272409, -160.25 DPS) [vendor]; Band of the Steadfast Hero (22331, -315.61 DPS) [dungeon] |
| finger2 | Naglering (11669) | Blackrock Depths: Golem Lord Argelmach [dungeon] | 26.6 stamina points (2766.53 DPS) | yes | Band of Resolution (22680, -16.29 DPS) [quest]; Ring of Awareness (272409, -40.77 DPS) [vendor]; Band of the Steadfast Hero (22331, -196.12 DPS) [dungeon] |
| trinket1 | Mark of Tyranny (13966) | General Drakkisath's Demise [quest] | sim-verified (46078.4 DPS) | yes | Stormpike Insignia Rank 6 (17904, -1821.70 DPS) [quest]; Vigilance Charm (18370, -1821.70 DPS) [dungeon]; Talisman of Arathor (20071, -3045.51 DPS) [rep] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (46078.4 DPS) | yes | Force of Will (11810, +0.00 DPS) [dungeon]; Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Vigilance Charm (18370, +0.00 DPS) [dungeon] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (46078.4 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Electrified Dagger (19100, -1710.76 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Faceguard; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Naglering; trinket1: Mark of Tyranny; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1449, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 36.1. Weights run: 2.2s. Verify run: 1.2s. 183 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.083 ± 0.002, defense=0.183 ± 0.035 per rating point (1 rating = 1%, 0.183 per %), dodge=0.129 ± 0.010 per rating point (12 rating = 1%, 1.544 per %), strength=0.070 ± 0.000, agility=0.172 ± 0.006, attack_power=0.035 ± 0.000, hit=0.057 ± 0.006 per rating point (10 rating = 1%, 0.572 per %), crit=0.023 ± 0.001 per rating point (14 rating = 1%, 0.321 per %), expertise=1.490 ± 0.069

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.6 stamina points (434.84 DPS) | yes | Brawler's Leather Hood (252504, -20.51 DPS) [crafted]; Defender's Leather Hood (252447, -39.67 DPS) [crafted]; Totemic Leather Hood (252448, -52.71 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 6.7 stamina points (156.52 DPS) | yes | Scholarly Pendant (277203, -16.10 DPS) [quest]; Scout's Medallion (20442, -85.56 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (218.32 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -65.69 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -72.10 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.7 stamina points (179.39 DPS) | yes | Sentry Cloak (2059, -5.35 DPS) [world_drop]; Grave Shroud (279865, -10.46 DPS) [quest]; Miner's Cape (5444, -20.14 DPS) [dungeon] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.2 stamina points (426.05 DPS) | yes | Gloomshroud Armor (1489, -28.55 DPS, sim-verified) [dungeon]; Loch Croc Hide Vest (6197, -75.15 DPS) [world]; Trapper's Leather Armor (252491, -115.23 DPS) [crafted] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.8 stamina points (207.04 DPS) | yes | Bear Bracers (4795, -17.92 DPS) [vendor]; Savannah Bracers (15453, -19.87 DPS) [quest]; Drakewing Bands (12999, -56.94 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 9.1 stamina points (212.99 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.95 DPS) [crafted]; Defender's Leather Gloves (252496, -9.58 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.3 stamina points (264.82 DPS) | yes | Belt of the Fang (10412, -22.74 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -61.57 DPS) [crafted]; Guardsman Belt (3429, -63.52 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 14.1 stamina points (330.98 DPS) | yes | Slick Deviate Leggings (6480, -4.91 DPS) [quest]; Smelting Pants (5199, -19.06 DPS) [dungeon]; Brawler's Leather Pants (252500, -31.57 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (350.99 DPS) | yes | Surfer Shoes (276274, -41.47 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -67.55 DPS) [dungeon]; Trapper's Leather Boots (252440, -94.98 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.4 stamina points (125.60 DPS) | yes | Ring of Scorn (3235, -31.99 DPS) [quest]; Ring of the Moon (12052, -53.76 DPS) [world_drop]; Deep Fathom Ring (6463, -55.39 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (117.02 DPS) | yes | Ring of Scorn (3235, -16.16 DPS, sim-verified) [quest]; Ring of the Moon (12052, -45.18 DPS) [world_drop]; Deep Fathom Ring (6463, -46.81 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (2199.8 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps: Fallenroot Shadowstalker [dungeon] | 18.6 stamina points (435.06 DPS) | yes | Crescent Staff (6505, -10.98 DPS) [quest]; Hammerbone (270018, -21.72 DPS) [quest]; Advisor's Gnarled Staff (20425, -23.52 DPS) [pvp] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 183, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 5821 Darkstalker Boots; 5968 Rugged Boots; 6478 Rat Stompers; 7187 VanCleef's Boots; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 42.1. Weights run: 2.1s. Verify run: 1.4s. 315 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.318 ± 0.057 per rating point (1 rating = 1%, 0.318 per %), dodge=0.139 ± 0.014 per rating point (12 rating = 1%, 1.662 per %), strength=0.065 ± 0.000, agility=0.173 ± 0.009, attack_power=0.032 ± 0.000, hit=0.072 ± 0.008 per rating point (10 rating = 1%, 0.724 per %), crit=0.024 ± 0.001 per rating point (14 rating = 1%, 0.331 per %), expertise=1.823 ± 0.100

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.7 stamina points (477.85 DPS) | yes | Defender's Leather Helm (252455, -25.90 DPS, sim-verified) [crafted]; Totemic Leather Helm (252456, -38.03 DPS) [crafted]; Trapper's Leather Helm (252513, -38.03 DPS) [crafted] |
| neck | Souvenier Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (285.66 DPS) | yes | Ghostshard Talisman (7731, -62.75 DPS, sim-verified) [dungeon]; River Pride Choker (13087, -82.20 DPS) [world_drop]; Senior Sergeant's Insignia (15200, -87.89 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery: Houndmaster Loksey [dungeon] | 17.9 stamina points (392.80 DPS) | yes | Tanned Shoulderpads (270023, -51.99 DPS, sim-verified) [quest]; Cloudy Gustwoven Spaulders (277043, -72.89 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -72.89 DPS) [crafted] |
| back | Sergeant's Cloak (18427) | PvP rank 7 · Sergeant · Horde [vendor] | 11.2 stamina points (245.48 DPS) | yes | Tigerstrike Mantle (13108, -15.36 DPS) [world_drop]; Enduring Cape (14763, -40.83 DPS) [world_drop]; Grimsteel Cape (4643, -63.82 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.2 stamina points (465.51 DPS) | yes | Spirewind Fetter (9406, -8.29 DPS) [dungeon]; Green Whelp Armor (7375, -36.61 DPS) [crafted]; Blackened Defias Armor (10399, -65.33 DPS) [dungeon] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (241.68 DPS) | yes | Barbaric Bracers (18948, -2.68 DPS) [crafted]; Black Wolf Bracers (3230, -6.11 DPS) [dungeon]; Technician's Bracers (270042, -38.44 DPS) [quest] |
| hands | Ebon Vise (7690) | Scarlet Monastery: Fallen Champion [dungeon] | 15.1 stamina points (332.77 DPS) | yes | Braced Handguards (6784, -34.88 DPS) [quest]; Brawler Gloves (720, -53.78 DPS) [world_drop]; Wolfclaw Gloves (1978, -59.04 DPS) [dungeon] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.7 stamina points (344.74 DPS) | yes | Skulker's Leather Belt (252520, -37.56 DPS, sim-verified) [crafted]; Prowler's Leather Belt (252459, -61.65 DPS) [crafted]; Stalker's Leather Belt (252521, -63.06 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.2 stamina points (399.52 DPS) | yes | Brawler's Leather Legguards (252516, -36.34 DPS) [crafted]; Dark Ritual Leggings (270031, -40.26 DPS) [quest]; Trapper's Leather Legguards (252517, -46.31 DPS) [crafted] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.3 stamina points (403.16 DPS) | yes | Harbinger Boots (7754, -28.30 DPS) [dungeon]; Grizzled Boots (6335, -54.36 DPS) [quest]; Highlander's Mail Greaves (20123, -61.14 DPS) [vendor] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.3 stamina points (182.20 DPS) | yes | Plains Ring (2039, -6.41 DPS) [dungeon]; Darkspear Signet (272071, -6.41 DPS) [vendor]; Defias Renegade Ring (1076, -53.31 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.2 stamina points (180.06 DPS) | yes | Plains Ring (2039, -4.27 DPS) [dungeon]; Darkspear Signet (272071, -4.27 DPS) [vendor]; Defias Renegade Ring (1076, -51.18 DPS) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (3429.0 DPS) | yes | Defiler's Talisman (21120, +0.00 DPS) [rep]; Rune of Perfection (21566, -43.95 DPS) [rep]; Rune of Duty (21568, -43.95 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272062) | Creeg Bothunk [vendor] | sim-verified (3429.0 DPS) | yes | Defiler's Talisman (21120, +0.00 DPS) [rep]; Rune of Perfection (21566, -43.95 DPS) [rep]; Rune of Duty (21568, -43.95 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul: Overlord Ramtusk [dungeon] | sim-verified (3429.0 DPS) | yes | Advisor's Gnarled Staff (19569, -33.57 DPS) [pvp]; Glimmering Staff (249392, -65.57 DPS) [crafted]; Manual Crowd Pummeler (9449, -229.51 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenier Sea Shell; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Ebon Vise; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Staff of the Shade

No-known-source sample (15 of 315, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 40 (tauren, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 73.1. Weights run: 2.6s. Verify run: 1.4s. 426 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.027, armor=0.125 ± 0.007, defense=0.391 ± 0.084 per rating point (1 rating = 1%, 0.391 per %), dodge=0.177 ± 0.019 per rating point (12 rating = 1%, 2.129 per %), strength=0.070 ± 0.000, agility=0.191 ± 0.012, attack_power=0.035 ± 0.000, hit=0.066 ± 0.015 per rating point (10 rating = 1%, 0.665 per %), crit=0.031 ± 0.006 per rating point (14 rating = 1%, 0.434 per %), expertise=2.366 ± 0.167

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman: Shadowforge Relic Hunter [dungeon] | 27.5 stamina points (818.13 DPS) | yes | Warden's Wizard Hat (14604, -47.95 DPS) [world_drop]; Brawler's Leather Helm (252512, -61.35 DPS) [crafted]; Nightscape Headband (8176, -83.44 DPS) [crafted] |
| neck | Shriveled Heart (9243) (or Souvenier Sea Shell (274749)) | Zul'Farrak: Sandarr Dunereaver [dungeon] | 13.0 stamina points (386.96 DPS) | yes | Souvenier Sea Shell (274749, +0.00 DPS) [vendor]; Dragon's Blood Necklace (10711, -29.77 DPS) [quest]; Gazlowe's Charm (13088, -51.16 DPS) [dungeon] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.4 stamina points (845.05 DPS) | yes | Flintrock Shoulders (7755, -130.52 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -213.23 DPS) [dungeon]; Tanned Shoulderpads (270023, -305.13 DPS) [quest] |
| back | Silky Spider Cape (10776) | Razorfen Downs: Tuten'kash [dungeon] | 14.8 stamina points (439.22 DPS) | yes | Wing of the Whelpling (13121, -14.50 DPS) [world_drop]; Well Oiled Cloak (12254, -33.49 DPS) [vendor]; Heraldic Cloak (8120, -36.81 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 31.5 stamina points (937.66 DPS) | yes | Warden's Wraps (14601, -65.14 DPS) [world_drop]; Robes of the Lich (10762, -111.30 DPS) [dungeon]; Insignia Chestguard (4057, -166.88 DPS) [world_drop] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 14.7 stamina points (436.21 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Barbaric Bracers (18948, -51.38 DPS) [crafted]; Branded Leather Bracers (19508, -57.63 DPS, sim-verified) [dungeon] |
| hands | Stalker's Leather Gloves (252526) | Leatherworking [crafted] | 20.4 stamina points (607.90 DPS) | yes | Warden's Leather Gloves (252527, -8.19 DPS) [crafted]; Razzeric's Racing Grips (6727, -33.72 DPS) [quest]; Bonefingers (10765, -38.17 DPS) [dungeon] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 20.2 stamina points (602.43 DPS) | yes | Warden's Leather Belt (252460, -53.13 DPS) [crafted]; Kolkar Hunter's Belt (6788, -61.21 DPS, sim-verified) [quest]; Scorpashi Sash (14652, -75.50 DPS) [world_drop] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 26.2 stamina points (778.49 DPS) | yes | Panther Hunter Leggings (4108, -36.53 DPS) [quest]; Warden's Woolies (14605, -37.48 DPS) [world_drop]; Imperial Leather Pants (4062, -50.18 DPS, sim-verified) [dungeon] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 24.3 stamina points (724.67 DPS) | yes | Skulker's Leather Shoes (252531, -96.35 DPS) [crafted]; Prowler's Leather Shoes (252465, -110.69 DPS) [crafted]; Gnomebot Operating Boots (9450, -123.33 DPS, sim-verified) [dungeon] |
| finger1 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.5 stamina points (342.09 DPS) | yes | Darkspear Signet (272070, -14.66 DPS) [vendor]; Underworld Band (1980, -44.42 DPS) [world_drop]; Dragonclaw Ring (10710, -44.42 DPS) [quest] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 11.4 stamina points (339.99 DPS) | yes | Darkspear Signet (272070, -12.56 DPS) [vendor]; Underworld Band (1980, -42.33 DPS) [world_drop]; Dragonclaw Ring (10710, -42.33 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | sim-verified (6613.2 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -29.77 DPS) [world] |
| trinket2 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (6613.2 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -29.77 DPS) [world] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (6613.2 DPS) | yes | Ironshod Bludgeon (9408, -15.12 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -162.92 DPS) [quest]; Manual Crowd Pummeler (9449, -583.49 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; neck: Shriveled Heart; shoulder: Fleshhide Shoulders; back: Silky Spider Cape; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Stalker's Leather Gloves; waist: Ogron's Sash; legs: Basilisk Hide Pants; feet: Warden's Leather Shoes; finger1: Suspicious Spare Part; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Cragwood Maul

No-known-source sample (15 of 426, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 5000 Coral Band; 5004 Mark of the Kirin Tor; 5005 Emberspark Pendant; 5008 Quicksilver Ring; 5010 Inscribed Gold Ring

### Band 50 (tauren, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 117.1. Weights run: 2.9s. Verify run: 1.9s. 561 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.035, armor=0.127 ± 0.011, defense=not significant (0.491 ± 0.126) per rating point (1 rating = 1%, 0.491 per %), dodge=0.246 ± 0.029 per rating point (12 rating = 1%, 2.951 per %), strength=0.061 ± 0.000, agility=0.252 ± 0.018, attack_power=0.030 ± 0.000, hit=not significant (0.069 ± 0.023) per rating point (10 rating = 1%, 0.694 per %), crit=not significant (0.035 ± 0.010) per rating point (14 rating = 1%, 0.493 per %), expertise=3.041 ± 0.263

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 39.7 stamina points (1531.50 DPS) | yes | Blood Guard's Leather Headband (220851, -182.07 DPS) [vendor]; Embrace of the Lycan (9479, -280.86 DPS, sim-verified) [dungeon]; Impractical Headwarmer (274756, -328.29 DPS) [vendor] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (540.40 DPS) | yes | Shriveled Heart (9243, -38.60 DPS) [dungeon]; Darkspear Warding Pendant (272073, -38.60 DPS) [vendor]; Souvenier Sea Shell (274749, -38.60 DPS) [vendor] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 28.9 stamina points (1115.80 DPS) | yes | Penance Spaulders (11963, -19.11 DPS) [quest]; Phytoskin Spaulders (17749, -28.62 DPS) [dungeon]; Blood Guard's Leather Shoulders (220853, -56.43 DPS) [vendor] |
| back | Graverot Cape (11677) | Blackrock Depths: Anub'shiah [dungeon] | 20.0 stamina points (770.57 DPS) | yes | Nightfall Drape (12465, -43.51 DPS) [dungeon]; Sergeant's Cloak (16341, -53.34 DPS) [vendor]; Grovekeeper's Drape (17739, -125.62 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 52.3 stamina points (2019.43 DPS) | yes | Mixologist's Tunic (12793, -378.89 DPS, sim-verified) [dungeon]; Jinxed Hoodoo Skin (9473, -516.74 DPS) [dungeon]; Heraldic Breastplate (8119, -574.96 DPS) [world_drop] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 23.7 stamina points (913.25 DPS) | yes | First Sergeant's Dragonhide Armguards (18436, -58.48 DPS) [vendor]; Forest Stalker's Bracers (19589, -87.94 DPS) [pvp]; Serpentskin Bracers (8257, -174.51 DPS) [world_drop] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 37.6 stamina points (1449.92 DPS) | yes | Warden's Leather Gauntlets (252549, -314.49 DPS) [crafted]; Raider Gloves (272100, -366.61 DPS, sim-verified) [vendor]; Feralheart Fists (226793, -388.43 DPS) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 26.1 stamina points (1005.60 DPS) | yes | Skulker's Leather Waistguard (252474, -17.95 DPS) [crafted]; Prowler's Leather Waistguard (252473, -32.77 DPS) [crafted]; Girdle of Beastial Fury (11686, -128.71 DPS) [dungeon] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 33.8 stamina points (1304.04 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -214.02 DPS) [vendor]; Scorpashi Leggings (14659, -231.52 DPS) [world_drop]; Windscale Sarong (10842, -586.29 DPS, sim-verified) [world] |
| feet | Shadefiend Boots (11675) | Blackrock Depths: Anub'shiah [dungeon] | 30.3 stamina points (1167.78 DPS) | yes | Slitherscale Boots (10801, -49.44 DPS) [dungeon]; Warden's Leather Boots (252470, -129.43 DPS) [crafted]; Skulker's Leather Boots (252469, -157.13 DPS) [crafted] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.5 stamina points (557.92 DPS) | yes | Ring of Saviors (1447, -17.52 DPS) [world_drop]; Darkspear Signet (272069, -17.52 DPS) [vendor]; Suspicious Spare Part (274754, -116.97 DPS) [vendor] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (540.40 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Suspicious Spare Part (274754, -99.45 DPS) [vendor] |
| trinket1 | Defiler's Talisman (21115) | The Defilers [rep] | sim-verified (12052.9 DPS) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -38.60 DPS) [quest]; Mark of the Chosen (17774, -145.35 DPS, sim-verified) [quest] |
| trinket2 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Relentless Raider's Seal (272060, +0.00 DPS) [vendor]; Guardian Talisman (1490, -38.60 DPS) [quest]; Rune of Perfection (21565, -77.20 DPS) [rep] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-decided (no score - a real sim tournament chose this pick) | yes | Advisor's Gnarled Staff (19567, +0.00 DPS) [pvp]; Dreamstaff (249454, +0.00 DPS) [crafted]; Shadowblade (2163, -670.82 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; back: Graverot Cape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Shadefiend Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Radiant Staff

No-known-source sample (15 of 561, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 60 (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 140.4. Weights run: 3.0s. Verify run: 4.6s. 1446 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.140 ± 0.015, defense=1.303 ± 0.250 per rating point (1 rating = 1%, 1.303 per %), dodge=0.708 ± 0.045 per rating point (12 rating = 1%, 8.495 per %), strength=0.119 ± 0.000, agility=0.547 ± 0.029, attack_power=0.059 ± 0.000, hit=0.184 ± 0.036 per rating point (10 rating = 1%, 1.839 per %), crit=0.086 ± 0.016 per rating point (14 rating = 1%, 1.198 per %), expertise=6.074 ± 0.393

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Faceguard (226801) | Mokvar [vendor] | sim-verified (26272.5 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -454.90 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 21.6 stamina points (1703.59 DPS) | yes | Medallion of Grand Marshal Morris (13091, -121.73 DPS) [world_drop]; Scout's Medallion (19534, -266.53 DPS) [rep]; Evil Eye Pendant (18381, -335.71 DPS) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 64.0 stamina points (5053.97 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -111.60 DPS) [vendor]; Feralheart Pauldrons (226798, -649.30 DPS, sim-verified) [vendor]; Warlord's Dragonhide Spaulders (231681, -761.00 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | 32.7 stamina points (2580.58 DPS) | yes | Stoneskin Gargoyle Cape (13397, -32.90 DPS) [dungeon]; Redoubt Cloak (18495, -176.82 DPS) [dungeon]; Shifting Cloak (18511, -755.00 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -25.59 DPS) [vendor]; Tunic of Undead Slaying (23089, -3149.23 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-decided (no score - a real sim tournament chose this pick) | yes | Bracers of Subterfuge (22668, -151.97 DPS) [quest]; Forest Stalker's Bracers (19587, -230.31 DPS) [rep]; Wristwraps of Undead Slaying (23093, -661.32 DPS, sim-verified) [world] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 52.9 stamina points (4178.53 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, +0.00 DPS) [vendor]; General's Dragonhide Gloves (231677, -335.01 DPS) [pvp] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 47.7 stamina points (3764.43 DPS) | yes | Shifter's Belt (272396, -595.06 DPS, sim-verified) [vendor]; Hivethrasher's Girdle (275614, -772.92 DPS) [crafted]; Belt of Preserved Heads (20216, -774.19 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 84.4 stamina points (6662.85 DPS) | yes | General's Dragonhide Leggings (231685, +0.00 DPS) [pvp]; Dire Warbear Woolies (227804, -417.87 DPS, sim-verified) [vendor]; General's Dragonhide Legguards (231673, -494.56 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 63.6 stamina points (5021.20 DPS) | yes | General's Dragonhide Treads (231683, -375.84 DPS) [vendor]; Feralheart Treads (226803, -491.76 DPS, sim-verified) [vendor]; Drudge Boots (21532, -856.38 DPS) [quest] |
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

Set DPS (verified): 318.7. Weights run: 3.0s. Verify run: 5.9s. 1446 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.163 ± 0.017, defense=1.681 ± 0.313 per rating point (1 rating = 1%, 1.681 per %), dodge=0.990 ± 0.058 per rating point (12 rating = 1%, 11.874 per %), strength=0.130 ± 0.000, agility=0.838 ± 0.037, attack_power=0.059 ± 0.000, hit=0.260 ± 0.044 per rating point (10 rating = 1%, 2.602 per %), crit=0.123 ± 0.020 per rating point (14 rating = 1%, 1.716 per %), expertise=7.837 ± 0.475

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Feralheart Faceguard (226801) | Mokvar [vendor] | sim-verified (48933.6 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Outlaw's Collar (279253, -1054.90 DPS, sim-verified) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 27.2 stamina points (2835.28 DPS) | yes | Medallion of Grand Marshal Morris (13091, -355.58 DPS) [world_drop]; Talisman of Evasion (13177, -464.07 DPS) [dungeon]; Evil Eye Pendant (18381, -683.51 DPS, sim-verified) [dungeon] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 75.6 stamina points (7874.55 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -529.99 DPS) [vendor]; Darkspear Pauldrons (272105, -1225.06 DPS, sim-verified) [vendor]; Warlord's Dragonhide Spaulders (231681, -1767.15 DPS) [vendor] |
| back | Stoneshield Cloak (12551) | Blackrock Depths: Anvilrage Overseer [dungeon] | sim-verified (48933.6 DPS) | yes | Stoneskin Gargoyle Cape (13397, -21.84 DPS) [dungeon]; Windshear Cape (20691, -128.13 DPS) [world]; Shifting Cloak (18511, -719.05 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | sim-verified (48933.6 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -911.89 DPS) [vendor]; Tunic of Undead Slaying (23089, -4899.48 DPS, sim-verified) [world] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | sim-verified (48933.6 DPS) | yes | Forest Stalker's Bracers (19587, +0.00 DPS) [rep]; Bracers of Subterfuge (22668, -8.13 DPS) [quest] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 63.6 stamina points (6620.05 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, -571.26 DPS, sim-verified) [vendor]; General's Dragonhide Gloves (231677, -1235.15 DPS) [pvp] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 55.4 stamina points (5768.21 DPS) | yes | Shifter's Belt (272396, -733.39 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -1017.30 DPS) [quest]; Ferocity of the Timbermaw (227805, -1108.22 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 97.8 stamina points (10191.25 DPS) | yes | General's Dragonhide Leggings (231685, -313.13 DPS) [pvp]; Dire Warbear Woolies (227804, -781.77 DPS, sim-verified) [vendor]; General's Dragonhide Legguards (231673, -1614.13 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 75.2 stamina points (7829.88 DPS) | yes | General's Dragonhide Treads (231683, -982.68 DPS) [vendor]; Drudge Boots (21532, -1009.72 DPS, sim-verified) [quest]; Feralheart Treads (226803, -1664.70 DPS) [vendor] |
| finger1 | Thrall's Resolve (12544) | The Princess Saved? [quest] | sim-verified (48933.6 DPS) | yes | Band of Resolution (22680, -582.05 DPS) [quest]; Ring of Awareness (272409, -606.53 DPS) [vendor]; Naglering (11669, -801.81 DPS, sim-verified) [dungeon] |
| finger2 | Signet Ring of the Bronze Dragonflight (21196) | The Path of the Protector [quest] | sim-verified (48933.6 DPS) | yes | Band of Resolution (22680, -135.78 DPS) [quest]; Ring of Awareness (272409, -160.25 DPS) [vendor]; Naglering (11669, -500.80 DPS, sim-verified) [dungeon] |
| trinket1 | Mark of Tyranny (13966) | For The Horde! [quest] | sim-verified (48933.6 DPS) | yes | Frostwolf Insignia Rank 6 (17909, -1821.70 DPS) [quest]; Vigilance Charm (18370, -1821.70 DPS) [dungeon]; Defiler's Talisman (20072, -3045.51 DPS) [rep] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (48933.6 DPS) | yes | Frostwolf Insignia Rank 6 (17909, +0.00 DPS) [quest]; Vigilance Charm (18370, +0.00 DPS) [dungeon]; Rune of the Guard Captain (19120, -1262.10 DPS, sim-verified) [quest] |
| main_hand | Headmaster's Charge (13937) | Scholomance: Darkmaster Gandling [dungeon] | sim-verified (48933.6 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glacial Blade (19099, -2372.68 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Feralheart Faceguard; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Stoneshield Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Thrall's Resolve; finger2: Signet Ring of the Bronze Dragonflight; trinket1: Mark of Tyranny; trinket2: Defender's Grip Stabilizer; main_hand: Headmaster's Charge

No-known-source sample (15 of 1446, see the JSON for more): 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 4990 Scorched Bands; 5000 Coral Band; 5004 Mark of the Kirin Tor

