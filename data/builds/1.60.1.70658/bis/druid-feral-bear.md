# Leveling BiS: Feral Bear

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 35.2. Weights run: 3.7s. Verify run: 1.8s. 186 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.082 ± 0.002, defense=0.140 ± 0.026 per rating point (1 rating = 1%, 0.140 per %), dodge=0.096 ± 0.008 per rating point (12 rating = 1%, 1.148 per %), strength=0.052 ± 0.000, agility=0.148 ± 0.005, attack_power=0.026 ± 0.000, hit=0.045 ± 0.005 per rating point (10 rating = 1%, 0.449 per %), crit=0.019 ± 0.001 per rating point (14 rating = 1%, 0.261 per %), expertise=1.200 ± 0.059

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.5 stamina points (352.66 DPS) | yes | Brawler's Leather Hood (252504, -21.07 DPS) [crafted]; Defender's Leather Hood (252447, -35.69 DPS) [crafted]; Totemic Leather Hood (252448, -43.63 DPS) [crafted] |
| neck | Tarnished Locket (279870) | Remember That I Love You [quest] | 5.0 stamina points (95.17 DPS) | yes | Erudite's Amulet (277204, -28.99 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -38.07 DPS) [quest]; Sentinel's Medallion (20444, -40.18 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (176.09 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -55.79 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -57.79 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.6 stamina points (145.44 DPS) | yes | Sentry Cloak (2059, -6.19 DPS) [world_drop]; Grave Shroud (279865, -10.42 DPS) [quest]; Firebane Cloak (12979, -20.60 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.0 stamina points (341.76 DPS) | yes | Gloomshroud Armor (1489, -27.28 DPS, sim-verified) [dungeon]; Tunic of Westfall (2041, -71.89 DPS) [quest]; Trapper's Leather Armor (252491, -94.10 DPS) [crafted] |
| wrist | Bear Bracers (4795) | Bernard Brubaker [vendor] | 8.0 stamina points (152.96 DPS) | yes | Drakewing Bands (12999, -18.65 DPS, sim-verified) [world_drop]; Wolf Bracers (4794, -45.82 DPS) [vendor]; Sanguine Cuffs (14375, -50.27 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 8.9 stamina points (170.19 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.56 DPS) [crafted]; Defender's Leather Gloves (252496, -7.31 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.1 stamina points (210.93 DPS) | yes | Belt of the Fang (10412, -28.35 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.55 DPS) [crafted]; Guardsman Belt (3429, -50.11 DPS) [world] |
| legs | Duty Bound Leggings (279868) | Bloodied Insignia [quest] | 16.5 stamina points (314.73 DPS) | yes | Defender's Leather Pants (252445, -46.40 DPS, sim-verified) [crafted]; Slick Deviate Leggings (6480, -53.08 DPS) [quest]; Smelting Pants (5199, -62.79 DPS) [dungeon] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (284.63 DPS) | yes | Surfer Shoes (276274, -37.67 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -58.24 DPS) [dungeon]; Trapper's Leather Boots (252440, -80.09 DPS) [crafted] |
| finger1 | Sustaining Ring (6743) | Knowledge in the Deeps [quest] | 6.0 stamina points (114.20 DPS) | yes | Blood Ring (4998, -19.03 DPS) [world_drop]; Ring of the Moon (12052, -56.11 DPS) [world_drop]; Deep Fathom Ring (6463, -57.10 DPS) [dungeon] |
| finger2 | Slain Baron's Signet (279867) | Abominable Creatures [quest] | 5.3 stamina points (100.50 DPS) | yes | Blood Ring (4998, -5.33 DPS) [world_drop]; Ring of the Moon (12052, -42.40 DPS) [world_drop]; Deep Fathom Ring (6463, -43.40 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (2168.8 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.2 stamina points (307.87 DPS) | yes | Rakzur Club (12983, -24.79 DPS) [world_drop]; Rhahk'Zor's Hammer (5187, -27.57 DPS) [dungeon]; Twisted Chanter's Staff (890, -51.75 DPS) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Tarnished Locket; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Bear Bracers; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Duty Bound Leggings; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Sustaining Ring; finger2: Slain Baron's Signet; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 186, see the JSON for more): 1189 Overseer's Ring; 3738 Brewing Rod; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5821 Darkstalker Boots; 6478 Rat Stompers; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads

### Band 30 (night-elf, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 41.6. Weights run: 3.5s. Verify run: 1.9s. 346 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.240 ± 0.054 per rating point (1 rating = 1%, 0.240 per %), dodge=0.106 ± 0.013 per rating point (12 rating = 1%, 1.268 per %), strength=0.049 ± 0.000, agility=0.150 ± 0.008, attack_power=0.025 ± 0.000, hit=0.047 ± 0.009 per rating point (10 rating = 1%, 0.466 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.276 per %), expertise=1.556 ± 0.106

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 stamina points (419.72 DPS) | yes | Totemic Leather Helm (252456, -29.23 DPS) [crafted]; Trapper's Leather Helm (252513, -29.23 DPS) [crafted]; Defender's Leather Helm (252455, -40.44 DPS, sim-verified) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (252.60 DPS) | yes | River Pride Choker (13087, -66.61 DPS, sim-verified) [world_drop]; Master Sergeant's Insignia (18442, -77.72 DPS) [vendor]; Darkspear Warding Pendant (272075, -97.15 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 17.9 stamina points (347.95 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -31.02 DPS, sim-verified) [crafted]; Azure Gustwoven Spaulders (277051, -66.71 DPS) [crafted]; Darkshire Shoulderguards (270044, -87.17 DPS) [quest] |
| back | Sergeant's Cape (18440) | PvP rank 7 · Sergeant · Alliance [vendor] | 11.2 stamina points (217.56 DPS) | yes | Tigerstrike Mantle (13108, +0.00 DPS, sim-verified) [world_drop]; Mourning Shawl (6751, -26.00 DPS) [quest]; Enduring Cape (14763, -37.66 DPS) [world_drop] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.3 stamina points (413.47 DPS) | yes | Spirewind Fetter (9406, -9.15 DPS) [dungeon]; Green Whelp Armor (7375, -32.30 DPS) [crafted]; Dokebi Chestguard (14581, -58.29 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (213.05 DPS) | yes | Barbaric Bracers (18948, -3.78 DPS) [crafted]; Black Wolf Bracers (3230, -4.52 DPS) [dungeon]; Staghide Armguards (270021, -32.33 DPS) [quest] |
| hands | Operator's Gloves (270045) | Gyrodrillmatic Excavationators [quest] | 13.0 stamina points (252.73 DPS) | yes | Brawler Gloves (720, -7.25 DPS) [world_drop]; Wolfclaw Gloves (1978, -13.66 DPS) [dungeon]; Gloves of Insight (9698, -16.58 DPS) [quest] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.5 stamina points (301.60 DPS) | yes | Skulker's Leather Belt (252520, -31.15 DPS, sim-verified) [crafted]; Stalker's Leather Belt (252521, -55.29 DPS) [crafted]; Prowler's Leather Belt (252459, -55.41 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.0 stamina points (349.81 DPS) | yes | Dark Ritual Leggings (270031, -30.47 DPS) [quest]; Brawler's Leather Legguards (252516, -33.02 DPS) [crafted]; Duty Bound Leggings (279868, -42.30 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.4 stamina points (357.93 DPS) | yes | Harbinger Boots (7754, -28.46 DPS, sim-verified) [dungeon]; Highlander's Mail Greaves (20123, -57.62 DPS) [vendor]; Nat Pagle's Extreme Anglin' Boots (19969, -65.66 DPS) [quest] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.2 stamina points (159.77 DPS) | yes | Plains Ring (2039, -4.32 DPS) [dungeon]; Sustaining Ring (6743, -43.19 DPS) [quest]; Defias Renegade Ring (1076, -48.00 DPS) [dungeon] |
| finger2 | Darkspear Signet (272071) (or Plains Ring (2039)) | Creeg Bothunk [vendor] | 8.0 stamina points (155.45 DPS) | yes | Plains Ring (2039, +0.00 DPS) [dungeon]; Sustaining Ring (6743, -38.86 DPS) [quest]; Defias Renegade Ring (1076, -43.68 DPS) [dungeon] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | sim-verified (3281.7 DPS) | yes | Rune of Duty (21568, -38.86 DPS) [rep] |
| trinket2 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | sim-verified (3281.7 DPS) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (3281.7 DPS) | yes | Spine of the Devourer (273040, -30.33 DPS) [world]; Glimmering Staff (249392, -62.65 DPS) [crafted]; Manual Crowd Pummeler (9449, -223.83 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenir Sea Shell; shoulder: Watchman Pauldrons; back: Sergeant's Cape; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Operator's Gloves; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Darkspear Signet; trinket1: Talisman of Arathor; trinket2: Rune of Perfection; main_hand: Staff of the Shade

No-known-source sample (15 of 346, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring

### Band 40 (night-elf, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 72.7. Weights run: 4.1s. Verify run: 2.1s. 459 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.018, armor=0.107 ± 0.006, defense=0.328 ± 0.080 per rating point (1 rating = 1%, 0.328 per %), dodge=0.113 ± 0.018 per rating point (12 rating = 1%, 1.360 per %), strength=0.046 ± 0.000, agility=0.136 ± 0.011, attack_power=0.023 ± 0.000, hit=not significant (0.055 ± 0.014) per rating point (10 rating = 1%, 0.550 per %), crit=0.029 ± 0.006 per rating point (14 rating = 1%, 0.407 per %), expertise=1.721 ± 0.160

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 25.3 stamina points (713.32 DPS) | yes | Warden's Wizard Hat (14604, -54.25 DPS) [world_drop]; Brawler's Leather Helm (252512, -54.53 DPS) [crafted]; Defender's Leather Helm (252455, -77.43 DPS) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (367.11 DPS) | yes | Gazlowe's Charm (13088, -51.28 DPS) [dungeon]; Darkspear Warding Pendant (272074, -56.48 DPS) [vendor]; River Pride Choker (13087, -107.76 DPS) [world_drop] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 26.2 stamina points (739.52 DPS) | yes | Flintrock Shoulders (7755, -135.44 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -183.86 DPS) [dungeon]; Pads of the Venom Spider (13103, -277.07 DPS) [world_drop] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.5 stamina points (382.18 DPS) | yes | Silky Spider Cape (10776, -9.37 DPS) [dungeon]; Well Oiled Cloak (12254, -12.38 DPS) [vendor]; Heraldic Cloak (8120, -22.21 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 29.1 stamina points (821.22 DPS) | yes | Warden's Wraps (14601, -74.47 DPS) [world_drop]; Insignia Chestguard (4057, -143.66 DPS) [world_drop]; Robes of the Lich (10762, -165.61 DPS, sim-verified) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 13.4 stamina points (377.50 DPS) | yes | Branded Leather Bracers (19508, -29.30 DPS) [dungeon]; Cultist's Armguards (270032, -40.71 DPS) [quest]; Barbaric Bracers (18948, -45.81 DPS) [crafted] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | 18.5 stamina points (523.16 DPS) | yes | Stalker's Leather Gloves (252526, -1.44 DPS) [crafted]; Bonefingers (10765, -24.87 DPS) [dungeon]; Razzeric's Racing Grips (6727, -27.22 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 18.2 stamina points (514.50 DPS) | yes | Warden's Leather Belt (252460, -39.56 DPS) [crafted]; Scorpashi Sash (14652, -68.17 DPS) [world_drop]; Kolkar Hunter's Belt (6788, -83.62 DPS, sim-verified) [quest] |
| legs | Imperial Leather Pants (4062) | Uldaman: Ancient Treasure [dungeon] | 23.3 stamina points (658.36 DPS) | yes | Warden's Woolies (14605, -9.04 DPS) [world_drop]; Basilisk Hide Pants (1718, -11.02 DPS) [world_drop]; Panther Hunter Leggings (4108, -19.88 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 22.1 stamina points (623.97 DPS) | yes | Worn Running Boots (9398, -54.40 DPS, sim-verified) [dungeon]; Gnomebot Operating Boots (9450, -56.03 DPS) [dungeon]; Harbinger Boots (7754, -96.74 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (338.87 DPS) | yes | Insurgent's Band (272066, -20.44 DPS) [vendor]; Darkspear Signet (272070, -28.24 DPS) [vendor]; Talvash's Gold Ring (9538, -56.48 DPS) [quest] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.3 stamina points (319.73 DPS) | yes | Insurgent's Band (272066, -1.30 DPS) [vendor]; Darkspear Signet (272070, -9.10 DPS) [vendor]; Talvash's Gold Ring (9538, -37.34 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | sim-verified (6287.1 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -28.24 DPS) [world]; Darkspear Voodoo Seal (272059, -28.24 DPS) [vendor] |
| trinket2 | Rune of Duty (21567) | Silverwing Sentinels [rep] | sim-verified (6287.1 DPS) | yes | Talisman of Arathor (21118, +0.00 DPS) [rep]; Arena Master (18706, -28.24 DPS) [world]; Darkspear Voodoo Seal (272059, -28.24 DPS) [vendor] |
| main_hand | Ironshod Bludgeon (9408) | Uldaman: Ironaya [dungeon] | sim-verified (6287.1 DPS) | yes | Windweaver Staff (7757, -138.60 DPS) [dungeon]; Mograine's Might (7723, -153.69 DPS) [dungeon]; Manual Crowd Pummeler (9449, -482.84 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Imperial Leather Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Suspicious Spare Part; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Ironshod Bludgeon

No-known-source sample (15 of 459, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 50 (night-elf, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 117.0. Weights run: 4.5s. Verify run: 4.5s. 598 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.032, armor=0.151 ± 0.012, defense=not significant (0.452 ± 0.148) per rating point (1 rating = 1%, 0.452 per %), dodge=0.255 ± 0.031 per rating point (12 rating = 1%, 3.057 per %), strength=0.057 ± 0.000, agility=0.252 ± 0.020, attack_power=0.029 ± 0.000, hit=not significant (0.096 ± 0.027) per rating point (10 rating = 1%, 0.959 per %), crit=not significant (0.035 ± 0.011) per rating point (14 rating = 1%, 0.488 per %), expertise=3.424 ± 0.295

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.5 stamina points (1513.01 DPS) | yes | Knight-Lieutenant's Leather Headband (220850, -265.17 DPS, sim-verified) [vendor]; Impractical Headwarmer (274756, -302.88 DPS) [vendor]; Scorpashi Skullcap (14658, -347.09 DPS) [world_drop] |
| neck | Master Sergeant's Insignia (18444) | PvP rank 8 · Master Sergeant · Alliance [vendor] | 14.0 stamina points (498.64 DPS) | yes | Darkspear Warding Pendant (272073, -35.62 DPS) [vendor]; Souvenir Sea Shell (274749, -35.62 DPS) [vendor]; Gazlowe's Charm (13088, -98.72 DPS) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) | Captain Dirgehammer [vendor] | sim-verified (11735.2 DPS) | yes | Fleshhide Shoulders (10774, +0.00 DPS) [dungeon]; Warden's Leather Shoulders (252536, -53.99 DPS) [crafted]; Penance Spaulders (11963, -455.86 DPS, sim-verified) [quest] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 19.7 stamina points (703.40 DPS) | yes | Sergeant's Cape (18441, -10.78 DPS) [vendor]; Emberscale Cape (11311, -21.55 DPS) [dungeon]; Grovekeeper's Drape (17739, -76.62 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.1 stamina points (1996.92 DPS) | yes | Jinxed Hoodoo Skin (9473, -525.32 DPS, sim-verified) [dungeon]; Heraldic Breastplate (8119, -549.00 DPS) [world_drop]; Knight's Leather Armor (220854, -568.40 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.2 stamina points (897.34 DPS) | yes | Sergeant Major's Dragonhide Armsplints (18455, -130.10 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -163.86 DPS) [crafted]; Serpentskin Bracers (8257, -166.10 DPS) [world_drop] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 39.9 stamina points (1422.84 DPS) | yes | Warden's Leather Gauntlets (252549, -291.69 DPS) [crafted]; Feralheart Fists (226793, -352.64 DPS) [vendor]; Raider Gloves (272100, -353.75 DPS, sim-verified) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.1 stamina points (1002.60 DPS) | yes | Skulker's Leather Waistguard (252474, -16.84 DPS) [crafted]; Prowler's Leather Waistguard (252473, -30.72 DPS) [crafted]; Pridelord Girdle (14674, -165.39 DPS) [world_drop] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 37.2 stamina points (1326.14 DPS) | yes | Knight's Restored Leather Leggings (220882, -230.53 DPS) [vendor]; Scorpashi Leggings (14659, -242.63 DPS) [world_drop]; Knight's Crackling Leather Leggings (220864, -754.31 DPS, sim-verified) [vendor] |
| feet | Sergeant Major's Leather Boots (220860) | PvP rank 9 · Sergeant Major · Alliance [vendor] | sim-verified (11735.2 DPS) | yes | Prowler's Leather Boots (252468, +0.00 DPS) [crafted]; Skulker's Leather Boots (252469, +0.00 DPS) [crafted]; Warden's Leather Boots (252470, -472.41 DPS, sim-verified) [crafted] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.4 stamina points (513.88 DPS) | yes | Ring of Saviors (1447, -15.24 DPS) [world_drop]; Darkspear Signet (272069, -15.24 DPS) [vendor]; Underworld Band (1980, -86.48 DPS) [world_drop] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (498.64 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Underworld Band (1980, -71.23 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (11735.2 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Guardian Talisman (1490, -35.62 DPS) [quest]; Rune of Perfection (21565, -71.23 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272060) | Creeg Bothunk [vendor] | sim-verified (11735.2 DPS) | yes | Talisman of Arathor (21117, +0.00 DPS) [rep]; Guardian Talisman (1490, -35.62 DPS) [quest]; Rune of Perfection (21565, -71.23 DPS) [rep] |
| main_hand | Radiant Staff (249453) | Enchanting [crafted] | sim-verified (11735.2 DPS) | yes | Dreamstaff (249454, +0.00 DPS) [crafted]; Glowing Brightwood Staff (812, -29.23 DPS) [world_drop]; The Hand of Antu'sul (9639, -524.20 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Master Sergeant's Insignia; shoulder: Knight-Lieutenant's Leather Shoulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Knight's Leather Pants; feet: Sergeant Major's Leather Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Radiant Staff

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 148.6. Weights run: 4.5s. Verify run: 6.2s. 1244 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.019, armor=0.185 ± 0.016, defense=0.985 ± 0.239 per rating point (1 rating = 1%, 0.985 per %), dodge=0.619 ± 0.048 per rating point (12 rating = 1%, 7.423 per %), strength=0.111 ± 0.000, agility=0.508 ± 0.030, attack_power=0.055 ± 0.000, hit=0.160 ± 0.039 per rating point (10 rating = 1%, 1.598 per %), crit=0.075 ± 0.016 per rating point (14 rating = 1%, 1.045 per %), expertise=5.806 ± 0.422

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.8 stamina points (4707.77 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 20.8 stamina points (1342.72 DPS) | yes | Sentinel's Medallion (19538, -203.00 DPS) [rep]; Master Sergeant's Insignia (18443, -242.75 DPS) [vendor]; Medallion of Grand Marshal Morris (13091, -252.29 DPS) [world_drop] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 66.9 stamina points (4327.97 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -348.73 DPS) [pvp]; Highlander's Leather Shoulders (20059, -440.75 DPS, sim-verified) [rep] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 32.9 stamina points (2129.56 DPS) | yes | Fel Cape (279269, -429.58 DPS) [crafted]; Sergeant's Cape (16342, -492.18 DPS) [vendor]; Cloak of Earth and Sky (279270, -497.17 DPS) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | 86.8 stamina points (5618.16 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, +0.00 DPS) [vendor]; Feralheart Armor (226804, -428.20 DPS, sim-verified) [vendor] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | 37.0 stamina points (2391.50 DPS) | yes | Forest Stalker's Bracers (19587, -143.97 DPS) [rep]; Bracers of Subterfuge (22668, -229.58 DPS, sim-verified) [quest]; Sergeant Major's Dragonhide Armsplints (18454, -371.95 DPS) [vendor] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 54.0 stamina points (3495.63 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, +0.00 DPS) [vendor]; Raider Gloves (272099, -237.45 DPS, sim-verified) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 49.5 stamina points (3203.46 DPS) | yes | Shifter's Belt (272396, -498.63 DPS, sim-verified) [vendor]; Hivethrasher's Girdle (275614, -504.10 DPS) [crafted]; Belt of Preserved Heads (20216, -505.58 DPS) [quest] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 91.9 stamina points (5945.92 DPS) | yes | Marshal's Dragonhide Leggings (231691, -6.01 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -359.98 DPS) [pvp]; Dire Warbear Woolies (227804, -547.77 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 65.1 stamina points (4212.67 DPS) | yes | Marshal's Dragonhide Treads (231692, -31.62 DPS) [vendor]; Drudge Boots (21532, -404.08 DPS) [quest]; Feralheart Treads (226803, -506.55 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 32.3 stamina points (2092.26 DPS) | yes | Ring of Awareness (272409, -814.17 DPS) [vendor]; Insurgent's Band (272064, -996.12 DPS) [vendor]; Cutthroat's Signet (272408, -1011.40 DPS) [vendor] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | 22.9 stamina points (1483.43 DPS) | yes | Ring of Awareness (272409, -205.34 DPS) [vendor]; Insurgent's Band (272064, -387.28 DPS) [vendor]; Cutthroat's Signet (272408, -402.57 DPS) [vendor] |
| trinket1 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (20969.1 DPS) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Relentless Raider's Seal (272060, -181.71 DPS) [vendor]; Darkspear Voodoo Seal (272061, -181.71 DPS) [vendor] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-verified (20969.1 DPS) | yes | Stormpike Insignia Rank 6 (17904, +0.00 DPS) [quest]; Relentless Raider's Seal (272060, -129.41 DPS) [vendor]; Darkspear Voodoo Seal (272061, -129.41 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (20969.1 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; The Lobotomizer (19324, -791.27 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Defender's Grip Stabilizer; trinket2: Talisman of Arathor; main_hand: The Unstoppable Force

No-known-source sample (15 of 1244, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60, raid preset (night-elf, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 333.4. Weights run: 4.7s. Verify run: 6.4s. 1244 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.015, armor=0.207 ± 0.018, defense=1.532 ± 0.310 per rating point (1 rating = 1%, 1.532 per %), dodge=0.858 ± 0.059 per rating point (12 rating = 1%, 10.293 per %), strength=0.120 ± 0.000, agility=0.762 ± 0.038, attack_power=0.055 ± 0.000, hit=0.262 ± 0.044 per rating point (10 rating = 1%, 2.623 per %), crit=0.125 ± 0.019 per rating point (14 rating = 1%, 1.755 per %), expertise=7.919 ± 0.482

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 82.7 stamina points (7173.83 DPS) | yes | Field Marshal's Dragonhide Headguard (231689, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Headdress (231701, +0.00 DPS) [vendor]; Feralheart Faceguard (226801, -586.48 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 25.7 stamina points (2226.96 DPS) | yes | Medallion of Grand Marshal Morris (13091, -290.40 DPS) [world_drop]; Sentinel's Medallion (19538, -368.54 DPS) [rep] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 79.7 stamina points (6911.05 DPS) | yes | Field Marshal's Dragonhide Shoulders (231693, -259.01 DPS) [vendor]; Darkspear Pauldrons (272105, -1183.83 DPS, sim-verified) [vendor]; Field Marshal's Dragonhide Spaulders (231699, -1197.77 DPS) [pvp] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 41.2 stamina points (3570.61 DPS) | yes | Cloak of the Honor Guard (20073, -1228.35 DPS) [rep]; Fel Cape (279269, -1242.28 DPS, sim-verified) [crafted]; Cloak of Earth and Sky (279270, -1286.10 DPS) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | 102.4 stamina points (8884.06 DPS) | yes | Field Marshal's Dragonhide Chestpiece (231690, +0.00 DPS) [vendor]; Field Marshal's Dragonhide Tunic (231702, -428.40 DPS) [vendor]; Feralheart Armor (226804, -1399.85 DPS, sim-verified) [vendor] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | 41.7 stamina points (3617.04 DPS) | yes | Forest Stalker's Bracers (19587, -11.10 DPS) [rep]; Bracers of Subterfuge (22668, -28.44 DPS) [quest] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 65.0 stamina points (5638.31 DPS) | yes | Marshal's Dragonhide Grips (231694, +0.00 DPS) [vendor]; Raider Gloves (272099, -233.18 DPS) [vendor]; Marshal's Dragonhide Gloves (231700, -651.08 DPS) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 58.0 stamina points (5030.22 DPS) | yes | Shifter's Belt (272396, -481.08 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -774.95 DPS) [quest]; Ferocity of the Timbermaw (227805, -864.95 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 104.6 stamina points (9075.48 DPS) | yes | Marshal's Dragonhide Leggings (231691, -257.36 DPS) [vendor]; Dire Warbear Woolies (227804, -627.39 DPS) [vendor]; Marshal's Dragonhide Legguards (231703, -1229.47 DPS) [pvp] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 78.5 stamina points (6809.75 DPS) | yes | Marshal's Dragonhide Treads (231692, -630.52 DPS) [vendor]; Drudge Boots (21532, -928.67 DPS, sim-verified) [quest]; Feralheart Treads (226803, -1290.70 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 36.3 stamina points (3148.05 DPS) | yes | Ring of Awareness (272409, -1000.75 DPS) [vendor]; Cutthroat's Signet (272408, -1301.62 DPS) [vendor]; Protector's Band (19514, -1602.23 DPS) [rep] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | 25.7 stamina points (2226.12 DPS) | yes | Ring of Awareness (272409, -78.82 DPS) [vendor]; Cutthroat's Signet (272408, -379.70 DPS) [vendor]; Protector's Band (19514, -680.30 DPS) [rep] |
| trinket1 | Stormpike Insignia Rank 6 (17904) | The Eye of Command [quest] | sim-verified (39819.9 DPS) | yes | Defender's Grip Stabilizer (272440, -57.62 DPS) [vendor]; Smotts' Compass (4130, -892.87 DPS) [quest]; Arena Grand Master (19024, -892.87 DPS) [quest] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | sim-verified (39819.9 DPS) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Smotts' Compass (4130, -148.10 DPS) [quest]; Arena Grand Master (19024, -148.10 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Stormpike Guard [rep] | sim-verified (39819.9 DPS) | yes | Grand Marshal's Battle Hammer (234567, +0.00 DPS) [pvp]; Grand Marshal's Stave (234571, +0.00 DPS) [pvp]; Hand of Edward the Odd (2243, -1178.95 DPS, sim-verified) [world_drop] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Stormpike Insignia Rank 6; trinket2: Talisman of Arathor; main_hand: The Unstoppable Force

No-known-source sample (15 of 1244, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4777 Ironwood Maul; 4778 Heavy Spiked Mace; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 4964 Goblin Smasher; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

## Horde

### Band 20 (tauren, 0000000000000000-55100000000000000000-0000000000000000)

Set DPS (verified): 36.0. Weights run: 3.7s. Verify run: 1.8s. 176 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.000, armor=0.082 ± 0.002, defense=0.140 ± 0.026 per rating point (1 rating = 1%, 0.140 per %), dodge=0.096 ± 0.008 per rating point (12 rating = 1%, 1.148 per %), strength=0.052 ± 0.000, agility=0.148 ± 0.005, attack_power=0.026 ± 0.000, hit=0.045 ± 0.005 per rating point (10 rating = 1%, 0.449 per %), crit=0.019 ± 0.001 per rating point (14 rating = 1%, 0.261 per %), expertise=1.200 ± 0.059

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Lucky Fishing Hat (19972) | Rare Fish - Keefer's Angelfish [quest] | 18.5 stamina points (352.66 DPS) | yes | Brawler's Leather Hood (252504, -21.07 DPS) [crafted]; Defender's Leather Hood (252447, -35.69 DPS) [crafted]; Totemic Leather Hood (252448, -43.63 DPS) [crafted] |
| neck | Erudite's Amulet (277204) | Friend of the Library [quest] | 3.3 stamina points (62.74 DPS) | yes | Scholarly Pendant (277203, -5.64 DPS) [quest]; Scout's Medallion (20442, -7.76 DPS) [rep] |
| shoulder | Forest Leather Mantle (4709) (or Prospector's Pads (14566)) | World drop [world_drop] | 9.3 stamina points (176.09 DPS) | yes | Prospector's Pads (14566, +0.00 DPS) [world_drop]; Serpent's Shoulders (5404, -55.79 DPS) [dungeon]; Double-Stitched Woolen Shoulders (4314, -57.79 DPS) [crafted] |
| back | Sporid Cape (6629) | Wailing Caverns: Verdan the Everliving [dungeon] | 7.6 stamina points (145.44 DPS) | yes | Sentry Cloak (2059, -6.19 DPS) [world_drop]; Grave Shroud (279865, -10.42 DPS) [quest]; Firebane Cloak (12979, -20.60 DPS) [world_drop] |
| chest | Blackened Defias Armor (10399) | The Deadmines: Edwin VanCleef [dungeon] | 18.0 stamina points (341.76 DPS) | yes | Gloomshroud Armor (1489, -38.44 DPS, sim-verified) [dungeon]; Trapper's Leather Armor (252491, -94.10 DPS) [crafted]; Camouflaged Tunic (3585, -95.42 DPS) [quest] |
| wrist | Spare Part Bindings (279875) | Light's Justice [quest] | 8.7 stamina points (166.10 DPS) | yes | Savannah Bracers (15453, -14.71 DPS) [quest]; Bear Bracers (4795, -17.61 DPS, sim-verified) [vendor]; Drakewing Bands (12999, -44.96 DPS) [world_drop] |
| hands | Forest Leather Gloves (3058) (or Trapper's Leather Gloves (252495)) | World drop [world_drop] | 8.9 stamina points (170.19 DPS) | yes | Trapper's Leather Gloves (252495, +0.00 DPS) [crafted]; Nimble Leather Gloves (7285, -1.56 DPS) [crafted]; Defender's Leather Gloves (252496, -7.31 DPS) [crafted] |
| waist | Deviate Scale Belt (6468) | Leatherworking [crafted] | 11.1 stamina points (210.93 DPS) | yes | Belt of the Fang (10412, -31.29 DPS, sim-verified) [dungeon]; Dark Leather Belt (4249, -48.55 DPS) [crafted]; Guardsman Belt (3429, -50.11 DPS) [world] |
| legs | Defender's Leather Pants (252445) | Leatherworking [crafted] | 13.8 stamina points (262.48 DPS) | yes | Slick Deviate Leggings (6480, -0.83 DPS) [quest]; Smelting Pants (5199, -10.54 DPS) [dungeon]; Brawler's Leather Pants (252500, -26.95 DPS) [crafted] |
| feet | Nat Pagle's Extreme Anglin' Boots (19969) | Rare Fish - Brownell's Blue Striped Racer [quest] | 15.0 stamina points (284.63 DPS) | yes | Surfer Shoes (276274, -52.52 DPS, sim-verified) [vendor]; Footpads of the Fang (10411, -58.24 DPS) [dungeon]; Trapper's Leather Boots (252440, -80.09 DPS) [crafted] |
| finger1 | Slain Baron's Signet (279867) | Unending Torment [quest] | 5.3 stamina points (100.50 DPS) | yes | Ring of Scorn (3235, -24.36 DPS) [quest]; Ring of the Moon (12052, -42.40 DPS) [world_drop]; Deep Fathom Ring (6463, -43.40 DPS) [dungeon] |
| finger2 | Blood Ring (4998) | World drop [world_drop] | 5.0 stamina points (95.17 DPS) | yes | Ring of Scorn (3235, -16.75 DPS, sim-verified) [quest]; Ring of the Moon (12052, -37.08 DPS) [world_drop]; Deep Fathom Ring (6463, -38.07 DPS) [dungeon] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (2229.7 DPS) | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | - |
| main_hand | Staff of the Blessed Seer (2271) | Blackfathom Deeps [dungeon] | 16.2 stamina points (307.87 DPS) | yes | Advisor's Gnarled Staff (20425, -19.10 DPS) [pvp]; Hammerbone (270018, -22.80 DPS) [quest]; Crescent Staff (6505, -23.24 DPS, sim-verified) [quest] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 20:** head: Lucky Fishing Hat; neck: Erudite's Amulet; shoulder: Forest Leather Mantle; back: Sporid Cape; chest: Blackened Defias Armor; wrist: Spare Part Bindings; hands: Forest Leather Gloves; waist: Deviate Scale Belt; legs: Defender's Leather Pants; feet: Nat Pagle's Extreme Anglin' Boots; finger1: Slain Baron's Signet; finger2: Blood Ring; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Staff of the Blessed Seer

No-known-source sample (15 of 176, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 3738 Brewing Rod; 5821 Darkstalker Boots; 6478 Rat Stompers; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 18863 Insignia of the Alliance; 209617 Insignia of the Alliance; 241089 Scarlet Dagger; 248005 Northshire Hammer; 248007 Militia Shortblade; 248008 Apprentice's Spellstaff; 263006 Scout Ranger's Tunic; 263007 Skyseer's Vest

### Band 30 (tauren, 0000000000000000-55230330000000000000-0000000000000000)

Set DPS (verified): 41.9. Weights run: 3.5s. Verify run: 1.9s. 336 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.001, armor=0.084 ± 0.004, defense=0.240 ± 0.054 per rating point (1 rating = 1%, 0.240 per %), dodge=0.106 ± 0.013 per rating point (12 rating = 1%, 1.268 per %), strength=0.049 ± 0.000, agility=0.150 ± 0.008, attack_power=0.025 ± 0.000, hit=0.047 ± 0.009 per rating point (10 rating = 1%, 0.466 per %), crit=0.020 ± 0.001 per rating point (14 rating = 1%, 0.276 per %), expertise=1.556 ± 0.106

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.6 stamina points (419.72 DPS) | yes | Totemic Leather Helm (252456, -29.23 DPS) [crafted]; Trapper's Leather Helm (252513, -29.23 DPS) [crafted]; Defender's Leather Helm (252455, -35.50 DPS, sim-verified) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (252.60 DPS) | yes | River Pride Choker (13087, -68.99 DPS, sim-verified) [world_drop]; Senior Sergeant's Insignia (15200, -77.72 DPS) [vendor]; Darkspear Warding Pendant (272075, -97.15 DPS) [vendor] |
| shoulder | Watchman Pauldrons (7727) | Scarlet Monastery [dungeon] | 17.9 stamina points (347.95 DPS) | yes | Cloudy Gustwoven Spaulders (277043, -66.71 DPS) [crafted]; Azure Gustwoven Spaulders (277051, -66.71 DPS) [crafted]; Tanned Shoulderpads (270023, -73.32 DPS, sim-verified) [quest] |
| back | Sergeant's Cloak (18427) | PvP rank 7 · Sergeant · Horde [vendor] | 11.2 stamina points (217.56 DPS) | yes | Tigerstrike Mantle (13108, -17.12 DPS) [world_drop]; Enduring Cape (14763, -37.66 DPS) [world_drop]; Grimsteel Cape (4643, -37.94 DPS) [quest] |
| chest | Raptor Hide Harness (4455) | Leatherworking [crafted] | 21.3 stamina points (413.47 DPS) | yes | Spirewind Fetter (9406, -9.15 DPS) [dungeon]; Green Whelp Armor (7375, -32.30 DPS) [crafted]; Dokebi Chestguard (14581, -58.29 DPS) [world_drop] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 11.0 stamina points (213.05 DPS) | yes | Barbaric Bracers (18948, -3.78 DPS) [crafted]; Black Wolf Bracers (3230, -4.52 DPS) [dungeon]; Technician's Bracers (270042, -33.94 DPS) [quest] |
| hands | Braced Handguards (6784) | Centaur Bounty [quest] | 13.5 stamina points (261.54 DPS) | yes | Wolfclaw Gloves (1978, -22.47 DPS) [dungeon]; Brawler Gloves (720, -27.98 DPS, sim-verified) [world_drop]; Emblazoned Gloves (6397, -36.77 DPS) [world_drop] |
| waist | Warden's Leather Belt (252460) | Leatherworking [crafted] | 15.5 stamina points (301.60 DPS) | yes | Larcenist's Belt (279846, -45.26 DPS, sim-verified) [quest]; Skulker's Leather Belt (252520, -49.52 DPS) [crafted]; Stalker's Leather Belt (252521, -55.29 DPS) [crafted] |
| legs | Defender's Leather Kilt (252457) | Leatherworking [crafted] | 18.0 stamina points (349.81 DPS) | yes | Brawler's Leather Legguards (252516, -33.02 DPS) [crafted]; Trapper's Leather Legguards (252517, -39.74 DPS) [crafted]; Dark Ritual Leggings (270031, -47.67 DPS, sim-verified) [quest] |
| feet | Gnomebot Operating Boots (9450) | Gnomeregan: Crowd Pummeler 9-60 [dungeon] | 18.4 stamina points (357.93 DPS) | yes | Harbinger Boots (7754, -34.70 DPS, sim-verified) [dungeon]; Grizzled Boots (6335, -49.11 DPS) [quest]; Highlander's Mail Greaves (20123, -57.62 DPS) [vendor] |
| finger1 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 8.2 stamina points (159.77 DPS) | yes | Plains Ring (2039, -4.32 DPS) [dungeon]; Darkspear Signet (272071, -4.32 DPS) [vendor]; Defias Renegade Ring (1076, -48.00 DPS) [dungeon] |
| finger2 | Seal of Sylvanas (6414) | Arugal Must Die [quest] | 8.1 stamina points (158.33 DPS) | yes | Plains Ring (2039, -2.88 DPS) [dungeon]; Darkspear Signet (272071, -2.88 DPS) [vendor]; Defias Renegade Ring (1076, -46.56 DPS) [dungeon] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | sim-verified (3498.1 DPS) | yes | Rune of Duty (21568, -38.86 DPS) [rep] |
| trinket2 | Rune of Perfection (21566) | Warsong Outriders [rep] | sim-verified (3498.1 DPS) | yes | Rune of Duty (21568, +0.00 DPS) [rep] |
| main_hand | Staff of the Shade (2549) | Razorfen Kraul [dungeon] | sim-verified (3498.1 DPS) | yes | Spine of the Devourer (273040, -30.33 DPS) [world]; Advisor's Gnarled Staff (19569, -41.06 DPS) [pvp]; Manual Crowd Pummeler (9449, -261.16 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 30:** head: Brawler's Leather Helm; neck: Souvenir Sea Shell; shoulder: Watchman Pauldrons; back: Sergeant's Cloak; chest: Raptor Hide Harness; wrist: Cultist's Armguards; hands: Braced Handguards; waist: Warden's Leather Belt; legs: Defender's Leather Kilt; feet: Gnomebot Operating Boots; finger1: Insurgent's Band; finger2: Seal of Sylvanas; trinket1: Defiler's Talisman; trinket2: Rune of Perfection; main_hand: Staff of the Shade

No-known-source sample (15 of 336, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 16315 Sergeant Major's Cape; 18440 Sergeant's Cape; 18442 Master Sergeant's Insignia

### Band 40 (tauren, 0000000000000000-55230332020132000000-0000000000000000)

Set DPS (verified): 72.3. Weights run: 4.1s. Verify run: 2.0s. 443 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.018, armor=0.107 ± 0.006, defense=0.328 ± 0.080 per rating point (1 rating = 1%, 0.328 per %), dodge=0.113 ± 0.018 per rating point (12 rating = 1%, 1.360 per %), strength=0.046 ± 0.000, agility=0.136 ± 0.011, attack_power=0.023 ± 0.000, hit=not significant (0.055 ± 0.014) per rating point (10 rating = 1%, 0.550 per %), crit=0.029 ± 0.006 per rating point (14 rating = 1%, 0.407 per %), expertise=1.721 ± 0.160

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Adventurer's Pith Helmet (9420) | Uldaman [dungeon] | 25.3 stamina points (713.32 DPS) | yes | Brawler's Leather Helm (252512, -54.53 DPS) [crafted]; Warden's Wizard Hat (14604, -58.52 DPS, sim-verified) [world_drop]; Defender's Leather Helm (252455, -77.43 DPS) [crafted] |
| neck | Souvenir Sea Shell (274749) | Gezzy Gunkgear [vendor] | 13.0 stamina points (367.11 DPS) | yes | Gazlowe's Charm (13088, -53.85 DPS, sim-verified) [dungeon]; Darkspear Warding Pendant (272074, -56.48 DPS) [vendor]; River Pride Choker (13087, -107.76 DPS) [world_drop] |
| shoulder | Fleshhide Shoulders (10774) | Razorfen Downs: Glutton [dungeon] | 26.2 stamina points (739.52 DPS) | yes | Flintrock Shoulders (7755, -161.70 DPS, sim-verified) [dungeon]; Watchman Pauldrons (7727, -183.86 DPS) [dungeon]; Tanned Shoulderpads (270023, -265.34 DPS) [quest] |
| back | Wing of the Whelpling (13121) | World drop [world_drop] | 13.5 stamina points (382.18 DPS) | yes | Silky Spider Cape (10776, -9.37 DPS) [dungeon]; Well Oiled Cloak (12254, -12.38 DPS) [vendor]; Heraldic Cloak (8120, -22.21 DPS) [dungeon] |
| chest | Raptor Hunter Tunic (4119) | Raptor Mastery [quest] | 29.1 stamina points (821.22 DPS) | yes | Warden's Wraps (14601, -74.47 DPS) [world_drop]; Insignia Chestguard (4057, -143.66 DPS) [world_drop]; Robes of the Lich (10762, -205.42 DPS, sim-verified) [dungeon] |
| wrist | Scorpashi Wristbands (14654) | World drop [world_drop] | 13.4 stamina points (377.50 DPS) | yes | Forest Stalker's Bracers (19590, +0.00 DPS) [pvp]; Branded Leather Bracers (19508, -29.30 DPS) [dungeon]; Cultist's Armguards (270032, -40.71 DPS) [quest] |
| hands | Warden's Leather Gloves (252527) | Leatherworking [crafted] | 18.5 stamina points (523.16 DPS) | yes | Stalker's Leather Gloves (252526, +0.00 DPS, sim-verified) [crafted]; Bonefingers (10765, -24.87 DPS) [dungeon]; Razzeric's Racing Grips (6727, -27.22 DPS) [quest] |
| waist | Ogron's Sash (13117) | World drop [world_drop] | 18.2 stamina points (514.50 DPS) | yes | Warden's Leather Belt (252460, -39.56 DPS) [crafted]; Scorpashi Sash (14652, -68.17 DPS) [world_drop]; Kolkar Hunter's Belt (6788, -80.67 DPS, sim-verified) [quest] |
| legs | Imperial Leather Pants (4062) | Uldaman: Ancient Treasure [dungeon] | 23.3 stamina points (658.36 DPS) | yes | Warden's Woolies (14605, -9.04 DPS) [world_drop]; Basilisk Hide Pants (1718, -11.02 DPS) [world_drop]; Panther Hunter Leggings (4108, -19.88 DPS) [quest] |
| feet | Warden's Leather Shoes (252466) | Leatherworking [crafted] | 22.1 stamina points (623.97 DPS) | yes | Worn Running Boots (9398, -29.31 DPS) [dungeon]; Gnomebot Operating Boots (9450, -56.03 DPS) [dungeon]; Harbinger Boots (7754, -96.74 DPS) [dungeon] |
| finger1 | Underworld Band (1980) | World drop [world_drop] | 12.0 stamina points (338.87 DPS) | yes | Insurgent's Band (272066, -20.44 DPS) [vendor]; Darkspear Signet (272070, -28.24 DPS) [vendor]; Nogg's Gold Ring (9588, -56.48 DPS) [quest] |
| finger2 | Suspicious Spare Part (274754) | Rettrick [vendor] | 11.3 stamina points (319.73 DPS) | yes | Insurgent's Band (272066, -1.30 DPS) [vendor]; Darkspear Signet (272070, -9.10 DPS) [vendor]; Nogg's Gold Ring (9588, -37.34 DPS) [quest] |
| trinket1 | Rune of Perfection (21565) | Warsong Outriders [rep] | sim-verified (6771.2 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -28.24 DPS) [world]; Darkspear Voodoo Seal (272059, -28.24 DPS) [vendor] |
| trinket2 | Rune of Duty (21567) | Warsong Outriders [rep] | sim-verified (6771.2 DPS) | yes | Defiler's Talisman (21116, +0.00 DPS) [rep]; Arena Master (18706, -28.24 DPS) [world]; Darkspear Voodoo Seal (272059, -28.24 DPS) [vendor] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (6771.2 DPS) | yes | Ironshod Bludgeon (9408, -48.42 DPS) [dungeon]; Tok'kar's Murloc Basher (9678, -159.70 DPS) [quest]; Manual Crowd Pummeler (9449, -575.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 40:** head: Adventurer's Pith Helmet; shoulder: Fleshhide Shoulders; back: Wing of the Whelpling; chest: Raptor Hunter Tunic; wrist: Scorpashi Wristbands; hands: Warden's Leather Gloves; waist: Ogron's Sash; legs: Imperial Leather Pants; feet: Warden's Leather Shoes; finger1: Underworld Band; finger2: Suspicious Spare Part; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Cragwood Maul

No-known-source sample (15 of 443, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 50 (tauren, 0000000000000000-55230332020132012511-0000000000000000)

Set DPS (verified): 116.0. Weights run: 4.5s. Verify run: 2.4s. 578 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.032, armor=0.151 ± 0.012, defense=not significant (0.452 ± 0.148) per rating point (1 rating = 1%, 0.452 per %), dodge=0.255 ± 0.031 per rating point (12 rating = 1%, 3.057 per %), strength=0.057 ± 0.000, agility=0.252 ± 0.020, attack_power=0.029 ± 0.000, hit=not significant (0.096 ± 0.027) per rating point (10 rating = 1%, 0.959 per %), crit=not significant (0.035 ± 0.011) per rating point (14 rating = 1%, 0.488 per %), expertise=3.424 ± 0.295

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Sprightring Helm (17776) | Twisted Evils [quest] | 42.5 stamina points (1513.01 DPS) | yes | Blood Guard's Leather Headband (220851, -150.00 DPS) [vendor]; Impractical Headwarmer (274756, -321.65 DPS, sim-verified) [vendor]; Scorpashi Skullcap (14658, -347.09 DPS) [world_drop] |
| neck | Senior Sergeant's Insignia (18428) | PvP rank 8 · Senior Sergeant · Horde [vendor] | 14.0 stamina points (498.64 DPS) | yes | Darkspear Warding Pendant (272073, -35.62 DPS) [vendor]; Souvenir Sea Shell (274749, -35.62 DPS) [vendor]; Woven Ivy Necklace (19159, -49.51 DPS) [quest] |
| shoulder | Penance Spaulders (11963) | Ribbly Screwspigot [quest] | 31.2 stamina points (1110.17 DPS) | yes | Fleshhide Shoulders (10774, -0.01 DPS) [dungeon]; Blood Guard's Leather Shoulders (220853, -22.88 DPS) [vendor]; Warden's Leather Shoulders (252536, -76.88 DPS) [crafted] |
| back | Nightfall Drape (12465) | Sunken Temple: Morphaz [dungeon] | 19.7 stamina points (703.40 DPS) | yes | Sergeant's Cloak (16341, -10.78 DPS) [vendor]; Emberscale Cape (11311, -21.55 DPS) [dungeon]; Grovekeeper's Drape (17739, -76.62 DPS) [dungeon] |
| chest | Warbear Harness (15064) | Leatherworking [crafted] | 56.1 stamina points (1996.92 DPS) | yes | Heraldic Breastplate (8119, -549.00 DPS) [world_drop]; Jinxed Hoodoo Skin (9473, -552.37 DPS, sim-verified) [dungeon]; Stone Guard's Leather Armor (220855, -568.40 DPS) [vendor] |
| wrist | Arena Bracers (18710) | Arena Treasure Chest [world] | 25.2 stamina points (897.34 DPS) | yes | Forest Stalker's Bracers (19589, -73.05 DPS) [pvp]; First Sergeant's Dragonhide Armguards (18436, -130.14 DPS, sim-verified) [vendor]; Warden's Leather Bracers (252541, -163.86 DPS) [crafted] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 39.9 stamina points (1422.84 DPS) | yes | Warden's Leather Gauntlets (252549, -291.69 DPS) [crafted]; Raider Gloves (272100, -338.61 DPS, sim-verified) [vendor]; Feralheart Fists (226793, -352.64 DPS) [vendor] |
| waist | Warden's Leather Waistguard (252475) | Leatherworking [crafted] | 28.1 stamina points (1002.60 DPS) | yes | Skulker's Leather Waistguard (252474, -16.84 DPS) [crafted]; Prowler's Leather Waistguard (252473, -30.72 DPS) [crafted]; Pridelord Girdle (14674, -165.39 DPS) [world_drop] |
| legs | Stone Guard's Leather Pants (220859) | Lady Palanseer [vendor] | 37.2 stamina points (1326.14 DPS) | yes | Stone Guard's Crackling Leather Leggings (220865, -219.15 DPS, sim-verified) [vendor]; Stone Guard's Restored Leather Leggings (220883, -230.53 DPS) [vendor]; Scorpashi Leggings (14659, -242.63 DPS) [world_drop] |
| feet | Warden's Leather Boots (252470) | Leatherworking [crafted] | 29.4 stamina points (1045.78 DPS) | yes | Skulker's Leather Boots (252469, -25.82 DPS) [crafted]; Prowler's Leather Boots (252468, -39.70 DPS) [crafted]; Slitherscale Boots (10801, -47.81 DPS) [dungeon] |
| finger1 | Insurgent's Band (272065) | Creeg Bothunk [vendor] | 14.4 stamina points (513.88 DPS) | yes | Ring of Saviors (1447, -15.24 DPS) [world_drop]; Darkspear Signet (272069, -15.24 DPS) [vendor]; Underworld Band (1980, -86.48 DPS) [world_drop] |
| finger2 | Darkmoon Ring (19302) (or Darkspear Signet (272069), Ring of Saviors (1447)) | Lhara [vendor] | 14.0 stamina points (498.64 DPS) | yes | Ring of Saviors (1447, +0.00 DPS) [world_drop]; Darkspear Signet (272069, +0.00 DPS) [vendor]; Underworld Band (1980, -71.23 DPS) [world_drop] |
| trinket1 | Darkspear Voodoo Seal (272061) | Creeg Bothunk [vendor] | sim-verified (12157.7 DPS) | yes | Defiler's Talisman (21115, +0.00 DPS) [rep]; Guardian Talisman (1490, -35.62 DPS) [quest]; Rune of Perfection (21565, -71.23 DPS) [rep] |
| trinket2 | Relentless Raider's Seal (272060) | Creeg Bothunk [vendor] | sim-verified (12157.7 DPS) | yes | Defiler's Talisman (21115, +0.00 DPS) [rep]; Guardian Talisman (1490, -35.62 DPS) [quest]; Rune of Perfection (21565, -71.23 DPS) [rep] |
| main_hand | Cragwood Maul (11265) | Nothing But The Truth [quest] | sim-verified (12157.7 DPS) | yes | Advisor's Gnarled Staff (19567, -5.86 DPS) [pvp]; Radiant Staff (249453, -6.23 DPS) [crafted]; Manual Crowd Pummeler (9449, -764.82 DPS, sim-verified) [dungeon] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 50:** head: Sprightring Helm; neck: Senior Sergeant's Insignia; shoulder: Penance Spaulders; back: Nightfall Drape; chest: Warbear Harness; wrist: Arena Bracers; hands: Feralheart Grips; waist: Warden's Leather Waistguard; legs: Stone Guard's Leather Pants; feet: Warden's Leather Boots; finger1: Insurgent's Band; finger2: Darkmoon Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal

No-known-source sample (15 of 578, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60 (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 151.5. Weights run: 4.5s. Verify run: 5.9s. 1238 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.019, armor=0.185 ± 0.016, defense=0.985 ± 0.239 per rating point (1 rating = 1%, 0.985 per %), dodge=0.619 ± 0.048 per rating point (12 rating = 1%, 7.423 per %), strength=0.111 ± 0.000, agility=0.508 ± 0.030, attack_power=0.055 ± 0.000, hit=0.160 ± 0.039 per rating point (10 rating = 1%, 1.598 per %), crit=0.075 ± 0.016 per rating point (14 rating = 1%, 1.045 per %), expertise=5.806 ± 0.422

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 72.8 stamina points (4707.77 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Warbear Helm (252485, +0.00 DPS) [crafted] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 20.8 stamina points (1342.72 DPS) | yes | Scout's Medallion (19534, -203.00 DPS) [rep]; Senior Sergeant's Insignia (16335, -242.75 DPS) [vendor]; Medallion of Grand Marshal Morris (13091, -252.29 DPS) [world_drop] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 66.9 stamina points (4327.97 DPS) | yes | Warlord's Dragonhide Shoulders (231684, +0.00 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, -348.73 DPS) [vendor]; Defiler's Leather Shoulders (20194, -1183.25 DPS, sim-verified) [rep] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 32.9 stamina points (2129.56 DPS) | yes | Sergeant's Cloak (18461, -492.18 DPS) [vendor]; Cloak of Earth and Sky (279270, -497.17 DPS) [crafted]; Fel Cape (279269, -563.87 DPS, sim-verified) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | 86.8 stamina points (5618.16 DPS) | yes | Warlord's Dragonhide Tunic (231674, +0.00 DPS) [vendor]; Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Feralheart Armor (226804, -876.10 DPS, sim-verified) [vendor] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | 37.0 stamina points (2391.50 DPS) | yes | Forest Stalker's Bracers (19587, -143.97 DPS) [rep]; Bracers of Subterfuge (22668, -370.59 DPS, sim-verified) [quest]; First Sergeant's Dragonhide Armguards (18434, -371.95 DPS) [vendor] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 54.0 stamina points (3495.63 DPS) | yes | General's Dragonhide Gloves (231677, +0.00 DPS) [pvp]; General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, -318.05 DPS, sim-verified) [vendor] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 49.5 stamina points (3203.46 DPS) | yes | Hivethrasher's Girdle (275614, -504.10 DPS) [crafted]; Belt of Preserved Heads (20216, -505.58 DPS) [quest]; Shifter's Belt (272396, -711.13 DPS, sim-verified) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 91.9 stamina points (5945.92 DPS) | yes | General's Dragonhide Leggings (231685, -6.01 DPS) [pvp]; General's Dragonhide Legguards (231673, -359.98 DPS) [vendor]; Dire Warbear Woolies (227804, -367.99 DPS, sim-verified) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 65.1 stamina points (4212.67 DPS) | yes | General's Dragonhide Treads (231683, -31.62 DPS) [vendor]; Feralheart Treads (226803, -506.55 DPS) [vendor]; Drudge Boots (21532, -725.81 DPS, sim-verified) [quest] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 32.3 stamina points (2092.26 DPS) | yes | Ring of Awareness (272409, -814.17 DPS) [vendor]; Insurgent's Band (272064, -996.12 DPS) [vendor]; Cutthroat's Signet (272408, -1011.40 DPS) [vendor] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | 22.9 stamina points (1483.43 DPS) | yes | Ring of Awareness (272409, -361.14 DPS, sim-verified) [vendor]; Insurgent's Band (272064, -387.28 DPS) [vendor]; Cutthroat's Signet (272408, -402.57 DPS) [vendor] |
| trinket1 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-verified (22279.3 DPS) | yes | Defender's Grip Stabilizer (272440, -131.89 DPS) [vendor]; Relentless Raider's Seal (272060, -313.60 DPS) [vendor]; Darkspear Voodoo Seal (272061, -313.60 DPS) [vendor] |
| trinket2 | Defiler's Talisman (20072) | The Defilers [rep] | sim-verified (22279.3 DPS) | yes | Defender's Grip Stabilizer (272440, +0.00 DPS) [vendor]; Relentless Raider's Seal (272060, -129.41 DPS) [vendor]; Darkspear Voodoo Seal (272061, -129.41 DPS) [vendor] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-verified (22279.3 DPS) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Glacial Blade (19099, -1050.70 DPS, sim-verified) [rep] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Frostwolf Insignia Rank 6; trinket2: Defiler's Talisman; main_hand: The Unstoppable Force

No-known-source sample (15 of 1238, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

### Band 60, raid preset (tauren, 0000000000000000-55230332020132012551-0510000000000000)

Set DPS (verified): 336.8. Weights run: 4.7s. Verify run: 6.1s. 1238 eligible items had no known source.

Stat weights (normalized to stamina = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant; a rating-family stat's weight is per RATING point, matching the item tooltip, not per 1% hit/crit/dodge/parry/block/defense): stamina=1.000 ± 0.015, armor=0.207 ± 0.018, defense=1.532 ± 0.310 per rating point (1 rating = 1%, 1.532 per %), dodge=0.858 ± 0.059 per rating point (12 rating = 1%, 10.293 per %), strength=0.120 ± 0.000, agility=0.762 ± 0.038, attack_power=0.055 ± 0.000, hit=0.262 ± 0.044 per rating point (10 rating = 1%, 2.623 per %), crit=0.125 ± 0.019 per rating point (14 rating = 1%, 1.755 per %), expertise=7.919 ± 0.482

| Slot | Item | Source | Score (stamina points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Outlaw's Collar (279253) | Leatherworking [crafted] | 82.7 stamina points (7173.83 DPS) | yes | Warlord's Dragonhide Headdress (231675, +0.00 DPS) [vendor]; Warlord's Dragonhide Headguard (231687, +0.00 DPS) [vendor]; Feralheart Faceguard (226801, -577.21 DPS, sim-verified) [vendor] |
| neck | Amulet of the Darkmoon (19491) | 1200 Tickets - Amulet of the Darkmoon [quest] | 25.7 stamina points (2226.96 DPS) | yes | Medallion of Grand Marshal Morris (13091, -290.40 DPS) [world_drop]; Scout's Medallion (19534, -368.54 DPS) [rep]; Woven Ivy Necklace (19159, -702.34 DPS) [quest] |
| shoulder | Glowing Mantle of the Dawn (227818) | Argent Quartermaster Hasana [vendor] | 79.7 stamina points (6911.05 DPS) | yes | Warlord's Dragonhide Shoulders (231684, -259.01 DPS) [vendor]; Warlord's Dragonhide Spaulders (231681, -1197.77 DPS) [vendor]; Darkspear Pauldrons (272105, -1290.53 DPS, sim-verified) [vendor] |
| back | Shifting Cloak (18511) | Leatherworking [crafted] | 41.2 stamina points (3570.61 DPS) | yes | Fel Cape (279269, -693.44 DPS, sim-verified) [crafted]; Deathguard's Cloak (20068, -1228.35 DPS) [rep]; Cloak of Earth and Sky (279270, -1286.10 DPS) [crafted] |
| chest | Dire Warbear Harness (227803) | Meilosh [vendor] | 102.4 stamina points (8884.06 DPS) | yes | Warlord's Dragonhide Chestpiece (231686, +0.00 DPS) [vendor]; Warlord's Dragonhide Tunic (231674, -428.40 DPS) [vendor]; Feralheart Armor (226804, -1005.38 DPS, sim-verified) [vendor] |
| wrist | Feralheart Wristguards (226796) | Mokvar [vendor] | 41.7 stamina points (3617.04 DPS) | yes | Forest Stalker's Bracers (19587, -11.10 DPS) [rep]; Bracers of Subterfuge (22668, -28.44 DPS) [quest] |
| hands | Feralheart Grips (226802) | Mokvar [vendor] | 65.0 stamina points (5638.31 DPS) | yes | General's Dragonhide Grips (231688, +0.00 DPS) [vendor]; Raider Gloves (272099, -233.18 DPS) [vendor]; General's Dragonhide Gloves (231677, -651.08 DPS) [pvp] |
| waist | Feralheart Waistguard (226797) | Mokvar [vendor] | 58.0 stamina points (5030.22 DPS) | yes | Shifter's Belt (272396, -544.78 DPS, sim-verified) [vendor]; Belt of Preserved Heads (20216, -774.95 DPS) [quest]; Ferocity of the Timbermaw (227805, -864.95 DPS) [vendor] |
| legs | Sentinel's Leather Pants (237818) | Illiyana Moonblaze [vendor] | 104.6 stamina points (9075.48 DPS) | yes | General's Dragonhide Leggings (231685, -257.36 DPS) [pvp]; Dire Warbear Woolies (227804, -786.76 DPS, sim-verified) [vendor]; General's Dragonhide Legguards (231673, -1229.47 DPS) [vendor] |
| feet | Fine Dawn Treaders (227815) | Argent Quartermaster Hasana [vendor] | 78.5 stamina points (6809.75 DPS) | yes | General's Dragonhide Treads (231683, -630.52 DPS) [vendor]; Drudge Boots (21532, -888.50 DPS, sim-verified) [quest]; Feralheart Treads (226803, -1290.70 DPS) [vendor] |
| finger1 | Signet Ring of the Bronze Dragonflight (234035) | Anachronos [vendor] | 36.3 stamina points (3148.05 DPS) | yes | Ring of Awareness (272409, -1000.75 DPS) [vendor]; Cutthroat's Signet (272408, -1301.62 DPS) [vendor]; Legionnaire's Band (19510, -1602.23 DPS) [rep] |
| finger2 | Band of Resolution (22680) | Superior Armaments of Battle - Friend of the Dawn [quest] | 25.7 stamina points (2226.12 DPS) | yes | Ring of Awareness (272409, -78.82 DPS) [vendor]; Cutthroat's Signet (272408, -379.70 DPS) [vendor]; Legionnaire's Band (19510, -680.30 DPS) [rep] |
| trinket1 | Frostwolf Insignia Rank 6 (17909) | The Eye of Command [quest] | sim-decided (no score - a real sim tournament chose this pick) | yes | Smotts' Compass (4130, -892.87 DPS) [quest]; Arena Grand Master (19024, -892.87 DPS) [quest]; Defiler's Talisman (20072, -1092.91 DPS, sim-verified) [rep] |
| trinket2 | Defender's Grip Stabilizer (272440) | Pix Xizzix [vendor] | sim-verified (42330.3 DPS) | yes | Defiler's Talisman (20072, -748.72 DPS, sim-verified) [rep]; Smotts' Compass (4130, -835.25 DPS) [quest]; Arena Grand Master (19024, -835.25 DPS) [quest] |
| main_hand | The Unstoppable Force (19323) | Frostwolf Clan [rep] | sim-decided (no score - a real sim tournament chose this pick) | yes | High Warlord's Pulverizer (234545, +0.00 DPS) [pvp]; High Warlord's War Staff (234549, +0.00 DPS) [pvp]; Masterwork Stormhammer (12794, -1079.48 DPS, sim-verified) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | - | - |  |  |  |

**New at 60:** head: Outlaw's Collar; neck: Amulet of the Darkmoon; shoulder: Glowing Mantle of the Dawn; back: Shifting Cloak; chest: Dire Warbear Harness; wrist: Feralheart Wristguards; waist: Feralheart Waistguard; legs: Sentinel's Leather Pants; feet: Fine Dawn Treaders; finger1: Signet Ring of the Bronze Dragonflight; finger2: Band of Resolution; trinket1: Frostwolf Insignia Rank 6; trinket2: Defender's Grip Stabilizer; main_hand: The Unstoppable Force

No-known-source sample (15 of 1238, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 3738 Brewing Rod; 4196 Feathered Mantle; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 8708 Hammer of Expertise; 9362 Brilliant Gold Ring

