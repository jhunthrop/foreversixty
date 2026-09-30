# Leveling BiS: Marksmanship

Prototype output of `sim/cmd/leveling-bis` (lane `bis-proto`). See the lane report for method, run times and gaps.

Forever unifies melee, ranged and spell hit into one stat, and likewise crit, one point higher on both tables at once (wowsims-forever's `proto/common.proto` StatHit/StatCrit, `data/pipeline/simdb/statmap.py`, `research/01-official-facts.md`) - a caster item carrying only Hit/Crit with no Intellect or Spell Power still raises that caster's own spell hit and crit chance, and is a correct pick, not a scoring bug.

## Alliance

### Band 20 (dwarf, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 75.1. Weights run: 1.1s. Verify run: 1.4s. 191 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.185 ± 0.045, crit=6.678 ± 0.268, hit=1.648 ± 0.065, melee_haste=3.462 ± 0.795

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.5 ranged_attack_power points (1.04 DPS) | yes | Flying Tiger Goggles (4368, -1.38 DPS, sim-verified) [crafted] |
| neck | Sentinel's Medallion (20444) | Silverwing Sentinels [rep] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Erudite's Amulet (277204, -0.26 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.9 ranged_attack_power points (0.65 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.61 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Cape of the Brotherhood (5193, +0.00 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Hide of Lupos (3018, -0.26 DPS) [world] |
| chest | Tunic of Westfall (2041) | The Defias Brotherhood [quest] | 24.0 ranged_attack_power points (1.44 DPS) | yes | Trapper's Leather Armor (252491, -0.52 DPS) [crafted]; Brawler's Leather Armor (252490, -0.64 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.65 DPS) [crafted] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.9 ranged_attack_power points (0.65 DPS) | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Bravo's Armbands (270015, -0.13 DPS) [quest]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 93.5 ranged_attack_power points (5.58 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -4.80 DPS) [dungeon]; Forest Leather Gloves (3058, -5.06 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.07 DPS) | yes | Dusty Belt (279897, -0.42 DPS) [quest]; Deviate Scale Belt (6468, -0.46 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.55 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.7 ranged_attack_power points (1.17 DPS) | yes | Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.13 DPS) [world]; Brawler's Leather Pants (252500, -0.35 DPS, sim-verified) [crafted] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.5 ranged_attack_power points (1.04 DPS) | yes | Blackened Defias Boots (10402, -0.26 DPS, sim-verified) [dungeon]; Footpads of the Fang (10411, -0.26 DPS) [dungeon]; Dark Leather Boots (2315, -0.39 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Protector's Band (20439) | Silverwing Sentinels [rep] | 8.7 ranged_attack_power points (0.52 DPS) | yes | Lavishly Jeweled Ring (1156, -0.08 DPS, sim-verified) [dungeon]; The 1 Ring (8350, -0.39 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 23.5 ranged_attack_power points (1.40 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.49 DPS) [world]; Lupine Axe (1220, -0.62 DPS) [world] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (10.65 DPS) | yes | Lil Timmy's Peashooter (13136, -1.63 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.30 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.53 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Sentinel's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Tunic of Westfall; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Protector's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 191, see the JSON for more): 1189 Overseer's Ring; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak; 5031 ZZZZZZZZ; 5036 ZZZZZ; 5255 Quilboar Tomahawk; 5748 Centaur Longbow

### Band 30 (dwarf, 0000000000000000-35305500000000000-000000000000000000)

Set DPS (verified): 90.5. Weights run: 1.2s. Verify run: 1.5s. 334 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.180 ± 0.048, crit=8.214 ± 0.316, hit=1.844 ± 0.073, melee_haste=7.178 ± 0.865

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.8 ranged_attack_power points (1.29 DPS) | yes | Tribal Worg Helm (6204, +0.00 DPS, sim-verified) [world]; Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Humbert's Helm (4724, -0.39 DPS) [world] |
| neck | Sentinel's Medallion (19541) | Silverwing Sentinels [rep] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Ghostshard Talisman (7731, -0.21 DPS, sim-verified) [dungeon]; Sentinel's Medallion (20444, -0.26 DPS) [rep]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.0 ranged_attack_power points (1.42 DPS) | yes | Dark Leather Shoulders (4252, -0.51 DPS) [crafted]; Insignia Mantle (4721, -0.51 DPS) [world_drop]; Mantle of Thieves (2264, -1.05 DPS, sim-verified) [dungeon] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Hawkeye's Cloak (14593, -0.13 DPS, sim-verified) [world_drop]; Cloak of Night (4447, -0.26 DPS) [world]; Fenrus' Hide (6340, -0.26 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.5 ranged_attack_power points (1.80 DPS) | yes | Tunic of Westfall (2041, -0.39 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.77 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.77 DPS) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410), Hawkeye's Bracers (14590)) | Razormaw Matriarch [world] | 13.1 ranged_attack_power points (0.77 DPS) | yes | Insignia Bracers (6410, +0.00 DPS, sim-verified) [world_drop]; Hawkeye's Bracers (14590, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.13 DPS) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 115.0 ranged_attack_power points (6.79 DPS) | yes | Pilferer's Gloves (7358, +0.00 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -5.84 DPS) [crafted]; Wolfclaw Gloves (1978, -6.02 DPS) [dungeon] |
| waist | Highlander's Chain Girdle (20090) (or Highlander's Leather Girdle (20117)) | The League of Arathor [rep] | 24.0 ranged_attack_power points (1.42 DPS) | yes | Highlander's Leather Girdle (20117, +0.00 DPS, sim-verified) [rep]; Skulker's Leather Belt (252520, -0.26 DPS) [crafted]; Stalker's Leather Belt (252521, -0.26 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Ambusher [dungeon] | 30.5 ranged_attack_power points (1.80 DPS) | yes | Troll's Bane Leggings (13114, +0.00 DPS, sim-verified) [world_drop]; Dusky Leather Leggings (7373, -0.13 DPS) [crafted]; Ferine Leggings (6690, -0.27 DPS) [dungeon] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Insignia Boots (4055, +0.00 DPS, sim-verified) [world_drop]; Highlander's Mail Greaves (20123, +0.00 DPS) [vendor]; Lancer Boots (6752, -0.13 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 ranged_attack_power points (1.16 DPS) | yes | Ring of Precision (1491, -0.39 DPS) [dungeon]; Protector's Band (19517, -0.39 DPS) [rep]; Signet of the Zhevra (285330, -0.39 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.3 ranged_attack_power points (0.90 DPS) | yes | Protector's Band (19517, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.31 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (90.5 DPS) | yes | Talisman of Arathor (21119, -1.65 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.9 ranged_attack_power points (1.00 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.1 ranged_attack_power points (0.77 DPS) | yes | Prison Shank (2941, +0.00 DPS, sim-verified) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (15.44 DPS) | yes | Glass Shooter (9456, -0.62 DPS) [dungeon]; Ironweaver (13137, -1.13 DPS) [world_drop]; Silver Star (3463, -1.34 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Sentinel's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Highlander's Chain Girdle; legs: Petrolspill Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 334, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 4797 Fiery Cloak; 4798 Heavy Runed Cloak; 4799 Antiquated Cloak

### Band 40 (dwarf, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 100.5. Weights run: 1.2s. Verify run: 1.6s. 563 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.222 ± 0.063, crit=11.681 ± 0.473, hit=2.043 ± 0.081, melee_haste=not significant (3.437 ± 0.971)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 163.5 ranged_attack_power points (9.59 DPS) | yes | Warden's Wizard Hat (14604, -0.00 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -8.03 DPS) [crafted]; Guard's Chain Helm (250499, -8.03 DPS) [crafted] |
| neck | Sentinel's Medallion (19540) | Silverwing Sentinels [rep] | 24.4 ranged_attack_power points (1.43 DPS) | yes | Sentinel's Medallion (19541, +0.00 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.61 DPS) [dungeon]; Sentinel's Medallion (20444, -0.65 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.4 ranged_attack_power points (2.14 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.70 DPS) [crafted]; Mantle of Thieves (2264, -0.83 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 22.2 ranged_attack_power points (1.30 DPS) | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.26 DPS) [crafted]; Tigerstrike Mantle (13108, -0.26 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 42.2 ranged_attack_power points (2.48 DPS) | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.52 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.52 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.13 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 183.5 ranged_attack_power points (10.77 DPS) | yes | Dragonscale Gauntlets (8347, +0.00 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20089) | The League of Arathor [rep] | 171.5 ranged_attack_power points (10.06 DPS) | yes | Highlander's Leather Girdle (20116, -0.01 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20090, -8.66 DPS) [rep]; Highlander's Leather Girdle (20117, -8.66 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.7 ranged_attack_power points (2.74 DPS) | yes | Triprunner Dungarees (9624, +0.00 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.91 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.91 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.9 ranged_attack_power points (1.69 DPS) | yes | Imperial Leather Boots (6431, +0.00 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.26 DPS) [crafted]; Worn Running Boots (9398, -0.26 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 22.2 ranged_attack_power points (1.30 DPS) | yes | Ironspine's Eye (7686, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop]; Protector's Band (19515, -0.26 DPS) [rep] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Assault Band (13095, -0.00 DPS) [world_drop]; Protector's Band (19515, -0.13 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Frost Tiger Blade (3854) | Blacksmithing [crafted] | sim-verified (31.1 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Gut Ripper (2164, -0.01 DPS, sim-verified) [world_drop]; Steel Spear (250605, -7.31 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (100.5 DPS) | yes | Shadowforge Bushmaster (9422, -0.46 DPS) [dungeon]; Swiftwind (13038, -0.79 DPS) [world_drop]; Bow of Searing Arrows (2825, -69.40 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Sentinel's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Highlander's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Frost Tiger Blade; ranged: The Silencer

No-known-source sample (15 of 563, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 50 (dwarf, 5500000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 132.1. Weights run: 1.3s. Verify run: 1.6s. 725 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.307 ± 0.077, crit=13.494 ± 0.534, hit=2.404 ± 0.097, melee_haste=13.786 ± 1.194

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Knight-Lieutenant's Chain Helmet (220822) | Captain Dirgehammer [vendor] | 225.8 ranged_attack_power points (13.20 DPS) | yes | Eye of Theradras (17715, -2.16 DPS) [dungeon]; Knight-Lieutenant's Mail Helmet (223075, -2.16 DPS) [vendor]; Raging Berserker's Helm (7719, -3.39 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.0 ranged_attack_power points (1.75 DPS) | yes | Sentinel's Medallion (19540, -0.27 DPS) [rep]; Sentinel's Medallion (19539, -0.32 DPS, sim-verified) [rep]; Sentinel's Medallion (19541, -0.67 DPS) [rep] |
| shoulder | Knight-Lieutenant's Chain Epaulets (220825) | Captain Dirgehammer [vendor] | 218.9 ranged_attack_power points (12.79 DPS) | yes | Knight-Lieutenant's Mail Epaulets (223073, -1.40 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -10.61 DPS) [vendor]; Phytoskin Spaulders (17749, -10.64 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (128.2 DPS) | yes | Blackveil Cape (11626, -0.13 DPS) [dungeon]; Duskbat Drape (19982, -0.13 DPS) [quest]; Blisterbane Wrap (12552, -2.21 DPS, sim-verified) [dungeon] |
| chest | Knight's Chain Armor (220828) | Captain Dirgehammer [vendor] | 223.5 ranged_attack_power points (13.06 DPS) | yes | Knight's Mail Armor (223078, -2.34 DPS, sim-verified) [vendor]; Fungus Shroud Armor (17742, -9.69 DPS) [dungeon]; Blazewind Breastplate (11193, -9.96 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.6 ranged_attack_power points (2.02 DPS) | yes | Bracers of the Stone Princess (17714, +0.00 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.54 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.54 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 208.9 ranged_attack_power points (12.21 DPS) | yes | Dragonscale Gauntlets (8347, -0.40 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Highlander's Chain Girdle (20088) (or Highlander's Leather Girdle (20115)) | The League of Arathor [rep] | 208.9 ranged_attack_power points (12.21 DPS) | yes | Highlander's Leather Girdle (20115, +0.00 DPS, sim-verified) [rep]; Highlander's Chain Girdle (20089, -0.70 DPS) [rep]; Highlander's Cloth Girdle (20097, -1.17 DPS) [rep] |
| legs | Knight's Chain Legplates (220832) | Captain Dirgehammer [vendor] | sim-verified (128.2 DPS) | yes | Knight's Mail Legplates (223074, -1.89 DPS) [vendor]; Stormshroud Pants (15057, -2.14 DPS, sim-verified) [crafted]; Basilisk Hide Pants (1718, -10.10 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 46.1 ranged_attack_power points (2.70 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.27 DPS) [world_drop]; Whisperwalk Boots (20255, -0.27 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.0 ranged_attack_power points (2.57 DPS) | yes | Ring of the Underwood (2951, -1.23 DPS) [world_drop]; Falcon's Hook (7552, -1.36 DPS) [world_drop]; Ironspine's Eye (7686, -1.36 DPS) [dungeon] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.3 ranged_attack_power points (1.89 DPS) | yes | Ring of the Underwood (2951, -0.60 DPS, sim-verified) [world_drop]; Falcon's Hook (7552, -0.67 DPS) [world_drop]; Ironspine's Eye (7686, -0.67 DPS) [dungeon] |
| trinket1 | - | - |  |  |  |
| trinket2 | Thunderbrew's Boot Flask (744) | Sweet Amber [quest] | sim-verified (126.0 DPS) | yes | Frozen Heart of the Mountain (249469, +0.00 DPS) [crafted]; Ankh of Life (1713, -0.36 DPS, sim-verified) [world_drop] |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (126.0 DPS) | yes | Shortsword of Vengeance (754, +0.00 DPS, sim-verified) [world_drop]; Frost Tiger Blade (3854, +0.00 DPS) [crafted]; Illusionary Rod (7713, +0.00 DPS) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the Light [quest] | sim-verified (127.5 DPS) | yes | Thermotastic Egg Timer (9644, -0.88 DPS) [quest]; Satyr's Rod (15962, -1.15 DPS) [world_drop]; Inventor's Focal Sword (17719, -1.42 DPS, sim-verified) [dungeon] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (126.0 DPS) | yes | Hurricane (2824, -1.42 DPS) [world_drop]; Precisely Calibrated Boomstick (2100, -1.70 DPS) [world_drop]; Dark Iron Rifle (16004, -2.19 DPS, sim-verified) [crafted] |

**New at 50:** head: Knight-Lieutenant's Chain Helmet; neck: Skibi's Pendant; shoulder: Knight-Lieutenant's Chain Epaulets; back: Dark Phantom Cape; chest: Knight's Chain Armor; wrist: Deepfury Bracers; waist: Highlander's Chain Girdle; legs: Knight's Chain Legplates; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Guardian Talisman; trinket2: Thunderbrew's Boot Flask; main_hand: Dawn's Edge; off_hand: Vanquisher's Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 725, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

### Band 60 (dwarf, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 198.5. Weights run: 1.2s. Verify run: 1.6s. 1505 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.339 ± 0.094, crit=20.135 ± 0.769, hit=3.434 ± 0.137, melee_haste=10.525 ± 1.657

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Headpiece (239540) | Leonid Barthalomew the Revered [vendor] | 678.4 ranged_attack_power points (38.50 DPS) | yes | Lieutenant Commander's Chain Helm (23306, -4.11 DPS) [vendor]; Lieutenant Commander's Chain Helm (227066, -4.11 DPS) [vendor]; Lieutenant Commander's Chain Greathelm (227086, -5.20 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (187.6 DPS) | yes | Blazefury Medallion (17111, -0.87 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -14.05 DPS) [quest]; Amulet of the Darkmoon (19491, -14.84 DPS) [quest] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | sim-verified (191.5 DPS) | yes | Dawnstalker Spaulders (239542, -3.83 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -5.15 DPS) [vendor]; Darkspear Epaulets (272106, -5.15 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 281.9 ranged_attack_power points (16.00 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Earthweave Cloak (21187, -12.06 DPS) [quest]; Howler's Furs (272414, -12.46 DPS) [vendor] |
| chest | Dawnstalker Tunic (239543) | Leonid Barthalomew the Revered [vendor] | sim-verified (187.6 DPS) | yes | Dawnstalker Breastplate (239529, -13.16 DPS) [vendor]; Tunic of Undead Slaying (23089, -13.20 DPS, sim-verified) [world]; Knight-Captain's Chain Armor (227089, -20.04 DPS) [vendor] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | sim-verified (187.6 DPS) | yes | Dawnstalker Wristguards (239544, -0.66 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -7.75 DPS, sim-verified) [world]; Primal Batskin Bracers (19687, -16.70 DPS) [crafted] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | sim-verified (190.4 DPS) | yes | Dawnstalker Handguards (239539, -2.75 DPS, sim-verified) [vendor]; Marshal's Chain Grips (16463, -4.34 DPS) [vendor]; Marshal's Chain Vices (231578, -4.34 DPS) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 384.1 ranged_attack_power points (21.80 DPS) | yes | Dawnstalker Girdle (239538, -1.47 DPS, sim-verified) [vendor]; Highlander's Chain Girdle (20043, -3.87 DPS) [rep]; Highlander's Leather Girdle (20045, -3.87 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 680.0 ranged_attack_power points (38.59 DPS) | yes | Dawnstalker Legguards (239541, -0.02 DPS, sim-verified) [vendor]; Sentinel's Leather Pants (237818, -3.01 DPS) [vendor]; Knight-Captain's Chain Legplates (227085, -4.13 DPS) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | sim-verified (192.2 DPS) | yes | Dawnstalker Boots (239537, -4.58 DPS, sim-verified) [vendor]; Marshal's Chain Boots (16462, -18.22 DPS) [vendor]; Marshal's Chain Greaves (231579, -18.22 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Stormpike Guard [rep] | sim-verified (187.6 DPS) | yes | Band of the Penitent (13217, -2.86 DPS) [quest]; Ring of Entropy (18543, -2.86 DPS) [world]; Naglering (11669, -4.84 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (187.6 DPS) | yes | Band of the Penitent (13217, -1.95 DPS) [quest]; Ring of Entropy (18543, -1.95 DPS) [world]; Naglering (11669, -3.57 DPS, sim-verified) [dungeon] |
| trinket1 | Frozen Heart of the Mountain (249469) | Enchanting [crafted] | sim-verified (187.6 DPS) | yes | - |
| trinket2 | Ankh of Life (1713) | World drop [world_drop] | sim-verified (187.6 DPS) | yes | Thunderbrew's Boot Flask (744, -0.70 DPS, sim-verified) [quest] |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (187.6 DPS) | yes | Grand Marshal's Glaive (234569, +0.00 DPS) [vendor]; Grand Marshal's Polearm (234570, +0.00 DPS) [vendor]; Electrified Dagger (19100, -0.23 DPS, sim-verified) [rep] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (187.6 DPS) | yes | Grand Marshal's Left Hand Blade (18847, +0.00 DPS) [vendor]; Grand Marshal's Left Hand Blade (234584, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -1.29 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (187.6 DPS) | yes | Grand Marshal's Bullseye (234585, -4.60 DPS) [vendor]; Grand Marshal's Repeater (234586, -4.60 DPS) [vendor]; Dark Iron Rifle (16004, -8.92 DPS, sim-verified) [crafted] |

**New at 60:** head: Dawnstalker Headpiece; neck: Medallion of the Dawn; shoulder: Dawnstalker Pauldrons; back: Chromatic Cloak; chest: Dawnstalker Tunic; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Sentinel's Chain Leggings; feet: Dawnstalker Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket1: Frozen Heart of the Mountain; trinket2: Ankh of Life; main_hand: Heartseeker; off_hand: Dawn's Edge; ranged: The Purifier

No-known-source sample (15 of 1505, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings; 4110 Master Hunter's Bow

## Horde

### Band 20 (troll, 0000000000000000-35300000000000000-000000000000000000)

Set DPS (verified): 74.6. Weights run: 1.1s. Verify run: 1.3s. 194 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.185 ± 0.045, crit=6.678 ± 0.268, hit=1.648 ± 0.065, melee_haste=3.462 ± 0.795

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Hood (252504) | Leatherworking [crafted] | 17.5 ranged_attack_power points (1.04 DPS) | yes | Flying Tiger Goggles (4368, -0.91 DPS, sim-verified) [crafted] |
| neck | Scout's Medallion (20442) | Warsong Outriders [rep] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Erudite's Amulet (277204, -0.25 DPS, sim-verified) [quest] |
| shoulder | Serpent's Shoulders (5404) | Wailing Caverns: Lady Anacondra [dungeon] | 10.9 ranged_attack_power points (0.65 DPS) | yes | Double-Stitched Woolen Shoulders (4314, -0.59 DPS, sim-verified) [crafted] |
| back | Glowing Lizardscale Cloak (6449) | Wailing Caverns: Skum [dungeon] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Cape of the Brotherhood (5193, -0.23 DPS, sim-verified) [dungeon]; Sentry Cloak (2059, -0.26 DPS) [world_drop]; Hide of Lupos (3018, -0.26 DPS) [world] |
| chest | Brawler's Leather Armor (252490) (or Trapper's Leather Armor (252491)) | Leatherworking [crafted] | 15.3 ranged_attack_power points (0.91 DPS) | yes | Trapper's Leather Armor (252491, +0.00 DPS, sim-verified) [crafted]; Dark Leather Tunic (2317, -0.13 DPS) [crafted]; Prospector's Chestpiece (14562, -0.13 DPS) [world_drop] |
| wrist | Forest Leather Bracers (3202) | World drop [world_drop] | 10.9 ranged_attack_power points (0.65 DPS) | yes | Wolf Bracers (4794, -0.13 DPS, sim-verified) [vendor]; Bristlebark Bindings (14569, -0.26 DPS) [world_drop]; Ratchet Wristwraps (274742, -0.26 DPS) [vendor] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 93.5 ranged_attack_power points (5.58 DPS) | yes | Serpent Gloves (5970, +0.00 DPS, sim-verified) [dungeon]; Gloves of the Fang (10413, -4.80 DPS) [dungeon]; Forest Leather Gloves (3058, -5.06 DPS) [world_drop] |
| waist | Blackened Defias Belt (10403) | Westfall: Captain Greenskin [dungeon] | 18.0 ranged_attack_power points (1.07 DPS) | yes | Dusty Belt (279897, -0.42 DPS) [quest]; Deviate Scale Belt (6468, -0.44 DPS, sim-verified) [crafted]; Guardsman Belt (3429, -0.55 DPS) [world_drop] |
| legs | Leggings of the Fang (10410) (or Brawler's Leather Pants (252500), Trapper's Leather Pants (252501)) | Wailing Caverns: Lord Cobrahn [dungeon] | 19.7 ranged_attack_power points (1.17 DPS) | yes | Brawler's Leather Pants (252500, +0.00 DPS, sim-verified) [crafted]; Trapper's Leather Pants (252501, +0.00 DPS) [crafted]; Bluegill Breeches (3022, -0.13 DPS) [world] |
| feet | Feet of the Lynx (1121) | World drop [world_drop] | 17.5 ranged_attack_power points (1.04 DPS) | yes | Blackened Defias Boots (10402, -0.25 DPS, sim-verified) [dungeon]; Footpads of the Fang (10411, -0.26 DPS) [dungeon]; Dark Leather Boots (2315, -0.39 DPS) [crafted] |
| finger1 | Signet of the Zhevra (285330) | Swiftmane [world] | 13.1 ranged_attack_power points (0.78 DPS) | yes | Bounty Hunter's Ring (5351, -0.39 DPS) [quest]; Lavishly Jeweled Ring (1156, -0.52 DPS) [dungeon]; The 1 Ring (8350, -0.65 DPS) [world] |
| finger2 | Legionnaire's Band (20429) | Warsong Outriders [rep] | 8.7 ranged_attack_power points (0.52 DPS) | yes | Bounty Hunter's Ring (5351, -0.13 DPS, sim-verified) [quest]; Lavishly Jeweled Ring (1156, -0.26 DPS) [dungeon]; The 1 Ring (8350, -0.39 DPS) [world] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Bronze Dory (250603) | Blacksmithing [crafted] | 23.5 ranged_attack_power points (1.40 DPS) | yes | Impaling Harpoon (5200, +0.00 DPS, sim-verified) [dungeon]; Scythe Axe (5749, -0.49 DPS) [world]; Crescent Staff (6505, -0.49 DPS) [quest] |
| off_hand | - | - |  |  |  |
| ranged | Ranger Bow (3021) | World drop [world_drop] | 178.4 ranged_attack_power points (10.65 DPS) | yes | Lil Timmy's Peashooter (13136, -0.75 DPS, sim-verified) [world_drop]; Cracked Blacksmith Hammer (285279, -2.30 DPS) [crafted]; Lovingly Crafted Boomstick (4372, -2.53 DPS) [crafted] |

**New at 20:** head: Brawler's Leather Hood; neck: Scout's Medallion; shoulder: Serpent's Shoulders; back: Glowing Lizardscale Cloak; chest: Brawler's Leather Armor; wrist: Forest Leather Bracers; hands: Fletcher's Gloves; waist: Blackened Defias Belt; legs: Leggings of the Fang; feet: Feet of the Lynx; finger1: Signet of the Zhevra; finger2: Legionnaire's Band; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Bronze Dory; ranged: Ranger Bow

No-known-source sample (15 of 194, see the JSON for more): 1189 Overseer's Ring; 1832 Lucky Trousers; 2664 Spinner Fang; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5031 ZZZZZZZZ; 5036 ZZZZZ; 5255 Quilboar Tomahawk; 5748 Centaur Longbow; 5821 Darkstalker Boots; 5968 Rugged Boots

### Band 30 (troll, 0000000000000000-35305500000000000-000000000000000000)

Set DPS (verified): 91.8. Weights run: 1.2s. Verify run: 1.5s. 336 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.180 ± 0.048, crit=8.214 ± 0.316, hit=1.844 ± 0.073, melee_haste=7.178 ± 0.865

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Brawler's Leather Helm (252512) | Leatherworking [crafted] | 21.8 ranged_attack_power points (1.29 DPS) | yes | Tribal Worg Helm (6204, +0.00 DPS, sim-verified) [world]; Brawler's Leather Hood (252504, -0.26 DPS) [crafted]; Humbert's Helm (4724, -0.39 DPS) [world] |
| neck | Scout's Medallion (19537) | Warsong Outriders [rep] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Ghostshard Talisman (7731, -0.22 DPS, sim-verified) [dungeon]; Scout's Medallion (20442, -0.26 DPS) [rep]; Kaleidoscope Chain (13084, -0.51 DPS) [world_drop] |
| shoulder | Forest Tracker Epaulets (2278) | World drop [world_drop] | 24.0 ranged_attack_power points (1.42 DPS) | yes | Mantle of Thieves (2264, -0.41 DPS, sim-verified) [dungeon]; Dark Leather Shoulders (4252, -0.51 DPS) [crafted]; Insignia Mantle (4721, -0.51 DPS) [world_drop] |
| back | Tigerstrike Mantle (13108) | World drop [world_drop] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Hawkeye's Cloak (14593, -0.13 DPS, sim-verified) [world_drop]; Cloak of Night (4447, -0.26 DPS) [world]; Fenrus' Hide (6340, -0.26 DPS) [dungeon] |
| chest | Dusky Leather Armor (7374) | Leatherworking [crafted] | 30.5 ranged_attack_power points (1.80 DPS) | yes | Panther Armor (6670, -0.69 DPS, sim-verified) [quest]; Green Leather Armor (4255, -0.77 DPS) [crafted]; Brawler's Leather Tunic (252508, -0.77 DPS) [crafted] |
| wrist | Jurassic Wristguards (6198) (or Insignia Bracers (6410), Hawkeye's Bracers (14590)) | Razormaw Matriarch [world] | 13.1 ranged_attack_power points (0.77 DPS) | yes | Insignia Bracers (6410, +0.00 DPS, sim-verified) [world_drop]; Hawkeye's Bracers (14590, +0.00 DPS) [world_drop]; Madwolf Bracers (897, -0.13 DPS) [world] |
| hands | Fletcher's Gloves (7348) | Leatherworking [crafted] | 115.0 ranged_attack_power points (6.79 DPS) | yes | Pilferer's Gloves (7358, +0.00 DPS, sim-verified) [crafted]; Heavy Earthen Gloves (7359, -5.84 DPS) [crafted]; Braced Handguards (6784, -5.89 DPS) [quest] |
| waist | Defiler's Chain Girdle (20152) (or Defiler's Leather Girdle (20191)) | The Defilers [rep] | 24.0 ranged_attack_power points (1.42 DPS) | yes | Defiler's Leather Girdle (20191, +0.00 DPS, sim-verified) [rep]; Deftkin Belt (16659, -0.19 DPS) [quest]; Skulker's Leather Belt (252520, -0.26 DPS) [crafted] |
| legs | Petrolspill Leggings (9509) (or Troll's Bane Leggings (13114)) | Gnomeregan: Caverndeep Ambusher [dungeon] | 30.5 ranged_attack_power points (1.80 DPS) | yes | Dusky Leather Leggings (7373, -0.13 DPS) [crafted]; Ferine Leggings (6690, -0.27 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.38 DPS, sim-verified) [world_drop] |
| feet | Feet of the Lynx (1121) (or Insignia Boots (4055), Vorrel's Boots (7751), Warsong Boots (16977), Highlander's Mail Greaves (20123)) | World drop [world_drop] | 17.4 ranged_attack_power points (1.03 DPS) | yes | Insignia Boots (4055, +0.00 DPS, sim-verified) [world_drop]; Vorrel's Boots (7751, +0.00 DPS) [quest]; Warsong Boots (16977, +0.00 DPS) [quest] |
| finger1 | Ironspine's Eye (7686) | Scarlet Monastery: Ironspine [dungeon] | 19.6 ranged_attack_power points (1.16 DPS) | yes | Ring of Precision (1491, -0.39 DPS) [dungeon]; Legionnaire's Band (19513, -0.39 DPS) [rep]; Signet of the Zhevra (285330, -0.39 DPS) [world] |
| finger2 | Monkey Ring (6748) | Willix the Importer [quest] | 15.3 ranged_attack_power points (0.90 DPS) | yes | Legionnaire's Band (19513, -0.13 DPS) [rep]; Signet of the Zhevra (285330, -0.13 DPS) [world]; Ring of Precision (1491, -0.22 DPS, sim-verified) [dungeon] |
| trinket1 | Darkspear Voodoo Seal (272059) | Creeg Bothunk [vendor] | sim-verified (91.8 DPS) | yes | Defiler's Talisman (21120, -2.41 DPS, sim-verified) [rep] |
| trinket2 | - | - |  |  |  |
| main_hand | Alliance Outrunner's Sword (285346) | Thora Feathermoon [world] | 16.9 ranged_attack_power points (1.00 DPS) | yes | Kam's Walking Stick (2280, +0.00 DPS) [dungeon]; Armor Piercer (6679, +0.00 DPS) [dungeon]; Bronze Dory (250603, +0.00 DPS) [crafted] |
| off_hand | Vendetta (776) (or Prison Shank (2941), Talon of Vultros (4454), Sentinel's Blade (212583), Scout's Blade (212587)) | Razorfen Kraul: Overlord Ramtusk [dungeon] | 13.1 ranged_attack_power points (0.77 DPS) | yes | Prison Shank (2941, +0.00 DPS, sim-verified) [dungeon]; Talon of Vultros (4454, +0.00 DPS) [world]; Sentinel's Blade (212583, +0.00 DPS) [vendor] |
| ranged | Booty Bay Bruiser's Buckshot (274748) | Gezzy Gunkgear [vendor] | 261.6 ranged_attack_power points (15.44 DPS) | yes | Glass Shooter (9456, -0.62 DPS) [dungeon]; Ironweaver (13137, -1.13 DPS) [world_drop]; Silver Star (3463, -1.92 DPS, sim-verified) [quest] |

**New at 30:** head: Brawler's Leather Helm; neck: Scout's Medallion; shoulder: Forest Tracker Epaulets; back: Tigerstrike Mantle; chest: Dusky Leather Armor; wrist: Jurassic Wristguards; waist: Defiler's Chain Girdle; legs: Petrolspill Leggings; finger1: Ironspine's Eye; finger2: Monkey Ring; trinket1: Darkspear Voodoo Seal; trinket2: Relentless Raider's Seal; main_hand: Alliance Outrunner's Sword; off_hand: Vendetta; ranged: Booty Bay Bruiser's Buckshot

No-known-source sample (15 of 336, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3738 Brewing Rod; 4196 Feathered Mantle; 4642 Star of Xil'yeh; 4763 Blackwood Recurve Bow; 5000 Coral Band; 5004 Mark of the Kirin Tor

### Band 40 (troll, 0000000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 101.3. Weights run: 1.2s. Verify run: 1.6s. 558 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.222 ± 0.063, crit=11.681 ± 0.473, hit=2.043 ± 0.081, melee_haste=not significant (3.437 ± 0.971)

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Raging Berserker's Helm (7719) | Scarlet Monastery: Herod [dungeon] | 163.5 ranged_attack_power points (9.59 DPS) | yes | Warden's Wizard Hat (14604, -0.00 DPS, sim-verified) [world_drop]; Nightscape Headband (8176, -8.03 DPS) [crafted]; Guard's Chain Helm (250499, -8.03 DPS) [crafted] |
| neck | Scout's Medallion (19536) | Warsong Outriders [rep] | 24.4 ranged_attack_power points (1.43 DPS) | yes | Scout's Medallion (19537, +0.00 DPS, sim-verified) [rep]; Ghostshard Talisman (7731, -0.61 DPS) [dungeon]; Scout's Medallion (20442, -0.65 DPS) [rep] |
| shoulder | Sunburn Spaulders (274751) | Rettrick [vendor] | 36.4 ranged_attack_power points (2.14 DPS) | yes | Forest Tracker Epaulets (2278, +0.00 DPS, sim-verified) [world_drop]; Nightscape Shoulders (8192, -0.70 DPS) [crafted]; Mantle of Thieves (2264, -0.83 DPS) [dungeon] |
| back | Dark Hooded Cape (5257) | Nimar the Slayer [world] | 22.2 ranged_attack_power points (1.30 DPS) | yes | Imperial Cloak (6432, +0.00 DPS, sim-verified) [world_drop]; Parachute Cloak (10518, -0.26 DPS) [crafted]; Tigerstrike Mantle (13108, -0.26 DPS) [world_drop] |
| chest | Quillward Harness (10583) | Razorfen Downs: Withered Warrior [dungeon] | 42.2 ranged_attack_power points (2.48 DPS) | yes | Wolffear Harness (13110, +0.00 DPS, sim-verified) [world_drop]; Nightscape Tunic (8175, -0.52 DPS) [crafted]; Tough Scorpid Breastplate (8203, -0.52 DPS) [crafted] |
| wrist | Branded Leather Bracers (19508) | Scarlet Monastery: High Inquisitor Fairbanks [dungeon] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Imperial Leather Bracers (4061, +0.00 DPS, sim-verified) [world_drop]; Dusky Bracers (7378, -0.13 DPS) [crafted]; Tough Scorpid Bracers (8205, -0.26 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 183.5 ranged_attack_power points (10.77 DPS) | yes | Dragonscale Gauntlets (8347, +0.00 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20153) | The Defilers [rep] | 171.5 ranged_attack_power points (10.06 DPS) | yes | Defiler's Leather Girdle (20192, -0.01 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20152, -8.66 DPS) [rep]; Defiler's Leather Girdle (20191, -8.66 DPS) [rep] |
| legs | Basilisk Hide Pants (1718) | World drop [world_drop] | 46.7 ranged_attack_power points (2.74 DPS) | yes | Triprunner Dungarees (9624, +0.00 DPS, sim-verified) [quest]; Petrolspill Leggings (9509, -0.91 DPS) [dungeon]; Troll's Bane Leggings (13114, -0.91 DPS) [world_drop] |
| feet | Swampwalker Boots (2276) | World drop [world_drop] | 28.9 ranged_attack_power points (1.69 DPS) | yes | Imperial Leather Boots (6431, +0.00 DPS, sim-verified) [world_drop]; Dusky Boots (7390, -0.26 DPS) [crafted]; Worn Running Boots (9398, -0.26 DPS) [dungeon] |
| finger1 | Ring of the Underwood (2951) | World drop [world_drop] | 22.2 ranged_attack_power points (1.30 DPS) | yes | Ironspine's Eye (7686, -0.13 DPS) [dungeon]; Assault Band (13095, -0.13 DPS) [world_drop]; Legionnaire's Band (19512, -0.26 DPS) [rep] |
| finger2 | Falcon's Hook (7552) (or Ironspine's Eye (7686)) | World drop [world_drop] | 20.0 ranged_attack_power points (1.17 DPS) | yes | Ironspine's Eye (7686, +0.00 DPS, sim-verified) [dungeon]; Assault Band (13095, -0.00 DPS) [world_drop]; Legionnaire's Band (19512, -0.13 DPS) [rep] |
| trinket1 | - | - |  |  |  |
| trinket2 | - | - |  |  |  |
| main_hand | Frost Tiger Blade (3854) | Blacksmithing [crafted] | sim-verified (31.1 DPS) | yes | Illusionary Rod (7713, +0.00 DPS) [dungeon]; Gut Ripper (2164, -0.01 DPS, sim-verified) [world_drop]; Steel Spear (250605, -7.31 DPS) [crafted] |
| off_hand | - | - |  |  |  |
| ranged | The Silencer (13138) | World drop [world_drop] | sim-verified (101.3 DPS) | yes | Shadowforge Bushmaster (9422, -0.46 DPS) [dungeon]; Swiftwind (13038, -0.79 DPS) [world_drop]; Bow of Searing Arrows (2825, -70.19 DPS, sim-verified) [world_drop] |

**New at 40:** head: Raging Berserker's Helm; neck: Scout's Medallion; shoulder: Sunburn Spaulders; back: Dark Hooded Cape; chest: Quillward Harness; wrist: Branded Leather Bracers; hands: Gloves of Holy Might; waist: Defiler's Chain Girdle; legs: Basilisk Hide Pants; feet: Swampwalker Boots; finger1: Ring of the Underwood; finger2: Falcon's Hook; trinket1: Rune of Perfection; trinket2: Rune of Duty; main_hand: Frost Tiger Blade; ranged: The Silencer

No-known-source sample (15 of 558, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 50 (troll, 5500000000000000-35305500115003000-000000000000000000)

Set DPS (verified): 138.3. Weights run: 1.3s. Verify run: 1.6s. 721 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.307 ± 0.077, crit=13.494 ± 0.534, hit=2.404 ± 0.097, melee_haste=13.786 ± 1.194

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Blood Guard's Chain Helmet (220821) | Lady Palanseer [vendor] | 225.8 ranged_attack_power points (13.20 DPS) | yes | Eye of Theradras (17715, -2.16 DPS) [dungeon]; Blood Guard's Mail Helmet (220820, -2.16 DPS) [vendor]; Raging Berserker's Helm (7719, -3.69 DPS, sim-verified) [dungeon] |
| neck | Skibi's Pendant (13089) | World drop [world_drop] | 30.0 ranged_attack_power points (1.75 DPS) | yes | Scout's Medallion (19536, -0.27 DPS) [rep]; Woven Ivy Necklace (19159, -0.54 DPS) [quest]; Scout's Medallion (19535, -0.78 DPS, sim-verified) [rep] |
| shoulder | Blood Guard's Chain Epaulets (220824) | Lady Palanseer [vendor] | 218.9 ranged_attack_power points (12.79 DPS) | yes | Blood Guard's Mail Epaulets (220823, -1.74 DPS, sim-verified) [vendor]; Sunburn Spaulders (274751, -10.61 DPS) [vendor]; Phytoskin Spaulders (17749, -10.64 DPS) [dungeon] |
| back | Dark Phantom Cape (13122) | World drop [world_drop] | sim-verified (135.0 DPS) | yes | Blackveil Cape (11626, -0.13 DPS) [dungeon]; Duskbat Drape (19982, -0.13 DPS) [quest]; Blisterbane Wrap (12552, -2.92 DPS, sim-verified) [dungeon] |
| chest | Stone Guard's Chain Armor (220827) | Lady Palanseer [vendor] | 223.5 ranged_attack_power points (13.06 DPS) | yes | Stone Guard's Mail Armor (220826, -2.21 DPS, sim-verified) [vendor]; Fungus Shroud Armor (17742, -9.69 DPS) [dungeon]; Blazewind Breastplate (11193, -9.96 DPS) [quest] |
| wrist | Deepfury Bracers (13120) | World drop [world_drop] | 34.6 ranged_attack_power points (2.02 DPS) | yes | Bracers of the Stone Princess (17714, -0.20 DPS, sim-verified) [dungeon]; Bloodlust Bracelets (14807, -0.54 DPS) [world_drop]; Wicked Leather Bracers (15084, -0.54 DPS) [crafted] |
| hands | Gloves of Holy Might (867) | World drop [world_drop] | 208.9 ranged_attack_power points (12.21 DPS) | yes | Dragonscale Gauntlets (8347, -0.36 DPS, sim-verified) [crafted]; Fletcher's Gloves (7348, -1.17 DPS) [crafted]; Shadowskin Gloves (18238, -1.17 DPS) [crafted] |
| waist | Defiler's Chain Girdle (20151) (or Defiler's Leather Girdle (20193)) | The Defilers [rep] | 208.9 ranged_attack_power points (12.21 DPS) | yes | Defiler's Leather Girdle (20193, +0.00 DPS, sim-verified) [rep]; Defiler's Chain Girdle (20153, -0.70 DPS) [rep]; Highlander's Mail Girdle (20118, -1.17 DPS) [vendor] |
| legs | Stone Guard's Chain Legplates (220833) | Lady Palanseer [vendor] | sim-verified (133.6 DPS) | yes | Stormshroud Pants (15057, -1.50 DPS, sim-verified) [crafted]; Stone Guard's Mail Legplates (220834, -1.89 DPS) [vendor]; Basilisk Hide Pants (1718, -10.10 DPS) [world_drop] |
| feet | Albino Crocscale Boots (17728) | Maraudon: Rotgrip [dungeon] | 46.1 ranged_attack_power points (2.70 DPS) | yes | Fleetfoot Greaves (11627, +0.00 DPS, sim-verified) [dungeon]; Elven Chain Boots (13125, -0.27 DPS) [world_drop]; Whisperwalk Boots (20255, -0.27 DPS) [quest] |
| finger1 | Blackstone Ring (17713) | Maraudon: Princess Theradras [dungeon] | 44.0 ranged_attack_power points (2.57 DPS) | yes | White Bone Band (11862, -1.17 DPS) [quest]; Ring of the Underwood (2951, -1.23 DPS) [world_drop]; Falcon's Hook (7552, -1.36 DPS) [world_drop] |
| finger2 | Masons Fraternity Ring (9533) | Divino-matic Rod [quest] | 32.3 ranged_attack_power points (1.89 DPS) | yes | Ring of the Underwood (2951, -0.54 DPS) [world_drop]; White Bone Band (11862, -0.55 DPS, sim-verified) [quest]; Falcon's Hook (7552, -0.67 DPS) [world_drop] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (132.1 DPS) | yes | Frozen Heart of the Mountain (249469, -4.63 DPS) [crafted]; Ankh of Life (1713, -5.88 DPS, sim-verified) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (132.1 DPS) | yes | Shortsword of Vengeance (754, +0.00 DPS, sim-verified) [world_drop]; Frost Tiger Blade (3854, +0.00 DPS) [crafted]; Illusionary Rod (7713, +0.00 DPS) [dungeon] |
| off_hand | Vanquisher's Sword (10823) | Bring the End [quest] | sim-verified (133.7 DPS) | yes | White Bone Shredder (11863, -0.34 DPS) [quest]; Thermotastic Egg Timer (9644, -0.88 DPS) [quest]; Inventor's Focal Sword (17719, -1.62 DPS, sim-verified) [dungeon] |
| ranged | Arcanite Blacksmith Hammer (285281) | Blacksmithing [crafted] | sim-verified (132.1 DPS) | yes | Hurricane (2824, -1.42 DPS) [world_drop]; Dark Iron Rifle (16004, -1.51 DPS, sim-verified) [crafted]; Precisely Calibrated Boomstick (2100, -1.70 DPS) [world_drop] |

**New at 50:** head: Blood Guard's Chain Helmet; neck: Skibi's Pendant; shoulder: Blood Guard's Chain Epaulets; back: Dark Phantom Cape; chest: Stone Guard's Chain Armor; wrist: Deepfury Bracers; waist: Defiler's Chain Girdle; legs: Stone Guard's Chain Legplates; feet: Albino Crocscale Boots; finger1: Blackstone Ring; finger2: Masons Fraternity Ring; trinket1: Rune of the Guard Captain; trinket2: Guardian Talisman; main_hand: Dawn's Edge; off_hand: Vanquisher's Sword; ranged: Arcanite Blacksmith Hammer

No-known-source sample (15 of 721, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

### Band 60 (troll, 5522000000000000-35305500115003000-510000000000000000)

Set DPS (verified): 204.1. Weights run: 1.2s. Verify run: 1.6s. 1500 eligible items had no known source.

Stat weights (normalized to ranged_attack_power = 1.0, error under 25% of the weight to publish - see report.go's isWeightSignificant): ranged_attack_power=1.000 ± 0.001, agility=2.339 ± 0.094, crit=20.135 ± 0.769, hit=3.434 ± 0.137, melee_haste=10.525 ± 1.657

| Slot | Item | Source | Score (ranged_attack_power points) | Verified | Alternatives |
|---|---|---|---|---|---|
| head | Dawnstalker Headpiece (239540) | Leonid Barthalomew the Revered [vendor] | 678.4 ranged_attack_power points (38.50 DPS) | yes | Champion's Chain Helm (23251, -4.11 DPS) [vendor]; Champion's Chain Helm (227067, -4.11 DPS) [vendor]; Champion's Chain Greathelm (227080, -5.16 DPS, sim-verified) [vendor] |
| neck | Medallion of the Dawn (22659) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (193.7 DPS) | yes | Blazefury Medallion (17111, -0.90 DPS, sim-verified) [world]; Beads of Ogre Might (22150, -14.05 DPS) [quest]; Amulet of the Darkmoon (19491, -14.84 DPS) [quest] |
| shoulder | Dawnstalker Pauldrons (239534) | Leonid Barthalomew the Revered [vendor] | sim-verified (197.7 DPS) | yes | Dawnstalker Spaulders (239542, -3.99 DPS, sim-verified) [vendor]; Darkspear Pauldrons (272105, -5.15 DPS) [vendor]; Darkspear Epaulets (272106, -5.15 DPS) [vendor] |
| back | Chromatic Cloak (18509) (or Fel Cape (279269)) | Leatherworking [crafted] | 281.9 ranged_attack_power points (16.00 DPS) | yes | Fel Cape (279269, +0.00 DPS, sim-verified) [crafted]; Earthweave Cloak (21187, -12.06 DPS) [quest]; Howler's Furs (272414, -12.46 DPS) [vendor] |
| chest | Dawnstalker Tunic (239543) | Leonid Barthalomew the Revered [vendor] | sim-verified (193.7 DPS) | yes | Dawnstalker Breastplate (239529, -13.16 DPS) [vendor]; Tunic of Undead Slaying (23089, -13.40 DPS, sim-verified) [world]; Legionnaire's Chain Armor (227083, -20.04 DPS) [vendor] |
| wrist | Dawnstalker Vambraces (239536) | Leonid Barthalomew the Revered [vendor] | sim-verified (193.7 DPS) | yes | Dawnstalker Wristguards (239544, -0.66 DPS) [vendor]; Wristwraps of Undead Slaying (23093, -7.75 DPS, sim-verified) [world]; Primal Batskin Bracers (19687, -16.70 DPS) [crafted] |
| hands | Dawnstalker Gauntlets (239531) | Leonid Barthalomew the Revered [vendor] | sim-verified (196.4 DPS) | yes | Dawnstalker Handguards (239539, -2.72 DPS, sim-verified) [vendor]; General's Chain Gloves (16571, -4.34 DPS) [vendor]; General's Chain Vices (231575, -4.34 DPS) [vendor] |
| waist | Dawnstalker Belt (239535) | Leonid Barthalomew the Revered [vendor] | 384.1 ranged_attack_power points (21.80 DPS) | yes | Dawnstalker Girdle (239538, -1.45 DPS, sim-verified) [vendor]; Defiler's Chain Girdle (20150, -3.87 DPS) [rep]; Defiler's Leather Girdle (20190, -3.87 DPS) [rep] |
| legs | Sentinel's Chain Leggings (237819) | Illiyana Moonblaze [vendor] | 680.0 ranged_attack_power points (38.59 DPS) | yes | Dawnstalker Legguards (239541, +0.00 DPS, sim-verified) [vendor]; Sentinel's Leather Pants (237818, -3.01 DPS) [vendor]; Legionnaire's Chain Legplates (227079, -4.13 DPS) [vendor] |
| feet | Dawnstalker Greaves (239530) | Leonid Barthalomew the Revered [vendor] | sim-verified (198.2 DPS) | yes | Dawnstalker Boots (239537, -4.47 DPS, sim-verified) [vendor]; General's Chain Sabatons (16569, -18.22 DPS) [vendor]; General's Chain Greaves (231570, -18.22 DPS) [vendor] |
| finger1 | Don Julio's Band (19325) | Frostwolf Clan [rep] | sim-verified (193.7 DPS) | yes | Band of the Penitent (13217, -2.86 DPS) [quest]; Ring of Entropy (18543, -2.86 DPS) [world]; Naglering (11669, -4.56 DPS, sim-verified) [dungeon] |
| finger2 | Band of Earthen Might (21182) | Veteran's Battlegear [quest] | sim-verified (193.7 DPS) | yes | Band of the Penitent (13217, -1.95 DPS) [quest]; Ring of Entropy (18543, -1.95 DPS) [world]; Naglering (11669, -3.30 DPS, sim-verified) [dungeon] |
| trinket1 | Rune of the Guard Captain (19120) | Job Opening: Guard Captain of Revantusk Village [quest] | sim-verified (193.7 DPS) | yes | Frozen Heart of the Mountain (249469, -4.38 DPS) [crafted]; Shard of the Fallen Star (21891, -8.04 DPS, sim-verified) [world_drop] |
| trinket2 | - | - |  |  |  |
| main_hand | Heartseeker (12783) | Blacksmithing [crafted] | sim-verified (193.7 DPS) | yes | High Warlord's Pig Sticker (234547, +0.00 DPS) [vendor]; High Warlord's Pig Poker (234548, +0.00 DPS) [vendor]; Glacial Blade (19099, -0.32 DPS, sim-verified) [rep] |
| off_hand | Dawn's Edge (12774) | Blacksmithing [crafted] | sim-verified (193.7 DPS) | yes | High Warlord's Left Claw (18848, +0.00 DPS) [vendor]; High Warlord's Left Claw (234558, +0.00 DPS) [vendor]; Eskhandar's Left Claw (18202, -1.02 DPS, sim-verified) [world] |
| ranged | The Purifier (22656) | Epic Armaments of Battle - Friend of the Dawn [quest] | sim-verified (193.7 DPS) | yes | High Warlord's Recurve (234559, -4.60 DPS) [vendor]; High Warlord's Crossbow (234560, -4.60 DPS) [vendor]; Dark Iron Rifle (16004, -9.12 DPS, sim-verified) [crafted] |

**New at 60:** head: Dawnstalker Headpiece; neck: Medallion of the Dawn; shoulder: Dawnstalker Pauldrons; back: Chromatic Cloak; chest: Dawnstalker Tunic; wrist: Dawnstalker Vambraces; hands: Dawnstalker Gauntlets; waist: Dawnstalker Belt; legs: Sentinel's Chain Leggings; feet: Dawnstalker Greaves; finger1: Don Julio's Band; finger2: Band of Earthen Might; trinket2: Darkmoon Card: Heroism; main_hand: Heartseeker; off_hand: Dawn's Edge; ranged: The Purifier

No-known-source sample (15 of 1500, see the JSON for more): 913 Huge Ogre Sword; 1189 Overseer's Ring; 1216 Frost Bracers; 1832 Lucky Trousers; 2016 Dusty Chain Armor; 2273 Guerrilla Armor; 2543 Militia Pants; 2664 Spinner Fang; 2944 Cursed Eye of Paleth; 2952 Fine Light Hide Jerkin; 3038 Archer's Longbow; 3222 Wicked Dagger; 3579 Ornate Copper Shoulders; 3738 Brewing Rod; 4081 Blackforge Leggings

