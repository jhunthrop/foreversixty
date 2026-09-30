# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.7. Weights run: 1.2s. Verify run: 1.4s. 284 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.290 ± 0.038, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.37 DPS) [crafted]; Lucky Fishing Hat (19972, -0.37 DPS) [quest]; Flying Tiger Goggles (4368, -0.46 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.12 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.28 DPS) [quest]; Tarnished Locket (279870, -0.28 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Forest Leather Mantle (4709, -0.23 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.28 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.01 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.09 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.2 | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.18 DPS) [crafted]; Dark Leather Tunic (2317, -0.23 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.09 DPS) [crafted]; Forest Leather Gloves (3058, -0.09 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.54 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.63 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.09 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.28 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Pyrewood Signet Ring (277210, -0.09 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; The 1 Ring (8350, -0.23 DPS) [world] |
| finger2 | Protector's Band (20439) (or Pyrewood Signet Ring (277210)) | Silverwing Sentinels [rep] | 4.1 | yes | Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; The 1 Ring (8350, -0.14 DPS) [world]; Pyrewood Signet Ring (277210, -0.20 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Silverwing Sentinels [rep] | 0.0 | yes | - |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 | yes | Assassin's Blade (1935, -0.32 DPS) [world_drop]; Evocator's Blade (2567, -0.65 DPS) [world_drop]; Buzzer Blade (2169, -1.48 DPS) [dungeon] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | 222.7 | yes | Grayson's Torch (1172, -10.09 DPS) [quest]; Pulsating Hydra Heart (5183, -10.09 DPS) [world]; Tear of Grief (5611, -10.09 DPS) [quest] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 279.0 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted]; Venomstrike (6469, -3.06 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Blackfang; off_hand: Edward's Knife; ranged: Ranger Bow

No-known-source sample (15 of 284, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak; 9771 Greenweave Gloves

### Band 30 (night-elf, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 38.8. Weights run: 1.3s. Verify run: 1.6s. 598 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.421 ± 0.047, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.1 | yes | Brawler's Leather Hood (252504, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.30 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.37 DPS) [rep]; Erudite's Amulet (277204, -0.46 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.31 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Fenrus' Hide (6340, -0.18 DPS) [dungeon]; Glowing Lizardscale Cloak (6449, -0.18 DPS) [dungeon]; Cloak of Night (4447, -0.20 DPS, sim-verified) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.07 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.23 DPS) [quest]; Green Leather Armor (4255, -0.37 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.09 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.18 DPS) [world_drop]; Madwolf Bracers (897, -0.23 DPS) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Pilferer's Gloves (7358, -0.40 DPS, sim-verified) [crafted]; Wolfclaw Gloves (1978, -0.46 DPS) [world_drop]; Toughened Leather Gloves (4253, -0.46 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.69 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Petrolspill Leggings (9509, -0.59 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.60 DPS) [crafted]; Insignia Leggings (4054, -0.79 DPS) [world_drop] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.04 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.1 | yes | Monkey Ring (6748, -0.09 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [world_drop]; Protector's Band (19517, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.09 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [world_drop]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | Talisman of Arathor (21119) | The League of Arathor [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Darkspear Insurgent's Spellblade (272086, +0.00 DPS, sim-verified) [vendor]; Thornspike (6681, -1.57 DPS) [dungeon]; Claw of the Shadowmancer (2912, -1.61 DPS) [world_drop] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | 0.0 | yes | Grayson's Torch (1172, -13.78 DPS) [quest]; Rod of Molten Fire (2565, -13.78 DPS) [world_drop]; Eye of Paleth (2943, -13.78 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 397.1 | yes | Silver Star (3463, -0.20 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.14 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.21 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 598, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring; 9395 Gloves of Old

### Band 40 (night-elf, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 75.0. Weights run: 1.4s. Verify run: 1.8s. 858 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.663 ± 0.081, hit=not significant (0.000 ± 0.000), melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.2 | yes | White Bandit Mask (10008, +0.00 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.05 DPS) [world]; Brawler's Leather Helm (252512, -0.10 DPS) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.29 DPS) [rep]; Sentinel's Medallion (20444, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -0.45 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted]; Yeti Fur Cloak (2805, -0.19 DPS) [quest] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Nightscape Tunic (8175, +0.00 DPS, sim-verified) [crafted]; Dusky Leather Armor (7374, -0.09 DPS) [crafted]; Hawkeye's Tunic (14592, -0.19 DPS) [world] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.64 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 29.3 | yes | Skulker's Leather Gloves (252525, -0.95 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.95 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.26 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.38 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest]; Petrolspill Leggings (9509, -0.59 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.2 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.54 DPS, sim-verified) [quest] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.14 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Protector's Band (19515, -0.19 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 10.1 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Protector's Band (19515, -0.10 DPS) [rep]; Disengagement Ring (276202, -0.10 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Silverwing Sentinels [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Talisman of Arathor (21118, +0.00 DPS) [rep] |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 432.8 | yes | Coldrage Dagger (10761, -0.86 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272085, -1.37 DPS) [vendor]; Hypnotic Blade (7714, -2.86 DPS) [dungeon] |
| off_hand | Black Menace (6831) (or Coldrage Dagger (10761)) | In the Name of the Light [quest] | 415.4 | yes | Coldrage Dagger (10761, +0.00 DPS, sim-verified) [dungeon]; Grayson's Torch (1172, -20.58 DPS) [quest]; Rod of Molten Fire (2565, -20.58 DPS) [world_drop] |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | World drop [world_drop] | 407.2 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Bow (17686, -0.49 DPS) [quest]; Master Hunter's Rifle (17687, -0.55 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Black Menace; ranged: Sniper Rifle

No-known-source sample (15 of 858, see the JSON for more): 1189 Overseer's Ring; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7470 Regal Wizard Hat

### Band 50 (night-elf, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 110.0. Weights run: 1.4s. Verify run: 2.0s. 1107 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.218 ± 0.061, crit=6.499 ± 0.352, hit=not significant (0.000 ± 0.000), melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 107.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.87 DPS) [dungeon]; Helm of Fire (8348, -4.67 DPS) [crafted] |
| neck | Sentinel's Medallion (19539) | Silverwing Sentinels [rep] | 14.6 | yes | Sentinel's Medallion (19540, -0.07 DPS) [rep]; Sentinel's Medallion (19541, -0.26 DPS) [rep]; Ghostshard Talisman (7731, -0.32 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 99.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -3.98 DPS) [vendor]; Forest Tracker Epaulets (2278, -4.63 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | World drop [world_drop] | 13.4 | yes | Nightscape Cloak (8195, -0.12 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.13 DPS) [world_drop]; Wolfmaster Cape (6314, -0.18 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 109.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -4.38 DPS) [quest]; Warbear Harness (15064, -4.71 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.42 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 111.0 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.44 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.08 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 111.0 | yes | Highlander's Lizardhide Girdle (20103, -1.08 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.47 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -4.38 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 0.0 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.17 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.49 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.4 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.34 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.34 DPS) [vendor] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Insurgent's Band (272065, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.42 DPS) [world_drop]; Insurgent's Band (272066, -0.43 DPS) [vendor] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 17.1 | yes | Ring of the Underwood (2951, -0.26 DPS) [world_drop]; Insurgent's Band (272066, -0.27 DPS) [vendor]; Insurgent's Band (272065, -0.47 DPS, sim-verified) [vendor] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| trinket2 | Smoking Heart of the Mountain (11811) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 | yes | Charstone Dirk (17710, -0.45 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272084, -1.17 DPS) [vendor]; Widowmaker (4091, -3.23 DPS) [world_drop] |
| off_hand | Lifeforce Dirk (10750) (or Charstone Dirk (17710)) | The God Hakkar [quest] | 503.2 | yes | Charstone Dirk (17710, -26.07 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -27.04 DPS) [quest]; Grayson's Torch (1172, -27.24 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 0.0 | yes | Dark Iron Rifle (16004, -0.44 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -3.46 DPS) [world_drop]; Houndmaster's Bow (11628, -4.18 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Sentinel's Medallion; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1107, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

### Band 60 (night-elf, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 217.3. Weights run: 1.4s. Verify run: 2.0s. 1643 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=8.118 ± 0.504, hit=not significant (5.733 ± 1.563), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 320.1 | yes | Bloodvine Goggles (19999, -4.88 DPS) [crafted]; Ragefury Eyepatch (11735, -4.93 DPS) [dungeon]; Mask of the Unforgiven (13404, -5.74 DPS, sim-verified) [dungeon] |
| neck | Onyxia Tooth Pendant (18404) | Celebrating Good Times [quest] | 0.0 | yes | Fury of the Forgotten Swarm (21809, -0.44 DPS) [world_drop]; Medallion of the Dawn (22659, -2.21 DPS) [quest]; Beads of Ogre Might (22150, -5.20 DPS) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 197.0 | yes | Champion's Leather Shoulders (23258, -0.21 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -0.21 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.21 DPS) [vendor] |
| back | Cloak of Veiled Shadows (21406) | Cloak of Veiled Shadows [quest] | 0.0 | yes | Earthweave Cloak (21187, -0.19 DPS) [quest]; Cloak of the Honor Guard (20073, -2.06 DPS) [rep]; Chromatic Cloak (18509, -3.49 DPS, sim-verified) [crafted] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 364.6 | yes | Zandalar Madcap's Tunic (19834, -5.40 DPS, sim-verified) [quest]; Stormshroud Armor (15056, -7.30 DPS) [crafted]; Deathdealer's Vest (21364, -7.97 DPS) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 144.4 | yes | Primal Batskin Bracers (19687, -1.30 DPS, sim-verified) [crafted]; Rockfury Bracers (21186, -4.63 DPS) [quest]; Marshal's Leather Armsplints (16460, -6.48 DPS) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 237.0 | yes | Stormshroud Gloves (21278, -4.69 DPS, sim-verified) [crafted]; Devilsaur Gauntlets (15063, -5.07 DPS) [crafted]; Marshal's Leather Handgrips (16454, -5.30 DPS) [vendor] |
| waist | Bonescythe Waistguard (22482) | Bonescythe Waistguard [quest] | 0.0 | yes | Highlander's Leather Girdle (20115, -0.44 DPS) [rep]; Belt of the Archmage (18405, -1.51 DPS) [crafted]; Highlander's Leather Girdle (20045, -2.99 DPS, sim-verified) [rep] |
| legs | Bonescythe Legplates (22477) | Bonescythe Legplates [quest] | 0.0 | yes | Marshal's Leather Leggings (16456, +0.00 DPS) [vendor]; General's Leather Legguards (16564, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [vendor] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 235.0 | yes | Deathdealer's Boots (21359, -2.51 DPS, sim-verified) [quest]; Fine Dawn Treaders (227815, -8.75 DPS) [vendor]; Darkmantle Footpads (226831, -8.85 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | 187.0 | yes | Band of the Penitent (13217, -3.90 DPS) [quest]; Dragonslayer's Signet (18403, -3.90 DPS) [quest]; Ring of Entropy (18543, -3.90 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 171.0 | yes | Band of the Penitent (13217, -2.60 DPS, sim-verified) [quest]; Dragonslayer's Signet (18403, -3.05 DPS) [quest]; Ring of Entropy (18543, -3.05 DPS) [world] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Thunderbrew's Boot Flask (744, -2.74 DPS) [quest]; Guardian Talisman (1490, -2.74 DPS) [quest]; Ankh of Life (1713, -2.74 DPS) [world_drop] |
| trinket2 | Onyxia Blood Talisman (18406) | Celebrating Good Times [quest] | 0.0 | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 | yes | Grand Marshal's Dirk (18838, +0.00 DPS) [vendor]; High Warlord's Razor (234556, +0.00 DPS) [vendor]; Grand Marshal's Dirk (234582, +0.00 DPS) [vendor] |
| off_hand | Dagger of Veiled Shadows (21404) | Dagger of Veiled Shadows [quest] | 798.4 | yes | Thermotastic Egg Timer (9644, -42.25 DPS) [quest]; Imperial Red Scepter (15930, -42.31 DPS) [world_drop]; Elunarian Sphere (15968, -42.31 DPS) [world_drop] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; High Warlord's Street Sweeper (234561, +0.00 DPS) [vendor] |

**New at 60:** head: Bonescythe Helmet; neck: Onyxia Tooth Pendant; shoulder: Bonescythe Pauldrons; back: Cloak of Veiled Shadows; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Bonescythe Waistguard; legs: Bonescythe Legplates; feet: Bonescythe Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket2: Onyxia Blood Talisman; main_hand: Shadowsong's Sorrow; off_hand: Dagger of Veiled Shadows; ranged: The Purifier

No-known-source sample (15 of 1643, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.5. Weights run: 1.2s. Verify run: 1.4s. 291 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.290 ± 0.038, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.37 DPS) [crafted]; Lucky Fishing Hat (19972, -0.37 DPS) [quest]; Flying Tiger Goggles (4368, -0.44 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.11 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.28 DPS) [quest]; Tarnished Locket (279870, -0.28 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Forest Leather Mantle (4709, -0.23 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.27 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.09 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Heckler's Hide (286536, -0.09 DPS) [world]; Trapper's Leather Armor (252491, -0.27 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor]; Spare Part Bindings (279875, -0.09 DPS) [quest] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.09 DPS) [crafted]; Forest Leather Gloves (3058, -0.09 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.54 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.63 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.09 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.27 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Pyrewood Signet Ring (277210, -0.09 DPS) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon] |
| finger2 | Legionnaire's Band (20429) (or Pyrewood Signet Ring (277210)) | Warsong Outriders [rep] | 4.1 | yes | Bounty Hunter's Ring (5351, -0.05 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.22 DPS, sim-verified) [quest] |
| trinket1 | Rune of Perfection (21566) | Warsong Outriders [rep] | 0.0 | yes | - |
| trinket2 | Rune of Duty (21568) | Warsong Outriders [rep] | 0.0 | yes | - |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 | yes | Assassin's Blade (1935, -0.32 DPS) [world_drop]; Evocator's Blade (2567, -0.65 DPS) [world_drop]; Buzzer Blade (2169, -1.48 DPS) [dungeon] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | 222.7 | yes | Grayson's Torch (1172, -10.09 DPS) [quest]; Nightglow Concoction (3451, -10.09 DPS) [quest]; Pulsating Hydra Heart (5183, -10.09 DPS) [world] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 279.0 | yes | Cracked Blacksmith Hammer (285279, +0.00 DPS, sim-verified) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted]; Venomstrike (6469, -3.06 DPS) [dungeon] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Blackfang; off_hand: Edward's Knife; ranged: Ranger Bow

No-known-source sample (15 of 291, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 7608 Seer's Fine Stein; 9747 Simple Britches; 9748 Simple Robe; 9749 Simple Blouse; 9756 Nomad Trousers; 9757 Nomad Tunic; 9766 Greenweave Sash; 9767 Greenweave Sandals; 9768 Greenweave Bracers; 9770 Greenweave Cloak

### Band 30 (troll, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 38.2. Weights run: 1.3s. Verify run: 1.6s. 609 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.421 ± 0.047, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.1 | yes | Brawler's Leather Hood (252504, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.10 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.31 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.37 DPS) [rep]; Erudite's Amulet (277204, -0.46 DPS) [quest] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.31 DPS, sim-verified) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Cloak of Night (4447, -0.18 DPS) [world]; Fenrus' Hide (6340, -0.18 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 | yes | Panther Armor (6670, -0.16 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.28 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.28 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.10 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.18 DPS) [world_drop]; Madwolf Bracers (897, -0.23 DPS) [world] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Pilferer's Gloves (7358, -0.41 DPS, sim-verified) [crafted]; Braced Handguards (6784, -0.42 DPS) [quest]; Wolfclaw Gloves (1978, -0.46 DPS) [world_drop] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.37 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Dusky Leather Leggings (7373, -0.60 DPS) [crafted]; Petrolspill Leggings (9509, -0.61 DPS, sim-verified) [world_drop]; Insignia Leggings (4054, -0.79 DPS) [world_drop] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.05 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Warsong Boots (16977, -0.18 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.1 | yes | Monkey Ring (6748, -0.09 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [world_drop]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [world_drop]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | Defiler's Talisman (21120) | The Defilers [rep] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS) [vendor] |
| trinket2 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | 0.0 | yes | Rune of Perfection (21566, +0.00 DPS) [rep]; Rune of Duty (21568, +0.00 DPS) [rep]; Relentless Raider's Seal (272062, +0.00 DPS, sim-verified) [vendor] |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Darkspear Insurgent's Spellblade (272086, +0.00 DPS, sim-verified) [vendor]; Thornspike (6681, -1.57 DPS) [dungeon]; Claw of the Shadowmancer (2912, -1.61 DPS) [world_drop] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | 0.0 | yes | Grayson's Torch (1172, -13.78 DPS) [quest]; Rod of Molten Fire (2565, -13.78 DPS) [world_drop]; Nightglow Concoction (3451, -13.78 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 397.1 | yes | Silver Star (3463, -0.20 DPS, sim-verified) [quest]; Nightstalker Bow (6696, -2.14 DPS) [dungeon]; Satchel of Bronze Bombs (285276, -3.21 DPS) [crafted] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 609, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7608 Seer's Fine Stein; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 8183 Precision Bow; 9362 Brilliant Gold Ring

### Band 40 (troll, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 74.4. Weights run: 1.4s. Verify run: 1.8s. 868 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.663 ± 0.081, hit=not significant (0.000 ± 0.000), melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Nightscape Headband (8176) | Leatherworking [crafted] | 12.2 | yes | White Bandit Mask (10008, +0.00 DPS, sim-verified) [crafted]; Hawkeye's Helm (14591, -0.05 DPS) [world]; Spirit Hunter Headdress (6720, -0.10 DPS) [quest] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.29 DPS) [rep]; Scout's Medallion (20442, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -0.45 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [world_drop] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.09 DPS) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted] |
| chest | Nightscape Tunic (8175) | Leatherworking [crafted] | 15.2 | yes | Dusky Leather Armor (7374, -0.10 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.15 DPS) [world]; Panther Armor (6670, -0.30 DPS) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.64 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 29.3 | yes | Skulker's Leather Gloves (252525, -0.95 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.95 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.27 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.38 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest]; Petrolspill Leggings (9509, -0.59 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.2 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.50 DPS, sim-verified) [quest] |
| finger1 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ironspine's Eye (7686, -0.14 DPS) [dungeon]; Insurgent's Band (272067, -0.15 DPS) [vendor]; Legionnaire's Band (19512, -0.19 DPS) [rep] |
| finger2 | Ring of the Underwood (2951) | World drop [world_drop] | 10.1 | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Legionnaire's Band (19512, -0.10 DPS) [rep]; Disengagement Ring (276202, -0.10 DPS) [vendor] |
| trinket1 | Ankh of Life (1713) | World drop [world_drop] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| trinket2 | Rune of Perfection (21565) | Warsong Outriders [rep] | 0.0 | yes | Blazing Emblem (2802, +0.00 DPS) [world_drop]; Cold Basilisk Eye (5079, +0.00 DPS) [world]; Defiler's Talisman (21116, +0.00 DPS) [rep] |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 432.8 | yes | Darkspear Insurgent's Spellblade (272085, -1.37 DPS) [vendor]; Hypnotic Blade (7714, -2.86 DPS) [dungeon]; Sacrificial Kris (3187, -3.61 DPS) [world_drop] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 415.4 | yes | Grayson's Torch (1172, -20.58 DPS) [quest]; Rod of Molten Fire (2565, -20.58 DPS) [world_drop]; Nightglow Concoction (3451, -20.58 DPS) [quest] |
| ranged | Sniper Rifle (3430) (or Mithril Heavy-bore Rifle (10510)) | World drop [world_drop] | 407.2 | yes | Mithril Heavy-bore Rifle (10510, +0.00 DPS, sim-verified) [crafted]; Master Hunter's Bow (17686, -0.49 DPS) [quest]; Master Hunter's Rifle (17687, -0.55 DPS) [quest] |

**New at 40:** head: Nightscape Headband; shoulder: Sunburn Spaulders; chest: Nightscape Tunic; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Insurgent's Band; finger2: Ring of the Underwood; trinket1: Ankh of Life; trinket2: Rune of Perfection; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Coldrage Dagger; ranged: Sniper Rifle

No-known-source sample (15 of 868, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 50 (troll, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 112.8. Weights run: 1.4s. Verify run: 2.1s. 1117 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.218 ± 0.061, crit=6.499 ± 0.352, hit=not significant (0.000 ± 0.000), melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 107.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.87 DPS) [dungeon]; Helm of Fire (8348, -4.67 DPS) [crafted] |
| neck | Scout's Medallion (19535) | Warsong Outriders [rep] | 14.6 | yes | Scout's Medallion (19536, -0.07 DPS) [rep]; Woven Ivy Necklace (19159, -0.20 DPS) [quest]; Ghostshard Talisman (7731, -0.35 DPS, sim-verified) [dungeon] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 99.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -3.98 DPS) [vendor]; Forest Tracker Epaulets (2278, -4.63 DPS) [world_drop] |
| back | Serpentskin Cloak (8259) | World drop [world_drop] | 13.4 | yes | Nightscape Cloak (8195, -0.11 DPS, sim-verified) [crafted]; Pridelord Cape (14673, -0.13 DPS) [world_drop]; Wolfmaster Cape (6314, -0.18 DPS) [dungeon] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 109.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -4.38 DPS) [quest]; Warbear Harness (15064, -4.71 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.42 DPS) [world_drop] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 111.0 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.44 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.08 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 111.0 | yes | Defiler's Lizardhide Girdle (20174, -1.08 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.47 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -4.38 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | 0.0 | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.62 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.49 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.4 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.34 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.34 DPS) [vendor] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -0.38 DPS) [quest]; Insurgent's Band (272065, -0.49 DPS) [vendor]; Ring of the Underwood (2951, -0.64 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Insurgent's Band (272065, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.42 DPS) [world_drop]; Masons Fraternity Ring (9533, -1.25 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Guardian Talisman (1490, -2.27 DPS) [quest]; Ankh of Life (1713, -2.27 DPS) [world_drop]; Blazing Emblem (2802, -2.27 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Guardian Talisman (1490, +0.00 DPS) [quest]; Ankh of Life (1713, +0.00 DPS) [world_drop]; Blazing Emblem (2802, +0.00 DPS) [world_drop] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 | yes | Charstone Dirk (17710, -0.45 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272084, -1.17 DPS) [vendor]; Widowmaker (4091, -3.23 DPS) [world_drop] |
| off_hand | Lifeforce Dirk (10750) (or Charstone Dirk (17710)) | The God Hakkar [quest] | 503.2 | yes | Charstone Dirk (17710, -26.87 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -27.04 DPS) [quest]; Grayson's Torch (1172, -27.24 DPS) [quest] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | 0.0 | yes | Dark Iron Rifle (16004, -0.44 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -3.46 DPS) [world_drop]; Houndmaster's Bow (11628, -4.18 DPS) [dungeon] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Scout's Medallion; shoulder: Knight-Lieutenant's Leather Shoulders; back: Serpentskin Cloak; chest: Knight's Leather Armor; waist: Defiler's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 1117, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

### Band 60 (troll, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 217.7. Weights run: 1.4s. Verify run: 2.1s. 1652 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=8.118 ± 0.504, hit=not significant (5.733 ± 1.563), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Bonescythe Helmet (22478) | Bonescythe Helmet [quest] | 320.1 | yes | Bloodvine Goggles (19999, -4.88 DPS) [crafted]; Ragefury Eyepatch (11735, -4.93 DPS) [dungeon]; Mask of the Unforgiven (13404, -7.35 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | 0.0 | yes | Fury of the Forgotten Swarm (21809, +0.00 DPS) [world_drop]; Medallion of the Dawn (22659, +0.00 DPS) [quest]; Onyxia Tooth Pendant (18404, -1.14 DPS, sim-verified) [quest] |
| shoulder | Bonescythe Pauldrons (22479) | Bonescythe Pauldrons [quest] | 197.0 | yes | Champion's Leather Shoulders (23258, -0.21 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, -0.21 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (227054, -0.21 DPS) [vendor] |
| back | Chromatic Cloak (18509) | Leatherworking [crafted] | 113.7 | yes | Cloak of Veiled Shadows (21406, +0.00 DPS, sim-verified) [quest]; Earthweave Cloak (21187, -2.05 DPS) [quest]; Deathguard's Cloak (20068, -3.92 DPS) [rep] |
| chest | Bonescythe Breastplate (22476) | Bonescythe Breastplate [quest] | 364.6 | yes | Zandalar Madcap's Tunic (19834, -6.01 DPS, sim-verified) [quest]; Stormshroud Armor (15056, -7.30 DPS) [crafted]; Deathdealer's Vest (21364, -7.97 DPS) [quest] |
| wrist | Bonescythe Bracers (22483) | Bonescythe Bracers [quest] | 144.4 | yes | Primal Batskin Bracers (19687, -2.86 DPS, sim-verified) [crafted]; Rockfury Bracers (21186, -4.63 DPS) [quest]; Marshal's Leather Armsplints (16460, -6.48 DPS) [pvp] |
| hands | Bonescythe Gauntlets (22481) | Bonescythe Gauntlets [quest] | 237.0 | yes | Devilsaur Gauntlets (15063, -5.07 DPS) [crafted]; Marshal's Leather Handgrips (16454, -5.30 DPS) [vendor]; Stormshroud Gloves (21278, -5.90 DPS, sim-verified) [crafted] |
| waist | Bonescythe Waistguard (22482) | Bonescythe Waistguard [quest] | 0.0 | yes | Defiler's Leather Girdle (20193, -0.44 DPS) [rep]; Belt of the Archmage (18405, -1.51 DPS) [crafted]; Defiler's Leather Girdle (20190, -2.73 DPS, sim-verified) [rep] |
| legs | Bonescythe Legplates (22477) | Bonescythe Legplates [quest] | 0.0 | yes | Marshal's Leather Leggings (16456, +0.00 DPS) [vendor]; General's Leather Legguards (16564, +0.00 DPS) [vendor]; Marshal's Leather Leggings (231548, +0.00 DPS) [vendor] |
| feet | Bonescythe Sabatons (22480) | Bonescythe Sabatons [quest] | 235.0 | yes | Deathdealer's Boots (21359, -2.49 DPS, sim-verified) [quest]; Fine Dawn Treaders (227815, -8.75 DPS) [vendor]; Darkmantle Footpads (226831, -8.85 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | 187.0 | yes | Band of the Penitent (13217, -3.90 DPS) [quest]; Dragonslayer's Signet (18403, -3.90 DPS) [quest]; Ring of Entropy (18543, -3.90 DPS) [world] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | 171.0 | yes | Dragonslayer's Signet (18403, -3.05 DPS) [quest]; Ring of Entropy (18543, -3.05 DPS) [world]; Band of the Penitent (13217, -3.24 DPS, sim-verified) [quest] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | 0.0 | yes | Guardian Talisman (1490, -4.37 DPS) [quest]; Ankh of Life (1713, -4.37 DPS) [world_drop]; Blazing Emblem (2802, -4.37 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | 0.0 | yes | Guardian Talisman (1490, -2.74 DPS) [quest]; Ankh of Life (1713, -2.74 DPS) [world_drop]; Blazing Emblem (2802, -2.74 DPS) [world_drop] |
| main_hand | Shadowsong's Sorrow (21522) | Treasure of the Timeless One [quest] | 0.0 | yes | Grand Marshal's Dirk (18838, +0.00 DPS) [vendor]; High Warlord's Razor (234556, +0.00 DPS) [vendor]; Grand Marshal's Dirk (234582, +0.00 DPS) [vendor] |
| off_hand | Dagger of Veiled Shadows (21404) | Dagger of Veiled Shadows [quest] | 798.4 | yes | Thermotastic Egg Timer (9644, -42.25 DPS) [quest]; Imperial Red Scepter (15930, -42.31 DPS) [world_drop]; Elunarian Sphere (15968, -42.31 DPS) [world_drop] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | 0.0 | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; High Warlord's Street Sweeper (234561, +0.00 DPS) [vendor] |

**New at 60:** head: Bonescythe Helmet; neck: Blazefury Medallion; shoulder: Bonescythe Pauldrons; back: Chromatic Cloak; chest: Bonescythe Breastplate; wrist: Bonescythe Bracers; hands: Bonescythe Gauntlets; waist: Bonescythe Waistguard; legs: Bonescythe Legplates; feet: Bonescythe Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; main_hand: Shadowsong's Sorrow; off_hand: Dagger of Veiled Shadows; ranged: The Purifier

No-known-source sample (15 of 1652, see the JSON for more): 1189 Overseer's Ring; 1447 Ring of Saviors; 1832 Lucky Trousers; 2277 Necromancer Leggings; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring

