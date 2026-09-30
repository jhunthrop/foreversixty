# Leveling BiS: Combat

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 35.3. Weights run: 1.1s. Verify run: 1.2s. 284 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.404 ± 0.062, hit=not significant (0.000 ± 0.000), melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.54 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.14 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.29 DPS) [quest]; Tarnished Locket (279870, -0.29 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.34 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.02 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.1 | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.19 DPS) [crafted]; Dark Leather Tunic (2317, -0.24 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.07 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.02 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.33 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Pyrewood Signet Ring (277210, -0.10 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon]; The 1 Ring (8350, -0.24 DPS) [world] |
| finger2 | Protector's Band (20439) (or Pyrewood Signet Ring (277210)) | Silverwing Sentinels [rep] | 4.1 | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The 1 Ring (8350, -0.14 DPS) [world]; Pyrewood Signet Ring (277210, -0.21 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Barrens Basher (274744, -1.21 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.85 DPS) [quest]; Pulsating Hydra Heart (5183, -10.85 DPS) [world]; Tear of Grief (5611, -10.85 DPS) [quest] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 267.4 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted]; Venomstrike (6469, -3.06 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Ranger Bow

No-known-source sample (15 of 284, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves

### Band 30 (night-elf, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.7. Weights run: 1.0s. Verify run: 1.6s. 598 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.485 ± 0.060, hit=2.379 ± 0.384, melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.37 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.38 DPS) [rep]; Erudite's Amulet (277204, -0.48 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Fenrus' Hide (6340, -0.19 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.19 DPS) [dungeon]; Cloak of Night (4447, -0.24 DPS, sim-verified) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.09 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.24 DPS) [quest]; Green Leather Armor (4255, -0.38 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.11 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.19 DPS) [world_drop]; Madwolf Bracers (897, -0.24 DPS) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Fletcher's Gloves (7348, -0.44 DPS) [crafted]; Wolfclaw Gloves (1978, -0.48 DPS) [dungeon]; Pilferer's Gloves (7358, -0.50 DPS, sim-verified) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.72 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.73 DPS, sim-verified) [dungeon]; Insignia Leggings (4054, -0.81 DPS) [world_drop] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.04 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.19 DPS) [vendor] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 13.5 | yes | Insurgent's Band (272067, -0.22 DPS) [vendor]; Monkey Ring (6748, -0.31 DPS) [quest]; Ring of Precision (1491, -0.36 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Insurgent's Band (272067, -0.29 DPS, sim-verified) [vendor] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.63 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.77 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 0.0 | yes | Grayson's Torch (1172, -15.28 DPS) [quest]; Rod of Molten Fire (2565, -15.28 DPS) [world_drop]; Eye of Paleth (2943, -15.28 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 384.9 | yes | Silver Star (3463, -0.25 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.15 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.22 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Pyrewood Signet Ring; finger2: Ironspine's Eye; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring; 9395 Gloves of Old

### Band 40 (night-elf, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 88.2. Weights run: 1.1s. Verify run: 1.5s. 858 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.925 ± 0.148, hit=not significant (3.384 ± 0.980), melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.3 | yes | White Bandit Mask (10008, +0.00 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.05 DPS) [world]; Brawler's Leather Helm (252512, -0.10 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.29 DPS) [rep]; Sentinel's Medallion (20444, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 | yes | Forest Tracker Epaulets (2278, -0.57 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted]; Yeti Fur Cloak (2805, -0.19 DPS) [quest] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Nightscape Tunic (8175, +0.00 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.08 DPS) [crafted]; Hawkeye's Tunic (14592, -0.18 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.81 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.0 | yes | Fletcher's Gloves (7348, -0.99 DPS) [crafted]; Shadowskin Gloves (18238, -0.99 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.12 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.49 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Petrolspill Leggings (9509, -0.58 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.57 DPS, sim-verified) [quest] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 17.6 | yes | Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Ironspine's Eye (7686, -0.42 DPS) [dungeon]; Protector's Band (19515, -0.47 DPS) [rep] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -1.35 DPS) [vendor]; Sword of Serenity (6829, -1.99 DPS) [quest]; Hand of Righteousness (7721, -2.12 DPS) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Grayson's Torch (1172, -21.93 DPS) [quest]; Rod of Molten Fire (2565, -21.93 DPS) [world_drop]; Eye of Paleth (2943, -21.93 DPS) [quest] |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | World drop [world_drop] | 406.5 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Bow (17686, -0.48 DPS) [quest]; Master Hunter's Rifle (17687, -0.55 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger2: Insurgent's Band; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 858, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 50 (night-elf, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 108.8. Weights run: 1.1s. Verify run: 1.4s. 1108 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=6.578 ± 0.398, hit=not significant (5.130 ± 1.364), melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 159.4 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.60 DPS) [dungeon]; Helm of Fire (8348, -7.50 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19540, -0.09 DPS) [rep]; Sentinel's Medallion (19541, -0.27 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 151.4 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -6.79 DPS) [vendor]; Forest Tracker Epaulets (2278, -7.43 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 12.4 | yes | Nightscape Cloak (8195, -0.11 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.12 DPS) [dungeon]; Wolfmaster Cape (6314, -0.13 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 161.4 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -7.25 DPS) [quest]; Warbear Harness (15064, -7.55 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 112.1 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.07 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 112.1 | yes | Highlander's Lizardhide Girdle (20103, -1.07 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.57 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -4.39 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 0.0 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.41 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -7.24 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.24 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.24 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 71.3 | yes | Masons Fraternity Ring (9533, -2.97 DPS) [quest]; Insurgent's Band (272065, -3.01 DPS) [vendor]; Insurgent's Band (272066, -3.17 DPS) [vendor] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 25.0 | yes | Masons Fraternity Ring (9533, +0.00 DPS, sim-verified) [quest]; Insurgent's Band (272065, -0.54 DPS) [vendor]; Insurgent's Band (272066, -0.70 DPS) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, -2.47 DPS) [quest]; Guardian Talisman (1490, -2.47 DPS) [quest]; Ankh of Life (1713, -2.47 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 0.0 | yes | Might of Hakkar (10838, -3.10 DPS) [world]; Thorium Cestus (250614, -3.15 DPS) [crafted]; Julie's Dagger (6660, -3.94 DPS) [world_drop] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Claw of Celebras (17738, -3.80 DPS) [dungeon]; Thermotastic Egg Timer (9644, -29.39 DPS) [quest]; Grayson's Torch (1172, -29.57 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 0.0 | yes | Dark Iron Rifle (16004, -0.47 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -3.54 DPS) [world_drop]; Houndmaster's Bow (11628, -4.19 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Pyrewood Signet Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1108, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 263.5. Weights run: 1.2s. Verify run: 1.7s. 1658 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 316.1 | yes | Bloodvine Lens (19998, -4.60 DPS) [crafted]; Mask of the Unforgiven (13404, -6.05 DPS) [dungeon]; Ragefury Eyepatch (11735, -9.80 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 0.0 | yes | Medallion of the Dawn (22659, -1.59 DPS) [quest]; Beads of Ogre Might (22150, -5.40 DPS) [quest]; Choker of the Shifting Sands (21505, -6.79 DPS) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 189.8 | yes | Champion's Leather Shoulders (23258, -0.47 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -0.47 DPS) [vendor]; Champion's Leather Shoulders (227056, -0.47 DPS) [vendor] |
| back | Cloak of Veiled Shadows (21406) | Cloak of Veiled Shadows [quest] | 0.0 | yes | Earthweave Cloak (21187, -0.22 DPS) [quest]; Cloak of the Honor Guard (20073, -1.51 DPS) [rep]; Chromatic Cloak (18509, -2.69 DPS, sim-verified) [crafted] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 354.2 | yes | Zandalar Madcap's Tunic (19834, -6.53 DPS, sim-verified) [quest]; Stormshroud Armor (15056, -6.64 DPS) [crafted]; Deathdealer's Vest (21364, -7.68 DPS) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 151.4 | yes | Primal Batskin Bracers (19687, -2.87 DPS, sim-verified) [crafted]; Rockfury Bracers (21186, -5.75 DPS) [quest]; Marshal's Leather Armsplints (16460, -6.69 DPS) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 225.1 | yes | Devilsaur Gauntlets (15063, -4.39 DPS) [crafted]; Marshal's Leather Handgrips (16454, -4.39 DPS) [vendor]; Stormshroud Gloves (21278, -7.80 DPS, sim-verified) [crafted] |
| waist | Bonescythe Waistguard (22482) | Bonescythe Waistguard [quest] | 0.0 | yes | Highlander's Leather Girdle (20115, -0.72 DPS) [rep]; Belt of the Archmage (18405, -1.79 DPS) [crafted]; Highlander's Leather Girdle (20045, -3.28 DPS, sim-verified) [rep] |
| legs | Bonescythe Legplates (22477) | Bonescythe Legplates [quest] | 0.0 | yes | Marshal's Leather Leggings (16456, +0.00 DPS) [vendor]; General's Leather Legguards (16564, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [vendor] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 223.1 | yes | Deathdealer's Boots (21359, -3.07 DPS, sim-verified) [quest]; Bloodvine Boots (19684, -9.59 DPS) [crafted]; Shadowcraft Boots (16711, -9.67 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 175.1 | yes | Band of the Penitent (13217, -3.21 DPS) [quest]; Dragonslayer's Signet (18403, -3.21 DPS) [quest]; Ring of Entropy (18543, -3.21 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 159.1 | yes | Dragonslayer's Signet (18403, -2.35 DPS) [quest]; Ring of Entropy (18543, -2.35 DPS) [world]; Band of the Penitent (13217, -3.06 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, -2.12 DPS) [quest]; Guardian Talisman (1490, -2.12 DPS) [quest]; Ankh of Life (1713, -2.12 DPS) [world_drop] |
| trinket2 | Onyxia Blood Talisman (18406) | Celebrating Good Times [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 | yes | High Warlord's Blade (234552, +0.00 DPS) [vendor]; High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 815.9 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; High Warlord's Street Sweeper (234561, +0.00 DPS) [vendor] |

**New at 60:** head: Bonescythe Helmet; neck: Onyxia Tooth Pendant; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Bonescythe Waistguard; legs: Bonescythe Legplates; feet: Bonescythe Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket2: Onyxia Blood Talisman; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1658, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 00000000000000000-32510000000000000-0000000000000000000)

Set DPS (verified): 34.8. Weights run: 1.1s. Verify run: 1.3s. 291 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.404 ± 0.062, hit=not significant (0.000 ± 0.000), melee_haste=not significant (0.665 ± 0.825)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.38 DPS) [crafted]; Lucky Fishing Hat (19972, -0.38 DPS) [quest]; Flying Tiger Goggles (4368, -0.51 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.13 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.29 DPS) [quest]; Tarnished Locket (279870, -0.29 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.24 DPS) [crafted]; Forest Leather Mantle (4709, -0.24 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.32 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Heckler's Hide (286536, -0.10 DPS) [world]; Trapper's Leather Armor (252491, -0.32 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor]; Spare Part Bindings (279875, -0.10 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.02 DPS) [crafted]; Forest Leather Gloves (3058, -0.10 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.61 DPS) [quest]; Deviate Scale Belt (6468, -0.64 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.66 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.32 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Pyrewood Signet Ring (277210, -0.10 DPS) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.19 DPS) [dungeon] |
| finger2 | Legionnaire's Band (20429) (or Pyrewood Signet Ring (277210)) | Warsong Outriders [rep] | 4.1 | yes | Bounty Hunter's Ring (5351, -0.05 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.25 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.96 DPS) [world_drop]; Diamond Hammer (2194, -1.05 DPS) [world_drop]; Wingblade (6504, -1.15 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.85 DPS) [quest]; Nightglow Concoction (3451, -10.85 DPS) [quest]; Pulsating Hydra Heart (5183, -10.85 DPS) [world] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 267.4 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted]; Venomstrike (6469, -3.06 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Ranger Bow

No-known-source sample (15 of 291, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (troll, 00000000000000000-32531300000400000-0000000000000000000)

Set DPS (verified): 46.4. Weights run: 1.0s. Verify run: 1.6s. 609 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.005 ± 0.002, crit=0.485 ± 0.060, hit=2.379 ± 0.384, melee_haste=not significant (0.814 ± 0.770)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Tribal Worg Helm (6204, -0.13 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.38 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.38 DPS) [rep]; Erudite's Amulet (277204, -0.48 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Cloak of Night (4447, -0.19 DPS) [world]; Fenrus' Hide (6340, -0.19 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 | yes | Panther Armor (6670, -0.20 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.29 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.29 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.13 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.19 DPS) [world_drop]; Madwolf Bracers (897, -0.24 DPS) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Braced Handguards (6784, -0.43 DPS) [quest]; Fletcher's Gloves (7348, -0.44 DPS) [crafted]; Pilferer's Gloves (7358, -0.51 DPS, sim-verified) [crafted] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.29 DPS) [dungeon]; Deftkin Belt (16659, -0.38 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -0.62 DPS) [crafted]; Petrolspill Leggings (9509, -0.76 DPS, sim-verified) [dungeon]; Insignia Leggings (4054, -0.81 DPS) [world_drop] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.06 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.19 DPS) [world_drop]; Warsong Boots (16977, -0.19 DPS) [quest] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 13.5 | yes | Insurgent's Band (272067, -0.22 DPS) [vendor]; Monkey Ring (6748, -0.31 DPS) [quest]; Ring of Precision (1491, -0.36 DPS) [dungeon] |
| finger2 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Insurgent's Band (272067, -0.26 DPS, sim-verified) [vendor] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.63 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.77 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 0.0 | yes | Grayson's Torch (1172, -15.28 DPS) [quest]; Rod of Molten Fire (2565, -15.28 DPS) [world_drop]; Nightglow Concoction (3451, -15.28 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 384.9 | yes | Silver Star (3463, -0.25 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.15 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.22 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Pyrewood Signet Ring; finger2: Ironspine's Eye; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 609, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 87.2. Weights run: 1.1s. Verify run: 1.5s. 868 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.025 ± 0.009, crit=0.925 ± 0.148, hit=not significant (3.384 ± 0.980), melee_haste=not significant (-0.099 ± 1.840)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.3 | yes | White Bandit Mask (10008, +0.00 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.05 DPS) [world]; Spirit Hunter Headdress (6720, -0.10 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.29 DPS) [rep]; Scout's Medallion (20442, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.3 | yes | Forest Tracker Epaulets (2278, -0.56 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.60 DPS) [crafted]; Mantle of Thieves (2264, -0.65 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.09 DPS) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 15.4 | yes | Dusky Leather Armor (7374, -0.10 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.15 DPS) [world]; Panther Armor (6670, -0.31 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.80 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 33.0 | yes | Fletcher's Gloves (7348, -0.99 DPS) [crafted]; Shadowskin Gloves (18238, -0.99 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.42 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.48 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.60 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.37 DPS) [quest]; Petrolspill Leggings (9509, -0.58 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.3 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.56 DPS, sim-verified) [quest] |
| finger1 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 17.6 | yes | Ring of the Underwood (2951, -0.37 DPS) [world_drop]; Ironspine's Eye (7686, -0.42 DPS) [dungeon]; Legionnaire's Band (19512, -0.47 DPS) [rep] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -1.35 DPS) [vendor]; Hand of Righteousness (7721, -2.12 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272093, -2.13 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Grayson's Torch (1172, -21.93 DPS) [quest]; Rod of Molten Fire (2565, -21.93 DPS) [world_drop]; Nightglow Concoction (3451, -21.93 DPS) [quest] |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | World drop [world_drop] | 406.5 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Bow (17686, -0.48 DPS) [quest]; Master Hunter's Rifle (17687, -0.55 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger2: Insurgent's Band; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 32500000000000000-32531300000515201-0000000000000000000)

Set DPS (verified): 110.1. Weights run: 1.1s. Verify run: 1.3s. 1118 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.123 ± 0.020, crit=6.578 ± 0.398, hit=not significant (5.130 ± 1.364), melee_haste=not significant (-2.593 ± 2.550)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 159.4 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -3.60 DPS) [dungeon]; Helm of Fire (8348, -7.50 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19535, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19536, -0.09 DPS) [rep]; Woven Ivy Necklace (19159, -0.21 DPS) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 151.4 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -6.79 DPS) [vendor]; Forest Tracker Epaulets (2278, -7.43 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 12.4 | yes | Nightscape Cloak (8195, -0.11 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.12 DPS) [dungeon]; Wolfmaster Cape (6314, -0.13 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 161.4 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -7.25 DPS) [quest]; Warbear Harness (15064, -7.55 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.41 DPS) [crafted]; Pridelord Bands (14672, -0.47 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 112.1 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.47 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.07 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 112.1 | yes | Defiler's Lizardhide Girdle (20174, -1.07 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.57 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -4.39 DPS) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 184.2 | yes | Knight's Leather Pants (220858, +0.00 DPS, sim-verified) [vendor]; Stone Guard's Leather Pants (220859, -1.22 DPS) [vendor]; Ferine Leggings (6690, -8.45 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.5 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.24 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.24 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 71.3 | yes | White Bone Band (11862, -2.53 DPS) [quest]; Masons Fraternity Ring (9533, -2.97 DPS) [quest]; Insurgent's Band (272065, -3.01 DPS) [vendor] |
| finger2 | Pyrewood Signet Ring (277210) | The Horn of Xelthos [quest] | 25.0 | yes | White Bone Band (11862, +0.00 DPS, sim-verified) [quest]; Masons Fraternity Ring (9533, -0.50 DPS) [quest]; Insurgent's Band (272065, -0.54 DPS) [vendor] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Guardian Talisman (1490, -4.16 DPS) [quest]; Ankh of Life (1713, -4.16 DPS) [world_drop]; Blazing Emblem (2802, -4.16 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Guardian Talisman (1490, -2.47 DPS) [quest]; Ankh of Life (1713, -2.47 DPS) [world_drop]; Blazing Emblem (2802, -2.47 DPS) [world_drop] |
| main_hand | Inventor's Focal Sword (17719) | Maraudon: Tinkerer Gizlock [dungeon] | 0.0 | yes | Might of Hakkar (10838, -3.10 DPS) [world]; Thorium Cestus (250614, -3.15 DPS) [crafted]; Julie's Dagger (6660, -3.94 DPS) [world_drop] |
| off_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 553.3 | yes | Claw of Celebras (17738, -3.80 DPS) [dungeon]; White Bone Shredder (11863, -5.90 DPS) [quest]; Thermotastic Egg Timer (9644, -29.39 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 0.0 | yes | Dark Iron Rifle (16004, -0.47 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -3.54 DPS) [world_drop]; Houndmaster's Bow (11628, -4.19 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Defiler's Leather Girdle; legs: Stormshroud Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Pyrewood Signet Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Inventor's Focal Sword; off_hand: Hammer of the Northern Wind; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 32531000000000000-32531300000515201-5100000000000000000)

Set DPS (verified): 266.5. Weights run: 1.2s. Verify run: 1.7s. 1667 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.397 ± 0.131, crit=8.223 ± 0.549, hit=not significant (4.397 ± 2.027), melee_haste=not significant (4.644 ± 4.093)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 316.1 | yes | Bloodvine Lens (19998, -4.60 DPS) [crafted]; Mask of the Unforgiven (13404, -6.05 DPS) [dungeon]; Ragefury Eyepatch (11735, -9.05 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | For All To See [quest] | 0.0 | yes | Medallion of the Dawn (22659, -1.59 DPS) [quest]; Beads of Ogre Might (22150, -5.40 DPS) [quest]; Choker of the Shifting Sands (21505, -6.79 DPS) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 189.8 | yes | Champion's Leather Shoulders (23258, -0.47 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -0.47 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.47 DPS) [vendor] |
| back | Cloak of Veiled Shadows (21406) | Cloak of Veiled Shadows [quest] | 0.0 | yes | Earthweave Cloak (21187, -0.22 DPS) [quest]; Deathguard's Cloak (20068, -1.51 DPS) [rep]; Chromatic Cloak (18509, -3.16 DPS, sim-verified) [crafted] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 354.2 | yes | Zandalar Madcap's Tunic (19834, -5.69 DPS, sim-verified) [quest]; Stormshroud Armor (15056, -6.64 DPS) [crafted]; Deathdealer's Vest (21364, -7.68 DPS) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 151.4 | yes | Primal Batskin Bracers (19687, -3.02 DPS, sim-verified) [crafted]; Rockfury Bracers (21186, -5.75 DPS) [quest]; Marshal's Leather Armsplints (16460, -6.69 DPS) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 225.1 | yes | Devilsaur Gauntlets (15063, -4.39 DPS) [crafted]; Marshal's Leather Handgrips (16454, -4.39 DPS) [vendor]; Stormshroud Gloves (21278, -6.25 DPS, sim-verified) [crafted] |
| waist | Bonescythe Waistguard (22482) | Bonescythe Waistguard [quest] | 0.0 | yes | Defiler's Leather Girdle (20193, -0.72 DPS) [rep]; Belt of the Archmage (18405, -1.79 DPS) [crafted]; Defiler's Leather Girdle (20190, -3.41 DPS, sim-verified) [rep] |
| legs | Bonescythe Legplates (22477) | Bonescythe Legplates [quest] | 0.0 | yes | Marshal's Leather Leggings (16456, +0.00 DPS) [vendor]; General's Leather Legguards (16564, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [vendor] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 223.1 | yes | Deathdealer's Boots (21359, -3.09 DPS, sim-verified) [quest]; Bloodvine Boots (19684, -9.59 DPS) [crafted]; Shadowcraft Boots (16711, -9.67 DPS) [dungeon] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 175.1 | yes | Band of the Penitent (13217, -3.21 DPS) [quest]; Dragonslayer's Signet (18403, -3.21 DPS) [quest]; Ring of Entropy (18543, -3.21 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 159.1 | yes | Band of the Penitent (13217, -2.24 DPS, sim-verified) [quest]; Dragonslayer's Signet (18403, -2.35 DPS) [quest]; Ring of Entropy (18543, -2.35 DPS) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Guardian Talisman (1490, -3.90 DPS) [quest]; Ankh of Life (1713, -3.90 DPS) [world_drop]; Blazing Emblem (2802, -3.90 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Guardian Talisman (1490, -2.12 DPS) [quest]; Ankh of Life (1713, -2.12 DPS) [world_drop]; Blazing Emblem (2802, -2.12 DPS) [world_drop] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 | yes | High Warlord's Blade (234552, +0.00 DPS) [vendor]; High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 815.9 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; High Warlord's Street Sweeper (234561, +0.00 DPS) [vendor] |

**New at 60:** head: Bonescythe Helmet; neck: Onyxia Tooth Pendant; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Bonescythe Waistguard; legs: Bonescythe Legplates; feet: Bonescythe Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: The Purifier

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

