# Leveling BiS: Assassination

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 35.6. Weights run: 1.1s. Verify run: 1.2s. 284 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.067 ± 0.020, crit=2.221 ± 0.111, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 | yes | Shadow Goggles (4373, -0.40 DPS) [crafted]; Lucky Fishing Hat (19972, -0.40 DPS) [quest]; Flying Tiger Goggles (4368, -0.66 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.4 | yes | Erudite's Amulet (277204, -0.16 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.30 DPS) [quest]; Tarnished Locket (279870, -0.30 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 | yes | Reinforced Woolen Shoulders (4315, -0.25 DPS) [crafted]; Forest Leather Mantle (4709, -0.25 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.41 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.10 DPS, sim-verified) [quest]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.7 | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.20 DPS) [crafted]; Dark Leather Tunic (2317, -0.25 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.08 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 31.1 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -1.15 DPS) [dungeon]; Forest Leather Gloves (3058, -1.25 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.58 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.64 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.6 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.5 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.15 DPS) [crafted]; Blackened Defias Boots (10402, -0.36 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 | yes | Pyrewood Signet Ring (277210, -0.10 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon]; The 1 Ring (8350, -0.25 DPS) [world] |
| finger2 | Protector's Band (20439) (or Pyrewood Signet Ring (277210)) | Silverwing Sentinels [rep] | 4.3 | yes | Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; The 1 Ring (8350, -0.15 DPS) [world]; Pyrewood Signet Ring (277210, -0.25 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Barrens Basher (274744, -1.20 DPS) [vendor] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.71 DPS) [quest]; Pulsating Hydra Heart (5183, -10.71 DPS) [world]; Tear of Grief (5611, -10.71 DPS) [quest] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 270.9 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted]; Venomstrike (6469, -3.06 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Ranger Bow

No-known-source sample (15 of 284, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves

### Band 30 (night-elf, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 43.4. Weights run: 1.1s. Verify run: 1.8s. 598 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.016, crit=3.306 ± 0.128, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Humbert's Helm (4724, -0.16 DPS) [world]; Tribal Worg Helm (6204, -0.16 DPS, sim-verified) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.22 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.35 DPS) [rep]; Erudite's Amulet (277204, -0.45 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 | yes | Dark Leather Shoulders (4252, -0.21 DPS) [crafted]; Insignia Mantle (4721, -0.21 DPS) [world_drop]; Mantle of Thieves (2264, -0.39 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Cloak of Night (4447, -0.14 DPS, sim-verified) [world]; Fenrus' Hide (6340, -0.16 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.16 DPS) [dungeon] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, +0.00 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.19 DPS) [quest]; Green Leather Armor (4255, -0.34 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.02 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.16 DPS) [world_drop]; Madwolf Bracers (897, -0.21 DPS) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 46.3 | yes | Heavy Earthen Gloves (7359, +0.00 DPS, sim-verified) [crafted]; Pilferer's Gloves (7358, -1.77 DPS) [crafted]; Wolfclaw Gloves (1978, -1.87 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.67 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Petrolspill Leggings (9509, -0.46 DPS, sim-verified) [dungeon]; Dusky Leather Leggings (7373, -0.55 DPS) [crafted]; Insignia Leggings (4054, -0.76 DPS) [world_drop] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.15 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.15 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon]; Protector's Band (19517, -0.16 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.00 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon]; Protector's Band (19517, -0.11 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.62 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.75 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 0.0 | yes | Grayson's Torch (1172, -15.01 DPS) [quest]; Rod of Molten Fire (2565, -15.01 DPS) [world_drop]; Eye of Paleth (2943, -15.01 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 391.5 | yes | Silver Star (3463, -0.16 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.13 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.21 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring; 9395 Gloves of Old

### Band 40 (night-elf, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 84.8. Weights run: 1.2s. Verify run: 1.7s. 858 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.105 ± 0.014, crit=2.769 ± 0.073, hit=not significant (0.000 ± 0.000), melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 13.3 | yes | White Bandit Mask (10008, +0.00 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.04 DPS) [world]; Brawler's Leather Helm (252512, -0.07 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.17 DPS) [rep]; Sentinel's Medallion (20444, -0.25 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 | yes | Nightscape Shoulders (8192, -0.40 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.49 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.04 DPS) [crafted]; Yeti Fur Cloak (2805, -0.11 DPS) [quest] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 16.6 | yes | Dusky Leather Armor (7374, -0.04 DPS) [crafted]; Hawkeye's Tunic (14592, -0.11 DPS) [world]; Raptorbane Armor (3566, -0.44 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.37 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted]; Cultist's Armguards (270032, -0.69 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 58.8 | yes | Shadowskin Gloves (18238, -0.67 DPS) [crafted]; Fletcher's Gloves (7348, -1.39 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -1.43 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.20 DPS) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon]; Highlander's Chain Girdle (20090, -0.42 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.20 DPS) [quest]; Petrolspill Leggings (9509, -0.35 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.4 | yes | Imperial Leather Boots (6431, +0.00 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.07 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.07 DPS) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.07 DPS) [dungeon]; Insurgent's Band (272067, -0.10 DPS) [vendor]; Protector's Band (19515, -0.11 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 11.0 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Protector's Band (19515, -0.07 DPS) [rep]; Disengagement Ring (276202, -0.07 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -0.89 DPS) [vendor]; Sword of Serenity (6829, -1.34 DPS) [quest]; Hand of Righteousness (7721, -1.43 DPS) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | 442.0 | yes | Grayson's Torch (1172, -14.80 DPS) [quest]; Rod of Molten Fire (2565, -14.80 DPS) [world_drop]; Eye of Paleth (2943, -14.80 DPS) [quest] |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | World drop [world_drop] | 602.2 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Bow (17686, -0.57 DPS) [quest]; Master Hunter's Rifle (17687, -0.60 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 858, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 50 (night-elf, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 120.9. Weights run: 1.2s. Verify run: 1.6s. 1108 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.107 ± 0.016, crit=3.272 ± 0.082, hit=not significant (0.000 ± 0.000), melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 61.8 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.54 DPS) [dungeon]; Helm of Fire (8348, -1.45 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19539, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19540, -0.06 DPS) [rep]; Sentinel's Medallion (19541, -0.17 DPS) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 53.8 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -1.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -1.40 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 12.2 | yes | Nightscape Cloak (8195, -0.06 DPS, sim-verified) [crafted]; Wolfmaster Cape (6314, -0.07 DPS) [dungeon]; Pridelord Cape (14673, -0.07 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 63.8 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -1.29 DPS) [quest]; Warbear Harness (15064, -1.48 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.26 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.8 | yes | First Sergeant's Leather Gauntlets (220857, -0.20 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.24 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -0.67 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 65.8 | yes | Highlander's Lizardhide Girdle (20103, -0.67 DPS) [rep]; Highlander's Cloth Girdle (20097, -0.79 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -1.21 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 0.0 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -0.98 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -1.28 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.14 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.14 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Insurgent's Band (272065, -0.17 DPS) [vendor]; Insurgent's Band (272066, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.30 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 15.5 | yes | Insurgent's Band (272066, -0.12 DPS) [vendor]; Insurgent's Band (272065, -0.14 DPS, sim-verified) [vendor]; Ring of the Underwood (2951, -0.15 DPS) [world_drop] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 0.0 | yes | Inventor's Focal Sword (17719, -0.48 DPS) [dungeon]; Julie's Dagger (6660, -1.41 DPS) [world_drop]; Lifeforce Dirk (10750, -1.69 DPS) [quest] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 0.0 | yes | Claw of Celebras (17738, -1.49 DPS) [dungeon]; Thermotastic Egg Timer (9644, -17.65 DPS) [quest]; Grayson's Torch (1172, -17.76 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 0.0 | yes | Dark Iron Rifle (16004, -0.24 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -3.86 DPS) [world_drop]; Houndmaster's Bow (11628, -4.43 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1108, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 218.4. Weights run: 1.2s. Verify run: 2.0s. 1658 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=not significant (0.000 ± 0.000), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 149.4 | yes | Bloodvine Lens (19998, -1.12 DPS) [crafted]; Champion's Leather Helm (23257, -1.85 DPS) [vendor]; Ragefury Eyepatch (11735, -8.43 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 0.0 | yes | Onyxia Tooth Pendant (18404, +0.00 DPS) [quest]; Choker of the Shifting Sands (21505, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -2.70 DPS, sim-verified) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 82.6 | yes | Champion's Leather Shoulders (23258, -0.10 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -0.10 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.10 DPS) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 57.7 | yes | Cloak of the Honor Guard (20073, +0.00 DPS, sim-verified) [rep]; Cape of the Black Baron (13340, -0.69 DPS) [dungeon]; Cloak of the Fallen God (21710, -0.94 DPS) [quest] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 195.5 | yes | Stormshroud Armor (15056, -2.65 DPS) [crafted]; Deathdealer's Vest (21364, -3.18 DPS) [quest]; Zandalar Madcap's Tunic (19834, -6.26 DPS, sim-verified) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 87.1 | yes | Forest Stalker's Bracers (19587, -2.05 DPS, sim-verified) [rep]; Marshal's Leather Armsplints (16460, -2.18 DPS) [pvp]; General's Leather Armsplints (16559, -2.18 DPS) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 123.7 | yes | Marshal's Leather Handgrips (16454, -1.44 DPS) [vendor]; General's Leather Mitts (16560, -1.44 DPS) [vendor]; Devilsaur Gauntlets (15063, -6.41 DPS, sim-verified) [crafted] |
| waist | Bonescythe Waistguard (22482) | Bonescythe Waistguard [quest] | 0.0 | yes | Highlander's Leather Girdle (20115, -0.24 DPS) [rep]; Belt of the Archmage (18405, -0.90 DPS) [crafted]; Highlander's Leather Girdle (20045, -2.47 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 115.5 | yes | Devilsaur Leggings (15062, +0.00 DPS, sim-verified) [crafted]; Knight-Captain's Leather Legguards (16419, +0.00 DPS) [pvp]; Legionnaire's Leather Leggings (16508, +0.00 DPS) [pvp] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 121.7 | yes | Deathdealer's Boots (21359, -3.06 DPS) [quest]; Blood Guard's Leather Walkers (22856, -3.11 DPS) [vendor]; Highlander's Leather Boots (20052, -7.19 DPS, sim-verified) [rep] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 73.7 | yes | Dragonslayer's Signet (18403, -0.53 DPS) [quest]; Ring of Entropy (18543, -0.53 DPS) [world]; Mindtear Band (20632, -0.53 DPS) [world] |
| finger2 | Band of the Penitent (13217) (or Dragonslayer's Signet (18403), Ring of Entropy (18543), Mindtear Band (20632), Band of Earthen Wrath (21179), Band of Earthen Might (21182), Don Rodrigo's Band (21563), Ritssyn's Ring of Chaos (21836), Ring of the Eternal Flame (23237)) | Houses of the Holy [quest] | 57.7 | yes | Dragonslayer's Signet (18403, +0.00 DPS, sim-verified) [quest]; Ring of Entropy (18543, +0.00 DPS) [world]; Mindtear Band (20632, +0.00 DPS) [world] |
| trinket1 | Onyxia Blood Talisman (18406) | Celebrating Good Times [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| trinket2 | Talisman of Arathor (20071) | The League of Arathor [rep] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 | yes | High Warlord's Blade (234552, +0.00 DPS) [vendor]; High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor] |
| off_hand | Ravencrest's Legacy (21520) | Treasure of the Timeless One [quest] | 813.5 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor] |
| ranged | Core Marksman Rifle (18282) | Engineering [crafted] | 0.0 | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; High Warlord's Street Sweeper (234561, +0.00 DPS) [vendor] |

**New at 60:** head: Bonescythe Helmet; neck: Blazefury Medallion; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Bonescythe Waistguard; legs: Stormshroud Pants; feet: Bonescythe Sabatons; finger1: Don Julio's Band; finger2: Band of the Penitent; trinket1: Onyxia Blood Talisman; trinket2: Talisman of Arathor; main_hand: Shadowsong's Sorrow; off_hand: Ravencrest's Legacy; ranged: Core Marksman Rifle

No-known-source sample (15 of 1658, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 32500000100000000-00000000000000000-0000000000000000000)

Set DPS (verified): 34.9. Weights run: 1.1s. Verify run: 1.3s. 291 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.067 ± 0.020, crit=2.221 ± 0.111, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.171 ± 0.745)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.5 | yes | Shadow Goggles (4373, -0.40 DPS) [crafted]; Lucky Fishing Hat (19972, -0.40 DPS) [quest]; Flying Tiger Goggles (4368, -0.64 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.4 | yes | Erudite's Amulet (277204, -0.15 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.30 DPS) [quest]; Tarnished Locket (279870, -0.30 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.3 | yes | Reinforced Woolen Shoulders (4315, -0.25 DPS) [crafted]; Forest Leather Mantle (4709, -0.25 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.39 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.4 | yes | Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Catacomb Cloak (279899, -0.09 DPS, sim-verified) [quest]; Sentry Cloak (2059, -0.10 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.5 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Heckler's Hide (286536, -0.10 DPS) [world]; Trapper's Leather Armor (252491, -0.33 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.3 | yes | Wolf Bracers (4794, -0.08 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.10 DPS) [vendor]; Spare Part Bindings (279875, -0.10 DPS) [quest] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 31.1 | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -1.15 DPS) [dungeon]; Forest Leather Gloves (3058, -1.25 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Dusty Belt (279897, -0.59 DPS) [quest]; Deviate Scale Belt (6468, -0.60 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.64 DPS) [world] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.6 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.5 | yes | Footpads of the Fang (10411, -0.10 DPS) [dungeon]; Dark Leather Boots (2315, -0.15 DPS) [crafted]; Blackened Defias Boots (10402, -0.35 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.4 | yes | Pyrewood Signet Ring (277210, -0.10 DPS) [quest]; Bounty Hunter's Ring (5351, -0.15 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.20 DPS) [dungeon] |
| finger2 | Legionnaire's Band (20429) (or Pyrewood Signet Ring (277210)) | Warsong Outriders [rep] | 4.3 | yes | Bounty Hunter's Ring (5351, -0.05 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.10 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.26 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Shadowfang (1482) | Shadowfang Keep: Son of Arugal [dungeon] | 248.9 | yes | Blackfang (2236, -0.95 DPS) [world_drop]; Diamond Hammer (2194, -1.03 DPS) [world_drop]; Wingblade (6504, -1.12 DPS) [quest] |
| off_hand | Cruel Barb (5191) | Westfall: Edwin VanCleef [dungeon] | 229.6 | yes | Grayson's Torch (1172, -10.71 DPS) [quest]; Nightglow Concoction (3451, -10.71 DPS) [quest]; Pulsating Hydra Heart (5183, -10.71 DPS) [world] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 270.9 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted]; Venomstrike (6469, -3.06 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Shadowfang; off_hand: Cruel Barb; ranged: Ranger Bow

No-known-source sample (15 of 291, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (troll, 32500000551000000-00000000000000000-0000000000000000000)

Set DPS (verified): 43.0. Weights run: 1.1s. Verify run: 1.7s. 609 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.097 ± 0.016, crit=3.306 ± 0.128, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.596 ± 0.802)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 11.0 | yes | Brawler's Leather Hood (252504, -0.10 DPS) [crafted]; Humbert's Helm (4724, -0.16 DPS) [world]; Tribal Worg Helm (6204, -0.17 DPS, sim-verified) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.21 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.35 DPS) [rep]; Erudite's Amulet (277204, -0.45 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 12.1 | yes | Dark Leather Shoulders (4252, -0.21 DPS) [crafted]; Insignia Mantle (4721, -0.21 DPS) [world_drop]; Mantle of Thieves (2264, -0.38 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Cloak of Night (4447, -0.16 DPS) [world]; Fenrus' Hide (6340, -0.16 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 15.4 | yes | Panther Armor (6670, -0.28 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.31 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.31 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.00 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.16 DPS) [world_drop]; Madwolf Bracers (897, -0.21 DPS) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 46.3 | yes | Heavy Earthen Gloves (7359, +0.00 DPS, sim-verified) [crafted]; Pilferer's Gloves (7358, -1.77 DPS) [crafted]; Braced Handguards (6784, -1.82 DPS) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.36 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Petrolspill Leggings (9509, -0.44 DPS, sim-verified) [dungeon]; Dusky Leather Leggings (7373, -0.55 DPS) [crafted]; Insignia Leggings (4054, -0.76 DPS) [world_drop] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, +0.00 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.15 DPS) [world_drop]; Warsong Boots (16977, -0.15 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.9 | yes | Monkey Ring (6748, -0.10 DPS) [quest]; Ring of Precision (1491, -0.16 DPS) [dungeon]; Legionnaire's Band (19513, -0.16 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, +0.00 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.11 DPS) [dungeon]; Legionnaire's Band (19513, -0.11 DPS) [rep] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Ironspine's Fist (7687, +0.00 DPS, sim-verified) [dungeon]; Electrocutioner Leg (9446, -0.62 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272094, -0.75 DPS) [vendor] |
| off_hand | Bloody Brass Knuckles (7683) | Scarlet Monastery: Interrogator Vishas [dungeon] | 0.0 | yes | Grayson's Torch (1172, -15.01 DPS) [quest]; Rod of Molten Fire (2565, -15.01 DPS) [world_drop]; Nightglow Concoction (3451, -15.01 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 391.5 | yes | Silver Star (3463, -0.14 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.13 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.21 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Bloody Brass Knuckles; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 609, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (troll, 32500000551501040-00000000000000000-0000000000000000000)

Set DPS (verified): 82.5. Weights run: 1.2s. Verify run: 1.7s. 868 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.105 ± 0.014, crit=2.769 ± 0.073, hit=not significant (0.000 ± 0.000), melee_haste=2.095 ± 0.041

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 13.3 | yes | White Bandit Mask (10008, +0.00 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.04 DPS) [world]; Spirit Hunter Headdress (6720, -0.07 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.17 DPS) [rep]; Scout's Medallion (20442, -0.25 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 24.2 | yes | Nightscape Shoulders (8192, -0.40 DPS) [crafted]; Mantle of Thieves (2264, -0.44 DPS) [dungeon]; Forest Tracker Epaulets (2278, -0.48 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.04 DPS) [world_drop]; Parachute Cloak (10518, -0.04 DPS) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 16.6 | yes | Dusky Leather Armor (7374, -0.11 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.11 DPS) [world]; Panther Armor (6670, -0.22 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.37 DPS) [world_drop]; Dusky Bracers (7378, -0.37 DPS) [crafted]; Cultist's Armguards (270032, -0.68 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 58.8 | yes | Shadowskin Gloves (18238, -0.67 DPS) [crafted]; Fletcher's Gloves (7348, -1.37 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -1.43 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.20 DPS) [rep]; Blackened Defias Belt (10403, -0.40 DPS) [dungeon]; Defiler's Chain Girdle (20152, -0.41 DPS, sim-verified) [rep] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.20 DPS) [quest]; Petrolspill Leggings (9509, -0.35 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 14.4 | yes | Imperial Leather Boots (6431, +0.00 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.07 DPS) [crafted]; Skulker's Leather Shoes (252531, -0.07 DPS) [crafted] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.07 DPS) [dungeon]; Insurgent's Band (272067, -0.10 DPS) [vendor]; Legionnaire's Band (19512, -0.11 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 11.0 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Legionnaire's Band (19512, -0.07 DPS) [rep]; Disengagement Ring (276202, -0.07 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Ardent Custodian (868) | World drop [world_drop] | 460.0 | yes | Jhordy's Misplaced Screwdriver (274753, -0.89 DPS) [vendor]; Hand of Righteousness (7721, -1.43 DPS) [dungeon]; Darkspear Skirmisher's Bludgeon (272093, -1.44 DPS) [vendor] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | 442.0 | yes | Grayson's Torch (1172, -14.80 DPS) [quest]; Rod of Molten Fire (2565, -14.80 DPS) [world_drop]; Nightglow Concoction (3451, -14.80 DPS) [quest] |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | World drop [world_drop] | 602.2 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Bow (17686, -0.57 DPS) [quest]; Master Hunter's Rifle (17687, -0.60 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Ardent Custodian; off_hand: Vanquisher's Sword; ranged: Sniper Rifle

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 32500000551501051-32300000000000000-0000000000000000000)

Set DPS (verified): 123.7. Weights run: 1.2s. Verify run: 1.6s. 1118 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.107 ± 0.016, crit=3.272 ± 0.082, hit=not significant (0.000 ± 0.000), melee_haste=2.543 ± 0.050

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 61.8 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.54 DPS) [dungeon]; Helm of Fire (8348, -1.45 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19535, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19536, -0.06 DPS) [rep]; Woven Ivy Necklace (19159, -0.14 DPS) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 53.8 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -1.00 DPS) [vendor]; Forest Tracker Epaulets (2278, -1.40 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | Maraudon: Princess Theradras [dungeon] | 12.2 | yes | Nightscape Cloak (8195, -0.05 DPS, sim-verified) [crafted]; Wolfmaster Cape (6314, -0.07 DPS) [dungeon]; Battlehard Cape (11858, -0.07 DPS) [quest] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 63.8 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -1.29 DPS) [quest]; Warbear Harness (15064, -1.48 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world]; Wicked Leather Bracers (15084, -0.26 DPS) [crafted]; Pridelord Bands (14672, -0.30 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 65.8 | yes | First Sergeant's Leather Gauntlets (220857, -0.20 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.24 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -0.67 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 65.8 | yes | Defiler's Lizardhide Girdle (20174, -0.67 DPS) [rep]; Defiler's Cloth Girdle (20165, -0.79 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -1.21 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 0.0 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -0.90 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -1.28 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 22.1 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.14 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.14 DPS) [vendor] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -0.29 DPS) [quest]; Insurgent's Band (272065, -0.30 DPS) [vendor]; Insurgent's Band (272066, -0.40 DPS) [vendor] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Insurgent's Band (272065, -0.17 DPS) [vendor]; Insurgent's Band (272066, -0.27 DPS) [vendor]; Masons Fraternity Ring (9533, -0.78 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Guardian Talisman (1490, -1.42 DPS) [quest]; Ankh of Life (1713, -1.42 DPS) [world_drop]; Blazing Emblem (2802, -1.42 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| main_hand | Hammer of the Northern Wind (810) | World drop [world_drop] | 0.0 | yes | Inventor's Focal Sword (17719, -0.48 DPS) [dungeon]; Julie's Dagger (6660, -1.41 DPS) [world_drop]; Lifeforce Dirk (10750, -1.69 DPS) [quest] |
| off_hand | Thorium Cestus (250614) | Blacksmithing [crafted] | 0.0 | yes | Claw of Celebras (17738, -1.49 DPS) [dungeon]; White Bone Shredder (11863, -2.82 DPS) [quest]; Thermotastic Egg Timer (9644, -17.65 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 0.0 | yes | Dark Iron Rifle (16004, -0.24 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -3.86 DPS) [world_drop]; Houndmaster's Bow (11628, -4.43 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Defiler's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Hammer of the Northern Wind; off_hand: Thorium Cestus; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1118, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 32500000551501051-32520000000000000-5100000000000000000)

Set DPS (verified): 222.9. Weights run: 1.2s. Verify run: 1.9s. 1667 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.002, agility=1.131 ± 0.020, crit=4.124 ± 0.104, hit=not significant (0.000 ± 0.000), melee_haste=3.276 ± 0.064

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 149.4 | yes | Bloodvine Lens (19998, -1.12 DPS) [crafted]; Champion's Leather Helm (23257, -1.85 DPS) [vendor]; Ragefury Eyepatch (11735, -8.29 DPS, sim-verified) [dungeon] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | Onyxia Tooth Pendant (18404, -0.53 DPS) [quest]; Choker of the Shifting Sands (21505, -1.32 DPS) [quest]; Dragonheart Necklace (20622, -1.91 DPS) [world] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 82.6 | yes | Champion's Leather Shoulders (23258, -0.10 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -0.10 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.10 DPS) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 57.7 | yes | Deathguard's Cloak (20068, +0.00 DPS, sim-verified) [rep]; Cape of the Black Baron (13340, -0.69 DPS) [dungeon]; Cloak of the Fallen God (21710, -0.94 DPS) [quest] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 195.5 | yes | Stormshroud Armor (15056, -2.65 DPS) [crafted]; Deathdealer's Vest (21364, -3.18 DPS) [quest]; Zandalar Madcap's Tunic (19834, -5.95 DPS, sim-verified) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 87.1 | yes | Forest Stalker's Bracers (19587, -2.09 DPS, sim-verified) [rep]; Marshal's Leather Armsplints (16460, -2.18 DPS) [pvp]; General's Leather Armsplints (16559, -2.18 DPS) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 123.7 | yes | Marshal's Leather Handgrips (16454, -1.44 DPS) [vendor]; General's Leather Mitts (16560, -1.44 DPS) [vendor]; Devilsaur Gauntlets (15063, -6.10 DPS, sim-verified) [crafted] |
| waist | Bonescythe Waistguard (22482) | Bonescythe Waistguard [quest] | 0.0 | yes | Defiler's Leather Girdle (20193, -0.24 DPS) [rep]; Belt of the Archmage (18405, -0.90 DPS) [crafted]; Defiler's Leather Girdle (20190, -2.67 DPS, sim-verified) [rep] |
| legs | Stormshroud Pants (15057) | Leatherworking [crafted] | 115.5 | yes | Devilsaur Leggings (15062, +0.00 DPS, sim-verified) [crafted]; Knight-Captain's Leather Legguards (16419, +0.00 DPS) [pvp]; Legionnaire's Leather Leggings (16508, +0.00 DPS) [pvp] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 121.7 | yes | Deathdealer's Boots (21359, -3.06 DPS) [quest]; Blood Guard's Leather Walkers (22856, -3.11 DPS) [vendor]; Defiler's Leather Boots (20186, -6.98 DPS, sim-verified) [rep] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 73.7 | yes | Dragonslayer's Signet (18403, -0.53 DPS) [quest]; Ring of Entropy (18543, -0.53 DPS) [world]; Mindtear Band (20632, -0.53 DPS) [world] |
| finger2 | Band of the Penitent (13217) (or Dragonslayer's Signet (18403), Ring of Entropy (18543), Mindtear Band (20632), Band of Earthen Wrath (21179), Band of Earthen Might (21182), Don Rodrigo's Band (21563), Ritssyn's Ring of Chaos (21836), Ring of the Eternal Flame (23237)) | Houses of the Holy [quest] | 57.7 | yes | Dragonslayer's Signet (18403, +0.00 DPS, sim-verified) [quest]; Ring of Entropy (18543, +0.00 DPS) [world]; Mindtear Band (20632, +0.00 DPS) [world] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Guardian Talisman (1490, -1.39 DPS) [quest]; Ankh of Life (1713, -1.39 DPS) [world_drop]; Blazing Emblem (2802, -1.39 DPS) [world_drop] |
| trinket2 | Onyxia Blood Talisman (18406) | For All To See [quest] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 | yes | High Warlord's Blade (234552, +0.00 DPS) [vendor]; High Warlord's Quickblade (234553, +0.00 DPS) [vendor]; Grand Marshal's Swiftblade (234579, +0.00 DPS) [vendor] |
| off_hand | Dagger of Veiled Shadows (21404) | Dagger of Veiled Shadows [quest] | 0.0 | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor] |
| ranged | Core Marksman Rifle (18282) | Engineering [crafted] | 0.0 | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; High Warlord's Street Sweeper (234561, +0.00 DPS) [vendor] |

**New at 60:** head: Bonescythe Helmet; neck: Medallion of the Dawn; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Bonescythe Waistguard; legs: Stormshroud Pants; feet: Bonescythe Sabatons; finger1: Don Julio's Band; finger2: Band of the Penitent; trinket2: Onyxia Blood Talisman; main_hand: Shadowsong's Sorrow; off_hand: Dagger of Veiled Shadows; ranged: Core Marksman Rifle

No-known-source sample (15 of 1667, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

