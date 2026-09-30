# Leveling BiS: Subtlety

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power (Onyxia Tooth Pendant, Earthweave Cloak) still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (night-elf, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.7. Weights run: 1.1s. Verify run: 1.5s. 159 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.290 ± 0.038, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.37 DPS) [crafted]; Lucky Fishing Hat (19972, -0.37 DPS) [quest]; Flying Tiger Goggles (4368, -0.46 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.12 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.28 DPS) [quest]; Tarnished Locket (279870, -0.28 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Forest Leather Mantle (4709, -0.23 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.28 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.01 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.09 DPS) [world_drop] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 11.2 | yes | Brawler's Leather Armor (252490, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Armor (252491, -0.18 DPS) [crafted]; Dark Leather Tunic (2317, -0.23 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Bravo's Armbands (270015, -0.05 DPS) [quest]; Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.09 DPS) [crafted]; Forest Leather Gloves (3058, -0.09 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.54 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.63 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.09 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.28 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Pyrewood Signet Ring (277210, -0.09 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon]; The 1 Ring (8350, -0.23 DPS) [world] |
| finger2 | Protector's Band (20439) (or Pyrewood Signet Ring (277210)) | Silverwing Sentinels [rep] | 4.1 | yes | Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; The 1 Ring (8350, -0.14 DPS) [world]; Pyrewood Signet Ring (277210, -0.20 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 | yes | Assassin's Blade (1935, -0.32 DPS) [dungeon]; Evocator's Blade (2567, -0.65 DPS) [dungeon]; Buzzer Blade (2169, -1.48 DPS) [dungeon] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | 222.7 | yes | Assassin's Blade (1935, -0.03 DPS, sim-verified) [dungeon]; Grayson's Torch (1172, -10.09 DPS) [quest]; Pulsating Hydra Heart (5183, -10.09 DPS) [world] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 279.0 | yes | Lil Timmy's Peashooter (13136, +0.00 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.64 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Blackfang; off_hand: Edward's Knife; ranged: Ranger Bow

No-known-source sample (15 of 159, see the JSON for more): 1189 Overseer's Ring; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape; 14150 Robe of Evocation; 14151 Chanting Blade; 14389 Durability Shoulderpads; 20438 Outrunner's Bow; 20440 Protector's Sword; 20443 Sentinel's Blade; 202256 Privateer's Ornate Pistol

### Band 30 (night-elf, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 38.8. Weights run: 1.2s. Verify run: 1.6s. 279 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.421 ± 0.047, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.1 | yes | Brawler's Leather Hood (252504, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.11 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19541, -0.30 DPS, sim-verified) [rep]; Sentinel's Medallion (20444, -0.37 DPS) [rep]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.31 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Tigerstrike Mantle (13108, -0.09 DPS, sim-verified) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop]; Cloak of Night (4447, -0.18 DPS) [world] |
| chest | Raptorbane Armor (3566) | Ormer's Revenge [quest] | 16.0 | yes | Dusky Leather Armor (7374, -0.07 DPS, sim-verified) [crafted]; Tunic of Westfall (2041, -0.23 DPS) [quest]; Green Leather Armor (4255, -0.37 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.09 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.18 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Pilferer's Gloves (7358, -0.40 DPS, sim-verified) [crafted]; Wolfclaw Gloves (1978, -0.46 DPS) [dungeon]; Toughened Leather Gloves (4253, -0.46 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Skulker's Leather Belt (252520, -0.69 DPS) [crafted] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Troll's Bane Leggings (13114, -0.55 DPS) [world_drop]; Petrolspill Leggings (9509, -0.59 DPS, sim-verified) [dungeon]; Dusky Leather Leggings (7373, -0.60 DPS) [crafted] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.04 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Highlander's Mail Greaves (20123, -0.18 DPS) [vendor] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.1 | yes | Monkey Ring (6748, -0.09 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.09 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Protector's Band (19517, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Darkspear Insurgent's Spellblade (272086, +0.00 DPS, sim-verified) [vendor]; Thornspike (6681, -1.57 DPS) [dungeon]; Claw of the Shadowmancer (2912, -1.61 DPS) [world_drop] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (38.8 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -4.45 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -13.74 DPS) [world_drop]; Totem of Infliction (1131, -13.78 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 397.1 | yes | Silver Star (3463, -0.20 DPS, sim-verified) [quest]; Ironweaver (13137, -1.15 DPS) [world_drop]; Nightstalker Bow (6696, -2.14 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Raptorbane Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Highlander's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Talisman of Arathor; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 279, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14145 Cursed Felblade; 14148 Crystalline Cuffs; 14149 Subterranean Cape

### Band 40 (night-elf, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 76.8. Weights run: 1.3s. Verify run: 1.8s. 402 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.663 ± 0.081, hit=not significant (0.000 ± 0.000), melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.2 | yes | White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop]; Nightscape Headband (8176, -0.11 DPS, sim-verified) [crafted] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Sentinel's Medallion (19540, +0.00 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.29 DPS) [rep]; Sentinel's Medallion (20444, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -0.45 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 17.2 | yes | Nightscape Tunic (8175, -0.10 DPS) [crafted]; Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Raptorbane Armor (3566, -0.72 DPS, sim-verified) [quest] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.64 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 29.3 | yes | Skulker's Leather Gloves (252525, -0.95 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.95 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.33 DPS, sim-verified) [crafted] |
| waist | Highlander's Leather Girdle (20116) | The League of Arathor [rep] | 30.0 | yes | Highlander's Leather Girdle (20117, -0.30 DPS) [rep]; Highlander's Chain Girdle (20090, -0.38 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest]; Petrolspill Leggings (9509, -0.59 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.2 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.58 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.54 DPS) [world_drop]; Ironspine's Eye (7686, -0.54 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.14 DPS) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 432.8 | yes | Coldrage Dagger (10761, -0.86 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272085, -1.37 DPS) [vendor]; Hypnotic Blade (7714, -2.86 DPS) [dungeon] |
| off_hand | Black Menace (6831) (or Coldrage Dagger (10761)) | In the Name of the Light [quest] | 415.4 | yes | Coldrage Dagger (10761, +0.00 DPS, sim-verified) [dungeon]; Satyr's Rod (15962, -20.53 DPS) [world_drop]; Totem of Infliction (1131, -20.58 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 464.6 | yes | Swiftwind (13038, -0.20 DPS, sim-verified) [world_drop]; Skystriker Bow (13020, -2.05 DPS) [world_drop]; Crusader Bow (15287, -2.10 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Black Menace; ranged: The Silencer

No-known-source sample (15 of 402, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7749 Omega Orb; 7948 Girdle of Thero-shan

### Band 50 (night-elf, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 110.8. Weights run: 1.3s. Verify run: 2.0s. 513 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.218 ± 0.061, crit=6.499 ± 0.352, hit=not significant (0.000 ± 0.000), melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 107.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.87 DPS) [dungeon]; Helm of Fire (8348, -4.67 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 15.8 | yes | Ghostshard Talisman (7731, -0.10 DPS) [dungeon]; Sentinel's Medallion (19540, -0.13 DPS) [rep]; Sentinel's Medallion (19539, -0.47 DPS, sim-verified) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 99.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -3.98 DPS) [vendor]; Forest Tracker Epaulets (2278, -4.63 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 18.3 | yes | Serpentskin Cloak (8259, -0.26 DPS) [dungeon]; Nightscape Cloak (8195, -0.33 DPS) [crafted]; Blackflame Cape (13109, -0.35 DPS, sim-verified) [world_drop] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 109.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -4.38 DPS) [quest]; Warbear Harness (15064, -4.71 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.42 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 111.0 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.44 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.08 DPS) [crafted] |
| waist | Highlander's Leather Girdle (20115) | The League of Arathor [rep] | 111.0 | yes | Highlander's Lizardhide Girdle (20103, -1.08 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.47 DPS, sim-verified) [rep]; Highlander's Leather Girdle (20116, -4.38 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (110.8 DPS) | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.15 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.49 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.4 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.34 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.34 DPS) [vendor] |
| finger1 | Assault Band (13095) (or Blackstone Ring (17713)) | World drop [world_drop] | 20.0 | yes | Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Insurgent's Band (272065, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.42 DPS) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 20.0 | yes | Insurgent's Band (272065, -0.27 DPS) [vendor]; Ring of the Underwood (2951, -0.42 DPS) [world_drop]; Masons Fraternity Ring (9533, -1.63 DPS, sim-verified) [quest] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (109.0 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest] |
| trinket2 | - | - |  |  |  |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 | yes | Charstone Dirk (17710, -0.45 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272084, -1.17 DPS) [vendor]; Widowmaker (4091, -3.23 DPS) [world_drop] |
| off_hand | Lifeforce Dirk (10750) (or Charstone Dirk (17710)) | The God Hakkar [quest] | 503.2 | yes | Charstone Dirk (17710, -26.06 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -27.04 DPS) [quest]; Satyr's Rod (15962, -27.17 DPS) [world_drop] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (109.0 DPS) | yes | Dark Iron Rifle (16004, -0.44 DPS, sim-verified) [crafted]; Gryphonwing Long Bow (13022, -3.19 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -3.46 DPS) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Highlander's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger2: Blackstone Ring; trinket1: Frozen Heart of the Mountain; trinket2: Smoking Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 513, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

### Band 60 (night-elf, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 247.7. Weights run: 1.3s. Verify run: 2.2s. 957 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=8.118 ± 0.504, hit=not significant (5.733 ± 1.563), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 286.2 | yes | Bloodvine Goggles (19999, -3.08 DPS) [crafted]; Ragefury Eyepatch (11735, -3.13 DPS) [dungeon]; Mask of the Unforgiven (13404, -10.23 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (244.2 DPS) | yes | Dragonheart Necklace (20622, +0.00 DPS) [world]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -1.47 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 179.0 | yes | Champion's Leather Shoulders (23258, +0.00 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, +0.00 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (247.7 DPS) | yes | Earthweave Cloak (21187, -0.55 DPS) [quest]; Arcanoweave Cloak (272411, -1.49 DPS) [vendor]; Chromatic Cloak (18509, -3.82 DPS, sim-verified) [crafted] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | 228.9 | yes | Dawn Armor (252483, -0.75 DPS) [crafted]; Legionnaire's Leather Chestpiece (22879, -1.27 DPS) [vendor]; Stormshroud Armor (15056, -11.33 DPS, sim-verified) [crafted] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 86.9 | yes | Duskwraith Wristguards (239547, -0.69 DPS) [vendor]; Rockfury Bracers (21186, -1.57 DPS) [quest]; Primal Batskin Bracers (19687, -2.38 DPS, sim-verified) [crafted] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 212.3 | yes | Devilsaur Gauntlets (15063, -3.76 DPS) [crafted]; Marshal's Leather Handgrips (16454, -3.99 DPS) [vendor]; Stormshroud Gloves (21278, -6.93 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 210.0 | yes | Assassin's Waistguard (272395, -3.75 DPS) [vendor]; Highlander's Leather Girdle (20115, -4.06 DPS) [rep]; Highlander's Leather Girdle (20045, -5.74 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 287.4 | yes | Marshal's Leather Leggings (16456, -1.44 DPS) [vendor]; General's Leather Legguards (16564, -1.44 DPS) [vendor]; Sentinel's Leather Pants (237818, -8.30 DPS, sim-verified) [vendor] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 212.3 | yes | Duskwraith Treads (239553, -5.19 DPS, sim-verified) [vendor]; Fine Dawn Treaders (227815, -7.55 DPS) [vendor]; Darkmantle Footpads (226831, -7.65 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (242.6 DPS) | yes | Band of the Penitent (13217, -3.90 DPS) [quest]; Ring of Entropy (18543, -3.90 DPS) [world]; Wrath of Cenarius (21190, -5.95 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (242.6 DPS) | yes | Band of the Penitent (13217, -3.05 DPS) [quest]; Ring of Entropy (18543, -3.05 DPS) [world]; Wrath of Cenarius (21190, -5.03 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (240.0 DPS) | yes | Thunderbrew's Boot Flask (744, +0.00 DPS) [quest]; Tidal Charm (1404, +0.00 DPS) [vendor]; Darkmoon Card: Heroism (19287, -9.64 DPS, sim-verified) [quest] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (242.6 DPS) | yes | Darkmoon Card: Heroism (19287, -2.19 DPS, sim-verified) [quest]; Thunderbrew's Boot Flask (744, -2.74 DPS) [quest]; Tidal Charm (1404, -2.74 DPS) [vendor] |
| main_hand | High Warlord's Shiv (235478) | Sergeant Thunderhorn [vendor] | sim-verified (242.6 DPS) | yes | High Warlord's Razor (234556, -0.22 DPS) [vendor]; Grand Marshal's Dirk (234582, -0.22 DPS) [vendor]; The Lobotomizer (19324, -12.52 DPS, sim-verified) [rep] |
| off_hand | Grand Marshal's Shiv (235479) | Captain O'Neal [vendor] | 1006.7 | yes | Shadowsong's Sorrow (21522, -7.26 DPS, sim-verified) [quest]; Thermotastic Egg Timer (9644, -53.32 DPS) [quest]; Highborne Star (15967, -53.32 DPS) [world_drop] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (244.2 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -1.89 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Blazefury Medallion; back: Howler's Furs; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Frozen Heart of the Mountain; main_hand: High Warlord's Shiv; off_hand: Grand Marshal's Shiv; ranged: The Purifier

No-known-source sample (15 of 957, see the JSON for more): 1189 Overseer's Ring; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers

## Horde

### Band 20 (troll, 00000000000000000-00000000000000000-5321000000000000000)

Set DPS (verified): 30.5. Weights run: 1.1s. Verify run: 1.6s. 166 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.015 ± 0.005, crit=0.290 ± 0.038, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.333 ± 0.713)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 8.1 | yes | Shadow Goggles (4373, -0.37 DPS) [crafted]; Lucky Fishing Hat (19972, -0.37 DPS) [quest]; Flying Tiger Goggles (4368, -0.44 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 6.1 | yes | Erudite's Amulet (277204, -0.11 DPS, sim-verified) [quest]; Scholarly Pendant (277203, -0.28 DPS) [quest]; Tarnished Locket (279870, -0.28 DPS) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 5.1 | yes | Reinforced Woolen Shoulders (4315, -0.23 DPS) [crafted]; Forest Leather Mantle (4709, -0.23 DPS) [world_drop]; Double-Stitched Woolen Shoulders (4314, -0.27 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 6.1 | yes | Catacomb Cloak (279899, -0.00 DPS, sim-verified) [quest]; Cape of the Brotherhood (5193, -0.05 DPS) [dungeon]; Sentry Cloak (2059, -0.09 DPS) [world_drop] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 7.1 | yes | Dark Leather Tunic (2317, -0.05 DPS) [crafted]; Prospector's Chestpiece (14562, -0.05 DPS) [world_drop]; Trapper's Leather Armor (252491, -0.27 DPS, sim-verified) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 5.1 | yes | Wolf Bracers (4794, -0.06 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.09 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.09 DPS) [vendor] |
| hands | Serpent Gloves (5970) (or Gloves of the Fang (10413)) | Wailing Caverns: Lord Serpentis [dungeon] | 6.1 | yes | Gloves of the Fang (10413, +0.00 DPS, sim-verified) [dungeon]; Fletcher's Gloves (7348, -0.09 DPS) [crafted]; Forest Leather Gloves (3058, -0.09 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 | yes | Deviate Scale Belt (6468, -0.54 DPS, sim-verified) [crafted]; Dusty Belt (279897, -0.59 DPS) [quest]; Guardsman Belt (3429, -0.63 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 9.1 | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.05 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 8.1 | yes | Footpads of the Fang (10411, -0.09 DPS) [dungeon]; Dark Leather Boots (2315, -0.14 DPS) [crafted]; Blackened Defias Boots (10402, -0.27 DPS, sim-verified) [dungeon] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 6.1 | yes | Pyrewood Signet Ring (277210, -0.09 DPS) [quest]; Bounty Hunter's Ring (5351, -0.14 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.18 DPS) [dungeon] |
| finger2 | Legionnaire's Band (20429) (or Pyrewood Signet Ring (277210)) | Warsong Outriders [rep] | 4.1 | yes | Bounty Hunter's Ring (5351, -0.05 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.09 DPS) [dungeon]; Pyrewood Signet Ring (277210, -0.22 DPS, sim-verified) [quest] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Blackfang (2236) | World drop [world_drop] | 228.6 | yes | Assassin's Blade (1935, -0.32 DPS) [dungeon]; Evocator's Blade (2567, -0.65 DPS) [dungeon]; Buzzer Blade (2169, -1.48 DPS) [dungeon] |
| off_hand | Edward's Knife (251485) | A Frightened Request [quest] | 222.7 | yes | Assassin's Blade (1935, -0.18 DPS, sim-verified) [dungeon]; Grayson's Torch (1172, -10.09 DPS) [quest]; Nightglow Concoction (3451, -10.09 DPS) [quest] |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 279.0 | yes | Lil Timmy's Peashooter (13136, +0.00 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.64 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.92 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Serpent Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Gnomish Universal Remote; trinket2: Rune of Perfection; main_hand: Blackfang; off_hand: Edward's Knife; ranged: Ranger Bow

No-known-source sample (15 of 166, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 4763 Blackwood Recurve Bow; 5821 Darkstalker Boots; 6478 Rat Stompers; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves; 20430 Legionnaire's Sword; 20437 Outrider's Bow; 20441 Scout's Blade; 202256 Privateer's Ornate Pistol; 209612 Insignia of the Alliance; 209622 Insignia of the Horde; 241089 Scarlet Dagger

### Band 30 (troll, 00000000000000000-00000000000000000-5322210310011000000)

Set DPS (verified): 38.2. Weights run: 1.2s. Verify run: 1.6s. 291 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.010 ± 0.004, crit=0.421 ± 0.047, hit=not significant (0.000 ± 0.000), melee_haste=not significant (1.102 ± 0.685)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 10.1 | yes | Brawler's Leather Hood (252504, -0.09 DPS) [crafted]; Tribal Worg Helm (6204, -0.10 DPS, sim-verified) [world]; Humbert's Helm (4724, -0.14 DPS) [world] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19537, -0.31 DPS, sim-verified) [rep]; Scout's Medallion (20442, -0.37 DPS) [rep]; Kaleidoscope Chain (13084, -0.46 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 11.1 | yes | Dark Leather Shoulders (4252, -0.19 DPS) [crafted]; Insignia Mantle (4721, -0.19 DPS) [world_drop]; Mantle of Thieves (2264, -0.31 DPS, sim-verified) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Tigerstrike Mantle (13108, -0.09 DPS) [world_drop]; Hawkeye's Cloak (14593, -0.14 DPS) [world_drop] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 14.1 | yes | Panther Armor (6670, -0.16 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.28 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.28 DPS) [crafted] |
| wrist | Cultist's Armguards (270032) | Blackfathom Villainy [quest] | 10.0 | yes | Jurassic Wristguards (6198, -0.10 DPS, sim-verified) [world]; Insignia Bracers (6410, -0.18 DPS) [world_drop]; Hawkeye's Bracers (14590, -0.18 DPS) [world_drop] |
| hands | Heavy Earthen Gloves (7359) | Leatherworking [crafted] | 16.0 | yes | Pilferer's Gloves (7358, -0.41 DPS, sim-verified) [crafted]; Braced Handguards (6784, -0.42 DPS) [quest]; Wolfclaw Gloves (1978, -0.46 DPS) [dungeon] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.28 DPS) [dungeon]; Deftkin Belt (16659, -0.37 DPS) [quest] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Troll's Bane Leggings (13114, -0.55 DPS) [world_drop]; Dusky Leather Leggings (7373, -0.60 DPS) [crafted]; Petrolspill Leggings (9509, -0.61 DPS, sim-verified) [dungeon] |
| feet | Disjointed Shoes (277226) | WANTED: Incinerator Gar'im [quest] | 12.0 | yes | Feet of the Lynx (1121, -0.05 DPS, sim-verified) [world_drop]; Insignia Boots (4055, -0.18 DPS) [world_drop]; Warsong Boots (16977, -0.18 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 9.1 | yes | Monkey Ring (6748, -0.09 DPS) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| finger2 | Insurgent's Band (272067) | Creeg Bothunk [vendor] | 9.0 | yes | Monkey Ring (6748, -0.10 DPS, sim-verified) [quest]; Ring of Precision (1491, -0.14 DPS) [dungeon]; Legionnaire's Band (19513, -0.14 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Swinetusk Shank (6691) | Razorfen Kraul: Agathelos the Raging [dungeon] | 322.0 | yes | Darkspear Insurgent's Spellblade (272086, +0.00 DPS, sim-verified) [vendor]; Thornspike (6681, -1.57 DPS) [dungeon]; Claw of the Shadowmancer (2912, -1.61 DPS) [world_drop] |
| off_hand | Torturing Poker (7682) | Scarlet Monastery: Interrogator Vishas [dungeon] | sim-verified (38.2 DPS) | yes | Darkspear Insurgent's Spellblade (272086, -4.39 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -13.74 DPS) [world_drop]; Grayson's Torch (1172, -13.78 DPS) [quest] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 397.1 | yes | Silver Star (3463, -0.20 DPS, sim-verified) [quest]; Ironweaver (13137, -1.15 DPS) [world_drop]; Nightstalker Bow (6696, -2.14 DPS) [dungeon] |

**New at 30:** head: Brawler's Leather Helm; neck: Ghostshard Talisman; shoulder: Forest Tracker Epaulets; back: Wolfmaster Cape; chest: Dusky Leather Armor; wrist: Cultist's Armguards; hands: Heavy Earthen Gloves; waist: Defiler's Chain Girdle; legs: Ferine Leggings; feet: Disjointed Shoes; finger1: Ironspine's Eye; finger2: Insurgent's Band; trinket1: Defiler's Talisman; trinket2: Darkspear Voodoo Seal; main_hand: Swinetusk Shank; off_hand: Torturing Poker; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 291, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan; 7951 Hands of Thero-shan; 9362 Brilliant Gold Ring; 10047 Simple Kilt; 14389 Durability Shoulderpads; 15401 Welldrip Gloves

### Band 40 (troll, 00000000000000000-00000000000000000-5322210310013011051)

Set DPS (verified): 75.8. Weights run: 1.3s. Verify run: 1.8s. 414 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.013 ± 0.005, crit=0.663 ± 0.081, hit=not significant (0.000 ± 0.000), melee_haste=not significant (2.989 ± 1.700)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Warden's Wizard Hat (14604) | World drop [world_drop] | 13.2 | yes | Nightscape Headband (8176, -0.09 DPS, sim-verified) [crafted]; White Bandit Mask (10008, -0.10 DPS) [crafted]; Hawkeye's Helm (14591, -0.10 DPS) [world_drop] |
| neck | Ghostshard Talisman (7731) | Scarlet Monastery: Azshir the Sleepless [dungeon] | 14.0 | yes | Scout's Medallion (19536, +0.00 DPS, sim-verified) [rep]; Scout's Medallion (19537, -0.29 DPS) [rep]; Scout's Medallion (20442, -0.39 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 23.1 | yes | Forest Tracker Epaulets (2278, -0.44 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.59 DPS) [crafted]; Mantle of Thieves (2264, -0.64 DPS) [dungeon] |
| back | Wolfmaster Cape (6314) (or Wildhunter Cloak (16658)) | Shadowfang Keep: Wolf Master Nandos [dungeon] | 10.0 | yes | Wildhunter Cloak (16658, +0.00 DPS, sim-verified) [quest]; Imperial Cloak (6432, -0.09 DPS) [world_drop]; Parachute Cloak (10518, -0.09 DPS) [crafted] |
| chest | Wolffear Harness (13110) | World drop [world_drop] | 17.2 | yes | Dusky Leather Armor (7374, -0.15 DPS) [crafted]; Nightscape Tunic (8175, -0.19 DPS, sim-verified) [crafted]; Hawkeye's Tunic (14592, -0.25 DPS) [world_drop] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Imperial Leather Bracers (4061, -0.59 DPS) [world_drop]; Dusky Bracers (7378, -0.59 DPS) [crafted]; Cultist's Armguards (270032, -0.63 DPS, sim-verified) [quest] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 29.3 | yes | Skulker's Leather Gloves (252525, -0.95 DPS) [crafted]; Stalker's Leather Gloves (252526, -0.95 DPS) [crafted]; Heavy Earthen Gloves (7359, -1.31 DPS, sim-verified) [crafted] |
| waist | Defiler's Leather Girdle (20192) | The Defilers [rep] | 30.0 | yes | Defiler's Leather Girdle (20191, -0.30 DPS) [rep]; Defiler's Chain Girdle (20152, -0.38 DPS, sim-verified) [rep]; Blackened Defias Belt (10403, -0.59 DPS) [dungeon] |
| legs | Ferine Leggings (6690) | Razorfen Kraul: Agathelos the Raging [dungeon] | 26.0 | yes | Basilisk Hide Pants (1718, +0.00 DPS, sim-verified) [world_drop]; Triprunner Dungarees (9624, -0.38 DPS) [quest]; Petrolspill Leggings (9509, -0.59 DPS) [dungeon] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 13.2 | yes | Imperial Leather Boots (6431, -0.10 DPS) [world_drop]; Dusky Boots (7390, -0.10 DPS) [crafted]; Disjointed Shoes (277226, -0.53 DPS, sim-verified) [quest] |
| finger1 | Assault Band (13095) | World drop [world_drop] | 20.0 | yes | Ring of the Underwood (2951, -0.49 DPS) [world_drop]; Falcon's Hook (7552, -0.54 DPS) [world_drop]; Ironspine's Eye (7686, -0.54 DPS) [dungeon] |
| finger2 | Insurgent's Band (272066) | Creeg Bothunk [vendor] | 12.0 | yes | Ring of the Underwood (2951, +0.00 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.14 DPS) [world_drop]; Ironspine's Eye (7686, -0.14 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Jhordy's Misplaced Screwdriver (274753) | Rettrick [vendor] | 432.8 | yes | Darkspear Insurgent's Spellblade (272085, -1.37 DPS) [vendor]; Hypnotic Blade (7714, -2.86 DPS) [dungeon]; The Ziggler (8006, -3.08 DPS) [world_drop] |
| off_hand | Coldrage Dagger (10761) | Razorfen Downs: Amnennar the Coldbringer [dungeon] | 415.4 | yes | Darkspear Insurgent's Spellblade (272085, -19.71 DPS, sim-verified) [vendor]; Satyr's Rod (15962, -20.53 DPS) [world_drop]; Grayson's Torch (1172, -20.58 DPS) [quest] |
| ranged | The Silencer (13138) | World drop [world_drop] | 464.6 | yes | Swiftwind (13038, -0.20 DPS, sim-verified) [world_drop]; Skystriker Bow (13020, -2.05 DPS) [world_drop]; Crusader Bow (15287, -2.10 DPS) [world_drop] |

**New at 40:** head: Warden's Wizard Hat; shoulder: Sunburn Spaulders; chest: Wolffear Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Leather Girdle; feet: Swampwalker Boots; finger1: Assault Band; finger2: Insurgent's Band; trinket1: Rune of Perfection; trinket2: Ankh of Life; main_hand: Jhordy's Misplaced Screwdriver; off_hand: Coldrage Dagger; ranged: The Silencer

No-known-source sample (15 of 414, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape; 6478 Rat Stompers; 7948 Girdle of Thero-shan

### Band 50 (troll, 00500000000000000-32000000000000000-5322210310013011051)

Set DPS (verified): 113.7. Weights run: 1.3s. Verify run: 2.0s. 525 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.218 ± 0.061, crit=6.499 ± 0.352, hit=not significant (0.000 ± 0.000), melee_haste=not significant (0.912 ± 2.693)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Leather Headband (220850) (or Blood Guard's Leather Headband (220851)) | Captain Dirgehammer [vendor] | 107.0 | yes | Blood Guard's Leather Headband (220851, +0.00 DPS, sim-verified) [vendor]; Eye of Theradras (17715, -0.87 DPS) [dungeon]; Helm of Fire (8348, -4.67 DPS) [crafted] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 15.8 | yes | Ghostshard Talisman (7731, -0.10 DPS) [dungeon]; Scout's Medallion (19536, -0.13 DPS) [rep]; Scout's Medallion (19535, -0.46 DPS, sim-verified) [rep] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 99.0 | yes | Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -3.98 DPS) [vendor]; Forest Tracker Epaulets (2278, -4.63 DPS) [world_drop] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | 18.3 | yes | Serpentskin Cloak (8259, -0.26 DPS) [dungeon]; Blackflame Cape (13109, -0.33 DPS, sim-verified) [world_drop]; Nightscape Cloak (8195, -0.33 DPS) [crafted] |
| chest | Knight's Leather Armor (220854) (or Stone Guard's Leather Armor (220855)) | Captain Dirgehammer [vendor] | 109.0 | yes | Stone Guard's Leather Armor (220855, +0.00 DPS, sim-verified) [vendor]; Blazewind Breastplate (11193, -4.38 DPS) [quest]; Warbear Harness (15064, -4.71 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 | yes | Deepfury Bracers (13120, +0.00 DPS, sim-verified) [world_drop]; Wicked Leather Bracers (15084, -0.36 DPS) [crafted]; Pridelord Bands (14672, -0.42 DPS) [dungeon] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 111.0 | yes | First Sergeant's Leather Gauntlets (220857, -0.32 DPS) [vendor]; Sergeant Major's Leather Gauntlets (220856, -0.44 DPS, sim-verified) [vendor]; Fletcher's Gloves (7348, -1.08 DPS) [crafted] |
| waist | Defiler's Leather Girdle (20193) | The Defilers [rep] | 111.0 | yes | Defiler's Lizardhide Girdle (20174, -1.08 DPS) [rep]; Defiler's Cloth Girdle (20165, -1.45 DPS, sim-verified) [rep]; Defiler's Leather Girdle (20192, -4.38 DPS) [rep] |
| legs | Knight's Leather Pants (220858) | Captain Dirgehammer [vendor] | sim-verified (112.2 DPS) | yes | Stone Guard's Leather Pants (220859, +0.00 DPS) [vendor]; Stormshroud Pants (15057, -1.47 DPS, sim-verified) [crafted]; Ferine Leggings (6690, -4.49 DPS) [dungeon] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 24.4 | yes | Sandstalker Ankleguards (12470, +0.00 DPS, sim-verified) [dungeon]; Sergeant Major's Leather Boots (220860, -0.34 DPS) [vendor]; First Sergeant's Leather Boots (220861, -0.34 DPS) [vendor] |
| finger1 | White Bone Band (11862) | Bone-Bladed Weapons [quest] | 24.0 | yes | Masons Fraternity Ring (9533, -0.38 DPS) [quest]; Insurgent's Band (272065, -0.49 DPS) [vendor]; Assault Band (13095, -1.74 DPS, sim-verified) [world_drop] |
| finger2 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | sim-verified (112.1 DPS) | yes | Masons Fraternity Ring (9533, -0.16 DPS) [quest]; Insurgent's Band (272065, -0.27 DPS) [vendor]; Assault Band (13095, -1.40 DPS, sim-verified) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (109.3 DPS) | yes | Tidal Charm (1404, -2.27 DPS) [vendor]; Guardian Talisman (1490, -2.27 DPS) [quest]; Ankh of Life (1713, -2.27 DPS) [world_drop] |
| trinket2 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (110.7 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Smoking Heart of the Mountain (11811, -1.61 DPS, sim-verified) [crafted] |
| main_hand | Julie's Dagger (6660) | World drop [world_drop] | 511.6 | yes | Charstone Dirk (17710, -0.45 DPS) [dungeon]; Darkspear Insurgent's Spellblade (272084, -1.17 DPS) [vendor]; Widowmaker (4091, -3.23 DPS) [world_drop] |
| off_hand | Lifeforce Dirk (10750) (or Charstone Dirk (17710)) | The God Hakkar [quest] | 503.2 | yes | Charstone Dirk (17710, -26.20 DPS, sim-verified) [dungeon]; Thermotastic Egg Timer (9644, -27.04 DPS) [quest]; Satyr's Rod (15962, -27.17 DPS) [world_drop] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (110.7 DPS) | yes | Dark Iron Rifle (16004, -0.44 DPS, sim-verified) [crafted]; Gryphonwing Long Bow (13022, -3.19 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -3.46 DPS) [world_drop] |

**New at 50:** head: Knight-Lieutenant's Leather Headband; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Leather Shoulders; back: Dark Phantom Cape; chest: Knight's Leather Armor; waist: Defiler's Leather Girdle; legs: Knight's Leather Pants; feet: Albino Crocscale Boots; finger1: White Bone Band; finger2: Blackstone Ring; trinket1: Rune of the Guard Captain; trinket2: Frozen Heart of the Mountain; main_hand: Julie's Dagger; off_hand: Lifeforce Dirk; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 525, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

### Band 60 (troll, 00500000000000000-32513100000000000-5322210310013011051)

Set DPS (verified): 248.5. Weights run: 1.3s. Verify run: 2.2s. 970 eligible items had no known source.

Stat weights (normalized to attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): attack_power=1.000 ± 0.001, agility=1.181 ± 0.023, crit=8.118 ± 0.504, hit=not significant (5.733 ± 1.563), melee_haste=not significant (-0.103 ± 4.302)

| Slot | Item | Source | Score | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Duskwraith Helmet (239560) | Leonid Barthalomew the Revered [vendor] | 286.2 | yes | Bloodvine Goggles (19999, -3.08 DPS) [crafted]; Ragefury Eyepatch (11735, -3.13 DPS) [dungeon]; Mask of the Unforgiven (13404, -10.43 DPS, sim-verified) [dungeon] |
| neck | Blazefury Medallion (17111) | Lord Kazzak [world] | sim-verified (245.8 DPS) | yes | Dragonheart Necklace (20622, +0.00 DPS) [world]; Beads of Ogre Might (22150, +0.00 DPS) [quest]; Medallion of the Dawn (22659, -0.84 DPS, sim-verified) [quest] |
| shoulder | Knight-Lieutenant's Leather Shoulders (220852) (or Blood Guard's Leather Shoulders (220853)) | Captain Dirgehammer [vendor] | 179.0 | yes | Champion's Leather Shoulders (23258, +0.00 DPS) [vendor]; Lieutenant Commander's Leather Shoulders (23313, +0.00 DPS) [vendor]; Blood Guard's Leather Shoulders (220853, +0.00 DPS, sim-verified) [vendor] |
| back | Howler's Furs (272414) | Pix Xizzix [vendor] | sim-verified (248.5 DPS) | yes | Earthweave Cloak (21187, -0.55 DPS) [quest]; Arcanoweave Cloak (272411, -1.49 DPS) [vendor]; Chromatic Cloak (18509, -3.46 DPS, sim-verified) [crafted] |
| chest | Duskwraith Breastplate (239562) | Leonid Barthalomew the Revered [vendor] | 228.9 | yes | Dawn Armor (252483, -0.75 DPS) [crafted]; Legionnaire's Leather Chestpiece (227060, -1.27 DPS) [vendor]; Stormshroud Armor (15056, -11.89 DPS, sim-verified) [crafted] |
| wrist | Duskwraith Bracers (239555) | Leonid Barthalomew the Revered [vendor] | 86.9 | yes | Duskwraith Wristguards (239547, -0.69 DPS) [vendor]; Rockfury Bracers (21186, -1.57 DPS) [quest]; Primal Batskin Bracers (19687, -2.52 DPS, sim-verified) [crafted] |
| hands | Duskwraith Gauntlets (239557) | Leonid Barthalomew the Revered [vendor] | 212.3 | yes | Devilsaur Gauntlets (15063, -3.76 DPS) [crafted]; Marshal's Leather Handgrips (16454, -3.99 DPS) [vendor]; Stormshroud Gloves (21278, -7.09 DPS, sim-verified) [crafted] |
| waist | Duskwraith Waistguard (239556) | Leonid Barthalomew the Revered [vendor] | 210.0 | yes | Assassin's Waistguard (272395, -3.75 DPS) [vendor]; Defiler's Leather Girdle (20193, -4.06 DPS) [rep]; Defiler's Leather Girdle (20190, -6.41 DPS, sim-verified) [rep] |
| legs | Duskwraith Legplates (239561) | Leonid Barthalomew the Revered [vendor] | 287.4 | yes | Marshal's Leather Leggings (16456, -1.44 DPS) [vendor]; General's Leather Legguards (16564, -1.44 DPS) [vendor]; Sentinel's Leather Pants (237818, -8.88 DPS, sim-verified) [vendor] |
| feet | Duskwraith Sabatons (239558) | Leonid Barthalomew the Revered [vendor] | 212.3 | yes | Duskwraith Treads (239553, -5.43 DPS, sim-verified) [vendor]; Fine Dawn Treaders (227815, -7.55 DPS) [vendor]; Darkmantle Footpads (226831, -7.65 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (245.6 DPS) | yes | Band of the Penitent (13217, -3.90 DPS) [quest]; Ring of Entropy (18543, -3.90 DPS) [world]; Wrath of Cenarius (21190, -6.53 DPS, sim-verified) [quest] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (245.6 DPS) | yes | Band of the Penitent (13217, -3.05 DPS) [quest]; Ring of Entropy (18543, -3.05 DPS) [world]; Wrath of Cenarius (21190, -5.61 DPS, sim-verified) [quest] |
| trinket1 | Darkmoon Card: Maelstrom (19289) | Darkmoon Elementals Deck [quest] | sim-verified (242.5 DPS) | yes | Tidal Charm (1404, +0.00 DPS) [vendor]; Guardian Talisman (1490, +0.00 DPS) [quest]; Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted] |
| trinket2 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (245.6 DPS) | yes | Frozen Heart of the Mountain (249469, -2.93 DPS, sim-verified) [crafted]; Tidal Charm (1404, -4.37 DPS) [vendor]; Guardian Talisman (1490, -4.37 DPS) [quest] |
| main_hand | High Warlord's Shiv (235478) | Sergeant Thunderhorn [vendor] | sim-verified (245.6 DPS) | yes | High Warlord's Razor (234556, -0.22 DPS) [vendor]; Grand Marshal's Dirk (234582, -0.22 DPS) [vendor]; The Lobotomizer (19324, -12.65 DPS, sim-verified) [rep] |
| off_hand | Grand Marshal's Shiv (235479) | Captain O'Neal [vendor] | 1006.7 | yes | Shadowsong's Sorrow (21522, -8.24 DPS, sim-verified) [quest]; Thermotastic Egg Timer (9644, -53.32 DPS) [quest]; Highborne Star (15967, -53.32 DPS) [world_drop] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (245.8 DPS) | yes | High Warlord's Recurve (234559, +0.00 DPS) [vendor]; High Warlord's Crossbow (234560, +0.00 DPS) [vendor]; Dark Iron Rifle (16004, -2.03 DPS, sim-verified) [crafted] |

**New at 60:** head: Duskwraith Helmet; neck: Blazefury Medallion; back: Howler's Furs; chest: Duskwraith Breastplate; wrist: Duskwraith Bracers; hands: Duskwraith Gauntlets; waist: Duskwraith Waistguard; legs: Duskwraith Legplates; feet: Duskwraith Sabatons; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Darkmoon Card: Maelstrom; trinket2: Rune of the Guard Captain; main_hand: High Warlord's Shiv; off_hand: Grand Marshal's Shiv; ranged: The Purifier

No-known-source sample (15 of 970, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2944 Cursed Eye of Paleth; 4110 Master Hunter's Bow; 4111 Master Hunter's Rifle; 4116 Olmann Sewar; 4196 Feathered Mantle; 4763 Blackwood Recurve Bow; 4988 Burning Obsidian Band; 4989 Mage Dragon Robe; 5000 Coral Band; 5008 Quicksilver Ring; 5743 Prismstone Ring; 5821 Darkstalker Boots; 5971 Feathered Cape

